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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
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

// 对请求图片进行五官定位（也称人脸关键点定位），计算构成人脸轮廓的 90 个点，包括眉毛（左右各 8 点）、眼睛（左右各 8 点）、鼻子（13 点）、嘴巴（22 点）、脸型轮廓（21 点）、眼珠[或瞳孔]（2点）。
func (c *Client) AnalyzeFace(request *AnalyzeFaceRequest) (response *AnalyzeFaceResponse, err error) {
    if request == nil {
        request = NewAnalyzeFaceRequest()
    }
    response = NewAnalyzeFaceResponse()
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

// 对两张图片中的人脸进行相似度比对，返回人脸相似度分数。
// 
// 若您需要判断 “此人是否是某人”，即验证某张图片中的人是否是已知身份的某人，如常见的人脸登录场景，建议使用[人脸验证](https://cloud.tencent.com/document/product/867/32806)接口。
// 
// 若您需要判断图片中人脸的具体身份信息，如是否是身份证上对应的人，建议使用[人脸核身·云智慧眼](https://cloud.tencent.com/product/facein)产品。
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

// 将已存在于某人员库的人员复制到其他人员库，该人员的描述信息不会被复制。单个人员最多只能同时存在100个人员库中。
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

// 将一组人脸图片添加到一个人员中。一个人员最多允许包含 5 张图片。若该人员存在多个人员库中，所有人员库中该人员图片均会增加。
// >
// - 增加人脸完成后，生效时间一般不超过 1 秒，极端情况最多不超过 5 秒，之后您可以进行[人脸搜索](https://cloud.tencent.com/document/product/867/32798)或[人脸验证](https://cloud.tencent.com/document/product/867/32806)。
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

// 用于创建一个空的人员库，如果人员库已存在返回错误。可根据需要创建自定义描述字段，用于辅助描述该人员库下的人员信息。1个APPID下最多创建2万个人员库（Group）、最多包含1000万张人脸（Face），单个人员库（Group）最多包含100万张人脸（Face）。
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

// 创建人员，添加人脸、姓名、性别及其他相关信息。
// >
// - 创建人员完成后，生效时间一般不超过 1 秒，极端情况最多不超过 5 秒，之后您可以进行[人脸搜索](https://cloud.tencent.com/document/product/867/32798)或[人脸验证](https://cloud.tencent.com/document/product/867/32806)。
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

// 删除一个人员下的人脸图片。如果该人员只有一张人脸图片，则返回错误。
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

// 删除该人员库及包含的所有的人员。同时，人员对应的所有人脸信息将被删除。若某人员同时存在多个人员库中，该人员不会被删除，但属于该人员库中的自定义描述字段信息会被删除。
// 
// 注：删除人员库的操作为异步执行，删除单张人脸时间约为10ms，即一小时内可以删除36万张。删除期间，无法向该人员库添加人员。
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

// 删除该人员信息，此操作会导致所有人员库均删除此人员。同时，该人员的所有人脸信息将被删除。
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

// 从某人员库中删除人员，此操作仅影响该人员库。若该人员仅存在于指定的人员库中，该人员将被删除，其所有的人脸信息也将被删除。
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

// 检测给定图片中的人脸（Face）的位置、相应的面部属性和人脸质量信息，位置包括 (x，y，w，h)，面部属性包括性别（gender）、年龄（age）、表情（expression）、魅力（beauty）、眼镜（glass）、发型（hair）、口罩（mask）和姿态 (pitch，roll，yaw)，人脸质量信息包括整体质量分（score）、模糊分（sharpness）、光照分（brightness）和五官遮挡分（completeness）。
// 
//  
// 其中，人脸质量信息主要用于评价输入的人脸图片的质量。在使用人脸识别服务时，建议您对输入的人脸图片进行质量检测，提升后续业务处理的效果。该功能的应用场景包括：
// 
// 1） 人员库[创建人员](https://cloud.tencent.com/document/product/867/32793)/[增加人脸](https://cloud.tencent.com/document/product/867/32795)：保证人员人脸信息的质量，便于后续的业务处理。
// 
// 2） [人脸搜索](https://cloud.tencent.com/document/product/867/32798)：保证输入的图片质量，快速准确匹配到对应的人员。
// 
// 3） [人脸验证](https://cloud.tencent.com/document/product/867/32806)：保证人脸信息的质量，避免明明是本人却认证不通过的情况。
// 
// 4） [人脸融合](https://cloud.tencent.com/product/facefusion)：保证上传的人脸质量，人脸融合的效果更好。
// 
func (c *Client) DetectFace(request *DetectFaceRequest) (response *DetectFaceResponse, err error) {
    if request == nil {
        request = NewDetectFaceRequest()
    }
    response = NewDetectFaceResponse()
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

// 用于对用户上传的静态图片进行人脸活体检测。与动态活体检测的区别是：静态活体检测中，用户不需要通过唇语或摇头眨眼等动作来识别。
// 
// 静态活体检测适用于手机自拍的场景，或对防攻击要求不高的场景。如果对活体检测有更高安全性要求，请使用[人脸核身·云智慧眼](https://cloud.tencent.com/product/faceid)产品。
// 
// >     
// - 图片的宽高比请接近3：4，不符合宽高比的图片返回的分值不具备参考意义。本接口适用于类手机自拍场景，非类手机自拍照返回的分值不具备参考意义。
func (c *Client) DetectLiveFace(request *DetectLiveFaceRequest) (response *DetectLiveFaceResponse, err error) {
    if request == nil {
        request = NewDetectLiveFaceRequest()
    }
    response = NewDetectLiveFaceResponse()
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

// 获取人员库列表。
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

// 获取指定人员的信息，包括姓名、性别、人脸等。
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

// 获取指定人员的信息，包括加入的人员库、描述内容等。
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

// 获取指定人员库中的人员列表。
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

// 获取指定人员库中人员数量。
func (c *Client) GetPersonListNum(request *GetPersonListNumRequest) (response *GetPersonListNumResponse, err error) {
    if request == nil {
        request = NewGetPersonListNumRequest()
    }
    response = NewGetPersonListNumResponse()
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

// 修改人员库名称、备注、自定义描述字段名称。
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

// 修改人员信息，包括名称、性别等。人员名称和性别修改会同步到包含该人员的所有人员库。
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

// 修改指定人员库人员描述内容。
func (c *Client) ModifyPersonGroupInfo(request *ModifyPersonGroupInfoRequest) (response *ModifyPersonGroupInfoResponse, err error) {
    if request == nil {
        request = NewModifyPersonGroupInfoRequest()
    }
    response = NewModifyPersonGroupInfoResponse()
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

// 用于对一张待识别的人脸图片，在一个或多个人员库中识别出最相似的 TopN 人员，识别结果按照相似度从大到小排序。单次搜索的人员库人脸总数量不得超过 100 万张。
// 此接口需与[人员库管理相关接口](https://cloud.tencent.com/document/product/867/32794)结合使用。
func (c *Client) SearchFaces(request *SearchFacesRequest) (response *SearchFacesResponse, err error) {
    if request == nil {
        request = NewSearchFacesRequest()
    }
    response = NewSearchFacesResponse()
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

// 给定一张人脸图片和一个 PersonId，判断图片中的人和 PersonId 对应的人是否为同一人。PersonId 请参考[人员库管理相关接口](https://cloud.tencent.com/document/product/867/32794)。 和[人脸比对](https://cloud.tencent.com/document/product/867/32802)接口不同的是，[人脸验证](https://cloud.tencent.com/document/product/867/32806)用于判断 “此人是否是此人”，“此人”的信息已存于人员库中，“此人”可能存在多张人脸图片；而[人脸比对](https://cloud.tencent.com/document/product/867/32802)用于判断两张人脸的相似度。
func (c *Client) VerifyFace(request *VerifyFaceRequest) (response *VerifyFaceResponse, err error) {
    if request == nil {
        request = NewVerifyFaceRequest()
    }
    response = NewVerifyFaceResponse()
    err = c.Send(request, response)
    return
}
