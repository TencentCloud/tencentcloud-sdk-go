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

package v20180813

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type AddProjectRequestParams struct {
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 项目描述
	Info *string `json:"Info,omitempty" name:"Info"`
}

type AddProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 项目描述
	Info *string `json:"Info,omitempty" name:"Info"`
}

func (r *AddProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectName")
	delete(f, "Info")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddProjectResponseParams struct {
	// 项目Id
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 是否为新项目
	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddProjectResponse struct {
	*tchttp.BaseResponse
	Response *AddProjectResponseParams `json:"Response"`
}

func (r *AddProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddResourceTagRequestParams struct {
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// [ 资源六段式描述 ](https://cloud.tencent.com/document/product/598/10606)
	Resource *string `json:"Resource,omitempty" name:"Resource"`
}

type AddResourceTagRequest struct {
	*tchttp.BaseRequest
	
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// [ 资源六段式描述 ](https://cloud.tencent.com/document/product/598/10606)
	Resource *string `json:"Resource,omitempty" name:"Resource"`
}

func (r *AddResourceTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddResourceTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKey")
	delete(f, "TagValue")
	delete(f, "Resource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddResourceTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddResourceTagResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddResourceTagResponse struct {
	*tchttp.BaseResponse
	Response *AddResourceTagResponseParams `json:"Response"`
}

func (r *AddResourceTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddResourceTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachResourcesTagRequestParams struct {
	// 资源所属业务名称（资源六段式中的第三段）
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 资源ID数组，资源个数最多为50
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// 资源所在地域，不区分地域的资源不需要传入该字段，区分地域的资源必填
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 资源前缀（资源六段式中最后一段"/"前面的部分），cos存储桶不需要传入该字段，其他云资源必填
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`
}

type AttachResourcesTagRequest struct {
	*tchttp.BaseRequest
	
	// 资源所属业务名称（资源六段式中的第三段）
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 资源ID数组，资源个数最多为50
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// 资源所在地域，不区分地域的资源不需要传入该字段，区分地域的资源必填
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 资源前缀（资源六段式中最后一段"/"前面的部分），cos存储桶不需要传入该字段，其他云资源必填
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`
}

func (r *AttachResourcesTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachResourcesTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceType")
	delete(f, "ResourceIds")
	delete(f, "TagKey")
	delete(f, "TagValue")
	delete(f, "ResourceRegion")
	delete(f, "ResourcePrefix")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachResourcesTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachResourcesTagResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AttachResourcesTagResponse struct {
	*tchttp.BaseResponse
	Response *AttachResourcesTagResponseParams `json:"Response"`
}

func (r *AttachResourcesTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachResourcesTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTagRequestParams struct {
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type CreateTagRequest struct {
	*tchttp.BaseRequest
	
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

func (r *CreateTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKey")
	delete(f, "TagValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTagResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTagResponse struct {
	*tchttp.BaseResponse
	Response *CreateTagResponseParams `json:"Response"`
}

func (r *CreateTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTagsRequestParams struct {
	// 标签列表。
	// N取值范围：0~9
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CreateTagsRequest struct {
	*tchttp.BaseRequest
	
	// 标签列表。
	// N取值范围：0~9
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTagsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTagsResponse struct {
	*tchttp.BaseResponse
	Response *CreateTagsResponseParams `json:"Response"`
}

func (r *CreateTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteResourceTagRequestParams struct {
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// [ 资源六段式描述 ](https://cloud.tencent.com/document/product/598/10606)
	Resource *string `json:"Resource,omitempty" name:"Resource"`
}

type DeleteResourceTagRequest struct {
	*tchttp.BaseRequest
	
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// [ 资源六段式描述 ](https://cloud.tencent.com/document/product/598/10606)
	Resource *string `json:"Resource,omitempty" name:"Resource"`
}

func (r *DeleteResourceTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourceTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKey")
	delete(f, "Resource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteResourceTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteResourceTagResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteResourceTagResponse struct {
	*tchttp.BaseResponse
	Response *DeleteResourceTagResponseParams `json:"Response"`
}

func (r *DeleteResourceTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourceTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTagRequestParams struct {
	// 需要删除的标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 需要删除的标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type DeleteTagRequest struct {
	*tchttp.BaseRequest
	
	// 需要删除的标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 需要删除的标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

func (r *DeleteTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKey")
	delete(f, "TagValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTagResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTagResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTagResponseParams `json:"Response"`
}

func (r *DeleteTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTagsRequestParams struct {
	// 标签列表。
	// N取值范围：0~9
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type DeleteTagsRequest struct {
	*tchttp.BaseRequest
	
	// 标签列表。
	// N取值范围：0~9
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *DeleteTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTagsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTagsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTagsResponseParams `json:"Response"`
}

func (r *DeleteTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectsRequestParams struct {
	// 传1拉取所有项目（包括隐藏项目），传0拉取显示项目
	AllList *uint64 `json:"AllList,omitempty" name:"AllList"`

	// 分页条数，固定值1000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeProjectsRequest struct {
	*tchttp.BaseRequest
	
	// 传1拉取所有项目（包括隐藏项目），传0拉取显示项目
	AllList *uint64 `json:"AllList,omitempty" name:"AllList"`

	// 分页条数，固定值1000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeProjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AllList")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectsResponseParams struct {
	// 数据总条数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 项目列表
	Projects []*Project `json:"Projects,omitempty" name:"Projects"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProjectsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectsResponseParams `json:"Response"`
}

func (r *DescribeProjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceTagsByResourceIdsRequestParams struct {
	// 业务类型
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 资源前缀
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`

	// 资源ID数组，大小不超过50
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 资源所在地域
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 数据偏移量，默认为 0, 必须为Limit参数的整数倍
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小，默认为 15
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeResourceTagsByResourceIdsRequest struct {
	*tchttp.BaseRequest
	
	// 业务类型
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 资源前缀
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`

	// 资源ID数组，大小不超过50
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 资源所在地域
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 数据偏移量，默认为 0, 必须为Limit参数的整数倍
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小，默认为 15
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeResourceTagsByResourceIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceTagsByResourceIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceType")
	delete(f, "ResourcePrefix")
	delete(f, "ResourceIds")
	delete(f, "ResourceRegion")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceTagsByResourceIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceTagsByResourceIdsResponseParams struct {
	// 结果总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 数据位移偏量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 标签列表
	Tags []*TagResource `json:"Tags,omitempty" name:"Tags"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResourceTagsByResourceIdsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceTagsByResourceIdsResponseParams `json:"Response"`
}

func (r *DescribeResourceTagsByResourceIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceTagsByResourceIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceTagsByResourceIdsSeqRequestParams struct {
	// 业务类型
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 资源前缀
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`

	// 资源唯一标记
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 资源所在地域
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 数据偏移量，默认为 0, 必须为Limit参数的整数倍
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小，默认为 15
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeResourceTagsByResourceIdsSeqRequest struct {
	*tchttp.BaseRequest
	
	// 业务类型
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 资源前缀
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`

	// 资源唯一标记
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 资源所在地域
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 数据偏移量，默认为 0, 必须为Limit参数的整数倍
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小，默认为 15
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeResourceTagsByResourceIdsSeqRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceTagsByResourceIdsSeqRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceType")
	delete(f, "ResourcePrefix")
	delete(f, "ResourceIds")
	delete(f, "ResourceRegion")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceTagsByResourceIdsSeqRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceTagsByResourceIdsSeqResponseParams struct {
	// 结果总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 数据位移偏量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 标签列表
	Tags []*TagResource `json:"Tags,omitempty" name:"Tags"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResourceTagsByResourceIdsSeqResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceTagsByResourceIdsSeqResponseParams `json:"Response"`
}

func (r *DescribeResourceTagsByResourceIdsSeqResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceTagsByResourceIdsSeqResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceTagsByTagKeysRequestParams struct {
	// 业务类型
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 资源前缀
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`

	// 资源地域
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 资源唯一标识
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 资源标签键
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// 每页大小，默认为 400
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 数据偏移量，默认为 0, 必须为Limit参数的整数倍
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeResourceTagsByTagKeysRequest struct {
	*tchttp.BaseRequest
	
	// 业务类型
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 资源前缀
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`

	// 资源地域
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 资源唯一标识
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 资源标签键
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// 每页大小，默认为 400
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 数据偏移量，默认为 0, 必须为Limit参数的整数倍
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeResourceTagsByTagKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceTagsByTagKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceType")
	delete(f, "ResourcePrefix")
	delete(f, "ResourceRegion")
	delete(f, "ResourceIds")
	delete(f, "TagKeys")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceTagsByTagKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceTagsByTagKeysResponseParams struct {
	// 结果总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 数据位移偏量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 资源标签
	Rows []*ResourceIdTag `json:"Rows,omitempty" name:"Rows"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResourceTagsByTagKeysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceTagsByTagKeysResponseParams `json:"Response"`
}

func (r *DescribeResourceTagsByTagKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceTagsByTagKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceTagsRequestParams struct {
	// 创建者uin
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// 资源所在地域
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 业务类型
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 资源前缀
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`

	// 资源唯一标识。只输入ResourceId进行查询可能会查询较慢，或者无法匹配到结果，建议在输入ResourceId的同时也输入ServiceType、ResourcePrefix和ResourceRegion（不区分地域的资源可忽略该参数）
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 数据偏移量，默认为 0, 必须为Limit参数的整数倍
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小，默认为 15
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 是否是cos的资源（0或者1），输入的ResourceId为cos资源时必填
	CosResourceId *uint64 `json:"CosResourceId,omitempty" name:"CosResourceId"`
}

type DescribeResourceTagsRequest struct {
	*tchttp.BaseRequest
	
	// 创建者uin
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// 资源所在地域
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 业务类型
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 资源前缀
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`

	// 资源唯一标识。只输入ResourceId进行查询可能会查询较慢，或者无法匹配到结果，建议在输入ResourceId的同时也输入ServiceType、ResourcePrefix和ResourceRegion（不区分地域的资源可忽略该参数）
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 数据偏移量，默认为 0, 必须为Limit参数的整数倍
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小，默认为 15
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 是否是cos的资源（0或者1），输入的ResourceId为cos资源时必填
	CosResourceId *uint64 `json:"CosResourceId,omitempty" name:"CosResourceId"`
}

func (r *DescribeResourceTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CreateUin")
	delete(f, "ResourceRegion")
	delete(f, "ServiceType")
	delete(f, "ResourcePrefix")
	delete(f, "ResourceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "CosResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceTagsResponseParams struct {
	// 结果总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 数据位移偏量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 资源标签
	Rows []*TagResource `json:"Rows,omitempty" name:"Rows"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResourceTagsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceTagsResponseParams `json:"Response"`
}

func (r *DescribeResourceTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcesByTagsRequestParams struct {
	// 标签过滤数组
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`

	// 创建标签者uin
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// 数据偏移量，默认为 0, 必须为Limit参数的整数倍
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小，默认为 15
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 资源前缀
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`

	// 资源唯一标记
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源所在地域
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 业务类型
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

type DescribeResourcesByTagsRequest struct {
	*tchttp.BaseRequest
	
	// 标签过滤数组
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`

	// 创建标签者uin
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// 数据偏移量，默认为 0, 必须为Limit参数的整数倍
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小，默认为 15
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 资源前缀
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`

	// 资源唯一标记
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源所在地域
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 业务类型
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DescribeResourcesByTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcesByTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagFilters")
	delete(f, "CreateUin")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ResourcePrefix")
	delete(f, "ResourceId")
	delete(f, "ResourceRegion")
	delete(f, "ServiceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourcesByTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcesByTagsResponseParams struct {
	// 结果总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 数据位移偏量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 资源标签
	Rows []*ResourceTag `json:"Rows,omitempty" name:"Rows"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResourcesByTagsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourcesByTagsResponseParams `json:"Response"`
}

func (r *DescribeResourcesByTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcesByTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcesByTagsUnionRequestParams struct {
	// 标签过滤数组
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`

	// 创建标签者uin
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// 数据偏移量，默认为 0, 必须为Limit参数的整数倍
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小，默认为 15
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 资源前缀
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`

	// 资源唯一标记
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源所在地域
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 业务类型
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

type DescribeResourcesByTagsUnionRequest struct {
	*tchttp.BaseRequest
	
	// 标签过滤数组
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`

	// 创建标签者uin
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// 数据偏移量，默认为 0, 必须为Limit参数的整数倍
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小，默认为 15
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 资源前缀
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`

	// 资源唯一标记
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源所在地域
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 业务类型
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DescribeResourcesByTagsUnionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcesByTagsUnionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagFilters")
	delete(f, "CreateUin")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ResourcePrefix")
	delete(f, "ResourceId")
	delete(f, "ResourceRegion")
	delete(f, "ServiceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourcesByTagsUnionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcesByTagsUnionResponseParams struct {
	// 结果总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 数据位移偏量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 资源标签
	Rows []*ResourceTag `json:"Rows,omitempty" name:"Rows"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResourcesByTagsUnionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourcesByTagsUnionResponseParams `json:"Response"`
}

func (r *DescribeResourcesByTagsUnionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcesByTagsUnionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagKeysRequestParams struct {
	// 创建者用户 Uin，不传或为空只将 Uin 作为条件查询
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// 数据偏移量，默认为 0, 必须为Limit参数的整数倍
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小，默认为 15
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 是否展现项目
	ShowProject *uint64 `json:"ShowProject,omitempty" name:"ShowProject"`
}

type DescribeTagKeysRequest struct {
	*tchttp.BaseRequest
	
	// 创建者用户 Uin，不传或为空只将 Uin 作为条件查询
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// 数据偏移量，默认为 0, 必须为Limit参数的整数倍
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小，默认为 15
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 是否展现项目
	ShowProject *uint64 `json:"ShowProject,omitempty" name:"ShowProject"`
}

func (r *DescribeTagKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CreateUin")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ShowProject")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagKeysResponseParams struct {
	// 结果总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 数据位移偏量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 标签列表
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTagKeysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTagKeysResponseParams `json:"Response"`
}

func (r *DescribeTagKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagValuesRequestParams struct {
	// 标签键列表
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// 创建者用户 Uin，不传或为空只将 Uin 作为条件查询
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// 数据偏移量，默认为 0, 必须为Limit参数的整数倍
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小，默认为 15
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeTagValuesRequest struct {
	*tchttp.BaseRequest
	
	// 标签键列表
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// 创建者用户 Uin，不传或为空只将 Uin 作为条件查询
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// 数据偏移量，默认为 0, 必须为Limit参数的整数倍
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小，默认为 15
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTagValuesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagValuesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKeys")
	delete(f, "CreateUin")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagValuesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagValuesResponseParams struct {
	// 结果总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 数据位移偏量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTagValuesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTagValuesResponseParams `json:"Response"`
}

func (r *DescribeTagValuesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagValuesSeqRequestParams struct {
	// 标签键列表
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// 创建者用户 Uin，不传或为空只将 Uin 作为条件查询
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// 数据偏移量，默认为 0, 必须为Limit参数的整数倍
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小，默认为 15
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeTagValuesSeqRequest struct {
	*tchttp.BaseRequest
	
	// 标签键列表
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// 创建者用户 Uin，不传或为空只将 Uin 作为条件查询
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// 数据偏移量，默认为 0, 必须为Limit参数的整数倍
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小，默认为 15
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTagValuesSeqRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagValuesSeqRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKeys")
	delete(f, "CreateUin")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagValuesSeqRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagValuesSeqResponseParams struct {
	// 结果总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 数据位移偏量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTagValuesSeqResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTagValuesSeqResponseParams `json:"Response"`
}

func (r *DescribeTagValuesSeqResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagValuesSeqResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagsRequestParams struct {
	// 标签键,与标签值同时存在或同时不存在，不存在时表示查询该用户所有标签
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值,与标签键同时存在或同时不存在，不存在时表示查询该用户所有标签
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// 数据偏移量，默认为 0, 必须为Limit参数的整数倍
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小，默认为 15
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 创建者用户 Uin，不传或为空只将 Uin 作为条件查询
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// 标签键数组,与标签值同时存在或同时不存在，不存在时表示查询该用户所有标签,当与TagKey同时传递时只取本值
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// 是否展现项目标签
	ShowProject *uint64 `json:"ShowProject,omitempty" name:"ShowProject"`
}

type DescribeTagsRequest struct {
	*tchttp.BaseRequest
	
	// 标签键,与标签值同时存在或同时不存在，不存在时表示查询该用户所有标签
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值,与标签键同时存在或同时不存在，不存在时表示查询该用户所有标签
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// 数据偏移量，默认为 0, 必须为Limit参数的整数倍
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小，默认为 15
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 创建者用户 Uin，不传或为空只将 Uin 作为条件查询
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// 标签键数组,与标签值同时存在或同时不存在，不存在时表示查询该用户所有标签,当与TagKey同时传递时只取本值
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// 是否展现项目标签
	ShowProject *uint64 `json:"ShowProject,omitempty" name:"ShowProject"`
}

func (r *DescribeTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKey")
	delete(f, "TagValue")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "CreateUin")
	delete(f, "TagKeys")
	delete(f, "ShowProject")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagsResponseParams struct {
	// 结果总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 数据位移偏量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 标签列表
	Tags []*TagWithDelete `json:"Tags,omitempty" name:"Tags"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTagsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTagsResponseParams `json:"Response"`
}

func (r *DescribeTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagsSeqRequestParams struct {
	// 标签键,与标签值同时存在或同时不存在，不存在时表示查询该用户所有标签
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值,与标签键同时存在或同时不存在，不存在时表示查询该用户所有标签
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// 数据偏移量，默认为 0, 必须为Limit参数的整数倍
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小，默认为 15
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 创建者用户 Uin，不传或为空只将 Uin 作为条件查询
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// 标签键数组,与标签值同时存在或同时不存在，不存在时表示查询该用户所有标签,当与TagKey同时传递时只取本值
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// 是否展现项目标签
	ShowProject *uint64 `json:"ShowProject,omitempty" name:"ShowProject"`
}

type DescribeTagsSeqRequest struct {
	*tchttp.BaseRequest
	
	// 标签键,与标签值同时存在或同时不存在，不存在时表示查询该用户所有标签
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值,与标签键同时存在或同时不存在，不存在时表示查询该用户所有标签
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// 数据偏移量，默认为 0, 必须为Limit参数的整数倍
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小，默认为 15
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 创建者用户 Uin，不传或为空只将 Uin 作为条件查询
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// 标签键数组,与标签值同时存在或同时不存在，不存在时表示查询该用户所有标签,当与TagKey同时传递时只取本值
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// 是否展现项目标签
	ShowProject *uint64 `json:"ShowProject,omitempty" name:"ShowProject"`
}

func (r *DescribeTagsSeqRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagsSeqRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKey")
	delete(f, "TagValue")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "CreateUin")
	delete(f, "TagKeys")
	delete(f, "ShowProject")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagsSeqRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagsSeqResponseParams struct {
	// 结果总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 数据位移偏量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 标签列表
	Tags []*TagWithDelete `json:"Tags,omitempty" name:"Tags"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTagsSeqResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTagsSeqResponseParams `json:"Response"`
}

func (r *DescribeTagsSeqResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagsSeqResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachResourcesTagRequestParams struct {
	// 资源所属业务名称（资源六段式中的第三段）
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 资源ID数组，资源个数最多为50
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 需要解绑的标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 资源所在地域，不区分地域的资源不需要传入该字段，区分地域的资源必填
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 资源前缀（资源六段式中最后一段"/"前面的部分），cos存储桶不需要传入该字段，其他云资源必填
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`
}

type DetachResourcesTagRequest struct {
	*tchttp.BaseRequest
	
	// 资源所属业务名称（资源六段式中的第三段）
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 资源ID数组，资源个数最多为50
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 需要解绑的标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 资源所在地域，不区分地域的资源不需要传入该字段，区分地域的资源必填
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 资源前缀（资源六段式中最后一段"/"前面的部分），cos存储桶不需要传入该字段，其他云资源必填
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`
}

func (r *DetachResourcesTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachResourcesTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceType")
	delete(f, "ResourceIds")
	delete(f, "TagKey")
	delete(f, "ResourceRegion")
	delete(f, "ResourcePrefix")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachResourcesTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachResourcesTagResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetachResourcesTagResponse struct {
	*tchttp.BaseResponse
	Response *DetachResourcesTagResponseParams `json:"Response"`
}

func (r *DetachResourcesTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachResourcesTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FailedResource struct {
	// 失败的资源六段式
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// 错误码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 错误信息
	Message *string `json:"Message,omitempty" name:"Message"`
}

// Predefined struct for user
type GetResourcesRequestParams struct {
	// 资源六段式列表。腾讯云使用资源六段式描述一个资源。
	// 例如：ResourceList.1 = qcs::${ServiceType}:${Region}:${Account}:${ResourcePreifx}/${ResourceId}。
	// 如果传入了此参数会返回所有匹配的资源列表，指定的MaxResults会失效。
	// N取值范围：0~9
	ResourceList []*string `json:"ResourceList,omitempty" name:"ResourceList"`

	// 标签键和标签值。
	// 指定多个标签，会查询同时绑定了该多个标签的资源。
	// N取值范围：0~5。
	// 每个TagFilters中的TagValue最多支持10个
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`

	// 从上一页的响应中获取的下一页的Token值。
	// 如果是第一次请求，设置为空。
	PaginationToken *string `json:"PaginationToken,omitempty" name:"PaginationToken"`

	// 每一页返回的数据最大条数，最大200。
	// 缺省值：50。
	MaxResults *uint64 `json:"MaxResults,omitempty" name:"MaxResults"`
}

type GetResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 资源六段式列表。腾讯云使用资源六段式描述一个资源。
	// 例如：ResourceList.1 = qcs::${ServiceType}:${Region}:${Account}:${ResourcePreifx}/${ResourceId}。
	// 如果传入了此参数会返回所有匹配的资源列表，指定的MaxResults会失效。
	// N取值范围：0~9
	ResourceList []*string `json:"ResourceList,omitempty" name:"ResourceList"`

	// 标签键和标签值。
	// 指定多个标签，会查询同时绑定了该多个标签的资源。
	// N取值范围：0~5。
	// 每个TagFilters中的TagValue最多支持10个
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`

	// 从上一页的响应中获取的下一页的Token值。
	// 如果是第一次请求，设置为空。
	PaginationToken *string `json:"PaginationToken,omitempty" name:"PaginationToken"`

	// 每一页返回的数据最大条数，最大200。
	// 缺省值：50。
	MaxResults *uint64 `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *GetResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceList")
	delete(f, "TagFilters")
	delete(f, "PaginationToken")
	delete(f, "MaxResults")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetResourcesResponseParams struct {
	// 获取的下一页的Token值
	PaginationToken *string `json:"PaginationToken,omitempty" name:"PaginationToken"`

	// 资源及关联的标签(键和值)列表
	ResourceTagMappingList []*ResourceTagMapping `json:"ResourceTagMappingList,omitempty" name:"ResourceTagMappingList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetResourcesResponse struct {
	*tchttp.BaseResponse
	Response *GetResourcesResponseParams `json:"Response"`
}

func (r *GetResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTagKeysRequestParams struct {
	// 从上一页的响应中获取的下一页的Token值。
	// 如果是第一次请求，设置为空。
	PaginationToken *string `json:"PaginationToken,omitempty" name:"PaginationToken"`

	// 每一页返回的数据最大条数，最大1000。
	// 缺省值：50。
	MaxResults *uint64 `json:"MaxResults,omitempty" name:"MaxResults"`
}

type GetTagKeysRequest struct {
	*tchttp.BaseRequest
	
	// 从上一页的响应中获取的下一页的Token值。
	// 如果是第一次请求，设置为空。
	PaginationToken *string `json:"PaginationToken,omitempty" name:"PaginationToken"`

	// 每一页返回的数据最大条数，最大1000。
	// 缺省值：50。
	MaxResults *uint64 `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *GetTagKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTagKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PaginationToken")
	delete(f, "MaxResults")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTagKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTagKeysResponseParams struct {
	// 获取的下一页的Token值
	PaginationToken *string `json:"PaginationToken,omitempty" name:"PaginationToken"`

	// 标签键信息。
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetTagKeysResponse struct {
	*tchttp.BaseResponse
	Response *GetTagKeysResponseParams `json:"Response"`
}

func (r *GetTagKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTagKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTagValuesRequestParams struct {
	// 标签键。
	// 返回所有标签键列表对应的标签值。
	// 最大长度：20
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// 从上一页的响应中获取的下一页的Token值。
	// 如果是第一次请求，设置为空。
	PaginationToken *string `json:"PaginationToken,omitempty" name:"PaginationToken"`

	// 每一页返回的数据最大条数，最大1000。
	// 缺省值：50。
	MaxResults *uint64 `json:"MaxResults,omitempty" name:"MaxResults"`
}

type GetTagValuesRequest struct {
	*tchttp.BaseRequest
	
	// 标签键。
	// 返回所有标签键列表对应的标签值。
	// 最大长度：20
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// 从上一页的响应中获取的下一页的Token值。
	// 如果是第一次请求，设置为空。
	PaginationToken *string `json:"PaginationToken,omitempty" name:"PaginationToken"`

	// 每一页返回的数据最大条数，最大1000。
	// 缺省值：50。
	MaxResults *uint64 `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *GetTagValuesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTagValuesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKeys")
	delete(f, "PaginationToken")
	delete(f, "MaxResults")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTagValuesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTagValuesResponseParams struct {
	// 获取的下一页的Token值
	PaginationToken *string `json:"PaginationToken,omitempty" name:"PaginationToken"`

	// 标签列表。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetTagValuesResponse struct {
	*tchttp.BaseResponse
	Response *GetTagValuesResponseParams `json:"Response"`
}

func (r *GetTagValuesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTagValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTagsRequestParams struct {
	// 从上一页的响应中获取的下一页的Token值。
	// 如果是第一次请求，设置为空。
	PaginationToken *string `json:"PaginationToken,omitempty" name:"PaginationToken"`

	// 每一页返回的数据最大条数，最大1000。
	// 缺省值：50。
	MaxResults *uint64 `json:"MaxResults,omitempty" name:"MaxResults"`

	// 标签键。
	// 返回所有标签键列表对应的标签。
	// 最大长度：20
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`
}

type GetTagsRequest struct {
	*tchttp.BaseRequest
	
	// 从上一页的响应中获取的下一页的Token值。
	// 如果是第一次请求，设置为空。
	PaginationToken *string `json:"PaginationToken,omitempty" name:"PaginationToken"`

	// 每一页返回的数据最大条数，最大1000。
	// 缺省值：50。
	MaxResults *uint64 `json:"MaxResults,omitempty" name:"MaxResults"`

	// 标签键。
	// 返回所有标签键列表对应的标签。
	// 最大长度：20
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`
}

func (r *GetTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PaginationToken")
	delete(f, "MaxResults")
	delete(f, "TagKeys")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTagsResponseParams struct {
	// 获取的下一页的Token值
	PaginationToken *string `json:"PaginationToken,omitempty" name:"PaginationToken"`

	// 标签列表。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetTagsResponse struct {
	*tchttp.BaseResponse
	Response *GetTagsResponseParams `json:"Response"`
}

func (r *GetTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceTagsRequestParams struct {
	// [ 资源六段式描述 ](https://cloud.tencent.com/document/product/598/10606)
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// 需要增加或修改的标签集合。如果Resource描述的资源未关联输入的标签键，则增加关联；若已关联，则将该资源关联的键对应的标签值修改为输入值。本接口中ReplaceTags和DeleteTags二者必须存在其一，且二者不能包含相同的标签键。可以不传该参数，但不能是空数组。
	ReplaceTags []*Tag `json:"ReplaceTags,omitempty" name:"ReplaceTags"`

	// 需要解关联的标签集合。本接口中ReplaceTags和DeleteTags二者必须存在其一，且二者不能包含相同的标签键。可以不传该参数，但不能是空数组。
	DeleteTags []*TagKeyObject `json:"DeleteTags,omitempty" name:"DeleteTags"`
}

type ModifyResourceTagsRequest struct {
	*tchttp.BaseRequest
	
	// [ 资源六段式描述 ](https://cloud.tencent.com/document/product/598/10606)
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// 需要增加或修改的标签集合。如果Resource描述的资源未关联输入的标签键，则增加关联；若已关联，则将该资源关联的键对应的标签值修改为输入值。本接口中ReplaceTags和DeleteTags二者必须存在其一，且二者不能包含相同的标签键。可以不传该参数，但不能是空数组。
	ReplaceTags []*Tag `json:"ReplaceTags,omitempty" name:"ReplaceTags"`

	// 需要解关联的标签集合。本接口中ReplaceTags和DeleteTags二者必须存在其一，且二者不能包含相同的标签键。可以不传该参数，但不能是空数组。
	DeleteTags []*TagKeyObject `json:"DeleteTags,omitempty" name:"DeleteTags"`
}

func (r *ModifyResourceTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Resource")
	delete(f, "ReplaceTags")
	delete(f, "DeleteTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourceTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceTagsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyResourceTagsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourceTagsResponseParams `json:"Response"`
}

func (r *ModifyResourceTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourcesTagValueRequestParams struct {
	// 资源所属业务名称（资源六段式中的第三段）
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 资源ID数组，资源个数最多为50
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// 资源所在地域，不区分地域的资源不需要传入该字段，区分地域的资源必填
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 资源前缀（资源六段式中最后一段"/"前面的部分），cos存储桶不需要传入该字段，其他云资源必填
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`
}

type ModifyResourcesTagValueRequest struct {
	*tchttp.BaseRequest
	
	// 资源所属业务名称（资源六段式中的第三段）
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 资源ID数组，资源个数最多为50
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// 资源所在地域，不区分地域的资源不需要传入该字段，区分地域的资源必填
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 资源前缀（资源六段式中最后一段"/"前面的部分），cos存储桶不需要传入该字段，其他云资源必填
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`
}

func (r *ModifyResourcesTagValueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourcesTagValueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceType")
	delete(f, "ResourceIds")
	delete(f, "TagKey")
	delete(f, "TagValue")
	delete(f, "ResourceRegion")
	delete(f, "ResourcePrefix")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourcesTagValueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourcesTagValueResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyResourcesTagValueResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourcesTagValueResponseParams `json:"Response"`
}

func (r *ModifyResourcesTagValueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourcesTagValueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Project struct {
	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 创建人uin
	CreatorUin *uint64 `json:"CreatorUin,omitempty" name:"CreatorUin"`

	// 项目描述
	ProjectInfo *string `json:"ProjectInfo,omitempty" name:"ProjectInfo"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type ResourceIdTag struct {
	// 资源唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 标签键值对
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKeyValues []*Tag `json:"TagKeyValues,omitempty" name:"TagKeyValues"`
}

type ResourceTag struct {
	// 资源所在地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 业务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 资源前缀
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`

	// 资源唯一标记
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type ResourceTagMapping struct {
	// 资源六段式。腾讯云使用资源六段式描述一个资源。
	// 例如：ResourceList.1 = qcs::${ServiceType}:${Region}:${Account}:${ResourcePreifx}/${ResourceId}。
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// 资源关联的标签列表
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type Tag struct {
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TagFilter struct {
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值数组 多个值的话是或的关系
	TagValue []*string `json:"TagValue,omitempty" name:"TagValue"`
}

type TagKeyObject struct {
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
}

type TagResource struct {
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 标签键MD5值
	TagKeyMd5 *string `json:"TagKeyMd5,omitempty" name:"TagKeyMd5"`

	// 标签值MD5值
	TagValueMd5 *string `json:"TagValueMd5,omitempty" name:"TagValueMd5"`

	// 资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

// Predefined struct for user
type TagResourcesRequestParams struct {
	// 资源六段式列表。腾讯云使用资源六段式描述一个资源。可参考[访问管理](https://cloud.tencent.com/document/product/598/67350)-概览-接口列表-资源六段式信息
	// 例如：ResourceList.1 = qcs::${ServiceType}:${Region}:uin/${Account}:${ResourcePrefix}/${ResourceId}。
	// N取值范围：0~9
	ResourceList []*string `json:"ResourceList,omitempty" name:"ResourceList"`

	// 标签键和标签值。
	// 如果指定多个标签，则会为指定资源同时创建并绑定该多个标签。
	// 同一个资源上的同一个标签键只能对应一个标签值。如果您尝试添加已有标签键，则对应的标签值会更新为新值。
	// 如果标签不存在会为您自动创建标签。
	// N取值范围：0~9
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type TagResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 资源六段式列表。腾讯云使用资源六段式描述一个资源。可参考[访问管理](https://cloud.tencent.com/document/product/598/67350)-概览-接口列表-资源六段式信息
	// 例如：ResourceList.1 = qcs::${ServiceType}:${Region}:uin/${Account}:${ResourcePrefix}/${ResourceId}。
	// N取值范围：0~9
	ResourceList []*string `json:"ResourceList,omitempty" name:"ResourceList"`

	// 标签键和标签值。
	// 如果指定多个标签，则会为指定资源同时创建并绑定该多个标签。
	// 同一个资源上的同一个标签键只能对应一个标签值。如果您尝试添加已有标签键，则对应的标签值会更新为新值。
	// 如果标签不存在会为您自动创建标签。
	// N取值范围：0~9
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *TagResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TagResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceList")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TagResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TagResourcesResponseParams struct {
	// 失败资源信息。
	// 创建并绑定标签成功时，返回的FailedResources为空。
	// 创建并绑定标签失败或部分失败时，返回的FailedResources会显示失败资源的详细信息。
	FailedResources []*FailedResource `json:"FailedResources,omitempty" name:"FailedResources"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TagResourcesResponse struct {
	*tchttp.BaseResponse
	Response *TagResourcesResponseParams `json:"Response"`
}

func (r *TagResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TagResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagWithDelete struct {
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// 是否可以删除
	CanDelete *uint64 `json:"CanDelete,omitempty" name:"CanDelete"`
}

// Predefined struct for user
type UnTagResourcesRequestParams struct {
	// 资源六段式列表。腾讯云使用资源六段式描述一个资源。可参考[访问管理](https://cloud.tencent.com/document/product/598/67350)-概览-接口列表-资源六段式信息
	// 例如：ResourceList.1 = qcs::${ServiceType}:${Region}:uin/${Account}:${ResourcePrefix}/${ResourceId}。
	// N取值范围：0~9
	ResourceList []*string `json:"ResourceList,omitempty" name:"ResourceList"`

	// 标签键。
	// 取值范围：0~9
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`
}

type UnTagResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 资源六段式列表。腾讯云使用资源六段式描述一个资源。可参考[访问管理](https://cloud.tencent.com/document/product/598/67350)-概览-接口列表-资源六段式信息
	// 例如：ResourceList.1 = qcs::${ServiceType}:${Region}:uin/${Account}:${ResourcePrefix}/${ResourceId}。
	// N取值范围：0~9
	ResourceList []*string `json:"ResourceList,omitempty" name:"ResourceList"`

	// 标签键。
	// 取值范围：0~9
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`
}

func (r *UnTagResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnTagResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceList")
	delete(f, "TagKeys")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnTagResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnTagResourcesResponseParams struct {
	// 失败资源信息。
	// 解绑标签成功时，返回的FailedResources为空。
	// 解绑标签失败或部分失败时，返回的FailedResources会显示失败资源的详细信息。
	FailedResources []*FailedResource `json:"FailedResources,omitempty" name:"FailedResources"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnTagResourcesResponse struct {
	*tchttp.BaseResponse
	Response *UnTagResourcesResponseParams `json:"Response"`
}

func (r *UnTagResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnTagResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateProjectRequestParams struct {
	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 禁用项目，1，禁用，0，启用
	Disable *int64 `json:"Disable,omitempty" name:"Disable"`

	// 备注
	Info *string `json:"Info,omitempty" name:"Info"`
}

type UpdateProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 禁用项目，1，禁用，0，启用
	Disable *int64 `json:"Disable,omitempty" name:"Disable"`

	// 备注
	Info *string `json:"Info,omitempty" name:"Info"`
}

func (r *UpdateProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ProjectName")
	delete(f, "Disable")
	delete(f, "Info")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateProjectResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateProjectResponse struct {
	*tchttp.BaseResponse
	Response *UpdateProjectResponseParams `json:"Response"`
}

func (r *UpdateProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateResourceTagValueRequestParams struct {
	// 资源关联的标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 修改后的标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// [ 资源六段式描述 ](https://cloud.tencent.com/document/product/598/10606)
	Resource *string `json:"Resource,omitempty" name:"Resource"`
}

type UpdateResourceTagValueRequest struct {
	*tchttp.BaseRequest
	
	// 资源关联的标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 修改后的标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// [ 资源六段式描述 ](https://cloud.tencent.com/document/product/598/10606)
	Resource *string `json:"Resource,omitempty" name:"Resource"`
}

func (r *UpdateResourceTagValueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateResourceTagValueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKey")
	delete(f, "TagValue")
	delete(f, "Resource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateResourceTagValueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateResourceTagValueResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateResourceTagValueResponse struct {
	*tchttp.BaseResponse
	Response *UpdateResourceTagValueResponseParams `json:"Response"`
}

func (r *UpdateResourceTagValueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateResourceTagValueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}