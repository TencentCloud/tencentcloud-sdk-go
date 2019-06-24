package cns

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-05-28"

type Client struct {
	common.Client
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
	client = &Client{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

func NewDomainListRequest() (request *DomainListRequest) {
	request = &DomainListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cns", APIVersion, "DomainList")

	return
}

func NewDomainListResponse() (response *DomainListResponse) {
	response = &DomainListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DomainList(request *DomainListRequest) (response *DomainListResponse, err error) {
	if request == nil {
		request = NewDomainListRequest()
	}
	response = NewDomainListResponse()
	err = c.Send(request, response)

	return
}

func NewDomainCreateRequest() (request *DomainCreateRequest) {
	request = &DomainCreateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cns", APIVersion, "DomainCreate")

	return
}

func NewDomainCreateResponse() (response *DomainCreateResponse) {
	response = &DomainCreateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DomainCreate(request *DomainCreateRequest) (response *DomainCreateResponse, err error) {
	if request == nil {
		request = NewDomainCreateRequest()
	}
	response = NewDomainCreateResponse()
	err = c.Send(request, response)

	return
}

func NewSetDomainStatusRequest() (request *SetDomainStatusRequest) {
	request = &SetDomainStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cns", APIVersion, "SetDomainStatus")

	return
}

func NewSetDomainStatusResponse() (response *SetDomainStatusResponse) {
	response = &SetDomainStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) SetDomainStatus(request *SetDomainStatusRequest) (response *SetDomainStatusResponse, err error) {
	if request == nil {
		request = NewSetDomainStatusRequest()
	}
	response = NewSetDomainStatusResponse()
	err = c.Send(request, response)

	return
}

func NewDomainDeleteRequest() (request *DomainDeleteRequest) {
	request = &DomainDeleteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cns", APIVersion, "DomainDelete")

	return
}

func NewDomainDeleteResponse() (response *DomainDeleteResponse) {
	response = &DomainDeleteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DomainDelete(request *DomainDeleteRequest) (response *DomainDeleteResponse, err error) {
	if request == nil {
		request = NewDomainDeleteRequest()
	}
	response = NewDomainDeleteResponse()
	err = c.Send(request, response)

	return
}

func NewRecordCreateRequest() (request *RecordCreateRequest) {
	request = &RecordCreateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cns", APIVersion, "RecordCreate")

	return
}

func NewRecordCreateResponse() (response *RecordCreateResponse) {
	response = &RecordCreateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) RecordCreate(request *RecordCreateRequest) (response *RecordCreateResponse, err error) {
	if request == nil {
		request = NewRecordCreateRequest()
	}
	response = NewRecordCreateResponse()
	err = c.Send(request, response)

	return
}

func NewRecordStatusRequest() (request *RecordStatusRequest) {
	request = &RecordStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cns", APIVersion, "RecordStatus")

	return
}

func NewRecordStatusResponse() (response *RecordStatusResponse) {
	response = &RecordStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) RecordStatus(request *RecordStatusRequest) (response *RecordStatusResponse, err error) {
	if request == nil {
		request = NewRecordStatusRequest()
	}
	response = NewRecordStatusResponse()
	err = c.Send(request, response)

	return
}

func NewRecordModifyRequest() (request *RecordModifyRequest) {
	request = &RecordModifyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cns", APIVersion, "RecordModify")

	return
}

func NewRecordModifyResponse() (response *RecordModifyResponse) {
	response = &RecordModifyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) RecordModify(request *RecordModifyRequest) (response *RecordModifyResponse, err error) {
	if request == nil {
		request = NewRecordModifyRequest()
	}
	response = NewRecordModifyResponse()
	err = c.Send(request, response)

	return
}

func NewRecordListRequest() (request *RecordListRequest) {
	request = &RecordListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cns", APIVersion, "RecordList")

	return
}

func NewRecordListResponse() (response *RecordListResponse) {
	response = &RecordListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) RecordList(request *RecordListRequest) (response *RecordListResponse, err error) {
	if request == nil {
		request = NewRecordListRequest()
	}
	response = NewRecordListResponse()
	err = c.Send(request, response)

	return
}

func NewRecordDeleteRequest() (request *RecordDeleteRequest) {
	request = &RecordDeleteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cns", APIVersion, "RecordDelete")

	return
}

func NewRecordDeleteResponse() (response *RecordDeleteResponse) {
	response = &RecordDeleteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) RecordDelete(request *RecordDeleteRequest) (response *RecordDeleteResponse, err error) {
	if request == nil {
		request = NewRecordDeleteRequest()
	}
	response = NewRecordDeleteResponse()
	err = c.Send(request, response)

	return
}
