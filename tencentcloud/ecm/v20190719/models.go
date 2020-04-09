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

package v20190719

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Area struct {

	// 区域ID
	AreaId *string `json:"AreaId,omitempty" name:"AreaId"`

	// 区域名称
	AreaName *string `json:"AreaName,omitempty" name:"AreaName"`
}

type City struct {

	// 城市ID
	CityId *string `json:"CityId,omitempty" name:"CityId"`

	// 城市名称
	CityName *string `json:"CityName,omitempty" name:"CityName"`
}

type Country struct {

	// 国家ID
	CountryId *string `json:"CountryId,omitempty" name:"CountryId"`

	// 国家名称
	CountryName *string `json:"CountryName,omitempty" name:"CountryName"`
}

type CreateModuleRequest struct {
	*tchttp.BaseRequest

	// 模块名称，如视频直播模块。限制：模块名称不得以空格开头，长度不得超过60个字符。
	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`

	// 默认带宽，单位：M。范围不得超过带宽上下限，详看DescribeConfig。
	DefaultBandWidth *int64 `json:"DefaultBandWidth,omitempty" name:"DefaultBandWidth"`

	// 默认镜像，如img-qsdf3ff2。
	DefaultImageId *string `json:"DefaultImageId,omitempty" name:"DefaultImageId"`

	// 机型ID。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 默认系统盘大小，单位：G，默认大小为50G。范围不得超过系统盘上下限制，详看DescribeConfig。
	DefaultSystemDiskSize *int64 `json:"DefaultSystemDiskSize,omitempty" name:"DefaultSystemDiskSize"`

	// 默认数据盘大小，单位：G。范围不得超过数据盘范围大小，详看DescribeConfig。
	DefaultDataDiskSize *int64 `json:"DefaultDataDiskSize,omitempty" name:"DefaultDataDiskSize"`
}

func (r *CreateModuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateModuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateModuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 模块ID，创建模块成功后分配给该模块的ID。
		ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateModuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateModuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImageRequest struct {
	*tchttp.BaseRequest

	// 镜像ID列表。
	ImageIDSet []*string `json:"ImageIDSet,omitempty" name:"ImageIDSet" list`
}

func (r *DeleteImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteModuleRequest struct {
	*tchttp.BaseRequest

	// 模块ID。如：es-qn46snq8
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`
}

func (r *DeleteModuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteModuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteModuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteModuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteModuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBaseOverviewRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBaseOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBaseOverviewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBaseOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 模块数量，单位：个
		ModuleNum *int64 `json:"ModuleNum,omitempty" name:"ModuleNum"`

		// 节点数量，单位：个
		NodeNum *int64 `json:"NodeNum,omitempty" name:"NodeNum"`

		// cpu核数，单位：个
		VcpuNum *int64 `json:"VcpuNum,omitempty" name:"VcpuNum"`

		// 内存大小，单位：G
		MemoryNum *int64 `json:"MemoryNum,omitempty" name:"MemoryNum"`

		// 硬盘大小，单位：G
		StorageNum *int64 `json:"StorageNum,omitempty" name:"StorageNum"`

		// 昨日网络峰值,单位：M
		NetworkNum *int64 `json:"NetworkNum,omitempty" name:"NetworkNum"`

		// 实例数量，单位：台
		InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`

		// 运行中数量，单位：台
		RunningNum *int64 `json:"RunningNum,omitempty" name:"RunningNum"`

		// 安全隔离数量，单位：台
		IsolationNum *int64 `json:"IsolationNum,omitempty" name:"IsolationNum"`

		// 过期实例数量，单位：台
		ExpiredNum *int64 `json:"ExpiredNum,omitempty" name:"ExpiredNum"`

		// 即将过期实例数量，单位：台
		WillExpireNum *int64 `json:"WillExpireNum,omitempty" name:"WillExpireNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaseOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBaseOverviewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 网络带宽硬盘大小的范围信息。
		NetworkStorageRange *NetworkStorageRange `json:"NetworkStorageRange,omitempty" name:"NetworkStorageRange"`

		// 镜像操作系统白名单
	// 注意：此字段可能返回 null，表示取不到有效值。
		ImageWhiteSet []*string `json:"ImageWhiteSet,omitempty" name:"ImageWhiteSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageRequest struct {
	*tchttp.BaseRequest

	// 过滤条件，每次请求的Filters的上限为10，详细的过滤条件如下：
	// image-id - String - 是否必填： 否 - （过滤条件）按照镜像ID进行过滤
	// image-type - String - 是否必填： 否 - （过滤条件）按照镜像类型进行过滤。取值范围：
	// PRIVATE_IMAGE: 私有镜像 (本帐户创建的镜像) 
	// PUBLIC_IMAGE: 公共镜像 (腾讯云官方镜像)
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 镜像总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 镜像数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		ImageSet []*Image `json:"ImageSet,omitempty" name:"ImageSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceTypeConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeInstanceTypeConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceTypeConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceTypeConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 机型配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		InstanceTypeConfigSet []*InstanceTypeConfig `json:"InstanceTypeConfigSet,omitempty" name:"InstanceTypeConfigSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceTypeConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceTypeConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesDeniedActionsRequest struct {
	*tchttp.BaseRequest

	// 无
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet" list`
}

func (r *DescribeInstancesDeniedActionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesDeniedActionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesDeniedActionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例对应的禁止操作
		InstanceOperatorSet []*InstanceOperator `json:"InstanceOperatorSet,omitempty" name:"InstanceOperatorSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesDeniedActionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesDeniedActionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// zone      String      是否必填：否     （过滤条件）按照可用区中文名过滤,支持模糊匹配。
	// module-id      String      是否必填：否     （过滤条件）按照模块ID过滤。
	// instance-id      String      是否必填：否      （过滤条件）按照实例ID过滤。
	// instance-name      String      是否必填：否      （过滤条件）按照实例名称过滤,支持模糊匹配。
	// ip-address      String      是否必填：否      （过滤条件）按照实例的内网/公网IP过滤。
	// instance-uuid   string 是否必填：否 （过滤条件）按照uuid过滤实例列表。
	// instance-state  string  是否必填：否 （过滤条件）按照实例状态更新实例列表。
	// internet-service-provider      String      是否必填：否      （过滤条件）按照实例公网IP所属的运营商进行过滤。
	// tag-key      String      是否必填：否      （过滤条件）按照标签键进行过滤。
	// tag:tag-key      String      是否必填：否      （过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。
	// 若不传Filters参数则表示查询所有相关的实例信息。
	// 单次请求的Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20(如果查询结果数目大于等于20)，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的实例相关信息列表的长度。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 返回的实例相关信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		InstanceSet []*Instance `json:"InstanceSet,omitempty" name:"InstanceSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeModuleDetailRequest struct {
	*tchttp.BaseRequest

	// 模块ID，如em-qn46snq8。
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`
}

func (r *DescribeModuleDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeModuleDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeModuleDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 模块的详细信息，详细见数据结构中的ModuleInfo。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Module *Module `json:"Module,omitempty" name:"Module"`

		// 模块的统计信息，详细见数据结构中的ModuleCounterInfo。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ModuleCounter *ModuleCounter `json:"ModuleCounter,omitempty" name:"ModuleCounter"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeModuleDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeModuleDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeModuleRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// module-name - string - 是否必填：否 - （过滤条件）按照模块名称过滤。
	// module-id - string - 是否必填：否 - （过滤条件）按照模块ID过滤。
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeModuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeModuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeModuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的模块数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 模块详情信息的列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ModuleItemSet []*ModuleItem `json:"ModuleItemSet,omitempty" name:"ModuleItemSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeModuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeModuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNodeRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 节点详细信息的列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		NodeSet []*Node `json:"NodeSet,omitempty" name:"NodeSet" list`

		// 所有的节点数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePeakBaseOverviewRequest struct {
	*tchttp.BaseRequest

	// 开始时间（xxxx-xx-xx）如2019-08-14，默认为一周之前的日期。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间（xxxx-xx-xx）如2019-08-14，默认为昨天。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribePeakBaseOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePeakBaseOverviewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePeakBaseOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 基础峰值列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		PeakFamilyInfoSet []*PeakFamilyInfo `json:"PeakFamilyInfoSet,omitempty" name:"PeakFamilyInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePeakBaseOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePeakBaseOverviewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePeakNetworkOverviewRequest struct {
	*tchttp.BaseRequest

	// 开始时间（xxxx-xx-xx）如2019-08-14，默认为一周之前的日期。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间（xxxx-xx-xx）如2019-08-14，默认为昨天。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 过滤条件。
	// region    String      是否必填：否     （过滤条件）按照region过滤,不支持模糊匹配。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribePeakNetworkOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePeakNetworkOverviewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePeakNetworkOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 网络峰值数组。
	// 注意：此字段可能返回 null，表示取不到有效值。
		PeakNetworkRegionSet []*PeakNetworkRegionInfo `json:"PeakNetworkRegionSet,omitempty" name:"PeakNetworkRegionSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePeakNetworkOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePeakNetworkOverviewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DiskInfo struct {

	// 磁盘类型：LOCAL_BASIC
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 磁盘ID
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// 磁盘大小（GB）
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

type EnhancedService struct {

	// 是否开启云镜服务。
	SecurityService *RunSecurityServiceEnabled `json:"SecurityService,omitempty" name:"SecurityService"`

	// 是否开启云监控服务。
	MonitorService *RunMonitorServiceEnabled `json:"MonitorService,omitempty" name:"MonitorService"`
}

type Filter struct {

	// 过滤字段名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤字段内容数组
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type ISP struct {

	// 运营商ID
	ISPId *string `json:"ISPId,omitempty" name:"ISPId"`

	// 运营商名称
	ISPName *string `json:"ISPName,omitempty" name:"ISPName"`
}

type ISPCounter struct {

	// 运营商名称
	ProviderName *string `json:"ProviderName,omitempty" name:"ProviderName"`

	// 节点数量
	ProviderNodeNum *int64 `json:"ProviderNodeNum,omitempty" name:"ProviderNodeNum"`

	// 实例数量
	ProvederInstanceNum *int64 `json:"ProvederInstanceNum,omitempty" name:"ProvederInstanceNum"`

	// Zone实例信息结构体数组
	ZoneInstanceInfoSet []*ZoneInstanceInfo `json:"ZoneInstanceInfoSet,omitempty" name:"ZoneInstanceInfoSet" list`
}

type Image struct {

	// 镜像ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 镜像状态
	ImageState *string `json:"ImageState,omitempty" name:"ImageState"`

	// 镜像类型
	ImageType *string `json:"ImageType,omitempty" name:"ImageType"`

	// 操作系统名称
	ImageOsName *string `json:"ImageOsName,omitempty" name:"ImageOsName"`

	// 镜像描述
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`

	// 镜像导入时间
	ImageCreateTime *string `json:"ImageCreateTime,omitempty" name:"ImageCreateTime"`

	// 操作系统位数
	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`

	// 操作系统类型
	OsType *string `json:"OsType,omitempty" name:"OsType"`

	// 操作系统版本
	OsVersion *string `json:"OsVersion,omitempty" name:"OsVersion"`

	// 操作系统平台
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 镜像所有者
	ImageOwner *int64 `json:"ImageOwner,omitempty" name:"ImageOwner"`

	// 镜像大小。单位：GB
	ImageSize *int64 `json:"ImageSize,omitempty" name:"ImageSize"`

	// 镜像来源信息
	SrcImage *SrcImage `json:"SrcImage,omitempty" name:"SrcImage"`
}

type ImportImageRequest struct {
	*tchttp.BaseRequest

	// 镜像的Id。
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 镜像的描述。
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`

	// 源地域
	SourceRegion *string `json:"SourceRegion,omitempty" name:"SourceRegion"`
}

func (r *ImportImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportImageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImportImageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportImageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Instance struct {

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称，如ens-34241f3s。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例状态。取值范围：
	// PENDING：表示创建中
	// LAUNCH_FAILED：表示创建失败
	// RUNNING：表示运行中
	// STOPPED：表示关机
	// STARTING：表示开机中
	// STOPPING：表示关机中
	// REBOOTING：表示重启中
	// SHUTDOWN：表示停止待销毁
	// TERMINATING：表示销毁中。
	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`

	// 实例当前使用的镜像的信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Image *Image `json:"Image,omitempty" name:"Image"`

	// 实例当前所属的模块简要信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SimpleModule *SimpleModule `json:"SimpleModule,omitempty" name:"SimpleModule"`

	// 实例所在的位置相关信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Position *Position `json:"Position,omitempty" name:"Position"`

	// 实例的网络相关信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Internet *Internet `json:"Internet,omitempty" name:"Internet"`

	// 实例的配置相关信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceTypeConfig *InstanceTypeConfig `json:"InstanceTypeConfig,omitempty" name:"InstanceTypeConfig"`

	// 实例的创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 实例的标签信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet" list`

	// 实例最后一次操作。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestOperation *string `json:"LatestOperation,omitempty" name:"LatestOperation"`

	// 实例最后一次操作结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestOperationState *string `json:"LatestOperationState,omitempty" name:"LatestOperationState"`

	// 实例业务状态。取值范围：
	// NORMAL：表示正常状态的实例
	// EXPIRED：表示过期的实例
	// PROTECTIVELY_ISOLATED：表示被安全隔离的实例。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestrictState *string `json:"RestrictState,omitempty" name:"RestrictState"`

	// 系统盘大小，单位GB。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SystemDiskSize *int64 `json:"SystemDiskSize,omitempty" name:"SystemDiskSize"`

	// 数据盘大小，单位GB。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataDiskSize *int64 `json:"DataDiskSize,omitempty" name:"DataDiskSize"`

	// UUID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UUID *string `json:"UUID,omitempty" name:"UUID"`

	// 付费方式。
	//     0为后付费。
	//     1为预付费。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 过期时间。格式为yyyy-mm-dd HH:mm:ss。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 隔离时间。格式为yyyy-mm-dd HH:mm:ss。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolatedTime *string `json:"IsolatedTime,omitempty" name:"IsolatedTime"`

	// 是否自动续费。
	//       0为不自动续费。
	//       1为自动续费。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 过期状态。
	//     NORMAL 表示机器运行正常。
	//     WILL_EXPIRE 表示即将过期。
	//     EXPIRED 表示已过期。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireState *string `json:"ExpireState,omitempty" name:"ExpireState"`

	// 系统盘信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SystemDisk *DiskInfo `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 数据盘信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataDisks []*DiskInfo `json:"DataDisks,omitempty" name:"DataDisks" list`
}

type InstanceFamilyConfig struct {

	// 机型名称
	InstanceFamilyName *string `json:"InstanceFamilyName,omitempty" name:"InstanceFamilyName"`

	// 机型ID
	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
}

type InstanceFamilyTypeConfig struct {

	// 实例机型系列类型Id
	InstanceFamilyType *string `json:"InstanceFamilyType,omitempty" name:"InstanceFamilyType"`

	// 实例机型系列类型名称
	InstanceFamilyTypeName *string `json:"InstanceFamilyTypeName,omitempty" name:"InstanceFamilyTypeName"`
}

type InstanceOperator struct {

	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例禁止的操作
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeniedActions []*OperatorAction `json:"DeniedActions,omitempty" name:"DeniedActions" list`
}

type InstanceTypeConfig struct {

	// 机型族配置信息
	InstanceFamilyConfig *InstanceFamilyConfig `json:"InstanceFamilyConfig,omitempty" name:"InstanceFamilyConfig"`

	// 机型
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// CPU核数
	Vcpu *int64 `json:"Vcpu,omitempty" name:"Vcpu"`

	// 内存大小
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 主频
	Frequency *string `json:"Frequency,omitempty" name:"Frequency"`

	// 处理器型号
	CpuModelName *string `json:"CpuModelName,omitempty" name:"CpuModelName"`

	// 机型族类别配置信息
	InstanceFamilyTypeConfig *InstanceFamilyTypeConfig `json:"InstanceFamilyTypeConfig,omitempty" name:"InstanceFamilyTypeConfig"`

	// 机型额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtInfo *string `json:"ExtInfo,omitempty" name:"ExtInfo"`
}

type Internet struct {

	// 实例的内网相关信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIPAddressSet []*PrivateIPAddressInfo `json:"PrivateIPAddressSet,omitempty" name:"PrivateIPAddressSet" list`

	// 实例的公网相关信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIPAddressSet []*PublicIPAddressInfo `json:"PublicIPAddressSet,omitempty" name:"PublicIPAddressSet" list`
}

type ModifyInstancesAttributeRequest struct {
	*tchttp.BaseRequest

	// 待修改的实例ID列表。在单次请求的过程中，请求实例数上限为100。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet" list`

	// 修改成功后显示的实例名称，不得超过60个字符，不传则名称显示为空。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *ModifyInstancesAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstancesAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstancesAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstancesAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyModuleImageRequest struct {
	*tchttp.BaseRequest

	// 默认镜像ID
	DefaultImageId *string `json:"DefaultImageId,omitempty" name:"DefaultImageId"`

	// 模块ID
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`
}

func (r *ModifyModuleImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyModuleImageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyModuleImageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyModuleImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyModuleImageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyModuleNameRequest struct {
	*tchttp.BaseRequest

	// 模块ID。
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// 模块名称。
	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

func (r *ModifyModuleNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyModuleNameRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyModuleNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyModuleNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyModuleNameResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyModuleNetworkRequest struct {
	*tchttp.BaseRequest

	// 模块Id
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// 默认带宽上限
	DefaultBandwidth *int64 `json:"DefaultBandwidth,omitempty" name:"DefaultBandwidth"`
}

func (r *ModifyModuleNetworkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyModuleNetworkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyModuleNetworkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyModuleNetworkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyModuleNetworkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Module struct {

	// 模块Id
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// 模块名称
	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`

	// 模块状态：
	// NORMAL：正常
	// DELETING：删除中 
	// DELETEFAILED：删除失败
	ModuleState *string `json:"ModuleState,omitempty" name:"ModuleState"`

	// 默认系统盘大小
	DefaultSystemDiskSize *int64 `json:"DefaultSystemDiskSize,omitempty" name:"DefaultSystemDiskSize"`

	// 默认数据盘大小
	DefaultDataDiskSize *int64 `json:"DefaultDataDiskSize,omitempty" name:"DefaultDataDiskSize"`

	// 默认机型
	InstanceTypeConfig *InstanceTypeConfig `json:"InstanceTypeConfig,omitempty" name:"InstanceTypeConfig"`

	// 默认镜像
	DefaultImage *Image `json:"DefaultImage,omitempty" name:"DefaultImage"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 默认带宽
	DefaultBandwidth *int64 `json:"DefaultBandwidth,omitempty" name:"DefaultBandwidth"`
}

type ModuleCounter struct {

	// 运营商统计信息列表
	ISPCounterSet []*ISPCounter `json:"ISPCounterSet,omitempty" name:"ISPCounterSet" list`

	// 省份数量
	ProvinceNum *int64 `json:"ProvinceNum,omitempty" name:"ProvinceNum"`

	// 城市数量
	CityNum *int64 `json:"CityNum,omitempty" name:"CityNum"`

	// 节点数量
	NodeNum *int64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// 实例数量
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`
}

type ModuleItem struct {

	// 节点实例统计信息
	NodeInstanceNum *NodeInstanceNum `json:"NodeInstanceNum,omitempty" name:"NodeInstanceNum"`

	// 模块信息
	Module *Module `json:"Module,omitempty" name:"Module"`
}

type NetworkStorageRange struct {

	// 网络带宽上限
	MaxBandwidth *int64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`

	// 数据盘上限
	MaxSystemDiskSize *int64 `json:"MaxSystemDiskSize,omitempty" name:"MaxSystemDiskSize"`

	// 网络带宽下限
	MinBandwidth *int64 `json:"MinBandwidth,omitempty" name:"MinBandwidth"`

	// 数据盘下限
	MinSystemDiskSize *int64 `json:"MinSystemDiskSize,omitempty" name:"MinSystemDiskSize"`

	// 最大数据盘大小
	MaxDataDiskSize *int64 `json:"MaxDataDiskSize,omitempty" name:"MaxDataDiskSize"`

	// 最小数据盘大小
	MinDataDiskSize *int64 `json:"MinDataDiskSize,omitempty" name:"MinDataDiskSize"`

	// 建议带宽
	SuggestBandwidth *int64 `json:"SuggestBandwidth,omitempty" name:"SuggestBandwidth"`

	// 建议硬盘大小
	SuggestDataDiskSize *int64 `json:"SuggestDataDiskSize,omitempty" name:"SuggestDataDiskSize"`

	// 建议系统盘大小
	SuggestSystemDiskSize *int64 `json:"SuggestSystemDiskSize,omitempty" name:"SuggestSystemDiskSize"`

	// Cpu核数峰值
	MaxVcpu *int64 `json:"MaxVcpu,omitempty" name:"MaxVcpu"`

	// Cpu核最小值
	MinVcpu *int64 `json:"MinVcpu,omitempty" name:"MinVcpu"`

	// 单次请求最大cpu核数
	MaxVcpuPerReq *int64 `json:"MaxVcpuPerReq,omitempty" name:"MaxVcpuPerReq"`

	// 带宽步长
	PerBandwidth *int64 `json:"PerBandwidth,omitempty" name:"PerBandwidth"`

	// 数据盘步长
	PerDataDisk *int64 `json:"PerDataDisk,omitempty" name:"PerDataDisk"`

	// 总模块数量
	MaxModuleNum *int64 `json:"MaxModuleNum,omitempty" name:"MaxModuleNum"`
}

type Node struct {

	// zone信息
	ZoneInfo *ZoneInfo `json:"ZoneInfo,omitempty" name:"ZoneInfo"`

	// 国家信息
	Country *Country `json:"Country,omitempty" name:"Country"`

	// 区域信息
	Area *Area `json:"Area,omitempty" name:"Area"`

	// 省份信息
	Province *Province `json:"Province,omitempty" name:"Province"`

	// 城市信息
	City *City `json:"City,omitempty" name:"City"`

	// Region信息
	RegionInfo *RegionInfo `json:"RegionInfo,omitempty" name:"RegionInfo"`

	// 运营商列表
	ISPSet []*ISP `json:"ISPSet,omitempty" name:"ISPSet" list`

	// 运营商数量
	ISPNum *int64 `json:"ISPNum,omitempty" name:"ISPNum"`
}

type NodeInstanceNum struct {

	// 节点数量
	NodeNum *int64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// 实例数量
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`
}

type OperatorAction struct {

	// 可执行操作
	Action *string `json:"Action,omitempty" name:"Action"`

	// 编码Code
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitempty" name:"Code"`

	// 具体信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`
}

type PeakBase struct {

	// CPU峰值
	PeakCpuNum *int64 `json:"PeakCpuNum,omitempty" name:"PeakCpuNum"`

	// 内存峰值
	PeakMemoryNum *int64 `json:"PeakMemoryNum,omitempty" name:"PeakMemoryNum"`

	// 硬盘峰值
	PeakStorageNum *int64 `json:"PeakStorageNum,omitempty" name:"PeakStorageNum"`

	// 记录时间
	RecordTime *string `json:"RecordTime,omitempty" name:"RecordTime"`
}

type PeakFamilyInfo struct {

	// 机型类别信息。
	InstanceFamily *InstanceFamilyTypeConfig `json:"InstanceFamily,omitempty" name:"InstanceFamily"`

	// 基础数据峰值信息。
	PeakBaseSet []*PeakBase `json:"PeakBaseSet,omitempty" name:"PeakBaseSet" list`
}

type PeakNetwork struct {

	// 记录时间。
	RecordTime *string `json:"RecordTime,omitempty" name:"RecordTime"`

	// 入带宽数据。
	PeakInNetwork *string `json:"PeakInNetwork,omitempty" name:"PeakInNetwork"`

	// 出带宽数据。
	PeakOutNetwork *string `json:"PeakOutNetwork,omitempty" name:"PeakOutNetwork"`
}

type PeakNetworkRegionInfo struct {

	// region信息
	Region *string `json:"Region,omitempty" name:"Region"`

	// 网络峰值集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeakNetworkSet []*PeakNetwork `json:"PeakNetworkSet,omitempty" name:"PeakNetworkSet" list`
}

type Position struct {

	// 实例所在的Zone的信息。
	ZoneInfo *ZoneInfo `json:"ZoneInfo,omitempty" name:"ZoneInfo"`

	// 实例所在的国家的信息。
	Country *Country `json:"Country,omitempty" name:"Country"`

	// 实例所在的Area的信息。
	Area *Area `json:"Area,omitempty" name:"Area"`

	// 实例所在的省份的信息。
	Province *Province `json:"Province,omitempty" name:"Province"`

	// 实例所在的城市的信息。
	City *City `json:"City,omitempty" name:"City"`

	// 实例所在的Region的信息。
	RegionInfo *RegionInfo `json:"RegionInfo,omitempty" name:"RegionInfo"`
}

type PrivateIPAddressInfo struct {

	// 实例的内网ip。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIPAddress *string `json:"PrivateIPAddress,omitempty" name:"PrivateIPAddress"`
}

type Province struct {

	// 省份Id
	ProvinceId *string `json:"ProvinceId,omitempty" name:"ProvinceId"`

	// 省份名称
	ProvinceName *string `json:"ProvinceName,omitempty" name:"ProvinceName"`
}

type PublicIPAddressInfo struct {

	// 计费模式。
	ChargeMode *string `json:"ChargeMode,omitempty" name:"ChargeMode"`

	// 实例的公网ip。
	PublicIPAddress *string `json:"PublicIPAddress,omitempty" name:"PublicIPAddress"`

	// 实例的公网ip所属的运营商。
	ISP *ISP `json:"ISP,omitempty" name:"ISP"`

	// 实例的最大出带宽上限。
	MaxBandwidthOut *int64 `json:"MaxBandwidthOut,omitempty" name:"MaxBandwidthOut"`
}

type RebootInstancesRequest struct {
	*tchttp.BaseRequest

	// 待重启的实例ID列表。在单次请求的过程中，单个region下的请求实例数上限为100。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet" list`

	// 是否在正常重启失败后选择强制重启实例。取值范围：
	// TRUE：表示在正常重启失败后进行强制重启；
	// FALSE：表示在正常重启失败后不进行强制重启；
	// 默认取值：FALSE。
	ForceReboot *bool `json:"ForceReboot,omitempty" name:"ForceReboot"`

	// 关机类型。取值范围：
	// SOFT：表示软关机
	// HARD：表示硬关机
	// SOFT_FIRST：表示优先软关机，失败再执行硬关机
	// 
	// 默认取值：SOFT。
	StopType *string `json:"StopType,omitempty" name:"StopType"`
}

func (r *RebootInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RebootInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RebootInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RebootInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RebootInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {

	// Region
	Region *string `json:"Region,omitempty" name:"Region"`

	// Region名称
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// RegionID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
}

type ResetInstancesMaxBandwidthRequest struct {
	*tchttp.BaseRequest

	// 待重置带宽上限的实例ID列表。在单次请求的过程中，单个region下的请求实例数上限为100。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet" list`

	// 修改后的最大带宽上限。
	MaxBandwidthOut *int64 `json:"MaxBandwidthOut,omitempty" name:"MaxBandwidthOut"`
}

func (r *ResetInstancesMaxBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstancesMaxBandwidthRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesMaxBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetInstancesMaxBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstancesMaxBandwidthResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesRequest struct {
	*tchttp.BaseRequest

	// 待重装的实例ID列表。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet" list`

	// 重装使用的镜像ID，若未指定，则使用各个实例当前的镜像进行重装。
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 密码设置，若未指定，则后续将以站内信的形式通知密码。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 是否开启云监控和云镜服务，未指定时默认开启。
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`
}

func (r *ResetInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunMonitorServiceEnabled struct {

	// 是否开启。
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type RunSecurityServiceEnabled struct {

	// 是否开启。
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// 云镜版本：0 基础版，1 专业版
	Version *int64 `json:"Version,omitempty" name:"Version"`
}

type SimpleModule struct {

	// 模块ID
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// 模块名称
	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

type SrcImage struct {

	// 镜像id
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 系统名称
	ImageOsName *string `json:"ImageOsName,omitempty" name:"ImageOsName"`

	// 镜像描述
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`

	// 区域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 区域ID
	RegionID *int64 `json:"RegionID,omitempty" name:"RegionID"`

	// 区域名称
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
}

type Tag struct {

	// 标签的键。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 标签的值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type TerminateInstancesRequest struct {
	*tchttp.BaseRequest

	// 待销毁的实例ID列表。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet" list`

	// 是否定时销毁，默认为否。
	TerminateDelay *bool `json:"TerminateDelay,omitempty" name:"TerminateDelay"`

	// 定时销毁的时间，格式形如："2019-08-05 12:01:30"，若非定时销毁，则此参数被忽略。
	TerminateTime *string `json:"TerminateTime,omitempty" name:"TerminateTime"`
}

func (r *TerminateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TerminateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ZoneInfo struct {

	// ZoneId
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// ZoneName
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Zone
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type ZoneInstanceInfo struct {

	// Zone名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 实例数量
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`
}
