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

package v20190819

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-08-19"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewCreateAclRequest() (request *CreateAclRequest) {
    request = &CreateAclRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "CreateAcl")
    return
}

func NewCreateAclResponse() (response *CreateAclResponse) {
    response = &CreateAclResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAcl
// 添加 ACL 策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateAcl(request *CreateAclRequest) (response *CreateAclResponse, err error) {
    if request == nil {
        request = NewCreateAclRequest()
    }
    response = NewCreateAclResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstancePreRequest() (request *CreateInstancePreRequest) {
    request = &CreateInstancePreRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "CreateInstancePre")
    return
}

func NewCreateInstancePreResponse() (response *CreateInstancePreResponse) {
    response = &CreateInstancePreResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateInstancePre
// 创建实例(预付费包年包月)
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateInstancePre(request *CreateInstancePreRequest) (response *CreateInstancePreResponse, err error) {
    if request == nil {
        request = NewCreateInstancePreRequest()
    }
    response = NewCreateInstancePreResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePartitionRequest() (request *CreatePartitionRequest) {
    request = &CreatePartitionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "CreatePartition")
    return
}

func NewCreatePartitionResponse() (response *CreatePartitionResponse) {
    response = &CreatePartitionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePartition
// 本接口用于增加主题中的分区
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreatePartition(request *CreatePartitionRequest) (response *CreatePartitionResponse, err error) {
    if request == nil {
        request = NewCreatePartitionRequest()
    }
    response = NewCreatePartitionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTopicRequest() (request *CreateTopicRequest) {
    request = &CreateTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "CreateTopic")
    return
}

func NewCreateTopicResponse() (response *CreateTopicResponse) {
    response = &CreateTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTopic
// 创建ckafka主题
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateTopic(request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
    if request == nil {
        request = NewCreateTopicRequest()
    }
    response = NewCreateTopicResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTopicIpWhiteListRequest() (request *CreateTopicIpWhiteListRequest) {
    request = &CreateTopicIpWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "CreateTopicIpWhiteList")
    return
}

func NewCreateTopicIpWhiteListResponse() (response *CreateTopicIpWhiteListResponse) {
    response = &CreateTopicIpWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTopicIpWhiteList
// 创建主题ip白名单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateTopicIpWhiteList(request *CreateTopicIpWhiteListRequest) (response *CreateTopicIpWhiteListResponse, err error) {
    if request == nil {
        request = NewCreateTopicIpWhiteListRequest()
    }
    response = NewCreateTopicIpWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserRequest() (request *CreateUserRequest) {
    request = &CreateUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "CreateUser")
    return
}

func NewCreateUserResponse() (response *CreateUserResponse) {
    response = &CreateUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateUser
// 添加用户
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateUser(request *CreateUserRequest) (response *CreateUserResponse, err error) {
    if request == nil {
        request = NewCreateUserRequest()
    }
    response = NewCreateUserResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAclRequest() (request *DeleteAclRequest) {
    request = &DeleteAclRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "DeleteAcl")
    return
}

func NewDeleteAclResponse() (response *DeleteAclResponse) {
    response = &DeleteAclResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAcl
// 删除ACL
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteAcl(request *DeleteAclRequest) (response *DeleteAclResponse, err error) {
    if request == nil {
        request = NewDeleteAclRequest()
    }
    response = NewDeleteAclResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAclRuleRequest() (request *DeleteAclRuleRequest) {
    request = &DeleteAclRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "DeleteAclRule")
    return
}

func NewDeleteAclRuleResponse() (response *DeleteAclRuleResponse) {
    response = &DeleteAclRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAclRule
// 删除ACL规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteAclRule(request *DeleteAclRuleRequest) (response *DeleteAclRuleResponse, err error) {
    if request == nil {
        request = NewDeleteAclRuleRequest()
    }
    response = NewDeleteAclRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTopicRequest() (request *DeleteTopicRequest) {
    request = &DeleteTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "DeleteTopic")
    return
}

func NewDeleteTopicResponse() (response *DeleteTopicResponse) {
    response = &DeleteTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTopic
// 删除ckafka主题
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteTopic(request *DeleteTopicRequest) (response *DeleteTopicResponse, err error) {
    if request == nil {
        request = NewDeleteTopicRequest()
    }
    response = NewDeleteTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTopicIpWhiteListRequest() (request *DeleteTopicIpWhiteListRequest) {
    request = &DeleteTopicIpWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "DeleteTopicIpWhiteList")
    return
}

func NewDeleteTopicIpWhiteListResponse() (response *DeleteTopicIpWhiteListResponse) {
    response = &DeleteTopicIpWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTopicIpWhiteList
// 删除主题IP白名单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteTopicIpWhiteList(request *DeleteTopicIpWhiteListRequest) (response *DeleteTopicIpWhiteListResponse, err error) {
    if request == nil {
        request = NewDeleteTopicIpWhiteListRequest()
    }
    response = NewDeleteTopicIpWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserRequest() (request *DeleteUserRequest) {
    request = &DeleteUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "DeleteUser")
    return
}

func NewDeleteUserResponse() (response *DeleteUserResponse) {
    response = &DeleteUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteUser
// 删除用户
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteUser(request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    if request == nil {
        request = NewDeleteUserRequest()
    }
    response = NewDeleteUserResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeACLRequest() (request *DescribeACLRequest) {
    request = &DescribeACLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeACL")
    return
}

func NewDescribeACLResponse() (response *DescribeACLResponse) {
    response = &DescribeACLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeACL
// 枚举ACL
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeACL(request *DescribeACLRequest) (response *DescribeACLResponse, err error) {
    if request == nil {
        request = NewDescribeACLRequest()
    }
    response = NewDescribeACLResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppInfoRequest() (request *DescribeAppInfoRequest) {
    request = &DescribeAppInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeAppInfo")
    return
}

func NewDescribeAppInfoResponse() (response *DescribeAppInfoResponse) {
    response = &DescribeAppInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAppInfo
// 查询用户列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeAppInfo(request *DescribeAppInfoRequest) (response *DescribeAppInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAppInfoRequest()
    }
    response = NewDescribeAppInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCkafkaZoneRequest() (request *DescribeCkafkaZoneRequest) {
    request = &DescribeCkafkaZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeCkafkaZone")
    return
}

func NewDescribeCkafkaZoneResponse() (response *DescribeCkafkaZoneResponse) {
    response = &DescribeCkafkaZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCkafkaZone
// 用于查看ckafka的可用区列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeCkafkaZone(request *DescribeCkafkaZoneRequest) (response *DescribeCkafkaZoneResponse, err error) {
    if request == nil {
        request = NewDescribeCkafkaZoneRequest()
    }
    response = NewDescribeCkafkaZoneResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConsumerGroupRequest() (request *DescribeConsumerGroupRequest) {
    request = &DescribeConsumerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeConsumerGroup")
    return
}

func NewDescribeConsumerGroupResponse() (response *DescribeConsumerGroupResponse) {
    response = &DescribeConsumerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeConsumerGroup
// 查询消费分组信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeConsumerGroup(request *DescribeConsumerGroupRequest) (response *DescribeConsumerGroupResponse, err error) {
    if request == nil {
        request = NewDescribeConsumerGroupRequest()
    }
    response = NewDescribeConsumerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupRequest() (request *DescribeGroupRequest) {
    request = &DescribeGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeGroup")
    return
}

func NewDescribeGroupResponse() (response *DescribeGroupResponse) {
    response = &DescribeGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGroup
// 枚举消费分组(精简版)
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeGroup(request *DescribeGroupRequest) (response *DescribeGroupResponse, err error) {
    if request == nil {
        request = NewDescribeGroupRequest()
    }
    response = NewDescribeGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupInfoRequest() (request *DescribeGroupInfoRequest) {
    request = &DescribeGroupInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeGroupInfo")
    return
}

func NewDescribeGroupInfoResponse() (response *DescribeGroupInfoResponse) {
    response = &DescribeGroupInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGroupInfo
// 获取消费分组信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeGroupInfo(request *DescribeGroupInfoRequest) (response *DescribeGroupInfoResponse, err error) {
    if request == nil {
        request = NewDescribeGroupInfoRequest()
    }
    response = NewDescribeGroupInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupOffsetsRequest() (request *DescribeGroupOffsetsRequest) {
    request = &DescribeGroupOffsetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeGroupOffsets")
    return
}

func NewDescribeGroupOffsetsResponse() (response *DescribeGroupOffsetsResponse) {
    response = &DescribeGroupOffsetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGroupOffsets
// 获取消费分组offset
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeGroupOffsets(request *DescribeGroupOffsetsRequest) (response *DescribeGroupOffsetsResponse, err error) {
    if request == nil {
        request = NewDescribeGroupOffsetsRequest()
    }
    response = NewDescribeGroupOffsetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceAttributesRequest() (request *DescribeInstanceAttributesRequest) {
    request = &DescribeInstanceAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeInstanceAttributes")
    return
}

func NewDescribeInstanceAttributesResponse() (response *DescribeInstanceAttributesResponse) {
    response = &DescribeInstanceAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceAttributes
// 获取实例属性
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeInstanceAttributes(request *DescribeInstanceAttributesRequest) (response *DescribeInstanceAttributesResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceAttributesRequest()
    }
    response = NewDescribeInstanceAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeInstances")
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstances
// 本接口（DescribeInstance）用于在用户账户下获取消息队列 CKafka 实例列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesDetailRequest() (request *DescribeInstancesDetailRequest) {
    request = &DescribeInstancesDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeInstancesDetail")
    return
}

func NewDescribeInstancesDetailResponse() (response *DescribeInstancesDetailResponse) {
    response = &DescribeInstancesDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstancesDetail
// 用户账户下获取实例列表详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeInstancesDetail(request *DescribeInstancesDetailRequest) (response *DescribeInstancesDetailResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesDetailRequest()
    }
    response = NewDescribeInstancesDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionRequest() (request *DescribeRegionRequest) {
    request = &DescribeRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeRegion")
    return
}

func NewDescribeRegionResponse() (response *DescribeRegionResponse) {
    response = &DescribeRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRegion
// 枚举地域,只支持广州地域
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeRegion(request *DescribeRegionRequest) (response *DescribeRegionResponse, err error) {
    if request == nil {
        request = NewDescribeRegionRequest()
    }
    response = NewDescribeRegionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRouteRequest() (request *DescribeRouteRequest) {
    request = &DescribeRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeRoute")
    return
}

func NewDescribeRouteResponse() (response *DescribeRouteResponse) {
    response = &DescribeRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRoute
// 查看路由信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeRoute(request *DescribeRouteRequest) (response *DescribeRouteResponse, err error) {
    if request == nil {
        request = NewDescribeRouteRequest()
    }
    response = NewDescribeRouteResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicRequest() (request *DescribeTopicRequest) {
    request = &DescribeTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeTopic")
    return
}

func NewDescribeTopicResponse() (response *DescribeTopicResponse) {
    response = &DescribeTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTopic
// 接口请求域名：https://ckafka.tencentcloudapi.com
//
// 本接口（DescribeTopic）用于在用户获取消息队列 CKafka 实例的主题列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeTopic(request *DescribeTopicRequest) (response *DescribeTopicResponse, err error) {
    if request == nil {
        request = NewDescribeTopicRequest()
    }
    response = NewDescribeTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicAttributesRequest() (request *DescribeTopicAttributesRequest) {
    request = &DescribeTopicAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeTopicAttributes")
    return
}

func NewDescribeTopicAttributesResponse() (response *DescribeTopicAttributesResponse) {
    response = &DescribeTopicAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTopicAttributes
// 获取主题属性
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeTopicAttributes(request *DescribeTopicAttributesRequest) (response *DescribeTopicAttributesResponse, err error) {
    if request == nil {
        request = NewDescribeTopicAttributesRequest()
    }
    response = NewDescribeTopicAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicDetailRequest() (request *DescribeTopicDetailRequest) {
    request = &DescribeTopicDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeTopicDetail")
    return
}

func NewDescribeTopicDetailResponse() (response *DescribeTopicDetailResponse) {
    response = &DescribeTopicDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTopicDetail
// 获取主题列表详情（仅控制台调用）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeTopicDetail(request *DescribeTopicDetailRequest) (response *DescribeTopicDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTopicDetailRequest()
    }
    response = NewDescribeTopicDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserRequest() (request *DescribeUserRequest) {
    request = &DescribeUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeUser")
    return
}

func NewDescribeUserResponse() (response *DescribeUserResponse) {
    response = &DescribeUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUser
// 查询用户信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeUser(request *DescribeUserRequest) (response *DescribeUserResponse, err error) {
    if request == nil {
        request = NewDescribeUserRequest()
    }
    response = NewDescribeUserResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGroupOffsetsRequest() (request *ModifyGroupOffsetsRequest) {
    request = &ModifyGroupOffsetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "ModifyGroupOffsets")
    return
}

func NewModifyGroupOffsetsResponse() (response *ModifyGroupOffsetsResponse) {
    response = &ModifyGroupOffsetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyGroupOffsets
// 设置Groups 消费分组offset
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) ModifyGroupOffsets(request *ModifyGroupOffsetsRequest) (response *ModifyGroupOffsetsResponse, err error) {
    if request == nil {
        request = NewModifyGroupOffsetsRequest()
    }
    response = NewModifyGroupOffsetsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceAttributesRequest() (request *ModifyInstanceAttributesRequest) {
    request = &ModifyInstanceAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "ModifyInstanceAttributes")
    return
}

func NewModifyInstanceAttributesResponse() (response *ModifyInstanceAttributesResponse) {
    response = &ModifyInstanceAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstanceAttributes
// 设置实例属性
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) ModifyInstanceAttributes(request *ModifyInstanceAttributesRequest) (response *ModifyInstanceAttributesResponse, err error) {
    if request == nil {
        request = NewModifyInstanceAttributesRequest()
    }
    response = NewModifyInstanceAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPasswordRequest() (request *ModifyPasswordRequest) {
    request = &ModifyPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "ModifyPassword")
    return
}

func NewModifyPasswordResponse() (response *ModifyPasswordResponse) {
    response = &ModifyPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPassword
// 修改密码
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) ModifyPassword(request *ModifyPasswordRequest) (response *ModifyPasswordResponse, err error) {
    if request == nil {
        request = NewModifyPasswordRequest()
    }
    response = NewModifyPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTopicAttributesRequest() (request *ModifyTopicAttributesRequest) {
    request = &ModifyTopicAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ckafka", APIVersion, "ModifyTopicAttributes")
    return
}

func NewModifyTopicAttributesResponse() (response *ModifyTopicAttributesResponse) {
    response = &ModifyTopicAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyTopicAttributes
// 本接口用于修改主题属性。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) ModifyTopicAttributes(request *ModifyTopicAttributesRequest) (response *ModifyTopicAttributesResponse, err error) {
    if request == nil {
        request = NewModifyTopicAttributesRequest()
    }
    response = NewModifyTopicAttributesResponse()
    err = c.Send(request, response)
    return
}
