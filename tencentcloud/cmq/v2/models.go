package v2

import tchttp "github.com/teamlint/tencentcloud-sdk-go/tencentcloud/common/http"

const (
	V2Path = "/v2/index.php"

	FilterTypeTag        = 1 // 消息过滤类型-标签
	FilterTypeBindingKey = 2 // 消息过滤类型-路由键

	SubscribeProtocolHTTP  = "http"  // 主题订阅协议HTTP, 使用 HTTP 协议，用户需自己搭建接受消息的 Web Server
	SubscribeProtocolQueue = "queue" // 主题订阅协议Queue, 使用 Queue，消息会自动推送到 CMQ Queue，用户可以并发地拉取消息

	NotifyStrategyBackoffRetry          = "BACKOFF_RETRY"           // 退避重试,重试三次丢弃
	NotifyStrategyExponentialDecayRetry = "EXPONENTIAL_DECAY_RETRY" // 指数衰退重试,1,2,4,8,最长一天
	NotifyContentFormatJSON             = "JSON"                    // 推送内容格式 JSON
	NotifyContentFormatSimplified       = "SIMPLIFIED"              // 推送内容格式 SIMPLIFIED，即 raw 格式。如果 protocol 是 queue，则取值必须为 SIMPLIFIED

)

// CreateQueueRequest 创建队列请求
type CreateQueueRequest struct {
	*tchttp.BaseRequest
	QueueName           *string `json:"queueName,omitempty" name:"queueName"`                     // 队列名字
	MaxMsgHeapNum       *int    `json:"maxMsgHeapNum,omitempty" name:"maxMsgHeapNum"`             // 最大堆积消息数, 取值范围在公测期间为 1,000,000 - 10,000,000
	PollingWaitSeconds  *int    `json:"pollingWaitSeconds,omitempty" name:"pollingWaitSeconds"`   // 消息接收长轮询等待时间(秒)
	VisibilityTimeout   *int    `json:"visibilityTimeout,omitempty" name:"visibilityTimeout"`     // 消息可见性超时时间(秒), 取值范围1 - 43200秒（即12小时内），默认值30
	MaxMsgSize          *int    `json:"maxMsgSize,omitempty" name:"maxMsgSize"`                   // 消息最大长度(byte),取值范围1024 - 65536 Byte（即1-64K），默认值65536
	MsgRetentionSeconds *int    `json:"msgRetentionSeconds,omitempty" name:"msgRetentionSeconds"` // 消息保留周期(秒),取值范围60 - 1296000秒（1min-15天），默认值345600 (4 天)
	RewindSeconds       *int    `json:"rewindSeconds,omitempty" name:"rewindSeconds"`             // 队列是否开启回溯消息能力，该参数取值范围 0-msgRetentionSeconds，即最大的回溯时间为消息在队列中的保留周期，0表示不开启
}

// CreateQueueResponse 创建队列响应
type CreateQueueResponse struct {
	*tchttp.BaseResponse
	// 0：表示成功，others：错误
	Code *int `json:"code,omitempty" name:"code"`
	// 错误提示信息
	Message *string `json:"message,omitempty" name:"message"`
	// 服务器生成的请求 ID
	RequestId *string `json:"requestId,omitempty" name:"requestId"`
	// 主题标识ID
	QueueId *string `json:"queueId,omitempty" name:"queueId"`
}

// ListQueueRequest 获取队列列表请求
type ListQueueRequest struct {
	*tchttp.BaseRequest
	SearchWord *string `json:"serarchWord,omitempty" name:"searchWord"` // 用于过滤队列列表模糊搜索,不填选择所有
	Offset     *int    `json:"offset,omitempty" name:"offset"`
	Limit      *int    `json:"limit,omitempty" name:"limit"`
}

// ListQueueResponse 获取队列列表响应
type ListQueueResponse struct {
	*tchttp.BaseResponse
	Code       *int    `json:"code,omitempty" name:"code"`
	Message    *string `json:"message,omitempty" name:"message"`
	RequestId  *string `json:"requestId,omitempty" name:"requestId"`
	TotalCount *int    `json:"totalCount,omitempty" name:"totalCount"`
	QueueList  []*struct {
		QueueId   string `json:"queueId,omitempty" name:"queueId"`
		QueueName string `json:"queueName,omitempty" name:"queueName"`
	} `json:"queueList,omitempty" name:"queueList"`
}

// GetQueueAttributesRequest 获取队列属性请求
type GetQueueAttributesRequest struct {
	*tchttp.BaseRequest
	QueueName *string `json:"queueName,omitempty" name:"queueName"`
}

// GetQueueAttributesResponse 获取队列属性响应
type GetQueueAttributesResponse struct {
	*tchttp.BaseResponse
	Code                *int    `json:"code,omitempty" name:"code"`
	Message             *string `json:"message,omitempty" name:"message"`
	RequestId           *string `json:"requestId,omitempty" name:"requestId"`
	MaxMsgHeapNum       *int    `json:"maxMsgHeapNum,omitempty" name:"maxMsgHeapNum"`             // 最大堆积消息数
	PollingWaitSeconds  *int    `json:"pollingWaitSeconds,omitempty" name:"pollingWaitSeconds"`   // 消息接收长轮询等待时间(秒)
	VisibilityTimeout   *int    `json:"visibilityTimeout,omitempty" name:"visibilityTimeout"`     // 消息可见性超时时间(秒)
	MaxMsgSize          *int    `json:"maxMsgSize,omitempty" name:"maxMsgSize"`                   // 消息最大长度(byte)
	MsgRetentionSeconds *int    `json:"msgRetentionSeconds,omitempty" name:"msgRetentionSeconds"` // 消息保留周期(秒)
	CreateTime          *int    `json:"createTime,omitempty" name:"createTime"`                   // 队列创建时间(unix秒)
	LastModifyTime      *int    `json:"lastModifyTime,omitempty" name:"lastModifyTime"`           // 最后修改时间
	ActiveMsgNum        *int    `json:"activeMsgNum,omitempty" name:"activeMsgNum"`               // 在队列中处于 Active 状态（不处于被消费状态）的消息总数，为近似值
	InactiveMsgNum      *int    `json:"inactiveMsgNum,omitempty" name:"inactiveMsgNum"`           // 在队列中处于 Inactive 状态（正处于被消费状态）的消息总数，为近似值
	RewindSeconds       *int    `json:"rewindSeconds,omitempty" name:"rewindSeconds"`             // 队列是否开启回溯消息周期
	RewindMsgNum        *int    `json:"rewindmsgNum,omitempty" name:"rewindmsgNum"`               // 已调用 DelMsg 接口删除，但还在回溯保留时间内的消息数量
	MinMsgTime          *int    `json:"minMsgTime,omitempty" name:"minMsgTime"`                   // 消息最小未消费时间，单位为秒
}

// SetQueueAttributesRequest 修改队列属性请求
type SetQueueAttributesRequest struct {
	*tchttp.BaseRequest
	QueueName           *string `json:"queueName,omitempty" name:"queueName"`
	MaxMsgHeapNum       *int    `json:"maxMsgHeapNum,omitempty" name:"maxMsgHeapNum"`
	PollingWaitSeconds  *int    `json:"pollingWaitSeconds,omitempty" name:"pollingWaitSeconds"`
	VisibilityTimeout   *int    `json:"visibilityTimeout,omitempty" name:"visibilityTimeout"`
	MaxMsgSize          *int    `json:"maxMsgSize,omitempty" name:"maxMsgSize"`
	MsgRetentionSeconds *int    `json:"msgRetentionSeconds,omitempty" name:"msgRetentionSeconds"`
	RewindSeconds       *int    `json:"rewindSeconds,omitempty" name:"rewindSeconds"`
}

// SetQueueAttributesResponse 修改队列属性响应
type SetQueueAttributesResponse struct {
	*tchttp.BaseResponse
	Code                *int    `json:"code,omitempty" name:"code"`
	Message             *string `json:"message,omitempty" name:"message"`
	RequestId           *string `json:"requestId,omitempty" name:"requestId"`
	MaxMsgHeapNum       *int    `json:"maxMsgHeapNum,omitempty" name:"maxMsgHeapNum"`
	PollingWaitSeconds  *int    `json:"pollingWaitSeconds,omitempty" name:"pollingWaitSeconds"`
	VisibilityTimeout   *int    `json:"visibilityTimeout,omitempty" name:"visibilityTimeout"`
	MaxMsgSize          *int    `json:"maxMsgSize,omitempty" name:"maxMsgSize"`
	MsgRetentionSeconds *int    `json:"msgRetentionSeconds,omitempty" name:"msgRetentionSeconds"`
	RewindSeconds       *int    `json:"rewindSeconds,omitempty" name:"rewindSeconds"`
}

// DeleteQueueRequest 删除队列请求
type DeleteQueueRequest struct {
	*tchttp.BaseRequest
	QueueName *string `json:"queueName,omitempty" name:"queueName"`
}

// DeleteQueueResponse 删除队列响应
type DeleteQueueResponse struct {
	*tchttp.BaseResponse
	Code      *int    `json:"code,omitempty" name:"code"`
	Message   *string `json:"message,omitempty" name:"message"`
	RequestId *string `json:"requestId,omitempty" name:"requestId"`
}

// SendMessage 发送消息请求(队列)
type SendMessageRequest struct {
	*tchttp.BaseRequest
	QueueName    *string `json:"queueName,omitempty" name:"queueName"`
	MsgBody      *string `json:"msgBody,omitempty" name:"msgBody"`
	DelaySeconds *int    `json:"delaySeconds,omitempty" name:"delaySeconds"`
}

// SendMessageResponse 发送消息响应(队列)
type SendMessageResponse struct {
	*tchttp.BaseResponse
	Code      *int    `json:"code,omitempty" name:"code"`
	Message   *string `json:"message,omitempty" name:"message"`
	RequestId *string `json:"requestId,omitempty" name:"requestId"`
	MsgId     *string `json:"msgId,omitempty" name:"msgId"`
}

// BatchSendMessage 批量发送消息请求(队列)
type BatchSendMessageRequest struct {
	*tchttp.BaseRequest
	QueueName    *string   `json:"queueName,omitempty" name:"queueName"`
	MsgBody      []*string `json:"msgBody,omitempty" name:"msgBody"`
	DelaySeconds *int      `json:"delaySeconds,omitempty" name:"delaySeconds"`
}

// BatchSendMessageResponse 批量发送消息响应(响应)
type BatchSendMessageResponse struct {
	*tchttp.BaseResponse
	Code      *int    `json:"code,omitempty" name:"code"`
	Message   *string `json:"message,omitempty" name:"message"`
	RequestId *string `json:"requestId,omitempty" name:"requestId"`
	MsgList   []*struct {
		MsgId string `json:"msgId,omitempty" name:"msgId"`
	} `json:"msgList,omitempty" name:"msgList"`
}

// ReceiveMessage 消费消息请求(队列)
type ReceiveMessageRequest struct {
	*tchttp.BaseRequest
	QueueName          *string `json:"queueName,omitempty" name:"queueName"`
	PollingWaitSeconds *int    `json:"pollingWaitSeconds,omitempty" name:"pollingWaitSeconds"`
}

// ReceiveMessageResponse 消费消息响应(队列)
type ReceiveMessageResponse struct {
	*tchttp.BaseResponse
	Code             *int    `json:"code,omitempty" name:"code"`
	Message          *string `json:"message,omitempty" name:"message"`
	RequestId        *string `json:"requestId,omitempty" name:"requestId"`
	MsgId            *string `json:"msgId,omitempty" name:"msgId"`
	MsgBody          *string `json:"msgBody,omitempty" name:"msgBody"`
	ReceiptHandle    *string `json:"receiptHandle,omitempty" name:"receiptHandle"`       // 消息句柄
	EnqueueTime      *int    `json:"enqueueTime,omitempty" name:"enqueueTime"`           // 入队时间
	FirstDequeueTime *int    `json:"firstDequeueTime,omitempty" name:"firstDequeueTime"` // 第一次消费时间
	NextVisibleTime  *int    `json:"nextVisibleTime,omitempty" name:"nextVisibleTime"`   // 消息下次可见时间
	DequeueCount     *int    `json:"dequeueCount,omitempty" name:"dequeueCount"`         // 消费次数
}

// BatchReceiveMessage 批量消费消息请求(队列)
type BatchReceiveMessageRequest struct {
	*tchttp.BaseRequest
	QueueName          *string `json:"queueName,omitempty" name:"queueName"`
	NumOfMsg           *int    `json:"numOfMsg,omitempty" name:"numOfMsg"`
	PollingWaitSeconds *int    `json:"pollingWaitSeconds,omitempty" name:"pollingWaitSeconds"`
}

// BatchReceiveMessageResponse 批量消费消息响应(队列)
type BatchReceiveMessageResponse struct {
	*tchttp.BaseResponse
	Code        *int    `json:"code,omitempty" name:"code"`
	Message     *string `json:"message,omitempty" name:"message"`
	RequestId   *string `json:"requestId,omitempty" name:"requestId"`
	MsgInfoList []*struct {
		MsgId            *string `json:"msgId,omitempty" name:"msgId"`
		MsgBody          *string `json:"msgBody,omitempty" name:"msgBody"`
		ReceiptHandle    *string `json:"receiptHandle,omitempty" name:"receiptHandle"`       // 消息句柄
		EnqueueTime      *int    `json:"enqueueTime,omitempty" name:"enqueueTime"`           // 入队时间
		FirstDequeueTime *int    `json:"firstDequeueTime,omitempty" name:"firstDequeueTime"` // 第一次消费时间
		NextVisibleTime  *int    `json:"nextVisibleTime,omitempty" name:"nextVisibleTime"`   // 消息下次可见时间
		DequeueCount     *int    `json:"dequeueCount,omitempty" name:"dequeueCount"`         // 消费次数
	} `json:"msgInfoList,omitempty" name:"msgInfoList"`
}

// DeleteMessageRequest 删除消息请求
type DeleteMessageRequest struct {
	*tchttp.BaseRequest
	QueueName *string `json:"queueName,omitempty" name:"queueName"`
}

// DeleteMessageResponse 删除消息响应
type DeleteMessageResponse struct {
	*tchttp.BaseResponse
	Code      *int    `json:"code,omitempty" name:"code"`
	Message   *string `json:"message,omitempty" name:"message"`
	RequestId *string `json:"requestId,omitempty" name:"requestId"`
}

// BatchDeleteMessage 批量删除消息请求(队列)
type BatchDeleteMessageRequest struct {
	*tchttp.BaseRequest
	QueueName     *string   `json:"queueName,omitempty" name:"queueName"`
	ReceiptHandle []*string `json:"receiptHandle,omitempty" name:"receiptHandle"`
}

// BatchDeleteMessageResponse 批量删除消息响应(队列)
type BatchDeleteMessageResponse struct {
	*tchttp.BaseResponse
	Code      *int    `json:"code,omitempty" name:"code"`
	Message   *string `json:"message,omitempty" name:"message"`
	RequestId *string `json:"requestId,omitempty" name:"requestId"`
	ErrorList []*struct {
		Code          *int    `json:"code,omitempty" name:"code"`
		Message       *string `json:"message,omitempty" name:"message"`
		ReceiptHandle *string `json:"receiptHandle,omitempty" name:"receiptHandle"`
	} `json:"msgInfoList,omitempty" name:"msgInfoList"`
}

// =================================================================================

// CreateTopicRequest 创建主题请求
type CreateTopicRequest struct {
	*tchttp.BaseRequest
	// 主题名字,主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)
	TopicName *string `json:"topicName,omitempty" name:"topicName"`
	// 消息最大长度,取值范围1024 - 65536 Byte（即1 - 64K）
	MaxMsgSize *int `json:"maxMsgSize,omitempty" name:"maxMsgSize"`
	// 用于指定主题的消息匹配策略： filterType =1 或为空， 表示该主题下所有订阅使用 filterTag 标签过滤； filterType =2 表示用户使用 bindingKey 过滤。
	FilterType *int `json:"filterType,omitempty" name:"filterType"`
}

// CreateTopicResponse 创建主题响应
type CreateTopicResponse struct {
	*tchttp.BaseResponse
	// 0：表示成功，others：错误
	Code *int `json:"code,omitempty" name:"code"`
	// 错误提示信息
	Message *string `json:"message,omitempty" name:"message"`
	// 服务器生成的请求 ID
	RequestId *string `json:"requestId,omitempty" name:"requestId"`
	// 主题标识ID
	TopicId *string `json:"topicId,omitempty" name:"topicId"`
}

// SetTopicAttributesRequest 修改主题属性请求
type SetTopicAttributesRequest struct {
	*tchttp.BaseRequest
	TopicName  *string `json:"topicName,omitempty" name:"topicName"`
	MaxMsgSize *int    `json:"maxMsgSize,omitempty" name:"maxMsgSize"`
}

// SetTopicAttributesResponse 修改主题属性响应
type SetTopicAttributesResponse struct {
	*tchttp.BaseResponse
	Code      *int    `json:"code,omitempty" name:"code"`
	Message   *string `json:"message,omitempty" name:"message"`
	RequestId *string `json:"requestId,omitempty" name:"requestId"`
}

// ListTopicRequest 获取主题列表请求
type ListTopicRequest struct {
	*tchttp.BaseRequest
	SearchWord *string `json:"serarchWord,omitempty" name:"searchWord"`
	Offset     *int    `json:"offset,omitempty" name:"offset"`
	Limit      *int    `json:"limit,omitempty" name:"limit"`
}

// ListTopicResponse 获取主题列表响应
type ListTopicResponse struct {
	*tchttp.BaseResponse
	Code       *int     `json:"code,omitempty" name:"code"`
	Message    *string  `json:"message,omitempty" name:"message"`
	RequestId  *string  `json:"requestId,omitempty" name:"requestId"`
	TotalCount *int     `json:"totalCount,omitempty" name:"totalCount"`
	TopicList  []*Topic `json:"topicList,omitempty" name:"topicList"`
}

// Topic 主题
type Topic struct {
	TopicId   string `json:"topicId,omitempty" name:"topicId"`
	TopicName string `json:"topicName,omitempty" name:"topicName"`
}

// GetTopicAttributesRequest 获取主题属性请求
type GetTopicAttributesRequest struct {
	*tchttp.BaseRequest
	TopicName *string `json:"topicName,omitempty" name:"topicName"`
}

// GetTopicAttributesResponse 获取主题属性响应
type GetTopicAttributesResponse struct {
	*tchttp.BaseResponse
	Code                *int    `json:"code,omitempty" name:"code"`
	Message             *string `json:"message,omitempty" name:"message"`
	RequestId           *string `json:"requestId,omitempty" name:"requestId"`
	MsgCount            *int    `json:"msgCount,omitempty" name:"msgCount"`
	MaxMsgSize          *int    `json:"maxMsgSize,omitempty" name:"maxMsgSize"`
	MsgRetentionSeconds *int    `json:"msgRetentionSeconds,omitempty" name:"maxMsgSize"` // 消息在主题中最长存活时间，单位为秒。固定为一天（86400秒），该属性不能修改
	CreateTime          *int    `json:"createTime,omitempty" name:"createTime"`          // 主题创建时间(unix秒)
	LastModifyTime      *int    `json:"lastModifyTime,omitempty" name:"lastModifyTime"`  // 最后修改时间
	FilterType          *int    `json:"filterType,omitempty" name:"filterType"`          // 消息过滤类型
}

// DeleteTopicRequest 删除主题请求
type DeleteTopicRequest struct {
	*tchttp.BaseRequest
	TopicName *string `json:"topicName,omitempty" name:"topicName"`
}

// DeleteTopicResponse 删除主题响应
type DeleteTopicResponse struct {
	*tchttp.BaseResponse
	Code      *int    `json:"code,omitempty" name:"code"`
	Message   *string `json:"message,omitempty" name:"message"`
	RequestId *string `json:"requestId,omitempty" name:"requestId"`
}

// PublishMessage 发布消息请求
type PublishMessageRequest struct {
	*tchttp.BaseRequest
	TopicName *string `json:"topicName,omitempty" name:"topicName"`
	// 消息正文。至少1Byte，最大长度受限于设置的主题消息最大长度属性
	MsgBody *string `json:"msgBody,omitempty" name:"msgBody"`
	// 消息过滤标签
	MsgTag []*string `json:"msgTag,omitempty" name:"msgTag"`
	// 路由键
	RoutingKey *string `json:"routingKey,omitempty" name:"routingKey"`
}

// PublishMessageResponse 发布消息响应
type PublishMessageResponse struct {
	*tchttp.BaseResponse
	// 0：表示成功，others：错误
	Code *int `json:"code,omitempty" name:"code"`
	// 错误提示信息
	Message *string `json:"message,omitempty" name:"message"`
	// 服务器生成的请求 ID
	RequestId *string `json:"requestId,omitempty" name:"requestId"`
	// 消息标识ID
	MsgId *string `json:"msgId,omitempty" name:"msgId"`
}

// BatchPublishMessage 批量发布消息请求
type BatchPublishMessageRequest struct {
	*tchttp.BaseRequest
	TopicName *string `json:"topicName,omitempty" name:"topicName"`
	// 消息正文,表示这一批量中的一条消息,目前批量消息数量不能超过16条
	MsgBody []*string `json:"msgBody,omitempty" name:"msgBody"`
	// 消息过滤标签,标签数量不能超过5个，每个标签不超过16个字符
	MsgTag []*string `json:"msgTag,omitempty" name:"msgTag"`
	// 路由键
	RoutingKey *string `json:"routingKey,omitempty" name:"routingKey"`
}

// BatchPublishMessageResponse 批量发布消息响应
type BatchPublishMessageResponse struct {
	*tchttp.BaseResponse
	// 0：表示成功，others：错误
	Code *int `json:"code,omitempty" name:"code"`
	// 错误提示信息
	Message *string `json:"message,omitempty" name:"message"`
	// 服务器生成的请求 ID
	RequestId *string `json:"requestId,omitempty" name:"requestId"`
	// 消息标识ID
	MsgList []*struct {
		MsgId string `json:"msgId,omitempty" name:"msgId"`
	} `json:"msgList,omitempty" name:"msgList"`
}

// Subscribe 创建订阅
type SubscribeRequest struct {
	*tchttp.BaseRequest
	TopicName           *string   `json:"topicName,omitempty" name:"topicName"`
	SubscriptionName    *string   `json:"subscriptionName,omitempty" name:"subscriptionName"` // 订阅名字
	Protocol            *string   `json:"protocol,omitempty" name:"protocol"`                 // 订阅协议,目前支持两种协议：HTTP、Queue。使用 HTTP 协议，用户需自己搭建接受消息的 Web Server。使用 Queue，消息会自动推送到 CMQ Queue，用户可以并发地拉取消息
	Endpoint            *string   `json:"endpoint,omitempty" name:"endpoint"`                 // 接收通知的 endpoint，根据协议 protocol 区分：对于 HTTP，endpoint 必须以 “http://” 开头，host 可以是域名或 IP；对于 Queue，则填 queueName
	NotifyStrategy      *string   `json:"notifyStrategy,omitempty" name:"notifyStrategy"`
	NotifyContentFormat *string   `json:"notifyContentFormat,omitempty" name:"notifyContentFormat"` // 推送内容的格式。取值：（1）JSON。（2）SIMPLIFIED，即 raw 格式。如果 protocol 是 Queue，则取值必须为 SIMPLIFIED。如果 protocol 是 HTTP，两个值均可以，默认值是 JSON
	FilterTag           []*string `json:"filterTag,omitempty" name:"filterTag"`                     // 消息标签
	BindingKey          []*string `json:"bindingKey,omitempty" name:"bindingKey"`                   // 绑定键
}

// SubscribeResponse 创建订阅响应
type SubscribeResponse struct {
	*tchttp.BaseResponse
	Code      *int    `json:"code,omitempty" name:"code"`
	Message   *string `json:"message,omitempty" name:"message"`
	RequestId *string `json:"requestId,omitempty" name:"requestId"`
}

// ListSubscriptionByTopicRequest 获取订阅列表请求
type ListSubscriptionByTopicRequest struct {
	*tchttp.BaseRequest
	TopicName  *string `json:"topicName,omitempty" name:"topicName"`
	SearchWord *string `json:"serarchWord,omitempty" name:"searchWord"`
	Offset     *int    `json:"offset,omitempty" name:"offset"`
	Limit      *int    `json:"limit,omitempty" name:"limit"`
}

// ListSubscriptionByTopicResponse 获取订阅列表响应
type ListSubscriptionByTopicResponse struct {
	*tchttp.BaseResponse
	Code             *int    `json:"code,omitempty" name:"code"`
	Message          *string `json:"message,omitempty" name:"message"`
	RequestId        *string `json:"requestId,omitempty" name:"requestId"`
	TotalCount       *int    `json:"totalCount,omitempty" name:"totalCount"` // 	用户帐号下本次请求返回的主题总数，而非分页后本页获取的主题数量
	SubscriptionList []*struct {
		SubscriptionId   *string `json:"subscriptionId,omitempty" name:"subscriptionId"`
		SubscriptionName *string `json:"subscriptionName,omitempty" name:"subscriptionName"`
		Protocol         *string `json:"protocol,omitempty" name:"protocol"`
		Endpoint         *string `json:"endpoint,omitempty" name:"endpoint"`
	} `json:"subscriptionList,omitempty" name:"subscriptionList"` // 主题列表信息
}

// SetSubscriptionAttributesRequest 修改订阅属性请求
type SetSubscriptionAttributesRequest struct {
	*tchttp.BaseRequest
	TopicName           *string   `json:"topicName,omitempty" name:"topicName"`                     // 主题名字
	SubscriptionName    *string   `json:"subscriptionName,omitempty" name:"subscriptionName"`       // 订阅名字
	NotifyStrategy      *string   `json:"notifyStrategy,omitempty" name:"notifyStrategy"`           // 重试策略
	NotifyContentFormat *string   `json:"notifyContentFormat,omitempty" name:"notifyContentFormat"` // 推送内容的格式：1JSON,2SIMPLIFIED
	FilterTag           []*string `json:"filterTag,omitempty" name:"filterTag"`                     // 消息标签
	BindingKey          []*string `json:"bindingKey,omitempty" name:"bindingKey"`                   // 绑定键
}

// SetSubcriptionAttributesResponse 修订订阅属性响应
type SetSubscriptionAttributesResponse struct {
	*tchttp.BaseResponse
	Code      *int    `json:"code,omitempty" name:"code"`
	Message   *string `json:"message,omitempty" name:"message"`
	RequestId *string `json:"requestId,omitempty" name:"requestId"`
}

// GetSubscriptionAttributesRequest 获取订阅属性请求
type GetSubscriptionAttributesRequest struct {
	*tchttp.BaseRequest
	TopicName        *string `json:"topicName,omitempty" name:"topicName"`               // 主题名字
	SubscriptionName *string `json:"subscriptionName,omitempty" name:"subscriptionName"` // 订阅名字
}

// GetSubcriptionAttributesResponse 获取订阅属性响应
type GetSubscriptionAttributesResponse struct {
	*tchttp.BaseResponse
	Code                *int      `json:"code,omitempty" name:"code"`
	Message             *string   `json:"message,omitempty" name:"message"`
	RequestId           *string   `json:"requestId,omitempty" name:"requestId"`
	TopicOwner          *string   `json:"topicOwner,omitempty" name:"topicOwner"` // 订阅拥有者的APPID
	MsgCount            *int      `json:"msgCount,omitempty" name:"msgCount"`
	Protocol            *string   `json:"protocol,omitempty" name:"protocol"` // 订阅协议
	Endpoint            *string   `json:"endpoint,omitempty" name:"endpoint"` // 接收通知的 endpoint
	NotifyStrategy      *string   `json:"notifyStrategy,omitempty" name:"notifyStrategy"`
	NotifyContentFormat *string   `json:"notifyContentFormat,omitempty" name:"notifyContentFormat"` // 推送内容的格式
	CreateTime          *int      `json:"createTime,omitempty" name:"createTime"`                   // 主题创建时间(unix秒)
	LastModifyTime      *int      `json:"lastModifyTime,omitempty" name:"lastModifyTime"`           // 最后修改时间
	BindingKey          []*string `json:"bindingKey,omitempty" name:"bindingKey"`                   // 表示订阅接收消息的过滤策略
}

// Unsubscribe 删除订阅请求
type UnsubscribeRequest struct {
	*tchttp.BaseRequest
	TopicName        *string `json:"topicName,omitempty" name:"topicName"`
	SubscriptionName *string `json:"subscriptionName,omitempty" name:"subscriptionName"` // 订阅名字
}

// UnsubscribeResponse 删除订阅响应
type UnsubscribeResponse struct {
	*tchttp.BaseResponse
	Code      *int    `json:"code,omitempty" name:"code"`
	Message   *string `json:"message,omitempty" name:"message"`
	RequestId *string `json:"requestId,omitempty" name:"requestId"`
}
