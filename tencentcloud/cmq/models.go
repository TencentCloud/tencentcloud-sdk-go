package cmq

import tchttp "github.com/teamlint/tencentcloud-sdk-go/tencentcloud/common/http"

const (
	FilterTypeTag     = 1 // 消息过滤类型-标签
	FilterTypeRouting = 2 // 消息过滤类型-路由匹配

	V2Path = "/v2/index.php"
)

// CreateTopicRequest 创建主题请求
type CreateTopicRequest struct {
	*tchttp.BaseRequest
	// 主题名字,主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)
	TopicName string `json:"topicName" name:"topicName"`
	// 消息最大长度,取值范围1024 - 65536 Byte（即1 - 64K）
	MaxMsgSize int `json:"maxMsgSize,omitempty" name:"maxMsgSize"`
	// 用于指定主题的消息匹配策略： filterType =1 或为空， 表示该主题下所有订阅使用 filterTag 标签过滤； filterType =2 表示用户使用 bindingKey 过滤。
	FilterType int `json:"filterType,omitempty" name:"filterType"`
}

// CreateTopicResponse 创建主题响应
type CreateTopicResponse struct {
	*tchttp.BaseResponse
	// 0：表示成功，others：错误
	Code int `json:"code" name:"code"`
	// 错误提示信息
	Message string `json:"message" name:"message"`
	// 服务器生成的请求 ID
	RequestId string `json:"requestId" name:"requestId"`
	// 主题标识ID
	TopicId string `json:"topicId" name:"topicId"`
}

// PublishMessage 发布消息请求
type PublishMessageRequest struct {
	*tchttp.BaseRequest
	TopicName string `json:"topicName" name:"topicName"`
	// 消息正文。至少1Byte，最大长度受限于设置的主题消息最大长度属性
	MsgBody string `json:"msgBody" name:"msgBody"`
	// 消息过滤标签
	MsgTag []string `json:"msgTag" name:"msgTag"`
	// 路由键
	RoutingKey string `json:"routingKey" name:"routingKey"`
}

// PublishMessageResponse 发布消息响应
type PublishMessageResponse struct {
	*tchttp.BaseResponse
	// 0：表示成功，others：错误
	Code int `json:"code" name:"code"`
	// 错误提示信息
	Message string `json:"message" name:"message"`
	// 服务器生成的请求 ID
	RequestId string `json:"requestId" name:"requestId"`
	// 消息标识ID
	MsgId string `json:"msgId" name:"msgId"`
}
