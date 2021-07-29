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

package v20180301

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-03-01"

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


func NewAnalyzeDenseLandmarksRequest() (request *AnalyzeDenseLandmarksRequest) {
    request = &AnalyzeDenseLandmarksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "AnalyzeDenseLandmarks")
    return
}

func NewAnalyzeDenseLandmarksResponse() (response *AnalyzeDenseLandmarksResponse) {
    response = &AnalyzeDenseLandmarksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AnalyzeDenseLandmarks
// 对请求图片进行五官定位（也称人脸关键点定位），获得人脸的精准信息，返回多达888点关键信息，对五官和脸部轮廓进行精确定位。
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) AnalyzeDenseLandmarks(request *AnalyzeDenseLandmarksRequest) (response *AnalyzeDenseLandmarksResponse, err error) {
    if request == nil {
        request = NewAnalyzeDenseLandmarksRequest()
    }
    response = NewAnalyzeDenseLandmarksResponse()
    err = c.Send(request, response)
    return
}

func NewAnalyzeFaceRequest() (request *AnalyzeFaceRequest) {
    request = &AnalyzeFaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "AnalyzeFace")
    return
}

func NewAnalyzeFaceResponse() (response *AnalyzeFaceResponse) {
    response = &AnalyzeFaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AnalyzeFace
// 对请求图片进行五官定位（也称人脸关键点定位），计算构成人脸轮廓的 90 个点，包括眉毛（左右各 8 点）、眼睛（左右各 8 点）、鼻子（13 点）、嘴巴（22 点）、脸型轮廓（21 点）、眼珠[或瞳孔]（2点）。
//
// 
//
// >     
//
// - 公共参数中的签名方式请使用V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_FACEMODELVERSIONILLEGAL = "InvalidParameterValue.FaceModelVersionIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) AnalyzeFace(request *AnalyzeFaceRequest) (response *AnalyzeFaceResponse, err error) {
    if request == nil {
        request = NewAnalyzeFaceRequest()
    }
    response = NewAnalyzeFaceResponse()
    err = c.Send(request, response)
    return
}

func NewCheckSimilarPersonRequest() (request *CheckSimilarPersonRequest) {
    request = &CheckSimilarPersonRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "CheckSimilarPerson")
    return
}

func NewCheckSimilarPersonResponse() (response *CheckSimilarPersonResponse) {
    response = &CheckSimilarPersonResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckSimilarPerson
// 对指定的人员库进行人员查重，给出疑似相同人的信息。
//
// 
//
// 可以使用本接口对已有的单个人员库进行人员查重，避免同一人在单个人员库中拥有多个身份；也可以使用本接口对已有的多个人员库进行人员查重，查询同一人是否同时存在多个人员库中。
//
// 
//
// 不支持跨算法模型版本查重，且目前仅支持算法模型为3.0的人员库使用查重功能。
//
// 
//
// >     
//
// - 若对完全相同的指定人员库进行查重操作，需等待上次操作完成才可。即，若两次请求输入的 GroupIds 相同，第一次请求若未完成，第二次请求将返回失败。
//
// 
//
// >     
//
// - 查重的人员库状态为腾讯云开始进行查重任务的那一刻，即您可以理解为当您发起查重请求后，若您的查重任务需要排队，在排队期间您对人员库的增删操作均会会影响查重的结果。腾讯云将以开始进行查重任务的那一刻人员库的状态进行查重。查重任务开始后，您对人员库的任何操作均不影响查重任务的进行。但建议查重任务开始后，请不要对人员库中人员和人脸进行增删操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACROSSVERSIONSERROR = "FailedOperation.AcrossVersionsError"
//  FAILEDOPERATION_CHECKDUPLICATEPERSONTASKNOTFINISHED = "FailedOperation.CheckDuplicatePersonTaskNotFinished"
//  FAILEDOPERATION_CHECKSIMILARPERSONTIMEEXCEED = "FailedOperation.CheckSimilarPersonTimeExceed"
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_UNIQUEPERSONCONTROLILLEGAL = "InvalidParameterValue.UniquePersonControlIllegal"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDGROUPFACEMODELVERSION = "InvalidParameterValue.UnsupportedGroupFaceModelVersion"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_GETAUTHINFOERROR = "ResourceUnavailable.GetAuthInfoError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) CheckSimilarPerson(request *CheckSimilarPersonRequest) (response *CheckSimilarPersonResponse, err error) {
    if request == nil {
        request = NewCheckSimilarPersonRequest()
    }
    response = NewCheckSimilarPersonResponse()
    err = c.Send(request, response)
    return
}

func NewCompareFaceRequest() (request *CompareFaceRequest) {
    request = &CompareFaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "CompareFace")
    return
}

func NewCompareFaceResponse() (response *CompareFaceResponse) {
    response = &CompareFaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CompareFace
// 对两张图片中的人脸进行相似度比对，返回人脸相似度分数。
//
// 
//
// 若您需要判断 “此人是否是某人”，即验证某张图片中的人是否是已知身份的某人，如常见的人脸登录场景，建议使用[人脸验证](https://cloud.tencent.com/document/product/867/32806)或[人员验证](https://cloud.tencent.com/document/product/867/38879)接口。
//
// 
//
// >     
//
// - 公共参数中的签名方式请使用V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACEQUALITYNOTQUALIFIED = "FailedOperation.FaceQualityNotQualified"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_FACEMODELVERSIONILLEGAL = "InvalidParameterValue.FaceModelVersionIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_QUALITYCONTROLILLEGAL = "InvalidParameterValue.QualityControlIllegal"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) CompareFace(request *CompareFaceRequest) (response *CompareFaceResponse, err error) {
    if request == nil {
        request = NewCompareFaceRequest()
    }
    response = NewCompareFaceResponse()
    err = c.Send(request, response)
    return
}

func NewCopyPersonRequest() (request *CopyPersonRequest) {
    request = &CopyPersonRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "CopyPerson")
    return
}

func NewCopyPersonResponse() (response *CopyPersonResponse) {
    response = &CopyPersonResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CopyPerson
// 将已存在于某人员库的人员复制到其他人员库，该人员的描述信息不会被复制。单个人员最多只能同时存在100个人员库中。
//
// >     
//
// - 注：若该人员创建时算法模型版本为2.0，复制到非2.0算法模型版本的Group中时，复制操作将会失败。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACROSSVERSIONSERROR = "FailedOperation.AcrossVersionsError"
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_CREATEFACECONCURRENT = "FailedOperation.CreateFaceConcurrent"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) CopyPerson(request *CopyPersonRequest) (response *CopyPersonResponse, err error) {
    if request == nil {
        request = NewCopyPersonRequest()
    }
    response = NewCopyPersonResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFaceRequest() (request *CreateFaceRequest) {
    request = &CreateFaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "CreateFace")
    return
}

func NewCreateFaceResponse() (response *CreateFaceResponse) {
    response = &CreateFaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFace
// 将一组人脸图片添加到一个人员中。一个人员最多允许包含 5 张图片。若该人员存在多个人员库中，所有人员库中该人员图片均会增加。
//
// 
//
// >     
//
// - 公共参数中的签名方式请使用V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_CREATEFACECONCURRENT = "FailedOperation.CreateFaceConcurrent"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_FACEMATCHTHRESHOLDILLEGAL = "InvalidParameterValue.FaceMatchThresholdIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_QUALITYCONTROLILLEGAL = "InvalidParameterValue.QualityControlIllegal"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) CreateFace(request *CreateFaceRequest) (response *CreateFaceResponse, err error) {
    if request == nil {
        request = NewCreateFaceRequest()
    }
    response = NewCreateFaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGroupRequest() (request *CreateGroupRequest) {
    request = &CreateGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "CreateGroup")
    return
}

func NewCreateGroupResponse() (response *CreateGroupResponse) {
    response = &CreateGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateGroup
// 用于创建一个空的人员库，如果人员库已存在返回错误。
//
// 可根据需要创建自定义描述字段，用于辅助描述该人员库下的人员信息。
//
// 
//
// 1个APPID下最多创建10万个人员库（Group）、最多包含5000万张人脸（Face）。
//
// 
//
// 不同算法模型版本（FaceModelVersion）的人员库（Group）最多可包含人脸（Face）数不同。算法模型版本为2.0的人员库最多包含100万张人脸，算法模型版本为3.0的人员库最多可包含300万张人脸。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_FACEMODELVERSIONILLEGAL = "InvalidParameterValue.FaceModelVersionIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) CreateGroup(request *CreateGroupRequest) (response *CreateGroupResponse, err error) {
    if request == nil {
        request = NewCreateGroupRequest()
    }
    response = NewCreateGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePersonRequest() (request *CreatePersonRequest) {
    request = &CreatePersonRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "CreatePerson")
    return
}

func NewCreatePersonResponse() (response *CreatePersonResponse) {
    response = &CreatePersonResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePerson
// 创建人员，添加人脸、姓名、性别及其他相关信息。
//
// 
//
// >     
//
// - 公共参数中的签名方式请使用V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_CREATEFACECONCURRENT = "FailedOperation.CreateFaceConcurrent"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACEQUALITYNOTQUALIFIED = "FailedOperation.FaceQualityNotQualified"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_QUALITYCONTROLILLEGAL = "InvalidParameterValue.QualityControlIllegal"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UNIQUEPERSONCONTROLILLEGAL = "InvalidParameterValue.UniquePersonControlIllegal"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) CreatePerson(request *CreatePersonRequest) (response *CreatePersonResponse, err error) {
    if request == nil {
        request = NewCreatePersonRequest()
    }
    response = NewCreatePersonResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFaceRequest() (request *DeleteFaceRequest) {
    request = &DeleteFaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "DeleteFace")
    return
}

func NewDeleteFaceResponse() (response *DeleteFaceResponse) {
    response = &DeleteFaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteFace
// 删除一个人员下的人脸图片。如果该人员只有一张人脸图片，则返回错误。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_CREATEFACECONCURRENT = "FailedOperation.CreateFaceConcurrent"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) DeleteFace(request *DeleteFaceRequest) (response *DeleteFaceResponse, err error) {
    if request == nil {
        request = NewDeleteFaceRequest()
    }
    response = NewDeleteFaceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGroupRequest() (request *DeleteGroupRequest) {
    request = &DeleteGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "DeleteGroup")
    return
}

func NewDeleteGroupResponse() (response *DeleteGroupResponse) {
    response = &DeleteGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteGroup
// 删除该人员库及包含的所有的人员。同时，人员对应的所有人脸信息将被删除。若某人员同时存在多个人员库中，该人员不会被删除，但属于该人员库中的自定义描述字段信息会被删除，属于其他人员库的自定义描述字段信息不受影响。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) DeleteGroup(request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    if request == nil {
        request = NewDeleteGroupRequest()
    }
    response = NewDeleteGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePersonRequest() (request *DeletePersonRequest) {
    request = &DeletePersonRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "DeletePerson")
    return
}

func NewDeletePersonResponse() (response *DeletePersonResponse) {
    response = &DeletePersonResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePerson
// 删除该人员信息，此操作会导致所有人员库均删除此人员。同时，该人员的所有人脸信息将被删除。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_CREATEFACECONCURRENT = "FailedOperation.CreateFaceConcurrent"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) DeletePerson(request *DeletePersonRequest) (response *DeletePersonResponse, err error) {
    if request == nil {
        request = NewDeletePersonRequest()
    }
    response = NewDeletePersonResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePersonFromGroupRequest() (request *DeletePersonFromGroupRequest) {
    request = &DeletePersonFromGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "DeletePersonFromGroup")
    return
}

func NewDeletePersonFromGroupResponse() (response *DeletePersonFromGroupResponse) {
    response = &DeletePersonFromGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePersonFromGroup
// 从某人员库中删除人员，此操作仅影响该人员库。若该人员仅存在于指定的人员库中，该人员将被删除，其所有的人脸信息也将被删除。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_CREATEFACECONCURRENT = "FailedOperation.CreateFaceConcurrent"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) DeletePersonFromGroup(request *DeletePersonFromGroupRequest) (response *DeletePersonFromGroupResponse, err error) {
    if request == nil {
        request = NewDeletePersonFromGroupRequest()
    }
    response = NewDeletePersonFromGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDetectFaceRequest() (request *DetectFaceRequest) {
    request = &DetectFaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "DetectFace")
    return
}

func NewDetectFaceResponse() (response *DetectFaceResponse) {
    response = &DetectFaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetectFace
// 检测给定图片中的人脸（Face）的位置、相应的面部属性和人脸质量信息，位置包括 (x，y，w，h)，面部属性包括性别（gender）、年龄（age）、表情（expression）、魅力（beauty）、眼镜（glass）、发型（hair）、口罩（mask）和姿态 (pitch，roll，yaw)，人脸质量信息包括整体质量分（score）、模糊分（sharpness）、光照分（brightness）和五官遮挡分（completeness）。
//
// 
//
//  
//
// 其中，人脸质量信息主要用于评价输入的人脸图片的质量。在使用人脸识别服务时，建议您对输入的人脸图片进行质量检测，提升后续业务处理的效果。该功能的应用场景包括：
//
// 
//
// 1） 人员库[创建人员](https://cloud.tencent.com/document/product/867/32793)/[增加人脸](https://cloud.tencent.com/document/product/867/32795)：保证人员人脸信息的质量，便于后续的业务处理。
//
// 
//
// 2） [人脸搜索](https://cloud.tencent.com/document/product/867/32798)：保证输入的图片质量，快速准确匹配到对应的人员。
//
// 
//
// 3） [人脸验证](https://cloud.tencent.com/document/product/867/32806)：保证人脸信息的质量，避免明明是本人却认证不通过的情况。
//
// 
//
// 4） [人脸融合](https://cloud.tencent.com/product/facefusion)：保证上传的人脸质量，人脸融合的效果更好。
//
// 
//
// >     
//
// - 公共参数中的签名方式请使用V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_FACEMODELVERSIONILLEGAL = "InvalidParameterValue.FaceModelVersionIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) DetectFace(request *DetectFaceRequest) (response *DetectFaceResponse, err error) {
    if request == nil {
        request = NewDetectFaceRequest()
    }
    response = NewDetectFaceResponse()
    err = c.Send(request, response)
    return
}

func NewDetectFaceAttributesRequest() (request *DetectFaceAttributesRequest) {
    request = &DetectFaceAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "DetectFaceAttributes")
    return
}

func NewDetectFaceAttributesResponse() (response *DetectFaceAttributesResponse) {
    response = &DetectFaceAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetectFaceAttributes
// 检测给定图片中的人脸（Face）的位置、相应的面部属性和人脸质量信息，位置包括 (x，y，w，h)，面部属性包括性别（gender）、年龄（age）、表情（expression）、魅力（beauty）、眼镜（glass）、发型（hair）、口罩（mask）和姿态 (pitch，roll，yaw)，人脸质量信息包括整体质量分（score）、模糊分（sharpness）、光照分（brightness）和五官遮挡分（completeness）。
//
// 
//
//  
//
// 其中，人脸质量信息主要用于评价输入的人脸图片的质量。在使用人脸识别服务时，建议您对输入的人脸图片进行质量检测，提升后续业务处理的效果。该功能的应用场景包括：
//
// 
//
// 1） 人员库[创建人员](https://cloud.tencent.com/document/product/867/32793)/[增加人脸](https://cloud.tencent.com/document/product/867/32795)：保证人员人脸信息的质量，便于后续的业务处理。
//
// 
//
// 2） [人脸搜索](https://cloud.tencent.com/document/product/867/32798)：保证输入的图片质量，快速准确匹配到对应的人员。
//
// 
//
// 3） [人脸验证](https://cloud.tencent.com/document/product/867/32806)：保证人脸信息的质量，避免明明是本人却认证不通过的情况。
//
// 
//
// 4） [人脸融合](https://cloud.tencent.com/product/facefusion)：保证上传的人脸质量，人脸融合的效果更好。
//
// 
//
// >     
//
// - 本接口是[人脸检测与分析](https://cloud.tencent.com/document/product/867/32800)的升级，具体在于：
//
// 
//
// 1.本接口可以指定需要计算返回的人脸属性，避免无效计算，降低耗时；
//
// 
//
// 2.本接口支持更多属性细项数，也会持续增加更多功能。
//
// 
//
// 请您使用本接口完成相应的人脸检测与属性分析需求。
//
// 
//
// - 公共参数中的签名方式请使用V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_FACEMODELVERSIONILLEGAL = "InvalidParameterValue.FaceModelVersionIllegal"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) DetectFaceAttributes(request *DetectFaceAttributesRequest) (response *DetectFaceAttributesResponse, err error) {
    if request == nil {
        request = NewDetectFaceAttributesRequest()
    }
    response = NewDetectFaceAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewDetectLiveFaceRequest() (request *DetectLiveFaceRequest) {
    request = &DetectLiveFaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "DetectLiveFace")
    return
}

func NewDetectLiveFaceResponse() (response *DetectLiveFaceResponse) {
    response = &DetectLiveFaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetectLiveFace
// 用于对用户上传的静态图片进行人脸活体检测。与动态活体检测的区别是：静态活体检测中，用户不需要通过唇语或摇头眨眼等动作来识别。
//
// 
//
// 静态活体检测适用于手机自拍的场景，或对防攻击要求不高的场景。如果对活体检测有更高安全性要求，请使用[人脸核身·云智慧眼](https://cloud.tencent.com/product/faceid)产品。
//
// 
//
// >     
//
// - 图片的宽高比请接近3：4，不符合宽高比的图片返回的分值不具备参考意义。本接口适用于类手机自拍场景，非类手机自拍照返回的分值不具备参考意义。
//
// 
//
// >
//
// - 使用过程中建议正对摄像头，不要距离太远，使面部可以完整地显示在识别的框内，识别过程中不要移动设备或遮挡面部。不要选择光线过强或过弱的环境进行面部识别，识别时不要添加任何滤镜。
//
// 
//
// >     
//
// - 公共参数中的签名方式请使用V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_FACEMODELVERSIONILLEGAL = "InvalidParameterValue.FaceModelVersionIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) DetectLiveFace(request *DetectLiveFaceRequest) (response *DetectLiveFaceResponse, err error) {
    if request == nil {
        request = NewDetectLiveFaceRequest()
    }
    response = NewDetectLiveFaceResponse()
    err = c.Send(request, response)
    return
}

func NewEstimateCheckSimilarPersonCostTimeRequest() (request *EstimateCheckSimilarPersonCostTimeRequest) {
    request = &EstimateCheckSimilarPersonCostTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "EstimateCheckSimilarPersonCostTime")
    return
}

func NewEstimateCheckSimilarPersonCostTimeResponse() (response *EstimateCheckSimilarPersonCostTimeResponse) {
    response = &EstimateCheckSimilarPersonCostTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EstimateCheckSimilarPersonCostTime
// 获取若要开始一个人员查重任务，这个任务结束的预估时间。
//
// 
//
// 若EndTimestamp符合您预期，请您尽快发起人员查重请求，否则导致可能需要更多处理时间。
//
// 
//
// 若预估时间超过5小时，则无法使用人员查重功能。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACROSSVERSIONSERROR = "FailedOperation.AcrossVersionsError"
//  FAILEDOPERATION_CHECKSIMILARPERSONTIMEEXCEED = "FailedOperation.CheckSimilarPersonTimeExceed"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDGROUPFACEMODELVERSION = "InvalidParameterValue.UnsupportedGroupFaceModelVersion"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_GETAUTHINFOERROR = "ResourceUnavailable.GetAuthInfoError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) EstimateCheckSimilarPersonCostTime(request *EstimateCheckSimilarPersonCostTimeRequest) (response *EstimateCheckSimilarPersonCostTimeResponse, err error) {
    if request == nil {
        request = NewEstimateCheckSimilarPersonCostTimeRequest()
    }
    response = NewEstimateCheckSimilarPersonCostTimeResponse()
    err = c.Send(request, response)
    return
}

func NewGetCheckSimilarPersonJobIdListRequest() (request *GetCheckSimilarPersonJobIdListRequest) {
    request = &GetCheckSimilarPersonJobIdListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "GetCheckSimilarPersonJobIdList")
    return
}

func NewGetCheckSimilarPersonJobIdListResponse() (response *GetCheckSimilarPersonJobIdListResponse) {
    response = &GetCheckSimilarPersonJobIdListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetCheckSimilarPersonJobIdList
// 获取人员查重任务列表，按任务创建时间逆序（最新的在前面）。
//
// 
//
// 只保留最近1年的数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACROSSVERSIONSERROR = "FailedOperation.AcrossVersionsError"
//  FAILEDOPERATION_JOBIDNOTEXIST = "FailedOperation.JobIdNotExist"
//  FAILEDOPERATION_JOBNOTEXIST = "FailedOperation.JobNotExist"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_GETAUTHINFOERROR = "ResourceUnavailable.GetAuthInfoError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) GetCheckSimilarPersonJobIdList(request *GetCheckSimilarPersonJobIdListRequest) (response *GetCheckSimilarPersonJobIdListResponse, err error) {
    if request == nil {
        request = NewGetCheckSimilarPersonJobIdListRequest()
    }
    response = NewGetCheckSimilarPersonJobIdListResponse()
    err = c.Send(request, response)
    return
}

func NewGetGroupInfoRequest() (request *GetGroupInfoRequest) {
    request = &GetGroupInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "GetGroupInfo")
    return
}

func NewGetGroupInfoResponse() (response *GetGroupInfoResponse) {
    response = &GetGroupInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetGroupInfo
// 获取人员库信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GetGroupInfo(request *GetGroupInfoRequest) (response *GetGroupInfoResponse, err error) {
    if request == nil {
        request = NewGetGroupInfoRequest()
    }
    response = NewGetGroupInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetGroupListRequest() (request *GetGroupListRequest) {
    request = &GetGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "GetGroupList")
    return
}

func NewGetGroupListResponse() (response *GetGroupListResponse) {
    response = &GetGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetGroupList
// 获取人员库列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) GetGroupList(request *GetGroupListRequest) (response *GetGroupListResponse, err error) {
    if request == nil {
        request = NewGetGroupListRequest()
    }
    response = NewGetGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewGetPersonBaseInfoRequest() (request *GetPersonBaseInfoRequest) {
    request = &GetPersonBaseInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "GetPersonBaseInfo")
    return
}

func NewGetPersonBaseInfoResponse() (response *GetPersonBaseInfoResponse) {
    response = &GetPersonBaseInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetPersonBaseInfo
// 获取指定人员的信息，包括姓名、性别、人脸等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) GetPersonBaseInfo(request *GetPersonBaseInfoRequest) (response *GetPersonBaseInfoResponse, err error) {
    if request == nil {
        request = NewGetPersonBaseInfoRequest()
    }
    response = NewGetPersonBaseInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetPersonGroupInfoRequest() (request *GetPersonGroupInfoRequest) {
    request = &GetPersonGroupInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "GetPersonGroupInfo")
    return
}

func NewGetPersonGroupInfoResponse() (response *GetPersonGroupInfoResponse) {
    response = &GetPersonGroupInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetPersonGroupInfo
// 获取指定人员的信息，包括加入的人员库、描述内容等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) GetPersonGroupInfo(request *GetPersonGroupInfoRequest) (response *GetPersonGroupInfoResponse, err error) {
    if request == nil {
        request = NewGetPersonGroupInfoRequest()
    }
    response = NewGetPersonGroupInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetPersonListRequest() (request *GetPersonListRequest) {
    request = &GetPersonListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "GetPersonList")
    return
}

func NewGetPersonListResponse() (response *GetPersonListResponse) {
    response = &GetPersonListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetPersonList
// 获取指定人员库中的人员列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) GetPersonList(request *GetPersonListRequest) (response *GetPersonListResponse, err error) {
    if request == nil {
        request = NewGetPersonListRequest()
    }
    response = NewGetPersonListResponse()
    err = c.Send(request, response)
    return
}

func NewGetPersonListNumRequest() (request *GetPersonListNumRequest) {
    request = &GetPersonListNumRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "GetPersonListNum")
    return
}

func NewGetPersonListNumResponse() (response *GetPersonListNumResponse) {
    response = &GetPersonListNumResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetPersonListNum
// 获取指定人员库中人员数量。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) GetPersonListNum(request *GetPersonListNumRequest) (response *GetPersonListNumResponse, err error) {
    if request == nil {
        request = NewGetPersonListNumRequest()
    }
    response = NewGetPersonListNumResponse()
    err = c.Send(request, response)
    return
}

func NewGetSimilarPersonResultRequest() (request *GetSimilarPersonResultRequest) {
    request = &GetSimilarPersonResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "GetSimilarPersonResult")
    return
}

func NewGetSimilarPersonResultResponse() (response *GetSimilarPersonResultResponse) {
    response = &GetSimilarPersonResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetSimilarPersonResult
// 获取人员查重接口（CheckSimilarPerson）结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_JOBIDNOTEXIST = "FailedOperation.JobIdNotExist"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_GETAUTHINFOERROR = "ResourceUnavailable.GetAuthInfoError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) GetSimilarPersonResult(request *GetSimilarPersonResultRequest) (response *GetSimilarPersonResultResponse, err error) {
    if request == nil {
        request = NewGetSimilarPersonResultRequest()
    }
    response = NewGetSimilarPersonResultResponse()
    err = c.Send(request, response)
    return
}

func NewGetUpgradeGroupFaceModelVersionJobListRequest() (request *GetUpgradeGroupFaceModelVersionJobListRequest) {
    request = &GetUpgradeGroupFaceModelVersionJobListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "GetUpgradeGroupFaceModelVersionJobList")
    return
}

func NewGetUpgradeGroupFaceModelVersionJobListResponse() (response *GetUpgradeGroupFaceModelVersionJobListResponse) {
    response = &GetUpgradeGroupFaceModelVersionJobListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetUpgradeGroupFaceModelVersionJobList
// 获取人员库升级任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GetUpgradeGroupFaceModelVersionJobList(request *GetUpgradeGroupFaceModelVersionJobListRequest) (response *GetUpgradeGroupFaceModelVersionJobListResponse, err error) {
    if request == nil {
        request = NewGetUpgradeGroupFaceModelVersionJobListRequest()
    }
    response = NewGetUpgradeGroupFaceModelVersionJobListResponse()
    err = c.Send(request, response)
    return
}

func NewGetUpgradeGroupFaceModelVersionResultRequest() (request *GetUpgradeGroupFaceModelVersionResultRequest) {
    request = &GetUpgradeGroupFaceModelVersionResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "GetUpgradeGroupFaceModelVersionResult")
    return
}

func NewGetUpgradeGroupFaceModelVersionResultResponse() (response *GetUpgradeGroupFaceModelVersionResultResponse) {
    response = &GetUpgradeGroupFaceModelVersionResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetUpgradeGroupFaceModelVersionResult
// 人员库升级结果查询
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UPGRADEJOBIDNOTEXIST = "FailedOperation.UpgradeJobIdNotExist"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GetUpgradeGroupFaceModelVersionResult(request *GetUpgradeGroupFaceModelVersionResultRequest) (response *GetUpgradeGroupFaceModelVersionResultResponse, err error) {
    if request == nil {
        request = NewGetUpgradeGroupFaceModelVersionResultRequest()
    }
    response = NewGetUpgradeGroupFaceModelVersionResultResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGroupRequest() (request *ModifyGroupRequest) {
    request = &ModifyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "ModifyGroup")
    return
}

func NewModifyGroupResponse() (response *ModifyGroupResponse) {
    response = &ModifyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyGroup
// 修改人员库名称、备注、自定义描述字段名称。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) ModifyGroup(request *ModifyGroupRequest) (response *ModifyGroupResponse, err error) {
    if request == nil {
        request = NewModifyGroupRequest()
    }
    response = NewModifyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPersonBaseInfoRequest() (request *ModifyPersonBaseInfoRequest) {
    request = &ModifyPersonBaseInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "ModifyPersonBaseInfo")
    return
}

func NewModifyPersonBaseInfoResponse() (response *ModifyPersonBaseInfoResponse) {
    response = &ModifyPersonBaseInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPersonBaseInfo
// 修改人员信息，包括名称、性别等。人员名称和性别修改会同步到包含该人员的所有人员库。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) ModifyPersonBaseInfo(request *ModifyPersonBaseInfoRequest) (response *ModifyPersonBaseInfoResponse, err error) {
    if request == nil {
        request = NewModifyPersonBaseInfoRequest()
    }
    response = NewModifyPersonBaseInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPersonGroupInfoRequest() (request *ModifyPersonGroupInfoRequest) {
    request = &ModifyPersonGroupInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "ModifyPersonGroupInfo")
    return
}

func NewModifyPersonGroupInfoResponse() (response *ModifyPersonGroupInfoResponse) {
    response = &ModifyPersonGroupInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPersonGroupInfo
// 修改指定人员库人员描述内容。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) ModifyPersonGroupInfo(request *ModifyPersonGroupInfoRequest) (response *ModifyPersonGroupInfoResponse, err error) {
    if request == nil {
        request = NewModifyPersonGroupInfoRequest()
    }
    response = NewModifyPersonGroupInfoResponse()
    err = c.Send(request, response)
    return
}

func NewRevertGroupFaceModelVersionRequest() (request *RevertGroupFaceModelVersionRequest) {
    request = &RevertGroupFaceModelVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "RevertGroupFaceModelVersion")
    return
}

func NewRevertGroupFaceModelVersionResponse() (response *RevertGroupFaceModelVersionResponse) {
    response = &RevertGroupFaceModelVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RevertGroupFaceModelVersion
// 本接口用于回滚人员库的人脸识别算法模型版本。单个人员库有且仅有一次回滚机会。
//
// 
//
// 回滚操作会在10s内生效，回滚操作中，您对人员库的操作可能会失效。
//
// 可能返回的错误码:
//  FAILEDOPERATION_JOBCANNNOTROLLBACK = "FailedOperation.JobCannnotRollback"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UPGRADEJOBIDNOTEXIST = "FailedOperation.UpgradeJobIdNotExist"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RevertGroupFaceModelVersion(request *RevertGroupFaceModelVersionRequest) (response *RevertGroupFaceModelVersionResponse, err error) {
    if request == nil {
        request = NewRevertGroupFaceModelVersionRequest()
    }
    response = NewRevertGroupFaceModelVersionResponse()
    err = c.Send(request, response)
    return
}

func NewSearchFacesRequest() (request *SearchFacesRequest) {
    request = &SearchFacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "SearchFaces")
    return
}

func NewSearchFacesResponse() (response *SearchFacesResponse) {
    response = &SearchFacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SearchFaces
// 用于对一张待识别的人脸图片，在一个或多个人员库中识别出最相似的 TopK 人员，识别结果按照相似度从大到小排序。
//
// 
//
// 支持一次性识别图片中的最多 10 张人脸，支持一次性跨 100 个人员库（Group）搜索。
//
// 
//
// 单次搜索的人员库人脸总数量和人员库的算法模型版本（FaceModelVersion）相关。算法模型版本为2.0的人员库，单次搜索人员库人脸总数量不得超过 100 万张；算法模型版本为3.0的人员库，单次搜索人员库人脸总数量不得超过 300 万张。
//
// 
//
// 与[人员搜索](https://cloud.tencent.com/document/product/867/38881)及[人员搜索按库返回](https://cloud.tencent.com/document/product/867/38880)接口不同的是，本接口将该人员（Person）下的每个人脸（Face）都作为单独个体进行验证，而人员搜索及人员搜索按库返回接口 会将该人员（Person）下的所有人脸（Face）进行融合特征处理，即若某个Person下有4张 Face，本接口会将4张 Face 的特征进行融合处理，生成对应这个 Person 的特征，使搜索更加准确。
//
// 
//
// 
//
// 本接口需与[人员库管理相关接口](https://cloud.tencent.com/document/product/867/32794)结合使用。
//
// 
//
// >     
//
// - 公共参数中的签名方式请使用V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACROSSVERSIONSERROR = "FailedOperation.AcrossVersionsError"
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_FACEMATCHTHRESHOLDILLEGAL = "InvalidParameterValue.FaceMatchThresholdIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_QUALITYCONTROLILLEGAL = "InvalidParameterValue.QualityControlIllegal"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) SearchFaces(request *SearchFacesRequest) (response *SearchFacesResponse, err error) {
    if request == nil {
        request = NewSearchFacesRequest()
    }
    response = NewSearchFacesResponse()
    err = c.Send(request, response)
    return
}

func NewSearchFacesReturnsByGroupRequest() (request *SearchFacesReturnsByGroupRequest) {
    request = &SearchFacesReturnsByGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "SearchFacesReturnsByGroup")
    return
}

func NewSearchFacesReturnsByGroupResponse() (response *SearchFacesReturnsByGroupResponse) {
    response = &SearchFacesReturnsByGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SearchFacesReturnsByGroup
// 用于对一张待识别的人脸图片，在一个或多个人员库中识别出最相似的 TopK 人员，按照**人员库的维度**以人员相似度从大到小顺序排列。
//
// 
//
// 支持一次性识别图片中的最多 10 张人脸，支持跨人员库（Group）搜索。
//
// 
//
// 单次搜索的人员库人脸总数量和人员库的算法模型版本（FaceModelVersion）相关。算法模型版本为2.0的人员库，单次搜索人员库人脸总数量不得超过 100 万张；算法模型版本为3.0的人员库，单次搜索人员库人脸总数量不得超过 300 万张。
//
// 
//
// 与[人员搜索](https://cloud.tencent.com/document/product/867/38881)及[人员搜索按库返回](https://cloud.tencent.com/document/product/867/38880)接口不同的是，本接口将该人员（Person）下的每个人脸（Face）都作为单独个体进行验证，而[人员搜索](https://cloud.tencent.com/document/product/867/38881)及[人员搜索按库返回](https://cloud.tencent.com/document/product/867/38880)接口 会将该人员（Person）下的所有人脸（Face）进行融合特征处理，即若某个Person下有4张 Face，本接口会将4张 Face 的特征进行融合处理，生成对应这个 Person 的特征，使搜索更加准确。
//
// 
//
// 本接口需与[人员库管理相关接口](https://cloud.tencent.com/document/product/867/32794)结合使用。
//
// 
//
// >     
//
// - 公共参数中的签名方式请使用V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACROSSVERSIONSERROR = "FailedOperation.AcrossVersionsError"
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_FACEMATCHTHRESHOLDILLEGAL = "InvalidParameterValue.FaceMatchThresholdIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_QUALITYCONTROLILLEGAL = "InvalidParameterValue.QualityControlIllegal"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) SearchFacesReturnsByGroup(request *SearchFacesReturnsByGroupRequest) (response *SearchFacesReturnsByGroupResponse, err error) {
    if request == nil {
        request = NewSearchFacesReturnsByGroupRequest()
    }
    response = NewSearchFacesReturnsByGroupResponse()
    err = c.Send(request, response)
    return
}

func NewSearchPersonsRequest() (request *SearchPersonsRequest) {
    request = &SearchPersonsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "SearchPersons")
    return
}

func NewSearchPersonsResponse() (response *SearchPersonsResponse) {
    response = &SearchPersonsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SearchPersons
// 用于对一张待识别的人脸图片，在一个或多个人员库中识别出最相似的 TopK 人员，按照相似度从大到小排列。
//
// 
//
// 支持一次性识别图片中的最多 10 张人脸，支持一次性跨 100 个人员库（Group）搜索。
//
// 
//
// 单次搜索的人员库人脸总数量和人员库的算法模型版本（FaceModelVersion）相关。算法模型版本为2.0的人员库，单次搜索人员库人脸总数量不得超过 100 万张；算法模型版本为3.0的人员库，单次搜索人员库人脸总数量不得超过 300 万张。
//
// 
//
// 本接口会将该人员（Person）下的所有人脸（Face）进行融合特征处理，即若某个 Person 下有4张 Face ，本接口会将4张 Face 的特征进行融合处理，生成对应这个 Person 的特征，使人员搜索（确定待识别的人脸图片是某人）更加准确。而[人脸搜索](https://cloud.tencent.com/document/product/867/32798)及[人脸搜索按库返回接口](https://cloud.tencent.com/document/product/867/38882)将该人员（Person）下的每个人脸（Face）都作为单独个体进行搜索。
//
// 
//
// >     
//
// - 公共参数中的签名方式请使用V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// - 仅支持算法模型版本（FaceModelVersion）为3.0的人员库。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACROSSVERSIONSERROR = "FailedOperation.AcrossVersionsError"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_FACEMATCHTHRESHOLDILLEGAL = "InvalidParameterValue.FaceMatchThresholdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_IMAGEEMPTYERROR = "InvalidParameterValue.ImageEmptyError"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_QUALITYCONTROLILLEGAL = "InvalidParameterValue.QualityControlIllegal"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDGROUPFACEMODELVERSION = "InvalidParameterValue.UnsupportedGroupFaceModelVersion"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_CHARGESTATUSEXCEPTION = "ResourceUnavailable.ChargeStatusException"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_GETAUTHINFOERROR = "ResourceUnavailable.GetAuthInfoError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) SearchPersons(request *SearchPersonsRequest) (response *SearchPersonsResponse, err error) {
    if request == nil {
        request = NewSearchPersonsRequest()
    }
    response = NewSearchPersonsResponse()
    err = c.Send(request, response)
    return
}

func NewSearchPersonsReturnsByGroupRequest() (request *SearchPersonsReturnsByGroupRequest) {
    request = &SearchPersonsReturnsByGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "SearchPersonsReturnsByGroup")
    return
}

func NewSearchPersonsReturnsByGroupResponse() (response *SearchPersonsReturnsByGroupResponse) {
    response = &SearchPersonsReturnsByGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SearchPersonsReturnsByGroup
// 用于对一张待识别的人脸图片，在一个或多个人员库中识别出最相似的 TopK 人员，按照**人员库的维度**以人员相似度从大到小顺序排列。
//
// 
//
// 支持一次性识别图片中的最多 10 张人脸，支持跨人员库（Group）搜索。
//
// 
//
// 单次搜索的人员库人脸总数量和人员库的算法模型版本（FaceModelVersion）相关。算法模型版本为2.0的人员库，单次搜索人员库人脸总数量不得超过 100 万张；算法模型版本为3.0的人员库，单次搜索人员库人脸总数量不得超过 300 万张。
//
// 
//
// 本接口会将该人员（Person）下的所有人脸（Face）进行融合特征处理，即若某个 Person 下有4张 Face ，本接口会将4张 Face 的特征进行融合处理，生成对应这个 Person 的特征，使人员搜索（确定待识别的人脸图片是某人）更加准确。而[人脸搜索](https://cloud.tencent.com/document/product/867/32798)及[人脸搜索按库返回接口](https://cloud.tencent.com/document/product/867/38882)将该人员（Person）下的每个人脸（Face）都作为单独个体进行搜索。
//
// >     
//
// - 公共参数中的签名方式请使用V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// - 仅支持算法模型版本（FaceModelVersion）为3.0的人员库。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACROSSVERSIONSERROR = "FailedOperation.AcrossVersionsError"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_FACEMATCHTHRESHOLDILLEGAL = "InvalidParameterValue.FaceMatchThresholdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXISTS = "InvalidParameterValue.GroupIdNotExists"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_QUALITYCONTROLILLEGAL = "InvalidParameterValue.QualityControlIllegal"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDGROUPFACEMODELVERSION = "InvalidParameterValue.UnsupportedGroupFaceModelVersion"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_CHARGESTATUSEXCEPTION = "ResourceUnavailable.ChargeStatusException"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) SearchPersonsReturnsByGroup(request *SearchPersonsReturnsByGroupRequest) (response *SearchPersonsReturnsByGroupResponse, err error) {
    if request == nil {
        request = NewSearchPersonsReturnsByGroupRequest()
    }
    response = NewSearchPersonsReturnsByGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeGroupFaceModelVersionRequest() (request *UpgradeGroupFaceModelVersionRequest) {
    request = &UpgradeGroupFaceModelVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "UpgradeGroupFaceModelVersion")
    return
}

func NewUpgradeGroupFaceModelVersionResponse() (response *UpgradeGroupFaceModelVersionResponse) {
    response = &UpgradeGroupFaceModelVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpgradeGroupFaceModelVersion
// 升级人员库。升级过程中，人员库仍然为原算法版本，人员库相关操作仍然支持。升级完成后，人员库为新算法版本。
//
// 单个人员库有且仅支持一次回滚操作。
//
// 注：此处QPS限制为10。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPCANNOTUPGRADE = "FailedOperation.GroupCannotUpgrade"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INVALIDPARAMETERVALUE_FACEMODELVERSIONILLEGAL = "InvalidParameterValue.FaceModelVersionIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) UpgradeGroupFaceModelVersion(request *UpgradeGroupFaceModelVersionRequest) (response *UpgradeGroupFaceModelVersionResponse, err error) {
    if request == nil {
        request = NewUpgradeGroupFaceModelVersionRequest()
    }
    response = NewUpgradeGroupFaceModelVersionResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyFaceRequest() (request *VerifyFaceRequest) {
    request = &VerifyFaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "VerifyFace")
    return
}

func NewVerifyFaceResponse() (response *VerifyFaceResponse) {
    response = &VerifyFaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// VerifyFace
// 给定一张人脸图片和一个 PersonId，判断图片中的人和 PersonId 对应的人是否为同一人。PersonId 请参考[人员库管理相关接口](https://cloud.tencent.com/document/product/867/32794)。 
//
// 
//
// 与[人脸比对](https://cloud.tencent.com/document/product/867/32802)接口不同的是，人脸验证用于判断 “此人是否是此人”，“此人”的信息已存于人员库中，“此人”可能存在多张人脸图片；而[人脸比对](https://cloud.tencent.com/document/product/867/32802)用于判断两张人脸的相似度。
//
// 
//
// 与[人员验证](https://cloud.tencent.com/document/product/867/38879)接口不同的是，人脸验证将该人员（Person）下的每个人脸（Face）都作为单独个体进行验证，而[人员验证](https://cloud.tencent.com/document/product/867/38879)会将该人员（Person）下的所有人脸（Face）进行融合特征处理，即若某个 Person下有4张 Face，人员验证接口会将4张 Face 的特征进行融合处理，生成对应这个 Person 的特征，使人员验证（确定待识别的人脸图片是某人员）更加准确。
//
// 
//
// >     
//
// - 公共参数中的签名方式请使用V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACEQUALITYNOTQUALIFIED = "FailedOperation.FaceQualityNotQualified"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_QUALITYCONTROLILLEGAL = "InvalidParameterValue.QualityControlIllegal"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) VerifyFace(request *VerifyFaceRequest) (response *VerifyFaceResponse, err error) {
    if request == nil {
        request = NewVerifyFaceRequest()
    }
    response = NewVerifyFaceResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyPersonRequest() (request *VerifyPersonRequest) {
    request = &VerifyPersonRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "VerifyPerson")
    return
}

func NewVerifyPersonResponse() (response *VerifyPersonResponse) {
    response = &VerifyPersonResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// VerifyPerson
// 给定一张人脸图片和一个 PersonId，判断图片中的人和 PersonId 对应的人是否为同一人。PersonId 请参考[人员库管理相关接口](https://cloud.tencent.com/document/product/867/32794)。
//
// 本接口会将该人员（Person）下的所有人脸（Face）进行融合特征处理，即若某个Person下有4张 Face，本接口会将4张 Face 的特征进行融合处理，生成对应这个 Person 的特征，使人员验证（确定待识别的人脸图片是某人员）更加准确。
//
// 
//
//  和人脸比对相关接口不同的是，人脸验证相关接口用于判断 “此人是否是此人”，“此人”的信息已存于人员库中，“此人”可能存在多张人脸图片；而人脸比对相关接口用于判断两张人脸的相似度。
//
// 
//
// 
//
// >     
//
// - 公共参数中的签名方式请使用V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// - 仅支持算法模型版本（FaceModelVersion）为3.0的人员库。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FACEQUALITYNOTQUALIFIED = "FailedOperation.FaceQualityNotQualified"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_QUALITYCONTROLILLEGAL = "InvalidParameterValue.QualityControlIllegal"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDGROUPFACEMODELVERSION = "InvalidParameterValue.UnsupportedGroupFaceModelVersion"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  RESOURCEUNAVAILABLE_CHARGESTATUSEXCEPTION = "ResourceUnavailable.ChargeStatusException"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_GETAUTHINFOERROR = "ResourceUnavailable.GetAuthInfoError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) VerifyPerson(request *VerifyPersonRequest) (response *VerifyPersonResponse, err error) {
    if request == nil {
        request = NewVerifyPersonRequest()
    }
    response = NewVerifyPersonResponse()
    err = c.Send(request, response)
    return
}
