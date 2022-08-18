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

package v20181106

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-11-06"

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


func NewEvaluationRequest() (request *EvaluationRequest) {
    request = &EvaluationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hcm", APIVersion, "Evaluation")
    
    
    return
}

func NewEvaluationResponse() (response *EvaluationResponse) {
    response = &EvaluationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Evaluation
// 速算题目批改接口，根据用户上传的图片或图片的URL识别图片中的数学算式，进而给出算式的正确性评估。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CHARGECOUNTERROR = "FailedOperation.ChargeCountError"
//  INTERNALERROR_ENGINEREQUESTFAILED = "InternalError.EngineRequestFailed"
//  INTERNALERROR_ENGINERESULTERROR = "InternalError.EngineResultError"
//  INTERNALERROR_INITIALPARAMETERERROR = "InternalError.InitialParameterError"
//  INTERNALERROR_SERVERINTERNALERROR = "InternalError.ServerInternalError"
//  INVALIDPARAMETERVALUE_CANNOTFINDIMAGEERROR = "InvalidParameterValue.CannotFindImageError"
//  INVALIDPARAMETERVALUE_CANNOTFINDSESSION = "InvalidParameterValue.CannotFindSession"
//  INVALIDPARAMETERVALUE_EMPTYIMAGEERROR = "InvalidParameterValue.EmptyImageError"
//  INVALIDPARAMETERVALUE_EMPTYINPUTERROR = "InvalidParameterValue.EmptyInputError"
//  INVALIDPARAMETERVALUE_EXCEEDDOWNLOADIMAGESIZEERROR = "InvalidParameterValue.ExceedDownloadImageSizeError"
//  INVALIDPARAMETERVALUE_FAILDECODEERROR = "InvalidParameterValue.FailDecodeError"
//  INVALIDPARAMETERVALUE_FAILDOWNLOADIMAGEERROR = "InvalidParameterValue.FailDownloadImageError"
//  INVALIDPARAMETERVALUE_FAILRECOGNIZEERROR = "InvalidParameterValue.FailRecognizeError"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEERROR = "InvalidParameterValue.InvalidImageError"
//  RESOURCENOTFOUND_CANNOTFINDUSER = "ResourceNotFound.CannotFindUser"
func (c *Client) Evaluation(request *EvaluationRequest) (response *EvaluationResponse, err error) {
    return c.EvaluationWithContext(context.Background(), request)
}

// Evaluation
// 速算题目批改接口，根据用户上传的图片或图片的URL识别图片中的数学算式，进而给出算式的正确性评估。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CHARGECOUNTERROR = "FailedOperation.ChargeCountError"
//  INTERNALERROR_ENGINEREQUESTFAILED = "InternalError.EngineRequestFailed"
//  INTERNALERROR_ENGINERESULTERROR = "InternalError.EngineResultError"
//  INTERNALERROR_INITIALPARAMETERERROR = "InternalError.InitialParameterError"
//  INTERNALERROR_SERVERINTERNALERROR = "InternalError.ServerInternalError"
//  INVALIDPARAMETERVALUE_CANNOTFINDIMAGEERROR = "InvalidParameterValue.CannotFindImageError"
//  INVALIDPARAMETERVALUE_CANNOTFINDSESSION = "InvalidParameterValue.CannotFindSession"
//  INVALIDPARAMETERVALUE_EMPTYIMAGEERROR = "InvalidParameterValue.EmptyImageError"
//  INVALIDPARAMETERVALUE_EMPTYINPUTERROR = "InvalidParameterValue.EmptyInputError"
//  INVALIDPARAMETERVALUE_EXCEEDDOWNLOADIMAGESIZEERROR = "InvalidParameterValue.ExceedDownloadImageSizeError"
//  INVALIDPARAMETERVALUE_FAILDECODEERROR = "InvalidParameterValue.FailDecodeError"
//  INVALIDPARAMETERVALUE_FAILDOWNLOADIMAGEERROR = "InvalidParameterValue.FailDownloadImageError"
//  INVALIDPARAMETERVALUE_FAILRECOGNIZEERROR = "InvalidParameterValue.FailRecognizeError"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEERROR = "InvalidParameterValue.InvalidImageError"
//  RESOURCENOTFOUND_CANNOTFINDUSER = "ResourceNotFound.CannotFindUser"
func (c *Client) EvaluationWithContext(ctx context.Context, request *EvaluationRequest) (response *EvaluationResponse, err error) {
    if request == nil {
        request = NewEvaluationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("Evaluation require credential")
    }

    request.SetContext(ctx)
    
    response = NewEvaluationResponse()
    err = c.Send(request, response)
    return
}
