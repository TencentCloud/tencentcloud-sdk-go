package cmq

// 创建主题

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

// SetTopicAttributes 修改主题属性
// func (c *Client) SetTopicAttributes(request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
// 	if request == nil {
// 		request = NewCreateTopicRequest()
// 	}
// 	response = NewCreateTopicResponse()
// 	err = c.Send(request, response)
// 	return
// }

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
