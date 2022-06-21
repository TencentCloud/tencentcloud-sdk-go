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

package v20180423

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type AttachCamRoleRequestParams struct {
	// 服务器ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 角色名称。
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

type AttachCamRoleRequest struct {
	*tchttp.BaseRequest
	
	// 服务器ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 角色名称。
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *AttachCamRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachCamRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RoleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachCamRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachCamRoleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AttachCamRoleResponse struct {
	*tchttp.BaseResponse
	Response *AttachCamRoleResponseParams `json:"Response"`
}

func (r *AttachCamRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachCamRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindPsaTagRequestParams struct {
	// 预授权规则ID
	PsaId *string `json:"PsaId,omitempty" name:"PsaId"`

	// 需要绑定的标签key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 需要绑定的标签value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type BindPsaTagRequest struct {
	*tchttp.BaseRequest
	
	// 预授权规则ID
	PsaId *string `json:"PsaId,omitempty" name:"PsaId"`

	// 需要绑定的标签key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 需要绑定的标签value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

func (r *BindPsaTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindPsaTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PsaId")
	delete(f, "TagKey")
	delete(f, "TagValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindPsaTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindPsaTagResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindPsaTagResponse struct {
	*tchttp.BaseResponse
	Response *BindPsaTagResponseParams `json:"Response"`
}

func (r *BindPsaTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindPsaTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BuyDevicesRequestParams struct {
	// 可用区ID。通过接口[查询地域以及可用区(DescribeRegions)](https://cloud.tencent.com/document/api/386/33564)获取可用区信息
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 部署服务器的操作系统ID。通过接口[查询操作系统信息(DescribeOsInfo)](https://cloud.tencent.com/document/product/386/32902)获取操作系统信息
	OsTypeId *uint64 `json:"OsTypeId,omitempty" name:"OsTypeId"`

	// RAID类型ID。通过接口[查询机型RAID方式以及系统盘大小(DescribeDeviceClassPartition)](https://cloud.tencent.com/document/api/386/32910)获取RAID信息
	RaidId *uint64 `json:"RaidId,omitempty" name:"RaidId"`

	// 购买数量
	GoodsCount *uint64 `json:"GoodsCount,omitempty" name:"GoodsCount"`

	// 购买至私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 购买至子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 购买的机型ID。通过接口[查询设备型号(DescribeDeviceClass)](https://cloud.tencent.com/document/api/386/32911)获取机型信息
	DeviceClassCode *string `json:"DeviceClassCode,omitempty" name:"DeviceClassCode"`

	// 购买时长单位，取值：M(月) D(天)
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 购买时长
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 是否安装安全Agent，取值：1(安装) 0(不安装)，默认取值0
	NeedSecurityAgent *uint64 `json:"NeedSecurityAgent,omitempty" name:"NeedSecurityAgent"`

	// 是否安装监控Agent，取值：1(安装) 0(不安装)，默认取值0
	NeedMonitorAgent *uint64 `json:"NeedMonitorAgent,omitempty" name:"NeedMonitorAgent"`

	// 是否安装EMR Agent，取值：1(安装) 0(不安装)，默认取值0
	NeedEMRAgent *uint64 `json:"NeedEMRAgent,omitempty" name:"NeedEMRAgent"`

	// 是否安装EMR软件包，取值：1(安装) 0(不安装)，默认取值0
	NeedEMRSoftware *uint64 `json:"NeedEMRSoftware,omitempty" name:"NeedEMRSoftware"`

	// 是否分配弹性公网IP，取值：1(分配) 0(不分配)，默认取值0
	ApplyEip *uint64 `json:"ApplyEip,omitempty" name:"ApplyEip"`

	// 弹性公网IP计费模式，取值：Flow(按流量计费) Bandwidth(按带宽计费)，默认取值Flow
	EipPayMode *string `json:"EipPayMode,omitempty" name:"EipPayMode"`

	// 弹性公网IP带宽限制，单位Mb
	EipBandwidth *uint64 `json:"EipBandwidth,omitempty" name:"EipBandwidth"`

	// 数据盘是否格式化，取值：1(格式化) 0(不格式化)，默认取值为1
	IsZoning *uint64 `json:"IsZoning,omitempty" name:"IsZoning"`

	// 物理机计费模式，取值：1(预付费) 2(后付费)，默认取值为1
	CpmPayMode *uint64 `json:"CpmPayMode,omitempty" name:"CpmPayMode"`

	// 自定义镜像ID，取值生效时用自定义镜像部署物理机
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 设置Linux root或Windows Administrator的密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 自动续费标志位，取值：1(自动续费) 0(不自动续费)，默认取值0
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 系统盘根分区大小，单位为G，默认取值10G。通过接口[查询机型RAID方式以及系统盘大小(DescribeDeviceClassPartition)](https://cloud.tencent.com/document/api/386/32910)获取根分区信息
	SysRootSpace *uint64 `json:"SysRootSpace,omitempty" name:"SysRootSpace"`

	// 系统盘swap分区或/boot/efi分区的大小，单位为G。若是uefi启动的机器，分区为/boot/efi，且此值是默认是2G。 普通机器为swap分区，可以不指定此分区。 机型是否是uefi启动，参见接口[查询设备型号(DescribeDeviceClass)](https://cloud.tencent.com/document/api/386/32911)
	SysSwaporuefiSpace *uint64 `json:"SysSwaporuefiSpace,omitempty" name:"SysSwaporuefiSpace"`

	// /usr/local分区大小，单位为G
	SysUsrlocalSpace *uint64 `json:"SysUsrlocalSpace,omitempty" name:"SysUsrlocalSpace"`

	// /data分区大小，单位为G。如果系统盘还有剩余大小，会分配给/data分区。（特殊情况：如果剩余空间不足10G，并且没有指定/data分区，则剩余空间会分配给Root分区）
	SysDataSpace *uint64 `json:"SysDataSpace,omitempty" name:"SysDataSpace"`

	// 是否开启超线程，取值：1(开启) 0(关闭)，默认取值1
	HyperThreading *uint64 `json:"HyperThreading,omitempty" name:"HyperThreading"`

	// 指定的内网IP列表，不指定时自动分配
	LanIps []*string `json:"LanIps,omitempty" name:"LanIps"`

	// 设备名称列表
	Aliases []*string `json:"Aliases,omitempty" name:"Aliases"`

	// CPU型号ID，自定义机型需要传入，取值：
	// <br/><li>1: E5-2620v3 (6核) &#42; 2</li><li>2: E5-2680v4 (14核) &#42; 2</li><li>3: E5-2670v3 (12核) &#42; 2</li><li>4: E5-2620v4 (8核) &#42; 2</li><li>5: 4110 (8核) &#42; 2</li><li>6: 6133 (20核) &#42; 2</li><br/>
	CpuId *uint64 `json:"CpuId,omitempty" name:"CpuId"`

	// 是否有RAID卡，取值：1(有) 0(无)，自定义机型需要传入
	ContainRaidCard *uint64 `json:"ContainRaidCard,omitempty" name:"ContainRaidCard"`

	// 内存大小，单位为G，自定义机型需要传入。取值参考接口[查询自定义机型部件信息(DescribeHardwareSpecification)](https://cloud.tencent.com/document/api/386/33565)返回值
	MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`

	// 系统盘ID，自定义机型需要传入。取值参考接口[查询自定义机型部件信息(DescribeHardwareSpecification)](https://cloud.tencent.com/document/api/386/33565)返回值
	SystemDiskTypeId *uint64 `json:"SystemDiskTypeId,omitempty" name:"SystemDiskTypeId"`

	// 系统盘数量，自定义机型需要传入。取值参考接口[查询自定义机型部件信息(DescribeHardwareSpecification)](https://cloud.tencent.com/document/api/386/33565)返回值
	SystemDiskCount *uint64 `json:"SystemDiskCount,omitempty" name:"SystemDiskCount"`

	// 数据盘ID，自定义机型需要传入。取值参考接口[查询自定义机型部件信息(DescribeHardwareSpecification)](https://cloud.tencent.com/document/api/386/33565)返回值
	DataDiskTypeId *uint64 `json:"DataDiskTypeId,omitempty" name:"DataDiskTypeId"`

	// 数据盘数量，自定义机型需要传入。取值参考接口[查询自定义机型部件信息(DescribeHardwareSpecification)](https://cloud.tencent.com/document/api/386/33565)返回值
	DataDiskCount *uint64 `json:"DataDiskCount,omitempty" name:"DataDiskCount"`

	// 绑定的标签列表
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 指定数据盘的文件系统格式，当前支持 EXT4和XFS选项， 默认为EXT4。 参数适用于数据盘和Linux， 且在IsZoning为1时生效
	FileSystem *string `json:"FileSystem,omitempty" name:"FileSystem"`

	// 此参数是为了防止重复发货。如果两次调用传入相同的BuySession，只会发货一次。 不要以设备别名作为BuySession，这样只会第一次购买成功。参数长度为128位，合法字符为大小字母，数字，下划线，横线。
	BuySession *string `json:"BuySession,omitempty" name:"BuySession"`

	// 绑定已有的安全组ID。仅在NeedSecurityAgent为1时生效
	SgId *string `json:"SgId,omitempty" name:"SgId"`

	// 安全组模板ID，由模板创建新安全组并绑定。TemplateId和SgId不能同时传入
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

type BuyDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 可用区ID。通过接口[查询地域以及可用区(DescribeRegions)](https://cloud.tencent.com/document/api/386/33564)获取可用区信息
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 部署服务器的操作系统ID。通过接口[查询操作系统信息(DescribeOsInfo)](https://cloud.tencent.com/document/product/386/32902)获取操作系统信息
	OsTypeId *uint64 `json:"OsTypeId,omitempty" name:"OsTypeId"`

	// RAID类型ID。通过接口[查询机型RAID方式以及系统盘大小(DescribeDeviceClassPartition)](https://cloud.tencent.com/document/api/386/32910)获取RAID信息
	RaidId *uint64 `json:"RaidId,omitempty" name:"RaidId"`

	// 购买数量
	GoodsCount *uint64 `json:"GoodsCount,omitempty" name:"GoodsCount"`

	// 购买至私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 购买至子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 购买的机型ID。通过接口[查询设备型号(DescribeDeviceClass)](https://cloud.tencent.com/document/api/386/32911)获取机型信息
	DeviceClassCode *string `json:"DeviceClassCode,omitempty" name:"DeviceClassCode"`

	// 购买时长单位，取值：M(月) D(天)
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 购买时长
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 是否安装安全Agent，取值：1(安装) 0(不安装)，默认取值0
	NeedSecurityAgent *uint64 `json:"NeedSecurityAgent,omitempty" name:"NeedSecurityAgent"`

	// 是否安装监控Agent，取值：1(安装) 0(不安装)，默认取值0
	NeedMonitorAgent *uint64 `json:"NeedMonitorAgent,omitempty" name:"NeedMonitorAgent"`

	// 是否安装EMR Agent，取值：1(安装) 0(不安装)，默认取值0
	NeedEMRAgent *uint64 `json:"NeedEMRAgent,omitempty" name:"NeedEMRAgent"`

	// 是否安装EMR软件包，取值：1(安装) 0(不安装)，默认取值0
	NeedEMRSoftware *uint64 `json:"NeedEMRSoftware,omitempty" name:"NeedEMRSoftware"`

	// 是否分配弹性公网IP，取值：1(分配) 0(不分配)，默认取值0
	ApplyEip *uint64 `json:"ApplyEip,omitempty" name:"ApplyEip"`

	// 弹性公网IP计费模式，取值：Flow(按流量计费) Bandwidth(按带宽计费)，默认取值Flow
	EipPayMode *string `json:"EipPayMode,omitempty" name:"EipPayMode"`

	// 弹性公网IP带宽限制，单位Mb
	EipBandwidth *uint64 `json:"EipBandwidth,omitempty" name:"EipBandwidth"`

	// 数据盘是否格式化，取值：1(格式化) 0(不格式化)，默认取值为1
	IsZoning *uint64 `json:"IsZoning,omitempty" name:"IsZoning"`

	// 物理机计费模式，取值：1(预付费) 2(后付费)，默认取值为1
	CpmPayMode *uint64 `json:"CpmPayMode,omitempty" name:"CpmPayMode"`

	// 自定义镜像ID，取值生效时用自定义镜像部署物理机
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 设置Linux root或Windows Administrator的密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 自动续费标志位，取值：1(自动续费) 0(不自动续费)，默认取值0
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 系统盘根分区大小，单位为G，默认取值10G。通过接口[查询机型RAID方式以及系统盘大小(DescribeDeviceClassPartition)](https://cloud.tencent.com/document/api/386/32910)获取根分区信息
	SysRootSpace *uint64 `json:"SysRootSpace,omitempty" name:"SysRootSpace"`

	// 系统盘swap分区或/boot/efi分区的大小，单位为G。若是uefi启动的机器，分区为/boot/efi，且此值是默认是2G。 普通机器为swap分区，可以不指定此分区。 机型是否是uefi启动，参见接口[查询设备型号(DescribeDeviceClass)](https://cloud.tencent.com/document/api/386/32911)
	SysSwaporuefiSpace *uint64 `json:"SysSwaporuefiSpace,omitempty" name:"SysSwaporuefiSpace"`

	// /usr/local分区大小，单位为G
	SysUsrlocalSpace *uint64 `json:"SysUsrlocalSpace,omitempty" name:"SysUsrlocalSpace"`

	// /data分区大小，单位为G。如果系统盘还有剩余大小，会分配给/data分区。（特殊情况：如果剩余空间不足10G，并且没有指定/data分区，则剩余空间会分配给Root分区）
	SysDataSpace *uint64 `json:"SysDataSpace,omitempty" name:"SysDataSpace"`

	// 是否开启超线程，取值：1(开启) 0(关闭)，默认取值1
	HyperThreading *uint64 `json:"HyperThreading,omitempty" name:"HyperThreading"`

	// 指定的内网IP列表，不指定时自动分配
	LanIps []*string `json:"LanIps,omitempty" name:"LanIps"`

	// 设备名称列表
	Aliases []*string `json:"Aliases,omitempty" name:"Aliases"`

	// CPU型号ID，自定义机型需要传入，取值：
	// <br/><li>1: E5-2620v3 (6核) &#42; 2</li><li>2: E5-2680v4 (14核) &#42; 2</li><li>3: E5-2670v3 (12核) &#42; 2</li><li>4: E5-2620v4 (8核) &#42; 2</li><li>5: 4110 (8核) &#42; 2</li><li>6: 6133 (20核) &#42; 2</li><br/>
	CpuId *uint64 `json:"CpuId,omitempty" name:"CpuId"`

	// 是否有RAID卡，取值：1(有) 0(无)，自定义机型需要传入
	ContainRaidCard *uint64 `json:"ContainRaidCard,omitempty" name:"ContainRaidCard"`

	// 内存大小，单位为G，自定义机型需要传入。取值参考接口[查询自定义机型部件信息(DescribeHardwareSpecification)](https://cloud.tencent.com/document/api/386/33565)返回值
	MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`

	// 系统盘ID，自定义机型需要传入。取值参考接口[查询自定义机型部件信息(DescribeHardwareSpecification)](https://cloud.tencent.com/document/api/386/33565)返回值
	SystemDiskTypeId *uint64 `json:"SystemDiskTypeId,omitempty" name:"SystemDiskTypeId"`

	// 系统盘数量，自定义机型需要传入。取值参考接口[查询自定义机型部件信息(DescribeHardwareSpecification)](https://cloud.tencent.com/document/api/386/33565)返回值
	SystemDiskCount *uint64 `json:"SystemDiskCount,omitempty" name:"SystemDiskCount"`

	// 数据盘ID，自定义机型需要传入。取值参考接口[查询自定义机型部件信息(DescribeHardwareSpecification)](https://cloud.tencent.com/document/api/386/33565)返回值
	DataDiskTypeId *uint64 `json:"DataDiskTypeId,omitempty" name:"DataDiskTypeId"`

	// 数据盘数量，自定义机型需要传入。取值参考接口[查询自定义机型部件信息(DescribeHardwareSpecification)](https://cloud.tencent.com/document/api/386/33565)返回值
	DataDiskCount *uint64 `json:"DataDiskCount,omitempty" name:"DataDiskCount"`

	// 绑定的标签列表
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 指定数据盘的文件系统格式，当前支持 EXT4和XFS选项， 默认为EXT4。 参数适用于数据盘和Linux， 且在IsZoning为1时生效
	FileSystem *string `json:"FileSystem,omitempty" name:"FileSystem"`

	// 此参数是为了防止重复发货。如果两次调用传入相同的BuySession，只会发货一次。 不要以设备别名作为BuySession，这样只会第一次购买成功。参数长度为128位，合法字符为大小字母，数字，下划线，横线。
	BuySession *string `json:"BuySession,omitempty" name:"BuySession"`

	// 绑定已有的安全组ID。仅在NeedSecurityAgent为1时生效
	SgId *string `json:"SgId,omitempty" name:"SgId"`

	// 安全组模板ID，由模板创建新安全组并绑定。TemplateId和SgId不能同时传入
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *BuyDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BuyDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "OsTypeId")
	delete(f, "RaidId")
	delete(f, "GoodsCount")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "DeviceClassCode")
	delete(f, "TimeUnit")
	delete(f, "TimeSpan")
	delete(f, "NeedSecurityAgent")
	delete(f, "NeedMonitorAgent")
	delete(f, "NeedEMRAgent")
	delete(f, "NeedEMRSoftware")
	delete(f, "ApplyEip")
	delete(f, "EipPayMode")
	delete(f, "EipBandwidth")
	delete(f, "IsZoning")
	delete(f, "CpmPayMode")
	delete(f, "ImageId")
	delete(f, "Password")
	delete(f, "AutoRenewFlag")
	delete(f, "SysRootSpace")
	delete(f, "SysSwaporuefiSpace")
	delete(f, "SysUsrlocalSpace")
	delete(f, "SysDataSpace")
	delete(f, "HyperThreading")
	delete(f, "LanIps")
	delete(f, "Aliases")
	delete(f, "CpuId")
	delete(f, "ContainRaidCard")
	delete(f, "MemSize")
	delete(f, "SystemDiskTypeId")
	delete(f, "SystemDiskCount")
	delete(f, "DataDiskTypeId")
	delete(f, "DataDiskCount")
	delete(f, "Tags")
	delete(f, "FileSystem")
	delete(f, "BuySession")
	delete(f, "SgId")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BuyDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BuyDevicesResponseParams struct {
	// 购买的物理机实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BuyDevicesResponse struct {
	*tchttp.BaseResponse
	Response *BuyDevicesResponseParams `json:"Response"`
}

func (r *BuyDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BuyDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CpuInfo struct {
	// CPU的ID
	CpuId *uint64 `json:"CpuId,omitempty" name:"CpuId"`

	// CPU型号描述
	CpuDescription *string `json:"CpuDescription,omitempty" name:"CpuDescription"`

	// 机型序列
	Series *uint64 `json:"Series,omitempty" name:"Series"`

	// 支持的RAID方式，0：有RAID卡，1：没有RAID卡
	ContainRaidCard []*uint64 `json:"ContainRaidCard,omitempty" name:"ContainRaidCard"`
}

// Predefined struct for user
type CreateCustomImageRequestParams struct {
	// 用于制作镜像的物理机ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 镜像别名
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 镜像描述
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`
}

type CreateCustomImageRequest struct {
	*tchttp.BaseRequest
	
	// 用于制作镜像的物理机ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 镜像别名
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 镜像描述
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`
}

func (r *CreateCustomImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ImageName")
	delete(f, "ImageDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCustomImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomImageResponseParams struct {
	// 黑石异步任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 镜像ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCustomImageResponse struct {
	*tchttp.BaseResponse
	Response *CreateCustomImageResponseParams `json:"Response"`
}

func (r *CreateCustomImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePsaRegulationRequestParams struct {
	// 规则别名
	PsaName *string `json:"PsaName,omitempty" name:"PsaName"`

	// 关联的故障类型ID列表
	TaskTypeIds []*uint64 `json:"TaskTypeIds,omitempty" name:"TaskTypeIds"`

	// 维修实例上限，默认为5
	RepairLimit *uint64 `json:"RepairLimit,omitempty" name:"RepairLimit"`

	// 规则备注
	PsaDescription *string `json:"PsaDescription,omitempty" name:"PsaDescription"`
}

type CreatePsaRegulationRequest struct {
	*tchttp.BaseRequest
	
	// 规则别名
	PsaName *string `json:"PsaName,omitempty" name:"PsaName"`

	// 关联的故障类型ID列表
	TaskTypeIds []*uint64 `json:"TaskTypeIds,omitempty" name:"TaskTypeIds"`

	// 维修实例上限，默认为5
	RepairLimit *uint64 `json:"RepairLimit,omitempty" name:"RepairLimit"`

	// 规则备注
	PsaDescription *string `json:"PsaDescription,omitempty" name:"PsaDescription"`
}

func (r *CreatePsaRegulationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePsaRegulationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PsaName")
	delete(f, "TaskTypeIds")
	delete(f, "RepairLimit")
	delete(f, "PsaDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePsaRegulationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePsaRegulationResponseParams struct {
	// 创建的预授权规则ID
	PsaId *string `json:"PsaId,omitempty" name:"PsaId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePsaRegulationResponse struct {
	*tchttp.BaseResponse
	Response *CreatePsaRegulationResponseParams `json:"Response"`
}

func (r *CreatePsaRegulationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePsaRegulationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSpotDeviceRequestParams struct {
	// 可用区名称。如ap-guangzhou-bls-1, 通过DescribeRegions获取
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 计算单元类型, 如v3.c2.medium，更详细的ComputeType参考[竞价实例产品文档](https://cloud.tencent.com/document/product/386/30256)
	ComputeType *string `json:"ComputeType,omitempty" name:"ComputeType"`

	// 操作系统类型ID
	OsTypeId *uint64 `json:"OsTypeId,omitempty" name:"OsTypeId"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 购买的计算单元个数
	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 出价策略。可取值为SpotWithPriceLimit和SpotAsPriceGo。SpotWithPriceLimit，用户设置价格上限，需要传SpotPriceLimit参数， 如果市场价高于用户的指定价格，则购买不成功;  SpotAsPriceGo 是随市场价的策略。
	SpotStrategy *string `json:"SpotStrategy,omitempty" name:"SpotStrategy"`

	// 用户设置的价格。当为SpotWithPriceLimit竞价策略时有效
	SpotPriceLimit *float64 `json:"SpotPriceLimit,omitempty" name:"SpotPriceLimit"`

	// 设置竞价实例密码。可选参数，没有指定会生成随机密码
	Passwd *string `json:"Passwd,omitempty" name:"Passwd"`
}

type CreateSpotDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 可用区名称。如ap-guangzhou-bls-1, 通过DescribeRegions获取
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 计算单元类型, 如v3.c2.medium，更详细的ComputeType参考[竞价实例产品文档](https://cloud.tencent.com/document/product/386/30256)
	ComputeType *string `json:"ComputeType,omitempty" name:"ComputeType"`

	// 操作系统类型ID
	OsTypeId *uint64 `json:"OsTypeId,omitempty" name:"OsTypeId"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 购买的计算单元个数
	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 出价策略。可取值为SpotWithPriceLimit和SpotAsPriceGo。SpotWithPriceLimit，用户设置价格上限，需要传SpotPriceLimit参数， 如果市场价高于用户的指定价格，则购买不成功;  SpotAsPriceGo 是随市场价的策略。
	SpotStrategy *string `json:"SpotStrategy,omitempty" name:"SpotStrategy"`

	// 用户设置的价格。当为SpotWithPriceLimit竞价策略时有效
	SpotPriceLimit *float64 `json:"SpotPriceLimit,omitempty" name:"SpotPriceLimit"`

	// 设置竞价实例密码。可选参数，没有指定会生成随机密码
	Passwd *string `json:"Passwd,omitempty" name:"Passwd"`
}

func (r *CreateSpotDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSpotDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "ComputeType")
	delete(f, "OsTypeId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "GoodsNum")
	delete(f, "SpotStrategy")
	delete(f, "SpotPriceLimit")
	delete(f, "Passwd")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSpotDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSpotDeviceResponseParams struct {
	// 创建的服务器ID
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 任务ID
	FlowId *uint64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSpotDeviceResponse struct {
	*tchttp.BaseResponse
	Response *CreateSpotDeviceResponseParams `json:"Response"`
}

func (r *CreateSpotDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSpotDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserCmdRequestParams struct {
	// 用户自定义脚本的名称
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 命令适用的操作系统类型，取值linux或xserver
	OsType *string `json:"OsType,omitempty" name:"OsType"`

	// 脚本内容，必须经过base64编码
	Content *string `json:"Content,omitempty" name:"Content"`
}

type CreateUserCmdRequest struct {
	*tchttp.BaseRequest
	
	// 用户自定义脚本的名称
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 命令适用的操作系统类型，取值linux或xserver
	OsType *string `json:"OsType,omitempty" name:"OsType"`

	// 脚本内容，必须经过base64编码
	Content *string `json:"Content,omitempty" name:"Content"`
}

func (r *CreateUserCmdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserCmdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Alias")
	delete(f, "OsType")
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserCmdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserCmdResponseParams struct {
	// 脚本ID
	CmdId *string `json:"CmdId,omitempty" name:"CmdId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateUserCmdResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserCmdResponseParams `json:"Response"`
}

func (r *CreateUserCmdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserCmdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomImage struct {
	// 镜像ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 镜像别名
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 镜像状态码
	ImageStatus *uint64 `json:"ImageStatus,omitempty" name:"ImageStatus"`

	// 镜像OS名
	OsClass *string `json:"OsClass,omitempty" name:"OsClass"`

	// 镜像OS版本
	OsVersion *string `json:"OsVersion,omitempty" name:"OsVersion"`

	// OS是64还是32位
	OsBit *uint64 `json:"OsBit,omitempty" name:"OsBit"`

	// 镜像大小(M)
	ImageSize *uint64 `json:"ImageSize,omitempty" name:"ImageSize"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 分区信息
	PartitionInfoSet []*PartitionInfo `json:"PartitionInfoSet,omitempty" name:"PartitionInfoSet"`

	// 适用机型
	DeviceClassCode *string `json:"DeviceClassCode,omitempty" name:"DeviceClassCode"`

	// 备注
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`

	// 原始镜像id
	OsTypeId *uint64 `json:"OsTypeId,omitempty" name:"OsTypeId"`
}

type CustomImageProcess struct {
	// 步骤
	StepName *string `json:"StepName,omitempty" name:"StepName"`

	// 此步骤开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 0: 已完成 1: 当前进行 2: 未开始
	StepType *uint64 `json:"StepType,omitempty" name:"StepType"`
}

// Predefined struct for user
type DeleteCustomImagesRequestParams struct {
	// 准备删除的镜像ID列表
	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`
}

type DeleteCustomImagesRequest struct {
	*tchttp.BaseRequest
	
	// 准备删除的镜像ID列表
	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`
}

func (r *DeleteCustomImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCustomImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomImagesResponseParams struct {
	// 黑石异步任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteCustomImagesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCustomImagesResponseParams `json:"Response"`
}

func (r *DeleteCustomImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePsaRegulationRequestParams struct {
	// 预授权规则ID
	PsaId *string `json:"PsaId,omitempty" name:"PsaId"`
}

type DeletePsaRegulationRequest struct {
	*tchttp.BaseRequest
	
	// 预授权规则ID
	PsaId *string `json:"PsaId,omitempty" name:"PsaId"`
}

func (r *DeletePsaRegulationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePsaRegulationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PsaId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePsaRegulationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePsaRegulationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeletePsaRegulationResponse struct {
	*tchttp.BaseResponse
	Response *DeletePsaRegulationResponseParams `json:"Response"`
}

func (r *DeletePsaRegulationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePsaRegulationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserCmdsRequestParams struct {
	// 需要删除的脚本ID
	CmdIds []*string `json:"CmdIds,omitempty" name:"CmdIds"`
}

type DeleteUserCmdsRequest struct {
	*tchttp.BaseRequest
	
	// 需要删除的脚本ID
	CmdIds []*string `json:"CmdIds,omitempty" name:"CmdIds"`
}

func (r *DeleteUserCmdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserCmdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CmdIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserCmdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserCmdsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteUserCmdsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserCmdsResponseParams `json:"Response"`
}

func (r *DeleteUserCmdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserCmdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomImageProcessRequestParams struct {
	// 镜像ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

type DescribeCustomImageProcessRequest struct {
	*tchttp.BaseRequest
	
	// 镜像ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *DescribeCustomImageProcessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomImageProcessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomImageProcessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomImageProcessResponseParams struct {
	// 镜像制作进度
	CustomImageProcessSet []*CustomImageProcess `json:"CustomImageProcessSet,omitempty" name:"CustomImageProcessSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCustomImageProcessResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomImageProcessResponseParams `json:"Response"`
}

func (r *DescribeCustomImageProcessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomImageProcessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomImagesRequestParams struct {
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数量限制
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段，仅支持CreateTime
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方式 0:递增(默认) 1:递减
	Order *uint64 `json:"Order,omitempty" name:"Order"`

	// 按ImageId查找指定镜像信息，ImageId字段存在时其他字段失效
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 模糊查询过滤，可以查询镜像ID或镜像名
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// <ul>
	// 镜像状态过滤列表，有效取值为：
	// <li>1：制作中</li>
	// <li>2：制作失败</li>
	// <li>3：正常</li>
	// <li>4：删除中</li>
	// </ul>
	ImageStatus []*uint64 `json:"ImageStatus,omitempty" name:"ImageStatus"`
}

type DescribeCustomImagesRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数量限制
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段，仅支持CreateTime
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方式 0:递增(默认) 1:递减
	Order *uint64 `json:"Order,omitempty" name:"Order"`

	// 按ImageId查找指定镜像信息，ImageId字段存在时其他字段失效
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 模糊查询过滤，可以查询镜像ID或镜像名
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// <ul>
	// 镜像状态过滤列表，有效取值为：
	// <li>1：制作中</li>
	// <li>2：制作失败</li>
	// <li>3：正常</li>
	// <li>4：删除中</li>
	// </ul>
	ImageStatus []*uint64 `json:"ImageStatus,omitempty" name:"ImageStatus"`
}

func (r *DescribeCustomImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "Order")
	delete(f, "ImageId")
	delete(f, "SearchKey")
	delete(f, "ImageStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomImagesResponseParams struct {
	// 返回镜像数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 镜像信息列表
	CustomImageSet []*CustomImage `json:"CustomImageSet,omitempty" name:"CustomImageSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCustomImagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomImagesResponseParams `json:"Response"`
}

func (r *DescribeCustomImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceClassPartitionRequestParams struct {
	// 设备类型代号。代号通过接口[查询设备型号(DescribeDeviceClass)](https://cloud.tencent.com/document/api/386/32911)查询。标准机型需要传入此参数。虽是可选参数，但DeviceClassCode和InstanceId参数，必须要填写一个。
	DeviceClassCode *string `json:"DeviceClassCode,omitempty" name:"DeviceClassCode"`

	// 需要查询自定义机型RAID信息时，传入自定义机型实例ID。InstanceId存在时其余参数失效。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// CPU型号ID，查询自定义机型时需要传入
	CpuId *uint64 `json:"CpuId,omitempty" name:"CpuId"`

	// 内存大小，单位为G，查询自定义机型时需要传入
	MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`

	// 是否有RAID卡，取值：1(有) 0(无)。查询自定义机型时需要传入
	ContainRaidCard *uint64 `json:"ContainRaidCard,omitempty" name:"ContainRaidCard"`

	// 系统盘类型ID，查询自定义机型时需要传入
	SystemDiskTypeId *uint64 `json:"SystemDiskTypeId,omitempty" name:"SystemDiskTypeId"`

	// 系统盘数量，查询自定义机型时需要传入
	SystemDiskCount *uint64 `json:"SystemDiskCount,omitempty" name:"SystemDiskCount"`

	// 数据盘类型ID，查询自定义机型时可传入
	DataDiskTypeId *uint64 `json:"DataDiskTypeId,omitempty" name:"DataDiskTypeId"`

	// 数据盘数量，查询自定义机型时可传入
	DataDiskCount *uint64 `json:"DataDiskCount,omitempty" name:"DataDiskCount"`
}

type DescribeDeviceClassPartitionRequest struct {
	*tchttp.BaseRequest
	
	// 设备类型代号。代号通过接口[查询设备型号(DescribeDeviceClass)](https://cloud.tencent.com/document/api/386/32911)查询。标准机型需要传入此参数。虽是可选参数，但DeviceClassCode和InstanceId参数，必须要填写一个。
	DeviceClassCode *string `json:"DeviceClassCode,omitempty" name:"DeviceClassCode"`

	// 需要查询自定义机型RAID信息时，传入自定义机型实例ID。InstanceId存在时其余参数失效。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// CPU型号ID，查询自定义机型时需要传入
	CpuId *uint64 `json:"CpuId,omitempty" name:"CpuId"`

	// 内存大小，单位为G，查询自定义机型时需要传入
	MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`

	// 是否有RAID卡，取值：1(有) 0(无)。查询自定义机型时需要传入
	ContainRaidCard *uint64 `json:"ContainRaidCard,omitempty" name:"ContainRaidCard"`

	// 系统盘类型ID，查询自定义机型时需要传入
	SystemDiskTypeId *uint64 `json:"SystemDiskTypeId,omitempty" name:"SystemDiskTypeId"`

	// 系统盘数量，查询自定义机型时需要传入
	SystemDiskCount *uint64 `json:"SystemDiskCount,omitempty" name:"SystemDiskCount"`

	// 数据盘类型ID，查询自定义机型时可传入
	DataDiskTypeId *uint64 `json:"DataDiskTypeId,omitempty" name:"DataDiskTypeId"`

	// 数据盘数量，查询自定义机型时可传入
	DataDiskCount *uint64 `json:"DataDiskCount,omitempty" name:"DataDiskCount"`
}

func (r *DescribeDeviceClassPartitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceClassPartitionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceClassCode")
	delete(f, "InstanceId")
	delete(f, "CpuId")
	delete(f, "MemSize")
	delete(f, "ContainRaidCard")
	delete(f, "SystemDiskTypeId")
	delete(f, "SystemDiskCount")
	delete(f, "DataDiskTypeId")
	delete(f, "DataDiskCount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceClassPartitionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceClassPartitionResponseParams struct {
	// 支持的RAID格式列表
	DeviceClassPartitionInfoSet []*DeviceClassPartitionInfo `json:"DeviceClassPartitionInfoSet,omitempty" name:"DeviceClassPartitionInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeviceClassPartitionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceClassPartitionResponseParams `json:"Response"`
}

func (r *DescribeDeviceClassPartitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceClassPartitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceClassRequestParams struct {
	// 是否仅查询在售标准机型配置信息。取值0：查询所有机型；1：查询在售机型。默认为1
	OnSale *uint64 `json:"OnSale,omitempty" name:"OnSale"`

	// 是否返回价格信息。取值0：不返回价格信息，接口返回速度更快；1：返回价格信息。默认为1
	NeedPriceInfo *uint64 `json:"NeedPriceInfo,omitempty" name:"NeedPriceInfo"`
}

type DescribeDeviceClassRequest struct {
	*tchttp.BaseRequest
	
	// 是否仅查询在售标准机型配置信息。取值0：查询所有机型；1：查询在售机型。默认为1
	OnSale *uint64 `json:"OnSale,omitempty" name:"OnSale"`

	// 是否返回价格信息。取值0：不返回价格信息，接口返回速度更快；1：返回价格信息。默认为1
	NeedPriceInfo *uint64 `json:"NeedPriceInfo,omitempty" name:"NeedPriceInfo"`
}

func (r *DescribeDeviceClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceClassRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OnSale")
	delete(f, "NeedPriceInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceClassRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceClassResponseParams struct {
	// 物理机设备类型列表
	DeviceClassSet []*DeviceClass `json:"DeviceClassSet,omitempty" name:"DeviceClassSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeviceClassResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceClassResponseParams `json:"Response"`
}

func (r *DescribeDeviceClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceHardwareInfoRequestParams struct {
	// 设备 ID 列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type DescribeDeviceHardwareInfoRequest struct {
	*tchttp.BaseRequest
	
	// 设备 ID 列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DescribeDeviceHardwareInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceHardwareInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceHardwareInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceHardwareInfoResponseParams struct {
	// 设备硬件配置信息
	DeviceHardwareInfoSet []*DeviceHardwareInfo `json:"DeviceHardwareInfoSet,omitempty" name:"DeviceHardwareInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeviceHardwareInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceHardwareInfoResponseParams `json:"Response"`
}

func (r *DescribeDeviceHardwareInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceHardwareInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceInventoryRequestParams struct {
	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 设备型号
	DeviceClassCode *string `json:"DeviceClassCode,omitempty" name:"DeviceClassCode"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// CPU型号ID，查询自定义机型时必填
	CpuId *uint64 `json:"CpuId,omitempty" name:"CpuId"`

	// 内存大小，单位为G，查询自定义机型时必填
	MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`

	// 是否有RAID卡，取值：1(有) 0(无)，查询自定义机型时必填
	ContainRaidCard *uint64 `json:"ContainRaidCard,omitempty" name:"ContainRaidCard"`

	// 系统盘类型ID，查询自定义机型时必填
	SystemDiskTypeId *uint64 `json:"SystemDiskTypeId,omitempty" name:"SystemDiskTypeId"`

	// 系统盘数量，查询自定义机型时必填
	SystemDiskCount *uint64 `json:"SystemDiskCount,omitempty" name:"SystemDiskCount"`

	// 数据盘类型ID，查询自定义机型时可填
	DataDiskTypeId *uint64 `json:"DataDiskTypeId,omitempty" name:"DataDiskTypeId"`

	// 数据盘数量，查询自定义机型时可填
	DataDiskCount *uint64 `json:"DataDiskCount,omitempty" name:"DataDiskCount"`
}

type DescribeDeviceInventoryRequest struct {
	*tchttp.BaseRequest
	
	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 设备型号
	DeviceClassCode *string `json:"DeviceClassCode,omitempty" name:"DeviceClassCode"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// CPU型号ID，查询自定义机型时必填
	CpuId *uint64 `json:"CpuId,omitempty" name:"CpuId"`

	// 内存大小，单位为G，查询自定义机型时必填
	MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`

	// 是否有RAID卡，取值：1(有) 0(无)，查询自定义机型时必填
	ContainRaidCard *uint64 `json:"ContainRaidCard,omitempty" name:"ContainRaidCard"`

	// 系统盘类型ID，查询自定义机型时必填
	SystemDiskTypeId *uint64 `json:"SystemDiskTypeId,omitempty" name:"SystemDiskTypeId"`

	// 系统盘数量，查询自定义机型时必填
	SystemDiskCount *uint64 `json:"SystemDiskCount,omitempty" name:"SystemDiskCount"`

	// 数据盘类型ID，查询自定义机型时可填
	DataDiskTypeId *uint64 `json:"DataDiskTypeId,omitempty" name:"DataDiskTypeId"`

	// 数据盘数量，查询自定义机型时可填
	DataDiskCount *uint64 `json:"DataDiskCount,omitempty" name:"DataDiskCount"`
}

func (r *DescribeDeviceInventoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceInventoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "DeviceClassCode")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "CpuId")
	delete(f, "MemSize")
	delete(f, "ContainRaidCard")
	delete(f, "SystemDiskTypeId")
	delete(f, "SystemDiskCount")
	delete(f, "DataDiskTypeId")
	delete(f, "DataDiskCount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceInventoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceInventoryResponseParams struct {
	// 库存设备数量
	DeviceCount *uint64 `json:"DeviceCount,omitempty" name:"DeviceCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeviceInventoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceInventoryResponseParams `json:"Response"`
}

func (r *DescribeDeviceInventoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceInventoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceOperationLogRequestParams struct {
	// 设备实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 查询开始日期
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束日期
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeDeviceOperationLogRequest struct {
	*tchttp.BaseRequest
	
	// 设备实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 查询开始日期
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束日期
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDeviceOperationLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceOperationLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceOperationLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceOperationLogResponseParams struct {
	// 操作日志列表
	DeviceOperationLogSet []*DeviceOperationLog `json:"DeviceOperationLogSet,omitempty" name:"DeviceOperationLogSet"`

	// 返回数目
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeviceOperationLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceOperationLogResponseParams `json:"Response"`
}

func (r *DescribeDeviceOperationLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceOperationLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicePartitionRequestParams struct {
	// 物理机ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeDevicePartitionRequest struct {
	*tchttp.BaseRequest
	
	// 物理机ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDevicePartitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicePartitionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDevicePartitionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicePartitionResponseParams struct {
	// 物理机分区格式
	DevicePartition *DevicePartition `json:"DevicePartition,omitempty" name:"DevicePartition"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDevicePartitionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDevicePartitionResponseParams `json:"Response"`
}

func (r *DescribeDevicePartitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicePartitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicePositionRequestParams struct {
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数量限制
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 实例别名
	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

type DescribeDevicePositionRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数量限制
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 实例别名
	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

func (r *DescribeDevicePositionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicePositionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "InstanceIds")
	delete(f, "Alias")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDevicePositionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicePositionResponseParams struct {
	// 返回数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 设备所在机架信息
	DevicePositionInfoSet []*DevicePositionInfo `json:"DevicePositionInfoSet,omitempty" name:"DevicePositionInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDevicePositionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDevicePositionResponseParams `json:"Response"`
}

func (r *DescribeDevicePositionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicePositionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicePriceInfoRequestParams struct {
	// 需要查询的实例列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 购买时长单位，当前只支持取值为m
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 购买时长
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
}

type DescribeDevicePriceInfoRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的实例列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 购买时长单位，当前只支持取值为m
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 购买时长
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
}

func (r *DescribeDevicePriceInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicePriceInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "TimeUnit")
	delete(f, "TimeSpan")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDevicePriceInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicePriceInfoResponseParams struct {
	// 服务器价格信息列表
	DevicePriceInfoSet []*DevicePriceInfo `json:"DevicePriceInfoSet,omitempty" name:"DevicePriceInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDevicePriceInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDevicePriceInfoResponseParams `json:"Response"`
}

func (r *DescribeDevicePriceInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicePriceInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicesRequestParams struct {
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 机型ID，通过接口[查询设备型号(DescribeDeviceClass)](https://cloud.tencent.com/document/api/386/32911)查询
	DeviceClassCode *string `json:"DeviceClassCode,omitempty" name:"DeviceClassCode"`

	// 设备ID数组
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 外网IP数组
	WanIps []*string `json:"WanIps,omitempty" name:"WanIps"`

	// 内网IP数组
	LanIps []*string `json:"LanIps,omitempty" name:"LanIps"`

	// 设备名称
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 模糊IP查询
	VagueIp *string `json:"VagueIp,omitempty" name:"VagueIp"`

	// 设备到期时间查询的起始时间
	DeadlineStartTime *string `json:"DeadlineStartTime,omitempty" name:"DeadlineStartTime"`

	// 设备到期时间查询的结束时间
	DeadlineEndTime *string `json:"DeadlineEndTime,omitempty" name:"DeadlineEndTime"`

	// 自动续费标志 0:不自动续费，1:自动续费
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 私有网络唯一ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网唯一ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 设备类型，取值有: compute(计算型), standard(标准型), storage(存储型) 等
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 竞价实例机器的过滤。如果未指定此参数，则不做过滤。0: 查询非竞价实例的机器; 1: 查询竞价实例的机器。
	IsLuckyDevice *uint64 `json:"IsLuckyDevice,omitempty" name:"IsLuckyDevice"`

	// 排序字段
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方式，取值：0:增序(默认)，1:降序
	Order *uint64 `json:"Order,omitempty" name:"Order"`

	// 按照维保方式过滤。可取值为 Maintain: 在保;  WillExpire: 即将过保; Expire: 已过保
	MaintainStatus *string `json:"MaintainStatus,omitempty" name:"MaintainStatus"`
}

type DescribeDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 机型ID，通过接口[查询设备型号(DescribeDeviceClass)](https://cloud.tencent.com/document/api/386/32911)查询
	DeviceClassCode *string `json:"DeviceClassCode,omitempty" name:"DeviceClassCode"`

	// 设备ID数组
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 外网IP数组
	WanIps []*string `json:"WanIps,omitempty" name:"WanIps"`

	// 内网IP数组
	LanIps []*string `json:"LanIps,omitempty" name:"LanIps"`

	// 设备名称
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 模糊IP查询
	VagueIp *string `json:"VagueIp,omitempty" name:"VagueIp"`

	// 设备到期时间查询的起始时间
	DeadlineStartTime *string `json:"DeadlineStartTime,omitempty" name:"DeadlineStartTime"`

	// 设备到期时间查询的结束时间
	DeadlineEndTime *string `json:"DeadlineEndTime,omitempty" name:"DeadlineEndTime"`

	// 自动续费标志 0:不自动续费，1:自动续费
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 私有网络唯一ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网唯一ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 设备类型，取值有: compute(计算型), standard(标准型), storage(存储型) 等
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 竞价实例机器的过滤。如果未指定此参数，则不做过滤。0: 查询非竞价实例的机器; 1: 查询竞价实例的机器。
	IsLuckyDevice *uint64 `json:"IsLuckyDevice,omitempty" name:"IsLuckyDevice"`

	// 排序字段
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方式，取值：0:增序(默认)，1:降序
	Order *uint64 `json:"Order,omitempty" name:"Order"`

	// 按照维保方式过滤。可取值为 Maintain: 在保;  WillExpire: 即将过保; Expire: 已过保
	MaintainStatus *string `json:"MaintainStatus,omitempty" name:"MaintainStatus"`
}

func (r *DescribeDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "DeviceClassCode")
	delete(f, "InstanceIds")
	delete(f, "WanIps")
	delete(f, "LanIps")
	delete(f, "Alias")
	delete(f, "VagueIp")
	delete(f, "DeadlineStartTime")
	delete(f, "DeadlineEndTime")
	delete(f, "AutoRenewFlag")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Tags")
	delete(f, "DeviceType")
	delete(f, "IsLuckyDevice")
	delete(f, "OrderField")
	delete(f, "Order")
	delete(f, "MaintainStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicesResponseParams struct {
	// 返回数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 物理机信息列表
	DeviceInfoSet []*DeviceInfo `json:"DeviceInfoSet,omitempty" name:"DeviceInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDevicesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDevicesResponseParams `json:"Response"`
}

func (r *DescribeDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHardwareSpecificationRequestParams struct {

}

type DescribeHardwareSpecificationRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeHardwareSpecificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHardwareSpecificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHardwareSpecificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHardwareSpecificationResponseParams struct {
	// CPU型号列表
	CpuInfoSet []*CpuInfo `json:"CpuInfoSet,omitempty" name:"CpuInfoSet"`

	// 内存的取值，单位为G
	MemSet []*uint64 `json:"MemSet,omitempty" name:"MemSet"`

	// 硬盘型号列表
	DiskInfoSet []*DiskInfo `json:"DiskInfoSet,omitempty" name:"DiskInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHardwareSpecificationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHardwareSpecificationResponseParams `json:"Response"`
}

func (r *DescribeHardwareSpecificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHardwareSpecificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostedDeviceOutBandInfoRequestParams struct {
	// 托管设备的唯一ID数组,数组个数不超过20
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 可用区ID
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type DescribeHostedDeviceOutBandInfoRequest struct {
	*tchttp.BaseRequest
	
	// 托管设备的唯一ID数组,数组个数不超过20
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 可用区ID
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *DescribeHostedDeviceOutBandInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostedDeviceOutBandInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "Zone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostedDeviceOutBandInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostedDeviceOutBandInfoResponseParams struct {
	// 托管设备带外信息
	HostedDeviceOutBandInfoSet []*HostedDeviceOutBandInfo `json:"HostedDeviceOutBandInfoSet,omitempty" name:"HostedDeviceOutBandInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHostedDeviceOutBandInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostedDeviceOutBandInfoResponseParams `json:"Response"`
}

func (r *DescribeHostedDeviceOutBandInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostedDeviceOutBandInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOperationResultRequestParams struct {
	// 异步任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeOperationResultRequest struct {
	*tchttp.BaseRequest
	
	// 异步任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeOperationResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOperationResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOperationResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOperationResultResponseParams struct {
	// 任务的整体状态，取值如下：<br>
	// 1：成功<br>
	// 2：失败<br>
	// 3：部分成功，部分失败<br>
	// 4：未完成<br>
	// 5：部分成功，部分未完成<br>
	// 6：部分未完成，部分失败<br>
	// 7：部分未完成，部分失败，部分成功
	TaskStatus *uint64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 各实例对应任务的状态ID
	SubtaskStatusSet []*SubtaskStatus `json:"SubtaskStatusSet,omitempty" name:"SubtaskStatusSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOperationResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOperationResultResponseParams `json:"Response"`
}

func (r *DescribeOperationResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOperationResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOsInfoRequestParams struct {
	// 设备类型代号。 可以从DescribeDeviceClass查询设备类型列表
	DeviceClassCode *string `json:"DeviceClassCode,omitempty" name:"DeviceClassCode"`
}

type DescribeOsInfoRequest struct {
	*tchttp.BaseRequest
	
	// 设备类型代号。 可以从DescribeDeviceClass查询设备类型列表
	DeviceClassCode *string `json:"DeviceClassCode,omitempty" name:"DeviceClassCode"`
}

func (r *DescribeOsInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOsInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceClassCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOsInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOsInfoResponseParams struct {
	// 操作系统信息列表
	OsInfoSet []*OsInfo `json:"OsInfoSet,omitempty" name:"OsInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOsInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOsInfoResponseParams `json:"Response"`
}

func (r *DescribeOsInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOsInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePsaRegulationsRequestParams struct {
	// 数量限制
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 规则ID过滤，支持模糊查询
	PsaIds []*string `json:"PsaIds,omitempty" name:"PsaIds"`

	// 规则别名过滤，支持模糊查询
	PsaNames []*string `json:"PsaNames,omitempty" name:"PsaNames"`

	// 标签过滤
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 排序字段，取值支持：CreateTime
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方式 0:递增(默认) 1:递减
	Order *uint64 `json:"Order,omitempty" name:"Order"`
}

type DescribePsaRegulationsRequest struct {
	*tchttp.BaseRequest
	
	// 数量限制
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 规则ID过滤，支持模糊查询
	PsaIds []*string `json:"PsaIds,omitempty" name:"PsaIds"`

	// 规则别名过滤，支持模糊查询
	PsaNames []*string `json:"PsaNames,omitempty" name:"PsaNames"`

	// 标签过滤
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 排序字段，取值支持：CreateTime
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方式 0:递增(默认) 1:递减
	Order *uint64 `json:"Order,omitempty" name:"Order"`
}

func (r *DescribePsaRegulationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePsaRegulationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "PsaIds")
	delete(f, "PsaNames")
	delete(f, "Tags")
	delete(f, "OrderField")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePsaRegulationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePsaRegulationsResponseParams struct {
	// 返回规则数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 返回规则列表
	PsaRegulations []*PsaRegulation `json:"PsaRegulations,omitempty" name:"PsaRegulations"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePsaRegulationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePsaRegulationsResponseParams `json:"Response"`
}

func (r *DescribePsaRegulationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePsaRegulationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionsRequestParams struct {
	// 地域整型ID，目前黑石可用地域包括：8-北京，4-上海，1-广州， 19-重庆
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
}

type DescribeRegionsRequest struct {
	*tchttp.BaseRequest
	
	// 地域整型ID，目前黑石可用地域包括：8-北京，4-上海，1-广州， 19-重庆
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *DescribeRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionsResponseParams struct {
	// 地域信息
	RegionInfoSet []*RegionInfo `json:"RegionInfoSet,omitempty" name:"RegionInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRegionsResponseParams `json:"Response"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRepairTaskConstantRequestParams struct {

}

type DescribeRepairTaskConstantRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRepairTaskConstantRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRepairTaskConstantRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRepairTaskConstantRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRepairTaskConstantResponseParams struct {
	// 故障类型ID与对应中文名列表
	TaskTypeSet []*TaskType `json:"TaskTypeSet,omitempty" name:"TaskTypeSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRepairTaskConstantResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRepairTaskConstantResponseParams `json:"Response"`
}

func (r *DescribeRepairTaskConstantResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRepairTaskConstantResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskInfoRequestParams struct {
	// 开始位置
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数据条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 时间过滤下限
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 时间过滤上限
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 任务状态ID过滤
	TaskStatus []*uint64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 排序字段，目前支持：CreateTime，AuthTime，EndTime
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方式 0:递增(默认) 1:递减
	Order *uint64 `json:"Order,omitempty" name:"Order"`

	// 任务ID过滤
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// 实例ID过滤
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 实例别名过滤
	Aliases []*string `json:"Aliases,omitempty" name:"Aliases"`

	// 故障类型ID过滤
	TaskTypeIds []*uint64 `json:"TaskTypeIds,omitempty" name:"TaskTypeIds"`
}

type DescribeTaskInfoRequest struct {
	*tchttp.BaseRequest
	
	// 开始位置
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数据条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 时间过滤下限
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 时间过滤上限
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 任务状态ID过滤
	TaskStatus []*uint64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 排序字段，目前支持：CreateTime，AuthTime，EndTime
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方式 0:递增(默认) 1:递减
	Order *uint64 `json:"Order,omitempty" name:"Order"`

	// 任务ID过滤
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// 实例ID过滤
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 实例别名过滤
	Aliases []*string `json:"Aliases,omitempty" name:"Aliases"`

	// 故障类型ID过滤
	TaskTypeIds []*uint64 `json:"TaskTypeIds,omitempty" name:"TaskTypeIds"`
}

func (r *DescribeTaskInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "TaskStatus")
	delete(f, "OrderField")
	delete(f, "Order")
	delete(f, "TaskIds")
	delete(f, "InstanceIds")
	delete(f, "Aliases")
	delete(f, "TaskTypeIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskInfoResponseParams struct {
	// 返回任务总数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 任务信息列表
	TaskInfoSet []*TaskInfo `json:"TaskInfoSet,omitempty" name:"TaskInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskInfoResponseParams `json:"Response"`
}

func (r *DescribeTaskInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskOperationLogRequestParams struct {
	// 维修任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 排序字段，目前支持：OperationTime
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方式 0:递增(默认) 1:递减
	Order *uint64 `json:"Order,omitempty" name:"Order"`
}

type DescribeTaskOperationLogRequest struct {
	*tchttp.BaseRequest
	
	// 维修任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 排序字段，目前支持：OperationTime
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方式 0:递增(默认) 1:递减
	Order *uint64 `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeTaskOperationLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskOperationLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "OrderField")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskOperationLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskOperationLogResponseParams struct {
	// 操作日志
	TaskOperationLogSet []*TaskOperationLog `json:"TaskOperationLogSet,omitempty" name:"TaskOperationLogSet"`

	// 日志条数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskOperationLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskOperationLogResponseParams `json:"Response"`
}

func (r *DescribeTaskOperationLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskOperationLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserCmdTaskInfoRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数量限制
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段，支持： RunBeginTime,RunEndTime,Status
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方式，取值: 1倒序，0顺序；默认倒序
	Order *uint64 `json:"Order,omitempty" name:"Order"`

	// 关键字搜索，可搜索ID或别名，支持模糊搜索
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

type DescribeUserCmdTaskInfoRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数量限制
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段，支持： RunBeginTime,RunEndTime,Status
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方式，取值: 1倒序，0顺序；默认倒序
	Order *uint64 `json:"Order,omitempty" name:"Order"`

	// 关键字搜索，可搜索ID或别名，支持模糊搜索
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *DescribeUserCmdTaskInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserCmdTaskInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "Order")
	delete(f, "SearchKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserCmdTaskInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserCmdTaskInfoResponseParams struct {
	// 返回数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 自定义脚本任务详细信息列表
	UserCmdTaskInfoSet []*UserCmdTaskInfo `json:"UserCmdTaskInfoSet,omitempty" name:"UserCmdTaskInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUserCmdTaskInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserCmdTaskInfoResponseParams `json:"Response"`
}

func (r *DescribeUserCmdTaskInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserCmdTaskInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserCmdTasksRequestParams struct {
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数量限制
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段，支持： RunBeginTime,RunEndTime,InstanceCount,SuccessCount,FailureCount
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方式，取值: 1倒序，0顺序；默认倒序
	Order *uint64 `json:"Order,omitempty" name:"Order"`
}

type DescribeUserCmdTasksRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数量限制
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段，支持： RunBeginTime,RunEndTime,InstanceCount,SuccessCount,FailureCount
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方式，取值: 1倒序，0顺序；默认倒序
	Order *uint64 `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeUserCmdTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserCmdTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserCmdTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserCmdTasksResponseParams struct {
	// 脚本任务信息数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 脚本任务信息列表
	UserCmdTasks []*UserCmdTask `json:"UserCmdTasks,omitempty" name:"UserCmdTasks"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUserCmdTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserCmdTasksResponseParams `json:"Response"`
}

func (r *DescribeUserCmdTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserCmdTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserCmdsRequestParams struct {
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数量限制
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段，支持： OsType,CreateTime,ModifyTime
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方式，取值: 1倒序，0顺序；默认倒序
	Order *uint64 `json:"Order,omitempty" name:"Order"`

	// 关键字搜索，可搜索ID或别名，支持模糊搜索
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// 查询的脚本ID
	CmdId *string `json:"CmdId,omitempty" name:"CmdId"`
}

type DescribeUserCmdsRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数量限制
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段，支持： OsType,CreateTime,ModifyTime
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方式，取值: 1倒序，0顺序；默认倒序
	Order *uint64 `json:"Order,omitempty" name:"Order"`

	// 关键字搜索，可搜索ID或别名，支持模糊搜索
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// 查询的脚本ID
	CmdId *string `json:"CmdId,omitempty" name:"CmdId"`
}

func (r *DescribeUserCmdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserCmdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "Order")
	delete(f, "SearchKey")
	delete(f, "CmdId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserCmdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserCmdsResponseParams struct {
	// 返回数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 脚本信息列表
	UserCmds []*UserCmd `json:"UserCmds,omitempty" name:"UserCmds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUserCmdsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserCmdsResponseParams `json:"Response"`
}

func (r *DescribeUserCmdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserCmdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachCamRoleRequestParams struct {
	// 服务器ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DetachCamRoleRequest struct {
	*tchttp.BaseRequest
	
	// 服务器ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DetachCamRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachCamRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachCamRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachCamRoleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetachCamRoleResponse struct {
	*tchttp.BaseResponse
	Response *DetachCamRoleResponseParams `json:"Response"`
}

func (r *DetachCamRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachCamRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeviceAlias struct {
	// 设备ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 设备别名
	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

type DeviceClass struct {
	// 机型ID
	DeviceClassCode *string `json:"DeviceClassCode,omitempty" name:"DeviceClassCode"`

	// CPU描述
	CpuDescription *string `json:"CpuDescription,omitempty" name:"CpuDescription"`

	// 内存描述
	MemDescription *string `json:"MemDescription,omitempty" name:"MemDescription"`

	// 硬盘描述
	DiskDescription *string `json:"DiskDescription,omitempty" name:"DiskDescription"`

	// 是否支持RAID. 0:不支持; 1:支持
	HaveRaidCard *uint64 `json:"HaveRaidCard,omitempty" name:"HaveRaidCard"`

	// 网卡描述
	NicDescription *string `json:"NicDescription,omitempty" name:"NicDescription"`

	// GPU描述
	GpuDescription *string `json:"GpuDescription,omitempty" name:"GpuDescription"`

	// 单价折扣
	// 注意：此字段可能返回 null，表示取不到有效值。
	Discount *float64 `json:"Discount,omitempty" name:"Discount"`

	// 用户刊例价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPrice *uint64 `json:"UnitPrice,omitempty" name:"UnitPrice"`

	// 实际价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealPrice *uint64 `json:"RealPrice,omitempty" name:"RealPrice"`

	// 官网刊例价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	NormalPrice *uint64 `json:"NormalPrice,omitempty" name:"NormalPrice"`

	// 设备使用场景类型
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 机型系列
	Series *uint64 `json:"Series,omitempty" name:"Series"`

	// cpu的核心数。仅是物理服务器未开启超线程的核心数， 超线程的核心数为Cpu*2
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存容量。单位G
	Mem *uint64 `json:"Mem,omitempty" name:"Mem"`
}

type DeviceClassPartitionInfo struct {
	// RAID类型ID
	RaidId *uint64 `json:"RaidId,omitempty" name:"RaidId"`

	// RAID名称
	Raid *string `json:"Raid,omitempty" name:"Raid"`

	// RAID名称（前台展示用）
	RaidDisplay *string `json:"RaidDisplay,omitempty" name:"RaidDisplay"`

	// 系统盘总大小（单位GiB）
	SystemDiskSize *uint64 `json:"SystemDiskSize,omitempty" name:"SystemDiskSize"`

	// 系统盘/分区默认大小（单位GiB）
	SysRootSpace *uint64 `json:"SysRootSpace,omitempty" name:"SysRootSpace"`

	// 系统盘swap分区默认大小（单位GiB）
	SysSwaporuefiSpace *uint64 `json:"SysSwaporuefiSpace,omitempty" name:"SysSwaporuefiSpace"`

	// 系统盘/usr/local分区默认大小（单位GiB）
	SysUsrlocalSpace *uint64 `json:"SysUsrlocalSpace,omitempty" name:"SysUsrlocalSpace"`

	// 系统盘/data分区默认大小（单位GiB）
	SysDataSpace *uint64 `json:"SysDataSpace,omitempty" name:"SysDataSpace"`

	// 设备是否是uefi启动方式。0:legacy启动; 1:uefi启动
	SysIsUefiType *uint64 `json:"SysIsUefiType,omitempty" name:"SysIsUefiType"`

	// 数据盘总大小
	DataDiskSize *uint64 `json:"DataDiskSize,omitempty" name:"DataDiskSize"`

	// 硬盘列表
	DeviceDiskSizeInfoSet []*DeviceDiskSizeInfo `json:"DeviceDiskSizeInfoSet,omitempty" name:"DeviceDiskSizeInfoSet"`
}

type DeviceDiskSizeInfo struct {
	// 硬盘名称
	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`

	// 硬盘大小（单位GiB）
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

type DeviceHardwareInfo struct {
	// 设备实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 是否自定义机型
	IsElastic *uint64 `json:"IsElastic,omitempty" name:"IsElastic"`

	// 机型计费模式，1 为预付费，2 为后付费
	CpmPayMode *uint64 `json:"CpmPayMode,omitempty" name:"CpmPayMode"`

	// 自定义机型，CPU 型号 ID（非自定义机型返回0）
	CpuId *uint64 `json:"CpuId,omitempty" name:"CpuId"`

	// 自定义机型，内存大小, 单位 GB（非自定义机型返回0）
	Mem *uint64 `json:"Mem,omitempty" name:"Mem"`

	// 是否有 RAID 卡，0：没有 RAID 卡； 1：有 RAID 卡
	ContainRaidCard *uint64 `json:"ContainRaidCard,omitempty" name:"ContainRaidCard"`

	// 自定义机型系统盘类型ID（若没有则返回0）
	SystemDiskTypeId *uint64 `json:"SystemDiskTypeId,omitempty" name:"SystemDiskTypeId"`

	// 自定义机型系统盘数量（若没有则返回0）
	SystemDiskCount *uint64 `json:"SystemDiskCount,omitempty" name:"SystemDiskCount"`

	// 自定义机型数据盘类型 ID（若没有则返回0）
	DataDiskTypeId *uint64 `json:"DataDiskTypeId,omitempty" name:"DataDiskTypeId"`

	// 自定义机型数据盘数量（若没有则返回0）
	DataDiskCount *uint64 `json:"DataDiskCount,omitempty" name:"DataDiskCount"`

	// CPU 型号描述
	CpuDescription *string `json:"CpuDescription,omitempty" name:"CpuDescription"`

	// 内存描述
	MemDescription *string `json:"MemDescription,omitempty" name:"MemDescription"`

	// 磁盘描述
	DiskDescription *string `json:"DiskDescription,omitempty" name:"DiskDescription"`

	// 网卡描述
	NicDescription *string `json:"NicDescription,omitempty" name:"NicDescription"`

	// 是否支持 RAID 的描述
	RaidDescription *string `json:"RaidDescription,omitempty" name:"RaidDescription"`

	// cpu的核心数。仅是物理服务器未开启超线程的核心数， 超线程的核心数为Cpu*2
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// 机型外部代号
	DeviceClassCode *string `json:"DeviceClassCode,omitempty" name:"DeviceClassCode"`
}

type DeviceInfo struct {
	// 设备唯一ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 设备状态ID，取值：<li>1：申领设备中</li><li>2：初始化中</li><li>4：运营中</li><li>7：隔离中</li><li>8：已隔离</li><li>10：解隔离中</li><li>16：故障中</li>
	DeviceStatus *uint64 `json:"DeviceStatus,omitempty" name:"DeviceStatus"`

	// 设备操作状态ID，取值：
	// <li>1：运行中</li><li>2：正在关机</li><li>3：已关机</li><li>5：正在开机</li><li>7：重启中</li><li>9：重装中</li><li>12：绑定EIP</li><li>13：解绑EIP</li><li>14：绑定LB</li><li>15：解绑LB</li><li>19：更换IP中</li><li>20：制作镜像中</li><li>21：制作镜像失败</li><li>23：故障待重装</li><li>24：无备件待退回</li>
	OperateStatus *uint64 `json:"OperateStatus,omitempty" name:"OperateStatus"`

	// 操作系统ID，参考接口[查询操作系统信息(DescribeOsInfo)](https://cloud.tencent.com/document/product/386/32902)
	OsTypeId *uint64 `json:"OsTypeId,omitempty" name:"OsTypeId"`

	// RAID类型ID，参考接口[查询机型RAID方式以及系统盘大小(DescribeDeviceClassPartition)](https://cloud.tencent.com/document/product/386/32910)
	RaidId *uint64 `json:"RaidId,omitempty" name:"RaidId"`

	// 设备别名
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 用户AppId
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 外网IP
	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`

	// 内网IP
	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`

	// 设备交付时间
	DeliverTime *string `json:"DeliverTime,omitempty" name:"DeliverTime"`

	// 设备到期时间
	Deadline *string `json:"Deadline,omitempty" name:"Deadline"`

	// 自动续费标识。0: 不自动续费; 1:自动续费
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 设备类型
	DeviceClassCode *string `json:"DeviceClassCode,omitempty" name:"DeviceClassCode"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 计费模式。1: 预付费; 2: 后付费; 3:预付费转后付费中
	CpmPayMode *uint64 `json:"CpmPayMode,omitempty" name:"CpmPayMode"`

	// 带外IP
	DhcpIp *string `json:"DhcpIp,omitempty" name:"DhcpIp"`

	// 所在私有网络别名
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 所在子网别名
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// 所在私有网络CIDR
	VpcCidrBlock *string `json:"VpcCidrBlock,omitempty" name:"VpcCidrBlock"`

	// 所在子网CIDR
	SubnetCidrBlock *string `json:"SubnetCidrBlock,omitempty" name:"SubnetCidrBlock"`

	// 标识是否是竞价实例。0: 普通设备; 1: 竞价实例设备
	IsLuckyDevice *uint64 `json:"IsLuckyDevice,omitempty" name:"IsLuckyDevice"`

	// 标识机器维保状态。Maintain: 在保;  WillExpire: 即将过保; Expire: 已过保
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaintainStatus *string `json:"MaintainStatus,omitempty" name:"MaintainStatus"`

	// 维保信息描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaintainMessage *string `json:"MaintainMessage,omitempty" name:"MaintainMessage"`
}

type DeviceOperationLog struct {
	// 日志的ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 设备ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 日志对应的操作任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 操作任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 操作任务中文名称
	TaskDescription *string `json:"TaskDescription,omitempty" name:"TaskDescription"`

	// 操作开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 操作结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 操作状态，0: 正在执行中；1：任务成功； 2: 任务失败。
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 操作者
	OpUin *string `json:"OpUin,omitempty" name:"OpUin"`

	// 操作描述
	LogDescription *string `json:"LogDescription,omitempty" name:"LogDescription"`
}

type DevicePartition struct {
	// 系统盘大小
	SystemDiskSize *uint64 `json:"SystemDiskSize,omitempty" name:"SystemDiskSize"`

	// 数据盘大小
	DataDiskSize *uint64 `json:"DataDiskSize,omitempty" name:"DataDiskSize"`

	// 是否兼容Uefi
	SysIsUefiType *bool `json:"SysIsUefiType,omitempty" name:"SysIsUefiType"`

	// root分区大小
	SysRootSpace *uint64 `json:"SysRootSpace,omitempty" name:"SysRootSpace"`

	// Swaporuefi分区大小
	SysSwaporuefiSpace *uint64 `json:"SysSwaporuefiSpace,omitempty" name:"SysSwaporuefiSpace"`

	// Usrlocal分区大小
	SysUsrlocalSpace *uint64 `json:"SysUsrlocalSpace,omitempty" name:"SysUsrlocalSpace"`

	// data分区大小
	SysDataSpace *uint64 `json:"SysDataSpace,omitempty" name:"SysDataSpace"`

	// 硬盘大小详情
	DeviceDiskSizeInfoSet []*DeviceDiskSizeInfo `json:"DeviceDiskSizeInfoSet,omitempty" name:"DeviceDiskSizeInfoSet"`
}

type DevicePositionInfo struct {
	// 设备ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 所在可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 业务IP
	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`

	// 实例别名
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 机架名称
	RckName *string `json:"RckName,omitempty" name:"RckName"`

	// 机位
	PosCode *uint64 `json:"PosCode,omitempty" name:"PosCode"`

	// 交换机名称
	SwitchName *string `json:"SwitchName,omitempty" name:"SwitchName"`

	// 设备交付时间
	DeliverTime *string `json:"DeliverTime,omitempty" name:"DeliverTime"`

	// 过期时间
	Deadline *string `json:"Deadline,omitempty" name:"Deadline"`
}

type DevicePriceInfo struct {
	// 物理机ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 设备型号
	DeviceClassCode *string `json:"DeviceClassCode,omitempty" name:"DeviceClassCode"`

	// 是否是弹性机型，1：是，0：否
	IsElastic *uint64 `json:"IsElastic,omitempty" name:"IsElastic"`

	// 付费模式ID, 1:预付费; 2:后付费; 3:预付费转后付费中
	CpmPayMode *uint64 `json:"CpmPayMode,omitempty" name:"CpmPayMode"`

	// Cpu信息描述
	CpuDescription *string `json:"CpuDescription,omitempty" name:"CpuDescription"`

	// 内存信息描述
	MemDescription *string `json:"MemDescription,omitempty" name:"MemDescription"`

	// 硬盘信息描述
	DiskDescription *string `json:"DiskDescription,omitempty" name:"DiskDescription"`

	// 网卡信息描述
	NicDescription *string `json:"NicDescription,omitempty" name:"NicDescription"`

	// Gpu信息描述
	GpuDescription *string `json:"GpuDescription,omitempty" name:"GpuDescription"`

	// Raid信息描述
	RaidDescription *string `json:"RaidDescription,omitempty" name:"RaidDescription"`

	// 客户的单价
	Price *uint64 `json:"Price,omitempty" name:"Price"`

	// 刊例单价
	NormalPrice *uint64 `json:"NormalPrice,omitempty" name:"NormalPrice"`

	// 原价
	TotalCost *uint64 `json:"TotalCost,omitempty" name:"TotalCost"`

	// 折扣价
	RealTotalCost *uint64 `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 计费时长
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 计费时长单位, M:按月计费; D:按天计费
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 商品数量
	GoodsCount *uint64 `json:"GoodsCount,omitempty" name:"GoodsCount"`
}

type DiskInfo struct {
	// 磁盘ID
	DiskTypeId *uint64 `json:"DiskTypeId,omitempty" name:"DiskTypeId"`

	// 磁盘的容量，单位为G
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 磁盘信息描述
	DiskDescription *string `json:"DiskDescription,omitempty" name:"DiskDescription"`
}

type FailedTaskInfo struct {
	// 运行脚本的设备ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 失败原因
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`
}

type HostedDeviceOutBandInfo struct {
	// 物理机ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 带外IP
	OutBandIp *string `json:"OutBandIp,omitempty" name:"OutBandIp"`

	// VPN的IP
	VpnIp *string `json:"VpnIp,omitempty" name:"VpnIp"`

	// VPN的端口
	VpnPort *uint64 `json:"VpnPort,omitempty" name:"VpnPort"`
}

// Predefined struct for user
type ModifyCustomImageAttributeRequestParams struct {
	// 镜像ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 设置新的镜像名
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 设置新的镜像描述
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`
}

type ModifyCustomImageAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 镜像ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 设置新的镜像名
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 设置新的镜像描述
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`
}

func (r *ModifyCustomImageAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomImageAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageId")
	delete(f, "ImageName")
	delete(f, "ImageDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomImageAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomImageAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCustomImageAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCustomImageAttributeResponseParams `json:"Response"`
}

func (r *ModifyCustomImageAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomImageAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDeviceAliasesRequestParams struct {
	// 需要改名的设备与别名列表
	DeviceAliases []*DeviceAlias `json:"DeviceAliases,omitempty" name:"DeviceAliases"`
}

type ModifyDeviceAliasesRequest struct {
	*tchttp.BaseRequest
	
	// 需要改名的设备与别名列表
	DeviceAliases []*DeviceAlias `json:"DeviceAliases,omitempty" name:"DeviceAliases"`
}

func (r *ModifyDeviceAliasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceAliasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceAliases")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDeviceAliasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDeviceAliasesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDeviceAliasesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDeviceAliasesResponseParams `json:"Response"`
}

func (r *ModifyDeviceAliasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceAliasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDeviceAutoRenewFlagRequestParams struct {
	// 自动续费标志位。0: 不自动续费; 1: 自动续费
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 需要修改的设备ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type ModifyDeviceAutoRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// 自动续费标志位。0: 不自动续费; 1: 自动续费
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 需要修改的设备ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *ModifyDeviceAutoRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceAutoRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoRenewFlag")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDeviceAutoRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDeviceAutoRenewFlagResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDeviceAutoRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDeviceAutoRenewFlagResponseParams `json:"Response"`
}

func (r *ModifyDeviceAutoRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceAutoRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLanIpRequestParams struct {
	// 物理机ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 指定新VPC
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 指定新子网
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 指定新内网IP
	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`

	// 是否需要重启机器，取值 1(需要) 0(不需要)，默认取值0
	RebootDevice *uint64 `json:"RebootDevice,omitempty" name:"RebootDevice"`
}

type ModifyLanIpRequest struct {
	*tchttp.BaseRequest
	
	// 物理机ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 指定新VPC
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 指定新子网
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 指定新内网IP
	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`

	// 是否需要重启机器，取值 1(需要) 0(不需要)，默认取值0
	RebootDevice *uint64 `json:"RebootDevice,omitempty" name:"RebootDevice"`
}

func (r *ModifyLanIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLanIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "LanIp")
	delete(f, "RebootDevice")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLanIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLanIpResponseParams struct {
	// 黑石异步任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyLanIpResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLanIpResponseParams `json:"Response"`
}

func (r *ModifyLanIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLanIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPayModePre2PostRequestParams struct {
	// 需要修改的设备ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type ModifyPayModePre2PostRequest struct {
	*tchttp.BaseRequest
	
	// 需要修改的设备ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *ModifyPayModePre2PostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPayModePre2PostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPayModePre2PostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPayModePre2PostResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPayModePre2PostResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPayModePre2PostResponseParams `json:"Response"`
}

func (r *ModifyPayModePre2PostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPayModePre2PostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPsaRegulationRequestParams struct {
	// 预授权规则ID
	PsaId *string `json:"PsaId,omitempty" name:"PsaId"`

	// 预授权规则别名
	PsaName *string `json:"PsaName,omitempty" name:"PsaName"`

	// 维修中的实例上限
	RepairLimit *uint64 `json:"RepairLimit,omitempty" name:"RepairLimit"`

	// 预授权规则备注
	PsaDescription *string `json:"PsaDescription,omitempty" name:"PsaDescription"`

	// 预授权规则关联故障类型ID列表
	TaskTypeIds []*uint64 `json:"TaskTypeIds,omitempty" name:"TaskTypeIds"`
}

type ModifyPsaRegulationRequest struct {
	*tchttp.BaseRequest
	
	// 预授权规则ID
	PsaId *string `json:"PsaId,omitempty" name:"PsaId"`

	// 预授权规则别名
	PsaName *string `json:"PsaName,omitempty" name:"PsaName"`

	// 维修中的实例上限
	RepairLimit *uint64 `json:"RepairLimit,omitempty" name:"RepairLimit"`

	// 预授权规则备注
	PsaDescription *string `json:"PsaDescription,omitempty" name:"PsaDescription"`

	// 预授权规则关联故障类型ID列表
	TaskTypeIds []*uint64 `json:"TaskTypeIds,omitempty" name:"TaskTypeIds"`
}

func (r *ModifyPsaRegulationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPsaRegulationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PsaId")
	delete(f, "PsaName")
	delete(f, "RepairLimit")
	delete(f, "PsaDescription")
	delete(f, "TaskTypeIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPsaRegulationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPsaRegulationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPsaRegulationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPsaRegulationResponseParams `json:"Response"`
}

func (r *ModifyPsaRegulationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPsaRegulationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserCmdRequestParams struct {
	// 待修改的脚本ID
	CmdId *string `json:"CmdId,omitempty" name:"CmdId"`

	// 待修改的脚本名称
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 脚本适用的操作系统类型
	OsType *string `json:"OsType,omitempty" name:"OsType"`

	// 待修改的脚本内容，必须经过base64编码
	Content *string `json:"Content,omitempty" name:"Content"`
}

type ModifyUserCmdRequest struct {
	*tchttp.BaseRequest
	
	// 待修改的脚本ID
	CmdId *string `json:"CmdId,omitempty" name:"CmdId"`

	// 待修改的脚本名称
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 脚本适用的操作系统类型
	OsType *string `json:"OsType,omitempty" name:"OsType"`

	// 待修改的脚本内容，必须经过base64编码
	Content *string `json:"Content,omitempty" name:"Content"`
}

func (r *ModifyUserCmdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserCmdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CmdId")
	delete(f, "Alias")
	delete(f, "OsType")
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserCmdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserCmdResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyUserCmdResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserCmdResponseParams `json:"Response"`
}

func (r *ModifyUserCmdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserCmdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OfflineDevicesRequestParams struct {
	// 需要退还的物理机ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type OfflineDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 需要退还的物理机ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *OfflineDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OfflineDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OfflineDevicesResponseParams struct {
	// 黑石异步任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type OfflineDevicesResponse struct {
	*tchttp.BaseResponse
	Response *OfflineDevicesResponseParams `json:"Response"`
}

func (r *OfflineDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OsInfo struct {
	// 操作系统ID
	OsTypeId *uint64 `json:"OsTypeId,omitempty" name:"OsTypeId"`

	// 操作系统名称
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// 操作系统名称描述
	OsDescription *string `json:"OsDescription,omitempty" name:"OsDescription"`

	// 操作系统英文名称
	OsEnglishDescription *string `json:"OsEnglishDescription,omitempty" name:"OsEnglishDescription"`

	// 操作系统的分类，如CentOs Debian
	OsClass *string `json:"OsClass,omitempty" name:"OsClass"`

	// 标识镜像分类。public:公共镜像; private: 专属镜像
	ImageTag *string `json:"ImageTag,omitempty" name:"ImageTag"`

	// 操作系统，ext4文件下所支持的最大的磁盘大小。单位为T
	MaxPartitionSize *uint64 `json:"MaxPartitionSize,omitempty" name:"MaxPartitionSize"`

	// 黑石版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OsMinorVersion *string `json:"OsMinorVersion,omitempty" name:"OsMinorVersion"`

	// 黑石版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	OsMinorClass *string `json:"OsMinorClass,omitempty" name:"OsMinorClass"`
}

type PartitionInfo struct {
	// 分区名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分区大小
	Size *uint64 `json:"Size,omitempty" name:"Size"`
}

type PsaRegulation struct {
	// 规则ID
	PsaId *string `json:"PsaId,omitempty" name:"PsaId"`

	// 规则别名
	PsaName *string `json:"PsaName,omitempty" name:"PsaName"`

	// 关联标签数量
	TagCount *uint64 `json:"TagCount,omitempty" name:"TagCount"`

	// 关联实例数量
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 故障实例数量
	RepairCount *uint64 `json:"RepairCount,omitempty" name:"RepairCount"`

	// 故障实例上限
	RepairLimit *uint64 `json:"RepairLimit,omitempty" name:"RepairLimit"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 规则备注
	PsaDescription *string `json:"PsaDescription,omitempty" name:"PsaDescription"`

	// 关联标签
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 关联故障类型id
	TaskTypeIds []*uint64 `json:"TaskTypeIds,omitempty" name:"TaskTypeIds"`
}

// Predefined struct for user
type RebootDevicesRequestParams struct {
	// 需要重启的设备ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type RebootDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 需要重启的设备ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *RebootDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RebootDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RebootDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RebootDevicesResponseParams struct {
	// 异步任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RebootDevicesResponse struct {
	*tchttp.BaseResponse
	Response *RebootDevicesResponseParams `json:"Response"`
}

func (r *RebootDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RebootDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecoverDevicesRequestParams struct {
	// 需要恢复的物理机ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type RecoverDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 需要恢复的物理机ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *RecoverDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecoverDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecoverDevicesResponseParams struct {
	// 黑石异步任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RecoverDevicesResponse struct {
	*tchttp.BaseResponse
	Response *RecoverDevicesResponseParams `json:"Response"`
}

func (r *RecoverDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {
	// 地域ID
	Region *string `json:"Region,omitempty" name:"Region"`

	// 地域整型ID
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// 地域描述
	RegionDescription *string `json:"RegionDescription,omitempty" name:"RegionDescription"`

	// 该地域下的可用区信息
	ZoneInfoSet []*ZoneInfo `json:"ZoneInfoSet,omitempty" name:"ZoneInfoSet"`
}

// Predefined struct for user
type ReloadDeviceOsRequestParams struct {
	// 设备的唯一ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 密码。 用户设置的linux root或Windows Administrator密码。密码校验规则: <li> Windows机器密码需12到16位，至少包括三项 `[a-z]`,`[A-Z]`,`[0-9]`和`[()`'`~!@#$%^&*-+=_`|`{}[]:;'<>,.?/]`的特殊符号, 密码不能包含Administrator(不区分大小写); <li> Linux机器密码需8到16位，至少包括两项`[a-z,A-Z]`,`[0-9]`和`[()`'`~!@#$%^&*-+=_`|`{}[]:;'<>,.?/]`的特殊符号
	Password *string `json:"Password,omitempty" name:"Password"`

	// 操作系统类型ID。通过接口[查询操作系统信息(DescribeOsInfo)](https://cloud.tencent.com/document/api/386/32902)获取操作系统信息
	OsTypeId *uint64 `json:"OsTypeId,omitempty" name:"OsTypeId"`

	// RAID类型ID。通过接口[查询机型RAID方式以及系统盘大小(DescribeDeviceClassPartition)](https://cloud.tencent.com/document/api/386/32910)获取RAID信息
	RaidId *uint64 `json:"RaidId,omitempty" name:"RaidId"`

	// 是否格式化数据盘。0: 不格式化（默认值）；1：格式化
	IsZoning *uint64 `json:"IsZoning,omitempty" name:"IsZoning"`

	// 系统盘根分区大小，默认是10G。系统盘的大小参考接口[查询机型RAID方式以及系统盘大小(DescribeDeviceClassPartition)](https://cloud.tencent.com/document/api/386/32910)
	SysRootSpace *uint64 `json:"SysRootSpace,omitempty" name:"SysRootSpace"`

	// 系统盘swap分区或/boot/efi分区的大小。若是uefi启动的机器，分区为/boot/efi ,且此值是默认是2G。普通机器为swap分区，可以不指定此分区。机型是否是uefi启动，参考接口[查询设备型号(DescribeDeviceClass)](https://cloud.tencent.com/document/api/386/32911)
	SysSwaporuefiSpace *uint64 `json:"SysSwaporuefiSpace,omitempty" name:"SysSwaporuefiSpace"`

	// /usr/local分区大小
	SysUsrlocalSpace *uint64 `json:"SysUsrlocalSpace,omitempty" name:"SysUsrlocalSpace"`

	// 重装到新的私有网络的ID。如果改变VPC子网，则要求与SubnetId同时传参，否则可不填
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 重装到新的子网的ID。如果改变VPC子网，则要求与VpcId同时传参，否则可不填
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 重装指定IP地址
	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`

	// 指定是否开启超线程。 0：关闭超线程；1：开启超线程（默认值）
	HyperThreading *uint64 `json:"HyperThreading,omitempty" name:"HyperThreading"`

	// 自定义镜像ID。传此字段则用自定义镜像重装
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 指定数据盘的文件系统格式，当前支持 EXT4和XFS选项， 默认为EXT4。 参数适用于数据盘和Linux， 且在IsZoning为1时生效
	FileSystem *string `json:"FileSystem,omitempty" name:"FileSystem"`

	// 是否安装安全Agent，取值：1(安装) 0(不安装)，默认取值0
	NeedSecurityAgent *uint64 `json:"NeedSecurityAgent,omitempty" name:"NeedSecurityAgent"`

	// 是否安装监控Agent，取值：1(安装) 0(不安装)，默认取值0
	NeedMonitorAgent *uint64 `json:"NeedMonitorAgent,omitempty" name:"NeedMonitorAgent"`

	// 是否安装EMR Agent，取值：1(安装) 0(不安装)，默认取值0
	NeedEMRAgent *uint64 `json:"NeedEMRAgent,omitempty" name:"NeedEMRAgent"`

	// 是否安装EMR软件包，取值：1(安装) 0(不安装)，默认取值0
	NeedEMRSoftware *uint64 `json:"NeedEMRSoftware,omitempty" name:"NeedEMRSoftware"`

	// 是否保留安全组配置，取值：1(保留) 0(不保留)，默认取值0
	ReserveSgConfig *uint64 `json:"ReserveSgConfig,omitempty" name:"ReserveSgConfig"`

	// /data分区大小，可不填。除root、swap、usr/local的剩余空间会自动分配到data分区
	SysDataSpace *uint64 `json:"SysDataSpace,omitempty" name:"SysDataSpace"`
}

type ReloadDeviceOsRequest struct {
	*tchttp.BaseRequest
	
	// 设备的唯一ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 密码。 用户设置的linux root或Windows Administrator密码。密码校验规则: <li> Windows机器密码需12到16位，至少包括三项 `[a-z]`,`[A-Z]`,`[0-9]`和`[()`'`~!@#$%^&*-+=_`|`{}[]:;'<>,.?/]`的特殊符号, 密码不能包含Administrator(不区分大小写); <li> Linux机器密码需8到16位，至少包括两项`[a-z,A-Z]`,`[0-9]`和`[()`'`~!@#$%^&*-+=_`|`{}[]:;'<>,.?/]`的特殊符号
	Password *string `json:"Password,omitempty" name:"Password"`

	// 操作系统类型ID。通过接口[查询操作系统信息(DescribeOsInfo)](https://cloud.tencent.com/document/api/386/32902)获取操作系统信息
	OsTypeId *uint64 `json:"OsTypeId,omitempty" name:"OsTypeId"`

	// RAID类型ID。通过接口[查询机型RAID方式以及系统盘大小(DescribeDeviceClassPartition)](https://cloud.tencent.com/document/api/386/32910)获取RAID信息
	RaidId *uint64 `json:"RaidId,omitempty" name:"RaidId"`

	// 是否格式化数据盘。0: 不格式化（默认值）；1：格式化
	IsZoning *uint64 `json:"IsZoning,omitempty" name:"IsZoning"`

	// 系统盘根分区大小，默认是10G。系统盘的大小参考接口[查询机型RAID方式以及系统盘大小(DescribeDeviceClassPartition)](https://cloud.tencent.com/document/api/386/32910)
	SysRootSpace *uint64 `json:"SysRootSpace,omitempty" name:"SysRootSpace"`

	// 系统盘swap分区或/boot/efi分区的大小。若是uefi启动的机器，分区为/boot/efi ,且此值是默认是2G。普通机器为swap分区，可以不指定此分区。机型是否是uefi启动，参考接口[查询设备型号(DescribeDeviceClass)](https://cloud.tencent.com/document/api/386/32911)
	SysSwaporuefiSpace *uint64 `json:"SysSwaporuefiSpace,omitempty" name:"SysSwaporuefiSpace"`

	// /usr/local分区大小
	SysUsrlocalSpace *uint64 `json:"SysUsrlocalSpace,omitempty" name:"SysUsrlocalSpace"`

	// 重装到新的私有网络的ID。如果改变VPC子网，则要求与SubnetId同时传参，否则可不填
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 重装到新的子网的ID。如果改变VPC子网，则要求与VpcId同时传参，否则可不填
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 重装指定IP地址
	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`

	// 指定是否开启超线程。 0：关闭超线程；1：开启超线程（默认值）
	HyperThreading *uint64 `json:"HyperThreading,omitempty" name:"HyperThreading"`

	// 自定义镜像ID。传此字段则用自定义镜像重装
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 指定数据盘的文件系统格式，当前支持 EXT4和XFS选项， 默认为EXT4。 参数适用于数据盘和Linux， 且在IsZoning为1时生效
	FileSystem *string `json:"FileSystem,omitempty" name:"FileSystem"`

	// 是否安装安全Agent，取值：1(安装) 0(不安装)，默认取值0
	NeedSecurityAgent *uint64 `json:"NeedSecurityAgent,omitempty" name:"NeedSecurityAgent"`

	// 是否安装监控Agent，取值：1(安装) 0(不安装)，默认取值0
	NeedMonitorAgent *uint64 `json:"NeedMonitorAgent,omitempty" name:"NeedMonitorAgent"`

	// 是否安装EMR Agent，取值：1(安装) 0(不安装)，默认取值0
	NeedEMRAgent *uint64 `json:"NeedEMRAgent,omitempty" name:"NeedEMRAgent"`

	// 是否安装EMR软件包，取值：1(安装) 0(不安装)，默认取值0
	NeedEMRSoftware *uint64 `json:"NeedEMRSoftware,omitempty" name:"NeedEMRSoftware"`

	// 是否保留安全组配置，取值：1(保留) 0(不保留)，默认取值0
	ReserveSgConfig *uint64 `json:"ReserveSgConfig,omitempty" name:"ReserveSgConfig"`

	// /data分区大小，可不填。除root、swap、usr/local的剩余空间会自动分配到data分区
	SysDataSpace *uint64 `json:"SysDataSpace,omitempty" name:"SysDataSpace"`
}

func (r *ReloadDeviceOsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReloadDeviceOsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Password")
	delete(f, "OsTypeId")
	delete(f, "RaidId")
	delete(f, "IsZoning")
	delete(f, "SysRootSpace")
	delete(f, "SysSwaporuefiSpace")
	delete(f, "SysUsrlocalSpace")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "LanIp")
	delete(f, "HyperThreading")
	delete(f, "ImageId")
	delete(f, "FileSystem")
	delete(f, "NeedSecurityAgent")
	delete(f, "NeedMonitorAgent")
	delete(f, "NeedEMRAgent")
	delete(f, "NeedEMRSoftware")
	delete(f, "ReserveSgConfig")
	delete(f, "SysDataSpace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReloadDeviceOsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReloadDeviceOsResponseParams struct {
	// 黑石异步任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReloadDeviceOsResponse struct {
	*tchttp.BaseResponse
	Response *ReloadDeviceOsResponseParams `json:"Response"`
}

func (r *ReloadDeviceOsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReloadDeviceOsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RepairTaskControlRequestParams struct {
	// 维修任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 操作
	Operate *string `json:"Operate,omitempty" name:"Operate"`

	// 需要重新维修操作的备注信息，可提供返场维修原因，以便驻场快速针对问题定位解决。
	OperateRemark *string `json:"OperateRemark,omitempty" name:"OperateRemark"`
}

type RepairTaskControlRequest struct {
	*tchttp.BaseRequest
	
	// 维修任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 操作
	Operate *string `json:"Operate,omitempty" name:"Operate"`

	// 需要重新维修操作的备注信息，可提供返场维修原因，以便驻场快速针对问题定位解决。
	OperateRemark *string `json:"OperateRemark,omitempty" name:"OperateRemark"`
}

func (r *RepairTaskControlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RepairTaskControlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Operate")
	delete(f, "OperateRemark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RepairTaskControlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RepairTaskControlResponseParams struct {
	// 出参TaskId是黑石异步任务ID，不同于入参TaskId字段。
	// 此字段可作为DescriptionOperationResult查询异步任务状态接口的入参，查询异步任务执行结果。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RepairTaskControlResponse struct {
	*tchttp.BaseResponse
	Response *RepairTaskControlResponseParams `json:"Response"`
}

func (r *RepairTaskControlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RepairTaskControlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetDevicePasswordRequestParams struct {
	// 需要重置密码的服务器ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 新密码
	Password *string `json:"Password,omitempty" name:"Password"`
}

type ResetDevicePasswordRequest struct {
	*tchttp.BaseRequest
	
	// 需要重置密码的服务器ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 新密码
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *ResetDevicePasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetDevicePasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetDevicePasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetDevicePasswordResponseParams struct {
	// 黑石异步任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResetDevicePasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetDevicePasswordResponseParams `json:"Response"`
}

func (r *ResetDevicePasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetDevicePasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReturnDevicesRequestParams struct {
	// 需要退还的物理机ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type ReturnDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 需要退还的物理机ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *ReturnDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReturnDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReturnDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReturnDevicesResponseParams struct {
	// 黑石异步任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReturnDevicesResponse struct {
	*tchttp.BaseResponse
	Response *ReturnDevicesResponseParams `json:"Response"`
}

func (r *ReturnDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReturnDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunUserCmdRequestParams struct {
	// 自定义脚本ID
	CmdId *string `json:"CmdId,omitempty" name:"CmdId"`

	// 执行脚本机器的用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 执行脚本机器的用户名的密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 执行脚本的服务器实例
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 执行脚本的参数，必须经过base64编码
	CmdParam *string `json:"CmdParam,omitempty" name:"CmdParam"`
}

type RunUserCmdRequest struct {
	*tchttp.BaseRequest
	
	// 自定义脚本ID
	CmdId *string `json:"CmdId,omitempty" name:"CmdId"`

	// 执行脚本机器的用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 执行脚本机器的用户名的密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 执行脚本的服务器实例
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 执行脚本的参数，必须经过base64编码
	CmdParam *string `json:"CmdParam,omitempty" name:"CmdParam"`
}

func (r *RunUserCmdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunUserCmdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CmdId")
	delete(f, "UserName")
	delete(f, "Password")
	delete(f, "InstanceIds")
	delete(f, "CmdParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunUserCmdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunUserCmdResponseParams struct {
	// 运行成功的任务信息列表
	SuccessTaskInfoSet []*SuccessTaskInfo `json:"SuccessTaskInfoSet,omitempty" name:"SuccessTaskInfoSet"`

	// 运行失败的任务信息列表
	FailedTaskInfoSet []*FailedTaskInfo `json:"FailedTaskInfoSet,omitempty" name:"FailedTaskInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RunUserCmdResponse struct {
	*tchttp.BaseResponse
	Response *RunUserCmdResponseParams `json:"Response"`
}

func (r *RunUserCmdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunUserCmdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetOutBandVpnAuthPasswordRequestParams struct {
	// 设置的Vpn认证密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 操作字段，取值为：Create（创建）或Update（修改）
	Operate *string `json:"Operate,omitempty" name:"Operate"`
}

type SetOutBandVpnAuthPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 设置的Vpn认证密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 操作字段，取值为：Create（创建）或Update（修改）
	Operate *string `json:"Operate,omitempty" name:"Operate"`
}

func (r *SetOutBandVpnAuthPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetOutBandVpnAuthPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Password")
	delete(f, "Operate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetOutBandVpnAuthPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetOutBandVpnAuthPasswordResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetOutBandVpnAuthPasswordResponse struct {
	*tchttp.BaseResponse
	Response *SetOutBandVpnAuthPasswordResponseParams `json:"Response"`
}

func (r *SetOutBandVpnAuthPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetOutBandVpnAuthPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ShutdownDevicesRequestParams struct {
	// 需要关闭的设备ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type ShutdownDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 需要关闭的设备ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *ShutdownDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ShutdownDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ShutdownDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ShutdownDevicesResponseParams struct {
	// 异步任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ShutdownDevicesResponse struct {
	*tchttp.BaseResponse
	Response *ShutdownDevicesResponseParams `json:"Response"`
}

func (r *ShutdownDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ShutdownDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartDevicesRequestParams struct {
	// 需要开机的设备ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type StartDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 需要开机的设备ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *StartDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartDevicesResponseParams struct {
	// 异步任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StartDevicesResponse struct {
	*tchttp.BaseResponse
	Response *StartDevicesResponseParams `json:"Response"`
}

func (r *StartDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubtaskStatus struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例ID对应任务的状态，取值如下：<br>
	// 1：成功<br>
	// 2：失败<br>
	// 3：部分成功，部分失败<br>
	// 4：未完成<br>
	// 5：部分成功，部分未完成<br>
	// 6：部分未完成，部分失败<br>
	// 7：部分未完成，部分失败，部分成功
	TaskStatus *uint64 `json:"TaskStatus,omitempty" name:"TaskStatus"`
}

type SuccessTaskInfo struct {
	// 运行脚本的设备ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 黑石异步任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 黑石自定义脚本运行任务ID
	CmdTaskId *string `json:"CmdTaskId,omitempty" name:"CmdTaskId"`
}

type Tag struct {
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签键对应的值
	TagValues []*string `json:"TagValues,omitempty" name:"TagValues"`
}

type TaskInfo struct {
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 主机id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主机别名
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 故障类型id
	TaskTypeId *uint64 `json:"TaskTypeId,omitempty" name:"TaskTypeId"`

	// 任务状态id
	TaskStatus *uint64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 授权时间
	AuthTime *string `json:"AuthTime,omitempty" name:"AuthTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 任务详情
	TaskDetail *string `json:"TaskDetail,omitempty" name:"TaskDetail"`

	// 设备状态
	DeviceStatus *uint64 `json:"DeviceStatus,omitempty" name:"DeviceStatus"`

	// 设备操作状态
	OperateStatus *uint64 `json:"OperateStatus,omitempty" name:"OperateStatus"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 所属网络
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 所在子网
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 子网名
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// VPC名
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// VpcCidrBlock
	VpcCidrBlock *string `json:"VpcCidrBlock,omitempty" name:"VpcCidrBlock"`

	// SubnetCidrBlock
	SubnetCidrBlock *string `json:"SubnetCidrBlock,omitempty" name:"SubnetCidrBlock"`

	// 公网ip
	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`

	// 内网IP
	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`

	// 管理IP
	MgtIp *string `json:"MgtIp,omitempty" name:"MgtIp"`

	// 故障类中文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeName *string `json:"TaskTypeName,omitempty" name:"TaskTypeName"`

	// 故障类型，取值：unconfirmed (不明确故障)；redundancy (有冗余故障)；nonredundancy (无冗余故障)
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskSubType *string `json:"TaskSubType,omitempty" name:"TaskSubType"`
}

type TaskOperationLog struct {
	// 操作步骤
	TaskStep *string `json:"TaskStep,omitempty" name:"TaskStep"`

	// 操作人
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 操作描述
	OperationDetail *string `json:"OperationDetail,omitempty" name:"OperationDetail"`

	// 操作时间
	OperationTime *string `json:"OperationTime,omitempty" name:"OperationTime"`
}

type TaskType struct {
	// 故障类ID
	TypeId *uint64 `json:"TypeId,omitempty" name:"TypeId"`

	// 故障类中文名
	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`

	// 故障类型父类
	TaskSubType *string `json:"TaskSubType,omitempty" name:"TaskSubType"`
}

// Predefined struct for user
type UnbindPsaTagRequestParams struct {
	// 预授权规则ID
	PsaId *string `json:"PsaId,omitempty" name:"PsaId"`

	// 需要解绑的标签key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 需要解绑的标签value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type UnbindPsaTagRequest struct {
	*tchttp.BaseRequest
	
	// 预授权规则ID
	PsaId *string `json:"PsaId,omitempty" name:"PsaId"`

	// 需要解绑的标签key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 需要解绑的标签value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

func (r *UnbindPsaTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindPsaTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PsaId")
	delete(f, "TagKey")
	delete(f, "TagValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindPsaTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindPsaTagResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnbindPsaTagResponse struct {
	*tchttp.BaseResponse
	Response *UnbindPsaTagResponseParams `json:"Response"`
}

func (r *UnbindPsaTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindPsaTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserCmd struct {
	// 用户自定义脚本名
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// AppId
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 脚本自增ID
	AutoId *uint64 `json:"AutoId,omitempty" name:"AutoId"`

	// 脚本ID
	CmdId *string `json:"CmdId,omitempty" name:"CmdId"`

	// 脚本内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 命令适用的操作系统类型
	OsType *string `json:"OsType,omitempty" name:"OsType"`
}

type UserCmdTask struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务状态ID，取值: -1(进行中) 0(结束)
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 脚本名称
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 脚本ID
	CmdId *string `json:"CmdId,omitempty" name:"CmdId"`

	// 运行实例数量
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 运行成功数量
	SuccessCount *uint64 `json:"SuccessCount,omitempty" name:"SuccessCount"`

	// 运行失败数量
	FailureCount *uint64 `json:"FailureCount,omitempty" name:"FailureCount"`

	// 执行开始时间
	RunBeginTime *string `json:"RunBeginTime,omitempty" name:"RunBeginTime"`

	// 执行结束时间
	RunEndTime *string `json:"RunEndTime,omitempty" name:"RunEndTime"`
}

type UserCmdTaskInfo struct {
	// 自动编号，可忽略
	AutoId *uint64 `json:"AutoId,omitempty" name:"AutoId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务开始时间
	RunBeginTime *string `json:"RunBeginTime,omitempty" name:"RunBeginTime"`

	// 任务结束时间
	RunEndTime *string `json:"RunEndTime,omitempty" name:"RunEndTime"`

	// 任务状态ID，取值为 -1：进行中；0：成功；>0：失败错误码
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 设备别名
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 设备ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 私有网络名
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 私有网络整型ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络Cidr
	VpcCidrBlock *string `json:"VpcCidrBlock,omitempty" name:"VpcCidrBlock"`

	// 子网名
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 子网Cidr
	SubnetCidrBlock *string `json:"SubnetCidrBlock,omitempty" name:"SubnetCidrBlock"`

	// 内网IP
	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`

	// 脚本内容，base64编码后的值
	CmdContent *string `json:"CmdContent,omitempty" name:"CmdContent"`

	// 脚本参数，base64编码后的值
	CmdParam *string `json:"CmdParam,omitempty" name:"CmdParam"`

	// 脚本执行结果，base64编码后的值
	CmdResult *string `json:"CmdResult,omitempty" name:"CmdResult"`

	// 用户AppId
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 用户执行脚本结束退出的返回值，没有返回值为-1
	LastShellExit *int64 `json:"LastShellExit,omitempty" name:"LastShellExit"`
}

type ZoneInfo struct {
	// 可用区ID
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 可用区整型ID
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 可用区描述
	ZoneDescription *string `json:"ZoneDescription,omitempty" name:"ZoneDescription"`
}