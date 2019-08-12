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

package v20181119

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-11-19"

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


func NewArithmeticOCRRequest() (request *ArithmeticOCRRequest) {
    request = &ArithmeticOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "ArithmeticOCR")
    return
}

func NewArithmeticOCRResponse() (response *ArithmeticOCRResponse) {
    response = &ArithmeticOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持作业算式题目的自动识别，目前覆盖 K12 学力范围内的 14 种题型，包括加减乘除四则运算、分数四则运算、竖式四则运算、脱式计算等。
func (c *Client) ArithmeticOCR(request *ArithmeticOCRRequest) (response *ArithmeticOCRResponse, err error) {
    if request == nil {
        request = NewArithmeticOCRRequest()
    }
    response = NewArithmeticOCRResponse()
    err = c.Send(request, response)
    return
}

func NewBankCardOCRRequest() (request *BankCardOCRRequest) {
    request = &BankCardOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "BankCardOCR")
    return
}

func NewBankCardOCRResponse() (response *BankCardOCRResponse) {
    response = &BankCardOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持对中国大陆主流银行卡的卡号、银行信息、有效期等关键字段的检测与识别。
func (c *Client) BankCardOCR(request *BankCardOCRRequest) (response *BankCardOCRResponse, err error) {
    if request == nil {
        request = NewBankCardOCRRequest()
    }
    response = NewBankCardOCRResponse()
    err = c.Send(request, response)
    return
}

func NewBizLicenseOCRRequest() (request *BizLicenseOCRRequest) {
    request = &BizLicenseOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "BizLicenseOCR")
    return
}

func NewBizLicenseOCRResponse() (response *BizLicenseOCRResponse) {
    response = &BizLicenseOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持快速精准识别营业执照上的字段，包括注册号、公司名称、经营场所、主体类型、法定代表人、注册资金、组成形式、成立日期、营业期限和经营范围等字段。
func (c *Client) BizLicenseOCR(request *BizLicenseOCRRequest) (response *BizLicenseOCRResponse, err error) {
    if request == nil {
        request = NewBizLicenseOCRRequest()
    }
    response = NewBizLicenseOCRResponse()
    err = c.Send(request, response)
    return
}

func NewBusinessCardOCRRequest() (request *BusinessCardOCRRequest) {
    request = &BusinessCardOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "BusinessCardOCR")
    return
}

func NewBusinessCardOCRResponse() (response *BusinessCardOCRResponse) {
    response = &BusinessCardOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持名片各字段的自动定位与识别，包含姓名、电话、手机号、邮箱、公司、部门、职位、网址、地址、QQ、微信、MSN等。
func (c *Client) BusinessCardOCR(request *BusinessCardOCRRequest) (response *BusinessCardOCRResponse, err error) {
    if request == nil {
        request = NewBusinessCardOCRRequest()
    }
    response = NewBusinessCardOCRResponse()
    err = c.Send(request, response)
    return
}

func NewCarInvoiceOCRRequest() (request *CarInvoiceOCRRequest) {
    request = &CarInvoiceOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "CarInvoiceOCR")
    return
}

func NewCarInvoiceOCRResponse() (response *CarInvoiceOCRResponse) {
    response = &CarInvoiceOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持机动车销售统一发票和二手车销售统一发票的识别，包括发票号码、发票代码、合计金额、合计税额等二十多个字段。
func (c *Client) CarInvoiceOCR(request *CarInvoiceOCRRequest) (response *CarInvoiceOCRResponse, err error) {
    if request == nil {
        request = NewCarInvoiceOCRRequest()
    }
    response = NewCarInvoiceOCRResponse()
    err = c.Send(request, response)
    return
}

func NewDriverLicenseOCRRequest() (request *DriverLicenseOCRRequest) {
    request = &DriverLicenseOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "DriverLicenseOCR")
    return
}

func NewDriverLicenseOCRResponse() (response *DriverLicenseOCRResponse) {
    response = &DriverLicenseOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持对驾驶证主页所有字段的自动定位与识别，包含证号、姓名、性别、国籍、住址、出生日期、初次领证日期、准驾车型、有效期限等。
func (c *Client) DriverLicenseOCR(request *DriverLicenseOCRRequest) (response *DriverLicenseOCRResponse, err error) {
    if request == nil {
        request = NewDriverLicenseOCRRequest()
    }
    response = NewDriverLicenseOCRResponse()
    err = c.Send(request, response)
    return
}

func NewEnglishOCRRequest() (request *EnglishOCRRequest) {
    request = &EnglishOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "EnglishOCR")
    return
}

func NewEnglishOCRResponse() (response *EnglishOCRResponse) {
    response = &EnglishOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持图像英文文字的检测和识别，返回文字框位置与文字内容。支持多场景、任意版面下的英文、字母、数字和常见字符的识别，同时覆盖英文印刷体和英文手写体识别。
func (c *Client) EnglishOCR(request *EnglishOCRRequest) (response *EnglishOCRResponse, err error) {
    if request == nil {
        request = NewEnglishOCRRequest()
    }
    response = NewEnglishOCRResponse()
    err = c.Send(request, response)
    return
}

func NewFlightInvoiceOCRRequest() (request *FlightInvoiceOCRRequest) {
    request = &FlightInvoiceOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "FlightInvoiceOCR")
    return
}

func NewFlightInvoiceOCRResponse() (response *FlightInvoiceOCRResponse) {
    response = &FlightInvoiceOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持机票行程单关键字段的识别，包括姓名、身份证件号码、航班号、票价 、合计、电子客票号码、填开日期等。
func (c *Client) FlightInvoiceOCR(request *FlightInvoiceOCRRequest) (response *FlightInvoiceOCRResponse, err error) {
    if request == nil {
        request = NewFlightInvoiceOCRRequest()
    }
    response = NewFlightInvoiceOCRResponse()
    err = c.Send(request, response)
    return
}

func NewGeneralAccurateOCRRequest() (request *GeneralAccurateOCRRequest) {
    request = &GeneralAccurateOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "GeneralAccurateOCR")
    return
}

func NewGeneralAccurateOCRResponse() (response *GeneralAccurateOCRResponse) {
    response = &GeneralAccurateOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持图像整体文字的检测和识别，返回文字框位置与文字内容。相比通用印刷体识别接口，准确率和召回率更高。
func (c *Client) GeneralAccurateOCR(request *GeneralAccurateOCRRequest) (response *GeneralAccurateOCRResponse, err error) {
    if request == nil {
        request = NewGeneralAccurateOCRRequest()
    }
    response = NewGeneralAccurateOCRResponse()
    err = c.Send(request, response)
    return
}

func NewGeneralBasicOCRRequest() (request *GeneralBasicOCRRequest) {
    request = &GeneralBasicOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "GeneralBasicOCR")
    return
}

func NewGeneralBasicOCRResponse() (response *GeneralBasicOCRResponse) {
    response = &GeneralBasicOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持多场景、任意版面下整图文字的识别。支持自动识别语言类型，同时支持自选语言种类（推荐），除中英文外，支持日语、韩语、西班牙语、法语、德语、葡萄牙语、越南语、马来语、俄语、意大利语、荷兰语、瑞典语、芬兰语、丹麦语、挪威语、匈牙利语、泰语等多种语言。应用场景包括：印刷文档识别、网络图片识别、广告图文字识别、街景店招识别、菜单识别、视频标题识别、头像文字识别等。
func (c *Client) GeneralBasicOCR(request *GeneralBasicOCRRequest) (response *GeneralBasicOCRResponse, err error) {
    if request == nil {
        request = NewGeneralBasicOCRRequest()
    }
    response = NewGeneralBasicOCRResponse()
    err = c.Send(request, response)
    return
}

func NewGeneralFastOCRRequest() (request *GeneralFastOCRRequest) {
    request = &GeneralFastOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "GeneralFastOCR")
    return
}

func NewGeneralFastOCRResponse() (response *GeneralFastOCRResponse) {
    response = &GeneralFastOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持图片中整体文字的检测和识别，返回文字框位置与文字内容。相比通用印刷体识别接口，识别速度更快、支持的 QPS 更高。
func (c *Client) GeneralFastOCR(request *GeneralFastOCRRequest) (response *GeneralFastOCRResponse, err error) {
    if request == nil {
        request = NewGeneralFastOCRRequest()
    }
    response = NewGeneralFastOCRResponse()
    err = c.Send(request, response)
    return
}

func NewGeneralHandwritingOCRRequest() (request *GeneralHandwritingOCRRequest) {
    request = &GeneralHandwritingOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "GeneralHandwritingOCR")
    return
}

func NewGeneralHandwritingOCRResponse() (response *GeneralHandwritingOCRResponse) {
    response = &GeneralHandwritingOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持图片内手写体文字的检测和识别，针对手写字体无规则、字迹潦草、模糊等特点进行了识别能力的增强。
func (c *Client) GeneralHandwritingOCR(request *GeneralHandwritingOCRRequest) (response *GeneralHandwritingOCRResponse, err error) {
    if request == nil {
        request = NewGeneralHandwritingOCRRequest()
    }
    response = NewGeneralHandwritingOCRResponse()
    err = c.Send(request, response)
    return
}

func NewIDCardOCRRequest() (request *IDCardOCRRequest) {
    request = &IDCardOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "IDCardOCR")
    return
}

func NewIDCardOCRResponse() (response *IDCardOCRResponse) {
    response = &IDCardOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持中国大陆居民二代身份证正反面所有字段的识别，包括姓名、性别、民族、出生日期、住址、公民身份证号、签发机关、有效期限；具备身份证照片、人像照片的裁剪功能和翻拍、PS、复印件告警功能，以及边框和框内遮挡告警、临时身份证告警和身份证有效期不合法告警等扩展功能。
func (c *Client) IDCardOCR(request *IDCardOCRRequest) (response *IDCardOCRResponse, err error) {
    if request == nil {
        request = NewIDCardOCRRequest()
    }
    response = NewIDCardOCRResponse()
    err = c.Send(request, response)
    return
}

func NewLicensePlateOCRRequest() (request *LicensePlateOCRRequest) {
    request = &LicensePlateOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "LicensePlateOCR")
    return
}

func NewLicensePlateOCRResponse() (response *LicensePlateOCRResponse) {
    response = &LicensePlateOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持对中国大陆机动车车牌的自动定位和识别，返回地域编号和车牌号信息。
func (c *Client) LicensePlateOCR(request *LicensePlateOCRRequest) (response *LicensePlateOCRResponse, err error) {
    if request == nil {
        request = NewLicensePlateOCRRequest()
    }
    response = NewLicensePlateOCRResponse()
    err = c.Send(request, response)
    return
}

func NewPermitOCRRequest() (request *PermitOCRRequest) {
    request = &PermitOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "PermitOCR")
    return
}

func NewPermitOCRResponse() (response *PermitOCRResponse) {
    response = &PermitOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持对卡式港澳台通行证的识别，包括签发地点、签发机关、有效期限、性别、出生日期、英文姓名、姓名、证件号等字段。
func (c *Client) PermitOCR(request *PermitOCRRequest) (response *PermitOCRResponse, err error) {
    if request == nil {
        request = NewPermitOCRRequest()
    }
    response = NewPermitOCRResponse()
    err = c.Send(request, response)
    return
}

func NewQuotaInvoiceOCRRequest() (request *QuotaInvoiceOCRRequest) {
    request = &QuotaInvoiceOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "QuotaInvoiceOCR")
    return
}

func NewQuotaInvoiceOCRResponse() (response *QuotaInvoiceOCRResponse) {
    response = &QuotaInvoiceOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持定额发票的发票号码、发票代码及金额等关键字段的识别。
func (c *Client) QuotaInvoiceOCR(request *QuotaInvoiceOCRRequest) (response *QuotaInvoiceOCRResponse, err error) {
    if request == nil {
        request = NewQuotaInvoiceOCRRequest()
    }
    response = NewQuotaInvoiceOCRResponse()
    err = c.Send(request, response)
    return
}

func NewTableOCRRequest() (request *TableOCRRequest) {
    request = &TableOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "TableOCR")
    return
}

func NewTableOCRResponse() (response *TableOCRResponse) {
    response = &TableOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持图片内表格文档的检测和识别，返回每个单元格的文字内容，支持将识别结果保存为 Excel 格式。
func (c *Client) TableOCR(request *TableOCRRequest) (response *TableOCRResponse, err error) {
    if request == nil {
        request = NewTableOCRRequest()
    }
    response = NewTableOCRResponse()
    err = c.Send(request, response)
    return
}

func NewTaxiInvoiceOCRRequest() (request *TaxiInvoiceOCRRequest) {
    request = &TaxiInvoiceOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "TaxiInvoiceOCR")
    return
}

func NewTaxiInvoiceOCRResponse() (response *TaxiInvoiceOCRResponse) {
    response = &TaxiInvoiceOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持出租车发票关键字段的识别，包括发票号码、发票代码、金额、日期等字段。
func (c *Client) TaxiInvoiceOCR(request *TaxiInvoiceOCRRequest) (response *TaxiInvoiceOCRResponse, err error) {
    if request == nil {
        request = NewTaxiInvoiceOCRRequest()
    }
    response = NewTaxiInvoiceOCRResponse()
    err = c.Send(request, response)
    return
}

func NewTrainTicketOCRRequest() (request *TrainTicketOCRRequest) {
    request = &TrainTicketOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "TrainTicketOCR")
    return
}

func NewTrainTicketOCRResponse() (response *TrainTicketOCRResponse) {
    response = &TrainTicketOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持火车票全字段的识别，包括编号、票价、姓名、座位号、出发时间、出发站、到达站、车次、席别等。
func (c *Client) TrainTicketOCR(request *TrainTicketOCRRequest) (response *TrainTicketOCRResponse, err error) {
    if request == nil {
        request = NewTrainTicketOCRRequest()
    }
    response = NewTrainTicketOCRResponse()
    err = c.Send(request, response)
    return
}

func NewVatInvoiceOCRRequest() (request *VatInvoiceOCRRequest) {
    request = &VatInvoiceOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "VatInvoiceOCR")
    return
}

func NewVatInvoiceOCRResponse() (response *VatInvoiceOCRResponse) {
    response = &VatInvoiceOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持增值税专用发票、增值税普通发票、增值税电子发票全字段的内容检测和识别，包括发票代码、发票号码、开票日期、合计金额、校验码、税率、合计税额、价税合计、购买方识别号、复核、销售方识别号、开票人、密码区1、密码区2、密码区3、密码区4、发票名称、购买方名称、销售方名称、服务名称、备注、规格型号、数量、单价、金额、税额、收款人等字段。
func (c *Client) VatInvoiceOCR(request *VatInvoiceOCRRequest) (response *VatInvoiceOCRResponse, err error) {
    if request == nil {
        request = NewVatInvoiceOCRRequest()
    }
    response = NewVatInvoiceOCRResponse()
    err = c.Send(request, response)
    return
}

func NewVehicleLicenseOCRRequest() (request *VehicleLicenseOCRRequest) {
    request = &VehicleLicenseOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "VehicleLicenseOCR")
    return
}

func NewVehicleLicenseOCRResponse() (response *VehicleLicenseOCRResponse) {
    response = &VehicleLicenseOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持行驶证主页和副页所有字段的自动定位与识别，包含车牌号码、车辆类型、所有人、住址、使用性质、品牌型号、车辆识别代码、发动机号、注册日期、发证日期等。
func (c *Client) VehicleLicenseOCR(request *VehicleLicenseOCRRequest) (response *VehicleLicenseOCRResponse, err error) {
    if request == nil {
        request = NewVehicleLicenseOCRRequest()
    }
    response = NewVehicleLicenseOCRResponse()
    err = c.Send(request, response)
    return
}

func NewVinOCRRequest() (request *VinOCRRequest) {
    request = &VinOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "VinOCR")
    return
}

func NewVinOCRResponse() (response *VinOCRResponse) {
    response = &VinOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持图片内车辆识别代号（VIN）的检测和识别。
func (c *Client) VinOCR(request *VinOCRRequest) (response *VinOCRResponse, err error) {
    if request == nil {
        request = NewVinOCRRequest()
    }
    response = NewVinOCRResponse()
    err = c.Send(request, response)
    return
}

func NewWaybillOCRRequest() (request *WaybillOCRRequest) {
    request = &WaybillOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "WaybillOCR")
    return
}

func NewWaybillOCRResponse() (response *WaybillOCRResponse) {
    response = &WaybillOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持市面上主流版式电子运单的识别，包括收件人和寄件人的姓名、电话、地址以及运单号等字段。
func (c *Client) WaybillOCR(request *WaybillOCRRequest) (response *WaybillOCRResponse, err error) {
    if request == nil {
        request = NewWaybillOCRRequest()
    }
    response = NewWaybillOCRResponse()
    err = c.Send(request, response)
    return
}
