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


func NewBankCard2EVerificationRequest() (request *BankCard2EVerificationRequest) {
    request = &BankCard2EVerificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("faceid", APIVersion, "BankCard2EVerification")
    return
}

func NewBankCard2EVerificationResponse() (response *BankCard2EVerificationResponse) {
    response = &BankCard2EVerificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用于校验姓名和银行卡号的真实性和一致性。
func (c *Client) BankCard2EVerification(request *BankCard2EVerificationRequest) (response *BankCard2EVerificationResponse, err error) {
    if request == nil {
        request = NewBankCard2EVerificationRequest()
    }
    response = NewBankCard2EVerificationResponse()
    err = c.Send(request, response)
    return
}

func NewBankCard4EVerificationRequest() (request *BankCard4EVerificationRequest) {
    request = &BankCard4EVerificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("faceid", APIVersion, "BankCard4EVerification")
    return
}

func NewBankCard4EVerificationResponse() (response *BankCard4EVerificationResponse) {
    response = &BankCard4EVerificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用于输入银行卡号、姓名、开户证件号、开户手机号，校验信息的真实性和一致性。
func (c *Client) BankCard4EVerification(request *BankCard4EVerificationRequest) (response *BankCard4EVerificationResponse, err error) {
    if request == nil {
        request = NewBankCard4EVerificationRequest()
    }
    response = NewBankCard4EVerificationResponse()
    err = c.Send(request, response)
    return
}

func NewBankCardVerificationRequest() (request *BankCardVerificationRequest) {
    request = &BankCardVerificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("faceid", APIVersion, "BankCardVerification")
    return
}

func NewBankCardVerificationResponse() (response *BankCardVerificationResponse) {
    response = &BankCardVerificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用于银行卡号、姓名、开户证件号信息的真实性和一致性。
func (c *Client) BankCardVerification(request *BankCardVerificationRequest) (response *BankCardVerificationResponse, err error) {
    if request == nil {
        request = NewBankCardVerificationRequest()
    }
    response = NewBankCardVerificationResponse()
    err = c.Send(request, response)
    return
}

func NewDetectAuthRequest() (request *DetectAuthRequest) {
    request = &DetectAuthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("faceid", APIVersion, "DetectAuth")
    return
}

func NewDetectAuthResponse() (response *DetectAuthResponse) {
    response = &DetectAuthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 每次调用人脸核身SaaS化服务前，需先调用本接口获取BizToken，用来串联核身流程，在验证完成后，用于获取验证结果信息。
func (c *Client) DetectAuth(request *DetectAuthRequest) (response *DetectAuthResponse, err error) {
    if request == nil {
        request = NewDetectAuthRequest()
    }
    response = NewDetectAuthResponse()
    err = c.Send(request, response)
    return
}

func NewGetActionSequenceRequest() (request *GetActionSequenceRequest) {
    request = &GetActionSequenceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("faceid", APIVersion, "GetActionSequence")
    return
}

func NewGetActionSequenceResponse() (response *GetActionSequenceResponse) {
    response = &GetActionSequenceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 使用动作活体检测模式前，需调用本接口获取动作顺序。
func (c *Client) GetActionSequence(request *GetActionSequenceRequest) (response *GetActionSequenceResponse, err error) {
    if request == nil {
        request = NewGetActionSequenceRequest()
    }
    response = NewGetActionSequenceResponse()
    err = c.Send(request, response)
    return
}

func NewGetDetectInfoRequest() (request *GetDetectInfoRequest) {
    request = &GetDetectInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("faceid", APIVersion, "GetDetectInfo")
    return
}

func NewGetDetectInfoResponse() (response *GetDetectInfoResponse) {
    response = &GetDetectInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 完成验证后，用BizToken调用本接口获取结果信息，BizToken生成后三天内（3\*24\*3,600秒）可多次拉取。
func (c *Client) GetDetectInfo(request *GetDetectInfoRequest) (response *GetDetectInfoResponse, err error) {
    if request == nil {
        request = NewGetDetectInfoRequest()
    }
    response = NewGetDetectInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetDetectInfoEnhancedRequest() (request *GetDetectInfoEnhancedRequest) {
    request = &GetDetectInfoEnhancedRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("faceid", APIVersion, "GetDetectInfoEnhanced")
    return
}

func NewGetDetectInfoEnhancedResponse() (response *GetDetectInfoEnhancedResponse) {
    response = &GetDetectInfoEnhancedResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 完成验证后，用BizToken调用本接口获取结果信息，BizToken生成后三天内（3\*24\*3,600秒）可多次拉取。
func (c *Client) GetDetectInfoEnhanced(request *GetDetectInfoEnhancedRequest) (response *GetDetectInfoEnhancedResponse, err error) {
    if request == nil {
        request = NewGetDetectInfoEnhancedRequest()
    }
    response = NewGetDetectInfoEnhancedResponse()
    err = c.Send(request, response)
    return
}

func NewGetLiveCodeRequest() (request *GetLiveCodeRequest) {
    request = &GetLiveCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("faceid", APIVersion, "GetLiveCode")
    return
}

func NewGetLiveCodeResponse() (response *GetLiveCodeResponse) {
    response = &GetLiveCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 使用数字活体检测模式前，需调用本接口获取数字验证码。
func (c *Client) GetLiveCode(request *GetLiveCodeRequest) (response *GetLiveCodeResponse, err error) {
    if request == nil {
        request = NewGetLiveCodeRequest()
    }
    response = NewGetLiveCodeResponse()
    err = c.Send(request, response)
    return
}

func NewIdCardOCRVerificationRequest() (request *IdCardOCRVerificationRequest) {
    request = &IdCardOCRVerificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("faceid", APIVersion, "IdCardOCRVerification")
    return
}

func NewIdCardOCRVerificationResponse() (response *IdCardOCRVerificationResponse) {
    response = &IdCardOCRVerificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用于校验姓名和身份证号的真实性和一致性，您可以通过输入姓名和身份证号或传入身份证人像面照片提供所需验证信息。
func (c *Client) IdCardOCRVerification(request *IdCardOCRVerificationRequest) (response *IdCardOCRVerificationResponse, err error) {
    if request == nil {
        request = NewIdCardOCRVerificationRequest()
    }
    response = NewIdCardOCRVerificationResponse()
    err = c.Send(request, response)
    return
}

func NewIdCardVerificationRequest() (request *IdCardVerificationRequest) {
    request = &IdCardVerificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("faceid", APIVersion, "IdCardVerification")
    return
}

func NewIdCardVerificationResponse() (response *IdCardVerificationResponse) {
    response = &IdCardVerificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 传入姓名和身份证号，校验两者的真实性和一致性。
func (c *Client) IdCardVerification(request *IdCardVerificationRequest) (response *IdCardVerificationResponse, err error) {
    if request == nil {
        request = NewIdCardVerificationRequest()
    }
    response = NewIdCardVerificationResponse()
    err = c.Send(request, response)
    return
}

func NewImageRecognitionRequest() (request *ImageRecognitionRequest) {
    request = &ImageRecognitionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("faceid", APIVersion, "ImageRecognition")
    return
}

func NewImageRecognitionResponse() (response *ImageRecognitionResponse) {
    response = &ImageRecognitionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 传入照片和身份信息，判断该照片与公安权威库的证件照是否属于同一个人。
func (c *Client) ImageRecognition(request *ImageRecognitionRequest) (response *ImageRecognitionResponse, err error) {
    if request == nil {
        request = NewImageRecognitionRequest()
    }
    response = NewImageRecognitionResponse()
    err = c.Send(request, response)
    return
}

func NewLivenessRequest() (request *LivenessRequest) {
    request = &LivenessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("faceid", APIVersion, "Liveness")
    return
}

func NewLivenessResponse() (response *LivenessResponse) {
    response = &LivenessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 活体检测
func (c *Client) Liveness(request *LivenessRequest) (response *LivenessResponse, err error) {
    if request == nil {
        request = NewLivenessRequest()
    }
    response = NewLivenessResponse()
    err = c.Send(request, response)
    return
}

func NewLivenessCompareRequest() (request *LivenessCompareRequest) {
    request = &LivenessCompareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("faceid", APIVersion, "LivenessCompare")
    return
}

func NewLivenessCompareResponse() (response *LivenessCompareResponse) {
    response = &LivenessCompareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 传入视频和照片，先判断视频中是否为真人，判断为真人后，再判断该视频中的人与上传照片是否属于同一个人。
func (c *Client) LivenessCompare(request *LivenessCompareRequest) (response *LivenessCompareResponse, err error) {
    if request == nil {
        request = NewLivenessCompareRequest()
    }
    response = NewLivenessCompareResponse()
    err = c.Send(request, response)
    return
}

func NewLivenessRecognitionRequest() (request *LivenessRecognitionRequest) {
    request = &LivenessRecognitionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("faceid", APIVersion, "LivenessRecognition")
    return
}

func NewLivenessRecognitionResponse() (response *LivenessRecognitionResponse) {
    response = &LivenessRecognitionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 传入视频和身份信息，先判断视频中是否为真人，判断为真人后，再判断该视频中的人与公安权威库的证件照是否属于同一个人。
func (c *Client) LivenessRecognition(request *LivenessRecognitionRequest) (response *LivenessRecognitionResponse, err error) {
    if request == nil {
        request = NewLivenessRecognitionRequest()
    }
    response = NewLivenessRecognitionResponse()
    err = c.Send(request, response)
    return
}

func NewMinorsVerificationRequest() (request *MinorsVerificationRequest) {
    request = &MinorsVerificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("faceid", APIVersion, "MinorsVerification")
    return
}

func NewMinorsVerificationResponse() (response *MinorsVerificationResponse) {
    response = &MinorsVerificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 未成年人守护接口是通过传入手机号或姓名和身份证号，结合权威数据源和腾讯健康守护可信模型，判断该信息是否真实且年满18周岁。腾讯健康守护可信模型覆盖了上十亿手机库源，覆盖率高、准确率高，如果不在库中的手机号，还可以通过姓名+身份证进行兜底验证。
func (c *Client) MinorsVerification(request *MinorsVerificationRequest) (response *MinorsVerificationResponse, err error) {
    if request == nil {
        request = NewMinorsVerificationRequest()
    }
    response = NewMinorsVerificationResponse()
    err = c.Send(request, response)
    return
}

func NewMobileNetworkTimeVerificationRequest() (request *MobileNetworkTimeVerificationRequest) {
    request = &MobileNetworkTimeVerificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("faceid", APIVersion, "MobileNetworkTimeVerification")
    return
}

func NewMobileNetworkTimeVerificationResponse() (response *MobileNetworkTimeVerificationResponse) {
    response = &MobileNetworkTimeVerificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用于查询手机号在网时长，输入手机号进行查询。
func (c *Client) MobileNetworkTimeVerification(request *MobileNetworkTimeVerificationRequest) (response *MobileNetworkTimeVerificationResponse, err error) {
    if request == nil {
        request = NewMobileNetworkTimeVerificationRequest()
    }
    response = NewMobileNetworkTimeVerificationResponse()
    err = c.Send(request, response)
    return
}

func NewMobileStatusRequest() (request *MobileStatusRequest) {
    request = &MobileStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("faceid", APIVersion, "MobileStatus")
    return
}

func NewMobileStatusResponse() (response *MobileStatusResponse) {
    response = &MobileStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用于验证手机号的状态，您可以输入手机号进行查询。
func (c *Client) MobileStatus(request *MobileStatusRequest) (response *MobileStatusResponse, err error) {
    if request == nil {
        request = NewMobileStatusRequest()
    }
    response = NewMobileStatusResponse()
    err = c.Send(request, response)
    return
}

func NewPhoneVerificationRequest() (request *PhoneVerificationRequest) {
    request = &PhoneVerificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("faceid", APIVersion, "PhoneVerification")
    return
}

func NewPhoneVerificationResponse() (response *PhoneVerificationResponse) {
    response = &PhoneVerificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用于校验手机号、姓名和身份证号的真实性和一致性。
func (c *Client) PhoneVerification(request *PhoneVerificationRequest) (response *PhoneVerificationResponse, err error) {
    if request == nil {
        request = NewPhoneVerificationRequest()
    }
    response = NewPhoneVerificationResponse()
    err = c.Send(request, response)
    return
}
