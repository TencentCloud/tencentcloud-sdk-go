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

package v20180228

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AccountStatistics struct {
	// 用户名。
	Username *string `json:"Username,omitnil" name:"Username"`

	// 主机数量。
	MachineNum *uint64 `json:"MachineNum,omitnil" name:"MachineNum"`
}

type AlarmInfo struct {
	// 该节点关联的告警，告警的table_name+id（t1:id1,t2:id2,...)
	AlarmId *string `json:"AlarmId,omitnil" name:"AlarmId"`

	// 告警事件表状态，当该节点为告警点时生效
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

type AssetAppBaseInfo struct {
	// 主机内网IP
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// 主机名称
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 主机外网IP
	MachineWanIp *string `json:"MachineWanIp,omitnil" name:"MachineWanIp"`

	// 主机Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 主机Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机业务组ID
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 主机标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*MachineTag `json:"Tag,omitnil" name:"Tag"`

	// 应用名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 应用类型	
	// 1: 运维
	// 2 : 数据库
	// 3 : 安全
	// 4 : 可疑应用
	// 5 : 系统架构
	// 6 : 系统应用
	// 7 : WEB服务
	// 99: 其他
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 二进制路径
	BinPath *string `json:"BinPath,omitnil" name:"BinPath"`

	// 操作系统信息
	OsInfo *string `json:"OsInfo,omitnil" name:"OsInfo"`

	// 关联进程数
	ProcessCount *uint64 `json:"ProcessCount,omitnil" name:"ProcessCount"`

	// 应用描述
	Desc *string `json:"Desc,omitnil" name:"Desc"`

	// 版本号
	Version *string `json:"Version,omitnil" name:"Version"`

	// 配置文件路径
	ConfigPath *string `json:"ConfigPath,omitnil" name:"ConfigPath"`

	// 首次采集时间
	FirstTime *string `json:"FirstTime,omitnil" name:"FirstTime"`

	// 数据更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 是否新增[0:否|1:是]
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNew *int64 `json:"IsNew,omitnil" name:"IsNew"`

	// 附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`
}

type AssetAppProcessInfo struct {
	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 进程状态
	Status *string `json:"Status,omitnil" name:"Status"`

	// 进程版本
	Version *string `json:"Version,omitnil" name:"Version"`

	// 路径
	Path *string `json:"Path,omitnil" name:"Path"`

	// 用户
	User *string `json:"User,omitnil" name:"User"`

	// 启动时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`
}

type AssetCoreModuleBaseInfo struct {
	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 描述
	Desc *string `json:"Desc,omitnil" name:"Desc"`

	// 路径
	Path *string `json:"Path,omitnil" name:"Path"`

	// 版本
	Version *string `json:"Version,omitnil" name:"Version"`

	// 服务器IP
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// 服务器名称
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 操作系统
	OsInfo *string `json:"OsInfo,omitnil" name:"OsInfo"`

	// 模块大小
	Size *uint64 `json:"Size,omitnil" name:"Size"`

	// 依赖进程数
	ProcessCount *uint64 `json:"ProcessCount,omitnil" name:"ProcessCount"`

	// 依赖模块数
	ModuleCount *uint64 `json:"ModuleCount,omitnil" name:"ModuleCount"`

	// 模块ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// 主机Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 数据更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 首次采集时间
	FirstTime *string `json:"FirstTime,omitnil" name:"FirstTime"`

	// 是否新增[0:否|1:是]
	IsNew *int64 `json:"IsNew,omitnil" name:"IsNew"`

	// 服务器外网IP
	MachineWanIp *string `json:"MachineWanIp,omitnil" name:"MachineWanIp"`

	//  附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`
}

type AssetCoreModuleDetail struct {
	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 描述
	Desc *string `json:"Desc,omitnil" name:"Desc"`

	// 路径
	Path *string `json:"Path,omitnil" name:"Path"`

	// 版本
	Version *string `json:"Version,omitnil" name:"Version"`

	// 大小
	Size *uint64 `json:"Size,omitnil" name:"Size"`

	// 依赖进程
	Processes *string `json:"Processes,omitnil" name:"Processes"`

	// 被依赖模块
	Modules *string `json:"Modules,omitnil" name:"Modules"`

	// 参数信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params []*AssetCoreModuleParam `json:"Params,omitnil" name:"Params"`

	// 数据更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

type AssetCoreModuleParam struct {
	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 数据
	Data *string `json:"Data,omitnil" name:"Data"`
}

type AssetDatabaseBaseInfo struct {
	// 主机内网IP
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// 主机外网IP
	MachineWanIp *string `json:"MachineWanIp,omitnil" name:"MachineWanIp"`

	// 主机Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 操作系统信息
	OsInfo *string `json:"OsInfo,omitnil" name:"OsInfo"`

	// 主机业务组ID
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 主机标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*MachineTag `json:"Tag,omitnil" name:"Tag"`

	// 数据库名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 版本
	Version *string `json:"Version,omitnil" name:"Version"`

	// 监听端口
	Port *string `json:"Port,omitnil" name:"Port"`

	// 协议
	Proto *string `json:"Proto,omitnil" name:"Proto"`

	// 运行用户
	User *string `json:"User,omitnil" name:"User"`

	// 绑定IP
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 配置文件路径
	ConfigPath *string `json:"ConfigPath,omitnil" name:"ConfigPath"`

	// 日志文件路径
	LogPath *string `json:"LogPath,omitnil" name:"LogPath"`

	// 数据路径
	DataPath *string `json:"DataPath,omitnil" name:"DataPath"`

	// 运行权限
	Permission *string `json:"Permission,omitnil" name:"Permission"`

	// 错误日志路径
	ErrorLogPath *string `json:"ErrorLogPath,omitnil" name:"ErrorLogPath"`

	// 插件路径
	PlugInPath *string `json:"PlugInPath,omitnil" name:"PlugInPath"`

	// 二进制路径
	BinPath *string `json:"BinPath,omitnil" name:"BinPath"`

	// 启动参数
	Param *string `json:"Param,omitnil" name:"Param"`

	// 数据库ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// 数据更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 首次采集时间
	FirstTime *string `json:"FirstTime,omitnil" name:"FirstTime"`

	// 是否新增[0:否|1:是]
	IsNew *int64 `json:"IsNew,omitnil" name:"IsNew"`

	// 主机名称
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	//  附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`
}

type AssetDatabaseDetail struct {
	// 主机内网IP
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// 主机外网IP
	MachineWanIp *string `json:"MachineWanIp,omitnil" name:"MachineWanIp"`

	// 主机Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 操作系统信息
	OsInfo *string `json:"OsInfo,omitnil" name:"OsInfo"`

	// 数据库名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 版本
	Version *string `json:"Version,omitnil" name:"Version"`

	// 监听端口
	Port *string `json:"Port,omitnil" name:"Port"`

	// 协议
	Proto *string `json:"Proto,omitnil" name:"Proto"`

	// 运行用户
	User *string `json:"User,omitnil" name:"User"`

	// 绑定IP
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 配置文件路径
	ConfigPath *string `json:"ConfigPath,omitnil" name:"ConfigPath"`

	// 日志文件路径
	LogPath *string `json:"LogPath,omitnil" name:"LogPath"`

	// 数据路径
	DataPath *string `json:"DataPath,omitnil" name:"DataPath"`

	// 运行权限
	Permission *string `json:"Permission,omitnil" name:"Permission"`

	// 错误日志路径
	ErrorLogPath *string `json:"ErrorLogPath,omitnil" name:"ErrorLogPath"`

	// 插件路径
	PlugInPath *string `json:"PlugInPath,omitnil" name:"PlugInPath"`

	// 二进制路径
	BinPath *string `json:"BinPath,omitnil" name:"BinPath"`

	// 启动参数
	Param *string `json:"Param,omitnil" name:"Param"`

	// 数据更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

type AssetDiskPartitionInfo struct {
	// 分区名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 分区大小：单位G
	Size *uint64 `json:"Size,omitnil" name:"Size"`

	// 分区使用率
	Percent *float64 `json:"Percent,omitnil" name:"Percent"`

	// 文件系统类型
	Type *string `json:"Type,omitnil" name:"Type"`

	// 挂载目录
	Path *string `json:"Path,omitnil" name:"Path"`

	// 已使用空间：单位G
	Used *uint64 `json:"Used,omitnil" name:"Used"`
}

type AssetEnvBaseInfo struct {
	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 类型：
	// 0:用户变量
	// 1:系统变量
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 启动用户
	User *string `json:"User,omitnil" name:"User"`

	// 环境变量值
	Value *string `json:"Value,omitnil" name:"Value"`

	// 服务器IP
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// 服务器名称
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 操作系统
	OsInfo *string `json:"OsInfo,omitnil" name:"OsInfo"`

	// 主机Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 数据更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 首次采集时间
	FirstTime *string `json:"FirstTime,omitnil" name:"FirstTime"`

	// 是否新增[0:否|1:是]
	IsNew *int64 `json:"IsNew,omitnil" name:"IsNew"`

	// 服务器外网IP
	MachineWanIp *string `json:"MachineWanIp,omitnil" name:"MachineWanIp"`

	//  附加信息
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`
}

type AssetFilters struct {
	// 过滤键的名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 一个或者多个过滤值。
	Values []*string `json:"Values,omitnil" name:"Values"`

	// 是否模糊查询
	ExactMatch *bool `json:"ExactMatch,omitnil" name:"ExactMatch"`
}

type AssetInitServiceBaseInfo struct {
	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 类型：
	// 1:编码器
	// 2:IE插件
	// 3:网络提供者
	// 4:镜像劫持
	// 5:LSA提供者
	// 6:KnownDLLs
	// 7:启动执行
	// 8:WMI
	// 9:计划任务
	// 10:Winsock提供者
	// 11:打印监控器
	// 12:资源管理器
	// 13:驱动服务
	// 14:登录
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 默认启用状态：0未启用，1启用
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 启动用户
	User *string `json:"User,omitnil" name:"User"`

	// 路径
	Path *string `json:"Path,omitnil" name:"Path"`

	// 服务器IP
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// 服务器名称
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 操作系统
	OsInfo *string `json:"OsInfo,omitnil" name:"OsInfo"`

	// 主机Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 数据更新时间
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 首次采集时间
	FirstTime *string `json:"FirstTime,omitnil" name:"FirstTime"`

	// 是否新增[0:否|1:是]
	IsNew *int64 `json:"IsNew,omitnil" name:"IsNew"`

	// 服务器外网IP
	MachineWanIp *string `json:"MachineWanIp,omitnil" name:"MachineWanIp"`

	//  附加信息
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`

	// 开机自启动[0:否|1:是]
	IsAutoRun *int64 `json:"IsAutoRun,omitnil" name:"IsAutoRun"`
}

type AssetJarBaseInfo struct {
	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 类型：1应用程序，2系统类库，3Web服务自带库，8:其他，
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// 是否可执行：0未知，1是，2否
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 版本
	Version *string `json:"Version,omitnil" name:"Version"`

	// 路径
	Path *string `json:"Path,omitnil" name:"Path"`

	// 服务器IP
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// 服务器名称
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 操作系统
	OsInfo *string `json:"OsInfo,omitnil" name:"OsInfo"`

	// Jar包ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// Jar包Md5
	Md5 *string `json:"Md5,omitnil" name:"Md5"`

	// 主机Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 数据更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 首次采集时间
	FirstTime *string `json:"FirstTime,omitnil" name:"FirstTime"`

	// 是否新增[0:否|1:是]
	IsNew *int64 `json:"IsNew,omitnil" name:"IsNew"`

	// 服务器外网IP
	MachineWanIp *string `json:"MachineWanIp,omitnil" name:"MachineWanIp"`

	//  附加信息
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`
}

type AssetJarDetail struct {
	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 类型：1应用程序，2系统类库，3Web服务自带库，8:其他，
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// 是否可执行：0未知，1是，2否
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 版本
	Version *string `json:"Version,omitnil" name:"Version"`

	// 路径
	Path *string `json:"Path,omitnil" name:"Path"`

	// 服务器IP
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// 服务器名称
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 操作系统
	OsInfo *string `json:"OsInfo,omitnil" name:"OsInfo"`

	// 引用进程列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Process []*AssetAppProcessInfo `json:"Process,omitnil" name:"Process"`

	// Jar包Md5
	// 注意：此字段可能返回 null，表示取不到有效值。
	Md5 *string `json:"Md5,omitnil" name:"Md5"`

	// 数据更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

type AssetKeyVal struct {
	// 标签
	Key *string `json:"Key,omitnil" name:"Key"`

	// 数量
	Value *int64 `json:"Value,omitnil" name:"Value"`

	// 描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitnil" name:"Desc"`

	// 今日新增数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewCount *int64 `json:"NewCount,omitnil" name:"NewCount"`
}

type AssetMachineBaseInfo struct {
	// 服务器Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 服务器uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 服务器内网IP
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// 服务器名称
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 操作系统名称
	OsInfo *string `json:"OsInfo,omitnil" name:"OsInfo"`

	// CPU信息
	Cpu *string `json:"Cpu,omitnil" name:"Cpu"`

	// 内存容量：单位G
	MemSize *uint64 `json:"MemSize,omitnil" name:"MemSize"`

	// 内存使用率百分比
	MemLoad *string `json:"MemLoad,omitnil" name:"MemLoad"`

	// 硬盘容量：单位G
	DiskSize *uint64 `json:"DiskSize,omitnil" name:"DiskSize"`

	// 硬盘使用率百分比
	DiskLoad *string `json:"DiskLoad,omitnil" name:"DiskLoad"`

	// 分区数
	PartitionCount *uint64 `json:"PartitionCount,omitnil" name:"PartitionCount"`

	// 主机外网IP
	MachineWanIp *string `json:"MachineWanIp,omitnil" name:"MachineWanIp"`

	// 业务组ID
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Cpu数量
	CpuSize *uint64 `json:"CpuSize,omitnil" name:"CpuSize"`

	// Cpu使用率百分比
	CpuLoad *string `json:"CpuLoad,omitnil" name:"CpuLoad"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*MachineTag `json:"Tag,omitnil" name:"Tag"`

	// 数据更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 是否新增[0:否|1:是]
	IsNew *int64 `json:"IsNew,omitnil" name:"IsNew"`

	// 首次采集时间
	FirstTime *string `json:"FirstTime,omitnil" name:"FirstTime"`

	// 附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`
}

type AssetMachineDetail struct {
	// 服务器Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 服务器uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 服务器内网IP
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// 服务器名称
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 操作系统名称
	OsInfo *string `json:"OsInfo,omitnil" name:"OsInfo"`

	// CPU信息
	Cpu *string `json:"Cpu,omitnil" name:"Cpu"`

	// 内存容量：单位G
	MemSize *uint64 `json:"MemSize,omitnil" name:"MemSize"`

	// 内存使用率百分比
	MemLoad *string `json:"MemLoad,omitnil" name:"MemLoad"`

	// 硬盘容量：单位G
	DiskSize *uint64 `json:"DiskSize,omitnil" name:"DiskSize"`

	// 硬盘使用率百分比
	DiskLoad *string `json:"DiskLoad,omitnil" name:"DiskLoad"`

	// 分区数
	PartitionCount *uint64 `json:"PartitionCount,omitnil" name:"PartitionCount"`

	// 主机外网IP
	MachineWanIp *string `json:"MachineWanIp,omitnil" name:"MachineWanIp"`

	// Cpu数量
	CpuSize *uint64 `json:"CpuSize,omitnil" name:"CpuSize"`

	// Cpu使用率百分比
	CpuLoad *string `json:"CpuLoad,omitnil" name:"CpuLoad"`

	// 防护级别：0基础版，1专业版，2旗舰版，3普惠版
	ProtectLevel *uint64 `json:"ProtectLevel,omitnil" name:"ProtectLevel"`

	// 风险状态：UNKNOW-未知，RISK-风险，SAFT-安全
	RiskStatus *string `json:"RiskStatus,omitnil" name:"RiskStatus"`

	// 已防护天数
	ProtectDays *uint64 `json:"ProtectDays,omitnil" name:"ProtectDays"`

	// 专业版开通时间
	BuyTime *string `json:"BuyTime,omitnil" name:"BuyTime"`

	// 专业版到期时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 内核版本
	CoreVersion *string `json:"CoreVersion,omitnil" name:"CoreVersion"`

	// linux/windows
	OsType *string `json:"OsType,omitnil" name:"OsType"`

	// agent版本
	AgentVersion *string `json:"AgentVersion,omitnil" name:"AgentVersion"`

	// 安装时间
	InstallTime *string `json:"InstallTime,omitnil" name:"InstallTime"`

	// 系统启动时间
	BootTime *string `json:"BootTime,omitnil" name:"BootTime"`

	// 最后上线时间
	LastLiveTime *string `json:"LastLiveTime,omitnil" name:"LastLiveTime"`

	// 生产商
	Producer *string `json:"Producer,omitnil" name:"Producer"`

	// 序列号
	SerialNumber *string `json:"SerialNumber,omitnil" name:"SerialNumber"`

	// 网卡
	NetCards []*AssetNetworkCardInfo `json:"NetCards,omitnil" name:"NetCards"`

	// 分区
	Disks []*AssetDiskPartitionInfo `json:"Disks,omitnil" name:"Disks"`

	// 0在线，1已离线
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 业务组ID
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 设备型号
	DeviceVersion *string `json:"DeviceVersion,omitnil" name:"DeviceVersion"`

	// 离线时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	OfflineTime *string `json:"OfflineTime,omitnil" name:"OfflineTime"`

	// 主机ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 数据更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 主机二外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`
}

type AssetNetworkCardInfo struct {
	// 网卡名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// Ipv4对应IP
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 网关
	GateWay *string `json:"GateWay,omitnil" name:"GateWay"`

	// MAC地址
	Mac *string `json:"Mac,omitnil" name:"Mac"`

	// Ipv6对应IP
	Ipv6 *string `json:"Ipv6,omitnil" name:"Ipv6"`

	// DNS服务器
	DnsServer *string `json:"DnsServer,omitnil" name:"DnsServer"`
}

type AssetPlanTask struct {
	// 默认启用状态：1启用，2未启用
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 执行周期
	Cycle *string `json:"Cycle,omitnil" name:"Cycle"`

	// 执行命令或脚本
	Command *string `json:"Command,omitnil" name:"Command"`

	// 启动用户
	User *string `json:"User,omitnil" name:"User"`

	// 配置文件路径
	ConfigPath *string `json:"ConfigPath,omitnil" name:"ConfigPath"`

	// 服务器IP
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// 服务器名称
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 操作系统
	OsInfo *string `json:"OsInfo,omitnil" name:"OsInfo"`

	// 主机Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 数据更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 首次采集时间
	FirstTime *string `json:"FirstTime,omitnil" name:"FirstTime"`

	// 是否新增[0:否|1:是]
	IsNew *int64 `json:"IsNew,omitnil" name:"IsNew"`

	// 服务器外网IP
	MachineWanIp *string `json:"MachineWanIp,omitnil" name:"MachineWanIp"`

	//  附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`
}

type AssetPortBaseInfo struct {
	// 主机内网IP
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// 主机外网IP
	MachineWanIp *string `json:"MachineWanIp,omitnil" name:"MachineWanIp"`

	// 主机Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 操作系统信息
	OsInfo *string `json:"OsInfo,omitnil" name:"OsInfo"`

	// 主机业务组ID
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 主机标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*MachineTag `json:"Tag,omitnil" name:"Tag"`

	// 进程名称
	ProcessName *string `json:"ProcessName,omitnil" name:"ProcessName"`

	// 进程版本
	ProcessVersion *string `json:"ProcessVersion,omitnil" name:"ProcessVersion"`

	// 进程路径
	ProcessPath *string `json:"ProcessPath,omitnil" name:"ProcessPath"`

	// 进程ID
	Pid *string `json:"Pid,omitnil" name:"Pid"`

	// 运行用户
	User *string `json:"User,omitnil" name:"User"`

	// 启动时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 启动参数
	Param *string `json:"Param,omitnil" name:"Param"`

	// 进程TTY
	Teletype *string `json:"Teletype,omitnil" name:"Teletype"`

	// 端口
	Port *string `json:"Port,omitnil" name:"Port"`

	// 所属用户组
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// 进程MD5
	Md5 *string `json:"Md5,omitnil" name:"Md5"`

	// 父进程ID
	Ppid *string `json:"Ppid,omitnil" name:"Ppid"`

	// 父进程名称
	ParentProcessName *string `json:"ParentProcessName,omitnil" name:"ParentProcessName"`

	// 端口协议
	Proto *string `json:"Proto,omitnil" name:"Proto"`

	// 绑定IP
	BindIp *string `json:"BindIp,omitnil" name:"BindIp"`

	// 主机名称
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 数据更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 首次采集时间
	FirstTime *string `json:"FirstTime,omitnil" name:"FirstTime"`

	// 是否新增[0:否|1:是]
	IsNew *int64 `json:"IsNew,omitnil" name:"IsNew"`

	//  附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`
}

type AssetProcessBaseInfo struct {
	// 主机内网IP
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// 主机外网IP
	MachineWanIp *string `json:"MachineWanIp,omitnil" name:"MachineWanIp"`

	// 主机Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 操作系统信息
	OsInfo *string `json:"OsInfo,omitnil" name:"OsInfo"`

	// 主机业务组ID
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 主机标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*MachineTag `json:"Tag,omitnil" name:"Tag"`

	// 进程名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 进程说明
	Desc *string `json:"Desc,omitnil" name:"Desc"`

	// 进程路径
	Path *string `json:"Path,omitnil" name:"Path"`

	// 进程ID
	Pid *string `json:"Pid,omitnil" name:"Pid"`

	// 运行用户
	User *string `json:"User,omitnil" name:"User"`

	// 启动时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 启动参数
	Param *string `json:"Param,omitnil" name:"Param"`

	// 进程TTY
	Tty *string `json:"Tty,omitnil" name:"Tty"`

	// 进程版本
	Version *string `json:"Version,omitnil" name:"Version"`

	// 进程用户组
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// 进程MD5
	Md5 *string `json:"Md5,omitnil" name:"Md5"`

	// 父进程ID
	Ppid *string `json:"Ppid,omitnil" name:"Ppid"`

	// 父进程名称
	ParentProcessName *string `json:"ParentProcessName,omitnil" name:"ParentProcessName"`

	// 进程状态
	Status *string `json:"Status,omitnil" name:"Status"`

	// 数字签名:0无，1有， 999 空，仅windows
	HasSign *uint64 `json:"HasSign,omitnil" name:"HasSign"`

	// 是否通过安装包安装：:0否，1是， 999 空，仅linux
	InstallByPackage *uint64 `json:"InstallByPackage,omitnil" name:"InstallByPackage"`

	// 软件包名
	PackageName *string `json:"PackageName,omitnil" name:"PackageName"`

	// 主机名称
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 数据更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 首次采集时间
	FirstTime *string `json:"FirstTime,omitnil" name:"FirstTime"`

	// 是否新增[0:否|1:是]
	IsNew *int64 `json:"IsNew,omitnil" name:"IsNew"`

	// 
	//  附加信息
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`
}

type AssetSystemPackageInfo struct {
	// 数据库名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 描述
	Desc *string `json:"Desc,omitnil" name:"Desc"`

	// 版本
	Version *string `json:"Version,omitnil" name:"Version"`

	// 安装时间
	InstallTime *string `json:"InstallTime,omitnil" name:"InstallTime"`

	// 类型
	Type *string `json:"Type,omitnil" name:"Type"`

	// 主机名称
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 主机IP
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// 操作系统
	OsInfo *string `json:"OsInfo,omitnil" name:"OsInfo"`

	// 数据更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 首次采集时间
	FirstTime *string `json:"FirstTime,omitnil" name:"FirstTime"`

	// 是否新增[0:否|1:是]
	IsNew *int64 `json:"IsNew,omitnil" name:"IsNew"`

	// 附加信息
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`

	// 主机Id
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// Agent Id
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

type AssetUserBaseInfo struct {
	// 主机内网IP
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// 主机外网IP
	MachineWanIp *string `json:"MachineWanIp,omitnil" name:"MachineWanIp"`

	// 主机名称
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 操作系统信息
	OsInfo *string `json:"OsInfo,omitnil" name:"OsInfo"`

	// 主机Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 主机Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 账号UID
	Uid *string `json:"Uid,omitnil" name:"Uid"`

	// 账号GID
	Gid *string `json:"Gid,omitnil" name:"Gid"`

	// 账号状态：0-禁用；1-启用
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 是否有root权限：0-否；1是，999为空: 仅linux
	IsRoot *uint64 `json:"IsRoot,omitnil" name:"IsRoot"`

	// 登录方式：0-不可登录；1-只允许key登录；2只允许密码登录；3-允许key和密码，999为空，仅linux
	LoginType *uint64 `json:"LoginType,omitnil" name:"LoginType"`

	// 上次登录时间
	LastLoginTime *string `json:"LastLoginTime,omitnil" name:"LastLoginTime"`

	// 账号名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 主机业务组ID
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 账号类型：0访客用户，1标准用户，2管理员用户 ,999为空,仅windows
	UserType *uint64 `json:"UserType,omitnil" name:"UserType"`

	// 是否域账号：0否， 1是，2否, 999为空  仅windows
	IsDomain *uint64 `json:"IsDomain,omitnil" name:"IsDomain"`

	// 是否有sudo权限，1是，0否, 999为空, 仅linux
	IsSudo *uint64 `json:"IsSudo,omitnil" name:"IsSudo"`

	// 是否允许ssh登录，1是，0否, 999为空, 仅linux
	IsSshLogin *uint64 `json:"IsSshLogin,omitnil" name:"IsSshLogin"`

	// Home目录
	HomePath *string `json:"HomePath,omitnil" name:"HomePath"`

	// Shell路径  仅linux
	Shell *string `json:"Shell,omitnil" name:"Shell"`

	// 是否shell登录性，0不是；1是 仅linux
	ShellLoginStatus *uint64 `json:"ShellLoginStatus,omitnil" name:"ShellLoginStatus"`

	// 密码修改时间
	PasswordChangeTime *string `json:"PasswordChangeTime,omitnil" name:"PasswordChangeTime"`

	// 密码过期时间  仅linux
	PasswordDueTime *string `json:"PasswordDueTime,omitnil" name:"PasswordDueTime"`

	// 密码锁定时间：单位天, -1为永不锁定 999为空，仅linux
	PasswordLockDays *int64 `json:"PasswordLockDays,omitnil" name:"PasswordLockDays"`

	// 密码状态：1正常 2即将过期 3已过期 4已锁定 999为空 仅linux
	PasswordStatus *int64 `json:"PasswordStatus,omitnil" name:"PasswordStatus"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 首次采集时间
	FirstTime *string `json:"FirstTime,omitnil" name:"FirstTime"`

	// 是否新增[0:否|1:是]
	IsNew *int64 `json:"IsNew,omitnil" name:"IsNew"`

	// 
	//  附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`
}

type AssetUserDetail struct {
	// 主机内网IP
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// 主机名称
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 主机Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 主机Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 账号UID
	Uid *string `json:"Uid,omitnil" name:"Uid"`

	// 账号GID
	Gid *string `json:"Gid,omitnil" name:"Gid"`

	// 账号状态：0-禁用；1-启用
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 是否有root权限：0-否；1是，999为空: 仅linux
	IsRoot *uint64 `json:"IsRoot,omitnil" name:"IsRoot"`

	// 上次登录时间
	LastLoginTime *string `json:"LastLoginTime,omitnil" name:"LastLoginTime"`

	// 账号名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 账号类型：0访客用户，1标准用户，2管理员用户 ,999为空,仅windows
	UserType *uint64 `json:"UserType,omitnil" name:"UserType"`

	// 是否域账号：0否， 1是, 999为空  仅windows
	IsDomain *uint64 `json:"IsDomain,omitnil" name:"IsDomain"`

	// 是否允许ssh登录，1是，0否, 999为空, 仅linux
	IsSshLogin *uint64 `json:"IsSshLogin,omitnil" name:"IsSshLogin"`

	// Home目录
	HomePath *string `json:"HomePath,omitnil" name:"HomePath"`

	// Shell路径  仅linux
	Shell *string `json:"Shell,omitnil" name:"Shell"`

	// 是否shell登录性，0不是；1是 仅linux
	ShellLoginStatus *uint64 `json:"ShellLoginStatus,omitnil" name:"ShellLoginStatus"`

	// 密码修改时间
	PasswordChangeTime *string `json:"PasswordChangeTime,omitnil" name:"PasswordChangeTime"`

	// 密码过期时间  仅linux
	PasswordDueTime *string `json:"PasswordDueTime,omitnil" name:"PasswordDueTime"`

	// 密码锁定时间：单位天, -1为永不锁定 999为空，仅linux
	PasswordLockDays *int64 `json:"PasswordLockDays,omitnil" name:"PasswordLockDays"`

	// 备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 用户组名
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// 账号到期时间
	DisableTime *string `json:"DisableTime,omitnil" name:"DisableTime"`

	// 最近登录终端
	LastLoginTerminal *string `json:"LastLoginTerminal,omitnil" name:"LastLoginTerminal"`

	// 最近登录位置
	LastLoginLoc *string `json:"LastLoginLoc,omitnil" name:"LastLoginLoc"`

	// 最近登录IP
	LastLoginIp *string `json:"LastLoginIp,omitnil" name:"LastLoginIp"`

	// 密码过期提醒：单位天
	PasswordWarnDays *uint64 `json:"PasswordWarnDays,omitnil" name:"PasswordWarnDays"`

	// 密码修改设置：0-不可修改，1-可修改
	PasswordChangeType *uint64 `json:"PasswordChangeType,omitnil" name:"PasswordChangeType"`

	// 用户公钥列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keys []*AssetUserKeyInfo `json:"Keys,omitnil" name:"Keys"`

	// 数据更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

type AssetUserKeyInfo struct {
	// 公钥值
	Value *string `json:"Value,omitnil" name:"Value"`

	// 公钥备注
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 加密方式
	EncryptType *string `json:"EncryptType,omitnil" name:"EncryptType"`
}

type AssetWebAppBaseInfo struct {
	// 主机内网IP
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// 主机外网IP
	MachineWanIp *string `json:"MachineWanIp,omitnil" name:"MachineWanIp"`

	// 主机Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 操作系统信息
	OsInfo *string `json:"OsInfo,omitnil" name:"OsInfo"`

	// 主机业务组ID
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 主机标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*MachineTag `json:"Tag,omitnil" name:"Tag"`

	// 应用名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 版本
	Version *string `json:"Version,omitnil" name:"Version"`

	// 根路径
	RootPath *string `json:"RootPath,omitnil" name:"RootPath"`

	// 服务类型
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`

	// 站点域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 虚拟路径
	VirtualPath *string `json:"VirtualPath,omitnil" name:"VirtualPath"`

	// 插件数
	PluginCount *uint64 `json:"PluginCount,omitnil" name:"PluginCount"`

	// 应用ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// 应用描述
	Desc *string `json:"Desc,omitnil" name:"Desc"`

	// 主机名称
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 数据更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 首次采集时间
	FirstTime *string `json:"FirstTime,omitnil" name:"FirstTime"`

	// 是否新增[0:否|1:是]
	IsNew *int64 `json:"IsNew,omitnil" name:"IsNew"`

	//  附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`
}

type AssetWebAppPluginInfo struct {
	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 描述
	Desc *string `json:"Desc,omitnil" name:"Desc"`

	// 版本
	Version *string `json:"Version,omitnil" name:"Version"`

	// 链接
	Link *string `json:"Link,omitnil" name:"Link"`
}

type AssetWebFrameBaseInfo struct {
	// 主机内网IP
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// 主机外网IP
	MachineWanIp *string `json:"MachineWanIp,omitnil" name:"MachineWanIp"`

	// 主机Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 操作系统信息
	OsInfo *string `json:"OsInfo,omitnil" name:"OsInfo"`

	// 主机业务组ID
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 主机标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*MachineTag `json:"Tag,omitnil" name:"Tag"`

	// 数据库名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 版本
	Version *string `json:"Version,omitnil" name:"Version"`

	// 语言
	Lang *string `json:"Lang,omitnil" name:"Lang"`

	// 服务类型
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`

	// 主机名称
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 数据更新时间
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 首次采集时间
	FirstTime *string `json:"FirstTime,omitnil" name:"FirstTime"`

	// 是否新增[0:否|1:是]
	IsNew *int64 `json:"IsNew,omitnil" name:"IsNew"`

	//  附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`

	// 应用路径
	Path *string `json:"Path,omitnil" name:"Path"`
}

type AssetWebLocationBaseInfo struct {
	// 主机Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 主机Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 内网IP
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// 外网IP
	MachineWanIp *string `json:"MachineWanIp,omitnil" name:"MachineWanIp"`

	// 主机名称
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 操作系统
	OsInfo *string `json:"OsInfo,omitnil" name:"OsInfo"`

	// 域名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 站点端口
	Port *string `json:"Port,omitnil" name:"Port"`

	// 站点协议
	Proto *string `json:"Proto,omitnil" name:"Proto"`

	// 服务类型
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`

	// 站点路经数
	PathCount *uint64 `json:"PathCount,omitnil" name:"PathCount"`

	// 运行用户
	User *string `json:"User,omitnil" name:"User"`

	// 主目录
	MainPath *string `json:"MainPath,omitnil" name:"MainPath"`

	// 主目录所有者
	MainPathOwner *string `json:"MainPathOwner,omitnil" name:"MainPathOwner"`

	// 拥有者权限
	Permission *string `json:"Permission,omitnil" name:"Permission"`

	// 主机业务组ID
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 主机标签
	Tag []*MachineTag `json:"Tag,omitnil" name:"Tag"`

	// Web站点Id
	Id *string `json:"Id,omitnil" name:"Id"`

	// 数据更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 首次采集时间
	FirstTime *string `json:"FirstTime,omitnil" name:"FirstTime"`

	// 是否新增[0:否|1:是]
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNew *int64 `json:"IsNew,omitnil" name:"IsNew"`

	//  附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`
}

type AssetWebLocationInfo struct {
	// 域名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 站点端口
	Port *string `json:"Port,omitnil" name:"Port"`

	// 站点协议
	Proto *string `json:"Proto,omitnil" name:"Proto"`

	// 服务类型
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`

	// 安全模块状态：0未启用，1启用，999空，仅nginx
	SafeStatus *uint64 `json:"SafeStatus,omitnil" name:"SafeStatus"`

	// 运行用户
	User *string `json:"User,omitnil" name:"User"`

	// 主目录
	MainPath *string `json:"MainPath,omitnil" name:"MainPath"`

	// 启动命令
	Command *string `json:"Command,omitnil" name:"Command"`

	// 绑定IP
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 数据更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

type AssetWebServiceBaseInfo struct {
	// 主机内网IP
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// 主机外网IP
	MachineWanIp *string `json:"MachineWanIp,omitnil" name:"MachineWanIp"`

	// 主机Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 操作系统信息
	OsInfo *string `json:"OsInfo,omitnil" name:"OsInfo"`

	// 主机业务组ID
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 主机标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*MachineTag `json:"Tag,omitnil" name:"Tag"`

	// 数据库名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 版本
	Version *string `json:"Version,omitnil" name:"Version"`

	// 二进制路径
	BinPath *string `json:"BinPath,omitnil" name:"BinPath"`

	// 启动用户
	User *string `json:"User,omitnil" name:"User"`

	// 安装路径
	InstallPath *string `json:"InstallPath,omitnil" name:"InstallPath"`

	// 配置路径
	ConfigPath *string `json:"ConfigPath,omitnil" name:"ConfigPath"`

	// 关联进程数
	ProcessCount *uint64 `json:"ProcessCount,omitnil" name:"ProcessCount"`

	// Web服务ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// 主机名称
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 描述
	Desc *string `json:"Desc,omitnil" name:"Desc"`

	// 数据更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 首次采集时间
	FirstTime *string `json:"FirstTime,omitnil" name:"FirstTime"`

	// 是否新增[0:否|1:是]
	IsNew *int64 `json:"IsNew,omitnil" name:"IsNew"`

	//  附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`
}

type BanWhiteListDetail struct {
	// 白名单ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// 白名单别名
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 阻断来源IP
	SrcIp *string `json:"SrcIp,omitnil" name:"SrcIp"`

	// 修改白名单时间
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// 创建白名单时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 白名单是否全局
	IsGlobal *bool `json:"IsGlobal,omitnil" name:"IsGlobal"`

	// 机器的UUID
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机安全程序的UUID
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 机器IP
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// 机器名称
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`
}

type BaselineBasicInfo struct {
	// 基线名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 基线id
	// 注意：此字段可能返回 null，表示取不到有效值。
	BaselineId *uint64 `json:"BaselineId,omitnil" name:"BaselineId"`

	// 父级id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentId *uint64 `json:"ParentId,omitnil" name:"ParentId"`
}

type BaselineCategory struct {
	// 分类Id
	CategoryId *int64 `json:"CategoryId,omitnil" name:"CategoryId"`

	// 分类名称
	CategoryName *string `json:"CategoryName,omitnil" name:"CategoryName"`

	// 父分类ID,如果为0则没有父分类
	ParentCategoryId *int64 `json:"ParentCategoryId,omitnil" name:"ParentCategoryId"`
}

type BaselineCustomRuleIdName struct {
	// 自定义规则ID　
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *int64 `json:"RuleId,omitnil" name:"RuleId"`

	// 自定义规则名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`
}

type BaselineDetail struct {
	// 基线描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 危害等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *uint64 `json:"Level,omitnil" name:"Level"`

	// package名
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageName *string `json:"PackageName,omitnil" name:"PackageName"`

	// 父级id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentId *uint64 `json:"ParentId,omitnil" name:"ParentId"`

	// 基线名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`
}

type BaselineDetectParam struct {
	// 检测的策略集合
	PolicyIds []*int64 `json:"PolicyIds,omitnil" name:"PolicyIds"`

	// 检测的规则集合
	RuleIds []*int64 `json:"RuleIds,omitnil" name:"RuleIds"`

	// 检测项集合
	ItemIds []*int64 `json:"ItemIds,omitnil" name:"ItemIds"`

	// 检测的主机ID集合
	HostIds []*string `json:"HostIds,omitnil" name:"HostIds"`
}

type BaselineDownload struct {
	// 任务Id
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil" name:"TaskName"`

	// 状态0:未完成 1:完成
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 完成时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 下载地址
	DownloadUrl *string `json:"DownloadUrl,omitnil" name:"DownloadUrl"`
}

type BaselineEffectHost struct {
	// 通过项
	// 注意：此字段可能返回 null，表示取不到有效值。
	PassCount *uint64 `json:"PassCount,omitnil" name:"PassCount"`

	// 风险项
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailCount *uint64 `json:"FailCount,omitnil" name:"FailCount"`

	// 首次检测事件
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstScanTime *string `json:"FirstScanTime,omitnil" name:"FirstScanTime"`

	// 最后检测时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastScanTime *string `json:"LastScanTime,omitnil" name:"LastScanTime"`

	// 风险项处理状态状态：0-未通过，1-通过
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 主机Quuid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostIp *string `json:"HostIp,omitnil" name:"HostIp"`

	// 主机别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AliasName *string `json:"AliasName,omitnil" name:"AliasName"`

	// 主机Uuid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 检测中状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxStatus *uint64 `json:"MaxStatus,omitnil" name:"MaxStatus"`
}

type BaselineEventLevelInfo struct {
	// 危害等级：1-低危；2-中危；3-高危；4-严重
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventLevel *uint64 `json:"EventLevel,omitnil" name:"EventLevel"`

	// 漏洞数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventCount *uint64 `json:"EventCount,omitnil" name:"EventCount"`
}

type BaselineFix struct {
	// 修复项名称
	ItemName *string `json:"ItemName,omitnil" name:"ItemName"`

	// 主机Ip
	HostIp *string `json:"HostIp,omitnil" name:"HostIp"`

	// 首次检测时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 最后检测时间
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// 修复时间
	FixTime *string `json:"FixTime,omitnil" name:"FixTime"`

	// 基线检测项结果ID
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// 主机额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`
}

type BaselineHost struct {
	// 主机Id
	HostId *string `json:"HostId,omitnil" name:"HostId"`

	// 主机名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostName *string `json:"HostName,omitnil" name:"HostName"`

	// 主机标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostTag *string `json:"HostTag,omitnil" name:"HostTag"`

	// 内网Ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostIp *string `json:"HostIp,omitnil" name:"HostIp"`

	// 外网Ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanIp *string `json:"WanIp,omitnil" name:"WanIp"`

	// 主机额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`
}

type BaselineHostDetect struct {
	// 主机Id
	HostId *string `json:"HostId,omitnil" name:"HostId"`

	// 内网Ip
	HostIp *string `json:"HostIp,omitnil" name:"HostIp"`

	// 主机名称
	HostName *string `json:"HostName,omitnil" name:"HostName"`

	// 外网Ip
	WanIp *string `json:"WanIp,omitnil" name:"WanIp"`

	// 0:未通过 1:忽略 3:通过 5:检测中
	DetectStatus *int64 `json:"DetectStatus,omitnil" name:"DetectStatus"`

	// 检测通过数
	PassedItemCount *int64 `json:"PassedItemCount,omitnil" name:"PassedItemCount"`

	// 关联检测项数
	ItemCount *int64 `json:"ItemCount,omitnil" name:"ItemCount"`

	// 检测未通过数
	NotPassedItemCount *int64 `json:"NotPassedItemCount,omitnil" name:"NotPassedItemCount"`

	// 首次检测时间
	FirstTime *string `json:"FirstTime,omitnil" name:"FirstTime"`

	// 最后检测时间
	LastTime *string `json:"LastTime,omitnil" name:"LastTime"`

	// 主机安全UUID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 主机额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`
}

type BaselineHostTopList struct {
	// 事件等级与次数列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventLevelList []*BaselineEventLevelInfo `json:"EventLevelList,omitnil" name:"EventLevelList"`

	// 主机名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostName *string `json:"HostName,omitnil" name:"HostName"`

	// 主机Quuid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 计算权重的分数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *uint64 `json:"Score,omitnil" name:"Score"`
}

type BaselineInfo struct {
	// 基线名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 危害等级：1-低危；2-中危；3-高危；4-严重
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *uint64 `json:"Level,omitnil" name:"Level"`

	// 检测项数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleCount *uint64 `json:"RuleCount,omitnil" name:"RuleCount"`

	// 影响服务器数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostCount *uint64 `json:"HostCount,omitnil" name:"HostCount"`

	// 通过状态:0:未通过,1:已通过
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 基线id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryId *uint64 `json:"CategoryId,omitnil" name:"CategoryId"`

	// 最后检测时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastScanTime *string `json:"LastScanTime,omitnil" name:"LastScanTime"`

	// 检测中状态: 5
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxStatus *uint64 `json:"MaxStatus,omitnil" name:"MaxStatus"`

	// 基线风险项
	// 注意：此字段可能返回 null，表示取不到有效值。
	BaselineFailCount *uint64 `json:"BaselineFailCount,omitnil" name:"BaselineFailCount"`
}

type BaselineItem struct {
	// 项Id
	ItemId *int64 `json:"ItemId,omitnil" name:"ItemId"`

	// 项名称
	ItemName *string `json:"ItemName,omitnil" name:"ItemName"`

	// 检测项分类
	CategoryId *int64 `json:"CategoryId,omitnil" name:"CategoryId"`

	// 项描述
	ItemDesc *string `json:"ItemDesc,omitnil" name:"ItemDesc"`

	// 修复方法
	FixMethod *string `json:"FixMethod,omitnil" name:"FixMethod"`

	// 所属规则
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 检测结果描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetectResultDesc *string `json:"DetectResultDesc,omitnil" name:"DetectResultDesc"`

	// 危险等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *int64 `json:"Level,omitnil" name:"Level"`

	// 检测状态：0 未通过，1：忽略，3：通过，5：检测中
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetectStatus *int64 `json:"DetectStatus,omitnil" name:"DetectStatus"`

	// 主机ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostId *string `json:"HostId,omitnil" name:"HostId"`

	// 主机名
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostName *string `json:"HostName,omitnil" name:"HostName"`

	// 主机IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostIp *string `json:"HostIp,omitnil" name:"HostIp"`

	// 外网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanIp *string `json:"WanIp,omitnil" name:"WanIp"`

	// 第一次出现时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstTime *string `json:"FirstTime,omitnil" name:"FirstTime"`

	// 最近出现时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastTime *string `json:"LastTime,omitnil" name:"LastTime"`

	// 是否可以修复
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanBeFixed *int64 `json:"CanBeFixed,omitnil" name:"CanBeFixed"`

	// 主机安全uuid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 主机额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`
}

type BaselineItemDetect struct {
	// 项Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemId *int64 `json:"ItemId,omitnil" name:"ItemId"`

	// 项名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemName *string `json:"ItemName,omitnil" name:"ItemName"`

	// 项描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemDesc *string `json:"ItemDesc,omitnil" name:"ItemDesc"`

	// 修复方法
	// 注意：此字段可能返回 null，表示取不到有效值。
	FixMethod *string `json:"FixMethod,omitnil" name:"FixMethod"`

	// 所属规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 0:未通过 1:忽略 3:通过 5:检测中
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetectStatus *int64 `json:"DetectStatus,omitnil" name:"DetectStatus"`

	// 风险等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *int64 `json:"Level,omitnil" name:"Level"`

	// 影响服务器数
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostCount *int64 `json:"HostCount,omitnil" name:"HostCount"`

	// 首次检测时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstTime *string `json:"FirstTime,omitnil" name:"FirstTime"`

	// 最后检测时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastTime *string `json:"LastTime,omitnil" name:"LastTime"`

	// 检测结果,Json字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetectResult *string `json:"DetectResult,omitnil" name:"DetectResult"`

	// 所属规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *int64 `json:"RuleId,omitnil" name:"RuleId"`

	// 通过的服务器数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PassedHostCount *int64 `json:"PassedHostCount,omitnil" name:"PassedHostCount"`

	// 未通过的服务器数
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotPassedHostCount *int64 `json:"NotPassedHostCount,omitnil" name:"NotPassedHostCount"`
}

type BaselineItemInfo struct {
	// 基线检测项ID
	ItemId *int64 `json:"ItemId,omitnil" name:"ItemId"`

	// 检测项名字
	ItemName *string `json:"ItemName,omitnil" name:"ItemName"`

	// 检测项所属规则的ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`

	// 检测项描述
	ItemDesc *string `json:"ItemDesc,omitnil" name:"ItemDesc"`

	// 检测项的修复方法
	FixMethod *string `json:"FixMethod,omitnil" name:"FixMethod"`

	// 检测项所属规则名字
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 危险等级
	Level *int64 `json:"Level,omitnil" name:"Level"`

	// 系统规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SysRuleId *int64 `json:"SysRuleId,omitnil" name:"SysRuleId"`

	// 被引自定义规则信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelatedCustomRuleInfo []*BaselineCustomRuleIdName `json:"RelatedCustomRuleInfo,omitnil" name:"RelatedCustomRuleInfo"`
}

type BaselinePolicy struct {
	// 策略名称,长度不超过128英文字符
	PolicyName *string `json:"PolicyName,omitnil" name:"PolicyName"`

	// 检测间隔[1:1天|3:3天|5:5天|7:7天]
	DetectInterval *int64 `json:"DetectInterval,omitnil" name:"DetectInterval"`

	// 检测时间
	DetectTime *string `json:"DetectTime,omitnil" name:"DetectTime"`

	// 是否开启[0:未开启|1:开启]
	IsEnabled *int64 `json:"IsEnabled,omitnil" name:"IsEnabled"`

	// 资产类型[0:所有专业版旗舰版|1:id|2:ip]
	AssetType *int64 `json:"AssetType,omitnil" name:"AssetType"`

	// 策略Id
	PolicyId *int64 `json:"PolicyId,omitnil" name:"PolicyId"`

	// 关联基线项数目
	RuleCount *int64 `json:"RuleCount,omitnil" name:"RuleCount"`

	// 关联基线项数目
	ItemCount *int64 `json:"ItemCount,omitnil" name:"ItemCount"`

	// 关联基线主机数目
	HostCount *int64 `json:"HostCount,omitnil" name:"HostCount"`

	// 规则Id
	RuleIds []*int64 `json:"RuleIds,omitnil" name:"RuleIds"`

	// 主机Id
	HostIds []*string `json:"HostIds,omitnil" name:"HostIds"`

	// 主机Ip
	HostIps []*string `json:"HostIps,omitnil" name:"HostIps"`

	// 是否是系统默认
	IsDefault *int64 `json:"IsDefault,omitnil" name:"IsDefault"`
}

type BaselinePolicyDetect struct {
	// 策略Id
	PolicyId *int64 `json:"PolicyId,omitnil" name:"PolicyId"`

	// 检测任务Id
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// 关联主机数
	HostCount *int64 `json:"HostCount,omitnil" name:"HostCount"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	FinishTime *string `json:"FinishTime,omitnil" name:"FinishTime"`

	// 策略名称
	PolicyName *string `json:"PolicyName,omitnil" name:"PolicyName"`

	// 成功主机数
	SuccessCount *int64 `json:"SuccessCount,omitnil" name:"SuccessCount"`

	// 失败主技数
	FailedCount *int64 `json:"FailedCount,omitnil" name:"FailedCount"`

	// 失败主机数
	TimeoutCount *int64 `json:"TimeoutCount,omitnil" name:"TimeoutCount"`

	// 1:检测中 2:检测完成
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyDetectStatus *int64 `json:"PolicyDetectStatus,omitnil" name:"PolicyDetectStatus"`
}

type BaselineRiskItem struct {
	// 检测项Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemId *int64 `json:"ItemId,omitnil" name:"ItemId"`

	// 检测项名字
	ItemName *string `json:"ItemName,omitnil" name:"ItemName"`

	// 风险等级
	Level *int64 `json:"Level,omitnil" name:"Level"`

	// 影响服务器数
	HostCount *int64 `json:"HostCount,omitnil" name:"HostCount"`
}

type BaselineRule struct {
	// 规则名称,长度不超过128英文字符
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 规则Id
	RuleId *int64 `json:"RuleId,omitnil" name:"RuleId"`

	// 规则类型 [0:系统|1:自定义]
	RuleType *int64 `json:"RuleType,omitnil" name:"RuleType"`

	// 规则分类
	CategoryId *int64 `json:"CategoryId,omitnil" name:"CategoryId"`

	// 规则描述
	RuleDesc *string `json:"RuleDesc,omitnil" name:"RuleDesc"`

	// 主机数
	HostCount *int64 `json:"HostCount,omitnil" name:"HostCount"`

	// 适配项ID列表
	Items []*Item `json:"Items,omitnil" name:"Items"`

	// [0:所有专业版旗舰版|1:hostID|2:ip]
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetType *int64 `json:"AssetType,omitnil" name:"AssetType"`

	// 主机Id集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostIds []*string `json:"HostIds,omitnil" name:"HostIds"`

	// 主机IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostIps []*string `json:"HostIps,omitnil" name:"HostIps"`
}

type BaselineRuleDetect struct {
	// 规则Id
	RuleId *int64 `json:"RuleId,omitnil" name:"RuleId"`

	// 规则名称
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 规则描述
	RuleDesc *string `json:"RuleDesc,omitnil" name:"RuleDesc"`

	// 关联项数
	ItemCount *int64 `json:"ItemCount,omitnil" name:"ItemCount"`

	// 关联主机数
	HostCount *int64 `json:"HostCount,omitnil" name:"HostCount"`

	// 首次检测时间
	FirstTime *string `json:"FirstTime,omitnil" name:"FirstTime"`

	// string
	LastTime *string `json:"LastTime,omitnil" name:"LastTime"`

	// 0:未通过 1:忽略 3:通过 5:检测中
	DetectStatus *int64 `json:"DetectStatus,omitnil" name:"DetectStatus"`

	// ItemID集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemIds []*int64 `json:"ItemIds,omitnil" name:"ItemIds"`
}

type BaselineRuleInfo struct {
	// 检测项名称
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 检测项描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 修复建议
	FixMessage *string `json:"FixMessage,omitnil" name:"FixMessage"`

	// 危害等级
	Level *uint64 `json:"Level,omitnil" name:"Level"`

	// 状态
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 检测项id
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`

	// 最后检测时间
	LastScanAt *string `json:"LastScanAt,omitnil" name:"LastScanAt"`

	// 具体原因说明
	RuleRemark *string `json:"RuleRemark,omitnil" name:"RuleRemark"`

	// 唯一Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 唯一事件ID
	EventId *uint64 `json:"EventId,omitnil" name:"EventId"`
}

type BaselineRuleTopInfo struct {
	// 基线检测项名
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 检测项危害等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *uint64 `json:"Level,omitnil" name:"Level"`

	// 事件总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventCount *uint64 `json:"EventCount,omitnil" name:"EventCount"`

	// 检测项id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`
}

type BaselineWeakPassword struct {
	// 密码Id
	PasswordId *int64 `json:"PasswordId,omitnil" name:"PasswordId"`

	// 密码
	WeakPassword *string `json:"WeakPassword,omitnil" name:"WeakPassword"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`
}

type BashEvent struct {
	// 数据ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 主机安全ID
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 主机ID
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机内网IP
	Hostip *string `json:"Hostip,omitnil" name:"Hostip"`

	// 执行用户名
	User *string `json:"User,omitnil" name:"User"`

	// 平台类型
	Platform *uint64 `json:"Platform,omitnil" name:"Platform"`

	// 执行命令
	BashCmd *string `json:"BashCmd,omitnil" name:"BashCmd"`

	// 规则ID
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`

	// 规则名称
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 规则等级：1-高 2-中 3-低
	RuleLevel *uint64 `json:"RuleLevel,omitnil" name:"RuleLevel"`

	// 处理状态： 0 = 待处理 1= 已处理, 2 = 已加白， 3 = 已忽略
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 发生时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 主机名
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 0: bash日志 1: 实时监控(雷霆版)
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetectBy *uint64 `json:"DetectBy,omitnil" name:"DetectBy"`

	// 进程id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pid *string `json:"Pid,omitnil" name:"Pid"`

	// 进程名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Exe *string `json:"Exe,omitnil" name:"Exe"`

	// 处理时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// 规则类别  0=系统规则，1=用户规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleCategory *uint64 `json:"RuleCategory,omitnil" name:"RuleCategory"`

	// 自动生成的正则表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegexBashCmd *string `json:"RegexBashCmd,omitnil" name:"RegexBashCmd"`
}

type BashEventNew struct {
	// 数据ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 主机安全ID
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 主机ID
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机内网IP
	HostIp *string `json:"HostIp,omitnil" name:"HostIp"`

	// 执行用户名
	User *string `json:"User,omitnil" name:"User"`

	// 平台类型
	Platform *uint64 `json:"Platform,omitnil" name:"Platform"`

	// 执行命令
	BashCmd *string `json:"BashCmd,omitnil" name:"BashCmd"`

	// 规则ID
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`

	// 规则名称
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 规则等级：1-高 2-中 3-低
	RuleLevel *uint64 `json:"RuleLevel,omitnil" name:"RuleLevel"`

	// 处理状态： 0 = 待处理 1= 已处理, 2 = 已加白， 3 = 已忽略
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 发生时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 主机名
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 0: bash日志 1: 实时监控(雷霆版)
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetectBy *uint64 `json:"DetectBy,omitnil" name:"DetectBy"`

	// 进程id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pid *string `json:"Pid,omitnil" name:"Pid"`

	// 进程名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Exe *string `json:"Exe,omitnil" name:"Exe"`

	// 处理时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// 规则类别  0=系统规则，1=用户规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleCategory *uint64 `json:"RuleCategory,omitnil" name:"RuleCategory"`

	// 自动生成的正则表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegexBashCmd *string `json:"RegexBashCmd,omitnil" name:"RegexBashCmd"`

	// 0:普通 1:专业版 2:旗舰版
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineType *int64 `json:"MachineType,omitnil" name:"MachineType"`

	// 机器额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`
}

type BashEventsInfoNew struct {
	// 数据ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 主机安全ID
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 主机ID
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机内网IP
	HostIp *string `json:"HostIp,omitnil" name:"HostIp"`

	// 平台类型
	Platform *uint64 `json:"Platform,omitnil" name:"Platform"`

	// 执行命令
	BashCmd *string `json:"BashCmd,omitnil" name:"BashCmd"`

	// 规则ID,等于0表示已规则已被删除或生效范围已修改
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`

	// 规则名称
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 规则等级：1-高 2-中 3-低
	RuleLevel *uint64 `json:"RuleLevel,omitnil" name:"RuleLevel"`

	// 处理状态： 0 = 待处理 1= 已处理, 2 = 已加白， 3= 已忽略
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 发生时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 主机名
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 进程名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Exe *string `json:"Exe,omitnil" name:"Exe"`

	// 处理时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// 规则类别  0=系统规则，1=用户规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleCategory *uint64 `json:"RuleCategory,omitnil" name:"RuleCategory"`

	// 自动生成的正则表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegexBashCmd *string `json:"RegexBashCmd,omitnil" name:"RegexBashCmd"`

	// 进程树 json  pid:进程id，exe:文件路径 ，account:进程所属用组和用户 ,cmdline:执行命令，ssh_service: SSH服务ip, ssh_soure:登录源
	// 注意：此字段可能返回 null，表示取不到有效值。
	PsTree *string `json:"PsTree,omitnil" name:"PsTree"`

	// 建议方案
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuggestScheme *string `json:"SuggestScheme,omitnil" name:"SuggestScheme"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	HarmDescribe *string `json:"HarmDescribe,omitnil" name:"HarmDescribe"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitnil" name:"Tags"`

	// 参考链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	References []*string `json:"References,omitnil" name:"References"`

	// 主机外网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineWanIp *string `json:"MachineWanIp,omitnil" name:"MachineWanIp"`

	// 主机在线状态 OFFLINE  ONLINE
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineStatus *string `json:"MachineStatus,omitnil" name:"MachineStatus"`

	// 登录用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	User *string `json:"User,omitnil" name:"User"`

	// 进程号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pid *string `json:"Pid,omitnil" name:"Pid"`

	// 0:普通 1:专业版 2:旗舰版
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineType *int64 `json:"MachineType,omitnil" name:"MachineType"`

	// 检测来源 0:bash日志 1:实时监控
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetectBy *int64 `json:"DetectBy,omitnil" name:"DetectBy"`
}

type BashRule struct {
	// 规则ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 客户端ID
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 规则名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 危险等级(0 ：无 1: 高危 2:中危 3: 低危)
	Level *uint64 `json:"Level,omitnil" name:"Level"`

	// 正则表达式
	Rule *string `json:"Rule,omitnil" name:"Rule"`

	// 规则描述
	//
	// Deprecated: Decription is deprecated.
	Decription *string `json:"Decription,omitnil" name:"Decription"`

	// 操作人
	Operator *string `json:"Operator,omitnil" name:"Operator"`

	// 是否全局规则
	IsGlobal *uint64 `json:"IsGlobal,omitnil" name:"IsGlobal"`

	// 状态 (0: 有效 1: 无效)
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// 主机IP
	Hostip *string `json:"Hostip,omitnil" name:"Hostip"`

	// 生效服务器的uuid数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uuids []*string `json:"Uuids,omitnil" name:"Uuids"`

	// 0=黑名单 1=白名单
	// 注意：此字段可能返回 null，表示取不到有效值。
	White *uint64 `json:"White,omitnil" name:"White"`

	// 是否处理之前的事件 0: 不处理 1:处理
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealOldEvents *uint64 `json:"DealOldEvents,omitnil" name:"DealOldEvents"`

	// 规则描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`
}

type Broadcasts struct {
	// 文章名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	Title *string `json:"Title,omitnil" name:"Title"`

	// 类型：0=紧急通知，1=功能更新，2=行业荣誉，3=版本发布
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 副标题
	Subtitle *string `json:"Subtitle,omitnil" name:"Subtitle"`

	// 发布时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 文章唯一id
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 危险程度  0：无， 1：严重， 2: 高危， 3:中危， 4: 低危
	Level *uint64 `json:"Level,omitnil" name:"Level"`
}

type BruteAttackInfo struct {
	// 唯一Id
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 主机安全客户端唯一标识UUID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 主机ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// 主机名
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 来源ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	SrcIp *string `json:"SrcIp,omitnil" name:"SrcIp"`

	// SUCCESS：破解成功；FAILED：破解失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 国家id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Country *uint64 `json:"Country,omitnil" name:"Country"`

	// 城市id
	// 注意：此字段可能返回 null，表示取不到有效值。
	City *uint64 `json:"City,omitnil" name:"City"`

	// 省份id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Province *uint64 `json:"Province,omitnil" name:"Province"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 阻断状态：1-阻断成功；非1-阻断失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	BanStatus *uint64 `json:"BanStatus,omitnil" name:"BanStatus"`

	// 事件类型：200-暴力破解事件，300-暴力破解成功事件（页面展示），400-暴力破解不存在的帐号事件
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventType *uint64 `json:"EventType,omitnil" name:"EventType"`

	// 发生次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// 机器UUID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 是否为专业版（true/false）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsProVersion *bool `json:"IsProVersion,omitnil" name:"IsProVersion"`

	// 被攻击的服务的用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *uint64 `json:"Port,omitnil" name:"Port"`

	// 最近攻击时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 0：待处理，1：忽略，5：已处理，6：加入白名单
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataStatus *uint64 `json:"DataStatus,omitnil" name:"DataStatus"`

	// 附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`

	// 地理位置中文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *string `json:"Location,omitnil" name:"Location"`
}

type BruteAttackRule struct {
	// 爆破事件发生的时间范围，单位：秒
	TimeRange *uint64 `json:"TimeRange,omitnil" name:"TimeRange"`

	// 爆破事件失败次数
	LoginFailTimes *uint64 `json:"LoginFailTimes,omitnil" name:"LoginFailTimes"`
}

type BruteAttackRuleList struct {
	// 爆破事件发生的时间范围，单位：秒
	TimeRange *uint64 `json:"TimeRange,omitnil" name:"TimeRange"`

	// 爆破事件失败次数
	LoginFailTimes *uint64 `json:"LoginFailTimes,omitnil" name:"LoginFailTimes"`

	// 规则是否为空，为空则填充默认规则
	Enable *bool `json:"Enable,omitnil" name:"Enable"`

	// 爆破事件发生的时间范围，单位：秒（默认规则）
	TimeRangeDefault *uint64 `json:"TimeRangeDefault,omitnil" name:"TimeRangeDefault"`

	// 爆破事件失败次数（默认规则）
	LoginFailTimesDefault *uint64 `json:"LoginFailTimesDefault,omitnil" name:"LoginFailTimesDefault"`
}

// Predefined struct for user
type CancelIgnoreVulRequestParams struct {
	// 漏洞事件id串，多个用英文逗号分隔
	EventIds *string `json:"EventIds,omitnil" name:"EventIds"`
}

type CancelIgnoreVulRequest struct {
	*tchttp.BaseRequest
	
	// 漏洞事件id串，多个用英文逗号分隔
	EventIds *string `json:"EventIds,omitnil" name:"EventIds"`
}

func (r *CancelIgnoreVulRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelIgnoreVulRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelIgnoreVulRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelIgnoreVulResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CancelIgnoreVulResponse struct {
	*tchttp.BaseResponse
	Response *CancelIgnoreVulResponseParams `json:"Response"`
}

func (r *CancelIgnoreVulResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelIgnoreVulResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeRuleEventsIgnoreStatusRequestParams struct {
	// 忽略状态 0:取消忽略 ； 1:忽略
	IgnoreStatus *uint64 `json:"IgnoreStatus,omitnil" name:"IgnoreStatus"`

	// 检测项id数组
	RuleIdList []*uint64 `json:"RuleIdList,omitnil" name:"RuleIdList"`

	// 事件id数组
	EventIdList []*uint64 `json:"EventIdList,omitnil" name:"EventIdList"`
}

type ChangeRuleEventsIgnoreStatusRequest struct {
	*tchttp.BaseRequest
	
	// 忽略状态 0:取消忽略 ； 1:忽略
	IgnoreStatus *uint64 `json:"IgnoreStatus,omitnil" name:"IgnoreStatus"`

	// 检测项id数组
	RuleIdList []*uint64 `json:"RuleIdList,omitnil" name:"RuleIdList"`

	// 事件id数组
	EventIdList []*uint64 `json:"EventIdList,omitnil" name:"EventIdList"`
}

func (r *ChangeRuleEventsIgnoreStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeRuleEventsIgnoreStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IgnoreStatus")
	delete(f, "RuleIdList")
	delete(f, "EventIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChangeRuleEventsIgnoreStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeRuleEventsIgnoreStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ChangeRuleEventsIgnoreStatusResponse struct {
	*tchttp.BaseResponse
	Response *ChangeRuleEventsIgnoreStatusResponseParams `json:"Response"`
}

func (r *ChangeRuleEventsIgnoreStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeRuleEventsIgnoreStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckBashRuleParamsRequestParams struct {
	// 校验内容 Name或Rule ，两个都要校验时逗号分割
	CheckField *string `json:"CheckField,omitnil" name:"CheckField"`

	// 在事件列表中新增白名时需要提交事件ID
	EventId *uint64 `json:"EventId,omitnil" name:"EventId"`

	// 填入的规则名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 用户填入的正则表达式："正则表达式" 需与 "提交EventId对应的命令内容" 相匹配
	Rule *string `json:"Rule,omitnil" name:"Rule"`

	// 编辑时传的规则id
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

type CheckBashRuleParamsRequest struct {
	*tchttp.BaseRequest
	
	// 校验内容 Name或Rule ，两个都要校验时逗号分割
	CheckField *string `json:"CheckField,omitnil" name:"CheckField"`

	// 在事件列表中新增白名时需要提交事件ID
	EventId *uint64 `json:"EventId,omitnil" name:"EventId"`

	// 填入的规则名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 用户填入的正则表达式："正则表达式" 需与 "提交EventId对应的命令内容" 相匹配
	Rule *string `json:"Rule,omitnil" name:"Rule"`

	// 编辑时传的规则id
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

func (r *CheckBashRuleParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckBashRuleParamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CheckField")
	delete(f, "EventId")
	delete(f, "Name")
	delete(f, "Rule")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckBashRuleParamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckBashRuleParamsResponseParams struct {
	// 0=校验通过  1=规则名称校验不通过 2=正则表达式校验不通过
	ErrCode *uint64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// 校验信息
	ErrMsg *string `json:"ErrMsg,omitnil" name:"ErrMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CheckBashRuleParamsResponse struct {
	*tchttp.BaseResponse
	Response *CheckBashRuleParamsResponseParams `json:"Response"`
}

func (r *CheckBashRuleParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckBashRuleParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComponentStatistics struct {
	// 组件ID。
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 主机数量。
	MachineNum *uint64 `json:"MachineNum,omitnil" name:"MachineNum"`

	// 组件名称。
	ComponentName *string `json:"ComponentName,omitnil" name:"ComponentName"`

	// 组件类型。
	// <li>WEB：Web组件</li>
	// <li>SYSTEM：系统组件</li>
	ComponentType *string `json:"ComponentType,omitnil" name:"ComponentType"`

	// 组件描述。
	Description *string `json:"Description,omitnil" name:"Description"`
}

// Predefined struct for user
type CreateBaselineStrategyRequestParams struct {
	// 策略名称
	StrategyName *string `json:"StrategyName,omitnil" name:"StrategyName"`

	// 检测周期, 表示每隔多少天进行检测.示例: 2, 表示每2天进行检测一次.
	ScanCycle *uint64 `json:"ScanCycle,omitnil" name:"ScanCycle"`

	// 定期检测时间，该时间下发扫描. 示例:“22:00”, 表示在22:00下发检测
	ScanAt *string `json:"ScanAt,omitnil" name:"ScanAt"`

	// 该策略下选择的基线id数组. 示例: [1,3,5,7]
	CategoryIds []*uint64 `json:"CategoryIds,omitnil" name:"CategoryIds"`

	// 扫描范围是否全部服务器, 1:是  0:否, 为1则为全部专业版主机
	IsGlobal *uint64 `json:"IsGlobal,omitnil" name:"IsGlobal"`

	// 云主机类型：
	// CVM：虚拟主机
	// BM：裸金属
	// ECM：边缘计算主机
	// LH：轻量应用服务器
	// Other：混合云机器
	MachineType *string `json:"MachineType,omitnil" name:"MachineType"`

	// 主机地域. 示例: "ap-guangzhou"
	RegionCode *string `json:"RegionCode,omitnil" name:"RegionCode"`

	// 主机id数组. 示例: ["quuid1","quuid2"]
	Quuids []*string `json:"Quuids,omitnil" name:"Quuids"`
}

type CreateBaselineStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 策略名称
	StrategyName *string `json:"StrategyName,omitnil" name:"StrategyName"`

	// 检测周期, 表示每隔多少天进行检测.示例: 2, 表示每2天进行检测一次.
	ScanCycle *uint64 `json:"ScanCycle,omitnil" name:"ScanCycle"`

	// 定期检测时间，该时间下发扫描. 示例:“22:00”, 表示在22:00下发检测
	ScanAt *string `json:"ScanAt,omitnil" name:"ScanAt"`

	// 该策略下选择的基线id数组. 示例: [1,3,5,7]
	CategoryIds []*uint64 `json:"CategoryIds,omitnil" name:"CategoryIds"`

	// 扫描范围是否全部服务器, 1:是  0:否, 为1则为全部专业版主机
	IsGlobal *uint64 `json:"IsGlobal,omitnil" name:"IsGlobal"`

	// 云主机类型：
	// CVM：虚拟主机
	// BM：裸金属
	// ECM：边缘计算主机
	// LH：轻量应用服务器
	// Other：混合云机器
	MachineType *string `json:"MachineType,omitnil" name:"MachineType"`

	// 主机地域. 示例: "ap-guangzhou"
	RegionCode *string `json:"RegionCode,omitnil" name:"RegionCode"`

	// 主机id数组. 示例: ["quuid1","quuid2"]
	Quuids []*string `json:"Quuids,omitnil" name:"Quuids"`
}

func (r *CreateBaselineStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBaselineStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StrategyName")
	delete(f, "ScanCycle")
	delete(f, "ScanAt")
	delete(f, "CategoryIds")
	delete(f, "IsGlobal")
	delete(f, "MachineType")
	delete(f, "RegionCode")
	delete(f, "Quuids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBaselineStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBaselineStrategyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateBaselineStrategyResponse struct {
	*tchttp.BaseResponse
	Response *CreateBaselineStrategyResponseParams `json:"Response"`
}

func (r *CreateBaselineStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBaselineStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEmergencyVulScanRequestParams struct {
	// 漏洞id
	VulId *uint64 `json:"VulId,omitnil" name:"VulId"`

	// 自选服务器时生效，主机uuid的string数组
	Uuids []*string `json:"Uuids,omitnil" name:"Uuids"`

	// 扫描超时时长 ，单位秒
	TimeoutPeriod *uint64 `json:"TimeoutPeriod,omitnil" name:"TimeoutPeriod"`
}

type CreateEmergencyVulScanRequest struct {
	*tchttp.BaseRequest
	
	// 漏洞id
	VulId *uint64 `json:"VulId,omitnil" name:"VulId"`

	// 自选服务器时生效，主机uuid的string数组
	Uuids []*string `json:"Uuids,omitnil" name:"Uuids"`

	// 扫描超时时长 ，单位秒
	TimeoutPeriod *uint64 `json:"TimeoutPeriod,omitnil" name:"TimeoutPeriod"`
}

func (r *CreateEmergencyVulScanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmergencyVulScanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VulId")
	delete(f, "Uuids")
	delete(f, "TimeoutPeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEmergencyVulScanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEmergencyVulScanResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateEmergencyVulScanResponse struct {
	*tchttp.BaseResponse
	Response *CreateEmergencyVulScanResponseParams `json:"Response"`
}

func (r *CreateEmergencyVulScanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmergencyVulScanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLicenseOrderRequestParams struct {
	// 标签数组, 空则表示不需要绑定标签
	Tags []*Tags `json:"Tags,omitnil" name:"Tags"`

	// 授权类型, 0 专业版-按量计费, 1专业版-包年包月 , 2 旗舰版-包年包月
	// 默认为0
	LicenseType *uint64 `json:"LicenseType,omitnil" name:"LicenseType"`

	// 授权数量 , 需要购买的数量.
	// 默认为1
	LicenseNum *uint64 `json:"LicenseNum,omitnil" name:"LicenseNum"`

	// 购买订单地域,这里仅支持 1 广州,9 新加坡. 推荐选择广州. 新加坡地域为白名单用户购买.
	// 默认为1
	RegionId *uint64 `json:"RegionId,omitnil" name:"RegionId"`

	// 项目ID .
	// 默认为0
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 购买时间长度,默认1 , 可选值为1,2,3,4,5,6,7,8,9,10,11,12,24,36
	// 该参数仅包年包月生效
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 是否自动续费, 默认不自动续费.
	// 该参数仅包年包月生效
	AutoRenewFlag *bool `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// 该字段作废
	AutoProtectOpenConfig *string `json:"AutoProtectOpenConfig,omitnil" name:"AutoProtectOpenConfig"`

	// 变配参数
	ModifyConfig *OrderModifyObject `json:"ModifyConfig,omitnil" name:"ModifyConfig"`
}

type CreateLicenseOrderRequest struct {
	*tchttp.BaseRequest
	
	// 标签数组, 空则表示不需要绑定标签
	Tags []*Tags `json:"Tags,omitnil" name:"Tags"`

	// 授权类型, 0 专业版-按量计费, 1专业版-包年包月 , 2 旗舰版-包年包月
	// 默认为0
	LicenseType *uint64 `json:"LicenseType,omitnil" name:"LicenseType"`

	// 授权数量 , 需要购买的数量.
	// 默认为1
	LicenseNum *uint64 `json:"LicenseNum,omitnil" name:"LicenseNum"`

	// 购买订单地域,这里仅支持 1 广州,9 新加坡. 推荐选择广州. 新加坡地域为白名单用户购买.
	// 默认为1
	RegionId *uint64 `json:"RegionId,omitnil" name:"RegionId"`

	// 项目ID .
	// 默认为0
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 购买时间长度,默认1 , 可选值为1,2,3,4,5,6,7,8,9,10,11,12,24,36
	// 该参数仅包年包月生效
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 是否自动续费, 默认不自动续费.
	// 该参数仅包年包月生效
	AutoRenewFlag *bool `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// 该字段作废
	AutoProtectOpenConfig *string `json:"AutoProtectOpenConfig,omitnil" name:"AutoProtectOpenConfig"`

	// 变配参数
	ModifyConfig *OrderModifyObject `json:"ModifyConfig,omitnil" name:"ModifyConfig"`
}

func (r *CreateLicenseOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLicenseOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tags")
	delete(f, "LicenseType")
	delete(f, "LicenseNum")
	delete(f, "RegionId")
	delete(f, "ProjectId")
	delete(f, "TimeSpan")
	delete(f, "AutoRenewFlag")
	delete(f, "AutoProtectOpenConfig")
	delete(f, "ModifyConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLicenseOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLicenseOrderResponseParams struct {
	// 订单号列表
	DealNames []*string `json:"DealNames,omitnil" name:"DealNames"`

	// 资源ID列表,预付费订单该字段空值
	ResourceIds []*string `json:"ResourceIds,omitnil" name:"ResourceIds"`

	// 大订单号 , 后付费该字段空值
	BigDealId *string `json:"BigDealId,omitnil" name:"BigDealId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateLicenseOrderResponse struct {
	*tchttp.BaseResponse
	Response *CreateLicenseOrderResponseParams `json:"Response"`
}

func (r *CreateLicenseOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLicenseOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProtectServerRequestParams struct {
	// 防护目录地址
	ProtectDir *string `json:"ProtectDir,omitnil" name:"ProtectDir"`

	// 防护机器 信息
	ProtectHostConfig []*ProtectHostConfig `json:"ProtectHostConfig,omitnil" name:"ProtectHostConfig"`
}

type CreateProtectServerRequest struct {
	*tchttp.BaseRequest
	
	// 防护目录地址
	ProtectDir *string `json:"ProtectDir,omitnil" name:"ProtectDir"`

	// 防护机器 信息
	ProtectHostConfig []*ProtectHostConfig `json:"ProtectHostConfig,omitnil" name:"ProtectHostConfig"`
}

func (r *CreateProtectServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProtectServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProtectDir")
	delete(f, "ProtectHostConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProtectServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProtectServerResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateProtectServerResponse struct {
	*tchttp.BaseResponse
	Response *CreateProtectServerResponseParams `json:"Response"`
}

func (r *CreateProtectServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProtectServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScanMalwareSettingRequestParams struct {
	// 扫描模式 0 全盘扫描, 1 快速扫描
	ScanPattern *uint64 `json:"ScanPattern,omitnil" name:"ScanPattern"`

	// 服务器分类：1:专业版服务器；2:自选服务器
	HostType *int64 `json:"HostType,omitnil" name:"HostType"`

	// 自选服务器时生效，主机quuid的string数组
	QuuidList []*string `json:"QuuidList,omitnil" name:"QuuidList"`

	// 超时时间单位 秒 默认3600 秒
	TimeoutPeriod *uint64 `json:"TimeoutPeriod,omitnil" name:"TimeoutPeriod"`

	// 1标准模式（只报严重、高危）、2增强模式（报严重、高危、中危）、3严格模式（报严重、高、中、低、提示）
	EngineType *uint64 `json:"EngineType,omitnil" name:"EngineType"`

	// 是否开启恶意进程查杀[0:未开启,1:开启]
	EnableMemShellScan *int64 `json:"EnableMemShellScan,omitnil" name:"EnableMemShellScan"`
}

type CreateScanMalwareSettingRequest struct {
	*tchttp.BaseRequest
	
	// 扫描模式 0 全盘扫描, 1 快速扫描
	ScanPattern *uint64 `json:"ScanPattern,omitnil" name:"ScanPattern"`

	// 服务器分类：1:专业版服务器；2:自选服务器
	HostType *int64 `json:"HostType,omitnil" name:"HostType"`

	// 自选服务器时生效，主机quuid的string数组
	QuuidList []*string `json:"QuuidList,omitnil" name:"QuuidList"`

	// 超时时间单位 秒 默认3600 秒
	TimeoutPeriod *uint64 `json:"TimeoutPeriod,omitnil" name:"TimeoutPeriod"`

	// 1标准模式（只报严重、高危）、2增强模式（报严重、高危、中危）、3严格模式（报严重、高、中、低、提示）
	EngineType *uint64 `json:"EngineType,omitnil" name:"EngineType"`

	// 是否开启恶意进程查杀[0:未开启,1:开启]
	EnableMemShellScan *int64 `json:"EnableMemShellScan,omitnil" name:"EnableMemShellScan"`
}

func (r *CreateScanMalwareSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScanMalwareSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScanPattern")
	delete(f, "HostType")
	delete(f, "QuuidList")
	delete(f, "TimeoutPeriod")
	delete(f, "EngineType")
	delete(f, "EnableMemShellScan")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateScanMalwareSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScanMalwareSettingResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateScanMalwareSettingResponse struct {
	*tchttp.BaseResponse
	Response *CreateScanMalwareSettingResponseParams `json:"Response"`
}

func (r *CreateScanMalwareSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScanMalwareSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSearchLogRequestParams struct {
	// 搜索内容
	SearchContent *string `json:"SearchContent,omitnil" name:"SearchContent"`
}

type CreateSearchLogRequest struct {
	*tchttp.BaseRequest
	
	// 搜索内容
	SearchContent *string `json:"SearchContent,omitnil" name:"SearchContent"`
}

func (r *CreateSearchLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSearchLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchContent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSearchLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSearchLogResponseParams struct {
	// 0：成功，非0：失败
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateSearchLogResponse struct {
	*tchttp.BaseResponse
	Response *CreateSearchLogResponseParams `json:"Response"`
}

func (r *CreateSearchLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSearchLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSearchTemplateRequestParams struct {
	// 搜索模板
	SearchTemplate *SearchTemplate `json:"SearchTemplate,omitnil" name:"SearchTemplate"`
}

type CreateSearchTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 搜索模板
	SearchTemplate *SearchTemplate `json:"SearchTemplate,omitnil" name:"SearchTemplate"`
}

func (r *CreateSearchTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSearchTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchTemplate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSearchTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSearchTemplateResponseParams struct {
	// 0：成功，非0：失败
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateSearchTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateSearchTemplateResponseParams `json:"Response"`
}

func (r *CreateSearchTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSearchTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DefendAttackLog struct {
	// 日志ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 客户端ID
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 来源IP
	SrcIp *string `json:"SrcIp,omitnil" name:"SrcIp"`

	// 来源端口
	SrcPort *uint64 `json:"SrcPort,omitnil" name:"SrcPort"`

	// 攻击方式
	HttpMethod *string `json:"HttpMethod,omitnil" name:"HttpMethod"`

	// 攻击描述
	HttpCgi *string `json:"HttpCgi,omitnil" name:"HttpCgi"`

	// 攻击参数
	HttpParam *string `json:"HttpParam,omitnil" name:"HttpParam"`

	// 威胁类型
	VulType *string `json:"VulType,omitnil" name:"VulType"`

	// 攻击时间
	CreatedAt *string `json:"CreatedAt,omitnil" name:"CreatedAt"`

	// 目标服务器IP
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// 目标服务器名称
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 目标IP
	DstIp *string `json:"DstIp,omitnil" name:"DstIp"`

	// 目标端口
	DstPort *uint64 `json:"DstPort,omitnil" name:"DstPort"`

	// 攻击内容
	HttpContent *string `json:"HttpContent,omitnil" name:"HttpContent"`

	// 主机额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`
}

// Predefined struct for user
type DeleteAttackLogsRequestParams struct {
	// 日志ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`

	// 是否全部删除
	IsAll *bool `json:"IsAll,omitnil" name:"IsAll"`
}

type DeleteAttackLogsRequest struct {
	*tchttp.BaseRequest
	
	// 日志ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`

	// 是否全部删除
	IsAll *bool `json:"IsAll,omitnil" name:"IsAll"`
}

func (r *DeleteAttackLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAttackLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	delete(f, "IsAll")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAttackLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAttackLogsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteAttackLogsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAttackLogsResponseParams `json:"Response"`
}

func (r *DeleteAttackLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAttackLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBaselinePolicyRequestParams struct {
	// 策略Id
	PolicyIds []*int64 `json:"PolicyIds,omitnil" name:"PolicyIds"`
}

type DeleteBaselinePolicyRequest struct {
	*tchttp.BaseRequest
	
	// 策略Id
	PolicyIds []*int64 `json:"PolicyIds,omitnil" name:"PolicyIds"`
}

func (r *DeleteBaselinePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBaselinePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBaselinePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBaselinePolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteBaselinePolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBaselinePolicyResponseParams `json:"Response"`
}

func (r *DeleteBaselinePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBaselinePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBaselineRuleIgnoreRequestParams struct {
	// 规则Id
	RuleIds []*int64 `json:"RuleIds,omitnil" name:"RuleIds"`
}

type DeleteBaselineRuleIgnoreRequest struct {
	*tchttp.BaseRequest
	
	// 规则Id
	RuleIds []*int64 `json:"RuleIds,omitnil" name:"RuleIds"`
}

func (r *DeleteBaselineRuleIgnoreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBaselineRuleIgnoreRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBaselineRuleIgnoreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBaselineRuleIgnoreResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteBaselineRuleIgnoreResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBaselineRuleIgnoreResponseParams `json:"Response"`
}

func (r *DeleteBaselineRuleIgnoreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBaselineRuleIgnoreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBaselineRuleRequestParams struct {
	// 规则Id
	RuleId *int64 `json:"RuleId,omitnil" name:"RuleId"`
}

type DeleteBaselineRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则Id
	RuleId *int64 `json:"RuleId,omitnil" name:"RuleId"`
}

func (r *DeleteBaselineRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBaselineRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBaselineRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBaselineRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteBaselineRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBaselineRuleResponseParams `json:"Response"`
}

func (r *DeleteBaselineRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBaselineRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBaselineStrategyRequestParams struct {
	// 基线策略id
	StrategyId *uint64 `json:"StrategyId,omitnil" name:"StrategyId"`
}

type DeleteBaselineStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 基线策略id
	StrategyId *uint64 `json:"StrategyId,omitnil" name:"StrategyId"`
}

func (r *DeleteBaselineStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBaselineStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StrategyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBaselineStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBaselineStrategyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteBaselineStrategyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBaselineStrategyResponseParams `json:"Response"`
}

func (r *DeleteBaselineStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBaselineStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBaselineWeakPasswordRequestParams struct {
	// 弱口令Id
	PasswordIds []*int64 `json:"PasswordIds,omitnil" name:"PasswordIds"`
}

type DeleteBaselineWeakPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 弱口令Id
	PasswordIds []*int64 `json:"PasswordIds,omitnil" name:"PasswordIds"`
}

func (r *DeleteBaselineWeakPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBaselineWeakPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PasswordIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBaselineWeakPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBaselineWeakPasswordResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteBaselineWeakPasswordResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBaselineWeakPasswordResponseParams `json:"Response"`
}

func (r *DeleteBaselineWeakPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBaselineWeakPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBashEventsRequestParams struct {
	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

type DeleteBashEventsRequest struct {
	*tchttp.BaseRequest
	
	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

func (r *DeleteBashEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBashEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBashEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBashEventsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteBashEventsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBashEventsResponseParams `json:"Response"`
}

func (r *DeleteBashEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBashEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBashRulesRequestParams struct {
	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

type DeleteBashRulesRequest struct {
	*tchttp.BaseRequest
	
	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

func (r *DeleteBashRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBashRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBashRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBashRulesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteBashRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBashRulesResponseParams `json:"Response"`
}

func (r *DeleteBashRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBashRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBruteAttacksRequestParams struct {
	// 暴力破解事件Id数组。(最大 100条)
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

type DeleteBruteAttacksRequest struct {
	*tchttp.BaseRequest
	
	// 暴力破解事件Id数组。(最大 100条)
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

func (r *DeleteBruteAttacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBruteAttacksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBruteAttacksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBruteAttacksResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteBruteAttacksResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBruteAttacksResponseParams `json:"Response"`
}

func (r *DeleteBruteAttacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBruteAttacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLicenseRecordRequestParams struct {
	// 授权ID ,可以用授权订单列表获取.
	LicenseId *uint64 `json:"LicenseId,omitnil" name:"LicenseId"`

	// 授权类型
	LicenseType *uint64 `json:"LicenseType,omitnil" name:"LicenseType"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`
}

type DeleteLicenseRecordRequest struct {
	*tchttp.BaseRequest
	
	// 授权ID ,可以用授权订单列表获取.
	LicenseId *uint64 `json:"LicenseId,omitnil" name:"LicenseId"`

	// 授权类型
	LicenseType *uint64 `json:"LicenseType,omitnil" name:"LicenseType"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`
}

func (r *DeleteLicenseRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLicenseRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LicenseId")
	delete(f, "LicenseType")
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLicenseRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLicenseRecordResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteLicenseRecordResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLicenseRecordResponseParams `json:"Response"`
}

func (r *DeleteLicenseRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLicenseRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoginWhiteListRequestParams struct {
	// 白名单ID (最大 100 条)
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

type DeleteLoginWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// 白名单ID (最大 100 条)
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

func (r *DeleteLoginWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoginWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLoginWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoginWhiteListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteLoginWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLoginWhiteListResponseParams `json:"Response"`
}

func (r *DeleteLoginWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoginWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMachineRequestParams struct {
	// 主机安全客户端Uuid。
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

type DeleteMachineRequest struct {
	*tchttp.BaseRequest
	
	// 主机安全客户端Uuid。
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

func (r *DeleteMachineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMachineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMachineResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteMachineResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMachineResponseParams `json:"Response"`
}

func (r *DeleteMachineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMachineTagRequestParams struct {
	// 关联的标签ID
	Rid *uint64 `json:"Rid,omitnil" name:"Rid"`
}

type DeleteMachineTagRequest struct {
	*tchttp.BaseRequest
	
	// 关联的标签ID
	Rid *uint64 `json:"Rid,omitnil" name:"Rid"`
}

func (r *DeleteMachineTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Rid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMachineTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMachineTagResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteMachineTagResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMachineTagResponseParams `json:"Response"`
}

func (r *DeleteMachineTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMaliciousRequestsRequestParams struct {
	// 恶意请求记录ID数组，(最大100条)
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

type DeleteMaliciousRequestsRequest struct {
	*tchttp.BaseRequest
	
	// 恶意请求记录ID数组，(最大100条)
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

func (r *DeleteMaliciousRequestsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMaliciousRequestsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMaliciousRequestsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMaliciousRequestsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteMaliciousRequestsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMaliciousRequestsResponseParams `json:"Response"`
}

func (r *DeleteMaliciousRequestsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMaliciousRequestsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMalwareScanTaskRequestParams struct {

}

type DeleteMalwareScanTaskRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DeleteMalwareScanTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMalwareScanTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMalwareScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMalwareScanTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteMalwareScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMalwareScanTaskResponseParams `json:"Response"`
}

func (r *DeleteMalwareScanTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMalwareScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMalwaresRequestParams struct {
	// 木马记录ID数组 (最大100条)
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

type DeleteMalwaresRequest struct {
	*tchttp.BaseRequest
	
	// 木马记录ID数组 (最大100条)
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

func (r *DeleteMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMalwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMalwaresResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMalwaresResponseParams `json:"Response"`
}

func (r *DeleteMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNonlocalLoginPlacesRequestParams struct {
	// 删除异地登录事件的方式，可选值："Ids"、"Ip"、"All"，默认为Ids
	DelType *string `json:"DelType,omitnil" name:"DelType"`

	// 异地登录事件ID数组。DelType为Ids或DelType未填时此项必填
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`

	// 异地登录事件的Ip。DelType为Ip时必填
	Ip []*string `json:"Ip,omitnil" name:"Ip"`

	// 主机Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

type DeleteNonlocalLoginPlacesRequest struct {
	*tchttp.BaseRequest
	
	// 删除异地登录事件的方式，可选值："Ids"、"Ip"、"All"，默认为Ids
	DelType *string `json:"DelType,omitnil" name:"DelType"`

	// 异地登录事件ID数组。DelType为Ids或DelType未填时此项必填
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`

	// 异地登录事件的Ip。DelType为Ip时必填
	Ip []*string `json:"Ip,omitnil" name:"Ip"`

	// 主机Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

func (r *DeleteNonlocalLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNonlocalLoginPlacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DelType")
	delete(f, "Ids")
	delete(f, "Ip")
	delete(f, "Uuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNonlocalLoginPlacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNonlocalLoginPlacesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteNonlocalLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNonlocalLoginPlacesResponseParams `json:"Response"`
}

func (r *DeleteNonlocalLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNonlocalLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrivilegeEventsRequestParams struct {
	// ID数组. (最大100条)
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

type DeletePrivilegeEventsRequest struct {
	*tchttp.BaseRequest
	
	// ID数组. (最大100条)
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

func (r *DeletePrivilegeEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivilegeEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrivilegeEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrivilegeEventsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeletePrivilegeEventsResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrivilegeEventsResponseParams `json:"Response"`
}

func (r *DeletePrivilegeEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivilegeEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrivilegeRulesRequestParams struct {
	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

type DeletePrivilegeRulesRequest struct {
	*tchttp.BaseRequest
	
	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

func (r *DeletePrivilegeRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivilegeRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrivilegeRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrivilegeRulesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeletePrivilegeRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrivilegeRulesResponseParams `json:"Response"`
}

func (r *DeletePrivilegeRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivilegeRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProtectDirRequestParams struct {
	// 删除的目录ID 最大100条
	Ids []*string `json:"Ids,omitnil" name:"Ids"`
}

type DeleteProtectDirRequest struct {
	*tchttp.BaseRequest
	
	// 删除的目录ID 最大100条
	Ids []*string `json:"Ids,omitnil" name:"Ids"`
}

func (r *DeleteProtectDirRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProtectDirRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProtectDirRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProtectDirResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteProtectDirResponse struct {
	*tchttp.BaseResponse
	Response *DeleteProtectDirResponseParams `json:"Response"`
}

func (r *DeleteProtectDirResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProtectDirResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReverseShellEventsRequestParams struct {
	// ID数组. (最大100条)
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

type DeleteReverseShellEventsRequest struct {
	*tchttp.BaseRequest
	
	// ID数组. (最大100条)
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

func (r *DeleteReverseShellEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReverseShellEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteReverseShellEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReverseShellEventsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteReverseShellEventsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteReverseShellEventsResponseParams `json:"Response"`
}

func (r *DeleteReverseShellEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReverseShellEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReverseShellRulesRequestParams struct {
	// ID数组. (最大100条)
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

type DeleteReverseShellRulesRequest struct {
	*tchttp.BaseRequest
	
	// ID数组. (最大100条)
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

func (r *DeleteReverseShellRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReverseShellRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteReverseShellRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReverseShellRulesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteReverseShellRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteReverseShellRulesResponseParams `json:"Response"`
}

func (r *DeleteReverseShellRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReverseShellRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteScanTaskRequestParams struct {
	// 任务Id
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// 模块类型 当前提供 Malware 木马 , Vul 漏洞 , Baseline 基线
	ModuleType *string `json:"ModuleType,omitnil" name:"ModuleType"`

	// 自选服务器时生效，主机quuid的string数组
	QuuidList []*string `json:"QuuidList,omitnil" name:"QuuidList"`
}

type DeleteScanTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务Id
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// 模块类型 当前提供 Malware 木马 , Vul 漏洞 , Baseline 基线
	ModuleType *string `json:"ModuleType,omitnil" name:"ModuleType"`

	// 自选服务器时生效，主机quuid的string数组
	QuuidList []*string `json:"QuuidList,omitnil" name:"QuuidList"`
}

func (r *DeleteScanTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScanTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ModuleType")
	delete(f, "QuuidList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteScanTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteScanTaskResponseParams `json:"Response"`
}

func (r *DeleteScanTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSearchTemplateRequestParams struct {
	// 模板ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

type DeleteSearchTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

func (r *DeleteSearchTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSearchTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSearchTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSearchTemplateResponseParams struct {
	// 0：成功，非0：失败
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteSearchTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSearchTemplateResponseParams `json:"Response"`
}

func (r *DeleteSearchTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSearchTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTagsRequestParams struct {
	// 标签ID (最大100 条)
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

type DeleteTagsRequest struct {
	*tchttp.BaseRequest
	
	// 标签ID (最大100 条)
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
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
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTagsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DeleteWebPageEventLogRequestParams struct {

}

type DeleteWebPageEventLogRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DeleteWebPageEventLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWebPageEventLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWebPageEventLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWebPageEventLogResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteWebPageEventLogResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWebPageEventLogResponseParams `json:"Response"`
}

func (r *DeleteWebPageEventLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWebPageEventLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountStatisticsRequestParams struct {
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>Username - String - 是否必填：否 - 帐号用户名</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeAccountStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>Username - String - 是否必填：否 - 帐号用户名</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeAccountStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountStatisticsResponseParams struct {
	// 帐号统计列表记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 帐号统计列表。
	AccountStatistics []*AccountStatistics `json:"AccountStatistics,omitnil" name:"AccountStatistics"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAccountStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountStatisticsResponseParams `json:"Response"`
}

func (r *DescribeAccountStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentInstallCommandRequestParams struct {
	// 是否腾讯云
	IsCloud *bool `json:"IsCloud,omitnil" name:"IsCloud"`

	// 网络类型：basic-基础网络，private-VPC, public-公网，direct-专线
	NetType *string `json:"NetType,omitnil" name:"NetType"`

	// 地域标示, NetType=direct时必填
	RegionCode *string `json:"RegionCode,omitnil" name:"RegionCode"`

	// VpcId, NetType=direct时必填
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 标签ID列表，IsCloud=false时才会生效
	TagIds []*uint64 `json:"TagIds,omitnil" name:"TagIds"`

	// 命令有效期，非腾讯云时必填
	ExpireDate *string `json:"ExpireDate,omitnil" name:"ExpireDate"`

	// 代理方式接入的vip
	Vip *string `json:"Vip,omitnil" name:"Vip"`
}

type DescribeAgentInstallCommandRequest struct {
	*tchttp.BaseRequest
	
	// 是否腾讯云
	IsCloud *bool `json:"IsCloud,omitnil" name:"IsCloud"`

	// 网络类型：basic-基础网络，private-VPC, public-公网，direct-专线
	NetType *string `json:"NetType,omitnil" name:"NetType"`

	// 地域标示, NetType=direct时必填
	RegionCode *string `json:"RegionCode,omitnil" name:"RegionCode"`

	// VpcId, NetType=direct时必填
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 标签ID列表，IsCloud=false时才会生效
	TagIds []*uint64 `json:"TagIds,omitnil" name:"TagIds"`

	// 命令有效期，非腾讯云时必填
	ExpireDate *string `json:"ExpireDate,omitnil" name:"ExpireDate"`

	// 代理方式接入的vip
	Vip *string `json:"Vip,omitnil" name:"Vip"`
}

func (r *DescribeAgentInstallCommandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentInstallCommandRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IsCloud")
	delete(f, "NetType")
	delete(f, "RegionCode")
	delete(f, "VpcId")
	delete(f, "TagIds")
	delete(f, "ExpireDate")
	delete(f, "Vip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentInstallCommandRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentInstallCommandResponseParams struct {
	// linux系统安装命令
	LinuxCommand *string `json:"LinuxCommand,omitnil" name:"LinuxCommand"`

	// windows系统安装命令（windows2008及以上）
	WindowsCommand *string `json:"WindowsCommand,omitnil" name:"WindowsCommand"`

	// windows系统安装命令第一步（windows2003）
	WindowsStepOne *string `json:"WindowsStepOne,omitnil" name:"WindowsStepOne"`

	// windows系统安装命令第二步（windows2003）
	WindowsStepTwo *string `json:"WindowsStepTwo,omitnil" name:"WindowsStepTwo"`

	// windows版agent下载链接
	WindowsDownloadUrl *string `json:"WindowsDownloadUrl,omitnil" name:"WindowsDownloadUrl"`

	// Arm安装命令
	// 注意：此字段可能返回 null，表示取不到有效值。
	ARMCommand *string `json:"ARMCommand,omitnil" name:"ARMCommand"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAgentInstallCommandResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentInstallCommandResponseParams `json:"Response"`
}

func (r *DescribeAgentInstallCommandResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentInstallCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmIncidentNodesRequestParams struct {
	// 机器uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 告警vid
	AlarmVid *string `json:"AlarmVid,omitnil" name:"AlarmVid"`

	// 告警时间
	AlarmTime *int64 `json:"AlarmTime,omitnil" name:"AlarmTime"`

	// 告警来源表ID
	TableId *int64 `json:"TableId,omitnil" name:"TableId"`
}

type DescribeAlarmIncidentNodesRequest struct {
	*tchttp.BaseRequest
	
	// 机器uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 告警vid
	AlarmVid *string `json:"AlarmVid,omitnil" name:"AlarmVid"`

	// 告警时间
	AlarmTime *int64 `json:"AlarmTime,omitnil" name:"AlarmTime"`

	// 告警来源表ID
	TableId *int64 `json:"TableId,omitnil" name:"TableId"`
}

func (r *DescribeAlarmIncidentNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmIncidentNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "AlarmVid")
	delete(f, "AlarmTime")
	delete(f, "TableId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmIncidentNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmIncidentNodesResponseParams struct {
	// 告警点所在事件的所有节点信息,可能包含多事件
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncidentNodes []*IncidentVertexInfo `json:"IncidentNodes,omitnil" name:"IncidentNodes"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAlarmIncidentNodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmIncidentNodesResponseParams `json:"Response"`
}

func (r *DescribeAlarmIncidentNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmIncidentNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmVertexIdRequestParams struct {
	// 机器uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 开始时间戳
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间戳
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`
}

type DescribeAlarmVertexIdRequest struct {
	*tchttp.BaseRequest
	
	// 机器uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 开始时间戳
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间戳
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`
}

func (r *DescribeAlarmVertexIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmVertexIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmVertexIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmVertexIdResponseParams struct {
	// 告警点id列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmVertexIds []*string `json:"AlarmVertexIds,omitnil" name:"AlarmVertexIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAlarmVertexIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmVertexIdResponseParams `json:"Response"`
}

func (r *DescribeAlarmVertexIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmVertexIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetAppListRequestParams struct {
	// 查询指定Quuid主机的信息
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>AppName- string - 是否必填：否 - 应用名搜索</li>
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Type - int - 是否必填：否 - 类型	: 仅linux
	// 0: 全部
	// 1: 运维
	// 2 : 数据库
	// 3 : 安全
	// 4 : 可疑应用
	// 5 : 系统架构
	// 6 : 系统应用
	// 7 : WEB服务
	// 99:其他</li>
	// <li>OsType - uint64 - 是否必填：否 - windows/linux</li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )</li>
	Filters []*AssetFilters `json:"Filters,omitnil" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序方式：[FirstTime|ProcessCount]
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeAssetAppListRequest struct {
	*tchttp.BaseRequest
	
	// 查询指定Quuid主机的信息
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>AppName- string - 是否必填：否 - 应用名搜索</li>
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Type - int - 是否必填：否 - 类型	: 仅linux
	// 0: 全部
	// 1: 运维
	// 2 : 数据库
	// 3 : 安全
	// 4 : 可疑应用
	// 5 : 系统架构
	// 6 : 系统应用
	// 7 : WEB服务
	// 99:其他</li>
	// <li>OsType - uint64 - 是否必填：否 - windows/linux</li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )</li>
	Filters []*AssetFilters `json:"Filters,omitnil" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序方式：[FirstTime|ProcessCount]
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeAssetAppListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetAppListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Quuid")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetAppListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetAppListResponseParams struct {
	// 应用列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Apps []*AssetAppBaseInfo `json:"Apps,omitnil" name:"Apps"`

	// 总数量
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAssetAppListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetAppListResponseParams `json:"Response"`
}

func (r *DescribeAssetAppListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetAppListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetAppProcessListRequestParams struct {
	// 主机Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// App名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeAssetAppProcessListRequest struct {
	*tchttp.BaseRequest
	
	// 主机Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// App名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeAssetAppProcessListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetAppProcessListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Quuid")
	delete(f, "Uuid")
	delete(f, "Name")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetAppProcessListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetAppProcessListResponseParams struct {
	// 进程列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Process []*AssetAppProcessInfo `json:"Process,omitnil" name:"Process"`

	// 分区总数
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAssetAppProcessListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetAppProcessListResponseParams `json:"Response"`
}

func (r *DescribeAssetAppProcessListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetAppProcessListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetCoreModuleInfoRequestParams struct {
	// 服务器Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 服务器Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 内核模块ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

type DescribeAssetCoreModuleInfoRequest struct {
	*tchttp.BaseRequest
	
	// 服务器Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 服务器Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 内核模块ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

func (r *DescribeAssetCoreModuleInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetCoreModuleInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Quuid")
	delete(f, "Uuid")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetCoreModuleInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetCoreModuleInfoResponseParams struct {
	// 内核模块详情
	Module *AssetCoreModuleDetail `json:"Module,omitnil" name:"Module"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAssetCoreModuleInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetCoreModuleInfoResponseParams `json:"Response"`
}

func (r *DescribeAssetCoreModuleInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetCoreModuleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetCoreModuleListRequestParams struct {
	// 服务器Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 服务器Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Name- string - 是否必填：否 - 包名</li>
	// <li>User- string - 是否必填：否 - 用户</li>
	Filters []*AssetFilters `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序依据[Size|FirstTime|ProcessCount|ModuleCount]
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeAssetCoreModuleListRequest struct {
	*tchttp.BaseRequest
	
	// 服务器Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 服务器Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Name- string - 是否必填：否 - 包名</li>
	// <li>User- string - 是否必填：否 - 用户</li>
	Filters []*AssetFilters `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序依据[Size|FirstTime|ProcessCount|ModuleCount]
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeAssetCoreModuleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetCoreModuleListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "Quuid")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetCoreModuleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetCoreModuleListResponseParams struct {
	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Modules []*AssetCoreModuleBaseInfo `json:"Modules,omitnil" name:"Modules"`

	// 总数量
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAssetCoreModuleListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetCoreModuleListResponseParams `json:"Response"`
}

func (r *DescribeAssetCoreModuleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetCoreModuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetDatabaseInfoRequestParams struct {
	// 服务器Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 服务器Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 数据库ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

type DescribeAssetDatabaseInfoRequest struct {
	*tchttp.BaseRequest
	
	// 服务器Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 服务器Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 数据库ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

func (r *DescribeAssetDatabaseInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetDatabaseInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Quuid")
	delete(f, "Uuid")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetDatabaseInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetDatabaseInfoResponseParams struct {
	// 数据库详情
	Database *AssetDatabaseDetail `json:"Database,omitnil" name:"Database"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAssetDatabaseInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetDatabaseInfoResponseParams `json:"Response"`
}

func (r *DescribeAssetDatabaseInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetDatabaseInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetDatabaseListRequestParams struct {
	// 查询指定Quuid主机的信息
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>User- string - 是否必填：否 - 运行用户</li>
	// <li>Ip - String - 是否必填：否 - 绑定IP</li>
	// <li>Port - Int - 是否必填：否 - 端口</li>
	// <li>Name - Int - 是否必填：否 - 数据库名称
	// 0:全部
	// 1:MySQL
	// 2:Redis
	// 3:Oracle
	// 4:MongoDB
	// 5:MemCache
	// 6:PostgreSQL
	// 7:HBase
	// 8:DB2
	// 9:Sybase
	// 10:TiDB</li>
	// <li>Proto - String - 是否必填：否 - 协议：1:TCP, 2:UDP, 3:未知</li>
	// <li>OsType - String - 是否必填：否 - 操作系统: linux/windows</li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )</li>
	Filters []*AssetFilters `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序方式：[FirstTime]
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeAssetDatabaseListRequest struct {
	*tchttp.BaseRequest
	
	// 查询指定Quuid主机的信息
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>User- string - 是否必填：否 - 运行用户</li>
	// <li>Ip - String - 是否必填：否 - 绑定IP</li>
	// <li>Port - Int - 是否必填：否 - 端口</li>
	// <li>Name - Int - 是否必填：否 - 数据库名称
	// 0:全部
	// 1:MySQL
	// 2:Redis
	// 3:Oracle
	// 4:MongoDB
	// 5:MemCache
	// 6:PostgreSQL
	// 7:HBase
	// 8:DB2
	// 9:Sybase
	// 10:TiDB</li>
	// <li>Proto - String - 是否必填：否 - 协议：1:TCP, 2:UDP, 3:未知</li>
	// <li>OsType - String - 是否必填：否 - 操作系统: linux/windows</li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )</li>
	Filters []*AssetFilters `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序方式：[FirstTime]
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeAssetDatabaseListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetDatabaseListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Quuid")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetDatabaseListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetDatabaseListResponseParams struct {
	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Databases []*AssetDatabaseBaseInfo `json:"Databases,omitnil" name:"Databases"`

	// 总数量
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAssetDatabaseListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetDatabaseListResponseParams `json:"Response"`
}

func (r *DescribeAssetDatabaseListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetDatabaseListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetEnvListRequestParams struct {
	// 服务器Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 服务器Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 该字段已废弃，由Filters代替
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Name- string - 是否必填：否 - 环境变量名</li>
	// <li>Type- int - 是否必填：否 - 类型：0用户变量，1系统变量</li>
	Filters []*AssetFilters `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序方式：[FirstTime]
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeAssetEnvListRequest struct {
	*tchttp.BaseRequest
	
	// 服务器Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 服务器Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 该字段已废弃，由Filters代替
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Name- string - 是否必填：否 - 环境变量名</li>
	// <li>Type- int - 是否必填：否 - 类型：0用户变量，1系统变量</li>
	Filters []*AssetFilters `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序方式：[FirstTime]
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeAssetEnvListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetEnvListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "Quuid")
	delete(f, "Type")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetEnvListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetEnvListResponseParams struct {
	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Envs []*AssetEnvBaseInfo `json:"Envs,omitnil" name:"Envs"`

	// 总数量
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAssetEnvListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetEnvListResponseParams `json:"Response"`
}

func (r *DescribeAssetEnvListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetEnvListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetHostTotalCountRequestParams struct {
	// 主机Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 主机Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`
}

type DescribeAssetHostTotalCountRequest struct {
	*tchttp.BaseRequest
	
	// 主机Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 主机Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`
}

func (r *DescribeAssetHostTotalCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetHostTotalCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "Quuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetHostTotalCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetHostTotalCountResponseParams struct {
	// 各项资源数量
	// system : 资源监控
	// account: 账号
	// port: 端口
	// process: 进程
	// app: 应用软件
	// database:数据库
	// webapp: Web应用
	// webframe: Web框架
	// webservice: Web服务
	// weblocation: Web站点
	// systempackage: 系统安装包
	// jar: jar包
	// initservice:启动服务
	// env: 环境变量
	// coremodule: 内核模块
	Types []*AssetKeyVal `json:"Types,omitnil" name:"Types"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAssetHostTotalCountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetHostTotalCountResponseParams `json:"Response"`
}

func (r *DescribeAssetHostTotalCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetHostTotalCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetInfoRequestParams struct {

}

type DescribeAssetInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAssetInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetInfoResponseParams struct {
	// 主机数
	MachineCount *uint64 `json:"MachineCount,omitnil" name:"MachineCount"`

	// 账号数
	AccountCount *uint64 `json:"AccountCount,omitnil" name:"AccountCount"`

	// 端口数
	PortCount *uint64 `json:"PortCount,omitnil" name:"PortCount"`

	// 进程数
	ProcessCount *uint64 `json:"ProcessCount,omitnil" name:"ProcessCount"`

	// 软件数
	SoftwareCount *uint64 `json:"SoftwareCount,omitnil" name:"SoftwareCount"`

	// 数据库数
	DatabaseCount *uint64 `json:"DatabaseCount,omitnil" name:"DatabaseCount"`

	// Web应用数
	WebAppCount *uint64 `json:"WebAppCount,omitnil" name:"WebAppCount"`

	// Web框架数
	WebFrameCount *uint64 `json:"WebFrameCount,omitnil" name:"WebFrameCount"`

	// Web服务数
	WebServiceCount *uint64 `json:"WebServiceCount,omitnil" name:"WebServiceCount"`

	// Web站点数
	WebLocationCount *uint64 `json:"WebLocationCount,omitnil" name:"WebLocationCount"`

	// 账号今日新增
	AccountNewCount *int64 `json:"AccountNewCount,omitnil" name:"AccountNewCount"`

	// 端口今日新增
	PortNewCount *int64 `json:"PortNewCount,omitnil" name:"PortNewCount"`

	// 进程今日新增
	ProcessNewCount *int64 `json:"ProcessNewCount,omitnil" name:"ProcessNewCount"`

	// 软件今日新增
	SoftwareNewCount *int64 `json:"SoftwareNewCount,omitnil" name:"SoftwareNewCount"`

	// 数据库今日新增
	DatabaseNewCount *int64 `json:"DatabaseNewCount,omitnil" name:"DatabaseNewCount"`

	// Web应用今日新增
	WebAppNewCount *int64 `json:"WebAppNewCount,omitnil" name:"WebAppNewCount"`

	// Web框架今日新增
	WebFrameNewCount *int64 `json:"WebFrameNewCount,omitnil" name:"WebFrameNewCount"`

	// Web服务今日新增
	WebServiceNewCount *int64 `json:"WebServiceNewCount,omitnil" name:"WebServiceNewCount"`

	// Web站点今日新增
	WebLocationNewCount *int64 `json:"WebLocationNewCount,omitnil" name:"WebLocationNewCount"`

	// 主机今日新增
	MachineNewCount *int64 `json:"MachineNewCount,omitnil" name:"MachineNewCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAssetInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetInfoResponseParams `json:"Response"`
}

func (r *DescribeAssetInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetInitServiceListRequestParams struct {
	// 服务器Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 服务器Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Name- string - 是否必填：否 - 包名</li>
	// <li>User- string - 是否必填：否 - 用户</li>
	// <li>IsAutoRun - string - 是否必填：否 - 是否开机自启动：0否，1是</li>
	// <li>Status- string - 是否必填：否 - 默认启用状态：0未启用， 1启用 仅linux</li>
	// <li>Type- string - 是否必填：否 - 类型：类型 仅windows：
	// 1:编码器
	// 2:IE插件
	// 3:网络提供者
	// 4:镜像劫持
	// 5:LSA提供者
	// 6:KnownDLLs
	// 7:启动执行
	// 8:WMI
	// 9:计划任务
	// 10:Winsock提供者
	// 11:打印监控器
	// 12:资源管理器
	// 13:驱动服务
	// 14:登录</li>
	Filters []*AssetFilters `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序方式：[FirstTime]
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeAssetInitServiceListRequest struct {
	*tchttp.BaseRequest
	
	// 服务器Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 服务器Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Name- string - 是否必填：否 - 包名</li>
	// <li>User- string - 是否必填：否 - 用户</li>
	// <li>IsAutoRun - string - 是否必填：否 - 是否开机自启动：0否，1是</li>
	// <li>Status- string - 是否必填：否 - 默认启用状态：0未启用， 1启用 仅linux</li>
	// <li>Type- string - 是否必填：否 - 类型：类型 仅windows：
	// 1:编码器
	// 2:IE插件
	// 3:网络提供者
	// 4:镜像劫持
	// 5:LSA提供者
	// 6:KnownDLLs
	// 7:启动执行
	// 8:WMI
	// 9:计划任务
	// 10:Winsock提供者
	// 11:打印监控器
	// 12:资源管理器
	// 13:驱动服务
	// 14:登录</li>
	Filters []*AssetFilters `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序方式：[FirstTime]
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeAssetInitServiceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetInitServiceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "Quuid")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetInitServiceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetInitServiceListResponseParams struct {
	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Services []*AssetInitServiceBaseInfo `json:"Services,omitnil" name:"Services"`

	// 总数量
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAssetInitServiceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetInitServiceListResponseParams `json:"Response"`
}

func (r *DescribeAssetInitServiceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetInitServiceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetJarInfoRequestParams struct {
	// 服务器Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 服务器Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Jar包ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

type DescribeAssetJarInfoRequest struct {
	*tchttp.BaseRequest
	
	// 服务器Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 服务器Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Jar包ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

func (r *DescribeAssetJarInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetJarInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Quuid")
	delete(f, "Uuid")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetJarInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetJarInfoResponseParams struct {
	// Jar包详情
	Jar *AssetJarDetail `json:"Jar,omitnil" name:"Jar"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAssetJarInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetJarInfoResponseParams `json:"Response"`
}

func (r *DescribeAssetJarInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetJarInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetJarListRequestParams struct {
	// 服务器Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 服务器Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Name- string - 是否必填：否 - 包名</li>
	// <li>Type- uint - 是否必填：否 - 类型	
	// 1: 应用程序
	// 2 : 系统类库
	// 3 : Web服务自带库
	// 4 : 其他依赖包</li>
	// <li>Status- string - 是否必填：否 - 是否可执行：0否，1是</li>
	Filters []*AssetFilters `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序方式：[FirstTime]
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeAssetJarListRequest struct {
	*tchttp.BaseRequest
	
	// 服务器Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 服务器Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Name- string - 是否必填：否 - 包名</li>
	// <li>Type- uint - 是否必填：否 - 类型	
	// 1: 应用程序
	// 2 : 系统类库
	// 3 : Web服务自带库
	// 4 : 其他依赖包</li>
	// <li>Status- string - 是否必填：否 - 是否可执行：0否，1是</li>
	Filters []*AssetFilters `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序方式：[FirstTime]
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeAssetJarListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetJarListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "Quuid")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetJarListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetJarListResponseParams struct {
	// 应用列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Jars []*AssetJarBaseInfo `json:"Jars,omitnil" name:"Jars"`

	// 总数量
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAssetJarListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetJarListResponseParams `json:"Response"`
}

func (r *DescribeAssetJarListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetJarListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetMachineDetailRequestParams struct {
	// 服务器Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 服务器Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

type DescribeAssetMachineDetailRequest struct {
	*tchttp.BaseRequest
	
	// 服务器Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 服务器Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

func (r *DescribeAssetMachineDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetMachineDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Quuid")
	delete(f, "Uuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetMachineDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetMachineDetailResponseParams struct {
	// 主机详情
	MachineDetail *AssetMachineDetail `json:"MachineDetail,omitnil" name:"MachineDetail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAssetMachineDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetMachineDetailResponseParams `json:"Response"`
}

func (r *DescribeAssetMachineDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetMachineDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetMachineListRequestParams struct {
	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>OsType - String - 是否必填：否 - windows或linux</li>
	// <li>CpuLoad - Int - 是否必填：否 - 
	// 0: 未知  1: 低负载
	// 2: 中负载  3: 高负载</li>
	// <li>DiskLoad - Int - 是否必填：否 - 
	// 0: 0%或未知  1: 0%～20%
	// 2: 20%～50%  3: 50%～80%
	// 4: 80%～100%</li>
	// <li>MemLoad - Int - 是否必填：否 - 
	// 0: 0%或未知  1: 0%～20%
	// 2: 20%～50%  3: 50%～80%
	// 4: 80%～100%</li>
	// <li>Quuid：主机Quuid</li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序[FirstTime|PartitionCount]
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeAssetMachineListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>OsType - String - 是否必填：否 - windows或linux</li>
	// <li>CpuLoad - Int - 是否必填：否 - 
	// 0: 未知  1: 低负载
	// 2: 中负载  3: 高负载</li>
	// <li>DiskLoad - Int - 是否必填：否 - 
	// 0: 0%或未知  1: 0%～20%
	// 2: 20%～50%  3: 50%～80%
	// 4: 80%～100%</li>
	// <li>MemLoad - Int - 是否必填：否 - 
	// 0: 0%或未知  1: 0%～20%
	// 2: 20%～50%  3: 50%～80%
	// 4: 80%～100%</li>
	// <li>Quuid：主机Quuid</li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序[FirstTime|PartitionCount]
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeAssetMachineListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetMachineListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetMachineListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetMachineListResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 记录列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Machines []*AssetMachineBaseInfo `json:"Machines,omitnil" name:"Machines"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAssetMachineListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetMachineListResponseParams `json:"Response"`
}

func (r *DescribeAssetMachineListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetMachineListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetPlanTaskListRequestParams struct {
	// 服务器Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 服务器Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>User- string - 是否必填：否 - 用户</li>
	// <li>Status- int - 是否必填：否 - 默认启用状态：0未启用， 1启用 </li>
	Filters []*AssetFilters `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序方式：[FirstTime]
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeAssetPlanTaskListRequest struct {
	*tchttp.BaseRequest
	
	// 服务器Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 服务器Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>User- string - 是否必填：否 - 用户</li>
	// <li>Status- int - 是否必填：否 - 默认启用状态：0未启用， 1启用 </li>
	Filters []*AssetFilters `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序方式：[FirstTime]
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeAssetPlanTaskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetPlanTaskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "Quuid")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetPlanTaskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetPlanTaskListResponseParams struct {
	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tasks []*AssetPlanTask `json:"Tasks,omitnil" name:"Tasks"`

	// 总数量
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAssetPlanTaskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetPlanTaskListResponseParams `json:"Response"`
}

func (r *DescribeAssetPlanTaskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetPlanTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetPortInfoListRequestParams struct {
	// 查询指定Quuid主机的信息
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>Port - uint64 - 是否必填：否 - 端口</li>
	// <li>Ip - String - 是否必填：否 - 绑定IP</li>
	// <li>ProcessName - String - 是否必填：否 - 监听进程</li>
	// <li>Pid - uint64 - 是否必填：否 - PID</li>
	// <li>User - String - 是否必填：否 - 运行用户</li>
	// <li>Group - String - 是否必填：否 - 所属用户组</li>
	// <li>Ppid - uint64 - 是否必填：否 - PPID</li>
	// <li>Proto - string - 是否必填：否 - tcp/udp或“”(空字符串筛选未知状态)</li>
	// <li>OsType - uint64 - 是否必填：否 - windows/linux</li>
	// <li>RunTimeStart - String - 是否必填：否 - 运行开始时间</li>
	// <li>RunTimeEnd - String - 是否必填：否 - 运行结束时间</li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序方式：[FirstTime|StartTime]
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeAssetPortInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 查询指定Quuid主机的信息
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>Port - uint64 - 是否必填：否 - 端口</li>
	// <li>Ip - String - 是否必填：否 - 绑定IP</li>
	// <li>ProcessName - String - 是否必填：否 - 监听进程</li>
	// <li>Pid - uint64 - 是否必填：否 - PID</li>
	// <li>User - String - 是否必填：否 - 运行用户</li>
	// <li>Group - String - 是否必填：否 - 所属用户组</li>
	// <li>Ppid - uint64 - 是否必填：否 - PPID</li>
	// <li>Proto - string - 是否必填：否 - tcp/udp或“”(空字符串筛选未知状态)</li>
	// <li>OsType - uint64 - 是否必填：否 - windows/linux</li>
	// <li>RunTimeStart - String - 是否必填：否 - 运行开始时间</li>
	// <li>RunTimeEnd - String - 是否必填：否 - 运行结束时间</li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序方式：[FirstTime|StartTime]
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeAssetPortInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetPortInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Quuid")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetPortInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetPortInfoListResponseParams struct {
	// 记录总数
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ports []*AssetPortBaseInfo `json:"Ports,omitnil" name:"Ports"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAssetPortInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetPortInfoListResponseParams `json:"Response"`
}

func (r *DescribeAssetPortInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetPortInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetProcessInfoListRequestParams struct {
	// 查询指定Quuid主机的信息
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Name - String - 是否必填：否 - 进程名</li>
	// <li>User - String - 是否必填：否 - 进程用户</li>
	// <li>Group - String - 是否必填：否 - 进程用户组</li>
	// <li>Pid - uint64 - 是否必填：否 - 进程ID</li>
	// <li>Ppid - uint64 - 是否必填：否 - 父进程ID</li>
	// <li>OsType - uint64 - 是否必填：否 - windows/linux</li>
	// <li>Status - string - 是否必填：否 - 进程状态：
	// 1:R 可执行
	// 2:S 可中断
	// 3:D 不可中断
	// 4:T 暂停状态或跟踪状态
	// 5:Z 僵尸状态
	// 6:X 将被销毁</li>
	// <li>RunTimeStart - String - 是否必填：否 - 运行开始时间</li>
	// <li>RunTimeEnd - String - 是否必填：否 - 运行结束时间</li>
	// <li>InstallByPackage - uint64 - 是否必填：否 - 是否包安装：0否，1是</li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序方式：[FirstTime|StartTime]
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeAssetProcessInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 查询指定Quuid主机的信息
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Name - String - 是否必填：否 - 进程名</li>
	// <li>User - String - 是否必填：否 - 进程用户</li>
	// <li>Group - String - 是否必填：否 - 进程用户组</li>
	// <li>Pid - uint64 - 是否必填：否 - 进程ID</li>
	// <li>Ppid - uint64 - 是否必填：否 - 父进程ID</li>
	// <li>OsType - uint64 - 是否必填：否 - windows/linux</li>
	// <li>Status - string - 是否必填：否 - 进程状态：
	// 1:R 可执行
	// 2:S 可中断
	// 3:D 不可中断
	// 4:T 暂停状态或跟踪状态
	// 5:Z 僵尸状态
	// 6:X 将被销毁</li>
	// <li>RunTimeStart - String - 是否必填：否 - 运行开始时间</li>
	// <li>RunTimeEnd - String - 是否必填：否 - 运行结束时间</li>
	// <li>InstallByPackage - uint64 - 是否必填：否 - 是否包安装：0否，1是</li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序方式：[FirstTime|StartTime]
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeAssetProcessInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetProcessInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Quuid")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetProcessInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetProcessInfoListResponseParams struct {
	// 记录总数
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Process []*AssetProcessBaseInfo `json:"Process,omitnil" name:"Process"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAssetProcessInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetProcessInfoListResponseParams `json:"Response"`
}

func (r *DescribeAssetProcessInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetProcessInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetRecentMachineInfoRequestParams struct {
	// 开始时间，如：2020-09-22
	BeginDate *string `json:"BeginDate,omitnil" name:"BeginDate"`

	// 结束时间，如：2020-09-22
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`
}

type DescribeAssetRecentMachineInfoRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间，如：2020-09-22
	BeginDate *string `json:"BeginDate,omitnil" name:"BeginDate"`

	// 结束时间，如：2020-09-22
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`
}

func (r *DescribeAssetRecentMachineInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetRecentMachineInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetRecentMachineInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetRecentMachineInfoResponseParams struct {
	// 总数量列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalList []*AssetKeyVal `json:"TotalList,omitnil" name:"TotalList"`

	// 在线数量列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveList []*AssetKeyVal `json:"LiveList,omitnil" name:"LiveList"`

	// 离线数量列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	OfflineList []*AssetKeyVal `json:"OfflineList,omitnil" name:"OfflineList"`

	// 风险数量列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskList []*AssetKeyVal `json:"RiskList,omitnil" name:"RiskList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAssetRecentMachineInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetRecentMachineInfoResponseParams `json:"Response"`
}

func (r *DescribeAssetRecentMachineInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetRecentMachineInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetSystemPackageListRequestParams struct {
	// 主机Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 主机Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>Name - String - 是否必填：否 - 包 名</li>
	// <li>StartTime - String - 是否必填：否 - 安装开始时间</li>
	// <li>EndTime - String - 是否必填：否 - 安装开始时间</li>
	// <li>Type - int - 是否必填：否 - 安装包类型：
	// 1:rmp
	// 2:dpkg
	// 3:java
	// 4:system</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式，asc-升序 或 desc-降序。默认：desc-降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序方式可选：[FistTime|InstallTime:安装时间]
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeAssetSystemPackageListRequest struct {
	*tchttp.BaseRequest
	
	// 主机Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 主机Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>Name - String - 是否必填：否 - 包 名</li>
	// <li>StartTime - String - 是否必填：否 - 安装开始时间</li>
	// <li>EndTime - String - 是否必填：否 - 安装开始时间</li>
	// <li>Type - int - 是否必填：否 - 安装包类型：
	// 1:rmp
	// 2:dpkg
	// 3:java
	// 4:system</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式，asc-升序 或 desc-降序。默认：desc-降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序方式可选：[FistTime|InstallTime:安装时间]
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeAssetSystemPackageListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetSystemPackageListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "Quuid")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetSystemPackageListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetSystemPackageListResponseParams struct {
	// 记录总数
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Packages []*AssetSystemPackageInfo `json:"Packages,omitnil" name:"Packages"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAssetSystemPackageListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetSystemPackageListResponseParams `json:"Response"`
}

func (r *DescribeAssetSystemPackageListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetSystemPackageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetUserInfoRequestParams struct {
	// 云服务器UUID
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机安全UUID
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 账户名
	Name *string `json:"Name,omitnil" name:"Name"`
}

type DescribeAssetUserInfoRequest struct {
	*tchttp.BaseRequest
	
	// 云服务器UUID
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机安全UUID
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 账户名
	Name *string `json:"Name,omitnil" name:"Name"`
}

func (r *DescribeAssetUserInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetUserInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Quuid")
	delete(f, "Uuid")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetUserInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetUserInfoResponseParams struct {
	// 用户详细信息
	User *AssetUserDetail `json:"User,omitnil" name:"User"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAssetUserInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetUserInfoResponseParams `json:"Response"`
}

func (r *DescribeAssetUserInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetUserInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetUserListRequestParams struct {
	// 查询指定Quuid主机的信息
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Name - String - 是否必填：否 - 账户名（模糊匹配）</li>
	// <li>NameStrict - String - 是否必填：否 - 账户名（严格匹配）</li>
	// <li>Uid - uint64 - 是否必填：否 - Uid</li>
	// <li>Guid - uint64 - 是否必填：否 - Guid</li>
	// <li>LoginTimeStart - String - 是否必填：否 - 开始时间，如：2021-01-11</li>
	// <li>LoginTimeEnd - String - 是否必填：否 - 结束时间，如：2021-01-11</li>
	// <li>LoginType - uint64 - 是否必填：否 - 0-不可登录；1-只允许key登录；2只允许密码登录；3-允许key和密码 仅linux</li>
	// <li>OsType - String - 是否必填：否 - windows或linux</li>
	// <li>Status - uint64 - 是否必填：否 - 账号状态：0-禁用；1-启用</li>
	// <li>UserType - uint64 - 是否必填：否 - 账号类型：0访客用户，1标准用户，2管理员用户 仅windows</li>
	// <li>IsDomain - uint64 - 是否必填：否 - 是否域账号：0 不是，1是 仅windows
	// <li>IsRoot - uint64 - 是否必填：否 - 是否Root权限：0 不是，1是 仅linux
	// <li>IsSudo - uint64 - 是否必填：否 - 是否Sudo权限：0 不是，1是 仅linux</li>
	// <li>IsSshLogin - uint64 - 是否必填：否 - 是否ssh登录：0 不是，1是 仅linux</li>
	// <li>ShellLoginStatus - uint64 - 是否必填：否 - 是否shell登录性，0不是；1是 仅linux</li>
	// <li>PasswordStatus - uint64 - 是否必填：否 - 密码状态：1正常 2即将过期 3已过期 4已锁定 仅linux</li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序：[FirstTime|LoginTime|PasswordChangeTime|PasswordDuaTime]
	// PasswordLockDays
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeAssetUserListRequest struct {
	*tchttp.BaseRequest
	
	// 查询指定Quuid主机的信息
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Name - String - 是否必填：否 - 账户名（模糊匹配）</li>
	// <li>NameStrict - String - 是否必填：否 - 账户名（严格匹配）</li>
	// <li>Uid - uint64 - 是否必填：否 - Uid</li>
	// <li>Guid - uint64 - 是否必填：否 - Guid</li>
	// <li>LoginTimeStart - String - 是否必填：否 - 开始时间，如：2021-01-11</li>
	// <li>LoginTimeEnd - String - 是否必填：否 - 结束时间，如：2021-01-11</li>
	// <li>LoginType - uint64 - 是否必填：否 - 0-不可登录；1-只允许key登录；2只允许密码登录；3-允许key和密码 仅linux</li>
	// <li>OsType - String - 是否必填：否 - windows或linux</li>
	// <li>Status - uint64 - 是否必填：否 - 账号状态：0-禁用；1-启用</li>
	// <li>UserType - uint64 - 是否必填：否 - 账号类型：0访客用户，1标准用户，2管理员用户 仅windows</li>
	// <li>IsDomain - uint64 - 是否必填：否 - 是否域账号：0 不是，1是 仅windows
	// <li>IsRoot - uint64 - 是否必填：否 - 是否Root权限：0 不是，1是 仅linux
	// <li>IsSudo - uint64 - 是否必填：否 - 是否Sudo权限：0 不是，1是 仅linux</li>
	// <li>IsSshLogin - uint64 - 是否必填：否 - 是否ssh登录：0 不是，1是 仅linux</li>
	// <li>ShellLoginStatus - uint64 - 是否必填：否 - 是否shell登录性，0不是；1是 仅linux</li>
	// <li>PasswordStatus - uint64 - 是否必填：否 - 密码状态：1正常 2即将过期 3已过期 4已锁定 仅linux</li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序：[FirstTime|LoginTime|PasswordChangeTime|PasswordDuaTime]
	// PasswordLockDays
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeAssetUserListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetUserListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Quuid")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetUserListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetUserListResponseParams struct {
	// 记录总数
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 账号列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Users []*AssetUserBaseInfo `json:"Users,omitnil" name:"Users"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAssetUserListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetUserListResponseParams `json:"Response"`
}

func (r *DescribeAssetUserListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetUserListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetWebAppListRequestParams struct {
	// 查询指定Quuid主机的信息
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Name - String - 是否必填：否 - 应用名</li>
	// <li>Domain - String - 是否必填：否 - 站点域名</li>
	// <li>Type - int - 是否必填：否 - 服务类型：
	// 0：全部
	// 1:Tomcat
	// 2:Apache
	// 3:Nginx
	// 4:WebLogic
	// 5:Websphere
	// 6:JBoss
	// 7:Jetty
	// 8:IHS
	// 9:Tengine</li>
	// <li>OsType - String - 是否必填：否 - windows/linux</li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序：[FirstTime|PluginCount]
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeAssetWebAppListRequest struct {
	*tchttp.BaseRequest
	
	// 查询指定Quuid主机的信息
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Name - String - 是否必填：否 - 应用名</li>
	// <li>Domain - String - 是否必填：否 - 站点域名</li>
	// <li>Type - int - 是否必填：否 - 服务类型：
	// 0：全部
	// 1:Tomcat
	// 2:Apache
	// 3:Nginx
	// 4:WebLogic
	// 5:Websphere
	// 6:JBoss
	// 7:Jetty
	// 8:IHS
	// 9:Tengine</li>
	// <li>OsType - String - 是否必填：否 - windows/linux</li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序：[FirstTime|PluginCount]
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeAssetWebAppListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetWebAppListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Quuid")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetWebAppListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetWebAppListResponseParams struct {
	// 记录总数
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebApps []*AssetWebAppBaseInfo `json:"WebApps,omitnil" name:"WebApps"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAssetWebAppListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetWebAppListResponseParams `json:"Response"`
}

func (r *DescribeAssetWebAppListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetWebAppListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetWebAppPluginListRequestParams struct {
	// 主机Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Web应用ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeAssetWebAppPluginListRequest struct {
	*tchttp.BaseRequest
	
	// 主机Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Web应用ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeAssetWebAppPluginListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetWebAppPluginListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Quuid")
	delete(f, "Uuid")
	delete(f, "Id")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetWebAppPluginListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetWebAppPluginListResponseParams struct {
	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Plugins []*AssetWebAppPluginInfo `json:"Plugins,omitnil" name:"Plugins"`

	// 分区总数
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAssetWebAppPluginListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetWebAppPluginListResponseParams `json:"Response"`
}

func (r *DescribeAssetWebAppPluginListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetWebAppPluginListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetWebFrameListRequestParams struct {
	// 查询指定Quuid主机的信息
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Name - String - 是否必填：否 - 框架名</li>
	// <li>NameStrict - String - 是否必填：否 - 框架名（严格匹配）</li>
	// <li>Lang - String - 是否必填：否 - 框架语言:java/python</li>
	// <li>Type - String - 是否必填：否 - 服务类型：
	// 0：全部
	// 1:Tomcat
	// 2:Apache
	// 3:Nginx
	// 4:WebLogic
	// 5:Websphere
	// 6:JBoss
	// 7:WildFly
	// 8:Jetty
	// 9:IHS
	// 10:Tengine</li>
	// <li>OsType - String - 是否必填：否 - windows/linux</li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序：[FirstTime|JarCount]
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeAssetWebFrameListRequest struct {
	*tchttp.BaseRequest
	
	// 查询指定Quuid主机的信息
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Name - String - 是否必填：否 - 框架名</li>
	// <li>NameStrict - String - 是否必填：否 - 框架名（严格匹配）</li>
	// <li>Lang - String - 是否必填：否 - 框架语言:java/python</li>
	// <li>Type - String - 是否必填：否 - 服务类型：
	// 0：全部
	// 1:Tomcat
	// 2:Apache
	// 3:Nginx
	// 4:WebLogic
	// 5:Websphere
	// 6:JBoss
	// 7:WildFly
	// 8:Jetty
	// 9:IHS
	// 10:Tengine</li>
	// <li>OsType - String - 是否必填：否 - windows/linux</li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序：[FirstTime|JarCount]
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeAssetWebFrameListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetWebFrameListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Quuid")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetWebFrameListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetWebFrameListResponseParams struct {
	// 记录总数
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebFrames []*AssetWebFrameBaseInfo `json:"WebFrames,omitnil" name:"WebFrames"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAssetWebFrameListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetWebFrameListResponseParams `json:"Response"`
}

func (r *DescribeAssetWebFrameListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetWebFrameListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetWebLocationInfoRequestParams struct {
	// 服务器Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 服务器Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 站点Id
	Id *string `json:"Id,omitnil" name:"Id"`
}

type DescribeAssetWebLocationInfoRequest struct {
	*tchttp.BaseRequest
	
	// 服务器Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 服务器Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 站点Id
	Id *string `json:"Id,omitnil" name:"Id"`
}

func (r *DescribeAssetWebLocationInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetWebLocationInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Quuid")
	delete(f, "Uuid")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetWebLocationInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetWebLocationInfoResponseParams struct {
	// 站点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebLocation *AssetWebLocationInfo `json:"WebLocation,omitnil" name:"WebLocation"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAssetWebLocationInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetWebLocationInfoResponseParams `json:"Response"`
}

func (r *DescribeAssetWebLocationInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetWebLocationInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetWebLocationListRequestParams struct {
	// 查询指定Quuid主机的信息
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Name - String - 是否必填：否 - 域名</li>
	// <li>User - String - 是否必填：否 - 运行用户</li>
	// <li>Port - uint64 - 是否必填：否 - 站点端口</li>
	// <li>Proto - uint64 - 是否必填：否 - 站点协议：1:HTTP,2:HTTPS</li>
	// <li>ServiceType - uint64 - 是否必填：否 - 服务类型：
	// 1:Tomcat
	// 2：Apache
	// 3:Nginx
	// 4:WebLogic
	// 5:Websphere
	// 6:JBoss
	// 7:WildFly
	// 8:Jetty
	// 9:IHS
	// 10:Tengine</li>
	// <li>OsType - String - 是否必填：否 - windows/linux</li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序：[FirstTime|PathCount]
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeAssetWebLocationListRequest struct {
	*tchttp.BaseRequest
	
	// 查询指定Quuid主机的信息
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Name - String - 是否必填：否 - 域名</li>
	// <li>User - String - 是否必填：否 - 运行用户</li>
	// <li>Port - uint64 - 是否必填：否 - 站点端口</li>
	// <li>Proto - uint64 - 是否必填：否 - 站点协议：1:HTTP,2:HTTPS</li>
	// <li>ServiceType - uint64 - 是否必填：否 - 服务类型：
	// 1:Tomcat
	// 2：Apache
	// 3:Nginx
	// 4:WebLogic
	// 5:Websphere
	// 6:JBoss
	// 7:WildFly
	// 8:Jetty
	// 9:IHS
	// 10:Tengine</li>
	// <li>OsType - String - 是否必填：否 - windows/linux</li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序：[FirstTime|PathCount]
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeAssetWebLocationListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetWebLocationListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Quuid")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetWebLocationListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetWebLocationListResponseParams struct {
	// 记录总数
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 站点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Locations []*AssetWebLocationBaseInfo `json:"Locations,omitnil" name:"Locations"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAssetWebLocationListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetWebLocationListResponseParams `json:"Response"`
}

func (r *DescribeAssetWebLocationListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetWebLocationListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetWebServiceInfoListRequestParams struct {
	// 查询指定Quuid主机的信息
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>User- string - 是否必填：否 - 运行用户</li>
	// <li>Name- string - 是否必填：否 - Web服务名：
	// 1:Tomcat
	// 2:Apache
	// 3:Nginx
	// 4:WebLogic
	// 5:Websphere
	// 6:JBoss
	// 7:WildFly
	// 8:Jetty
	// 9:IHS
	// 10:Tengine</li>
	// <li>OsType- string - 是否必填：否 - Windows/linux</li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )</li>
	Filters []*AssetFilters `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序：[FirstTime|ProcessCount]
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeAssetWebServiceInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 查询指定Quuid主机的信息
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>User- string - 是否必填：否 - 运行用户</li>
	// <li>Name- string - 是否必填：否 - Web服务名：
	// 1:Tomcat
	// 2:Apache
	// 3:Nginx
	// 4:WebLogic
	// 5:Websphere
	// 6:JBoss
	// 7:WildFly
	// 8:Jetty
	// 9:IHS
	// 10:Tengine</li>
	// <li>OsType- string - 是否必填：否 - Windows/linux</li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )</li>
	Filters []*AssetFilters `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序：[FirstTime|ProcessCount]
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeAssetWebServiceInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetWebServiceInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Quuid")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetWebServiceInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetWebServiceInfoListResponseParams struct {
	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebServices []*AssetWebServiceBaseInfo `json:"WebServices,omitnil" name:"WebServices"`

	// 总数量
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAssetWebServiceInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetWebServiceInfoListResponseParams `json:"Response"`
}

func (r *DescribeAssetWebServiceInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetWebServiceInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetWebServiceProcessListRequestParams struct {
	// 主机Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Web服务ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeAssetWebServiceProcessListRequest struct {
	*tchttp.BaseRequest
	
	// 主机Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Web服务ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeAssetWebServiceProcessListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetWebServiceProcessListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Quuid")
	delete(f, "Uuid")
	delete(f, "Id")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetWebServiceProcessListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetWebServiceProcessListResponseParams struct {
	// 进程列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Process []*AssetAppProcessInfo `json:"Process,omitnil" name:"Process"`

	// 总数
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAssetWebServiceProcessListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetWebServiceProcessListResponseParams `json:"Response"`
}

func (r *DescribeAssetWebServiceProcessListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetWebServiceProcessListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAttackEventsRequestParams struct {
	// 返回数量，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	//  过滤条件。
	// <li>Type - String 攻击状态 0: 尝试攻击 1: 攻击成功 - 是否必填: 否</li>
	// <li>Status - String 事件处理状态 0：待处理 1：已处理 2： 已加白 3： 已忽略 4：已删除  - 是否必填: 否</li>
	// <li>SrcIP - String 来源IP - 是否必填: 否</li>
	// <li>Uuids - String 主机安全uuid - 是否必填: 否</li>
	// <li>Quuids - String cvm uuid - 是否必填: 否</li>
	// <li>DstPort - String 攻击目标端口 - 是否必填: 否</li>
	// <li>MachineName - String 主机名称 - 是否必填: 否</li>
	// <li>InstanceID - String 主机实例ID - 是否必填: 否</li>
	// <li>AttackTimeBegin - String 攻击开始时间 - 是否必填: 否</li>
	// <li>AttackTimeEnd - String 攻击结束时间 - 是否必填: 否</li>
	// <li>VulSupportDefense - String 漏洞是否支持防御 0不支持，1支持 - 是否必填: 否</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 排序
	By *string `json:"By,omitnil" name:"By"`

	// 排序方式 ASC,DESC
	Order *string `json:"Order,omitnil" name:"Order"`
}

type DescribeAttackEventsRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	//  过滤条件。
	// <li>Type - String 攻击状态 0: 尝试攻击 1: 攻击成功 - 是否必填: 否</li>
	// <li>Status - String 事件处理状态 0：待处理 1：已处理 2： 已加白 3： 已忽略 4：已删除  - 是否必填: 否</li>
	// <li>SrcIP - String 来源IP - 是否必填: 否</li>
	// <li>Uuids - String 主机安全uuid - 是否必填: 否</li>
	// <li>Quuids - String cvm uuid - 是否必填: 否</li>
	// <li>DstPort - String 攻击目标端口 - 是否必填: 否</li>
	// <li>MachineName - String 主机名称 - 是否必填: 否</li>
	// <li>InstanceID - String 主机实例ID - 是否必填: 否</li>
	// <li>AttackTimeBegin - String 攻击开始时间 - 是否必填: 否</li>
	// <li>AttackTimeEnd - String 攻击结束时间 - 是否必填: 否</li>
	// <li>VulSupportDefense - String 漏洞是否支持防御 0不支持，1支持 - 是否必填: 否</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 排序
	By *string `json:"By,omitnil" name:"By"`

	// 排序方式 ASC,DESC
	Order *string `json:"Order,omitnil" name:"Order"`
}

func (r *DescribeAttackEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttackEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAttackEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAttackEventsResponseParams struct {
	// 攻击事件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*NetAttackEvent `json:"List,omitnil" name:"List"`

	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAttackEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAttackEventsResponseParams `json:"Response"`
}

func (r *DescribeAttackEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttackEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAttackLogInfoRequestParams struct {
	// 日志ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

type DescribeAttackLogInfoRequest struct {
	*tchttp.BaseRequest
	
	// 日志ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

func (r *DescribeAttackLogInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttackLogInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAttackLogInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAttackLogInfoResponseParams struct {
	// 日志ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 主机ID
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 攻击来源端口
	SrcPort *uint64 `json:"SrcPort,omitnil" name:"SrcPort"`

	// 攻击来源IP
	SrcIp *string `json:"SrcIp,omitnil" name:"SrcIp"`

	// 攻击目标端口
	DstPort *uint64 `json:"DstPort,omitnil" name:"DstPort"`

	// 攻击目标IP
	DstIp *string `json:"DstIp,omitnil" name:"DstIp"`

	// 攻击方法
	HttpMethod *string `json:"HttpMethod,omitnil" name:"HttpMethod"`

	// 攻击目标主机
	HttpHost *string `json:"HttpHost,omitnil" name:"HttpHost"`

	// 攻击头信息
	HttpHead *string `json:"HttpHead,omitnil" name:"HttpHead"`

	// 攻击者浏览器标识
	HttpUserAgent *string `json:"HttpUserAgent,omitnil" name:"HttpUserAgent"`

	// 请求源
	HttpReferer *string `json:"HttpReferer,omitnil" name:"HttpReferer"`

	// 威胁类型
	VulType *string `json:"VulType,omitnil" name:"VulType"`

	// 攻击路径
	HttpCgi *string `json:"HttpCgi,omitnil" name:"HttpCgi"`

	// 攻击参数
	HttpParam *string `json:"HttpParam,omitnil" name:"HttpParam"`

	// 攻击时间
	CreatedAt *string `json:"CreatedAt,omitnil" name:"CreatedAt"`

	// 攻击内容
	HttpContent *string `json:"HttpContent,omitnil" name:"HttpContent"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAttackLogInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAttackLogInfoResponseParams `json:"Response"`
}

func (r *DescribeAttackLogInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttackLogInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAttackLogsRequestParams struct {
	// 返回数量，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>HttpMethod - String - 是否必填：否 - 攻击方法(POST|GET)</li>
	// <li>DateRange - String - 是否必填：否 - 时间范围(存储最近3个月的数据)，如最近一个月["2019-11-17", "2019-12-17"]</li>
	// <li>VulType - String 威胁类型 - 是否必填: 否</li>
	// <li>SrcIp - String 攻击源IP - 是否必填: 否</li>
	// <li>DstIp - String 攻击目标IP - 是否必填: 否</li>
	// <li>SrcPort - String 攻击源端口 - 是否必填: 否</li>
	// <li>DstPort - String 攻击目标端口 - 是否必填: 否</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 主机安全客户端ID
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 云主机机器ID
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`
}

type DescribeAttackLogsRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>HttpMethod - String - 是否必填：否 - 攻击方法(POST|GET)</li>
	// <li>DateRange - String - 是否必填：否 - 时间范围(存储最近3个月的数据)，如最近一个月["2019-11-17", "2019-12-17"]</li>
	// <li>VulType - String 威胁类型 - 是否必填: 否</li>
	// <li>SrcIp - String 攻击源IP - 是否必填: 否</li>
	// <li>DstIp - String 攻击目标IP - 是否必填: 否</li>
	// <li>SrcPort - String 攻击源端口 - 是否必填: 否</li>
	// <li>DstPort - String 攻击目标端口 - 是否必填: 否</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 主机安全客户端ID
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 云主机机器ID
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`
}

func (r *DescribeAttackLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttackLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Uuid")
	delete(f, "Quuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAttackLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAttackLogsResponseParams struct {
	// 日志列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackLogs []*DefendAttackLog `json:"AttackLogs,omitnil" name:"AttackLogs"`

	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAttackLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAttackLogsResponseParams `json:"Response"`
}

func (r *DescribeAttackLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttackLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAttackVulTypeListRequestParams struct {

}

type DescribeAttackVulTypeListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAttackVulTypeListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttackVulTypeListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAttackVulTypeListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAttackVulTypeListResponseParams struct {
	// 威胁类型列表
	List []*string `json:"List,omitnil" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAttackVulTypeListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAttackVulTypeListResponseParams `json:"Response"`
}

func (r *DescribeAttackVulTypeListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttackVulTypeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAvailableExpertServiceDetailRequestParams struct {

}

type DescribeAvailableExpertServiceDetailRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAvailableExpertServiceDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailableExpertServiceDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAvailableExpertServiceDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAvailableExpertServiceDetailResponseParams struct {
	// 安全管家订单
	ExpertService []*ExpertServiceOrderInfo `json:"ExpertService,omitnil" name:"ExpertService"`

	// 应急响应可用次数
	EmergencyResponse *uint64 `json:"EmergencyResponse,omitnil" name:"EmergencyResponse"`

	// 旗舰护网可用次数
	ProtectNet *uint64 `json:"ProtectNet,omitnil" name:"ProtectNet"`

	// 是否购买过安全管家
	ExpertServiceBuy *bool `json:"ExpertServiceBuy,omitnil" name:"ExpertServiceBuy"`

	// 是否购买过应急响应
	EmergencyResponseBuy *bool `json:"EmergencyResponseBuy,omitnil" name:"EmergencyResponseBuy"`

	// 是否购买过旗舰护网
	ProtectNetBuy *bool `json:"ProtectNetBuy,omitnil" name:"ProtectNetBuy"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAvailableExpertServiceDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAvailableExpertServiceDetailResponseParams `json:"Response"`
}

func (r *DescribeAvailableExpertServiceDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailableExpertServiceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBanModeRequestParams struct {

}

type DescribeBanModeRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeBanModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBanModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBanModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBanModeResponseParams struct {
	// 阻断模式，STANDARD_MODE：标准阻断，DEEP_MODE：深度阻断
	Mode *string `json:"Mode,omitnil" name:"Mode"`

	// 标准阻断模式的配置
	StandardModeConfig *StandardModeConfig `json:"StandardModeConfig,omitnil" name:"StandardModeConfig"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBanModeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBanModeResponseParams `json:"Response"`
}

func (r *DescribeBanModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBanModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBanRegionsRequestParams struct {
	// 阻断模式，STANDARD_MODE：标准阻断，DEEP_MODE：深度阻断
	Mode *string `json:"Mode,omitnil" name:"Mode"`
}

type DescribeBanRegionsRequest struct {
	*tchttp.BaseRequest
	
	// 阻断模式，STANDARD_MODE：标准阻断，DEEP_MODE：深度阻断
	Mode *string `json:"Mode,omitnil" name:"Mode"`
}

func (r *DescribeBanRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBanRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Mode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBanRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBanRegionsResponseParams struct {
	// 地域信息列表
	RegionSet []*RegionSet `json:"RegionSet,omitnil" name:"RegionSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBanRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBanRegionsResponseParams `json:"Response"`
}

func (r *DescribeBanRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBanRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBanStatusRequestParams struct {

}

type DescribeBanStatusRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeBanStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBanStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBanStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBanStatusResponseParams struct {
	// 阻断开关状态 0:关闭 1:开启
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 是否弹窗提示信息 false: 关闭，true: 开启
	ShowTips *bool `json:"ShowTips,omitnil" name:"ShowTips"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBanStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBanStatusResponseParams `json:"Response"`
}

func (r *DescribeBanStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBanStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBanWhiteListRequestParams struct {
	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字 </li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeBanWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字 </li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeBanWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBanWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBanWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBanWhiteListResponseParams struct {
	// 总记录数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 白名单列表
	WhiteList []*BanWhiteListDetail `json:"WhiteList,omitnil" name:"WhiteList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBanWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBanWhiteListResponseParams `json:"Response"`
}

func (r *DescribeBanWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBanWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineAnalysisDataRequestParams struct {
	// 基线策略id
	StrategyId *uint64 `json:"StrategyId,omitnil" name:"StrategyId"`
}

type DescribeBaselineAnalysisDataRequest struct {
	*tchttp.BaseRequest
	
	// 基线策略id
	StrategyId *uint64 `json:"StrategyId,omitnil" name:"StrategyId"`
}

func (r *DescribeBaselineAnalysisDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineAnalysisDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StrategyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaselineAnalysisDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineAnalysisDataResponseParams struct {
	// 最后检测时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestScanTime *string `json:"LatestScanTime,omitnil" name:"LatestScanTime"`

	// 是否全部服务器：1-是 0-否
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsGlobal *uint64 `json:"IsGlobal,omitnil" name:"IsGlobal"`

	// 服务器总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanHostCount *uint64 `json:"ScanHostCount,omitnil" name:"ScanHostCount"`

	// 检测项总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanRuleCount *uint64 `json:"ScanRuleCount,omitnil" name:"ScanRuleCount"`

	// 是否是第一次检测  1是 0不是
	// 注意：此字段可能返回 null，表示取不到有效值。
	IfFirstScan *uint64 `json:"IfFirstScan,omitnil" name:"IfFirstScan"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBaselineAnalysisDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaselineAnalysisDataResponseParams `json:"Response"`
}

func (r *DescribeBaselineAnalysisDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineAnalysisDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineBasicInfoRequestParams struct {
	// 基线名称
	BaselineName *string `json:"BaselineName,omitnil" name:"BaselineName"`
}

type DescribeBaselineBasicInfoRequest struct {
	*tchttp.BaseRequest
	
	// 基线名称
	BaselineName *string `json:"BaselineName,omitnil" name:"BaselineName"`
}

func (r *DescribeBaselineBasicInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineBasicInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BaselineName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaselineBasicInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineBasicInfoResponseParams struct {
	// 基线基础信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	BaselineBasicInfoList []*BaselineBasicInfo `json:"BaselineBasicInfoList,omitnil" name:"BaselineBasicInfoList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBaselineBasicInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaselineBasicInfoResponseParams `json:"Response"`
}

func (r *DescribeBaselineBasicInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineBasicInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineDetailRequestParams struct {
	// 基线id
	BaselineId *uint64 `json:"BaselineId,omitnil" name:"BaselineId"`
}

type DescribeBaselineDetailRequest struct {
	*tchttp.BaseRequest
	
	// 基线id
	BaselineId *uint64 `json:"BaselineId,omitnil" name:"BaselineId"`
}

func (r *DescribeBaselineDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BaselineId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaselineDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineDetailResponseParams struct {
	// 基线详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	BaselineDetail *BaselineDetail `json:"BaselineDetail,omitnil" name:"BaselineDetail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBaselineDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaselineDetailResponseParams `json:"Response"`
}

func (r *DescribeBaselineDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineDetectListRequestParams struct {
	// <li>PolicyName - string - 是否必填：否 - 策略名称</li>
	// <li>PolicyDetectStatus - int - 是否必填：否 - 1:检测中 2:检测完成</li>
	// <li>FirstTime - string - 是否必填：否 - 开始时间</li>
	// <li>LastTime - string - 是否必填：否 - 结束时间</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 限制条数,默认10,最大100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式: [ASC:升序|DESC:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序列: [HostCount|StartTime|StopTime]
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeBaselineDetectListRequest struct {
	*tchttp.BaseRequest
	
	// <li>PolicyName - string - 是否必填：否 - 策略名称</li>
	// <li>PolicyDetectStatus - int - 是否必填：否 - 1:检测中 2:检测完成</li>
	// <li>FirstTime - string - 是否必填：否 - 开始时间</li>
	// <li>LastTime - string - 是否必填：否 - 结束时间</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 限制条数,默认10,最大100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式: [ASC:升序|DESC:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序列: [HostCount|StartTime|StopTime]
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeBaselineDetectListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineDetectListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaselineDetectListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineDetectListResponseParams struct {
	// 无
	List []*BaselinePolicyDetect `json:"List,omitnil" name:"List"`

	// 总数
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBaselineDetectListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaselineDetectListResponseParams `json:"Response"`
}

func (r *DescribeBaselineDetectListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineDetectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineDetectOverviewRequestParams struct {
	// 策略Id
	PolicyId *int64 `json:"PolicyId,omitnil" name:"PolicyId"`
}

type DescribeBaselineDetectOverviewRequest struct {
	*tchttp.BaseRequest
	
	// 策略Id
	PolicyId *int64 `json:"PolicyId,omitnil" name:"PolicyId"`
}

func (r *DescribeBaselineDetectOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineDetectOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaselineDetectOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineDetectOverviewResponseParams struct {
	// 检测服务器数
	HostCount *int64 `json:"HostCount,omitnil" name:"HostCount"`

	// 检测项
	ItemCount *int64 `json:"ItemCount,omitnil" name:"ItemCount"`

	// 检测策略项
	PolicyCount *int64 `json:"PolicyCount,omitnil" name:"PolicyCount"`

	// 通过率*100%
	PassRate *int64 `json:"PassRate,omitnil" name:"PassRate"`

	// 最近一次检测通过个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestPassCount *int64 `json:"LatestPassCount,omitnil" name:"LatestPassCount"`

	// 最近一次检测未通过个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestNotPassCount *int64 `json:"LatestNotPassCount,omitnil" name:"LatestNotPassCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBaselineDetectOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaselineDetectOverviewResponseParams `json:"Response"`
}

func (r *DescribeBaselineDetectOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineDetectOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineDownloadListRequestParams struct {
	// <li>Status - int - 是否必填：否 - 0:导出中 1:已完成</li>
	// <li>StartTime - string - 是否必填：否 - 开始时间</li>
	// <li>EndTime - string - 是否必填：否 - 结束时间</li>
	// <li>TaskName - string - 是否必填：否 - 任务名称</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 限制条数,默认10,最大100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式: [ASC:升序|DESC:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序列: [StartTime|EndTime]
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeBaselineDownloadListRequest struct {
	*tchttp.BaseRequest
	
	// <li>Status - int - 是否必填：否 - 0:导出中 1:已完成</li>
	// <li>StartTime - string - 是否必填：否 - 开始时间</li>
	// <li>EndTime - string - 是否必填：否 - 结束时间</li>
	// <li>TaskName - string - 是否必填：否 - 任务名称</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 限制条数,默认10,最大100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式: [ASC:升序|DESC:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序列: [StartTime|EndTime]
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeBaselineDownloadListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineDownloadListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaselineDownloadListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineDownloadListResponseParams struct {
	// 无
	List []*BaselineDownload `json:"List,omitnil" name:"List"`

	// 总数
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBaselineDownloadListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaselineDownloadListResponseParams `json:"Response"`
}

func (r *DescribeBaselineDownloadListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineDownloadListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineEffectHostListRequestParams struct {
	// 分页参数 最大100条
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页参数
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 基线id
	BaselineId *uint64 `json:"BaselineId,omitnil" name:"BaselineId"`

	// 过滤条件。
	// <li>AliasName- String- 主机别名</li>
	// <li>Status- Uint- 1已通过  0未通过 5检测中</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 策略id
	StrategyId *uint64 `json:"StrategyId,omitnil" name:"StrategyId"`

	// 主机uuid数组
	UuidList []*string `json:"UuidList,omitnil" name:"UuidList"`
}

type DescribeBaselineEffectHostListRequest struct {
	*tchttp.BaseRequest
	
	// 分页参数 最大100条
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页参数
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 基线id
	BaselineId *uint64 `json:"BaselineId,omitnil" name:"BaselineId"`

	// 过滤条件。
	// <li>AliasName- String- 主机别名</li>
	// <li>Status- Uint- 1已通过  0未通过 5检测中</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 策略id
	StrategyId *uint64 `json:"StrategyId,omitnil" name:"StrategyId"`

	// 主机uuid数组
	UuidList []*string `json:"UuidList,omitnil" name:"UuidList"`
}

func (r *DescribeBaselineEffectHostListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineEffectHostListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "BaselineId")
	delete(f, "Filters")
	delete(f, "StrategyId")
	delete(f, "UuidList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaselineEffectHostListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineEffectHostListResponseParams struct {
	// 记录总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 影响服务器列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	EffectHostList []*BaselineEffectHost `json:"EffectHostList,omitnil" name:"EffectHostList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBaselineEffectHostListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaselineEffectHostListResponseParams `json:"Response"`
}

func (r *DescribeBaselineEffectHostListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineEffectHostListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineFixListRequestParams struct {
	// <li>ItemName- string - 是否必填：否 - 项名称</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 限制条数,默认10,最大100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式: [ASC:升序|DESC:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序列: [CreateTime|MoifyTime|FixTime]
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeBaselineFixListRequest struct {
	*tchttp.BaseRequest
	
	// <li>ItemName- string - 是否必填：否 - 项名称</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 限制条数,默认10,最大100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式: [ASC:升序|DESC:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序列: [CreateTime|MoifyTime|FixTime]
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeBaselineFixListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineFixListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaselineFixListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineFixListResponseParams struct {
	// 无
	List []*BaselineFix `json:"List,omitnil" name:"List"`

	// 总数
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBaselineFixListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaselineFixListResponseParams `json:"Response"`
}

func (r *DescribeBaselineFixListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineFixListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineHostDetectListRequestParams struct {
	// <li>PolicyId - int64 - 是否必填：否 - 策略Id</li>
	// <li>HostName - string - 是否必填：否 - 主机名称</i>
	// <li>HostIp - string - 是否必填：否 - 主机Ip</i>
	// <li>ItemId - int64 - 是否必填：否 - 项Id</i>
	// <li>RuleId - int64 - 是否必填：否 - 规则Id</li>
	// <li>DetectStatus - int - 是否必填：否 - 检测状态</li>
	// <li>Level - int - 是否必填：否 - 风险等级</li>
	// <li>StartTime - string - 是否必填：否 - 开时时间</li>
	// <li>EndTime - string - 是否必填：否 - 结束时间</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 限制条数,默认10,最大100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式: [ASC:升序|DESC:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序列: [LastTime|ItemCount|PassedItemCount|NotPassedItemCount|FirstTime]
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeBaselineHostDetectListRequest struct {
	*tchttp.BaseRequest
	
	// <li>PolicyId - int64 - 是否必填：否 - 策略Id</li>
	// <li>HostName - string - 是否必填：否 - 主机名称</i>
	// <li>HostIp - string - 是否必填：否 - 主机Ip</i>
	// <li>ItemId - int64 - 是否必填：否 - 项Id</i>
	// <li>RuleId - int64 - 是否必填：否 - 规则Id</li>
	// <li>DetectStatus - int - 是否必填：否 - 检测状态</li>
	// <li>Level - int - 是否必填：否 - 风险等级</li>
	// <li>StartTime - string - 是否必填：否 - 开时时间</li>
	// <li>EndTime - string - 是否必填：否 - 结束时间</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 限制条数,默认10,最大100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式: [ASC:升序|DESC:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序列: [LastTime|ItemCount|PassedItemCount|NotPassedItemCount|FirstTime]
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeBaselineHostDetectListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineHostDetectListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaselineHostDetectListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineHostDetectListResponseParams struct {
	// 无
	List []*BaselineHostDetect `json:"List,omitnil" name:"List"`

	// 总数
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBaselineHostDetectListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaselineHostDetectListResponseParams `json:"Response"`
}

func (r *DescribeBaselineHostDetectListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineHostDetectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineHostIgnoreListRequestParams struct {
	// 请求的规则
	RuleID *int64 `json:"RuleID,omitnil" name:"RuleID"`

	// 限制条数,默认10,最大100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeBaselineHostIgnoreListRequest struct {
	*tchttp.BaseRequest
	
	// 请求的规则
	RuleID *int64 `json:"RuleID,omitnil" name:"RuleID"`

	// 限制条数,默认10,最大100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeBaselineHostIgnoreListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineHostIgnoreListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleID")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaselineHostIgnoreListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineHostIgnoreListResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 无
	List []*BaselineHost `json:"List,omitnil" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBaselineHostIgnoreListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaselineHostIgnoreListResponseParams `json:"Response"`
}

func (r *DescribeBaselineHostIgnoreListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineHostIgnoreListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineHostRiskTopRequestParams struct {
	// 策略ID
	PolicyId *int64 `json:"PolicyId,omitnil" name:"PolicyId"`
}

type DescribeBaselineHostRiskTopRequest struct {
	*tchttp.BaseRequest
	
	// 策略ID
	PolicyId *int64 `json:"PolicyId,omitnil" name:"PolicyId"`
}

func (r *DescribeBaselineHostRiskTopRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineHostRiskTopRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaselineHostRiskTopRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineHostRiskTopResponseParams struct {
	// 风险主机top5
	HostRiskTop5 []*HostRiskLevelCount `json:"HostRiskTop5,omitnil" name:"HostRiskTop5"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBaselineHostRiskTopResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaselineHostRiskTopResponseParams `json:"Response"`
}

func (r *DescribeBaselineHostRiskTopResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineHostRiskTopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineHostTopRequestParams struct {
	// 动态top值
	Top *uint64 `json:"Top,omitnil" name:"Top"`

	// 策略id
	StrategyId *uint64 `json:"StrategyId,omitnil" name:"StrategyId"`
}

type DescribeBaselineHostTopRequest struct {
	*tchttp.BaseRequest
	
	// 动态top值
	Top *uint64 `json:"Top,omitnil" name:"Top"`

	// 策略id
	StrategyId *uint64 `json:"StrategyId,omitnil" name:"StrategyId"`
}

func (r *DescribeBaselineHostTopRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineHostTopRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Top")
	delete(f, "StrategyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaselineHostTopRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineHostTopResponseParams struct {
	// 主机基线策略事件Top
	// 注意：此字段可能返回 null，表示取不到有效值。
	BaselineHostTopList []*BaselineHostTopList `json:"BaselineHostTopList,omitnil" name:"BaselineHostTopList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBaselineHostTopResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaselineHostTopResponseParams `json:"Response"`
}

func (r *DescribeBaselineHostTopResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineHostTopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineItemDetectListRequestParams struct {
	// <li>HostId - string - 是否必填：否 - 主机Id</li>
	// <li>RuleId - int64 - 是否必填：否 - 规则Id</li>
	// <li>PolicyId - int64 - 是否必填：否 - 规则Id</li>
	// <li>ItemName - string - 是否必填：否 - 项名称</li>
	// <li>DetectStatus - int - 是否必填：否 - 检测状态</li>
	// <li>Level - int - 是否必填：否 - 风险等级</li>
	// <li>StartTime - string - 是否必填：否 - 开始时间</li>
	// <li>EndTime - string - 是否必填：否 - 结束时间</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 限制条数,默认10,最大100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式: [ASC:升序|DESC:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序列: [HostCount|FirstTime|LastTime]
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeBaselineItemDetectListRequest struct {
	*tchttp.BaseRequest
	
	// <li>HostId - string - 是否必填：否 - 主机Id</li>
	// <li>RuleId - int64 - 是否必填：否 - 规则Id</li>
	// <li>PolicyId - int64 - 是否必填：否 - 规则Id</li>
	// <li>ItemName - string - 是否必填：否 - 项名称</li>
	// <li>DetectStatus - int - 是否必填：否 - 检测状态</li>
	// <li>Level - int - 是否必填：否 - 风险等级</li>
	// <li>StartTime - string - 是否必填：否 - 开始时间</li>
	// <li>EndTime - string - 是否必填：否 - 结束时间</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 限制条数,默认10,最大100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式: [ASC:升序|DESC:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序列: [HostCount|FirstTime|LastTime]
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeBaselineItemDetectListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineItemDetectListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaselineItemDetectListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineItemDetectListResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 无
	List []*BaselineItemDetect `json:"List,omitnil" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBaselineItemDetectListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaselineItemDetectListResponseParams `json:"Response"`
}

func (r *DescribeBaselineItemDetectListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineItemDetectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineItemIgnoreListRequestParams struct {
	// 忽略规则ID
	RuleID *int64 `json:"RuleID,omitnil" name:"RuleID"`

	// 限制条数,默认10,最大100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 请求偏移默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// <li>CatgoryId - int64 - 是否必填：否 - 规则Id</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 排序方式 [ASC:升序|DESC:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序列 [ID]
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeBaselineItemIgnoreListRequest struct {
	*tchttp.BaseRequest
	
	// 忽略规则ID
	RuleID *int64 `json:"RuleID,omitnil" name:"RuleID"`

	// 限制条数,默认10,最大100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 请求偏移默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// <li>CatgoryId - int64 - 是否必填：否 - 规则Id</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 排序方式 [ASC:升序|DESC:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序列 [ID]
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeBaselineItemIgnoreListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineItemIgnoreListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleID")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaselineItemIgnoreListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineItemIgnoreListResponseParams struct {
	// 无
	List []*BaselineItemInfo `json:"List,omitnil" name:"List"`

	// 总数
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBaselineItemIgnoreListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaselineItemIgnoreListResponseParams `json:"Response"`
}

func (r *DescribeBaselineItemIgnoreListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineItemIgnoreListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineItemInfoRequestParams struct {
	// <li>ItemId - int64 - 是否必填：否 - 项Id</i>
	// <li>PolicyId - int64 - 是否必填：否 - 项Id</i>
	// <li>Level - int - 是否必填：否 - 风险等级</li>
	// <li>ItemName - string - 是否必填：否 - 检测项名字</li>
	// <li>RuleId - int - 是否必填：否 - 规则Id</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 限制条数,默认10,最大100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式: [ASC:升序|DESC:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序列
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeBaselineItemInfoRequest struct {
	*tchttp.BaseRequest
	
	// <li>ItemId - int64 - 是否必填：否 - 项Id</i>
	// <li>PolicyId - int64 - 是否必填：否 - 项Id</i>
	// <li>Level - int - 是否必填：否 - 风险等级</li>
	// <li>ItemName - string - 是否必填：否 - 检测项名字</li>
	// <li>RuleId - int - 是否必填：否 - 规则Id</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 限制条数,默认10,最大100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式: [ASC:升序|DESC:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序列
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeBaselineItemInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineItemInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaselineItemInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineItemInfoResponseParams struct {
	// 结果列表
	List []*BaselineItemInfo `json:"List,omitnil" name:"List"`

	// 总条目数
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBaselineItemInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaselineItemInfoResponseParams `json:"Response"`
}

func (r *DescribeBaselineItemInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineItemInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineItemListRequestParams struct {
	// <li>PolicyId - int64 - 是否必填：否 - 策略Id</li>
	// <li>RuleId - int64 - 是否必填：否 - 规则Id</li>
	// <li>HostId - string - 是否必填：否 - 主机Id</li>
	// <li>HostName - string - 是否必填：否 - 主机名</li>
	// <li>HostIp - string - 是否必填：否 - 主机IP</li>
	// <li>ItemId - String - 是否必填：否 - 检测项Id</li>
	// <li>ItemName - String - 是否必填：否 - 项名称</li>
	// <li>DetectStatus - int - 是否必填：否 - 检测状态[0:未通过|3:通过|5:检测中]</li>
	// <li>Level - int - 是否必填：否 - 风险等级</li>
	// <li>StartTime - string - 是否必填：否 - 开始时间</li>
	// <li>EndTime - string - 是否必填：否 - 结束时间</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 限制条数,默认10,最大100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式: [ASC:升序|DESC:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序列
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeBaselineItemListRequest struct {
	*tchttp.BaseRequest
	
	// <li>PolicyId - int64 - 是否必填：否 - 策略Id</li>
	// <li>RuleId - int64 - 是否必填：否 - 规则Id</li>
	// <li>HostId - string - 是否必填：否 - 主机Id</li>
	// <li>HostName - string - 是否必填：否 - 主机名</li>
	// <li>HostIp - string - 是否必填：否 - 主机IP</li>
	// <li>ItemId - String - 是否必填：否 - 检测项Id</li>
	// <li>ItemName - String - 是否必填：否 - 项名称</li>
	// <li>DetectStatus - int - 是否必填：否 - 检测状态[0:未通过|3:通过|5:检测中]</li>
	// <li>Level - int - 是否必填：否 - 风险等级</li>
	// <li>StartTime - string - 是否必填：否 - 开始时间</li>
	// <li>EndTime - string - 是否必填：否 - 结束时间</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 限制条数,默认10,最大100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式: [ASC:升序|DESC:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序列
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeBaselineItemListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineItemListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaselineItemListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineItemListResponseParams struct {
	// 无
	List []*BaselineItem `json:"List,omitnil" name:"List"`

	// 总数
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBaselineItemListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaselineItemListResponseParams `json:"Response"`
}

func (r *DescribeBaselineItemListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineItemListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineItemRiskTopRequestParams struct {
	// 策略ID
	PolicyId *int64 `json:"PolicyId,omitnil" name:"PolicyId"`
}

type DescribeBaselineItemRiskTopRequest struct {
	*tchttp.BaseRequest
	
	// 策略ID
	PolicyId *int64 `json:"PolicyId,omitnil" name:"PolicyId"`
}

func (r *DescribeBaselineItemRiskTopRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineItemRiskTopRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaselineItemRiskTopRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineItemRiskTopResponseParams struct {
	// 结果数组
	RiskItemTop5 []*BaselineRiskItem `json:"RiskItemTop5,omitnil" name:"RiskItemTop5"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBaselineItemRiskTopResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaselineItemRiskTopResponseParams `json:"Response"`
}

func (r *DescribeBaselineItemRiskTopResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineItemRiskTopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineListRequestParams struct {
	// 分页参数 最大100条
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页参数
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>StrategyId- Uint64 - 基线策略id</li>
	// <li>Status - Uint64 - 处理状态1已通过 0未通过</li>
	// <li>Level - Uint64[] - 处理状态1已通过 0未通过</li>BaselineName 
	// <li>BaselineName  - String - 基线名称</li>
	// <li>Quuid- String - 主机quuid</li>
	// <li>Uuid- String - 主机uuid</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`
}

type DescribeBaselineListRequest struct {
	*tchttp.BaseRequest
	
	// 分页参数 最大100条
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页参数
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>StrategyId- Uint64 - 基线策略id</li>
	// <li>Status - Uint64 - 处理状态1已通过 0未通过</li>
	// <li>Level - Uint64[] - 处理状态1已通过 0未通过</li>BaselineName 
	// <li>BaselineName  - String - 基线名称</li>
	// <li>Quuid- String - 主机quuid</li>
	// <li>Uuid- String - 主机uuid</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeBaselineListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaselineListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineListResponseParams struct {
	// 基线信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	BaselineList []*BaselineInfo `json:"BaselineList,omitnil" name:"BaselineList"`

	// 分页查询记录总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBaselineListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaselineListResponseParams `json:"Response"`
}

func (r *DescribeBaselineListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselinePolicyListRequestParams struct {
	// <li>PolicyName - String - 是否必填：否 - 策略名称</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 限制条数,默认10,最大100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式: [ASC:升序|DESC:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序列: [RuleCount|ItemCount|HostCount]
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeBaselinePolicyListRequest struct {
	*tchttp.BaseRequest
	
	// <li>PolicyName - String - 是否必填：否 - 策略名称</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 限制条数,默认10,最大100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式: [ASC:升序|DESC:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序列: [RuleCount|ItemCount|HostCount]
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeBaselinePolicyListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselinePolicyListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaselinePolicyListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselinePolicyListResponseParams struct {
	// 无
	List []*BaselinePolicy `json:"List,omitnil" name:"List"`

	// 总数
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBaselinePolicyListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaselinePolicyListResponseParams `json:"Response"`
}

func (r *DescribeBaselinePolicyListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselinePolicyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineRuleCategoryListRequestParams struct {

}

type DescribeBaselineRuleCategoryListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeBaselineRuleCategoryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineRuleCategoryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaselineRuleCategoryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineRuleCategoryListResponseParams struct {
	// 无
	List []*BaselineCategory `json:"List,omitnil" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBaselineRuleCategoryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaselineRuleCategoryListResponseParams `json:"Response"`
}

func (r *DescribeBaselineRuleCategoryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineRuleCategoryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineRuleDetectListRequestParams struct {
	// <li>PolicyId - int64 - 是否必填：否 - 策略Id</li>
	// <li>ItemId - int64 - 是否必填：否 - 策略Id</li>
	// <li>RuleName - string - 是否必填：否 - 规则名称</li>
	// <li>DetectStatus - int - 是否必填：否 - 检测状态</li>
	// <li>StartTime - string - 是否必填：否 - 开时时间</li>
	// <li>EndTime - string - 是否必填：否 - 结束时间</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 限制条数,默认10,最大100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式: [ASC:升序|DESC:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序列: [HostCount|FirstTime|LastTime]
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeBaselineRuleDetectListRequest struct {
	*tchttp.BaseRequest
	
	// <li>PolicyId - int64 - 是否必填：否 - 策略Id</li>
	// <li>ItemId - int64 - 是否必填：否 - 策略Id</li>
	// <li>RuleName - string - 是否必填：否 - 规则名称</li>
	// <li>DetectStatus - int - 是否必填：否 - 检测状态</li>
	// <li>StartTime - string - 是否必填：否 - 开时时间</li>
	// <li>EndTime - string - 是否必填：否 - 结束时间</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 限制条数,默认10,最大100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式: [ASC:升序|DESC:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序列: [HostCount|FirstTime|LastTime]
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeBaselineRuleDetectListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineRuleDetectListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaselineRuleDetectListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineRuleDetectListResponseParams struct {
	// 无
	List []*BaselineRuleDetect `json:"List,omitnil" name:"List"`

	// 总数
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBaselineRuleDetectListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaselineRuleDetectListResponseParams `json:"Response"`
}

func (r *DescribeBaselineRuleDetectListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineRuleDetectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineRuleIgnoreListRequestParams struct {
	// <li>RuleName - String - 是否必填：否 - 规则名称</li>
	// <li>ItemId- int - 是否必填：否 - 检测项Id</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 限制条数,默认10,最大100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式: [ASC:升序|DESC:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序列: [HostCount]
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeBaselineRuleIgnoreListRequest struct {
	*tchttp.BaseRequest
	
	// <li>RuleName - String - 是否必填：否 - 规则名称</li>
	// <li>ItemId- int - 是否必填：否 - 检测项Id</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 限制条数,默认10,最大100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式: [ASC:升序|DESC:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序列: [HostCount]
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeBaselineRuleIgnoreListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineRuleIgnoreListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaselineRuleIgnoreListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineRuleIgnoreListResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 列表
	List []*BaselineRule `json:"List,omitnil" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBaselineRuleIgnoreListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaselineRuleIgnoreListResponseParams `json:"Response"`
}

func (r *DescribeBaselineRuleIgnoreListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineRuleIgnoreListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineRuleListRequestParams struct {
	// <li>RuleName - String - 是否必填：否 - 规则名称</li>
	// <li>CategoryId - int64 - 是否必填：否 自定义筛选为-1 - 规则分类</li>
	// <li>RuleType - int - 是否必填：否 0:系统 1:自定义 - 规则类型</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 限制条数,默认10,最大100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式: [ASC:升序|DESC:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序列
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeBaselineRuleListRequest struct {
	*tchttp.BaseRequest
	
	// <li>RuleName - String - 是否必填：否 - 规则名称</li>
	// <li>CategoryId - int64 - 是否必填：否 自定义筛选为-1 - 规则分类</li>
	// <li>RuleType - int - 是否必填：否 0:系统 1:自定义 - 规则类型</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 限制条数,默认10,最大100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式: [ASC:升序|DESC:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序列
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeBaselineRuleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineRuleListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaselineRuleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineRuleListResponseParams struct {
	// 无
	List []*BaselineRule `json:"List,omitnil" name:"List"`

	// 总数
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBaselineRuleListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaselineRuleListResponseParams `json:"Response"`
}

func (r *DescribeBaselineRuleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineRuleRequestParams struct {
	// 基线id
	BaselineId *uint64 `json:"BaselineId,omitnil" name:"BaselineId"`

	// 分页参数 最大100条
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页参数
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 危害等级
	Level []*uint64 `json:"Level,omitnil" name:"Level"`

	// 状态
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 主机quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

type DescribeBaselineRuleRequest struct {
	*tchttp.BaseRequest
	
	// 基线id
	BaselineId *uint64 `json:"BaselineId,omitnil" name:"BaselineId"`

	// 分页参数 最大100条
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页参数
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 危害等级
	Level []*uint64 `json:"Level,omitnil" name:"Level"`

	// 状态
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 主机quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

func (r *DescribeBaselineRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BaselineId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Level")
	delete(f, "Status")
	delete(f, "Quuid")
	delete(f, "Uuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaselineRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineRuleResponseParams struct {
	// 分页查询记录总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 基线检测项列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	BaselineRuleList []*BaselineRuleInfo `json:"BaselineRuleList,omitnil" name:"BaselineRuleList"`

	// 是否显示说明列：true-是，false-否
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShowRuleRemark *bool `json:"ShowRuleRemark,omitnil" name:"ShowRuleRemark"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBaselineRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaselineRuleResponseParams `json:"Response"`
}

func (r *DescribeBaselineRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineScanScheduleRequestParams struct {
	// 任务id
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`
}

type DescribeBaselineScanScheduleRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`
}

func (r *DescribeBaselineScanScheduleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineScanScheduleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaselineScanScheduleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineScanScheduleResponseParams struct {
	// 检测进度(百分比)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Schedule *uint64 `json:"Schedule,omitnil" name:"Schedule"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBaselineScanScheduleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaselineScanScheduleResponseParams `json:"Response"`
}

func (r *DescribeBaselineScanScheduleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineScanScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineStrategyDetailRequestParams struct {
	// 用户基线策略id
	StrategyId *uint64 `json:"StrategyId,omitnil" name:"StrategyId"`
}

type DescribeBaselineStrategyDetailRequest struct {
	*tchttp.BaseRequest
	
	// 用户基线策略id
	StrategyId *uint64 `json:"StrategyId,omitnil" name:"StrategyId"`
}

func (r *DescribeBaselineStrategyDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineStrategyDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StrategyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaselineStrategyDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineStrategyDetailResponseParams struct {
	// 策略扫描通过率
	// 注意：此字段可能返回 null，表示取不到有效值。
	PassRate *uint64 `json:"PassRate,omitnil" name:"PassRate"`

	// 策略名
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrategyName *string `json:"StrategyName,omitnil" name:"StrategyName"`

	// 策略扫描周期(天)
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanCycle *string `json:"ScanCycle,omitnil" name:"ScanCycle"`

	// 定期检测时间, 该时间下发扫描
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanAt *string `json:"ScanAt,omitnil" name:"ScanAt"`

	// 扫描范围是否全部服务器, 1:是  0:否, 为1则为全部专业版主机
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsGlobal *uint64 `json:"IsGlobal,omitnil" name:"IsGlobal"`

	// 云服务器类型：
	// cvm：腾讯云服务器
	// bm：裸金属
	// ecm：边缘计算主机
	// lh: 轻量应用服务器
	// ohter: 混合云机器
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineType *string `json:"MachineType,omitnil" name:"MachineType"`

	// 主机地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// 用户该策略下的所有主机id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Quuids []*string `json:"Quuids,omitnil" name:"Quuids"`

	// 用户该策略下所有的基线id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryIds []*string `json:"CategoryIds,omitnil" name:"CategoryIds"`

	// 1 表示扫描过, 0没扫描过
	// 注意：此字段可能返回 null，表示取不到有效值。
	IfScanned *uint64 `json:"IfScanned,omitnil" name:"IfScanned"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBaselineStrategyDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaselineStrategyDetailResponseParams `json:"Response"`
}

func (r *DescribeBaselineStrategyDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineStrategyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineStrategyListRequestParams struct {
	// 分页参数 最大100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页参数
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 规则开关，1：打开 0：关闭  2:全部
	Enabled *uint64 `json:"Enabled,omitnil" name:"Enabled"`
}

type DescribeBaselineStrategyListRequest struct {
	*tchttp.BaseRequest
	
	// 分页参数 最大100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页参数
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 规则开关，1：打开 0：关闭  2:全部
	Enabled *uint64 `json:"Enabled,omitnil" name:"Enabled"`
}

func (r *DescribeBaselineStrategyListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineStrategyListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Enabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaselineStrategyListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineStrategyListResponseParams struct {
	// 分页查询记录的总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 用户策略信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrategyList []*Strategy `json:"StrategyList,omitnil" name:"StrategyList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBaselineStrategyListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaselineStrategyListResponseParams `json:"Response"`
}

func (r *DescribeBaselineStrategyListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineStrategyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineTopRequestParams struct {
	// 动态top值
	Top *uint64 `json:"Top,omitnil" name:"Top"`

	// 策略id
	StrategyId *uint64 `json:"StrategyId,omitnil" name:"StrategyId"`
}

type DescribeBaselineTopRequest struct {
	*tchttp.BaseRequest
	
	// 动态top值
	Top *uint64 `json:"Top,omitnil" name:"Top"`

	// 策略id
	StrategyId *uint64 `json:"StrategyId,omitnil" name:"StrategyId"`
}

func (r *DescribeBaselineTopRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineTopRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Top")
	delete(f, "StrategyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaselineTopRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineTopResponseParams struct {
	// 检测项Top列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleTopList []*BaselineRuleTopInfo `json:"RuleTopList,omitnil" name:"RuleTopList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBaselineTopResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaselineTopResponseParams `json:"Response"`
}

func (r *DescribeBaselineTopResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineTopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineWeakPasswordListRequestParams struct {
	// <li>WeakPassword - string - 是否必填：否 - 弱口令</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 限制条数,默认10,最大100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式 [ASC:升序|DESC:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序列 [CreateTime|ModifyTime]
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeBaselineWeakPasswordListRequest struct {
	*tchttp.BaseRequest
	
	// <li>WeakPassword - string - 是否必填：否 - 弱口令</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 限制条数,默认10,最大100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式 [ASC:升序|DESC:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序列 [CreateTime|ModifyTime]
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeBaselineWeakPasswordListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineWeakPasswordListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaselineWeakPasswordListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaselineWeakPasswordListResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 无
	List []*BaselineWeakPassword `json:"List,omitnil" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBaselineWeakPasswordListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaselineWeakPasswordListResponseParams `json:"Response"`
}

func (r *DescribeBaselineWeakPasswordListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaselineWeakPasswordListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBashEventsInfoNewRequestParams struct {
	// 事件id
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

type DescribeBashEventsInfoNewRequest struct {
	*tchttp.BaseRequest
	
	// 事件id
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

func (r *DescribeBashEventsInfoNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBashEventsInfoNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBashEventsInfoNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBashEventsInfoNewResponseParams struct {
	// 事件详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	BashEventsInfo *BashEventsInfoNew `json:"BashEventsInfo,omitnil" name:"BashEventsInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBashEventsInfoNewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBashEventsInfoNewResponseParams `json:"Response"`
}

func (r *DescribeBashEventsInfoNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBashEventsInfoNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBashEventsNewRequestParams struct {
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件。
	// <li>HostName - String - 是否必填：否 - 主机名</li>
	// <li>Hostip - String - 是否必填：否 - 主机内网IP</li>
	// <li>HostIp - String - 是否必填：否 - 主机内网IP</li>
	// <li>RuleCategory - Int - 是否必填：否 - 策略类型,全部或者单选(0:系统 1:用户)</li>
	// <li>RuleName - String - 是否必填：否 - 策略名称</li>
	// <li>RuleLevel - Int - 是否必填：否 - 威胁等级,可以多选</li>
	// <li>Status - Int - 是否必填：否 - 处理状态,可多选(0:待处理 1:已处理 2:已加白  3:已忽略 4:已删除 5:已拦截)</li>
	// <li>DetectBy - Int - 是否必填：否 - 数据来源,可多选(0:bash日志 1:实时监控)</li>
	// <li>StartTime - String - 是否必填：否 - 开始时间</li>
	// <li>EndTime - String - 是否必填：否 - 结束时间</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式：根据请求次数排序：asc-升序/desc-降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段：CreateTime-发生时间。ModifyTime-处理时间
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeBashEventsNewRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件。
	// <li>HostName - String - 是否必填：否 - 主机名</li>
	// <li>Hostip - String - 是否必填：否 - 主机内网IP</li>
	// <li>HostIp - String - 是否必填：否 - 主机内网IP</li>
	// <li>RuleCategory - Int - 是否必填：否 - 策略类型,全部或者单选(0:系统 1:用户)</li>
	// <li>RuleName - String - 是否必填：否 - 策略名称</li>
	// <li>RuleLevel - Int - 是否必填：否 - 威胁等级,可以多选</li>
	// <li>Status - Int - 是否必填：否 - 处理状态,可多选(0:待处理 1:已处理 2:已加白  3:已忽略 4:已删除 5:已拦截)</li>
	// <li>DetectBy - Int - 是否必填：否 - 数据来源,可多选(0:bash日志 1:实时监控)</li>
	// <li>StartTime - String - 是否必填：否 - 开始时间</li>
	// <li>EndTime - String - 是否必填：否 - 结束时间</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式：根据请求次数排序：asc-升序/desc-降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段：CreateTime-发生时间。ModifyTime-处理时间
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeBashEventsNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBashEventsNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBashEventsNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBashEventsNewResponseParams struct {
	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 高危命令事件列表
	List []*BashEventNew `json:"List,omitnil" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBashEventsNewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBashEventsNewResponseParams `json:"Response"`
}

func (r *DescribeBashEventsNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBashEventsNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBashEventsRequestParams struct {
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件。
	// <li>HostName - String - 是否必填：否 - 主机名</li>
	// <li>Hostip - String - 是否必填：否 - 主机内网IP</li>
	// <li>RuleCategory - Int - 是否必填：否 - 策略类型,全部或者单选(0:系统 1:用户)</li>
	// <li>RuleName - String - 是否必填：否 - 策略名称</li>
	// <li>RuleLevel - Int - 是否必填：否 - 威胁等级,可以多选</li>
	// <li>Status - Int - 是否必填：否 - 处理状态,可多选(0:待处理 1:已处理 2:已加白  3:已忽略 4:已删除 5:已拦截)</li>
	// <li>DetectBy - Int - 是否必填：否 - 数据来源,可多选(0:bash日志 1:实时监控)</li>
	// <li>StartTime - String - 是否必填：否 - 开始时间</li>
	// <li>EndTime - String - 是否必填：否 - 结束时间</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式：根据请求次数排序：asc-升序/desc-降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段：CreateTime-发生时间。ModifyTime-处理时间
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeBashEventsRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件。
	// <li>HostName - String - 是否必填：否 - 主机名</li>
	// <li>Hostip - String - 是否必填：否 - 主机内网IP</li>
	// <li>RuleCategory - Int - 是否必填：否 - 策略类型,全部或者单选(0:系统 1:用户)</li>
	// <li>RuleName - String - 是否必填：否 - 策略名称</li>
	// <li>RuleLevel - Int - 是否必填：否 - 威胁等级,可以多选</li>
	// <li>Status - Int - 是否必填：否 - 处理状态,可多选(0:待处理 1:已处理 2:已加白  3:已忽略 4:已删除 5:已拦截)</li>
	// <li>DetectBy - Int - 是否必填：否 - 数据来源,可多选(0:bash日志 1:实时监控)</li>
	// <li>StartTime - String - 是否必填：否 - 开始时间</li>
	// <li>EndTime - String - 是否必填：否 - 结束时间</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式：根据请求次数排序：asc-升序/desc-降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段：CreateTime-发生时间。ModifyTime-处理时间
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeBashEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBashEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBashEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBashEventsResponseParams struct {
	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 高危命令事件列表
	List []*BashEvent `json:"List,omitnil" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBashEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBashEventsResponseParams `json:"Response"`
}

func (r *DescribeBashEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBashEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBashRulesRequestParams struct {
	// 0-系统规则; 1-用户规则
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 返回数量，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键字(规则名称)</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeBashRulesRequest struct {
	*tchttp.BaseRequest
	
	// 0-系统规则; 1-用户规则
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 返回数量，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键字(规则名称)</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeBashRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBashRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBashRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBashRulesResponseParams struct {
	// 列表内容
	List []*BashRule `json:"List,omitnil" name:"List"`

	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBashRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBashRulesResponseParams `json:"Response"`
}

func (r *DescribeBashRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBashRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBruteAttackListRequestParams struct {
	// 需要返回的数量，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Uuid - String - 是否必填：否 - 主机安全唯一Uuid</li>
	// <li>Quuid - String - 是否必填：否 - 云服务器uuid</li>
	// <li>Status - String - 是否必填：否 - 状态筛选：失败：FAILED 成功：SUCCESS</li>
	// <li>UserName - String - 是否必填：否 - UserName筛选</li>
	// <li>SrcIp - String - 是否必填：否 - 来源ip筛选</li>
	// <li>CreateBeginTime - String - 是否必填：否 - 首次攻击时间筛选，开始时间</li>
	// <li>CreateEndTime - String - 是否必填：否 - 首次攻击时间筛选，结束时间</li>
	// <li>ModifyBeginTime - String - 是否必填：否 - 最近攻击时间筛选，开始时间</li>
	// <li>ModifyEndTime - String - 是否必填：否 - 最近攻击时间筛选，结束时间</li>
	// <li>Banned - String - 是否必填：否 - 阻断状态筛选，多个用","分割：0-未阻断（全局ZK开关关闭），82-未阻断(非专业版)，83-未阻断(已加白名单)，1-已阻断，2-未阻断-程序异常，3-未阻断-内网攻击暂不支持阻断，4-未阻断-安平暂不支持阻断</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 排序方式：根据请求次数排序：asc-升序/desc-降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段：CreateTime-首次攻击时间
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeBruteAttackListRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Uuid - String - 是否必填：否 - 主机安全唯一Uuid</li>
	// <li>Quuid - String - 是否必填：否 - 云服务器uuid</li>
	// <li>Status - String - 是否必填：否 - 状态筛选：失败：FAILED 成功：SUCCESS</li>
	// <li>UserName - String - 是否必填：否 - UserName筛选</li>
	// <li>SrcIp - String - 是否必填：否 - 来源ip筛选</li>
	// <li>CreateBeginTime - String - 是否必填：否 - 首次攻击时间筛选，开始时间</li>
	// <li>CreateEndTime - String - 是否必填：否 - 首次攻击时间筛选，结束时间</li>
	// <li>ModifyBeginTime - String - 是否必填：否 - 最近攻击时间筛选，开始时间</li>
	// <li>ModifyEndTime - String - 是否必填：否 - 最近攻击时间筛选，结束时间</li>
	// <li>Banned - String - 是否必填：否 - 阻断状态筛选，多个用","分割：0-未阻断（全局ZK开关关闭），82-未阻断(非专业版)，83-未阻断(已加白名单)，1-已阻断，2-未阻断-程序异常，3-未阻断-内网攻击暂不支持阻断，4-未阻断-安平暂不支持阻断</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 排序方式：根据请求次数排序：asc-升序/desc-降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段：CreateTime-首次攻击时间
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeBruteAttackListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBruteAttackListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBruteAttackListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBruteAttackListResponseParams struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 密码破解列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	BruteAttackList []*BruteAttackInfo `json:"BruteAttackList,omitnil" name:"BruteAttackList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBruteAttackListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBruteAttackListResponseParams `json:"Response"`
}

func (r *DescribeBruteAttackListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBruteAttackListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBruteAttackRulesRequestParams struct {

}

type DescribeBruteAttackRulesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeBruteAttackRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBruteAttackRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBruteAttackRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBruteAttackRulesResponseParams struct {
	// 爆破阻断规则列表
	Rules []*BruteAttackRuleList `json:"Rules,omitnil" name:"Rules"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBruteAttackRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBruteAttackRulesResponseParams `json:"Response"`
}

func (r *DescribeBruteAttackRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBruteAttackRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClientExceptionRequestParams struct {
	// 客户端异常类型 1:客户端离线，2:客户端卸载
	ExceptionType *int64 `json:"ExceptionType,omitnil" name:"ExceptionType"`

	// 分页的偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 分页单页限制数目，不为0，最大值100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 起始时间 `2006-01-02 15:04:05` 格式
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间 `2006-01-02 15:04:05` 格式
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

type DescribeClientExceptionRequest struct {
	*tchttp.BaseRequest
	
	// 客户端异常类型 1:客户端离线，2:客户端卸载
	ExceptionType *int64 `json:"ExceptionType,omitnil" name:"ExceptionType"`

	// 分页的偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 分页单页限制数目，不为0，最大值100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 起始时间 `2006-01-02 15:04:05` 格式
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间 `2006-01-02 15:04:05` 格式
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

func (r *DescribeClientExceptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClientExceptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExceptionType")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClientExceptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClientExceptionResponseParams struct {
	// 事件总数量
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 事件详情
	Records []*RecordInfo `json:"Records,omitnil" name:"Records"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeClientExceptionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClientExceptionResponseParams `json:"Response"`
}

func (r *DescribeClientExceptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClientExceptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComponentStatisticsRequestParams struct {
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// ComponentName - String - 是否必填：否 - 组件名称
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeComponentStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// ComponentName - String - 是否必填：否 - 组件名称
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeComponentStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComponentStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeComponentStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComponentStatisticsResponseParams struct {
	// 组件统计列表记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 组件统计列表数据数组。
	ComponentStatistics []*ComponentStatistics `json:"ComponentStatistics,omitnil" name:"ComponentStatistics"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeComponentStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeComponentStatisticsResponseParams `json:"Response"`
}

func (r *DescribeComponentStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComponentStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeESAggregationsRequestParams struct {
	// ES聚合条件JSON
	Query *string `json:"Query,omitnil" name:"Query"`
}

type DescribeESAggregationsRequest struct {
	*tchttp.BaseRequest
	
	// ES聚合条件JSON
	Query *string `json:"Query,omitnil" name:"Query"`
}

func (r *DescribeESAggregationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeESAggregationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Query")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeESAggregationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeESAggregationsResponseParams struct {
	// ES聚合结果JSON
	Data *string `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeESAggregationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeESAggregationsResponseParams `json:"Response"`
}

func (r *DescribeESAggregationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeESAggregationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEmergencyResponseListRequestParams struct {
	// 过滤条件。
	// <li>Keyword- String - 是否必填：否 - 关键词过滤，</li>
	// <li>Uuids - String - 是否必填：否 - 主机id过滤</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 需要返回的数量，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序步长
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方法
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段 StartTime，EndTime
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeEmergencyResponseListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>Keyword- String - 是否必填：否 - 关键词过滤，</li>
	// <li>Uuids - String - 是否必填：否 - 主机id过滤</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 需要返回的数量，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序步长
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方法
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段 StartTime，EndTime
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeEmergencyResponseListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEmergencyResponseListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEmergencyResponseListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEmergencyResponseListResponseParams struct {
	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 应急响应列表
	List []*EmergencyResponseInfo `json:"List,omitnil" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeEmergencyResponseListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEmergencyResponseListResponseParams `json:"Response"`
}

func (r *DescribeEmergencyResponseListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEmergencyResponseListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEmergencyVulListRequestParams struct {
	// 返回数量，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>Status - String - 是否必填：是 - 漏洞状态筛选，0//未检测 1有风险 ，2无风险 ，3 检查中展示progress</li>
	// <li>Level - String - 是否必填：否 - 漏洞等级筛选 1:低 2:中 3:高 4:提示</li>
	// <li>VulName- String - 是否必填：否 - 漏洞名称搜索</li>
	// <li>Uuids- String - 是否必填：否 - 主机uuid</li>
	// <li>IsSupportDefense - int- 是否必填：否 - 是否支持防御 0:不支持 1:支持</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 排序方式 desc , asc
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段 PublishDate  LastScanTime HostCount
	By *string `json:"By,omitnil" name:"By"`

	// 是否热点漏洞
	HotspotAttack *bool `json:"HotspotAttack,omitnil" name:"HotspotAttack"`
}

type DescribeEmergencyVulListRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>Status - String - 是否必填：是 - 漏洞状态筛选，0//未检测 1有风险 ，2无风险 ，3 检查中展示progress</li>
	// <li>Level - String - 是否必填：否 - 漏洞等级筛选 1:低 2:中 3:高 4:提示</li>
	// <li>VulName- String - 是否必填：否 - 漏洞名称搜索</li>
	// <li>Uuids- String - 是否必填：否 - 主机uuid</li>
	// <li>IsSupportDefense - int- 是否必填：否 - 是否支持防御 0:不支持 1:支持</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 排序方式 desc , asc
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段 PublishDate  LastScanTime HostCount
	By *string `json:"By,omitnil" name:"By"`

	// 是否热点漏洞
	HotspotAttack *bool `json:"HotspotAttack,omitnil" name:"HotspotAttack"`
}

func (r *DescribeEmergencyVulListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEmergencyVulListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	delete(f, "HotspotAttack")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEmergencyVulListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEmergencyVulListResponseParams struct {
	// 漏洞列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*EmergencyVul `json:"List,omitnil" name:"List"`

	// 漏洞总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 是否存在风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExistsRisk *bool `json:"ExistsRisk,omitnil" name:"ExistsRisk"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeEmergencyVulListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEmergencyVulListResponseParams `json:"Response"`
}

func (r *DescribeEmergencyVulListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEmergencyVulListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventByTableRequestParams struct {
	// 事件表名
	TableName *string `json:"TableName,omitnil" name:"TableName"`

	// 事件表id号
	Ids []*int64 `json:"Ids,omitnil" name:"Ids"`
}

type DescribeEventByTableRequest struct {
	*tchttp.BaseRequest
	
	// 事件表名
	TableName *string `json:"TableName,omitnil" name:"TableName"`

	// 事件表id号
	Ids []*int64 `json:"Ids,omitnil" name:"Ids"`
}

func (r *DescribeEventByTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventByTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableName")
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEventByTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventByTableResponseParams struct {
	// 告警类型，爆破bruteattack，高危命令bash，恶意文件malware，恶意请求risk_dns，本地提权privilege_escalation，反弹shell reverse_shell，内存马java_shell
	Type *string `json:"Type,omitnil" name:"Type"`

	// 事件内容的json编码字符串，字段结构对齐事件表
	Value *string `json:"Value,omitnil" name:"Value"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeEventByTableResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEventByTableResponseParams `json:"Response"`
}

func (r *DescribeEventByTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventByTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExpertServiceListRequestParams struct {
	// 过滤条件。
	// <li>Keyword- String - 是否必填：否 - 关键词过滤，</li>
	// <li>Uuids - String - 是否必填：否 - 主机id过滤</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 需要返回的数量，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序步长
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方法
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段 StartTime，EndTime
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeExpertServiceListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>Keyword- String - 是否必填：否 - 关键词过滤，</li>
	// <li>Uuids - String - 是否必填：否 - 主机id过滤</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 需要返回的数量，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序步长
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方法
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段 StartTime，EndTime
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeExpertServiceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExpertServiceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExpertServiceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExpertServiceListResponseParams struct {
	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 安全管家数据
	List []*SecurityButlerInfo `json:"List,omitnil" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeExpertServiceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExpertServiceListResponseParams `json:"Response"`
}

func (r *DescribeExpertServiceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExpertServiceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExpertServiceOrderListRequestParams struct {
	// <li>InquireType- String - 是否必填：否 - 订单类型过滤，</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 分页条数 最大100条
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页步长
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeExpertServiceOrderListRequest struct {
	*tchttp.BaseRequest
	
	// <li>InquireType- String - 是否必填：否 - 订单类型过滤，</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 分页条数 最大100条
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页步长
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeExpertServiceOrderListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExpertServiceOrderListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExpertServiceOrderListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExpertServiceOrderListResponseParams struct {
	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 订单列表
	List []*ExpertServiceOrderInfo `json:"List,omitnil" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeExpertServiceOrderListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExpertServiceOrderListResponseParams `json:"Response"`
}

func (r *DescribeExpertServiceOrderListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExpertServiceOrderListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExportMachinesRequestParams struct {
	// 云主机类型。
	// <li>CVM：表示虚拟主机</li>
	// <li>BM:  表示黑石物理机</li>
	MachineType *string `json:"MachineType,omitnil" name:"MachineType"`

	// 机器所属地域。如：ap-guangzhou，ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitnil" name:"MachineRegion"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字 </li>
	// <li>Status - String - 是否必填：否 - 客户端在线状态（OFFLINE: 离线 | ONLINE: 在线 | UNINSTALLED：未安装）</li>
	// <li>Version - String  是否必填：否 - 当前防护版本（ PRO_VERSION：专业版 | BASIC_VERSION：基础版）</li>
	// 每个过滤条件只支持一个值，暂不支持多个值“或”关系查询
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 机器所属业务ID列表
	ProjectIds []*uint64 `json:"ProjectIds,omitnil" name:"ProjectIds"`
}

type DescribeExportMachinesRequest struct {
	*tchttp.BaseRequest
	
	// 云主机类型。
	// <li>CVM：表示虚拟主机</li>
	// <li>BM:  表示黑石物理机</li>
	MachineType *string `json:"MachineType,omitnil" name:"MachineType"`

	// 机器所属地域。如：ap-guangzhou，ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitnil" name:"MachineRegion"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字 </li>
	// <li>Status - String - 是否必填：否 - 客户端在线状态（OFFLINE: 离线 | ONLINE: 在线 | UNINSTALLED：未安装）</li>
	// <li>Version - String  是否必填：否 - 当前防护版本（ PRO_VERSION：专业版 | BASIC_VERSION：基础版）</li>
	// 每个过滤条件只支持一个值，暂不支持多个值“或”关系查询
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 机器所属业务ID列表
	ProjectIds []*uint64 `json:"ProjectIds,omitnil" name:"ProjectIds"`
}

func (r *DescribeExportMachinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExportMachinesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MachineType")
	delete(f, "MachineRegion")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "ProjectIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExportMachinesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExportMachinesResponseParams struct {
	// 任务ID,需要到接口“异步导出任务”ExportTasks获取DownloadUrl下载地址
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeExportMachinesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExportMachinesResponseParams `json:"Response"`
}

func (r *DescribeExportMachinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExportMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileTamperEventsRequestParams struct {
	// 过滤条件。
	// <li>Status - String - 是否必填：否 - 处理状态  0 -- 待处理 1 -- 已加白 2 -- 已删除 3 - 已忽略</li>
	// <li>ModifyTime - String - 是否必填：否 - 最近发生时间</li>
	// <li>Uuid- String - 是否必填：否 - 主机uuid查询</li>
	// <li>RuleCategory- string - 是否必填：否 - 规则类别 0 系统规则 1 自定义规则</li>
	// <li>FileAction- string - 是否必填：否 - 威胁行为 read 读取文件 write 写文件</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式 ASC,DESC
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段 CreateTime、ModifyTime
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeFileTamperEventsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>Status - String - 是否必填：否 - 处理状态  0 -- 待处理 1 -- 已加白 2 -- 已删除 3 - 已忽略</li>
	// <li>ModifyTime - String - 是否必填：否 - 最近发生时间</li>
	// <li>Uuid- String - 是否必填：否 - 主机uuid查询</li>
	// <li>RuleCategory- string - 是否必填：否 - 规则类别 0 系统规则 1 自定义规则</li>
	// <li>FileAction- string - 是否必填：否 - 威胁行为 read 读取文件 write 写文件</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式 ASC,DESC
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段 CreateTime、ModifyTime
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeFileTamperEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileTamperEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileTamperEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileTamperEventsResponseParams struct {
	// 核心文件事件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*FileTamperEvent `json:"List,omitnil" name:"List"`

	// 数据总条数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeFileTamperEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileTamperEventsResponseParams `json:"Response"`
}

func (r *DescribeFileTamperEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileTamperEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGeneralStatRequestParams struct {
	// 云主机类型。
	// <li>CVM：表示腾讯云服务器</li>
	// <li>BM:  表示黑石物理机</li>
	// <li>ECM:  表示边缘计算服务器</li>
	// <li>LH:  表示轻量应用服务器</li>
	// <li>Other:  表示混合云机器</li>
	MachineType *string `json:"MachineType,omitnil" name:"MachineType"`

	// 机器所属地域。如：ap-guangzhou，ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitnil" name:"MachineRegion"`
}

type DescribeGeneralStatRequest struct {
	*tchttp.BaseRequest
	
	// 云主机类型。
	// <li>CVM：表示腾讯云服务器</li>
	// <li>BM:  表示黑石物理机</li>
	// <li>ECM:  表示边缘计算服务器</li>
	// <li>LH:  表示轻量应用服务器</li>
	// <li>Other:  表示混合云机器</li>
	MachineType *string `json:"MachineType,omitnil" name:"MachineType"`

	// 机器所属地域。如：ap-guangzhou，ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitnil" name:"MachineRegion"`
}

func (r *DescribeGeneralStatRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGeneralStatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MachineType")
	delete(f, "MachineRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGeneralStatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGeneralStatResponseParams struct {
	// 云主机总数
	MachinesAll *uint64 `json:"MachinesAll,omitnil" name:"MachinesAll"`

	// 云主机没有安装主机安全客户端的总数
	MachinesUninstalled *uint64 `json:"MachinesUninstalled,omitnil" name:"MachinesUninstalled"`

	// 主机安全客户端总数的总数
	AgentsAll *uint64 `json:"AgentsAll,omitnil" name:"AgentsAll"`

	// 主机安全客户端在线的总数
	AgentsOnline *uint64 `json:"AgentsOnline,omitnil" name:"AgentsOnline"`

	// 主机安全客户端 离线+关机 的总数
	AgentsOffline *uint64 `json:"AgentsOffline,omitnil" name:"AgentsOffline"`

	// 主机安全客户端专业版的总数
	AgentsPro *uint64 `json:"AgentsPro,omitnil" name:"AgentsPro"`

	// 主机安全客户端基础版的总数
	AgentsBasic *uint64 `json:"AgentsBasic,omitnil" name:"AgentsBasic"`

	// 7天内到期的预付费专业版总数
	AgentsProExpireWithInSevenDays *uint64 `json:"AgentsProExpireWithInSevenDays,omitnil" name:"AgentsProExpireWithInSevenDays"`

	// 风险主机总数
	RiskMachine *uint64 `json:"RiskMachine,omitnil" name:"RiskMachine"`

	// 已关机总数
	Shutdown *uint64 `json:"Shutdown,omitnil" name:"Shutdown"`

	// 已离线总数
	Offline *uint64 `json:"Offline,omitnil" name:"Offline"`

	// 旗舰版主机数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlagshipMachineCnt *uint64 `json:"FlagshipMachineCnt,omitnil" name:"FlagshipMachineCnt"`

	// 保护天数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtectDays *uint64 `json:"ProtectDays,omitnil" name:"ProtectDays"`

	// 15天内新增的主机数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddedOnTheFifteen *uint64 `json:"AddedOnTheFifteen,omitnil" name:"AddedOnTheFifteen"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeGeneralStatResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGeneralStatResponseParams `json:"Response"`
}

func (r *DescribeGeneralStatResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGeneralStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHistoryAccountsRequestParams struct {
	// 云镜客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>Username - String - 是否必填：否 - 帐号名</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeHistoryAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 云镜客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>Username - String - 是否必填：否 - 帐号名</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeHistoryAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHistoryAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHistoryAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHistoryAccountsResponseParams struct {
	// 帐号变更历史列表记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 帐号变更历史数据数组。
	HistoryAccounts []*HistoryAccount `json:"HistoryAccounts,omitnil" name:"HistoryAccounts"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeHistoryAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHistoryAccountsResponseParams `json:"Response"`
}

func (r *DescribeHistoryAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHistoryAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHistoryServiceRequestParams struct {

}

type DescribeHistoryServiceRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeHistoryServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHistoryServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHistoryServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHistoryServiceResponseParams struct {
	// 1 可购买 2 只能升降配 3 只能跳到续费管理页
	BuyStatus *uint64 `json:"BuyStatus,omitnil" name:"BuyStatus"`

	// 用户已购容量 单位 G
	InquireNum *uint64 `json:"InquireNum,omitnil" name:"InquireNum"`

	// 到期时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 是否自动续费,0 初始值, 1 开通 2 没开通
	IsAutoOpenRenew *uint64 `json:"IsAutoOpenRenew,omitnil" name:"IsAutoOpenRenew"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 0 没开通 1 正常 2隔离 3销毁
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeHistoryServiceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHistoryServiceResponseParams `json:"Response"`
}

func (r *DescribeHistoryServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHistoryServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostLoginListRequestParams struct {
	// 需要返回的数量，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Uuid - String - 是否必填：否 - 主机安全唯一Uuid</li>
	// <li>Quuid - String - 是否必填：否 - 云服务器uuid</li>
	// <li>UserName - String - 是否必填：否 - 用户名筛选</li>
	// <li>LoginTimeBegin - String - 是否必填：否 - 按照修改时间段筛选，开始时间</li>
	// <li>LoginTimeEnd - String - 是否必填：否 - 按照修改时间段筛选，结束时间</li>
	// <li>SrcIp - String - 是否必填：否 - 来源ip筛选</li>
	// <li>Status - int - 是否必填：否 - 状态筛选1:正常登录；5：已加白,14:已处理，15：已忽略</li>
	// <li>RiskLevel - int - 是否必填：否 - 状态筛选0:高危；1：可疑</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 排序方式：根据请求次数排序：asc-升序/desc-降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段：LoginTime-发生时间
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeHostLoginListRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Uuid - String - 是否必填：否 - 主机安全唯一Uuid</li>
	// <li>Quuid - String - 是否必填：否 - 云服务器uuid</li>
	// <li>UserName - String - 是否必填：否 - 用户名筛选</li>
	// <li>LoginTimeBegin - String - 是否必填：否 - 按照修改时间段筛选，开始时间</li>
	// <li>LoginTimeEnd - String - 是否必填：否 - 按照修改时间段筛选，结束时间</li>
	// <li>SrcIp - String - 是否必填：否 - 来源ip筛选</li>
	// <li>Status - int - 是否必填：否 - 状态筛选1:正常登录；5：已加白,14:已处理，15：已忽略</li>
	// <li>RiskLevel - int - 是否必填：否 - 状态筛选0:高危；1：可疑</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 排序方式：根据请求次数排序：asc-升序/desc-降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段：LoginTime-发生时间
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeHostLoginListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostLoginListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostLoginListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostLoginListResponseParams struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 登录审计列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostLoginList []*HostLoginList `json:"HostLoginList,omitnil" name:"HostLoginList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeHostLoginListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostLoginListResponseParams `json:"Response"`
}

func (r *DescribeHostLoginListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostLoginListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIgnoreBaselineRuleRequestParams struct {
	// 分页参数 最大100条
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页参数
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 检测项名称
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`
}

type DescribeIgnoreBaselineRuleRequest struct {
	*tchttp.BaseRequest
	
	// 分页参数 最大100条
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页参数
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 检测项名称
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`
}

func (r *DescribeIgnoreBaselineRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIgnoreBaselineRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "RuleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIgnoreBaselineRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIgnoreBaselineRuleResponseParams struct {
	// 忽略基线检测项列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoreBaselineRuleList []*IgnoreBaselineRule `json:"IgnoreBaselineRuleList,omitnil" name:"IgnoreBaselineRuleList"`

	// 分页查询记录总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeIgnoreBaselineRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIgnoreBaselineRuleResponseParams `json:"Response"`
}

func (r *DescribeIgnoreBaselineRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIgnoreBaselineRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIgnoreHostAndItemConfigRequestParams struct {
	// <li>ItemId - int64 - 是否必填：否 - 项Id</i>
	// <li>RuleId - int64 - 是否必填：否 - 规则Id</li>
	// <li>HostId - string - 是否必填：否 - 主机Id</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeIgnoreHostAndItemConfigRequest struct {
	*tchttp.BaseRequest
	
	// <li>ItemId - int64 - 是否必填：否 - 项Id</i>
	// <li>RuleId - int64 - 是否必填：否 - 规则Id</li>
	// <li>HostId - string - 是否必填：否 - 主机Id</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeIgnoreHostAndItemConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIgnoreHostAndItemConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIgnoreHostAndItemConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIgnoreHostAndItemConfigResponseParams struct {
	// 受影响检测项
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemSet []*BaselineItemInfo `json:"ItemSet,omitnil" name:"ItemSet"`

	// 受影响主机
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostSet []*BaselineHost `json:"HostSet,omitnil" name:"HostSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeIgnoreHostAndItemConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIgnoreHostAndItemConfigResponseParams `json:"Response"`
}

func (r *DescribeIgnoreHostAndItemConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIgnoreHostAndItemConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIgnoreRuleEffectHostListRequestParams struct {
	// 分页参数 最大100条
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页参数
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 检测项id
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`

	// 过滤条件。
	// <li>AliasName- String- 主机别名</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 主机标签名
	TagNames []*string `json:"TagNames,omitnil" name:"TagNames"`
}

type DescribeIgnoreRuleEffectHostListRequest struct {
	*tchttp.BaseRequest
	
	// 分页参数 最大100条
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页参数
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 检测项id
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`

	// 过滤条件。
	// <li>AliasName- String- 主机别名</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 主机标签名
	TagNames []*string `json:"TagNames,omitnil" name:"TagNames"`
}

func (r *DescribeIgnoreRuleEffectHostListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIgnoreRuleEffectHostListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "RuleId")
	delete(f, "Filters")
	delete(f, "TagNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIgnoreRuleEffectHostListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIgnoreRuleEffectHostListResponseParams struct {
	// 忽略检测项影响主机列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoreRuleEffectHostList []*IgnoreRuleEffectHostInfo `json:"IgnoreRuleEffectHostList,omitnil" name:"IgnoreRuleEffectHostList"`

	// 分页查询记录总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeIgnoreRuleEffectHostListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIgnoreRuleEffectHostListResponseParams `json:"Response"`
}

func (r *DescribeIgnoreRuleEffectHostListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIgnoreRuleEffectHostListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImportMachineInfoRequestParams struct {
	// 服务器内网IP（默认）/ 服务器名称 / 服务器ID 数组 (最大 1000条)
	MachineList []*string `json:"MachineList,omitnil" name:"MachineList"`

	// 批量导入的数据类型：Ip、Name、Id 三选一
	ImportType *string `json:"ImportType,omitnil" name:"ImportType"`

	// 该参数已作废.
	IsQueryProMachine *bool `json:"IsQueryProMachine,omitnil" name:"IsQueryProMachine"`

	// 过滤条件：
	// <li>Version - String  是否必填：否 - 当前防护版本（ PRO_VERSION：专业版 | BASIC_VERSION：基础版 | Flagship：旗舰版 | ProtectedMachines：专业版+旗舰版） | BASIC_PROPOST_GENERAL_DISCOUNT：普惠版+专业版按量计费+基础版主机 | UnFlagship：专业版预付费+专业版后付费+基础版+普惠版</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`
}

type DescribeImportMachineInfoRequest struct {
	*tchttp.BaseRequest
	
	// 服务器内网IP（默认）/ 服务器名称 / 服务器ID 数组 (最大 1000条)
	MachineList []*string `json:"MachineList,omitnil" name:"MachineList"`

	// 批量导入的数据类型：Ip、Name、Id 三选一
	ImportType *string `json:"ImportType,omitnil" name:"ImportType"`

	// 该参数已作废.
	IsQueryProMachine *bool `json:"IsQueryProMachine,omitnil" name:"IsQueryProMachine"`

	// 过滤条件：
	// <li>Version - String  是否必填：否 - 当前防护版本（ PRO_VERSION：专业版 | BASIC_VERSION：基础版 | Flagship：旗舰版 | ProtectedMachines：专业版+旗舰版） | BASIC_PROPOST_GENERAL_DISCOUNT：普惠版+专业版按量计费+基础版主机 | UnFlagship：专业版预付费+专业版后付费+基础版+普惠版</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeImportMachineInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImportMachineInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MachineList")
	delete(f, "ImportType")
	delete(f, "IsQueryProMachine")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImportMachineInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImportMachineInfoResponseParams struct {
	// 有效的机器信息列表：机器名称、机器公网/内网ip、机器标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	EffectiveMachineInfoList []*EffectiveMachineInfo `json:"EffectiveMachineInfoList,omitnil" name:"EffectiveMachineInfoList"`

	// 用户批量导入失败的机器列表（例如机器不存在等...）
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvalidMachineList []*string `json:"InvalidMachineList,omitnil" name:"InvalidMachineList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeImportMachineInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImportMachineInfoResponseParams `json:"Response"`
}

func (r *DescribeImportMachineInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImportMachineInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIndexListRequestParams struct {

}

type DescribeIndexListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeIndexListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIndexListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIndexListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIndexListResponseParams struct {
	// ES 索引信息
	Data *string `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeIndexListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIndexListResponseParams `json:"Response"`
}

func (r *DescribeIndexListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIndexListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJavaMemShellListRequestParams struct {
	// 过滤条件：Keywords: ip或者主机名模糊查询, Type，Status精确匹配，CreateBeginTime，CreateEndTime时间段
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeJavaMemShellListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件：Keywords: ip或者主机名模糊查询, Type，Status精确匹配，CreateBeginTime，CreateEndTime时间段
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeJavaMemShellListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJavaMemShellListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeJavaMemShellListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJavaMemShellListResponseParams struct {
	// 事件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*JavaMemShellInfo `json:"List,omitnil" name:"List"`

	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeJavaMemShellListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeJavaMemShellListResponseParams `json:"Response"`
}

func (r *DescribeJavaMemShellListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJavaMemShellListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLicenseBindListRequestParams struct {
	// 授权ID
	LicenseId *uint64 `json:"LicenseId,omitnil" name:"LicenseId"`

	// 授权类型
	LicenseType *uint64 `json:"LicenseType,omitnil" name:"LicenseType"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// <li>InstanceID、IP、
	// 
	// MachineName 模糊查询</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 限制条数,默认10.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeLicenseBindListRequest struct {
	*tchttp.BaseRequest
	
	// 授权ID
	LicenseId *uint64 `json:"LicenseId,omitnil" name:"LicenseId"`

	// 授权类型
	LicenseType *uint64 `json:"LicenseType,omitnil" name:"LicenseType"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// <li>InstanceID、IP、
	// 
	// MachineName 模糊查询</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 限制条数,默认10.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeLicenseBindListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLicenseBindListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LicenseId")
	delete(f, "LicenseType")
	delete(f, "ResourceId")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLicenseBindListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLicenseBindListResponseParams struct {
	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 绑定机器列表信息
	List []*LicenseBindDetail `json:"List,omitnil" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeLicenseBindListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLicenseBindListResponseParams `json:"Response"`
}

func (r *DescribeLicenseBindListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLicenseBindListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLicenseBindScheduleRequestParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// 限制条数,默认10.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤参数
	// Status 绑定进度状态 0 进行中 1 已完成 2 失败
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeLicenseBindScheduleRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// 限制条数,默认10.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤参数
	// Status 绑定进度状态 0 进行中 1 已完成 2 失败
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeLicenseBindScheduleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLicenseBindScheduleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLicenseBindScheduleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLicenseBindScheduleResponseParams struct {
	// 进度
	Schedule *uint64 `json:"Schedule,omitnil" name:"Schedule"`

	// 绑定任务详情
	List []*LicenseBindTaskDetail `json:"List,omitnil" name:"List"`

	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeLicenseBindScheduleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLicenseBindScheduleResponseParams `json:"Response"`
}

func (r *DescribeLicenseBindScheduleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLicenseBindScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLicenseGeneralRequestParams struct {

}

type DescribeLicenseGeneralRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeLicenseGeneralRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLicenseGeneralRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLicenseGeneralRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLicenseGeneralResponseParams struct {
	// 总授权数 (包含隔离,过期等不可用状态)
	LicenseCnt *uint64 `json:"LicenseCnt,omitnil" name:"LicenseCnt"`

	// 可用授权数
	AvailableLicenseCnt *uint64 `json:"AvailableLicenseCnt,omitnil" name:"AvailableLicenseCnt"`

	// 可用专业版授权数(包含后付费).
	AvailableProVersionLicenseCnt *uint64 `json:"AvailableProVersionLicenseCnt,omitnil" name:"AvailableProVersionLicenseCnt"`

	// 可用旗舰版授权数
	AvailableFlagshipVersionLicenseCnt *uint64 `json:"AvailableFlagshipVersionLicenseCnt,omitnil" name:"AvailableFlagshipVersionLicenseCnt"`

	// 即将到期授权数 (15天内到期的)
	NearExpiryLicenseCnt *uint64 `json:"NearExpiryLicenseCnt,omitnil" name:"NearExpiryLicenseCnt"`

	// 已到期授权数(不包含已删除的记录)
	ExpireLicenseCnt *uint64 `json:"ExpireLicenseCnt,omitnil" name:"ExpireLicenseCnt"`

	// 自动升级开关状态,默认 false,  true 开启, false 关闭
	AutoOpenStatus *bool `json:"AutoOpenStatus,omitnil" name:"AutoOpenStatus"`

	// PROVERSION_POSTPAY 专业版-后付费, PROVERSION_PREPAY 专业版-预付费, FLAGSHIP_PREPAY 旗舰版-预付费
	ProtectType *string `json:"ProtectType,omitnil" name:"ProtectType"`

	// 历史是否开通过自动升级开关
	IsOpenStatusHistory *bool `json:"IsOpenStatusHistory,omitnil" name:"IsOpenStatusHistory"`

	// 已使用授权数
	UsedLicenseCnt *uint64 `json:"UsedLicenseCnt,omitnil" name:"UsedLicenseCnt"`

	// 未到期授权数
	NotExpiredLicenseCnt *uint64 `json:"NotExpiredLicenseCnt,omitnil" name:"NotExpiredLicenseCnt"`

	// 旗舰版总授权数(有效订单)
	FlagshipVersionLicenseCnt *uint64 `json:"FlagshipVersionLicenseCnt,omitnil" name:"FlagshipVersionLicenseCnt"`

	// 专业版总授权数(有效订单)
	ProVersionLicenseCnt *uint64 `json:"ProVersionLicenseCnt,omitnil" name:"ProVersionLicenseCnt"`

	// 普惠版总授权数(有效订单的授权数)
	CwpVersionLicenseCnt *uint64 `json:"CwpVersionLicenseCnt,omitnil" name:"CwpVersionLicenseCnt"`

	// 可用惠普版授权数
	AvailableLHLicenseCnt *uint64 `json:"AvailableLHLicenseCnt,omitnil" name:"AvailableLHLicenseCnt"`

	// 自动加购开关, true 开启, false 关闭
	AutoRepurchaseSwitch *bool `json:"AutoRepurchaseSwitch,omitnil" name:"AutoRepurchaseSwitch"`

	// 自动加购订单是否自动续费 ,true 开启, false 关闭
	AutoRepurchaseRenewSwitch *bool `json:"AutoRepurchaseRenewSwitch,omitnil" name:"AutoRepurchaseRenewSwitch"`

	// 已销毁订单数
	DestroyOrderNum *uint64 `json:"DestroyOrderNum,omitnil" name:"DestroyOrderNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeLicenseGeneralResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLicenseGeneralResponseParams `json:"Response"`
}

func (r *DescribeLicenseGeneralResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLicenseGeneralResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLicenseListRequestParams struct {
	// 多个条件筛选时 LicenseStatus,DeadlineStatus,ResourceId,Keywords 取交集
	// <li> LicenseStatus 授权状态信息,0 未使用,1 部分使用, 2 已用完, 3 不可用  4 可使用</li>
	// <li> BuyTime 购买时间</li>
	// <li> LicenseType  授权类型, 0 专业版-按量计费, 1专业版-包年包月 , 2 旗舰版-包年包月</li>
	// <li>DeadlineStatus 到期状态 NotExpired 未过期, Expire 已过期(包含已销毁) NearExpiry 即将到期</li>
	// <li>ResourceId 资源ID</li>
	// <li>Keywords IP筛选</li>
	// <li>PayMode 付费模式 0 按量计费 , 1 包年包月</li>
	// <li>OrderStatus 订单状态 1 正常 2 隔离 3 销毁</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 限制条数,默认10.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 标签筛选,平台标签能力,这里传入 标签键,标签值作为一个对象
	Tags []*Tags `json:"Tags,omitnil" name:"Tags"`
}

type DescribeLicenseListRequest struct {
	*tchttp.BaseRequest
	
	// 多个条件筛选时 LicenseStatus,DeadlineStatus,ResourceId,Keywords 取交集
	// <li> LicenseStatus 授权状态信息,0 未使用,1 部分使用, 2 已用完, 3 不可用  4 可使用</li>
	// <li> BuyTime 购买时间</li>
	// <li> LicenseType  授权类型, 0 专业版-按量计费, 1专业版-包年包月 , 2 旗舰版-包年包月</li>
	// <li>DeadlineStatus 到期状态 NotExpired 未过期, Expire 已过期(包含已销毁) NearExpiry 即将到期</li>
	// <li>ResourceId 资源ID</li>
	// <li>Keywords IP筛选</li>
	// <li>PayMode 付费模式 0 按量计费 , 1 包年包月</li>
	// <li>OrderStatus 订单状态 1 正常 2 隔离 3 销毁</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 限制条数,默认10.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 标签筛选,平台标签能力,这里传入 标签键,标签值作为一个对象
	Tags []*Tags `json:"Tags,omitnil" name:"Tags"`
}

func (r *DescribeLicenseListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLicenseListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLicenseListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLicenseListResponseParams struct {
	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 授权数列表信息
	List []*LicenseDetail `json:"List,omitnil" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeLicenseListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLicenseListResponseParams `json:"Response"`
}

func (r *DescribeLicenseListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLicenseListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogStorageConfigRequestParams struct {

}

type DescribeLogStorageConfigRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeLogStorageConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogStorageConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogStorageConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogStorageConfigResponseParams struct {
	// 存储类型，string数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type []*string `json:"Type,omitnil" name:"Type"`

	// 日志存储天数，3640表示不限
	// 注意：此字段可能返回 null，表示取不到有效值。
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// 本月Period的修改次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeriodModifyCount *int64 `json:"PeriodModifyCount,omitnil" name:"PeriodModifyCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeLogStorageConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogStorageConfigResponseParams `json:"Response"`
}

func (r *DescribeLogStorageConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogStorageConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogStorageRecordRequestParams struct {

}

type DescribeLogStorageRecordRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeLogStorageRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogStorageRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogStorageRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogStorageRecordResponseParams struct {
	// 存储量记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	Records []*LogStorageRecord `json:"Records,omitnil" name:"Records"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeLogStorageRecordResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogStorageRecordResponseParams `json:"Response"`
}

func (r *DescribeLogStorageRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogStorageRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogStorageStatisticRequestParams struct {

}

type DescribeLogStorageStatisticRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeLogStorageStatisticRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogStorageStatisticRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogStorageStatisticRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogStorageStatisticResponseParams struct {
	// 总容量（单位：GB）
	TotalSize *uint64 `json:"TotalSize,omitnil" name:"TotalSize"`

	// 已使用容量（单位：GB）
	UsedSize *uint64 `json:"UsedSize,omitnil" name:"UsedSize"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeLogStorageStatisticResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogStorageStatisticResponseParams `json:"Response"`
}

func (r *DescribeLogStorageStatisticResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogStorageStatisticResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoginWhiteCombinedListRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>UserName - String - 是否必填：否 - 用户名筛选</li>
	// <li>ModifyBeginTime - String - 是否必填：否 - 按照修改时间段筛选，开始时间</li>
	// <li>ModifyEndTime - String - 是否必填：否 - 按照修改时间段筛选，结束时间</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeLoginWhiteCombinedListRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>UserName - String - 是否必填：否 - 用户名筛选</li>
	// <li>ModifyBeginTime - String - 是否必填：否 - 按照修改时间段筛选，开始时间</li>
	// <li>ModifyEndTime - String - 是否必填：否 - 按照修改时间段筛选，结束时间</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeLoginWhiteCombinedListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoginWhiteCombinedListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoginWhiteCombinedListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoginWhiteCombinedListResponseParams struct {
	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 合并后的白名单列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoginWhiteCombinedInfos []*LoginWhiteCombinedInfo `json:"LoginWhiteCombinedInfos,omitnil" name:"LoginWhiteCombinedInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeLoginWhiteCombinedListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoginWhiteCombinedListResponseParams `json:"Response"`
}

func (r *DescribeLoginWhiteCombinedListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoginWhiteCombinedListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoginWhiteListRequestParams struct {
	// 返回数量，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 查询关键字 </li>
	// <li>UserName - String - 是否必填：否 - 用户名筛选 </li>
	// <li>ModifyBeginTime - String - 是否必填：否 - 按照修改时间段筛选，开始时间 </li>
	// <li>ModifyEndTime - String - 是否必填：否 - 按照修改时间段筛选，结束时间 </li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeLoginWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 查询关键字 </li>
	// <li>UserName - String - 是否必填：否 - 用户名筛选 </li>
	// <li>ModifyBeginTime - String - 是否必填：否 - 按照修改时间段筛选，开始时间 </li>
	// <li>ModifyEndTime - String - 是否必填：否 - 按照修改时间段筛选，结束时间 </li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeLoginWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoginWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoginWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoginWhiteListResponseParams struct {
	// 记录总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 异地登录白名单数组
	LoginWhiteLists []*LoginWhiteLists `json:"LoginWhiteLists,omitnil" name:"LoginWhiteLists"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeLoginWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoginWhiteListResponseParams `json:"Response"`
}

func (r *DescribeLoginWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoginWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachineInfoRequestParams struct {
	// 主机安全客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Quuid , Uuid 必填一项
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`
}

type DescribeMachineInfoRequest struct {
	*tchttp.BaseRequest
	
	// 主机安全客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Quuid , Uuid 必填一项
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`
}

func (r *DescribeMachineInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "Quuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachineInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachineInfoResponseParams struct {
	// 机器ip。
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// 受主机安全保护天数。
	ProtectDays *uint64 `json:"ProtectDays,omitnil" name:"ProtectDays"`

	// 操作系统。
	MachineOs *string `json:"MachineOs,omitnil" name:"MachineOs"`

	// 主机名称。
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 在线状态。
	// <li>ONLINE： 在线</li>
	// <li>OFFLINE：离线</li>
	MachineStatus *string `json:"MachineStatus,omitnil" name:"MachineStatus"`

	// CVM或BM主机唯一标识。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 主机外网IP。
	MachineWanIp *string `json:"MachineWanIp,omitnil" name:"MachineWanIp"`

	// CVM或BM主机唯一Uuid。
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机安全客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 是否开通专业版。
	// <li>true：是</li>
	// <li>false：否</li>
	IsProVersion *bool `json:"IsProVersion,omitnil" name:"IsProVersion"`

	// 专业版开通时间。
	ProVersionOpenDate *string `json:"ProVersionOpenDate,omitnil" name:"ProVersionOpenDate"`

	// 云服务器类型。
	// <li>CVM: 腾讯云服务器</li>
	// <li>BM: 黑石物理机</li>
	// <li>ECM: 边缘计算服务器</li>
	// <li>LH: 轻量应用服务器</li>
	// <li>Other: 混合云机器</li>
	MachineType *string `json:"MachineType,omitnil" name:"MachineType"`

	// 机器所属地域。如：ap-guangzhou，ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitnil" name:"MachineRegion"`

	// 主机状态。
	// <li>POSTPAY: 表示后付费，即按量计费  </li>
	// <li>PREPAY: 表示预付费，即包年包月</li>
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// 免费木马剩余检测数量。
	FreeMalwaresLeft *uint64 `json:"FreeMalwaresLeft,omitnil" name:"FreeMalwaresLeft"`

	// 免费漏洞剩余检测数量。
	FreeVulsLeft *uint64 `json:"FreeVulsLeft,omitnil" name:"FreeVulsLeft"`

	// agent版本号
	AgentVersion *string `json:"AgentVersion,omitnil" name:"AgentVersion"`

	// 专业版到期时间(仅预付费)
	ProVersionDeadline *string `json:"ProVersionDeadline,omitnil" name:"ProVersionDeadline"`

	// 是否有资产扫描记录，0无，1有
	HasAssetScan *uint64 `json:"HasAssetScan,omitnil" name:"HasAssetScan"`

	// 防护版本：BASIC_VERSION 基础版，PRO_VERSION 专业版，Flagship 旗舰版，GENERAL_DISCOUNT 普惠版
	ProtectType *string `json:"ProtectType,omitnil" name:"ProtectType"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMachineInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMachineInfoResponseParams `json:"Response"`
}

func (r *DescribeMachineInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachineListRequestParams struct {
	// 云主机类型。
	// <li>CVM：表示虚拟主机</li>
	// <li>BM:  表示黑石物理机</li>
	// <li>ECM:  表示边缘计算服务器</li>
	// <li>LH:  表示轻量应用服务器</li>
	// <li>Other:  表示混合云机器</li>
	MachineType *string `json:"MachineType,omitnil" name:"MachineType"`

	// 机器所属地域。如：ap-guangzhou，ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitnil" name:"MachineRegion"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字 </li>
	// <li>Status - String - 是否必填：否 - 客户端在线状态（OFFLINE: 离线 | ONLINE: 在线 | UNINSTALLED：未安装）</li>
	// <li>Version - String  是否必填：否 - 当前防护版本（ PRO_VERSION：专业版 | BASIC_VERSION：基础版）</li>
	// 每个过滤条件只支持一个值，暂不支持多个值“或”关系查询
	Filters []*AssetFilters `json:"Filters,omitnil" name:"Filters"`
}

type DescribeMachineListRequest struct {
	*tchttp.BaseRequest
	
	// 云主机类型。
	// <li>CVM：表示虚拟主机</li>
	// <li>BM:  表示黑石物理机</li>
	// <li>ECM:  表示边缘计算服务器</li>
	// <li>LH:  表示轻量应用服务器</li>
	// <li>Other:  表示混合云机器</li>
	MachineType *string `json:"MachineType,omitnil" name:"MachineType"`

	// 机器所属地域。如：ap-guangzhou，ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitnil" name:"MachineRegion"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字 </li>
	// <li>Status - String - 是否必填：否 - 客户端在线状态（OFFLINE: 离线 | ONLINE: 在线 | UNINSTALLED：未安装）</li>
	// <li>Version - String  是否必填：否 - 当前防护版本（ PRO_VERSION：专业版 | BASIC_VERSION：基础版）</li>
	// 每个过滤条件只支持一个值，暂不支持多个值“或”关系查询
	Filters []*AssetFilters `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeMachineListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MachineType")
	delete(f, "MachineRegion")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachineListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachineListResponseParams struct {
	// 主机列表
	Machines []*Machine `json:"Machines,omitnil" name:"Machines"`

	// 主机数量
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMachineListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMachineListResponseParams `json:"Response"`
}

func (r *DescribeMachineListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachineOsListRequestParams struct {

}

type DescribeMachineOsListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeMachineOsListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineOsListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachineOsListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachineOsListResponseParams struct {
	// 操作系统列表
	List []*OsName `json:"List,omitnil" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMachineOsListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMachineOsListResponseParams `json:"Response"`
}

func (r *DescribeMachineOsListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineOsListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachineRegionsRequestParams struct {

}

type DescribeMachineRegionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeMachineRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachineRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachineRegionsResponseParams struct {
	// CVM 云服务器地域列表
	CVM []*RegionInfo `json:"CVM,omitnil" name:"CVM"`

	// BM 黑石机器地域列表
	BM []*RegionInfo `json:"BM,omitnil" name:"BM"`

	// LH 轻量应用服务器地域列表
	LH []*RegionInfo `json:"LH,omitnil" name:"LH"`

	// ECM 边缘计算服务器地域列表
	ECM []*RegionInfo `json:"ECM,omitnil" name:"ECM"`

	// Other 混合云地域列表
	Other []*RegionInfo `json:"Other,omitnil" name:"Other"`

	// 所有地域列表(包含以上所有地域)
	ALL []*RegionInfo `json:"ALL,omitnil" name:"ALL"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMachineRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMachineRegionsResponseParams `json:"Response"`
}

func (r *DescribeMachineRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachineRiskCntRequestParams struct {
	// 过滤条件。
	// <li>Uuids- String - 是否必填：否 - 主机uuid</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeMachineRiskCntRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>Uuids- String - 是否必填：否 - 主机uuid</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeMachineRiskCntRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineRiskCntRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachineRiskCntRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachineRiskCntResponseParams struct {
	// 异地登录
	HostLogin *uint64 `json:"HostLogin,omitnil" name:"HostLogin"`

	// 密码破解
	BruteAttack *uint64 `json:"BruteAttack,omitnil" name:"BruteAttack"`

	// 恶意请求
	MaliciousRequest *uint64 `json:"MaliciousRequest,omitnil" name:"MaliciousRequest"`

	// 反弹shell
	ReverseShell *uint64 `json:"ReverseShell,omitnil" name:"ReverseShell"`

	// 高危命令
	Bash *uint64 `json:"Bash,omitnil" name:"Bash"`

	// 本地提权
	PrivilegeEscalation *uint64 `json:"PrivilegeEscalation,omitnil" name:"PrivilegeEscalation"`

	// 木马
	Malware *uint64 `json:"Malware,omitnil" name:"Malware"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMachineRiskCntResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMachineRiskCntResponseParams `json:"Response"`
}

func (r *DescribeMachineRiskCntResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineRiskCntResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachinesRequestParams struct {
	// 机器所属专区类型 
	// CVM 云服务器
	// BM 黑石
	// ECM 边缘计算
	// LH 轻量应用服务器
	// Other 混合云专区
	MachineType *string `json:"MachineType,omitnil" name:"MachineType"`

	// 机器所属地域。如：ap-guangzhou，ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitnil" name:"MachineRegion"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>Ips - String - 是否必填：否 - 通过ip查询 </li>
	// <li>Names - String - 是否必填：否 - 通过实例名查询 </li>
	// <li>InstanceIds - String - 是否必填：否 - 通过实例id查询 </li>
	// <li>Status - String - 是否必填：否 - 客户端在线状态（OFFLINE: 离线/关机 | ONLINE: 在线 | UNINSTALLED：未安装 | AGENT_OFFLINE 离线| AGENT_SHUTDOWN 已关机）</li>
	// <li>Version - String  是否必填：否 - 当前防护版本（ PRO_VERSION：专业版 | BASIC_VERSION：基础版 | Flagship : 旗舰版 | ProtectedMachines: 专业版+旗舰版）</li>
	// <li>Risk - String 是否必填: 否 - 风险主机( yes ) </li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )
	// 每个过滤条件只支持一个值，暂不支持多个值“或”关系查询
	// <li>Quuid - String - 是否必填: 否 - 云服务器uuid  最大100条.</li>
	// <li>AddedOnTheFifteen- String 是否必填: 否 - 是否只查询15天内新增的主机( 1：是) </li>
	// <li> TagId- String 是否必填: 否 - 查询指定标签关联的主机列表 </li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 机器所属业务ID列表
	ProjectIds []*uint64 `json:"ProjectIds,omitnil" name:"ProjectIds"`
}

type DescribeMachinesRequest struct {
	*tchttp.BaseRequest
	
	// 机器所属专区类型 
	// CVM 云服务器
	// BM 黑石
	// ECM 边缘计算
	// LH 轻量应用服务器
	// Other 混合云专区
	MachineType *string `json:"MachineType,omitnil" name:"MachineType"`

	// 机器所属地域。如：ap-guangzhou，ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitnil" name:"MachineRegion"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>Ips - String - 是否必填：否 - 通过ip查询 </li>
	// <li>Names - String - 是否必填：否 - 通过实例名查询 </li>
	// <li>InstanceIds - String - 是否必填：否 - 通过实例id查询 </li>
	// <li>Status - String - 是否必填：否 - 客户端在线状态（OFFLINE: 离线/关机 | ONLINE: 在线 | UNINSTALLED：未安装 | AGENT_OFFLINE 离线| AGENT_SHUTDOWN 已关机）</li>
	// <li>Version - String  是否必填：否 - 当前防护版本（ PRO_VERSION：专业版 | BASIC_VERSION：基础版 | Flagship : 旗舰版 | ProtectedMachines: 专业版+旗舰版）</li>
	// <li>Risk - String 是否必填: 否 - 风险主机( yes ) </li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )
	// 每个过滤条件只支持一个值，暂不支持多个值“或”关系查询
	// <li>Quuid - String - 是否必填: 否 - 云服务器uuid  最大100条.</li>
	// <li>AddedOnTheFifteen- String 是否必填: 否 - 是否只查询15天内新增的主机( 1：是) </li>
	// <li> TagId- String 是否必填: 否 - 查询指定标签关联的主机列表 </li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 机器所属业务ID列表
	ProjectIds []*uint64 `json:"ProjectIds,omitnil" name:"ProjectIds"`
}

func (r *DescribeMachinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachinesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MachineType")
	delete(f, "MachineRegion")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "ProjectIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachinesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachinesResponseParams struct {
	// 主机列表
	Machines []*Machine `json:"Machines,omitnil" name:"Machines"`

	// 主机数量
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMachinesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMachinesResponseParams `json:"Response"`
}

func (r *DescribeMachinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMalWareListRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>FilePath - String - 是否必填：否 - 路径筛选</li>
	// <li>VirusName - String - 是否必填：否 - 描述筛选</li>
	// <li>CreateBeginTime - String - 是否必填：否 - 创建时间筛选-开始时间</li>
	// <li>CreateEndTime - String - 是否必填：否 - 创建时间筛选-结束时间</li>
	// <li>Status - String - 是否必填：否 - 状态筛选 4待处理,5信任,6已隔离,10隔离中,11恢复隔离中</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 检测排序 CreateTime
	By *string `json:"By,omitnil" name:"By"`

	// 排序方式 ASC,DESC
	Order *string `json:"Order,omitnil" name:"Order"`
}

type DescribeMalWareListRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>FilePath - String - 是否必填：否 - 路径筛选</li>
	// <li>VirusName - String - 是否必填：否 - 描述筛选</li>
	// <li>CreateBeginTime - String - 是否必填：否 - 创建时间筛选-开始时间</li>
	// <li>CreateEndTime - String - 是否必填：否 - 创建时间筛选-结束时间</li>
	// <li>Status - String - 是否必填：否 - 状态筛选 4待处理,5信任,6已隔离,10隔离中,11恢复隔离中</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 检测排序 CreateTime
	By *string `json:"By,omitnil" name:"By"`

	// 排序方式 ASC,DESC
	Order *string `json:"Order,omitnil" name:"Order"`
}

func (r *DescribeMalWareListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMalWareListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMalWareListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMalWareListResponseParams struct {
	// 木马列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MalWareList []*MalWareList `json:"MalWareList,omitnil" name:"MalWareList"`

	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMalWareListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMalWareListResponseParams `json:"Response"`
}

func (r *DescribeMalWareListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMalWareListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMaliciousRequestWhiteListRequestParams struct {
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// 
	// <li>Domain  - String - 基线名称</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 排序方式 [asc:升序|desc:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeMaliciousRequestWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// 
	// <li>Domain  - String - 基线名称</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 排序方式 [asc:升序|desc:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeMaliciousRequestWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaliciousRequestWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMaliciousRequestWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMaliciousRequestWhiteListResponseParams struct {
	// 白名单信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*MaliciousRequestWhiteListInfo `json:"List,omitnil" name:"List"`

	// 分页查询记录总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMaliciousRequestWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMaliciousRequestWhiteListResponseParams `json:"Response"`
}

func (r *DescribeMaliciousRequestWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaliciousRequestWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMalwareFileRequestParams struct {
	// 木马记录ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

type DescribeMalwareFileRequest struct {
	*tchttp.BaseRequest
	
	// 木马记录ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

func (r *DescribeMalwareFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMalwareFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMalwareFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMalwareFileResponseParams struct {
	// 木马文件下载地址
	FileUrl *string `json:"FileUrl,omitnil" name:"FileUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMalwareFileResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMalwareFileResponseParams `json:"Response"`
}

func (r *DescribeMalwareFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMalwareFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMalwareInfoRequestParams struct {
	// 唯一ID
	Id *int64 `json:"Id,omitnil" name:"Id"`
}

type DescribeMalwareInfoRequest struct {
	*tchttp.BaseRequest
	
	// 唯一ID
	Id *int64 `json:"Id,omitnil" name:"Id"`
}

func (r *DescribeMalwareInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMalwareInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMalwareInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMalwareInfoResponseParams struct {
	// 恶意文件详情信息
	MalwareInfo *MalwareInfo `json:"MalwareInfo,omitnil" name:"MalwareInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMalwareInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMalwareInfoResponseParams `json:"Response"`
}

func (r *DescribeMalwareInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMalwareInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMalwareRiskWarningRequestParams struct {

}

type DescribeMalwareRiskWarningRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeMalwareRiskWarningRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMalwareRiskWarningRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMalwareRiskWarningRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMalwareRiskWarningResponseParams struct {
	// 是否开启自动扫描：true-开启，false-未开启
	IsCheckRisk *bool `json:"IsCheckRisk,omitnil" name:"IsCheckRisk"`

	// 风险文件列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*MalwareRisk `json:"List,omitnil" name:"List"`

	// 是否弹出提示 true 弹出, false不弹
	IsPop *bool `json:"IsPop,omitnil" name:"IsPop"`

	// 异常进程列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessList []*MalwareRisk `json:"ProcessList,omitnil" name:"ProcessList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMalwareRiskWarningResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMalwareRiskWarningResponseParams `json:"Response"`
}

func (r *DescribeMalwareRiskWarningResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMalwareRiskWarningResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMalwareTimingScanSettingRequestParams struct {

}

type DescribeMalwareTimingScanSettingRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeMalwareTimingScanSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMalwareTimingScanSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMalwareTimingScanSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMalwareTimingScanSettingResponseParams struct {
	// 检测模式 0 全盘检测  1快速检测
	CheckPattern *uint64 `json:"CheckPattern,omitnil" name:"CheckPattern"`

	// 检测周期 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 检测周期 超时结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 是否全部服务器 1 全部 2 自选
	IsGlobal *uint64 `json:"IsGlobal,omitnil" name:"IsGlobal"`

	// 自选服务器时必须 主机quuid的string数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuuidList []*string `json:"QuuidList,omitnil" name:"QuuidList"`

	// 监控模式 0 标准 1深度
	MonitoringPattern *uint64 `json:"MonitoringPattern,omitnil" name:"MonitoringPattern"`

	// 周期 1每天
	Cycle *uint64 `json:"Cycle,omitnil" name:"Cycle"`

	// 定时检测开关 0 关闭1 开启
	EnableScan *int64 `json:"EnableScan,omitnil" name:"EnableScan"`

	// 唯一ID
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// 实时监控0 关闭 1开启
	RealTimeMonitoring *int64 `json:"RealTimeMonitoring,omitnil" name:"RealTimeMonitoring"`

	// 是否自动隔离：1-是，0-否
	AutoIsolation *uint64 `json:"AutoIsolation,omitnil" name:"AutoIsolation"`

	// 一键扫描超时时长，如：1800秒（s）
	ClickTimeout *uint64 `json:"ClickTimeout,omitnil" name:"ClickTimeout"`

	// 是否杀掉进程 1杀掉 0不杀掉 只有开启自动隔离才生效
	KillProcess *uint64 `json:"KillProcess,omitnil" name:"KillProcess"`

	// 1标准模式（只报严重、高危）、2增强模式（报严重、高危、中危）、3严格模式（报严重、高、中、低、提示）
	EngineType *uint64 `json:"EngineType,omitnil" name:"EngineType"`

	// 启发引擎 0 关闭 1开启
	EnableInspiredEngine *uint64 `json:"EnableInspiredEngine,omitnil" name:"EnableInspiredEngine"`

	// 是否开启恶意进程查杀[0:未开启,1:开启]
	EnableMemShellScan *uint64 `json:"EnableMemShellScan,omitnil" name:"EnableMemShellScan"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMalwareTimingScanSettingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMalwareTimingScanSettingResponseParams `json:"Response"`
}

func (r *DescribeMalwareTimingScanSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMalwareTimingScanSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMonthInspectionReportRequestParams struct {
	// 分页大小
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页步长
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeMonthInspectionReportRequest struct {
	*tchttp.BaseRequest
	
	// 分页大小
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页步长
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeMonthInspectionReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMonthInspectionReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMonthInspectionReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMonthInspectionReportResponseParams struct {
	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 巡检报告列表
	List []*MonthInspectionReport `json:"List,omitnil" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMonthInspectionReportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMonthInspectionReportResponseParams `json:"Response"`
}

func (r *DescribeMonthInspectionReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMonthInspectionReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOpenPortStatisticsRequestParams struct {
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>Port - Uint64 - 是否必填：否 - 端口号</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeOpenPortStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>Port - Uint64 - 是否必填：否 - 端口号</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeOpenPortStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpenPortStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOpenPortStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOpenPortStatisticsResponseParams struct {
	// 端口统计列表总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 端口统计数据列表
	OpenPortStatistics []*OpenPortStatistics `json:"OpenPortStatistics,omitnil" name:"OpenPortStatistics"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeOpenPortStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOpenPortStatisticsResponseParams `json:"Response"`
}

func (r *DescribeOpenPortStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpenPortStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewStatisticsRequestParams struct {

}

type DescribeOverviewStatisticsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeOverviewStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOverviewStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOverviewStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewStatisticsResponseParams struct {
	// 服务器在线数。
	OnlineMachineNum *uint64 `json:"OnlineMachineNum,omitnil" name:"OnlineMachineNum"`

	// 专业服务器数。
	ProVersionMachineNum *uint64 `json:"ProVersionMachineNum,omitnil" name:"ProVersionMachineNum"`

	// 木马文件数。
	MalwareNum *uint64 `json:"MalwareNum,omitnil" name:"MalwareNum"`

	// 异地登录数。
	NonlocalLoginNum *uint64 `json:"NonlocalLoginNum,omitnil" name:"NonlocalLoginNum"`

	// 暴力破解成功数。
	BruteAttackSuccessNum *uint64 `json:"BruteAttackSuccessNum,omitnil" name:"BruteAttackSuccessNum"`

	// 漏洞数。
	VulNum *uint64 `json:"VulNum,omitnil" name:"VulNum"`

	// 安全基线数。
	BaseLineNum *uint64 `json:"BaseLineNum,omitnil" name:"BaseLineNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeOverviewStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOverviewStatisticsResponseParams `json:"Response"`
}

func (r *DescribeOverviewStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOverviewStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrivilegeEventInfoRequestParams struct {
	// 事件id
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

type DescribePrivilegeEventInfoRequest struct {
	*tchttp.BaseRequest
	
	// 事件id
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

func (r *DescribePrivilegeEventInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivilegeEventInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrivilegeEventInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrivilegeEventInfoResponseParams struct {
	// 本地提权详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivilegeEventInfo *PrivilegeEventInfo `json:"PrivilegeEventInfo,omitnil" name:"PrivilegeEventInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribePrivilegeEventInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrivilegeEventInfoResponseParams `json:"Response"`
}

func (r *DescribePrivilegeEventInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivilegeEventInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrivilegeEventsRequestParams struct {
	// 返回数量，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键词(主机IP)</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 排序方式：根据请求次数排序：asc-升序/desc-降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段：CreateTime-发现时间
	By *string `json:"By,omitnil" name:"By"`
}

type DescribePrivilegeEventsRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键词(主机IP)</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 排序方式：根据请求次数排序：asc-升序/desc-降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段：CreateTime-发现时间
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribePrivilegeEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivilegeEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrivilegeEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrivilegeEventsResponseParams struct {
	// 数据列表
	List []*PrivilegeEscalationProcess `json:"List,omitnil" name:"List"`

	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribePrivilegeEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrivilegeEventsResponseParams `json:"Response"`
}

func (r *DescribePrivilegeEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivilegeEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrivilegeRulesRequestParams struct {
	// 返回数量，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键字(进程名称)</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribePrivilegeRulesRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键字(进程名称)</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribePrivilegeRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivilegeRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrivilegeRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrivilegeRulesResponseParams struct {
	// 列表内容
	List []*PrivilegeRule `json:"List,omitnil" name:"List"`

	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribePrivilegeRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrivilegeRulesResponseParams `json:"Response"`
}

func (r *DescribePrivilegeRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivilegeRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProVersionInfoRequestParams struct {

}

type DescribeProVersionInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeProVersionInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProVersionInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProVersionInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProVersionInfoResponseParams struct {
	// 后付费昨日扣费
	PostPayCost *uint64 `json:"PostPayCost,omitnil" name:"PostPayCost"`

	// 新增主机是否自动开通专业版
	IsAutoOpenProVersion *bool `json:"IsAutoOpenProVersion,omitnil" name:"IsAutoOpenProVersion"`

	// 开通专业版主机数
	ProVersionNum *uint64 `json:"ProVersionNum,omitnil" name:"ProVersionNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeProVersionInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProVersionInfoResponseParams `json:"Response"`
}

func (r *DescribeProVersionInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProVersionInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProVersionStatusRequestParams struct {
	// 主机安全客户端UUID、填写"all"表示所有主机。
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

type DescribeProVersionStatusRequest struct {
	*tchttp.BaseRequest
	
	// 主机安全客户端UUID、填写"all"表示所有主机。
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

func (r *DescribeProVersionStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProVersionStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProVersionStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProVersionStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeProVersionStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProVersionStatusResponseParams `json:"Response"`
}

func (r *DescribeProVersionStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProVersionStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProcessStatisticsRequestParams struct {
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>ProcessName - String - 是否必填：否 - 进程名</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeProcessStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>ProcessName - String - 是否必填：否 - 进程名</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeProcessStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProcessStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProcessStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProcessStatisticsResponseParams struct {
	// 进程统计列表记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 进程统计列表数据数组。
	ProcessStatistics []*ProcessStatistics `json:"ProcessStatistics,omitnil" name:"ProcessStatistics"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeProcessStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProcessStatisticsResponseParams `json:"Response"`
}

func (r *DescribeProcessStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProcessStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProtectDirListRequestParams struct {
	// 分页条数 最大100条
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// DirName 网站名称
	// DirPath 网站防护目录地址
	Filters []*AssetFilters `json:"Filters,omitnil" name:"Filters"`

	// asc：升序/desc：降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeProtectDirListRequest struct {
	*tchttp.BaseRequest
	
	// 分页条数 最大100条
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// DirName 网站名称
	// DirPath 网站防护目录地址
	Filters []*AssetFilters `json:"Filters,omitnil" name:"Filters"`

	// asc：升序/desc：降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeProtectDirListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProtectDirListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProtectDirListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProtectDirListResponseParams struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 防护目录列表信息
	List []*ProtectDirInfo `json:"List,omitnil" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeProtectDirListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProtectDirListResponseParams `json:"Response"`
}

func (r *DescribeProtectDirListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProtectDirListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProtectDirRelatedServerRequestParams struct {
	// 唯一ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// 分页条数 最大100条
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤参数 ProtectStatus
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序值
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeProtectDirRelatedServerRequest struct {
	*tchttp.BaseRequest
	
	// 唯一ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// 分页条数 最大100条
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤参数 ProtectStatus
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序值
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeProtectDirRelatedServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProtectDirRelatedServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProtectDirRelatedServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProtectDirRelatedServerResponseParams struct {
	// 网站关联服务器列表信息
	List []*ProtectDirRelatedServer `json:"List,omitnil" name:"List"`

	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 已开启防护总数
	ProtectServerCount *uint64 `json:"ProtectServerCount,omitnil" name:"ProtectServerCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeProtectDirRelatedServerResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProtectDirRelatedServerResponseParams `json:"Response"`
}

func (r *DescribeProtectDirRelatedServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProtectDirRelatedServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProtectNetListRequestParams struct {
	// 过滤条件。
	// <li>Keyword- String - 是否必填：否 - 关键词过滤，</li>
	// <li>Uuids - String - 是否必填：否 - 主机id过滤</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 需要返回的数量，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序步长
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方法
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段 StartTime，EndTime
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeProtectNetListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>Keyword- String - 是否必填：否 - 关键词过滤，</li>
	// <li>Uuids - String - 是否必填：否 - 主机id过滤</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 需要返回的数量，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序步长
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方法
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段 StartTime，EndTime
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeProtectNetListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProtectNetListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProtectNetListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProtectNetListResponseParams struct {
	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 安全管家数据
	List []*ProtectNetInfo `json:"List,omitnil" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeProtectNetListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProtectNetListResponseParams `json:"Response"`
}

func (r *DescribeProtectNetListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProtectNetListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReverseShellEventInfoRequestParams struct {
	// 事件id
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

type DescribeReverseShellEventInfoRequest struct {
	*tchttp.BaseRequest
	
	// 事件id
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

func (r *DescribeReverseShellEventInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReverseShellEventInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReverseShellEventInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReverseShellEventInfoResponseParams struct {
	// 反弹shell详情信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReverseShellEventInfo *ReverseShellEventInfo `json:"ReverseShellEventInfo,omitnil" name:"ReverseShellEventInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeReverseShellEventInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReverseShellEventInfoResponseParams `json:"Response"`
}

func (r *DescribeReverseShellEventInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReverseShellEventInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReverseShellEventsRequestParams struct {
	// 返回数量，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键字(主机内网IP|进程名)</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 排序方式：根据请求次数排序：asc-升序/desc-降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段：CreateTime-发生时间
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeReverseShellEventsRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键字(主机内网IP|进程名)</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 排序方式：根据请求次数排序：asc-升序/desc-降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段：CreateTime-发生时间
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeReverseShellEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReverseShellEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReverseShellEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReverseShellEventsResponseParams struct {
	// 列表内容
	List []*ReverseShell `json:"List,omitnil" name:"List"`

	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeReverseShellEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReverseShellEventsResponseParams `json:"Response"`
}

func (r *DescribeReverseShellEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReverseShellEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReverseShellRulesRequestParams struct {
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键字(进程名称)</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeReverseShellRulesRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键字(进程名称)</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeReverseShellRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReverseShellRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReverseShellRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReverseShellRulesResponseParams struct {
	// 列表内容
	List []*ReverseShellRule `json:"List,omitnil" name:"List"`

	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeReverseShellRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReverseShellRulesResponseParams `json:"Response"`
}

func (r *DescribeReverseShellRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReverseShellRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskDnsEventInfoRequestParams struct {
	// 恶意请求事件Id
	Id *int64 `json:"Id,omitnil" name:"Id"`
}

type DescribeRiskDnsEventInfoRequest struct {
	*tchttp.BaseRequest
	
	// 恶意请求事件Id
	Id *int64 `json:"Id,omitnil" name:"Id"`
}

func (r *DescribeRiskDnsEventInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskDnsEventInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskDnsEventInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskDnsEventInfoResponseParams struct {
	// 恶意请求事件详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Info *RiskDnsEvent `json:"Info,omitnil" name:"Info"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRiskDnsEventInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskDnsEventInfoResponseParams `json:"Response"`
}

func (r *DescribeRiskDnsEventInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskDnsEventInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskDnsEventListRequestParams struct {
	// <li>IpOrName - String - 是否必填：否 - 主机Ip或别名筛选</li>
	// <li>HostId - String - 是否必填：否 - 主机Id</li>
	// <li>AgentId - String - 是否必填：否 - 客户端Id</li>
	// <li>PolicyType - String - 是否必填：否 - 策略类型,0:系统策略1:用户自定义策略</li>
	// <li>Domain - String - 是否必填：否 - 域名(先对域名做urlencode,再base64)</li>
	// <li>HandleStatus - String - 是否必填：否 - 状态筛选0:待处理；2:信任；3:不信任</li>
	// <li>BeginTime - String - 是否必填：否 - 最近访问开始时间</li>
	// <li>EndTime - String - 是否必填：否 - 最近访问结束时间</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式：根据请求次数排序：[asc:升序|desc:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段：[AccessCount:请求次数|LastTime:最近请求时间]
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeRiskDnsEventListRequest struct {
	*tchttp.BaseRequest
	
	// <li>IpOrName - String - 是否必填：否 - 主机Ip或别名筛选</li>
	// <li>HostId - String - 是否必填：否 - 主机Id</li>
	// <li>AgentId - String - 是否必填：否 - 客户端Id</li>
	// <li>PolicyType - String - 是否必填：否 - 策略类型,0:系统策略1:用户自定义策略</li>
	// <li>Domain - String - 是否必填：否 - 域名(先对域名做urlencode,再base64)</li>
	// <li>HandleStatus - String - 是否必填：否 - 状态筛选0:待处理；2:信任；3:不信任</li>
	// <li>BeginTime - String - 是否必填：否 - 最近访问开始时间</li>
	// <li>EndTime - String - 是否必填：否 - 最近访问结束时间</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式：根据请求次数排序：[asc:升序|desc:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段：[AccessCount:请求次数|LastTime:最近请求时间]
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeRiskDnsEventListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskDnsEventListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskDnsEventListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskDnsEventListResponseParams struct {
	// 恶意请求事件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*RiskDnsEvent `json:"List,omitnil" name:"List"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRiskDnsEventListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskDnsEventListResponseParams `json:"Response"`
}

func (r *DescribeRiskDnsEventListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskDnsEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskDnsListRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Url - String - 是否必填：否 - Url筛选</li>
	// <li>Status - String - 是否必填：否 - 状态筛选0:待处理；2:信任；3:不信任</li>
	// <li>MergeBeginTime - String - 是否必填：否 - 最近访问开始时间</li>
	// <li>MergeEndTime - String - 是否必填：否 - 最近访问结束时间</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 排序方式：根据请求次数排序：asc-升序/desc-降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段：AccessCount-请求次数。MergeTime-最近请求时间
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeRiskDnsListRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Url - String - 是否必填：否 - Url筛选</li>
	// <li>Status - String - 是否必填：否 - 状态筛选0:待处理；2:信任；3:不信任</li>
	// <li>MergeBeginTime - String - 是否必填：否 - 最近访问开始时间</li>
	// <li>MergeEndTime - String - 是否必填：否 - 最近访问结束时间</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 排序方式：根据请求次数排序：asc-升序/desc-降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段：AccessCount-请求次数。MergeTime-最近请求时间
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeRiskDnsListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskDnsListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskDnsListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskDnsListResponseParams struct {
	// 恶意请求列表数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskDnsList []*RiskDnsList `json:"RiskDnsList,omitnil" name:"RiskDnsList"`

	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRiskDnsListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskDnsListResponseParams `json:"Response"`
}

func (r *DescribeRiskDnsListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskDnsListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanMalwareScheduleRequestParams struct {

}

type DescribeScanMalwareScheduleRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeScanMalwareScheduleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanMalwareScheduleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanMalwareScheduleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanMalwareScheduleResponseParams struct {
	// 扫描进度（单位：%）
	Schedule *int64 `json:"Schedule,omitnil" name:"Schedule"`

	// 风险文件数,当进度满了以后才有该值
	RiskFileNumber *int64 `json:"RiskFileNumber,omitnil" name:"RiskFileNumber"`

	// 是否正在扫描中
	IsSchedule *bool `json:"IsSchedule,omitnil" name:"IsSchedule"`

	// 0 从未扫描过、 1 扫描中、 2扫描完成、 3停止中、 4停止完成
	ScanStatus *uint64 `json:"ScanStatus,omitnil" name:"ScanStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeScanMalwareScheduleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScanMalwareScheduleResponseParams `json:"Response"`
}

func (r *DescribeScanMalwareScheduleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanMalwareScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanScheduleRequestParams struct {
	// 任务id
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`
}

type DescribeScanScheduleRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`
}

func (r *DescribeScanScheduleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanScheduleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanScheduleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanScheduleResponseParams struct {
	// 检测进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Schedule *uint64 `json:"Schedule,omitnil" name:"Schedule"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeScanScheduleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScanScheduleResponseParams `json:"Response"`
}

func (r *DescribeScanScheduleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanStateRequestParams struct {
	// 模块类型 当前提供 Malware 木马 , Vul 漏洞 , Baseline 基线
	ModuleType *string `json:"ModuleType,omitnil" name:"ModuleType"`

	// 过滤参数;
	// <li>StrategyId 基线策略ID ,仅ModuleType 为 Baseline 时需要</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`
}

type DescribeScanStateRequest struct {
	*tchttp.BaseRequest
	
	// 模块类型 当前提供 Malware 木马 , Vul 漏洞 , Baseline 基线
	ModuleType *string `json:"ModuleType,omitnil" name:"ModuleType"`

	// 过滤参数;
	// <li>StrategyId 基线策略ID ,仅ModuleType 为 Baseline 时需要</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeScanStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModuleType")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanStateResponseParams struct {
	// 0 从未扫描过、 1 扫描中、 2扫描完成、 3停止中、 4停止完成
	ScanState *uint64 `json:"ScanState,omitnil" name:"ScanState"`

	// 扫描进度
	Schedule *uint64 `json:"Schedule,omitnil" name:"Schedule"`

	// 任务Id
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// 任务扫描的漏洞id
	VulId []*uint64 `json:"VulId,omitnil" name:"VulId"`

	// 0一键检测 1定时检测
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 开始扫描时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanBeginTime *string `json:"ScanBeginTime,omitnil" name:"ScanBeginTime"`

	// 扫描漏洞数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskEventCount *uint64 `json:"RiskEventCount,omitnil" name:"RiskEventCount"`

	// 扫描结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanEndTime *string `json:"ScanEndTime,omitnil" name:"ScanEndTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeScanStateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScanStateResponseParams `json:"Response"`
}

func (r *DescribeScanStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanTaskDetailsRequestParams struct {
	// 模块类型 当前提供 Malware 木马 , Vul 漏洞 , Baseline 基线
	ModuleType *string `json:"ModuleType,omitnil" name:"ModuleType"`

	// 任务ID
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// 过滤参数
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 需要返回的数量，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeScanTaskDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 模块类型 当前提供 Malware 木马 , Vul 漏洞 , Baseline 基线
	ModuleType *string `json:"ModuleType,omitnil" name:"ModuleType"`

	// 任务ID
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// 过滤参数
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 需要返回的数量，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeScanTaskDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanTaskDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModuleType")
	delete(f, "TaskId")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanTaskDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanTaskDetailsResponseParams struct {
	// 扫描任务信息列表
	ScanTaskDetailList []*ScanTaskDetails `json:"ScanTaskDetailList,omitnil" name:"ScanTaskDetailList"`

	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 扫描机器总数
	ScanMachineCount *uint64 `json:"ScanMachineCount,omitnil" name:"ScanMachineCount"`

	// 发现风险机器数
	RiskMachineCount *uint64 `json:"RiskMachineCount,omitnil" name:"RiskMachineCount"`

	// 扫描开始时间
	ScanBeginTime *string `json:"ScanBeginTime,omitnil" name:"ScanBeginTime"`

	// 扫描结束时间
	ScanEndTime *string `json:"ScanEndTime,omitnil" name:"ScanEndTime"`

	// 检测时间
	ScanTime *uint64 `json:"ScanTime,omitnil" name:"ScanTime"`

	// 扫描进度
	ScanProgress *uint64 `json:"ScanProgress,omitnil" name:"ScanProgress"`

	// 扫描剩余时间
	ScanLeftTime *uint64 `json:"ScanLeftTime,omitnil" name:"ScanLeftTime"`

	// 扫描内容
	ScanContent []*string `json:"ScanContent,omitnil" name:"ScanContent"`

	// 漏洞信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulInfo []*VulDetailInfo `json:"VulInfo,omitnil" name:"VulInfo"`

	// 风险事件个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskEventCount *uint64 `json:"RiskEventCount,omitnil" name:"RiskEventCount"`

	// 0一键检测 1定时检测
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 任务是否全部正在被停止 ture是
	// 注意：此字段可能返回 null，表示取不到有效值。
	StoppingAll *bool `json:"StoppingAll,omitnil" name:"StoppingAll"`

	// 扫描出漏洞个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulCount *uint64 `json:"VulCount,omitnil" name:"VulCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeScanTaskDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScanTaskDetailsResponseParams `json:"Response"`
}

func (r *DescribeScanTaskDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanTaskDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanTaskStatusRequestParams struct {
	// 模块类型 当前提供 Malware 木马 , Vul 漏洞 , Baseline 基线
	ModuleType *string `json:"ModuleType,omitnil" name:"ModuleType"`
}

type DescribeScanTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 模块类型 当前提供 Malware 木马 , Vul 漏洞 , Baseline 基线
	ModuleType *string `json:"ModuleType,omitnil" name:"ModuleType"`
}

func (r *DescribeScanTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModuleType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanTaskStatusResponseParams struct {
	// 任务扫描状态列表
	State *TaskStatus `json:"State,omitnil" name:"State"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeScanTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScanTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeScanTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanVulSettingRequestParams struct {

}

type DescribeScanVulSettingRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeScanVulSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanVulSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanVulSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanVulSettingResponseParams struct {
	// 漏洞类型：1: web-cms漏洞 2:应用漏洞  4: Linux软件漏洞 5: Windows系统漏洞
	VulCategories *string `json:"VulCategories,omitnil" name:"VulCategories"`

	// 危害等级：1-低危；2-中危；3-高危；4-严重 (多选英文逗号分隔)
	VulLevels *string `json:"VulLevels,omitnil" name:"VulLevels"`

	// 定期检测间隔时间（天）
	TimerInterval *uint64 `json:"TimerInterval,omitnil" name:"TimerInterval"`

	// 定期检测时间，如：00:00
	TimerTime *string `json:"TimerTime,omitnil" name:"TimerTime"`

	// 是否紧急漏洞：0-否 1-是
	VulEmergency *uint64 `json:"VulEmergency,omitnil" name:"VulEmergency"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 是否开启
	EnableScan *uint64 `json:"EnableScan,omitnil" name:"EnableScan"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 一键扫描超时时长，如：1800秒（s）
	ClickTimeout *uint64 `json:"ClickTimeout,omitnil" name:"ClickTimeout"`

	// 为空默认扫描全部专业版、旗舰版、普惠版主机，不为空只扫描选中主机
	Uuids []*string `json:"Uuids,omitnil" name:"Uuids"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeScanVulSettingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScanVulSettingResponseParams `json:"Response"`
}

func (r *DescribeScanVulSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanVulSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSearchLogsRequestParams struct {

}

type DescribeSearchLogsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSearchLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSearchLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSearchLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSearchLogsResponseParams struct {
	// 历史搜索记录 保留最新的10条
	Data []*string `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSearchLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSearchLogsResponseParams `json:"Response"`
}

func (r *DescribeSearchLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSearchLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSearchTemplatesRequestParams struct {
	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeSearchTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeSearchTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSearchTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSearchTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSearchTemplatesResponseParams struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 模板列表
	List []*SearchTemplate `json:"List,omitnil" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSearchTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSearchTemplatesResponseParams `json:"Response"`
}

func (r *DescribeSearchTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSearchTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityBroadcastsRequestParams struct {
	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10 ，0=全部
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 筛选发布日期：开始时间
	BeginDate *string `json:"BeginDate,omitnil" name:"BeginDate"`

	// 筛选发布日期：结束时间
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`

	// 过滤安全播报类型：0-紧急通知，1-功能更新，2-行业荣誉，3-版本发布
	BroadcastType *string `json:"BroadcastType,omitnil" name:"BroadcastType"`
}

type DescribeSecurityBroadcastsRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要返回的数量，默认为10 ，0=全部
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 筛选发布日期：开始时间
	BeginDate *string `json:"BeginDate,omitnil" name:"BeginDate"`

	// 筛选发布日期：结束时间
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`

	// 过滤安全播报类型：0-紧急通知，1-功能更新，2-行业荣誉，3-版本发布
	BroadcastType *string `json:"BroadcastType,omitnil" name:"BroadcastType"`
}

func (r *DescribeSecurityBroadcastsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityBroadcastsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "BeginDate")
	delete(f, "EndDate")
	delete(f, "BroadcastType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityBroadcastsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityBroadcastsResponseParams struct {
	// 列表
	List []*Broadcasts `json:"List,omitnil" name:"List"`

	// 总共多少条
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSecurityBroadcastsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityBroadcastsResponseParams `json:"Response"`
}

func (r *DescribeSecurityBroadcastsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityBroadcastsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityDynamicsRequestParams struct {
	// 返回数量，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeSecurityDynamicsRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeSecurityDynamicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityDynamicsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityDynamicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityDynamicsResponseParams struct {
	// 安全事件消息数组。
	SecurityDynamics []*SecurityDynamic `json:"SecurityDynamics,omitnil" name:"SecurityDynamics"`

	// 记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSecurityDynamicsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityDynamicsResponseParams `json:"Response"`
}

func (r *DescribeSecurityDynamicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityDynamicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityEventStatRequestParams struct {
	// 该接口无过滤条件
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeSecurityEventStatRequest struct {
	*tchttp.BaseRequest
	
	// 该接口无过滤条件
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeSecurityEventStatRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityEventStatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityEventStatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityEventStatResponseParams struct {
	// 木马事件统计
	MalwareStat *EventStat `json:"MalwareStat,omitnil" name:"MalwareStat"`

	// 异地事件统计
	HostLoginStat *EventStat `json:"HostLoginStat,omitnil" name:"HostLoginStat"`

	// 爆破事件统计
	BruteAttackStat *EventStat `json:"BruteAttackStat,omitnil" name:"BruteAttackStat"`

	// 恶意请求事件统计
	MaliciousRequestStat *EventStat `json:"MaliciousRequestStat,omitnil" name:"MaliciousRequestStat"`

	// 本地提权事件统计
	PrivilegeStat *EventStat `json:"PrivilegeStat,omitnil" name:"PrivilegeStat"`

	// 反弹Shell事件统计
	ReverseShellStat *EventStat `json:"ReverseShellStat,omitnil" name:"ReverseShellStat"`

	// 高危命令事件统计
	HighRiskBashStat *EventStat `json:"HighRiskBashStat,omitnil" name:"HighRiskBashStat"`

	// 网络攻击事件统计
	AttackLogsStat *EventStat `json:"AttackLogsStat,omitnil" name:"AttackLogsStat"`

	// 高危漏洞事件统计
	VulHighStat *EventStat `json:"VulHighStat,omitnil" name:"VulHighStat"`

	// 中危漏洞事件统计
	VulNormalStat *EventStat `json:"VulNormalStat,omitnil" name:"VulNormalStat"`

	// 低危漏洞事件统计
	VulLowStat *EventStat `json:"VulLowStat,omitnil" name:"VulLowStat"`

	// 高危基线漏洞事件统计
	BaselineHighStat *EventStat `json:"BaselineHighStat,omitnil" name:"BaselineHighStat"`

	// 中危基线漏事件统计
	BaselineNormalStat *EventStat `json:"BaselineNormalStat,omitnil" name:"BaselineNormalStat"`

	// 低危基线漏事件统计
	BaselineLowStat *EventStat `json:"BaselineLowStat,omitnil" name:"BaselineLowStat"`

	// 有未处理安全事件的机器总数
	MachineTotalAffectNum *uint64 `json:"MachineTotalAffectNum,omitnil" name:"MachineTotalAffectNum"`

	// 有未处理入侵安全事件的机器总数
	InvasionTotalAffectNum *uint64 `json:"InvasionTotalAffectNum,omitnil" name:"InvasionTotalAffectNum"`

	// 有未处理漏洞安全事件的机器总数
	VulTotalAffectNum *uint64 `json:"VulTotalAffectNum,omitnil" name:"VulTotalAffectNum"`

	// 有未处理基线安全事件的机器总数
	BaseLineTotalAffectNum *uint64 `json:"BaseLineTotalAffectNum,omitnil" name:"BaseLineTotalAffectNum"`

	// 有未处理网络攻击安全事件的机器总数
	CyberAttackTotalAffectNum *uint64 `json:"CyberAttackTotalAffectNum,omitnil" name:"CyberAttackTotalAffectNum"`

	// 严重漏洞事件统计
	VulRiskStat *EventStat `json:"VulRiskStat,omitnil" name:"VulRiskStat"`

	// 严重基线漏洞事件统计
	BaselineRiskStat *EventStat `json:"BaselineRiskStat,omitnil" name:"BaselineRiskStat"`

	// 漏洞数统计
	VulStat *EventStat `json:"VulStat,omitnil" name:"VulStat"`

	// 安全得分
	Score *uint64 `json:"Score,omitnil" name:"Score"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSecurityEventStatResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityEventStatResponseParams `json:"Response"`
}

func (r *DescribeSecurityEventStatResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityEventStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityEventsCntRequestParams struct {

}

type DescribeSecurityEventsCntRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSecurityEventsCntRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityEventsCntRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityEventsCntRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityEventsCntResponseParams struct {
	// 木马文件相关风险事件
	Malware *SecurityEventInfo `json:"Malware,omitnil" name:"Malware"`

	// 登录审计相关风险事件
	HostLogin *SecurityEventInfo `json:"HostLogin,omitnil" name:"HostLogin"`

	// 密码破解相关风险事件
	BruteAttack *SecurityEventInfo `json:"BruteAttack,omitnil" name:"BruteAttack"`

	// 恶意请求相关风险事件
	RiskDns *SecurityEventInfo `json:"RiskDns,omitnil" name:"RiskDns"`

	// 高危命令相关风险事件
	Bash *SecurityEventInfo `json:"Bash,omitnil" name:"Bash"`

	// 本地提权相关风险事件
	PrivilegeRules *SecurityEventInfo `json:"PrivilegeRules,omitnil" name:"PrivilegeRules"`

	// 反弹Shell相关风险事件
	ReverseShell *SecurityEventInfo `json:"ReverseShell,omitnil" name:"ReverseShell"`

	// 应用漏洞风险事件
	SysVul *SecurityEventInfo `json:"SysVul,omitnil" name:"SysVul"`

	// Web应用漏洞相关风险事件
	WebVul *SecurityEventInfo `json:"WebVul,omitnil" name:"WebVul"`

	// 应急漏洞相关风险事件
	EmergencyVul *SecurityEventInfo `json:"EmergencyVul,omitnil" name:"EmergencyVul"`

	// 安全基线相关风险事件
	BaseLine *SecurityEventInfo `json:"BaseLine,omitnil" name:"BaseLine"`

	// 攻击检测相关风险事件
	AttackLogs *SecurityEventInfo `json:"AttackLogs,omitnil" name:"AttackLogs"`

	// 受影响机器数
	EffectMachineCount *uint64 `json:"EffectMachineCount,omitnil" name:"EffectMachineCount"`

	// 所有事件总数
	EventsCount *uint64 `json:"EventsCount,omitnil" name:"EventsCount"`

	// window 系统漏洞事件总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	WindowVul *SecurityEventInfo `json:"WindowVul,omitnil" name:"WindowVul"`

	// linux系统漏洞事件总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	LinuxVul *SecurityEventInfo `json:"LinuxVul,omitnil" name:"LinuxVul"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSecurityEventsCntResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityEventsCntResponseParams `json:"Response"`
}

func (r *DescribeSecurityEventsCntResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityEventsCntResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityTrendsRequestParams struct {
	// 开始时间，如：2021-07-10
	BeginDate *string `json:"BeginDate,omitnil" name:"BeginDate"`

	// 结束时间，如：2021-07-10
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`
}

type DescribeSecurityTrendsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间，如：2021-07-10
	BeginDate *string `json:"BeginDate,omitnil" name:"BeginDate"`

	// 结束时间，如：2021-07-10
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`
}

func (r *DescribeSecurityTrendsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityTrendsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityTrendsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityTrendsResponseParams struct {
	// 木马事件统计数据数组。
	Malwares []*SecurityTrend `json:"Malwares,omitnil" name:"Malwares"`

	// 异地登录事件统计数据数组。
	NonLocalLoginPlaces []*SecurityTrend `json:"NonLocalLoginPlaces,omitnil" name:"NonLocalLoginPlaces"`

	// 密码破解事件统计数据数组。
	BruteAttacks []*SecurityTrend `json:"BruteAttacks,omitnil" name:"BruteAttacks"`

	// 漏洞统计数据数组。
	Vuls []*SecurityTrend `json:"Vuls,omitnil" name:"Vuls"`

	// 基线统计数据数组。
	BaseLines []*SecurityTrend `json:"BaseLines,omitnil" name:"BaseLines"`

	// 恶意请求统计数据数组。
	MaliciousRequests []*SecurityTrend `json:"MaliciousRequests,omitnil" name:"MaliciousRequests"`

	// 高危命令统计数据数组。
	HighRiskBashs []*SecurityTrend `json:"HighRiskBashs,omitnil" name:"HighRiskBashs"`

	// 反弹shell统计数据数组。
	ReverseShells []*SecurityTrend `json:"ReverseShells,omitnil" name:"ReverseShells"`

	// 本地提权统计数据数组。
	PrivilegeEscalations []*SecurityTrend `json:"PrivilegeEscalations,omitnil" name:"PrivilegeEscalations"`

	// 网络攻击统计数据数组。
	CyberAttacks []*SecurityTrend `json:"CyberAttacks,omitnil" name:"CyberAttacks"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSecurityTrendsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityTrendsResponseParams `json:"Response"`
}

func (r *DescribeSecurityTrendsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityTrendsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServerRelatedDirInfoRequestParams struct {
	// 唯一ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

type DescribeServerRelatedDirInfoRequest struct {
	*tchttp.BaseRequest
	
	// 唯一ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

func (r *DescribeServerRelatedDirInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServerRelatedDirInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServerRelatedDirInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServerRelatedDirInfoResponseParams struct {
	// 服务器名称
	HostName *string `json:"HostName,omitnil" name:"HostName"`

	// 服务器IP
	HostIp *string `json:"HostIp,omitnil" name:"HostIp"`

	// 防护目录数量
	ProtectDirNum *uint64 `json:"ProtectDirNum,omitnil" name:"ProtectDirNum"`

	// 防护文件数量
	ProtectFileNum *uint64 `json:"ProtectFileNum,omitnil" name:"ProtectFileNum"`

	// 防篡改数量
	ProtectTamperNum *uint64 `json:"ProtectTamperNum,omitnil" name:"ProtectTamperNum"`

	// 防护软链数量
	ProtectLinkNum *uint64 `json:"ProtectLinkNum,omitnil" name:"ProtectLinkNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeServerRelatedDirInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServerRelatedDirInfoResponseParams `json:"Response"`
}

func (r *DescribeServerRelatedDirInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServerRelatedDirInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServersAndRiskAndFirstInfoRequestParams struct {

}

type DescribeServersAndRiskAndFirstInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeServersAndRiskAndFirstInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServersAndRiskAndFirstInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServersAndRiskAndFirstInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServersAndRiskAndFirstInfoResponseParams struct {
	// 风险文件数
	RiskFileCount *uint64 `json:"RiskFileCount,omitnil" name:"RiskFileCount"`

	// 今日新增风险文件数
	AddRiskFileCount *uint64 `json:"AddRiskFileCount,omitnil" name:"AddRiskFileCount"`

	// 受影响服务器台数
	ServersCount *uint64 `json:"ServersCount,omitnil" name:"ServersCount"`

	// 是否试用：true-是，false-否
	IsFirstCheck *bool `json:"IsFirstCheck,omitnil" name:"IsFirstCheck"`

	// 木马最近检测时间
	ScanTime *string `json:"ScanTime,omitnil" name:"ScanTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeServersAndRiskAndFirstInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServersAndRiskAndFirstInfoResponseParams `json:"Response"`
}

func (r *DescribeServersAndRiskAndFirstInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServersAndRiskAndFirstInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStrategyExistRequestParams struct {
	// 策略名
	StrategyName *string `json:"StrategyName,omitnil" name:"StrategyName"`
}

type DescribeStrategyExistRequest struct {
	*tchttp.BaseRequest
	
	// 策略名
	StrategyName *string `json:"StrategyName,omitnil" name:"StrategyName"`
}

func (r *DescribeStrategyExistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStrategyExistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StrategyName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStrategyExistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStrategyExistResponseParams struct {
	// 策略是否存在, 1是 0否
	// 注意：此字段可能返回 null，表示取不到有效值。
	IfExist *uint64 `json:"IfExist,omitnil" name:"IfExist"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStrategyExistResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStrategyExistResponseParams `json:"Response"`
}

func (r *DescribeStrategyExistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStrategyExistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagMachinesRequestParams struct {
	// 标签ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

type DescribeTagMachinesRequest struct {
	*tchttp.BaseRequest
	
	// 标签ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

func (r *DescribeTagMachinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagMachinesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagMachinesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagMachinesResponseParams struct {
	// 列表数据
	List []*TagMachine `json:"List,omitnil" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTagMachinesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTagMachinesResponseParams `json:"Response"`
}

func (r *DescribeTagMachinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagsRequestParams struct {
	// 云主机类型。
	// <li>CVM：表示云服务器</li>
	// <li>BM:  表示黑石物理机</li>
	// <li>ECM:  表示边缘计算服务器</li>
	// <li>LH:  表示轻量应用服务器</li>
	// <li>Other:  表示混合云服务器</li>
	MachineType *string `json:"MachineType,omitnil" name:"MachineType"`

	// 机器所属地域。如：ap-guangzhou
	MachineRegion *string `json:"MachineRegion,omitnil" name:"MachineRegion"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字(机器名称/机器IP </li>
	// <li>Status - String - 是否必填：否 - 客户端在线状态（OFFLINE: 离线 | ONLINE: 在线 | UNINSTALLED：未安装 | SHUTDOWN 已关机）</li>
	// <li>Version - String  是否必填：否 - 当前防护版本（ PRO_VERSION：专业版 | BASIC_VERSION：基础版）</li>
	// <li>Risk - String 是否必填: 否 - 风险主机( yes ) </li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )
	// 每个过滤条件只支持一个值，暂不支持多个值“或”关系查询</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`
}

type DescribeTagsRequest struct {
	*tchttp.BaseRequest
	
	// 云主机类型。
	// <li>CVM：表示云服务器</li>
	// <li>BM:  表示黑石物理机</li>
	// <li>ECM:  表示边缘计算服务器</li>
	// <li>LH:  表示轻量应用服务器</li>
	// <li>Other:  表示混合云服务器</li>
	MachineType *string `json:"MachineType,omitnil" name:"MachineType"`

	// 机器所属地域。如：ap-guangzhou
	MachineRegion *string `json:"MachineRegion,omitnil" name:"MachineRegion"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字(机器名称/机器IP </li>
	// <li>Status - String - 是否必填：否 - 客户端在线状态（OFFLINE: 离线 | ONLINE: 在线 | UNINSTALLED：未安装 | SHUTDOWN 已关机）</li>
	// <li>Version - String  是否必填：否 - 当前防护版本（ PRO_VERSION：专业版 | BASIC_VERSION：基础版）</li>
	// <li>Risk - String 是否必填: 否 - 风险主机( yes ) </li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )
	// 每个过滤条件只支持一个值，暂不支持多个值“或”关系查询</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`
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
	delete(f, "MachineType")
	delete(f, "MachineRegion")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagsResponseParams struct {
	// 列表信息
	List []*Tag `json:"List,omitnil" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeUndoVulCountsRequestParams struct {
	// 漏洞分类，1: web-cms漏洞 2:应用漏洞  4: Linux软件漏洞 5: Windows系统漏洞
	VulCategory *uint64 `json:"VulCategory,omitnil" name:"VulCategory"`

	// 是否应急漏洞筛选, 是 : yes
	IfEmergency *string `json:"IfEmergency,omitnil" name:"IfEmergency"`
}

type DescribeUndoVulCountsRequest struct {
	*tchttp.BaseRequest
	
	// 漏洞分类，1: web-cms漏洞 2:应用漏洞  4: Linux软件漏洞 5: Windows系统漏洞
	VulCategory *uint64 `json:"VulCategory,omitnil" name:"VulCategory"`

	// 是否应急漏洞筛选, 是 : yes
	IfEmergency *string `json:"IfEmergency,omitnil" name:"IfEmergency"`
}

func (r *DescribeUndoVulCountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUndoVulCountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VulCategory")
	delete(f, "IfEmergency")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUndoVulCountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUndoVulCountsResponseParams struct {
	// 未处理的漏洞数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UndoVulCount *uint64 `json:"UndoVulCount,omitnil" name:"UndoVulCount"`

	// 未处理的主机数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UndoHostCount *int64 `json:"UndoHostCount,omitnil" name:"UndoHostCount"`

	// 普通版主机数
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotProfessionCount *uint64 `json:"NotProfessionCount,omitnil" name:"NotProfessionCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeUndoVulCountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUndoVulCountsResponseParams `json:"Response"`
}

func (r *DescribeUndoVulCountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUndoVulCountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsualLoginPlacesRequestParams struct {
	// 主机安全客户端UUID
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

type DescribeUsualLoginPlacesRequest struct {
	*tchttp.BaseRequest
	
	// 主机安全客户端UUID
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

func (r *DescribeUsualLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsualLoginPlacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsualLoginPlacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsualLoginPlacesResponseParams struct {
	// 常用登录地数组
	UsualLoginPlaces []*UsualPlace `json:"UsualLoginPlaces,omitnil" name:"UsualLoginPlaces"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeUsualLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsualLoginPlacesResponseParams `json:"Response"`
}

func (r *DescribeUsualLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsualLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVersionStatisticsRequestParams struct {

}

type DescribeVersionStatisticsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeVersionStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVersionStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVersionStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVersionStatisticsResponseParams struct {
	// 基础版数量
	BasicVersionNum *uint64 `json:"BasicVersionNum,omitnil" name:"BasicVersionNum"`

	// 专业版数量
	ProVersionNum *uint64 `json:"ProVersionNum,omitnil" name:"ProVersionNum"`

	// 旗舰版数量
	UltimateVersionNum *uint64 `json:"UltimateVersionNum,omitnil" name:"UltimateVersionNum"`

	// 普惠版数量
	GeneralVersionNum *uint64 `json:"GeneralVersionNum,omitnil" name:"GeneralVersionNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVersionStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVersionStatisticsResponseParams `json:"Response"`
}

func (r *DescribeVersionStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVersionStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVertexDetailRequestParams struct {
	// 点id列表
	VertexIds []*string `json:"VertexIds,omitnil" name:"VertexIds"`

	// 事件id
	IncidentId *string `json:"IncidentId,omitnil" name:"IncidentId"`

	// 事件所在表名
	TableName *string `json:"TableName,omitnil" name:"TableName"`
}

type DescribeVertexDetailRequest struct {
	*tchttp.BaseRequest
	
	// 点id列表
	VertexIds []*string `json:"VertexIds,omitnil" name:"VertexIds"`

	// 事件id
	IncidentId *string `json:"IncidentId,omitnil" name:"IncidentId"`

	// 事件所在表名
	TableName *string `json:"TableName,omitnil" name:"TableName"`
}

func (r *DescribeVertexDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVertexDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VertexIds")
	delete(f, "IncidentId")
	delete(f, "TableName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVertexDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVertexDetailResponseParams struct {
	// 指定点列表的属性信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VertexDetails []*VertexDetail `json:"VertexDetails,omitnil" name:"VertexDetails"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVertexDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVertexDetailResponseParams `json:"Response"`
}

func (r *DescribeVertexDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVertexDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulCountByDatesRequestParams struct {
	// 需要查询最近几天的数据，需要都 -1后传入
	LastDays []*uint64 `json:"LastDays,omitnil" name:"LastDays"`

	// 漏洞的分类: 1: web-cms漏洞 2:应用漏洞  4: Linux软件漏洞 5: Windows系统漏洞
	VulCategory *uint64 `json:"VulCategory,omitnil" name:"VulCategory"`

	// 是否为应急漏洞筛选  是: yes
	IfEmergency *string `json:"IfEmergency,omitnil" name:"IfEmergency"`
}

type DescribeVulCountByDatesRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询最近几天的数据，需要都 -1后传入
	LastDays []*uint64 `json:"LastDays,omitnil" name:"LastDays"`

	// 漏洞的分类: 1: web-cms漏洞 2:应用漏洞  4: Linux软件漏洞 5: Windows系统漏洞
	VulCategory *uint64 `json:"VulCategory,omitnil" name:"VulCategory"`

	// 是否为应急漏洞筛选  是: yes
	IfEmergency *string `json:"IfEmergency,omitnil" name:"IfEmergency"`
}

func (r *DescribeVulCountByDatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulCountByDatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LastDays")
	delete(f, "VulCategory")
	delete(f, "IfEmergency")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulCountByDatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulCountByDatesResponseParams struct {
	// 批量获得对应天数的漏洞数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulCount []*uint64 `json:"VulCount,omitnil" name:"VulCount"`

	// 批量获得对应天数的主机数量
	HostCount []*uint64 `json:"HostCount,omitnil" name:"HostCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVulCountByDatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulCountByDatesResponseParams `json:"Response"`
}

func (r *DescribeVulCountByDatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulCountByDatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulEffectHostListRequestParams struct {
	// 分页limit 最大100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页Offset
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 漏洞id
	VulId *uint64 `json:"VulId,omitnil" name:"VulId"`

	// 过滤条件：
	// <li>AliasName - String - 主机名筛选</li>
	// <li>TagIds - String - 主机标签id串，多个用英文用逗号分隔</li>
	// <li>Status - String - 状态：0-待处理 1-忽略  3-已修复  5-检测中  6-修复中  8-修复失败</li>
	// <li>Uuid - String数组 - Uuid串数组</li>
	// <li>Version - String数组 - 付费版本数组："Flagship"-旗舰版 "PRO_VERSION"-专业版 "BASIC_VERSION"-基础版</li>
	// <li>InstanceState - String数组 - 实例状态数组："PENDING"-创建中 "LAUNCH_FAILED"-创建失败 "RUNNING"-运行中 "STOPPED"-关机 "STARTING"-开机中 "STOPPING"-关机中 "REBOOTING"-重启中 "SHUTDOWN"-待销毁 "TERMINATING"-销毁中 "UNKNOWN"-未知（针对非腾讯云机器，且客户端离线的场景） </li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeVulEffectHostListRequest struct {
	*tchttp.BaseRequest
	
	// 分页limit 最大100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页Offset
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 漏洞id
	VulId *uint64 `json:"VulId,omitnil" name:"VulId"`

	// 过滤条件：
	// <li>AliasName - String - 主机名筛选</li>
	// <li>TagIds - String - 主机标签id串，多个用英文用逗号分隔</li>
	// <li>Status - String - 状态：0-待处理 1-忽略  3-已修复  5-检测中  6-修复中  8-修复失败</li>
	// <li>Uuid - String数组 - Uuid串数组</li>
	// <li>Version - String数组 - 付费版本数组："Flagship"-旗舰版 "PRO_VERSION"-专业版 "BASIC_VERSION"-基础版</li>
	// <li>InstanceState - String数组 - 实例状态数组："PENDING"-创建中 "LAUNCH_FAILED"-创建失败 "RUNNING"-运行中 "STOPPED"-关机 "STARTING"-开机中 "STOPPING"-关机中 "REBOOTING"-重启中 "SHUTDOWN"-待销毁 "TERMINATING"-销毁中 "UNKNOWN"-未知（针对非腾讯云机器，且客户端离线的场景） </li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeVulEffectHostListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulEffectHostListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "VulId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulEffectHostListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulEffectHostListResponseParams struct {
	// 列表总数量
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 影响主机列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulEffectHostList []*VulEffectHostList `json:"VulEffectHostList,omitnil" name:"VulEffectHostList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVulEffectHostListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulEffectHostListResponseParams `json:"Response"`
}

func (r *DescribeVulEffectHostListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulEffectHostListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulHostCountScanTimeRequestParams struct {

}

type DescribeVulHostCountScanTimeRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeVulHostCountScanTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulHostCountScanTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulHostCountScanTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulHostCountScanTimeResponseParams struct {
	// 总漏洞数
	TotalVulCount *uint64 `json:"TotalVulCount,omitnil" name:"TotalVulCount"`

	// 漏洞影响主机数
	VulHostCount *uint64 `json:"VulHostCount,omitnil" name:"VulHostCount"`

	// 扫描时间
	ScanTime *string `json:"ScanTime,omitnil" name:"ScanTime"`

	// 是否第一次检测
	IfFirstScan *bool `json:"IfFirstScan,omitnil" name:"IfFirstScan"`

	// 运行中的任务号, 没有任务则为0
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// 最后一次修复漏洞的时间
	LastFixTime *string `json:"LastFixTime,omitnil" name:"LastFixTime"`

	// 是否有支持自动修复的漏洞事件
	hadAutoFixVul *bool `json:"hadAutoFixVul,omitnil" name:"hadAutoFixVul"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVulHostCountScanTimeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulHostCountScanTimeResponseParams `json:"Response"`
}

func (r *DescribeVulHostCountScanTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulHostCountScanTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulHostTopRequestParams struct {
	// 获取top值，1-100
	Top *uint64 `json:"Top,omitnil" name:"Top"`

	// 1:web-cms 漏洞，2.应用漏洞   4: Linux软件漏洞 5: windows系统漏洞 6:应急漏洞，不填或者填0时返回 1，2，4，5 的总统计数据
	VulCategory *uint64 `json:"VulCategory,omitnil" name:"VulCategory"`

	// 是否仅统计重点关注漏洞 1=仅统计重点关注漏洞, 0=统计全部漏洞
	IsFollowVul *uint64 `json:"IsFollowVul,omitnil" name:"IsFollowVul"`
}

type DescribeVulHostTopRequest struct {
	*tchttp.BaseRequest
	
	// 获取top值，1-100
	Top *uint64 `json:"Top,omitnil" name:"Top"`

	// 1:web-cms 漏洞，2.应用漏洞   4: Linux软件漏洞 5: windows系统漏洞 6:应急漏洞，不填或者填0时返回 1，2，4，5 的总统计数据
	VulCategory *uint64 `json:"VulCategory,omitnil" name:"VulCategory"`

	// 是否仅统计重点关注漏洞 1=仅统计重点关注漏洞, 0=统计全部漏洞
	IsFollowVul *uint64 `json:"IsFollowVul,omitnil" name:"IsFollowVul"`
}

func (r *DescribeVulHostTopRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulHostTopRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Top")
	delete(f, "VulCategory")
	delete(f, "IsFollowVul")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulHostTopRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulHostTopResponseParams struct {
	// 服务器风险top列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulHostTopList []*VulHostTopInfo `json:"VulHostTopList,omitnil" name:"VulHostTopList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVulHostTopResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulHostTopResponseParams `json:"Response"`
}

func (r *DescribeVulHostTopResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulHostTopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulInfoCvssRequestParams struct {
	// 漏洞id
	VulId *uint64 `json:"VulId,omitnil" name:"VulId"`
}

type DescribeVulInfoCvssRequest struct {
	*tchttp.BaseRequest
	
	// 漏洞id
	VulId *uint64 `json:"VulId,omitnil" name:"VulId"`
}

func (r *DescribeVulInfoCvssRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulInfoCvssRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VulId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulInfoCvssRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulInfoCvssResponseParams struct {
	// 漏洞id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulId *uint64 `json:"VulId,omitnil" name:"VulId"`

	// 漏洞名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulName *string `json:"VulName,omitnil" name:"VulName"`

	// 危害等级：1-低危；2-中危；3-高危；4-严重
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulLevel *uint64 `json:"VulLevel,omitnil" name:"VulLevel"`

	// 漏洞分类 1: web-cms漏洞 2:应用漏洞  4: Linux软件漏洞 5: Windows系统漏洞
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulType *uint64 `json:"VulType,omitnil" name:"VulType"`

	// 漏洞描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 修复方案
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepairPlan *string `json:"RepairPlan,omitnil" name:"RepairPlan"`

	// 漏洞CVEID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CveId *string `json:"CveId,omitnil" name:"CveId"`

	// 参考链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reference *string `json:"Reference,omitnil" name:"Reference"`

	// CVSS信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CVSS *string `json:"CVSS,omitnil" name:"CVSS"`

	// 发布时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicDate *string `json:"PublicDate,omitnil" name:"PublicDate"`

	// Cvss分数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CvssScore *uint64 `json:"CvssScore,omitnil" name:"CvssScore"`

	// cvss详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	CveInfo *string `json:"CveInfo,omitnil" name:"CveInfo"`

	// cvss 分数 浮点型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CvssScoreFloat *float64 `json:"CvssScoreFloat,omitnil" name:"CvssScoreFloat"`

	// 漏洞标签 多个逗号分割
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels *string `json:"Labels,omitnil" name:"Labels"`

	// 已防御的攻击次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefenseAttackCount *uint64 `json:"DefenseAttackCount,omitnil" name:"DefenseAttackCount"`

	// 全网修复成功次数, 不支持自动修复的漏洞默认返回0
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessFixCount *uint64 `json:"SuccessFixCount,omitnil" name:"SuccessFixCount"`

	// 修复是否支持：0-windows/linux均不支持修复 ;1-windows/linux 均支持修复 ;2-仅linux支持修复;3-仅windows支持修复
	// 注意：此字段可能返回 null，表示取不到有效值。
	FixSwitch *int64 `json:"FixSwitch,omitnil" name:"FixSwitch"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVulInfoCvssResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulInfoCvssResponseParams `json:"Response"`
}

func (r *DescribeVulInfoCvssResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulInfoCvssResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulLevelCountRequestParams struct {
	// 1:web-cms 漏洞，2.应用漏洞 3:安全基线 4: Linux软件漏洞 5: windows系统漏洞 6:应急漏洞，不填或者填0时返回 1，2，4，5 的总统计数据
	VulCategory *uint64 `json:"VulCategory,omitnil" name:"VulCategory"`

	// 是否仅统计重点关注漏洞 1=仅统计重点关注漏洞, 0=统计全部漏洞
	IsFollowVul *uint64 `json:"IsFollowVul,omitnil" name:"IsFollowVul"`
}

type DescribeVulLevelCountRequest struct {
	*tchttp.BaseRequest
	
	// 1:web-cms 漏洞，2.应用漏洞 3:安全基线 4: Linux软件漏洞 5: windows系统漏洞 6:应急漏洞，不填或者填0时返回 1，2，4，5 的总统计数据
	VulCategory *uint64 `json:"VulCategory,omitnil" name:"VulCategory"`

	// 是否仅统计重点关注漏洞 1=仅统计重点关注漏洞, 0=统计全部漏洞
	IsFollowVul *uint64 `json:"IsFollowVul,omitnil" name:"IsFollowVul"`
}

func (r *DescribeVulLevelCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulLevelCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VulCategory")
	delete(f, "IsFollowVul")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulLevelCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulLevelCountResponseParams struct {
	// 统计结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulLevelList []*VulLevelInfo `json:"VulLevelList,omitnil" name:"VulLevelList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVulLevelCountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulLevelCountResponseParams `json:"Response"`
}

func (r *DescribeVulLevelCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulLevelCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulListRequestParams struct {
	// 返回数量，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>Status - String - 是否必填：否 - 处理状态  0 -- 待处理 1 -- 已加白 2 -- 已删除 3 - 已忽略</li>
	// <li>ModifyTime - String - 是否必填：否 - 最近发生时间</li>
	// <li>Uuid- String - 是否必填：否 - 主机uuid查询</li>
	// <li>VulName- string -</li>
	// <li>VulCategory- string - 是否必填：否 - 漏洞类别 1: web-cms漏洞 2:应用漏洞  4: Linux软件漏洞 5: Windows系统漏洞</li>
	// <li>IsSupportDefense - int- 是否必填：否 - 是否支持防御 0:不支持 1:支持</li>
	// <li>Labels- string- 是否必填：否 - 标签搜索</li>
	// <li>IsSupportAutoFix- string- 是否必填：否 - 是否支持自动修复 0:不支持 1:支持</li>
	// <li>CvssScore- string- 是否必填：否 - CvssScore大于多少</li>
	// <li>AttackLevel- string- 是否必填：否 - 攻击热度大于多少</li>
	// 
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 可选排序字段 Level，LastTime，HostCount
	By *string `json:"By,omitnil" name:"By"`

	// 排序顺序：desc  默认asc
	Order *string `json:"Order,omitnil" name:"Order"`
}

type DescribeVulListRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件。
	// <li>Status - String - 是否必填：否 - 处理状态  0 -- 待处理 1 -- 已加白 2 -- 已删除 3 - 已忽略</li>
	// <li>ModifyTime - String - 是否必填：否 - 最近发生时间</li>
	// <li>Uuid- String - 是否必填：否 - 主机uuid查询</li>
	// <li>VulName- string -</li>
	// <li>VulCategory- string - 是否必填：否 - 漏洞类别 1: web-cms漏洞 2:应用漏洞  4: Linux软件漏洞 5: Windows系统漏洞</li>
	// <li>IsSupportDefense - int- 是否必填：否 - 是否支持防御 0:不支持 1:支持</li>
	// <li>Labels- string- 是否必填：否 - 标签搜索</li>
	// <li>IsSupportAutoFix- string- 是否必填：否 - 是否支持自动修复 0:不支持 1:支持</li>
	// <li>CvssScore- string- 是否必填：否 - CvssScore大于多少</li>
	// <li>AttackLevel- string- 是否必填：否 - 攻击热度大于多少</li>
	// 
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 可选排序字段 Level，LastTime，HostCount
	By *string `json:"By,omitnil" name:"By"`

	// 排序顺序：desc  默认asc
	Order *string `json:"Order,omitnil" name:"Order"`
}

func (r *DescribeVulListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulListResponseParams struct {
	// 漏洞列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulInfoList []*VulInfoList `json:"VulInfoList,omitnil" name:"VulInfoList"`

	// 漏洞总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 重点关注漏洞总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FollowVulCount *uint64 `json:"FollowVulCount,omitnil" name:"FollowVulCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVulListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulListResponseParams `json:"Response"`
}

func (r *DescribeVulListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulStoreListRequestParams struct {
	// 
	// <li>VulName- string - 是否必填：否 - 漏洞名称</li>
	// <li>CveId- string - 是否必填：否 - cveid</li>
	// <li>VulCategory- string - 是否必填：否 - 漏洞分类  1 Web-CMS漏洞 ,2 应用漏洞 ,4 Linux软件漏洞,5 Windows系统漏洞</li>
	// <li>Method- string - 是否必填：否 - 检测方法 0版本对比,1 poc检测 </li>
	// <li>SupportDefense- string - 是否必填：否 - 是否支持防御 0不支持,1支持</li>
	// <li>FixSwitch- string - 是否必填：否 - 是否支持自动修复 0不支持,1支持</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 限制条数,默认10,最大100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式: [ASC:升序|DESC:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序列: [PublishDate]
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeVulStoreListRequest struct {
	*tchttp.BaseRequest
	
	// 
	// <li>VulName- string - 是否必填：否 - 漏洞名称</li>
	// <li>CveId- string - 是否必填：否 - cveid</li>
	// <li>VulCategory- string - 是否必填：否 - 漏洞分类  1 Web-CMS漏洞 ,2 应用漏洞 ,4 Linux软件漏洞,5 Windows系统漏洞</li>
	// <li>Method- string - 是否必填：否 - 检测方法 0版本对比,1 poc检测 </li>
	// <li>SupportDefense- string - 是否必填：否 - 是否支持防御 0不支持,1支持</li>
	// <li>FixSwitch- string - 是否必填：否 - 是否支持自动修复 0不支持,1支持</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 限制条数,默认10,最大100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序方式: [ASC:升序|DESC:降序]
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序列: [PublishDate]
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeVulStoreListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulStoreListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulStoreListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulStoreListResponseParams struct {
	// 漏洞信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*VulStoreListInfo `json:"List,omitnil" name:"List"`

	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 今日剩余搜索此时
	Remaining *uint64 `json:"Remaining,omitnil" name:"Remaining"`

	// 免费搜索次数
	FreeSearchTimes *uint64 `json:"FreeSearchTimes,omitnil" name:"FreeSearchTimes"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVulStoreListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulStoreListResponseParams `json:"Response"`
}

func (r *DescribeVulStoreListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulStoreListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulTopRequestParams struct {
	// 漏洞风险服务器top，1-100
	Top *uint64 `json:"Top,omitnil" name:"Top"`

	// 1:web-cms 漏洞，2.应用漏洞 4: Linux软件漏洞 5: windows系统漏洞 6:应急漏洞，不填或者填0时返回 1，2，4，5 的总统计数据
	VulCategory *uint64 `json:"VulCategory,omitnil" name:"VulCategory"`

	// 是否仅统计重点关注漏洞 1=仅统计重点关注漏洞, 0=统计全部漏洞
	IsFollowVul *uint64 `json:"IsFollowVul,omitnil" name:"IsFollowVul"`
}

type DescribeVulTopRequest struct {
	*tchttp.BaseRequest
	
	// 漏洞风险服务器top，1-100
	Top *uint64 `json:"Top,omitnil" name:"Top"`

	// 1:web-cms 漏洞，2.应用漏洞 4: Linux软件漏洞 5: windows系统漏洞 6:应急漏洞，不填或者填0时返回 1，2，4，5 的总统计数据
	VulCategory *uint64 `json:"VulCategory,omitnil" name:"VulCategory"`

	// 是否仅统计重点关注漏洞 1=仅统计重点关注漏洞, 0=统计全部漏洞
	IsFollowVul *uint64 `json:"IsFollowVul,omitnil" name:"IsFollowVul"`
}

func (r *DescribeVulTopRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulTopRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Top")
	delete(f, "VulCategory")
	delete(f, "IsFollowVul")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulTopRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulTopResponseParams struct {
	// 漏洞top列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulTopList []*VulTopInfo `json:"VulTopList,omitnil" name:"VulTopList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVulTopResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulTopResponseParams `json:"Response"`
}

func (r *DescribeVulTopResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulTopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWarningListRequestParams struct {

}

type DescribeWarningListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeWarningListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWarningListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWarningListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWarningListResponseParams struct {
	// 获取告警列表
	WarningInfoList []*WarningInfoObj `json:"WarningInfoList,omitnil" name:"WarningInfoList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWarningListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWarningListResponseParams `json:"Response"`
}

func (r *DescribeWarningListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWarningListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebPageEventListRequestParams struct {
	// 过滤条件
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>EventType - String - 是否必填：否 - 事件类型</li>
	// <li>EventStatus - String - 是否必填：否 - 事件状态</li>
	Filters []*AssetFilters `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式：CreateTime 或 RestoreTime，默认为CreateTime
	By *string `json:"By,omitnil" name:"By"`

	// 排序方式，0降序，1升序，默认为0
	Order *uint64 `json:"Order,omitnil" name:"Order"`
}

type DescribeWebPageEventListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>EventType - String - 是否必填：否 - 事件类型</li>
	// <li>EventStatus - String - 是否必填：否 - 事件状态</li>
	Filters []*AssetFilters `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式：CreateTime 或 RestoreTime，默认为CreateTime
	By *string `json:"By,omitnil" name:"By"`

	// 排序方式，0降序，1升序，默认为0
	Order *uint64 `json:"Order,omitnil" name:"Order"`
}

func (r *DescribeWebPageEventListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebPageEventListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebPageEventListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebPageEventListResponseParams struct {
	// 防护事件列表信息
	List []*ProtectEventLists `json:"List,omitnil" name:"List"`

	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWebPageEventListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebPageEventListResponseParams `json:"Response"`
}

func (r *DescribeWebPageEventListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebPageEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebPageGeneralizeRequestParams struct {

}

type DescribeWebPageGeneralizeRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeWebPageGeneralizeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebPageGeneralizeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebPageGeneralizeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebPageGeneralizeResponseParams struct {
	// 防护监测 0 未开启 1 已开启 2 异常
	ProtectMonitor *uint64 `json:"ProtectMonitor,omitnil" name:"ProtectMonitor"`

	// 防护目录数
	ProtectDirNum *uint64 `json:"ProtectDirNum,omitnil" name:"ProtectDirNum"`

	// 防护文件数
	ProtectFileNum *uint64 `json:"ProtectFileNum,omitnil" name:"ProtectFileNum"`

	// 篡改文件数
	TamperFileNum *uint64 `json:"TamperFileNum,omitnil" name:"TamperFileNum"`

	// 篡改数
	TamperNum *uint64 `json:"TamperNum,omitnil" name:"TamperNum"`

	// 今日防护数
	ProtectToday *uint64 `json:"ProtectToday,omitnil" name:"ProtectToday"`

	// 防护主机数
	ProtectHostNum *uint64 `json:"ProtectHostNum,omitnil" name:"ProtectHostNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWebPageGeneralizeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebPageGeneralizeResponseParams `json:"Response"`
}

func (r *DescribeWebPageGeneralizeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebPageGeneralizeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebPageProtectStatRequestParams struct {

}

type DescribeWebPageProtectStatRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeWebPageProtectStatRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebPageProtectStatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebPageProtectStatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebPageProtectStatResponseParams struct {
	// 文件篡改信息
	FileTamperNum []*ProtectStat `json:"FileTamperNum,omitnil" name:"FileTamperNum"`

	// 防护文件分类信息
	ProtectFileType []*ProtectStat `json:"ProtectFileType,omitnil" name:"ProtectFileType"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWebPageProtectStatResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebPageProtectStatResponseParams `json:"Response"`
}

func (r *DescribeWebPageProtectStatResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebPageProtectStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebPageServiceInfoRequestParams struct {

}

type DescribeWebPageServiceInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeWebPageServiceInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebPageServiceInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebPageServiceInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebPageServiceInfoResponseParams struct {
	// 是否已购服务：true-是，false-否
	Status *bool `json:"Status,omitnil" name:"Status"`

	// 已使用授权数
	UsedNum *uint64 `json:"UsedNum,omitnil" name:"UsedNum"`

	// 剩余授权数
	ResidueNum *uint64 `json:"ResidueNum,omitnil" name:"ResidueNum"`

	// 已购授权数
	BuyNum *uint64 `json:"BuyNum,omitnil" name:"BuyNum"`

	// 临近到期数量
	ExpireNum *uint64 `json:"ExpireNum,omitnil" name:"ExpireNum"`

	// 所有授权机器信息
	AllAuthorizedMachines []*ProtectMachineInfo `json:"AllAuthorizedMachines,omitnil" name:"AllAuthorizedMachines"`

	// 临近到期授权机器信息
	ExpireAuthorizedMachines []*ProtectMachine `json:"ExpireAuthorizedMachines,omitnil" name:"ExpireAuthorizedMachines"`

	// 已过期授权数
	ExpiredNum *uint64 `json:"ExpiredNum,omitnil" name:"ExpiredNum"`

	// 防护目录数
	ProtectDirNum *uint64 `json:"ProtectDirNum,omitnil" name:"ProtectDirNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWebPageServiceInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebPageServiceInfoResponseParams `json:"Response"`
}

func (r *DescribeWebPageServiceInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebPageServiceInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyOrderRequestParams struct {
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 授权类型 0 专业版-按量计费, 1专业版-包年包月 , 2 旗舰版-包年包月
	LicenseType *uint64 `json:"LicenseType,omitnil" name:"LicenseType"`
}

type DestroyOrderRequest struct {
	*tchttp.BaseRequest
	
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 授权类型 0 专业版-按量计费, 1专业版-包年包月 , 2 旗舰版-包年包月
	LicenseType *uint64 `json:"LicenseType,omitnil" name:"LicenseType"`
}

func (r *DestroyOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "LicenseType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyOrderResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DestroyOrderResponse struct {
	*tchttp.BaseResponse
	Response *DestroyOrderResponseParams `json:"Response"`
}

func (r *DestroyOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EditBashRulesRequestParams struct {
	// 规则ID（新增时不填）
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 客户端ID数组
	Uuids []*string `json:"Uuids,omitnil" name:"Uuids"`

	// 主机IP
	HostIp *string `json:"HostIp,omitnil" name:"HostIp"`

	// 规则名称，编辑时不可修改规则名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 危险等级(0:无，1: 高危 2:中危 3: 低危)
	Level *uint64 `json:"Level,omitnil" name:"Level"`

	// 正则表达式 ，编辑时不可修改正则表达式，需要对内容QueryEscape后再base64
	Rule *string `json:"Rule,omitnil" name:"Rule"`

	// 是否全局规则(默认否)：1-全局，0-非全局
	IsGlobal *uint64 `json:"IsGlobal,omitnil" name:"IsGlobal"`

	// 0=黑名单， 1=白名单
	White *uint64 `json:"White,omitnil" name:"White"`

	// 事件列表点击“加入白名单”时,需要传EventId 事件的id
	EventId *uint64 `json:"EventId,omitnil" name:"EventId"`

	// 是否处理旧事件为白名单 0=不处理 1=处理
	DealOldEvents *uint64 `json:"DealOldEvents,omitnil" name:"DealOldEvents"`
}

type EditBashRulesRequest struct {
	*tchttp.BaseRequest
	
	// 规则ID（新增时不填）
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 客户端ID数组
	Uuids []*string `json:"Uuids,omitnil" name:"Uuids"`

	// 主机IP
	HostIp *string `json:"HostIp,omitnil" name:"HostIp"`

	// 规则名称，编辑时不可修改规则名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 危险等级(0:无，1: 高危 2:中危 3: 低危)
	Level *uint64 `json:"Level,omitnil" name:"Level"`

	// 正则表达式 ，编辑时不可修改正则表达式，需要对内容QueryEscape后再base64
	Rule *string `json:"Rule,omitnil" name:"Rule"`

	// 是否全局规则(默认否)：1-全局，0-非全局
	IsGlobal *uint64 `json:"IsGlobal,omitnil" name:"IsGlobal"`

	// 0=黑名单， 1=白名单
	White *uint64 `json:"White,omitnil" name:"White"`

	// 事件列表点击“加入白名单”时,需要传EventId 事件的id
	EventId *uint64 `json:"EventId,omitnil" name:"EventId"`

	// 是否处理旧事件为白名单 0=不处理 1=处理
	DealOldEvents *uint64 `json:"DealOldEvents,omitnil" name:"DealOldEvents"`
}

func (r *EditBashRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditBashRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Uuids")
	delete(f, "HostIp")
	delete(f, "Name")
	delete(f, "Level")
	delete(f, "Rule")
	delete(f, "IsGlobal")
	delete(f, "White")
	delete(f, "EventId")
	delete(f, "DealOldEvents")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EditBashRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EditBashRulesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type EditBashRulesResponse struct {
	*tchttp.BaseResponse
	Response *EditBashRulesResponseParams `json:"Response"`
}

func (r *EditBashRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditBashRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EditTagsRequestParams struct {
	// 标签名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 标签ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// Quuid
	Quuids []*string `json:"Quuids,omitnil" name:"Quuids"`
}

type EditTagsRequest struct {
	*tchttp.BaseRequest
	
	// 标签名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 标签ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// Quuid
	Quuids []*string `json:"Quuids,omitnil" name:"Quuids"`
}

func (r *EditTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Id")
	delete(f, "Quuids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EditTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EditTagsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type EditTagsResponse struct {
	*tchttp.BaseResponse
	Response *EditTagsResponseParams `json:"Response"`
}

func (r *EditTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EffectiveMachineInfo struct {
	// 机器名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 机器公网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachinePublicIp *string `json:"MachinePublicIp,omitnil" name:"MachinePublicIp"`

	// 机器内网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachinePrivateIp *string `json:"MachinePrivateIp,omitnil" name:"MachinePrivateIp"`

	// 机器标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineTag []*MachineTag `json:"MachineTag,omitnil" name:"MachineTag"`

	// 机器Quuid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 云镜Uuid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 内核版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	KernelVersion *string `json:"KernelVersion,omitnil" name:"KernelVersion"`

	// 在线状态 OFFLINE，ONLINE
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineStatus *string `json:"MachineStatus,omitnil" name:"MachineStatus"`

	// 授权订单对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	LicenseOrder *LicenseOrder `json:"LicenseOrder,omitnil" name:"LicenseOrder"`

	// 漏洞数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulNum *uint64 `json:"VulNum,omitnil" name:"VulNum"`

	// 云标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CloudTags []*Tags `json:"CloudTags,omitnil" name:"CloudTags"`

	// 机器instance ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
}

type EmergencyResponseInfo struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 主机个数
	HostNum *uint64 `json:"HostNum,omitnil" name:"HostNum"`

	// 服务状态 0未启动，·响应中，2响应完成
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 服务开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 服务结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 报告下载地址
	ReportPath *string `json:"ReportPath,omitnil" name:"ReportPath"`
}

type EmergencyVul struct {
	// 漏洞id
	VulId *uint64 `json:"VulId,omitnil" name:"VulId"`

	// 漏洞级别
	Level *uint64 `json:"Level,omitnil" name:"Level"`

	// 漏洞名称
	VulName *string `json:"VulName,omitnil" name:"VulName"`

	// 发布日期
	PublishDate *string `json:"PublishDate,omitnil" name:"PublishDate"`

	// 漏洞分类
	Category *uint64 `json:"Category,omitnil" name:"Category"`

	// 漏洞状态 0未检测 1有风险 ，2无风险 ，3 检查中展示progress
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 最后扫描时间
	LastScanTime *string `json:"LastScanTime,omitnil" name:"LastScanTime"`

	// 扫描进度
	Progress *uint64 `json:"Progress,omitnil" name:"Progress"`

	// cve编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	CveId *string `json:"CveId,omitnil" name:"CveId"`

	// CVSS评分
	// 注意：此字段可能返回 null，表示取不到有效值。
	CvssScore *float64 `json:"CvssScore,omitnil" name:"CvssScore"`

	// 漏洞标签 多个逗号分割
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels *string `json:"Labels,omitnil" name:"Labels"`

	// 影响机器数
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostCount *uint64 `json:"HostCount,omitnil" name:"HostCount"`

	// 是否支持防御， 0:不支持 1:支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsSupportDefense *uint64 `json:"IsSupportDefense,omitnil" name:"IsSupportDefense"`

	// 已防御的攻击次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefenseAttackCount *uint64 `json:"DefenseAttackCount,omitnil" name:"DefenseAttackCount"`

	// 检测规则 0 - 版本比对, 1 - POC验证
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *uint64 `json:"Method,omitnil" name:"Method"`

	// 攻击热度级别
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackLevel *uint64 `json:"AttackLevel,omitnil" name:"AttackLevel"`

	// 是否有漏洞主机开启漏洞防御
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefenseState *bool `json:"DefenseState,omitnil" name:"DefenseState"`
}

type EventStat struct {
	// 事件数
	EventsNum *uint64 `json:"EventsNum,omitnil" name:"EventsNum"`

	// 受影响的主机数
	MachineAffectNum *uint64 `json:"MachineAffectNum,omitnil" name:"MachineAffectNum"`
}

type ExpertServiceOrderInfo struct {
	// 订单id
	OrderId *uint64 `json:"OrderId,omitnil" name:"OrderId"`

	// 订单类型 1应急 2 旗舰重保 3 安全管家
	InquireType *uint64 `json:"InquireType,omitnil" name:"InquireType"`

	// 服务数量
	InquireNum *uint64 `json:"InquireNum,omitnil" name:"InquireNum"`

	// 服务开始时间
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 服务结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 服务时长几个月
	ServiceTime *uint64 `json:"ServiceTime,omitnil" name:"ServiceTime"`

	// 订单状态 0 未启动 1 服务中 2已过期 3完成，4退费销毁
	Status *uint64 `json:"Status,omitnil" name:"Status"`
}

// Predefined struct for user
type ExportAssetCoreModuleListRequestParams struct {
	// 服务器Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 服务器Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>Name- string - 是否必填：否 - 包名</li>
	// <li>User- string - 是否必填：否 - 用户</li>
	Filters []*AssetFilters `json:"Filters,omitnil" name:"Filters"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序依据[FirstTime|Size|ProcessCount|ModuleCount]
	By *string `json:"By,omitnil" name:"By"`
}

type ExportAssetCoreModuleListRequest struct {
	*tchttp.BaseRequest
	
	// 服务器Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 服务器Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>Name- string - 是否必填：否 - 包名</li>
	// <li>User- string - 是否必填：否 - 用户</li>
	Filters []*AssetFilters `json:"Filters,omitnil" name:"Filters"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序依据[FirstTime|Size|ProcessCount|ModuleCount]
	By *string `json:"By,omitnil" name:"By"`
}

func (r *ExportAssetCoreModuleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportAssetCoreModuleListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "Quuid")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportAssetCoreModuleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportAssetCoreModuleListResponseParams struct {
	// 异步下载任务ID，需要配合ExportTasks接口使用
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportAssetCoreModuleListResponse struct {
	*tchttp.BaseResponse
	Response *ExportAssetCoreModuleListResponseParams `json:"Response"`
}

func (r *ExportAssetCoreModuleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportAssetCoreModuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportAssetWebServiceInfoListRequestParams struct {
	// 查询指定Quuid主机的信息
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>User- string - 是否必填：否 - 运行用户</li>
	// <li>Name- string - 是否必填：否 - Web服务名：
	// 1:Tomcat
	// 2:Apache
	// 3:Nginx
	// 4:WebLogic
	// 5:Websphere
	// 6:JBoss
	// 7:WildFly
	// 8:Jetty
	// 9:IHS
	// 10:Tengine</li>
	// <li>OsType- string - 是否必填：否 - Windows/linux</li>
	Filters []*AssetFilters `json:"Filters,omitnil" name:"Filters"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序：[FirstTime|ProcessCount]
	By *string `json:"By,omitnil" name:"By"`
}

type ExportAssetWebServiceInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 查询指定Quuid主机的信息
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 过滤条件。
	// <li>User- string - 是否必填：否 - 运行用户</li>
	// <li>Name- string - 是否必填：否 - Web服务名：
	// 1:Tomcat
	// 2:Apache
	// 3:Nginx
	// 4:WebLogic
	// 5:Websphere
	// 6:JBoss
	// 7:WildFly
	// 8:Jetty
	// 9:IHS
	// 10:Tengine</li>
	// <li>OsType- string - 是否必填：否 - Windows/linux</li>
	Filters []*AssetFilters `json:"Filters,omitnil" name:"Filters"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 可选排序：[FirstTime|ProcessCount]
	By *string `json:"By,omitnil" name:"By"`
}

func (r *ExportAssetWebServiceInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportAssetWebServiceInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Quuid")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportAssetWebServiceInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportAssetWebServiceInfoListResponseParams struct {
	// 异步下载任务ID，需要配合ExportTasks接口使用
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportAssetWebServiceInfoListResponse struct {
	*tchttp.BaseResponse
	Response *ExportAssetWebServiceInfoListResponseParams `json:"Response"`
}

func (r *ExportAssetWebServiceInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportAssetWebServiceInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportAttackLogsRequestParams struct {
	// 过滤条件。
	// <li>HttpMethod - String - 是否必填：否 - 攻击方法(POST|GET)</li>
	// <li>DateRange - String - 是否必填：否 - 时间范围(存储最近3个月的数据)，如最近一个月["2019-11-17", "2019-12-17"]</li>
	// <li>VulType - String 威胁类型 - 是否必填: 否</li>
	// <li>SrcIp - String 攻击源IP - 是否必填: 否</li>
	// <li>DstIp - String 攻击目标IP - 是否必填: 否</li>
	// <li>SrcPort - String 攻击源端口 - 是否必填: 否</li>
	// <li>DstPort - String 攻击目标端口 - 是否必填: 否</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 主机安全客户端ID
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 云主机机器ID
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`
}

type ExportAttackLogsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>HttpMethod - String - 是否必填：否 - 攻击方法(POST|GET)</li>
	// <li>DateRange - String - 是否必填：否 - 时间范围(存储最近3个月的数据)，如最近一个月["2019-11-17", "2019-12-17"]</li>
	// <li>VulType - String 威胁类型 - 是否必填: 否</li>
	// <li>SrcIp - String 攻击源IP - 是否必填: 否</li>
	// <li>DstIp - String 攻击目标IP - 是否必填: 否</li>
	// <li>SrcPort - String 攻击源端口 - 是否必填: 否</li>
	// <li>DstPort - String 攻击目标端口 - 是否必填: 否</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 主机安全客户端ID
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 云主机机器ID
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`
}

func (r *ExportAttackLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportAttackLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Uuid")
	delete(f, "Quuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportAttackLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportAttackLogsResponseParams struct {
	// 已废弃
	DownloadUrl *string `json:"DownloadUrl,omitnil" name:"DownloadUrl"`

	// 任务ID,需要到接口“异步导出任务”ExportTasks获取DownloadUrl下载地址
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportAttackLogsResponse struct {
	*tchttp.BaseResponse
	Response *ExportAttackLogsResponseParams `json:"Response"`
}

func (r *ExportAttackLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportAttackLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportBaselineEffectHostListRequestParams struct {
	// 基线id
	BaselineId *uint64 `json:"BaselineId,omitnil" name:"BaselineId"`

	// 筛选条件
	// <li>AliasName- String- 主机别名</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 策略id
	StrategyId *uint64 `json:"StrategyId,omitnil" name:"StrategyId"`

	// 主机uuid数组
	UuidList []*string `json:"UuidList,omitnil" name:"UuidList"`

	// 基线名称
	BaselineName *string `json:"BaselineName,omitnil" name:"BaselineName"`
}

type ExportBaselineEffectHostListRequest struct {
	*tchttp.BaseRequest
	
	// 基线id
	BaselineId *uint64 `json:"BaselineId,omitnil" name:"BaselineId"`

	// 筛选条件
	// <li>AliasName- String- 主机别名</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 策略id
	StrategyId *uint64 `json:"StrategyId,omitnil" name:"StrategyId"`

	// 主机uuid数组
	UuidList []*string `json:"UuidList,omitnil" name:"UuidList"`

	// 基线名称
	BaselineName *string `json:"BaselineName,omitnil" name:"BaselineName"`
}

func (r *ExportBaselineEffectHostListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBaselineEffectHostListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BaselineId")
	delete(f, "Filters")
	delete(f, "StrategyId")
	delete(f, "UuidList")
	delete(f, "BaselineName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportBaselineEffectHostListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportBaselineEffectHostListResponseParams struct {
	// 该参数已废弃
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadUrl *string `json:"DownloadUrl,omitnil" name:"DownloadUrl"`

	// 任务ID,需要到接口“异步导出任务”ExportTasks获取DownloadUrl下载地址
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportBaselineEffectHostListResponse struct {
	*tchttp.BaseResponse
	Response *ExportBaselineEffectHostListResponseParams `json:"Response"`
}

func (r *ExportBaselineEffectHostListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBaselineEffectHostListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportBaselineFixListRequestParams struct {
	// <li>ItemName - String - 是否必填：否 - 项名称</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 0:过滤的结果导出；1:全部导出
	ExportAll *int64 `json:"ExportAll,omitnil" name:"ExportAll"`
}

type ExportBaselineFixListRequest struct {
	*tchttp.BaseRequest
	
	// <li>ItemName - String - 是否必填：否 - 项名称</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 0:过滤的结果导出；1:全部导出
	ExportAll *int64 `json:"ExportAll,omitnil" name:"ExportAll"`
}

func (r *ExportBaselineFixListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBaselineFixListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "ExportAll")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportBaselineFixListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportBaselineFixListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportBaselineFixListResponse struct {
	*tchttp.BaseResponse
	Response *ExportBaselineFixListResponseParams `json:"Response"`
}

func (r *ExportBaselineFixListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBaselineFixListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportBaselineHostDetectListRequestParams struct {
	// <li>HostTag - string - 是否必填：否 - 主机标签</i>
	// <li>ItemId - int64 - 是否必填：否 - 项Id</i>
	// <li>RuleId - int64 - 是否必填：否 - 规则Id</li>
	// <li>IsPassed - int - 是否必填：否 - 是否通过</li>
	// <li>RiskTier - int - 是否必填：否 - 风险等级</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 0:过滤的结果导出；1:全部导出
	ExportAll *int64 `json:"ExportAll,omitnil" name:"ExportAll"`
}

type ExportBaselineHostDetectListRequest struct {
	*tchttp.BaseRequest
	
	// <li>HostTag - string - 是否必填：否 - 主机标签</i>
	// <li>ItemId - int64 - 是否必填：否 - 项Id</i>
	// <li>RuleId - int64 - 是否必填：否 - 规则Id</li>
	// <li>IsPassed - int - 是否必填：否 - 是否通过</li>
	// <li>RiskTier - int - 是否必填：否 - 风险等级</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 0:过滤的结果导出；1:全部导出
	ExportAll *int64 `json:"ExportAll,omitnil" name:"ExportAll"`
}

func (r *ExportBaselineHostDetectListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBaselineHostDetectListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "ExportAll")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportBaselineHostDetectListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportBaselineHostDetectListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportBaselineHostDetectListResponse struct {
	*tchttp.BaseResponse
	Response *ExportBaselineHostDetectListResponseParams `json:"Response"`
}

func (r *ExportBaselineHostDetectListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBaselineHostDetectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportBaselineItemDetectListRequestParams struct {
	// <li>HostId - string - 是否必填：否 - 主机Id</i>
	// <li>RuleId - int64 - 是否必填：否 - 规则Id</i>
	// <li>IsPassed - int - 是否必填：否 - 是否通过</li>
	// <li>RiskTier - int - 是否必填：否 - 风险等级</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 0:过滤的结果导出；1:全部导出
	ExportAll *int64 `json:"ExportAll,omitnil" name:"ExportAll"`
}

type ExportBaselineItemDetectListRequest struct {
	*tchttp.BaseRequest
	
	// <li>HostId - string - 是否必填：否 - 主机Id</i>
	// <li>RuleId - int64 - 是否必填：否 - 规则Id</i>
	// <li>IsPassed - int - 是否必填：否 - 是否通过</li>
	// <li>RiskTier - int - 是否必填：否 - 风险等级</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 0:过滤的结果导出；1:全部导出
	ExportAll *int64 `json:"ExportAll,omitnil" name:"ExportAll"`
}

func (r *ExportBaselineItemDetectListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBaselineItemDetectListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "ExportAll")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportBaselineItemDetectListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportBaselineItemDetectListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportBaselineItemDetectListResponse struct {
	*tchttp.BaseResponse
	Response *ExportBaselineItemDetectListResponseParams `json:"Response"`
}

func (r *ExportBaselineItemDetectListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBaselineItemDetectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportBaselineItemListRequestParams struct {
	// <li>PolicyId - int64 - 是否必填：否 - 策略Id</li>
	// <li>RuleId - int64 - 是否必填：否 - 规则Id</li>
	// <li>HostId - string - 是否必填：否 - 主机Id</li>
	// <li>HostName - string - 是否必填：否 - 主机名</li>
	// <li>HostIp - string - 是否必填：否 - 主机IP</li>
	// <li>ItemId - String - 是否必填：否 - 检测项Id</li>
	// <li>ItemName - String - 是否必填：否 - 项名称</li>
	// <li>DetectStatus - int - 是否必填：否 - 检测状态[0:未通过|3:通过|5:检测中]</li>
	// <li>Level - int - 是否必填：否 - 风险等级</li>
	// <li>StartTime - string - 是否必填：否 - 开时时间</li>
	// <li>EndTime - string - 是否必填：否 - 结束时间</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 0:过滤的结果导出；1:全部导出
	ExportAll *int64 `json:"ExportAll,omitnil" name:"ExportAll"`
}

type ExportBaselineItemListRequest struct {
	*tchttp.BaseRequest
	
	// <li>PolicyId - int64 - 是否必填：否 - 策略Id</li>
	// <li>RuleId - int64 - 是否必填：否 - 规则Id</li>
	// <li>HostId - string - 是否必填：否 - 主机Id</li>
	// <li>HostName - string - 是否必填：否 - 主机名</li>
	// <li>HostIp - string - 是否必填：否 - 主机IP</li>
	// <li>ItemId - String - 是否必填：否 - 检测项Id</li>
	// <li>ItemName - String - 是否必填：否 - 项名称</li>
	// <li>DetectStatus - int - 是否必填：否 - 检测状态[0:未通过|3:通过|5:检测中]</li>
	// <li>Level - int - 是否必填：否 - 风险等级</li>
	// <li>StartTime - string - 是否必填：否 - 开时时间</li>
	// <li>EndTime - string - 是否必填：否 - 结束时间</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 0:过滤的结果导出；1:全部导出
	ExportAll *int64 `json:"ExportAll,omitnil" name:"ExportAll"`
}

func (r *ExportBaselineItemListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBaselineItemListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "ExportAll")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportBaselineItemListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportBaselineItemListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportBaselineItemListResponse struct {
	*tchttp.BaseResponse
	Response *ExportBaselineItemListResponseParams `json:"Response"`
}

func (r *ExportBaselineItemListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBaselineItemListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportBaselineListRequestParams struct {
	// 过滤条件：
	// <li>StrategyId- Uint64 - 基线策略id</li>
	// <li>Status - Uint64 - 事件状态：0-未通过，1-忽略，3-通过，5-检测中</li>
	// <li>BaselineName  - String - 基线名称</li>
	// <li>AliasName- String - 服务器名称/服务器ip</li>
	// <li>Uuid- String - 主机uuid</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 已废弃
	IfDetail *uint64 `json:"IfDetail,omitnil" name:"IfDetail"`
}

type ExportBaselineListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件：
	// <li>StrategyId- Uint64 - 基线策略id</li>
	// <li>Status - Uint64 - 事件状态：0-未通过，1-忽略，3-通过，5-检测中</li>
	// <li>BaselineName  - String - 基线名称</li>
	// <li>AliasName- String - 服务器名称/服务器ip</li>
	// <li>Uuid- String - 主机uuid</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 已废弃
	IfDetail *uint64 `json:"IfDetail,omitnil" name:"IfDetail"`
}

func (r *ExportBaselineListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBaselineListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "IfDetail")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportBaselineListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportBaselineListResponseParams struct {
	// 导出文件下载地址（已弃用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadUrl *string `json:"DownloadUrl,omitnil" name:"DownloadUrl"`

	// 任务ID,需要到接口“异步导出任务”ExportTasks获取DownloadUrl下载地址
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportBaselineListResponse struct {
	*tchttp.BaseResponse
	Response *ExportBaselineListResponseParams `json:"Response"`
}

func (r *ExportBaselineListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBaselineListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportBaselineRuleDetectListRequestParams struct {
	// <li>RuleName - string - 是否必填：否 - 规则名称</i>
	// <li>IsPassed - int - 是否必填：否 - 是否通过</li>
	// <li>RiskTier - int - 是否必填：否 - 风险等级</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 0:过滤的结果导出；1:全部导出
	ExportAll *int64 `json:"ExportAll,omitnil" name:"ExportAll"`
}

type ExportBaselineRuleDetectListRequest struct {
	*tchttp.BaseRequest
	
	// <li>RuleName - string - 是否必填：否 - 规则名称</i>
	// <li>IsPassed - int - 是否必填：否 - 是否通过</li>
	// <li>RiskTier - int - 是否必填：否 - 风险等级</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 0:过滤的结果导出；1:全部导出
	ExportAll *int64 `json:"ExportAll,omitnil" name:"ExportAll"`
}

func (r *ExportBaselineRuleDetectListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBaselineRuleDetectListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "ExportAll")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportBaselineRuleDetectListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportBaselineRuleDetectListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportBaselineRuleDetectListResponse struct {
	*tchttp.BaseResponse
	Response *ExportBaselineRuleDetectListResponseParams `json:"Response"`
}

func (r *ExportBaselineRuleDetectListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBaselineRuleDetectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportBaselineWeakPasswordListRequestParams struct {
	// <li>WeakPassword - string - 是否必填：否 - 弱口令</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 0:过滤的结果导出；1:全部导出
	ExportAll *int64 `json:"ExportAll,omitnil" name:"ExportAll"`
}

type ExportBaselineWeakPasswordListRequest struct {
	*tchttp.BaseRequest
	
	// <li>WeakPassword - string - 是否必填：否 - 弱口令</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 0:过滤的结果导出；1:全部导出
	ExportAll *int64 `json:"ExportAll,omitnil" name:"ExportAll"`
}

func (r *ExportBaselineWeakPasswordListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBaselineWeakPasswordListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "ExportAll")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportBaselineWeakPasswordListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportBaselineWeakPasswordListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportBaselineWeakPasswordListResponse struct {
	*tchttp.BaseResponse
	Response *ExportBaselineWeakPasswordListResponseParams `json:"Response"`
}

func (r *ExportBaselineWeakPasswordListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBaselineWeakPasswordListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportBashEventsRequestParams struct {
	// 过滤参数
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`
}

type ExportBashEventsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤参数
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`
}

func (r *ExportBashEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBashEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportBashEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportBashEventsResponseParams struct {
	// 导出文件下载链接地址。
	DownloadUrl *string `json:"DownloadUrl,omitnil" name:"DownloadUrl"`

	// 任务ID,需要到接口“异步导出任务”ExportTasks获取DownloadUrl下载地址
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportBashEventsResponse struct {
	*tchttp.BaseResponse
	Response *ExportBashEventsResponseParams `json:"Response"`
}

func (r *ExportBashEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBashEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportBruteAttacksRequestParams struct {
	// 过滤参数
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`
}

type ExportBruteAttacksRequest struct {
	*tchttp.BaseRequest
	
	// 过滤参数
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`
}

func (r *ExportBruteAttacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBruteAttacksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportBruteAttacksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportBruteAttacksResponseParams struct {
	// 该参数已废弃
	DownloadUrl *string `json:"DownloadUrl,omitnil" name:"DownloadUrl"`

	// 任务ID,需要到接口“异步导出任务”ExportTasks获取DownloadUrl下载地址
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportBruteAttacksResponse struct {
	*tchttp.BaseResponse
	Response *ExportBruteAttacksResponseParams `json:"Response"`
}

func (r *ExportBruteAttacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBruteAttacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportIgnoreBaselineRuleRequestParams struct {
	// 检测项名称
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`
}

type ExportIgnoreBaselineRuleRequest struct {
	*tchttp.BaseRequest
	
	// 检测项名称
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`
}

func (r *ExportIgnoreBaselineRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportIgnoreBaselineRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportIgnoreBaselineRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportIgnoreBaselineRuleResponseParams struct {
	// 该参数已废弃
	DownloadUrl *string `json:"DownloadUrl,omitnil" name:"DownloadUrl"`

	// 任务ID,需要到接口“异步导出任务”ExportTasks获取DownloadUrl下载地址
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportIgnoreBaselineRuleResponse struct {
	*tchttp.BaseResponse
	Response *ExportIgnoreBaselineRuleResponseParams `json:"Response"`
}

func (r *ExportIgnoreBaselineRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportIgnoreBaselineRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportIgnoreRuleEffectHostListRequestParams struct {
	// 检测项id
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`

	// 过滤条件。
	// <li>AliasName- String- 主机别名</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`
}

type ExportIgnoreRuleEffectHostListRequest struct {
	*tchttp.BaseRequest
	
	// 检测项id
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`

	// 过滤条件。
	// <li>AliasName- String- 主机别名</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`
}

func (r *ExportIgnoreRuleEffectHostListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportIgnoreRuleEffectHostListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportIgnoreRuleEffectHostListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportIgnoreRuleEffectHostListResponseParams struct {
	// 该参数已废弃
	DownloadUrl *string `json:"DownloadUrl,omitnil" name:"DownloadUrl"`

	// 任务ID,需要到接口“异步导出任务”ExportTasks获取DownloadUrl下载地址
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportIgnoreRuleEffectHostListResponse struct {
	*tchttp.BaseResponse
	Response *ExportIgnoreRuleEffectHostListResponseParams `json:"Response"`
}

func (r *ExportIgnoreRuleEffectHostListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportIgnoreRuleEffectHostListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportLicenseDetailRequestParams struct {
	// 多个条件筛选时 LicenseStatus,DeadlineStatus,ResourceId,Keywords 取交集
	// <li> LicenseType  授权类型, 0 专业版-按量计费, 1专业版-包年包月 , 2 旗舰版-包年包月</li>
	// <li>ResourceId 资源ID</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 是否导出全部授权详情
	IsHistory *bool `json:"IsHistory,omitnil" name:"IsHistory"`

	// 标签筛选,平台标签能力,这里传入 标签键,标签值作为一个对象
	Tags []*Tags `json:"Tags,omitnil" name:"Tags"`

	// 导出月份, 该参数仅在IsHistory 时可选.
	ExportMonth *string `json:"ExportMonth,omitnil" name:"ExportMonth"`
}

type ExportLicenseDetailRequest struct {
	*tchttp.BaseRequest
	
	// 多个条件筛选时 LicenseStatus,DeadlineStatus,ResourceId,Keywords 取交集
	// <li> LicenseType  授权类型, 0 专业版-按量计费, 1专业版-包年包月 , 2 旗舰版-包年包月</li>
	// <li>ResourceId 资源ID</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 是否导出全部授权详情
	IsHistory *bool `json:"IsHistory,omitnil" name:"IsHistory"`

	// 标签筛选,平台标签能力,这里传入 标签键,标签值作为一个对象
	Tags []*Tags `json:"Tags,omitnil" name:"Tags"`

	// 导出月份, 该参数仅在IsHistory 时可选.
	ExportMonth *string `json:"ExportMonth,omitnil" name:"ExportMonth"`
}

func (r *ExportLicenseDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportLicenseDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "IsHistory")
	delete(f, "Tags")
	delete(f, "ExportMonth")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportLicenseDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportLicenseDetailResponseParams struct {
	// 下载地址,该字段废弃
	DownloadUrl *string `json:"DownloadUrl,omitnil" name:"DownloadUrl"`

	// 任务ID,需要到接口“异步导出任务”ExportTasks获取DownloadUrl下载地址
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportLicenseDetailResponse struct {
	*tchttp.BaseResponse
	Response *ExportLicenseDetailResponseParams `json:"Response"`
}

func (r *ExportLicenseDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportLicenseDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportMaliciousRequestsRequestParams struct {
	// 过滤参数
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`
}

type ExportMaliciousRequestsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤参数
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`
}

func (r *ExportMaliciousRequestsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportMaliciousRequestsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportMaliciousRequestsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportMaliciousRequestsResponseParams struct {
	// 该参数已废弃
	DownloadUrl *string `json:"DownloadUrl,omitnil" name:"DownloadUrl"`

	// 任务ID,需要到接口“异步导出任务”ExportTasks获取DownloadUrl下载地址
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportMaliciousRequestsResponse struct {
	*tchttp.BaseResponse
	Response *ExportMaliciousRequestsResponseParams `json:"Response"`
}

func (r *ExportMaliciousRequestsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportMaliciousRequestsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportMalwaresRequestParams struct {
	// 限制条数,默认10
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量 默认0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤参数。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>FilePath - String - 是否必填：否 - 路径筛选</li>
	// <li>VirusName - String - 是否必填：否 - 描述筛选</li>
	// <li>CreateBeginTime - String - 是否必填：否 - 创建时间筛选-开始时间</li>
	// <li>CreateEndTime - String - 是否必填：否 - 创建时间筛选-结束时间</li>
	// <li>Status - String - 是否必填：否 - 状态筛选</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 排序值 CreateTime
	By *string `json:"By,omitnil" name:"By"`

	// 排序 方式 ，ASC，DESC
	Order *string `json:"Order,omitnil" name:"Order"`
}

type ExportMalwaresRequest struct {
	*tchttp.BaseRequest
	
	// 限制条数,默认10
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量 默认0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤参数。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>FilePath - String - 是否必填：否 - 路径筛选</li>
	// <li>VirusName - String - 是否必填：否 - 描述筛选</li>
	// <li>CreateBeginTime - String - 是否必填：否 - 创建时间筛选-开始时间</li>
	// <li>CreateEndTime - String - 是否必填：否 - 创建时间筛选-结束时间</li>
	// <li>Status - String - 是否必填：否 - 状态筛选</li>
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 排序值 CreateTime
	By *string `json:"By,omitnil" name:"By"`

	// 排序 方式 ，ASC，DESC
	Order *string `json:"Order,omitnil" name:"Order"`
}

func (r *ExportMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportMalwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportMalwaresResponseParams struct {
	// 该参数已废弃
	DownloadUrl *string `json:"DownloadUrl,omitnil" name:"DownloadUrl"`

	// 任务ID,需要到接口“异步导出任务”ExportTasks获取DownloadUrl下载地址
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *ExportMalwaresResponseParams `json:"Response"`
}

func (r *ExportMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportNonlocalLoginPlacesRequestParams struct {
	// <li>Status - int - 是否必填：否 - 状态筛选1:正常登录；2：异地登录</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type ExportNonlocalLoginPlacesRequest struct {
	*tchttp.BaseRequest
	
	// <li>Status - int - 是否必填：否 - 状态筛选1:正常登录；2：异地登录</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *ExportNonlocalLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportNonlocalLoginPlacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportNonlocalLoginPlacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportNonlocalLoginPlacesResponseParams struct {
	// 该参数已废弃
	DownloadUrl *string `json:"DownloadUrl,omitnil" name:"DownloadUrl"`

	// 任务ID,需要到接口“异步导出任务”ExportTasks获取DownloadUrl下载地址
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportNonlocalLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *ExportNonlocalLoginPlacesResponseParams `json:"Response"`
}

func (r *ExportNonlocalLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportNonlocalLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportPrivilegeEventsRequestParams struct {
	// 过滤参数
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`
}

type ExportPrivilegeEventsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤参数
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`
}

func (r *ExportPrivilegeEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportPrivilegeEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportPrivilegeEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportPrivilegeEventsResponseParams struct {
	// 该参数已废弃
	DownloadUrl *string `json:"DownloadUrl,omitnil" name:"DownloadUrl"`

	// 任务ID,需要到接口“异步导出任务”ExportTasks获取DownloadUrl下载地址
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportPrivilegeEventsResponse struct {
	*tchttp.BaseResponse
	Response *ExportPrivilegeEventsResponseParams `json:"Response"`
}

func (r *ExportPrivilegeEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportPrivilegeEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportProtectDirListRequestParams struct {
	// DirName 网站名称
	// DirPath 网站防护目录地址
	Filters []*AssetFilters `json:"Filters,omitnil" name:"Filters"`

	// asc：升序/desc：降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段
	By *string `json:"By,omitnil" name:"By"`
}

type ExportProtectDirListRequest struct {
	*tchttp.BaseRequest
	
	// DirName 网站名称
	// DirPath 网站防护目录地址
	Filters []*AssetFilters `json:"Filters,omitnil" name:"Filters"`

	// asc：升序/desc：降序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段
	By *string `json:"By,omitnil" name:"By"`
}

func (r *ExportProtectDirListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportProtectDirListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportProtectDirListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportProtectDirListResponseParams struct {
	// 任务ID,需要到接口“异步导出任务”ExportTasks获取DownloadUrl下载地址
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportProtectDirListResponse struct {
	*tchttp.BaseResponse
	Response *ExportProtectDirListResponseParams `json:"Response"`
}

func (r *ExportProtectDirListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportProtectDirListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportReverseShellEventsRequestParams struct {
	// 过滤参数
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`
}

type ExportReverseShellEventsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤参数
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`
}

func (r *ExportReverseShellEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportReverseShellEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportReverseShellEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportReverseShellEventsResponseParams struct {
	// 该参数已废弃
	DownloadUrl *string `json:"DownloadUrl,omitnil" name:"DownloadUrl"`

	// 任务ID,需要到接口“异步导出任务”ExportTasks获取DownloadUrl下载地址
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportReverseShellEventsResponse struct {
	*tchttp.BaseResponse
	Response *ExportReverseShellEventsResponseParams `json:"Response"`
}

func (r *ExportReverseShellEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportReverseShellEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportScanTaskDetailsRequestParams struct {
	// 本次检测的任务id（不同于出参的导出本次检测Excel的任务Id）
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// 模块类型，当前提供：Malware 木马 , Vul 漏洞 , Baseline 基线
	ModuleType *string `json:"ModuleType,omitnil" name:"ModuleType"`

	// 过滤参数：ipOrAlias（服务器名/ip）
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`
}

type ExportScanTaskDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 本次检测的任务id（不同于出参的导出本次检测Excel的任务Id）
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// 模块类型，当前提供：Malware 木马 , Vul 漏洞 , Baseline 基线
	ModuleType *string `json:"ModuleType,omitnil" name:"ModuleType"`

	// 过滤参数：ipOrAlias（服务器名/ip）
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`
}

func (r *ExportScanTaskDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportScanTaskDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ModuleType")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportScanTaskDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportScanTaskDetailsResponseParams struct {
	// 任务ID,需要到接口“异步导出任务”ExportTasks获取DownloadUrl下载地址(不同于入参的本次检测任务id)
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportScanTaskDetailsResponse struct {
	*tchttp.BaseResponse
	Response *ExportScanTaskDetailsResponseParams `json:"Response"`
}

func (r *ExportScanTaskDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportScanTaskDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportSecurityTrendsRequestParams struct {
	// 开始时间。
	BeginDate *string `json:"BeginDate,omitnil" name:"BeginDate"`

	// 结束时间。
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`
}

type ExportSecurityTrendsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	BeginDate *string `json:"BeginDate,omitnil" name:"BeginDate"`

	// 结束时间。
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`
}

func (r *ExportSecurityTrendsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportSecurityTrendsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportSecurityTrendsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportSecurityTrendsResponseParams struct {
	// 导出文件下载链接地址。
	DownloadUrl *string `json:"DownloadUrl,omitnil" name:"DownloadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportSecurityTrendsResponse struct {
	*tchttp.BaseResponse
	Response *ExportSecurityTrendsResponseParams `json:"Response"`
}

func (r *ExportSecurityTrendsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportSecurityTrendsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportTasksRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

type ExportTasksRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

func (r *ExportTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportTasksResponseParams struct {
	// PENDING：正在生成下载链接，FINISHED：下载链接已生成，ERROR：网络异常等异常情况
	Status *string `json:"Status,omitnil" name:"Status"`

	// 下载链接
	DownloadUrl *string `json:"DownloadUrl,omitnil" name:"DownloadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportTasksResponse struct {
	*tchttp.BaseResponse
	Response *ExportTasksResponseParams `json:"Response"`
}

func (r *ExportTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportVulDetectionExcelRequestParams struct {
	// 本次漏洞检测任务id（不同于出参的导出本次漏洞检测Excel的任务Id）
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`
}

type ExportVulDetectionExcelRequest struct {
	*tchttp.BaseRequest
	
	// 本次漏洞检测任务id（不同于出参的导出本次漏洞检测Excel的任务Id）
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`
}

func (r *ExportVulDetectionExcelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportVulDetectionExcelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportVulDetectionExcelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportVulDetectionExcelResponseParams struct {
	// 该参数已废弃
	DownloadUrl *string `json:"DownloadUrl,omitnil" name:"DownloadUrl"`

	// 任务ID,需要到接口“异步导出任务”ExportTasks获取DownloadUrl下载地址（不同于入参的本次漏洞检测任务id）
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportVulDetectionExcelResponse struct {
	*tchttp.BaseResponse
	Response *ExportVulDetectionExcelResponseParams `json:"Response"`
}

func (r *ExportVulDetectionExcelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportVulDetectionExcelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportVulDetectionReportRequestParams struct {
	// 漏洞扫描任务id（不同于出参的导出检测报告的任务Id）
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// 过滤参数
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type ExportVulDetectionReportRequest struct {
	*tchttp.BaseRequest
	
	// 漏洞扫描任务id（不同于出参的导出检测报告的任务Id）
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// 过滤参数
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *ExportVulDetectionReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportVulDetectionReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportVulDetectionReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportVulDetectionReportResponseParams struct {
	// 导出文件下载链接地址
	DownloadUrl *string `json:"DownloadUrl,omitnil" name:"DownloadUrl"`

	// 导出检测报告的任务Id（不同于入参的漏洞扫描任务id）
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportVulDetectionReportResponse struct {
	*tchttp.BaseResponse
	Response *ExportVulDetectionReportResponseParams `json:"Response"`
}

func (r *ExportVulDetectionReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportVulDetectionReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportVulEffectHostListRequestParams struct {
	// 漏洞id
	VulId *uint64 `json:"VulId,omitnil" name:"VulId"`

	// 过滤条件。
	// <li>AliasName - String - 主机名筛选</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type ExportVulEffectHostListRequest struct {
	*tchttp.BaseRequest
	
	// 漏洞id
	VulId *uint64 `json:"VulId,omitnil" name:"VulId"`

	// 过滤条件。
	// <li>AliasName - String - 主机名筛选</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *ExportVulEffectHostListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportVulEffectHostListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VulId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportVulEffectHostListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportVulEffectHostListResponseParams struct {
	// 已废弃
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadUrl *string `json:"DownloadUrl,omitnil" name:"DownloadUrl"`

	// 导出任务Id , 可通过ExportTasks 接口下载
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportVulEffectHostListResponse struct {
	*tchttp.BaseResponse
	Response *ExportVulEffectHostListResponseParams `json:"Response"`
}

func (r *ExportVulEffectHostListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportVulEffectHostListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportVulListRequestParams struct {
	// 过滤条件。
	// <li>VulCategory - int - 是否必填：否 - 漏洞分类筛选1: web-cms漏洞 2:应用漏洞  4: Linux软件漏洞 5: Windows系统漏洞</li>
	// <li>IfEmergency - String - 是否必填：否 - 是否为应急漏洞，查询应急漏洞传:yes</li>
	// <li>Status - String - 是否必填：是 - 漏洞状态筛选，0: 待处理 1:忽略  3:已修复  5:检测中， 控制台仅处理0,1,3,5四种状态</li>
	// <li>Level - String - 是否必填：否 - 漏洞等级筛选 1:低 2:中 3:高 4:提示</li>
	// <li>VulName- String - 是否必填：否 - 漏洞名称搜索</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 是否导出详情,1是 0不是
	IfDetail *uint64 `json:"IfDetail,omitnil" name:"IfDetail"`
}

type ExportVulListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>VulCategory - int - 是否必填：否 - 漏洞分类筛选1: web-cms漏洞 2:应用漏洞  4: Linux软件漏洞 5: Windows系统漏洞</li>
	// <li>IfEmergency - String - 是否必填：否 - 是否为应急漏洞，查询应急漏洞传:yes</li>
	// <li>Status - String - 是否必填：是 - 漏洞状态筛选，0: 待处理 1:忽略  3:已修复  5:检测中， 控制台仅处理0,1,3,5四种状态</li>
	// <li>Level - String - 是否必填：否 - 漏洞等级筛选 1:低 2:中 3:高 4:提示</li>
	// <li>VulName- String - 是否必填：否 - 漏洞名称搜索</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 是否导出详情,1是 0不是
	IfDetail *uint64 `json:"IfDetail,omitnil" name:"IfDetail"`
}

func (r *ExportVulListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportVulListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "IfDetail")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportVulListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportVulListResponseParams struct {
	// 导出的文件下载url（已弃用！）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadUrl *string `json:"DownloadUrl,omitnil" name:"DownloadUrl"`

	// 导出文件Id 可通过ExportTasks接口下载
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportVulListResponse struct {
	*tchttp.BaseResponse
	Response *ExportVulListResponseParams `json:"Response"`
}

func (r *ExportVulListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportVulListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportWebPageEventListRequestParams struct {
	// 过滤条件
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>EventType - String - 是否必填：否 - 事件类型</li>
	// <li>EventStatus - String - 是否必填：否 - 事件状态</li>
	Filters []*AssetFilters `json:"Filters,omitnil" name:"Filters"`

	// 排序方式：CreateTime 或 RestoreTime，默认为CreateTime
	By *string `json:"By,omitnil" name:"By"`

	// 排序方式，0降序，1升序，默认为0
	Order *uint64 `json:"Order,omitnil" name:"Order"`
}

type ExportWebPageEventListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>EventType - String - 是否必填：否 - 事件类型</li>
	// <li>EventStatus - String - 是否必填：否 - 事件状态</li>
	Filters []*AssetFilters `json:"Filters,omitnil" name:"Filters"`

	// 排序方式：CreateTime 或 RestoreTime，默认为CreateTime
	By *string `json:"By,omitnil" name:"By"`

	// 排序方式，0降序，1升序，默认为0
	Order *uint64 `json:"Order,omitnil" name:"Order"`
}

func (r *ExportWebPageEventListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportWebPageEventListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportWebPageEventListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportWebPageEventListResponseParams struct {
	// 任务id 可通过 ExportTasks接口下载
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportWebPageEventListResponse struct {
	*tchttp.BaseResponse
	Response *ExportWebPageEventListResponseParams `json:"Response"`
}

func (r *ExportWebPageEventListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportWebPageEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileTamperEvent struct {
	// 机器名称
	HostName *string `json:"HostName,omitnil" name:"HostName"`

	// 机器IP
	HostIp *string `json:"HostIp,omitnil" name:"HostIp"`

	// 发生时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 最近发生时间
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// 事件id
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 主机uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// cvm id
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 事件类型/动作  0 -- 告警
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 进程路径
	ProcessExe *string `json:"ProcessExe,omitnil" name:"ProcessExe"`

	// 进程参数
	ProcessArgv *string `json:"ProcessArgv,omitnil" name:"ProcessArgv"`

	// 目标文件路径
	Target *string `json:"Target,omitnil" name:"Target"`

	// 处理状态  0 -- 待处理 1 -- 已加白 2 -- 已删除 3 - 已忽略 4-已手动处理
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 事件产生次数
	EventCount *uint64 `json:"EventCount,omitnil" name:"EventCount"`

	// 规则id
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`

	// 规则名称
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 事件详情: json格式
	Pstree *string `json:"Pstree,omitnil" name:"Pstree"`

	// 规则类型 0系统规则 1自定义规则
	RuleCategory *uint64 `json:"RuleCategory,omitnil" name:"RuleCategory"`

	// 主机在线信息 ONLINE、OFFLINE
	MachineStatus *string `json:"MachineStatus,omitnil" name:"MachineStatus"`

	// 危害描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 修护建议
	Suggestion *string `json:"Suggestion,omitnil" name:"Suggestion"`

	// 内网ip
	PrivateIp *string `json:"PrivateIp,omitnil" name:"PrivateIp"`

	// 进程权限
	ExePermission *string `json:"ExePermission,omitnil" name:"ExePermission"`

	// 用户名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 用户组
	UserGroup *string `json:"UserGroup,omitnil" name:"UserGroup"`

	// 进程名
	ExeMd5 *string `json:"ExeMd5,omitnil" name:"ExeMd5"`

	// 进程文件大小
	ExeSize *uint64 `json:"ExeSize,omitnil" name:"ExeSize"`

	// 进程执行时长
	ExeTime *uint64 `json:"ExeTime,omitnil" name:"ExeTime"`

	// 目标文件大小
	TargetSize *uint64 `json:"TargetSize,omitnil" name:"TargetSize"`

	// 目标文件权限
	TargetPermission *string `json:"TargetPermission,omitnil" name:"TargetPermission"`

	// 目标文件更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetModifyTime *string `json:"TargetModifyTime,omitnil" name:"TargetModifyTime"`

	// 目标文件创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetCreatTime *string `json:"TargetCreatTime,omitnil" name:"TargetCreatTime"`

	// 进程pid
	ExePid *uint64 `json:"ExePid,omitnil" name:"ExePid"`

	// 文件名称
	TargetName *string `json:"TargetName,omitnil" name:"TargetName"`

	// 参考链接
	Reference *string `json:"Reference,omitnil" name:"Reference"`

	// 风险等级 0：无， 1: 高危， 2:中危， 3: 低危
	Level *uint64 `json:"Level,omitnil" name:"Level"`

	// 进程名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExeName *string `json:"ExeName,omitnil" name:"ExeName"`

	//  主机额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`

	// 文件威胁行为
	// <li>read 读取文件</li>
	// <li>write 修改文件</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileAction *string `json:"FileAction,omitnil" name:"FileAction"`
}

type Filter struct {
	// 过滤键的名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 一个或者多个过滤值。
	Values []*string `json:"Values,omitnil" name:"Values"`

	// 模糊搜索
	ExactMatch *bool `json:"ExactMatch,omitnil" name:"ExactMatch"`
}

type Filters struct {
	// 过滤键的名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 一个或者多个过滤值。
	Values []*string `json:"Values,omitnil" name:"Values"`

	// 是否模糊匹配，前端框架会带上，可以不管
	ExactMatch *bool `json:"ExactMatch,omitnil" name:"ExactMatch"`
}

// Predefined struct for user
type FixBaselineDetectRequestParams struct {
	// 主机Id
	HostId *string `json:"HostId,omitnil" name:"HostId"`

	// 项Id
	ItemId *int64 `json:"ItemId,omitnil" name:"ItemId"`

	// 修复内容
	Data []*string `json:"Data,omitnil" name:"Data"`
}

type FixBaselineDetectRequest struct {
	*tchttp.BaseRequest
	
	// 主机Id
	HostId *string `json:"HostId,omitnil" name:"HostId"`

	// 项Id
	ItemId *int64 `json:"ItemId,omitnil" name:"ItemId"`

	// 修复内容
	Data []*string `json:"Data,omitnil" name:"Data"`
}

func (r *FixBaselineDetectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FixBaselineDetectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HostId")
	delete(f, "ItemId")
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FixBaselineDetectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FixBaselineDetectResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type FixBaselineDetectResponse struct {
	*tchttp.BaseResponse
	Response *FixBaselineDetectResponseParams `json:"Response"`
}

func (r *FixBaselineDetectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FixBaselineDetectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HistoryAccount struct {
	// 唯一ID。
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 主机安全客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 主机内网IP。
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// 主机名。
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 帐号名。
	Username *string `json:"Username,omitnil" name:"Username"`

	// 帐号变更类型。
	// <li>CREATE：表示新增帐号</li>
	// <li>MODIFY：表示修改帐号</li>
	// <li>DELETE：表示删除帐号</li>
	ModifyType *string `json:"ModifyType,omitnil" name:"ModifyType"`

	// 变更时间。
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`
}

type HostLoginList struct {
	// 记录Id
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// Uuid串
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 主机ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// 主机名
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 来源ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	SrcIp *string `json:"SrcIp,omitnil" name:"SrcIp"`

	// 1:正常登录；2异地登录； 5已加白； 14：已处理；15：已忽略。
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 国家id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Country *uint64 `json:"Country,omitnil" name:"Country"`

	// 城市id
	// 注意：此字段可能返回 null，表示取不到有效值。
	City *uint64 `json:"City,omitnil" name:"City"`

	// 省份id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Province *uint64 `json:"Province,omitnil" name:"Province"`

	// 登录时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoginTime *string `json:"LoginTime,omitnil" name:"LoginTime"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// 是否命中异地登录异常  1表示命中此类异常, 0表示未命中
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsRiskArea *uint64 `json:"IsRiskArea,omitnil" name:"IsRiskArea"`

	// 是否命中异常用户异常 1表示命中此类异常, 0表示未命中
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsRiskUser *uint64 `json:"IsRiskUser,omitnil" name:"IsRiskUser"`

	// 是否命中异常时间异常 1表示命中此类异常, 0表示未命中
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsRiskTime *uint64 `json:"IsRiskTime,omitnil" name:"IsRiskTime"`

	// 是否命中异常IP异常 1表示命中此类异常, 0表示未命中
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsRiskSrcIp *uint64 `json:"IsRiskSrcIp,omitnil" name:"IsRiskSrcIp"`

	// 危险等级：
	// 0 高危
	// 1 可疑
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevel *uint64 `json:"RiskLevel,omitnil" name:"RiskLevel"`

	// 位置名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *string `json:"Location,omitnil" name:"Location"`

	// 主机quuid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 高危信息说明：
	// ABROAD - 海外IP；
	// XTI - 威胁情报
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitnil" name:"Desc"`

	// 附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`

	// 请求目的端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitnil" name:"Port"`
}

type HostRiskLevelCount struct {
	// 主机ID
	HostId *string `json:"HostId,omitnil" name:"HostId"`

	// 主机名
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostName *string `json:"HostName,omitnil" name:"HostName"`

	// 严重个数
	SeriousCount *int64 `json:"SeriousCount,omitnil" name:"SeriousCount"`

	// 高危个数
	HighCount *int64 `json:"HighCount,omitnil" name:"HighCount"`

	// 中危个数
	MediumCount *int64 `json:"MediumCount,omitnil" name:"MediumCount"`

	// 低危个数
	LowCount *int64 `json:"LowCount,omitnil" name:"LowCount"`
}

type IgnoreBaselineRule struct {
	// 基线检测项名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 基线检测项id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// 修复建议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fix *string `json:"Fix,omitnil" name:"Fix"`

	// 影响主机数
	// 注意：此字段可能返回 null，表示取不到有效值。
	EffectHostCount *uint64 `json:"EffectHostCount,omitnil" name:"EffectHostCount"`
}

// Predefined struct for user
type IgnoreImpactedHostsRequestParams struct {
	// 漏洞ID数组。
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

type IgnoreImpactedHostsRequest struct {
	*tchttp.BaseRequest
	
	// 漏洞ID数组。
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

func (r *IgnoreImpactedHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IgnoreImpactedHostsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IgnoreImpactedHostsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IgnoreImpactedHostsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type IgnoreImpactedHostsResponse struct {
	*tchttp.BaseResponse
	Response *IgnoreImpactedHostsResponseParams `json:"Response"`
}

func (r *IgnoreImpactedHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IgnoreImpactedHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IgnoreRuleEffectHostInfo struct {
	// 主机名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostName *string `json:"HostName,omitnil" name:"HostName"`

	// 危害等级：1-低位，2-中危，3-高危，4-严重
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *uint64 `json:"Level,omitnil" name:"Level"`

	// 主机标签数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagList []*string `json:"TagList,omitnil" name:"TagList"`

	// 状态：0-未通过，1-忽略，3-已通过，5-检测中
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 最后检测时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastScanTime *string `json:"LastScanTime,omitnil" name:"LastScanTime"`

	// 事件id
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventId *uint64 `json:"EventId,omitnil" name:"EventId"`

	// 主机quuid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`
}

type IncidentVertexInfo struct {
	// 事件id
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncidentId *string `json:"IncidentId,omitnil" name:"IncidentId"`

	// 事件所在表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitnil" name:"TableName"`

	// 节点信息列表，数组项中包含节点详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vertex []*VertexInfo `json:"Vertex,omitnil" name:"Vertex"`

	// 节点总个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	VertexCount *int64 `json:"VertexCount,omitnil" name:"VertexCount"`
}

type Item struct {
	// Id
	ItemId *int64 `json:"ItemId,omitnil" name:"ItemId"`

	// 名称
	ItemName *string `json:"ItemName,omitnil" name:"ItemName"`
}

type JavaMemShellInfo struct {
	// 事件ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 服务器名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitnil" name:"Alias"`

	// 服务器IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostIp *string `json:"HostIp,omitnil" name:"HostIp"`

	// 内存马类型  0:Filter型 1:Listener型 2:Servlet型 3:Interceptors型 4:Agent型 5:其他
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 说明
	Description *string `json:"Description,omitnil" name:"Description"`

	// 首次发现时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 最近检测时间
	RecentFoundTime *string `json:"RecentFoundTime,omitnil" name:"RecentFoundTime"`

	// 处理状态  0 -- 待处理 1 -- 已加白 2 -- 已删除 3 - 已忽略  4 - 已手动处理
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 服务器quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`

	// 服务器uuid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

type LicenseBindDetail struct {
	// 机器别名
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 机器公网IP
	MachineWanIp *string `json:"MachineWanIp,omitnil" name:"MachineWanIp"`

	// 机器内网IP
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// 云服务器UUID
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机安全客户端UUID
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 标签信息
	Tags []*string `json:"Tags,omitnil" name:"Tags"`

	// 主机安全客户端状态,OFFLINE 离线,ONLINE 在线,UNINSTALL 未安装
	AgentStatus *string `json:"AgentStatus,omitnil" name:"AgentStatus"`

	// 是否允许解绑,false 不允许解绑
	IsUnBind *bool `json:"IsUnBind,omitnil" name:"IsUnBind"`

	// 是否允许换绑,false 不允许换绑
	IsSwitchBind *bool `json:"IsSwitchBind,omitnil" name:"IsSwitchBind"`

	// 主机额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`
}

type LicenseBindTaskDetail struct {
	// 云服务器UUID
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 错误信息
	ErrMsg *string `json:"ErrMsg,omitnil" name:"ErrMsg"`

	// 0 执行中, 1 成功,2失败
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 修复建议
	FixMessage *string `json:"FixMessage,omitnil" name:"FixMessage"`

	// 机器额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`
}

type LicenseDetail struct {
	// 授权ID
	LicenseId *uint64 `json:"LicenseId,omitnil" name:"LicenseId"`

	// 授权类型,0 专业版-按量计费, 1专业版-包年包月 , 2 旗舰版-包年包月
	LicenseType *uint64 `json:"LicenseType,omitnil" name:"LicenseType"`

	// 授权状态 0 未使用,1 部分使用, 2 已用完, 3 不可用
	// 注意：此字段可能返回 null，表示取不到有效值。
	LicenseStatus *uint64 `json:"LicenseStatus,omitnil" name:"LicenseStatus"`

	// 总授权数
	LicenseCnt *uint64 `json:"LicenseCnt,omitnil" name:"LicenseCnt"`

	// 已使用授权数
	UsedLicenseCnt *uint64 `json:"UsedLicenseCnt,omitnil" name:"UsedLicenseCnt"`

	// 订单状态 1 正常 2隔离, 3销毁
	OrderStatus *uint64 `json:"OrderStatus,omitnil" name:"OrderStatus"`

	// 截止日期
	Deadline *string `json:"Deadline,omitnil" name:"Deadline"`

	// 订单资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 0 初始化,1 自动续费,2 不自动续费
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 任务ID ,默认0 ,查询绑定进度用
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// 购买时间
	BuyTime *string `json:"BuyTime,omitnil" name:"BuyTime"`

	// 是否试用订单.
	SourceType *uint64 `json:"SourceType,omitnil" name:"SourceType"`

	// 资源别名
	Alias *string `json:"Alias,omitnil" name:"Alias"`

	// 平台标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tags `json:"Tags,omitnil" name:"Tags"`

	// 冻结数,当为0时 为未冻结,非0 则表示冻结授权数额
	// 注意：此字段可能返回 null，表示取不到有效值。
	FreezeNum *int64 `json:"FreezeNum,omitnil" name:"FreezeNum"`
}

type LicenseOrder struct {
	// 授权ID
	LicenseId *uint64 `json:"LicenseId,omitnil" name:"LicenseId"`

	// 授权类型
	LicenseType *uint64 `json:"LicenseType,omitnil" name:"LicenseType"`

	// 授权订单资源状态
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 订单类型
	SourceType *uint64 `json:"SourceType,omitnil" name:"SourceType"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`
}

type LicenseUnBindRsp struct {
	// QUUID 云服务器uuid,轻量服务器uuid,边缘计算 uuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 失败原因
	ErrMsg *string `json:"ErrMsg,omitnil" name:"ErrMsg"`
}

type LogStorageRecord struct {
	// 年月份
	// 注意：此字段可能返回 null，表示取不到有效值。
	Month *string `json:"Month,omitnil" name:"Month"`

	// 存储量，字节
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedSize *uint64 `json:"UsedSize,omitnil" name:"UsedSize"`

	// 总量，字节
	// 注意：此字段可能返回 null，表示取不到有效值。
	InquireSize *uint64 `json:"InquireSize,omitnil" name:"InquireSize"`
}

type LoginWhiteCombinedInfo struct {
	// 白名单地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Places []*Place `json:"Places,omitnil" name:"Places"`

	// 白名单用户（多个用户逗号隔开）
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 白名单IP（多个IP逗号隔开）
	SrcIp *string `json:"SrcIp,omitnil" name:"SrcIp"`

	// 地域字符串
	Locale *string `json:"Locale,omitnil" name:"Locale"`

	// 备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 是否对全局生效, 1：全局有效 0: 对指定主机列表生效'
	IsGlobal *uint64 `json:"IsGlobal,omitnil" name:"IsGlobal"`

	// 白名单名字：IsLocal=1时固定为：全部服务器；单台机器时为机器内网IP，多台服务器时为服务器数量，如：11台
	Name *string `json:"Name,omitnil" name:"Name"`

	// 仅在单台服务器时，返回服务器名称
	Desc *string `json:"Desc,omitnil" name:"Desc"`

	// 白名单ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 最近修改时间
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// 服务器Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 登陆地
	Locations *string `json:"Locations,omitnil" name:"Locations"`
}

type LoginWhiteLists struct {
	// 记录ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 云镜客户端ID
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 白名单地域
	Places []*Place `json:"Places,omitnil" name:"Places"`

	// 白名单用户（多个用户逗号隔开）
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 白名单IP（多个IP逗号隔开）
	SrcIp *string `json:"SrcIp,omitnil" name:"SrcIp"`

	// 是否为全局规则
	IsGlobal *bool `json:"IsGlobal,omitnil" name:"IsGlobal"`

	// 创建白名单时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 修改白名单时间
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// 机器名
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 机器IP
	HostIp *string `json:"HostIp,omitnil" name:"HostIp"`

	// 起始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

type Machine struct {
	// 主机名称。
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 主机系统。
	MachineOs *string `json:"MachineOs,omitnil" name:"MachineOs"`

	// 主机状态。
	// <li>OFFLINE: 离线  </li>
	// <li>ONLINE: 在线</li>
	// <li>SHUTDOWN: 已关机</li>
	// <li>UNINSTALLED: 未防护</li>
	MachineStatus *string `json:"MachineStatus,omitnil" name:"MachineStatus"`

	// 云镜客户端唯一Uuid，若客户端长时间不在线将返回空字符。
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// CVM或BM机器唯一Uuid。
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 漏洞数。
	VulNum *int64 `json:"VulNum,omitnil" name:"VulNum"`

	// 主机IP。
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// 是否是专业版。
	// <li>true： 是</li>
	// <li>false：否</li>
	IsProVersion *bool `json:"IsProVersion,omitnil" name:"IsProVersion"`

	// 主机外网IP。
	MachineWanIp *string `json:"MachineWanIp,omitnil" name:"MachineWanIp"`

	// 主机状态。
	// <li>POSTPAY: 表示后付费，即按量计费  </li>
	// <li>PREPAY: 表示预付费，即包年包月</li>
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// 木马数。
	MalwareNum *int64 `json:"MalwareNum,omitnil" name:"MalwareNum"`

	// 标签信息
	Tag []*MachineTag `json:"Tag,omitnil" name:"Tag"`

	// 基线风险数。
	BaselineNum *int64 `json:"BaselineNum,omitnil" name:"BaselineNum"`

	// 网络风险数。
	CyberAttackNum *int64 `json:"CyberAttackNum,omitnil" name:"CyberAttackNum"`

	// 风险状态。
	// <li>SAFE：安全</li>
	// <li>RISK：风险</li>
	// <li>UNKNOWN：未知</li>
	SecurityStatus *string `json:"SecurityStatus,omitnil" name:"SecurityStatus"`

	// 入侵事件数
	InvasionNum *int64 `json:"InvasionNum,omitnil" name:"InvasionNum"`

	// 地域信息
	RegionInfo *RegionInfo `json:"RegionInfo,omitnil" name:"RegionInfo"`

	// 实例状态 TERMINATED_PRO_VERSION 已销毁
	InstanceState *string `json:"InstanceState,omitnil" name:"InstanceState"`

	// 防篡改 授权状态 1 授权 0 未授权
	LicenseStatus *uint64 `json:"LicenseStatus,omitnil" name:"LicenseStatus"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 是否有资产扫描接口，0无，1有
	HasAssetScan *uint64 `json:"HasAssetScan,omitnil" name:"HasAssetScan"`

	// 机器所属专区类型 CVM 云服务器, BM 黑石, ECM 边缘计算, LH 轻量应用服务器 ,Other 混合云专区
	MachineType *string `json:"MachineType,omitnil" name:"MachineType"`

	// 内核版本
	KernelVersion *string `json:"KernelVersion,omitnil" name:"KernelVersion"`

	// 防护版本：BASIC_VERSION 基础版， PRO_VERSION 专业版，Flagship 旗舰版，GENERAL_DISCOUNT 普惠版
	ProtectType *string `json:"ProtectType,omitnil" name:"ProtectType"`

	// 云标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CloudTags []*Tags `json:"CloudTags,omitnil" name:"CloudTags"`

	// 是否15天内新增的主机 0：非15天内新增的主机，1：15天内增加的主机
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAddedOnTheFifteen *uint64 `json:"IsAddedOnTheFifteen,omitnil" name:"IsAddedOnTheFifteen"`

	// 主机ip列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpList *string `json:"IpList,omitnil" name:"IpList"`

	// 所属网络
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

type MachineExtraInfo struct {
	// 公网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanIP *string `json:"WanIP,omitnil" name:"WanIP"`

	// 内网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIP *string `json:"PrivateIP,omitnil" name:"PrivateIP"`

	// 网络类型，1:vpc网络 2:基础网络 3:非腾讯云网络
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkType *int64 `json:"NetworkType,omitnil" name:"NetworkType"`

	// 网络名，vpc网络情况下会返回vpc_id
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkName *string `json:"NetworkName,omitnil" name:"NetworkName"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`

	// 主机名
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostName *string `json:"HostName,omitnil" name:"HostName"`
}

type MachineTag struct {
	// 关联标签ID
	Rid *int64 `json:"Rid,omitnil" name:"Rid"`

	// 标签名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 标签ID
	TagId *uint64 `json:"TagId,omitnil" name:"TagId"`
}

type MalWareList struct {
	// 服务器ip
	HostIp *string `json:"HostIp,omitnil" name:"HostIp"`

	// 唯一UUID
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 路径
	FilePath *string `json:"FilePath,omitnil" name:"FilePath"`

	// 描述
	VirusName *string `json:"VirusName,omitnil" name:"VirusName"`

	// 状态；4-:待处理，5-已信任，6-已隔离，8-文件已删除, 14:已处理
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 唯一ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 主机别名
	Alias *string `json:"Alias,omitnil" name:"Alias"`

	// 特性标签，已废弃字段，不会再返回标签，详情中才会返回标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitnil" name:"Tags"`

	// 首次运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileCreateTime *string `json:"FileCreateTime,omitnil" name:"FileCreateTime"`

	// 最近运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileModifierTime *string `json:"FileModifierTime,omitnil" name:"FileModifierTime"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 最近扫描时间
	LatestScanTime *string `json:"LatestScanTime,omitnil" name:"LatestScanTime"`

	// 风险等级 0未知、1低、2中、3高、4严重
	Level *uint64 `json:"Level,omitnil" name:"Level"`

	// '木马检测平台用,分割 1云查杀引擎、2TAV、3binaryAi、4异常行为、5威胁情报
	CheckPlatform *string `json:"CheckPlatform,omitnil" name:"CheckPlatform"`

	// 木马进程是否存在 0:不存在，1:存在
	ProcessExists *uint64 `json:"ProcessExists,omitnil" name:"ProcessExists"`

	// 木马文件是否存在 0:不存在，1:存在
	FileExists *uint64 `json:"FileExists,omitnil" name:"FileExists"`

	// cvm quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 木马样本md5
	MD5 *string `json:"MD5,omitnil" name:"MD5"`

	// 附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`
}

type MaliciousRequestWhiteListInfo struct {
	// 白名单id
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 备注
	Mark *string `json:"Mark,omitnil" name:"Mark"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 更新时间
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`
}

type MalwareInfo struct {
	// 病毒名称
	VirusName *string `json:"VirusName,omitnil" name:"VirusName"`

	// 文件大小
	FileSize *int64 `json:"FileSize,omitnil" name:"FileSize"`

	// 文件MD5
	MD5 *string `json:"MD5,omitnil" name:"MD5"`

	// 文件地址
	FilePath *string `json:"FilePath,omitnil" name:"FilePath"`

	// 首次运行时间
	FileCreateTime *string `json:"FileCreateTime,omitnil" name:"FileCreateTime"`

	// 最近一次运行时间
	FileModifierTime *string `json:"FileModifierTime,omitnil" name:"FileModifierTime"`

	// 危害描述
	HarmDescribe *string `json:"HarmDescribe,omitnil" name:"HarmDescribe"`

	// 建议方案
	SuggestScheme *string `json:"SuggestScheme,omitnil" name:"SuggestScheme"`

	// 服务器名称
	ServersName *string `json:"ServersName,omitnil" name:"ServersName"`

	// 服务器IP
	HostIp *string `json:"HostIp,omitnil" name:"HostIp"`

	// 进程名称
	ProcessName *string `json:"ProcessName,omitnil" name:"ProcessName"`

	// 进程ID
	ProcessID *string `json:"ProcessID,omitnil" name:"ProcessID"`

	// 标签特性
	Tags []*string `json:"Tags,omitnil" name:"Tags"`

	// 影响广度 // 暂时不提供
	// 注意：此字段可能返回 null，表示取不到有效值。
	Breadth *string `json:"Breadth,omitnil" name:"Breadth"`

	// 查询热度 // 暂时不提供
	// 注意：此字段可能返回 null，表示取不到有效值。
	Heat *string `json:"Heat,omitnil" name:"Heat"`

	// 唯一ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 文件名称
	FileName *string `json:"FileName,omitnil" name:"FileName"`

	// 首次发现时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 最近扫描时间
	LatestScanTime *string `json:"LatestScanTime,omitnil" name:"LatestScanTime"`

	// 参考链接
	Reference *string `json:"Reference,omitnil" name:"Reference"`

	// 外网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineWanIp *string `json:"MachineWanIp,omitnil" name:"MachineWanIp"`

	// 进程树 json  pid:进程id，exe:文件路径 ，account:进程所属用组和用户 ,cmdline:执行命令，ssh_service: SSH服务ip, ssh_soure:登录源
	// 注意：此字段可能返回 null，表示取不到有效值。
	PsTree *string `json:"PsTree,omitnil" name:"PsTree"`

	// 主机在线状态 OFFLINE  ONLINE
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineStatus *string `json:"MachineStatus,omitnil" name:"MachineStatus"`

	// 状态；4-:待处理，5-已信任，6-已隔离
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 风险等级 0提示、1低、2中、3高、4严重
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *uint64 `json:"Level,omitnil" name:"Level"`

	// 木马检测平台用,分割 1云查杀引擎、2TAV、3binaryAi、4异常行为、5威胁情报
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckPlatform *string `json:"CheckPlatform,omitnil" name:"CheckPlatform"`

	// 主机uuid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 最近修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// 最近访问时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrFileAccessTime *string `json:"StrFileAccessTime,omitnil" name:"StrFileAccessTime"`

	// 附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`
}

type MalwareRisk struct {
	// 机器IP
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// 病毒名
	VirusName *string `json:"VirusName,omitnil" name:"VirusName"`

	// 发现时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 唯一ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

// Predefined struct for user
type ModifyAutoOpenProVersionConfigRequestParams struct {
	// 设置自动开通状态。
	// <li>CLOSE：关闭</li>
	// <li>OPEN：打开</li>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 自动加购/扩容授权开关,默认 1, 0关闭, 1开启
	AutoRepurchaseSwitch *uint64 `json:"AutoRepurchaseSwitch,omitnil" name:"AutoRepurchaseSwitch"`

	// 自动加购的订单是否自动续费,默认0 ,0关闭, 1开启
	AutoRepurchaseRenewSwitch *uint64 `json:"AutoRepurchaseRenewSwitch,omitnil" name:"AutoRepurchaseRenewSwitch"`
}

type ModifyAutoOpenProVersionConfigRequest struct {
	*tchttp.BaseRequest
	
	// 设置自动开通状态。
	// <li>CLOSE：关闭</li>
	// <li>OPEN：打开</li>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 自动加购/扩容授权开关,默认 1, 0关闭, 1开启
	AutoRepurchaseSwitch *uint64 `json:"AutoRepurchaseSwitch,omitnil" name:"AutoRepurchaseSwitch"`

	// 自动加购的订单是否自动续费,默认0 ,0关闭, 1开启
	AutoRepurchaseRenewSwitch *uint64 `json:"AutoRepurchaseRenewSwitch,omitnil" name:"AutoRepurchaseRenewSwitch"`
}

func (r *ModifyAutoOpenProVersionConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoOpenProVersionConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "AutoRepurchaseSwitch")
	delete(f, "AutoRepurchaseRenewSwitch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAutoOpenProVersionConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoOpenProVersionConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAutoOpenProVersionConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAutoOpenProVersionConfigResponseParams `json:"Response"`
}

func (r *ModifyAutoOpenProVersionConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoOpenProVersionConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBanModeRequestParams struct {
	// 阻断模式，STANDARD_MODE：标准阻断，DEEP_MODE：深度阻断
	Mode *string `json:"Mode,omitnil" name:"Mode"`

	// 阻断时间，用于标准阻断模式
	Ttl *uint64 `json:"Ttl,omitnil" name:"Ttl"`
}

type ModifyBanModeRequest struct {
	*tchttp.BaseRequest
	
	// 阻断模式，STANDARD_MODE：标准阻断，DEEP_MODE：深度阻断
	Mode *string `json:"Mode,omitnil" name:"Mode"`

	// 阻断时间，用于标准阻断模式
	Ttl *uint64 `json:"Ttl,omitnil" name:"Ttl"`
}

func (r *ModifyBanModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBanModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Mode")
	delete(f, "Ttl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBanModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBanModeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyBanModeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBanModeResponseParams `json:"Response"`
}

func (r *ModifyBanModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBanModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBanStatusRequestParams struct {
	// 阻断状态 0:关闭 1:开启
	Status *uint64 `json:"Status,omitnil" name:"Status"`
}

type ModifyBanStatusRequest struct {
	*tchttp.BaseRequest
	
	// 阻断状态 0:关闭 1:开启
	Status *uint64 `json:"Status,omitnil" name:"Status"`
}

func (r *ModifyBanStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBanStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBanStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBanStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyBanStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBanStatusResponseParams `json:"Response"`
}

func (r *ModifyBanStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBanStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBaselinePolicyRequestParams struct {
	// 无
	Data *BaselinePolicy `json:"Data,omitnil" name:"Data"`

	// <li>RuleName - String - 是否必填：否 - 规则名称</li>
	// <li>CategoryId - int64 - 是否必填：否 自定义筛选为-1 - 规则分类</li>
	// <li>RuleType - int - 是否必填：否 0:系统 1:自定义 - 规则类型</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 是否按照过滤的全选
	SelectAll *int64 `json:"SelectAll,omitnil" name:"SelectAll"`
}

type ModifyBaselinePolicyRequest struct {
	*tchttp.BaseRequest
	
	// 无
	Data *BaselinePolicy `json:"Data,omitnil" name:"Data"`

	// <li>RuleName - String - 是否必填：否 - 规则名称</li>
	// <li>CategoryId - int64 - 是否必填：否 自定义筛选为-1 - 规则分类</li>
	// <li>RuleType - int - 是否必填：否 0:系统 1:自定义 - 规则类型</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 是否按照过滤的全选
	SelectAll *int64 `json:"SelectAll,omitnil" name:"SelectAll"`
}

func (r *ModifyBaselinePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBaselinePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Data")
	delete(f, "Filters")
	delete(f, "SelectAll")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBaselinePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBaselinePolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyBaselinePolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBaselinePolicyResponseParams `json:"Response"`
}

func (r *ModifyBaselinePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBaselinePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBaselinePolicyStateRequestParams struct {
	// 策略Id
	PolicyId *int64 `json:"PolicyId,omitnil" name:"PolicyId"`

	// 开启状态[1:开启|0:未开启]
	IsEnabled *int64 `json:"IsEnabled,omitnil" name:"IsEnabled"`
}

type ModifyBaselinePolicyStateRequest struct {
	*tchttp.BaseRequest
	
	// 策略Id
	PolicyId *int64 `json:"PolicyId,omitnil" name:"PolicyId"`

	// 开启状态[1:开启|0:未开启]
	IsEnabled *int64 `json:"IsEnabled,omitnil" name:"IsEnabled"`
}

func (r *ModifyBaselinePolicyStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBaselinePolicyStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "IsEnabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBaselinePolicyStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBaselinePolicyStateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyBaselinePolicyStateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBaselinePolicyStateResponseParams `json:"Response"`
}

func (r *ModifyBaselinePolicyStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBaselinePolicyStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBaselineRuleIgnoreRequestParams struct {
	// 规则名称
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 资产类型[0:所有专业版旗舰版|1:id|2:ip]
	AssetType *int64 `json:"AssetType,omitnil" name:"AssetType"`

	// 规则Id
	RuleId *int64 `json:"RuleId,omitnil" name:"RuleId"`

	// 关联项
	ItemIds []*int64 `json:"ItemIds,omitnil" name:"ItemIds"`

	// 主机Id
	HostIds []*string `json:"HostIds,omitnil" name:"HostIds"`

	// 主机Ip
	HostIps []*string `json:"HostIps,omitnil" name:"HostIps"`

	// 是否全选过滤
	SelectAll *int64 `json:"SelectAll,omitnil" name:"SelectAll"`

	// <li>ItemName - string - 是否必填：否 - 项名称</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type ModifyBaselineRuleIgnoreRequest struct {
	*tchttp.BaseRequest
	
	// 规则名称
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 资产类型[0:所有专业版旗舰版|1:id|2:ip]
	AssetType *int64 `json:"AssetType,omitnil" name:"AssetType"`

	// 规则Id
	RuleId *int64 `json:"RuleId,omitnil" name:"RuleId"`

	// 关联项
	ItemIds []*int64 `json:"ItemIds,omitnil" name:"ItemIds"`

	// 主机Id
	HostIds []*string `json:"HostIds,omitnil" name:"HostIds"`

	// 主机Ip
	HostIps []*string `json:"HostIps,omitnil" name:"HostIps"`

	// 是否全选过滤
	SelectAll *int64 `json:"SelectAll,omitnil" name:"SelectAll"`

	// <li>ItemName - string - 是否必填：否 - 项名称</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *ModifyBaselineRuleIgnoreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBaselineRuleIgnoreRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleName")
	delete(f, "AssetType")
	delete(f, "RuleId")
	delete(f, "ItemIds")
	delete(f, "HostIds")
	delete(f, "HostIps")
	delete(f, "SelectAll")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBaselineRuleIgnoreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBaselineRuleIgnoreResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyBaselineRuleIgnoreResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBaselineRuleIgnoreResponseParams `json:"Response"`
}

func (r *ModifyBaselineRuleIgnoreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBaselineRuleIgnoreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBaselineRuleRequestParams struct {
	// 无
	Data *BaselineRule `json:"Data,omitnil" name:"Data"`

	// 是否过滤全选
	SelectAll *int64 `json:"SelectAll,omitnil" name:"SelectAll"`

	// <li>ItemName - string - 是否必填：否 - 项名称</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type ModifyBaselineRuleRequest struct {
	*tchttp.BaseRequest
	
	// 无
	Data *BaselineRule `json:"Data,omitnil" name:"Data"`

	// 是否过滤全选
	SelectAll *int64 `json:"SelectAll,omitnil" name:"SelectAll"`

	// <li>ItemName - string - 是否必填：否 - 项名称</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *ModifyBaselineRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBaselineRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Data")
	delete(f, "SelectAll")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBaselineRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBaselineRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyBaselineRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBaselineRuleResponseParams `json:"Response"`
}

func (r *ModifyBaselineRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBaselineRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBaselineWeakPasswordRequestParams struct {
	// 无
	Data []*BaselineWeakPassword `json:"Data,omitnil" name:"Data"`
}

type ModifyBaselineWeakPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 无
	Data []*BaselineWeakPassword `json:"Data,omitnil" name:"Data"`
}

func (r *ModifyBaselineWeakPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBaselineWeakPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBaselineWeakPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBaselineWeakPasswordResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyBaselineWeakPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBaselineWeakPasswordResponseParams `json:"Response"`
}

func (r *ModifyBaselineWeakPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBaselineWeakPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBruteAttackRulesRequestParams struct {
	// 暴力破解判断规则
	Rules []*BruteAttackRule `json:"Rules,omitnil" name:"Rules"`
}

type ModifyBruteAttackRulesRequest struct {
	*tchttp.BaseRequest
	
	// 暴力破解判断规则
	Rules []*BruteAttackRule `json:"Rules,omitnil" name:"Rules"`
}

func (r *ModifyBruteAttackRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBruteAttackRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBruteAttackRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBruteAttackRulesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyBruteAttackRulesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBruteAttackRulesResponseParams `json:"Response"`
}

func (r *ModifyBruteAttackRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBruteAttackRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLicenseBindsRequestParams struct {
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 授权类型
	LicenseType *uint64 `json:"LicenseType,omitnil" name:"LicenseType"`

	// 是否全部机器(当全部机器数大于当前订单可用授权数时,多余机器会被跳过)
	IsAll *bool `json:"IsAll,omitnil" name:"IsAll"`

	// 需要绑定的机器quuid列表, 当IsAll = false 时必填,反之忽略该参数. 最大长度=2000
	QuuidList []*string `json:"QuuidList,omitnil" name:"QuuidList"`
}

type ModifyLicenseBindsRequest struct {
	*tchttp.BaseRequest
	
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 授权类型
	LicenseType *uint64 `json:"LicenseType,omitnil" name:"LicenseType"`

	// 是否全部机器(当全部机器数大于当前订单可用授权数时,多余机器会被跳过)
	IsAll *bool `json:"IsAll,omitnil" name:"IsAll"`

	// 需要绑定的机器quuid列表, 当IsAll = false 时必填,反之忽略该参数. 最大长度=2000
	QuuidList []*string `json:"QuuidList,omitnil" name:"QuuidList"`
}

func (r *ModifyLicenseBindsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLicenseBindsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "LicenseType")
	delete(f, "IsAll")
	delete(f, "QuuidList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLicenseBindsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLicenseBindsResponseParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyLicenseBindsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLicenseBindsResponseParams `json:"Response"`
}

func (r *ModifyLicenseBindsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLicenseBindsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLicenseOrderRequestParams struct {
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 预期值,如果当前为10,扩容则输入原来大的值, 缩容则比原来小的值(缩容时不允许预期值比使用量小),如果保持不变则填写原值,
	InquireNum *uint64 `json:"InquireNum,omitnil" name:"InquireNum"`

	// 项目ID,不修改则输入原值.
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 资源别名,不修改则输入原值.
	Alias *string `json:"Alias,omitnil" name:"Alias"`
}

type ModifyLicenseOrderRequest struct {
	*tchttp.BaseRequest
	
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 预期值,如果当前为10,扩容则输入原来大的值, 缩容则比原来小的值(缩容时不允许预期值比使用量小),如果保持不变则填写原值,
	InquireNum *uint64 `json:"InquireNum,omitnil" name:"InquireNum"`

	// 项目ID,不修改则输入原值.
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 资源别名,不修改则输入原值.
	Alias *string `json:"Alias,omitnil" name:"Alias"`
}

func (r *ModifyLicenseOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLicenseOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "InquireNum")
	delete(f, "ProjectId")
	delete(f, "Alias")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLicenseOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLicenseOrderResponseParams struct {
	// 订单号
	DealNames []*string `json:"DealNames,omitnil" name:"DealNames"`

	// 资源ID
	ResourceIds []*string `json:"ResourceIds,omitnil" name:"ResourceIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyLicenseOrderResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLicenseOrderResponseParams `json:"Response"`
}

func (r *ModifyLicenseOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLicenseOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLicenseUnBindsRequestParams struct {
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 授权类型
	LicenseType *uint64 `json:"LicenseType,omitnil" name:"LicenseType"`

	// 是否全部机器(当全部机器数大于当前订单可用授权数时,多余机器会被跳过)
	IsAll *bool `json:"IsAll,omitnil" name:"IsAll"`

	// 需要绑定的机器quuid列表, 当IsAll = false 时必填,反之忽略该参数.
	// 最大长度=100
	QuuidList []*string `json:"QuuidList,omitnil" name:"QuuidList"`
}

type ModifyLicenseUnBindsRequest struct {
	*tchttp.BaseRequest
	
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 授权类型
	LicenseType *uint64 `json:"LicenseType,omitnil" name:"LicenseType"`

	// 是否全部机器(当全部机器数大于当前订单可用授权数时,多余机器会被跳过)
	IsAll *bool `json:"IsAll,omitnil" name:"IsAll"`

	// 需要绑定的机器quuid列表, 当IsAll = false 时必填,反之忽略该参数.
	// 最大长度=100
	QuuidList []*string `json:"QuuidList,omitnil" name:"QuuidList"`
}

func (r *ModifyLicenseUnBindsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLicenseUnBindsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "LicenseType")
	delete(f, "IsAll")
	delete(f, "QuuidList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLicenseUnBindsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLicenseUnBindsResponseParams struct {
	// 只有解绑失败的才有该值.
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg []*LicenseUnBindRsp `json:"ErrMsg,omitnil" name:"ErrMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyLicenseUnBindsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLicenseUnBindsResponseParams `json:"Response"`
}

func (r *ModifyLicenseUnBindsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLicenseUnBindsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLogStorageConfigRequestParams struct {
	// 是否修改有效期
	IsModifyPeriod *bool `json:"IsModifyPeriod,omitnil" name:"IsModifyPeriod"`

	// 存储类型，string数组
	Type []*string `json:"Type,omitnil" name:"Type"`

	// 日志存储天数，3640表示不限
	Period *int64 `json:"Period,omitnil" name:"Period"`
}

type ModifyLogStorageConfigRequest struct {
	*tchttp.BaseRequest
	
	// 是否修改有效期
	IsModifyPeriod *bool `json:"IsModifyPeriod,omitnil" name:"IsModifyPeriod"`

	// 存储类型，string数组
	Type []*string `json:"Type,omitnil" name:"Type"`

	// 日志存储天数，3640表示不限
	Period *int64 `json:"Period,omitnil" name:"Period"`
}

func (r *ModifyLogStorageConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLogStorageConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IsModifyPeriod")
	delete(f, "Type")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLogStorageConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLogStorageConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyLogStorageConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLogStorageConfigResponseParams `json:"Response"`
}

func (r *ModifyLogStorageConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLogStorageConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMachineRemarkRequestParams struct {
	// 主机Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 备注信息
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

type ModifyMachineRemarkRequest struct {
	*tchttp.BaseRequest
	
	// 主机Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 备注信息
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

func (r *ModifyMachineRemarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMachineRemarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Quuid")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMachineRemarkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMachineRemarkResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyMachineRemarkResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMachineRemarkResponseParams `json:"Response"`
}

func (r *ModifyMachineRemarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMachineRemarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMalwareTimingScanSettingsRequestParams struct {
	// 检测模式 0 全盘检测  1快速检测
	CheckPattern *uint64 `json:"CheckPattern,omitnil" name:"CheckPattern"`

	// 检测周期 开始时间，如：02:00:00
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 检测周期 超时结束时间，如：04:00:00
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 是否全部服务器 1 全部 2 自选
	IsGlobal *uint64 `json:"IsGlobal,omitnil" name:"IsGlobal"`

	// 定时检测开关 0 关闭 1开启
	EnableScan *uint64 `json:"EnableScan,omitnil" name:"EnableScan"`

	// 监控模式 0 标准 1深度
	MonitoringPattern *uint64 `json:"MonitoringPattern,omitnil" name:"MonitoringPattern"`

	// 扫描周期 默认每天 1
	Cycle *uint64 `json:"Cycle,omitnil" name:"Cycle"`

	// 实时监控 0 关闭 1开启
	RealTimeMonitoring *uint64 `json:"RealTimeMonitoring,omitnil" name:"RealTimeMonitoring"`

	// 自选服务器时必须 主机quuid的string数组
	QuuidList []*string `json:"QuuidList,omitnil" name:"QuuidList"`

	// 是否自动隔离 1隔离 0 不隔离
	AutoIsolation *uint64 `json:"AutoIsolation,omitnil" name:"AutoIsolation"`

	// 是否杀掉进程 1杀掉 0不杀掉
	KillProcess *uint64 `json:"KillProcess,omitnil" name:"KillProcess"`

	// 1标准模式（只报严重、高危）、2增强模式（报严重、高危、中危）、3严格模式（报严重、高、中、低、提示）
	EngineType *uint64 `json:"EngineType,omitnil" name:"EngineType"`

	// 启发引擎开关 0 关闭 1开启
	EnableInspiredEngine *uint64 `json:"EnableInspiredEngine,omitnil" name:"EnableInspiredEngine"`

	// 是否开启恶意进程查杀[0:未开启,1:开启]
	EnableMemShellScan *uint64 `json:"EnableMemShellScan,omitnil" name:"EnableMemShellScan"`
}

type ModifyMalwareTimingScanSettingsRequest struct {
	*tchttp.BaseRequest
	
	// 检测模式 0 全盘检测  1快速检测
	CheckPattern *uint64 `json:"CheckPattern,omitnil" name:"CheckPattern"`

	// 检测周期 开始时间，如：02:00:00
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 检测周期 超时结束时间，如：04:00:00
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 是否全部服务器 1 全部 2 自选
	IsGlobal *uint64 `json:"IsGlobal,omitnil" name:"IsGlobal"`

	// 定时检测开关 0 关闭 1开启
	EnableScan *uint64 `json:"EnableScan,omitnil" name:"EnableScan"`

	// 监控模式 0 标准 1深度
	MonitoringPattern *uint64 `json:"MonitoringPattern,omitnil" name:"MonitoringPattern"`

	// 扫描周期 默认每天 1
	Cycle *uint64 `json:"Cycle,omitnil" name:"Cycle"`

	// 实时监控 0 关闭 1开启
	RealTimeMonitoring *uint64 `json:"RealTimeMonitoring,omitnil" name:"RealTimeMonitoring"`

	// 自选服务器时必须 主机quuid的string数组
	QuuidList []*string `json:"QuuidList,omitnil" name:"QuuidList"`

	// 是否自动隔离 1隔离 0 不隔离
	AutoIsolation *uint64 `json:"AutoIsolation,omitnil" name:"AutoIsolation"`

	// 是否杀掉进程 1杀掉 0不杀掉
	KillProcess *uint64 `json:"KillProcess,omitnil" name:"KillProcess"`

	// 1标准模式（只报严重、高危）、2增强模式（报严重、高危、中危）、3严格模式（报严重、高、中、低、提示）
	EngineType *uint64 `json:"EngineType,omitnil" name:"EngineType"`

	// 启发引擎开关 0 关闭 1开启
	EnableInspiredEngine *uint64 `json:"EnableInspiredEngine,omitnil" name:"EnableInspiredEngine"`

	// 是否开启恶意进程查杀[0:未开启,1:开启]
	EnableMemShellScan *uint64 `json:"EnableMemShellScan,omitnil" name:"EnableMemShellScan"`
}

func (r *ModifyMalwareTimingScanSettingsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMalwareTimingScanSettingsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CheckPattern")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "IsGlobal")
	delete(f, "EnableScan")
	delete(f, "MonitoringPattern")
	delete(f, "Cycle")
	delete(f, "RealTimeMonitoring")
	delete(f, "QuuidList")
	delete(f, "AutoIsolation")
	delete(f, "KillProcess")
	delete(f, "EngineType")
	delete(f, "EnableInspiredEngine")
	delete(f, "EnableMemShellScan")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMalwareTimingScanSettingsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMalwareTimingScanSettingsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyMalwareTimingScanSettingsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMalwareTimingScanSettingsResponseParams `json:"Response"`
}

func (r *ModifyMalwareTimingScanSettingsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMalwareTimingScanSettingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOrderAttributeRequestParams struct {
	// 授权类型 0 专业版-按量计费, 1专业版-包年包月 , 2 旗舰版-包年包月
	LicenseType *uint64 `json:"LicenseType,omitnil" name:"LicenseType"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 可编辑的属性名称 ,当前支持的有: alias 资源别名
	AttrName *string `json:"AttrName,omitnil" name:"AttrName"`

	// 属性值
	AttrValue *string `json:"AttrValue,omitnil" name:"AttrValue"`
}

type ModifyOrderAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 授权类型 0 专业版-按量计费, 1专业版-包年包月 , 2 旗舰版-包年包月
	LicenseType *uint64 `json:"LicenseType,omitnil" name:"LicenseType"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 可编辑的属性名称 ,当前支持的有: alias 资源别名
	AttrName *string `json:"AttrName,omitnil" name:"AttrName"`

	// 属性值
	AttrValue *string `json:"AttrValue,omitnil" name:"AttrValue"`
}

func (r *ModifyOrderAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyOrderAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LicenseType")
	delete(f, "ResourceId")
	delete(f, "AttrName")
	delete(f, "AttrValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyOrderAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOrderAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyOrderAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyOrderAttributeResponseParams `json:"Response"`
}

func (r *ModifyOrderAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyOrderAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWarningSettingRequestParams struct {
	// 告警设置的修改内容
	WarningObjects []*WarningObject `json:"WarningObjects,omitnil" name:"WarningObjects"`
}

type ModifyWarningSettingRequest struct {
	*tchttp.BaseRequest
	
	// 告警设置的修改内容
	WarningObjects []*WarningObject `json:"WarningObjects,omitnil" name:"WarningObjects"`
}

func (r *ModifyWarningSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWarningSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WarningObjects")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWarningSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWarningSettingResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyWarningSettingResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWarningSettingResponseParams `json:"Response"`
}

func (r *ModifyWarningSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWarningSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWebPageProtectDirRequestParams struct {
	// 网站防护目录地址
	ProtectDirAddr *string `json:"ProtectDirAddr,omitnil" name:"ProtectDirAddr"`

	// 网站防护目录名称
	ProtectDirName *string `json:"ProtectDirName,omitnil" name:"ProtectDirName"`

	// 防护文件类型,分号分割 ;
	ProtectFileType *string `json:"ProtectFileType,omitnil" name:"ProtectFileType"`

	// 防护机器列表信息
	HostConfig []*ProtectHostConfig `json:"HostConfig,omitnil" name:"HostConfig"`
}

type ModifyWebPageProtectDirRequest struct {
	*tchttp.BaseRequest
	
	// 网站防护目录地址
	ProtectDirAddr *string `json:"ProtectDirAddr,omitnil" name:"ProtectDirAddr"`

	// 网站防护目录名称
	ProtectDirName *string `json:"ProtectDirName,omitnil" name:"ProtectDirName"`

	// 防护文件类型,分号分割 ;
	ProtectFileType *string `json:"ProtectFileType,omitnil" name:"ProtectFileType"`

	// 防护机器列表信息
	HostConfig []*ProtectHostConfig `json:"HostConfig,omitnil" name:"HostConfig"`
}

func (r *ModifyWebPageProtectDirRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWebPageProtectDirRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProtectDirAddr")
	delete(f, "ProtectDirName")
	delete(f, "ProtectFileType")
	delete(f, "HostConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWebPageProtectDirRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWebPageProtectDirResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyWebPageProtectDirResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWebPageProtectDirResponseParams `json:"Response"`
}

func (r *ModifyWebPageProtectDirResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWebPageProtectDirResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWebPageProtectSettingRequestParams struct {
	// 需要操作的类型1 目录名称 2 防护文件类型
	ModifyType *uint64 `json:"ModifyType,omitnil" name:"ModifyType"`

	// 提交值
	Value *string `json:"Value,omitnil" name:"Value"`

	// 配置对应的protect_path
	Id *string `json:"Id,omitnil" name:"Id"`
}

type ModifyWebPageProtectSettingRequest struct {
	*tchttp.BaseRequest
	
	// 需要操作的类型1 目录名称 2 防护文件类型
	ModifyType *uint64 `json:"ModifyType,omitnil" name:"ModifyType"`

	// 提交值
	Value *string `json:"Value,omitnil" name:"Value"`

	// 配置对应的protect_path
	Id *string `json:"Id,omitnil" name:"Id"`
}

func (r *ModifyWebPageProtectSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWebPageProtectSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModifyType")
	delete(f, "Value")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWebPageProtectSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWebPageProtectSettingResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyWebPageProtectSettingResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWebPageProtectSettingResponseParams `json:"Response"`
}

func (r *ModifyWebPageProtectSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWebPageProtectSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWebPageProtectSwitchRequestParams struct {
	// 开关类型 1 防护开关  2 自动恢复开关 3 移除防护目录
	SwitchType *uint64 `json:"SwitchType,omitnil" name:"SwitchType"`

	// 需要操作开关的网站 最大100条
	Ids []*string `json:"Ids,omitnil" name:"Ids"`

	// 1 开启 0 关闭 SwitchType 为 1 | 2 必填;
	Status *uint64 `json:"Status,omitnil" name:"Status"`
}

type ModifyWebPageProtectSwitchRequest struct {
	*tchttp.BaseRequest
	
	// 开关类型 1 防护开关  2 自动恢复开关 3 移除防护目录
	SwitchType *uint64 `json:"SwitchType,omitnil" name:"SwitchType"`

	// 需要操作开关的网站 最大100条
	Ids []*string `json:"Ids,omitnil" name:"Ids"`

	// 1 开启 0 关闭 SwitchType 为 1 | 2 必填;
	Status *uint64 `json:"Status,omitnil" name:"Status"`
}

func (r *ModifyWebPageProtectSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWebPageProtectSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SwitchType")
	delete(f, "Ids")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWebPageProtectSwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWebPageProtectSwitchResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyWebPageProtectSwitchResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWebPageProtectSwitchResponseParams `json:"Response"`
}

func (r *ModifyWebPageProtectSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWebPageProtectSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonthInspectionReport struct {
	// 巡检报告名称
	ReportName *string `json:"ReportName,omitnil" name:"ReportName"`

	// 巡检报告下载地址
	ReportPath *string `json:"ReportPath,omitnil" name:"ReportPath"`

	// 巡检报告更新时间
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`
}

type NetAttackEvent struct {
	// 日志ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 客户端ID
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 目标端口
	DstPort *uint64 `json:"DstPort,omitnil" name:"DstPort"`

	// 来源IP
	SrcIP *string `json:"SrcIP,omitnil" name:"SrcIP"`

	// 来源地
	Location *string `json:"Location,omitnil" name:"Location"`

	// 漏洞id
	VulId *uint64 `json:"VulId,omitnil" name:"VulId"`

	// 漏洞名称
	VulName *string `json:"VulName,omitnil" name:"VulName"`

	// 攻击时间
	MergeTime *string `json:"MergeTime,omitnil" name:"MergeTime"`

	// 主机额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`

	// 攻击状态，0: 尝试攻击 1: 实锤攻击(攻击成功)
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 处理状态，0 待处理 1 已处理 2 已加白  3 已忽略 4 已删除 5: 已开启防御
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 漏洞是否支持防御，0:不支持 1:支持
	VulSupportDefense *uint64 `json:"VulSupportDefense,omitnil" name:"VulSupportDefense"`

	// 是否开启漏洞防御，0关1开
	VulDefenceStatus *uint64 `json:"VulDefenceStatus,omitnil" name:"VulDefenceStatus"`

	// 机器付费版本，0 基础版，1专业版，2旗舰版，3普惠版
	PayVersion *uint64 `json:"PayVersion,omitnil" name:"PayVersion"`

	// cvm uuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 攻击次数
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// 是否今日新增主机
	New *bool `json:"New,omitnil" name:"New"`
}

type OpenPortStatistics struct {
	// 端口号
	Port *uint64 `json:"Port,omitnil" name:"Port"`

	// 主机数量
	MachineNum *uint64 `json:"MachineNum,omitnil" name:"MachineNum"`
}

type OrderModifyObject struct {
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 新产品标识,这里支持PRO_VERSION 专业版,FLAGSHIP 旗舰版
	NewSubProductCode *string `json:"NewSubProductCode,omitnil" name:"NewSubProductCode"`

	// 扩容/缩容数,变配子产品忽略该参数
	InquireNum *int64 `json:"InquireNum,omitnil" name:"InquireNum"`
}

type OsName struct {
	// 系统名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 操作系统类型枚举值
	MachineOSType *uint64 `json:"MachineOSType,omitnil" name:"MachineOSType"`
}

type Place struct {
	// 城市 ID。
	CityId *uint64 `json:"CityId,omitnil" name:"CityId"`

	// 省份 ID。
	ProvinceId *uint64 `json:"ProvinceId,omitnil" name:"ProvinceId"`

	// 国家ID，暂只支持国内：1。
	CountryId *uint64 `json:"CountryId,omitnil" name:"CountryId"`

	// 位置名称
	Location *string `json:"Location,omitnil" name:"Location"`
}

type PrivilegeEscalationProcess struct {
	// 数据ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 主机安全ID
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 主机ID
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机内网IP
	Hostip *string `json:"Hostip,omitnil" name:"Hostip"`

	// 进程名
	ProcessName *string `json:"ProcessName,omitnil" name:"ProcessName"`

	// 进程路径
	FullPath *string `json:"FullPath,omitnil" name:"FullPath"`

	// 执行命令
	CmdLine *string `json:"CmdLine,omitnil" name:"CmdLine"`

	// 用户名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 用户组
	UserGroup *string `json:"UserGroup,omitnil" name:"UserGroup"`

	// 进程文件权限
	ProcFilePrivilege *string `json:"ProcFilePrivilege,omitnil" name:"ProcFilePrivilege"`

	// 父进程名
	ParentProcName *string `json:"ParentProcName,omitnil" name:"ParentProcName"`

	// 父进程用户名
	ParentProcUser *string `json:"ParentProcUser,omitnil" name:"ParentProcUser"`

	// 父进程用户组
	ParentProcGroup *string `json:"ParentProcGroup,omitnil" name:"ParentProcGroup"`

	// 父进程路径
	ParentProcPath *string `json:"ParentProcPath,omitnil" name:"ParentProcPath"`

	// 进程树
	ProcTree *string `json:"ProcTree,omitnil" name:"ProcTree"`

	// 处理状态：0-待处理 2-白名单 3-已处理 4-已忽略
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 发生时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 机器名
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`

	// 进程id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pid *int64 `json:"Pid,omitnil" name:"Pid"`
}

type PrivilegeEventInfo struct {
	// 数据ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 云镜ID
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 主机ID
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机内网IP
	HostIp *string `json:"HostIp,omitnil" name:"HostIp"`

	// 进程名
	ProcessName *string `json:"ProcessName,omitnil" name:"ProcessName"`

	// 进程路径
	FullPath *string `json:"FullPath,omitnil" name:"FullPath"`

	// 执行命令
	CmdLine *string `json:"CmdLine,omitnil" name:"CmdLine"`

	// 用户名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 用户组
	UserGroup *string `json:"UserGroup,omitnil" name:"UserGroup"`

	// 进程文件权限
	ProcFilePrivilege *string `json:"ProcFilePrivilege,omitnil" name:"ProcFilePrivilege"`

	// 父进程名
	ParentProcName *string `json:"ParentProcName,omitnil" name:"ParentProcName"`

	// 父进程用户名
	ParentProcUser *string `json:"ParentProcUser,omitnil" name:"ParentProcUser"`

	// 父进程用户组
	ParentProcGroup *string `json:"ParentProcGroup,omitnil" name:"ParentProcGroup"`

	// 父进程路径
	ParentProcPath *string `json:"ParentProcPath,omitnil" name:"ParentProcPath"`

	// 进程树 json  pid:进程id，exe:文件路径 ，account:进程所属用组和用户 ,cmdline:执行命令，ssh_service: SSH服务ip, ssh_soure:登录源
	PsTree *string `json:"PsTree,omitnil" name:"PsTree"`

	// 处理状态：0-待处理 2-白名单 3-已处理 4-已忽略
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 发生时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 机器名
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 建议方案
	SuggestScheme *string `json:"SuggestScheme,omitnil" name:"SuggestScheme"`

	// 危害描述信息
	HarmDescribe *string `json:"HarmDescribe,omitnil" name:"HarmDescribe"`

	// 标签
	Tags []*string `json:"Tags,omitnil" name:"Tags"`

	// 参考链接
	References []*string `json:"References,omitnil" name:"References"`

	// 主机外网ip
	MachineWanIp *string `json:"MachineWanIp,omitnil" name:"MachineWanIp"`

	// 权限列表|隔开
	NewCaps *string `json:"NewCaps,omitnil" name:"NewCaps"`

	// 主机在线状态 OFFLINE  ONLINE
	MachineStatus *string `json:"MachineStatus,omitnil" name:"MachineStatus"`

	// 处理时间
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`
}

type PrivilegeRule struct {
	// 规则ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 客户端ID
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 进程名
	ProcessName *string `json:"ProcessName,omitnil" name:"ProcessName"`

	// 是否S权限
	SMode *uint64 `json:"SMode,omitnil" name:"SMode"`

	// 操作人
	Operator *string `json:"Operator,omitnil" name:"Operator"`

	// 是否全局规则
	IsGlobal *uint64 `json:"IsGlobal,omitnil" name:"IsGlobal"`

	// 状态(0: 有效 1: 无效)
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// 主机IP
	Hostip *string `json:"Hostip,omitnil" name:"Hostip"`
}

type ProcessStatistics struct {
	// 进程名。
	ProcessName *string `json:"ProcessName,omitnil" name:"ProcessName"`

	// 主机数量。
	MachineNum *uint64 `json:"MachineNum,omitnil" name:"MachineNum"`
}

type ProtectDirInfo struct {
	// 网站名称
	DirName *string `json:"DirName,omitnil" name:"DirName"`

	// 网站防护目录地址
	DirPath *string `json:"DirPath,omitnil" name:"DirPath"`

	// 关联服务器数
	RelatedServerNum *uint64 `json:"RelatedServerNum,omitnil" name:"RelatedServerNum"`

	// 防护服务器数
	ProtectServerNum *uint64 `json:"ProtectServerNum,omitnil" name:"ProtectServerNum"`

	// 未防护服务器数
	NoProtectServerNum *uint64 `json:"NoProtectServerNum,omitnil" name:"NoProtectServerNum"`

	// 唯一ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// 防护状态
	ProtectStatus *uint64 `json:"ProtectStatus,omitnil" name:"ProtectStatus"`

	// 防护异常
	ProtectException *uint64 `json:"ProtectException,omitnil" name:"ProtectException"`

	// 自动恢复开关 (Filters 过滤Quuid 时 返回) 默认0
	AutoRestoreSwitchStatus *uint64 `json:"AutoRestoreSwitchStatus,omitnil" name:"AutoRestoreSwitchStatus"`

	// 首次开启防护时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstProtectTime *string `json:"FirstProtectTime,omitnil" name:"FirstProtectTime"`

	// 最近开启防护时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestProtectTime *string `json:"LatestProtectTime,omitnil" name:"LatestProtectTime"`

	// 防护文件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtectFileType *string `json:"ProtectFileType,omitnil" name:"ProtectFileType"`

	// 防护文件总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtectFilesCount *int64 `json:"ProtectFilesCount,omitnil" name:"ProtectFilesCount"`
}

type ProtectDirRelatedServer struct {
	// 唯一ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// 服务器名称
	HostName *string `json:"HostName,omitnil" name:"HostName"`

	// 服务器IP
	HostIp *string `json:"HostIp,omitnil" name:"HostIp"`

	// 服务器系统
	MachineOs *string `json:"MachineOs,omitnil" name:"MachineOs"`

	// 关联目录数
	RelateDirNum *uint64 `json:"RelateDirNum,omitnil" name:"RelateDirNum"`

	// 防护状态
	ProtectStatus *uint64 `json:"ProtectStatus,omitnil" name:"ProtectStatus"`

	// 防护开关
	ProtectSwitch *uint64 `json:"ProtectSwitch,omitnil" name:"ProtectSwitch"`

	// 自动恢复开关
	AutoRestoreSwitchStatus *uint64 `json:"AutoRestoreSwitchStatus,omitnil" name:"AutoRestoreSwitchStatus"`

	// 服务器唯一ID
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 是否已经授权
	Authorization *bool `json:"Authorization,omitnil" name:"Authorization"`

	// 异常状态
	Exception *uint64 `json:"Exception,omitnil" name:"Exception"`

	// 过渡进度
	Progress *uint64 `json:"Progress,omitnil" name:"Progress"`

	// 异常信息
	ExceptionMessage *string `json:"ExceptionMessage,omitnil" name:"ExceptionMessage"`

	// 主机额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`
}

type ProtectEventLists struct {
	// 服务器名称
	HostName *string `json:"HostName,omitnil" name:"HostName"`

	// 服务器ip
	HostIp *string `json:"HostIp,omitnil" name:"HostIp"`

	// 事件地址
	EventDir *string `json:"EventDir,omitnil" name:"EventDir"`

	// 事件类型 0-内容被修改恢复；1-权限被修改恢复；2-归属被修改恢复；3-被删除恢复；4-新增删除
	EventType *uint64 `json:"EventType,omitnil" name:"EventType"`

	// 事件状态 1 已恢复 0 未恢复
	EventStatus *uint64 `json:"EventStatus,omitnil" name:"EventStatus"`

	// 发现时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 恢复时间
	RestoreTime *string `json:"RestoreTime,omitnil" name:"RestoreTime"`

	// 唯一ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 文件类型 0-常规文件；1-目录；2-软链
	FileType *uint64 `json:"FileType,omitnil" name:"FileType"`

	// 主机额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`

	// 机器实例uuid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`
}

type ProtectHostConfig struct {
	// 机器唯一ID
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 防护开关 0  关闭 1开启
	ProtectSwitch *uint64 `json:"ProtectSwitch,omitnil" name:"ProtectSwitch"`

	// 自动恢复开关 0 关闭 1开启
	AutoRecovery *uint64 `json:"AutoRecovery,omitnil" name:"AutoRecovery"`
}

type ProtectMachine struct {
	// 机器名称
	HostName *string `json:"HostName,omitnil" name:"HostName"`

	// 机器IP
	HostIp *string `json:"HostIp,omitnil" name:"HostIp"`

	// 防护目录数
	SafeguardDirNum *uint64 `json:"SafeguardDirNum,omitnil" name:"SafeguardDirNum"`
}

type ProtectMachineInfo struct {
	// 机器名称
	HostName *string `json:"HostName,omitnil" name:"HostName"`

	// 机器IP
	HostIp *string `json:"HostIp,omitnil" name:"HostIp"`

	// 开通时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 到期时间
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`
}

type ProtectNetInfo struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 护网天数
	ProtectDays *uint64 `json:"ProtectDays,omitnil" name:"ProtectDays"`

	// 护网状态 0未启动，1护网中，2已完成
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 护网启动时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 护网完成时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 报告下载地址
	ReportPath *string `json:"ReportPath,omitnil" name:"ReportPath"`
}

type ProtectStat struct {
	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 数量
	Num *uint64 `json:"Num,omitnil" name:"Num"`
}

type RecordInfo struct {
	// 主机ip
	HostIP *string `json:"HostIP,omitnil" name:"HostIP"`

	// 主机实例id
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`

	// 客户端离线时间
	OfflineTime *string `json:"OfflineTime,omitnil" name:"OfflineTime"`

	// 客户端卸载时间
	UninstallTime *string `json:"UninstallTime,omitnil" name:"UninstallTime"`

	// 客户端卸载调用链
	UninstallCmd *string `json:"UninstallCmd,omitnil" name:"UninstallCmd"`

	// 客户端uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

// Predefined struct for user
type RecoverMalwaresRequestParams struct {
	// 木马Id数组（最大100条）
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

type RecoverMalwaresRequest struct {
	*tchttp.BaseRequest
	
	// 木马Id数组（最大100条）
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

func (r *RecoverMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecoverMalwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecoverMalwaresResponseParams struct {
	// 恢复成功id数组，若无则返回空数组
	SuccessIds []*uint64 `json:"SuccessIds,omitnil" name:"SuccessIds"`

	// 恢复失败id数组，若无则返回空数组
	FailedIds []*uint64 `json:"FailedIds,omitnil" name:"FailedIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecoverMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *RecoverMalwaresResponseParams `json:"Response"`
}

func (r *RecoverMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {
	// 地域标志，如 ap-guangzhou，ap-shanghai，ap-beijing
	Region *string `json:"Region,omitnil" name:"Region"`

	// 地域中文名，如华南地区（广州），华东地区（上海金融），华北地区（北京）
	RegionName *string `json:"RegionName,omitnil" name:"RegionName"`

	// 地域ID
	RegionId *uint64 `json:"RegionId,omitnil" name:"RegionId"`

	// 地域代码，如 gz，sh，bj
	RegionCode *string `json:"RegionCode,omitnil" name:"RegionCode"`

	// 地域英文名
	RegionNameEn *string `json:"RegionNameEn,omitnil" name:"RegionNameEn"`
}

type RegionSet struct {
	// 地域名称
	RegionName *string `json:"RegionName,omitnil" name:"RegionName"`

	// 可用区信息
	ZoneSet []*ZoneInfo `json:"ZoneSet,omitnil" name:"ZoneSet"`
}

type ReverseShell struct {
	// ID 主键
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 主机安全UUID
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 主机ID
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机内网IP
	Hostip *string `json:"Hostip,omitnil" name:"Hostip"`

	// 目标IP
	DstIp *string `json:"DstIp,omitnil" name:"DstIp"`

	// 目标端口
	DstPort *uint64 `json:"DstPort,omitnil" name:"DstPort"`

	// 进程名
	ProcessName *string `json:"ProcessName,omitnil" name:"ProcessName"`

	// 进程路径
	FullPath *string `json:"FullPath,omitnil" name:"FullPath"`

	// 命令详情
	CmdLine *string `json:"CmdLine,omitnil" name:"CmdLine"`

	// 执行用户
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 执行用户组
	UserGroup *string `json:"UserGroup,omitnil" name:"UserGroup"`

	// 父进程名
	ParentProcName *string `json:"ParentProcName,omitnil" name:"ParentProcName"`

	// 父进程用户
	ParentProcUser *string `json:"ParentProcUser,omitnil" name:"ParentProcUser"`

	// 父进程用户组
	ParentProcGroup *string `json:"ParentProcGroup,omitnil" name:"ParentProcGroup"`

	// 父进程路径
	ParentProcPath *string `json:"ParentProcPath,omitnil" name:"ParentProcPath"`

	// 处理状态：0-待处理 2-白名单 3-已处理 4-已忽略
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 产生时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 主机名
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 进程树
	ProcTree *string `json:"ProcTree,omitnil" name:"ProcTree"`

	// 检测方法
	DetectBy *uint64 `json:"DetectBy,omitnil" name:"DetectBy"`

	//  主机额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`

	// 进程id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pid *int64 `json:"Pid,omitnil" name:"Pid"`
}

type ReverseShellEventInfo struct {
	// ID 主键
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 云镜UUID
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 主机ID
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机内网IP
	HostIp *string `json:"HostIp,omitnil" name:"HostIp"`

	// 目标IP
	DstIp *string `json:"DstIp,omitnil" name:"DstIp"`

	// 目标端口
	DstPort *uint64 `json:"DstPort,omitnil" name:"DstPort"`

	// 进程名
	ProcessName *string `json:"ProcessName,omitnil" name:"ProcessName"`

	// 进程路径
	FullPath *string `json:"FullPath,omitnil" name:"FullPath"`

	// 命令详情
	CmdLine *string `json:"CmdLine,omitnil" name:"CmdLine"`

	// 执行用户
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 执行用户组
	UserGroup *string `json:"UserGroup,omitnil" name:"UserGroup"`

	// 父进程名
	ParentProcName *string `json:"ParentProcName,omitnil" name:"ParentProcName"`

	// 父进程用户
	ParentProcUser *string `json:"ParentProcUser,omitnil" name:"ParentProcUser"`

	// 父进程用户组
	ParentProcGroup *string `json:"ParentProcGroup,omitnil" name:"ParentProcGroup"`

	// 父进程路径
	ParentProcPath *string `json:"ParentProcPath,omitnil" name:"ParentProcPath"`

	// 处理状态：0-待处理 2-白名单 3-已处理 4-已忽略
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 产生时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 主机名
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 检测方法
	DetectBy *uint64 `json:"DetectBy,omitnil" name:"DetectBy"`

	// 进程树 json  pid:进程id，exe:文件路径 ，account:进程所属用组和用户 ,cmdline:执行命令，ssh_service: SSH服务ip, ssh_soure:登录源
	// 注意：此字段可能返回 null，表示取不到有效值。
	PsTree *string `json:"PsTree,omitnil" name:"PsTree"`

	// 建议方案
	SuggestScheme *string `json:"SuggestScheme,omitnil" name:"SuggestScheme"`

	// 描述
	HarmDescribe *string `json:"HarmDescribe,omitnil" name:"HarmDescribe"`

	// 标签
	Tags []*string `json:"Tags,omitnil" name:"Tags"`

	// 参考链接
	References []*string `json:"References,omitnil" name:"References"`

	// 主机外网ip
	MachineWanIp *string `json:"MachineWanIp,omitnil" name:"MachineWanIp"`

	// 主机在线状态 OFFLINE  ONLINE
	MachineStatus *string `json:"MachineStatus,omitnil" name:"MachineStatus"`

	// 处理时间
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`
}

type ReverseShellRule struct {
	// 规则ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 客户端ID
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 进程名称
	ProcessName *string `json:"ProcessName,omitnil" name:"ProcessName"`

	// 目标IP
	DestIp *string `json:"DestIp,omitnil" name:"DestIp"`

	// 目标端口
	DestPort *string `json:"DestPort,omitnil" name:"DestPort"`

	// 操作人
	Operator *string `json:"Operator,omitnil" name:"Operator"`

	// 是否全局规则
	IsGlobal *uint64 `json:"IsGlobal,omitnil" name:"IsGlobal"`

	// 状态 (0: 有效 1: 无效)
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// 主机IP
	Hostip *string `json:"Hostip,omitnil" name:"Hostip"`
}

type RiskDnsEvent struct {
	// 事件Id
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// 策略ID
	PolicyId *int64 `json:"PolicyId,omitnil" name:"PolicyId"`

	// 命中策略类型[-1:未知|0系统|1:用户]
	PolicyType *int64 `json:"PolicyType,omitnil" name:"PolicyType"`

	// 命中策略名称
	PolicyName *string `json:"PolicyName,omitnil" name:"PolicyName"`

	// 保护级别[0:基础版|1:专业版|2:旗舰版]
	ProtectLevel *int64 `json:"ProtectLevel,omitnil" name:"ProtectLevel"`

	// 主机ID
	HostId *string `json:"HostId,omitnil" name:"HostId"`

	// 主机名称
	HostName *string `json:"HostName,omitnil" name:"HostName"`

	// 主机IP
	HostIp *string `json:"HostIp,omitnil" name:"HostIp"`

	// 外网IP
	WanIp *string `json:"WanIp,omitnil" name:"WanIp"`

	// 客户端ID
	AgentId *string `json:"AgentId,omitnil" name:"AgentId"`

	// 访问域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 标签特性
	Tags []*string `json:"Tags,omitnil" name:"Tags"`

	// 访问次数
	AccessCount *int64 `json:"AccessCount,omitnil" name:"AccessCount"`

	// 威胁描述
	ThreatDesc *string `json:"ThreatDesc,omitnil" name:"ThreatDesc"`

	// 修复方案
	SuggestSolution *string `json:"SuggestSolution,omitnil" name:"SuggestSolution"`

	// 参考链接
	ReferenceLink *string `json:"ReferenceLink,omitnil" name:"ReferenceLink"`

	// 处理状态；[0:待处理|2:已加白|3:非信任状态|4:已处理|5:已忽略]
	HandleStatus *int64 `json:"HandleStatus,omitnil" name:"HandleStatus"`

	// 进程ID
	Pid *int64 `json:"Pid,omitnil" name:"Pid"`

	// 进程名
	ProcessName *string `json:"ProcessName,omitnil" name:"ProcessName"`

	// 进程MD5
	ProcessMd5 *string `json:"ProcessMd5,omitnil" name:"ProcessMd5"`

	// 命令行
	CmdLine *string `json:"CmdLine,omitnil" name:"CmdLine"`

	// 首次访问时间
	FirstTime *string `json:"FirstTime,omitnil" name:"FirstTime"`

	// 最近访问时间
	LastTime *string `json:"LastTime,omitnil" name:"LastTime"`

	// 主机在线状态[OFFLINE:离线|ONLINE:在线|UNKNOWN:未知]
	HostStatus *string `json:"HostStatus,omitnil" name:"HostStatus"`

	// 附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`
}

type RiskDnsList struct {
	// 对外访问域名
	Url *string `json:"Url,omitnil" name:"Url"`

	// 访问次数
	AccessCount *uint64 `json:"AccessCount,omitnil" name:"AccessCount"`

	// 进程名
	ProcessName *string `json:"ProcessName,omitnil" name:"ProcessName"`

	// 进程MD5
	ProcessMd5 *string `json:"ProcessMd5,omitnil" name:"ProcessMd5"`

	// 是否为全局规则，0否，1是
	GlobalRuleId *uint64 `json:"GlobalRuleId,omitnil" name:"GlobalRuleId"`

	// 用户规则id
	UserRuleId *uint64 `json:"UserRuleId,omitnil" name:"UserRuleId"`

	// 状态；0-待处理，2-已加白，3-非信任状态，4-已处理，5-已忽略
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 首次访问时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 最近访问时间
	MergeTime *string `json:"MergeTime,omitnil" name:"MergeTime"`

	// 唯一 Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机ip
	HostIp *string `json:"HostIp,omitnil" name:"HostIp"`

	// 别名
	Alias *string `json:"Alias,omitnil" name:"Alias"`

	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 唯一ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 参考
	Reference *string `json:"Reference,omitnil" name:"Reference"`

	// 命令行
	CmdLine *string `json:"CmdLine,omitnil" name:"CmdLine"`

	// 进程号
	Pid *uint64 `json:"Pid,omitnil" name:"Pid"`

	// 唯一UUID
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 建议方案
	SuggestScheme *string `json:"SuggestScheme,omitnil" name:"SuggestScheme"`

	// 标签特性
	Tags []*string `json:"Tags,omitnil" name:"Tags"`

	// 外网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineWanIp *string `json:"MachineWanIp,omitnil" name:"MachineWanIp"`

	// 主机在线状态[OFFLINE:离线|ONLINE:在线|UNKNOWN:未知]
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineStatus *string `json:"MachineStatus,omitnil" name:"MachineStatus"`
}

// Predefined struct for user
type ScanAssetRequestParams struct {
	// 资产指纹类型id列表
	AssetTypeIds []*uint64 `json:"AssetTypeIds,omitnil" name:"AssetTypeIds"`

	// Quuid列表
	Quuids []*string `json:"Quuids,omitnil" name:"Quuids"`
}

type ScanAssetRequest struct {
	*tchttp.BaseRequest
	
	// 资产指纹类型id列表
	AssetTypeIds []*uint64 `json:"AssetTypeIds,omitnil" name:"AssetTypeIds"`

	// Quuid列表
	Quuids []*string `json:"Quuids,omitnil" name:"Quuids"`
}

func (r *ScanAssetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanAssetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetTypeIds")
	delete(f, "Quuids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScanAssetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScanAssetResponseParams struct {
	// 任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ScanAssetResponse struct {
	*tchttp.BaseResponse
	Response *ScanAssetResponseParams `json:"Response"`
}

func (r *ScanAssetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanTaskDetails struct {
	// 服务器IP
	HostIp *string `json:"HostIp,omitnil" name:"HostIp"`

	// 服务器名称
	HostName *string `json:"HostName,omitnil" name:"HostName"`

	// 操作系统
	OsName *string `json:"OsName,omitnil" name:"OsName"`

	// 风险数量
	RiskNum *uint64 `json:"RiskNum,omitnil" name:"RiskNum"`

	// 扫描开始时间
	ScanBeginTime *string `json:"ScanBeginTime,omitnil" name:"ScanBeginTime"`

	// 扫描结束时间
	ScanEndTime *string `json:"ScanEndTime,omitnil" name:"ScanEndTime"`

	// 唯一Uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 唯一Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 状态码
	Status *string `json:"Status,omitnil" name:"Status"`

	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// id唯一
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 失败详情
	FailType *uint64 `json:"FailType,omitnil" name:"FailType"`

	// 外网ip
	MachineWanIp *string `json:"MachineWanIp,omitnil" name:"MachineWanIp"`

	// 附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`
}

// Predefined struct for user
type ScanVulAgainRequestParams struct {
	// 漏洞事件id串，多个用英文逗号分隔
	EventIds *string `json:"EventIds,omitnil" name:"EventIds"`

	// 重新检查的机器uuid,多个逗号分隔
	Uuids *string `json:"Uuids,omitnil" name:"Uuids"`
}

type ScanVulAgainRequest struct {
	*tchttp.BaseRequest
	
	// 漏洞事件id串，多个用英文逗号分隔
	EventIds *string `json:"EventIds,omitnil" name:"EventIds"`

	// 重新检查的机器uuid,多个逗号分隔
	Uuids *string `json:"Uuids,omitnil" name:"Uuids"`
}

func (r *ScanVulAgainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanVulAgainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventIds")
	delete(f, "Uuids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScanVulAgainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScanVulAgainResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ScanVulAgainResponse struct {
	*tchttp.BaseResponse
	Response *ScanVulAgainResponseParams `json:"Response"`
}

func (r *ScanVulAgainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanVulAgainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScanVulRequestParams struct {
	// 危害等级：1-低危；2-中危；3-高危；4-严重 (多选英文;分隔)
	VulLevels *string `json:"VulLevels,omitnil" name:"VulLevels"`

	// 服务器分类：1:专业版服务器；2:自选服务器
	HostType *uint64 `json:"HostType,omitnil" name:"HostType"`

	// 漏洞类型：1: web-cms漏洞 2:应用漏洞  4: Linux软件漏洞 5: Windows系统漏洞 (多选英文;分隔)
	VulCategories *string `json:"VulCategories,omitnil" name:"VulCategories"`

	// 自选服务器时生效，主机quuid的string数组
	QuuidList []*string `json:"QuuidList,omitnil" name:"QuuidList"`

	// 是否是应急漏洞 0 否 1 是
	VulEmergency *uint64 `json:"VulEmergency,omitnil" name:"VulEmergency"`

	// 超时时长 单位秒 默认 3600 秒
	TimeoutPeriod *uint64 `json:"TimeoutPeriod,omitnil" name:"TimeoutPeriod"`

	// 需要扫描的漏洞id
	VulIds []*uint64 `json:"VulIds,omitnil" name:"VulIds"`
}

type ScanVulRequest struct {
	*tchttp.BaseRequest
	
	// 危害等级：1-低危；2-中危；3-高危；4-严重 (多选英文;分隔)
	VulLevels *string `json:"VulLevels,omitnil" name:"VulLevels"`

	// 服务器分类：1:专业版服务器；2:自选服务器
	HostType *uint64 `json:"HostType,omitnil" name:"HostType"`

	// 漏洞类型：1: web-cms漏洞 2:应用漏洞  4: Linux软件漏洞 5: Windows系统漏洞 (多选英文;分隔)
	VulCategories *string `json:"VulCategories,omitnil" name:"VulCategories"`

	// 自选服务器时生效，主机quuid的string数组
	QuuidList []*string `json:"QuuidList,omitnil" name:"QuuidList"`

	// 是否是应急漏洞 0 否 1 是
	VulEmergency *uint64 `json:"VulEmergency,omitnil" name:"VulEmergency"`

	// 超时时长 单位秒 默认 3600 秒
	TimeoutPeriod *uint64 `json:"TimeoutPeriod,omitnil" name:"TimeoutPeriod"`

	// 需要扫描的漏洞id
	VulIds []*uint64 `json:"VulIds,omitnil" name:"VulIds"`
}

func (r *ScanVulRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanVulRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VulLevels")
	delete(f, "HostType")
	delete(f, "VulCategories")
	delete(f, "QuuidList")
	delete(f, "VulEmergency")
	delete(f, "TimeoutPeriod")
	delete(f, "VulIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScanVulRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScanVulResponseParams struct {
	// 任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ScanVulResponse struct {
	*tchttp.BaseResponse
	Response *ScanVulResponseParams `json:"Response"`
}

func (r *ScanVulResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanVulResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScanVulSettingRequestParams struct {
	// 定期检测间隔时间（天）
	TimerInterval *uint64 `json:"TimerInterval,omitnil" name:"TimerInterval"`

	// 漏洞类型：1: web-cms漏洞 2:应用漏洞  4: Linux软件漏洞 5: Windows系统漏洞, 以数组方式传参[1,2]
	VulCategories []*uint64 `json:"VulCategories,omitnil" name:"VulCategories"`

	// 危害等级：1-低危；2-中危；3-高危；4-严重,以数组方式传参[1,2,3]
	VulLevels []*uint64 `json:"VulLevels,omitnil" name:"VulLevels"`

	// 定期检测时间，如：02:10:50
	TimerTime *string `json:"TimerTime,omitnil" name:"TimerTime"`

	// 是否是应急漏洞 0 否 1 是
	VulEmergency *uint64 `json:"VulEmergency,omitnil" name:"VulEmergency"`

	// 扫描开始时间，如：00:00
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 扫描结束时间，如：08:00
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 是否开启扫描 1开启 0不开启
	EnableScan *uint64 `json:"EnableScan,omitnil" name:"EnableScan"`

	// 为空默认扫描全部专业版、旗舰版、普惠版主机，不为空只扫描选中主机
	Uuids []*string `json:"Uuids,omitnil" name:"Uuids"`
}

type ScanVulSettingRequest struct {
	*tchttp.BaseRequest
	
	// 定期检测间隔时间（天）
	TimerInterval *uint64 `json:"TimerInterval,omitnil" name:"TimerInterval"`

	// 漏洞类型：1: web-cms漏洞 2:应用漏洞  4: Linux软件漏洞 5: Windows系统漏洞, 以数组方式传参[1,2]
	VulCategories []*uint64 `json:"VulCategories,omitnil" name:"VulCategories"`

	// 危害等级：1-低危；2-中危；3-高危；4-严重,以数组方式传参[1,2,3]
	VulLevels []*uint64 `json:"VulLevels,omitnil" name:"VulLevels"`

	// 定期检测时间，如：02:10:50
	TimerTime *string `json:"TimerTime,omitnil" name:"TimerTime"`

	// 是否是应急漏洞 0 否 1 是
	VulEmergency *uint64 `json:"VulEmergency,omitnil" name:"VulEmergency"`

	// 扫描开始时间，如：00:00
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 扫描结束时间，如：08:00
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 是否开启扫描 1开启 0不开启
	EnableScan *uint64 `json:"EnableScan,omitnil" name:"EnableScan"`

	// 为空默认扫描全部专业版、旗舰版、普惠版主机，不为空只扫描选中主机
	Uuids []*string `json:"Uuids,omitnil" name:"Uuids"`
}

func (r *ScanVulSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanVulSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimerInterval")
	delete(f, "VulCategories")
	delete(f, "VulLevels")
	delete(f, "TimerTime")
	delete(f, "VulEmergency")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "EnableScan")
	delete(f, "Uuids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScanVulSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScanVulSettingResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ScanVulSettingResponse struct {
	*tchttp.BaseResponse
	Response *ScanVulSettingResponseParams `json:"Response"`
}

func (r *ScanVulSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanVulSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchTemplate struct {
	// 检索名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 检索索引类型
	LogType *string `json:"LogType,omitnil" name:"LogType"`

	// 检索语句
	Condition *string `json:"Condition,omitnil" name:"Condition"`

	// 时间范围
	TimeRange *string `json:"TimeRange,omitnil" name:"TimeRange"`

	// 转换的检索语句内容
	Query *string `json:"Query,omitnil" name:"Query"`

	// 检索方式。输入框检索：standard,过滤，检索：simple
	Flag *string `json:"Flag,omitnil" name:"Flag"`

	// 展示数据
	DisplayData *string `json:"DisplayData,omitnil" name:"DisplayData"`

	// 规则ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

type SecurityButlerInfo struct {
	// 数据id
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 订单id
	OrderId *uint64 `json:"OrderId,omitnil" name:"OrderId"`

	// cvm id
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 服务状态 0-服务中,1-已到期 2已销毁
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 服务开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 服务结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 主机名称
	HostName *string `json:"HostName,omitnil" name:"HostName"`

	// 主机Ip
	HostIp *string `json:"HostIp,omitnil" name:"HostIp"`

	// 主机 uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 主机风险数
	RiskCount *uint64 `json:"RiskCount,omitnil" name:"RiskCount"`
}

type SecurityDynamic struct {
	// 主机安全客户端UUID。
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 安全事件发生时间。
	EventTime *string `json:"EventTime,omitnil" name:"EventTime"`

	// 安全事件类型。
	// <li>MALWARE：木马事件</li>
	// <li>NON_LOCAL_LOGIN：异地登录</li>
	// <li>BRUTEATTACK_SUCCESS：密码破解成功</li>
	// <li>VUL：漏洞</li>
	// <li>BASELINE：安全基线</li>
	EventType *string `json:"EventType,omitnil" name:"EventType"`

	// 安全事件消息。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 安全事件等级。
	// <li>RISK: 严重</li>
	// <li>HIGH: 高危</li>
	// <li>NORMAL: 中危</li>
	// <li>LOW: 低危</li>
	// <li>UNKNOWNED: 可疑</li>
	SecurityLevel *string `json:"SecurityLevel,omitnil" name:"SecurityLevel"`
}

type SecurityEventInfo struct {
	// 安全事件数
	EventCnt *uint64 `json:"EventCnt,omitnil" name:"EventCnt"`

	// 受影响机器数
	UuidCnt *uint64 `json:"UuidCnt,omitnil" name:"UuidCnt"`
}

type SecurityTrend struct {
	// 事件时间。
	Date *string `json:"Date,omitnil" name:"Date"`

	// 事件数量。
	EventNum *uint64 `json:"EventNum,omitnil" name:"EventNum"`
}

// Predefined struct for user
type SeparateMalwaresRequestParams struct {
	// 木马事件ID数组。(最大100条)
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`

	// 是否杀掉进程
	KillProcess *bool `json:"KillProcess,omitnil" name:"KillProcess"`
}

type SeparateMalwaresRequest struct {
	*tchttp.BaseRequest
	
	// 木马事件ID数组。(最大100条)
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`

	// 是否杀掉进程
	KillProcess *bool `json:"KillProcess,omitnil" name:"KillProcess"`
}

func (r *SeparateMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SeparateMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	delete(f, "KillProcess")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SeparateMalwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SeparateMalwaresResponseParams struct {
	// 隔离成功的id数组，若无则返回空数组
	SuccessIds []*uint64 `json:"SuccessIds,omitnil" name:"SuccessIds"`

	// 隔离失败的id数组，若无则返回空数组
	FailedIds []*uint64 `json:"FailedIds,omitnil" name:"FailedIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SeparateMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *SeparateMalwaresResponseParams `json:"Response"`
}

func (r *SeparateMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SeparateMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetBashEventsStatusRequestParams struct {
	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`

	// 新状态(0-待处理 1-高危 2-正常)
	Status *uint64 `json:"Status,omitnil" name:"Status"`
}

type SetBashEventsStatusRequest struct {
	*tchttp.BaseRequest
	
	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`

	// 新状态(0-待处理 1-高危 2-正常)
	Status *uint64 `json:"Status,omitnil" name:"Status"`
}

func (r *SetBashEventsStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetBashEventsStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetBashEventsStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetBashEventsStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SetBashEventsStatusResponse struct {
	*tchttp.BaseResponse
	Response *SetBashEventsStatusResponseParams `json:"Response"`
}

func (r *SetBashEventsStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetBashEventsStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StandardModeConfig struct {
	// 阻断时长，单位：秒
	Ttl *uint64 `json:"Ttl,omitnil" name:"Ttl"`
}

// Predefined struct for user
type StartBaselineDetectRequestParams struct {
	// 基线检测参数
	Param *BaselineDetectParam `json:"Param,omitnil" name:"Param"`
}

type StartBaselineDetectRequest struct {
	*tchttp.BaseRequest
	
	// 基线检测参数
	Param *BaselineDetectParam `json:"Param,omitnil" name:"Param"`
}

func (r *StartBaselineDetectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartBaselineDetectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Param")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartBaselineDetectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartBaselineDetectResponseParams struct {
	// 扫描任务ID
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StartBaselineDetectResponse struct {
	*tchttp.BaseResponse
	Response *StartBaselineDetectResponseParams `json:"Response"`
}

func (r *StartBaselineDetectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartBaselineDetectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopBaselineDetectRequestParams struct {
	// 取消任务ID集合
	TaskIds []*int64 `json:"TaskIds,omitnil" name:"TaskIds"`
}

type StopBaselineDetectRequest struct {
	*tchttp.BaseRequest
	
	// 取消任务ID集合
	TaskIds []*int64 `json:"TaskIds,omitnil" name:"TaskIds"`
}

func (r *StopBaselineDetectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopBaselineDetectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopBaselineDetectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopBaselineDetectResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StopBaselineDetectResponse struct {
	*tchttp.BaseResponse
	Response *StopBaselineDetectResponseParams `json:"Response"`
}

func (r *StopBaselineDetectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopBaselineDetectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopNoticeBanTipsRequestParams struct {

}

type StopNoticeBanTipsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *StopNoticeBanTipsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopNoticeBanTipsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopNoticeBanTipsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopNoticeBanTipsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StopNoticeBanTipsResponse struct {
	*tchttp.BaseResponse
	Response *StopNoticeBanTipsResponseParams `json:"Response"`
}

func (r *StopNoticeBanTipsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopNoticeBanTipsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Strategy struct {
	// 策略名
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrategyName *string `json:"StrategyName,omitnil" name:"StrategyName"`

	// 策略id
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrategyId *uint64 `json:"StrategyId,omitnil" name:"StrategyId"`

	// 基线检测项总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleCount *uint64 `json:"RuleCount,omitnil" name:"RuleCount"`

	// 主机数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostCount *uint64 `json:"HostCount,omitnil" name:"HostCount"`

	// 扫描周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanCycle *uint64 `json:"ScanCycle,omitnil" name:"ScanCycle"`

	// 扫描时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanAt *string `json:"ScanAt,omitnil" name:"ScanAt"`

	// 是否可用
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enabled *uint64 `json:"Enabled,omitnil" name:"Enabled"`

	// 通过率
	// 注意：此字段可能返回 null，表示取不到有效值。
	PassRate *uint64 `json:"PassRate,omitnil" name:"PassRate"`

	// 基线id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryIds *string `json:"CategoryIds,omitnil" name:"CategoryIds"`

	// 是否默认策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDefault *uint64 `json:"IsDefault,omitnil" name:"IsDefault"`
}

// Predefined struct for user
type SwitchBashRulesRequestParams struct {
	// 规则ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 是否禁用
	Disabled *uint64 `json:"Disabled,omitnil" name:"Disabled"`
}

type SwitchBashRulesRequest struct {
	*tchttp.BaseRequest
	
	// 规则ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 是否禁用
	Disabled *uint64 `json:"Disabled,omitnil" name:"Disabled"`
}

func (r *SwitchBashRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchBashRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Disabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchBashRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchBashRulesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SwitchBashRulesResponse struct {
	*tchttp.BaseResponse
	Response *SwitchBashRulesResponseParams `json:"Response"`
}

func (r *SwitchBashRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchBashRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncAssetScanRequestParams struct {
	// 是否同步：true-是 false-否；默认false
	Sync *bool `json:"Sync,omitnil" name:"Sync"`
}

type SyncAssetScanRequest struct {
	*tchttp.BaseRequest
	
	// 是否同步：true-是 false-否；默认false
	Sync *bool `json:"Sync,omitnil" name:"Sync"`
}

func (r *SyncAssetScanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncAssetScanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Sync")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncAssetScanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncAssetScanResponseParams struct {
	// 枚举值有(大写)：NOTASK（没有同步任务），SYNCING（同步中），FINISHED（同步完成）
	State *string `json:"State,omitnil" name:"State"`

	// 最新开始同步时间
	LatestStartTime *string `json:"LatestStartTime,omitnil" name:"LatestStartTime"`

	// 最新结束同步时间
	LatestEndTime *string `json:"LatestEndTime,omitnil" name:"LatestEndTime"`

	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SyncAssetScanResponse struct {
	*tchttp.BaseResponse
	Response *SyncAssetScanResponseParams `json:"Response"`
}

func (r *SyncAssetScanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncAssetScanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncBaselineDetectSummaryRequestParams struct {

}

type SyncBaselineDetectSummaryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *SyncBaselineDetectSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncBaselineDetectSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncBaselineDetectSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncBaselineDetectSummaryResponseParams struct {
	// 处理进度
	ProgressRate *int64 `json:"ProgressRate,omitnil" name:"ProgressRate"`

	// 未通过策略总数
	NotPassPolicyCount *int64 `json:"NotPassPolicyCount,omitnil" name:"NotPassPolicyCount"`

	// 主机总数
	HostCount *int64 `json:"HostCount,omitnil" name:"HostCount"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 1:即将进行首次扫描   0:已经扫描过了
	WillFirstScan *int64 `json:"WillFirstScan,omitnil" name:"WillFirstScan"`

	// 正在检测的任务ID
	DetectingTaskIds []*int64 `json:"DetectingTaskIds,omitnil" name:"DetectingTaskIds"`

	// 扫描中剩余时间(分钟)
	LeftMins *int64 `json:"LeftMins,omitnil" name:"LeftMins"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SyncBaselineDetectSummaryResponse struct {
	*tchttp.BaseResponse
	Response *SyncBaselineDetectSummaryResponseParams `json:"Response"`
}

func (r *SyncBaselineDetectSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncBaselineDetectSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// 标签ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 标签名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 服务器数
	Count *uint64 `json:"Count,omitnil" name:"Count"`
}

type TagMachine struct {
	// ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// 主机ID
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机名称
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// 主机内网IP
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// 主机外网IP
	MachineWanIp *string `json:"MachineWanIp,omitnil" name:"MachineWanIp"`

	// 主机区域
	MachineRegion *string `json:"MachineRegion,omitnil" name:"MachineRegion"`

	// 主机区域类型
	MachineType *string `json:"MachineType,omitnil" name:"MachineType"`
}

type Tags struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}

type TaskStatus struct {
	// 扫描中（包含初始化）
	Scanning *string `json:"Scanning,omitnil" name:"Scanning"`

	// 扫描终止（包含终止中）
	Ok *string `json:"Ok,omitnil" name:"Ok"`

	// 扫描失败
	Fail *string `json:"Fail,omitnil" name:"Fail"`

	// 扫描失败（提示具体原因：扫描超时、客户端版本低、客户端离线）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Stop *string `json:"Stop,omitnil" name:"Stop"`
}

// Predefined struct for user
type TrustMalwaresRequestParams struct {
	// 木马ID数组（单次不超过的最大条数：100）
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

type TrustMalwaresRequest struct {
	*tchttp.BaseRequest
	
	// 木马ID数组（单次不超过的最大条数：100）
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

func (r *TrustMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TrustMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TrustMalwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TrustMalwaresResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TrustMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *TrustMalwaresResponseParams `json:"Response"`
}

func (r *TrustMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TrustMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UntrustMalwaresRequestParams struct {
	// 木马ID数组 (最大100条)
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

type UntrustMalwaresRequest struct {
	*tchttp.BaseRequest
	
	// 木马ID数组 (最大100条)
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

func (r *UntrustMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UntrustMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UntrustMalwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UntrustMalwaresResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UntrustMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *UntrustMalwaresResponseParams `json:"Response"`
}

func (r *UntrustMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UntrustMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateBaselineStrategyRequestParams struct {
	// 策略id
	StrategyId *uint64 `json:"StrategyId,omitnil" name:"StrategyId"`

	// 策略名称
	StrategyName *string `json:"StrategyName,omitnil" name:"StrategyName"`

	// 检测周期
	ScanCycle *uint64 `json:"ScanCycle,omitnil" name:"ScanCycle"`

	// 定期检测时间，该时间下发扫描
	ScanAt *string `json:"ScanAt,omitnil" name:"ScanAt"`

	// 该策略下选择的基线id数组
	CategoryIds []*string `json:"CategoryIds,omitnil" name:"CategoryIds"`

	// 扫描范围是否全部服务器, 1:是  0:否, 为1则为全部专业版主机
	IsGlobal *uint64 `json:"IsGlobal,omitnil" name:"IsGlobal"`

	// 云主机类型：
	// cvm：腾讯云服务器
	// bm：裸金属
	// ecm：边缘计算主机
	// lh:轻量应用服务器
	// other:混合云机器
	MachineType *string `json:"MachineType,omitnil" name:"MachineType"`

	// 主机地域 ap-guangzhou
	RegionCode *string `json:"RegionCode,omitnil" name:"RegionCode"`

	// 主机id数组
	Quuids []*string `json:"Quuids,omitnil" name:"Quuids"`
}

type UpdateBaselineStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 策略id
	StrategyId *uint64 `json:"StrategyId,omitnil" name:"StrategyId"`

	// 策略名称
	StrategyName *string `json:"StrategyName,omitnil" name:"StrategyName"`

	// 检测周期
	ScanCycle *uint64 `json:"ScanCycle,omitnil" name:"ScanCycle"`

	// 定期检测时间，该时间下发扫描
	ScanAt *string `json:"ScanAt,omitnil" name:"ScanAt"`

	// 该策略下选择的基线id数组
	CategoryIds []*string `json:"CategoryIds,omitnil" name:"CategoryIds"`

	// 扫描范围是否全部服务器, 1:是  0:否, 为1则为全部专业版主机
	IsGlobal *uint64 `json:"IsGlobal,omitnil" name:"IsGlobal"`

	// 云主机类型：
	// cvm：腾讯云服务器
	// bm：裸金属
	// ecm：边缘计算主机
	// lh:轻量应用服务器
	// other:混合云机器
	MachineType *string `json:"MachineType,omitnil" name:"MachineType"`

	// 主机地域 ap-guangzhou
	RegionCode *string `json:"RegionCode,omitnil" name:"RegionCode"`

	// 主机id数组
	Quuids []*string `json:"Quuids,omitnil" name:"Quuids"`
}

func (r *UpdateBaselineStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateBaselineStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StrategyId")
	delete(f, "StrategyName")
	delete(f, "ScanCycle")
	delete(f, "ScanAt")
	delete(f, "CategoryIds")
	delete(f, "IsGlobal")
	delete(f, "MachineType")
	delete(f, "RegionCode")
	delete(f, "Quuids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateBaselineStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateBaselineStrategyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateBaselineStrategyResponse struct {
	*tchttp.BaseResponse
	Response *UpdateBaselineStrategyResponseParams `json:"Response"`
}

func (r *UpdateBaselineStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateBaselineStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateMachineTagsRequestParams struct {
	// 机器 Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 标签ID，该操作会覆盖原有的标签列表
	TagIds []*uint64 `json:"TagIds,omitnil" name:"TagIds"`

	// 服务器地区 如: ap-guangzhou
	MachineRegion *string `json:"MachineRegion,omitnil" name:"MachineRegion"`

	// 服务器类型(CVM|BM|ECM|LH|Other)
	MachineArea *string `json:"MachineArea,omitnil" name:"MachineArea"`
}

type UpdateMachineTagsRequest struct {
	*tchttp.BaseRequest
	
	// 机器 Quuid
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 标签ID，该操作会覆盖原有的标签列表
	TagIds []*uint64 `json:"TagIds,omitnil" name:"TagIds"`

	// 服务器地区 如: ap-guangzhou
	MachineRegion *string `json:"MachineRegion,omitnil" name:"MachineRegion"`

	// 服务器类型(CVM|BM|ECM|LH|Other)
	MachineArea *string `json:"MachineArea,omitnil" name:"MachineArea"`
}

func (r *UpdateMachineTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateMachineTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Quuid")
	delete(f, "TagIds")
	delete(f, "MachineRegion")
	delete(f, "MachineArea")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateMachineTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateMachineTagsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateMachineTagsResponse struct {
	*tchttp.BaseResponse
	Response *UpdateMachineTagsResponseParams `json:"Response"`
}

func (r *UpdateMachineTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateMachineTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UsualPlace struct {
	// ID。
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 主机安全客户端唯一标识UUID。
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 国家 ID。
	CountryId *uint64 `json:"CountryId,omitnil" name:"CountryId"`

	// 省份 ID。
	ProvinceId *uint64 `json:"ProvinceId,omitnil" name:"ProvinceId"`

	// 城市 ID。
	CityId *uint64 `json:"CityId,omitnil" name:"CityId"`
}

type VertexDetail struct {
	// 该节点类型，进程:1；网络:2；文件:3；ssh:4
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// 各节点类型用到的时间，2022-11-29 00:00:00 格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *string `json:"Time,omitnil" name:"Time"`

	// 告警信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmInfo []*AlarmInfo `json:"AlarmInfo,omitnil" name:"AlarmInfo"`

	// 进程名，当该节点为进程时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcName *string `json:"ProcName,omitnil" name:"ProcName"`

	// 命令行，当该节点为进程时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	CmdLine *string `json:"CmdLine,omitnil" name:"CmdLine"`

	// 进程id，当该节点为进程时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pid *string `json:"Pid,omitnil" name:"Pid"`

	// 文件md5，当该节点为文件时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileMd5 *string `json:"FileMd5,omitnil" name:"FileMd5"`

	// 文件写入内容，当该节点为文件时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileContent *string `json:"FileContent,omitnil" name:"FileContent"`

	// 文件路径，当该节点为文件时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilePath *string `json:"FilePath,omitnil" name:"FilePath"`

	// 文件创建时间，当该节点为文件时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileCreateTime *string `json:"FileCreateTime,omitnil" name:"FileCreateTime"`

	// 请求目的地址，当该节点为网络时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address *string `json:"Address,omitnil" name:"Address"`

	// 目标端口，当该节点为网络时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	DstPort *uint64 `json:"DstPort,omitnil" name:"DstPort"`

	// 登录源ip，当该节点为ssh时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	SrcIP *string `json:"SrcIP,omitnil" name:"SrcIP"`

	// 登录用户名用户组，当该节点为ssh时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	User *string `json:"User,omitnil" name:"User"`

	// 漏洞名称，当该节点为漏洞时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulName *string `json:"VulName,omitnil" name:"VulName"`

	// 漏洞利用时间，当该节点为漏洞时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulTime *string `json:"VulTime,omitnil" name:"VulTime"`

	// http请求内容，当该节点为漏洞时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpContent *string `json:"HttpContent,omitnil" name:"HttpContent"`

	// 漏洞利用者来源ip，当该节点为漏洞时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulSrcIP *string `json:"VulSrcIP,omitnil" name:"VulSrcIP"`

	// 点id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VertexId *string `json:"VertexId,omitnil" name:"VertexId"`
}

type VertexInfo struct {
	// 该结点类型，进程:1；网络:2；文件:3；ssh:4；
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// 该节点包含的vid
	Vid *string `json:"Vid,omitnil" name:"Vid"`

	// 该节点的父节点vid
	ParentVid *string `json:"ParentVid,omitnil" name:"ParentVid"`

	// 是否叶子
	IsLeaf *bool `json:"IsLeaf,omitnil" name:"IsLeaf"`

	// 进程名，当Type=1时使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcNamePrefix *string `json:"ProcNamePrefix,omitnil" name:"ProcNamePrefix"`

	// 进程名md5，当Type=1时使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcNameMd5 *string `json:"ProcNameMd5,omitnil" name:"ProcNameMd5"`

	// 命令行，当Type=1时使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	CmdLinePrefix *string `json:"CmdLinePrefix,omitnil" name:"CmdLinePrefix"`

	// 命令行md5，当Type=1时使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	CmdLineMd5 *string `json:"CmdLineMd5,omitnil" name:"CmdLineMd5"`

	// 文件路径，当Type=3时使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilePathPrefix *string `json:"FilePathPrefix,omitnil" name:"FilePathPrefix"`

	// 请求目的地址，当Type=2时使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddressPrefix *string `json:"AddressPrefix,omitnil" name:"AddressPrefix"`

	// 是否漏洞节点
	IsWeDetect *bool `json:"IsWeDetect,omitnil" name:"IsWeDetect"`

	// 是否告警节点
	IsAlarm *bool `json:"IsAlarm,omitnil" name:"IsAlarm"`

	// 文件路径md5，当Type=3时使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilePathMd5 *string `json:"FilePathMd5,omitnil" name:"FilePathMd5"`

	// 请求目的地址md5，当Type=2时使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddressMd5 *string `json:"AddressMd5,omitnil" name:"AddressMd5"`
}

type VulDetailInfo struct {
	// 漏洞ID
	VulId *uint64 `json:"VulId,omitnil" name:"VulId"`

	// 漏洞级别
	Level *uint64 `json:"Level,omitnil" name:"Level"`

	// 漏洞名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// cve编号
	CveId *string `json:"CveId,omitnil" name:"CveId"`

	// 1: web-cms漏洞 2:应用漏洞  4: Linux软件漏洞 5: Windows系统漏洞 0= 应急漏洞
	VulCategory *uint64 `json:"VulCategory,omitnil" name:"VulCategory"`

	// 漏洞描述
	Descript *string `json:"Descript,omitnil" name:"Descript"`

	// 修复建议
	Fix *string `json:"Fix,omitnil" name:"Fix"`

	// 参考链接
	Reference *string `json:"Reference,omitnil" name:"Reference"`

	// CVSS评分
	CvssScore *float64 `json:"CvssScore,omitnil" name:"CvssScore"`

	// CVSS详情
	Cvss *string `json:"Cvss,omitnil" name:"Cvss"`

	// 发布时间
	PublishTime *string `json:"PublishTime,omitnil" name:"PublishTime"`
}

type VulEffectHostList struct {
	// 事件id
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventId *uint64 `json:"EventId,omitnil" name:"EventId"`

	// 状态：0: 待处理 1:忽略  3:已修复  5:检测中 6:修复中 7: 回滚中 8:修复失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 最后检测时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastTime *string `json:"LastTime,omitnil" name:"LastTime"`

	// 危害等级：1-低危；2-中危；3-高危；4-严重
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *uint64 `json:"Level,omitnil" name:"Level"`

	// 主机Quuid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// 主机Uuid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 主机HostIp
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostIp *string `json:"HostIp,omitnil" name:"HostIp"`

	// 主机别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AliasName *string `json:"AliasName,omitnil" name:"AliasName"`

	// 主机标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitnil" name:"Tags"`

	// 说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 版本信息：0-基础版 1-专业版 2-旗舰版 3-普惠版
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostVersion *uint64 `json:"HostVersion,omitnil" name:"HostVersion"`

	// 是否能自动修复 0 :漏洞不可自动修复，  1：可自动修复， 2：客户端已离线， 3：主机不是旗舰版只能手动修复， 4：机型不允许 ，5：修复中 ，6：已修复， 7：检测中  9:修复失败，10:已忽略 11:漏洞只支持linux不支持Windows 12：漏洞只支持Windows不支持linux，13:修复失败但此时主机已离线，14:修复失败但此时主机不是旗舰版， 15:已手动修复
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsSupportAutoFix *uint64 `json:"IsSupportAutoFix,omitnil" name:"IsSupportAutoFix"`

	// 失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FixStatusMsg *string `json:"FixStatusMsg,omitnil" name:"FixStatusMsg"`

	// 首次发现时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstDiscoveryTime *string `json:"FirstDiscoveryTime,omitnil" name:"FirstDiscoveryTime"`

	// 实例状态："PENDING"-创建中 "LAUNCH_FAILED"-创建失败 "RUNNING"-运行中 "STOPPED"-关机 "STARTING"-表示开机中 "STOPPING"-表示关机中 "REBOOTING"-重启中 "SHUTDOWN"-表示停止待销毁 "TERMINATING"-表示销毁中 "
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceState *string `json:"InstanceState,omitnil" name:"InstanceState"`

	// 外网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIpAddresses *string `json:"PublicIpAddresses,omitnil" name:"PublicIpAddresses"`

	// 云标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CloudTags []*Tags `json:"CloudTags,omitnil" name:"CloudTags"`

	// 主机额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitnil" name:"MachineExtraInfo"`
}

type VulHostTopInfo struct {
	// 主机名
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostName *string `json:"HostName,omitnil" name:"HostName"`

	// 漏洞等级与数量统计列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulLevelList []*VulLevelCountInfo `json:"VulLevelList,omitnil" name:"VulLevelList"`

	// 主机Quuid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// top评分
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *uint64 `json:"Score,omitnil" name:"Score"`
}

type VulInfoList struct {
	// 漏洞包含的事件id串，多个用“,”分割
	Ids *string `json:"Ids,omitnil" name:"Ids"`

	// 漏洞名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 0: 待处理 1:忽略  3:已修复  5:检测中 6:修复中  8:修复失败
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 漏洞id
	VulId *uint64 `json:"VulId,omitnil" name:"VulId"`

	// 漏洞披露事件
	PublishTime *string `json:"PublishTime,omitnil" name:"PublishTime"`

	// 最后检测时间
	LastTime *string `json:"LastTime,omitnil" name:"LastTime"`

	// 影响主机数
	HostCount *uint64 `json:"HostCount,omitnil" name:"HostCount"`

	// 漏洞等级 1:低 2:中 3:高 4:严重
	Level *uint64 `json:"Level,omitnil" name:"Level"`

	// 废弃字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	From *uint64 `json:"From,omitnil" name:"From"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Descript *string `json:"Descript,omitnil" name:"Descript"`

	// 废弃字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublishTimeWisteria *string `json:"PublishTimeWisteria,omitnil" name:"PublishTimeWisteria"`

	// 废弃字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	NameWisteria *string `json:"NameWisteria,omitnil" name:"NameWisteria"`

	// 废弃字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	DescriptWisteria *string `json:"DescriptWisteria,omitnil" name:"DescriptWisteria"`

	// 聚合后事件状态串
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusStr *string `json:"StatusStr,omitnil" name:"StatusStr"`

	// cve编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	CveId *string `json:"CveId,omitnil" name:"CveId"`

	// CVSS评分
	// 注意：此字段可能返回 null，表示取不到有效值。
	CvssScore *float64 `json:"CvssScore,omitnil" name:"CvssScore"`

	// 漏洞标签 多个逗号分割
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels *string `json:"Labels,omitnil" name:"Labels"`

	// 是否能自动修复且包含能自动修复的主机， 0=否  1=是
	// 注意：此字段可能返回 null，表示取不到有效值。
	FixSwitch *uint64 `json:"FixSwitch,omitnil" name:"FixSwitch"`

	// 最后扫描任务的id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// 是否支持防御， 0:不支持 1:支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsSupportDefense *uint64 `json:"IsSupportDefense,omitnil" name:"IsSupportDefense"`

	// 已防御的攻击次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefenseAttackCount *uint64 `json:"DefenseAttackCount,omitnil" name:"DefenseAttackCount"`

	// 首次出现时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstAppearTime *string `json:"FirstAppearTime,omitnil" name:"FirstAppearTime"`

	// 漏洞类别 1: web-cms漏洞 2:应用漏洞  4: Linux软件漏洞 5: Windows系统漏洞
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulCategory *uint64 `json:"VulCategory,omitnil" name:"VulCategory"`

	// 攻击热度级别
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackLevel *uint64 `json:"AttackLevel,omitnil" name:"AttackLevel"`

	// 漏洞修复后是否需要重启
	// 注意：此字段可能返回 null，表示取不到有效值。
	FixNoNeedRestart *bool `json:"FixNoNeedRestart,omitnil" name:"FixNoNeedRestart"`

	// 检测方式0 - 版本比对, 1 - POC验证
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *uint64 `json:"Method,omitnil" name:"Method"`
}

type VulLevelCountInfo struct {
	// 漏洞等级
	VulLevel *uint64 `json:"VulLevel,omitnil" name:"VulLevel"`

	// 漏洞数量
	VulCount *uint64 `json:"VulCount,omitnil" name:"VulCount"`
}

type VulLevelInfo struct {
	// // 危害等级：1-低危；2-中危；3-高危；4-严重
	VulLevel *uint64 `json:"VulLevel,omitnil" name:"VulLevel"`

	// 数量
	Count *uint64 `json:"Count,omitnil" name:"Count"`
}

type VulStoreListInfo struct {
	// 漏洞ID
	VulId *uint64 `json:"VulId,omitnil" name:"VulId"`

	// 漏洞级别
	Level *uint64 `json:"Level,omitnil" name:"Level"`

	// 漏洞名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// cve编号
	CveId *string `json:"CveId,omitnil" name:"CveId"`

	// 1: web-cms漏洞 2:应用漏洞  4: Linux软件漏洞 5: Windows系统漏洞 0= 应急漏洞
	VulCategory *uint64 `json:"VulCategory,omitnil" name:"VulCategory"`

	// 发布时间
	PublishDate *string `json:"PublishDate,omitnil" name:"PublishDate"`

	// 漏洞检测方法 0 - 版本比对, 1 - POC验证
	Method *uint64 `json:"Method,omitnil" name:"Method"`

	// 漏洞攻击热度
	AttackLevel *uint64 `json:"AttackLevel,omitnil" name:"AttackLevel"`

	// 漏洞是否支持自动修复
	// 0-windows/linux均关闭; 1-windows/linux均打开; 2-仅linux; 3-仅windows
	FixSwitch *uint64 `json:"FixSwitch,omitnil" name:"FixSwitch"`

	// 漏洞是否支持防御
	// 0:不支持 1:支持
	SupportDefense *uint64 `json:"SupportDefense,omitnil" name:"SupportDefense"`
}

type VulTopInfo struct {
	// 漏洞 名
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulName *string `json:"VulName,omitnil" name:"VulName"`

	// 危害等级：1-低危；2-中危；3-高危；4-严重
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulLevel *uint64 `json:"VulLevel,omitnil" name:"VulLevel"`

	// 漏洞数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulCount *uint64 `json:"VulCount,omitnil" name:"VulCount"`

	// 漏洞id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulId *uint64 `json:"VulId,omitnil" name:"VulId"`
}

type WarningInfoObj struct {
	// 事件告警类型；1：离线，2：木马，3：异常登录，4：爆破，5：漏洞（已拆分为9-12四种类型）6：高危命令，7：反弹sell，8：本地提权，9：应用漏洞，10：web-cms漏洞，11：应急漏洞，12：安全基线 ,13: 防篡改，14：恶意请求，15: 网络攻击，16：Windows系统漏洞，17：Linux软件漏洞，18：核心文件监控告警，19：客户端卸载告警。20：客户端离线告警
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 1: 关闭告警 0: 开启告警
	DisablePhoneWarning *uint64 `json:"DisablePhoneWarning,omitnil" name:"DisablePhoneWarning"`

	// 开始时间，格式: HH:mm
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 结束时间，格式: HH:mm
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 时区信息
	TimeZone *string `json:"TimeZone,omitnil" name:"TimeZone"`

	// 漏洞等级控制位（对应DB的十进制存储）
	ControlBit *uint64 `json:"ControlBit,omitnil" name:"ControlBit"`

	// 漏洞等级控制位二进制，每一位对应页面漏洞等级的开启关闭：低中高（0:关闭；1：开启），例如：101 → 同时勾选低+高
	ControlBits *string `json:"ControlBits,omitnil" name:"ControlBits"`

	// 告警主机范围类型，0:全部主机，1:按所属项目选，2:按腾讯云标签选，3:按主机安全标签选，4:自选主机
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostRange *int64 `json:"HostRange,omitnil" name:"HostRange"`

	// 配置的告警范围主机个数，前端用此判断展示提示信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitnil" name:"Count"`
}

type WarningObject struct {
	// 事件告警类型；1：离线，2：木马，3：异常登录，4：爆破，5：漏洞（已拆分为9-12四种类型）6：高位命令，7：反弹sell，8：本地提权，9：系统组件漏洞，10：web应用漏洞，11：应急漏洞，12：安全基线，14：恶意请求，15: 网络攻击，16：Windows系统漏洞，17：Linux软件漏洞
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 1: 关闭告警 0: 开启告警
	DisablePhoneWarning *uint64 `json:"DisablePhoneWarning,omitnil" name:"DisablePhoneWarning"`

	// 开始时间，格式: HH:mm
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 结束时间，格式: HH:mm
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 漏洞等级控制位二进制，每一位对应页面漏洞等级的开启关闭：低中高（0:关闭；1：开启），例如：101 → 同时勾选低+高；01→(登录审计)疑似不告警，高危告警
	ControlBits *string `json:"ControlBits,omitnil" name:"ControlBits"`

	// 告警主机范围类型，0:全部主机，1:按所属项目选，2:按腾讯云标签选，3:按主机安全标签选，4:自选主机
	HostRange *int64 `json:"HostRange,omitnil" name:"HostRange"`
}

type ZoneInfo struct {
	// 可用区名称
	ZoneName *string `json:"ZoneName,omitnil" name:"ZoneName"`
}