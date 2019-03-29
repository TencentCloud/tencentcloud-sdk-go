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
