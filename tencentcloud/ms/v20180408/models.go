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

package v20180408

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AndroidAppInfo struct {
	// app文件的md5算法值，需要正确传递，在线加固必输。
	// 例如linux环境下执行算法命令md5sum ：
	// #md5sum test.apk 
	// d40cc11e4bddd643ecdf29cde729a12b
	AppMd5 *string `json:"AppMd5,omitempty" name:"AppMd5"`

	// app的大小，非必输。
	AppSize *int64 `json:"AppSize,omitempty" name:"AppSize"`

	// app下载链接，在线加固必输。
	AppUrl *string `json:"AppUrl,omitempty" name:"AppUrl"`

	// app名称，非必输
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// app的包名，本次操作的包名。
	// 当Android是按年收费、免费试用加固时，在线加固和输出工具要求该字段必输，且与AndroidPlan.AppPkgName值相等。
	AppPkgName *string `json:"AppPkgName,omitempty" name:"AppPkgName"`

	// app的文件名，非必输。
	AppFileName *string `json:"AppFileName,omitempty" name:"AppFileName"`

	// app版本号，非必输。
	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`

	// Android app的文件类型，本次加固操作的应用类型 。
	// Android在线加固和输出工具加固必输，其值需等于“apk”或“aab”，且与AndroidAppInfo.AppType值相等。
	AppType *string `json:"AppType,omitempty" name:"AppType"`
}

type AndroidPlan struct {
	// 非必输字段，PlanId 是指本次加固使用的配置策略Id，可通过载入上次配置接口获取。其值非0时，代表引用对应的策略。
	PlanId *int64 `json:"PlanId,omitempty" name:"PlanId"`

	// 本次操作的包名。
	// 当收费模式是安卓按年收费和安卓免费试用的在线加固和输出工具加固时，要求该字段必输，且与AndroidAppInfo.AppPkgName值相等。
	AppPkgName *string `json:"AppPkgName,omitempty" name:"AppPkgName"`

	// 安卓app的文件类型，本次加固操作的应用类型 。 
	// 安卓在线加固和输出工具加固必输，其值需等于“apk”或“aab”，且与AndroidAppInfo.AppType值相等。
	AppType *string `json:"AppType,omitempty" name:"AppType"`

	// 安卓加固必输字段。
	// 加固策略，json格式字符串。
	// 字段说明（0-关闭，1-开启）：
	//         "enable"=1 #DEX整体加固;
	//         "antiprotect"=1 #反调试;
	//         "antirepack"=1 #防重打包、防篡改;
	//         "dexsig"=1       #签名校验;
	//         "antimonitor"=1 #防模拟器运行保护;
	//         "ptrace"=1 #防动态注入、动态调试;
	//         "so"."enable" = 1 #文件加密;
	//         "vmp"."enable" = 1 #VMP虚拟化保护;
	//         "respro"."assets"."enable" = 1 #assets资源文件加密
	//        "respro"."res"."enable" = 1 #res资源文件加密
	// 
	// so文件加密：
	// 支持5种架构:
	// apk 格式: /lib/armeabi/libxxx.so,/lib/arm64-v8a/libxxx.so,/lib/armeabi-v7a/libxxx.so,/lib/x86/libxxx.so,/lib/x86_64/libxxx.so
	// aab格式: /base/lib/armeabi/libxxx.so,/base/lib/arm64-v8a/libxxx.so,/base/lib/armeabi-v7a/libxxx.so,/base/lib/x86/libxxx.so,/base/lib/x86_64/libxxx.so
	// 请列举 SO 库在 apk 文件解压后的完整有效路径，如:/lib/armeabi/libxxx.so；
	// 需要加固的 SO 库需确认为自研的 SO 库，不要加固第三方 SO 库，否则会增加 crash 风险
	// 
	// res资源文件加密注意事项：
	// 请指定需要加密的文件全路径，如：res/layout/1.xml;
	// res资源文件加密不能加密APP图标
	// res目录文件，不能加密以下后缀规则的文件".wav", ".mp2", ".mp3", ".ogg", ".aac", ".mpg",".mpeg", ".mid", ".midi", ".smf", ".jet", ".rtttl", ".imy", ".xmf", ".mp4", ".m4a", ".m4v", ".3gp",".3gpp", ".3g2", ".3gpp2", ".amr", ".awb", ".wma", ".wmv"
	// 
	// assets资源文件加密注意事项:
	// 请指定需要加密的文件全路径，如：assets/main.js；可以完整路径，也可以相对路径。
	// 如果有通配符需要完整路径 ":all"或者"*"代表所有文件
	// assets资源文件加密不能加密APP图标
	// assets目录文件，不能加密以下后缀规则的文件".wav", ".mp2", ".mp3", ".ogg", ".aac", ".mpg",".mpeg", ".mid", ".midi", ".smf", ".jet", ".rtttl", ".imy", ".xmf", ".mp4", ".m4a", ".m4v", ".3gp",".3gpp", ".3g2", ".3gpp2", ".amr", ".awb", ".wma", ".wmv"
	// 
	// 
	// apk[dex+so+vmp+res+assets]加固参数示例：
	// ‘{
	//     "dex": {
	//         "enable": 1,
	//         "antiprotect": 1,
	//         "antirepack": 1,
	//         "dexsig": 1,
	//         "antimonitor": 1,
	//         "ptrace": 1
	//     },
	//     "so": {
	//         "enable": 1,
	//         "ver": "1.3.3",
	//         "file": [
	//             "/lib/armeabi/libtest.so"
	//         ]
	//     },
	//     "vmp": {
	//         "enable": 1,
	//         "ndkpath": "/xxx/android-ndk-r10e",
	//         "profile": "/xxx/vmpprofile.txt",
	//         "mapping": "/xxx/mapping.txt"
	//     },
	//     "respro": {
	//         "assets": {
	//             "enable": 1,
	//             "file": [
	//                 "assets/1.js",
	//                 "assets/2.jpg"
	//             ]
	//         },
	//         "res": {
	//             "enable": 1,
	//             "file": [
	//                 "res/layout/1.xml",
	//                 "res/layout/2.xml"
	//             ]
	//         }
	//     }
	// }’
	// 
	// aab加固方案一 
	// [dex+res+assets]加固json字符串：
	// ‘{
	//     "dex": {
	//         "enable": 1,
	//         "antiprotect": 1,
	//         "antimonitor": 1
	//     },
	//     "respro": {
	//         "assets": {
	//             "enable": 1,
	//             "file": [
	//                 "assets/1.js",
	//                 "assets/2.jpg"
	//             ]
	//         },
	//         "res": {
	//             "enable": 1,
	//             "file": [
	//                 "res/layout/1.xml",
	//                 "res/layout/2.xml"
	//             ]
	//         }
	//     }
	// }’
	// 
	// aab加固方案二
	// 单独vmp加固：
	// ‘{
	//     "vmp": {
	//         "enable": 1,
	//         "ndkpath": "/xxx/android-ndk-r10e",
	//         "profile": "/xxx/vmpprofile.txt",
	//         "mapping": "/xxx/mapping.txt",
	//         "antiprotect": 1,
	//         "antimonitor": 1
	//     }
	// }’
	EncryptParam *string `json:"EncryptParam,omitempty" name:"EncryptParam"`
}

type AndroidResult struct {
	// 结果Id,用于查询加固结果
	ResultId *string `json:"ResultId,omitempty" name:"ResultId"`

	// 与当前任务关联的订单id
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 与当前任务关联的资源Id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 本次任务发起者
	OpUin *int64 `json:"OpUin,omitempty" name:"OpUin"`

	// 应用类型：安卓-apk; 安卓-aab;
	AppType *string `json:"AppType,omitempty" name:"AppType"`

	// 应用包名
	AppPkgName *string `json:"AppPkgName,omitempty" name:"AppPkgName"`

	// 后台资源绑定的包名
	BindAppPkgName *string `json:"BindAppPkgName,omitempty" name:"BindAppPkgName"`

	// 加固结果
	EncryptState *int64 `json:"EncryptState,omitempty" name:"EncryptState"`

	// 加固结果描述
	EncryptStateDesc *string `json:"EncryptStateDesc,omitempty" name:"EncryptStateDesc"`

	// 加固失败错误码
	EncryptErrCode *int64 `json:"EncryptErrCode,omitempty" name:"EncryptErrCode"`

	// 加固失败原因
	EncryptErrDesc *string `json:"EncryptErrDesc,omitempty" name:"EncryptErrDesc"`

	// 加固失败解决方案
	EncryptErrRef *string `json:"EncryptErrRef,omitempty" name:"EncryptErrRef"`

	// 任务创建时间
	CreatTime *string `json:"CreatTime,omitempty" name:"CreatTime"`

	// 任务开始处理时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 任务处理结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 加固耗时（秒单位）
	CostTime *int64 `json:"CostTime,omitempty" name:"CostTime"`

	// 在线加固-安卓应用原包下载链接
	AppUrl *string `json:"AppUrl,omitempty" name:"AppUrl"`

	// 在线加固-安卓应用文件MD5算法值
	AppMd5 *string `json:"AppMd5,omitempty" name:"AppMd5"`

	// 在线加固-安卓应用应用名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 在线加固-安卓应用版本；
	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`

	// 在线加固-安卓应用大小
	AppSize *int64 `json:"AppSize,omitempty" name:"AppSize"`

	// 在线加固-安卓加固-腾讯云应用加固工具版本
	OnlineToolVersion *string `json:"OnlineToolVersion,omitempty" name:"OnlineToolVersion"`

	// 在线加固-安卓加固，加固成功后文件md5算法值
	EncryptAppMd5 *string `json:"EncryptAppMd5,omitempty" name:"EncryptAppMd5"`

	// 在线加固-安卓加固，加固成功后应用大小
	EncryptAppSize *int64 `json:"EncryptAppSize,omitempty" name:"EncryptAppSize"`

	// 在线加固-安卓加固，加固包下载链接。
	EncryptPkgUrl *string `json:"EncryptPkgUrl,omitempty" name:"EncryptPkgUrl"`

	// 输出工具-安卓加固-腾讯云输出工具版本
	OutputToolVersion *string `json:"OutputToolVersion,omitempty" name:"OutputToolVersion"`

	// 输出工具-安卓加固-工具大小
	OutputToolSize *int64 `json:"OutputToolSize,omitempty" name:"OutputToolSize"`

	// 输出工具-安卓加固-工具输出时间
	ToolOutputTime *string `json:"ToolOutputTime,omitempty" name:"ToolOutputTime"`

	// 输出工具-安卓加固-工具到期时间
	ToolExpireTime *string `json:"ToolExpireTime,omitempty" name:"ToolExpireTime"`

	// 输出工具-安卓加固-输出工具下载链接
	OutputToolUrl *string `json:"OutputToolUrl,omitempty" name:"OutputToolUrl"`

	// 本次安卓加固策略信息
	AndroidPlan *AndroidPlan `json:"AndroidPlan,omitempty" name:"AndroidPlan"`
}

type AppDetailInfo struct {
	// app的名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// app的包名
	AppPkgName *string `json:"AppPkgName,omitempty" name:"AppPkgName"`

	// app的版本号
	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`

	// app的大小
	AppSize *uint64 `json:"AppSize,omitempty" name:"AppSize"`

	// app的md5
	AppMd5 *string `json:"AppMd5,omitempty" name:"AppMd5"`

	// app的图标url
	AppIconUrl *string `json:"AppIconUrl,omitempty" name:"AppIconUrl"`

	// app的文件名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

type AppInfo struct {
	// app的url，必须保证不用权限校验就可以下载
	AppUrl *string `json:"AppUrl,omitempty" name:"AppUrl"`

	// app的md5，需要正确传递
	AppMd5 *string `json:"AppMd5,omitempty" name:"AppMd5"`

	// app的大小
	AppSize *uint64 `json:"AppSize,omitempty" name:"AppSize"`

	// app的文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// app的包名，需要正确的传递此字段
	AppPkgName *string `json:"AppPkgName,omitempty" name:"AppPkgName"`

	// app的版本号
	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`

	// app的图标url
	AppIconUrl *string `json:"AppIconUrl,omitempty" name:"AppIconUrl"`

	// app的名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`
}

type AppSetInfo struct {
	// 任务唯一标识
	ItemId *string `json:"ItemId,omitempty" name:"ItemId"`

	// app的名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// app的包名
	AppPkgName *string `json:"AppPkgName,omitempty" name:"AppPkgName"`

	// app的版本号
	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`

	// app的md5
	AppMd5 *string `json:"AppMd5,omitempty" name:"AppMd5"`

	// app的大小
	AppSize *uint64 `json:"AppSize,omitempty" name:"AppSize"`

	// 加固服务版本
	ServiceEdition *string `json:"ServiceEdition,omitempty" name:"ServiceEdition"`

	// 加固结果返回码
	ShieldCode *uint64 `json:"ShieldCode,omitempty" name:"ShieldCode"`

	// 加固后的APP下载地址
	AppUrl *string `json:"AppUrl,omitempty" name:"AppUrl"`

	// 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
	TaskStatus *uint64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 请求的客户端ip
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// 提交加固时间
	TaskTime *uint64 `json:"TaskTime,omitempty" name:"TaskTime"`

	// app的图标url
	AppIconUrl *string `json:"AppIconUrl,omitempty" name:"AppIconUrl"`

	// 加固后app的md5
	ShieldMd5 *string `json:"ShieldMd5,omitempty" name:"ShieldMd5"`

	// 加固后app的大小
	ShieldSize *uint64 `json:"ShieldSize,omitempty" name:"ShieldSize"`
}

type AppletInfo struct {
	// 客户JS包
	AppletJsUrl *string `json:"AppletJsUrl,omitempty" name:"AppletJsUrl"`

	// 小程序加固等级配置
	// 1 - 开启代码混淆、代码压缩、代码反调试保护。 2 - 开启字符串编码和代码变换，代码膨胀，随机插入冗余代码，开启代码控制流平坦化，保证业务逻辑正常前提下，扁平化代码逻辑分支，破坏代码简单的线性结构。 3 - 开启代码加密，对字符串、函数、变量、属性、类、数组等结构进行加密保护，更多得代码控制流平坦化，扁平化逻辑分支。
	AppletLevel *int64 `json:"AppletLevel,omitempty" name:"AppletLevel"`

	// 本次加固输出产物名称，如”test.zip“,非空必须是 ”.zip“结尾
	Name *string `json:"Name,omitempty" name:"Name"`
}

type AppletPlan struct {
	// 策略Id
	PlanId *int64 `json:"PlanId,omitempty" name:"PlanId"`

	// 1 - 开启代码混淆、代码压缩、代码反调试保护。
	// 2 - 开启字符串编码和代码变换，代码膨胀，随机插入冗余代码，开启代码控制流平坦化，保证业务逻辑正常前提下，扁平化代码逻辑分支，破坏代码简单的线性结构。
	// 3 - 开启代码加密，对字符串、函数、变量、属性、类、数组等结构进行加密保护，更多得代码控制流平坦化，扁平化逻辑分支。
	AppletLevel *int64 `json:"AppletLevel,omitempty" name:"AppletLevel"`
}

type AppletResult struct {
	// 加固任务结果id
	ResultId *string `json:"ResultId,omitempty" name:"ResultId"`

	// 资源id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 订单id
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 操作账号
	OpUin *int64 `json:"OpUin,omitempty" name:"OpUin"`

	// 加固结果
	EncryptState *int64 `json:"EncryptState,omitempty" name:"EncryptState"`

	// 加固结果描述
	EncryptStateDesc *string `json:"EncryptStateDesc,omitempty" name:"EncryptStateDesc"`

	// 失败错误码
	EncryptErrCode *int64 `json:"EncryptErrCode,omitempty" name:"EncryptErrCode"`

	// 失败原因
	EncryptErrDesc *string `json:"EncryptErrDesc,omitempty" name:"EncryptErrDesc"`

	// 解决方案
	EncryptErrRef *string `json:"EncryptErrRef,omitempty" name:"EncryptErrRef"`

	// 任务创建时间
	CreatTime *string `json:"CreatTime,omitempty" name:"CreatTime"`

	// 任务开始处理时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 任务处理结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 加固耗时（秒单位）
	CostTime *int64 `json:"CostTime,omitempty" name:"CostTime"`

	// 在线加固成功下载包
	EncryptPkgUrl *string `json:"EncryptPkgUrl,omitempty" name:"EncryptPkgUrl"`

	// 本次加固配置
	AppletInfo *AppletInfo `json:"AppletInfo,omitempty" name:"AppletInfo"`
}

type BindInfo struct {
	// app的icon的url
	AppIconUrl *string `json:"AppIconUrl,omitempty" name:"AppIconUrl"`

	// app的名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// app的包名
	AppPkgName *string `json:"AppPkgName,omitempty" name:"AppPkgName"`
}

// Predefined struct for user
type CancelEncryptTaskRequestParams struct {
	// 加固任务结果Id 
	ResultId *string `json:"ResultId,omitempty" name:"ResultId"`
}

type CancelEncryptTaskRequest struct {
	*tchttp.BaseRequest
	
	// 加固任务结果Id 
	ResultId *string `json:"ResultId,omitempty" name:"ResultId"`
}

func (r *CancelEncryptTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelEncryptTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResultId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelEncryptTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelEncryptTaskResponseParams struct {
	// 1: 取消任务成功 ； -1 ：取消任务失败，原因为任务进程已结束，不能取消。
	State *int64 `json:"State,omitempty" name:"State"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CancelEncryptTaskResponse struct {
	*tchttp.BaseResponse
	Response *CancelEncryptTaskResponseParams `json:"Response"`
}

func (r *CancelEncryptTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelEncryptTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBindInstanceRequestParams struct {
	// 资源id，全局唯一
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// app的icon的url
	AppIconUrl *string `json:"AppIconUrl,omitempty" name:"AppIconUrl"`

	// app的名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// app的包名
	AppPkgName *string `json:"AppPkgName,omitempty" name:"AppPkgName"`
}

type CreateBindInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 资源id，全局唯一
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// app的icon的url
	AppIconUrl *string `json:"AppIconUrl,omitempty" name:"AppIconUrl"`

	// app的名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// app的包名
	AppPkgName *string `json:"AppPkgName,omitempty" name:"AppPkgName"`
}

func (r *CreateBindInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBindInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "AppIconUrl")
	delete(f, "AppName")
	delete(f, "AppPkgName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBindInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBindInstanceResponseParams struct {
	// 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBindInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateBindInstanceResponseParams `json:"Response"`
}

func (r *CreateBindInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBindInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCosSecKeyInstanceRequestParams struct {
	// 地域信息，例如广州：ap-guangzhou，上海：ap-shanghai，默认为广州。
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// 密钥有效时间，默认为1小时。
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`
}

type CreateCosSecKeyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 地域信息，例如广州：ap-guangzhou，上海：ap-shanghai，默认为广州。
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// 密钥有效时间，默认为1小时。
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`
}

func (r *CreateCosSecKeyInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosSecKeyInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CosRegion")
	delete(f, "Duration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCosSecKeyInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCosSecKeyInstanceResponseParams struct {
	// COS密钥对应的AppId
	CosAppid *uint64 `json:"CosAppid,omitempty" name:"CosAppid"`

	// COS密钥对应的存储桶名
	CosBucket *string `json:"CosBucket,omitempty" name:"CosBucket"`

	// 存储桶对应的地域
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// 密钥过期时间
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 密钥ID信息
	CosId *string `json:"CosId,omitempty" name:"CosId"`

	// 密钥KEY信息
	CosKey *string `json:"CosKey,omitempty" name:"CosKey"`

	// 密钥TOCKEN信息
	//
	// Deprecated: CosTocken is deprecated.
	CosTocken *string `json:"CosTocken,omitempty" name:"CosTocken"`

	// 密钥可访问的文件前缀人。例如：CosPrefix=test/123/666，则该密钥只能操作test/123/666为前缀的文件，例如test/123/666/1.txt
	CosPrefix *string `json:"CosPrefix,omitempty" name:"CosPrefix"`

	// 密钥TOCKEN信息
	CosToken *string `json:"CosToken,omitempty" name:"CosToken"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCosSecKeyInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateCosSecKeyInstanceResponseParams `json:"Response"`
}

func (r *CreateCosSecKeyInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosSecKeyInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEncryptInstanceRequestParams struct {
	// 平台类型  1.android加固   2.ios源码混淆  3.sdk加固  4.applet小程序加固
	PlatformType *int64 `json:"PlatformType,omitempty" name:"PlatformType"`

	// 订单采购类型 1-免费试用 2-按年收费 3-按次收费
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 1-在线加固、  2-输出工具加固
	EncryptOpType *int64 `json:"EncryptOpType,omitempty" name:"EncryptOpType"`

	// 本次加固使用的资源id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 渠道合作android加固App信息 
	AndroidAppInfo *AndroidAppInfo `json:"AndroidAppInfo,omitempty" name:"AndroidAppInfo"`

	// 渠道合作android加固策略信息
	AndroidPlan *AndroidPlan `json:"AndroidPlan,omitempty" name:"AndroidPlan"`

	// 小程序加固信息
	AppletInfo *AppletInfo `json:"AppletInfo,omitempty" name:"AppletInfo"`
}

type CreateEncryptInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 平台类型  1.android加固   2.ios源码混淆  3.sdk加固  4.applet小程序加固
	PlatformType *int64 `json:"PlatformType,omitempty" name:"PlatformType"`

	// 订单采购类型 1-免费试用 2-按年收费 3-按次收费
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 1-在线加固、  2-输出工具加固
	EncryptOpType *int64 `json:"EncryptOpType,omitempty" name:"EncryptOpType"`

	// 本次加固使用的资源id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 渠道合作android加固App信息 
	AndroidAppInfo *AndroidAppInfo `json:"AndroidAppInfo,omitempty" name:"AndroidAppInfo"`

	// 渠道合作android加固策略信息
	AndroidPlan *AndroidPlan `json:"AndroidPlan,omitempty" name:"AndroidPlan"`

	// 小程序加固信息
	AppletInfo *AppletInfo `json:"AppletInfo,omitempty" name:"AppletInfo"`
}

func (r *CreateEncryptInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEncryptInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlatformType")
	delete(f, "OrderType")
	delete(f, "EncryptOpType")
	delete(f, "ResourceId")
	delete(f, "AndroidAppInfo")
	delete(f, "AndroidPlan")
	delete(f, "AppletInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEncryptInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEncryptInstanceResponseParams struct {
	// 加固任务Id
	ResultId *string `json:"ResultId,omitempty" name:"ResultId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateEncryptInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateEncryptInstanceResponseParams `json:"Response"`
}

func (r *CreateEncryptInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEncryptInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrderInstanceRequestParams struct {
	// 平台类型枚举值：1-android加固  ；2-ios源码混淆 ； 3-sdk加固 ； 4-applet小程序加固
	PlatformType *int64 `json:"PlatformType,omitempty" name:"PlatformType"`

	// 订单采购类型 1-免费试用 ；2-按年收费 ；3-按次收费
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 代表应用包名列表，值为单个包名（例如：“a.b.xxx”）或多个包名用逗号隔开(例如：“a.b.xxx,b.c.xxx”)。
	// 当android按年收费加固或android免费试用加固时，该字段要求非空，即PlatformType=1 并且 OrderType=2时，AppPkgNameList必传值。
	AppPkgNameList *string `json:"AppPkgNameList,omitempty" name:"AppPkgNameList"`
}

type CreateOrderInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 平台类型枚举值：1-android加固  ；2-ios源码混淆 ； 3-sdk加固 ； 4-applet小程序加固
	PlatformType *int64 `json:"PlatformType,omitempty" name:"PlatformType"`

	// 订单采购类型 1-免费试用 ；2-按年收费 ；3-按次收费
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 代表应用包名列表，值为单个包名（例如：“a.b.xxx”）或多个包名用逗号隔开(例如：“a.b.xxx,b.c.xxx”)。
	// 当android按年收费加固或android免费试用加固时，该字段要求非空，即PlatformType=1 并且 OrderType=2时，AppPkgNameList必传值。
	AppPkgNameList *string `json:"AppPkgNameList,omitempty" name:"AppPkgNameList"`
}

func (r *CreateOrderInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrderInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlatformType")
	delete(f, "OrderType")
	delete(f, "AppPkgNameList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrderInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrderInstanceResponseParams struct {
	// 订单Id
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 与订单关联的资源id
	ResourceId []*string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateOrderInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrderInstanceResponseParams `json:"Response"`
}

func (r *CreateOrderInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrderInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResourceInstancesRequestParams struct {
	// 资源类型id。13624：加固专业版。
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`

	// 时间单位，取值为d，m，y，分别表示天，月，年。
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 时间数量。
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 资源数量。
	ResourceNum *uint64 `json:"ResourceNum,omitempty" name:"ResourceNum"`
}

type CreateResourceInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 资源类型id。13624：加固专业版。
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`

	// 时间单位，取值为d，m，y，分别表示天，月，年。
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 时间数量。
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 资源数量。
	ResourceNum *uint64 `json:"ResourceNum,omitempty" name:"ResourceNum"`
}

func (r *CreateResourceInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResourceInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Pid")
	delete(f, "TimeUnit")
	delete(f, "TimeSpan")
	delete(f, "ResourceNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateResourceInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResourceInstancesResponseParams struct {
	// 新创建的资源列表。
	ResourceSet []*string `json:"ResourceSet,omitempty" name:"ResourceSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateResourceInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateResourceInstancesResponseParams `json:"Response"`
}

func (r *CreateResourceInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResourceInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateShieldInstanceRequestParams struct {
	// 待加固的应用信息
	AppInfo *AppInfo `json:"AppInfo,omitempty" name:"AppInfo"`

	// 加固服务信息
	ServiceInfo *ServiceInfo `json:"ServiceInfo,omitempty" name:"ServiceInfo"`
}

type CreateShieldInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 待加固的应用信息
	AppInfo *AppInfo `json:"AppInfo,omitempty" name:"AppInfo"`

	// 加固服务信息
	ServiceInfo *ServiceInfo `json:"ServiceInfo,omitempty" name:"ServiceInfo"`
}

func (r *CreateShieldInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateShieldInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppInfo")
	delete(f, "ServiceInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateShieldInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateShieldInstanceResponseParams struct {
	// 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`

	// 任务唯一标识
	ItemId *string `json:"ItemId,omitempty" name:"ItemId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateShieldInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateShieldInstanceResponseParams `json:"Response"`
}

func (r *CreateShieldInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateShieldInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateShieldPlanInstanceRequestParams struct {
	// 资源id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 策略名称
	PlanName *string `json:"PlanName,omitempty" name:"PlanName"`

	// 策略具体信息
	PlanInfo *PlanInfo `json:"PlanInfo,omitempty" name:"PlanInfo"`
}

type CreateShieldPlanInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 资源id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 策略名称
	PlanName *string `json:"PlanName,omitempty" name:"PlanName"`

	// 策略具体信息
	PlanInfo *PlanInfo `json:"PlanInfo,omitempty" name:"PlanInfo"`
}

func (r *CreateShieldPlanInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateShieldPlanInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "PlanName")
	delete(f, "PlanInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateShieldPlanInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateShieldPlanInstanceResponseParams struct {
	// 策略id
	PlanId *uint64 `json:"PlanId,omitempty" name:"PlanId"`

	// 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateShieldPlanInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateShieldPlanInstanceResponseParams `json:"Response"`
}

func (r *CreateShieldPlanInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateShieldPlanInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteShieldInstancesRequestParams struct {
	// 任务唯一标识ItemId的列表
	ItemIds []*string `json:"ItemIds,omitempty" name:"ItemIds"`
}

type DeleteShieldInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 任务唯一标识ItemId的列表
	ItemIds []*string `json:"ItemIds,omitempty" name:"ItemIds"`
}

func (r *DeleteShieldInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteShieldInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ItemIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteShieldInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteShieldInstancesResponseParams struct {
	// 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteShieldInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteShieldInstancesResponseParams `json:"Response"`
}

func (r *DeleteShieldInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteShieldInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApkDetectionResultRequestParams struct {
	// 软件包的下载链接
	ApkUrl *string `json:"ApkUrl,omitempty" name:"ApkUrl"`

	// 软件包的md5值，具有唯一性。腾讯APK云检测服务会根据md5值来判断该包是否为库中已收集的样本，已存在，则返回检测结果，反之，需要一定时间检测该样本。
	ApkMd5 *string `json:"ApkMd5,omitempty" name:"ApkMd5"`
}

type DescribeApkDetectionResultRequest struct {
	*tchttp.BaseRequest
	
	// 软件包的下载链接
	ApkUrl *string `json:"ApkUrl,omitempty" name:"ApkUrl"`

	// 软件包的md5值，具有唯一性。腾讯APK云检测服务会根据md5值来判断该包是否为库中已收集的样本，已存在，则返回检测结果，反之，需要一定时间检测该样本。
	ApkMd5 *string `json:"ApkMd5,omitempty" name:"ApkMd5"`
}

func (r *DescribeApkDetectionResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApkDetectionResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApkUrl")
	delete(f, "ApkMd5")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApkDetectionResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApkDetectionResultResponseParams struct {
	// 响应结果，ok表示正常，error表示错误
	Result *string `json:"Result,omitempty" name:"Result"`

	// Result为error错误时的原因说明
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// APK检测结果数组
	ResultList []*ResultListItem `json:"ResultList,omitempty" name:"ResultList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApkDetectionResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApkDetectionResultResponseParams `json:"Response"`
}

func (r *DescribeApkDetectionResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApkDetectionResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEncryptInstancesRequestParams struct {
	// 多记录查询时使用，页码
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 多记录每页展示数量
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 多记录查询时排序使用  仅支持CreateTime 任务创建时间排序
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 升序（asc）还是降序（desc），默认：desc。
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`

	// (条件过滤字段) 平台类型  1.android加固   2.ios源码混淆  3.sdk加固  4.applet小程序加固
	PlatformType *int64 `json:"PlatformType,omitempty" name:"PlatformType"`

	// (条件过滤字段) 订单采购类型 1-免费试用 2-按年收费 3-按次收费
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// (条件过滤字段) 1-在线加固 或 2-输出工具加固
	EncryptOpType *int64 `json:"EncryptOpType,omitempty" name:"EncryptOpType"`

	// (条件过滤字段) 单记录查询时使用，结果ID该字段非空时，结构会根据结果ID进行单记录查询，符合查询条件时，只返回一条记录。
	ResultId *string `json:"ResultId,omitempty" name:"ResultId"`

	// (条件过滤字段) 查询与订单Id关联的任务
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// (条件过滤字段) 查询与资源Id关联的任务
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// (条件过滤字段) 应用类型：android-apk; android-aab;
	AppType *string `json:"AppType,omitempty" name:"AppType"`

	// （条件过滤字段）应用的包名
	AppPkgName *string `json:"AppPkgName,omitempty" name:"AppPkgName"`

	// 加固结果，
	// 0：正在排队；
	// 1：加固成功；
	// 2：加固中；
	// 3：加固失败；
	// 5：重试；
	// 多记录查询时，根据查询结果进行检索使用。
	EncryptState []*int64 `json:"EncryptState,omitempty" name:"EncryptState"`
}

type DescribeEncryptInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 多记录查询时使用，页码
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 多记录每页展示数量
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 多记录查询时排序使用  仅支持CreateTime 任务创建时间排序
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 升序（asc）还是降序（desc），默认：desc。
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`

	// (条件过滤字段) 平台类型  1.android加固   2.ios源码混淆  3.sdk加固  4.applet小程序加固
	PlatformType *int64 `json:"PlatformType,omitempty" name:"PlatformType"`

	// (条件过滤字段) 订单采购类型 1-免费试用 2-按年收费 3-按次收费
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// (条件过滤字段) 1-在线加固 或 2-输出工具加固
	EncryptOpType *int64 `json:"EncryptOpType,omitempty" name:"EncryptOpType"`

	// (条件过滤字段) 单记录查询时使用，结果ID该字段非空时，结构会根据结果ID进行单记录查询，符合查询条件时，只返回一条记录。
	ResultId *string `json:"ResultId,omitempty" name:"ResultId"`

	// (条件过滤字段) 查询与订单Id关联的任务
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// (条件过滤字段) 查询与资源Id关联的任务
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// (条件过滤字段) 应用类型：android-apk; android-aab;
	AppType *string `json:"AppType,omitempty" name:"AppType"`

	// （条件过滤字段）应用的包名
	AppPkgName *string `json:"AppPkgName,omitempty" name:"AppPkgName"`

	// 加固结果，
	// 0：正在排队；
	// 1：加固成功；
	// 2：加固中；
	// 3：加固失败；
	// 5：重试；
	// 多记录查询时，根据查询结果进行检索使用。
	EncryptState []*int64 `json:"EncryptState,omitempty" name:"EncryptState"`
}

func (r *DescribeEncryptInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEncryptInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "OrderField")
	delete(f, "OrderDirection")
	delete(f, "PlatformType")
	delete(f, "OrderType")
	delete(f, "EncryptOpType")
	delete(f, "ResultId")
	delete(f, "OrderId")
	delete(f, "ResourceId")
	delete(f, "AppType")
	delete(f, "AppPkgName")
	delete(f, "EncryptState")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEncryptInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEncryptInstancesResponseParams struct {
	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 渠道合作加固信息数组
	EncryptResults []*EncryptResults `json:"EncryptResults,omitempty" name:"EncryptResults"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEncryptInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEncryptInstancesResponseParams `json:"Response"`
}

func (r *DescribeEncryptInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEncryptInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEncryptPlanRequestParams struct {
	// 平台类型  1.android加固   2.ios源码混淆  3.sdk加固  4.applet小程序加固
	PlatformType *int64 `json:"PlatformType,omitempty" name:"PlatformType"`

	// 订单采购类型 1-免费试用 2-按年收费 3-按次收费
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 1-在线加固；2-输出工具
	EncryptOpType *int64 `json:"EncryptOpType,omitempty" name:"EncryptOpType"`

	// 本次加固使用的资源id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// （条件过滤字段）加固查询时，根据包名查询
	AppPkgName *string `json:"AppPkgName,omitempty" name:"AppPkgName"`

	// （条件过滤字段）加固查询时，根据应用格式查询，枚举值：“apk”、“aab”
	AppType *string `json:"AppType,omitempty" name:"AppType"`
}

type DescribeEncryptPlanRequest struct {
	*tchttp.BaseRequest
	
	// 平台类型  1.android加固   2.ios源码混淆  3.sdk加固  4.applet小程序加固
	PlatformType *int64 `json:"PlatformType,omitempty" name:"PlatformType"`

	// 订单采购类型 1-免费试用 2-按年收费 3-按次收费
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 1-在线加固；2-输出工具
	EncryptOpType *int64 `json:"EncryptOpType,omitempty" name:"EncryptOpType"`

	// 本次加固使用的资源id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// （条件过滤字段）加固查询时，根据包名查询
	AppPkgName *string `json:"AppPkgName,omitempty" name:"AppPkgName"`

	// （条件过滤字段）加固查询时，根据应用格式查询，枚举值：“apk”、“aab”
	AppType *string `json:"AppType,omitempty" name:"AppType"`
}

func (r *DescribeEncryptPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEncryptPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlatformType")
	delete(f, "OrderType")
	delete(f, "EncryptOpType")
	delete(f, "ResourceId")
	delete(f, "AppPkgName")
	delete(f, "AppType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEncryptPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEncryptPlanResponseParams struct {
	// 平台类型整型值  
	PlatformType *int64 `json:"PlatformType,omitempty" name:"PlatformType"`

	// 平台类型描述 1.android加固   2.ios源码混淆  3.sdk加固  4.applet小程序加固
	PlatformTypeDesc *string `json:"PlatformTypeDesc,omitempty" name:"PlatformTypeDesc"`

	// 1- 在线加固 2-输出工具加固
	EncryptOpType *int64 `json:"EncryptOpType,omitempty" name:"EncryptOpType"`

	// 1- 在线加固 2-输出工具加固
	EncryptOpTypeDesc *string `json:"EncryptOpTypeDesc,omitempty" name:"EncryptOpTypeDesc"`

	// 订单收费类型枚举值
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 订单收费类型描述
	OrderTypeDesc *string `json:"OrderTypeDesc,omitempty" name:"OrderTypeDesc"`

	// 资源id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 上次加固策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	AndroidPlan *AndroidPlan `json:"AndroidPlan,omitempty" name:"AndroidPlan"`

	// 上次小程序加固策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppletPlan *AppletPlan `json:"AppletPlan,omitempty" name:"AppletPlan"`

	// 上次ios源码混淆加固配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	IOSPlan *IOSPlan `json:"IOSPlan,omitempty" name:"IOSPlan"`

	// 上次sdk加固配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	SDKPlan *SDKPlan `json:"SDKPlan,omitempty" name:"SDKPlan"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEncryptPlanResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEncryptPlanResponseParams `json:"Response"`
}

func (r *DescribeEncryptPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEncryptPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrderInstancesRequestParams struct {
	// 页码
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每页展示数量
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 按某个字段排序，目前仅支持CreateTime排序。
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 升序（asc）还是降序（desc），默认：desc。
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`

	// （条件过滤字段）平台类型  1.android加固   2.ios源码混淆  3.sdk加固  4.applet小程序加固
	PlatformType *int64 `json:"PlatformType,omitempty" name:"PlatformType"`

	// （条件过滤字段）订单采购类型 1-免费试用 2-按年收费 3-按次收费
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// （条件过滤字段）订单审批状态：
	ApprovalStatus *int64 `json:"ApprovalStatus,omitempty" name:"ApprovalStatus"`

	// （条件过滤字段）资源状态：
	ResourceStatus *int64 `json:"ResourceStatus,omitempty" name:"ResourceStatus"`

	// （条件过滤字段）订单ID
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// （条件过滤字段）资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// （条件过滤字段）包名，查询android加固订单时使用
	AppPkgName *string `json:"AppPkgName,omitempty" name:"AppPkgName"`
}

type DescribeOrderInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 页码
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每页展示数量
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 按某个字段排序，目前仅支持CreateTime排序。
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 升序（asc）还是降序（desc），默认：desc。
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`

	// （条件过滤字段）平台类型  1.android加固   2.ios源码混淆  3.sdk加固  4.applet小程序加固
	PlatformType *int64 `json:"PlatformType,omitempty" name:"PlatformType"`

	// （条件过滤字段）订单采购类型 1-免费试用 2-按年收费 3-按次收费
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// （条件过滤字段）订单审批状态：
	ApprovalStatus *int64 `json:"ApprovalStatus,omitempty" name:"ApprovalStatus"`

	// （条件过滤字段）资源状态：
	ResourceStatus *int64 `json:"ResourceStatus,omitempty" name:"ResourceStatus"`

	// （条件过滤字段）订单ID
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// （条件过滤字段）资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// （条件过滤字段）包名，查询android加固订单时使用
	AppPkgName *string `json:"AppPkgName,omitempty" name:"AppPkgName"`
}

func (r *DescribeOrderInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrderInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "OrderField")
	delete(f, "OrderDirection")
	delete(f, "PlatformType")
	delete(f, "OrderType")
	delete(f, "ApprovalStatus")
	delete(f, "ResourceStatus")
	delete(f, "OrderId")
	delete(f, "ResourceId")
	delete(f, "AppPkgName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrderInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrderInstancesResponseParams struct {
	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 订单信息
	Orders []*Orders `json:"Orders,omitempty" name:"Orders"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOrderInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrderInstancesResponseParams `json:"Response"`
}

func (r *DescribeOrderInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrderInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceInstancesRequestParams struct {
	// 支持CreateTime、ExpireTime、AppName、AppPkgName、BindValue、IsBind过滤
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数量限制，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 资源类别id数组，13624：加固专业版，12750：企业版。空数组表示返回全部资源。
	Pids []*uint64 `json:"Pids,omitempty" name:"Pids"`

	// 按某个字段排序，目前支持CreateTime、ExpireTime其中的一个排序。
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 升序（asc）还是降序（desc），默认：desc。
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

type DescribeResourceInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 支持CreateTime、ExpireTime、AppName、AppPkgName、BindValue、IsBind过滤
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数量限制，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 资源类别id数组，13624：加固专业版，12750：企业版。空数组表示返回全部资源。
	Pids []*uint64 `json:"Pids,omitempty" name:"Pids"`

	// 按某个字段排序，目前支持CreateTime、ExpireTime其中的一个排序。
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 升序（asc）还是降序（desc），默认：desc。
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeResourceInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Pids")
	delete(f, "OrderField")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceInstancesResponseParams struct {
	// 符合要求的资源数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 符合要求的资源数组
	ResourceSet []*ResourceInfo `json:"ResourceSet,omitempty" name:"ResourceSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResourceInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceInstancesResponseParams `json:"Response"`
}

func (r *DescribeResourceInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShieldInstancesRequestParams struct {
	// 支持通过app名称，app包名，加固的服务版本，提交的渠道进行筛选。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数量限制，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 可以提供ItemId数组来查询一个或者多个结果。注意不可以同时指定ItemIds和Filters。
	ItemIds []*string `json:"ItemIds,omitempty" name:"ItemIds"`

	// 按某个字段排序，目前仅支持TaskTime排序。
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 升序（asc）还是降序（desc），默认：desc。
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

type DescribeShieldInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 支持通过app名称，app包名，加固的服务版本，提交的渠道进行筛选。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数量限制，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 可以提供ItemId数组来查询一个或者多个结果。注意不可以同时指定ItemIds和Filters。
	ItemIds []*string `json:"ItemIds,omitempty" name:"ItemIds"`

	// 按某个字段排序，目前仅支持TaskTime排序。
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 升序（asc）还是降序（desc），默认：desc。
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeShieldInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShieldInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ItemIds")
	delete(f, "OrderField")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeShieldInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShieldInstancesResponseParams struct {
	// 符合要求的app数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 一个关于app详细信息的结构体，主要包括app的基本信息和加固信息。
	AppSet []*AppSetInfo `json:"AppSet,omitempty" name:"AppSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeShieldInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeShieldInstancesResponseParams `json:"Response"`
}

func (r *DescribeShieldInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShieldInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShieldPlanInstanceRequestParams struct {
	// 资源id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 服务类别id
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
}

type DescribeShieldPlanInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 资源id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 服务类别id
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
}

func (r *DescribeShieldPlanInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShieldPlanInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "Pid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeShieldPlanInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShieldPlanInstanceResponseParams struct {
	// 绑定资源信息
	BindInfo *BindInfo `json:"BindInfo,omitempty" name:"BindInfo"`

	// 加固策略信息
	ShieldPlanInfo *ShieldPlanInfo `json:"ShieldPlanInfo,omitempty" name:"ShieldPlanInfo"`

	// 加固资源信息
	ResourceServiceInfo *ResourceServiceInfo `json:"ResourceServiceInfo,omitempty" name:"ResourceServiceInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeShieldPlanInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeShieldPlanInstanceResponseParams `json:"Response"`
}

func (r *DescribeShieldPlanInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShieldPlanInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShieldResultRequestParams struct {
	// 任务唯一标识
	ItemId *string `json:"ItemId,omitempty" name:"ItemId"`
}

type DescribeShieldResultRequest struct {
	*tchttp.BaseRequest
	
	// 任务唯一标识
	ItemId *string `json:"ItemId,omitempty" name:"ItemId"`
}

func (r *DescribeShieldResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShieldResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ItemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeShieldResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShieldResultResponseParams struct {
	// 任务状态: 0-请返回,1-已完成,2-处理中,3-处理出错,4-处理超时
	TaskStatus *uint64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// app加固前的详细信息
	AppDetailInfo *AppDetailInfo `json:"AppDetailInfo,omitempty" name:"AppDetailInfo"`

	// app加固后的详细信息
	ShieldInfo *ShieldInfo `json:"ShieldInfo,omitempty" name:"ShieldInfo"`

	// 状态描述
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// 状态指引
	StatusRef *string `json:"StatusRef,omitempty" name:"StatusRef"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeShieldResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeShieldResultResponseParams `json:"Response"`
}

func (r *DescribeShieldResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShieldResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUrlDetectionResultRequestParams struct {
	// 查询的网址
	Url *string `json:"Url,omitempty" name:"Url"`
}

type DescribeUrlDetectionResultRequest struct {
	*tchttp.BaseRequest
	
	// 查询的网址
	Url *string `json:"Url,omitempty" name:"Url"`
}

func (r *DescribeUrlDetectionResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUrlDetectionResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Url")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUrlDetectionResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUrlDetectionResultResponseParams struct {
	// [查询结果]查询结果；枚举值：0 查询成功，否则查询失败
	ResultCode *int64 `json:"ResultCode,omitempty" name:"ResultCode"`

	// [固定信息]响应协议版本号
	RespVer *int64 `json:"RespVer,omitempty" name:"RespVer"`

	// [查询结果]url恶意状态
	// 枚举值：
	// 0-1：未知，访问暂无风险。
	// 2 ：	风险网址，具体的恶意类型定义参考恶意大类EvilClass字段。
	// 3-4：安全，访问无风险。
	// 
	// 注意：查询结果EvilClass字段在Urltype=2时，才有意义。
	UrlType *int64 `json:"UrlType,omitempty" name:"UrlType"`

	// [查询结果]url恶意类型大类:{
	//     "1": "社工欺诈（仿冒、账号钓鱼、中奖诈骗）",
	//     "2": "信息诈骗（虚假充值、虚假兼职、虚假金融投资、虚假信用卡代办、网络赌博诈骗）",
	//     "3": "虚假销售（男女保健美容减肥产品、电子产品、虚假广告、违法销售）",
	//     "4": "恶意文件（病毒文件，木马文件，恶意apk文件的下载链接以及站点，挂马网站）",
	//     "5": "博彩网站（博彩网站，在线赌博网站）",
	//     "6": "色情网站（涉嫌传播色情内容，提供色情服务的网站）"
	//   }
	EvilClass *int64 `json:"EvilClass,omitempty" name:"EvilClass"`

	// 该字段暂为空
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 该字段暂为空
	Level *int64 `json:"Level,omitempty" name:"Level"`

	// [查询详情]url检出时间；时间戳
	DetectTime *int64 `json:"DetectTime,omitempty" name:"DetectTime"`

	// 该字段暂为空
	Wording *string `json:"Wording,omitempty" name:"Wording"`

	// 该字段暂为空
	WordingTitle *string `json:"WordingTitle,omitempty" name:"WordingTitle"`

	// [查询结果]url恶意状态说明；为UrlType字段值对应的说明
	UrlTypeDesc *string `json:"UrlTypeDesc,omitempty" name:"UrlTypeDesc"`

	// [查询结果]url恶意大类说明；为EvilClass字段值对应的说明
	EvilClassDesc *string `json:"EvilClassDesc,omitempty" name:"EvilClassDesc"`

	// 该字段暂为空
	EvilTypeDesc *string `json:"EvilTypeDesc,omitempty" name:"EvilTypeDesc"`

	// 该字段暂为空
	LevelDesc *string `json:"LevelDesc,omitempty" name:"LevelDesc"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUrlDetectionResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUrlDetectionResultResponseParams `json:"Response"`
}

func (r *DescribeUrlDetectionResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUrlDetectionResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserBaseInfoInstanceRequestParams struct {

}

type DescribeUserBaseInfoInstanceRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeUserBaseInfoInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserBaseInfoInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserBaseInfoInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserBaseInfoInstanceResponseParams struct {
	// 用户uin信息
	UserUin *uint64 `json:"UserUin,omitempty" name:"UserUin"`

	// 用户APPID信息
	UserAppid *uint64 `json:"UserAppid,omitempty" name:"UserAppid"`

	// 系统时间戳
	TimeStamp *uint64 `json:"TimeStamp,omitempty" name:"TimeStamp"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUserBaseInfoInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserBaseInfoInstanceResponseParams `json:"Response"`
}

func (r *DescribeUserBaseInfoInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserBaseInfoInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EncryptResults struct {
	// 平台类型枚举值  1-android安卓加固   2-ios源码混淆  3-sdk加固  4-applet小程序加固
	PlatformType *int64 `json:"PlatformType,omitempty" name:"PlatformType"`

	// 平台类型描述  1-android安卓加固   2-ios源码混淆  3-sdk加固  4-applet小程序加固
	PlatformDesc *string `json:"PlatformDesc,omitempty" name:"PlatformDesc"`

	// 订单采购类型枚举值， 1-免费试用 2-按年收费 3-按次收费
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 订单采购类型 描述：1-免费试用 2-按年收费 3-按次收费
	OrderTypeDesc *string `json:"OrderTypeDesc,omitempty" name:"OrderTypeDesc"`

	// 枚举值：1-在线加固 或 2-输出工具加固
	EncryptOpType *int64 `json:"EncryptOpType,omitempty" name:"EncryptOpType"`

	// 描述：1-在线加固 或 2-输出工具加固
	EncryptOpTypeDesc *string `json:"EncryptOpTypeDesc,omitempty" name:"EncryptOpTypeDesc"`

	// 与当前任务关联的资源Id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 与当前任务关联的订单Id
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 对应PlatformType平台类型值   1-android安卓加固结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	AndroidResult *AndroidResult `json:"AndroidResult,omitempty" name:"AndroidResult"`

	// 对应PlatformType平台类型值   2-ios源码混淆加固结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	IOSResult *IOSResult `json:"IOSResult,omitempty" name:"IOSResult"`

	// 对应PlatformType平台类型值   3-sdk加固结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	SDKResult *SDKResult `json:"SDKResult,omitempty" name:"SDKResult"`

	// 对应PlatformType平台类型值   4-applet小程序加固结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppletResult *AppletResult `json:"AppletResult,omitempty" name:"AppletResult"`
}

type Filter struct {
	// 需要过滤的字段
	Name *string `json:"Name,omitempty" name:"Name"`

	// 需要过滤字段的值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type IOSPlan struct {
	// 策略id
	PlanId *int64 `json:"PlanId,omitempty" name:"PlanId"`
}

type IOSResult struct {
	// 加固任务结果Id
	ResultId *string `json:"ResultId,omitempty" name:"ResultId"`
}

type OptPluginListItem struct {
	// 非广告类型
	PluginType *string `json:"PluginType,omitempty" name:"PluginType"`

	// 非广告插件名称
	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`

	// 非广告插件描述
	PluginDesc *string `json:"PluginDesc,omitempty" name:"PluginDesc"`
}

type Orders struct {
	// 订单号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 平台类型整型值 
	PlatformType *int64 `json:"PlatformType,omitempty" name:"PlatformType"`

	// 平台类型描述：  1.android安卓加固   2.ios源码混淆  3.sdk加固  4.applet小程序加固
	PlatformTypeDesc *string `json:"PlatformTypeDesc,omitempty" name:"PlatformTypeDesc"`

	// 订单采购类型整型值
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 订单采购类型描述： 1-免费试用 2-按年收费 3-按次收费
	OrderTypeDesc *string `json:"OrderTypeDesc,omitempty" name:"OrderTypeDesc"`

	// 安卓包年收费加固的包名
	AppPkgName *string `json:"AppPkgName,omitempty" name:"AppPkgName"`

	// 资源号
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源状态整型值
	ResourceStatus *int64 `json:"ResourceStatus,omitempty" name:"ResourceStatus"`

	// 资源状态描述
	// 0-未生效、1-生效中、2-已失效。
	ResourceStatusDesc *string `json:"ResourceStatusDesc,omitempty" name:"ResourceStatusDesc"`

	// 订单类型为免费试用时的免费加固次数。
	TestTimes *int64 `json:"TestTimes,omitempty" name:"TestTimes"`

	// 资源生效时间
	ValidTime *string `json:"ValidTime,omitempty" name:"ValidTime"`

	// 资源过期时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 资源创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 订单审批人
	Approver *string `json:"Approver,omitempty" name:"Approver"`

	// 订单审批状态整型值
	ApprovalStatus *int64 `json:"ApprovalStatus,omitempty" name:"ApprovalStatus"`

	// 订单审批状态整型值描述：0-未审批、1-审批通过、2-驳回。
	ApprovalStatusDesc *string `json:"ApprovalStatusDesc,omitempty" name:"ApprovalStatusDesc"`

	// 订单审批时间
	ApprovalTime *string `json:"ApprovalTime,omitempty" name:"ApprovalTime"`

	// 按次收费加固资源，其关联的总任务数
	TimesTaskTotalCount *int64 `json:"TimesTaskTotalCount,omitempty" name:"TimesTaskTotalCount"`

	// 按次收费加固资源，其关联的任务成功数
	TimesTaskSuccessCount *int64 `json:"TimesTaskSuccessCount,omitempty" name:"TimesTaskSuccessCount"`

	// 按次收费加固资源，其关联的任务失败数
	TimesTaskFailCount *int64 `json:"TimesTaskFailCount,omitempty" name:"TimesTaskFailCount"`
}

type PlanDetailInfo struct {
	// 默认策略，1为默认，0为非默认
	IsDefault *uint64 `json:"IsDefault,omitempty" name:"IsDefault"`

	// 策略id
	PlanId *uint64 `json:"PlanId,omitempty" name:"PlanId"`

	// 策略名称
	PlanName *string `json:"PlanName,omitempty" name:"PlanName"`

	// 策略信息
	PlanInfo *PlanInfo `json:"PlanInfo,omitempty" name:"PlanInfo"`
}

type PlanInfo struct {
	// apk大小优化，0关闭，1开启
	ApkSizeOpt *uint64 `json:"ApkSizeOpt,omitempty" name:"ApkSizeOpt"`

	// Dex加固，0关闭，1开启
	Dex *uint64 `json:"Dex,omitempty" name:"Dex"`

	// So加固，0关闭，1开启
	So *uint64 `json:"So,omitempty" name:"So"`

	// 数据收集，0关闭，1开启
	Bugly *uint64 `json:"Bugly,omitempty" name:"Bugly"`

	// 防止重打包，0关闭，1开启
	AntiRepack *uint64 `json:"AntiRepack,omitempty" name:"AntiRepack"`

	// Dex分离，0关闭，1开启
	SeperateDex *uint64 `json:"SeperateDex,omitempty" name:"SeperateDex"`

	// 内存保护，0关闭，1开启
	Db *uint64 `json:"Db,omitempty" name:"Db"`

	// Dex签名校验，0关闭，1开启
	DexSig *uint64 `json:"DexSig,omitempty" name:"DexSig"`

	// So文件信息
	SoInfo *SoInfo `json:"SoInfo,omitempty" name:"SoInfo"`

	// vmp，0关闭，1开启
	AntiVMP *uint64 `json:"AntiVMP,omitempty" name:"AntiVMP"`

	// 保护so的强度，
	SoType []*string `json:"SoType,omitempty" name:"SoType"`

	// 防日志泄漏，0关闭，1开启
	AntiLogLeak *uint64 `json:"AntiLogLeak,omitempty" name:"AntiLogLeak"`

	// root检测，0关闭，1开启
	AntiQemuRoot *uint64 `json:"AntiQemuRoot,omitempty" name:"AntiQemuRoot"`

	// 资源防篡改，0关闭，1开启
	AntiAssets *uint64 `json:"AntiAssets,omitempty" name:"AntiAssets"`

	// 防止截屏，0关闭，1开启
	AntiScreenshot *uint64 `json:"AntiScreenshot,omitempty" name:"AntiScreenshot"`

	// SSL证书防窃取，0关闭，1开启
	AntiSSL *uint64 `json:"AntiSSL,omitempty" name:"AntiSSL"`
}

type PluginListItem struct {
	// 数字类型，分别为 1-通知栏广告，2-积分墙广告，3-banner广告，4- 悬浮窗图标广告，5-精品推荐列表广告, 6-插播广告
	PluginType *string `json:"PluginType,omitempty" name:"PluginType"`

	// 广告插件名称
	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`

	// 广告插件描述
	PluginDesc *string `json:"PluginDesc,omitempty" name:"PluginDesc"`
}

type ResourceInfo struct {
	// 用户购买的资源id，全局唯一
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源的pid，MTP加固-12767，应用加固-12750 MTP反作弊-12766 源代码混淆-12736
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`

	// 购买时间戳
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 到期时间戳
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 0-未绑定，1-已绑定
	IsBind *int64 `json:"IsBind,omitempty" name:"IsBind"`

	// 用户绑定app的基本信息
	BindInfo *BindInfo `json:"BindInfo,omitempty" name:"BindInfo"`

	// 资源名称，如应用加固，漏洞扫描
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
}

type ResourceServiceInfo struct {
	// 创建时间戳
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 到期时间戳
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 资源名称，如应用加固，源码混淆
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
}

type ResultListItem struct {
	// banner广告软件标记，分别为-1-不确定，0-否，1-是
	Banner *string `json:"Banner,omitempty" name:"Banner"`

	// 精品推荐列表广告标记，分别为-1-不确定，0-否，1-是
	BoutiqueRecommand *string `json:"BoutiqueRecommand,omitempty" name:"BoutiqueRecommand"`

	// 悬浮窗图标广告标记,分别为-1-不确定，0-否，1-是
	FloatWindows *string `json:"FloatWindows,omitempty" name:"FloatWindows"`

	// 积分墙广告软件标记，分别为 -1 -不确定，0-否，1-是
	IntegralWall *string `json:"IntegralWall,omitempty" name:"IntegralWall"`

	// 安装包的md5
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// 通知栏广告软件标记，分别为-1-不确定，0-否，1-是
	NotifyBar *string `json:"NotifyBar,omitempty" name:"NotifyBar"`

	// 1表示官方，0表示非官方
	Official *string `json:"Official,omitempty" name:"Official"`

	// 广告插件结果列表
	PluginList []*PluginListItem `json:"PluginList,omitempty" name:"PluginList"`

	// 非广告插件结果列表(SDK、风险插件等)
	OptPluginList []*OptPluginListItem `json:"OptPluginList,omitempty" name:"OptPluginList"`

	// 数字类型，分别为0-未知， 1-安全软件，2-风险软件，3-病毒软件
	SafeType *string `json:"SafeType,omitempty" name:"SafeType"`

	// Session id，合作方可以用来区分回调数据，需要唯一。
	Sid *string `json:"Sid,omitempty" name:"Sid"`

	// 安装包名称
	SoftName *string `json:"SoftName,omitempty" name:"SoftName"`

	// 插播广告软件标记，取值：-1 不确定，0否， 1 是
	Spot *string `json:"Spot,omitempty" name:"Spot"`

	// 病毒名称，utf8编码
	VirusName *string `json:"VirusName,omitempty" name:"VirusName"`

	// 病毒描述，utf8编码
	VirusDesc *string `json:"VirusDesc,omitempty" name:"VirusDesc"`

	// 二次打包状态：0-表示默认；1-表示二次
	RepackageStatus *string `json:"RepackageStatus,omitempty" name:"RepackageStatus"`

	// 应用错误码：0、1-表示正常；                  
	// 
	// 2表示System Error(engine analysis error).
	// 
	// 3表示App analysis error, please confirm it.
	// 
	// 4表示App have not cert, please confirm it.
	// 
	// 5表示App size is zero, please confirm it.
	// 
	// 6表示App have not package name, please confirm it.
	// 
	// 7表示App build time is empty, please confirm it.
	// 
	// 8表示App have not valid cert, please confirm it.
	// 
	// 99表示Other error.
	// 
	// 1000表示App downloadlink download fail, please confirm it.
	// 
	// 1001表示APP md5 different between real md5, please confirm it.
	// 
	// 1002表示App md5 uncollect, please offer downloadlink.
	Errno *string `json:"Errno,omitempty" name:"Errno"`

	// 对应errno的错误信息描述
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
}

type SDKPlan struct {
	// 策略id
	PlanId *int64 `json:"PlanId,omitempty" name:"PlanId"`
}

type SDKResult struct {
	// 加固任务结果Id
	ResultId *string `json:"ResultId,omitempty" name:"ResultId"`
}

type ServiceInfo struct {
	// 服务版本，基础版basic，专业版professional，企业版enterprise。
	ServiceEdition *string `json:"ServiceEdition,omitempty" name:"ServiceEdition"`

	// 任务处理完成后的反向通知回调地址，如果不需要通知请传递空字符串。通知为POST请求，post包体数据示例{"Response":{"ItemId":"4cdad8fb86f036b06bccb3f58971c306","ShieldCode":0,"ShieldMd5":"78701576793c4a5f04e1c9660de0aa0b","ShieldSize":11997354,"TaskStatus":1,"TaskTime":1539148141}}，调用方需要返回如下信息，{"Result":"ok","Reason":"xxxxx"}，如果Result字段值不等于ok会继续回调。
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 提交来源 YYB-应用宝 RDM-rdm MC-控制台 MAC_TOOL-mac工具 WIN_TOOL-window工具。
	SubmitSource *string `json:"SubmitSource,omitempty" name:"SubmitSource"`

	// 加固策略编号，如果不传则使用系统默认加固策略。如果指定的plan不存在会返回错误。
	PlanId *uint64 `json:"PlanId,omitempty" name:"PlanId"`
}

type ShieldInfo struct {
	// 加固结果的返回码
	ShieldCode *uint64 `json:"ShieldCode,omitempty" name:"ShieldCode"`

	// 加固后app的大小
	ShieldSize *uint64 `json:"ShieldSize,omitempty" name:"ShieldSize"`

	// 加固后app的md5
	ShieldMd5 *string `json:"ShieldMd5,omitempty" name:"ShieldMd5"`

	// 加固后的APP下载地址，该地址有效期为20分钟，请及时下载
	AppUrl *string `json:"AppUrl,omitempty" name:"AppUrl"`

	// 加固的提交时间
	TaskTime *uint64 `json:"TaskTime,omitempty" name:"TaskTime"`

	// 任务唯一标识
	ItemId *string `json:"ItemId,omitempty" name:"ItemId"`

	// 加固版本，basic基础版，professional专业版，enterprise企业版
	ServiceEdition *string `json:"ServiceEdition,omitempty" name:"ServiceEdition"`
}

type ShieldPlanInfo struct {
	// 加固策略数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 加固策略具体信息数组
	PlanSet []*PlanDetailInfo `json:"PlanSet,omitempty" name:"PlanSet"`
}

type SoInfo struct {
	// so文件列表
	SoFileNames []*string `json:"SoFileNames,omitempty" name:"SoFileNames"`
}