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

package v20210903

import (
    "context"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-09-03"

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


func NewAddCustomPersonImageRequest() (request *AddCustomPersonImageRequest) {
    request = &AddCustomPersonImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ivld", APIVersion, "AddCustomPersonImage")
    
    
    return
}

func NewAddCustomPersonImageResponse() (response *AddCustomPersonImageResponse) {
    response = &AddCustomPersonImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddCustomPersonImage
// 增加自定义人脸图片，每个自定义人物最多可包含5张人脸图片
//
// 
//
// 请注意，与创建自定义人物一样，图片数据优先级优于图片URL优先级
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  FAILEDOPERATION_DOWNLOADFAILED = "FailedOperation.DownloadFailed"
//  FAILEDOPERATION_FEATUREALGOFAILED = "FailedOperation.FeatureAlgoFailed"
//  FAILEDOPERATION_IMAGENUMEXCEEDED = "FailedOperation.ImageNumExceeded"
//  FAILEDOPERATION_MULTIPLEFACESINIMAGE = "FailedOperation.MultipleFacesInImage"
//  FAILEDOPERATION_NOFACEINIMAGE = "FailedOperation.NoFaceInImage"
//  FAILEDOPERATION_PERSONNOTMATCHED = "FailedOperation.PersonNotMatched"
//  FAILEDOPERATION_QUALITYALGOFAILED = "FailedOperation.QualityAlgoFailed"
//  FAILEDOPERATION_QUALITYTOOLOW = "FailedOperation.QualityTooLow"
//  INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  INTERNALERROR_INTERNALOVERFLOW = "InternalError.InternalOverflow"
//  INVALIDPARAMETER_INVALIDIMAGE = "InvalidParameter.InvalidImage"
//  INVALIDPARAMETER_INVALIDPERSONID = "InvalidParameter.InvalidPersonId"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidURL"
//  RESOURCENOTFOUND_RECORDNOTFOUND = "ResourceNotFound.RecordNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) AddCustomPersonImage(request *AddCustomPersonImageRequest) (response *AddCustomPersonImageResponse, err error) {
    if request == nil {
        request = NewAddCustomPersonImageRequest()
    }
    
    response = NewAddCustomPersonImageResponse()
    err = c.Send(request, response)
    return
}

// AddCustomPersonImage
// 增加自定义人脸图片，每个自定义人物最多可包含5张人脸图片
//
// 
//
// 请注意，与创建自定义人物一样，图片数据优先级优于图片URL优先级
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  FAILEDOPERATION_DOWNLOADFAILED = "FailedOperation.DownloadFailed"
//  FAILEDOPERATION_FEATUREALGOFAILED = "FailedOperation.FeatureAlgoFailed"
//  FAILEDOPERATION_IMAGENUMEXCEEDED = "FailedOperation.ImageNumExceeded"
//  FAILEDOPERATION_MULTIPLEFACESINIMAGE = "FailedOperation.MultipleFacesInImage"
//  FAILEDOPERATION_NOFACEINIMAGE = "FailedOperation.NoFaceInImage"
//  FAILEDOPERATION_PERSONNOTMATCHED = "FailedOperation.PersonNotMatched"
//  FAILEDOPERATION_QUALITYALGOFAILED = "FailedOperation.QualityAlgoFailed"
//  FAILEDOPERATION_QUALITYTOOLOW = "FailedOperation.QualityTooLow"
//  INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  INTERNALERROR_INTERNALOVERFLOW = "InternalError.InternalOverflow"
//  INVALIDPARAMETER_INVALIDIMAGE = "InvalidParameter.InvalidImage"
//  INVALIDPARAMETER_INVALIDPERSONID = "InvalidParameter.InvalidPersonId"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidURL"
//  RESOURCENOTFOUND_RECORDNOTFOUND = "ResourceNotFound.RecordNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) AddCustomPersonImageWithContext(ctx context.Context, request *AddCustomPersonImageRequest) (response *AddCustomPersonImageResponse, err error) {
    if request == nil {
        request = NewAddCustomPersonImageRequest()
    }
    request.SetContext(ctx)
    
    response = NewAddCustomPersonImageResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCustomCategoryRequest() (request *CreateCustomCategoryRequest) {
    request = &CreateCustomCategoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ivld", APIVersion, "CreateCustomCategory")
    
    
    return
}

func NewCreateCustomCategoryResponse() (response *CreateCustomCategoryResponse) {
    response = &CreateCustomCategoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCustomCategory
// 创建自定义人物分类信息
//
// 
//
// 当L2Category为空时，将创建一级自定义分类。
//
// 当L1Category与L2Category均不为空时，将创建二级自定义分类。请注意，**只有当一级自定义分类存在时，才可创建二级自定义分类**。
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  FAILEDOPERATION_CATEGORYEXIST = "FailedOperation.CategoryExist"
//  INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  INVALIDPARAMETER_INVALIDL1CATEGORY = "InvalidParameter.InvalidL1Category"
//  INVALIDPARAMETER_INVALIDL2CATEGORY = "InvalidParameter.InvalidL2Category"
//  RESOURCENOTFOUND_CUSTOMCATEGORYNOTFOUND = "ResourceNotFound.CustomCategoryNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) CreateCustomCategory(request *CreateCustomCategoryRequest) (response *CreateCustomCategoryResponse, err error) {
    if request == nil {
        request = NewCreateCustomCategoryRequest()
    }
    
    response = NewCreateCustomCategoryResponse()
    err = c.Send(request, response)
    return
}

// CreateCustomCategory
// 创建自定义人物分类信息
//
// 
//
// 当L2Category为空时，将创建一级自定义分类。
//
// 当L1Category与L2Category均不为空时，将创建二级自定义分类。请注意，**只有当一级自定义分类存在时，才可创建二级自定义分类**。
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  FAILEDOPERATION_CATEGORYEXIST = "FailedOperation.CategoryExist"
//  INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  INVALIDPARAMETER_INVALIDL1CATEGORY = "InvalidParameter.InvalidL1Category"
//  INVALIDPARAMETER_INVALIDL2CATEGORY = "InvalidParameter.InvalidL2Category"
//  RESOURCENOTFOUND_CUSTOMCATEGORYNOTFOUND = "ResourceNotFound.CustomCategoryNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) CreateCustomCategoryWithContext(ctx context.Context, request *CreateCustomCategoryRequest) (response *CreateCustomCategoryResponse, err error) {
    if request == nil {
        request = NewCreateCustomCategoryRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateCustomCategoryResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCustomGroupRequest() (request *CreateCustomGroupRequest) {
    request = &CreateCustomGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ivld", APIVersion, "CreateCustomGroup")
    
    
    return
}

func NewCreateCustomGroupResponse() (response *CreateCustomGroupResponse) {
    response = &CreateCustomGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCustomGroup
// 创建自定义人物库
//
// 
//
// Bucket的格式参考为 `bucketName-123456.cos.ap-shanghai.myqcloud.com`
//
// 
//
// 在调用CreateCustomPerson和AddCustomPersonImage接口之前，请先确保本接口成功调用。当前每个用户只支持一个自定义人物库，一旦自定义人物库创建成功，后续接口调用均会返回人物库已存在错误。
//
// 
//
// 由于人脸图片对于自定义人物识别至关重要，因此自定义人物识别功能需要用户显式指定COS存储桶方可使用。具体来说，自定义人物识别功能接口(主要是CreateCustomPerson和AddCustomPersonImage)会在此COS桶下面新建IVLDCustomPersonImage目录，并在此目录下存储自定义人物图片数据以支持后续潜在的特征更新。
//
// 
//
// 请注意：本接口指定的COS桶仅用于**备份存储自定义人物图片**，CreateCustomPerson和AddCustomPersonImage接口入参URL可使用任意COS存储桶下的任意图片。
//
// 
//
// **重要**：请务必确保本接口指定的COS存储桶存在(不要手动删除COS桶)。COS存储桶一旦指定，将不能修改。
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  FAILEDOPERATION_CUSTOMGROUPALREADYEXIST = "FailedOperation.CustomGroupAlreadyExist"
//  INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidURL"
//  INVALIDPARAMETER_UNSUPPORTURL = "InvalidParameter.UnsupportURL"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) CreateCustomGroup(request *CreateCustomGroupRequest) (response *CreateCustomGroupResponse, err error) {
    if request == nil {
        request = NewCreateCustomGroupRequest()
    }
    
    response = NewCreateCustomGroupResponse()
    err = c.Send(request, response)
    return
}

// CreateCustomGroup
// 创建自定义人物库
//
// 
//
// Bucket的格式参考为 `bucketName-123456.cos.ap-shanghai.myqcloud.com`
//
// 
//
// 在调用CreateCustomPerson和AddCustomPersonImage接口之前，请先确保本接口成功调用。当前每个用户只支持一个自定义人物库，一旦自定义人物库创建成功，后续接口调用均会返回人物库已存在错误。
//
// 
//
// 由于人脸图片对于自定义人物识别至关重要，因此自定义人物识别功能需要用户显式指定COS存储桶方可使用。具体来说，自定义人物识别功能接口(主要是CreateCustomPerson和AddCustomPersonImage)会在此COS桶下面新建IVLDCustomPersonImage目录，并在此目录下存储自定义人物图片数据以支持后续潜在的特征更新。
//
// 
//
// 请注意：本接口指定的COS桶仅用于**备份存储自定义人物图片**，CreateCustomPerson和AddCustomPersonImage接口入参URL可使用任意COS存储桶下的任意图片。
//
// 
//
// **重要**：请务必确保本接口指定的COS存储桶存在(不要手动删除COS桶)。COS存储桶一旦指定，将不能修改。
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  FAILEDOPERATION_CUSTOMGROUPALREADYEXIST = "FailedOperation.CustomGroupAlreadyExist"
//  INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidURL"
//  INVALIDPARAMETER_UNSUPPORTURL = "InvalidParameter.UnsupportURL"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) CreateCustomGroupWithContext(ctx context.Context, request *CreateCustomGroupRequest) (response *CreateCustomGroupResponse, err error) {
    if request == nil {
        request = NewCreateCustomGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateCustomGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCustomPersonRequest() (request *CreateCustomPersonRequest) {
    request = &CreateCustomPersonRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ivld", APIVersion, "CreateCustomPerson")
    
    
    return
}

func NewCreateCustomPersonResponse() (response *CreateCustomPersonResponse) {
    response = &CreateCustomPersonResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCustomPerson
// 创建自定义人物。
//
// 
//
// 输入人物名称，基本信息，分类信息与人脸图片，创建自定义人物
//
// 
//
// 人脸图片可使用图片数据(base64编码的图片数据)或者图片URL(推荐使用COS以减少下载时间，其他地址也支持)，原始图片优先，也即如果同时指定了图片数据和图片URL，接口将仅使用图片数据
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDSECRETID = "AuthFailure.InvalidSecretId"
//  AUTHFAILURE_SECRETIDNOTFOUND = "AuthFailure.SecretIdNotFound"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  FAILEDOPERATION_DOWNLOADFAILED = "FailedOperation.DownloadFailed"
//  FAILEDOPERATION_FEATUREALGOFAILED = "FailedOperation.FeatureAlgoFailed"
//  FAILEDOPERATION_MULTIPLEFACESINIMAGE = "FailedOperation.MultipleFacesInImage"
//  FAILEDOPERATION_NOFACEINIMAGE = "FailedOperation.NoFaceInImage"
//  FAILEDOPERATION_PERSONDUPLICATED = "FailedOperation.PersonDuplicated"
//  FAILEDOPERATION_PERSONNUMEXCEEDED = "FailedOperation.PersonNumExceeded"
//  FAILEDOPERATION_QUALITYALGOFAILED = "FailedOperation.QualityAlgoFailed"
//  FAILEDOPERATION_QUALITYTOOLOW = "FailedOperation.QualityTooLow"
//  INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  INTERNALERROR_INTERNALOVERFLOW = "InternalError.InternalOverflow"
//  INVALIDPARAMETER_INVALIDCATEGORYID = "InvalidParameter.InvalidCategoryId"
//  INVALIDPARAMETER_INVALIDIMAGE = "InvalidParameter.InvalidImage"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_INVALIDPARAM = "InvalidParameter.InvalidParam"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidURL"
//  INVALIDPARAMETER_UNSUPPORTURL = "InvalidParameter.UnsupportURL"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) CreateCustomPerson(request *CreateCustomPersonRequest) (response *CreateCustomPersonResponse, err error) {
    if request == nil {
        request = NewCreateCustomPersonRequest()
    }
    
    response = NewCreateCustomPersonResponse()
    err = c.Send(request, response)
    return
}

// CreateCustomPerson
// 创建自定义人物。
//
// 
//
// 输入人物名称，基本信息，分类信息与人脸图片，创建自定义人物
//
// 
//
// 人脸图片可使用图片数据(base64编码的图片数据)或者图片URL(推荐使用COS以减少下载时间，其他地址也支持)，原始图片优先，也即如果同时指定了图片数据和图片URL，接口将仅使用图片数据
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDSECRETID = "AuthFailure.InvalidSecretId"
//  AUTHFAILURE_SECRETIDNOTFOUND = "AuthFailure.SecretIdNotFound"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  FAILEDOPERATION_DOWNLOADFAILED = "FailedOperation.DownloadFailed"
//  FAILEDOPERATION_FEATUREALGOFAILED = "FailedOperation.FeatureAlgoFailed"
//  FAILEDOPERATION_MULTIPLEFACESINIMAGE = "FailedOperation.MultipleFacesInImage"
//  FAILEDOPERATION_NOFACEINIMAGE = "FailedOperation.NoFaceInImage"
//  FAILEDOPERATION_PERSONDUPLICATED = "FailedOperation.PersonDuplicated"
//  FAILEDOPERATION_PERSONNUMEXCEEDED = "FailedOperation.PersonNumExceeded"
//  FAILEDOPERATION_QUALITYALGOFAILED = "FailedOperation.QualityAlgoFailed"
//  FAILEDOPERATION_QUALITYTOOLOW = "FailedOperation.QualityTooLow"
//  INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  INTERNALERROR_INTERNALOVERFLOW = "InternalError.InternalOverflow"
//  INVALIDPARAMETER_INVALIDCATEGORYID = "InvalidParameter.InvalidCategoryId"
//  INVALIDPARAMETER_INVALIDIMAGE = "InvalidParameter.InvalidImage"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_INVALIDPARAM = "InvalidParameter.InvalidParam"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidURL"
//  INVALIDPARAMETER_UNSUPPORTURL = "InvalidParameter.UnsupportURL"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) CreateCustomPersonWithContext(ctx context.Context, request *CreateCustomPersonRequest) (response *CreateCustomPersonResponse, err error) {
    if request == nil {
        request = NewCreateCustomPersonRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateCustomPersonResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDefaultCategoriesRequest() (request *CreateDefaultCategoriesRequest) {
    request = &CreateDefaultCategoriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ivld", APIVersion, "CreateDefaultCategories")
    
    
    return
}

func NewCreateDefaultCategoriesResponse() (response *CreateDefaultCategoriesResponse) {
    response = &CreateDefaultCategoriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDefaultCategories
// 创建默认自定义人物类型
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) CreateDefaultCategories(request *CreateDefaultCategoriesRequest) (response *CreateDefaultCategoriesResponse, err error) {
    if request == nil {
        request = NewCreateDefaultCategoriesRequest()
    }
    
    response = NewCreateDefaultCategoriesResponse()
    err = c.Send(request, response)
    return
}

// CreateDefaultCategories
// 创建默认自定义人物类型
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) CreateDefaultCategoriesWithContext(ctx context.Context, request *CreateDefaultCategoriesRequest) (response *CreateDefaultCategoriesResponse, err error) {
    if request == nil {
        request = NewCreateDefaultCategoriesRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateDefaultCategoriesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTaskRequest() (request *CreateTaskRequest) {
    request = &CreateTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ivld", APIVersion, "CreateTask")
    
    
    return
}

func NewCreateTaskResponse() (response *CreateTaskResponse) {
    response = &CreateTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTask
// 创建智能标签任务。
//
// 
//
// 请注意，本接口为异步接口，**返回TaskId只代表任务创建成功，不代表任务执行成功**。
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  FAILEDOPERATION_AITEMPLATENOTEXIST = "FailedOperation.AiTemplateNotExist"
//  FAILEDOPERATION_MEDIAEXPIRED = "FailedOperation.MediaExpired"
//  FAILEDOPERATION_MEDIANOTREADY = "FailedOperation.MediaNotReady"
//  FAILEDOPERATION_TASKALREADYEXIST = "FailedOperation.TaskAlreadyExist"
//  FAILEDOPERATION_TASKNOTFINISHED = "FailedOperation.TaskNotFinished"
//  INVALIDPARAMETER_INVALIDMEDIAID = "InvalidParameter.InvalidMediaId"
//  INVALIDPARAMETER_INVALIDMEDIALABEL = "InvalidParameter.InvalidMediaLabel"
//  INVALIDPARAMETER_INVALIDMEDIALANG = "InvalidParameter.InvalidMediaLang"
//  INVALIDPARAMETER_INVALIDMEDIANAME = "InvalidParameter.InvalidMediaName"
//  INVALIDPARAMETER_INVALIDMEDIATYPE = "InvalidParameter.InvalidMediaType"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidURL"
//  RESOURCENOTFOUND_MEDIANOTFOUND = "ResourceNotFound.MediaNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) CreateTask(request *CreateTaskRequest) (response *CreateTaskResponse, err error) {
    if request == nil {
        request = NewCreateTaskRequest()
    }
    
    response = NewCreateTaskResponse()
    err = c.Send(request, response)
    return
}

// CreateTask
// 创建智能标签任务。
//
// 
//
// 请注意，本接口为异步接口，**返回TaskId只代表任务创建成功，不代表任务执行成功**。
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  FAILEDOPERATION_AITEMPLATENOTEXIST = "FailedOperation.AiTemplateNotExist"
//  FAILEDOPERATION_MEDIAEXPIRED = "FailedOperation.MediaExpired"
//  FAILEDOPERATION_MEDIANOTREADY = "FailedOperation.MediaNotReady"
//  FAILEDOPERATION_TASKALREADYEXIST = "FailedOperation.TaskAlreadyExist"
//  FAILEDOPERATION_TASKNOTFINISHED = "FailedOperation.TaskNotFinished"
//  INVALIDPARAMETER_INVALIDMEDIAID = "InvalidParameter.InvalidMediaId"
//  INVALIDPARAMETER_INVALIDMEDIALABEL = "InvalidParameter.InvalidMediaLabel"
//  INVALIDPARAMETER_INVALIDMEDIALANG = "InvalidParameter.InvalidMediaLang"
//  INVALIDPARAMETER_INVALIDMEDIANAME = "InvalidParameter.InvalidMediaName"
//  INVALIDPARAMETER_INVALIDMEDIATYPE = "InvalidParameter.InvalidMediaType"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidURL"
//  RESOURCENOTFOUND_MEDIANOTFOUND = "ResourceNotFound.MediaNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) CreateTaskWithContext(ctx context.Context, request *CreateTaskRequest) (response *CreateTaskResponse, err error) {
    if request == nil {
        request = NewCreateTaskRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCustomCategoryRequest() (request *DeleteCustomCategoryRequest) {
    request = &DeleteCustomCategoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ivld", APIVersion, "DeleteCustomCategory")
    
    
    return
}

func NewDeleteCustomCategoryResponse() (response *DeleteCustomCategoryResponse) {
    response = &DeleteCustomCategoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCustomCategory
// 删除自定义分类信息
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  FAILEDOPERATION_CATEGORYREFERRED = "FailedOperation.CategoryReferred"
//  INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  INVALIDPARAMETER_INVALIDCATEGORYID = "InvalidParameter.InvalidCategoryId"
//  RESOURCENOTFOUND_CUSTOMCATEGORYNOTFOUND = "ResourceNotFound.CustomCategoryNotFound"
//  RESOURCENOTFOUND_RECORDNOTFOUND = "ResourceNotFound.RecordNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) DeleteCustomCategory(request *DeleteCustomCategoryRequest) (response *DeleteCustomCategoryResponse, err error) {
    if request == nil {
        request = NewDeleteCustomCategoryRequest()
    }
    
    response = NewDeleteCustomCategoryResponse()
    err = c.Send(request, response)
    return
}

// DeleteCustomCategory
// 删除自定义分类信息
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  FAILEDOPERATION_CATEGORYREFERRED = "FailedOperation.CategoryReferred"
//  INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  INVALIDPARAMETER_INVALIDCATEGORYID = "InvalidParameter.InvalidCategoryId"
//  RESOURCENOTFOUND_CUSTOMCATEGORYNOTFOUND = "ResourceNotFound.CustomCategoryNotFound"
//  RESOURCENOTFOUND_RECORDNOTFOUND = "ResourceNotFound.RecordNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) DeleteCustomCategoryWithContext(ctx context.Context, request *DeleteCustomCategoryRequest) (response *DeleteCustomCategoryResponse, err error) {
    if request == nil {
        request = NewDeleteCustomCategoryRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteCustomCategoryResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCustomPersonRequest() (request *DeleteCustomPersonRequest) {
    request = &DeleteCustomPersonRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ivld", APIVersion, "DeleteCustomPerson")
    
    
    return
}

func NewDeleteCustomPersonResponse() (response *DeleteCustomPersonResponse) {
    response = &DeleteCustomPersonResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCustomPerson
// 删除自定义人物
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  INVALIDPARAMETER_INVALIDPERSONID = "InvalidParameter.InvalidPersonId"
//  RESOURCENOTFOUND_RECORDNOTFOUND = "ResourceNotFound.RecordNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) DeleteCustomPerson(request *DeleteCustomPersonRequest) (response *DeleteCustomPersonResponse, err error) {
    if request == nil {
        request = NewDeleteCustomPersonRequest()
    }
    
    response = NewDeleteCustomPersonResponse()
    err = c.Send(request, response)
    return
}

// DeleteCustomPerson
// 删除自定义人物
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  INVALIDPARAMETER_INVALIDPERSONID = "InvalidParameter.InvalidPersonId"
//  RESOURCENOTFOUND_RECORDNOTFOUND = "ResourceNotFound.RecordNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) DeleteCustomPersonWithContext(ctx context.Context, request *DeleteCustomPersonRequest) (response *DeleteCustomPersonResponse, err error) {
    if request == nil {
        request = NewDeleteCustomPersonRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteCustomPersonResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCustomPersonImageRequest() (request *DeleteCustomPersonImageRequest) {
    request = &DeleteCustomPersonImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ivld", APIVersion, "DeleteCustomPersonImage")
    
    
    return
}

func NewDeleteCustomPersonImageResponse() (response *DeleteCustomPersonImageResponse) {
    response = &DeleteCustomPersonImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCustomPersonImage
// 删除自定义人脸数据
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  FAILEDOPERATION_IMAGENUMEXCEEDED = "FailedOperation.ImageNumExceeded"
//  INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  INTERNALERROR_INTERNALOVERFLOW = "InternalError.InternalOverflow"
//  INVALIDPARAMETER_INVALIDIMAGEID = "InvalidParameter.InvalidImageId"
//  INVALIDPARAMETER_INVALIDPERSONID = "InvalidParameter.InvalidPersonId"
//  RESOURCENOTFOUND_RECORDNOTFOUND = "ResourceNotFound.RecordNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) DeleteCustomPersonImage(request *DeleteCustomPersonImageRequest) (response *DeleteCustomPersonImageResponse, err error) {
    if request == nil {
        request = NewDeleteCustomPersonImageRequest()
    }
    
    response = NewDeleteCustomPersonImageResponse()
    err = c.Send(request, response)
    return
}

// DeleteCustomPersonImage
// 删除自定义人脸数据
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  FAILEDOPERATION_IMAGENUMEXCEEDED = "FailedOperation.ImageNumExceeded"
//  INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  INTERNALERROR_INTERNALOVERFLOW = "InternalError.InternalOverflow"
//  INVALIDPARAMETER_INVALIDIMAGEID = "InvalidParameter.InvalidImageId"
//  INVALIDPARAMETER_INVALIDPERSONID = "InvalidParameter.InvalidPersonId"
//  RESOURCENOTFOUND_RECORDNOTFOUND = "ResourceNotFound.RecordNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) DeleteCustomPersonImageWithContext(ctx context.Context, request *DeleteCustomPersonImageRequest) (response *DeleteCustomPersonImageResponse, err error) {
    if request == nil {
        request = NewDeleteCustomPersonImageRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteCustomPersonImageResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMediaRequest() (request *DeleteMediaRequest) {
    request = &DeleteMediaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ivld", APIVersion, "DeleteMedia")
    
    
    return
}

func NewDeleteMediaResponse() (response *DeleteMediaResponse) {
    response = &DeleteMediaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMedia
// 将MediaId对应的媒资文件从系统中删除。
//
// 
//
// **请注意，本接口仅删除媒资文件，媒资文件对应的视频分析结果不会被删除**。如您需要删除结构化分析结果，请调用DeleteTask接口。
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  FAILEDOPERATION_MEDIAALREADYEXIST = "FailedOperation.MediaAlreadyExist"
//  FAILEDOPERATION_MEDIAINUSE = "FailedOperation.MediaInUse"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDMEDIAID = "InvalidParameter.InvalidMediaId"
//  INVALIDPARAMETER_INVALIDMEDIASTATUS = "InvalidParameter.InvalidMediaStatus"
//  RESOURCENOTFOUND_MEDIANOTFOUND = "ResourceNotFound.MediaNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_MEDIANOTACCESSIBLE = "UnsupportedOperation.MediaNotAccessible"
func (c *Client) DeleteMedia(request *DeleteMediaRequest) (response *DeleteMediaResponse, err error) {
    if request == nil {
        request = NewDeleteMediaRequest()
    }
    
    response = NewDeleteMediaResponse()
    err = c.Send(request, response)
    return
}

// DeleteMedia
// 将MediaId对应的媒资文件从系统中删除。
//
// 
//
// **请注意，本接口仅删除媒资文件，媒资文件对应的视频分析结果不会被删除**。如您需要删除结构化分析结果，请调用DeleteTask接口。
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  FAILEDOPERATION_MEDIAALREADYEXIST = "FailedOperation.MediaAlreadyExist"
//  FAILEDOPERATION_MEDIAINUSE = "FailedOperation.MediaInUse"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDMEDIAID = "InvalidParameter.InvalidMediaId"
//  INVALIDPARAMETER_INVALIDMEDIASTATUS = "InvalidParameter.InvalidMediaStatus"
//  RESOURCENOTFOUND_MEDIANOTFOUND = "ResourceNotFound.MediaNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_MEDIANOTACCESSIBLE = "UnsupportedOperation.MediaNotAccessible"
func (c *Client) DeleteMediaWithContext(ctx context.Context, request *DeleteMediaRequest) (response *DeleteMediaResponse, err error) {
    if request == nil {
        request = NewDeleteMediaRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteMediaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomCategoriesRequest() (request *DescribeCustomCategoriesRequest) {
    request = &DescribeCustomCategoriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ivld", APIVersion, "DescribeCustomCategories")
    
    
    return
}

func NewDescribeCustomCategoriesResponse() (response *DescribeCustomCategoriesResponse) {
    response = &DescribeCustomCategoriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCustomCategories
// 批量描述自定义人物分类信息
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) DescribeCustomCategories(request *DescribeCustomCategoriesRequest) (response *DescribeCustomCategoriesResponse, err error) {
    if request == nil {
        request = NewDescribeCustomCategoriesRequest()
    }
    
    response = NewDescribeCustomCategoriesResponse()
    err = c.Send(request, response)
    return
}

// DescribeCustomCategories
// 批量描述自定义人物分类信息
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) DescribeCustomCategoriesWithContext(ctx context.Context, request *DescribeCustomCategoriesRequest) (response *DescribeCustomCategoriesResponse, err error) {
    if request == nil {
        request = NewDescribeCustomCategoriesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCustomCategoriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomGroupRequest() (request *DescribeCustomGroupRequest) {
    request = &DescribeCustomGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ivld", APIVersion, "DescribeCustomGroup")
    
    
    return
}

func NewDescribeCustomGroupResponse() (response *DescribeCustomGroupResponse) {
    response = &DescribeCustomGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCustomGroup
// 描述自定义人物库信息，当前库大小(库中有多少人脸)，以及库中的存储桶
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  RESOURCENOTFOUND_CUSTOMGROUPNOTFOUND = "ResourceNotFound.CustomGroupNotFound"
//  RESOURCENOTFOUND_RECORDNOTFOUND = "ResourceNotFound.RecordNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) DescribeCustomGroup(request *DescribeCustomGroupRequest) (response *DescribeCustomGroupResponse, err error) {
    if request == nil {
        request = NewDescribeCustomGroupRequest()
    }
    
    response = NewDescribeCustomGroupResponse()
    err = c.Send(request, response)
    return
}

// DescribeCustomGroup
// 描述自定义人物库信息，当前库大小(库中有多少人脸)，以及库中的存储桶
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  RESOURCENOTFOUND_CUSTOMGROUPNOTFOUND = "ResourceNotFound.CustomGroupNotFound"
//  RESOURCENOTFOUND_RECORDNOTFOUND = "ResourceNotFound.RecordNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) DescribeCustomGroupWithContext(ctx context.Context, request *DescribeCustomGroupRequest) (response *DescribeCustomGroupResponse, err error) {
    if request == nil {
        request = NewDescribeCustomGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCustomGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomPersonDetailRequest() (request *DescribeCustomPersonDetailRequest) {
    request = &DescribeCustomPersonDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ivld", APIVersion, "DescribeCustomPersonDetail")
    
    
    return
}

func NewDescribeCustomPersonDetailResponse() (response *DescribeCustomPersonDetailResponse) {
    response = &DescribeCustomPersonDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCustomPersonDetail
// 描述自定义人物详细信息，包括人物信息与人物信息
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  INVALIDPARAMETER_INVALIDPERSONID = "InvalidParameter.InvalidPersonId"
//  RESOURCENOTFOUND_RECORDNOTFOUND = "ResourceNotFound.RecordNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) DescribeCustomPersonDetail(request *DescribeCustomPersonDetailRequest) (response *DescribeCustomPersonDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCustomPersonDetailRequest()
    }
    
    response = NewDescribeCustomPersonDetailResponse()
    err = c.Send(request, response)
    return
}

// DescribeCustomPersonDetail
// 描述自定义人物详细信息，包括人物信息与人物信息
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  INVALIDPARAMETER_INVALIDPERSONID = "InvalidParameter.InvalidPersonId"
//  RESOURCENOTFOUND_RECORDNOTFOUND = "ResourceNotFound.RecordNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) DescribeCustomPersonDetailWithContext(ctx context.Context, request *DescribeCustomPersonDetailRequest) (response *DescribeCustomPersonDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCustomPersonDetailRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCustomPersonDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomPersonsRequest() (request *DescribeCustomPersonsRequest) {
    request = &DescribeCustomPersonsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ivld", APIVersion, "DescribeCustomPersons")
    
    
    return
}

func NewDescribeCustomPersonsResponse() (response *DescribeCustomPersonsResponse) {
    response = &DescribeCustomPersonsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCustomPersons
// 批量描述自定义人物
//
// 
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  INVALIDPARAMETER_INVALIDCATEGORYID = "InvalidParameter.InvalidCategoryId"
//  INVALIDPARAMETER_INVALIDL1CATEGORY = "InvalidParameter.InvalidL1Category"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_INVALIDPAGENUMBER = "InvalidParameter.InvalidPageNumber"
//  INVALIDPARAMETER_INVALIDPAGESIZE = "InvalidParameter.InvalidPageSize"
//  INVALIDPARAMETER_INVALIDPARAM = "InvalidParameter.InvalidParam"
//  INVALIDPARAMETER_INVALIDPERSONID = "InvalidParameter.InvalidPersonId"
//  INVALIDPARAMETER_INVALIDSORTBY = "InvalidParameter.InvalidSortBy"
//  RESOURCENOTFOUND_CUSTOMCATEGORYNOTFOUND = "ResourceNotFound.CustomCategoryNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) DescribeCustomPersons(request *DescribeCustomPersonsRequest) (response *DescribeCustomPersonsResponse, err error) {
    if request == nil {
        request = NewDescribeCustomPersonsRequest()
    }
    
    response = NewDescribeCustomPersonsResponse()
    err = c.Send(request, response)
    return
}

// DescribeCustomPersons
// 批量描述自定义人物
//
// 
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  INVALIDPARAMETER_INVALIDCATEGORYID = "InvalidParameter.InvalidCategoryId"
//  INVALIDPARAMETER_INVALIDL1CATEGORY = "InvalidParameter.InvalidL1Category"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_INVALIDPAGENUMBER = "InvalidParameter.InvalidPageNumber"
//  INVALIDPARAMETER_INVALIDPAGESIZE = "InvalidParameter.InvalidPageSize"
//  INVALIDPARAMETER_INVALIDPARAM = "InvalidParameter.InvalidParam"
//  INVALIDPARAMETER_INVALIDPERSONID = "InvalidParameter.InvalidPersonId"
//  INVALIDPARAMETER_INVALIDSORTBY = "InvalidParameter.InvalidSortBy"
//  RESOURCENOTFOUND_CUSTOMCATEGORYNOTFOUND = "ResourceNotFound.CustomCategoryNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) DescribeCustomPersonsWithContext(ctx context.Context, request *DescribeCustomPersonsRequest) (response *DescribeCustomPersonsResponse, err error) {
    if request == nil {
        request = NewDescribeCustomPersonsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCustomPersonsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMediaRequest() (request *DescribeMediaRequest) {
    request = &DescribeMediaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ivld", APIVersion, "DescribeMedia")
    
    
    return
}

func NewDescribeMediaResponse() (response *DescribeMediaResponse) {
    response = &DescribeMediaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMedia
// 描述媒资文件信息，包括媒资状态，分辨率，帧率等。
//
// 
//
// 如果媒资文件未完成导入，本接口将仅输出媒资文件的状态信息；导入完成后，本接口还将输出媒资文件的其他元信息。
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  FAILEDOPERATION_DOWNLOADFAILED = "FailedOperation.DownloadFailed"
//  FAILEDOPERATION_MD5MISMATCH = "FailedOperation.MD5Mismatch"
//  FAILEDOPERATION_MEDIANOTREADY = "FailedOperation.MediaNotReady"
//  FAILEDOPERATION_TRANSCODEFAILED = "FailedOperation.TranscodeFailed"
//  INVALIDPARAMETER_INVALIDMEDIAID = "InvalidParameter.InvalidMediaId"
//  RESOURCENOTFOUND_MEDIANOTFOUND = "ResourceNotFound.MediaNotFound"
func (c *Client) DescribeMedia(request *DescribeMediaRequest) (response *DescribeMediaResponse, err error) {
    if request == nil {
        request = NewDescribeMediaRequest()
    }
    
    response = NewDescribeMediaResponse()
    err = c.Send(request, response)
    return
}

// DescribeMedia
// 描述媒资文件信息，包括媒资状态，分辨率，帧率等。
//
// 
//
// 如果媒资文件未完成导入，本接口将仅输出媒资文件的状态信息；导入完成后，本接口还将输出媒资文件的其他元信息。
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  FAILEDOPERATION_DOWNLOADFAILED = "FailedOperation.DownloadFailed"
//  FAILEDOPERATION_MD5MISMATCH = "FailedOperation.MD5Mismatch"
//  FAILEDOPERATION_MEDIANOTREADY = "FailedOperation.MediaNotReady"
//  FAILEDOPERATION_TRANSCODEFAILED = "FailedOperation.TranscodeFailed"
//  INVALIDPARAMETER_INVALIDMEDIAID = "InvalidParameter.InvalidMediaId"
//  RESOURCENOTFOUND_MEDIANOTFOUND = "ResourceNotFound.MediaNotFound"
func (c *Client) DescribeMediaWithContext(ctx context.Context, request *DescribeMediaRequest) (response *DescribeMediaResponse, err error) {
    if request == nil {
        request = NewDescribeMediaRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeMediaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMediasRequest() (request *DescribeMediasRequest) {
    request = &DescribeMediasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ivld", APIVersion, "DescribeMedias")
    
    
    return
}

func NewDescribeMediasResponse() (response *DescribeMediasResponse) {
    response = &DescribeMediasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMedias
// 依照输入条件，描述命中的媒资文件信息，包括媒资状态，分辨率，帧率等。
//
// 
//
// 请注意，本接口最多支持同时描述**50**个媒资文件
//
// 
//
// 如果媒资文件未完成导入，本接口将仅输出媒资文件的状态信息；导入完成后，本接口还将输出媒资文件的其他元信息。
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  FAILEDOPERATION_MEDIAALREADYEXIST = "FailedOperation.MediaAlreadyExist"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  INVALIDPARAMETER_INVALIDMD5 = "InvalidParameter.InvalidMD5"
//  INVALIDPARAMETER_INVALIDMEDIAID = "InvalidParameter.InvalidMediaId"
//  INVALIDPARAMETER_INVALIDMEDIANAME = "InvalidParameter.InvalidMediaName"
//  INVALIDPARAMETER_INVALIDMEDIASTATUS = "InvalidParameter.InvalidMediaStatus"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_INVALIDPAGENUMBER = "InvalidParameter.InvalidPageNumber"
//  INVALIDPARAMETER_INVALIDPAGESIZE = "InvalidParameter.InvalidPageSize"
//  INVALIDPARAMETER_INVALIDSORTBY = "InvalidParameter.InvalidSortBy"
//  INVALIDPARAMETER_INVALIDSORTORDER = "InvalidParameter.InvalidSortOrder"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidURL"
//  INVALIDPARAMETER_INVALIDUIN = "InvalidParameter.InvalidUin"
//  INVALIDPARAMETER_NAMETOOLONG = "InvalidParameter.NameTooLong"
//  INVALIDPARAMETER_UNSUPPORTURL = "InvalidParameter.UnsupportURL"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) DescribeMedias(request *DescribeMediasRequest) (response *DescribeMediasResponse, err error) {
    if request == nil {
        request = NewDescribeMediasRequest()
    }
    
    response = NewDescribeMediasResponse()
    err = c.Send(request, response)
    return
}

// DescribeMedias
// 依照输入条件，描述命中的媒资文件信息，包括媒资状态，分辨率，帧率等。
//
// 
//
// 请注意，本接口最多支持同时描述**50**个媒资文件
//
// 
//
// 如果媒资文件未完成导入，本接口将仅输出媒资文件的状态信息；导入完成后，本接口还将输出媒资文件的其他元信息。
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  FAILEDOPERATION_MEDIAALREADYEXIST = "FailedOperation.MediaAlreadyExist"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  INVALIDPARAMETER_INVALIDMD5 = "InvalidParameter.InvalidMD5"
//  INVALIDPARAMETER_INVALIDMEDIAID = "InvalidParameter.InvalidMediaId"
//  INVALIDPARAMETER_INVALIDMEDIANAME = "InvalidParameter.InvalidMediaName"
//  INVALIDPARAMETER_INVALIDMEDIASTATUS = "InvalidParameter.InvalidMediaStatus"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_INVALIDPAGENUMBER = "InvalidParameter.InvalidPageNumber"
//  INVALIDPARAMETER_INVALIDPAGESIZE = "InvalidParameter.InvalidPageSize"
//  INVALIDPARAMETER_INVALIDSORTBY = "InvalidParameter.InvalidSortBy"
//  INVALIDPARAMETER_INVALIDSORTORDER = "InvalidParameter.InvalidSortOrder"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidURL"
//  INVALIDPARAMETER_INVALIDUIN = "InvalidParameter.InvalidUin"
//  INVALIDPARAMETER_NAMETOOLONG = "InvalidParameter.NameTooLong"
//  INVALIDPARAMETER_UNSUPPORTURL = "InvalidParameter.UnsupportURL"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) DescribeMediasWithContext(ctx context.Context, request *DescribeMediasRequest) (response *DescribeMediasResponse, err error) {
    if request == nil {
        request = NewDescribeMediasRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeMediasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskRequest() (request *DescribeTaskRequest) {
    request = &DescribeTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ivld", APIVersion, "DescribeTask")
    
    
    return
}

func NewDescribeTaskResponse() (response *DescribeTaskResponse) {
    response = &DescribeTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTask
// 描述智能标签任务进度。
//
// 
//
// 请注意，**此接口仅返回任务执行状态信息，不返回任务执行结果**
//
// 
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  INVALIDPARAMETER_INVALIDTASKID = "InvalidParameter.InvalidTaskId"
//  RESOURCENOTFOUND_RECORDNOTFOUND = "ResourceNotFound.RecordNotFound"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) DescribeTask(request *DescribeTaskRequest) (response *DescribeTaskResponse, err error) {
    if request == nil {
        request = NewDescribeTaskRequest()
    }
    
    response = NewDescribeTaskResponse()
    err = c.Send(request, response)
    return
}

// DescribeTask
// 描述智能标签任务进度。
//
// 
//
// 请注意，**此接口仅返回任务执行状态信息，不返回任务执行结果**
//
// 
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  INVALIDPARAMETER_INVALIDTASKID = "InvalidParameter.InvalidTaskId"
//  RESOURCENOTFOUND_RECORDNOTFOUND = "ResourceNotFound.RecordNotFound"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) DescribeTaskWithContext(ctx context.Context, request *DescribeTaskRequest) (response *DescribeTaskResponse, err error) {
    if request == nil {
        request = NewDescribeTaskRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskDetailRequest() (request *DescribeTaskDetailRequest) {
    request = &DescribeTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ivld", APIVersion, "DescribeTaskDetail")
    
    
    return
}

func NewDescribeTaskDetailResponse() (response *DescribeTaskDetailResponse) {
    response = &DescribeTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskDetail
// 描述任务信息，如果任务成功完成，还将返回任务结果
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  FAILEDOPERATION_TASKNOTFINISHED = "FailedOperation.TaskNotFinished"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETER_INVALIDTASKID = "InvalidParameter.InvalidTaskId"
//  RESOURCENOTFOUND_RECORDNOTFOUND = "ResourceNotFound.RecordNotFound"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) DescribeTaskDetail(request *DescribeTaskDetailRequest) (response *DescribeTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTaskDetailRequest()
    }
    
    response = NewDescribeTaskDetailResponse()
    err = c.Send(request, response)
    return
}

// DescribeTaskDetail
// 描述任务信息，如果任务成功完成，还将返回任务结果
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  FAILEDOPERATION_TASKNOTFINISHED = "FailedOperation.TaskNotFinished"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETER_INVALIDTASKID = "InvalidParameter.InvalidTaskId"
//  RESOURCENOTFOUND_RECORDNOTFOUND = "ResourceNotFound.RecordNotFound"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) DescribeTaskDetailWithContext(ctx context.Context, request *DescribeTaskDetailRequest) (response *DescribeTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTaskDetailRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeTaskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTasksRequest() (request *DescribeTasksRequest) {
    request = &DescribeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ivld", APIVersion, "DescribeTasks")
    
    
    return
}

func NewDescribeTasksResponse() (response *DescribeTasksResponse) {
    response = &DescribeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTasks
// 依照输入条件，描述命中的任务信息，包括任务创建时间，处理时间信息等。
//
// 
//
// 请注意，本接口最多支持同时描述**50**个任务信息
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  FAILEDOPERATION_GETTASKLISTFAILED = "FailedOperation.GetTaskListFailed"
//  INVALIDPARAMETER_INVALIDPAGENUMBER = "InvalidParameter.InvalidPageNumber"
//  INVALIDPARAMETER_INVALIDPAGESIZE = "InvalidParameter.InvalidPageSize"
//  INVALIDPARAMETER_INVALIDSORTBY = "InvalidParameter.InvalidSortBy"
//  INVALIDPARAMETER_INVALIDTASKID = "InvalidParameter.InvalidTaskId"
//  INVALIDPARAMETER_INVALIDTASKNAME = "InvalidParameter.InvalidTaskName"
//  INVALIDPARAMETER_INVALIDTASKSTATUS = "InvalidParameter.InvalidTaskStatus"
//  INVALIDPARAMETER_INVALIDUIN = "InvalidParameter.InvalidUin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTasksRequest()
    }
    
    response = NewDescribeTasksResponse()
    err = c.Send(request, response)
    return
}

// DescribeTasks
// 依照输入条件，描述命中的任务信息，包括任务创建时间，处理时间信息等。
//
// 
//
// 请注意，本接口最多支持同时描述**50**个任务信息
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  FAILEDOPERATION_GETTASKLISTFAILED = "FailedOperation.GetTaskListFailed"
//  INVALIDPARAMETER_INVALIDPAGENUMBER = "InvalidParameter.InvalidPageNumber"
//  INVALIDPARAMETER_INVALIDPAGESIZE = "InvalidParameter.InvalidPageSize"
//  INVALIDPARAMETER_INVALIDSORTBY = "InvalidParameter.InvalidSortBy"
//  INVALIDPARAMETER_INVALIDTASKID = "InvalidParameter.InvalidTaskId"
//  INVALIDPARAMETER_INVALIDTASKNAME = "InvalidParameter.InvalidTaskName"
//  INVALIDPARAMETER_INVALIDTASKSTATUS = "InvalidParameter.InvalidTaskStatus"
//  INVALIDPARAMETER_INVALIDUIN = "InvalidParameter.InvalidUin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) DescribeTasksWithContext(ctx context.Context, request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTasksRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewImportMediaRequest() (request *ImportMediaRequest) {
    request = &ImportMediaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ivld", APIVersion, "ImportMedia")
    
    
    return
}

func NewImportMediaResponse() (response *ImportMediaResponse) {
    response = &ImportMediaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ImportMedia
// 将URL指向的媒资视频文件导入系统之中。
//
// 
//
// **请注意，本接口为异步接口**。接口返回MediaId仅代表导入视频任务发起，不代表任务完成，您可调用读接口(DescribeMedia/DescribeMedias)接口查询MediaId对应的媒资文件的状态。
//
// 
//
// 当前URL只支持COS地址，其形式为`https://${Bucket}-${AppId}.cos.${Region}.myqcloud.com/${ObjectKey}`，其中`${Bucket}`为您的COS桶名称，Region为COS桶所在[可用区](https://cloud.tencent.com/document/product/213/6091)，`${ObjectKey}`为指向存储在COS桶内的待分析的视频的[ObjectKey](https://cloud.tencent.com/document/product/436/13324)
//
// 
//
// 分析完成后，本产品将在您的`${Bucket}`桶内创建名为`${ObjectKey}-${task-start-time}`的目录(`task-start-time`形式为1970-01-01T08:08:08)并将分析结果将回传回该目录，也即，结构化分析结果(包括图片，JSON等数据)将会写回`https://${Bucket}-${AppId}.cos.${Region}.myqcloud.com/${ObjectKey}-${task-start-time}`目录
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDSECRETID = "AuthFailure.InvalidSecretId"
//  AUTHFAILURE_MFAFAILURE = "AuthFailure.MFAFailure"
//  AUTHFAILURE_SECRETIDNOTFOUND = "AuthFailure.SecretIdNotFound"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TASKFINISHED = "AuthFailure.TaskFinished"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  AUTHFAILURE_USERACTIVATED = "AuthFailure.UserActivated"
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  FAILEDOPERATION_AITEMPLATENOTEXIST = "FailedOperation.AiTemplateNotExist"
//  FAILEDOPERATION_CATEGORYEXIST = "FailedOperation.CategoryExist"
//  FAILEDOPERATION_CATEGORYLEVELCHANGED = "FailedOperation.CategoryLevelChanged"
//  FAILEDOPERATION_CATEGORYREFERRED = "FailedOperation.CategoryReferred"
//  FAILEDOPERATION_CUSTOMGROUPALREADYEXIST = "FailedOperation.CustomGroupAlreadyExist"
//  FAILEDOPERATION_DBCONNECTIONERROR = "FailedOperation.DBConnectionError"
//  FAILEDOPERATION_DOWNLOADFAILED = "FailedOperation.DownloadFailed"
//  FAILEDOPERATION_FEATUREALGOFAILED = "FailedOperation.FeatureAlgoFailed"
//  FAILEDOPERATION_GETCAMTOKENFAILED = "FailedOperation.GetCAMTokenFailed"
//  FAILEDOPERATION_GETTASKLISTFAILED = "FailedOperation.GetTaskListFailed"
//  FAILEDOPERATION_GETVIDEOMETADATAFAILED = "FailedOperation.GetVideoMetadataFailed"
//  FAILEDOPERATION_IMAGENUMEXCEEDED = "FailedOperation.ImageNumExceeded"
//  FAILEDOPERATION_MD5MISMATCH = "FailedOperation.MD5Mismatch"
//  FAILEDOPERATION_MEDIAALREADYEXIST = "FailedOperation.MediaAlreadyExist"
//  FAILEDOPERATION_MEDIAEXPIRED = "FailedOperation.MediaExpired"
//  FAILEDOPERATION_MEDIAINUSE = "FailedOperation.MediaInUse"
//  FAILEDOPERATION_MEDIANOTREADY = "FailedOperation.MediaNotReady"
//  FAILEDOPERATION_MULTIPLEFACESINIMAGE = "FailedOperation.MultipleFacesInImage"
//  FAILEDOPERATION_NOFACEINIMAGE = "FailedOperation.NoFaceInImage"
//  FAILEDOPERATION_OPENCHARGEFAILED = "FailedOperation.OpenChargeFailed"
//  FAILEDOPERATION_PERSONDUPLICATED = "FailedOperation.PersonDuplicated"
//  FAILEDOPERATION_PERSONNOTMATCHED = "FailedOperation.PersonNotMatched"
//  FAILEDOPERATION_PERSONNUMEXCEEDED = "FailedOperation.PersonNumExceeded"
//  FAILEDOPERATION_QUALITYALGOFAILED = "FailedOperation.QualityAlgoFailed"
//  FAILEDOPERATION_QUALITYTOOLOW = "FailedOperation.QualityTooLow"
//  FAILEDOPERATION_SNAPSHOTDESERIALIZEFAILED = "FailedOperation.SnapshotDeserializeFailed"
//  FAILEDOPERATION_STOPFLOWFAILED = "FailedOperation.StopFlowFailed"
//  FAILEDOPERATION_TASKALREADYEXIST = "FailedOperation.TaskAlreadyExist"
//  FAILEDOPERATION_TASKNOTFINISHED = "FailedOperation.TaskNotFinished"
//  FAILEDOPERATION_TRANSCODEFAILED = "FailedOperation.TranscodeFailed"
//  FAILEDOPERATION_UPLOADFAILED = "FailedOperation.UploadFailed"
//  INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  INTERNALERROR_INTERNALOVERFLOW = "InternalError.InternalOverflow"
//  INVALIDPARAMETER_INVALIDCATEGORYID = "InvalidParameter.InvalidCategoryId"
//  INVALIDPARAMETER_INVALIDFILEPATH = "InvalidParameter.InvalidFilePath"
//  INVALIDPARAMETER_INVALIDIMAGE = "InvalidParameter.InvalidImage"
//  INVALIDPARAMETER_INVALIDIMAGEID = "InvalidParameter.InvalidImageId"
//  INVALIDPARAMETER_INVALIDL1CATEGORY = "InvalidParameter.InvalidL1Category"
//  INVALIDPARAMETER_INVALIDL2CATEGORY = "InvalidParameter.InvalidL2Category"
//  INVALIDPARAMETER_INVALIDMD5 = "InvalidParameter.InvalidMD5"
//  INVALIDPARAMETER_INVALIDMEDIAID = "InvalidParameter.InvalidMediaId"
//  INVALIDPARAMETER_INVALIDMEDIALABEL = "InvalidParameter.InvalidMediaLabel"
//  INVALIDPARAMETER_INVALIDMEDIALANG = "InvalidParameter.InvalidMediaLang"
//  INVALIDPARAMETER_INVALIDMEDIANAME = "InvalidParameter.InvalidMediaName"
//  INVALIDPARAMETER_INVALIDMEDIAPREKNOWNINFO = "InvalidParameter.InvalidMediaPreknownInfo"
//  INVALIDPARAMETER_INVALIDMEDIASTATUS = "InvalidParameter.InvalidMediaStatus"
//  INVALIDPARAMETER_INVALIDMEDIATYPE = "InvalidParameter.InvalidMediaType"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_INVALIDPAGENUMBER = "InvalidParameter.InvalidPageNumber"
//  INVALIDPARAMETER_INVALIDPAGESIZE = "InvalidParameter.InvalidPageSize"
//  INVALIDPARAMETER_INVALIDPARAM = "InvalidParameter.InvalidParam"
//  INVALIDPARAMETER_INVALIDPERSONID = "InvalidParameter.InvalidPersonId"
//  INVALIDPARAMETER_INVALIDSORTBY = "InvalidParameter.InvalidSortBy"
//  INVALIDPARAMETER_INVALIDSORTORDER = "InvalidParameter.InvalidSortOrder"
//  INVALIDPARAMETER_INVALIDTASKID = "InvalidParameter.InvalidTaskId"
//  INVALIDPARAMETER_INVALIDTASKNAME = "InvalidParameter.InvalidTaskName"
//  INVALIDPARAMETER_INVALIDTASKSTATUS = "InvalidParameter.InvalidTaskStatus"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidURL"
//  INVALIDPARAMETER_INVALIDUIN = "InvalidParameter.InvalidUin"
//  INVALIDPARAMETER_NAMETOOLONG = "InvalidParameter.NameTooLong"
//  INVALIDPARAMETER_UNSUPPORTURL = "InvalidParameter.UnsupportURL"
//  LIMITEXCEEDED_USAGELIMITEXCEEDED = "LimitExceeded.UsageLimitExceeded"
//  REQUESTLIMITEXCEEDED_BATCHIMPORTOVERFLOW = "RequestLimitExceeded.BatchImportOverflow"
//  REQUESTLIMITEXCEEDED_CONCURRENCYOVERFLOW = "RequestLimitExceeded.ConcurrencyOverflow"
//  RESOURCENOTFOUND_CUSTOMCATEGORYNOTFOUND = "ResourceNotFound.CustomCategoryNotFound"
//  RESOURCENOTFOUND_CUSTOMGROUPNOTFOUND = "ResourceNotFound.CustomGroupNotFound"
//  RESOURCENOTFOUND_MEDIANOTFOUND = "ResourceNotFound.MediaNotFound"
//  RESOURCENOTFOUND_RECORDNOTFOUND = "ResourceNotFound.RecordNotFound"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
//  UNSUPPORTEDOPERATION_MEDIANOTACCESSIBLE = "UnsupportedOperation.MediaNotAccessible"
//  UNSUPPORTEDOPERATION_TASKNOTACCESSIBLE = "UnsupportedOperation.TaskNotAccessible"
func (c *Client) ImportMedia(request *ImportMediaRequest) (response *ImportMediaResponse, err error) {
    if request == nil {
        request = NewImportMediaRequest()
    }
    
    response = NewImportMediaResponse()
    err = c.Send(request, response)
    return
}

// ImportMedia
// 将URL指向的媒资视频文件导入系统之中。
//
// 
//
// **请注意，本接口为异步接口**。接口返回MediaId仅代表导入视频任务发起，不代表任务完成，您可调用读接口(DescribeMedia/DescribeMedias)接口查询MediaId对应的媒资文件的状态。
//
// 
//
// 当前URL只支持COS地址，其形式为`https://${Bucket}-${AppId}.cos.${Region}.myqcloud.com/${ObjectKey}`，其中`${Bucket}`为您的COS桶名称，Region为COS桶所在[可用区](https://cloud.tencent.com/document/product/213/6091)，`${ObjectKey}`为指向存储在COS桶内的待分析的视频的[ObjectKey](https://cloud.tencent.com/document/product/436/13324)
//
// 
//
// 分析完成后，本产品将在您的`${Bucket}`桶内创建名为`${ObjectKey}-${task-start-time}`的目录(`task-start-time`形式为1970-01-01T08:08:08)并将分析结果将回传回该目录，也即，结构化分析结果(包括图片，JSON等数据)将会写回`https://${Bucket}-${AppId}.cos.${Region}.myqcloud.com/${ObjectKey}-${task-start-time}`目录
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDSECRETID = "AuthFailure.InvalidSecretId"
//  AUTHFAILURE_MFAFAILURE = "AuthFailure.MFAFailure"
//  AUTHFAILURE_SECRETIDNOTFOUND = "AuthFailure.SecretIdNotFound"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TASKFINISHED = "AuthFailure.TaskFinished"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  AUTHFAILURE_USERACTIVATED = "AuthFailure.UserActivated"
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  FAILEDOPERATION_AITEMPLATENOTEXIST = "FailedOperation.AiTemplateNotExist"
//  FAILEDOPERATION_CATEGORYEXIST = "FailedOperation.CategoryExist"
//  FAILEDOPERATION_CATEGORYLEVELCHANGED = "FailedOperation.CategoryLevelChanged"
//  FAILEDOPERATION_CATEGORYREFERRED = "FailedOperation.CategoryReferred"
//  FAILEDOPERATION_CUSTOMGROUPALREADYEXIST = "FailedOperation.CustomGroupAlreadyExist"
//  FAILEDOPERATION_DBCONNECTIONERROR = "FailedOperation.DBConnectionError"
//  FAILEDOPERATION_DOWNLOADFAILED = "FailedOperation.DownloadFailed"
//  FAILEDOPERATION_FEATUREALGOFAILED = "FailedOperation.FeatureAlgoFailed"
//  FAILEDOPERATION_GETCAMTOKENFAILED = "FailedOperation.GetCAMTokenFailed"
//  FAILEDOPERATION_GETTASKLISTFAILED = "FailedOperation.GetTaskListFailed"
//  FAILEDOPERATION_GETVIDEOMETADATAFAILED = "FailedOperation.GetVideoMetadataFailed"
//  FAILEDOPERATION_IMAGENUMEXCEEDED = "FailedOperation.ImageNumExceeded"
//  FAILEDOPERATION_MD5MISMATCH = "FailedOperation.MD5Mismatch"
//  FAILEDOPERATION_MEDIAALREADYEXIST = "FailedOperation.MediaAlreadyExist"
//  FAILEDOPERATION_MEDIAEXPIRED = "FailedOperation.MediaExpired"
//  FAILEDOPERATION_MEDIAINUSE = "FailedOperation.MediaInUse"
//  FAILEDOPERATION_MEDIANOTREADY = "FailedOperation.MediaNotReady"
//  FAILEDOPERATION_MULTIPLEFACESINIMAGE = "FailedOperation.MultipleFacesInImage"
//  FAILEDOPERATION_NOFACEINIMAGE = "FailedOperation.NoFaceInImage"
//  FAILEDOPERATION_OPENCHARGEFAILED = "FailedOperation.OpenChargeFailed"
//  FAILEDOPERATION_PERSONDUPLICATED = "FailedOperation.PersonDuplicated"
//  FAILEDOPERATION_PERSONNOTMATCHED = "FailedOperation.PersonNotMatched"
//  FAILEDOPERATION_PERSONNUMEXCEEDED = "FailedOperation.PersonNumExceeded"
//  FAILEDOPERATION_QUALITYALGOFAILED = "FailedOperation.QualityAlgoFailed"
//  FAILEDOPERATION_QUALITYTOOLOW = "FailedOperation.QualityTooLow"
//  FAILEDOPERATION_SNAPSHOTDESERIALIZEFAILED = "FailedOperation.SnapshotDeserializeFailed"
//  FAILEDOPERATION_STOPFLOWFAILED = "FailedOperation.StopFlowFailed"
//  FAILEDOPERATION_TASKALREADYEXIST = "FailedOperation.TaskAlreadyExist"
//  FAILEDOPERATION_TASKNOTFINISHED = "FailedOperation.TaskNotFinished"
//  FAILEDOPERATION_TRANSCODEFAILED = "FailedOperation.TranscodeFailed"
//  FAILEDOPERATION_UPLOADFAILED = "FailedOperation.UploadFailed"
//  INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  INTERNALERROR_INTERNALOVERFLOW = "InternalError.InternalOverflow"
//  INVALIDPARAMETER_INVALIDCATEGORYID = "InvalidParameter.InvalidCategoryId"
//  INVALIDPARAMETER_INVALIDFILEPATH = "InvalidParameter.InvalidFilePath"
//  INVALIDPARAMETER_INVALIDIMAGE = "InvalidParameter.InvalidImage"
//  INVALIDPARAMETER_INVALIDIMAGEID = "InvalidParameter.InvalidImageId"
//  INVALIDPARAMETER_INVALIDL1CATEGORY = "InvalidParameter.InvalidL1Category"
//  INVALIDPARAMETER_INVALIDL2CATEGORY = "InvalidParameter.InvalidL2Category"
//  INVALIDPARAMETER_INVALIDMD5 = "InvalidParameter.InvalidMD5"
//  INVALIDPARAMETER_INVALIDMEDIAID = "InvalidParameter.InvalidMediaId"
//  INVALIDPARAMETER_INVALIDMEDIALABEL = "InvalidParameter.InvalidMediaLabel"
//  INVALIDPARAMETER_INVALIDMEDIALANG = "InvalidParameter.InvalidMediaLang"
//  INVALIDPARAMETER_INVALIDMEDIANAME = "InvalidParameter.InvalidMediaName"
//  INVALIDPARAMETER_INVALIDMEDIAPREKNOWNINFO = "InvalidParameter.InvalidMediaPreknownInfo"
//  INVALIDPARAMETER_INVALIDMEDIASTATUS = "InvalidParameter.InvalidMediaStatus"
//  INVALIDPARAMETER_INVALIDMEDIATYPE = "InvalidParameter.InvalidMediaType"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_INVALIDPAGENUMBER = "InvalidParameter.InvalidPageNumber"
//  INVALIDPARAMETER_INVALIDPAGESIZE = "InvalidParameter.InvalidPageSize"
//  INVALIDPARAMETER_INVALIDPARAM = "InvalidParameter.InvalidParam"
//  INVALIDPARAMETER_INVALIDPERSONID = "InvalidParameter.InvalidPersonId"
//  INVALIDPARAMETER_INVALIDSORTBY = "InvalidParameter.InvalidSortBy"
//  INVALIDPARAMETER_INVALIDSORTORDER = "InvalidParameter.InvalidSortOrder"
//  INVALIDPARAMETER_INVALIDTASKID = "InvalidParameter.InvalidTaskId"
//  INVALIDPARAMETER_INVALIDTASKNAME = "InvalidParameter.InvalidTaskName"
//  INVALIDPARAMETER_INVALIDTASKSTATUS = "InvalidParameter.InvalidTaskStatus"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidURL"
//  INVALIDPARAMETER_INVALIDUIN = "InvalidParameter.InvalidUin"
//  INVALIDPARAMETER_NAMETOOLONG = "InvalidParameter.NameTooLong"
//  INVALIDPARAMETER_UNSUPPORTURL = "InvalidParameter.UnsupportURL"
//  LIMITEXCEEDED_USAGELIMITEXCEEDED = "LimitExceeded.UsageLimitExceeded"
//  REQUESTLIMITEXCEEDED_BATCHIMPORTOVERFLOW = "RequestLimitExceeded.BatchImportOverflow"
//  REQUESTLIMITEXCEEDED_CONCURRENCYOVERFLOW = "RequestLimitExceeded.ConcurrencyOverflow"
//  RESOURCENOTFOUND_CUSTOMCATEGORYNOTFOUND = "ResourceNotFound.CustomCategoryNotFound"
//  RESOURCENOTFOUND_CUSTOMGROUPNOTFOUND = "ResourceNotFound.CustomGroupNotFound"
//  RESOURCENOTFOUND_MEDIANOTFOUND = "ResourceNotFound.MediaNotFound"
//  RESOURCENOTFOUND_RECORDNOTFOUND = "ResourceNotFound.RecordNotFound"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
//  UNSUPPORTEDOPERATION_MEDIANOTACCESSIBLE = "UnsupportedOperation.MediaNotAccessible"
//  UNSUPPORTEDOPERATION_TASKNOTACCESSIBLE = "UnsupportedOperation.TaskNotAccessible"
func (c *Client) ImportMediaWithContext(ctx context.Context, request *ImportMediaRequest) (response *ImportMediaResponse, err error) {
    if request == nil {
        request = NewImportMediaRequest()
    }
    request.SetContext(ctx)
    
    response = NewImportMediaResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCustomCategoryRequest() (request *UpdateCustomCategoryRequest) {
    request = &UpdateCustomCategoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ivld", APIVersion, "UpdateCustomCategory")
    
    
    return
}

func NewUpdateCustomCategoryResponse() (response *UpdateCustomCategoryResponse) {
    response = &UpdateCustomCategoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateCustomCategory
// 更新自定义人物分类
//
// 
//
// 当L2Category为空时，代表更新CategoryId对应的一级自定义人物类型以及所有二级自定义人物类型所从属的一级自定义人物类型；
//
// 当L2Category非空时，仅更新CategoryId对应的二级自定义人物类型
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  FAILEDOPERATION_CATEGORYEXIST = "FailedOperation.CategoryExist"
//  FAILEDOPERATION_CATEGORYLEVELCHANGED = "FailedOperation.CategoryLevelChanged"
//  INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  INTERNALERROR_INTERNALOVERFLOW = "InternalError.InternalOverflow"
//  INVALIDPARAMETER_INVALIDCATEGORYID = "InvalidParameter.InvalidCategoryId"
//  INVALIDPARAMETER_INVALIDL1CATEGORY = "InvalidParameter.InvalidL1Category"
//  INVALIDPARAMETER_INVALIDL2CATEGORY = "InvalidParameter.InvalidL2Category"
//  RESOURCENOTFOUND_CUSTOMCATEGORYNOTFOUND = "ResourceNotFound.CustomCategoryNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) UpdateCustomCategory(request *UpdateCustomCategoryRequest) (response *UpdateCustomCategoryResponse, err error) {
    if request == nil {
        request = NewUpdateCustomCategoryRequest()
    }
    
    response = NewUpdateCustomCategoryResponse()
    err = c.Send(request, response)
    return
}

// UpdateCustomCategory
// 更新自定义人物分类
//
// 
//
// 当L2Category为空时，代表更新CategoryId对应的一级自定义人物类型以及所有二级自定义人物类型所从属的一级自定义人物类型；
//
// 当L2Category非空时，仅更新CategoryId对应的二级自定义人物类型
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  FAILEDOPERATION_CATEGORYEXIST = "FailedOperation.CategoryExist"
//  FAILEDOPERATION_CATEGORYLEVELCHANGED = "FailedOperation.CategoryLevelChanged"
//  INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  INTERNALERROR_INTERNALOVERFLOW = "InternalError.InternalOverflow"
//  INVALIDPARAMETER_INVALIDCATEGORYID = "InvalidParameter.InvalidCategoryId"
//  INVALIDPARAMETER_INVALIDL1CATEGORY = "InvalidParameter.InvalidL1Category"
//  INVALIDPARAMETER_INVALIDL2CATEGORY = "InvalidParameter.InvalidL2Category"
//  RESOURCENOTFOUND_CUSTOMCATEGORYNOTFOUND = "ResourceNotFound.CustomCategoryNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) UpdateCustomCategoryWithContext(ctx context.Context, request *UpdateCustomCategoryRequest) (response *UpdateCustomCategoryResponse, err error) {
    if request == nil {
        request = NewUpdateCustomCategoryRequest()
    }
    request.SetContext(ctx)
    
    response = NewUpdateCustomCategoryResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCustomPersonRequest() (request *UpdateCustomPersonRequest) {
    request = &UpdateCustomPersonRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ivld", APIVersion, "UpdateCustomPerson")
    
    
    return
}

func NewUpdateCustomPersonResponse() (response *UpdateCustomPersonResponse) {
    response = &UpdateCustomPersonResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateCustomPerson
// 更新自定义人物信息，包括姓名，简要信息，分类信息等
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  INVALIDPARAMETER_INVALIDCATEGORYID = "InvalidParameter.InvalidCategoryId"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_INVALIDPARAM = "InvalidParameter.InvalidParam"
//  INVALIDPARAMETER_INVALIDPERSONID = "InvalidParameter.InvalidPersonId"
//  RESOURCENOTFOUND_RECORDNOTFOUND = "ResourceNotFound.RecordNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) UpdateCustomPerson(request *UpdateCustomPersonRequest) (response *UpdateCustomPersonResponse, err error) {
    if request == nil {
        request = NewUpdateCustomPersonRequest()
    }
    
    response = NewUpdateCustomPersonResponse()
    err = c.Send(request, response)
    return
}

// UpdateCustomPerson
// 更新自定义人物信息，包括姓名，简要信息，分类信息等
//
// 可能返回的错误码:
//  AUTHFAILURE_USERINVALIDSTATUS = "AuthFailure.UserInvalidStatus"
//  AUTHFAILURE_USERNOTFOUND = "AuthFailure.UserNotFound"
//  AUTHFAILURE_USERSTOPARREAR = "AuthFailure.UserStopArrear"
//  INTERNALERROR_DBCONNECTIONERROR = "InternalError.DBConnectionError"
//  INTERNALERROR_INNERERROR = "InternalError.InnerError"
//  INVALIDPARAMETER_INVALIDCATEGORYID = "InvalidParameter.InvalidCategoryId"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_INVALIDPARAM = "InvalidParameter.InvalidParam"
//  INVALIDPARAMETER_INVALIDPERSONID = "InvalidParameter.InvalidPersonId"
//  RESOURCENOTFOUND_RECORDNOTFOUND = "ResourceNotFound.RecordNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDPRODUCT = "UnauthorizedOperation.UnauthorizedProduct"
func (c *Client) UpdateCustomPersonWithContext(ctx context.Context, request *UpdateCustomPersonRequest) (response *UpdateCustomPersonResponse, err error) {
    if request == nil {
        request = NewUpdateCustomPersonRequest()
    }
    request.SetContext(ctx)
    
    response = NewUpdateCustomPersonResponse()
    err = c.Send(request, response)
    return
}
