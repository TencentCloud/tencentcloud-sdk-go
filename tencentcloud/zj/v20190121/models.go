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

package v20190121

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AddCrowdPackInfoRequest struct {
	*tchttp.BaseRequest

	// 商户证书
	License *string `json:"License,omitempty" name:"License"`

	// 人群包名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 人群包文件名称,人群包文件必须为utf8编码，动态参数只能是汉字、数字、英文字母的组合，不能包含其他字符
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 人群包描述
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 已经上传好的人群包cos地址
	CosUrl *string `json:"CosUrl,omitempty" name:"CosUrl"`

	// 人群包手机号数量
	PhoneNum *int64 `json:"PhoneNum,omitempty" name:"PhoneNum"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCrowdPackInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "License")
	delete(f, "Name")
	delete(f, "FileName")
	delete(f, "Desc")
	delete(f, "CosUrl")
	delete(f, "PhoneNum")
	if len(f) > 0 {
		return errors.New("AddCrowdPackInfoRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AddCrowdPackInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 接口返回
		Data *SmsAddCrowdPackInfoResponse `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCrowdPackInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddSmsSignRequest struct {
	*tchttp.BaseRequest

	// 商户证书
	License *string `json:"License,omitempty" name:"License"`

	// 签名类型。其中每种类型后面标注了其可选的 DocumentType（证明类型）：
	// 0：公司（0，1，2，3）。
	// 1：APP（0，1，2，3，4） 。
	// 2：网站（0，1，2，3，5）。
	// 3：公众号或者小程序（0，1，2，3，6）。
	// 4：商标（7）。
	// 5：政府/机关事业单位/其他机构（2，3）。
	// 注：必须按照对应关系选择证明类型，否则会审核失败。
	SignType *uint64 `json:"SignType,omitempty" name:"SignType"`

	// 证明类型：
	// 0：三证合一。
	// 1：企业营业执照。
	// 2：组织机构代码证书。
	// 3：社会信用代码证书。
	// 4：应用后台管理截图（个人开发APP）。
	// 5：网站备案后台截图（个人开发网站）。
	// 6：小程序设置页面截图（个人认证小程序）。
	// 7：商标注册书
	DocumentType *uint64 `json:"DocumentType,omitempty" name:"DocumentType"`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitempty" name:"International"`

	// 资质图片url
	ProofImage *string `json:"ProofImage,omitempty" name:"ProofImage"`

	// 签名内容
	SignName *string `json:"SignName,omitempty" name:"SignName"`

	// 签名备注，比如申请原因，使用场景等,可以填空
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddSmsSignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "License")
	delete(f, "SignType")
	delete(f, "DocumentType")
	delete(f, "International")
	delete(f, "ProofImage")
	delete(f, "SignName")
	delete(f, "Remark")
	if len(f) > 0 {
		return errors.New("AddSmsSignRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AddSmsSignResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 签名id数组
		Data *PaasCreateSignResp `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddSmsSignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddSmsTemplateDataStruct struct {

	// 短信模板ID
	TemplateId *uint64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

type AddSmsTemplateRequest struct {
	*tchttp.BaseRequest

	// 商户证书
	License *string `json:"License,omitempty" name:"License"`

	// 短信签名，创建签名时返回
	SignID *uint64 `json:"SignID,omitempty" name:"SignID"`

	// 模板名称
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 短信内容，动态内容使用占位符{1}，{2}等表示
	TemplateContent *string `json:"TemplateContent,omitempty" name:"TemplateContent"`

	// 短信类型：{0:普通短信，1:营销短信}
	SmsType *uint64 `json:"SmsType,omitempty" name:"SmsType"`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitempty" name:"International"`

	// 短信模板标签
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 发送短信活动时配置的落地链接地址,仅用作短信活动
	Urls []*string `json:"Urls,omitempty" name:"Urls" list`

	// 发送短信活动时用于展示人群包动态参数模板占位符序号或接口发送时变量占位符序号
	CommonParams []*int64 `json:"CommonParams,omitempty" name:"CommonParams" list`

	// 发送短信活动时用于展示短连接模板占位符序号,仅用作短信活动
	UrlParams []*int64 `json:"UrlParams,omitempty" name:"UrlParams" list`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddSmsTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "License")
	delete(f, "SignID")
	delete(f, "TemplateName")
	delete(f, "TemplateContent")
	delete(f, "SmsType")
	delete(f, "International")
	delete(f, "Remark")
	delete(f, "Urls")
	delete(f, "CommonParams")
	delete(f, "UrlParams")
	if len(f) > 0 {
		return errors.New("AddSmsTemplateRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AddSmsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 短信模板创建接口返回
		Data *AddSmsTemplateDataStruct `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddSmsTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelActivityData struct {

	// 成功返回时的文字描述
	Message *string `json:"Message,omitempty" name:"Message"`
}

type CancelCampaignRequest struct {
	*tchttp.BaseRequest

	// 商户证书
	License *string `json:"License,omitempty" name:"License"`

	// 短信活动id
	CampaignId *int64 `json:"CampaignId,omitempty" name:"CampaignId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelCampaignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "License")
	delete(f, "CampaignId")
	if len(f) > 0 {
		return errors.New("CancelCampaignRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CancelCampaignResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 处理结果
		Data *CancelActivityData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelCampaignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCampaignRequest struct {
	*tchttp.BaseRequest

	// 商户证书
	License *string `json:"License,omitempty" name:"License"`

	// 短信活动发送时间
	SendTime *int64 `json:"SendTime,omitempty" name:"SendTime"`

	// 短信活动名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 发送策略
	Strategies []*PaasStrategy `json:"Strategies,omitempty" name:"Strategies" list`

	// 废弃
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 废弃
	CrowdID *int64 `json:"CrowdID,omitempty" name:"CrowdID"`

	// 活动类型(0-短信,1-超短,不填默认为超短)
	SmsType *uint64 `json:"SmsType,omitempty" name:"SmsType"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCampaignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "License")
	delete(f, "SendTime")
	delete(f, "Name")
	delete(f, "Strategies")
	delete(f, "TemplateId")
	delete(f, "CrowdID")
	delete(f, "SmsType")
	if len(f) > 0 {
		return errors.New("CreateCampaignRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateCampaignResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 活动信息
		Data *SmsCreateCampaignResponse `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCampaignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMmsInstanceItem struct {

	// 素材类型：1-文本 2-图片 3-视频 4-音频
	ContentType *int64 `json:"ContentType,omitempty" name:"ContentType"`

	// 素材内容：如果素材是文本类型，直接填写文本内容，否则填写素材文件上传到cos后的url地址
	Content *string `json:"Content,omitempty" name:"Content"`
}

type CreateMmsInstanceRequest struct {
	*tchttp.BaseRequest

	// 商户证书
	License *string `json:"License,omitempty" name:"License"`

	// 样例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 标题
	Title *string `json:"Title,omitempty" name:"Title"`

	// 签名
	Sign *string `json:"Sign,omitempty" name:"Sign"`

	// 素材内容
	Contents []*CreateMmsInstanceItem `json:"Contents,omitempty" name:"Contents" list`

	// 样例中链接动态变量对应的链接，和占位符顺序一致
	Urls []*string `json:"Urls,omitempty" name:"Urls" list`

	// 机型列表
	PhoneType []*uint64 `json:"PhoneType,omitempty" name:"PhoneType" list`

	// 发送超短活动时用于展示人群包动态参数模板占位符序号或接口发送时变量占位符序号
	CommonParams []*uint64 `json:"CommonParams,omitempty" name:"CommonParams" list`

	// 发送超短活动时用于展示短连接模板占位符序号,仅用作超短活动
	UrlParams []*uint64 `json:"UrlParams,omitempty" name:"UrlParams" list`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMmsInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "License")
	delete(f, "InstanceName")
	delete(f, "Title")
	delete(f, "Sign")
	delete(f, "Contents")
	delete(f, "Urls")
	delete(f, "PhoneType")
	delete(f, "CommonParams")
	delete(f, "UrlParams")
	if len(f) > 0 {
		return errors.New("CreateMmsInstanceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateMmsInstanceResp struct {

	// 返回码：0-成功 其它-失败
	ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`

	// 返回信息
	ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`

	// 样例id
	InstanceId *int64 `json:"InstanceId,omitempty" name:"InstanceId"`
}

type CreateMmsInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建样例返回信息
		Data *CreateMmsInstanceResp `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMmsInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DelCrowdPackRequest struct {
	*tchttp.BaseRequest

	// 商户证书
	License *string `json:"License,omitempty" name:"License"`

	// 人群包id
	ID *uint64 `json:"ID,omitempty" name:"ID"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DelCrowdPackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "License")
	delete(f, "ID")
	if len(f) > 0 {
		return errors.New("DelCrowdPackRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DelCrowdPackResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 接口返回
		Data *SmsSuccessResponse `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DelCrowdPackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DelMmsInstanceData struct {

	// 样例id
	InstanceId *uint64 `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DelTemplateRequest struct {
	*tchttp.BaseRequest

	// 商户证书
	License *string `json:"License,omitempty" name:"License"`

	// 短信模板ID
	TemplateID *int64 `json:"TemplateID,omitempty" name:"TemplateID"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DelTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "License")
	delete(f, "TemplateID")
	if len(f) > 0 {
		return errors.New("DelTemplateRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DelTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 接口返回
		Data *SmsSuccessResponse `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DelTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMmsInstanceRequest struct {
	*tchttp.BaseRequest

	// 商户证书
	License *string `json:"License,omitempty" name:"License"`

	// 超级短信样例id
	InstanceId *uint64 `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMmsInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "License")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return errors.New("DeleteMmsInstanceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMmsInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除信息返回
		Data *DelMmsInstanceData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMmsInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMmsInstanceInfoRequest struct {
	*tchttp.BaseRequest

	// 商户证书
	License *string `json:"License,omitempty" name:"License"`

	// 彩信实例id
	InstanceId *int64 `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMmsInstanceInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "License")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return errors.New("DescribeMmsInstanceInfoRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMmsInstanceInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 彩信实例信息
		Data *MmsInstanceInfo `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMmsInstanceInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMmsInstanceListRequest struct {
	*tchttp.BaseRequest

	// 商户证书
	License *string `json:"License,omitempty" name:"License"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 业务代码
	AppSubId *string `json:"AppSubId,omitempty" name:"AppSubId"`

	// 实例标题
	Title *string `json:"Title,omitempty" name:"Title"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMmsInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "License")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "AppSubId")
	delete(f, "Title")
	if len(f) > 0 {
		return errors.New("DescribeMmsInstanceListRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMmsInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 彩信实例信息列表返回
		Data *MmsInstanceInfoList `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMmsInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSmsCampaignStatisticsRequest struct {
	*tchttp.BaseRequest

	// 活动id
	CampaignId *uint64 `json:"CampaignId,omitempty" name:"CampaignId"`

	// 商户证书
	License *string `json:"License,omitempty" name:"License"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmsCampaignStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CampaignId")
	delete(f, "License")
	if len(f) > 0 {
		return errors.New("DescribeSmsCampaignStatisticsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSmsCampaignStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 响应数据
		Data *SmsCampaignStatisticsData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmsCampaignStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSmsSignListDataStruct struct {

	// 签名Id
	SignId *uint64 `json:"SignId,omitempty" name:"SignId"`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitempty" name:"International"`

	// 申请签名状态。其中：
	// 0：表示审核通过。
	// -1：表示审核未通过或审核失败。
	StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// 审核回复，审核人员审核后给出的回复，通常是审核未通过的原因。
	ReviewReply *string `json:"ReviewReply,omitempty" name:"ReviewReply"`

	// 签名名称。
	SignName *string `json:"SignName,omitempty" name:"SignName"`

	// 提交审核时间，UNIX 时间戳（单位：秒）。
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DescribeSmsSignListRequest struct {
	*tchttp.BaseRequest

	// 商户证书
	License *string `json:"License,omitempty" name:"License"`

	// 签名ID数组
	SignIdSet []*uint64 `json:"SignIdSet,omitempty" name:"SignIdSet" list`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitempty" name:"International"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmsSignListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "License")
	delete(f, "SignIdSet")
	delete(f, "International")
	if len(f) > 0 {
		return errors.New("DescribeSmsSignListRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSmsSignListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data []*DescribeSmsSignListDataStruct `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmsSignListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSmsTemplateListDataStruct struct {

	// 模板Id
	TemplateId *uint64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitempty" name:"International"`

	// 申请签名状态。其中：
	// 0：表示审核通过。
	// -1：表示审核未通过或审核失败。
	StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// 审核回复，审核人员审核后给出的回复，通常是审核未通过的原因。
	ReviewReply *string `json:"ReviewReply,omitempty" name:"ReviewReply"`

	// 模板名称。
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 提交审核时间，UNIX 时间戳（单位：秒）。
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DescribeSmsTemplateListRequest struct {
	*tchttp.BaseRequest

	// 商户证书
	License *string `json:"License,omitempty" name:"License"`

	// 短信模板id数组
	TemplateIdSet []*uint64 `json:"TemplateIdSet,omitempty" name:"TemplateIdSet" list`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitempty" name:"International"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmsTemplateListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "License")
	delete(f, "TemplateIdSet")
	delete(f, "International")
	if len(f) > 0 {
		return errors.New("DescribeSmsTemplateListRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSmsTemplateListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据信息
		Data []*DescribeSmsTemplateListDataStruct `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmsTemplateListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCrowdPackListRequest struct {
	*tchttp.BaseRequest

	// 商户证书
	License *string `json:"License,omitempty" name:"License"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制返回数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 人群包名称，用于过滤人群包
	Name *string `json:"Name,omitempty" name:"Name"`

	// 人群包状态，默认-1，用于过滤人群包
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCrowdPackListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "License")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Name")
	delete(f, "Status")
	if len(f) > 0 {
		return errors.New("GetCrowdPackListRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetCrowdPackListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 人群包信息列表
		Data *SmsGetCrowdPackListResponse `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCrowdPackListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCrowdUploadInfoRequest struct {
	*tchttp.BaseRequest

	// 商户证书
	License *string `json:"License,omitempty" name:"License"`

	// 上传文件名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCrowdUploadInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "License")
	delete(f, "FileName")
	if len(f) > 0 {
		return errors.New("GetCrowdUploadInfoRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetCrowdUploadInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *SmsGetCrowdUploadInfoResponse `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCrowdUploadInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSmsAmountInfoRequest struct {
	*tchttp.BaseRequest

	// 商户证书
	License *string `json:"License,omitempty" name:"License"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSmsAmountInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "License")
	if len(f) > 0 {
		return errors.New("GetSmsAmountInfoRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetSmsAmountInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 短信账号额度接口
		Data *SmsAmountDataStruct `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSmsAmountInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSmsCampaignStatusRequest struct {
	*tchttp.BaseRequest

	// 商户证书
	License *string `json:"License,omitempty" name:"License"`

	// 活动ID
	CampaignId *int64 `json:"CampaignId,omitempty" name:"CampaignId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSmsCampaignStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "License")
	delete(f, "CampaignId")
	if len(f) > 0 {
		return errors.New("GetSmsCampaignStatusRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetSmsCampaignStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 活动状态
		Data *PaasSmsCampaignStatusResp `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSmsCampaignStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MmsInstanceInfo struct {

	// 彩信实例id
	InstanceId *int64 `json:"InstanceId,omitempty" name:"InstanceId"`

	// 彩信实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 状态是否通知
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 实例审核状态信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusInfo []*MmsInstanceStateInfo `json:"StatusInfo,omitempty" name:"StatusInfo" list`

	// 业务码
	AppSubId *string `json:"AppSubId,omitempty" name:"AppSubId"`

	// 彩信标题
	Title *string `json:"Title,omitempty" name:"Title"`

	// 签名
	Sign *string `json:"Sign,omitempty" name:"Sign"`

	// 彩信内容
	Contents *string `json:"Contents,omitempty" name:"Contents"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 样例配置的链接地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Urls []*string `json:"Urls,omitempty" name:"Urls" list`

	// 机型列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneType []*uint64 `json:"PhoneType,omitempty" name:"PhoneType" list`

	// 普通参数序号数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	CommonParams []*uint64 `json:"CommonParams,omitempty" name:"CommonParams" list`

	// 链接参数序号数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	UrlParams []*uint64 `json:"UrlParams,omitempty" name:"UrlParams" list`
}

type MmsInstanceInfoList struct {

	// 总数据量
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 彩信实例状态信息列表
	List []*MmsInstanceInfo `json:"List,omitempty" name:"List" list`
}

type MmsInstanceStateInfo struct {

	// 运营商
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 审核状态：0未审核，1审核通过，2审核拒绝
	State *int64 `json:"State,omitempty" name:"State"`
}

type ModifySmsTemplateDataStruct struct {

	// 短信模板id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

type ModifySmsTemplateRequest struct {
	*tchttp.BaseRequest

	// 商户证书
	License *string `json:"License,omitempty" name:"License"`

	// 短信模板id
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 短信签名，创建签名时返回
	SignID *uint64 `json:"SignID,omitempty" name:"SignID"`

	// 模板名称
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 短信内容，动态内容使用占位符{1}，{2}等表示
	TemplateContent *string `json:"TemplateContent,omitempty" name:"TemplateContent"`

	// 短信类型：{0:普通短信，1:营销短信}
	SmsType *uint64 `json:"SmsType,omitempty" name:"SmsType"`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitempty" name:"International"`

	// 短信模板标签
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 发送短信活动时配置的落地链接地址,仅用作短信活动
	Urls []*string `json:"Urls,omitempty" name:"Urls" list`

	// 发送短信活动时用于展示人群包动态参数模板占位符序号,仅用作短信活动
	CommonParams []*int64 `json:"CommonParams,omitempty" name:"CommonParams" list`

	// 发送短信活动时用于展示短连接模板占位符序号,仅用作短信活动
	UrlParams []*int64 `json:"UrlParams,omitempty" name:"UrlParams" list`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySmsTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "License")
	delete(f, "TemplateId")
	delete(f, "SignID")
	delete(f, "TemplateName")
	delete(f, "TemplateContent")
	delete(f, "SmsType")
	delete(f, "International")
	delete(f, "Remark")
	delete(f, "Urls")
	delete(f, "CommonParams")
	delete(f, "UrlParams")
	if len(f) > 0 {
		return errors.New("ModifySmsTemplateRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifySmsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回
		Data *ModifySmsTemplateDataStruct `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySmsTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PaasCreateSignResp struct {

	// 签名id
	SignId *int64 `json:"SignId,omitempty" name:"SignId"`
}

type PaasSmsCampaignStatusResp struct {

	// 0-未发送 1-发送中 2-发送结束 3-发送取消
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type PaasStrategy struct {

	// 人群包id
	CrowdID *int64 `json:"CrowdID,omitempty" name:"CrowdID"`

	// 待选素材数组
	Items []*PaasStrategyItem `json:"Items,omitempty" name:"Items" list`
}

type PaasStrategyItem struct {

	// 短信模板id或超级短信样例id
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 素材类型 0-普短 1-超短
	ContentType *int64 `json:"ContentType,omitempty" name:"ContentType"`
}

type PushMmsContentRequest struct {
	*tchttp.BaseRequest

	// 商户证书
	License *string `json:"License,omitempty" name:"License"`

	// 素材样例id
	InstanceId *int64 `json:"InstanceId,omitempty" name:"InstanceId"`

	// 手机号
	Tel *string `json:"Tel,omitempty" name:"Tel"`

	// 附带数据字段
	Session *string `json:"Session,omitempty" name:"Session"`

	// 动态参数key(即申请样例时设置的u_或p_开头的动态参数,要求序号有序)
	DynamicParaKey []*string `json:"DynamicParaKey,omitempty" name:"DynamicParaKey" list`

	// 动态参数值,和DynamicParaKey对应
	DynamicParaValue []*string `json:"DynamicParaValue,omitempty" name:"DynamicParaValue" list`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PushMmsContentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "License")
	delete(f, "InstanceId")
	delete(f, "Tel")
	delete(f, "Session")
	delete(f, "DynamicParaKey")
	delete(f, "DynamicParaValue")
	if len(f) > 0 {
		return errors.New("PushMmsContentRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type PushMmsContentResp struct {

	// 返回码：0-成功 其它-失败
	ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`

	// 返回信息
	ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`

	// 消息回执id
	MessageId *int64 `json:"MessageId,omitempty" name:"MessageId"`
}

type PushMmsContentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 推送短信返回信息
		Data *PushMmsContentResp `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PushMmsContentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendSmsPaasDataStruct struct {

	// 发送流水号
	SerialNo *string `json:"SerialNo,omitempty" name:"SerialNo"`

	// 手机号码,e.164标准，+[国家或地区码][手机号] ，示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 计费条数
	Fee *uint64 `json:"Fee,omitempty" name:"Fee"`

	// OK为成功
	Code *string `json:"Code,omitempty" name:"Code"`

	// 短信请求错误码描述
	Message *string `json:"Message,omitempty" name:"Message"`
}

type SendSmsRequest struct {
	*tchttp.BaseRequest

	// 商户证书
	License *string `json:"License,omitempty" name:"License"`

	// 手机号码,采用 e.164 标准，格式为+[国家或地区码][手机号]，单次请求最多支持200个手机号且要求全为境内手机号,如:+8613800138000
	Phone []*string `json:"Phone,omitempty" name:"Phone" list`

	// 短信模板id(推荐使用模板id发送,使用内容发送时模板id留空)
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 模板参数，若无模板参数，则设置为空。
	Params []*string `json:"Params,omitempty" name:"Params" list`

	// 短信签名内容，使用 UTF-8 编码，必须填写已审核通过的签名。注：国内短信为必填参数。
	Sign *string `json:"Sign,omitempty" name:"Sign"`

	// 国际/港澳台短信 senderid，国内短信填空
	SenderId *string `json:"SenderId,omitempty" name:"SenderId"`

	// 短信类型：{0:普通短信，1:营销短信}，使用内容发送时必填
	SmsType *uint64 `json:"SmsType,omitempty" name:"SmsType"`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。使用内容发送时必填
	International *uint64 `json:"International,omitempty" name:"International"`

	// 发送使用的模板内容,如果有占位符,此处也包括占位符,占位符的实际内容通过Params参数传递,使用模板id发送时此字段为空
	Content *string `json:"Content,omitempty" name:"Content"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendSmsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "License")
	delete(f, "Phone")
	delete(f, "TemplateId")
	delete(f, "Params")
	delete(f, "Sign")
	delete(f, "SenderId")
	delete(f, "SmsType")
	delete(f, "International")
	delete(f, "Content")
	if len(f) > 0 {
		return errors.New("SendSmsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SendSmsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 出参数据
		Data []*SendSmsPaasDataStruct `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendSmsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SmsAddCrowdPackInfoResponse struct {

	// 人群包id
	ID *uint64 `json:"ID,omitempty" name:"ID"`
}

type SmsAmountDataStruct struct {

	// 短信活动配额
	SmsCampaignAmount *uint64 `json:"SmsCampaignAmount,omitempty" name:"SmsCampaignAmount"`

	// 短信活动消耗配额
	SmsCampaignConsume *uint64 `json:"SmsCampaignConsume,omitempty" name:"SmsCampaignConsume"`

	// 短信发送额度
	SmsSendAmount *uint64 `json:"SmsSendAmount,omitempty" name:"SmsSendAmount"`

	// 短信发送消耗额度
	SmsSendConsume *uint64 `json:"SmsSendConsume,omitempty" name:"SmsSendConsume"`

	// 超短活动额度
	MmsCampaignAmount *uint64 `json:"MmsCampaignAmount,omitempty" name:"MmsCampaignAmount"`

	// 超短活动消耗额度
	MmsCampaignConsume *uint64 `json:"MmsCampaignConsume,omitempty" name:"MmsCampaignConsume"`

	// 超短短信额度
	MmsSendAmount *uint64 `json:"MmsSendAmount,omitempty" name:"MmsSendAmount"`

	// 超短短信消耗额度
	MmsSendConsume *uint64 `json:"MmsSendConsume,omitempty" name:"MmsSendConsume"`
}

type SmsCampaignStatisticsCrowdData struct {

	// 人群包id
	CrowdId *uint64 `json:"CrowdId,omitempty" name:"CrowdId"`

	// 人群包名称
	CrowdName *string `json:"CrowdName,omitempty" name:"CrowdName"`

	// 人群包目标触达总数
	CrowdCount *uint64 `json:"CrowdCount,omitempty" name:"CrowdCount"`

	// 模板列表
	TemplateList []*SmsCampaignStatisticsTemplateData `json:"TemplateList,omitempty" name:"TemplateList" list`
}

type SmsCampaignStatisticsData struct {

	// 活动Id
	CampaignId *uint64 `json:"CampaignId,omitempty" name:"CampaignId"`

	// 统计数据
	Statistics []*SmsCampaignStatisticsCrowdData `json:"Statistics,omitempty" name:"Statistics" list`
}

type SmsCampaignStatisticsTemplateData struct {

	// 模板或样例id
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 模板内容
	TemplateContent *string `json:"TemplateContent,omitempty" name:"TemplateContent"`

	// 触达成功数
	SendCount *uint64 `json:"SendCount,omitempty" name:"SendCount"`

	// 短链点击数
	ClickCount *uint64 `json:"ClickCount,omitempty" name:"ClickCount"`
}

type SmsCreateCampaignResponse struct {

	// 活动id
	CampaignId *int64 `json:"CampaignId,omitempty" name:"CampaignId"`
}

type SmsGetCrowdPackList struct {

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 人群包id
	ID *uint64 `json:"ID,omitempty" name:"ID"`

	// 人群包名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 人群包状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 人群包手机号数量
	PhoneNum *int64 `json:"PhoneNum,omitempty" name:"PhoneNum"`

	// 人群包标签信息
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 人群包md5
	MD5 *string `json:"MD5,omitempty" name:"MD5"`

	// 人群包文件名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 人群包描述
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

type SmsGetCrowdPackListResponse struct {

	// 人群包总数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 人群包返回数据列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*SmsGetCrowdPackList `json:"List,omitempty" name:"List" list`
}

type SmsGetCrowdUploadInfoResponse struct {

	// 过期时间
	ExpiredTime *uint64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// 会话token
	SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`

	// 临时密钥id
	TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`

	// 临时密钥
	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`

	// cos信息
	CosInfo *UploadFansInfoCosInfo `json:"CosInfo,omitempty" name:"CosInfo"`
}

type SmsSuccessResponse struct {

	// 成功返回信息
	Message *string `json:"Message,omitempty" name:"Message"`
}

type UploadFansInfoCosInfo struct {

	// COS bucket
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// COS路径
	Key *string `json:"Key,omitempty" name:"Key"`

	// COS区域
	Region *string `json:"Region,omitempty" name:"Region"`
}
