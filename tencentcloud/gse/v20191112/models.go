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

package v20191112

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Alias struct {
	// 别名的唯一标识符
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`

	// 别名的全局唯一资源标识符
	AliasArn *string `json:"AliasArn,omitempty" name:"AliasArn"`

	// 名字，长度不小于1字符不超过1024字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 别名的可读说明，长度不小于1字符不超过1024字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 别名的路由配置
	RoutingStrategy *RoutingStrategy `json:"RoutingStrategy,omitempty" name:"RoutingStrategy"`

	// 创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 上次修改此数据对象的时间
	LastUpdatedTime *string `json:"LastUpdatedTime,omitempty" name:"LastUpdatedTime"`

	// 标签列表，最大长度50组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type Asset struct {
	// 生成包ID
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// 生成包名字，最小长度为1，最大长度为64
	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`

	// 生成包版本，最小长度为1，最大长度为64
	AssetVersion *string `json:"AssetVersion,omitempty" name:"AssetVersion"`

	// 生成包可运行的操作系统，暂时只支持CentOS7.16
	OperateSystem *string `json:"OperateSystem,omitempty" name:"OperateSystem"`

	// 生成包状态，0代表上传中，1代表上传失败，2代表上传成功
	Stauts *int64 `json:"Stauts,omitempty" name:"Stauts"`

	// 生成包大小
	Size *string `json:"Size,omitempty" name:"Size"`

	// 生成包创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 生成包绑定的Fleet个数，最小值为0
	BindFleetNum *int64 `json:"BindFleetNum,omitempty" name:"BindFleetNum"`

	// 生成包的全局唯一资源标识符
	AssetArn *string `json:"AssetArn,omitempty" name:"AssetArn"`

	// 生成包支持的操作系统镜像id
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 生成包支持的操作系统类型
	OsType *string `json:"OsType,omitempty" name:"OsType"`

	// 生成包资源类型，ASSET 或者 IMAGE；ASSET 代表是原有生成包类型，IMAGE 为扩充使用镜像类型
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 镜像资源共享类型，当 ResourceType 为 IMAGE 时该字段有意义，SHARED 表示共享、SHARED_IMAGE 表示未共享；ResourceType 为 ASSET 时这里返回 UNKNOWN_SHARED 用于占位
	SharingStatus *string `json:"SharingStatus,omitempty" name:"SharingStatus"`

	// 标签列表，最大长度50组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type AssetCredentials struct {
	// 临时证书密钥ID
	TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`

	// 临时证书密钥Key
	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`

	// 临时证书Token
	Token *string `json:"Token,omitempty" name:"Token"`
}

type AssetSupportSys struct {
	// 生成包操作系统的镜像Id
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 生成包操作系统的类型
	OsType *string `json:"OsType,omitempty" name:"OsType"`

	// 生成包操作系统的位数
	OsBit *int64 `json:"OsBit,omitempty" name:"OsBit"`

	// 生成包操作系统的版本
	OsVersion *string `json:"OsVersion,omitempty" name:"OsVersion"`
}

// Predefined struct for user
type AttachCcnInstancesRequestParams struct {
	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 云联网账号 Uin
	AccountId *string `json:"AccountId,omitempty" name:"AccountId"`

	// 云联网 Id
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

type AttachCcnInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 云联网账号 Uin
	AccountId *string `json:"AccountId,omitempty" name:"AccountId"`

	// 云联网 Id
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

func (r *AttachCcnInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachCcnInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	delete(f, "AccountId")
	delete(f, "CcnId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachCcnInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachCcnInstancesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AttachCcnInstancesResponse struct {
	*tchttp.BaseResponse
	Response *AttachCcnInstancesResponseParams `json:"Response"`
}

func (r *AttachCcnInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachCcnInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnInfo struct {
	// 云联网所属账号
	AccountId *string `json:"AccountId,omitempty" name:"AccountId"`

	// 云联网id
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

type CcnInstanceSets struct {
	// 云联网账号 Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountId *string `json:"AccountId,omitempty" name:"AccountId"`

	// 云联网 Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// 云联网关联时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 云联网实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 云联网状态：申请中、已连接、已过期、已拒绝、已删除、失败的、关联中、解关联中、解关联失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitempty" name:"State"`
}

// Predefined struct for user
type CopyFleetRequestParams struct {
	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 复制数量，最小值1，最大值为剩余配额，可以根据[获取用户配额](https://cloud.tencent.com/document/product/1165/48732)接口获取。
	CopyNumber *int64 `json:"CopyNumber,omitempty" name:"CopyNumber"`

	// 生成包 Id
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// 描述，最小长度0，最大长度100
	Description *string `json:"Description,omitempty" name:"Description"`

	// 网络配置
	InboundPermissions []*InboundPermission `json:"InboundPermissions,omitempty" name:"InboundPermissions"`

	// 服务器类型，参数根据[获取服务器实例类型列表](https://cloud.tencent.com/document/product/1165/48732)接口获取。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 服务器舰队类型，目前只支持ON_DEMAND类型
	FleetType *string `json:"FleetType,omitempty" name:"FleetType"`

	// 服务器舰队名称，最小长度1，最大长度50
	Name *string `json:"Name,omitempty" name:"Name"`

	// 保护策略：不保护NoProtection、完全保护FullProtection、时限保护TimeLimitProtection
	NewGameServerSessionProtectionPolicy *string `json:"NewGameServerSessionProtectionPolicy,omitempty" name:"NewGameServerSessionProtectionPolicy"`

	// 资源创建限制策略
	ResourceCreationLimitPolicy *ResourceCreationLimitPolicy `json:"ResourceCreationLimitPolicy,omitempty" name:"ResourceCreationLimitPolicy"`

	// 进程配置
	RuntimeConfiguration *RuntimeConfiguration `json:"RuntimeConfiguration,omitempty" name:"RuntimeConfiguration"`

	// 时限保护超时时间，默认60分钟，最小值5，最大值1440；当NewGameSessionProtectionPolicy为TimeLimitProtection时参数有效
	GameServerSessionProtectionTimeLimit *int64 `json:"GameServerSessionProtectionTimeLimit,omitempty" name:"GameServerSessionProtectionTimeLimit"`

	// 是否选择扩缩容：SCALING_SELECTED 或者 SCALING_UNSELECTED；默认是 SCALING_UNSELECTED
	SelectedScalingType *string `json:"SelectedScalingType,omitempty" name:"SelectedScalingType"`

	// 是否选择云联网：CCN_SELECTED_BEFORE_CREATE（创建前关联）， CCN_SELECTED_AFTER_CREATE（创建后关联）或者 CCN_UNSELECTED（不关联）；默认是 CCN_UNSELECTED
	SelectedCcnType *string `json:"SelectedCcnType,omitempty" name:"SelectedCcnType"`

	// 标签列表，最大长度50组
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 系统盘，储存类型为 SSD 云硬盘（CLOUD_SSD）时，100-500GB；储存类型为高性能云硬盘（CLOUD_PREMIUM）时，50-500GB；容量以1为单位
	SystemDiskInfo *DiskInfo `json:"SystemDiskInfo,omitempty" name:"SystemDiskInfo"`

	// 数据盘，储存类型为 SSD 云硬盘（CLOUD_SSD）时，100-32000GB；储存类型为高性能云硬盘（CLOUD_PREMIUM）时，10-32000GB；容量以10为单位
	DataDiskInfo []*DiskInfo `json:"DataDiskInfo,omitempty" name:"DataDiskInfo"`

	// 是否选择复制定时器策略：TIMER_SELECTED 或者 TIMER_UNSELECTED；默认是 TIMER_UNSELECTED
	SelectedTimerType *string `json:"SelectedTimerType,omitempty" name:"SelectedTimerType"`

	// 云联网信息，包含对应的账号信息及所属id
	CcnInfos []*CcnInfo `json:"CcnInfos,omitempty" name:"CcnInfos"`

	// fleet公网出带宽最大值，默认100Mbps，范围1-200Mbps
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
}

type CopyFleetRequest struct {
	*tchttp.BaseRequest
	
	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 复制数量，最小值1，最大值为剩余配额，可以根据[获取用户配额](https://cloud.tencent.com/document/product/1165/48732)接口获取。
	CopyNumber *int64 `json:"CopyNumber,omitempty" name:"CopyNumber"`

	// 生成包 Id
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// 描述，最小长度0，最大长度100
	Description *string `json:"Description,omitempty" name:"Description"`

	// 网络配置
	InboundPermissions []*InboundPermission `json:"InboundPermissions,omitempty" name:"InboundPermissions"`

	// 服务器类型，参数根据[获取服务器实例类型列表](https://cloud.tencent.com/document/product/1165/48732)接口获取。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 服务器舰队类型，目前只支持ON_DEMAND类型
	FleetType *string `json:"FleetType,omitempty" name:"FleetType"`

	// 服务器舰队名称，最小长度1，最大长度50
	Name *string `json:"Name,omitempty" name:"Name"`

	// 保护策略：不保护NoProtection、完全保护FullProtection、时限保护TimeLimitProtection
	NewGameServerSessionProtectionPolicy *string `json:"NewGameServerSessionProtectionPolicy,omitempty" name:"NewGameServerSessionProtectionPolicy"`

	// 资源创建限制策略
	ResourceCreationLimitPolicy *ResourceCreationLimitPolicy `json:"ResourceCreationLimitPolicy,omitempty" name:"ResourceCreationLimitPolicy"`

	// 进程配置
	RuntimeConfiguration *RuntimeConfiguration `json:"RuntimeConfiguration,omitempty" name:"RuntimeConfiguration"`

	// 时限保护超时时间，默认60分钟，最小值5，最大值1440；当NewGameSessionProtectionPolicy为TimeLimitProtection时参数有效
	GameServerSessionProtectionTimeLimit *int64 `json:"GameServerSessionProtectionTimeLimit,omitempty" name:"GameServerSessionProtectionTimeLimit"`

	// 是否选择扩缩容：SCALING_SELECTED 或者 SCALING_UNSELECTED；默认是 SCALING_UNSELECTED
	SelectedScalingType *string `json:"SelectedScalingType,omitempty" name:"SelectedScalingType"`

	// 是否选择云联网：CCN_SELECTED_BEFORE_CREATE（创建前关联）， CCN_SELECTED_AFTER_CREATE（创建后关联）或者 CCN_UNSELECTED（不关联）；默认是 CCN_UNSELECTED
	SelectedCcnType *string `json:"SelectedCcnType,omitempty" name:"SelectedCcnType"`

	// 标签列表，最大长度50组
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 系统盘，储存类型为 SSD 云硬盘（CLOUD_SSD）时，100-500GB；储存类型为高性能云硬盘（CLOUD_PREMIUM）时，50-500GB；容量以1为单位
	SystemDiskInfo *DiskInfo `json:"SystemDiskInfo,omitempty" name:"SystemDiskInfo"`

	// 数据盘，储存类型为 SSD 云硬盘（CLOUD_SSD）时，100-32000GB；储存类型为高性能云硬盘（CLOUD_PREMIUM）时，10-32000GB；容量以10为单位
	DataDiskInfo []*DiskInfo `json:"DataDiskInfo,omitempty" name:"DataDiskInfo"`

	// 是否选择复制定时器策略：TIMER_SELECTED 或者 TIMER_UNSELECTED；默认是 TIMER_UNSELECTED
	SelectedTimerType *string `json:"SelectedTimerType,omitempty" name:"SelectedTimerType"`

	// 云联网信息，包含对应的账号信息及所属id
	CcnInfos []*CcnInfo `json:"CcnInfos,omitempty" name:"CcnInfos"`

	// fleet公网出带宽最大值，默认100Mbps，范围1-200Mbps
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
}

func (r *CopyFleetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyFleetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	delete(f, "CopyNumber")
	delete(f, "AssetId")
	delete(f, "Description")
	delete(f, "InboundPermissions")
	delete(f, "InstanceType")
	delete(f, "FleetType")
	delete(f, "Name")
	delete(f, "NewGameServerSessionProtectionPolicy")
	delete(f, "ResourceCreationLimitPolicy")
	delete(f, "RuntimeConfiguration")
	delete(f, "GameServerSessionProtectionTimeLimit")
	delete(f, "SelectedScalingType")
	delete(f, "SelectedCcnType")
	delete(f, "Tags")
	delete(f, "SystemDiskInfo")
	delete(f, "DataDiskInfo")
	delete(f, "SelectedTimerType")
	delete(f, "CcnInfos")
	delete(f, "InternetMaxBandwidthOut")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CopyFleetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopyFleetResponseParams struct {
	// 服务器舰队属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetAttributes []*FleetAttributes `json:"FleetAttributes,omitempty" name:"FleetAttributes"`

	// 服务器舰队数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CopyFleetResponse struct {
	*tchttp.BaseResponse
	Response *CopyFleetResponseParams `json:"Response"`
}

func (r *CopyFleetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyFleetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAliasRequestParams struct {
	// 名字，长度不小于1字符不超过1024字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 别名的路由配置
	RoutingStrategy *RoutingStrategy `json:"RoutingStrategy,omitempty" name:"RoutingStrategy"`

	// 别名的可读说明，长度不小于1字符不超过1024字符
	Description *string `json:"Description,omitempty" name:"Description"`

	// 标签列表，最大长度50组
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CreateAliasRequest struct {
	*tchttp.BaseRequest
	
	// 名字，长度不小于1字符不超过1024字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 别名的路由配置
	RoutingStrategy *RoutingStrategy `json:"RoutingStrategy,omitempty" name:"RoutingStrategy"`

	// 别名的可读说明，长度不小于1字符不超过1024字符
	Description *string `json:"Description,omitempty" name:"Description"`

	// 标签列表，最大长度50组
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateAliasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAliasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "RoutingStrategy")
	delete(f, "Description")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAliasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAliasResponseParams struct {
	// 别名对象
	Alias *Alias `json:"Alias,omitempty" name:"Alias"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAliasResponse struct {
	*tchttp.BaseResponse
	Response *CreateAliasResponseParams `json:"Response"`
}

func (r *CreateAliasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAliasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAssetRequestParams struct {
	// 生成包的ZIP包名，例如：server.zip
	BucketKey *string `json:"BucketKey,omitempty" name:"BucketKey"`

	// 生成包名字，最小长度为1，最大长度为64
	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`

	// 生成包版本，最小长度为1，最大长度为64
	AssetVersion *string `json:"AssetVersion,omitempty" name:"AssetVersion"`

	// 生成包所在地域，详见产品支持的 [地域列表](https://cloud.tencent.com/document/api/1165/42053#.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8)
	AssetRegion *string `json:"AssetRegion,omitempty" name:"AssetRegion"`

	// 生成包可运行的操作系统，若传入参数为CentOS7.16则不需要传入ImageId字段，否则，需要传入Imageid字段（该方式是为了兼容之前的版本，后续建议使用ImageId来替代该字段）。这里可通过[DescribeAssetSystems](https://cloud.tencent.com/document/product/1165/49191)接口获取asset支持的操作系统进行传入（使用AssetSupportSys的OsVersion字段）
	OperateSystem *string `json:"OperateSystem,omitempty" name:"OperateSystem"`

	// 生成包支持的操作系统镜像id，若传入OperateSystem字段的值是CentOS7.16，则不需要传入该值；如果不是，则需要通过[DescribeAssetSystems](https://cloud.tencent.com/document/product/1165/49191)接口获取asset支持的操作系统ImageId进行传入
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 标签列表，最大长度50组
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CreateAssetRequest struct {
	*tchttp.BaseRequest
	
	// 生成包的ZIP包名，例如：server.zip
	BucketKey *string `json:"BucketKey,omitempty" name:"BucketKey"`

	// 生成包名字，最小长度为1，最大长度为64
	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`

	// 生成包版本，最小长度为1，最大长度为64
	AssetVersion *string `json:"AssetVersion,omitempty" name:"AssetVersion"`

	// 生成包所在地域，详见产品支持的 [地域列表](https://cloud.tencent.com/document/api/1165/42053#.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8)
	AssetRegion *string `json:"AssetRegion,omitempty" name:"AssetRegion"`

	// 生成包可运行的操作系统，若传入参数为CentOS7.16则不需要传入ImageId字段，否则，需要传入Imageid字段（该方式是为了兼容之前的版本，后续建议使用ImageId来替代该字段）。这里可通过[DescribeAssetSystems](https://cloud.tencent.com/document/product/1165/49191)接口获取asset支持的操作系统进行传入（使用AssetSupportSys的OsVersion字段）
	OperateSystem *string `json:"OperateSystem,omitempty" name:"OperateSystem"`

	// 生成包支持的操作系统镜像id，若传入OperateSystem字段的值是CentOS7.16，则不需要传入该值；如果不是，则需要通过[DescribeAssetSystems](https://cloud.tencent.com/document/product/1165/49191)接口获取asset支持的操作系统ImageId进行传入
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 标签列表，最大长度50组
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateAssetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAssetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BucketKey")
	delete(f, "AssetName")
	delete(f, "AssetVersion")
	delete(f, "AssetRegion")
	delete(f, "OperateSystem")
	delete(f, "ImageId")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAssetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAssetResponseParams struct {
	// 生成包ID
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// 生成包的全局唯一资源标识符
	AssetArn *string `json:"AssetArn,omitempty" name:"AssetArn"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAssetResponse struct {
	*tchttp.BaseResponse
	Response *CreateAssetResponseParams `json:"Response"`
}

func (r *CreateAssetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAssetWithImageRequestParams struct {
	// 生成包名字，最小长度为1，最大长度为64
	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`

	// 生成包版本，最小长度为1，最大长度为64
	AssetVersion *string `json:"AssetVersion,omitempty" name:"AssetVersion"`

	// 生成包所在地域，详见产品支持的 [地域列表](https://cloud.tencent.com/document/api/1165/42053#.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8)
	AssetRegion *string `json:"AssetRegion,omitempty" name:"AssetRegion"`

	// 生成包支持的操作系统镜像id
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 操作系统镜像包大小，比如：40GB，支持单位 KB、MB、GB
	ImageSize *string `json:"ImageSize,omitempty" name:"ImageSize"`

	// 操作系统镜像包名称，最小长度为1，最大长度为64
	ImageOs *string `json:"ImageOs,omitempty" name:"ImageOs"`

	// 操作系统镜像包类型，CentOS 或者 Windows
	OsType *string `json:"OsType,omitempty" name:"OsType"`

	// 操作系统镜像包类型，当前只支持 SHARED_IMAGE
	ImageType *string `json:"ImageType,omitempty" name:"ImageType"`

	// 操作系统镜像包位数，32 或者 64
	OsBit *uint64 `json:"OsBit,omitempty" name:"OsBit"`
}

type CreateAssetWithImageRequest struct {
	*tchttp.BaseRequest
	
	// 生成包名字，最小长度为1，最大长度为64
	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`

	// 生成包版本，最小长度为1，最大长度为64
	AssetVersion *string `json:"AssetVersion,omitempty" name:"AssetVersion"`

	// 生成包所在地域，详见产品支持的 [地域列表](https://cloud.tencent.com/document/api/1165/42053#.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8)
	AssetRegion *string `json:"AssetRegion,omitempty" name:"AssetRegion"`

	// 生成包支持的操作系统镜像id
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 操作系统镜像包大小，比如：40GB，支持单位 KB、MB、GB
	ImageSize *string `json:"ImageSize,omitempty" name:"ImageSize"`

	// 操作系统镜像包名称，最小长度为1，最大长度为64
	ImageOs *string `json:"ImageOs,omitempty" name:"ImageOs"`

	// 操作系统镜像包类型，CentOS 或者 Windows
	OsType *string `json:"OsType,omitempty" name:"OsType"`

	// 操作系统镜像包类型，当前只支持 SHARED_IMAGE
	ImageType *string `json:"ImageType,omitempty" name:"ImageType"`

	// 操作系统镜像包位数，32 或者 64
	OsBit *uint64 `json:"OsBit,omitempty" name:"OsBit"`
}

func (r *CreateAssetWithImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAssetWithImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetName")
	delete(f, "AssetVersion")
	delete(f, "AssetRegion")
	delete(f, "ImageId")
	delete(f, "ImageSize")
	delete(f, "ImageOs")
	delete(f, "OsType")
	delete(f, "ImageType")
	delete(f, "OsBit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAssetWithImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAssetWithImageResponseParams struct {
	// 生成包ID
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// 生成包的全局唯一资源标识符
	AssetArn *string `json:"AssetArn,omitempty" name:"AssetArn"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAssetWithImageResponse struct {
	*tchttp.BaseResponse
	Response *CreateAssetWithImageResponseParams `json:"Response"`
}

func (r *CreateAssetWithImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAssetWithImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFleetRequestParams struct {
	// 生成包 Id
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// 描述，最小长度0，最大长度100
	Description *string `json:"Description,omitempty" name:"Description"`

	// 网络配置
	InboundPermissions []*InboundPermission `json:"InboundPermissions,omitempty" name:"InboundPermissions"`

	// 服务器类型，参数根据[获取服务器实例类型列表](https://cloud.tencent.com/document/product/1165/48732)接口获取。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 服务器舰队类型，目前只支持ON_DEMAND类型
	FleetType *string `json:"FleetType,omitempty" name:"FleetType"`

	// 服务器舰队名称，最小长度1，最大长度50
	Name *string `json:"Name,omitempty" name:"Name"`

	// 保护策略：不保护NoProtection、完全保护FullProtection、时限保护TimeLimitProtection
	NewGameServerSessionProtectionPolicy *string `json:"NewGameServerSessionProtectionPolicy,omitempty" name:"NewGameServerSessionProtectionPolicy"`

	// VPC 网络 Id，对等连接已不再使用
	PeerVpcId *string `json:"PeerVpcId,omitempty" name:"PeerVpcId"`

	// 资源创建限制策略
	ResourceCreationLimitPolicy *ResourceCreationLimitPolicy `json:"ResourceCreationLimitPolicy,omitempty" name:"ResourceCreationLimitPolicy"`

	// 进程配置
	RuntimeConfiguration *RuntimeConfiguration `json:"RuntimeConfiguration,omitempty" name:"RuntimeConfiguration"`

	// VPC 子网，对等连接已不再使用
	SubNetId *string `json:"SubNetId,omitempty" name:"SubNetId"`

	// 时限保护超时时间，默认60分钟，最小值5，最大值1440；当NewGameSessionProtectionPolicy为TimeLimitProtection时参数有效
	GameServerSessionProtectionTimeLimit *int64 `json:"GameServerSessionProtectionTimeLimit,omitempty" name:"GameServerSessionProtectionTimeLimit"`

	// 标签列表，最大长度50组
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 系统盘，储存类型为 SSD 云硬盘（CLOUD_SSD）时，100-500GB；储存类型为高性能云硬盘（CLOUD_PREMIUM）时，50-500GB；容量以1为单位
	SystemDiskInfo *DiskInfo `json:"SystemDiskInfo,omitempty" name:"SystemDiskInfo"`

	// 数据盘，储存类型为 SSD 云硬盘（CLOUD_SSD）时，100-32000GB；储存类型为高性能云硬盘（CLOUD_PREMIUM）时，10-32000GB；容量以10为单位
	DataDiskInfo []*DiskInfo `json:"DataDiskInfo,omitempty" name:"DataDiskInfo"`

	// 云联网信息，包含对应的账号信息及所属id
	CcnInfos []*CcnInfo `json:"CcnInfos,omitempty" name:"CcnInfos"`

	// fleet公网出带宽最大值，默认100Mbps，范围1-200Mbps
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
}

type CreateFleetRequest struct {
	*tchttp.BaseRequest
	
	// 生成包 Id
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// 描述，最小长度0，最大长度100
	Description *string `json:"Description,omitempty" name:"Description"`

	// 网络配置
	InboundPermissions []*InboundPermission `json:"InboundPermissions,omitempty" name:"InboundPermissions"`

	// 服务器类型，参数根据[获取服务器实例类型列表](https://cloud.tencent.com/document/product/1165/48732)接口获取。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 服务器舰队类型，目前只支持ON_DEMAND类型
	FleetType *string `json:"FleetType,omitempty" name:"FleetType"`

	// 服务器舰队名称，最小长度1，最大长度50
	Name *string `json:"Name,omitempty" name:"Name"`

	// 保护策略：不保护NoProtection、完全保护FullProtection、时限保护TimeLimitProtection
	NewGameServerSessionProtectionPolicy *string `json:"NewGameServerSessionProtectionPolicy,omitempty" name:"NewGameServerSessionProtectionPolicy"`

	// VPC 网络 Id，对等连接已不再使用
	PeerVpcId *string `json:"PeerVpcId,omitempty" name:"PeerVpcId"`

	// 资源创建限制策略
	ResourceCreationLimitPolicy *ResourceCreationLimitPolicy `json:"ResourceCreationLimitPolicy,omitempty" name:"ResourceCreationLimitPolicy"`

	// 进程配置
	RuntimeConfiguration *RuntimeConfiguration `json:"RuntimeConfiguration,omitempty" name:"RuntimeConfiguration"`

	// VPC 子网，对等连接已不再使用
	SubNetId *string `json:"SubNetId,omitempty" name:"SubNetId"`

	// 时限保护超时时间，默认60分钟，最小值5，最大值1440；当NewGameSessionProtectionPolicy为TimeLimitProtection时参数有效
	GameServerSessionProtectionTimeLimit *int64 `json:"GameServerSessionProtectionTimeLimit,omitempty" name:"GameServerSessionProtectionTimeLimit"`

	// 标签列表，最大长度50组
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 系统盘，储存类型为 SSD 云硬盘（CLOUD_SSD）时，100-500GB；储存类型为高性能云硬盘（CLOUD_PREMIUM）时，50-500GB；容量以1为单位
	SystemDiskInfo *DiskInfo `json:"SystemDiskInfo,omitempty" name:"SystemDiskInfo"`

	// 数据盘，储存类型为 SSD 云硬盘（CLOUD_SSD）时，100-32000GB；储存类型为高性能云硬盘（CLOUD_PREMIUM）时，10-32000GB；容量以10为单位
	DataDiskInfo []*DiskInfo `json:"DataDiskInfo,omitempty" name:"DataDiskInfo"`

	// 云联网信息，包含对应的账号信息及所属id
	CcnInfos []*CcnInfo `json:"CcnInfos,omitempty" name:"CcnInfos"`

	// fleet公网出带宽最大值，默认100Mbps，范围1-200Mbps
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
}

func (r *CreateFleetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFleetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetId")
	delete(f, "Description")
	delete(f, "InboundPermissions")
	delete(f, "InstanceType")
	delete(f, "FleetType")
	delete(f, "Name")
	delete(f, "NewGameServerSessionProtectionPolicy")
	delete(f, "PeerVpcId")
	delete(f, "ResourceCreationLimitPolicy")
	delete(f, "RuntimeConfiguration")
	delete(f, "SubNetId")
	delete(f, "GameServerSessionProtectionTimeLimit")
	delete(f, "Tags")
	delete(f, "SystemDiskInfo")
	delete(f, "DataDiskInfo")
	delete(f, "CcnInfos")
	delete(f, "InternetMaxBandwidthOut")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFleetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFleetResponseParams struct {
	// 服务器舰队属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetAttributes *FleetAttributes `json:"FleetAttributes,omitempty" name:"FleetAttributes"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateFleetResponse struct {
	*tchttp.BaseResponse
	Response *CreateFleetResponseParams `json:"Response"`
}

func (r *CreateFleetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFleetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGameServerSessionQueueRequestParams struct {
	// 游戏服务器会话队列名称，长度1~128
	Name *string `json:"Name,omitempty" name:"Name"`

	// 目的服务器舰队（可为别名）列表
	Destinations []*GameServerSessionQueueDestination `json:"Destinations,omitempty" name:"Destinations"`

	// 延迟策略集合
	PlayerLatencyPolicies []*PlayerLatencyPolicy `json:"PlayerLatencyPolicies,omitempty" name:"PlayerLatencyPolicies"`

	// 超时时间（单位秒，默认值为600秒）
	TimeoutInSeconds *uint64 `json:"TimeoutInSeconds,omitempty" name:"TimeoutInSeconds"`

	// 标签列表，最大长度50组
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CreateGameServerSessionQueueRequest struct {
	*tchttp.BaseRequest
	
	// 游戏服务器会话队列名称，长度1~128
	Name *string `json:"Name,omitempty" name:"Name"`

	// 目的服务器舰队（可为别名）列表
	Destinations []*GameServerSessionQueueDestination `json:"Destinations,omitempty" name:"Destinations"`

	// 延迟策略集合
	PlayerLatencyPolicies []*PlayerLatencyPolicy `json:"PlayerLatencyPolicies,omitempty" name:"PlayerLatencyPolicies"`

	// 超时时间（单位秒，默认值为600秒）
	TimeoutInSeconds *uint64 `json:"TimeoutInSeconds,omitempty" name:"TimeoutInSeconds"`

	// 标签列表，最大长度50组
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateGameServerSessionQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGameServerSessionQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Destinations")
	delete(f, "PlayerLatencyPolicies")
	delete(f, "TimeoutInSeconds")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGameServerSessionQueueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGameServerSessionQueueResponseParams struct {
	// 游戏服务器会话队列
	GameServerSessionQueue *GameServerSessionQueue `json:"GameServerSessionQueue,omitempty" name:"GameServerSessionQueue"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateGameServerSessionQueueResponse struct {
	*tchttp.BaseResponse
	Response *CreateGameServerSessionQueueResponseParams `json:"Response"`
}

func (r *CreateGameServerSessionQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGameServerSessionQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGameServerSessionRequestParams struct {
	// 最大玩家数量，最小值不小于0
	MaximumPlayerSessionCount *uint64 `json:"MaximumPlayerSessionCount,omitempty" name:"MaximumPlayerSessionCount"`

	// 别名ID。每个请求需要指定别名ID 或者舰队 ID，如果两个同时指定时，优先选择舰队 ID
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`

	// 创建者ID，最大长度不超过1024个ASCII字符
	CreatorId *string `json:"CreatorId,omitempty" name:"CreatorId"`

	// 舰队ID。每个请求需要指定别名ID 或者舰队 ID，如果两个同时指定时，优先选择舰队 ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 游戏属性，最大长度不超过16组
	GameProperties []*GameProperty `json:"GameProperties,omitempty" name:"GameProperties"`

	// 游戏服务器会话属性详情，最大长度不超过4096个ASCII字符
	GameServerSessionData *string `json:"GameServerSessionData,omitempty" name:"GameServerSessionData"`

	// 游戏服务器会话自定义ID，最大长度不超过4096个ASCII字符
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 幂等token，最大长度不超过48个ASCII字符
	IdempotencyToken *string `json:"IdempotencyToken,omitempty" name:"IdempotencyToken"`

	// 游戏服务器会话名称，最大长度不超过1024个ASCII字符
	Name *string `json:"Name,omitempty" name:"Name"`
}

type CreateGameServerSessionRequest struct {
	*tchttp.BaseRequest
	
	// 最大玩家数量，最小值不小于0
	MaximumPlayerSessionCount *uint64 `json:"MaximumPlayerSessionCount,omitempty" name:"MaximumPlayerSessionCount"`

	// 别名ID。每个请求需要指定别名ID 或者舰队 ID，如果两个同时指定时，优先选择舰队 ID
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`

	// 创建者ID，最大长度不超过1024个ASCII字符
	CreatorId *string `json:"CreatorId,omitempty" name:"CreatorId"`

	// 舰队ID。每个请求需要指定别名ID 或者舰队 ID，如果两个同时指定时，优先选择舰队 ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 游戏属性，最大长度不超过16组
	GameProperties []*GameProperty `json:"GameProperties,omitempty" name:"GameProperties"`

	// 游戏服务器会话属性详情，最大长度不超过4096个ASCII字符
	GameServerSessionData *string `json:"GameServerSessionData,omitempty" name:"GameServerSessionData"`

	// 游戏服务器会话自定义ID，最大长度不超过4096个ASCII字符
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 幂等token，最大长度不超过48个ASCII字符
	IdempotencyToken *string `json:"IdempotencyToken,omitempty" name:"IdempotencyToken"`

	// 游戏服务器会话名称，最大长度不超过1024个ASCII字符
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CreateGameServerSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGameServerSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MaximumPlayerSessionCount")
	delete(f, "AliasId")
	delete(f, "CreatorId")
	delete(f, "FleetId")
	delete(f, "GameProperties")
	delete(f, "GameServerSessionData")
	delete(f, "GameServerSessionId")
	delete(f, "IdempotencyToken")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGameServerSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGameServerSessionResponseParams struct {
	// 游戏服务器会话
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameServerSession *GameServerSession `json:"GameServerSession,omitempty" name:"GameServerSession"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateGameServerSessionResponse struct {
	*tchttp.BaseResponse
	Response *CreateGameServerSessionResponseParams `json:"Response"`
}

func (r *CreateGameServerSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGameServerSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Credentials struct {
	// ssh私钥
	Secret *string `json:"Secret,omitempty" name:"Secret"`

	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

// Predefined struct for user
type DeleteAliasRequestParams struct {
	// 要删除的别名ID
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`
}

type DeleteAliasRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的别名ID
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`
}

func (r *DeleteAliasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAliasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AliasId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAliasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAliasResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAliasResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAliasResponseParams `json:"Response"`
}

func (r *DeleteAliasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAliasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAssetRequestParams struct {
	// 生成包ID
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
}

type DeleteAssetRequest struct {
	*tchttp.BaseRequest
	
	// 生成包ID
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
}

func (r *DeleteAssetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAssetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAssetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAssetResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAssetResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAssetResponseParams `json:"Response"`
}

func (r *DeleteAssetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFleetRequestParams struct {
	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`
}

type DeleteFleetRequest struct {
	*tchttp.BaseRequest
	
	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`
}

func (r *DeleteFleetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFleetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFleetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFleetResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteFleetResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFleetResponseParams `json:"Response"`
}

func (r *DeleteFleetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFleetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGameServerSessionQueueRequestParams struct {
	// 游戏服务器会话队列名字，长度1~128
	Name *string `json:"Name,omitempty" name:"Name"`
}

type DeleteGameServerSessionQueueRequest struct {
	*tchttp.BaseRequest
	
	// 游戏服务器会话队列名字，长度1~128
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteGameServerSessionQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGameServerSessionQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGameServerSessionQueueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGameServerSessionQueueResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteGameServerSessionQueueResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGameServerSessionQueueResponseParams `json:"Response"`
}

func (r *DeleteGameServerSessionQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGameServerSessionQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteScalingPolicyRequestParams struct {
	// 服务器舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 扩缩容策略名称，最小长度为0，最大长度为1024
	Name *string `json:"Name,omitempty" name:"Name"`
}

type DeleteScalingPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 服务器舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 扩缩容策略名称，最小长度为0，最大长度为1024
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteScalingPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScalingPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteScalingPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteScalingPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteScalingPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteScalingPolicyResponseParams `json:"Response"`
}

func (r *DeleteScalingPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScalingPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTimerScalingPolicyRequestParams struct {
	// 定时器ID, 进行encode
	TimerId *string `json:"TimerId,omitempty" name:"TimerId"`

	// 扩缩容配置服务器舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 定时器名称
	TimerName *string `json:"TimerName,omitempty" name:"TimerName"`
}

type DeleteTimerScalingPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 定时器ID, 进行encode
	TimerId *string `json:"TimerId,omitempty" name:"TimerId"`

	// 扩缩容配置服务器舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 定时器名称
	TimerName *string `json:"TimerName,omitempty" name:"TimerName"`
}

func (r *DeleteTimerScalingPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTimerScalingPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimerId")
	delete(f, "FleetId")
	delete(f, "TimerName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTimerScalingPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTimerScalingPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTimerScalingPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTimerScalingPolicyResponseParams `json:"Response"`
}

func (r *DeleteTimerScalingPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTimerScalingPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAliasRequestParams struct {
	// 要检索的队列别名的唯一标识符
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`
}

type DescribeAliasRequest struct {
	*tchttp.BaseRequest
	
	// 要检索的队列别名的唯一标识符
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`
}

func (r *DescribeAliasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAliasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AliasId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAliasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAliasResponseParams struct {
	// 别名对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *Alias `json:"Alias,omitempty" name:"Alias"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAliasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAliasResponseParams `json:"Response"`
}

func (r *DescribeAliasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAliasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetRequestParams struct {
	// 生成包ID
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
}

type DescribeAssetRequest struct {
	*tchttp.BaseRequest
	
	// 生成包ID
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
}

func (r *DescribeAssetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetResponseParams struct {
	// 生成包信息
	Asset *Asset `json:"Asset,omitempty" name:"Asset"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetResponseParams `json:"Response"`
}

func (r *DescribeAssetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetSystemsRequestParams struct {
	// 生成包支持的操作系统类型
	OsType *string `json:"OsType,omitempty" name:"OsType"`

	// 生成包支持的操作系统位数
	OsBit *int64 `json:"OsBit,omitempty" name:"OsBit"`
}

type DescribeAssetSystemsRequest struct {
	*tchttp.BaseRequest
	
	// 生成包支持的操作系统类型
	OsType *string `json:"OsType,omitempty" name:"OsType"`

	// 生成包支持的操作系统位数
	OsBit *int64 `json:"OsBit,omitempty" name:"OsBit"`
}

func (r *DescribeAssetSystemsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetSystemsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OsType")
	delete(f, "OsBit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetSystemsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetSystemsResponseParams struct {
	// 生成包支持的操作系统类型列表
	AssetSupportSys []*AssetSupportSys `json:"AssetSupportSys,omitempty" name:"AssetSupportSys"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetSystemsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetSystemsResponseParams `json:"Response"`
}

func (r *DescribeAssetSystemsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetSystemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetsRequestParams struct {
	// 生成包支持的可部署 [地域列表](https://cloud.tencent.com/document/api/1165/42053#.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8)
	AssetRegion *string `json:"AssetRegion,omitempty" name:"AssetRegion"`

	// 偏移，代表页数，与asset实际数量相关
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 前端界面每页显示的最大条数，不超过100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索条件，支持包ID或包名字过滤，该字段会逐步废弃，建议使用 Filters 字段
	Filter *string `json:"Filter,omitempty" name:"Filter"`

	// 资源过滤字段，可以按照资源名称、资源ID和标签进行过滤- 资源名称过滤    - Key: 固定字符串 "resource:name"    - Values: 资源名称数组（生成包当前仅支持单个名称的过滤）- 资源ID过滤    - Key: 固定字符串 "resource:resourceId"    - Values: 生成包ID数组（生成包当前仅支持单个生成包ID的过滤）- 标签过滤    - 通过标签键过滤        - Key: 固定字符串 "tag:key"        - Values 不传    - 通过标签键值过滤        - Key: 固定字符串 "tag:key-value"        - Values: 标签键值对数组，例如 ["key1:value1", "key1:value2", "key2:value2"]
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeAssetsRequest struct {
	*tchttp.BaseRequest
	
	// 生成包支持的可部署 [地域列表](https://cloud.tencent.com/document/api/1165/42053#.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8)
	AssetRegion *string `json:"AssetRegion,omitempty" name:"AssetRegion"`

	// 偏移，代表页数，与asset实际数量相关
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 前端界面每页显示的最大条数，不超过100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索条件，支持包ID或包名字过滤，该字段会逐步废弃，建议使用 Filters 字段
	Filter *string `json:"Filter,omitempty" name:"Filter"`

	// 资源过滤字段，可以按照资源名称、资源ID和标签进行过滤- 资源名称过滤    - Key: 固定字符串 "resource:name"    - Values: 资源名称数组（生成包当前仅支持单个名称的过滤）- 资源ID过滤    - Key: 固定字符串 "resource:resourceId"    - Values: 生成包ID数组（生成包当前仅支持单个生成包ID的过滤）- 标签过滤    - 通过标签键过滤        - Key: 固定字符串 "tag:key"        - Values 不传    - 通过标签键值过滤        - Key: 固定字符串 "tag:key-value"        - Values: 标签键值对数组，例如 ["key1:value1", "key1:value2", "key2:value2"]
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetRegion")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filter")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetsResponseParams struct {
	// 生成包总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 生成包列表
	Assets []*Asset `json:"Assets,omitempty" name:"Assets"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetsResponseParams `json:"Response"`
}

func (r *DescribeAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcnInstancesRequestParams struct {
	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`
}

type DescribeCcnInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`
}

func (r *DescribeCcnInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcnInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCcnInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcnInstancesResponseParams struct {
	// 云联网实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CcnInstanceSets []*CcnInstanceSets `json:"CcnInstanceSets,omitempty" name:"CcnInstanceSets"`

	// 云联网实例个数，最小值为0
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCcnInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCcnInstancesResponseParams `json:"Response"`
}

func (r *DescribeCcnInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcnInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFleetAttributesRequestParams struct {
	// 服务器舰队 Ids
	FleetIds []*string `json:"FleetIds,omitempty" name:"FleetIds"`

	// 结果返回最大数量，默认值20，最大值100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回结果偏移，最小值0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeFleetAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 服务器舰队 Ids
	FleetIds []*string `json:"FleetIds,omitempty" name:"FleetIds"`

	// 结果返回最大数量，默认值20，最大值100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回结果偏移，最小值0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeFleetAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFleetAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetIds")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFleetAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFleetAttributesResponseParams struct {
	// 服务器舰队属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetAttributes []*FleetAttributes `json:"FleetAttributes,omitempty" name:"FleetAttributes"`

	// 服务器舰队总数，最小值0
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFleetAttributesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFleetAttributesResponseParams `json:"Response"`
}

func (r *DescribeFleetAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFleetAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFleetCapacityRequestParams struct {
	// 服务器舰队ID列表
	FleetIds []*string `json:"FleetIds,omitempty" name:"FleetIds"`

	// 结果返回最大数量，最大值 100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回结果偏移，最小值 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeFleetCapacityRequest struct {
	*tchttp.BaseRequest
	
	// 服务器舰队ID列表
	FleetIds []*string `json:"FleetIds,omitempty" name:"FleetIds"`

	// 结果返回最大数量，最大值 100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回结果偏移，最小值 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeFleetCapacityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFleetCapacityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetIds")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFleetCapacityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFleetCapacityResponseParams struct {
	// 服务器舰队的容量配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetCapacity []*FleetCapacity `json:"FleetCapacity,omitempty" name:"FleetCapacity"`

	// 结果返回最大数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFleetCapacityResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFleetCapacityResponseParams `json:"Response"`
}

func (r *DescribeFleetCapacityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFleetCapacityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFleetEventsRequestParams struct {
	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 分页时返回服务器舰队事件的数量，默认为20，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页时的数据偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 事件代码
	EventCode *string `json:"EventCode,omitempty" name:"EventCode"`

	// 发生事件的开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 发生事件的结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeFleetEventsRequest struct {
	*tchttp.BaseRequest
	
	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 分页时返回服务器舰队事件的数量，默认为20，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页时的数据偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 事件代码
	EventCode *string `json:"EventCode,omitempty" name:"EventCode"`

	// 发生事件的开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 发生事件的结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeFleetEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFleetEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EventCode")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFleetEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFleetEventsResponseParams struct {
	// 返回的事件列表
	Events []*Event `json:"Events,omitempty" name:"Events"`

	// 返回的事件总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFleetEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFleetEventsResponseParams `json:"Response"`
}

func (r *DescribeFleetEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFleetEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFleetPortSettingsRequestParams struct {
	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`
}

type DescribeFleetPortSettingsRequest struct {
	*tchttp.BaseRequest
	
	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`
}

func (r *DescribeFleetPortSettingsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFleetPortSettingsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFleetPortSettingsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFleetPortSettingsResponseParams struct {
	// 安全组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InboundPermissions []*InboundPermission `json:"InboundPermissions,omitempty" name:"InboundPermissions"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFleetPortSettingsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFleetPortSettingsResponseParams `json:"Response"`
}

func (r *DescribeFleetPortSettingsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFleetPortSettingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFleetRelatedResourcesRequestParams struct {
	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`
}

type DescribeFleetRelatedResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`
}

func (r *DescribeFleetRelatedResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFleetRelatedResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFleetRelatedResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFleetRelatedResourcesResponseParams struct {
	// 与服务器舰队关联的资源信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resources []*FleetRelatedResource `json:"Resources,omitempty" name:"Resources"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFleetRelatedResourcesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFleetRelatedResourcesResponseParams `json:"Response"`
}

func (r *DescribeFleetRelatedResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFleetRelatedResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFleetStatisticDetailsRequestParams struct {
	// 服务器舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 查询开始时间，时间格式：YYYY-MM-DD hh:mm:ss
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 查询结束时间，时间格式：YYYY-MM-DD hh:mm:ss
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 结果返回最大数量，最小值0，最大值100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回结果偏移，最小值0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeFleetStatisticDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 服务器舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 查询开始时间，时间格式：YYYY-MM-DD hh:mm:ss
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 查询结束时间，时间格式：YYYY-MM-DD hh:mm:ss
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 结果返回最大数量，最小值0，最大值100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回结果偏移，最小值0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeFleetStatisticDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFleetStatisticDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFleetStatisticDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFleetStatisticDetailsResponseParams struct {
	// 服务部署统计详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetailList []*FleetStatisticDetail `json:"DetailList,omitempty" name:"DetailList"`

	// 记录总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 统计时间类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeType *string `json:"TimeType,omitempty" name:"TimeType"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFleetStatisticDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFleetStatisticDetailsResponseParams `json:"Response"`
}

func (r *DescribeFleetStatisticDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFleetStatisticDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFleetStatisticFlowsRequestParams struct {
	// 服务器舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 查询开始时间，时间格式：YYYY-MM-DD hh:mm:ss
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 查询结束时间，时间格式：YYYY-MM-DD hh:mm:ss
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 结果返回最大数量，最小值0，最大值100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回结果偏移，最小值0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeFleetStatisticFlowsRequest struct {
	*tchttp.BaseRequest
	
	// 服务器舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 查询开始时间，时间格式：YYYY-MM-DD hh:mm:ss
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 查询结束时间，时间格式：YYYY-MM-DD hh:mm:ss
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 结果返回最大数量，最小值0，最大值100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回结果偏移，最小值0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeFleetStatisticFlowsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFleetStatisticFlowsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFleetStatisticFlowsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFleetStatisticFlowsResponseParams struct {
	// 流量统计列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedFlowList []*FleetStatisticFlows `json:"UsedFlowList,omitempty" name:"UsedFlowList"`

	// 时长统计列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedTimeList []*FleetStatisticTimes `json:"UsedTimeList,omitempty" name:"UsedTimeList"`

	// 记录总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 统计时间类型，取值：小时和天
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeType *string `json:"TimeType,omitempty" name:"TimeType"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFleetStatisticFlowsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFleetStatisticFlowsResponseParams `json:"Response"`
}

func (r *DescribeFleetStatisticFlowsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFleetStatisticFlowsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFleetStatisticSummaryRequestParams struct {
	// 服务器舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 查询开始时间，时间格式: YYYY-MM-DD hh:mm:ss
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 查询结束时间，时间格式: YYYY-MM-DD hh:mm:ss
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeFleetStatisticSummaryRequest struct {
	*tchttp.BaseRequest
	
	// 服务器舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 查询开始时间，时间格式: YYYY-MM-DD hh:mm:ss
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 查询结束时间，时间格式: YYYY-MM-DD hh:mm:ss
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeFleetStatisticSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFleetStatisticSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFleetStatisticSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFleetStatisticSummaryResponseParams struct {
	// 总时长，单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalUsedTimeSeconds *string `json:"TotalUsedTimeSeconds,omitempty" name:"TotalUsedTimeSeconds"`

	// 总流量，单位MB
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalUsedFlowMegaBytes *float64 `json:"TotalUsedFlowMegaBytes,omitempty" name:"TotalUsedFlowMegaBytes"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFleetStatisticSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFleetStatisticSummaryResponseParams `json:"Response"`
}

func (r *DescribeFleetStatisticSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFleetStatisticSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFleetUtilizationRequestParams struct {
	// 服务器舰队 Ids
	FleetIds []*string `json:"FleetIds,omitempty" name:"FleetIds"`
}

type DescribeFleetUtilizationRequest struct {
	*tchttp.BaseRequest
	
	// 服务器舰队 Ids
	FleetIds []*string `json:"FleetIds,omitempty" name:"FleetIds"`
}

func (r *DescribeFleetUtilizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFleetUtilizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFleetUtilizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFleetUtilizationResponseParams struct {
	// 服务器舰队利用率
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetUtilization []*FleetUtilization `json:"FleetUtilization,omitempty" name:"FleetUtilization"`

	// 总数，最小值0
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFleetUtilizationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFleetUtilizationResponseParams `json:"Response"`
}

func (r *DescribeFleetUtilizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFleetUtilizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGameServerSessionDetailsRequestParams struct {
	// 别名ID
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`

	// 舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 游戏服务器会话ID，最小长度不小于1个ASCII字符，最大长度不超过48个ASCII字符
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 单次查询记录数上限
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页偏移，用于查询下一页，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// 游戏服务器会话状态(ACTIVE,ACTIVATING,TERMINATED,TERMINATING,ERROR)
	StatusFilter *string `json:"StatusFilter,omitempty" name:"StatusFilter"`
}

type DescribeGameServerSessionDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 别名ID
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`

	// 舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 游戏服务器会话ID，最小长度不小于1个ASCII字符，最大长度不超过48个ASCII字符
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 单次查询记录数上限
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页偏移，用于查询下一页，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// 游戏服务器会话状态(ACTIVE,ACTIVATING,TERMINATED,TERMINATING,ERROR)
	StatusFilter *string `json:"StatusFilter,omitempty" name:"StatusFilter"`
}

func (r *DescribeGameServerSessionDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGameServerSessionDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AliasId")
	delete(f, "FleetId")
	delete(f, "GameServerSessionId")
	delete(f, "Limit")
	delete(f, "NextToken")
	delete(f, "StatusFilter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGameServerSessionDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGameServerSessionDetailsResponseParams struct {
	// 游戏服务器会话详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameServerSessionDetails []*GameServerSessionDetail `json:"GameServerSessionDetails,omitempty" name:"GameServerSessionDetails"`

	// 页偏移，用于查询下一页，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGameServerSessionDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGameServerSessionDetailsResponseParams `json:"Response"`
}

func (r *DescribeGameServerSessionDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGameServerSessionDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGameServerSessionPlacementRequestParams struct {
	// 游戏服务器会话放置的唯一标识符
	PlacementId *string `json:"PlacementId,omitempty" name:"PlacementId"`
}

type DescribeGameServerSessionPlacementRequest struct {
	*tchttp.BaseRequest
	
	// 游戏服务器会话放置的唯一标识符
	PlacementId *string `json:"PlacementId,omitempty" name:"PlacementId"`
}

func (r *DescribeGameServerSessionPlacementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGameServerSessionPlacementRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlacementId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGameServerSessionPlacementRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGameServerSessionPlacementResponseParams struct {
	// 游戏服务器会话放置
	GameServerSessionPlacement *GameServerSessionPlacement `json:"GameServerSessionPlacement,omitempty" name:"GameServerSessionPlacement"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGameServerSessionPlacementResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGameServerSessionPlacementResponseParams `json:"Response"`
}

func (r *DescribeGameServerSessionPlacementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGameServerSessionPlacementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGameServerSessionQueuesRequestParams struct {
	// 游戏服务器会话队列名称数组，单个名字长度1~128
	Names []*string `json:"Names,omitempty" name:"Names"`

	// 结果返回最大数量，最小值0，最大值100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回结果偏移，最小值0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 资源过滤字段，可以按照资源名称、资源ID和标签进行过滤- 资源名称过滤    - Key: 固定字符串 "resource:name"    - Values: 资源名称数组（游戏服务器会话队列支持多个名称的过滤）- 标签过滤    - 通过标签键过滤        - Key: 固定字符串 "tag:key"        - Values 不传    - 通过标签键值过滤        - Key: 固定字符串 "tag:key-value"        - Values: 标签键值对数组，例如 ["key1:value1", "key1:value2", "key2:value2"]
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeGameServerSessionQueuesRequest struct {
	*tchttp.BaseRequest
	
	// 游戏服务器会话队列名称数组，单个名字长度1~128
	Names []*string `json:"Names,omitempty" name:"Names"`

	// 结果返回最大数量，最小值0，最大值100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回结果偏移，最小值0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 资源过滤字段，可以按照资源名称、资源ID和标签进行过滤- 资源名称过滤    - Key: 固定字符串 "resource:name"    - Values: 资源名称数组（游戏服务器会话队列支持多个名称的过滤）- 标签过滤    - 通过标签键过滤        - Key: 固定字符串 "tag:key"        - Values 不传    - 通过标签键值过滤        - Key: 固定字符串 "tag:key-value"        - Values: 标签键值对数组，例如 ["key1:value1", "key1:value2", "key2:value2"]
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeGameServerSessionQueuesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGameServerSessionQueuesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Names")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGameServerSessionQueuesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGameServerSessionQueuesResponseParams struct {
	// 游戏服务器会话队列数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameServerSessionQueues []*GameServerSessionQueue `json:"GameServerSessionQueues,omitempty" name:"GameServerSessionQueues"`

	// 游戏服务器会话队列总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGameServerSessionQueuesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGameServerSessionQueuesResponseParams `json:"Response"`
}

func (r *DescribeGameServerSessionQueuesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGameServerSessionQueuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGameServerSessionsRequestParams struct {
	// 别名ID
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`

	// 舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 游戏服务器会话ID，最小长度不小于1个ASCII字符，最大长度不超过48个ASCII字符
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 单次查询记录数上限
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页偏移，用于查询下一页，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// 游戏服务器会话状态(ACTIVE,ACTIVATING,TERMINATED,TERMINATING,ERROR)
	StatusFilter *string `json:"StatusFilter,omitempty" name:"StatusFilter"`
}

type DescribeGameServerSessionsRequest struct {
	*tchttp.BaseRequest
	
	// 别名ID
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`

	// 舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 游戏服务器会话ID，最小长度不小于1个ASCII字符，最大长度不超过48个ASCII字符
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 单次查询记录数上限
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页偏移，用于查询下一页，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// 游戏服务器会话状态(ACTIVE,ACTIVATING,TERMINATED,TERMINATING,ERROR)
	StatusFilter *string `json:"StatusFilter,omitempty" name:"StatusFilter"`
}

func (r *DescribeGameServerSessionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGameServerSessionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AliasId")
	delete(f, "FleetId")
	delete(f, "GameServerSessionId")
	delete(f, "Limit")
	delete(f, "NextToken")
	delete(f, "StatusFilter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGameServerSessionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGameServerSessionsResponseParams struct {
	// 游戏服务器会话列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameServerSessions []*GameServerSession `json:"GameServerSessions,omitempty" name:"GameServerSessions"`

	// 页偏移，用于查询下一页，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGameServerSessionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGameServerSessionsResponseParams `json:"Response"`
}

func (r *DescribeGameServerSessionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGameServerSessionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceLimitRequestParams struct {

}

type DescribeInstanceLimitRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeInstanceLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceLimitResponseParams struct {
	// 限额
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 详细信息
	ExtraInfos []*ExtraInfos `json:"ExtraInfos,omitempty" name:"ExtraInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceLimitResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceLimitResponseParams `json:"Response"`
}

func (r *DescribeInstanceLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceTypesRequestParams struct {

}

type DescribeInstanceTypesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeInstanceTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceTypesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceTypesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceTypesResponseParams struct {
	// 服务器实例类型列表
	InstanceTypeList []*InstanceTypeInfo `json:"InstanceTypeList,omitempty" name:"InstanceTypeList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceTypesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceTypesResponseParams `json:"Response"`
}

func (r *DescribeInstanceTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesExtendRequestParams struct {
	// 服务器舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 返回结果偏移，最小值0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 结果返回最大数量，最小值0，最大值100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// CVM实例公网IP
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
}

type DescribeInstancesExtendRequest struct {
	*tchttp.BaseRequest
	
	// 服务器舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 返回结果偏移，最小值0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 结果返回最大数量，最小值0，最大值100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// CVM实例公网IP
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
}

func (r *DescribeInstancesExtendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesExtendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "IpAddress")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesExtendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesExtendResponseParams struct {
	// 实例信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Instances []*InstanceExtend `json:"Instances,omitempty" name:"Instances"`

	// 梳理信息总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstancesExtendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesExtendResponseParams `json:"Response"`
}

func (r *DescribeInstancesExtendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesExtendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// 服务器舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// CVM实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 结果返回最大数量，最小值0，最大值100
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回结果偏移，最小值0
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// CVM实例公网IP
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 服务器舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// CVM实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 结果返回最大数量，最小值0，最大值100
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回结果偏移，最小值0
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// CVM实例公网IP
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "IpAddress")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// 实例信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Instances []*Instance `json:"Instances,omitempty" name:"Instances"`

	// 结果返回最大数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesResponseParams `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePlayerSessionsRequestParams struct {
	// 游戏服务器会话ID，最小长度不小于1个ASCII字符，最大长度不超过48个ASCII字符
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 单次查询记录数上限
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页偏移，用于查询下一页，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// 玩家ID，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// 玩家会话ID，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	PlayerSessionId *string `json:"PlayerSessionId,omitempty" name:"PlayerSessionId"`

	// 玩家会话状态（RESERVED,ACTIVE,COMPLETED,TIMEDOUT）
	PlayerSessionStatusFilter *string `json:"PlayerSessionStatusFilter,omitempty" name:"PlayerSessionStatusFilter"`
}

type DescribePlayerSessionsRequest struct {
	*tchttp.BaseRequest
	
	// 游戏服务器会话ID，最小长度不小于1个ASCII字符，最大长度不超过48个ASCII字符
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 单次查询记录数上限
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页偏移，用于查询下一页，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// 玩家ID，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// 玩家会话ID，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	PlayerSessionId *string `json:"PlayerSessionId,omitempty" name:"PlayerSessionId"`

	// 玩家会话状态（RESERVED,ACTIVE,COMPLETED,TIMEDOUT）
	PlayerSessionStatusFilter *string `json:"PlayerSessionStatusFilter,omitempty" name:"PlayerSessionStatusFilter"`
}

func (r *DescribePlayerSessionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePlayerSessionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GameServerSessionId")
	delete(f, "Limit")
	delete(f, "NextToken")
	delete(f, "PlayerId")
	delete(f, "PlayerSessionId")
	delete(f, "PlayerSessionStatusFilter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePlayerSessionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePlayerSessionsResponseParams struct {
	// 玩家会话列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlayerSessions []*PlayerSession `json:"PlayerSessions,omitempty" name:"PlayerSessions"`

	// 页偏移，用于查询下一页，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePlayerSessionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePlayerSessionsResponseParams `json:"Response"`
}

func (r *DescribePlayerSessionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePlayerSessionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuntimeConfigurationRequestParams struct {
	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`
}

type DescribeRuntimeConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`
}

func (r *DescribeRuntimeConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuntimeConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRuntimeConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuntimeConfigurationResponseParams struct {
	// 服务器舰队运行配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuntimeConfiguration *RuntimeConfiguration `json:"RuntimeConfiguration,omitempty" name:"RuntimeConfiguration"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRuntimeConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRuntimeConfigurationResponseParams `json:"Response"`
}

func (r *DescribeRuntimeConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuntimeConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScalingPoliciesRequestParams struct {
	// 服务器舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 状态过滤条件，取值：ACTIVE表示活跃
	StatusFilter *string `json:"StatusFilter,omitempty" name:"StatusFilter"`

	// 返回结果偏移，最小值0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 结果返回最大数量，最小值0，最大值100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeScalingPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 服务器舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 状态过滤条件，取值：ACTIVE表示活跃
	StatusFilter *string `json:"StatusFilter,omitempty" name:"StatusFilter"`

	// 返回结果偏移，最小值0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 结果返回最大数量，最小值0，最大值100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeScalingPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScalingPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	delete(f, "StatusFilter")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScalingPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScalingPoliciesResponseParams struct {
	// 动态扩缩容配置策略数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScalingPolicies []*ScalingPolicy `json:"ScalingPolicies,omitempty" name:"ScalingPolicies"`

	// 动态扩缩容配置策略总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeScalingPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScalingPoliciesResponseParams `json:"Response"`
}

func (r *DescribeScalingPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScalingPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimerScalingPoliciesRequestParams struct {
	// 扩缩容配置服务器舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 定时器名称
	TimerName *string `json:"TimerName,omitempty" name:"TimerName"`

	// 定时器开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 定时器结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeTimerScalingPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 扩缩容配置服务器舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 定时器名称
	TimerName *string `json:"TimerName,omitempty" name:"TimerName"`

	// 定时器开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 定时器结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTimerScalingPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimerScalingPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	delete(f, "TimerName")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTimerScalingPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimerScalingPoliciesResponseParams struct {
	// 定时器扩缩容策略配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimerScalingPolicies []*TimerScalingPolicy `json:"TimerScalingPolicies,omitempty" name:"TimerScalingPolicies"`

	// 定时器总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTimerScalingPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTimerScalingPoliciesResponseParams `json:"Response"`
}

func (r *DescribeTimerScalingPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimerScalingPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserQuotaRequestParams struct {
	// 资源类型
	ResourceType *uint64 `json:"ResourceType,omitempty" name:"ResourceType"`
}

type DescribeUserQuotaRequest struct {
	*tchttp.BaseRequest
	
	// 资源类型
	ResourceType *uint64 `json:"ResourceType,omitempty" name:"ResourceType"`
}

func (r *DescribeUserQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserQuotaResponseParams struct {
	// 配额资源信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuotaResource *QuotaResource `json:"QuotaResource,omitempty" name:"QuotaResource"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUserQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserQuotaResponseParams `json:"Response"`
}

func (r *DescribeUserQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserQuotasRequestParams struct {

}

type DescribeUserQuotasRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeUserQuotasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserQuotasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserQuotasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserQuotasResponseParams struct {
	// 配额信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuotaResource []*QuotaResource `json:"QuotaResource,omitempty" name:"QuotaResource"`

	// 配额信息列表总数，最小值0
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUserQuotasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserQuotasResponseParams `json:"Response"`
}

func (r *DescribeUserQuotasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserQuotasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DesiredPlayerSession struct {
	// 与玩家会话关联的唯一玩家标识
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// 开发人员定义的玩家数据
	PlayerData *string `json:"PlayerData,omitempty" name:"PlayerData"`
}

// Predefined struct for user
type DetachCcnInstancesRequestParams struct {
	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`
}

type DetachCcnInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`
}

func (r *DetachCcnInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachCcnInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachCcnInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachCcnInstancesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetachCcnInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DetachCcnInstancesResponseParams `json:"Response"`
}

func (r *DetachCcnInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachCcnInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskInfo struct {
	// 磁盘类型，支持：高性能云硬盘（CLOUD_PREMIUM）、SSD云硬盘（CLOUD_SSD）
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 系统盘：可选硬盘容量，50-500GB，数字以1为单位，数据盘：可选硬盘容量：10-32000GB，数字以10为单位；当磁盘类型为SSD云硬盘（CLOUD_SSD）最小大小为 100GB
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

// Predefined struct for user
type EndGameServerSessionAndProcessRequestParams struct {
	// 游戏服务器会话ID，如果传入游戏服务器会话ID，结束对应进程以及游戏服务器会话和玩家会话。
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// CVM的公网IP地址，需同时传入IpAddress和Port，结束IpAddress和Port对应的进程以及游戏服务器会话（如果存在）和玩家会话（如果存在），单独传入IpAddress不生效。
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// 端口号，取值范围1025-60000，需同时传入IpAddress和Port，结束IpAddress和Port对应的进程以及游戏服务器会话（如果存在）和玩家会话（如果存在），单独传入Port不生效。
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

type EndGameServerSessionAndProcessRequest struct {
	*tchttp.BaseRequest
	
	// 游戏服务器会话ID，如果传入游戏服务器会话ID，结束对应进程以及游戏服务器会话和玩家会话。
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// CVM的公网IP地址，需同时传入IpAddress和Port，结束IpAddress和Port对应的进程以及游戏服务器会话（如果存在）和玩家会话（如果存在），单独传入IpAddress不生效。
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// 端口号，取值范围1025-60000，需同时传入IpAddress和Port，结束IpAddress和Port对应的进程以及游戏服务器会话（如果存在）和玩家会话（如果存在），单独传入Port不生效。
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

func (r *EndGameServerSessionAndProcessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EndGameServerSessionAndProcessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GameServerSessionId")
	delete(f, "IpAddress")
	delete(f, "Port")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EndGameServerSessionAndProcessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EndGameServerSessionAndProcessResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EndGameServerSessionAndProcessResponse struct {
	*tchttp.BaseResponse
	Response *EndGameServerSessionAndProcessResponseParams `json:"Response"`
}

func (r *EndGameServerSessionAndProcessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EndGameServerSessionAndProcessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Event struct {
	// 事件代码，支持以下的事件代码
	// 
	// - FLEET_CREATED 
	// - FLEET_STATE_DOWNLOADING 
	// - FLEET_BINARY_DOWNLOAD_FAILED 
	// - FLEET_CREATION_EXTRACTING_BUILD
	// - FLEET_CREATION_VALIDATING_RUNTIME_CONFIG
	// - FLEET_STATE_VALIDATING
	// - FLEET_STATE_BUILDING 
	// - FLEET_STATE_ACTIVATING
	// - FLEET_STATE_ACTIVE
	// - FLEET_SCALING_EVENT
	// - FLEET_STATE_ERROR
	// - FLEET_VALIDATION_LAUNCH_PATH_NOT_FOUND
	// - FLEET_ACTIVATION_FAILED_NO_INSTANCES
	// - FLEET_VPC_PEERING_SUCCEEDED
	// - FLEET_VPC_PEERING_FAILED
	// - FLEET_VPC_PEERING_DELETE
	// - FLEET_INITIALIZATION_FAILED
	// - FLEET_DELETED
	// - FLEET_STATE_DELETING
	// - FLEET_ACTIVATION_FAILED
	// - GAME_SESSION_ACTIVATION_TIMEOUT
	EventCode *string `json:"EventCode,omitempty" name:"EventCode"`

	// 事件的唯一标识 ID
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 事件的发生时间，UTC 时间格式
	EventTime *string `json:"EventTime,omitempty" name:"EventTime"`

	// 事件的消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 事件相关的日志存储路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreSignedLogUrl *string `json:"PreSignedLogUrl,omitempty" name:"PreSignedLogUrl"`

	// 事件对应的资源对象唯一标识 ID，例如服务器舰队 ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

type ExtraInfos struct {
	// 实例类型，例如S5.LARGE8
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 实例限额数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalInstances *uint64 `json:"TotalInstances,omitempty" name:"TotalInstances"`
}

type Filter struct {
	// 过滤属性的 key
	Key *string `json:"Key,omitempty" name:"Key"`

	// 过滤属性的 values 值
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type FleetAttributes struct {
	// 生成包 Id
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// 创建服务器舰队时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 服务器舰队资源描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetArn *string `json:"FleetArn,omitempty" name:"FleetArn"`

	// 服务器舰队 Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 服务器舰队类型，目前只支持ON_DEMAND
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetType *string `json:"FleetType,omitempty" name:"FleetType"`

	// 服务器类型，例如S5.LARGE8
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 服务器舰队名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 游戏会话保护策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewGameServerSessionProtectionPolicy *string `json:"NewGameServerSessionProtectionPolicy,omitempty" name:"NewGameServerSessionProtectionPolicy"`

	// 操作系统类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatingSystem *string `json:"OperatingSystem,omitempty" name:"OperatingSystem"`

	// 资源创建限制策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceCreationLimitPolicy *ResourceCreationLimitPolicy `json:"ResourceCreationLimitPolicy,omitempty" name:"ResourceCreationLimitPolicy"`

	// 状态：新建、下载中、验证中、生成中、激活中、活跃、异常、删除中、结束
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 服务器舰队停止状态，为空时表示自动扩缩容
	// 注意：此字段可能返回 null，表示取不到有效值。
	StoppedActions []*string `json:"StoppedActions,omitempty" name:"StoppedActions"`

	// 服务器舰队终止时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TerminationTime *string `json:"TerminationTime,omitempty" name:"TerminationTime"`

	// 时限保护超时时间，默认60分钟，最小值5，最大值1440
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameServerSessionProtectionTimeLimit *uint64 `json:"GameServerSessionProtectionTimeLimit,omitempty" name:"GameServerSessionProtectionTimeLimit"`

	// 计费状态：未开通、已开通、异常、欠费隔离、销毁、解冻
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingStatus *string `json:"BillingStatus,omitempty" name:"BillingStatus"`

	// 标签列表，最大长度50组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 数据盘，储存类型为 SSD 云硬盘（CLOUD_SSD）时，100-32000GB；储存类型为高性能云硬盘（CLOUD_PREMIUM）时，10-32000GB；容量以10为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataDiskInfo []*DiskInfo `json:"DataDiskInfo,omitempty" name:"DataDiskInfo"`

	// 系统盘，储存类型为 SSD 云硬盘（CLOUD_SSD）时，100-500GB；储存类型为高性能云硬盘（CLOUD_PREMIUM）时，50-500GB；容量以1为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	SystemDiskInfo *DiskInfo `json:"SystemDiskInfo,omitempty" name:"SystemDiskInfo"`

	// 云联网相关信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelatedCcnInfos []*RelatedCcnInfo `json:"RelatedCcnInfos,omitempty" name:"RelatedCcnInfos"`

	// fleet公网出带宽最大值，默认100Mbps，范围1-200Mbps
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
}

type FleetCapacity struct {
	// 服务部署 Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 服务器类型，如S3.LARGE8,S2.LARGE8,S5.LARGE8等
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 服务器实例统计数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCounts *InstanceCounts `json:"InstanceCounts,omitempty" name:"InstanceCounts"`

	// 服务器伸缩容间隔，单位分钟，最小值3，最大值30，默认值10
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScalingInterval *uint64 `json:"ScalingInterval,omitempty" name:"ScalingInterval"`
}

type FleetRelatedResource struct {
	// 资源类型。
	// - ALIAS：别名
	// - QUEUE：队列
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 资源ID，目前仅支持别名ID和队列名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源所在区域，如ap-shanghai、na-siliconvalley等
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`
}

type FleetStatisticDetail struct {
	// 舰队ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceIP *string `json:"InstanceIP,omitempty" name:"InstanceIP"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 总时长，单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalUsedTimeSeconds *string `json:"TotalUsedTimeSeconds,omitempty" name:"TotalUsedTimeSeconds"`

	// 总流量，单位MB
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalUsedFlowMegaBytes *float64 `json:"TotalUsedFlowMegaBytes,omitempty" name:"TotalUsedFlowMegaBytes"`
}

type FleetStatisticFlows struct {
	// 总流量，单位MB
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalUsedFlowMegaBytes *float64 `json:"TotalUsedFlowMegaBytes,omitempty" name:"TotalUsedFlowMegaBytes"`

	// 统计开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
}

type FleetStatisticTimes struct {
	// 统计开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 统计总时长，单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalUsedTimeSeconds *string `json:"TotalUsedTimeSeconds,omitempty" name:"TotalUsedTimeSeconds"`
}

type FleetUtilization struct {
	// 游戏会话数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveGameServerSessionCount *uint64 `json:"ActiveGameServerSessionCount,omitempty" name:"ActiveGameServerSessionCount"`

	// 活跃进程数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveServerProcessCount *uint64 `json:"ActiveServerProcessCount,omitempty" name:"ActiveServerProcessCount"`

	// 当前游戏玩家数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentPlayerSessionCount *uint64 `json:"CurrentPlayerSessionCount,omitempty" name:"CurrentPlayerSessionCount"`

	// 服务部署 Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 最大玩家会话数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaximumPlayerSessionCount *uint64 `json:"MaximumPlayerSessionCount,omitempty" name:"MaximumPlayerSessionCount"`
}

type GameProperty struct {
	// 属性名称，最大长度不超过32个ASCII字符
	Key *string `json:"Key,omitempty" name:"Key"`

	// 属性值，最大长度不超过96个ASCII字符
	Value *string `json:"Value,omitempty" name:"Value"`
}

type GameServerSession struct {
	// 游戏服务器会话创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 创建者ID，最大长度不超过1024个ASCII字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorId *string `json:"CreatorId,omitempty" name:"CreatorId"`

	// 当前玩家数量，最小值不小于0
	CurrentPlayerSessionCount *uint64 `json:"CurrentPlayerSessionCount,omitempty" name:"CurrentPlayerSessionCount"`

	// CVM的DNS标识符
	// 注意：此字段可能返回 null，表示取不到有效值。
	DnsName *string `json:"DnsName,omitempty" name:"DnsName"`

	// 舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 游戏属性，最大长度不超过16组
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameProperties []*GameProperty `json:"GameProperties,omitempty" name:"GameProperties"`

	// 游戏服务器会话属性详情，最大长度不超过4096个ASCII字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameServerSessionData *string `json:"GameServerSessionData,omitempty" name:"GameServerSessionData"`

	// 游戏服务器会话ID，最小长度不小于1个ASCII字符，最大长度不超过48个ASCII字符
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// CVM IP地址
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// 对战进程详情，最大长度不超过400000个ASCII字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	MatchmakerData *string `json:"MatchmakerData,omitempty" name:"MatchmakerData"`

	// 最大玩家数量，最小值不小于0
	MaximumPlayerSessionCount *uint64 `json:"MaximumPlayerSessionCount,omitempty" name:"MaximumPlayerSessionCount"`

	// 游戏服务器会话名称，最大长度不超过1024个ASCII字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 玩家会话创建策略（ACCEPT_ALL,DENY_ALL）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlayerSessionCreationPolicy *string `json:"PlayerSessionCreationPolicy,omitempty" name:"PlayerSessionCreationPolicy"`

	// 端口号，最小值不小于1，最大值不超过60000
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 游戏服务器会话状态（ACTIVE,ACTIVATING,TERMINATED,TERMINATING,ERROR）
	Status *string `json:"Status,omitempty" name:"Status"`

	// 游戏服务器会话状态附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusReason *string `json:"StatusReason,omitempty" name:"StatusReason"`

	// 终止的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TerminationTime *string `json:"TerminationTime,omitempty" name:"TerminationTime"`

	// 实例类型，最大长度不超过128个ASCII字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 当前自定义数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentCustomCount *int64 `json:"CurrentCustomCount,omitempty" name:"CurrentCustomCount"`

	// 最大自定义数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxCustomCount *int64 `json:"MaxCustomCount,omitempty" name:"MaxCustomCount"`

	// 权重
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// 会话可用性状态，是否被屏蔽（Enable,Disable）
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailabilityStatus *string `json:"AvailabilityStatus,omitempty" name:"AvailabilityStatus"`
}

type GameServerSessionDetail struct {
	// 游戏服务器会话
	GameServerSession *GameServerSession `json:"GameServerSession,omitempty" name:"GameServerSession"`

	// 保护策略，可选（NoProtection,FullProtection）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtectionPolicy *string `json:"ProtectionPolicy,omitempty" name:"ProtectionPolicy"`
}

type GameServerSessionPlacement struct {
	// 部署Id
	PlacementId *string `json:"PlacementId,omitempty" name:"PlacementId"`

	// 服务部署组名称
	GameServerSessionQueueName *string `json:"GameServerSessionQueueName,omitempty" name:"GameServerSessionQueueName"`

	// 玩家延迟
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlayerLatencies []*PlayerLatency `json:"PlayerLatencies,omitempty" name:"PlayerLatencies"`

	// 服务部署状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 分配给正在运行游戏会话的实例的DNS标识符
	// 注意：此字段可能返回 null，表示取不到有效值。
	DnsName *string `json:"DnsName,omitempty" name:"DnsName"`

	// 游戏会话Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 游戏会话名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameServerSessionName *string `json:"GameServerSessionName,omitempty" name:"GameServerSessionName"`

	// 服务部署区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameServerSessionRegion *string `json:"GameServerSessionRegion,omitempty" name:"GameServerSessionRegion"`

	// 游戏属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameProperties []*GameProperty `json:"GameProperties,omitempty" name:"GameProperties"`

	// 游戏服务器允许同时连接到游戏会话的最大玩家数量，最小值1，最大值为玩家会话最大限额
	MaximumPlayerSessionCount *uint64 `json:"MaximumPlayerSessionCount,omitempty" name:"MaximumPlayerSessionCount"`

	// 游戏会话数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameServerSessionData *string `json:"GameServerSessionData,omitempty" name:"GameServerSessionData"`

	// 运行游戏会话的实例的IP地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// 运行游戏会话的实例的端口号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 游戏匹配数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	MatchmakerData *string `json:"MatchmakerData,omitempty" name:"MatchmakerData"`

	// 部署的玩家游戏数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlacedPlayerSessions []*PlacedPlayerSession `json:"PlacedPlayerSessions,omitempty" name:"PlacedPlayerSessions"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type GameServerSessionQueue struct {
	// 服务部署组名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 服务部署组资源
	GameServerSessionQueueArn *string `json:"GameServerSessionQueueArn,omitempty" name:"GameServerSessionQueueArn"`

	// 目的fleet（可为别名）列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Destinations []*GameServerSessionQueueDestination `json:"Destinations,omitempty" name:"Destinations"`

	// 延迟策略集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlayerLatencyPolicies []*PlayerLatencyPolicy `json:"PlayerLatencyPolicies,omitempty" name:"PlayerLatencyPolicies"`

	// 超时时间
	TimeoutInSeconds *uint64 `json:"TimeoutInSeconds,omitempty" name:"TimeoutInSeconds"`

	// 标签列表，最大长度50组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type GameServerSessionQueueDestination struct {
	// 服务部署组目的的资源描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	DestinationArn *string `json:"DestinationArn,omitempty" name:"DestinationArn"`

	// 服务部署组目的的状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetStatus *string `json:"FleetStatus,omitempty" name:"FleetStatus"`
}

// Predefined struct for user
type GetGameServerInstanceLogUrlRequestParams struct {
	// 游戏舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例IP
	ServerIp *string `json:"ServerIp,omitempty" name:"ServerIp"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每次条数
	Size *uint64 `json:"Size,omitempty" name:"Size"`
}

type GetGameServerInstanceLogUrlRequest struct {
	*tchttp.BaseRequest
	
	// 游戏舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例IP
	ServerIp *string `json:"ServerIp,omitempty" name:"ServerIp"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每次条数
	Size *uint64 `json:"Size,omitempty" name:"Size"`
}

func (r *GetGameServerInstanceLogUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGameServerInstanceLogUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	delete(f, "InstanceId")
	delete(f, "ServerIp")
	delete(f, "Offset")
	delete(f, "Size")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetGameServerInstanceLogUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetGameServerInstanceLogUrlResponseParams struct {
	// 日志下载URL的数组，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	PresignedUrls []*string `json:"PresignedUrls,omitempty" name:"PresignedUrls"`

	// 总条数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 是否还有没拉取完的
	HasNext *bool `json:"HasNext,omitempty" name:"HasNext"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetGameServerInstanceLogUrlResponse struct {
	*tchttp.BaseResponse
	Response *GetGameServerInstanceLogUrlResponseParams `json:"Response"`
}

func (r *GetGameServerInstanceLogUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGameServerInstanceLogUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetGameServerSessionLogUrlRequestParams struct {
	// 游戏服务器会话ID，最小长度不小于1个ASCII字符，最大长度不超过48个ASCII字符
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`
}

type GetGameServerSessionLogUrlRequest struct {
	*tchttp.BaseRequest
	
	// 游戏服务器会话ID，最小长度不小于1个ASCII字符，最大长度不超过48个ASCII字符
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`
}

func (r *GetGameServerSessionLogUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGameServerSessionLogUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GameServerSessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetGameServerSessionLogUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetGameServerSessionLogUrlResponseParams struct {
	// 日志下载URL，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreSignedUrl *string `json:"PreSignedUrl,omitempty" name:"PreSignedUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetGameServerSessionLogUrlResponse struct {
	*tchttp.BaseResponse
	Response *GetGameServerSessionLogUrlResponseParams `json:"Response"`
}

func (r *GetGameServerSessionLogUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGameServerSessionLogUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetInstanceAccessRequestParams struct {
	// 服务器舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type GetInstanceAccessRequest struct {
	*tchttp.BaseRequest
	
	// 服务器舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *GetInstanceAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetInstanceAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetInstanceAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetInstanceAccessResponseParams struct {
	// 实例登录所需要的凭据
	InstanceAccess *InstanceAccess `json:"InstanceAccess,omitempty" name:"InstanceAccess"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetInstanceAccessResponse struct {
	*tchttp.BaseResponse
	Response *GetInstanceAccessResponseParams `json:"Response"`
}

func (r *GetInstanceAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetInstanceAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUploadCredentialsRequestParams struct {
	// 生成包所在地域，详见产品支持的 [地域列表](https://cloud.tencent.com/document/api/1165/42053#.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8)
	AssetRegion *string `json:"AssetRegion,omitempty" name:"AssetRegion"`

	// 生成包的ZIP包名，例如：server.zip
	BucketKey *string `json:"BucketKey,omitempty" name:"BucketKey"`
}

type GetUploadCredentialsRequest struct {
	*tchttp.BaseRequest
	
	// 生成包所在地域，详见产品支持的 [地域列表](https://cloud.tencent.com/document/api/1165/42053#.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8)
	AssetRegion *string `json:"AssetRegion,omitempty" name:"AssetRegion"`

	// 生成包的ZIP包名，例如：server.zip
	BucketKey *string `json:"BucketKey,omitempty" name:"BucketKey"`
}

func (r *GetUploadCredentialsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUploadCredentialsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetRegion")
	delete(f, "BucketKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetUploadCredentialsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUploadCredentialsResponseParams struct {
	// 上传文件授权信息Auth
	BucketAuth *string `json:"BucketAuth,omitempty" name:"BucketAuth"`

	// Bucket名字
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`

	// 生成包所在地域
	AssetRegion *string `json:"AssetRegion,omitempty" name:"AssetRegion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetUploadCredentialsResponse struct {
	*tchttp.BaseResponse
	Response *GetUploadCredentialsResponseParams `json:"Response"`
}

func (r *GetUploadCredentialsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUploadCredentialsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUploadFederationTokenRequestParams struct {

}

type GetUploadFederationTokenRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetUploadFederationTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUploadFederationTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetUploadFederationTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUploadFederationTokenResponseParams struct {
	// 临时证书的过期时间，Unix 时间戳，精确到秒
	ExpiredTime *uint64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// 临时证书
	AssetCredentials *AssetCredentials `json:"AssetCredentials,omitempty" name:"AssetCredentials"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetUploadFederationTokenResponse struct {
	*tchttp.BaseResponse
	Response *GetUploadFederationTokenResponseParams `json:"Response"`
}

func (r *GetUploadFederationTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUploadFederationTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InboundPermission struct {
	// 起始端口号，最小值1025
	FromPort *uint64 `json:"FromPort,omitempty" name:"FromPort"`

	// IP 段范围，合法的 CIDR 地址类型，如所有IPv4来源：0.0.0.0/0
	IpRange *string `json:"IpRange,omitempty" name:"IpRange"`

	// 协议类型：TCP或者UDP
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 终止端口号，最大值60000
	ToPort *uint64 `json:"ToPort,omitempty" name:"ToPort"`
}

type InboundPermissionAuthorization struct {
	// 起始端口号
	FromPort *uint64 `json:"FromPort,omitempty" name:"FromPort"`

	// IP 端范围，CIDR方式划分
	IpRange *string `json:"IpRange,omitempty" name:"IpRange"`

	// 协议类型
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 终止端口号
	ToPort *uint64 `json:"ToPort,omitempty" name:"ToPort"`
}

type InboundPermissionRevocations struct {
	// 起始端口号
	FromPort *uint64 `json:"FromPort,omitempty" name:"FromPort"`

	// IP 端范围，CIDR 方式换分
	IpRange *string `json:"IpRange,omitempty" name:"IpRange"`

	// 协议类型：UDP或者TCP
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 终止端口号
	ToPort *uint64 `json:"ToPort,omitempty" name:"ToPort"`
}

type Instance struct {
	// 服务部署ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// dns
	// 注意：此字段可能返回 null，表示取不到有效值。
	DnsName *string `json:"DnsName,omitempty" name:"DnsName"`

	// 操作系统
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatingSystem *string `json:"OperatingSystem,omitempty" name:"OperatingSystem"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 实例权重
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// 实例是否保留, 1-保留，0-不保留,默认
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReserveValue *int64 `json:"ReserveValue,omitempty" name:"ReserveValue"`

	// 实例的私有IP地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
}

type InstanceAccess struct {
	// 访问实例所需要的凭据
	Credentials *Credentials `json:"Credentials,omitempty" name:"Credentials"`

	// 服务部署Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例公网IP
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// 操作系统
	OperatingSystem *string `json:"OperatingSystem,omitempty" name:"OperatingSystem"`
}

type InstanceCounts struct {
	// 活跃的服务器实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Active *uint64 `json:"Active,omitempty" name:"Active"`

	// 期望的服务器实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desired *uint64 `json:"Desired,omitempty" name:"Desired"`

	// 空闲的服务器实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Idle *uint64 `json:"Idle,omitempty" name:"Idle"`

	// 服务器实例数最大限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxiNum *uint64 `json:"MaxiNum,omitempty" name:"MaxiNum"`

	// 服务器实例数最小限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	MiniNum *uint64 `json:"MiniNum,omitempty" name:"MiniNum"`

	// 已开始创建，但未激活的服务器实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pending *uint64 `json:"Pending,omitempty" name:"Pending"`

	// 结束中的服务器实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Terminating *uint64 `json:"Terminating,omitempty" name:"Terminating"`
}

type InstanceExtend struct {
	// 实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Instance *Instance `json:"Instance,omitempty" name:"Instance"`

	// 实例状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitempty" name:"State"`

	// 健康进程数
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthyProcessCnt *int64 `json:"HealthyProcessCnt,omitempty" name:"HealthyProcessCnt"`

	// 活跃进程数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveProcessCnt *int64 `json:"ActiveProcessCnt,omitempty" name:"ActiveProcessCnt"`

	// 当前游戏会话总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameSessionCnt *int64 `json:"GameSessionCnt,omitempty" name:"GameSessionCnt"`

	// 最大游戏会话数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxGameSessionCnt *int64 `json:"MaxGameSessionCnt,omitempty" name:"MaxGameSessionCnt"`

	// 当前玩家会话数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlayerSessionCnt *int64 `json:"PlayerSessionCnt,omitempty" name:"PlayerSessionCnt"`

	// 最大玩家会话数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxPlayerSessionCnt *int64 `json:"MaxPlayerSessionCnt,omitempty" name:"MaxPlayerSessionCnt"`
}

type InstanceTypeInfo struct {
	// 类型名，例如“标准型SA1”
	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`

	// 类型，例如"SA1.SMALL1"
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// CPU，例如1核就是1
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存，例如2G就是2
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 网络收到包,例如25万PPS就是25
	NetworkCard *uint64 `json:"NetworkCard,omitempty" name:"NetworkCard"`
}

// Predefined struct for user
type JoinGameServerSessionBatchRequestParams struct {
	// 游戏服务器会话ID，最小长度1个ASCII字符，最大长度不超过256个ASCII字符
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 玩家ID列表，最小1组，最大25组
	PlayerIds []*string `json:"PlayerIds,omitempty" name:"PlayerIds"`

	// 玩家自定义数据
	PlayerDataMap *PlayerDataMap `json:"PlayerDataMap,omitempty" name:"PlayerDataMap"`
}

type JoinGameServerSessionBatchRequest struct {
	*tchttp.BaseRequest
	
	// 游戏服务器会话ID，最小长度1个ASCII字符，最大长度不超过256个ASCII字符
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 玩家ID列表，最小1组，最大25组
	PlayerIds []*string `json:"PlayerIds,omitempty" name:"PlayerIds"`

	// 玩家自定义数据
	PlayerDataMap *PlayerDataMap `json:"PlayerDataMap,omitempty" name:"PlayerDataMap"`
}

func (r *JoinGameServerSessionBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *JoinGameServerSessionBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GameServerSessionId")
	delete(f, "PlayerIds")
	delete(f, "PlayerDataMap")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "JoinGameServerSessionBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type JoinGameServerSessionBatchResponseParams struct {
	// 玩家会话列表，最大25组
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlayerSessions []*PlayerSession `json:"PlayerSessions,omitempty" name:"PlayerSessions"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type JoinGameServerSessionBatchResponse struct {
	*tchttp.BaseResponse
	Response *JoinGameServerSessionBatchResponseParams `json:"Response"`
}

func (r *JoinGameServerSessionBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *JoinGameServerSessionBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type JoinGameServerSessionRequestParams struct {
	// 游戏服务器会话ID，最小长度1个ASCII字符，最大长度不超过256个ASCII字符
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 玩家ID，最大长度1024个ASCII字符
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// 玩家自定义数据，最大长度2048个ASCII字符
	PlayerData *string `json:"PlayerData,omitempty" name:"PlayerData"`
}

type JoinGameServerSessionRequest struct {
	*tchttp.BaseRequest
	
	// 游戏服务器会话ID，最小长度1个ASCII字符，最大长度不超过256个ASCII字符
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 玩家ID，最大长度1024个ASCII字符
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// 玩家自定义数据，最大长度2048个ASCII字符
	PlayerData *string `json:"PlayerData,omitempty" name:"PlayerData"`
}

func (r *JoinGameServerSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *JoinGameServerSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GameServerSessionId")
	delete(f, "PlayerId")
	delete(f, "PlayerData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "JoinGameServerSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type JoinGameServerSessionResponseParams struct {
	// 玩家会话
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlayerSession *PlayerSession `json:"PlayerSession,omitempty" name:"PlayerSession"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type JoinGameServerSessionResponse struct {
	*tchttp.BaseResponse
	Response *JoinGameServerSessionResponseParams `json:"Response"`
}

func (r *JoinGameServerSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *JoinGameServerSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAliasesRequestParams struct {
	// 名字，长度不小于1字符不超过1024字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 路由策略类型，有效值常规别名(SIMPLE)、终止别名(TERMINAL)
	RoutingStrategyType *string `json:"RoutingStrategyType,omitempty" name:"RoutingStrategyType"`

	// 要返回的最大结果数，最小值1
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移，默认0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段，例如CreationTime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，有效值asc|desc
	OrderWay *string `json:"OrderWay,omitempty" name:"OrderWay"`

	// 资源过滤字段，可以按照资源名称和标签进行过滤- 资源名称过滤    - Key: 固定字符串 "resource:name"    - Values: 资源名称数组（舰队当前仅支持单个名称的过滤）- 标签过滤    - 通过标签键过滤        - Key: 固定字符串 "tag:key"        - Values 不传    - 通过标签键值过滤        - Key: 固定字符串 "tag:key-value"        - Values: 标签键值对数组，例如 ["key1:value1", "key1:value2", "key2:value2"]
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type ListAliasesRequest struct {
	*tchttp.BaseRequest
	
	// 名字，长度不小于1字符不超过1024字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 路由策略类型，有效值常规别名(SIMPLE)、终止别名(TERMINAL)
	RoutingStrategyType *string `json:"RoutingStrategyType,omitempty" name:"RoutingStrategyType"`

	// 要返回的最大结果数，最小值1
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移，默认0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段，例如CreationTime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，有效值asc|desc
	OrderWay *string `json:"OrderWay,omitempty" name:"OrderWay"`

	// 资源过滤字段，可以按照资源名称和标签进行过滤- 资源名称过滤    - Key: 固定字符串 "resource:name"    - Values: 资源名称数组（舰队当前仅支持单个名称的过滤）- 标签过滤    - 通过标签键过滤        - Key: 固定字符串 "tag:key"        - Values 不传    - 通过标签键值过滤        - Key: 固定字符串 "tag:key-value"        - Values: 标签键值对数组，例如 ["key1:value1", "key1:value2", "key2:value2"]
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ListAliasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAliasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "RoutingStrategyType")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderWay")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAliasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAliasesResponseParams struct {
	// 别名对象数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Aliases []*Alias `json:"Aliases,omitempty" name:"Aliases"`

	// 总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListAliasesResponse struct {
	*tchttp.BaseResponse
	Response *ListAliasesResponseParams `json:"Response"`
}

func (r *ListAliasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAliasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListFleetsRequestParams struct {
	// 生成包 Id
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// 结果返回最大值，暂未使用
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 结果返回偏移，暂未使用
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 资源过滤字段，可以按照资源名称和标签进行过滤- 资源名称过滤    - Key: 固定字符串 "resource:name"    - Values: 资源名称数组（当前仅支持单个名称的过滤）- 标签过滤    - 通过标签键过滤        - Key: 固定字符串 "tag:key"        - Values 不传    - 通过标签键值过滤        - Key: 固定字符串 "tag:key-value"        - Values: 标签键值对数组，例如 ["key1:value1", "key1:value2", "key2:value2"]
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type ListFleetsRequest struct {
	*tchttp.BaseRequest
	
	// 生成包 Id
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// 结果返回最大值，暂未使用
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 结果返回偏移，暂未使用
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 资源过滤字段，可以按照资源名称和标签进行过滤- 资源名称过滤    - Key: 固定字符串 "resource:name"    - Values: 资源名称数组（当前仅支持单个名称的过滤）- 标签过滤    - 通过标签键过滤        - Key: 固定字符串 "tag:key"        - Values 不传    - 通过标签键值过滤        - Key: 固定字符串 "tag:key-value"        - Values: 标签键值对数组，例如 ["key1:value1", "key1:value2", "key2:value2"]
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ListFleetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListFleetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListFleetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListFleetsResponseParams struct {
	// 服务器舰队 Id 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetIds []*string `json:"FleetIds,omitempty" name:"FleetIds"`

	// 服务器舰队 Id 总数，最小值0
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListFleetsResponse struct {
	*tchttp.BaseResponse
	Response *ListFleetsResponseParams `json:"Response"`
}

func (r *ListFleetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListFleetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PlacedPlayerSession struct {
	// 玩家Id
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// 玩家会话Id
	PlayerSessionId *string `json:"PlayerSessionId,omitempty" name:"PlayerSessionId"`
}

type PlayerDataMap struct {
	// 玩家自定义数据键，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	Key *string `json:"Key,omitempty" name:"Key"`

	// 玩家自定义数据值，最小长度不小于1个ASCII字符，最大长度不超过2048个ASCII字符
	Value *string `json:"Value,omitempty" name:"Value"`
}

type PlayerLatency struct {
	// 玩家Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// 延迟对应的区域名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionIdentifier *string `json:"RegionIdentifier,omitempty" name:"RegionIdentifier"`

	// 毫秒级延迟
	LatencyInMilliseconds *uint64 `json:"LatencyInMilliseconds,omitempty" name:"LatencyInMilliseconds"`
}

type PlayerLatencyPolicy struct {
	// 任意player允许的最大延迟，单位：毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaximumIndividualPlayerLatencyMilliseconds *uint64 `json:"MaximumIndividualPlayerLatencyMilliseconds,omitempty" name:"MaximumIndividualPlayerLatencyMilliseconds"`

	// 放置新GameServerSession时强制实施策略的时间长度，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyDurationSeconds *uint64 `json:"PolicyDurationSeconds,omitempty" name:"PolicyDurationSeconds"`
}

type PlayerSession struct {
	// 玩家会话创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 游戏服务器会话运行的DNS标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	DnsName *string `json:"DnsName,omitempty" name:"DnsName"`

	// 舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 游戏服务器会话ID，最小长度1个ASCII字符，最大长度不超过256个ASCII字符
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 游戏服务器会话运行的CVM地址
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// 玩家自定义数据，最大长度2048个ASCII字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlayerData *string `json:"PlayerData,omitempty" name:"PlayerData"`

	// 玩家ID，最大长度1024个ASCII字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// 玩家会话ID
	PlayerSessionId *string `json:"PlayerSessionId,omitempty" name:"PlayerSessionId"`

	// 端口号，最小值不小于1，最大值不超过60000
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 玩家会话的状态（RESERVED = 1,ACTIVE = 2,COMPLETED = 3,TIMEDOUT = 4）
	Status *string `json:"Status,omitempty" name:"Status"`

	// 玩家会话终止时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TerminationTime *string `json:"TerminationTime,omitempty" name:"TerminationTime"`
}

// Predefined struct for user
type PutScalingPolicyRequestParams struct {
	// 扩缩容配置服务器舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 扩缩容策略名称，最小长度为1，最大长度为1024
	Name *string `json:"Name,omitempty" name:"Name"`

	// 扩缩容调整值，ScalingAdjustmentType取值PercentChangeInCapacity时，取值范围-99~99
	// ScalingAdjustmentType取值ChangeInCapacity或ExactCapacity时，最小值要缩容的最多CVM个数，最大值为实际最大的CVM个数限额
	ScalingAdjustment *int64 `json:"ScalingAdjustment,omitempty" name:"ScalingAdjustment"`

	// 扩缩容调整类型，取值（ChangeInCapacity，ExactCapacity，PercentChangeInCapacity）
	ScalingAdjustmentType *string `json:"ScalingAdjustmentType,omitempty" name:"ScalingAdjustmentType"`

	// 扩缩容指标阈值
	Threshold *float64 `json:"Threshold,omitempty" name:"Threshold"`

	// 扩缩容策略比较符，取值：>,>=,<,<=
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" name:"ComparisonOperator"`

	// 单个策略持续时间长度（分钟）
	EvaluationPeriods *int64 `json:"EvaluationPeriods,omitempty" name:"EvaluationPeriods"`

	// 扩缩容参与计算的指标名称，PolicyType取值RuleBased，
	// MetricName取值（AvailableGameServerSessions，AvailableCustomCount，PercentAvailableCustomCount，ActiveInstances，IdleInstances，CurrentPlayerSessions和PercentIdleInstances）；
	// PolicyType取值TargetBased时，MetricName取值PercentAvailableGameSessions
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 策略类型，取值：TargetBased表示基于目标的策略；RuleBased表示基于规则的策略
	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`

	// 扩缩容目标值配置，只有TargetBased类型的策略生效
	TargetConfiguration *TargetConfiguration `json:"TargetConfiguration,omitempty" name:"TargetConfiguration"`
}

type PutScalingPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 扩缩容配置服务器舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 扩缩容策略名称，最小长度为1，最大长度为1024
	Name *string `json:"Name,omitempty" name:"Name"`

	// 扩缩容调整值，ScalingAdjustmentType取值PercentChangeInCapacity时，取值范围-99~99
	// ScalingAdjustmentType取值ChangeInCapacity或ExactCapacity时，最小值要缩容的最多CVM个数，最大值为实际最大的CVM个数限额
	ScalingAdjustment *int64 `json:"ScalingAdjustment,omitempty" name:"ScalingAdjustment"`

	// 扩缩容调整类型，取值（ChangeInCapacity，ExactCapacity，PercentChangeInCapacity）
	ScalingAdjustmentType *string `json:"ScalingAdjustmentType,omitempty" name:"ScalingAdjustmentType"`

	// 扩缩容指标阈值
	Threshold *float64 `json:"Threshold,omitempty" name:"Threshold"`

	// 扩缩容策略比较符，取值：>,>=,<,<=
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" name:"ComparisonOperator"`

	// 单个策略持续时间长度（分钟）
	EvaluationPeriods *int64 `json:"EvaluationPeriods,omitempty" name:"EvaluationPeriods"`

	// 扩缩容参与计算的指标名称，PolicyType取值RuleBased，
	// MetricName取值（AvailableGameServerSessions，AvailableCustomCount，PercentAvailableCustomCount，ActiveInstances，IdleInstances，CurrentPlayerSessions和PercentIdleInstances）；
	// PolicyType取值TargetBased时，MetricName取值PercentAvailableGameSessions
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 策略类型，取值：TargetBased表示基于目标的策略；RuleBased表示基于规则的策略
	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`

	// 扩缩容目标值配置，只有TargetBased类型的策略生效
	TargetConfiguration *TargetConfiguration `json:"TargetConfiguration,omitempty" name:"TargetConfiguration"`
}

func (r *PutScalingPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutScalingPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	delete(f, "Name")
	delete(f, "ScalingAdjustment")
	delete(f, "ScalingAdjustmentType")
	delete(f, "Threshold")
	delete(f, "ComparisonOperator")
	delete(f, "EvaluationPeriods")
	delete(f, "MetricName")
	delete(f, "PolicyType")
	delete(f, "TargetConfiguration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PutScalingPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PutScalingPolicyResponseParams struct {
	// 规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type PutScalingPolicyResponse struct {
	*tchttp.BaseResponse
	Response *PutScalingPolicyResponseParams `json:"Response"`
}

func (r *PutScalingPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutScalingPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PutTimerScalingPolicyRequestParams struct {
	// 定时器策略消息
	TimerScalingPolicy *TimerScalingPolicy `json:"TimerScalingPolicy,omitempty" name:"TimerScalingPolicy"`
}

type PutTimerScalingPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 定时器策略消息
	TimerScalingPolicy *TimerScalingPolicy `json:"TimerScalingPolicy,omitempty" name:"TimerScalingPolicy"`
}

func (r *PutTimerScalingPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutTimerScalingPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimerScalingPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PutTimerScalingPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PutTimerScalingPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type PutTimerScalingPolicyResponse struct {
	*tchttp.BaseResponse
	Response *PutTimerScalingPolicyResponseParams `json:"Response"`
}

func (r *PutTimerScalingPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutTimerScalingPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuotaResource struct {
	// 资源类型，1生成包、2服务部署、3别名、4游戏服务器队列、5实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType *uint64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// 总额度
	// 注意：此字段可能返回 null，表示取不到有效值。
	HardLimit *uint64 `json:"HardLimit,omitempty" name:"HardLimit"`

	// 剩余额度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remaining *uint64 `json:"Remaining,omitempty" name:"Remaining"`

	// 额外信息，可能为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraInfo *string `json:"ExtraInfo,omitempty" name:"ExtraInfo"`
}

type RelatedCcnInfo struct {
	// 云联网所属账号
	AccountId *string `json:"AccountId,omitempty" name:"AccountId"`

	// 云联网 ID
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// 关联云联网状态
	AttachType *string `json:"AttachType,omitempty" name:"AttachType"`
}

// Predefined struct for user
type ResolveAliasRequestParams struct {
	// 要获取fleetId的别名ID
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`
}

type ResolveAliasRequest struct {
	*tchttp.BaseRequest
	
	// 要获取fleetId的别名ID
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`
}

func (r *ResolveAliasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResolveAliasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AliasId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResolveAliasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResolveAliasResponseParams struct {
	// 别名指向的fleet的唯一标识符
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResolveAliasResponse struct {
	*tchttp.BaseResponse
	Response *ResolveAliasResponseParams `json:"Response"`
}

func (r *ResolveAliasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResolveAliasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceCreationLimitPolicy struct {
	// 创建数量，最小值1，默认2
	NewGameServerSessionsPerCreator *uint64 `json:"NewGameServerSessionsPerCreator,omitempty" name:"NewGameServerSessionsPerCreator"`

	// 单位时间，最小值1，默认3，单位分钟
	PolicyPeriodInMinutes *uint64 `json:"PolicyPeriodInMinutes,omitempty" name:"PolicyPeriodInMinutes"`
}

type RoutingStrategy struct {
	// 别名的路由策略的类型，有效值常规别名(SIMPLE)、终止别名(TERMINAL)
	Type *string `json:"Type,omitempty" name:"Type"`

	// 别名指向的队列的唯一标识符
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 与终端路由策略一起使用的消息文本，长度不小于1字符不超过1024字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`
}

type RuntimeConfiguration struct {
	// 游戏会话进程超时，最小值1，最大值600，单位秒
	GameServerSessionActivationTimeoutSeconds *uint64 `json:"GameServerSessionActivationTimeoutSeconds,omitempty" name:"GameServerSessionActivationTimeoutSeconds"`

	// 最大游戏会话数，最小值1，最大值2147483647
	MaxConcurrentGameServerSessionActivations *uint64 `json:"MaxConcurrentGameServerSessionActivations,omitempty" name:"MaxConcurrentGameServerSessionActivations"`

	// 服务进程配置，至少有一个进程配置
	ServerProcesses []*ServerProcesse `json:"ServerProcesses,omitempty" name:"ServerProcesses"`
}

type ScalingPolicy struct {
	// 服务部署ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 保留参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScalingAdjustment *string `json:"ScalingAdjustment,omitempty" name:"ScalingAdjustment"`

	// 保留参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScalingAdjustmentType *string `json:"ScalingAdjustmentType,omitempty" name:"ScalingAdjustmentType"`

	// 保留参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" name:"ComparisonOperator"`

	// 保留参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Threshold *string `json:"Threshold,omitempty" name:"Threshold"`

	// 保留参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	EvaluationPeriods *string `json:"EvaluationPeriods,omitempty" name:"EvaluationPeriods"`

	// 保留参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 策略类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`

	// 基于规则的配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetConfiguration *TargetConfiguration `json:"TargetConfiguration,omitempty" name:"TargetConfiguration"`
}

// Predefined struct for user
type SearchGameServerSessionsRequestParams struct {
	// 别名ID
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`

	// 舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 单次查询记录数上限
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页偏移，用于查询下一页，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// 搜索条件表达式，支持如下变量
	// gameServerSessionName 游戏会话名称 String
	// gameServerSessionId 游戏会话ID String
	// maximumSessions 最大的玩家会话数 Number
	// creationTimeMillis 创建时间，单位：毫秒 Number
	// playerSessionCount 当前玩家会话数 Number
	// hasAvailablePlayerSessions 是否有可用玩家数 String 取值true或false
	// gameServerSessionProperties 游戏会话属性 String
	// 
	// 表达式String类型 等于=，不等于<>判断
	// 表示Number类型支持 =,<>,>,>=,<,<=
	// 
	// 例如：
	// FilterExpression取值
	// playerSessionCount>=2 AND hasAvailablePlayerSessions=true"
	// 表示查找至少有两个玩家，而且有可用玩家会话的游戏会话。
	// FilterExpression取值
	// gameServerSessionProperties.K1 = 'V1' AND gameServerSessionProperties.K2 = 'V2' OR gameServerSessionProperties.K3 = 'V3'
	// 
	// 表示
	// 查询满足如下游戏服务器会话属性的游戏会话
	// {
	//     "GameProperties":[
	//         {
	//             "Key":"K1",
	//             "Value":"V1"
	//         },
	//         {
	//             "Key":"K2",
	//             "Value":"V2"
	//         },
	//         {
	//             "Key":"K3",
	//             "Value":"V3"
	//         }
	//     ]
	// }
	FilterExpression *string `json:"FilterExpression,omitempty" name:"FilterExpression"`

	// 排序条件关键字
	// 支持排序字段
	// gameServerSessionName 游戏会话名称 String
	// gameServerSessionId 游戏会话ID String
	// maximumSessions 最大的玩家会话数 Number
	// creationTimeMillis 创建时间，单位：毫秒 Number
	// playerSessionCount 当前玩家会话数 Number
	SortExpression *string `json:"SortExpression,omitempty" name:"SortExpression"`
}

type SearchGameServerSessionsRequest struct {
	*tchttp.BaseRequest
	
	// 别名ID
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`

	// 舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 单次查询记录数上限
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页偏移，用于查询下一页，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// 搜索条件表达式，支持如下变量
	// gameServerSessionName 游戏会话名称 String
	// gameServerSessionId 游戏会话ID String
	// maximumSessions 最大的玩家会话数 Number
	// creationTimeMillis 创建时间，单位：毫秒 Number
	// playerSessionCount 当前玩家会话数 Number
	// hasAvailablePlayerSessions 是否有可用玩家数 String 取值true或false
	// gameServerSessionProperties 游戏会话属性 String
	// 
	// 表达式String类型 等于=，不等于<>判断
	// 表示Number类型支持 =,<>,>,>=,<,<=
	// 
	// 例如：
	// FilterExpression取值
	// playerSessionCount>=2 AND hasAvailablePlayerSessions=true"
	// 表示查找至少有两个玩家，而且有可用玩家会话的游戏会话。
	// FilterExpression取值
	// gameServerSessionProperties.K1 = 'V1' AND gameServerSessionProperties.K2 = 'V2' OR gameServerSessionProperties.K3 = 'V3'
	// 
	// 表示
	// 查询满足如下游戏服务器会话属性的游戏会话
	// {
	//     "GameProperties":[
	//         {
	//             "Key":"K1",
	//             "Value":"V1"
	//         },
	//         {
	//             "Key":"K2",
	//             "Value":"V2"
	//         },
	//         {
	//             "Key":"K3",
	//             "Value":"V3"
	//         }
	//     ]
	// }
	FilterExpression *string `json:"FilterExpression,omitempty" name:"FilterExpression"`

	// 排序条件关键字
	// 支持排序字段
	// gameServerSessionName 游戏会话名称 String
	// gameServerSessionId 游戏会话ID String
	// maximumSessions 最大的玩家会话数 Number
	// creationTimeMillis 创建时间，单位：毫秒 Number
	// playerSessionCount 当前玩家会话数 Number
	SortExpression *string `json:"SortExpression,omitempty" name:"SortExpression"`
}

func (r *SearchGameServerSessionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchGameServerSessionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AliasId")
	delete(f, "FleetId")
	delete(f, "Limit")
	delete(f, "NextToken")
	delete(f, "FilterExpression")
	delete(f, "SortExpression")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchGameServerSessionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchGameServerSessionsResponseParams struct {
	// 游戏服务器会话列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameServerSessions []*GameServerSession `json:"GameServerSessions,omitempty" name:"GameServerSessions"`

	// 页偏移，用于查询下一页，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SearchGameServerSessionsResponse struct {
	*tchttp.BaseResponse
	Response *SearchGameServerSessionsResponseParams `json:"Response"`
}

func (r *SearchGameServerSessionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchGameServerSessionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerProcesse struct {
	// 并发执行数量，所有进程并发执行总数最小值1，最大值50
	ConcurrentExecutions *uint64 `json:"ConcurrentExecutions,omitempty" name:"ConcurrentExecutions"`

	// 启动路径：Linux路径/local/game/ 或WIndows路径C:\game\，最小长度1，最大长度1024
	LaunchPath *string `json:"LaunchPath,omitempty" name:"LaunchPath"`

	// 启动参数，最小长度0，最大长度1024
	Parameters *string `json:"Parameters,omitempty" name:"Parameters"`
}

// Predefined struct for user
type SetServerReservedRequestParams struct {
	// 扩缩容配置服务器舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例是否保留, 1-保留，0-不保留,默认
	ReserveValue *int64 `json:"ReserveValue,omitempty" name:"ReserveValue"`
}

type SetServerReservedRequest struct {
	*tchttp.BaseRequest
	
	// 扩缩容配置服务器舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例是否保留, 1-保留，0-不保留,默认
	ReserveValue *int64 `json:"ReserveValue,omitempty" name:"ReserveValue"`
}

func (r *SetServerReservedRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetServerReservedRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	delete(f, "InstanceId")
	delete(f, "ReserveValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetServerReservedRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetServerReservedResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetServerReservedResponse struct {
	*tchttp.BaseResponse
	Response *SetServerReservedResponseParams `json:"Response"`
}

func (r *SetServerReservedResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetServerReservedResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetServerWeightRequestParams struct {
	// 服务器舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 权重，最小值0，最大值10，默认值5
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

type SetServerWeightRequest struct {
	*tchttp.BaseRequest
	
	// 服务器舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 权重，最小值0，最大值10，默认值5
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

func (r *SetServerWeightRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetServerWeightRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	delete(f, "InstanceId")
	delete(f, "Weight")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetServerWeightRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetServerWeightResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetServerWeightResponse struct {
	*tchttp.BaseResponse
	Response *SetServerWeightResponseParams `json:"Response"`
}

func (r *SetServerWeightResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetServerWeightResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartFleetActionsRequestParams struct {
	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 服务器舰队扩展策略，值为["AUTO_SCALING"]
	Actions []*string `json:"Actions,omitempty" name:"Actions"`
}

type StartFleetActionsRequest struct {
	*tchttp.BaseRequest
	
	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 服务器舰队扩展策略，值为["AUTO_SCALING"]
	Actions []*string `json:"Actions,omitempty" name:"Actions"`
}

func (r *StartFleetActionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartFleetActionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	delete(f, "Actions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartFleetActionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartFleetActionsResponseParams struct {
	// 服务器舰队 Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StartFleetActionsResponse struct {
	*tchttp.BaseResponse
	Response *StartFleetActionsResponseParams `json:"Response"`
}

func (r *StartFleetActionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartFleetActionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartGameServerSessionPlacementRequestParams struct {
	// 开始部署游戏服务器会话的唯一标识符，最大值48个ASCII字符，模式：[a-zA-Z0-9-]+
	PlacementId *string `json:"PlacementId,omitempty" name:"PlacementId"`

	// 游戏服务器会话队列名称
	GameServerSessionQueueName *string `json:"GameServerSessionQueueName,omitempty" name:"GameServerSessionQueueName"`

	// 游戏服务器允许同时连接到游戏会话的最大玩家数量，最小值1，最大值为玩家会话最大限额
	MaximumPlayerSessionCount *uint64 `json:"MaximumPlayerSessionCount,omitempty" name:"MaximumPlayerSessionCount"`

	// 玩家游戏会话信息
	DesiredPlayerSessions []*DesiredPlayerSession `json:"DesiredPlayerSessions,omitempty" name:"DesiredPlayerSessions"`

	// 玩家游戏会话属性
	GameProperties []*GameProperty `json:"GameProperties,omitempty" name:"GameProperties"`

	// 游戏服务器会话数据，最大长度不超过4096个ASCII字符
	GameServerSessionData *string `json:"GameServerSessionData,omitempty" name:"GameServerSessionData"`

	// 游戏服务器会话名称，最大长度不超过4096个ASCII字符
	GameServerSessionName *string `json:"GameServerSessionName,omitempty" name:"GameServerSessionName"`

	// 玩家延迟
	PlayerLatencies []*PlayerLatency `json:"PlayerLatencies,omitempty" name:"PlayerLatencies"`
}

type StartGameServerSessionPlacementRequest struct {
	*tchttp.BaseRequest
	
	// 开始部署游戏服务器会话的唯一标识符，最大值48个ASCII字符，模式：[a-zA-Z0-9-]+
	PlacementId *string `json:"PlacementId,omitempty" name:"PlacementId"`

	// 游戏服务器会话队列名称
	GameServerSessionQueueName *string `json:"GameServerSessionQueueName,omitempty" name:"GameServerSessionQueueName"`

	// 游戏服务器允许同时连接到游戏会话的最大玩家数量，最小值1，最大值为玩家会话最大限额
	MaximumPlayerSessionCount *uint64 `json:"MaximumPlayerSessionCount,omitempty" name:"MaximumPlayerSessionCount"`

	// 玩家游戏会话信息
	DesiredPlayerSessions []*DesiredPlayerSession `json:"DesiredPlayerSessions,omitempty" name:"DesiredPlayerSessions"`

	// 玩家游戏会话属性
	GameProperties []*GameProperty `json:"GameProperties,omitempty" name:"GameProperties"`

	// 游戏服务器会话数据，最大长度不超过4096个ASCII字符
	GameServerSessionData *string `json:"GameServerSessionData,omitempty" name:"GameServerSessionData"`

	// 游戏服务器会话名称，最大长度不超过4096个ASCII字符
	GameServerSessionName *string `json:"GameServerSessionName,omitempty" name:"GameServerSessionName"`

	// 玩家延迟
	PlayerLatencies []*PlayerLatency `json:"PlayerLatencies,omitempty" name:"PlayerLatencies"`
}

func (r *StartGameServerSessionPlacementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartGameServerSessionPlacementRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlacementId")
	delete(f, "GameServerSessionQueueName")
	delete(f, "MaximumPlayerSessionCount")
	delete(f, "DesiredPlayerSessions")
	delete(f, "GameProperties")
	delete(f, "GameServerSessionData")
	delete(f, "GameServerSessionName")
	delete(f, "PlayerLatencies")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartGameServerSessionPlacementRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartGameServerSessionPlacementResponseParams struct {
	// 游戏服务器会话放置
	GameServerSessionPlacement *GameServerSessionPlacement `json:"GameServerSessionPlacement,omitempty" name:"GameServerSessionPlacement"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StartGameServerSessionPlacementResponse struct {
	*tchttp.BaseResponse
	Response *StartGameServerSessionPlacementResponseParams `json:"Response"`
}

func (r *StartGameServerSessionPlacementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartGameServerSessionPlacementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopFleetActionsRequestParams struct {
	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 服务器舰队扩展策略，值为["AUTO_SCALING"]
	Actions []*string `json:"Actions,omitempty" name:"Actions"`
}

type StopFleetActionsRequest struct {
	*tchttp.BaseRequest
	
	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 服务器舰队扩展策略，值为["AUTO_SCALING"]
	Actions []*string `json:"Actions,omitempty" name:"Actions"`
}

func (r *StopFleetActionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopFleetActionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	delete(f, "Actions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopFleetActionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopFleetActionsResponseParams struct {
	// 服务器舰队 Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopFleetActionsResponse struct {
	*tchttp.BaseResponse
	Response *StopFleetActionsResponseParams `json:"Response"`
}

func (r *StopFleetActionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopFleetActionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopGameServerSessionPlacementRequestParams struct {
	// 游戏服务器会话放置的唯一标识符
	PlacementId *string `json:"PlacementId,omitempty" name:"PlacementId"`
}

type StopGameServerSessionPlacementRequest struct {
	*tchttp.BaseRequest
	
	// 游戏服务器会话放置的唯一标识符
	PlacementId *string `json:"PlacementId,omitempty" name:"PlacementId"`
}

func (r *StopGameServerSessionPlacementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopGameServerSessionPlacementRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlacementId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopGameServerSessionPlacementRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopGameServerSessionPlacementResponseParams struct {
	// 游戏服务器会话放置
	GameServerSessionPlacement *GameServerSessionPlacement `json:"GameServerSessionPlacement,omitempty" name:"GameServerSessionPlacement"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopGameServerSessionPlacementResponse struct {
	*tchttp.BaseResponse
	Response *StopGameServerSessionPlacementResponseParams `json:"Response"`
}

func (r *StopGameServerSessionPlacementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopGameServerSessionPlacementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// 标签键，最大长度127字节
	Key *string `json:"Key,omitempty" name:"Key"`

	// 标签值，最大长度255字节
	Value *string `json:"Value,omitempty" name:"Value"`
}

type TargetConfiguration struct {
	// 预留存率
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetValue *uint64 `json:"TargetValue,omitempty" name:"TargetValue"`
}

type TimerConfiguration struct {
	// 定时器重复周期类型（未定义0，单次1、按天2、按月3、按周4）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimerType *int64 `json:"TimerType,omitempty" name:"TimerType"`

	// 定时器取值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimerValue *TimerValue `json:"TimerValue,omitempty" name:"TimerValue"`

	// 定时器开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 定时器结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type TimerFleetCapacity struct {
	// 扩缩容配置服务器舰队ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 期望实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DesiredInstances *int64 `json:"DesiredInstances,omitempty" name:"DesiredInstances"`

	// 最小实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinSize *int64 `json:"MinSize,omitempty" name:"MinSize"`

	// 最大实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`

	// 伸缩容间隔，单位：分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScalingInterval *int64 `json:"ScalingInterval,omitempty" name:"ScalingInterval"`

	// 扩缩容类型（手动1，自动2、未定义0）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScalingType *int64 `json:"ScalingType,omitempty" name:"ScalingType"`

	// 基于目标的扩展策略的设置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetConfiguration *TargetConfiguration `json:"TargetConfiguration,omitempty" name:"TargetConfiguration"`
}

type TimerScalingPolicy struct {
	// 定时器ID，进行encode，填写时更新
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimerId *string `json:"TimerId,omitempty" name:"TimerId"`

	// 定时器名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimerName *string `json:"TimerName,omitempty" name:"TimerName"`

	// 定时器状态(未定义0、未生效1、生效中2、已停止3、已过期4)
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimerStatus *int64 `json:"TimerStatus,omitempty" name:"TimerStatus"`

	// 定时器弹性伸缩策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimerFleetCapacity *TimerFleetCapacity `json:"TimerFleetCapacity,omitempty" name:"TimerFleetCapacity"`

	// 重复周期配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimerConfiguration *TimerConfiguration `json:"TimerConfiguration,omitempty" name:"TimerConfiguration"`
}

type TimerValue struct {
	// 每X天，执行一次(重复周期-按天/单次)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Day *int64 `json:"Day,omitempty" name:"Day"`

	// 每月从第x天，执行一次(重复周期-按月)
	// 注意：此字段可能返回 null，表示取不到有效值。
	FromDay *int64 `json:"FromDay,omitempty" name:"FromDay"`

	// 每月到第x天，执行一次(重复周期-按月)
	// 注意：此字段可能返回 null，表示取不到有效值。
	ToDay *int64 `json:"ToDay,omitempty" name:"ToDay"`

	// 重复周期-按周，周几（多个值,取值周一(1,2,3,4,5,6,7)周日）
	// 注意：此字段可能返回 null，表示取不到有效值。
	WeekDays []*int64 `json:"WeekDays,omitempty" name:"WeekDays"`
}

// Predefined struct for user
type UpdateAliasRequestParams struct {
	// 要更新的别名的唯一标识符
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`

	// 名字，长度不小于1字符不超过1024字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 别名的可读说明，长度不小于1字符不超过1024字符
	Description *string `json:"Description,omitempty" name:"Description"`

	// 别名的路由配置
	RoutingStrategy *RoutingStrategy `json:"RoutingStrategy,omitempty" name:"RoutingStrategy"`
}

type UpdateAliasRequest struct {
	*tchttp.BaseRequest
	
	// 要更新的别名的唯一标识符
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`

	// 名字，长度不小于1字符不超过1024字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 别名的可读说明，长度不小于1字符不超过1024字符
	Description *string `json:"Description,omitempty" name:"Description"`

	// 别名的路由配置
	RoutingStrategy *RoutingStrategy `json:"RoutingStrategy,omitempty" name:"RoutingStrategy"`
}

func (r *UpdateAliasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAliasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AliasId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "RoutingStrategy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAliasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAliasResponseParams struct {
	// 别名对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *Alias `json:"Alias,omitempty" name:"Alias"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateAliasResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAliasResponseParams `json:"Response"`
}

func (r *UpdateAliasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAliasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAssetRequestParams struct {
	// 生成包ID
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// 生成包名字，最小长度为1，最大长度为64
	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`

	// 生成包版本，最小长度为1，最大长度为64
	AssetVersion *string `json:"AssetVersion,omitempty" name:"AssetVersion"`
}

type UpdateAssetRequest struct {
	*tchttp.BaseRequest
	
	// 生成包ID
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// 生成包名字，最小长度为1，最大长度为64
	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`

	// 生成包版本，最小长度为1，最大长度为64
	AssetVersion *string `json:"AssetVersion,omitempty" name:"AssetVersion"`
}

func (r *UpdateAssetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAssetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetId")
	delete(f, "AssetName")
	delete(f, "AssetVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAssetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAssetResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateAssetResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAssetResponseParams `json:"Response"`
}

func (r *UpdateAssetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateBucketAccelerateOptRequestParams struct {
	// true为开启全球加速，false为关闭
	Allowed *bool `json:"Allowed,omitempty" name:"Allowed"`
}

type UpdateBucketAccelerateOptRequest struct {
	*tchttp.BaseRequest
	
	// true为开启全球加速，false为关闭
	Allowed *bool `json:"Allowed,omitempty" name:"Allowed"`
}

func (r *UpdateBucketAccelerateOptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateBucketAccelerateOptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Allowed")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateBucketAccelerateOptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateBucketAccelerateOptResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateBucketAccelerateOptResponse struct {
	*tchttp.BaseResponse
	Response *UpdateBucketAccelerateOptResponseParams `json:"Response"`
}

func (r *UpdateBucketAccelerateOptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateBucketAccelerateOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateBucketCORSOptRequestParams struct {
	// 允许的访问来源;具体参见 [cos文档](https://cloud.tencent.com/document/product/436/8279)
	AllowedOrigins []*string `json:"AllowedOrigins,omitempty" name:"AllowedOrigins"`

	// 允许的 HTTP 操作方法;可以配置多个:PUT、GET、POST、HEAD。[cos文档](https://cloud.tencent.com/document/product/436/8279)
	AllowedMethods []*string `json:"AllowedMethods,omitempty" name:"AllowedMethods"`

	// 用于指定允许浏览器发送 CORS 请求时携带的自定义 HTTP 请求头部;可以配置*，代表允许所有头部，为了避免遗漏，推荐配置为*。[cos文档](https://cloud.tencent.com/document/product/436/8279)
	AllowedHeaders []*string `json:"AllowedHeaders,omitempty" name:"AllowedHeaders"`

	// 跨域资源共享配置的有效时间，单位为秒。[cos文档](https://cloud.tencent.com/document/product/436/8279)
	MaxAgeSeconds *int64 `json:"MaxAgeSeconds,omitempty" name:"MaxAgeSeconds"`

	// 允许浏览器获取的 CORS 请求响应中的头部，不区分大小写；默认情况下浏览器只能访问简单响应头部：Cache-Control、Content-Type、Expires、Last-Modified，如果需要访问其他响应头部，需要添加 ExposeHeader 配置。[cos文档](https://cloud.tencent.com/document/product/436/8279)
	ExposeHeaders []*string `json:"ExposeHeaders,omitempty" name:"ExposeHeaders"`
}

type UpdateBucketCORSOptRequest struct {
	*tchttp.BaseRequest
	
	// 允许的访问来源;具体参见 [cos文档](https://cloud.tencent.com/document/product/436/8279)
	AllowedOrigins []*string `json:"AllowedOrigins,omitempty" name:"AllowedOrigins"`

	// 允许的 HTTP 操作方法;可以配置多个:PUT、GET、POST、HEAD。[cos文档](https://cloud.tencent.com/document/product/436/8279)
	AllowedMethods []*string `json:"AllowedMethods,omitempty" name:"AllowedMethods"`

	// 用于指定允许浏览器发送 CORS 请求时携带的自定义 HTTP 请求头部;可以配置*，代表允许所有头部，为了避免遗漏，推荐配置为*。[cos文档](https://cloud.tencent.com/document/product/436/8279)
	AllowedHeaders []*string `json:"AllowedHeaders,omitempty" name:"AllowedHeaders"`

	// 跨域资源共享配置的有效时间，单位为秒。[cos文档](https://cloud.tencent.com/document/product/436/8279)
	MaxAgeSeconds *int64 `json:"MaxAgeSeconds,omitempty" name:"MaxAgeSeconds"`

	// 允许浏览器获取的 CORS 请求响应中的头部，不区分大小写；默认情况下浏览器只能访问简单响应头部：Cache-Control、Content-Type、Expires、Last-Modified，如果需要访问其他响应头部，需要添加 ExposeHeader 配置。[cos文档](https://cloud.tencent.com/document/product/436/8279)
	ExposeHeaders []*string `json:"ExposeHeaders,omitempty" name:"ExposeHeaders"`
}

func (r *UpdateBucketCORSOptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateBucketCORSOptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AllowedOrigins")
	delete(f, "AllowedMethods")
	delete(f, "AllowedHeaders")
	delete(f, "MaxAgeSeconds")
	delete(f, "ExposeHeaders")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateBucketCORSOptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateBucketCORSOptResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateBucketCORSOptResponse struct {
	*tchttp.BaseResponse
	Response *UpdateBucketCORSOptResponseParams `json:"Response"`
}

func (r *UpdateBucketCORSOptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateBucketCORSOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateFleetAttributesRequestParams struct {
	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 服务器舰队描述，最小长度0，最大长度100
	Description *string `json:"Description,omitempty" name:"Description"`

	// 服务器舰队名称，最小长度1，最大长度50
	Name *string `json:"Name,omitempty" name:"Name"`

	// 保护策略：不保护NoProtection、完全保护FullProtection、时限保护TimeLimitProtection
	NewGameSessionProtectionPolicy *string `json:"NewGameSessionProtectionPolicy,omitempty" name:"NewGameSessionProtectionPolicy"`

	// 资源创建限制策略
	ResourceCreationLimitPolicy *ResourceCreationLimitPolicy `json:"ResourceCreationLimitPolicy,omitempty" name:"ResourceCreationLimitPolicy"`

	// 时限保护超时时间，默认60分钟，最小值5，最大值1440；当NewGameSessionProtectionPolicy为TimeLimitProtection时参数有效
	GameServerSessionProtectionTimeLimit *int64 `json:"GameServerSessionProtectionTimeLimit,omitempty" name:"GameServerSessionProtectionTimeLimit"`
}

type UpdateFleetAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 服务器舰队描述，最小长度0，最大长度100
	Description *string `json:"Description,omitempty" name:"Description"`

	// 服务器舰队名称，最小长度1，最大长度50
	Name *string `json:"Name,omitempty" name:"Name"`

	// 保护策略：不保护NoProtection、完全保护FullProtection、时限保护TimeLimitProtection
	NewGameSessionProtectionPolicy *string `json:"NewGameSessionProtectionPolicy,omitempty" name:"NewGameSessionProtectionPolicy"`

	// 资源创建限制策略
	ResourceCreationLimitPolicy *ResourceCreationLimitPolicy `json:"ResourceCreationLimitPolicy,omitempty" name:"ResourceCreationLimitPolicy"`

	// 时限保护超时时间，默认60分钟，最小值5，最大值1440；当NewGameSessionProtectionPolicy为TimeLimitProtection时参数有效
	GameServerSessionProtectionTimeLimit *int64 `json:"GameServerSessionProtectionTimeLimit,omitempty" name:"GameServerSessionProtectionTimeLimit"`
}

func (r *UpdateFleetAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateFleetAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	delete(f, "Description")
	delete(f, "Name")
	delete(f, "NewGameSessionProtectionPolicy")
	delete(f, "ResourceCreationLimitPolicy")
	delete(f, "GameServerSessionProtectionTimeLimit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateFleetAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateFleetAttributesResponseParams struct {
	// 服务器舰队Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateFleetAttributesResponse struct {
	*tchttp.BaseResponse
	Response *UpdateFleetAttributesResponseParams `json:"Response"`
}

func (r *UpdateFleetAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateFleetAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateFleetCapacityRequestParams struct {
	// 服务器舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 期望的服务器实例数
	DesiredInstances *uint64 `json:"DesiredInstances,omitempty" name:"DesiredInstances"`

	// 服务器实例数最小限制，最小值0，最大值不超过最高配额查看各地区最高配额减1
	MinSize *uint64 `json:"MinSize,omitempty" name:"MinSize"`

	// 服务器实例数最大限制，最小值1，最大值不超过最高配额查看各地区最高配额
	MaxSize *uint64 `json:"MaxSize,omitempty" name:"MaxSize"`

	// 服务器伸缩容间隔，单位分钟，最小值3，最大值30，默认值10
	ScalingInterval *uint64 `json:"ScalingInterval,omitempty" name:"ScalingInterval"`
}

type UpdateFleetCapacityRequest struct {
	*tchttp.BaseRequest
	
	// 服务器舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 期望的服务器实例数
	DesiredInstances *uint64 `json:"DesiredInstances,omitempty" name:"DesiredInstances"`

	// 服务器实例数最小限制，最小值0，最大值不超过最高配额查看各地区最高配额减1
	MinSize *uint64 `json:"MinSize,omitempty" name:"MinSize"`

	// 服务器实例数最大限制，最小值1，最大值不超过最高配额查看各地区最高配额
	MaxSize *uint64 `json:"MaxSize,omitempty" name:"MaxSize"`

	// 服务器伸缩容间隔，单位分钟，最小值3，最大值30，默认值10
	ScalingInterval *uint64 `json:"ScalingInterval,omitempty" name:"ScalingInterval"`
}

func (r *UpdateFleetCapacityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateFleetCapacityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	delete(f, "DesiredInstances")
	delete(f, "MinSize")
	delete(f, "MaxSize")
	delete(f, "ScalingInterval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateFleetCapacityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateFleetCapacityResponseParams struct {
	// 服务器舰队ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateFleetCapacityResponse struct {
	*tchttp.BaseResponse
	Response *UpdateFleetCapacityResponseParams `json:"Response"`
}

func (r *UpdateFleetCapacityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateFleetCapacityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateFleetNameRequestParams struct {
	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 服务器舰队名称，最小长度1，最大长度50
	Name *string `json:"Name,omitempty" name:"Name"`
}

type UpdateFleetNameRequest struct {
	*tchttp.BaseRequest
	
	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 服务器舰队名称，最小长度1，最大长度50
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *UpdateFleetNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateFleetNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateFleetNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateFleetNameResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateFleetNameResponse struct {
	*tchttp.BaseResponse
	Response *UpdateFleetNameResponseParams `json:"Response"`
}

func (r *UpdateFleetNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateFleetNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateFleetPortSettingsRequestParams struct {
	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 新增安全组
	InboundPermissionAuthorizations []*InboundPermissionAuthorization `json:"InboundPermissionAuthorizations,omitempty" name:"InboundPermissionAuthorizations"`

	// 移除安全组
	InboundPermissionRevocations []*InboundPermissionRevocations `json:"InboundPermissionRevocations,omitempty" name:"InboundPermissionRevocations"`
}

type UpdateFleetPortSettingsRequest struct {
	*tchttp.BaseRequest
	
	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 新增安全组
	InboundPermissionAuthorizations []*InboundPermissionAuthorization `json:"InboundPermissionAuthorizations,omitempty" name:"InboundPermissionAuthorizations"`

	// 移除安全组
	InboundPermissionRevocations []*InboundPermissionRevocations `json:"InboundPermissionRevocations,omitempty" name:"InboundPermissionRevocations"`
}

func (r *UpdateFleetPortSettingsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateFleetPortSettingsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	delete(f, "InboundPermissionAuthorizations")
	delete(f, "InboundPermissionRevocations")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateFleetPortSettingsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateFleetPortSettingsResponseParams struct {
	// 服务部署 Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateFleetPortSettingsResponse struct {
	*tchttp.BaseResponse
	Response *UpdateFleetPortSettingsResponseParams `json:"Response"`
}

func (r *UpdateFleetPortSettingsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateFleetPortSettingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGameServerSessionQueueRequestParams struct {
	// 游戏服务器会话队列名字，长度1~128
	Name *string `json:"Name,omitempty" name:"Name"`

	// 目的服务器舰队（可为别名）列表
	Destinations []*GameServerSessionQueueDestination `json:"Destinations,omitempty" name:"Destinations"`

	// 延迟策略集合
	PlayerLatencyPolicies []*PlayerLatencyPolicy `json:"PlayerLatencyPolicies,omitempty" name:"PlayerLatencyPolicies"`

	// 超时时间
	TimeoutInSeconds *uint64 `json:"TimeoutInSeconds,omitempty" name:"TimeoutInSeconds"`
}

type UpdateGameServerSessionQueueRequest struct {
	*tchttp.BaseRequest
	
	// 游戏服务器会话队列名字，长度1~128
	Name *string `json:"Name,omitempty" name:"Name"`

	// 目的服务器舰队（可为别名）列表
	Destinations []*GameServerSessionQueueDestination `json:"Destinations,omitempty" name:"Destinations"`

	// 延迟策略集合
	PlayerLatencyPolicies []*PlayerLatencyPolicy `json:"PlayerLatencyPolicies,omitempty" name:"PlayerLatencyPolicies"`

	// 超时时间
	TimeoutInSeconds *uint64 `json:"TimeoutInSeconds,omitempty" name:"TimeoutInSeconds"`
}

func (r *UpdateGameServerSessionQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGameServerSessionQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Destinations")
	delete(f, "PlayerLatencyPolicies")
	delete(f, "TimeoutInSeconds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateGameServerSessionQueueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGameServerSessionQueueResponseParams struct {
	// 部署服务组对象
	GameServerSessionQueue *GameServerSessionQueue `json:"GameServerSessionQueue,omitempty" name:"GameServerSessionQueue"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateGameServerSessionQueueResponse struct {
	*tchttp.BaseResponse
	Response *UpdateGameServerSessionQueueResponseParams `json:"Response"`
}

func (r *UpdateGameServerSessionQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGameServerSessionQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGameServerSessionRequestParams struct {
	// 游戏服务器会话ID，最小长度1个ASCII字符，最大长度不超过256个ASCII字符
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 最大玩家数量，最小值不小于0
	MaximumPlayerSessionCount *uint64 `json:"MaximumPlayerSessionCount,omitempty" name:"MaximumPlayerSessionCount"`

	// 游戏服务器会话名称，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 玩家会话创建策略，包括允许所有玩家加入和禁止所有玩家加入（ACCEPT_ALL,DENY_ALL）
	PlayerSessionCreationPolicy *string `json:"PlayerSessionCreationPolicy,omitempty" name:"PlayerSessionCreationPolicy"`

	// 保护策略，包括不保护、时限保护和完全保护(NoProtection,TimeLimitProtection,FullProtection)
	ProtectionPolicy *string `json:"ProtectionPolicy,omitempty" name:"ProtectionPolicy"`
}

type UpdateGameServerSessionRequest struct {
	*tchttp.BaseRequest
	
	// 游戏服务器会话ID，最小长度1个ASCII字符，最大长度不超过256个ASCII字符
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 最大玩家数量，最小值不小于0
	MaximumPlayerSessionCount *uint64 `json:"MaximumPlayerSessionCount,omitempty" name:"MaximumPlayerSessionCount"`

	// 游戏服务器会话名称，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 玩家会话创建策略，包括允许所有玩家加入和禁止所有玩家加入（ACCEPT_ALL,DENY_ALL）
	PlayerSessionCreationPolicy *string `json:"PlayerSessionCreationPolicy,omitempty" name:"PlayerSessionCreationPolicy"`

	// 保护策略，包括不保护、时限保护和完全保护(NoProtection,TimeLimitProtection,FullProtection)
	ProtectionPolicy *string `json:"ProtectionPolicy,omitempty" name:"ProtectionPolicy"`
}

func (r *UpdateGameServerSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGameServerSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GameServerSessionId")
	delete(f, "MaximumPlayerSessionCount")
	delete(f, "Name")
	delete(f, "PlayerSessionCreationPolicy")
	delete(f, "ProtectionPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateGameServerSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGameServerSessionResponseParams struct {
	// 更新后的游戏会话
	GameServerSession *GameServerSession `json:"GameServerSession,omitempty" name:"GameServerSession"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateGameServerSessionResponse struct {
	*tchttp.BaseResponse
	Response *UpdateGameServerSessionResponseParams `json:"Response"`
}

func (r *UpdateGameServerSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGameServerSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRuntimeConfigurationRequestParams struct {
	// 服务器舰队Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 服务器舰队配置
	RuntimeConfiguration *RuntimeConfiguration `json:"RuntimeConfiguration,omitempty" name:"RuntimeConfiguration"`
}

type UpdateRuntimeConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// 服务器舰队Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 服务器舰队配置
	RuntimeConfiguration *RuntimeConfiguration `json:"RuntimeConfiguration,omitempty" name:"RuntimeConfiguration"`
}

func (r *UpdateRuntimeConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRuntimeConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	delete(f, "RuntimeConfiguration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRuntimeConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRuntimeConfigurationResponseParams struct {
	// 服务器舰队配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuntimeConfiguration *RuntimeConfiguration `json:"RuntimeConfiguration,omitempty" name:"RuntimeConfiguration"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateRuntimeConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRuntimeConfigurationResponseParams `json:"Response"`
}

func (r *UpdateRuntimeConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRuntimeConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}