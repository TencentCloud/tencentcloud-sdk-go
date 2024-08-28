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

package v20230110

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2023-01-10"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewBatchApplyAccountBaselinesRequest() (request *BatchApplyAccountBaselinesRequest) {
    request = &BatchApplyAccountBaselinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("controlcenter", APIVersion, "BatchApplyAccountBaselines")
    
    
    return
}

func NewBatchApplyAccountBaselinesResponse() (response *BatchApplyAccountBaselinesResponse) {
    response = &BatchApplyAccountBaselinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchApplyAccountBaselines
// 批量对存量账号应用基线
//
// 可能返回的错误码:
//  FAILEDOPERATION_AFREQUIREDITEMNOTSELECT = "FailedOperation.AFRequiredItemNotSelect"
//  FAILEDOPERATION_ACCOUNTFACTORYMEMBERUINNUMEXCEED = "FailedOperation.AccountFactoryMemberUinNumExceed"
//  FAILEDOPERATION_ACCOUNTFACTORYTASKISDEPLOYING = "FailedOperation.AccountFactoryTaskIsDeploying"
//  FAILEDOPERATION_CONTROLCENTERNOTOPEN = "FailedOperation.ControlCenterNotOpen"
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_DEPENDONITEMNOTDEPLOY = "FailedOperation.DependOnItemNotDeploy"
//  FAILEDOPERATION_REMOTECALLERROR = "FailedOperation.RemoteCallError"
//  RESOURCENOTFOUND_ACCOUNTFACTORYBASELINENOTEXIST = "ResourceNotFound.AccountFactoryBaselineNotExist"
//  RESOURCENOTFOUND_ACCOUNTFACTORYITEMNOTCONFIG = "ResourceNotFound.AccountFactoryItemNotConfig"
//  RESOURCENOTFOUND_ACCOUNTFACTORYITEMNOTEXIST = "ResourceNotFound.AccountFactoryItemNotExist"
func (c *Client) BatchApplyAccountBaselines(request *BatchApplyAccountBaselinesRequest) (response *BatchApplyAccountBaselinesResponse, err error) {
    return c.BatchApplyAccountBaselinesWithContext(context.Background(), request)
}

// BatchApplyAccountBaselines
// 批量对存量账号应用基线
//
// 可能返回的错误码:
//  FAILEDOPERATION_AFREQUIREDITEMNOTSELECT = "FailedOperation.AFRequiredItemNotSelect"
//  FAILEDOPERATION_ACCOUNTFACTORYMEMBERUINNUMEXCEED = "FailedOperation.AccountFactoryMemberUinNumExceed"
//  FAILEDOPERATION_ACCOUNTFACTORYTASKISDEPLOYING = "FailedOperation.AccountFactoryTaskIsDeploying"
//  FAILEDOPERATION_CONTROLCENTERNOTOPEN = "FailedOperation.ControlCenterNotOpen"
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_DEPENDONITEMNOTDEPLOY = "FailedOperation.DependOnItemNotDeploy"
//  FAILEDOPERATION_REMOTECALLERROR = "FailedOperation.RemoteCallError"
//  RESOURCENOTFOUND_ACCOUNTFACTORYBASELINENOTEXIST = "ResourceNotFound.AccountFactoryBaselineNotExist"
//  RESOURCENOTFOUND_ACCOUNTFACTORYITEMNOTCONFIG = "ResourceNotFound.AccountFactoryItemNotConfig"
//  RESOURCENOTFOUND_ACCOUNTFACTORYITEMNOTEXIST = "ResourceNotFound.AccountFactoryItemNotExist"
func (c *Client) BatchApplyAccountBaselinesWithContext(ctx context.Context, request *BatchApplyAccountBaselinesRequest) (response *BatchApplyAccountBaselinesResponse, err error) {
    if request == nil {
        request = NewBatchApplyAccountBaselinesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchApplyAccountBaselines require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchApplyAccountBaselinesResponse()
    err = c.Send(request, response)
    return
}
