// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20191118

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AndroidApp struct {
	// 安卓应用 Id
	AndroidAppId *string `json:"AndroidAppId,omitnil,omitempty" name:"AndroidAppId"`

	// 安卓应用名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 安卓应用状态（上架、下架）
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 安卓应用版本列表
	AndroidAppVersionInfo []*AndroidAppVersionInfo `json:"AndroidAppVersionInfo,omitnil,omitempty" name:"AndroidAppVersionInfo"`

	// 安卓应用创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 用户 Id
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 应用模式（NORMAL : 普通模式；ADVANCED : 高级模式）
	AppMode *string `json:"AppMode,omitnil,omitempty" name:"AppMode"`

	// 应用更新状态，取值：UPLOADING 上传中、CREATING 创建中、CREATE_FAIL 创建失败、CREATE_SUCCESS 创建成功、PACKAGE_NAME_MISMATCH 包名不匹配、VERSION_ALREADY_EXISTS 版本已存在、APP_PARSE_FAIL app 解析失败、APP_EXISTS_SECURITY_RISK app 存在安全风险、NORMAL 默认状态
	UpdateState *string `json:"UpdateState,omitnil,omitempty" name:"UpdateState"`

	// 安卓应用包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

type AndroidAppCosInfo struct {
	// 安卓应用ID
	AndroidAppId *string `json:"AndroidAppId,omitnil,omitempty" name:"AndroidAppId"`

	// 应用名称（支持 apk 和 tgz 两种格式文件，当应用 AppMode 为 NORMAL 时，只支持上传 apk 类型文件，当应用 AppMode 为 ADVANCED 高级模式时，只支持上传  tgz 类型文件）
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`
}

type AndroidAppVersionInfo struct {
	// 安卓应用版本
	AndroidAppVersion *string `json:"AndroidAppVersion,omitnil,omitempty" name:"AndroidAppVersion"`

	// 安卓应用版本创建状态，取值：NORMAL：无（默认）、UPLOADING：上传中、CREATING： 创建中、CREATE_FAIL：创建失败、PACKAGE_NAME_MISMATCH：包名不匹配、VERSION_ALREADY_EXISTS：版本已存在、APP_PARSE_FAIL： app 解析失败、APP_EXISTS_SECURITY_RISK：app 存在安全风险、CREATE_SUCCESS：创建成功
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 安卓应用版本创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// shell 安装命令（支持多条命令执行，通过 && 组合；只在应用 AppMode 为 ADVANCED 高级模式下 才会生效）
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// shell 卸载命令（支持多条命令执行，通过 && 组合；只在应用 AppMode 为 ADVANCED 高级模式下 才会生效）
	UninstallCommand *string `json:"UninstallCommand,omitnil,omitempty" name:"UninstallCommand"`

	// 应用资源清理模式（实例安装应用所用资源），取值：CLEANUP_ON_UNINSTALL（默认值），卸载 App 时清理；CLEANUP_AFTER_INSTALL，安装 App 后立即清理。普通应用只有 CLEANUP_AFTER_INSTALL 模式。
	CleanupMode *string `json:"CleanupMode,omitnil,omitempty" name:"CleanupMode"`

	// 安卓应用版本名称（版本描述、备注）
	AndroidAppVersionName *string `json:"AndroidAppVersionName,omitnil,omitempty" name:"AndroidAppVersionName"`

	// 安卓应用启动页
	Activity *string `json:"Activity,omitnil,omitempty" name:"Activity"`

	// 应用版本号（Version Name）
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`

	// 应用包 MD5
	MD5 *string `json:"MD5,omitnil,omitempty" name:"MD5"`

	// 应用包文件大小（字节）
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 安卓应用包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

type AndroidInstance struct {
	// 实例 ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`

	// 实例所在区域
	AndroidInstanceRegion *string `json:"AndroidInstanceRegion,omitnil,omitempty" name:"AndroidInstanceRegion"`

	// 实例可用区
	AndroidInstanceZone *string `json:"AndroidInstanceZone,omitnil,omitempty" name:"AndroidInstanceZone"`

	// 实例状态：INITIALIZING，NORMAL，PROCESSING
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 实例规格
	AndroidInstanceType *string `json:"AndroidInstanceType,omitnil,omitempty" name:"AndroidInstanceType"`

	// 实例镜像 ID
	AndroidInstanceImageId *string `json:"AndroidInstanceImageId,omitnil,omitempty" name:"AndroidInstanceImageId"`

	// 分辨率宽度
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 分辨率高度
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 宿主机 ID
	HostSerialNumber *string `json:"HostSerialNumber,omitnil,omitempty" name:"HostSerialNumber"`

	// 分组 ID
	AndroidInstanceGroupId *string `json:"AndroidInstanceGroupId,omitnil,omitempty" name:"AndroidInstanceGroupId"`

	// 标签列表
	AndroidInstanceLabels []*AndroidInstanceLabel `json:"AndroidInstanceLabels,omitnil,omitempty" name:"AndroidInstanceLabels"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 内网 IP
	PrivateIP *string `json:"PrivateIP,omitnil,omitempty" name:"PrivateIP"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 机箱 ID
	HostServerSerialNumber *string `json:"HostServerSerialNumber,omitnil,omitempty" name:"HostServerSerialNumber"`

	// 服务状态。
	// IDLE：未连接
	// ESTABLISHED：连接中
	ServiceStatus *string `json:"ServiceStatus,omitnil,omitempty" name:"ServiceStatus"`
}

type AndroidInstanceAppBlacklist struct {
	// 安卓实例 ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`

	// 应用黑名单
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppBlacklist []*string `json:"AppBlacklist,omitnil,omitempty" name:"AppBlacklist"`
}

type AndroidInstanceAppInfo struct {
	// 应用id
	AndroidAppId *string `json:"AndroidAppId,omitnil,omitempty" name:"AndroidAppId"`

	// 应用名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 应用版本
	AndroidAppVersion *string `json:"AndroidAppVersion,omitnil,omitempty" name:"AndroidAppVersion"`

	// 应用包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 应用包版本
	//
	// Deprecated: PackageVersion is deprecated.
	PackageVersion *string `json:"PackageVersion,omitnil,omitempty" name:"PackageVersion"`

	// 应用包标签
	PackageLabel *string `json:"PackageLabel,omitnil,omitempty" name:"PackageLabel"`

	// 应用包版本号
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`
}

type AndroidInstanceBackup struct {
	// 备份ID
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 备份状态
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 大小，单位 Byte
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 备份的安卓实例 ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type AndroidInstanceDevice struct {
	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 型号
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`
}

type AndroidInstanceError struct {
	// 安卓实例 ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`

	// 错误信息
	Error *Error `json:"Error,omitnil,omitempty" name:"Error"`
}

type AndroidInstanceHostTask struct {
	// 任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 宿主机序列号
	HostSerialNumber *string `json:"HostSerialNumber,omitnil,omitempty" name:"HostSerialNumber"`
}

type AndroidInstanceImage struct {
	// 镜像 ID
	AndroidInstanceImageId *string `json:"AndroidInstanceImageId,omitnil,omitempty" name:"AndroidInstanceImageId"`

	// 镜像名称，由业务方自定义，仅用于展示
	AndroidInstanceImageName *string `json:"AndroidInstanceImageName,omitnil,omitempty" name:"AndroidInstanceImageName"`

	// 镜像状态
	AndroidInstanceImageState *string `json:"AndroidInstanceImageState,omitnil,omitempty" name:"AndroidInstanceImageState"`

	// 镜像可用区
	AndroidInstanceImageZone *string `json:"AndroidInstanceImageZone,omitnil,omitempty" name:"AndroidInstanceImageZone"`

	// 镜像描述
	AndroidInstanceImageDescription *string `json:"AndroidInstanceImageDescription,omitnil,omitempty" name:"AndroidInstanceImageDescription"`

	// 安卓10
	AndroidVersion *string `json:"AndroidVersion,omitnil,omitempty" name:"AndroidVersion"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type AndroidInstanceInformation struct {
	// 安卓实例 ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`

	// 实例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type AndroidInstanceLabel struct {
	// 标签键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type AndroidInstanceLabelDetail struct {
	// 标签
	Label *AndroidInstanceLabel `json:"Label,omitnil,omitempty" name:"Label"`

	// 标签描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 标签创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type AndroidInstanceProperty struct {
	// 属性键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 属性值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type AndroidInstanceTask struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 实例ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`
}

type AndroidInstanceTaskStatus struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务状态：SUCCESS，FAILED，PROCESSING，PENDING,CANCELED
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`

	// 任务执行结果描述，针对某些任务，可以是可解析的 json
	TaskResult *string `json:"TaskResult,omitnil,omitempty" name:"TaskResult"`

	// 任务类型
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 任务创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务完成时间
	CompleteTime *string `json:"CompleteTime,omitnil,omitempty" name:"CompleteTime"`
}

type AndroidInstanceUploadFile struct {
	// 安卓实例 ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`

	// 文件上传 URL
	FileURL *string `json:"FileURL,omitnil,omitempty" name:"FileURL"`

	// 上传目标目录，只能上传到 /sdcard/ 目录或其子目录下
	DestinationDirectory *string `json:"DestinationDirectory,omitnil,omitempty" name:"DestinationDirectory"`

	// 目标文件名
	DestinationFileName *string `json:"DestinationFileName,omitnil,omitempty" name:"DestinationFileName"`
}

// Predefined struct for user
type BackUpAndroidInstanceRequestParams struct {
	// 安卓实例ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`

	// 包含的路径，支持仅含一个通配符*，通配符不能出现在路径开始
	Includes []*string `json:"Includes,omitnil,omitempty" name:"Includes"`

	// 需要排除路径，支持仅含一个通配符*，通配符不能出现在路径开始
	Excludes []*string `json:"Excludes,omitnil,omitempty" name:"Excludes"`
}

type BackUpAndroidInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 安卓实例ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`

	// 包含的路径，支持仅含一个通配符*，通配符不能出现在路径开始
	Includes []*string `json:"Includes,omitnil,omitempty" name:"Includes"`

	// 需要排除路径，支持仅含一个通配符*，通配符不能出现在路径开始
	Excludes []*string `json:"Excludes,omitnil,omitempty" name:"Excludes"`
}

func (r *BackUpAndroidInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BackUpAndroidInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceId")
	delete(f, "Includes")
	delete(f, "Excludes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BackUpAndroidInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BackUpAndroidInstanceResponseParams struct {
	// 备份 ID
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BackUpAndroidInstanceResponse struct {
	*tchttp.BaseResponse
	Response *BackUpAndroidInstanceResponseParams `json:"Response"`
}

func (r *BackUpAndroidInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BackUpAndroidInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BackUpAndroidInstanceToStorageRequestParams struct {
	// 安卓实例ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`

	// 存储服务器类型，如 COS、S3。注意：使用 COS 和 S3 都将占用外网带宽。
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 自定义对象Key
	ObjectKey *string `json:"ObjectKey,omitnil,omitempty" name:"ObjectKey"`

	// 包含的路径，支持仅含一个通配符*，通配符不能出现在路径开始
	Includes []*string `json:"Includes,omitnil,omitempty" name:"Includes"`

	// 需要排除路径，支持仅含一个通配符*，通配符不能出现在路径开始
	Excludes []*string `json:"Excludes,omitnil,omitempty" name:"Excludes"`

	// COS协议选项
	COSOptions *COSOptions `json:"COSOptions,omitnil,omitempty" name:"COSOptions"`

	// S3存储协议选项
	S3Options *S3Options `json:"S3Options,omitnil,omitempty" name:"S3Options"`
}

type BackUpAndroidInstanceToStorageRequest struct {
	*tchttp.BaseRequest
	
	// 安卓实例ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`

	// 存储服务器类型，如 COS、S3。注意：使用 COS 和 S3 都将占用外网带宽。
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 自定义对象Key
	ObjectKey *string `json:"ObjectKey,omitnil,omitempty" name:"ObjectKey"`

	// 包含的路径，支持仅含一个通配符*，通配符不能出现在路径开始
	Includes []*string `json:"Includes,omitnil,omitempty" name:"Includes"`

	// 需要排除路径，支持仅含一个通配符*，通配符不能出现在路径开始
	Excludes []*string `json:"Excludes,omitnil,omitempty" name:"Excludes"`

	// COS协议选项
	COSOptions *COSOptions `json:"COSOptions,omitnil,omitempty" name:"COSOptions"`

	// S3存储协议选项
	S3Options *S3Options `json:"S3Options,omitnil,omitempty" name:"S3Options"`
}

func (r *BackUpAndroidInstanceToStorageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BackUpAndroidInstanceToStorageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceId")
	delete(f, "StorageType")
	delete(f, "ObjectKey")
	delete(f, "Includes")
	delete(f, "Excludes")
	delete(f, "COSOptions")
	delete(f, "S3Options")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BackUpAndroidInstanceToStorageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BackUpAndroidInstanceToStorageResponseParams struct {
	// 实例任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BackUpAndroidInstanceToStorageResponse struct {
	*tchttp.BaseResponse
	Response *BackUpAndroidInstanceToStorageResponseParams `json:"Response"`
}

func (r *BackUpAndroidInstanceToStorageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BackUpAndroidInstanceToStorageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type COSOptions struct {
	// 存储桶
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 存储区域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

// Predefined struct for user
type CleanAndroidInstancesAppDataRequestParams struct {
	// 安卓实例 ID 列表（最大100条数据）
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 应用包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

type CleanAndroidInstancesAppDataRequest struct {
	*tchttp.BaseRequest
	
	// 安卓实例 ID 列表（最大100条数据）
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 应用包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

func (r *CleanAndroidInstancesAppDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CleanAndroidInstancesAppDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	delete(f, "PackageName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CleanAndroidInstancesAppDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CleanAndroidInstancesAppDataResponseParams struct {
	// 错误列表。如果实例操作都成功，则响应没有这个字段；如果有实例操作失败，该字段包含了实例操作的错误信息
	AndroidInstanceErrors []*AndroidInstanceError `json:"AndroidInstanceErrors,omitnil,omitempty" name:"AndroidInstanceErrors"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CleanAndroidInstancesAppDataResponse struct {
	*tchttp.BaseResponse
	Response *CleanAndroidInstancesAppDataResponseParams `json:"Response"`
}

func (r *CleanAndroidInstancesAppDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CleanAndroidInstancesAppDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ConnectAndroidInstanceRequestParams struct {
	// 用户Session信息
	ClientSession *string `json:"ClientSession,omitnil,omitempty" name:"ClientSession"`

	// 实例ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`

	// 用户IP，用户客户端的公网IP，用于选择最佳网络链路
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`
}

type ConnectAndroidInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 用户Session信息
	ClientSession *string `json:"ClientSession,omitnil,omitempty" name:"ClientSession"`

	// 实例ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`

	// 用户IP，用户客户端的公网IP，用于选择最佳网络链路
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`
}

func (r *ConnectAndroidInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConnectAndroidInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientSession")
	delete(f, "AndroidInstanceId")
	delete(f, "UserIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ConnectAndroidInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ConnectAndroidInstanceResponseParams struct {
	// 服务端session信息
	ServerSession *string `json:"ServerSession,omitnil,omitempty" name:"ServerSession"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ConnectAndroidInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ConnectAndroidInstanceResponseParams `json:"Response"`
}

func (r *ConnectAndroidInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConnectAndroidInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopyAndroidInstanceRequestParams struct {
	// 源安卓实例 ID
	SourceAndroidInstanceId *string `json:"SourceAndroidInstanceId,omitnil,omitempty" name:"SourceAndroidInstanceId"`

	// 目的安卓实例 ID
	TargetAndroidInstanceId *string `json:"TargetAndroidInstanceId,omitnil,omitempty" name:"TargetAndroidInstanceId"`

	// 包含的路径，支持仅含一个通配符*，通配符不能出现在路径开始
	Includes []*string `json:"Includes,omitnil,omitempty" name:"Includes"`

	// 需要排除路径，支持仅含一个通配符*，通配符不能出现在路径开始
	Excludes []*string `json:"Excludes,omitnil,omitempty" name:"Excludes"`
}

type CopyAndroidInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 源安卓实例 ID
	SourceAndroidInstanceId *string `json:"SourceAndroidInstanceId,omitnil,omitempty" name:"SourceAndroidInstanceId"`

	// 目的安卓实例 ID
	TargetAndroidInstanceId *string `json:"TargetAndroidInstanceId,omitnil,omitempty" name:"TargetAndroidInstanceId"`

	// 包含的路径，支持仅含一个通配符*，通配符不能出现在路径开始
	Includes []*string `json:"Includes,omitnil,omitempty" name:"Includes"`

	// 需要排除路径，支持仅含一个通配符*，通配符不能出现在路径开始
	Excludes []*string `json:"Excludes,omitnil,omitempty" name:"Excludes"`
}

func (r *CopyAndroidInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyAndroidInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SourceAndroidInstanceId")
	delete(f, "TargetAndroidInstanceId")
	delete(f, "Includes")
	delete(f, "Excludes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CopyAndroidInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopyAndroidInstanceResponseParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CopyAndroidInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CopyAndroidInstanceResponseParams `json:"Response"`
}

func (r *CopyAndroidInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyAndroidInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAndroidAppRequestParams struct {
	// 安卓应用名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户 Id
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 应用模式（NORMAL : 普通模式、只支持 apk 文件上传，为默认值；ADVANCED : 高级模式、只支持上传 tgz 文件 和 自定义 shell 命令执行）
	AppMode *string `json:"AppMode,omitnil,omitempty" name:"AppMode"`
}

type CreateAndroidAppRequest struct {
	*tchttp.BaseRequest
	
	// 安卓应用名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户 Id
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 应用模式（NORMAL : 普通模式、只支持 apk 文件上传，为默认值；ADVANCED : 高级模式、只支持上传 tgz 文件 和 自定义 shell 命令执行）
	AppMode *string `json:"AppMode,omitnil,omitempty" name:"AppMode"`
}

func (r *CreateAndroidAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAndroidAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "UserId")
	delete(f, "AppMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAndroidAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAndroidAppResponseParams struct {
	// 应用ID
	AndroidAppId *string `json:"AndroidAppId,omitnil,omitempty" name:"AndroidAppId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAndroidAppResponse struct {
	*tchttp.BaseResponse
	Response *CreateAndroidAppResponseParams `json:"Response"`
}

func (r *CreateAndroidAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAndroidAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAndroidAppVersionRequestParams struct {
	// 应用ID
	AndroidAppId *string `json:"AndroidAppId,omitnil,omitempty" name:"AndroidAppId"`

	// 应用包下载地址
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// 应用 shell 安装命令（支持多条命令执行，通过 && 组合；只在应用 AppMode 为 ADVANCED 高级模式下 才会生效）
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 应用 shell 卸载命令（支持多条命令执行，通过 && 组合；只在应用 AppMode 为 ADVANCED 高级模式下 才会生效）
	UninstallCommand *string `json:"UninstallCommand,omitnil,omitempty" name:"UninstallCommand"`

	// 应用资源清理模式（实例安装应用所用资源），取值：CLEANUP_ON_UNINSTALL（默认值），卸载 App 时清理；CLEANUP_AFTER_INSTALL，安装 App 后立即清理。普通应用只有 CLEANUP_AFTER_INSTALL 模式。
	CleanupMode *string `json:"CleanupMode,omitnil,omitempty" name:"CleanupMode"`
}

type CreateAndroidAppVersionRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	AndroidAppId *string `json:"AndroidAppId,omitnil,omitempty" name:"AndroidAppId"`

	// 应用包下载地址
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// 应用 shell 安装命令（支持多条命令执行，通过 && 组合；只在应用 AppMode 为 ADVANCED 高级模式下 才会生效）
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 应用 shell 卸载命令（支持多条命令执行，通过 && 组合；只在应用 AppMode 为 ADVANCED 高级模式下 才会生效）
	UninstallCommand *string `json:"UninstallCommand,omitnil,omitempty" name:"UninstallCommand"`

	// 应用资源清理模式（实例安装应用所用资源），取值：CLEANUP_ON_UNINSTALL（默认值），卸载 App 时清理；CLEANUP_AFTER_INSTALL，安装 App 后立即清理。普通应用只有 CLEANUP_AFTER_INSTALL 模式。
	CleanupMode *string `json:"CleanupMode,omitnil,omitempty" name:"CleanupMode"`
}

func (r *CreateAndroidAppVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAndroidAppVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidAppId")
	delete(f, "DownloadUrl")
	delete(f, "Command")
	delete(f, "UninstallCommand")
	delete(f, "CleanupMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAndroidAppVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAndroidAppVersionResponseParams struct {
	// 应用版本
	AndroidAppVersion *string `json:"AndroidAppVersion,omitnil,omitempty" name:"AndroidAppVersion"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAndroidAppVersionResponse struct {
	*tchttp.BaseResponse
	Response *CreateAndroidAppVersionResponseParams `json:"Response"`
}

func (r *CreateAndroidAppVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAndroidAppVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAndroidInstanceADBRequestParams struct {
	// 安卓实例 ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`
}

type CreateAndroidInstanceADBRequest struct {
	*tchttp.BaseRequest
	
	// 安卓实例 ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`
}

func (r *CreateAndroidInstanceADBRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAndroidInstanceADBRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAndroidInstanceADBRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAndroidInstanceADBResponseParams struct {
	// 连接私钥，需要保存为文件形式，例如 private_key.pem
	PrivateKey *string `json:"PrivateKey,omitnil,omitempty" name:"PrivateKey"`

	// 用户名称
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 连接地址
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 连接端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 连接参考命令
	ConnectCommand *string `json:"ConnectCommand,omitnil,omitempty" name:"ConnectCommand"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAndroidInstanceADBResponse struct {
	*tchttp.BaseResponse
	Response *CreateAndroidInstanceADBResponseParams `json:"Response"`
}

func (r *CreateAndroidInstanceADBResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAndroidInstanceADBResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAndroidInstanceAcceleratorTokenRequestParams struct {
	// 实例 ID 列表。每次请求的实例的上限为 500。
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 用户IP，可以根据该 IP 选择就近加速点。如果不填，将自动选择就近加速点。
	UserIP *string `json:"UserIP,omitnil,omitempty" name:"UserIP"`
}

type CreateAndroidInstanceAcceleratorTokenRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表。每次请求的实例的上限为 500。
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 用户IP，可以根据该 IP 选择就近加速点。如果不填，将自动选择就近加速点。
	UserIP *string `json:"UserIP,omitnil,omitempty" name:"UserIP"`
}

func (r *CreateAndroidInstanceAcceleratorTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAndroidInstanceAcceleratorTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	delete(f, "UserIP")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAndroidInstanceAcceleratorTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAndroidInstanceAcceleratorTokenResponseParams struct {
	// token
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 加速信息
	AcceleratorInfo *string `json:"AcceleratorInfo,omitnil,omitempty" name:"AcceleratorInfo"`

	// 安卓实例错误列表。列表包含有问题的安卓实例 ID 以及发生的错误信息。
	AndroidInstanceErrors []*AndroidInstanceError `json:"AndroidInstanceErrors,omitnil,omitempty" name:"AndroidInstanceErrors"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAndroidInstanceAcceleratorTokenResponse struct {
	*tchttp.BaseResponse
	Response *CreateAndroidInstanceAcceleratorTokenResponseParams `json:"Response"`
}

func (r *CreateAndroidInstanceAcceleratorTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAndroidInstanceAcceleratorTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAndroidInstanceImageRequestParams struct {
	// 安卓实例镜像名称
	AndroidInstanceImageName *string `json:"AndroidInstanceImageName,omitnil,omitempty" name:"AndroidInstanceImageName"`

	// 安卓实例 ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`

	// 安卓实例镜像描述
	AndroidInstanceImageDescription *string `json:"AndroidInstanceImageDescription,omitnil,omitempty" name:"AndroidInstanceImageDescription"`
}

type CreateAndroidInstanceImageRequest struct {
	*tchttp.BaseRequest
	
	// 安卓实例镜像名称
	AndroidInstanceImageName *string `json:"AndroidInstanceImageName,omitnil,omitempty" name:"AndroidInstanceImageName"`

	// 安卓实例 ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`

	// 安卓实例镜像描述
	AndroidInstanceImageDescription *string `json:"AndroidInstanceImageDescription,omitnil,omitempty" name:"AndroidInstanceImageDescription"`
}

func (r *CreateAndroidInstanceImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAndroidInstanceImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceImageName")
	delete(f, "AndroidInstanceId")
	delete(f, "AndroidInstanceImageDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAndroidInstanceImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAndroidInstanceImageResponseParams struct {
	// 安卓实例镜像 ID
	AndroidInstanceImageId *string `json:"AndroidInstanceImageId,omitnil,omitempty" name:"AndroidInstanceImageId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAndroidInstanceImageResponse struct {
	*tchttp.BaseResponse
	Response *CreateAndroidInstanceImageResponseParams `json:"Response"`
}

func (r *CreateAndroidInstanceImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAndroidInstanceImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAndroidInstanceLabelRequestParams struct {
	// 标签键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值。普通场景下，该值不需要填写；高级场景下，需要两个层级进行分组时才填写。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 标签描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateAndroidInstanceLabelRequest struct {
	*tchttp.BaseRequest
	
	// 标签键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值。普通场景下，该值不需要填写；高级场景下，需要两个层级进行分组时才填写。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 标签描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateAndroidInstanceLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAndroidInstanceLabelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Key")
	delete(f, "Value")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAndroidInstanceLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAndroidInstanceLabelResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAndroidInstanceLabelResponse struct {
	*tchttp.BaseResponse
	Response *CreateAndroidInstanceLabelResponseParams `json:"Response"`
}

func (r *CreateAndroidInstanceLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAndroidInstanceLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAndroidInstanceSSHRequestParams struct {
	// 实例ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`

	// 连接过期时间，最长可设置7天
	ExpiredTime *string `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`
}

type CreateAndroidInstanceSSHRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`

	// 连接过期时间，最长可设置7天
	ExpiredTime *string `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`
}

func (r *CreateAndroidInstanceSSHRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAndroidInstanceSSHRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceId")
	delete(f, "ExpiredTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAndroidInstanceSSHRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAndroidInstanceSSHResponseParams struct {
	// 连接私钥，需要保存为文件形式，例如 private_key.pem
	PrivateKey *string `json:"PrivateKey,omitnil,omitempty" name:"PrivateKey"`

	// 用户名称
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 连接地址
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 连接端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 连接参考命令
	ConnectCommand *string `json:"ConnectCommand,omitnil,omitempty" name:"ConnectCommand"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAndroidInstanceSSHResponse struct {
	*tchttp.BaseResponse
	Response *CreateAndroidInstanceSSHResponseParams `json:"Response"`
}

func (r *CreateAndroidInstanceSSHResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAndroidInstanceSSHResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAndroidInstanceWebShellRequestParams struct {
	// 实例 ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`
}

type CreateAndroidInstanceWebShellRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`
}

func (r *CreateAndroidInstanceWebShellRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAndroidInstanceWebShellRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAndroidInstanceWebShellRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAndroidInstanceWebShellResponseParams struct {
	// 鉴权密钥
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 连接地址
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 连接区域
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 访问链接，可以直接使用此链接访问 WebShell
	ConnectUrl *string `json:"ConnectUrl,omitnil,omitempty" name:"ConnectUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAndroidInstanceWebShellResponse struct {
	*tchttp.BaseResponse
	Response *CreateAndroidInstanceWebShellResponseParams `json:"Response"`
}

func (r *CreateAndroidInstanceWebShellResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAndroidInstanceWebShellResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAndroidInstancesAccessTokenRequestParams struct {
	// 实例 ID 列表。每次请求的实例的上限为 500。
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 有效期，默认为 12 小时，最大为 24 小时。支持 s（秒）、m（分）、h（小时）等单位，比如 12h 表示 12 小时，1h2m3s 表示一小时两分三秒
	ExpirationDuration *string `json:"ExpirationDuration,omitnil,omitempty" name:"ExpirationDuration"`

	// 模式。
	// STANDARD：默认值，标准模式
	// ACCELERATED：加速模式，该模式需要开通加速服务才能生效
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 用户 IP。在加速模式下，该字段必填。
	UserIP *string `json:"UserIP,omitnil,omitempty" name:"UserIP"`
}

type CreateAndroidInstancesAccessTokenRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表。每次请求的实例的上限为 500。
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 有效期，默认为 12 小时，最大为 24 小时。支持 s（秒）、m（分）、h（小时）等单位，比如 12h 表示 12 小时，1h2m3s 表示一小时两分三秒
	ExpirationDuration *string `json:"ExpirationDuration,omitnil,omitempty" name:"ExpirationDuration"`

	// 模式。
	// STANDARD：默认值，标准模式
	// ACCELERATED：加速模式，该模式需要开通加速服务才能生效
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 用户 IP。在加速模式下，该字段必填。
	UserIP *string `json:"UserIP,omitnil,omitempty" name:"UserIP"`
}

func (r *CreateAndroidInstancesAccessTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAndroidInstancesAccessTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	delete(f, "ExpirationDuration")
	delete(f, "Mode")
	delete(f, "UserIP")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAndroidInstancesAccessTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAndroidInstancesAccessTokenResponseParams struct {
	// token
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 访问信息
	AccessInfo *string `json:"AccessInfo,omitnil,omitempty" name:"AccessInfo"`

	// 安卓实例错误列表。列表包含有问题的安卓实例 ID，生成的 Token 对这些有问题的实例无效。
	AndroidInstanceErrors []*AndroidInstanceError `json:"AndroidInstanceErrors,omitnil,omitempty" name:"AndroidInstanceErrors"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAndroidInstancesAccessTokenResponse struct {
	*tchttp.BaseResponse
	Response *CreateAndroidInstancesAccessTokenResponseParams `json:"Response"`
}

func (r *CreateAndroidInstancesAccessTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAndroidInstancesAccessTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAndroidInstancesRequestParams struct {
	// 安卓实例可用区。
	// ap-guangzhou-3：广州三区
	// ap-shenzhen-1：深圳一区
	// ap-xian-ec-1：西安一区
	// ap-hangzhou-ec-1：杭州一区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 安卓实例类型。
	// A1：单开
	// A2：双开
	// A3：三开
	// A4：四开
	// A5：五开
	// A6：六开
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 当 HostSerialNumbers 不为空时，该参数表示每个宿主机要创建的安卓实例数量；
	// 当 HostSerialNumbers 为空时，该参数表示要创建安卓实例的总数量，最大值为 100。
	Number *uint64 `json:"Number,omitnil,omitempty" name:"Number"`

	// 宿主机 ID 列表。可以指定宿主机 ID 进行创建；也可以不指定由系统自动分配宿主机。
	HostSerialNumbers []*string `json:"HostSerialNumbers,omitnil,omitempty" name:"HostSerialNumbers"`

	// 镜像 ID。如果不填，将使用默认的系统镜像
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 安卓实例标签列表
	Labels []*AndroidInstanceLabel `json:"Labels,omitnil,omitempty" name:"Labels"`
}

type CreateAndroidInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 安卓实例可用区。
	// ap-guangzhou-3：广州三区
	// ap-shenzhen-1：深圳一区
	// ap-xian-ec-1：西安一区
	// ap-hangzhou-ec-1：杭州一区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 安卓实例类型。
	// A1：单开
	// A2：双开
	// A3：三开
	// A4：四开
	// A5：五开
	// A6：六开
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 当 HostSerialNumbers 不为空时，该参数表示每个宿主机要创建的安卓实例数量；
	// 当 HostSerialNumbers 为空时，该参数表示要创建安卓实例的总数量，最大值为 100。
	Number *uint64 `json:"Number,omitnil,omitempty" name:"Number"`

	// 宿主机 ID 列表。可以指定宿主机 ID 进行创建；也可以不指定由系统自动分配宿主机。
	HostSerialNumbers []*string `json:"HostSerialNumbers,omitnil,omitempty" name:"HostSerialNumbers"`

	// 镜像 ID。如果不填，将使用默认的系统镜像
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 安卓实例标签列表
	Labels []*AndroidInstanceLabel `json:"Labels,omitnil,omitempty" name:"Labels"`
}

func (r *CreateAndroidInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAndroidInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "Type")
	delete(f, "Number")
	delete(f, "HostSerialNumbers")
	delete(f, "ImageId")
	delete(f, "Labels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAndroidInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAndroidInstancesResponseParams struct {
	// 安卓实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAndroidInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateAndroidInstancesResponseParams `json:"Response"`
}

func (r *CreateAndroidInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAndroidInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAndroidInstancesScreenshotRequestParams struct {
	// 实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`
}

type CreateAndroidInstancesScreenshotRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`
}

func (r *CreateAndroidInstancesScreenshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAndroidInstancesScreenshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAndroidInstancesScreenshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAndroidInstancesScreenshotResponseParams struct {
	// 任务列表
	TaskSet []*AndroidInstanceTask `json:"TaskSet,omitnil,omitempty" name:"TaskSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAndroidInstancesScreenshotResponse struct {
	*tchttp.BaseResponse
	Response *CreateAndroidInstancesScreenshotResponseParams `json:"Response"`
}

func (r *CreateAndroidInstancesScreenshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAndroidInstancesScreenshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCosCredentialRequestParams struct {
	// Cos 密钥类型，取值： Mobile 云手游、PC 云端游、AndroidApp 云手机应用管理、AndroidAppFile 云手机文件管理、AndroidAppBackup 云手机备份还原
	CosType *string `json:"CosType,omitnil,omitempty" name:"CosType"`

	// 云手机应用管理 Cos 数据
	AndroidAppCosInfo *AndroidAppCosInfo `json:"AndroidAppCosInfo,omitnil,omitempty" name:"AndroidAppCosInfo"`

	// 云手机文件管理 Cos 数据
	AndroidAppFileCosInfo *FileCosInfo `json:"AndroidAppFileCosInfo,omitnil,omitempty" name:"AndroidAppFileCosInfo"`
}

type CreateCosCredentialRequest struct {
	*tchttp.BaseRequest
	
	// Cos 密钥类型，取值： Mobile 云手游、PC 云端游、AndroidApp 云手机应用管理、AndroidAppFile 云手机文件管理、AndroidAppBackup 云手机备份还原
	CosType *string `json:"CosType,omitnil,omitempty" name:"CosType"`

	// 云手机应用管理 Cos 数据
	AndroidAppCosInfo *AndroidAppCosInfo `json:"AndroidAppCosInfo,omitnil,omitempty" name:"AndroidAppCosInfo"`

	// 云手机文件管理 Cos 数据
	AndroidAppFileCosInfo *FileCosInfo `json:"AndroidAppFileCosInfo,omitnil,omitempty" name:"AndroidAppFileCosInfo"`
}

func (r *CreateCosCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CosType")
	delete(f, "AndroidAppCosInfo")
	delete(f, "AndroidAppFileCosInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCosCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCosCredentialResponseParams struct {
	// Cos SecretID
	SecretID *string `json:"SecretID,omitnil,omitempty" name:"SecretID"`

	// Cos SecretKey
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// Cos Session
	SessionToken *string `json:"SessionToken,omitnil,omitempty" name:"SessionToken"`

	// Cos Bucket
	CosBucket *string `json:"CosBucket,omitnil,omitempty" name:"CosBucket"`

	// Cos Region
	CosRegion *string `json:"CosRegion,omitnil,omitempty" name:"CosRegion"`

	// Cos 操作路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// Cos 密钥的起始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Cos 密钥的失效时间
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCosCredentialResponse struct {
	*tchttp.BaseResponse
	Response *CreateCosCredentialResponseParams `json:"Response"`
}

func (r *CreateCosCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSessionRequestParams struct {
	// 唯一用户身份标识，由业务方自定义，平台不予理解。（可根据业务需要决定使用用户的唯一身份标识或是使用时间戳随机生成；在用户重连时应保持UserId不变）
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 【已废弃】只在TrylockWorker时生效
	GameId *string `json:"GameId,omitnil,omitempty" name:"GameId"`

	// 【已废弃】只在TrylockWorker时生效
	GameRegion *string `json:"GameRegion,omitnil,omitempty" name:"GameRegion"`

	// 游戏参数
	GameParas *string `json:"GameParas,omitnil,omitempty" name:"GameParas"`

	// 客户端session信息，从JSSDK请求中获得。特殊的，当 RunMode 参数为 RunWithoutClient 时，该字段可以为空
	ClientSession *string `json:"ClientSession,omitnil,omitempty" name:"ClientSession"`

	// 分辨率,，可设置为1080p或720p或1920x1080格式
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 背景图url，格式为png或jpeg，宽高1920*1080
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 【已废弃】
	SetNo *uint64 `json:"SetNo,omitnil,omitempty" name:"SetNo"`

	// 【已废弃】
	Bitrate *uint64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// 单位Mbps，动态调整最大码率建议值，会按实际情况调整
	MaxBitrate *uint64 `json:"MaxBitrate,omitnil,omitempty" name:"MaxBitrate"`

	// 单位Mbps，动态调整最小码率建议值，会按实际情况调整
	MinBitrate *uint64 `json:"MinBitrate,omitnil,omitempty" name:"MinBitrate"`

	// 帧率，可设置为30、45、60、90、120、144
	Fps *uint64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// 【推荐填写】用户IP，用户客户端的公网IP，用于就近调度，不填将严重影响用户体验
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// 【已废弃】优化项，便于客户灰度开启新的优化项，默认为0
	Optimization *uint64 `json:"Optimization,omitnil,omitempty" name:"Optimization"`

	// 【互动云游】游戏主机用户ID
	HostUserId *string `json:"HostUserId,omitnil,omitempty" name:"HostUserId"`

	// 【互动云游】角色；Player表示玩家；Viewer表示观察者
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 游戏相关参数
	GameContext *string `json:"GameContext,omitnil,omitempty" name:"GameContext"`

	// 云端运行模式。
	// RunWithoutClient：允许无客户端连接的情况下仍保持云端 App 运行
	// 默认值（空）：要求必须有客户端连接才会保持云端 App 运行。
	RunMode *string `json:"RunMode,omitnil,omitempty" name:"RunMode"`
}

type CreateSessionRequest struct {
	*tchttp.BaseRequest
	
	// 唯一用户身份标识，由业务方自定义，平台不予理解。（可根据业务需要决定使用用户的唯一身份标识或是使用时间戳随机生成；在用户重连时应保持UserId不变）
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 【已废弃】只在TrylockWorker时生效
	GameId *string `json:"GameId,omitnil,omitempty" name:"GameId"`

	// 【已废弃】只在TrylockWorker时生效
	GameRegion *string `json:"GameRegion,omitnil,omitempty" name:"GameRegion"`

	// 游戏参数
	GameParas *string `json:"GameParas,omitnil,omitempty" name:"GameParas"`

	// 客户端session信息，从JSSDK请求中获得。特殊的，当 RunMode 参数为 RunWithoutClient 时，该字段可以为空
	ClientSession *string `json:"ClientSession,omitnil,omitempty" name:"ClientSession"`

	// 分辨率,，可设置为1080p或720p或1920x1080格式
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 背景图url，格式为png或jpeg，宽高1920*1080
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 【已废弃】
	SetNo *uint64 `json:"SetNo,omitnil,omitempty" name:"SetNo"`

	// 【已废弃】
	Bitrate *uint64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// 单位Mbps，动态调整最大码率建议值，会按实际情况调整
	MaxBitrate *uint64 `json:"MaxBitrate,omitnil,omitempty" name:"MaxBitrate"`

	// 单位Mbps，动态调整最小码率建议值，会按实际情况调整
	MinBitrate *uint64 `json:"MinBitrate,omitnil,omitempty" name:"MinBitrate"`

	// 帧率，可设置为30、45、60、90、120、144
	Fps *uint64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// 【推荐填写】用户IP，用户客户端的公网IP，用于就近调度，不填将严重影响用户体验
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// 【已废弃】优化项，便于客户灰度开启新的优化项，默认为0
	Optimization *uint64 `json:"Optimization,omitnil,omitempty" name:"Optimization"`

	// 【互动云游】游戏主机用户ID
	HostUserId *string `json:"HostUserId,omitnil,omitempty" name:"HostUserId"`

	// 【互动云游】角色；Player表示玩家；Viewer表示观察者
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 游戏相关参数
	GameContext *string `json:"GameContext,omitnil,omitempty" name:"GameContext"`

	// 云端运行模式。
	// RunWithoutClient：允许无客户端连接的情况下仍保持云端 App 运行
	// 默认值（空）：要求必须有客户端连接才会保持云端 App 运行。
	RunMode *string `json:"RunMode,omitnil,omitempty" name:"RunMode"`
}

func (r *CreateSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "GameId")
	delete(f, "GameRegion")
	delete(f, "GameParas")
	delete(f, "ClientSession")
	delete(f, "Resolution")
	delete(f, "ImageUrl")
	delete(f, "SetNo")
	delete(f, "Bitrate")
	delete(f, "MaxBitrate")
	delete(f, "MinBitrate")
	delete(f, "Fps")
	delete(f, "UserIp")
	delete(f, "Optimization")
	delete(f, "HostUserId")
	delete(f, "Role")
	delete(f, "GameContext")
	delete(f, "RunMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSessionResponseParams struct {
	// 服务端session信息，返回给JSSDK
	ServerSession *string `json:"ServerSession,omitnil,omitempty" name:"ServerSession"`

	// 【已废弃】
	RoleNumber *string `json:"RoleNumber,omitnil,omitempty" name:"RoleNumber"`

	// 【互动云游】角色；Player表示玩家；Viewer表示观察者
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSessionResponse struct {
	*tchttp.BaseResponse
	Response *CreateSessionResponseParams `json:"Response"`
}

func (r *CreateSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAndroidAppRequestParams struct {
	// 应用ID
	AndroidAppId *string `json:"AndroidAppId,omitnil,omitempty" name:"AndroidAppId"`
}

type DeleteAndroidAppRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	AndroidAppId *string `json:"AndroidAppId,omitnil,omitempty" name:"AndroidAppId"`
}

func (r *DeleteAndroidAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAndroidAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAndroidAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAndroidAppResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAndroidAppResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAndroidAppResponseParams `json:"Response"`
}

func (r *DeleteAndroidAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAndroidAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAndroidAppVersionRequestParams struct {
	// 安卓应用 Id
	AndroidAppId *string `json:"AndroidAppId,omitnil,omitempty" name:"AndroidAppId"`

	// 安卓应用版本
	AndroidAppVersion *string `json:"AndroidAppVersion,omitnil,omitempty" name:"AndroidAppVersion"`
}

type DeleteAndroidAppVersionRequest struct {
	*tchttp.BaseRequest
	
	// 安卓应用 Id
	AndroidAppId *string `json:"AndroidAppId,omitnil,omitempty" name:"AndroidAppId"`

	// 安卓应用版本
	AndroidAppVersion *string `json:"AndroidAppVersion,omitnil,omitempty" name:"AndroidAppVersion"`
}

func (r *DeleteAndroidAppVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAndroidAppVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidAppId")
	delete(f, "AndroidAppVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAndroidAppVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAndroidAppVersionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAndroidAppVersionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAndroidAppVersionResponseParams `json:"Response"`
}

func (r *DeleteAndroidAppVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAndroidAppVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAndroidInstanceBackupFilesRequestParams struct {
	// 文件对象键列表
	ObjectKeys []*string `json:"ObjectKeys,omitnil,omitempty" name:"ObjectKeys"`

	// 存储服务器类型，如 COS、S3。注意：使用 COS 和 S3 都将占用外网带宽。
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// COS协议选项
	COSOptions *COSOptions `json:"COSOptions,omitnil,omitempty" name:"COSOptions"`

	// S3存储协议选项
	S3Options *S3Options `json:"S3Options,omitnil,omitempty" name:"S3Options"`

	// 安卓实例可用区。StorageType 为 S3 时，需要填写该字段；StorageType 为 COS 时，不需要填写该字段
	AndroidInstanceZone *string `json:"AndroidInstanceZone,omitnil,omitempty" name:"AndroidInstanceZone"`
}

type DeleteAndroidInstanceBackupFilesRequest struct {
	*tchttp.BaseRequest
	
	// 文件对象键列表
	ObjectKeys []*string `json:"ObjectKeys,omitnil,omitempty" name:"ObjectKeys"`

	// 存储服务器类型，如 COS、S3。注意：使用 COS 和 S3 都将占用外网带宽。
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// COS协议选项
	COSOptions *COSOptions `json:"COSOptions,omitnil,omitempty" name:"COSOptions"`

	// S3存储协议选项
	S3Options *S3Options `json:"S3Options,omitnil,omitempty" name:"S3Options"`

	// 安卓实例可用区。StorageType 为 S3 时，需要填写该字段；StorageType 为 COS 时，不需要填写该字段
	AndroidInstanceZone *string `json:"AndroidInstanceZone,omitnil,omitempty" name:"AndroidInstanceZone"`
}

func (r *DeleteAndroidInstanceBackupFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAndroidInstanceBackupFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ObjectKeys")
	delete(f, "StorageType")
	delete(f, "COSOptions")
	delete(f, "S3Options")
	delete(f, "AndroidInstanceZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAndroidInstanceBackupFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAndroidInstanceBackupFilesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAndroidInstanceBackupFilesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAndroidInstanceBackupFilesResponseParams `json:"Response"`
}

func (r *DeleteAndroidInstanceBackupFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAndroidInstanceBackupFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAndroidInstanceBackupsRequestParams struct {
	// 备份ID列表
	BackupIds []*string `json:"BackupIds,omitnil,omitempty" name:"BackupIds"`
}

type DeleteAndroidInstanceBackupsRequest struct {
	*tchttp.BaseRequest
	
	// 备份ID列表
	BackupIds []*string `json:"BackupIds,omitnil,omitempty" name:"BackupIds"`
}

func (r *DeleteAndroidInstanceBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAndroidInstanceBackupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAndroidInstanceBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAndroidInstanceBackupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAndroidInstanceBackupsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAndroidInstanceBackupsResponseParams `json:"Response"`
}

func (r *DeleteAndroidInstanceBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAndroidInstanceBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAndroidInstanceImagesRequestParams struct {
	// 镜像 ID 列表
	AndroidInstanceImageIds []*string `json:"AndroidInstanceImageIds,omitnil,omitempty" name:"AndroidInstanceImageIds"`
}

type DeleteAndroidInstanceImagesRequest struct {
	*tchttp.BaseRequest
	
	// 镜像 ID 列表
	AndroidInstanceImageIds []*string `json:"AndroidInstanceImageIds,omitnil,omitempty" name:"AndroidInstanceImageIds"`
}

func (r *DeleteAndroidInstanceImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAndroidInstanceImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceImageIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAndroidInstanceImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAndroidInstanceImagesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAndroidInstanceImagesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAndroidInstanceImagesResponseParams `json:"Response"`
}

func (r *DeleteAndroidInstanceImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAndroidInstanceImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAndroidInstanceLabelRequestParams struct {
	// 标签键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type DeleteAndroidInstanceLabelRequest struct {
	*tchttp.BaseRequest
	
	// 标签键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

func (r *DeleteAndroidInstanceLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAndroidInstanceLabelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Key")
	delete(f, "Value")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAndroidInstanceLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAndroidInstanceLabelResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAndroidInstanceLabelResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAndroidInstanceLabelResponseParams `json:"Response"`
}

func (r *DeleteAndroidInstanceLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAndroidInstanceLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAndroidAppsRequestParams struct {
	// 分页偏移
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 应用ID数组
	AndroidAppIds []*string `json:"AndroidAppIds,omitnil,omitempty" name:"AndroidAppIds"`

	// 过滤条件，支持过滤的字段有：UserId、State、UpdateState、Name、AppMode 。其中 Name 为模糊匹配，其他参数为精确匹配。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeAndroidAppsRequest struct {
	*tchttp.BaseRequest
	
	// 分页偏移
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 应用ID数组
	AndroidAppIds []*string `json:"AndroidAppIds,omitnil,omitempty" name:"AndroidAppIds"`

	// 过滤条件，支持过滤的字段有：UserId、State、UpdateState、Name、AppMode 。其中 Name 为模糊匹配，其他参数为精确匹配。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeAndroidAppsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAndroidAppsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "AndroidAppIds")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAndroidAppsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAndroidAppsResponseParams struct {
	// 安卓应用列表
	Apps []*AndroidApp `json:"Apps,omitnil,omitempty" name:"Apps"`

	// 安卓应用列表长度
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAndroidAppsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAndroidAppsResponseParams `json:"Response"`
}

func (r *DescribeAndroidAppsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAndroidAppsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAndroidInstanceAppsRequestParams struct {
	// 实例ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`
}

type DescribeAndroidInstanceAppsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`
}

func (r *DescribeAndroidInstanceAppsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAndroidInstanceAppsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAndroidInstanceAppsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAndroidInstanceAppsResponseParams struct {
	// 安卓应用列表
	Apps []*AndroidInstanceAppInfo `json:"Apps,omitnil,omitempty" name:"Apps"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAndroidInstanceAppsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAndroidInstanceAppsResponseParams `json:"Response"`
}

func (r *DescribeAndroidInstanceAppsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAndroidInstanceAppsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAndroidInstanceBackupsRequestParams struct {
	// 备份ID列表
	BackupIds []*string `json:"BackupIds,omitnil,omitempty" name:"BackupIds"`

	// 分页偏移
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeAndroidInstanceBackupsRequest struct {
	*tchttp.BaseRequest
	
	// 备份ID列表
	BackupIds []*string `json:"BackupIds,omitnil,omitempty" name:"BackupIds"`

	// 分页偏移
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeAndroidInstanceBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAndroidInstanceBackupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAndroidInstanceBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAndroidInstanceBackupsResponseParams struct {
	// 备份列表
	Backups []*AndroidInstanceBackup `json:"Backups,omitnil,omitempty" name:"Backups"`

	// 备份总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAndroidInstanceBackupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAndroidInstanceBackupsResponseParams `json:"Response"`
}

func (r *DescribeAndroidInstanceBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAndroidInstanceBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAndroidInstanceImagesRequestParams struct {
	// 镜像 ID 列表
	AndroidInstanceImageIds []*string `json:"AndroidInstanceImageIds,omitnil,omitempty" name:"AndroidInstanceImageIds"`

	// 镜像可用区列表
	AndroidInstanceImageZones []*string `json:"AndroidInstanceImageZones,omitnil,omitempty" name:"AndroidInstanceImageZones"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制量，默认为20，最大值为100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 字段过滤器。Filter 的 Name 有以下值：
	// ImageName：镜像名称
	// ImageState：镜像状态
	// AndroidVersion：安卓版本
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeAndroidInstanceImagesRequest struct {
	*tchttp.BaseRequest
	
	// 镜像 ID 列表
	AndroidInstanceImageIds []*string `json:"AndroidInstanceImageIds,omitnil,omitempty" name:"AndroidInstanceImageIds"`

	// 镜像可用区列表
	AndroidInstanceImageZones []*string `json:"AndroidInstanceImageZones,omitnil,omitempty" name:"AndroidInstanceImageZones"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制量，默认为20，最大值为100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 字段过滤器。Filter 的 Name 有以下值：
	// ImageName：镜像名称
	// ImageState：镜像状态
	// AndroidVersion：安卓版本
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeAndroidInstanceImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAndroidInstanceImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceImageIds")
	delete(f, "AndroidInstanceImageZones")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAndroidInstanceImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAndroidInstanceImagesResponseParams struct {
	// 镜像总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 镜像列表
	AndroidInstanceImages []*AndroidInstanceImage `json:"AndroidInstanceImages,omitnil,omitempty" name:"AndroidInstanceImages"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAndroidInstanceImagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAndroidInstanceImagesResponseParams `json:"Response"`
}

func (r *DescribeAndroidInstanceImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAndroidInstanceImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAndroidInstanceLabelsRequestParams struct {
	// 标签键列表
	Keys []*string `json:"Keys,omitnil,omitempty" name:"Keys"`

	// 标签值列表
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 偏移量，默认为 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeAndroidInstanceLabelsRequest struct {
	*tchttp.BaseRequest
	
	// 标签键列表
	Keys []*string `json:"Keys,omitnil,omitempty" name:"Keys"`

	// 标签值列表
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 偏移量，默认为 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeAndroidInstanceLabelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAndroidInstanceLabelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Keys")
	delete(f, "Values")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAndroidInstanceLabelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAndroidInstanceLabelsResponseParams struct {
	// 安卓实例标签总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 安卓实例标签列表
	//
	// Deprecated: Labels is deprecated.
	Labels []*AndroidInstanceLabel `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 安卓实例标签详情列表
	AndroidInstanceLabels []*AndroidInstanceLabelDetail `json:"AndroidInstanceLabels,omitnil,omitempty" name:"AndroidInstanceLabels"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAndroidInstanceLabelsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAndroidInstanceLabelsResponseParams `json:"Response"`
}

func (r *DescribeAndroidInstanceLabelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAndroidInstanceLabelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAndroidInstanceTasksStatusRequestParams struct {
	// 任务 ID 列表
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 条件过滤器
	Filter []*Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 偏移量，默认为 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制量，默认为20，最大值为100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 时间范围限制，以天数为单位
	RecentDays *int64 `json:"RecentDays,omitnil,omitempty" name:"RecentDays"`
}

type DescribeAndroidInstanceTasksStatusRequest struct {
	*tchttp.BaseRequest
	
	// 任务 ID 列表
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 条件过滤器
	Filter []*Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 偏移量，默认为 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制量，默认为20，最大值为100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 时间范围限制，以天数为单位
	RecentDays *int64 `json:"RecentDays,omitnil,omitempty" name:"RecentDays"`
}

func (r *DescribeAndroidInstanceTasksStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAndroidInstanceTasksStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIds")
	delete(f, "Filter")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "RecentDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAndroidInstanceTasksStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAndroidInstanceTasksStatusResponseParams struct {
	// 任务状态集合
	TaskStatusSet []*AndroidInstanceTaskStatus `json:"TaskStatusSet,omitnil,omitempty" name:"TaskStatusSet"`

	// 任务总数量
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAndroidInstanceTasksStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAndroidInstanceTasksStatusResponseParams `json:"Response"`
}

func (r *DescribeAndroidInstanceTasksStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAndroidInstanceTasksStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAndroidInstancesAppBlacklistRequestParams struct {
	// 实例 ID 列表，数量上限 100
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`
}

type DescribeAndroidInstancesAppBlacklistRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表，数量上限 100
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`
}

func (r *DescribeAndroidInstancesAppBlacklistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAndroidInstancesAppBlacklistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAndroidInstancesAppBlacklistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAndroidInstancesAppBlacklistResponseParams struct {
	// 黑名单集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppBlacklistSet []*AndroidInstanceAppBlacklist `json:"AppBlacklistSet,omitnil,omitempty" name:"AppBlacklistSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAndroidInstancesAppBlacklistResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAndroidInstancesAppBlacklistResponseParams `json:"Response"`
}

func (r *DescribeAndroidInstancesAppBlacklistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAndroidInstancesAppBlacklistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAndroidInstancesByAppsRequestParams struct {
	// 偏移量，默认为 0	
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制量，默认为 20，最大值为 500	
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 应用 ID 列表。当 AndroidIds 为多条数据时（例如 app1, app2），返回的实例列表为：安装了 app1 应用的实例和安装了 app2 应用的实例集合（并集）。
	AndroidAppIds []*string `json:"AndroidAppIds,omitnil,omitempty" name:"AndroidAppIds"`

	// 字段过滤器，Filter 的 Name 有以下值： AndroidInstanceId：实例 Id
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeAndroidInstancesByAppsRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为 0	
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制量，默认为 20，最大值为 500	
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 应用 ID 列表。当 AndroidIds 为多条数据时（例如 app1, app2），返回的实例列表为：安装了 app1 应用的实例和安装了 app2 应用的实例集合（并集）。
	AndroidAppIds []*string `json:"AndroidAppIds,omitnil,omitempty" name:"AndroidAppIds"`

	// 字段过滤器，Filter 的 Name 有以下值： AndroidInstanceId：实例 Id
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeAndroidInstancesByAppsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAndroidInstancesByAppsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "AndroidAppIds")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAndroidInstancesByAppsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAndroidInstancesByAppsResponseParams struct {
	// 实例总数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例列表	
	AndroidInstances []*AndroidInstance `json:"AndroidInstances,omitnil,omitempty" name:"AndroidInstances"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAndroidInstancesByAppsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAndroidInstancesByAppsResponseParams `json:"Response"`
}

func (r *DescribeAndroidInstancesByAppsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAndroidInstancesByAppsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAndroidInstancesRequestParams struct {
	// 偏移量，默认为 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 实例ID。每次请求的实例的上限为100。
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 实例地域。目前还不支持按地域进行聚合查询
	AndroidInstanceRegion *string `json:"AndroidInstanceRegion,omitnil,omitempty" name:"AndroidInstanceRegion"`

	// 实例可用区
	AndroidInstanceZone *string `json:"AndroidInstanceZone,omitnil,omitempty" name:"AndroidInstanceZone"`

	// 实例分组 ID 列表
	AndroidInstanceGroupIds []*string `json:"AndroidInstanceGroupIds,omitnil,omitempty" name:"AndroidInstanceGroupIds"`

	// 实例标签选择器
	LabelSelector []*LabelRequirement `json:"LabelSelector,omitnil,omitempty" name:"LabelSelector"`

	// 字段过滤器。Filter 的 Name 有以下值：
	// Name：实例名称
	// UserId：实例用户ID
	// HostSerialNumber：宿主机序列号
	// HostServerSerialNumber：机箱序列号
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeAndroidInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 实例ID。每次请求的实例的上限为100。
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 实例地域。目前还不支持按地域进行聚合查询
	AndroidInstanceRegion *string `json:"AndroidInstanceRegion,omitnil,omitempty" name:"AndroidInstanceRegion"`

	// 实例可用区
	AndroidInstanceZone *string `json:"AndroidInstanceZone,omitnil,omitempty" name:"AndroidInstanceZone"`

	// 实例分组 ID 列表
	AndroidInstanceGroupIds []*string `json:"AndroidInstanceGroupIds,omitnil,omitempty" name:"AndroidInstanceGroupIds"`

	// 实例标签选择器
	LabelSelector []*LabelRequirement `json:"LabelSelector,omitnil,omitempty" name:"LabelSelector"`

	// 字段过滤器。Filter 的 Name 有以下值：
	// Name：实例名称
	// UserId：实例用户ID
	// HostSerialNumber：宿主机序列号
	// HostServerSerialNumber：机箱序列号
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeAndroidInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAndroidInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "AndroidInstanceIds")
	delete(f, "AndroidInstanceRegion")
	delete(f, "AndroidInstanceZone")
	delete(f, "AndroidInstanceGroupIds")
	delete(f, "LabelSelector")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAndroidInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAndroidInstancesResponseParams struct {
	// 实例总数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例列表
	AndroidInstances []*AndroidInstance `json:"AndroidInstances,omitnil,omitempty" name:"AndroidInstances"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAndroidInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAndroidInstancesResponseParams `json:"Response"`
}

func (r *DescribeAndroidInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAndroidInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesCountRequestParams struct {
	// 游戏ID
	GameId *string `json:"GameId,omitnil,omitempty" name:"GameId"`

	// 实例分组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 游戏区域
	GameRegion *string `json:"GameRegion,omitnil,omitempty" name:"GameRegion"`

	// 游戏类型。
	// MOBILE：手游
	// PC：默认值，端游
	GameType *string `json:"GameType,omitnil,omitempty" name:"GameType"`
}

type DescribeInstancesCountRequest struct {
	*tchttp.BaseRequest
	
	// 游戏ID
	GameId *string `json:"GameId,omitnil,omitempty" name:"GameId"`

	// 实例分组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 游戏区域
	GameRegion *string `json:"GameRegion,omitnil,omitempty" name:"GameRegion"`

	// 游戏类型。
	// MOBILE：手游
	// PC：默认值，端游
	GameType *string `json:"GameType,omitnil,omitempty" name:"GameType"`
}

func (r *DescribeInstancesCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GameId")
	delete(f, "GroupId")
	delete(f, "GameRegion")
	delete(f, "GameType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesCountResponseParams struct {
	// 客户的实例总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 客户的实例运行数
	Running *uint64 `json:"Running,omitnil,omitempty" name:"Running"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstancesCountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesCountResponseParams `json:"Response"`
}

func (r *DescribeInstancesCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyAndroidInstancesRequestParams struct {
	// 安卓实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`
}

type DestroyAndroidInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 安卓实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`
}

func (r *DestroyAndroidInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyAndroidInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyAndroidInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyAndroidInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyAndroidInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DestroyAndroidInstancesResponseParams `json:"Response"`
}

func (r *DestroyAndroidInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyAndroidInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableAndroidInstancesAppRequestParams struct {
	// 安卓实例 ID 列表（最大100条数据）
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 应用包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

type DisableAndroidInstancesAppRequest struct {
	*tchttp.BaseRequest
	
	// 安卓实例 ID 列表（最大100条数据）
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 应用包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

func (r *DisableAndroidInstancesAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableAndroidInstancesAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	delete(f, "PackageName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableAndroidInstancesAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableAndroidInstancesAppResponseParams struct {
	// 错误列表。如果实例操作都成功，则响应没有这个字段；如果有实例操作失败，该字段包含了实例操作的错误信息
	AndroidInstanceErrors []*AndroidInstanceError `json:"AndroidInstanceErrors,omitnil,omitempty" name:"AndroidInstanceErrors"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableAndroidInstancesAppResponse struct {
	*tchttp.BaseResponse
	Response *DisableAndroidInstancesAppResponseParams `json:"Response"`
}

func (r *DisableAndroidInstancesAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableAndroidInstancesAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisconnectAndroidInstanceAcceleratorRequestParams struct {
	// 实例ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`

	// 用户 ID。用户 ID 为空，将断开该实例的所有用户连接；用户 ID 不为空，只断开该用户的连接。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DisconnectAndroidInstanceAcceleratorRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`

	// 用户 ID。用户 ID 为空，将断开该实例的所有用户连接；用户 ID 不为空，只断开该用户的连接。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *DisconnectAndroidInstanceAcceleratorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisconnectAndroidInstanceAcceleratorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisconnectAndroidInstanceAcceleratorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisconnectAndroidInstanceAcceleratorResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisconnectAndroidInstanceAcceleratorResponse struct {
	*tchttp.BaseResponse
	Response *DisconnectAndroidInstanceAcceleratorResponseParams `json:"Response"`
}

func (r *DisconnectAndroidInstanceAcceleratorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisconnectAndroidInstanceAcceleratorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisconnectAndroidInstanceRequestParams struct {
	// 实例ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`
}

type DisconnectAndroidInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`
}

func (r *DisconnectAndroidInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisconnectAndroidInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisconnectAndroidInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisconnectAndroidInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisconnectAndroidInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DisconnectAndroidInstanceResponseParams `json:"Response"`
}

func (r *DisconnectAndroidInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisconnectAndroidInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DistributeAndroidInstanceImageToHostsRequestParams struct {
	// 宿主机序列号数组
	HostSerialNumbers []*string `json:"HostSerialNumbers,omitnil,omitempty" name:"HostSerialNumbers"`

	// 实例镜像 ID
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`
}

type DistributeAndroidInstanceImageToHostsRequest struct {
	*tchttp.BaseRequest
	
	// 宿主机序列号数组
	HostSerialNumbers []*string `json:"HostSerialNumbers,omitnil,omitempty" name:"HostSerialNumbers"`

	// 实例镜像 ID
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`
}

func (r *DistributeAndroidInstanceImageToHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DistributeAndroidInstanceImageToHostsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HostSerialNumbers")
	delete(f, "ImageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DistributeAndroidInstanceImageToHostsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DistributeAndroidInstanceImageToHostsResponseParams struct {
	// 任务集合
	TaskSet []*AndroidInstanceHostTask `json:"TaskSet,omitnil,omitempty" name:"TaskSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DistributeAndroidInstanceImageToHostsResponse struct {
	*tchttp.BaseResponse
	Response *DistributeAndroidInstanceImageToHostsResponseParams `json:"Response"`
}

func (r *DistributeAndroidInstanceImageToHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DistributeAndroidInstanceImageToHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DistributeFileToAndroidInstancesRequestParams struct {
	// 安卓实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 文件下载 URL
	FileURL *string `json:"FileURL,omitnil,omitempty" name:"FileURL"`

	// 上传目标目录，只能上传到 /sdcard/ 目录或其子目录下
	DestinationDirectory *string `json:"DestinationDirectory,omitnil,omitempty" name:"DestinationDirectory"`

	// 目标文件名
	DestinationFileName *string `json:"DestinationFileName,omitnil,omitempty" name:"DestinationFileName"`
}

type DistributeFileToAndroidInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 安卓实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 文件下载 URL
	FileURL *string `json:"FileURL,omitnil,omitempty" name:"FileURL"`

	// 上传目标目录，只能上传到 /sdcard/ 目录或其子目录下
	DestinationDirectory *string `json:"DestinationDirectory,omitnil,omitempty" name:"DestinationDirectory"`

	// 目标文件名
	DestinationFileName *string `json:"DestinationFileName,omitnil,omitempty" name:"DestinationFileName"`
}

func (r *DistributeFileToAndroidInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DistributeFileToAndroidInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	delete(f, "FileURL")
	delete(f, "DestinationDirectory")
	delete(f, "DestinationFileName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DistributeFileToAndroidInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DistributeFileToAndroidInstancesResponseParams struct {
	// 实例任务集合
	TaskSet []*AndroidInstanceTask `json:"TaskSet,omitnil,omitempty" name:"TaskSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DistributeFileToAndroidInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DistributeFileToAndroidInstancesResponseParams `json:"Response"`
}

func (r *DistributeFileToAndroidInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DistributeFileToAndroidInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DistributePhotoToAndroidInstancesRequestParams struct {
	// 安卓实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 照片下载 URL
	PhotoURL *string `json:"PhotoURL,omitnil,omitempty" name:"PhotoURL"`
}

type DistributePhotoToAndroidInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 安卓实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 照片下载 URL
	PhotoURL *string `json:"PhotoURL,omitnil,omitempty" name:"PhotoURL"`
}

func (r *DistributePhotoToAndroidInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DistributePhotoToAndroidInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	delete(f, "PhotoURL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DistributePhotoToAndroidInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DistributePhotoToAndroidInstancesResponseParams struct {
	// 实例任务集合
	TaskSet []*AndroidInstanceTask `json:"TaskSet,omitnil,omitempty" name:"TaskSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DistributePhotoToAndroidInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DistributePhotoToAndroidInstancesResponseParams `json:"Response"`
}

func (r *DistributePhotoToAndroidInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DistributePhotoToAndroidInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableAndroidInstancesAppRequestParams struct {
	// 安卓实例 ID 列表（最大100条数据）
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 应用包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

type EnableAndroidInstancesAppRequest struct {
	*tchttp.BaseRequest
	
	// 安卓实例 ID 列表（最大100条数据）
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 应用包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

func (r *EnableAndroidInstancesAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableAndroidInstancesAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	delete(f, "PackageName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableAndroidInstancesAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableAndroidInstancesAppResponseParams struct {
	// 错误列表。如果实例操作都成功，则响应没有这个字段；如果有实例操作失败，该字段包含了实例操作的错误信息
	AndroidInstanceErrors []*AndroidInstanceError `json:"AndroidInstanceErrors,omitnil,omitempty" name:"AndroidInstanceErrors"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableAndroidInstancesAppResponse struct {
	*tchttp.BaseResponse
	Response *EnableAndroidInstancesAppResponseParams `json:"Response"`
}

func (r *EnableAndroidInstancesAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableAndroidInstancesAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Error struct {
	// 错误码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 错误详细信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

// Predefined struct for user
type ExecuteCommandOnAndroidInstancesRequestParams struct {
	// 安卓实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// shell 命令
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`
}

type ExecuteCommandOnAndroidInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 安卓实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// shell 命令
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`
}

func (r *ExecuteCommandOnAndroidInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteCommandOnAndroidInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	delete(f, "Command")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExecuteCommandOnAndroidInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExecuteCommandOnAndroidInstancesResponseParams struct {
	// 任务集合，可异步查询任务状态
	TaskSet []*AndroidInstanceTask `json:"TaskSet,omitnil,omitempty" name:"TaskSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExecuteCommandOnAndroidInstancesResponse struct {
	*tchttp.BaseResponse
	Response *ExecuteCommandOnAndroidInstancesResponseParams `json:"Response"`
}

func (r *ExecuteCommandOnAndroidInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteCommandOnAndroidInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FetchAndroidInstancesLogsRequestParams struct {
	// 安卓实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// cos 桶名称
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// cos 桶区域
	BucketRegion *string `json:"BucketRegion,omitnil,omitempty" name:"BucketRegion"`

	// cos 桶目录，默认为 /log/
	BucketDirectory *string `json:"BucketDirectory,omitnil,omitempty" name:"BucketDirectory"`

	// 下载最近几天的日志，默认值为 1
	RecentDays *uint64 `json:"RecentDays,omitnil,omitempty" name:"RecentDays"`
}

type FetchAndroidInstancesLogsRequest struct {
	*tchttp.BaseRequest
	
	// 安卓实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// cos 桶名称
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// cos 桶区域
	BucketRegion *string `json:"BucketRegion,omitnil,omitempty" name:"BucketRegion"`

	// cos 桶目录，默认为 /log/
	BucketDirectory *string `json:"BucketDirectory,omitnil,omitempty" name:"BucketDirectory"`

	// 下载最近几天的日志，默认值为 1
	RecentDays *uint64 `json:"RecentDays,omitnil,omitempty" name:"RecentDays"`
}

func (r *FetchAndroidInstancesLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FetchAndroidInstancesLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	delete(f, "BucketName")
	delete(f, "BucketRegion")
	delete(f, "BucketDirectory")
	delete(f, "RecentDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FetchAndroidInstancesLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FetchAndroidInstancesLogsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FetchAndroidInstancesLogsResponse struct {
	*tchttp.BaseResponse
	Response *FetchAndroidInstancesLogsResponseParams `json:"Response"`
}

func (r *FetchAndroidInstancesLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FetchAndroidInstancesLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileCosInfo struct {
	// 文件 Id
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`
}

type Filter struct {
	// 字段名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字段值列表
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

// Predefined struct for user
type ImportAndroidInstanceImageRequestParams struct {
	// 镜像名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 镜像文件下载地址，要求是 tgz 压缩文件
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`

	// 镜像文件 MD5 值
	MD5 *string `json:"MD5,omitnil,omitempty" name:"MD5"`

	// 安卓版本。
	// ANDROID10：默认值，安卓10
	// ANDROID12：安卓12
	// ANDROID14：安卓14
	AndroidVersion *string `json:"AndroidVersion,omitnil,omitempty" name:"AndroidVersion"`

	// 镜像可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type ImportAndroidInstanceImageRequest struct {
	*tchttp.BaseRequest
	
	// 镜像名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 镜像文件下载地址，要求是 tgz 压缩文件
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`

	// 镜像文件 MD5 值
	MD5 *string `json:"MD5,omitnil,omitempty" name:"MD5"`

	// 安卓版本。
	// ANDROID10：默认值，安卓10
	// ANDROID12：安卓12
	// ANDROID14：安卓14
	AndroidVersion *string `json:"AndroidVersion,omitnil,omitempty" name:"AndroidVersion"`

	// 镜像可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

func (r *ImportAndroidInstanceImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportAndroidInstanceImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "URL")
	delete(f, "MD5")
	delete(f, "AndroidVersion")
	delete(f, "Zone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportAndroidInstanceImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportAndroidInstanceImageResponseParams struct {
	// 安卓实例镜像 ID
	AndroidInstanceImageId *string `json:"AndroidInstanceImageId,omitnil,omitempty" name:"AndroidInstanceImageId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImportAndroidInstanceImageResponse struct {
	*tchttp.BaseResponse
	Response *ImportAndroidInstanceImageResponseParams `json:"Response"`
}

func (r *ImportAndroidInstanceImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportAndroidInstanceImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InstallAndroidInstancesAppRequestParams struct {
	// 实例ID
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 应用ID
	AndroidAppId *string `json:"AndroidAppId,omitnil,omitempty" name:"AndroidAppId"`

	// 应用版本
	AndroidAppVersion *string `json:"AndroidAppVersion,omitnil,omitempty" name:"AndroidAppVersion"`

	// 安装方式。
	// CLEAR_DATA 默认，清理数据
	// KEEP_DATA 保留数据
	InstallationMethod *string `json:"InstallationMethod,omitnil,omitempty" name:"InstallationMethod"`
}

type InstallAndroidInstancesAppRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 应用ID
	AndroidAppId *string `json:"AndroidAppId,omitnil,omitempty" name:"AndroidAppId"`

	// 应用版本
	AndroidAppVersion *string `json:"AndroidAppVersion,omitnil,omitempty" name:"AndroidAppVersion"`

	// 安装方式。
	// CLEAR_DATA 默认，清理数据
	// KEEP_DATA 保留数据
	InstallationMethod *string `json:"InstallationMethod,omitnil,omitempty" name:"InstallationMethod"`
}

func (r *InstallAndroidInstancesAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InstallAndroidInstancesAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	delete(f, "AndroidAppId")
	delete(f, "AndroidAppVersion")
	delete(f, "InstallationMethod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InstallAndroidInstancesAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InstallAndroidInstancesAppResponseParams struct {
	// 任务集合
	TaskSet []*AndroidInstanceTask `json:"TaskSet,omitnil,omitempty" name:"TaskSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InstallAndroidInstancesAppResponse struct {
	*tchttp.BaseResponse
	Response *InstallAndroidInstancesAppResponseParams `json:"Response"`
}

func (r *InstallAndroidInstancesAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InstallAndroidInstancesAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InstallAndroidInstancesAppWithURLRequestParams struct {
	// 实例ID
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 安卓应用下载 URL
	AndroidAppURL *string `json:"AndroidAppURL,omitnil,omitempty" name:"AndroidAppURL"`
}

type InstallAndroidInstancesAppWithURLRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 安卓应用下载 URL
	AndroidAppURL *string `json:"AndroidAppURL,omitnil,omitempty" name:"AndroidAppURL"`
}

func (r *InstallAndroidInstancesAppWithURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InstallAndroidInstancesAppWithURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	delete(f, "AndroidAppURL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InstallAndroidInstancesAppWithURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InstallAndroidInstancesAppWithURLResponseParams struct {
	// 任务集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskSet []*AndroidInstanceTask `json:"TaskSet,omitnil,omitempty" name:"TaskSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InstallAndroidInstancesAppWithURLResponse struct {
	*tchttp.BaseResponse
	Response *InstallAndroidInstancesAppWithURLResponseParams `json:"Response"`
}

func (r *InstallAndroidInstancesAppWithURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InstallAndroidInstancesAppWithURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LabelRequirement struct {
	// 标签键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 运算符类型。
	// IN：要求对象的标签键 Key 对应的标签值需满足 Values 中的一个
	// NOT_IN：要求对象的标签键 Key 对应的标签值不满足 Values 中的任何一个
	// EXISTS：要求对象标签存在标签键 Key
	// NOT_EXISTS: 要求对象标签不存在标签键 Key
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 标签值列表
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

// Predefined struct for user
type ModifyAndroidAppRequestParams struct {
	// 安卓应用 Id
	AndroidAppId *string `json:"AndroidAppId,omitnil,omitempty" name:"AndroidAppId"`

	// 安卓应用名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户 Id
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type ModifyAndroidAppRequest struct {
	*tchttp.BaseRequest
	
	// 安卓应用 Id
	AndroidAppId *string `json:"AndroidAppId,omitnil,omitempty" name:"AndroidAppId"`

	// 安卓应用名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户 Id
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *ModifyAndroidAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAndroidAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidAppId")
	delete(f, "Name")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAndroidAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAndroidAppResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAndroidAppResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAndroidAppResponseParams `json:"Response"`
}

func (r *ModifyAndroidAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAndroidAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAndroidAppVersionRequestParams struct {
	// 安卓应用 Id
	AndroidAppId *string `json:"AndroidAppId,omitnil,omitempty" name:"AndroidAppId"`

	// 安卓应用版本 Id
	AndroidAppVersion *string `json:"AndroidAppVersion,omitnil,omitempty" name:"AndroidAppVersion"`

	// 安卓应用版本名称
	AndroidAppVersionName *string `json:"AndroidAppVersionName,omitnil,omitempty" name:"AndroidAppVersionName"`

	// 应用 shell 安装命令（支持多条命令执行，通过 && 组合；只在应用 AppMode 为 ADVANCED 高级模式下 才会生效）
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 应用 shell 卸载命令（支持多条命令执行，通过 && 组合；只在应用 AppMode 为 ADVANCED 高级模式下 才会生效）
	UninstallCommand *string `json:"UninstallCommand,omitnil,omitempty" name:"UninstallCommand"`
}

type ModifyAndroidAppVersionRequest struct {
	*tchttp.BaseRequest
	
	// 安卓应用 Id
	AndroidAppId *string `json:"AndroidAppId,omitnil,omitempty" name:"AndroidAppId"`

	// 安卓应用版本 Id
	AndroidAppVersion *string `json:"AndroidAppVersion,omitnil,omitempty" name:"AndroidAppVersion"`

	// 安卓应用版本名称
	AndroidAppVersionName *string `json:"AndroidAppVersionName,omitnil,omitempty" name:"AndroidAppVersionName"`

	// 应用 shell 安装命令（支持多条命令执行，通过 && 组合；只在应用 AppMode 为 ADVANCED 高级模式下 才会生效）
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 应用 shell 卸载命令（支持多条命令执行，通过 && 组合；只在应用 AppMode 为 ADVANCED 高级模式下 才会生效）
	UninstallCommand *string `json:"UninstallCommand,omitnil,omitempty" name:"UninstallCommand"`
}

func (r *ModifyAndroidAppVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAndroidAppVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidAppId")
	delete(f, "AndroidAppVersion")
	delete(f, "AndroidAppVersionName")
	delete(f, "Command")
	delete(f, "UninstallCommand")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAndroidAppVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAndroidAppVersionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAndroidAppVersionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAndroidAppVersionResponseParams `json:"Response"`
}

func (r *ModifyAndroidAppVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAndroidAppVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAndroidInstanceInformationRequestParams struct {
	// 安卓实例 ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`

	// 实例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type ModifyAndroidInstanceInformationRequest struct {
	*tchttp.BaseRequest
	
	// 安卓实例 ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`

	// 实例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *ModifyAndroidInstanceInformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAndroidInstanceInformationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAndroidInstanceInformationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAndroidInstanceInformationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAndroidInstanceInformationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAndroidInstanceInformationResponseParams `json:"Response"`
}

func (r *ModifyAndroidInstanceInformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAndroidInstanceInformationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAndroidInstanceResolutionRequestParams struct {
	// 安卓实例 ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`

	// 分辨率宽度。建议按照以下数值设置，避免出现性能不足问题：
	// 实例类型为单开（A1）：建议设置为 1080
	// 实例类型为双开（A2） 及以上：建议设置为 720
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 分辨率高度。建议按照以下数值设置，避免出现性能不足问题：
	// 实例类型为单开（A1）：建议设置为 1920
	// 实例类型为双开（A2） 及以上：建议设置为 1280
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 每英寸像素点。如果不填，系统将会计算一个合理的数值。修改 DPI 可能会导致 App 异常退出，请谨慎使用！
	// 分辨率为 720x1280：建议配置为 320
	// 分辨率为  1080x1920：建议配置为 480
	DPI *uint64 `json:"DPI,omitnil,omitempty" name:"DPI"`

	// 帧率。ResolutionType 为 PHYSICAL 时才会修改帧率。另外建议按照以下数值设置，避免出现性能不足问题： 实例类型为单开（A1）：建议设置为 60 实例类型为双开（A2） 及以上：建议设置为 30
	FPS *uint64 `json:"FPS,omitnil,omitempty" name:"FPS"`

	// 修改分辨率类型。修改物理分辨率，需要重启才能生效。
	// OVERRIDE：默认值，修改覆盖（显示）分辨率
	// PHYSICAL：修改物理分辨率
	ResolutionType *string `json:"ResolutionType,omitnil,omitempty" name:"ResolutionType"`
}

type ModifyAndroidInstanceResolutionRequest struct {
	*tchttp.BaseRequest
	
	// 安卓实例 ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`

	// 分辨率宽度。建议按照以下数值设置，避免出现性能不足问题：
	// 实例类型为单开（A1）：建议设置为 1080
	// 实例类型为双开（A2） 及以上：建议设置为 720
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 分辨率高度。建议按照以下数值设置，避免出现性能不足问题：
	// 实例类型为单开（A1）：建议设置为 1920
	// 实例类型为双开（A2） 及以上：建议设置为 1280
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 每英寸像素点。如果不填，系统将会计算一个合理的数值。修改 DPI 可能会导致 App 异常退出，请谨慎使用！
	// 分辨率为 720x1280：建议配置为 320
	// 分辨率为  1080x1920：建议配置为 480
	DPI *uint64 `json:"DPI,omitnil,omitempty" name:"DPI"`

	// 帧率。ResolutionType 为 PHYSICAL 时才会修改帧率。另外建议按照以下数值设置，避免出现性能不足问题： 实例类型为单开（A1）：建议设置为 60 实例类型为双开（A2） 及以上：建议设置为 30
	FPS *uint64 `json:"FPS,omitnil,omitempty" name:"FPS"`

	// 修改分辨率类型。修改物理分辨率，需要重启才能生效。
	// OVERRIDE：默认值，修改覆盖（显示）分辨率
	// PHYSICAL：修改物理分辨率
	ResolutionType *string `json:"ResolutionType,omitnil,omitempty" name:"ResolutionType"`
}

func (r *ModifyAndroidInstanceResolutionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAndroidInstanceResolutionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceId")
	delete(f, "Width")
	delete(f, "Height")
	delete(f, "DPI")
	delete(f, "FPS")
	delete(f, "ResolutionType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAndroidInstanceResolutionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAndroidInstanceResolutionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAndroidInstanceResolutionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAndroidInstanceResolutionResponseParams `json:"Response"`
}

func (r *ModifyAndroidInstanceResolutionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAndroidInstanceResolutionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAndroidInstancesAppBlacklistRequestParams struct {
	// 实例ID列表，数量上限100
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 应用列表
	AppList []*string `json:"AppList,omitnil,omitempty" name:"AppList"`

	// ADD、REMOVE、CLEAR
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`
}

type ModifyAndroidInstancesAppBlacklistRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID列表，数量上限100
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 应用列表
	AppList []*string `json:"AppList,omitnil,omitempty" name:"AppList"`

	// ADD、REMOVE、CLEAR
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`
}

func (r *ModifyAndroidInstancesAppBlacklistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAndroidInstancesAppBlacklistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	delete(f, "AppList")
	delete(f, "Operation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAndroidInstancesAppBlacklistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAndroidInstancesAppBlacklistResponseParams struct {
	// 任务集合
	TaskSet []*AndroidInstanceTask `json:"TaskSet,omitnil,omitempty" name:"TaskSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAndroidInstancesAppBlacklistResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAndroidInstancesAppBlacklistResponseParams `json:"Response"`
}

func (r *ModifyAndroidInstancesAppBlacklistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAndroidInstancesAppBlacklistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAndroidInstancesInformationRequestParams struct {
	// 安卓实例信息数据
	AndroidInstanceInformations []*AndroidInstanceInformation `json:"AndroidInstanceInformations,omitnil,omitempty" name:"AndroidInstanceInformations"`
}

type ModifyAndroidInstancesInformationRequest struct {
	*tchttp.BaseRequest
	
	// 安卓实例信息数据
	AndroidInstanceInformations []*AndroidInstanceInformation `json:"AndroidInstanceInformations,omitnil,omitempty" name:"AndroidInstanceInformations"`
}

func (r *ModifyAndroidInstancesInformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAndroidInstancesInformationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceInformations")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAndroidInstancesInformationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAndroidInstancesInformationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAndroidInstancesInformationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAndroidInstancesInformationResponseParams `json:"Response"`
}

func (r *ModifyAndroidInstancesInformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAndroidInstancesInformationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAndroidInstancesLabelsRequestParams struct {
	// 安卓实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 操作类型。ADD：标签键不存在的添加新标签，标签键存在的将覆盖原有标签REMOVE：根据标签键删除标签REPLACE：使用请求标签列表替换原来所有标签CLEAR：清除所有标签
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 安卓实例标签列表
	AndroidInstanceLabels []*AndroidInstanceLabel `json:"AndroidInstanceLabels,omitnil,omitempty" name:"AndroidInstanceLabels"`
}

type ModifyAndroidInstancesLabelsRequest struct {
	*tchttp.BaseRequest
	
	// 安卓实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 操作类型。ADD：标签键不存在的添加新标签，标签键存在的将覆盖原有标签REMOVE：根据标签键删除标签REPLACE：使用请求标签列表替换原来所有标签CLEAR：清除所有标签
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 安卓实例标签列表
	AndroidInstanceLabels []*AndroidInstanceLabel `json:"AndroidInstanceLabels,omitnil,omitempty" name:"AndroidInstanceLabels"`
}

func (r *ModifyAndroidInstancesLabelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAndroidInstancesLabelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	delete(f, "Operation")
	delete(f, "AndroidInstanceLabels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAndroidInstancesLabelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAndroidInstancesLabelsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAndroidInstancesLabelsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAndroidInstancesLabelsResponseParams `json:"Response"`
}

func (r *ModifyAndroidInstancesLabelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAndroidInstancesLabelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAndroidInstancesPropertiesRequestParams struct {
	// 安卓实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 安卓实例设备
	AndroidInstanceDevice *AndroidInstanceDevice `json:"AndroidInstanceDevice,omitnil,omitempty" name:"AndroidInstanceDevice"`

	// 安卓实例其它属性列表
	AndroidInstanceProperties []*AndroidInstanceProperty `json:"AndroidInstanceProperties,omitnil,omitempty" name:"AndroidInstanceProperties"`
}

type ModifyAndroidInstancesPropertiesRequest struct {
	*tchttp.BaseRequest
	
	// 安卓实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 安卓实例设备
	AndroidInstanceDevice *AndroidInstanceDevice `json:"AndroidInstanceDevice,omitnil,omitempty" name:"AndroidInstanceDevice"`

	// 安卓实例其它属性列表
	AndroidInstanceProperties []*AndroidInstanceProperty `json:"AndroidInstanceProperties,omitnil,omitempty" name:"AndroidInstanceProperties"`
}

func (r *ModifyAndroidInstancesPropertiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAndroidInstancesPropertiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	delete(f, "AndroidInstanceDevice")
	delete(f, "AndroidInstanceProperties")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAndroidInstancesPropertiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAndroidInstancesPropertiesResponseParams struct {
	// 安卓实例错误列表
	AndroidInstanceErrors []*AndroidInstanceError `json:"AndroidInstanceErrors,omitnil,omitempty" name:"AndroidInstanceErrors"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAndroidInstancesPropertiesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAndroidInstancesPropertiesResponseParams `json:"Response"`
}

func (r *ModifyAndroidInstancesPropertiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAndroidInstancesPropertiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAndroidInstancesResolutionRequestParams struct {
	// 安卓实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 分辨率宽度。建议按照以下数值设置，避免出现性能不足问题：
	// 实例类型为单开（A1）：建议设置为 1080
	// 实例类型为双开（A2） 及以上：建议设置为 720
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 分辨率高度。建议按照以下数值设置，避免出现性能不足问题：
	// 实例类型为单开（A1）：建议设置为 1920
	// 实例类型为双开（A2） 及以上：建议设置为 1280
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 每英寸像素点。
	// 分辨率为 720x1280：建议配置为 320
	// 分辨率为  1080x1920：建议配置为 480
	DPI *uint64 `json:"DPI,omitnil,omitempty" name:"DPI"`

	// 帧率。ResolutionType 为 PHYSICAL 时才会修改帧率。另外建议按照以下数值设置，避免出现性能不足问题：
	// 实例类型为单开（A1）：建议设置为 60
	// 实例类型为双开（A2） 及以上：建议设置为 30
	FPS *uint64 `json:"FPS,omitnil,omitempty" name:"FPS"`

	// 修改分辨率类型。修改物理分辨率，需要重启才能生效。
	// OVERRIDE：默认值，修改覆盖（显示）分辨率
	// PHYSICAL：修改物理分辨率
	ResolutionType *string `json:"ResolutionType,omitnil,omitempty" name:"ResolutionType"`
}

type ModifyAndroidInstancesResolutionRequest struct {
	*tchttp.BaseRequest
	
	// 安卓实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 分辨率宽度。建议按照以下数值设置，避免出现性能不足问题：
	// 实例类型为单开（A1）：建议设置为 1080
	// 实例类型为双开（A2） 及以上：建议设置为 720
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 分辨率高度。建议按照以下数值设置，避免出现性能不足问题：
	// 实例类型为单开（A1）：建议设置为 1920
	// 实例类型为双开（A2） 及以上：建议设置为 1280
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 每英寸像素点。
	// 分辨率为 720x1280：建议配置为 320
	// 分辨率为  1080x1920：建议配置为 480
	DPI *uint64 `json:"DPI,omitnil,omitempty" name:"DPI"`

	// 帧率。ResolutionType 为 PHYSICAL 时才会修改帧率。另外建议按照以下数值设置，避免出现性能不足问题：
	// 实例类型为单开（A1）：建议设置为 60
	// 实例类型为双开（A2） 及以上：建议设置为 30
	FPS *uint64 `json:"FPS,omitnil,omitempty" name:"FPS"`

	// 修改分辨率类型。修改物理分辨率，需要重启才能生效。
	// OVERRIDE：默认值，修改覆盖（显示）分辨率
	// PHYSICAL：修改物理分辨率
	ResolutionType *string `json:"ResolutionType,omitnil,omitempty" name:"ResolutionType"`
}

func (r *ModifyAndroidInstancesResolutionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAndroidInstancesResolutionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	delete(f, "Width")
	delete(f, "Height")
	delete(f, "DPI")
	delete(f, "FPS")
	delete(f, "ResolutionType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAndroidInstancesResolutionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAndroidInstancesResolutionResponseParams struct {
	// 安卓实例错误列表
	AndroidInstanceErrors []*AndroidInstanceError `json:"AndroidInstanceErrors,omitnil,omitempty" name:"AndroidInstanceErrors"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAndroidInstancesResolutionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAndroidInstancesResolutionResponseParams `json:"Response"`
}

func (r *ModifyAndroidInstancesResolutionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAndroidInstancesResolutionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAndroidInstancesResourcesRequestParams struct {
	// 安卓实例 ID 列表（最大100条数据）
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 内存配额（单位 MB）
	MemoryQuota *int64 `json:"MemoryQuota,omitnil,omitempty" name:"MemoryQuota"`
}

type ModifyAndroidInstancesResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 安卓实例 ID 列表（最大100条数据）
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 内存配额（单位 MB）
	MemoryQuota *int64 `json:"MemoryQuota,omitnil,omitempty" name:"MemoryQuota"`
}

func (r *ModifyAndroidInstancesResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAndroidInstancesResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	delete(f, "MemoryQuota")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAndroidInstancesResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAndroidInstancesResourcesResponseParams struct {
	// 任务集合
	TaskSet []*AndroidInstanceTask `json:"TaskSet,omitnil,omitempty" name:"TaskSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAndroidInstancesResourcesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAndroidInstancesResourcesResponseParams `json:"Response"`
}

func (r *ModifyAndroidInstancesResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAndroidInstancesResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAndroidInstancesUserIdRequestParams struct {
	// 安卓实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 用户 ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 有效时长。如果不填该字段，默认为永久。支持 s（秒）、m（分）、h（小时）、d（天）等单位，比如 12h 表示 12 小时，1h2m3s 表示一小时两分三秒
	ExpirationDuration *string `json:"ExpirationDuration,omitnil,omitempty" name:"ExpirationDuration"`
}

type ModifyAndroidInstancesUserIdRequest struct {
	*tchttp.BaseRequest
	
	// 安卓实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 用户 ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 有效时长。如果不填该字段，默认为永久。支持 s（秒）、m（分）、h（小时）、d（天）等单位，比如 12h 表示 12 小时，1h2m3s 表示一小时两分三秒
	ExpirationDuration *string `json:"ExpirationDuration,omitnil,omitempty" name:"ExpirationDuration"`
}

func (r *ModifyAndroidInstancesUserIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAndroidInstancesUserIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	delete(f, "UserId")
	delete(f, "ExpirationDuration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAndroidInstancesUserIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAndroidInstancesUserIdResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAndroidInstancesUserIdResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAndroidInstancesUserIdResponseParams `json:"Response"`
}

func (r *ModifyAndroidInstancesUserIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAndroidInstancesUserIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RebootAndroidInstanceHostsRequestParams struct {
	// 宿主机序列号集合
	HostSerialNumbers []*string `json:"HostSerialNumbers,omitnil,omitempty" name:"HostSerialNumbers"`
}

type RebootAndroidInstanceHostsRequest struct {
	*tchttp.BaseRequest
	
	// 宿主机序列号集合
	HostSerialNumbers []*string `json:"HostSerialNumbers,omitnil,omitempty" name:"HostSerialNumbers"`
}

func (r *RebootAndroidInstanceHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RebootAndroidInstanceHostsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HostSerialNumbers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RebootAndroidInstanceHostsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RebootAndroidInstanceHostsResponseParams struct {
	// 任务 ID 集合，以供任务状态查询，其中 InstanceId 为宿主机序列号
	TaskSet []*AndroidInstanceTask `json:"TaskSet,omitnil,omitempty" name:"TaskSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RebootAndroidInstanceHostsResponse struct {
	*tchttp.BaseResponse
	Response *RebootAndroidInstanceHostsResponseParams `json:"Response"`
}

func (r *RebootAndroidInstanceHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RebootAndroidInstanceHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RebootAndroidInstancesRequestParams struct {
	// 实例ID
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`
}

type RebootAndroidInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`
}

func (r *RebootAndroidInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RebootAndroidInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RebootAndroidInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RebootAndroidInstancesResponseParams struct {
	// 任务集合
	TaskSet []*AndroidInstanceTask `json:"TaskSet,omitnil,omitempty" name:"TaskSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RebootAndroidInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RebootAndroidInstancesResponseParams `json:"Response"`
}

func (r *RebootAndroidInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RebootAndroidInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewAndroidInstancesAccessTokenRequestParams struct {
	// token
	AccessToken *string `json:"AccessToken,omitnil,omitempty" name:"AccessToken"`

	// 有效期，默认为 12 小时，最大为 24 小时。支持 s（秒）、m（分）、h（小时）等单位，比如 12h 表示 12 小时，1h2m3s 表示一小时两分三秒
	ExpirationDuration *string `json:"ExpirationDuration,omitnil,omitempty" name:"ExpirationDuration"`
}

type RenewAndroidInstancesAccessTokenRequest struct {
	*tchttp.BaseRequest
	
	// token
	AccessToken *string `json:"AccessToken,omitnil,omitempty" name:"AccessToken"`

	// 有效期，默认为 12 小时，最大为 24 小时。支持 s（秒）、m（分）、h（小时）等单位，比如 12h 表示 12 小时，1h2m3s 表示一小时两分三秒
	ExpirationDuration *string `json:"ExpirationDuration,omitnil,omitempty" name:"ExpirationDuration"`
}

func (r *RenewAndroidInstancesAccessTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewAndroidInstancesAccessTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessToken")
	delete(f, "ExpirationDuration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewAndroidInstancesAccessTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewAndroidInstancesAccessTokenResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewAndroidInstancesAccessTokenResponse struct {
	*tchttp.BaseResponse
	Response *RenewAndroidInstancesAccessTokenResponseParams `json:"Response"`
}

func (r *RenewAndroidInstancesAccessTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewAndroidInstancesAccessTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetAndroidInstancesRequestParams struct {
	// 实例ID列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 指定有效的镜像 ID。
	// 默认取值：默认使用当前镜像。
	AndroidInstanceImageId *string `json:"AndroidInstanceImageId,omitnil,omitempty" name:"AndroidInstanceImageId"`

	// 重置模式。在 AndroidInstanceImageId 不为空时才生效。
	// 
	// CleanData：默认选项，清理系统属性和用户数据
	// KeepSystemProperties：只保留系统属性
	// KeepData: 保留系统属性和用户数据
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`
}

type ResetAndroidInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 指定有效的镜像 ID。
	// 默认取值：默认使用当前镜像。
	AndroidInstanceImageId *string `json:"AndroidInstanceImageId,omitnil,omitempty" name:"AndroidInstanceImageId"`

	// 重置模式。在 AndroidInstanceImageId 不为空时才生效。
	// 
	// CleanData：默认选项，清理系统属性和用户数据
	// KeepSystemProperties：只保留系统属性
	// KeepData: 保留系统属性和用户数据
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`
}

func (r *ResetAndroidInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAndroidInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	delete(f, "AndroidInstanceImageId")
	delete(f, "Mode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetAndroidInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetAndroidInstancesResponseParams struct {
	// 任务集合
	TaskSet []*AndroidInstanceTask `json:"TaskSet,omitnil,omitempty" name:"TaskSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetAndroidInstancesResponse struct {
	*tchttp.BaseResponse
	Response *ResetAndroidInstancesResponseParams `json:"Response"`
}

func (r *ResetAndroidInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAndroidInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartAndroidInstancesAppRequestParams struct {
	// 实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 应用包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 启动页。建议指定启动页来启动应用，避免启动失败。如果启动页为空，系统尝试根据 PackageName 启动，但不保证成功。
	Activity *string `json:"Activity,omitnil,omitempty" name:"Activity"`
}

type RestartAndroidInstancesAppRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 应用包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 启动页。建议指定启动页来启动应用，避免启动失败。如果启动页为空，系统尝试根据 PackageName 启动，但不保证成功。
	Activity *string `json:"Activity,omitnil,omitempty" name:"Activity"`
}

func (r *RestartAndroidInstancesAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartAndroidInstancesAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	delete(f, "PackageName")
	delete(f, "Activity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartAndroidInstancesAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartAndroidInstancesAppResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartAndroidInstancesAppResponse struct {
	*tchttp.BaseResponse
	Response *RestartAndroidInstancesAppResponseParams `json:"Response"`
}

func (r *RestartAndroidInstancesAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartAndroidInstancesAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestoreAndroidInstanceFromStorageRequestParams struct {
	// 安卓实例ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`

	// 自定义备份对象Key
	ObjectKey *string `json:"ObjectKey,omitnil,omitempty" name:"ObjectKey"`

	// 存储服务器类型，如 COS、S3。注意：使用 COS 和 S3 都将占用外网带宽。
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// COS协议选项
	COSOptions *COSOptions `json:"COSOptions,omitnil,omitempty" name:"COSOptions"`

	// S3存储协议选项
	S3Options *S3Options `json:"S3Options,omitnil,omitempty" name:"S3Options"`
}

type RestoreAndroidInstanceFromStorageRequest struct {
	*tchttp.BaseRequest
	
	// 安卓实例ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`

	// 自定义备份对象Key
	ObjectKey *string `json:"ObjectKey,omitnil,omitempty" name:"ObjectKey"`

	// 存储服务器类型，如 COS、S3。注意：使用 COS 和 S3 都将占用外网带宽。
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// COS协议选项
	COSOptions *COSOptions `json:"COSOptions,omitnil,omitempty" name:"COSOptions"`

	// S3存储协议选项
	S3Options *S3Options `json:"S3Options,omitnil,omitempty" name:"S3Options"`
}

func (r *RestoreAndroidInstanceFromStorageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreAndroidInstanceFromStorageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceId")
	delete(f, "ObjectKey")
	delete(f, "StorageType")
	delete(f, "COSOptions")
	delete(f, "S3Options")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestoreAndroidInstanceFromStorageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestoreAndroidInstanceFromStorageResponseParams struct {
	// 实例任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestoreAndroidInstanceFromStorageResponse struct {
	*tchttp.BaseResponse
	Response *RestoreAndroidInstanceFromStorageResponseParams `json:"Response"`
}

func (r *RestoreAndroidInstanceFromStorageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreAndroidInstanceFromStorageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestoreAndroidInstanceRequestParams struct {
	// 安卓实例 ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`

	// 备份 ID
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`
}

type RestoreAndroidInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 安卓实例 ID
	AndroidInstanceId *string `json:"AndroidInstanceId,omitnil,omitempty" name:"AndroidInstanceId"`

	// 备份 ID
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`
}

func (r *RestoreAndroidInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreAndroidInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceId")
	delete(f, "BackupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestoreAndroidInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestoreAndroidInstanceResponseParams struct {
	// 实例任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestoreAndroidInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RestoreAndroidInstanceResponseParams `json:"Response"`
}

func (r *RestoreAndroidInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreAndroidInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type S3Options struct {
	// 存储节点
	EndPoint *string `json:"EndPoint,omitnil,omitempty" name:"EndPoint"`

	// 存储桶
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 密钥 ID
	AccessKeyId *string `json:"AccessKeyId,omitnil,omitempty" name:"AccessKeyId"`

	// 密钥 Key
	SecretAccessKey *string `json:"SecretAccessKey,omitnil,omitempty" name:"SecretAccessKey"`
}

// Predefined struct for user
type SaveGameArchiveRequestParams struct {
	// 游戏用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 游戏ID
	GameId *string `json:"GameId,omitnil,omitempty" name:"GameId"`
}

type SaveGameArchiveRequest struct {
	*tchttp.BaseRequest
	
	// 游戏用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 游戏ID
	GameId *string `json:"GameId,omitnil,omitempty" name:"GameId"`
}

func (r *SaveGameArchiveRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SaveGameArchiveRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "GameId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SaveGameArchiveRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SaveGameArchiveResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SaveGameArchiveResponse struct {
	*tchttp.BaseResponse
	Response *SaveGameArchiveResponseParams `json:"Response"`
}

func (r *SaveGameArchiveResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SaveGameArchiveResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetAndroidInstancesBGAppKeepAliveRequestParams struct {
	// 安卓实例 ID 列表（最大100条数据）
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 操作类型，取值：ADD 添加应用到后台保活列表、REMOVE 从后台保活列表中移除应用、SET 全量设置后台保活列表，替换当前列表。
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 应用包名列表
	PackageNames []*string `json:"PackageNames,omitnil,omitempty" name:"PackageNames"`
}

type SetAndroidInstancesBGAppKeepAliveRequest struct {
	*tchttp.BaseRequest
	
	// 安卓实例 ID 列表（最大100条数据）
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 操作类型，取值：ADD 添加应用到后台保活列表、REMOVE 从后台保活列表中移除应用、SET 全量设置后台保活列表，替换当前列表。
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 应用包名列表
	PackageNames []*string `json:"PackageNames,omitnil,omitempty" name:"PackageNames"`
}

func (r *SetAndroidInstancesBGAppKeepAliveRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetAndroidInstancesBGAppKeepAliveRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	delete(f, "Operation")
	delete(f, "PackageNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetAndroidInstancesBGAppKeepAliveRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetAndroidInstancesBGAppKeepAliveResponseParams struct {
	// 错误列表。如果实例操作都成功，则响应没有这个字段；如果有实例操作失败，该字段包含了实例操作的错误信息
	AndroidInstanceErrors []*AndroidInstanceError `json:"AndroidInstanceErrors,omitnil,omitempty" name:"AndroidInstanceErrors"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetAndroidInstancesBGAppKeepAliveResponse struct {
	*tchttp.BaseResponse
	Response *SetAndroidInstancesBGAppKeepAliveResponseParams `json:"Response"`
}

func (r *SetAndroidInstancesBGAppKeepAliveResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetAndroidInstancesBGAppKeepAliveResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetAndroidInstancesFGAppKeepAliveRequestParams struct {
	// 安卓实例 ID 列表（最大100条数据）
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 操作类型，取值：ENABLE 开启保活、DISABLE 关闭保活。当关闭保活时，PackageName 参数传空即可
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 应用包名，开启保活时，必须传入 PackageName
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

type SetAndroidInstancesFGAppKeepAliveRequest struct {
	*tchttp.BaseRequest
	
	// 安卓实例 ID 列表（最大100条数据）
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 操作类型，取值：ENABLE 开启保活、DISABLE 关闭保活。当关闭保活时，PackageName 参数传空即可
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 应用包名，开启保活时，必须传入 PackageName
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

func (r *SetAndroidInstancesFGAppKeepAliveRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetAndroidInstancesFGAppKeepAliveRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	delete(f, "Operation")
	delete(f, "PackageName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetAndroidInstancesFGAppKeepAliveRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetAndroidInstancesFGAppKeepAliveResponseParams struct {
	// 错误列表。如果实例操作都成功，则响应没有这个字段；如果有实例操作失败，该字段包含了实例操作的错误信息
	AndroidInstanceErrors []*AndroidInstanceError `json:"AndroidInstanceErrors,omitnil,omitempty" name:"AndroidInstanceErrors"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetAndroidInstancesFGAppKeepAliveResponse struct {
	*tchttp.BaseResponse
	Response *SetAndroidInstancesFGAppKeepAliveResponseParams `json:"Response"`
}

func (r *SetAndroidInstancesFGAppKeepAliveResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetAndroidInstancesFGAppKeepAliveResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartAndroidInstancesAppRequestParams struct {
	// 实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 应用包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 启动页。建议指定启动页来启动应用，避免启动失败。如果启动页为空，系统尝试根据 PackageName 启动，但不保证成功。
	Activity *string `json:"Activity,omitnil,omitempty" name:"Activity"`
}

type StartAndroidInstancesAppRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 应用包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 启动页。建议指定启动页来启动应用，避免启动失败。如果启动页为空，系统尝试根据 PackageName 启动，但不保证成功。
	Activity *string `json:"Activity,omitnil,omitempty" name:"Activity"`
}

func (r *StartAndroidInstancesAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartAndroidInstancesAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	delete(f, "PackageName")
	delete(f, "Activity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartAndroidInstancesAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartAndroidInstancesAppResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartAndroidInstancesAppResponse struct {
	*tchttp.BaseResponse
	Response *StartAndroidInstancesAppResponseParams `json:"Response"`
}

func (r *StartAndroidInstancesAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartAndroidInstancesAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartAndroidInstancesRequestParams struct {
	// 实例ID
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`
}

type StartAndroidInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`
}

func (r *StartAndroidInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartAndroidInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartAndroidInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartAndroidInstancesResponseParams struct {
	// 任务集合
	TaskSet []*AndroidInstanceTask `json:"TaskSet,omitnil,omitempty" name:"TaskSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartAndroidInstancesResponse struct {
	*tchttp.BaseResponse
	Response *StartAndroidInstancesResponseParams `json:"Response"`
}

func (r *StartAndroidInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartAndroidInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartPublishStreamRequestParams struct {
	// 唯一用户身份标识，由业务方自定义，平台不予理解。（可根据业务需要决定使用用户的唯一身份标识或是使用时间戳随机生成；在用户重连时应保持UserId不变）
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 推流地址，仅支持rtmp协议
	PublishUrl *string `json:"PublishUrl,omitnil,omitempty" name:"PublishUrl"`
}

type StartPublishStreamRequest struct {
	*tchttp.BaseRequest
	
	// 唯一用户身份标识，由业务方自定义，平台不予理解。（可根据业务需要决定使用用户的唯一身份标识或是使用时间戳随机生成；在用户重连时应保持UserId不变）
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 推流地址，仅支持rtmp协议
	PublishUrl *string `json:"PublishUrl,omitnil,omitempty" name:"PublishUrl"`
}

func (r *StartPublishStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartPublishStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "PublishUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartPublishStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartPublishStreamResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartPublishStreamResponse struct {
	*tchttp.BaseResponse
	Response *StartPublishStreamResponseParams `json:"Response"`
}

func (r *StartPublishStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartPublishStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartPublishStreamToCSSRequestParams struct {
	// 唯一用户身份标识，由业务方自定义，平台不予理解。（可根据业务需要决定使用用户的唯一身份标识或是使用时间戳随机生成；在用户重连时应保持UserId不变）
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 推流参数，推流时携带自定义参数。
	PublishStreamArgs *string `json:"PublishStreamArgs,omitnil,omitempty" name:"PublishStreamArgs"`
}

type StartPublishStreamToCSSRequest struct {
	*tchttp.BaseRequest
	
	// 唯一用户身份标识，由业务方自定义，平台不予理解。（可根据业务需要决定使用用户的唯一身份标识或是使用时间戳随机生成；在用户重连时应保持UserId不变）
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 推流参数，推流时携带自定义参数。
	PublishStreamArgs *string `json:"PublishStreamArgs,omitnil,omitempty" name:"PublishStreamArgs"`
}

func (r *StartPublishStreamToCSSRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartPublishStreamToCSSRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "PublishStreamArgs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartPublishStreamToCSSRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartPublishStreamToCSSResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartPublishStreamToCSSResponse struct {
	*tchttp.BaseResponse
	Response *StartPublishStreamToCSSResponseParams `json:"Response"`
}

func (r *StartPublishStreamToCSSResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartPublishStreamToCSSResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopAndroidInstancesAppRequestParams struct {
	// 实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 应用包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

type StopAndroidInstancesAppRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 应用包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

func (r *StopAndroidInstancesAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopAndroidInstancesAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	delete(f, "PackageName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopAndroidInstancesAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopAndroidInstancesAppResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopAndroidInstancesAppResponse struct {
	*tchttp.BaseResponse
	Response *StopAndroidInstancesAppResponseParams `json:"Response"`
}

func (r *StopAndroidInstancesAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopAndroidInstancesAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopAndroidInstancesRequestParams struct {
	// 实例ID
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`
}

type StopAndroidInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`
}

func (r *StopAndroidInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopAndroidInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopAndroidInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopAndroidInstancesResponseParams struct {
	// 任务集合
	TaskSet []*AndroidInstanceTask `json:"TaskSet,omitnil,omitempty" name:"TaskSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopAndroidInstancesResponse struct {
	*tchttp.BaseResponse
	Response *StopAndroidInstancesResponseParams `json:"Response"`
}

func (r *StopAndroidInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopAndroidInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopGameRequestParams struct {
	// 唯一用户身份标识，由业务方自定义，平台不予理解。（可根据业务需要决定使用用户的唯一身份标识或是使用时间戳随机生成；在用户重连时应保持UserId不变）
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 【多人游戏】游戏主机用户ID
	HostUserId *string `json:"HostUserId,omitnil,omitempty" name:"HostUserId"`
}

type StopGameRequest struct {
	*tchttp.BaseRequest
	
	// 唯一用户身份标识，由业务方自定义，平台不予理解。（可根据业务需要决定使用用户的唯一身份标识或是使用时间戳随机生成；在用户重连时应保持UserId不变）
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 【多人游戏】游戏主机用户ID
	HostUserId *string `json:"HostUserId,omitnil,omitempty" name:"HostUserId"`
}

func (r *StopGameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopGameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "HostUserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopGameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopGameResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopGameResponse struct {
	*tchttp.BaseResponse
	Response *StopGameResponseParams `json:"Response"`
}

func (r *StopGameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopGameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopPublishStreamRequestParams struct {
	// 唯一用户身份标识，由业务方自定义，平台不予理解。（可根据业务需要决定使用用户的唯一身份标识或是使用时间戳随机生成；在用户重连时应保持UserId不变）
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type StopPublishStreamRequest struct {
	*tchttp.BaseRequest
	
	// 唯一用户身份标识，由业务方自定义，平台不予理解。（可根据业务需要决定使用用户的唯一身份标识或是使用时间戳随机生成；在用户重连时应保持UserId不变）
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *StopPublishStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopPublishStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopPublishStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopPublishStreamResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopPublishStreamResponse struct {
	*tchttp.BaseResponse
	Response *StopPublishStreamResponseParams `json:"Response"`
}

func (r *StopPublishStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopPublishStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchGameArchiveRequestParams struct {
	// 游戏用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 游戏ID
	GameId *string `json:"GameId,omitnil,omitempty" name:"GameId"`

	// 游戏存档Url
	GameArchiveUrl *string `json:"GameArchiveUrl,omitnil,omitempty" name:"GameArchiveUrl"`

	// 游戏相关参数
	GameContext *string `json:"GameContext,omitnil,omitempty" name:"GameContext"`
}

type SwitchGameArchiveRequest struct {
	*tchttp.BaseRequest
	
	// 游戏用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 游戏ID
	GameId *string `json:"GameId,omitnil,omitempty" name:"GameId"`

	// 游戏存档Url
	GameArchiveUrl *string `json:"GameArchiveUrl,omitnil,omitempty" name:"GameArchiveUrl"`

	// 游戏相关参数
	GameContext *string `json:"GameContext,omitnil,omitempty" name:"GameContext"`
}

func (r *SwitchGameArchiveRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchGameArchiveRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "GameId")
	delete(f, "GameArchiveUrl")
	delete(f, "GameContext")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchGameArchiveRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchGameArchiveResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SwitchGameArchiveResponse struct {
	*tchttp.BaseResponse
	Response *SwitchGameArchiveResponseParams `json:"Response"`
}

func (r *SwitchGameArchiveResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchGameArchiveResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncAndroidInstanceImage struct {
	// 镜像 ID
	AndroidInstanceImageId *string `json:"AndroidInstanceImageId,omitnil,omitempty" name:"AndroidInstanceImageId"`

	// 镜像可用区
	AndroidInstanceImageZone *string `json:"AndroidInstanceImageZone,omitnil,omitempty" name:"AndroidInstanceImageZone"`
}

// Predefined struct for user
type SyncAndroidInstanceImageRequestParams struct {
	// 安卓实例镜像 ID
	AndroidInstanceImageId *string `json:"AndroidInstanceImageId,omitnil,omitempty" name:"AndroidInstanceImageId"`

	// 目的同步可用区列表
	DestinationZones []*string `json:"DestinationZones,omitnil,omitempty" name:"DestinationZones"`
}

type SyncAndroidInstanceImageRequest struct {
	*tchttp.BaseRequest
	
	// 安卓实例镜像 ID
	AndroidInstanceImageId *string `json:"AndroidInstanceImageId,omitnil,omitempty" name:"AndroidInstanceImageId"`

	// 目的同步可用区列表
	DestinationZones []*string `json:"DestinationZones,omitnil,omitempty" name:"DestinationZones"`
}

func (r *SyncAndroidInstanceImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncAndroidInstanceImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceImageId")
	delete(f, "DestinationZones")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncAndroidInstanceImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncAndroidInstanceImageResponseParams struct {
	// 同步安卓实例镜像列表
	SyncAndroidInstanceImages []*SyncAndroidInstanceImage `json:"SyncAndroidInstanceImages,omitnil,omitempty" name:"SyncAndroidInstanceImages"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SyncAndroidInstanceImageResponse struct {
	*tchttp.BaseResponse
	Response *SyncAndroidInstanceImageResponseParams `json:"Response"`
}

func (r *SyncAndroidInstanceImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncAndroidInstanceImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncExecuteCommandOnAndroidInstancesRequestParams struct {
	// 安卓实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// shell 命令，必须是1秒内能自动结束的命令
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`
}

type SyncExecuteCommandOnAndroidInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 安卓实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// shell 命令，必须是1秒内能自动结束的命令
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`
}

func (r *SyncExecuteCommandOnAndroidInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncExecuteCommandOnAndroidInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	delete(f, "Command")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncExecuteCommandOnAndroidInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncExecuteCommandOnAndroidInstancesResponseParams struct {
	// 命令执行结果列表
	CommandResults []*SyncExecuteCommandResult `json:"CommandResults,omitnil,omitempty" name:"CommandResults"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SyncExecuteCommandOnAndroidInstancesResponse struct {
	*tchttp.BaseResponse
	Response *SyncExecuteCommandOnAndroidInstancesResponseParams `json:"Response"`
}

func (r *SyncExecuteCommandOnAndroidInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncExecuteCommandOnAndroidInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncExecuteCommandResult struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 命令执行输出内容
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`

	// 命令执行结果
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type TrylockWorkerRequestParams struct {
	// 唯一用户身份标识，由业务方自定义，平台不予理解。（可根据业务需要决定使用用户的唯一身份标识或是使用时间戳随机生成；在用户重连时应保持UserId不变）
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 游戏ID
	GameId *string `json:"GameId,omitnil,omitempty" name:"GameId"`

	// 游戏区域，ap-guangzhou、ap-shanghai、ap-beijing等，如果不为空，优先按照该区域进行调度分配机器
	GameRegion *string `json:"GameRegion,omitnil,omitempty" name:"GameRegion"`

	// 【废弃】资源池编号
	SetNo *uint64 `json:"SetNo,omitnil,omitempty" name:"SetNo"`

	// 【推荐填写】用户IP，用户客户端的公网IP，用于就近调度，不填将严重影响用户体验
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// 分组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type TrylockWorkerRequest struct {
	*tchttp.BaseRequest
	
	// 唯一用户身份标识，由业务方自定义，平台不予理解。（可根据业务需要决定使用用户的唯一身份标识或是使用时间戳随机生成；在用户重连时应保持UserId不变）
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 游戏ID
	GameId *string `json:"GameId,omitnil,omitempty" name:"GameId"`

	// 游戏区域，ap-guangzhou、ap-shanghai、ap-beijing等，如果不为空，优先按照该区域进行调度分配机器
	GameRegion *string `json:"GameRegion,omitnil,omitempty" name:"GameRegion"`

	// 【废弃】资源池编号
	SetNo *uint64 `json:"SetNo,omitnil,omitempty" name:"SetNo"`

	// 【推荐填写】用户IP，用户客户端的公网IP，用于就近调度，不填将严重影响用户体验
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// 分组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *TrylockWorkerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TrylockWorkerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "GameId")
	delete(f, "GameRegion")
	delete(f, "SetNo")
	delete(f, "UserIp")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TrylockWorkerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TrylockWorkerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TrylockWorkerResponse struct {
	*tchttp.BaseResponse
	Response *TrylockWorkerResponseParams `json:"Response"`
}

func (r *TrylockWorkerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TrylockWorkerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UninstallAndroidInstancesAppRequestParams struct {
	// 实例ID
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 应用ID
	AndroidAppId *string `json:"AndroidAppId,omitnil,omitempty" name:"AndroidAppId"`

	// 包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

type UninstallAndroidInstancesAppRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 应用ID
	AndroidAppId *string `json:"AndroidAppId,omitnil,omitempty" name:"AndroidAppId"`

	// 包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

func (r *UninstallAndroidInstancesAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UninstallAndroidInstancesAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	delete(f, "AndroidAppId")
	delete(f, "PackageName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UninstallAndroidInstancesAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UninstallAndroidInstancesAppResponseParams struct {
	// 任务集合
	TaskSet []*AndroidInstanceTask `json:"TaskSet,omitnil,omitempty" name:"TaskSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UninstallAndroidInstancesAppResponse struct {
	*tchttp.BaseResponse
	Response *UninstallAndroidInstancesAppResponseParams `json:"Response"`
}

func (r *UninstallAndroidInstancesAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UninstallAndroidInstancesAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadFileToAndroidInstancesRequestParams struct {
	// 安卓实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 文件下载 URL
	FileURL *string `json:"FileURL,omitnil,omitempty" name:"FileURL"`

	// 上传目标目录，只能上传到 /sdcard/ 目录或其子目录下
	DestinationDirectory *string `json:"DestinationDirectory,omitnil,omitempty" name:"DestinationDirectory"`

	// 目标文件名
	DestinationFileName *string `json:"DestinationFileName,omitnil,omitempty" name:"DestinationFileName"`
}

type UploadFileToAndroidInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 安卓实例 ID 列表
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitnil,omitempty" name:"AndroidInstanceIds"`

	// 文件下载 URL
	FileURL *string `json:"FileURL,omitnil,omitempty" name:"FileURL"`

	// 上传目标目录，只能上传到 /sdcard/ 目录或其子目录下
	DestinationDirectory *string `json:"DestinationDirectory,omitnil,omitempty" name:"DestinationDirectory"`

	// 目标文件名
	DestinationFileName *string `json:"DestinationFileName,omitnil,omitempty" name:"DestinationFileName"`
}

func (r *UploadFileToAndroidInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadFileToAndroidInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AndroidInstanceIds")
	delete(f, "FileURL")
	delete(f, "DestinationDirectory")
	delete(f, "DestinationFileName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadFileToAndroidInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadFileToAndroidInstancesResponseParams struct {
	// 实例任务集合
	TaskSet []*AndroidInstanceTask `json:"TaskSet,omitnil,omitempty" name:"TaskSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UploadFileToAndroidInstancesResponse struct {
	*tchttp.BaseResponse
	Response *UploadFileToAndroidInstancesResponseParams `json:"Response"`
}

func (r *UploadFileToAndroidInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadFileToAndroidInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadFilesToAndroidInstancesRequestParams struct {
	// 上传文件信息列表
	Files []*AndroidInstanceUploadFile `json:"Files,omitnil,omitempty" name:"Files"`
}

type UploadFilesToAndroidInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 上传文件信息列表
	Files []*AndroidInstanceUploadFile `json:"Files,omitnil,omitempty" name:"Files"`
}

func (r *UploadFilesToAndroidInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadFilesToAndroidInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Files")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadFilesToAndroidInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadFilesToAndroidInstancesResponseParams struct {
	// 实例任务集合
	TaskSet []*AndroidInstanceTask `json:"TaskSet,omitnil,omitempty" name:"TaskSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UploadFilesToAndroidInstancesResponse struct {
	*tchttp.BaseResponse
	Response *UploadFilesToAndroidInstancesResponseParams `json:"Response"`
}

func (r *UploadFilesToAndroidInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadFilesToAndroidInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}