package v2

// 主题相关接口
// =================================================================================

// 创建主题 CreateTopic

func NewCreateTopicRequest() (request *CreateTopicRequest) {
	return &CreateTopicRequest{BaseRequest: newCmqBaseRequest("CreateTopic")}
}

func NewCreateTopicResponse() (response *CreateTopicResponse) {
	return &CreateTopicResponse{BaseResponse: newCmqBaseResponse()}
}

// CreateTopic 创建主题
func (c *Client) CreateTopic(request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
	if request == nil {
		request = NewCreateTopicRequest()
	}
	response = NewCreateTopicResponse()
	err = c.Send(request, response)
	return
}

func NewSetTopicAttributesRequest() (request *SetTopicAttributesRequest) {
	return &SetTopicAttributesRequest{BaseRequest: newCmqBaseRequest("SetTopicAttributes")}
}

func NewSetTopicAttributesResponse() (response *SetTopicAttributesResponse) {
	return &SetTopicAttributesResponse{BaseResponse: newCmqBaseResponse()}
}

// SetTopicAttributes 修改主题属性
func (c *Client) SetTopicAttributes(request *SetTopicAttributesRequest) (response *SetTopicAttributesResponse, err error) {
	if request == nil {
		request = NewSetTopicAttributesRequest()
	}
	response = NewSetTopicAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewListTopicRequest() (request *ListTopicRequest) {
	return &ListTopicRequest{BaseRequest: newCmqBaseRequest("ListTopic")}
}

func NewListTopicResponse() (response *ListTopicResponse) {
	return &ListTopicResponse{BaseResponse: newCmqBaseResponse()}
}

// ListTopic 获取主题列表
func (c *Client) ListTopic(request *ListTopicRequest) (response *ListTopicResponse, err error) {
	if request == nil {
		request = NewListTopicRequest()
	}
	response = NewListTopicResponse()
	err = c.Send(request, response)
	return
}

func NewGetTopicAttributesRequest() (request *GetTopicAttributesRequest) {
	return &GetTopicAttributesRequest{BaseRequest: newCmqBaseRequest("GetTopicAttributes")}
}

func NewGetTopicAttributesResponse() (response *GetTopicAttributesResponse) {
	return &GetTopicAttributesResponse{BaseResponse: newCmqBaseResponse()}
}

// GetTopicAttributes 获取主题属性
func (c *Client) GetTopicAttributes(request *GetTopicAttributesRequest) (response *GetTopicAttributesResponse, err error) {
	if request == nil {
		request = NewGetTopicAttributesRequest()
	}
	response = NewGetTopicAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTopicRequest() (request *DeleteTopicRequest) {
	return &DeleteTopicRequest{BaseRequest: newCmqBaseRequest("DeleteTopic")}
}

func NewDeleteTopicResponse() (response *DeleteTopicResponse) {
	return &DeleteTopicResponse{BaseResponse: newCmqBaseResponse()}
}

// DeleteTopic 删除主题
func (c *Client) DeleteTopic(request *DeleteTopicRequest) (response *DeleteTopicResponse, err error) {
	if request == nil {
		request = NewDeleteTopicRequest()
	}
	response = NewDeleteTopicResponse()
	err = c.Send(request, response)
	return
}

// 消息相关接口
// =================================================================================

// 发送消息

func NewPublishMessageRequest() (request *PublishMessageRequest) {
	return &PublishMessageRequest{BaseRequest: newCmqBaseRequest("PublishMessage")}
}

func NewPublishMessageResponse() (response *PublishMessageResponse) {
	return &PublishMessageResponse{BaseResponse: newCmqBaseResponse()}
}

// PublishMessage 发布消息
func (c *Client) PublishMessage(request *PublishMessageRequest) (response *PublishMessageResponse, err error) {
	if request == nil {
		request = NewPublishMessageRequest()
	}
	response = NewPublishMessageResponse()
	err = c.Send(request, response)
	return
}

// 批量发送消息
func NewBatchPublishMessageRequest() (request *BatchPublishMessageRequest) {
	return &BatchPublishMessageRequest{BaseRequest: newCmqBaseRequest("BatchPublishMessage")}
}

func NewBatchPublishMessageResponse() (response *BatchPublishMessageResponse) {
	return &BatchPublishMessageResponse{BaseResponse: newCmqBaseResponse()}
}

// BatchPublishMessage 批量发布消息
func (c *Client) BatchPublishMessage(request *BatchPublishMessageRequest) (response *BatchPublishMessageResponse, err error) {
	if request == nil {
		request = NewBatchPublishMessageRequest()
	}
	response = NewBatchPublishMessageResponse()
	err = c.Send(request, response)
	return
}

// 订阅相关接口
// =================================================================================

// Subscribe 创建订阅
func NewSubscribeRequest() (request *SubscribeRequest) {
	return &SubscribeRequest{BaseRequest: newCmqBaseRequest("Subscribe")}
}

func NewSubscribeResponse() (response *SubscribeResponse) {
	return &SubscribeResponse{BaseResponse: newCmqBaseResponse()}
}

// Subscribe 创建订阅
func (c *Client) Subscribe(request *SubscribeRequest) (response *SubscribeResponse, err error) {
	if request == nil {
		request = NewSubscribeRequest()
	}
	response = NewSubscribeResponse()
	err = c.Send(request, response)
	return
}

func NewListSubscriptionByTopicRequest() (request *ListSubscriptionByTopicRequest) {
	return &ListSubscriptionByTopicRequest{BaseRequest: newCmqBaseRequest("ListSubscriptionByTopic")}
}

func NewListSubscriptionByTopicResponse() (response *ListSubscriptionByTopicResponse) {
	return &ListSubscriptionByTopicResponse{BaseResponse: newCmqBaseResponse()}
}

// ListSubscriptionByTopic 获取订阅列表
func (c *Client) ListSubscriptionByTopic(request *ListSubscriptionByTopicRequest) (response *ListSubscriptionByTopicResponse, err error) {
	if request == nil {
		request = NewListSubscriptionByTopicRequest()
	}
	response = NewListSubscriptionByTopicResponse()
	err = c.Send(request, response)
	return
}

func NewSetSubscriptionAttributesRequest() (request *SetSubscriptionAttributesRequest) {
	return &SetSubscriptionAttributesRequest{BaseRequest: newCmqBaseRequest("SetSubscriptionAttributes")}
}

func NewSetSubscriptionAttributesResponse() (response *SetSubscriptionAttributesResponse) {
	return &SetSubscriptionAttributesResponse{BaseResponse: newCmqBaseResponse()}
}

// SetSubscriptionAttributes 修改订阅属性
func (c *Client) SetSubscriptionAttributes(request *SetSubscriptionAttributesRequest) (response *SetSubscriptionAttributesResponse, err error) {
	if request == nil {
		request = NewSetSubscriptionAttributesRequest()
	}
	response = NewSetSubscriptionAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewGetSubscriptionAttributesRequest() (request *GetSubscriptionAttributesRequest) {
	return &GetSubscriptionAttributesRequest{BaseRequest: newCmqBaseRequest("GetSubscriptionAttributes")}
}

func NewGetSubscriptionAttributesResponse() (response *GetSubscriptionAttributesResponse) {
	return &GetSubscriptionAttributesResponse{BaseResponse: newCmqBaseResponse()}
}

// GetSubscriptionAttributes 获取订阅属性
func (c *Client) GetSubscriptionAttributes(request *GetSubscriptionAttributesRequest) (response *GetSubscriptionAttributesResponse, err error) {
	if request == nil {
		request = NewGetSubscriptionAttributesRequest()
	}
	response = NewGetSubscriptionAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewUnsubscribeRequest() (request *UnsubscribeRequest) {
	return &UnsubscribeRequest{BaseRequest: newCmqBaseRequest("Unsubscribe")}
}

func NewUnsubscribeResponse() (response *UnsubscribeResponse) {
	return &UnsubscribeResponse{BaseResponse: newCmqBaseResponse()}
}

// Unsubscribe 删除订阅
func (c *Client) Unsubscribe(request *UnsubscribeRequest) (response *UnsubscribeResponse, err error) {
	if request == nil {
		request = NewUnsubscribeRequest()
	}
	response = NewUnsubscribeResponse()
	err = c.Send(request, response)
	return
}
