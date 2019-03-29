package v2

func NewCreateQueueRequest() (request *CreateQueueRequest) {
	return &CreateQueueRequest{BaseRequest: newCmqBaseRequest("CreateQueue")}
}

func NewCreateQueueResponse() (response *CreateQueueResponse) {
	return &CreateQueueResponse{BaseResponse: newCmqBaseResponse()}
}

// CreateQueue 创建队列
func (c *Client) CreateQueue(request *CreateQueueRequest) (response *CreateQueueResponse, err error) {
	if request == nil {
		request = NewCreateQueueRequest()
	}
	response = NewCreateQueueResponse()
	err = c.Send(request, response)
	return
}

func NewListQueueRequest() (request *ListQueueRequest) {
	return &ListQueueRequest{BaseRequest: newCmqBaseRequest("ListQueue")}
}

func NewListQueueResponse() (response *ListQueueResponse) {
	return &ListQueueResponse{BaseResponse: newCmqBaseResponse()}
}

// ListQueue 获取队列列表
func (c *Client) ListQueue(request *ListQueueRequest) (response *ListQueueResponse, err error) {
	if request == nil {
		request = NewListQueueRequest()
	}
	response = NewListQueueResponse()
	err = c.Send(request, response)
	return
}

func NewGetQueueAttributesRequest() (request *GetQueueAttributesRequest) {
	return &GetQueueAttributesRequest{BaseRequest: newCmqBaseRequest("GetQueueAttributes")}
}

func NewGetQueueAttributesResponse() (response *GetQueueAttributesResponse) {
	return &GetQueueAttributesResponse{BaseResponse: newCmqBaseResponse()}
}

// GetQueueAttributes 修改队列属性
func (c *Client) GetQueueAttributes(request *GetQueueAttributesRequest) (response *GetQueueAttributesResponse, err error) {
	if request == nil {
		request = NewGetQueueAttributesRequest()
	}
	response = NewGetQueueAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewSetQueueAttributesRequest() (request *SetQueueAttributesRequest) {
	return &SetQueueAttributesRequest{BaseRequest: newCmqBaseRequest("SetQueueAttributes")}
}

func NewSetQueueAttributesResponse() (response *SetQueueAttributesResponse) {
	return &SetQueueAttributesResponse{BaseResponse: newCmqBaseResponse()}
}

// SetQueueAttributes 修改队列属性
func (c *Client) SetQueueAttributes(request *SetQueueAttributesRequest) (response *SetQueueAttributesResponse, err error) {
	if request == nil {
		request = NewSetQueueAttributesRequest()
	}
	response = NewSetQueueAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteQueueRequest() (request *DeleteQueueRequest) {
	return &DeleteQueueRequest{BaseRequest: newCmqBaseRequest("DeleteQueue")}
}

func NewDeleteQueueResponse() (response *DeleteQueueResponse) {
	return &DeleteQueueResponse{BaseResponse: newCmqBaseResponse()}
}

// DeleteQueue 删除队列
func (c *Client) DeleteQueue(request *DeleteQueueRequest) (response *DeleteQueueResponse, err error) {
	if request == nil {
		request = NewDeleteQueueRequest()
	}
	response = NewDeleteQueueResponse()
	err = c.Send(request, response)
	return
}

func NewSendMessageRequest() (request *SendMessageRequest) {
	return &SendMessageRequest{BaseRequest: newCmqBaseRequest("SendMessage")}
}

func NewSendMessageResponse() (response *SendMessageResponse) {
	return &SendMessageResponse{BaseResponse: newCmqBaseResponse()}
}

// SendMessage 发送消息
func (c *Client) SendMessage(request *SendMessageRequest) (response *SendMessageResponse, err error) {
	if request == nil {
		request = NewSendMessageRequest()
	}
	response = NewSendMessageResponse()
	err = c.Send(request, response)
	return
}

func NewBatchSendMessageRequest() (request *BatchSendMessageRequest) {
	return &BatchSendMessageRequest{BaseRequest: newCmqBaseRequest("BatchSendMessage")}
}

func NewBatchSendMessageResponse() (response *BatchSendMessageResponse) {
	return &BatchSendMessageResponse{BaseResponse: newCmqBaseResponse()}
}

// BatchSendMessage 批量发送消息
func (c *Client) BatchSendMessage(request *BatchSendMessageRequest) (response *BatchSendMessageResponse, err error) {
	if request == nil {
		request = NewBatchSendMessageRequest()
	}
	response = NewBatchSendMessageResponse()
	err = c.Send(request, response)
	return
}

func NewReceiveMessageRequest() (request *ReceiveMessageRequest) {
	return &ReceiveMessageRequest{BaseRequest: newCmqBaseRequest("ReceiveMessage")}
}

func NewReceiveMessageResponse() (response *ReceiveMessageResponse) {
	return &ReceiveMessageResponse{BaseResponse: newCmqBaseResponse()}
}

// ReceiveMessage 接收消息
func (c *Client) ReceiveMessage(request *ReceiveMessageRequest) (response *ReceiveMessageResponse, err error) {
	if request == nil {
		request = NewReceiveMessageRequest()
	}
	response = NewReceiveMessageResponse()
	err = c.Send(request, response)
	return
}

func NewBatchReceiveMessageRequest() (request *BatchReceiveMessageRequest) {
	return &BatchReceiveMessageRequest{BaseRequest: newCmqBaseRequest("BatchReceiveMessage")}
}

func NewBatchReceiveMessageResponse() (response *BatchReceiveMessageResponse) {
	return &BatchReceiveMessageResponse{BaseResponse: newCmqBaseResponse()}
}

// BatchReceiveMessage 批量接收消息
func (c *Client) BatchReceiveMessage(request *BatchReceiveMessageRequest) (response *BatchReceiveMessageResponse, err error) {
	if request == nil {
		request = NewBatchReceiveMessageRequest()
	}
	response = NewBatchReceiveMessageResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMessageRequest() (request *DeleteMessageRequest) {
	return &DeleteMessageRequest{BaseRequest: newCmqBaseRequest("DeleteMessage")}
}

func NewDeleteMessageResponse() (response *DeleteMessageResponse) {
	return &DeleteMessageResponse{BaseResponse: newCmqBaseResponse()}
}

// DeleteMessage  删除消息
func (c *Client) DeleteMessage(request *DeleteMessageRequest) (response *DeleteMessageResponse, err error) {
	if request == nil {
		request = NewDeleteMessageRequest()
	}
	response = NewDeleteMessageResponse()
	err = c.Send(request, response)
	return
}

func NewBatchDeleteMessageRequest() (request *BatchDeleteMessageRequest) {
	return &BatchDeleteMessageRequest{BaseRequest: newCmqBaseRequest("BatchDeleteMessage")}
}

func NewBatchDeleteMessageResponse() (response *BatchDeleteMessageResponse) {
	return &BatchDeleteMessageResponse{BaseResponse: newCmqBaseResponse()}
}

// BatchDeleteMessage  批量删除消息
func (c *Client) BatchDeleteMessage(request *BatchDeleteMessageRequest) (response *BatchDeleteMessageResponse, err error) {
	if request == nil {
		request = NewBatchDeleteMessageRequest()
	}
	response = NewBatchDeleteMessageResponse()
	err = c.Send(request, response)
	return
}
