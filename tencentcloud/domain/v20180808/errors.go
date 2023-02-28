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

package v20180808

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 域名查询失败，请稍后重试该功能。
	FAILEDOPERATION_CHECKDOMAINFAILED = "FailedOperation.CheckDomainFailed"

	// 创建模板操作失败。
	FAILEDOPERATION_CREATETEMPLATEFAILED = "FailedOperation.CreateTemplateFailed"

	// 删除模版操作失败，请稍后重试。
	FAILEDOPERATION_DELETETEMPLATEFAILED = "FailedOperation.DeleteTemplateFailed"

	// 获取域名信息操作失败，请稍后重试。
	FAILEDOPERATION_DESCRIBEDOMAINFAILED = "FailedOperation.DescribeDomainFailed"

	// 获取域名信息操作失败，请稍后重试。
	FAILEDOPERATION_DESCRIBEDOMAINLISTFAILED = "FailedOperation.DescribeDomainListFailed"

	// 查询模板操作失败。
	FAILEDOPERATION_DESCRIBETEMPLATEFAILED = "FailedOperation.DescribeTemplateFailed"

	// 域名已过期，不允许操作。
	FAILEDOPERATION_DOMAINEXPIREDUNSUPPORTED = "FailedOperation.DomainExpiredUnsupported"

	// 获取域名价格列表失败。
	FAILEDOPERATION_DOMAINPRICELISTFAILED = "FailedOperation.DomainPriceListFailed"

	// 当前账号下已有相同的手机/邮箱，无需重复添加。
	FAILEDOPERATION_DUPLICATEPHONEEMAIL = "FailedOperation.DuplicatePhoneEmail"

	// 获取域名价格失败。
	FAILEDOPERATION_GETDOMAINPRICEFAILED = "FailedOperation.GetDomainPriceFailed"

	// 域名过户失败。
	FAILEDOPERATION_MODIFYDOMAINOWNERFAILED = "FailedOperation.ModifyDomainOwnerFailed"

	// 禁止的手机或邮箱。
	FAILEDOPERATION_PROHIBITPHONEEMAIL = "FailedOperation.ProhibitPhoneEmail"

	// 域名注册操作失败，请稍后重试。
	FAILEDOPERATION_REGISTERDOMAIN = "FailedOperation.RegisterDomain"

	// 域名注册操作失败，请稍后重试。
	FAILEDOPERATION_REGISTERDOMAINFAILED = "FailedOperation.RegisterDomainFailed"

	// 当前账号为云开发（TCB）账号，无法使用验证功能，请切换登录小程序公众号后重新操作。
	FAILEDOPERATION_SENDTCBPHONEEMAILCODEFAILED = "FailedOperation.SendTcbPhoneEmailCodeFailed"

	// 发送验证码过于频繁，请稍后重试。
	FAILEDOPERATION_SENDVERIFYCODEISLIMITED = "FailedOperation.SendVerifyCodeIsLimited"

	// 修改 DNS 失败，请输入正确的 DNS 服务器地址。
	FAILEDOPERATION_SETDOMAINDNSFAILED = "FailedOperation.SetDomainDnsFailed"

	// 信息模板超过可用数量上限，建议删除已有模板后重试。
	FAILEDOPERATION_TEMPLATEMAXNUMFAILED = "FailedOperation.TemplateMaxNumFailed"

	// 上传图片操作失败。
	FAILEDOPERATION_UPLOADIMAGEFAILED = "FailedOperation.UploadImageFailed"

	// 当前用户未在官网进行实名操作。
	FAILEDOPERATION_VERIFYUINISREALNAME = "FailedOperation.VerifyUinIsRealname"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 网络报错，请稍后重试。
	INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 资质信息输入不正确。
	INVALIDPARAMETER_CERTIFICATECODEISINVALID = "InvalidParameter.CertificateCodeIsInvalid"

	// 资质照片输入不正确。
	INVALIDPARAMETER_CERTIFICATEIMAGEISINVALID = "InvalidParameter.CertificateImageIsInvalid"

	// 类型只能为手机或者邮箱。
	INVALIDPARAMETER_CODETYPEISINVALID = "InvalidParameter.CodeTypeIsInvalid"

	// 无权限自定义DNS。
	INVALIDPARAMETER_CUSTOMDNSNOTALLOWED = "InvalidParameter.CustomDnsNotAllowed"

	// 域名输入为空或者不正确。
	INVALIDPARAMETER_DOMAINNAMEISINVALID = "InvalidParameter.DomainNameIsInvalid"

	// 存在重复域名，请检查后重新提交。
	INVALIDPARAMETER_DUPLICATEDOMAINEXISTS = "InvalidParameter.DuplicateDomainExists"

	// 邮箱为空或者不正确。
	INVALIDPARAMETER_EMAILISINVALID = "InvalidParameter.EmailIsInvalid"

	// 仅支持已验证的电子邮箱，请先在控制台创建后使用
	INVALIDPARAMETER_EMAILISUNVERIFIED = "InvalidParameter.EmailIsUnverified"

	// 不支持该格式文件，请上传 JPG、JPEG 格式图片（可使用第三方图片格式转换工具）。
	INVALIDPARAMETER_IMAGEEXTINVALID = "InvalidParameter.ImageExtInvalid"

	// 上传的照片参数为空或者不合法。
	INVALIDPARAMETER_IMAGEFILEISINVALID = "InvalidParameter.ImageFileIsInvalid"

	// 非标准的 JPG、JPEG 格式图片，请使用工具转换格式后重新上传（可使用第三方图片格式转换工具）。
	INVALIDPARAMETER_IMAGEFORMATISINVALID = "InvalidParameter.ImageFormatIsInvalid"

	// 图片大小低于最小限制(55KB)，请重新上传。
	INVALIDPARAMETER_IMAGESIZEBELOW = "InvalidParameter.ImageSizeBelow"

	// 图片过大，请减小后重试。
	INVALIDPARAMETER_IMAGESIZEEXCEED = "InvalidParameter.ImageSizeExceed"

	// 图片大小超过限制(1M)，请重新上传。
	INVALIDPARAMETER_IMAGESIZELIMIT = "InvalidParameter.ImageSizeLimit"

	// 联系人为空或者不合法。
	INVALIDPARAMETER_NAMEISINVALID = "InvalidParameter.NameIsInvalid"

	// 联系人填写有误，或因其他原因无法使用，请更换其他联系人。
	INVALIDPARAMETER_NAMEISKEYWORD = "InvalidParameter.NameIsKeyword"

	// 注册人为空或者不合法。
	INVALIDPARAMETER_ORGISINVALID = "InvalidParameter.OrgIsInvalid"

	// 域名所有者填写有误，或因其他原因无法使用，请更换其他域名所有者。
	INVALIDPARAMETER_ORGISKEYWORD = "InvalidParameter.OrgIsKeyword"

	// 特惠包ID无效。
	INVALIDPARAMETER_PACKAGERESOURCEIDINVALID = "InvalidParameter.PackageResourceIdInvalid"

	// 请求类型错误。
	INVALIDPARAMETER_REPTYPEISINVALID = "InvalidParameter.RepTypeIsInvalid"

	// 地址有误，请传入正确的地址。
	INVALIDPARAMETER_STREETISINVALID = "InvalidParameter.StreetIsInvalid"

	// 电话为空或者不合法。
	INVALIDPARAMETER_TELEPHONEISINVALID = "InvalidParameter.TelephoneIsInvalid"

	// 仅支持已验证的手机号码，请先在控制台创建后使用。
	INVALIDPARAMETER_TELEPHONEISUNVERIFIED = "InvalidParameter.TelephoneIsUnverified"

	// 域名数量不能超过4000个。
	INVALIDPARAMETER_UPTO4000 = "InvalidParameter.UpTo4000"

	// 用户类型为空或者不合法。
	INVALIDPARAMETER_USERTYPEISINVALID = "InvalidParameter.UserTypeIsInvalid"

	// 验证码错误，请重新输入。
	INVALIDPARAMETER_VERIFYCODEISINVALID = "InvalidParameter.VerifyCodeIsInvalid"

	// 邮编为空或者不合法。
	INVALIDPARAMETER_ZIPCODEISINVALID = "InvalidParameter.ZipCodeIsInvalid"

	// 请求频率超过限制。
	LIMITEXCEEDED_REQUESTLIMIT = "LimitExceeded.RequestLimit"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 域名不能为空。
	MISSINGPARAMETER_DOMAINISEMPTY = "MissingParameter.DomainIsEmpty"

	// 请求数据不能为空。
	MISSINGPARAMETER_REPDATAISNONE = "MissingParameter.RepDataIsNone"

	// 模板ID为空或者不合法。
	MISSINGPARAMETER_TEMPLATEIDISEMPTY = "MissingParameter.TemplateIdIsEmpty"

	// 模板已存在。
	MISSINGPARAMETER_TEMPLATEIDISEXIST = "MissingParameter.TemplateIdIsExist"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 当前正在执行中的任务过多，请稍后再提交新的任务。
	RESOURCEINSUFFICIENT_OVERWORK = "ResourceInsufficient.Overwork"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 模板未实名。
	RESOURCENOTFOUND_APPROVEDTEMPLATENOTFOUND = "ResourceNotFound.ApprovedTemplateNotFound"

	// 域名地址有误，请输入正确的域名地址。
	RESOURCENOTFOUND_DOMAINNOTFOUND = "ResourceNotFound.DomainNotFound"

	// 模板信息有误，请输入正确的信息。
	RESOURCENOTFOUND_TEMPLATENOTFOUND = "ResourceNotFound.TemplateNotFound"

	// 该域名已有同类型操作未完成，无法执行该操作。
	RESOURCEUNAVAILABLE_DOMAINISMODIFYINGDNS = "ResourceUnavailable.DomainIsModifyingDNS"

	// 账户实名认证未通过。
	UNSUPPORTEDOPERATION_ACCOUNTREALNAME = "UnsupportedOperation.AccountRealName"

	// 当前域名未完成实名认证，无法完成该操作。
	UNSUPPORTEDOPERATION_DOMAINNOTVERIFIED = "UnsupportedOperation.DomainNotVerified"

	// 当前域名状态不支持修改。
	UNSUPPORTEDOPERATION_MODIFYDOMAININFOUNSUPPORTED = "UnsupportedOperation.ModifyDomainInfoUnsupported"

	// 当前域名状态不支持修改。
	UNSUPPORTEDOPERATION_MODIFYDOMAINUNSUPPORTED = "UnsupportedOperation.ModifyDomainUnsupported"
)
