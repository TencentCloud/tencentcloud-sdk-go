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
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type EvaluationRequestParams struct {
	// 图片唯一标识，一张图片一个SessionId；
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 图片数据，需要使用base64对图片的二进制数据进行编码，与url参数二者填一即可；
	Image *string `json:"Image,omitempty" name:"Image"`

	// 业务应用ID，与账号应用APPID无关，是用来方便客户管理服务的参数，新的 HcmAppid 可以在[控制台](https://console.cloud.tencent.com/hcm)【应用管理】下新建。
	HcmAppid *string `json:"HcmAppid,omitempty" name:"HcmAppid"`

	// 图片url，与Image参数二者填一即可；
	Url *string `json:"Url,omitempty" name:"Url"`

	// 横屏拍摄开关，若开启则支持传输横屏拍摄的图片；
	SupportHorizontalImage *bool `json:"SupportHorizontalImage,omitempty" name:"SupportHorizontalImage"`

	// 拒绝非速算图（如风景图、人物图）开关，若开启，则遇到非速算图会快速返回拒绝的结果，但极端情况下可能会影响评估结果（比如算式截图贴到风景画里可能被判为非速算图直接返回了）。
	RejectNonArithmeticImage *bool `json:"RejectNonArithmeticImage,omitempty" name:"RejectNonArithmeticImage"`

	// 异步模式标识，0：同步模式，1：异步模式。默认为同步模式
	IsAsync *int64 `json:"IsAsync,omitempty" name:"IsAsync"`

	// 是否展开耦合算式中的竖式计算
	EnableDispRelatedVertical *bool `json:"EnableDispRelatedVertical,omitempty" name:"EnableDispRelatedVertical"`

	// 是否展示竖式算式的中间结果和格式控制字符
	EnableDispMidresult *bool `json:"EnableDispMidresult,omitempty" name:"EnableDispMidresult"`

	// 是否开启pdf识别，默认开启
	EnablePdfRecognize *bool `json:"EnablePdfRecognize,omitempty" name:"EnablePdfRecognize"`

	// pdf页码，从0开始，默认为0
	PdfPageIndex *int64 `json:"PdfPageIndex,omitempty" name:"PdfPageIndex"`

	// 是否返回LaTex，默认为0返回普通格式，设置成1返回LaTex格式
	LaTex *int64 `json:"LaTex,omitempty" name:"LaTex"`

	// 用于选择是否拒绝模糊题 目。打开则丢弃模糊题目， 不进行后续的判题返回结 果。
	RejectVagueArithmetic *bool `json:"RejectVagueArithmetic,omitempty" name:"RejectVagueArithmetic"`
}

type EvaluationRequest struct {
	*tchttp.BaseRequest
	
	// 图片唯一标识，一张图片一个SessionId；
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 图片数据，需要使用base64对图片的二进制数据进行编码，与url参数二者填一即可；
	Image *string `json:"Image,omitempty" name:"Image"`

	// 业务应用ID，与账号应用APPID无关，是用来方便客户管理服务的参数，新的 HcmAppid 可以在[控制台](https://console.cloud.tencent.com/hcm)【应用管理】下新建。
	HcmAppid *string `json:"HcmAppid,omitempty" name:"HcmAppid"`

	// 图片url，与Image参数二者填一即可；
	Url *string `json:"Url,omitempty" name:"Url"`

	// 横屏拍摄开关，若开启则支持传输横屏拍摄的图片；
	SupportHorizontalImage *bool `json:"SupportHorizontalImage,omitempty" name:"SupportHorizontalImage"`

	// 拒绝非速算图（如风景图、人物图）开关，若开启，则遇到非速算图会快速返回拒绝的结果，但极端情况下可能会影响评估结果（比如算式截图贴到风景画里可能被判为非速算图直接返回了）。
	RejectNonArithmeticImage *bool `json:"RejectNonArithmeticImage,omitempty" name:"RejectNonArithmeticImage"`

	// 异步模式标识，0：同步模式，1：异步模式。默认为同步模式
	IsAsync *int64 `json:"IsAsync,omitempty" name:"IsAsync"`

	// 是否展开耦合算式中的竖式计算
	EnableDispRelatedVertical *bool `json:"EnableDispRelatedVertical,omitempty" name:"EnableDispRelatedVertical"`

	// 是否展示竖式算式的中间结果和格式控制字符
	EnableDispMidresult *bool `json:"EnableDispMidresult,omitempty" name:"EnableDispMidresult"`

	// 是否开启pdf识别，默认开启
	EnablePdfRecognize *bool `json:"EnablePdfRecognize,omitempty" name:"EnablePdfRecognize"`

	// pdf页码，从0开始，默认为0
	PdfPageIndex *int64 `json:"PdfPageIndex,omitempty" name:"PdfPageIndex"`

	// 是否返回LaTex，默认为0返回普通格式，设置成1返回LaTex格式
	LaTex *int64 `json:"LaTex,omitempty" name:"LaTex"`

	// 用于选择是否拒绝模糊题 目。打开则丢弃模糊题目， 不进行后续的判题返回结 果。
	RejectVagueArithmetic *bool `json:"RejectVagueArithmetic,omitempty" name:"RejectVagueArithmetic"`
}

func (r *EvaluationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EvaluationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	delete(f, "Image")
	delete(f, "HcmAppid")
	delete(f, "Url")
	delete(f, "SupportHorizontalImage")
	delete(f, "RejectNonArithmeticImage")
	delete(f, "IsAsync")
	delete(f, "EnableDispRelatedVertical")
	delete(f, "EnableDispMidresult")
	delete(f, "EnablePdfRecognize")
	delete(f, "PdfPageIndex")
	delete(f, "LaTex")
	delete(f, "RejectVagueArithmetic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EvaluationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EvaluationResponseParams struct {
	// 图片唯一标识，一张图片一个SessionId；
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 识别出的算式信息；
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*Item `json:"Items,omitempty" name:"Items"`

	// 任务 id，用于查询接口
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EvaluationResponse struct {
	*tchttp.BaseResponse
	Response *EvaluationResponseParams `json:"Response"`
}

func (r *EvaluationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EvaluationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Item struct {
	// 识别的算式是否正确，算式运算结果:
	// ‘YES’:正确 
	// ‘NO’: 错误 
	// ‘NA’: 非法参数
	// ‘EMPTY’: 未作答
	Item *string `json:"Item,omitempty" name:"Item"`

	// 识别出的算式，识别出的文本行字符串
	ItemString *string `json:"ItemString,omitempty" name:"ItemString"`

	// 识别的算式在图片上的位置信息，文本行在旋转纠正之后的图像中的像素坐 标，表示为(左上角 x, 左上角 y，宽 width， 高 height)
	ItemCoord *ItemCoord `json:"ItemCoord,omitempty" name:"ItemCoord"`

	// 错题推荐答案，算式运算结果正确返回为 ""，算式运算结果错误返回推荐答案 (注:暂不支持多个关系运算符(如 1<10<7)、 无关系运算符(如 frac(1,2)+frac(2,3))、单 位换算(如 1 元=100 角)错题的推荐答案 返回)
	// (注:使用@@标记答案填写区域)
	Answer *string `json:"Answer,omitempty" name:"Answer"`

	// 算式题型编号，如加减乘除四则题型，具体题型及编号如下：1 加减乘除四则 2 加减乘除已知结果求运算因子3 判断大小 4 约等于估算 5 带余数除法 6 分数四则运算 7 单位换算 8 竖式加减法 9 竖式乘除法 10 脱式计算 11 解方程
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpressionType *string `json:"ExpressionType,omitempty" name:"ExpressionType"`

	// 文本行置信度
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemConf *float64 `json:"ItemConf,omitempty" name:"ItemConf"`

	// 用于标识题目 id，如果有若干算式属于同一 题，则其对应的 id 相同。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuestionId *string `json:"QuestionId,omitempty" name:"QuestionId"`
}

type ItemCoord struct {
	// 算式高度
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 算式宽度
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 算式图的左上角横坐标
	X *int64 `json:"X,omitempty" name:"X"`

	// 算式图的左上角纵坐标
	Y *int64 `json:"Y,omitempty" name:"Y"`
}