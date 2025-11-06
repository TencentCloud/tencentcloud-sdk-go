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

package v20230110

const (
	// 此产品的特有错误码

	// 账号工厂必选基线项未配置
	FAILEDOPERATION_AFREQUIREDITEMNOTSELECT = "FailedOperation.AFRequiredItemNotSelect"

	// 联系人邮箱未验证。
	FAILEDOPERATION_ACCOUNTFACTORYCONTACTEMAILNOTVERIFY = "FailedOperation.AccountFactoryContactEmailNotVerify"

	// 联系人手机未验证。
	FAILEDOPERATION_ACCOUNTFACTORYCONTACTPHONENOTVERIFY = "FailedOperation.AccountFactoryContactPhoneNotVerify"

	// 账号工厂基线项的依赖项未配置
	FAILEDOPERATION_ACCOUNTFACTORYDEPENDONITEMNOTCONFIG = "FailedOperation.AccountFactoryDependOnItemNotConfig"

	// 账号工厂批量应用基线的账号数量超过限制
	FAILEDOPERATION_ACCOUNTFACTORYMEMBERUINNUMEXCEED = "FailedOperation.AccountFactoryMemberUinNumExceed"

	// 用户正在部署账号基线中，无法重复部署
	FAILEDOPERATION_ACCOUNTFACTORYTASKISDEPLOYING = "FailedOperation.AccountFactoryTaskIsDeploying"

	// Control Center服务未开通
	FAILEDOPERATION_CONTROLCENTERNOTOPEN = "FailedOperation.ControlCenterNotOpen"

	// 数据库操作异常
	FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"

	// 依赖的功能项未部署
	FAILEDOPERATION_DEPENDONITEMNOTDEPLOY = "FailedOperation.DependOnItemNotDeploy"

	// 远程调用失败
	FAILEDOPERATION_REMOTECALLERROR = "FailedOperation.RemoteCallError"

	// 基线项预设标签数量超过最大数量
	INVALIDPARAMETER_ACCOUNTFACTORYTAGEXCEEDMAXNUM = "InvalidParameter.AccountFactoryTagExceedMaxNum"

	// 入参校验错误
	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"

	// 用户基线配置数据不存在
	RESOURCENOTFOUND_ACCOUNTFACTORYBASELINENOTEXIST = "ResourceNotFound.AccountFactoryBaselineNotExist"

	// 用户所部署的基线项未配置
	RESOURCENOTFOUND_ACCOUNTFACTORYITEMNOTCONFIG = "ResourceNotFound.AccountFactoryItemNotConfig"

	// 账号工厂基线不存在
	RESOURCENOTFOUND_ACCOUNTFACTORYITEMNOTEXIST = "ResourceNotFound.AccountFactoryItemNotExist"
)
