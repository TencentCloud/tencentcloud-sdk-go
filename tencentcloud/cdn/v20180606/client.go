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

package v20180606

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-06-06"

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


func NewAddCLSTopicDomainsRequest() (request *AddCLSTopicDomainsRequest) {
    request = &AddCLSTopicDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "AddCLSTopicDomains")
    
    
    return
}

func NewAddCLSTopicDomainsResponse() (response *AddCLSTopicDomainsResponse) {
    response = &AddCLSTopicDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddCLSTopicDomains
// AddCLSTopicDomains 用于新增域名到某日志主题下
//
// 可能返回的错误码:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNAMEINVALID = "InvalidParameter.CdnClsTopicNameInvalid"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSINVALIDPARAM = "InvalidParameter.ClsInvalidParam"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSSERVICENOTACTIVATED = "UnauthorizedOperation.ClsServiceNotActivated"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNSUPPORTEDOPERATION_CLSNOTALLOWED = "UnsupportedOperation.ClsNotAllowed"
func (c *Client) AddCLSTopicDomains(request *AddCLSTopicDomainsRequest) (response *AddCLSTopicDomainsResponse, err error) {
    return c.AddCLSTopicDomainsWithContext(context.Background(), request)
}

// AddCLSTopicDomains
// AddCLSTopicDomains 用于新增域名到某日志主题下
//
// 可能返回的错误码:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNAMEINVALID = "InvalidParameter.CdnClsTopicNameInvalid"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSINVALIDPARAM = "InvalidParameter.ClsInvalidParam"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSSERVICENOTACTIVATED = "UnauthorizedOperation.ClsServiceNotActivated"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNSUPPORTEDOPERATION_CLSNOTALLOWED = "UnsupportedOperation.ClsNotAllowed"
func (c *Client) AddCLSTopicDomainsWithContext(ctx context.Context, request *AddCLSTopicDomainsRequest) (response *AddCLSTopicDomainsResponse, err error) {
    if request == nil {
        request = NewAddCLSTopicDomainsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddCLSTopicDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddCLSTopicDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewAddCdnDomainRequest() (request *AddCdnDomainRequest) {
    request = &AddCdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "AddCdnDomain")
    
    
    return
}

func NewAddCdnDomainResponse() (response *AddCdnDomainResponse) {
    response = &AddCdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddCdnDomain
// AddCdnDomain 用于新增内容分发网络加速域名。1分钟内最多可新增100个域名。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CAMRESOURCEBELONGTODIFFERENTUSER = "InvalidParameter.CamResourceBelongToDifferentUser"
//  INVALIDPARAMETER_CAMRESOURCESIXSTAGEERROR = "InvalidParameter.CamResourceSixStageError"
//  INVALIDPARAMETER_CAMTAGKEYALREADYATTACHED = "InvalidParameter.CamTagKeyAlreadyAttached"
//  INVALIDPARAMETER_CAMTAGKEYILLEGAL = "InvalidParameter.CamTagKeyIllegal"
//  INVALIDPARAMETER_CAMTAGKEYNOTEXIST = "InvalidParameter.CamTagKeyNotExist"
//  INVALIDPARAMETER_CAMTAGVALUEILLEGAL = "InvalidParameter.CamTagValueIllegal"
//  INVALIDPARAMETER_CDNCERTINFONOTFOUND = "InvalidParameter.CdnCertInfoNotFound"
//  INVALIDPARAMETER_CDNCONFIGINVALIDHOST = "InvalidParameter.CdnConfigInvalidHost"
//  INVALIDPARAMETER_CDNCONFIGINVALIDTAG = "InvalidParameter.CdnConfigInvalidTag"
//  INVALIDPARAMETER_CDNCONFIGTAGREQUIRED = "InvalidParameter.CdnConfigTagRequired"
//  INVALIDPARAMETER_CDNHOSTINTERNALHOST = "InvalidParameter.CdnHostInternalHost"
//  INVALIDPARAMETER_CDNHOSTISCOSDEFAULTACCESS = "InvalidParameter.CdnHostIsCosDefaultAccess"
//  INVALIDPARAMETER_CDNHOSTTOOLONGHOST = "InvalidParameter.CdnHostTooLongHost"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNKEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.CdnKeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  LIMITEXCEEDED_CAMRESOURCEARRAYTOOLONG = "LimitExceeded.CamResourceArrayTooLong"
//  LIMITEXCEEDED_CAMRESOURCETOOMANYTAGKEY = "LimitExceeded.CamResourceTooManyTagKey"
//  LIMITEXCEEDED_CAMTAGKEYTOOLONG = "LimitExceeded.CamTagKeyTooLong"
//  LIMITEXCEEDED_CAMTAGKEYTOOMANYTAGVALUE = "LimitExceeded.CamTagKeyTooManyTagValue"
//  LIMITEXCEEDED_CAMTAGQUOTAEXCEEDLIMIT = "LimitExceeded.CamTagQuotaExceedLimit"
//  LIMITEXCEEDED_CAMUSERTOOMANYTAGKEY = "LimitExceeded.CamUserTooManyTagKey"
//  LIMITEXCEEDED_CDNCONFIGTOOMANYCACHERULES = "LimitExceeded.CdnConfigTooManyCacheRules"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  LIMITEXCEEDED_CDNUSERTOOMANYHOSTS = "LimitExceeded.CdnUserTooManyHosts"
//  RESOURCEINUSE_CDNCONFLICTHOSTEXISTS = "ResourceInUse.CdnConflictHostExists"
//  RESOURCEINUSE_CDNHOSTEXISTS = "ResourceInUse.CdnHostExists"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CAMTAGKEYNOTEXIST = "ResourceNotFound.CamTagKeyNotExist"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTEXISTSINDSA = "ResourceUnavailable.CdnHostExistsInDsa"
//  RESOURCEUNAVAILABLE_CDNHOSTEXISTSINTCB = "ResourceUnavailable.CdnHostExistsInTcb"
//  RESOURCEUNAVAILABLE_CDNHOSTISMALICIOUS = "ResourceUnavailable.CdnHostIsMalicious"
//  RESOURCEUNAVAILABLE_CDNHOSTNOICP = "ResourceUnavailable.CdnHostNoIcp"
//  RESOURCEUNAVAILABLE_HOSTEXISTINVOD = "ResourceUnavailable.HostExistInVod"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNDOMAINRECORDNOTVERIFIED = "UnauthorizedOperation.CdnDomainRecordNotVerified"
//  UNAUTHORIZEDOPERATION_CDNHOSTEXISTSININTERNAL = "UnauthorizedOperation.CdnHostExistsInInternal"
//  UNAUTHORIZEDOPERATION_CDNHOSTINICPBLACKLIST = "UnauthorizedOperation.CdnHostInIcpBlacklist"
//  UNAUTHORIZEDOPERATION_CDNHOSTISOWNEDBYOTHER = "UnauthorizedOperation.CdnHostIsOwnedByOther"
//  UNAUTHORIZEDOPERATION_CDNHOSTISUSEDBYOTHER = "UnauthorizedOperation.CdnHostIsUsedByOther"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNTAGUNAUTHORIZED = "UnauthorizedOperation.CdnTagUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISISOLATED = "UnauthorizedOperation.CdnUserIsIsolated"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNSUPPORTEDOPERATION_OPNOAUTH = "UnsupportedOperation.OpNoAuth"
func (c *Client) AddCdnDomain(request *AddCdnDomainRequest) (response *AddCdnDomainResponse, err error) {
    return c.AddCdnDomainWithContext(context.Background(), request)
}

// AddCdnDomain
// AddCdnDomain 用于新增内容分发网络加速域名。1分钟内最多可新增100个域名。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CAMRESOURCEBELONGTODIFFERENTUSER = "InvalidParameter.CamResourceBelongToDifferentUser"
//  INVALIDPARAMETER_CAMRESOURCESIXSTAGEERROR = "InvalidParameter.CamResourceSixStageError"
//  INVALIDPARAMETER_CAMTAGKEYALREADYATTACHED = "InvalidParameter.CamTagKeyAlreadyAttached"
//  INVALIDPARAMETER_CAMTAGKEYILLEGAL = "InvalidParameter.CamTagKeyIllegal"
//  INVALIDPARAMETER_CAMTAGKEYNOTEXIST = "InvalidParameter.CamTagKeyNotExist"
//  INVALIDPARAMETER_CAMTAGVALUEILLEGAL = "InvalidParameter.CamTagValueIllegal"
//  INVALIDPARAMETER_CDNCERTINFONOTFOUND = "InvalidParameter.CdnCertInfoNotFound"
//  INVALIDPARAMETER_CDNCONFIGINVALIDHOST = "InvalidParameter.CdnConfigInvalidHost"
//  INVALIDPARAMETER_CDNCONFIGINVALIDTAG = "InvalidParameter.CdnConfigInvalidTag"
//  INVALIDPARAMETER_CDNCONFIGTAGREQUIRED = "InvalidParameter.CdnConfigTagRequired"
//  INVALIDPARAMETER_CDNHOSTINTERNALHOST = "InvalidParameter.CdnHostInternalHost"
//  INVALIDPARAMETER_CDNHOSTISCOSDEFAULTACCESS = "InvalidParameter.CdnHostIsCosDefaultAccess"
//  INVALIDPARAMETER_CDNHOSTTOOLONGHOST = "InvalidParameter.CdnHostTooLongHost"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNKEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.CdnKeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  LIMITEXCEEDED_CAMRESOURCEARRAYTOOLONG = "LimitExceeded.CamResourceArrayTooLong"
//  LIMITEXCEEDED_CAMRESOURCETOOMANYTAGKEY = "LimitExceeded.CamResourceTooManyTagKey"
//  LIMITEXCEEDED_CAMTAGKEYTOOLONG = "LimitExceeded.CamTagKeyTooLong"
//  LIMITEXCEEDED_CAMTAGKEYTOOMANYTAGVALUE = "LimitExceeded.CamTagKeyTooManyTagValue"
//  LIMITEXCEEDED_CAMTAGQUOTAEXCEEDLIMIT = "LimitExceeded.CamTagQuotaExceedLimit"
//  LIMITEXCEEDED_CAMUSERTOOMANYTAGKEY = "LimitExceeded.CamUserTooManyTagKey"
//  LIMITEXCEEDED_CDNCONFIGTOOMANYCACHERULES = "LimitExceeded.CdnConfigTooManyCacheRules"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  LIMITEXCEEDED_CDNUSERTOOMANYHOSTS = "LimitExceeded.CdnUserTooManyHosts"
//  RESOURCEINUSE_CDNCONFLICTHOSTEXISTS = "ResourceInUse.CdnConflictHostExists"
//  RESOURCEINUSE_CDNHOSTEXISTS = "ResourceInUse.CdnHostExists"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CAMTAGKEYNOTEXIST = "ResourceNotFound.CamTagKeyNotExist"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTEXISTSINDSA = "ResourceUnavailable.CdnHostExistsInDsa"
//  RESOURCEUNAVAILABLE_CDNHOSTEXISTSINTCB = "ResourceUnavailable.CdnHostExistsInTcb"
//  RESOURCEUNAVAILABLE_CDNHOSTISMALICIOUS = "ResourceUnavailable.CdnHostIsMalicious"
//  RESOURCEUNAVAILABLE_CDNHOSTNOICP = "ResourceUnavailable.CdnHostNoIcp"
//  RESOURCEUNAVAILABLE_HOSTEXISTINVOD = "ResourceUnavailable.HostExistInVod"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNDOMAINRECORDNOTVERIFIED = "UnauthorizedOperation.CdnDomainRecordNotVerified"
//  UNAUTHORIZEDOPERATION_CDNHOSTEXISTSININTERNAL = "UnauthorizedOperation.CdnHostExistsInInternal"
//  UNAUTHORIZEDOPERATION_CDNHOSTINICPBLACKLIST = "UnauthorizedOperation.CdnHostInIcpBlacklist"
//  UNAUTHORIZEDOPERATION_CDNHOSTISOWNEDBYOTHER = "UnauthorizedOperation.CdnHostIsOwnedByOther"
//  UNAUTHORIZEDOPERATION_CDNHOSTISUSEDBYOTHER = "UnauthorizedOperation.CdnHostIsUsedByOther"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNTAGUNAUTHORIZED = "UnauthorizedOperation.CdnTagUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISISOLATED = "UnauthorizedOperation.CdnUserIsIsolated"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNSUPPORTEDOPERATION_OPNOAUTH = "UnsupportedOperation.OpNoAuth"
func (c *Client) AddCdnDomainWithContext(ctx context.Context, request *AddCdnDomainRequest) (response *AddCdnDomainResponse, err error) {
    if request == nil {
        request = NewAddCdnDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddCdnDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddCdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClsLogTopicRequest() (request *CreateClsLogTopicRequest) {
    request = &CreateClsLogTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "CreateClsLogTopic")
    
    
    return
}

func NewCreateClsLogTopicResponse() (response *CreateClsLogTopicResponse) {
    response = &CreateClsLogTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateClsLogTopic
// CreateClsLogTopic 用于创建日志主题。注意：一个日志集下至多可创建10个日志主题。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNAMEINVALID = "InvalidParameter.CdnClsTopicNameInvalid"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CDNUSERISISOLATED = "UnauthorizedOperation.CdnUserIsIsolated"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSSERVICENOTACTIVATED = "UnauthorizedOperation.ClsServiceNotActivated"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) CreateClsLogTopic(request *CreateClsLogTopicRequest) (response *CreateClsLogTopicResponse, err error) {
    return c.CreateClsLogTopicWithContext(context.Background(), request)
}

// CreateClsLogTopic
// CreateClsLogTopic 用于创建日志主题。注意：一个日志集下至多可创建10个日志主题。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNAMEINVALID = "InvalidParameter.CdnClsTopicNameInvalid"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CDNUSERISISOLATED = "UnauthorizedOperation.CdnUserIsIsolated"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSSERVICENOTACTIVATED = "UnauthorizedOperation.ClsServiceNotActivated"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) CreateClsLogTopicWithContext(ctx context.Context, request *CreateClsLogTopicRequest) (response *CreateClsLogTopicResponse, err error) {
    if request == nil {
        request = NewCreateClsLogTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateClsLogTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClsLogTopicResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDiagnoseUrlRequest() (request *CreateDiagnoseUrlRequest) {
    request = &CreateDiagnoseUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "CreateDiagnoseUrl")
    
    
    return
}

func NewCreateDiagnoseUrlResponse() (response *CreateDiagnoseUrlResponse) {
    response = &CreateDiagnoseUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDiagnoseUrl
// CreateDiagnoseUrl 用于添加域名诊断任务URL， <font color=red>将于 **2023年5月31日** 下线</font><br>
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYPARAMERROR = "InternalError.CdnQueryParamError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) CreateDiagnoseUrl(request *CreateDiagnoseUrlRequest) (response *CreateDiagnoseUrlResponse, err error) {
    return c.CreateDiagnoseUrlWithContext(context.Background(), request)
}

// CreateDiagnoseUrl
// CreateDiagnoseUrl 用于添加域名诊断任务URL， <font color=red>将于 **2023年5月31日** 下线</font><br>
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYPARAMERROR = "InternalError.CdnQueryParamError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) CreateDiagnoseUrlWithContext(ctx context.Context, request *CreateDiagnoseUrlRequest) (response *CreateDiagnoseUrlResponse, err error) {
    if request == nil {
        request = NewCreateDiagnoseUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDiagnoseUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDiagnoseUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEdgePackTaskRequest() (request *CreateEdgePackTaskRequest) {
    request = &CreateEdgePackTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "CreateEdgePackTask")
    
    
    return
}

func NewCreateEdgePackTaskResponse() (response *CreateEdgePackTaskResponse) {
    response = &CreateEdgePackTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateEdgePackTask
// 动态打包任务提交接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
func (c *Client) CreateEdgePackTask(request *CreateEdgePackTaskRequest) (response *CreateEdgePackTaskResponse, err error) {
    return c.CreateEdgePackTaskWithContext(context.Background(), request)
}

// CreateEdgePackTask
// 动态打包任务提交接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
func (c *Client) CreateEdgePackTaskWithContext(ctx context.Context, request *CreateEdgePackTaskRequest) (response *CreateEdgePackTaskResponse, err error) {
    if request == nil {
        request = NewCreateEdgePackTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEdgePackTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEdgePackTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateScdnDomainRequest() (request *CreateScdnDomainRequest) {
    request = &CreateScdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "CreateScdnDomain")
    
    
    return
}

func NewCreateScdnDomainResponse() (response *CreateScdnDomainResponse) {
    response = &CreateScdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateScdnDomain
// CreateScdnDomain 用于创建 SCDN 加速域名
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_SCDNUSERNOPACKAGE = "InternalError.ScdnUserNoPackage"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNHOSTHASSPECIALCONFIG = "InvalidParameter.CdnHostHasSpecialConfig"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  LIMITEXCEEDED_CDNUSERTOOMANYHOSTS = "LimitExceeded.CdnUserTooManyHosts"
//  RESOURCEINUSE_CDNHOSTEXISTS = "ResourceInUse.CdnHostExists"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  RESOURCEUNAVAILABLE_SCDNUSERNOPACKAGE = "ResourceUnavailable.ScdnUserNoPackage"
//  RESOURCEUNAVAILABLE_SCDNUSERSUSPEND = "ResourceUnavailable.ScdnUserSuspend"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
func (c *Client) CreateScdnDomain(request *CreateScdnDomainRequest) (response *CreateScdnDomainResponse, err error) {
    return c.CreateScdnDomainWithContext(context.Background(), request)
}

// CreateScdnDomain
// CreateScdnDomain 用于创建 SCDN 加速域名
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_SCDNUSERNOPACKAGE = "InternalError.ScdnUserNoPackage"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNHOSTHASSPECIALCONFIG = "InvalidParameter.CdnHostHasSpecialConfig"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  LIMITEXCEEDED_CDNUSERTOOMANYHOSTS = "LimitExceeded.CdnUserTooManyHosts"
//  RESOURCEINUSE_CDNHOSTEXISTS = "ResourceInUse.CdnHostExists"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  RESOURCEUNAVAILABLE_SCDNUSERNOPACKAGE = "ResourceUnavailable.ScdnUserNoPackage"
//  RESOURCEUNAVAILABLE_SCDNUSERSUSPEND = "ResourceUnavailable.ScdnUserSuspend"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
func (c *Client) CreateScdnDomainWithContext(ctx context.Context, request *CreateScdnDomainRequest) (response *CreateScdnDomainResponse, err error) {
    if request == nil {
        request = NewCreateScdnDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateScdnDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateScdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateScdnFailedLogTaskRequest() (request *CreateScdnFailedLogTaskRequest) {
    request = &CreateScdnFailedLogTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "CreateScdnFailedLogTask")
    
    
    return
}

func NewCreateScdnFailedLogTaskResponse() (response *CreateScdnFailedLogTaskResponse) {
    response = &CreateScdnFailedLogTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateScdnFailedLogTask
// CreateScdnFailedLogTask 用于重试创建失败的事件日志任务
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_SCDNLOGTASKEXPIRED = "InvalidParameter.ScdnLogTaskExpired"
//  INVALIDPARAMETER_SCDNLOGTASKNOTFOUNDORNOTFAIL = "InvalidParameter.ScdnLogTaskNotFoundOrNotFail"
//  INVALIDPARAMETER_SCDNLOGTASKTIMERANGEINVALID = "InvalidParameter.ScdnLogTaskTimeRangeInvalid"
//  LIMITEXCEEDED_SCDNLOGTASKEXCEEDDAYLIMIT = "LimitExceeded.ScdnLogTaskExceedDayLimit"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
func (c *Client) CreateScdnFailedLogTask(request *CreateScdnFailedLogTaskRequest) (response *CreateScdnFailedLogTaskResponse, err error) {
    return c.CreateScdnFailedLogTaskWithContext(context.Background(), request)
}

// CreateScdnFailedLogTask
// CreateScdnFailedLogTask 用于重试创建失败的事件日志任务
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_SCDNLOGTASKEXPIRED = "InvalidParameter.ScdnLogTaskExpired"
//  INVALIDPARAMETER_SCDNLOGTASKNOTFOUNDORNOTFAIL = "InvalidParameter.ScdnLogTaskNotFoundOrNotFail"
//  INVALIDPARAMETER_SCDNLOGTASKTIMERANGEINVALID = "InvalidParameter.ScdnLogTaskTimeRangeInvalid"
//  LIMITEXCEEDED_SCDNLOGTASKEXCEEDDAYLIMIT = "LimitExceeded.ScdnLogTaskExceedDayLimit"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
func (c *Client) CreateScdnFailedLogTaskWithContext(ctx context.Context, request *CreateScdnFailedLogTaskRequest) (response *CreateScdnFailedLogTaskResponse, err error) {
    if request == nil {
        request = NewCreateScdnFailedLogTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateScdnFailedLogTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateScdnFailedLogTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateScdnLogTaskRequest() (request *CreateScdnLogTaskRequest) {
    request = &CreateScdnLogTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "CreateScdnLogTask")
    
    
    return
}

func NewCreateScdnLogTaskResponse() (response *CreateScdnLogTaskResponse) {
    response = &CreateScdnLogTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateScdnLogTask
// CreateScdnLogTask 用于创建事件日志任务
//
// 可能返回的错误码:
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_SCDNLOGTASKEXPIRED = "InvalidParameter.ScdnLogTaskExpired"
//  INVALIDPARAMETER_SCDNLOGTASKNOTFOUNDORNOTFAIL = "InvalidParameter.ScdnLogTaskNotFoundOrNotFail"
//  INVALIDPARAMETER_SCDNLOGTASKTIMERANGEINVALID = "InvalidParameter.ScdnLogTaskTimeRangeInvalid"
//  LIMITEXCEEDED_SCDNLOGTASKEXCEEDDAYLIMIT = "LimitExceeded.ScdnLogTaskExceedDayLimit"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
func (c *Client) CreateScdnLogTask(request *CreateScdnLogTaskRequest) (response *CreateScdnLogTaskResponse, err error) {
    return c.CreateScdnLogTaskWithContext(context.Background(), request)
}

// CreateScdnLogTask
// CreateScdnLogTask 用于创建事件日志任务
//
// 可能返回的错误码:
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_SCDNLOGTASKEXPIRED = "InvalidParameter.ScdnLogTaskExpired"
//  INVALIDPARAMETER_SCDNLOGTASKNOTFOUNDORNOTFAIL = "InvalidParameter.ScdnLogTaskNotFoundOrNotFail"
//  INVALIDPARAMETER_SCDNLOGTASKTIMERANGEINVALID = "InvalidParameter.ScdnLogTaskTimeRangeInvalid"
//  LIMITEXCEEDED_SCDNLOGTASKEXCEEDDAYLIMIT = "LimitExceeded.ScdnLogTaskExceedDayLimit"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
func (c *Client) CreateScdnLogTaskWithContext(ctx context.Context, request *CreateScdnLogTaskRequest) (response *CreateScdnLogTaskResponse, err error) {
    if request == nil {
        request = NewCreateScdnLogTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateScdnLogTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateScdnLogTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVerifyRecordRequest() (request *CreateVerifyRecordRequest) {
    request = &CreateVerifyRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "CreateVerifyRecord")
    
    
    return
}

func NewCreateVerifyRecordResponse() (response *CreateVerifyRecordResponse) {
    response = &CreateVerifyRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateVerifyRecord
// 生成一条子域名解析，提示客户添加到域名解析上，用于泛域名及域名取回校验归属权
//
// 可能返回的错误码:
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INVALIDPARAMETER_CDNCONFIGINVALIDHOST = "InvalidParameter.CdnConfigInvalidHost"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) CreateVerifyRecord(request *CreateVerifyRecordRequest) (response *CreateVerifyRecordResponse, err error) {
    return c.CreateVerifyRecordWithContext(context.Background(), request)
}

// CreateVerifyRecord
// 生成一条子域名解析，提示客户添加到域名解析上，用于泛域名及域名取回校验归属权
//
// 可能返回的错误码:
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INVALIDPARAMETER_CDNCONFIGINVALIDHOST = "InvalidParameter.CdnConfigInvalidHost"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) CreateVerifyRecordWithContext(ctx context.Context, request *CreateVerifyRecordRequest) (response *CreateVerifyRecordResponse, err error) {
    if request == nil {
        request = NewCreateVerifyRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVerifyRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVerifyRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCdnDomainRequest() (request *DeleteCdnDomainRequest) {
    request = &DeleteCdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DeleteCdnDomain")
    
    
    return
}

func NewDeleteCdnDomainResponse() (response *DeleteCdnDomainResponse) {
    response = &DeleteCdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCdnDomain
// DeleteCdnDomain 用于删除指定加速域名
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_TAGSYSTEMERROR = "InternalError.TagSystemError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNHOSTHASSPECIALCONFIG = "InvalidParameter.CdnHostHasSpecialConfig"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  RESOURCEUNAVAILABLE_CDNHOSTISNOTOFFLINE = "ResourceUnavailable.CdnHostIsNotOffline"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNSUPPORTEDOPERATION_OPNOAUTH = "UnsupportedOperation.OpNoAuth"
func (c *Client) DeleteCdnDomain(request *DeleteCdnDomainRequest) (response *DeleteCdnDomainResponse, err error) {
    return c.DeleteCdnDomainWithContext(context.Background(), request)
}

// DeleteCdnDomain
// DeleteCdnDomain 用于删除指定加速域名
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_TAGSYSTEMERROR = "InternalError.TagSystemError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNHOSTHASSPECIALCONFIG = "InvalidParameter.CdnHostHasSpecialConfig"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  RESOURCEUNAVAILABLE_CDNHOSTISNOTOFFLINE = "ResourceUnavailable.CdnHostIsNotOffline"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNSUPPORTEDOPERATION_OPNOAUTH = "UnsupportedOperation.OpNoAuth"
func (c *Client) DeleteCdnDomainWithContext(ctx context.Context, request *DeleteCdnDomainRequest) (response *DeleteCdnDomainResponse, err error) {
    if request == nil {
        request = NewDeleteCdnDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCdnDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClsLogTopicRequest() (request *DeleteClsLogTopicRequest) {
    request = &DeleteClsLogTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DeleteClsLogTopic")
    
    
    return
}

func NewDeleteClsLogTopicResponse() (response *DeleteClsLogTopicResponse) {
    response = &DeleteClsLogTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteClsLogTopic
// DeleteClsLogTopic 用于删除日志主题。注意：删除后，所有该日志主题下绑定域名的日志将不再继续投递至该主题，已经投递的日志将会被全部清空。生效时间约为 5~15 分钟。
//
// 可能返回的错误码:
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DeleteClsLogTopic(request *DeleteClsLogTopicRequest) (response *DeleteClsLogTopicResponse, err error) {
    return c.DeleteClsLogTopicWithContext(context.Background(), request)
}

// DeleteClsLogTopic
// DeleteClsLogTopic 用于删除日志主题。注意：删除后，所有该日志主题下绑定域名的日志将不再继续投递至该主题，已经投递的日志将会被全部清空。生效时间约为 5~15 分钟。
//
// 可能返回的错误码:
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DeleteClsLogTopicWithContext(ctx context.Context, request *DeleteClsLogTopicRequest) (response *DeleteClsLogTopicResponse, err error) {
    if request == nil {
        request = NewDeleteClsLogTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteClsLogTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClsLogTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteScdnDomainRequest() (request *DeleteScdnDomainRequest) {
    request = &DeleteScdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DeleteScdnDomain")
    
    
    return
}

func NewDeleteScdnDomainResponse() (response *DeleteScdnDomainResponse) {
    response = &DeleteScdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteScdnDomain
// 删除SCDN域名
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_SCDNUSERNOPACKAGE = "InternalError.ScdnUserNoPackage"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  RESOURCEUNAVAILABLE_SCDNUSERNOPACKAGE = "ResourceUnavailable.ScdnUserNoPackage"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
func (c *Client) DeleteScdnDomain(request *DeleteScdnDomainRequest) (response *DeleteScdnDomainResponse, err error) {
    return c.DeleteScdnDomainWithContext(context.Background(), request)
}

// DeleteScdnDomain
// 删除SCDN域名
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_SCDNUSERNOPACKAGE = "InternalError.ScdnUserNoPackage"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  RESOURCEUNAVAILABLE_SCDNUSERNOPACKAGE = "ResourceUnavailable.ScdnUserNoPackage"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
func (c *Client) DeleteScdnDomainWithContext(ctx context.Context, request *DeleteScdnDomainRequest) (response *DeleteScdnDomainResponse, err error) {
    if request == nil {
        request = NewDeleteScdnDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteScdnDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteScdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillingDataRequest() (request *DescribeBillingDataRequest) {
    request = &DescribeBillingDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeBillingData")
    
    
    return
}

func NewDescribeBillingDataResponse() (response *DescribeBillingDataResponse) {
    response = &DescribeBillingDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBillingData
// DescribeBillingData 用于查询实际计费数据明细。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_COSTDATASYSTEMERROR = "InternalError.CostDataSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNPROJECTNOTEXISTS = "ResourceNotFound.CdnProjectNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNTAGUNAUTHORIZED = "UnauthorizedOperation.CdnTagUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeBillingData(request *DescribeBillingDataRequest) (response *DescribeBillingDataResponse, err error) {
    return c.DescribeBillingDataWithContext(context.Background(), request)
}

// DescribeBillingData
// DescribeBillingData 用于查询实际计费数据明细。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_COSTDATASYSTEMERROR = "InternalError.CostDataSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNPROJECTNOTEXISTS = "ResourceNotFound.CdnProjectNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNTAGUNAUTHORIZED = "UnauthorizedOperation.CdnTagUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeBillingDataWithContext(ctx context.Context, request *DescribeBillingDataRequest) (response *DescribeBillingDataResponse, err error) {
    if request == nil {
        request = NewDescribeBillingDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillingData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillingDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCcDataRequest() (request *DescribeCcDataRequest) {
    request = &DescribeCcDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeCcData")
    
    
    return
}

func NewDescribeCcDataResponse() (response *DescribeCcDataResponse) {
    response = &DescribeCcDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCcData
// CC统计数据查询
//
// 可能返回的错误码:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_DOMAINEMPTY = "UnauthorizedOperation.DomainEmpty"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeCcData(request *DescribeCcDataRequest) (response *DescribeCcDataResponse, err error) {
    return c.DescribeCcDataWithContext(context.Background(), request)
}

// DescribeCcData
// CC统计数据查询
//
// 可能返回的错误码:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_DOMAINEMPTY = "UnauthorizedOperation.DomainEmpty"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeCcDataWithContext(ctx context.Context, request *DescribeCcDataRequest) (response *DescribeCcDataResponse, err error) {
    if request == nil {
        request = NewDescribeCcDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCcData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCcDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCdnDataRequest() (request *DescribeCdnDataRequest) {
    request = &DescribeCdnDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeCdnData")
    
    
    return
}

func NewDescribeCdnDataResponse() (response *DescribeCdnDataResponse) {
    response = &DescribeCdnDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCdnData
// DescribeCdnData 用于查询 CDN 实时访问监控数据，支持以下指标查询：
//
// 
//
// + 流量（单位为 byte）
//
// + 带宽（单位为 bps）
//
// + 请求数（单位为 次）
//
// + 命中请求数（单位为 次）
//
// + 请求命中率（单位为 %）
//
// + 命中流量（单位为 byte）
//
// + 流量命中率（单位为 %）
//
// + 状态码 2xx 汇总及各 2 开头状态码明细（单位为 个）
//
// + 状态码 3xx 汇总及各 3 开头状态码明细（单位为 个）
//
// + 状态码 4xx 汇总及各 4 开头状态码明细（单位为 个）
//
// + 状态码 5xx 汇总及各 5 开头状态码明细（单位为 个）
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_COSTDATASYSTEMERROR = "InternalError.CostDataSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_INVALIDERRORCODE = "InternalError.InvalidErrorCode"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_TAGSYSTEMERROR = "InternalError.TagSystemError"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNINVALIDPARAMMETRIC = "InvalidParameter.CdnInvalidParamMetric"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNPROJECTNOTEXISTS = "ResourceNotFound.CdnProjectNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNTAGUNAUTHORIZED = "UnauthorizedOperation.CdnTagUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_CSRFERROR = "UnauthorizedOperation.CsrfError"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeCdnData(request *DescribeCdnDataRequest) (response *DescribeCdnDataResponse, err error) {
    return c.DescribeCdnDataWithContext(context.Background(), request)
}

// DescribeCdnData
// DescribeCdnData 用于查询 CDN 实时访问监控数据，支持以下指标查询：
//
// 
//
// + 流量（单位为 byte）
//
// + 带宽（单位为 bps）
//
// + 请求数（单位为 次）
//
// + 命中请求数（单位为 次）
//
// + 请求命中率（单位为 %）
//
// + 命中流量（单位为 byte）
//
// + 流量命中率（单位为 %）
//
// + 状态码 2xx 汇总及各 2 开头状态码明细（单位为 个）
//
// + 状态码 3xx 汇总及各 3 开头状态码明细（单位为 个）
//
// + 状态码 4xx 汇总及各 4 开头状态码明细（单位为 个）
//
// + 状态码 5xx 汇总及各 5 开头状态码明细（单位为 个）
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_COSTDATASYSTEMERROR = "InternalError.CostDataSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_INVALIDERRORCODE = "InternalError.InvalidErrorCode"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_TAGSYSTEMERROR = "InternalError.TagSystemError"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNINVALIDPARAMMETRIC = "InvalidParameter.CdnInvalidParamMetric"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNPROJECTNOTEXISTS = "ResourceNotFound.CdnProjectNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNTAGUNAUTHORIZED = "UnauthorizedOperation.CdnTagUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_CSRFERROR = "UnauthorizedOperation.CsrfError"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeCdnDataWithContext(ctx context.Context, request *DescribeCdnDataRequest) (response *DescribeCdnDataResponse, err error) {
    if request == nil {
        request = NewDescribeCdnDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCdnData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCdnDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCdnDomainLogsRequest() (request *DescribeCdnDomainLogsRequest) {
    request = &DescribeCdnDomainLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeCdnDomainLogs")
    
    
    return
}

func NewDescribeCdnDomainLogsResponse() (response *DescribeCdnDomainLogsResponse) {
    response = &DescribeCdnDomainLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCdnDomainLogs
// DescribeCdnDomainLogs 用于查询访问日志下载地址，仅支持 30 天以内的境内、境外访问日志下载链接查询。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNSUPPORTEDOPERATION_OPNOAUTH = "UnsupportedOperation.OpNoAuth"
func (c *Client) DescribeCdnDomainLogs(request *DescribeCdnDomainLogsRequest) (response *DescribeCdnDomainLogsResponse, err error) {
    return c.DescribeCdnDomainLogsWithContext(context.Background(), request)
}

// DescribeCdnDomainLogs
// DescribeCdnDomainLogs 用于查询访问日志下载地址，仅支持 30 天以内的境内、境外访问日志下载链接查询。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNSUPPORTEDOPERATION_OPNOAUTH = "UnsupportedOperation.OpNoAuth"
func (c *Client) DescribeCdnDomainLogsWithContext(ctx context.Context, request *DescribeCdnDomainLogsRequest) (response *DescribeCdnDomainLogsResponse, err error) {
    if request == nil {
        request = NewDescribeCdnDomainLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCdnDomainLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCdnDomainLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCdnIpRequest() (request *DescribeCdnIpRequest) {
    request = &DescribeCdnIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeCdnIp")
    
    
    return
}

func NewDescribeCdnIpResponse() (response *DescribeCdnIpResponse) {
    response = &DescribeCdnIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCdnIp
// DescribeCdnIp 用于查询 CDN IP 归属。
//
// （注意：此接口请求频率限制以 CDN 侧限制为准：200次/10分钟）  
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNCALLINGQUERYIPTOOOFTEN = "LimitExceeded.CdnCallingQueryIpTooOften"
//  LIMITEXCEEDED_CDNQUERYIPBATCHTOOMANY = "LimitExceeded.CdnQueryIpBatchTooMany"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
func (c *Client) DescribeCdnIp(request *DescribeCdnIpRequest) (response *DescribeCdnIpResponse, err error) {
    return c.DescribeCdnIpWithContext(context.Background(), request)
}

// DescribeCdnIp
// DescribeCdnIp 用于查询 CDN IP 归属。
//
// （注意：此接口请求频率限制以 CDN 侧限制为准：200次/10分钟）  
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNCALLINGQUERYIPTOOOFTEN = "LimitExceeded.CdnCallingQueryIpTooOften"
//  LIMITEXCEEDED_CDNQUERYIPBATCHTOOMANY = "LimitExceeded.CdnQueryIpBatchTooMany"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
func (c *Client) DescribeCdnIpWithContext(ctx context.Context, request *DescribeCdnIpRequest) (response *DescribeCdnIpResponse, err error) {
    if request == nil {
        request = NewDescribeCdnIpRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCdnIp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCdnIpResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCdnOriginIpRequest() (request *DescribeCdnOriginIpRequest) {
    request = &DescribeCdnOriginIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeCdnOriginIp")
    
    
    return
}

func NewDescribeCdnOriginIpResponse() (response *DescribeCdnOriginIpResponse) {
    response = &DescribeCdnOriginIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCdnOriginIp
// 本接口（DescribeCdnOriginIp）用于查询 CDN 回源节点的IP信息。（注：此接口即将下线，不再进行维护，请通过DescribeIpStatus 接口进行查询）
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNCALLINGQUERYIPTOOOFTEN = "LimitExceeded.CdnCallingQueryIpTooOften"
//  LIMITEXCEEDED_CDNQUERYIPBATCHTOOMANY = "LimitExceeded.CdnQueryIpBatchTooMany"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
func (c *Client) DescribeCdnOriginIp(request *DescribeCdnOriginIpRequest) (response *DescribeCdnOriginIpResponse, err error) {
    return c.DescribeCdnOriginIpWithContext(context.Background(), request)
}

// DescribeCdnOriginIp
// 本接口（DescribeCdnOriginIp）用于查询 CDN 回源节点的IP信息。（注：此接口即将下线，不再进行维护，请通过DescribeIpStatus 接口进行查询）
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNCALLINGQUERYIPTOOOFTEN = "LimitExceeded.CdnCallingQueryIpTooOften"
//  LIMITEXCEEDED_CDNQUERYIPBATCHTOOMANY = "LimitExceeded.CdnQueryIpBatchTooMany"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
func (c *Client) DescribeCdnOriginIpWithContext(ctx context.Context, request *DescribeCdnOriginIpRequest) (response *DescribeCdnOriginIpResponse, err error) {
    if request == nil {
        request = NewDescribeCdnOriginIpRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCdnOriginIp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCdnOriginIpResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertDomainsRequest() (request *DescribeCertDomainsRequest) {
    request = &DescribeCertDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeCertDomains")
    
    
    return
}

func NewDescribeCertDomainsResponse() (response *DescribeCertDomainsResponse) {
    response = &DescribeCertDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCertDomains
// DescribeCertDomains 用于校验SSL证书并提取证书中包含的域名。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INVALIDPARAMETER_CDNCERTNOCERTINFO = "InvalidParameter.CdnCertNoCertInfo"
//  INVALIDPARAMETER_CDNCERTNOTPEM = "InvalidParameter.CdnCertNotPem"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNHOSTISUSEDBYOTHER = "UnauthorizedOperation.CdnHostIsUsedByOther"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNSUPPORTEDOPERATION_CLSNOTALLOWED = "UnsupportedOperation.ClsNotAllowed"
//  UNSUPPORTEDOPERATION_OPNOAUTH = "UnsupportedOperation.OpNoAuth"
func (c *Client) DescribeCertDomains(request *DescribeCertDomainsRequest) (response *DescribeCertDomainsResponse, err error) {
    return c.DescribeCertDomainsWithContext(context.Background(), request)
}

// DescribeCertDomains
// DescribeCertDomains 用于校验SSL证书并提取证书中包含的域名。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INVALIDPARAMETER_CDNCERTNOCERTINFO = "InvalidParameter.CdnCertNoCertInfo"
//  INVALIDPARAMETER_CDNCERTNOTPEM = "InvalidParameter.CdnCertNotPem"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNHOSTISUSEDBYOTHER = "UnauthorizedOperation.CdnHostIsUsedByOther"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNSUPPORTEDOPERATION_CLSNOTALLOWED = "UnsupportedOperation.ClsNotAllowed"
//  UNSUPPORTEDOPERATION_OPNOAUTH = "UnsupportedOperation.OpNoAuth"
func (c *Client) DescribeCertDomainsWithContext(ctx context.Context, request *DescribeCertDomainsRequest) (response *DescribeCertDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeCertDomainsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCertDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCertDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSDataRequest() (request *DescribeDDoSDataRequest) {
    request = &DescribeDDoSDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeDDoSData")
    
    
    return
}

func NewDescribeDDoSDataResponse() (response *DescribeDDoSDataResponse) {
    response = &DescribeDDoSDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDoSData
// DDoS统计数据查询
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMDBERROR = "InternalError.SystemDBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDDoSData(request *DescribeDDoSDataRequest) (response *DescribeDDoSDataResponse, err error) {
    return c.DescribeDDoSDataWithContext(context.Background(), request)
}

// DescribeDDoSData
// DDoS统计数据查询
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMDBERROR = "InternalError.SystemDBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDDoSDataWithContext(ctx context.Context, request *DescribeDDoSDataRequest) (response *DescribeDDoSDataResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDoSDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDiagnoseReportRequest() (request *DescribeDiagnoseReportRequest) {
    request = &DescribeDiagnoseReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeDiagnoseReport")
    
    
    return
}

func NewDescribeDiagnoseReportResponse() (response *DescribeDiagnoseReportResponse) {
    response = &DescribeDiagnoseReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDiagnoseReport
// DescribeDiagnoseReport 用于获取指定报告id的内容， <font color=red>将于 **2023年5月31日** 下线</font><br>
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
func (c *Client) DescribeDiagnoseReport(request *DescribeDiagnoseReportRequest) (response *DescribeDiagnoseReportResponse, err error) {
    return c.DescribeDiagnoseReportWithContext(context.Background(), request)
}

// DescribeDiagnoseReport
// DescribeDiagnoseReport 用于获取指定报告id的内容， <font color=red>将于 **2023年5月31日** 下线</font><br>
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
func (c *Client) DescribeDiagnoseReportWithContext(ctx context.Context, request *DescribeDiagnoseReportRequest) (response *DescribeDiagnoseReportResponse, err error) {
    if request == nil {
        request = NewDescribeDiagnoseReportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDiagnoseReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDiagnoseReportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDistrictIspDataRequest() (request *DescribeDistrictIspDataRequest) {
    request = &DescribeDistrictIspDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeDistrictIspData")
    
    
    return
}

func NewDescribeDistrictIspDataResponse() (response *DescribeDistrictIspDataResponse) {
    response = &DescribeDistrictIspDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDistrictIspData
// 查询指定域名的区域、运营商明细数据
//
// 注意事项：接口尚未全量开放，未在内测名单中的账号不支持调用
//
// 可能返回的错误码:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNINVALIDPARAMINTERVAL = "InvalidParameter.CdnInvalidParamInterval"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeDistrictIspData(request *DescribeDistrictIspDataRequest) (response *DescribeDistrictIspDataResponse, err error) {
    return c.DescribeDistrictIspDataWithContext(context.Background(), request)
}

// DescribeDistrictIspData
// 查询指定域名的区域、运营商明细数据
//
// 注意事项：接口尚未全量开放，未在内测名单中的账号不支持调用
//
// 可能返回的错误码:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNINVALIDPARAMINTERVAL = "InvalidParameter.CdnInvalidParamInterval"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeDistrictIspDataWithContext(ctx context.Context, request *DescribeDistrictIspDataRequest) (response *DescribeDistrictIspDataResponse, err error) {
    if request == nil {
        request = NewDescribeDistrictIspDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDistrictIspData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDistrictIspDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainsRequest() (request *DescribeDomainsRequest) {
    request = &DescribeDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeDomains")
    
    
    return
}

func NewDescribeDomainsResponse() (response *DescribeDomainsResponse) {
    response = &DescribeDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDomains
// DescribeDomains 用于查询内容分发网络加速域名（含境内、境外）基本配置信息，包括项目ID、服务状态，业务类型、创建时间、更新时间等信息。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNTAGUNAUTHORIZED = "UnauthorizedOperation.CdnTagUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeDomains(request *DescribeDomainsRequest) (response *DescribeDomainsResponse, err error) {
    return c.DescribeDomainsWithContext(context.Background(), request)
}

// DescribeDomains
// DescribeDomains 用于查询内容分发网络加速域名（含境内、境外）基本配置信息，包括项目ID、服务状态，业务类型、创建时间、更新时间等信息。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNTAGUNAUTHORIZED = "UnauthorizedOperation.CdnTagUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeDomainsWithContext(ctx context.Context, request *DescribeDomainsRequest) (response *DescribeDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeDomainsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainsConfigRequest() (request *DescribeDomainsConfigRequest) {
    request = &DescribeDomainsConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeDomainsConfig")
    
    
    return
}

func NewDescribeDomainsConfigResponse() (response *DescribeDomainsConfigResponse) {
    response = &DescribeDomainsConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDomainsConfig
// DescribeDomainsConfig 用于查询内容分发网络加速域名（含境内、境外）的所有配置信息。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNCERTNOCERTINFO = "InvalidParameter.CdnCertNoCertInfo"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNTAGUNAUTHORIZED = "UnauthorizedOperation.CdnTagUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeDomainsConfig(request *DescribeDomainsConfigRequest) (response *DescribeDomainsConfigResponse, err error) {
    return c.DescribeDomainsConfigWithContext(context.Background(), request)
}

// DescribeDomainsConfig
// DescribeDomainsConfig 用于查询内容分发网络加速域名（含境内、境外）的所有配置信息。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNCERTNOCERTINFO = "InvalidParameter.CdnCertNoCertInfo"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNTAGUNAUTHORIZED = "UnauthorizedOperation.CdnTagUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeDomainsConfigWithContext(ctx context.Context, request *DescribeDomainsConfigRequest) (response *DescribeDomainsConfigResponse, err error) {
    if request == nil {
        request = NewDescribeDomainsConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainsConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainsConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgePackTaskStatusRequest() (request *DescribeEdgePackTaskStatusRequest) {
    request = &DescribeEdgePackTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeEdgePackTaskStatus")
    
    
    return
}

func NewDescribeEdgePackTaskStatusResponse() (response *DescribeEdgePackTaskStatusResponse) {
    response = &DescribeEdgePackTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgePackTaskStatus
// DescribeEdgePackTaskStatus 用于查询动态打包任务状态列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
func (c *Client) DescribeEdgePackTaskStatus(request *DescribeEdgePackTaskStatusRequest) (response *DescribeEdgePackTaskStatusResponse, err error) {
    return c.DescribeEdgePackTaskStatusWithContext(context.Background(), request)
}

// DescribeEdgePackTaskStatus
// DescribeEdgePackTaskStatus 用于查询动态打包任务状态列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
func (c *Client) DescribeEdgePackTaskStatusWithContext(ctx context.Context, request *DescribeEdgePackTaskStatusRequest) (response *DescribeEdgePackTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeEdgePackTaskStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgePackTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgePackTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEventLogDataRequest() (request *DescribeEventLogDataRequest) {
    request = &DescribeEventLogDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeEventLogData")
    
    
    return
}

func NewDescribeEventLogDataResponse() (response *DescribeEventLogDataResponse) {
    response = &DescribeEventLogDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEventLogData
// DescribeEventLogData 用于查询事件日志统计曲线
//
// 可能返回的错误码:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMDBERROR = "InternalError.SystemDBError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeEventLogData(request *DescribeEventLogDataRequest) (response *DescribeEventLogDataResponse, err error) {
    return c.DescribeEventLogDataWithContext(context.Background(), request)
}

// DescribeEventLogData
// DescribeEventLogData 用于查询事件日志统计曲线
//
// 可能返回的错误码:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMDBERROR = "InternalError.SystemDBError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeEventLogDataWithContext(ctx context.Context, request *DescribeEventLogDataRequest) (response *DescribeEventLogDataResponse, err error) {
    if request == nil {
        request = NewDescribeEventLogDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEventLogData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEventLogDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHttpsPackagesRequest() (request *DescribeHttpsPackagesRequest) {
    request = &DescribeHttpsPackagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeHttpsPackages")
    
    
    return
}

func NewDescribeHttpsPackagesResponse() (response *DescribeHttpsPackagesResponse) {
    response = &DescribeHttpsPackagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeHttpsPackages
// DescribeHttpsPackages 用于查询 CDN HTTPS请求包详情。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeHttpsPackages(request *DescribeHttpsPackagesRequest) (response *DescribeHttpsPackagesResponse, err error) {
    return c.DescribeHttpsPackagesWithContext(context.Background(), request)
}

// DescribeHttpsPackages
// DescribeHttpsPackages 用于查询 CDN HTTPS请求包详情。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeHttpsPackagesWithContext(ctx context.Context, request *DescribeHttpsPackagesRequest) (response *DescribeHttpsPackagesResponse, err error) {
    if request == nil {
        request = NewDescribeHttpsPackagesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHttpsPackages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHttpsPackagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageConfigRequest() (request *DescribeImageConfigRequest) {
    request = &DescribeImageConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeImageConfig")
    
    
    return
}

func NewDescribeImageConfigResponse() (response *DescribeImageConfigResponse) {
    response = &DescribeImageConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImageConfig
// DescribeImageConfig 用于获取域名图片优化的当前配置，支持Webp、TPG、 Guetzli 和 Avif。 
//
// 可能返回的错误码:
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeImageConfig(request *DescribeImageConfigRequest) (response *DescribeImageConfigResponse, err error) {
    return c.DescribeImageConfigWithContext(context.Background(), request)
}

// DescribeImageConfig
// DescribeImageConfig 用于获取域名图片优化的当前配置，支持Webp、TPG、 Guetzli 和 Avif。 
//
// 可能返回的错误码:
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeImageConfigWithContext(ctx context.Context, request *DescribeImageConfigRequest) (response *DescribeImageConfigResponse, err error) {
    if request == nil {
        request = NewDescribeImageConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIpStatusRequest() (request *DescribeIpStatusRequest) {
    request = &DescribeIpStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeIpStatus")
    
    
    return
}

func NewDescribeIpStatusResponse() (response *DescribeIpStatusResponse) {
    response = &DescribeIpStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIpStatus
// DescribeIpStatus 用于查询域名所在加速平台的边缘节点、回源节点明细。注意事项：暂不支持查询边缘节点信息并且数据会存在一定延迟。
//
// 
//
// >?  若您的业务已迁移至 CDN 控制台，请参考<a href="https://cloud.tencent.com/document/api/228/41954"> CDN 接口文档</a>，使用  CDN 相关API 进行操作。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNCONFIGINVALIDHOST = "InvalidParameter.CdnConfigInvalidHost"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeIpStatus(request *DescribeIpStatusRequest) (response *DescribeIpStatusResponse, err error) {
    return c.DescribeIpStatusWithContext(context.Background(), request)
}

// DescribeIpStatus
// DescribeIpStatus 用于查询域名所在加速平台的边缘节点、回源节点明细。注意事项：暂不支持查询边缘节点信息并且数据会存在一定延迟。
//
// 
//
// >?  若您的业务已迁移至 CDN 控制台，请参考<a href="https://cloud.tencent.com/document/api/228/41954"> CDN 接口文档</a>，使用  CDN 相关API 进行操作。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNCONFIGINVALIDHOST = "InvalidParameter.CdnConfigInvalidHost"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeIpStatusWithContext(ctx context.Context, request *DescribeIpStatusRequest) (response *DescribeIpStatusResponse, err error) {
    if request == nil {
        request = NewDescribeIpStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIpStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIpStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIpVisitRequest() (request *DescribeIpVisitRequest) {
    request = &DescribeIpVisitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeIpVisit")
    
    
    return
}

func NewDescribeIpVisitResponse() (response *DescribeIpVisitResponse) {
    response = &DescribeIpVisitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIpVisit
// DescribeIpVisit 用于查询 5 分钟活跃用户数，及日活跃用户数明细
//
// 
//
// + 5 分钟活跃用户数：根据日志中客户端 IP，5 分钟粒度去重统计
//
// + 日活跃用户数：根据日志中客户端 IP，按天粒度去重统计
//
// 可能返回的错误码:
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNINVALIDPARAMINTERVAL = "InvalidParameter.CdnInvalidParamInterval"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNPROJECTNOTEXISTS = "ResourceNotFound.CdnProjectNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeIpVisit(request *DescribeIpVisitRequest) (response *DescribeIpVisitResponse, err error) {
    return c.DescribeIpVisitWithContext(context.Background(), request)
}

// DescribeIpVisit
// DescribeIpVisit 用于查询 5 分钟活跃用户数，及日活跃用户数明细
//
// 
//
// + 5 分钟活跃用户数：根据日志中客户端 IP，5 分钟粒度去重统计
//
// + 日活跃用户数：根据日志中客户端 IP，按天粒度去重统计
//
// 可能返回的错误码:
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNINVALIDPARAMINTERVAL = "InvalidParameter.CdnInvalidParamInterval"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNPROJECTNOTEXISTS = "ResourceNotFound.CdnProjectNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeIpVisitWithContext(ctx context.Context, request *DescribeIpVisitRequest) (response *DescribeIpVisitResponse, err error) {
    if request == nil {
        request = NewDescribeIpVisitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIpVisit require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIpVisitResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMapInfoRequest() (request *DescribeMapInfoRequest) {
    request = &DescribeMapInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeMapInfo")
    
    
    return
}

func NewDescribeMapInfoResponse() (response *DescribeMapInfoResponse) {
    response = &DescribeMapInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMapInfo
// DescribeMapInfo 用于查询省份对应的 ID，运营商对应的 ID 信息。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeMapInfo(request *DescribeMapInfoRequest) (response *DescribeMapInfoResponse, err error) {
    return c.DescribeMapInfoWithContext(context.Background(), request)
}

// DescribeMapInfo
// DescribeMapInfo 用于查询省份对应的 ID，运营商对应的 ID 信息。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeMapInfoWithContext(ctx context.Context, request *DescribeMapInfoRequest) (response *DescribeMapInfoResponse, err error) {
    if request == nil {
        request = NewDescribeMapInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMapInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMapInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOriginDataRequest() (request *DescribeOriginDataRequest) {
    request = &DescribeOriginDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeOriginData")
    
    
    return
}

func NewDescribeOriginDataResponse() (response *DescribeOriginDataResponse) {
    response = &DescribeOriginDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOriginData
// DescribeOriginData 用于查询 CDN 实时回源监控数据，支持以下指标查询：
//
// 
//
// + 回源流量（单位为 byte）
//
// + 回源带宽（单位为 bps）
//
// + 回源请求数（单位为 次）
//
// + 回源失败请求数（单位为 次）
//
// + 回源失败率（单位为 %，小数点后保留两位）
//
// + 回源状态码 2xx 汇总及各 2 开头回源状态码明细（单位为 个）
//
// + 回源状态码 3xx 汇总及各 3 开头回源状态码明细（单位为 个）
//
// + 回源状态码 4xx 汇总及各 4 开头回源状态码明细（单位为 个）
//
// + 回源状态码 5xx 汇总及各 5 开头回源状态码明细（单位为 个）
//
// 可能返回的错误码:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_INVALIDERRORCODE = "InternalError.InvalidErrorCode"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_TAGSYSTEMERROR = "InternalError.TagSystemError"
//  INVALIDPARAMETER_CDNHOSTINVALIDMIDDLECONFIG = "InvalidParameter.CdnHostInvalidMiddleConfig"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNINVALIDPARAMMETRIC = "InvalidParameter.CdnInvalidParamMetric"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNPROJECTNOTEXISTS = "ResourceNotFound.CdnProjectNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CDNUSERTOOMANYHOSTS = "ResourceNotFound.CdnUserTooManyHosts"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeOriginData(request *DescribeOriginDataRequest) (response *DescribeOriginDataResponse, err error) {
    return c.DescribeOriginDataWithContext(context.Background(), request)
}

// DescribeOriginData
// DescribeOriginData 用于查询 CDN 实时回源监控数据，支持以下指标查询：
//
// 
//
// + 回源流量（单位为 byte）
//
// + 回源带宽（单位为 bps）
//
// + 回源请求数（单位为 次）
//
// + 回源失败请求数（单位为 次）
//
// + 回源失败率（单位为 %，小数点后保留两位）
//
// + 回源状态码 2xx 汇总及各 2 开头回源状态码明细（单位为 个）
//
// + 回源状态码 3xx 汇总及各 3 开头回源状态码明细（单位为 个）
//
// + 回源状态码 4xx 汇总及各 4 开头回源状态码明细（单位为 个）
//
// + 回源状态码 5xx 汇总及各 5 开头回源状态码明细（单位为 个）
//
// 可能返回的错误码:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_INVALIDERRORCODE = "InternalError.InvalidErrorCode"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_TAGSYSTEMERROR = "InternalError.TagSystemError"
//  INVALIDPARAMETER_CDNHOSTINVALIDMIDDLECONFIG = "InvalidParameter.CdnHostInvalidMiddleConfig"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNINVALIDPARAMMETRIC = "InvalidParameter.CdnInvalidParamMetric"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNPROJECTNOTEXISTS = "ResourceNotFound.CdnProjectNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CDNUSERTOOMANYHOSTS = "ResourceNotFound.CdnUserTooManyHosts"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeOriginDataWithContext(ctx context.Context, request *DescribeOriginDataRequest) (response *DescribeOriginDataResponse, err error) {
    if request == nil {
        request = NewDescribeOriginDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOriginData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOriginDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePayTypeRequest() (request *DescribePayTypeRequest) {
    request = &DescribePayTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribePayType")
    
    
    return
}

func NewDescribePayTypeResponse() (response *DescribePayTypeResponse) {
    response = &DescribePayTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePayType
// DescribePayType 用于查询用户的计费类型，计费周期等信息。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribePayType(request *DescribePayTypeRequest) (response *DescribePayTypeResponse, err error) {
    return c.DescribePayTypeWithContext(context.Background(), request)
}

// DescribePayType
// DescribePayType 用于查询用户的计费类型，计费周期等信息。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribePayTypeWithContext(ctx context.Context, request *DescribePayTypeRequest) (response *DescribePayTypeResponse, err error) {
    if request == nil {
        request = NewDescribePayTypeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePayType require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePayTypeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePurgeQuotaRequest() (request *DescribePurgeQuotaRequest) {
    request = &DescribePurgeQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribePurgeQuota")
    
    
    return
}

func NewDescribePurgeQuotaResponse() (response *DescribePurgeQuotaResponse) {
    response = &DescribePurgeQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePurgeQuota
// DescribePurgeQuota 用于查询账户刷新配额和每日可用量。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribePurgeQuota(request *DescribePurgeQuotaRequest) (response *DescribePurgeQuotaResponse, err error) {
    return c.DescribePurgeQuotaWithContext(context.Background(), request)
}

// DescribePurgeQuota
// DescribePurgeQuota 用于查询账户刷新配额和每日可用量。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribePurgeQuotaWithContext(ctx context.Context, request *DescribePurgeQuotaRequest) (response *DescribePurgeQuotaResponse, err error) {
    if request == nil {
        request = NewDescribePurgeQuotaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePurgeQuota require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePurgeQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePurgeTasksRequest() (request *DescribePurgeTasksRequest) {
    request = &DescribePurgeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribePurgeTasks")
    
    
    return
}

func NewDescribePurgeTasksResponse() (response *DescribePurgeTasksResponse) {
    response = &DescribePurgeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePurgeTasks
// DescribePurgeTasks 用于查询提交的 URL 刷新、目录刷新记录及执行进度，通过 PurgePathCache 与 PurgeUrlsCache 接口提交的任务均可通过此接口进行查询。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribePurgeTasks(request *DescribePurgeTasksRequest) (response *DescribePurgeTasksResponse, err error) {
    return c.DescribePurgeTasksWithContext(context.Background(), request)
}

// DescribePurgeTasks
// DescribePurgeTasks 用于查询提交的 URL 刷新、目录刷新记录及执行进度，通过 PurgePathCache 与 PurgeUrlsCache 接口提交的任务均可通过此接口进行查询。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribePurgeTasksWithContext(ctx context.Context, request *DescribePurgeTasksRequest) (response *DescribePurgeTasksResponse, err error) {
    if request == nil {
        request = NewDescribePurgeTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePurgeTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePurgeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePushQuotaRequest() (request *DescribePushQuotaRequest) {
    request = &DescribePushQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribePushQuota")
    
    
    return
}

func NewDescribePushQuotaResponse() (response *DescribePushQuotaResponse) {
    response = &DescribePushQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePushQuota
// DescribePushQuota  用于查询预热配额和每日可用量。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribePushQuota(request *DescribePushQuotaRequest) (response *DescribePushQuotaResponse, err error) {
    return c.DescribePushQuotaWithContext(context.Background(), request)
}

// DescribePushQuota
// DescribePushQuota  用于查询预热配额和每日可用量。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribePushQuotaWithContext(ctx context.Context, request *DescribePushQuotaRequest) (response *DescribePushQuotaResponse, err error) {
    if request == nil {
        request = NewDescribePushQuotaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePushQuota require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePushQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePushTasksRequest() (request *DescribePushTasksRequest) {
    request = &DescribePushTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribePushTasks")
    
    
    return
}

func NewDescribePushTasksResponse() (response *DescribePushTasksResponse) {
    response = &DescribePushTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePushTasks
// DescribePushTasks  用于查询预热任务提交历史记录及执行进度。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISNOTONLINE = "ResourceUnavailable.CdnHostIsNotOnline"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribePushTasks(request *DescribePushTasksRequest) (response *DescribePushTasksResponse, err error) {
    return c.DescribePushTasksWithContext(context.Background(), request)
}

// DescribePushTasks
// DescribePushTasks  用于查询预热任务提交历史记录及执行进度。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISNOTONLINE = "ResourceUnavailable.CdnHostIsNotOnline"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribePushTasksWithContext(ctx context.Context, request *DescribePushTasksRequest) (response *DescribePushTasksResponse, err error) {
    if request == nil {
        request = NewDescribePushTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePushTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePushTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReportDataRequest() (request *DescribeReportDataRequest) {
    request = &DescribeReportDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeReportData")
    
    
    return
}

func NewDescribeReportDataResponse() (response *DescribeReportDataResponse) {
    response = &DescribeReportDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeReportData
// DescribeReportData 用于查询域名/项目维度的日/周/月报表数据。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_INVALIDERRORCODE = "InternalError.InvalidErrorCode"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNSUPPORTEDOPERATION_OPNOAUTH = "UnsupportedOperation.OpNoAuth"
func (c *Client) DescribeReportData(request *DescribeReportDataRequest) (response *DescribeReportDataResponse, err error) {
    return c.DescribeReportDataWithContext(context.Background(), request)
}

// DescribeReportData
// DescribeReportData 用于查询域名/项目维度的日/周/月报表数据。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_INVALIDERRORCODE = "InternalError.InvalidErrorCode"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNSUPPORTEDOPERATION_OPNOAUTH = "UnsupportedOperation.OpNoAuth"
func (c *Client) DescribeReportDataWithContext(ctx context.Context, request *DescribeReportDataRequest) (response *DescribeReportDataResponse, err error) {
    if request == nil {
        request = NewDescribeReportDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReportData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReportDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScdnBotDataRequest() (request *DescribeScdnBotDataRequest) {
    request = &DescribeScdnBotDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeScdnBotData")
    
    
    return
}

func NewDescribeScdnBotDataResponse() (response *DescribeScdnBotDataResponse) {
    response = &DescribeScdnBotDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScdnBotData
// 获取BOT统计数据列表
//
// 可能返回的错误码:
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMDBERROR = "InternalError.SystemDBError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_DOMAINEMPTY = "UnauthorizedOperation.DomainEmpty"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeScdnBotData(request *DescribeScdnBotDataRequest) (response *DescribeScdnBotDataResponse, err error) {
    return c.DescribeScdnBotDataWithContext(context.Background(), request)
}

// DescribeScdnBotData
// 获取BOT统计数据列表
//
// 可能返回的错误码:
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMDBERROR = "InternalError.SystemDBError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_DOMAINEMPTY = "UnauthorizedOperation.DomainEmpty"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeScdnBotDataWithContext(ctx context.Context, request *DescribeScdnBotDataRequest) (response *DescribeScdnBotDataResponse, err error) {
    if request == nil {
        request = NewDescribeScdnBotDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScdnBotData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScdnBotDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScdnBotRecordsRequest() (request *DescribeScdnBotRecordsRequest) {
    request = &DescribeScdnBotRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeScdnBotRecords")
    
    
    return
}

func NewDescribeScdnBotRecordsResponse() (response *DescribeScdnBotRecordsResponse) {
    response = &DescribeScdnBotRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScdnBotRecords
// 查询BOT会话记录列表
//
// 可能返回的错误码:
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_SYSTEMDBERROR = "InternalError.SystemDBError"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
func (c *Client) DescribeScdnBotRecords(request *DescribeScdnBotRecordsRequest) (response *DescribeScdnBotRecordsResponse, err error) {
    return c.DescribeScdnBotRecordsWithContext(context.Background(), request)
}

// DescribeScdnBotRecords
// 查询BOT会话记录列表
//
// 可能返回的错误码:
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_SYSTEMDBERROR = "InternalError.SystemDBError"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
func (c *Client) DescribeScdnBotRecordsWithContext(ctx context.Context, request *DescribeScdnBotRecordsRequest) (response *DescribeScdnBotRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeScdnBotRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScdnBotRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScdnBotRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScdnConfigRequest() (request *DescribeScdnConfigRequest) {
    request = &DescribeScdnConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeScdnConfig")
    
    
    return
}

func NewDescribeScdnConfigResponse() (response *DescribeScdnConfigResponse) {
    response = &DescribeScdnConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScdnConfig
// DescribeScdnConfig 用于查询指定 SCDN 加速域名的安全相关配置
//
// 可能返回的错误码:
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
func (c *Client) DescribeScdnConfig(request *DescribeScdnConfigRequest) (response *DescribeScdnConfigResponse, err error) {
    return c.DescribeScdnConfigWithContext(context.Background(), request)
}

// DescribeScdnConfig
// DescribeScdnConfig 用于查询指定 SCDN 加速域名的安全相关配置
//
// 可能返回的错误码:
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
func (c *Client) DescribeScdnConfigWithContext(ctx context.Context, request *DescribeScdnConfigRequest) (response *DescribeScdnConfigResponse, err error) {
    if request == nil {
        request = NewDescribeScdnConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScdnConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScdnConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScdnIpStrategyRequest() (request *DescribeScdnIpStrategyRequest) {
    request = &DescribeScdnIpStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeScdnIpStrategy")
    
    
    return
}

func NewDescribeScdnIpStrategyResponse() (response *DescribeScdnIpStrategyResponse) {
    response = &DescribeScdnIpStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScdnIpStrategy
// 查询在SCDN IP安全策略
//
// 可能返回的错误码:
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_SCDNUSERNOPACKAGE = "InternalError.ScdnUserNoPackage"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
func (c *Client) DescribeScdnIpStrategy(request *DescribeScdnIpStrategyRequest) (response *DescribeScdnIpStrategyResponse, err error) {
    return c.DescribeScdnIpStrategyWithContext(context.Background(), request)
}

// DescribeScdnIpStrategy
// 查询在SCDN IP安全策略
//
// 可能返回的错误码:
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_SCDNUSERNOPACKAGE = "InternalError.ScdnUserNoPackage"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
func (c *Client) DescribeScdnIpStrategyWithContext(ctx context.Context, request *DescribeScdnIpStrategyRequest) (response *DescribeScdnIpStrategyResponse, err error) {
    if request == nil {
        request = NewDescribeScdnIpStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScdnIpStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScdnIpStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScdnTopDataRequest() (request *DescribeScdnTopDataRequest) {
    request = &DescribeScdnTopDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeScdnTopData")
    
    
    return
}

func NewDescribeScdnTopDataResponse() (response *DescribeScdnTopDataResponse) {
    response = &DescribeScdnTopDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScdnTopData
// 获取SCDN的Top数据
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDFILTER = "InvalidParameter.CdnStatInvalidFilter"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
func (c *Client) DescribeScdnTopData(request *DescribeScdnTopDataRequest) (response *DescribeScdnTopDataResponse, err error) {
    return c.DescribeScdnTopDataWithContext(context.Background(), request)
}

// DescribeScdnTopData
// 获取SCDN的Top数据
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDFILTER = "InvalidParameter.CdnStatInvalidFilter"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
func (c *Client) DescribeScdnTopDataWithContext(ctx context.Context, request *DescribeScdnTopDataRequest) (response *DescribeScdnTopDataResponse, err error) {
    if request == nil {
        request = NewDescribeScdnTopDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScdnTopData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScdnTopDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopDataRequest() (request *DescribeTopDataRequest) {
    request = &DescribeTopDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeTopData")
    
    
    return
}

func NewDescribeTopDataResponse() (response *DescribeTopDataResponse) {
    response = &DescribeTopDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTopData
// DescribeTopData 通过入参 Metric 和 Filter 组合不同，可以查询以下排序数据：
//
// 
//
// + 依据总流量、总请求数对访问 IP 排序，从大至小返回 TOP 100 IP
//
// + 依据总流量、总请求数对访问 Refer 排序，从大至小返回 TOP 100 Refer
//
// + 依据总流量、总请求数对访问 设备 排序，从大至小返回 设备类型
//
// + 依据总流量、总请求数对访问 操作系统 排序，从大至小返回 操作系统
//
// + 依据总流量、总请求数对访问 浏览器 排序，从大至小返回 浏览器
//
// 
//
// 注意：
//
// + 仅支持 90 天内数据查询，且从2021年09月20日开始有数据
//
// + 本接口为beta版，尚未正式全量发布
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYPARAMERROR = "InternalError.CdnQueryParamError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDFILTER = "InvalidParameter.CdnStatInvalidFilter"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNPROJECTNOTEXISTS = "ResourceNotFound.CdnProjectNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeTopData(request *DescribeTopDataRequest) (response *DescribeTopDataResponse, err error) {
    return c.DescribeTopDataWithContext(context.Background(), request)
}

// DescribeTopData
// DescribeTopData 通过入参 Metric 和 Filter 组合不同，可以查询以下排序数据：
//
// 
//
// + 依据总流量、总请求数对访问 IP 排序，从大至小返回 TOP 100 IP
//
// + 依据总流量、总请求数对访问 Refer 排序，从大至小返回 TOP 100 Refer
//
// + 依据总流量、总请求数对访问 设备 排序，从大至小返回 设备类型
//
// + 依据总流量、总请求数对访问 操作系统 排序，从大至小返回 操作系统
//
// + 依据总流量、总请求数对访问 浏览器 排序，从大至小返回 浏览器
//
// 
//
// 注意：
//
// + 仅支持 90 天内数据查询，且从2021年09月20日开始有数据
//
// + 本接口为beta版，尚未正式全量发布
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYPARAMERROR = "InternalError.CdnQueryParamError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDFILTER = "InvalidParameter.CdnStatInvalidFilter"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNPROJECTNOTEXISTS = "ResourceNotFound.CdnProjectNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeTopDataWithContext(ctx context.Context, request *DescribeTopDataRequest) (response *DescribeTopDataResponse, err error) {
    if request == nil {
        request = NewDescribeTopDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrafficPackagesRequest() (request *DescribeTrafficPackagesRequest) {
    request = &DescribeTrafficPackagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeTrafficPackages")
    
    
    return
}

func NewDescribeTrafficPackagesResponse() (response *DescribeTrafficPackagesResponse) {
    response = &DescribeTrafficPackagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTrafficPackages
// DescribeTrafficPackages 用于查询 CDN 流量包详情。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeTrafficPackages(request *DescribeTrafficPackagesRequest) (response *DescribeTrafficPackagesResponse, err error) {
    return c.DescribeTrafficPackagesWithContext(context.Background(), request)
}

// DescribeTrafficPackages
// DescribeTrafficPackages 用于查询 CDN 流量包详情。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeTrafficPackagesWithContext(ctx context.Context, request *DescribeTrafficPackagesRequest) (response *DescribeTrafficPackagesResponse, err error) {
    if request == nil {
        request = NewDescribeTrafficPackagesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrafficPackages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrafficPackagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUrlViolationsRequest() (request *DescribeUrlViolationsRequest) {
    request = &DescribeUrlViolationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeUrlViolations")
    
    
    return
}

func NewDescribeUrlViolationsResponse() (response *DescribeUrlViolationsResponse) {
    response = &DescribeUrlViolationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUrlViolations
// DescribeUrlViolations 用于查询被 CDN 系统扫描到的域名违规 URL 列表及当前状态。
//
// 对应内容分发网络控制台【图片鉴黄】页面。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeUrlViolations(request *DescribeUrlViolationsRequest) (response *DescribeUrlViolationsResponse, err error) {
    return c.DescribeUrlViolationsWithContext(context.Background(), request)
}

// DescribeUrlViolations
// DescribeUrlViolations 用于查询被 CDN 系统扫描到的域名违规 URL 列表及当前状态。
//
// 对应内容分发网络控制台【图片鉴黄】页面。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeUrlViolationsWithContext(ctx context.Context, request *DescribeUrlViolationsRequest) (response *DescribeUrlViolationsResponse, err error) {
    if request == nil {
        request = NewDescribeUrlViolationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUrlViolations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUrlViolationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWafDataRequest() (request *DescribeWafDataRequest) {
    request = &DescribeWafDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeWafData")
    
    
    return
}

func NewDescribeWafDataResponse() (response *DescribeWafDataResponse) {
    response = &DescribeWafDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWafData
// Waf统计数据查询
//
// 可能返回的错误码:
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMDBERROR = "InternalError.SystemDBError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDFILTER = "InvalidParameter.CdnStatInvalidFilter"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_DOMAINEMPTY = "UnauthorizedOperation.DomainEmpty"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeWafData(request *DescribeWafDataRequest) (response *DescribeWafDataResponse, err error) {
    return c.DescribeWafDataWithContext(context.Background(), request)
}

// DescribeWafData
// Waf统计数据查询
//
// 可能返回的错误码:
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMDBERROR = "InternalError.SystemDBError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDFILTER = "InvalidParameter.CdnStatInvalidFilter"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_DOMAINEMPTY = "UnauthorizedOperation.DomainEmpty"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeWafDataWithContext(ctx context.Context, request *DescribeWafDataRequest) (response *DescribeWafDataResponse, err error) {
    if request == nil {
        request = NewDescribeWafDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWafData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWafDataResponse()
    err = c.Send(request, response)
    return
}

func NewDisableCachesRequest() (request *DisableCachesRequest) {
    request = &DisableCachesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DisableCaches")
    
    
    return
}

func NewDisableCachesResponse() (response *DisableCachesResponse) {
    response = &DisableCachesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableCaches
// DisableCaches 用于禁用 CDN 上指定 URL 的访问，禁用完成后，中国境内访问会直接返回 403。（注：接口尚在内测中，暂未全量开放；封禁URL并非无限期永久封禁）
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  LIMITEXCEEDED_CDNPUSHEXCEEDDAYLIMIT = "LimitExceeded.CdnPushExceedDayLimit"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DisableCaches(request *DisableCachesRequest) (response *DisableCachesResponse, err error) {
    return c.DisableCachesWithContext(context.Background(), request)
}

// DisableCaches
// DisableCaches 用于禁用 CDN 上指定 URL 的访问，禁用完成后，中国境内访问会直接返回 403。（注：接口尚在内测中，暂未全量开放；封禁URL并非无限期永久封禁）
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  LIMITEXCEEDED_CDNPUSHEXCEEDDAYLIMIT = "LimitExceeded.CdnPushExceedDayLimit"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DisableCachesWithContext(ctx context.Context, request *DisableCachesRequest) (response *DisableCachesResponse, err error) {
    if request == nil {
        request = NewDisableCachesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableCaches require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableCachesResponse()
    err = c.Send(request, response)
    return
}

func NewDisableClsLogTopicRequest() (request *DisableClsLogTopicRequest) {
    request = &DisableClsLogTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DisableClsLogTopic")
    
    
    return
}

func NewDisableClsLogTopicResponse() (response *DisableClsLogTopicResponse) {
    response = &DisableClsLogTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableClsLogTopic
// DisableClsLogTopic 用于停止日志主题投递。注意：停止后，所有绑定该日志主题域名的日志将不再继续投递至该主题，已经投递的日志将会继续保留。生效时间约为 5~15 分钟。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DisableClsLogTopic(request *DisableClsLogTopicRequest) (response *DisableClsLogTopicResponse, err error) {
    return c.DisableClsLogTopicWithContext(context.Background(), request)
}

// DisableClsLogTopic
// DisableClsLogTopic 用于停止日志主题投递。注意：停止后，所有绑定该日志主题域名的日志将不再继续投递至该主题，已经投递的日志将会继续保留。生效时间约为 5~15 分钟。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DisableClsLogTopicWithContext(ctx context.Context, request *DisableClsLogTopicRequest) (response *DisableClsLogTopicResponse, err error) {
    if request == nil {
        request = NewDisableClsLogTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableClsLogTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableClsLogTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDuplicateDomainConfigRequest() (request *DuplicateDomainConfigRequest) {
    request = &DuplicateDomainConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DuplicateDomainConfig")
    
    
    return
}

func NewDuplicateDomainConfigResponse() (response *DuplicateDomainConfigResponse) {
    response = &DuplicateDomainConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DuplicateDomainConfig
// 拷贝参考域名的配置至新域名。暂不支持自有证书以及定制化配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CDNCERTINFONOTFOUND = "InvalidParameter.CdnCertInfoNotFound"
//  INVALIDPARAMETER_CDNCONFIGINVALIDHOST = "InvalidParameter.CdnConfigInvalidHost"
//  INVALIDPARAMETER_CDNHOSTINTERNALHOST = "InvalidParameter.CdnHostInternalHost"
//  INVALIDPARAMETER_CDNKEYRULESEXCLUDECUSTOMREQUIRESFULLLEGO = "InvalidParameter.CdnKeyRulesExcludeCustomRequiresFullLego"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CAMTAGQUOTAEXCEEDLIMIT = "LimitExceeded.CamTagQuotaExceedLimit"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  LIMITEXCEEDED_CDNUSERTOOMANYHOSTS = "LimitExceeded.CdnUserTooManyHosts"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE_CDNCONFLICTHOSTEXISTS = "ResourceInUse.CdnConflictHostExists"
//  RESOURCEINUSE_CDNHOSTEXISTS = "ResourceInUse.CdnHostExists"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCEINUSE_TCBHOSTEXISTS = "ResourceInUse.TcbHostExists"
//  RESOURCENOTFOUND_CAMTAGKEYNOTEXIST = "ResourceNotFound.CamTagKeyNotExist"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CDNUSERTOOMANYHOSTS = "ResourceNotFound.CdnUserTooManyHosts"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  RESOURCEUNAVAILABLE_CDNHOSTISMALICIOUS = "ResourceUnavailable.CdnHostIsMalicious"
//  RESOURCEUNAVAILABLE_CDNHOSTNOICP = "ResourceUnavailable.CdnHostNoIcp"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNDOMAINRECORDNOTVERIFIED = "UnauthorizedOperation.CdnDomainRecordNotVerified"
//  UNAUTHORIZEDOPERATION_CDNHOSTEXISTSININTERNAL = "UnauthorizedOperation.CdnHostExistsInInternal"
//  UNAUTHORIZEDOPERATION_CDNHOSTISOWNEDBYOTHER = "UnauthorizedOperation.CdnHostIsOwnedByOther"
//  UNAUTHORIZEDOPERATION_CDNHOSTISTOAPPLYHOST = "UnauthorizedOperation.CdnHostIsToApplyHost"
//  UNAUTHORIZEDOPERATION_CDNHOSTISUSEDBYOTHER = "UnauthorizedOperation.CdnHostIsUsedByOther"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNINVALIDUSERSTATUS = "UnauthorizedOperation.CdnInvalidUserStatus"
//  UNAUTHORIZEDOPERATION_CDNTAGUNAUTHORIZED = "UnauthorizedOperation.CdnTagUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERISISOLATED = "UnauthorizedOperation.CdnUserIsIsolated"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DuplicateDomainConfig(request *DuplicateDomainConfigRequest) (response *DuplicateDomainConfigResponse, err error) {
    return c.DuplicateDomainConfigWithContext(context.Background(), request)
}

// DuplicateDomainConfig
// 拷贝参考域名的配置至新域名。暂不支持自有证书以及定制化配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CDNCERTINFONOTFOUND = "InvalidParameter.CdnCertInfoNotFound"
//  INVALIDPARAMETER_CDNCONFIGINVALIDHOST = "InvalidParameter.CdnConfigInvalidHost"
//  INVALIDPARAMETER_CDNHOSTINTERNALHOST = "InvalidParameter.CdnHostInternalHost"
//  INVALIDPARAMETER_CDNKEYRULESEXCLUDECUSTOMREQUIRESFULLLEGO = "InvalidParameter.CdnKeyRulesExcludeCustomRequiresFullLego"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CAMTAGQUOTAEXCEEDLIMIT = "LimitExceeded.CamTagQuotaExceedLimit"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  LIMITEXCEEDED_CDNUSERTOOMANYHOSTS = "LimitExceeded.CdnUserTooManyHosts"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE_CDNCONFLICTHOSTEXISTS = "ResourceInUse.CdnConflictHostExists"
//  RESOURCEINUSE_CDNHOSTEXISTS = "ResourceInUse.CdnHostExists"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCEINUSE_TCBHOSTEXISTS = "ResourceInUse.TcbHostExists"
//  RESOURCENOTFOUND_CAMTAGKEYNOTEXIST = "ResourceNotFound.CamTagKeyNotExist"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CDNUSERTOOMANYHOSTS = "ResourceNotFound.CdnUserTooManyHosts"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  RESOURCEUNAVAILABLE_CDNHOSTISMALICIOUS = "ResourceUnavailable.CdnHostIsMalicious"
//  RESOURCEUNAVAILABLE_CDNHOSTNOICP = "ResourceUnavailable.CdnHostNoIcp"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNDOMAINRECORDNOTVERIFIED = "UnauthorizedOperation.CdnDomainRecordNotVerified"
//  UNAUTHORIZEDOPERATION_CDNHOSTEXISTSININTERNAL = "UnauthorizedOperation.CdnHostExistsInInternal"
//  UNAUTHORIZEDOPERATION_CDNHOSTISOWNEDBYOTHER = "UnauthorizedOperation.CdnHostIsOwnedByOther"
//  UNAUTHORIZEDOPERATION_CDNHOSTISTOAPPLYHOST = "UnauthorizedOperation.CdnHostIsToApplyHost"
//  UNAUTHORIZEDOPERATION_CDNHOSTISUSEDBYOTHER = "UnauthorizedOperation.CdnHostIsUsedByOther"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNINVALIDUSERSTATUS = "UnauthorizedOperation.CdnInvalidUserStatus"
//  UNAUTHORIZEDOPERATION_CDNTAGUNAUTHORIZED = "UnauthorizedOperation.CdnTagUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERISISOLATED = "UnauthorizedOperation.CdnUserIsIsolated"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DuplicateDomainConfigWithContext(ctx context.Context, request *DuplicateDomainConfigRequest) (response *DuplicateDomainConfigResponse, err error) {
    if request == nil {
        request = NewDuplicateDomainConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DuplicateDomainConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDuplicateDomainConfigResponse()
    err = c.Send(request, response)
    return
}

func NewEnableCachesRequest() (request *EnableCachesRequest) {
    request = &EnableCachesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "EnableCaches")
    
    
    return
}

func NewEnableCachesResponse() (response *EnableCachesResponse) {
    response = &EnableCachesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableCaches
// EnableCaches 用于解禁手工封禁的 URL，解禁成功后，全网生效时间约 5~10 分钟。（接口尚在内测中，暂未全量开放使用）
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) EnableCaches(request *EnableCachesRequest) (response *EnableCachesResponse, err error) {
    return c.EnableCachesWithContext(context.Background(), request)
}

// EnableCaches
// EnableCaches 用于解禁手工封禁的 URL，解禁成功后，全网生效时间约 5~10 分钟。（接口尚在内测中，暂未全量开放使用）
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) EnableCachesWithContext(ctx context.Context, request *EnableCachesRequest) (response *EnableCachesResponse, err error) {
    if request == nil {
        request = NewEnableCachesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableCaches require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableCachesResponse()
    err = c.Send(request, response)
    return
}

func NewEnableClsLogTopicRequest() (request *EnableClsLogTopicRequest) {
    request = &EnableClsLogTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "EnableClsLogTopic")
    
    
    return
}

func NewEnableClsLogTopicResponse() (response *EnableClsLogTopicResponse) {
    response = &EnableClsLogTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableClsLogTopic
// EnableClsLogTopic 用于启动日志主题投递。注意：启动后，所有绑定该日志主题域名的日志将继续投递至该主题。生效时间约为 5~15 分钟。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) EnableClsLogTopic(request *EnableClsLogTopicRequest) (response *EnableClsLogTopicResponse, err error) {
    return c.EnableClsLogTopicWithContext(context.Background(), request)
}

// EnableClsLogTopic
// EnableClsLogTopic 用于启动日志主题投递。注意：启动后，所有绑定该日志主题域名的日志将继续投递至该主题。生效时间约为 5~15 分钟。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) EnableClsLogTopicWithContext(ctx context.Context, request *EnableClsLogTopicRequest) (response *EnableClsLogTopicResponse, err error) {
    if request == nil {
        request = NewEnableClsLogTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableClsLogTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableClsLogTopicResponse()
    err = c.Send(request, response)
    return
}

func NewGetDisableRecordsRequest() (request *GetDisableRecordsRequest) {
    request = &GetDisableRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "GetDisableRecords")
    
    
    return
}

func NewGetDisableRecordsResponse() (response *GetDisableRecordsResponse) {
    response = &GetDisableRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetDisableRecords
// GetDisableRecords 用于查询资源禁用历史，及 URL 当前状态。（接口尚在内测中，暂未全量开放使用）
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) GetDisableRecords(request *GetDisableRecordsRequest) (response *GetDisableRecordsResponse, err error) {
    return c.GetDisableRecordsWithContext(context.Background(), request)
}

// GetDisableRecords
// GetDisableRecords 用于查询资源禁用历史，及 URL 当前状态。（接口尚在内测中，暂未全量开放使用）
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) GetDisableRecordsWithContext(ctx context.Context, request *GetDisableRecordsRequest) (response *GetDisableRecordsResponse, err error) {
    if request == nil {
        request = NewGetDisableRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDisableRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDisableRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewListClsLogTopicsRequest() (request *ListClsLogTopicsRequest) {
    request = &ListClsLogTopicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "ListClsLogTopics")
    
    
    return
}

func NewListClsLogTopicsResponse() (response *ListClsLogTopicsResponse) {
    response = &ListClsLogTopicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListClsLogTopics
// ListClsLogTopics 用于显示日志主题列表。注意：一个日志集下至多含10个日志主题。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSSERVICENOTACTIVATED = "UnauthorizedOperation.ClsServiceNotActivated"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) ListClsLogTopics(request *ListClsLogTopicsRequest) (response *ListClsLogTopicsResponse, err error) {
    return c.ListClsLogTopicsWithContext(context.Background(), request)
}

// ListClsLogTopics
// ListClsLogTopics 用于显示日志主题列表。注意：一个日志集下至多含10个日志主题。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSSERVICENOTACTIVATED = "UnauthorizedOperation.ClsServiceNotActivated"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) ListClsLogTopicsWithContext(ctx context.Context, request *ListClsLogTopicsRequest) (response *ListClsLogTopicsResponse, err error) {
    if request == nil {
        request = NewListClsLogTopicsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListClsLogTopics require credential")
    }

    request.SetContext(ctx)
    
    response = NewListClsLogTopicsResponse()
    err = c.Send(request, response)
    return
}

func NewListClsTopicDomainsRequest() (request *ListClsTopicDomainsRequest) {
    request = &ListClsTopicDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "ListClsTopicDomains")
    
    
    return
}

func NewListClsTopicDomainsResponse() (response *ListClsTopicDomainsResponse) {
    response = &ListClsTopicDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListClsTopicDomains
// ListClsTopicDomains 用于获取某日志主题下绑定的域名列表。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) ListClsTopicDomains(request *ListClsTopicDomainsRequest) (response *ListClsTopicDomainsResponse, err error) {
    return c.ListClsTopicDomainsWithContext(context.Background(), request)
}

// ListClsTopicDomains
// ListClsTopicDomains 用于获取某日志主题下绑定的域名列表。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) ListClsTopicDomainsWithContext(ctx context.Context, request *ListClsTopicDomainsRequest) (response *ListClsTopicDomainsResponse, err error) {
    if request == nil {
        request = NewListClsTopicDomainsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListClsTopicDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewListClsTopicDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewListDiagnoseReportRequest() (request *ListDiagnoseReportRequest) {
    request = &ListDiagnoseReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "ListDiagnoseReport")
    
    
    return
}

func NewListDiagnoseReportResponse() (response *ListDiagnoseReportResponse) {
    response = &ListDiagnoseReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListDiagnoseReport
// ListDiagnoseReport 用于获取用户诊断URL访问后各个子任务的简要详情， <font color=red>将于 **2023年5月31日** 下线</font><br>
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) ListDiagnoseReport(request *ListDiagnoseReportRequest) (response *ListDiagnoseReportResponse, err error) {
    return c.ListDiagnoseReportWithContext(context.Background(), request)
}

// ListDiagnoseReport
// ListDiagnoseReport 用于获取用户诊断URL访问后各个子任务的简要详情， <font color=red>将于 **2023年5月31日** 下线</font><br>
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) ListDiagnoseReportWithContext(ctx context.Context, request *ListDiagnoseReportRequest) (response *ListDiagnoseReportResponse, err error) {
    if request == nil {
        request = NewListDiagnoseReportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListDiagnoseReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewListDiagnoseReportResponse()
    err = c.Send(request, response)
    return
}

func NewListScdnDomainsRequest() (request *ListScdnDomainsRequest) {
    request = &ListScdnDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "ListScdnDomains")
    
    
    return
}

func NewListScdnDomainsResponse() (response *ListScdnDomainsResponse) {
    response = &ListScdnDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListScdnDomains
// ListScdnDomains 用于查询 SCDN 安全加速域名列表，及域名基本配置信息
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
func (c *Client) ListScdnDomains(request *ListScdnDomainsRequest) (response *ListScdnDomainsResponse, err error) {
    return c.ListScdnDomainsWithContext(context.Background(), request)
}

// ListScdnDomains
// ListScdnDomains 用于查询 SCDN 安全加速域名列表，及域名基本配置信息
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
func (c *Client) ListScdnDomainsWithContext(ctx context.Context, request *ListScdnDomainsRequest) (response *ListScdnDomainsResponse, err error) {
    if request == nil {
        request = NewListScdnDomainsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListScdnDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewListScdnDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewListScdnLogTasksRequest() (request *ListScdnLogTasksRequest) {
    request = &ListScdnLogTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "ListScdnLogTasks")
    
    
    return
}

func NewListScdnLogTasksResponse() (response *ListScdnLogTasksResponse) {
    response = &ListScdnLogTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListScdnLogTasks
// ListScdnLogTasks 用于查询SCDN日志下载任务列表,以及展示下载任务基本信息
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ListScdnLogTasks(request *ListScdnLogTasksRequest) (response *ListScdnLogTasksResponse, err error) {
    return c.ListScdnLogTasksWithContext(context.Background(), request)
}

// ListScdnLogTasks
// ListScdnLogTasks 用于查询SCDN日志下载任务列表,以及展示下载任务基本信息
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ListScdnLogTasksWithContext(ctx context.Context, request *ListScdnLogTasksRequest) (response *ListScdnLogTasksResponse, err error) {
    if request == nil {
        request = NewListScdnLogTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListScdnLogTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListScdnLogTasksResponse()
    err = c.Send(request, response)
    return
}

func NewListScdnTopBotDataRequest() (request *ListScdnTopBotDataRequest) {
    request = &ListScdnTopBotDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "ListScdnTopBotData")
    
    
    return
}

func NewListScdnTopBotDataResponse() (response *ListScdnTopBotDataResponse) {
    response = &ListScdnTopBotDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListScdnTopBotData
// 获取Bot攻击的Top数据列表
//
// 可能返回的错误码:
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_SYSTEMDBERROR = "InternalError.SystemDBError"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_DOMAINEMPTY = "UnauthorizedOperation.DomainEmpty"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ListScdnTopBotData(request *ListScdnTopBotDataRequest) (response *ListScdnTopBotDataResponse, err error) {
    return c.ListScdnTopBotDataWithContext(context.Background(), request)
}

// ListScdnTopBotData
// 获取Bot攻击的Top数据列表
//
// 可能返回的错误码:
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_SYSTEMDBERROR = "InternalError.SystemDBError"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_DOMAINEMPTY = "UnauthorizedOperation.DomainEmpty"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ListScdnTopBotDataWithContext(ctx context.Context, request *ListScdnTopBotDataRequest) (response *ListScdnTopBotDataResponse, err error) {
    if request == nil {
        request = NewListScdnTopBotDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListScdnTopBotData require credential")
    }

    request.SetContext(ctx)
    
    response = NewListScdnTopBotDataResponse()
    err = c.Send(request, response)
    return
}

func NewListTopBotDataRequest() (request *ListTopBotDataRequest) {
    request = &ListTopBotDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "ListTopBotData")
    
    
    return
}

func NewListTopBotDataResponse() (response *ListTopBotDataResponse) {
    response = &ListTopBotDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListTopBotData
// 获取Bot攻击的Top信息
//
// 可能返回的错误码:
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
func (c *Client) ListTopBotData(request *ListTopBotDataRequest) (response *ListTopBotDataResponse, err error) {
    return c.ListTopBotDataWithContext(context.Background(), request)
}

// ListTopBotData
// 获取Bot攻击的Top信息
//
// 可能返回的错误码:
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
func (c *Client) ListTopBotDataWithContext(ctx context.Context, request *ListTopBotDataRequest) (response *ListTopBotDataResponse, err error) {
    if request == nil {
        request = NewListTopBotDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTopBotData require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTopBotDataResponse()
    err = c.Send(request, response)
    return
}

func NewListTopCcDataRequest() (request *ListTopCcDataRequest) {
    request = &ListTopCcDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "ListTopCcData")
    
    
    return
}

func NewListTopCcDataResponse() (response *ListTopCcDataResponse) {
    response = &ListTopCcDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListTopCcData
// 获取CC攻击Top数据
//
// 可能返回的错误码:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_DOMAINEMPTY = "UnauthorizedOperation.DomainEmpty"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ListTopCcData(request *ListTopCcDataRequest) (response *ListTopCcDataResponse, err error) {
    return c.ListTopCcDataWithContext(context.Background(), request)
}

// ListTopCcData
// 获取CC攻击Top数据
//
// 可能返回的错误码:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_DOMAINEMPTY = "UnauthorizedOperation.DomainEmpty"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ListTopCcDataWithContext(ctx context.Context, request *ListTopCcDataRequest) (response *ListTopCcDataResponse, err error) {
    if request == nil {
        request = NewListTopCcDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTopCcData require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTopCcDataResponse()
    err = c.Send(request, response)
    return
}

func NewListTopClsLogDataRequest() (request *ListTopClsLogDataRequest) {
    request = &ListTopClsLogDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "ListTopClsLogData")
    
    
    return
}

func NewListTopClsLogDataResponse() (response *ListTopClsLogDataResponse) {
    response = &ListTopClsLogDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListTopClsLogData
// 通过CLS日志计算Top信息。支持近7天的日志数据。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSINVALIDPARAM = "InvalidParameter.ClsInvalidParam"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) ListTopClsLogData(request *ListTopClsLogDataRequest) (response *ListTopClsLogDataResponse, err error) {
    return c.ListTopClsLogDataWithContext(context.Background(), request)
}

// ListTopClsLogData
// 通过CLS日志计算Top信息。支持近7天的日志数据。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSINVALIDPARAM = "InvalidParameter.ClsInvalidParam"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) ListTopClsLogDataWithContext(ctx context.Context, request *ListTopClsLogDataRequest) (response *ListTopClsLogDataResponse, err error) {
    if request == nil {
        request = NewListTopClsLogDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTopClsLogData require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTopClsLogDataResponse()
    err = c.Send(request, response)
    return
}

func NewListTopDDoSDataRequest() (request *ListTopDDoSDataRequest) {
    request = &ListTopDDoSDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "ListTopDDoSData")
    
    
    return
}

func NewListTopDDoSDataResponse() (response *ListTopDDoSDataResponse) {
    response = &ListTopDDoSDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListTopDDoSData
// 获取DDoS攻击Top数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ListTopDDoSData(request *ListTopDDoSDataRequest) (response *ListTopDDoSDataResponse, err error) {
    return c.ListTopDDoSDataWithContext(context.Background(), request)
}

// ListTopDDoSData
// 获取DDoS攻击Top数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ListTopDDoSDataWithContext(ctx context.Context, request *ListTopDDoSDataRequest) (response *ListTopDDoSDataResponse, err error) {
    if request == nil {
        request = NewListTopDDoSDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTopDDoSData require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTopDDoSDataResponse()
    err = c.Send(request, response)
    return
}

func NewListTopDataRequest() (request *ListTopDataRequest) {
    request = &ListTopDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "ListTopData")
    
    
    return
}

func NewListTopDataResponse() (response *ListTopDataResponse) {
    response = &ListTopDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListTopData
// ListTopData 通过入参 Metric 和 Filter 组合不同，可以查询以下排序数据：
//
// 
//
// + 依据总流量、总请求数对访问 URL 排序，从大至小返回 TOP 1000 URL
//
// + 依据总流量、总请求数对客户端省份排序，从大至小返回省份列表
//
// + 依据总流量、总请求数对客户端运营商排序，从大至小返回运营商列表
//
// + 依据总流量、峰值带宽、总请求数、平均命中率、2XX/3XX/4XX/5XX 状态码对域名排序，从大至小返回域名列表
//
// + 依据总回源流量、回源峰值带宽、总回源请求数、平均回源失败率、2XX/3XX/4XX/5XX 回源状态码对域名排序，从大至小返回域名列表
//
// 
//
// 注意：仅支持 90 天内数据查询
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_TAGSYSTEMERROR = "InternalError.TagSystemError"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDFILTER = "InvalidParameter.CdnStatInvalidFilter"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNPROJECTNOTEXISTS = "ResourceNotFound.CdnProjectNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) ListTopData(request *ListTopDataRequest) (response *ListTopDataResponse, err error) {
    return c.ListTopDataWithContext(context.Background(), request)
}

// ListTopData
// ListTopData 通过入参 Metric 和 Filter 组合不同，可以查询以下排序数据：
//
// 
//
// + 依据总流量、总请求数对访问 URL 排序，从大至小返回 TOP 1000 URL
//
// + 依据总流量、总请求数对客户端省份排序，从大至小返回省份列表
//
// + 依据总流量、总请求数对客户端运营商排序，从大至小返回运营商列表
//
// + 依据总流量、峰值带宽、总请求数、平均命中率、2XX/3XX/4XX/5XX 状态码对域名排序，从大至小返回域名列表
//
// + 依据总回源流量、回源峰值带宽、总回源请求数、平均回源失败率、2XX/3XX/4XX/5XX 回源状态码对域名排序，从大至小返回域名列表
//
// 
//
// 注意：仅支持 90 天内数据查询
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_TAGSYSTEMERROR = "InternalError.TagSystemError"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDFILTER = "InvalidParameter.CdnStatInvalidFilter"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNPROJECTNOTEXISTS = "ResourceNotFound.CdnProjectNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) ListTopDataWithContext(ctx context.Context, request *ListTopDataRequest) (response *ListTopDataResponse, err error) {
    if request == nil {
        request = NewListTopDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTopData require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTopDataResponse()
    err = c.Send(request, response)
    return
}

func NewListTopWafDataRequest() (request *ListTopWafDataRequest) {
    request = &ListTopWafDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "ListTopWafData")
    
    
    return
}

func NewListTopWafDataResponse() (response *ListTopWafDataResponse) {
    response = &ListTopWafDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListTopWafData
// 获取Waf攻击Top数据
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDFILTER = "InvalidParameter.CdnStatInvalidFilter"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_DOMAINEMPTY = "UnauthorizedOperation.DomainEmpty"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ListTopWafData(request *ListTopWafDataRequest) (response *ListTopWafDataResponse, err error) {
    return c.ListTopWafDataWithContext(context.Background(), request)
}

// ListTopWafData
// 获取Waf攻击Top数据
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDFILTER = "InvalidParameter.CdnStatInvalidFilter"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_DOMAINEMPTY = "UnauthorizedOperation.DomainEmpty"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ListTopWafDataWithContext(ctx context.Context, request *ListTopWafDataRequest) (response *ListTopWafDataResponse, err error) {
    if request == nil {
        request = NewListTopWafDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTopWafData require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTopWafDataResponse()
    err = c.Send(request, response)
    return
}

func NewManageClsTopicDomainsRequest() (request *ManageClsTopicDomainsRequest) {
    request = &ManageClsTopicDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "ManageClsTopicDomains")
    
    
    return
}

func NewManageClsTopicDomainsResponse() (response *ManageClsTopicDomainsResponse) {
    response = &ManageClsTopicDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ManageClsTopicDomains
// ManageClsTopicDomains 用于管理某日志主题下绑定的域名列表。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) ManageClsTopicDomains(request *ManageClsTopicDomainsRequest) (response *ManageClsTopicDomainsResponse, err error) {
    return c.ManageClsTopicDomainsWithContext(context.Background(), request)
}

// ManageClsTopicDomains
// ManageClsTopicDomains 用于管理某日志主题下绑定的域名列表。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) ManageClsTopicDomainsWithContext(ctx context.Context, request *ManageClsTopicDomainsRequest) (response *ManageClsTopicDomainsResponse, err error) {
    if request == nil {
        request = NewManageClsTopicDomainsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ManageClsTopicDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewManageClsTopicDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPurgeFetchTaskStatusRequest() (request *ModifyPurgeFetchTaskStatusRequest) {
    request = &ModifyPurgeFetchTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "ModifyPurgeFetchTaskStatus")
    
    
    return
}

func NewModifyPurgeFetchTaskStatusResponse() (response *ModifyPurgeFetchTaskStatusResponse) {
    response = &ModifyPurgeFetchTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPurgeFetchTaskStatus
// ModifyPurgeFetchTaskStatus 用于上报定时刷新预热任务执行状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifyPurgeFetchTaskStatus(request *ModifyPurgeFetchTaskStatusRequest) (response *ModifyPurgeFetchTaskStatusResponse, err error) {
    return c.ModifyPurgeFetchTaskStatusWithContext(context.Background(), request)
}

// ModifyPurgeFetchTaskStatus
// ModifyPurgeFetchTaskStatus 用于上报定时刷新预热任务执行状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifyPurgeFetchTaskStatusWithContext(ctx context.Context, request *ModifyPurgeFetchTaskStatusRequest) (response *ModifyPurgeFetchTaskStatusResponse, err error) {
    if request == nil {
        request = NewModifyPurgeFetchTaskStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPurgeFetchTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPurgeFetchTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewPurgePathCacheRequest() (request *PurgePathCacheRequest) {
    request = &PurgePathCacheRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "PurgePathCache")
    
    
    return
}

func NewPurgePathCacheResponse() (response *PurgePathCacheResponse) {
    response = &PurgePathCacheResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PurgePathCache
// PurgePathCache 用于批量提交目录刷新，根据域名的加速区域进行对应区域的刷新。
//
// 默认情况下境内、境外加速区域每日目录刷新额度为各 100 条，每次最多可提交 500 条。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYPARAMERROR = "InternalError.CdnQueryParamError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNPURGEWILDCARDNOTALLOWED = "InvalidParameter.CdnPurgeWildcardNotAllowed"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNURLEXCEEDLENGTH = "InvalidParameter.CdnUrlExceedLength"
//  LIMITEXCEEDED_CDNPURGEPATHEXCEEDBATCHLIMIT = "LimitExceeded.CdnPurgePathExceedBatchLimit"
//  LIMITEXCEEDED_CDNPURGEPATHEXCEEDDAYLIMIT = "LimitExceeded.CdnPurgePathExceedDayLimit"
//  LIMITEXCEEDED_CDNPURGEURLEXCEEDBATCHLIMIT = "LimitExceeded.CdnPurgeUrlExceedBatchLimit"
//  LIMITEXCEEDED_CDNPURGEURLEXCEEDDAYLIMIT = "LimitExceeded.CdnPurgeUrlExceedDayLimit"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) PurgePathCache(request *PurgePathCacheRequest) (response *PurgePathCacheResponse, err error) {
    return c.PurgePathCacheWithContext(context.Background(), request)
}

// PurgePathCache
// PurgePathCache 用于批量提交目录刷新，根据域名的加速区域进行对应区域的刷新。
//
// 默认情况下境内、境外加速区域每日目录刷新额度为各 100 条，每次最多可提交 500 条。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYPARAMERROR = "InternalError.CdnQueryParamError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNPURGEWILDCARDNOTALLOWED = "InvalidParameter.CdnPurgeWildcardNotAllowed"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNURLEXCEEDLENGTH = "InvalidParameter.CdnUrlExceedLength"
//  LIMITEXCEEDED_CDNPURGEPATHEXCEEDBATCHLIMIT = "LimitExceeded.CdnPurgePathExceedBatchLimit"
//  LIMITEXCEEDED_CDNPURGEPATHEXCEEDDAYLIMIT = "LimitExceeded.CdnPurgePathExceedDayLimit"
//  LIMITEXCEEDED_CDNPURGEURLEXCEEDBATCHLIMIT = "LimitExceeded.CdnPurgeUrlExceedBatchLimit"
//  LIMITEXCEEDED_CDNPURGEURLEXCEEDDAYLIMIT = "LimitExceeded.CdnPurgeUrlExceedDayLimit"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) PurgePathCacheWithContext(ctx context.Context, request *PurgePathCacheRequest) (response *PurgePathCacheResponse, err error) {
    if request == nil {
        request = NewPurgePathCacheRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PurgePathCache require credential")
    }

    request.SetContext(ctx)
    
    response = NewPurgePathCacheResponse()
    err = c.Send(request, response)
    return
}

func NewPurgeUrlsCacheRequest() (request *PurgeUrlsCacheRequest) {
    request = &PurgeUrlsCacheRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "PurgeUrlsCache")
    
    
    return
}

func NewPurgeUrlsCacheResponse() (response *PurgeUrlsCacheResponse) {
    response = &PurgeUrlsCacheResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PurgeUrlsCache
// PurgeUrlsCache 用于批量提交 URL 进行刷新，根据 URL 中域名的当前加速区域进行对应区域的刷新。
//
// 默认情况下境内、境外加速区域每日 URL 刷新额度各为 10000 条，每次最多可提交 1000 条。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYPARAMERROR = "InternalError.CdnQueryParamError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNPURGEWILDCARDNOTALLOWED = "InvalidParameter.CdnPurgeWildcardNotAllowed"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNURLEXCEEDLENGTH = "InvalidParameter.CdnUrlExceedLength"
//  LIMITEXCEEDED_CDNPURGEURLEXCEEDBATCHLIMIT = "LimitExceeded.CdnPurgeUrlExceedBatchLimit"
//  LIMITEXCEEDED_CDNPURGEURLEXCEEDDAYLIMIT = "LimitExceeded.CdnPurgeUrlExceedDayLimit"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISNOTONLINE = "ResourceUnavailable.CdnHostIsNotOnline"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) PurgeUrlsCache(request *PurgeUrlsCacheRequest) (response *PurgeUrlsCacheResponse, err error) {
    return c.PurgeUrlsCacheWithContext(context.Background(), request)
}

// PurgeUrlsCache
// PurgeUrlsCache 用于批量提交 URL 进行刷新，根据 URL 中域名的当前加速区域进行对应区域的刷新。
//
// 默认情况下境内、境外加速区域每日 URL 刷新额度各为 10000 条，每次最多可提交 1000 条。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYPARAMERROR = "InternalError.CdnQueryParamError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNPURGEWILDCARDNOTALLOWED = "InvalidParameter.CdnPurgeWildcardNotAllowed"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNURLEXCEEDLENGTH = "InvalidParameter.CdnUrlExceedLength"
//  LIMITEXCEEDED_CDNPURGEURLEXCEEDBATCHLIMIT = "LimitExceeded.CdnPurgeUrlExceedBatchLimit"
//  LIMITEXCEEDED_CDNPURGEURLEXCEEDDAYLIMIT = "LimitExceeded.CdnPurgeUrlExceedDayLimit"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISNOTONLINE = "ResourceUnavailable.CdnHostIsNotOnline"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) PurgeUrlsCacheWithContext(ctx context.Context, request *PurgeUrlsCacheRequest) (response *PurgeUrlsCacheResponse, err error) {
    if request == nil {
        request = NewPurgeUrlsCacheRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PurgeUrlsCache require credential")
    }

    request.SetContext(ctx)
    
    response = NewPurgeUrlsCacheResponse()
    err = c.Send(request, response)
    return
}

func NewPushUrlsCacheRequest() (request *PushUrlsCacheRequest) {
    request = &PushUrlsCacheRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "PushUrlsCache")
    
    
    return
}

func NewPushUrlsCacheResponse() (response *PushUrlsCacheResponse) {
    response = &PushUrlsCacheResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PushUrlsCache
// PushUrlsCache 用于将指定 URL 资源列表加载至 CDN 节点，支持指定加速区域预热。
//
// 默认情况下境内、境外每日预热 URL 限额为各 1000 条，每次最多可提交 500 条。注意：中国境外区域预热，资源默认加载至中国境外边缘节点，所产生的边缘层流量会计入计费流量。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNHOSTINVALIDMIDDLECONFIG = "InvalidParameter.CdnHostInvalidMiddleConfig"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNPUSHWILDCARDNOTALLOWED = "InvalidParameter.CdnPushWildcardNotAllowed"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  LIMITEXCEEDED_CDNPUSHEXCEEDBATCHLIMIT = "LimitExceeded.CdnPushExceedBatchLimit"
//  LIMITEXCEEDED_CDNPUSHEXCEEDDAYLIMIT = "LimitExceeded.CdnPushExceedDayLimit"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  RESOURCEUNAVAILABLE_CDNHOSTISNOTONLINE = "ResourceUnavailable.CdnHostIsNotOnline"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) PushUrlsCache(request *PushUrlsCacheRequest) (response *PushUrlsCacheResponse, err error) {
    return c.PushUrlsCacheWithContext(context.Background(), request)
}

// PushUrlsCache
// PushUrlsCache 用于将指定 URL 资源列表加载至 CDN 节点，支持指定加速区域预热。
//
// 默认情况下境内、境外每日预热 URL 限额为各 1000 条，每次最多可提交 500 条。注意：中国境外区域预热，资源默认加载至中国境外边缘节点，所产生的边缘层流量会计入计费流量。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNHOSTINVALIDMIDDLECONFIG = "InvalidParameter.CdnHostInvalidMiddleConfig"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNPUSHWILDCARDNOTALLOWED = "InvalidParameter.CdnPushWildcardNotAllowed"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  LIMITEXCEEDED_CDNPUSHEXCEEDBATCHLIMIT = "LimitExceeded.CdnPushExceedBatchLimit"
//  LIMITEXCEEDED_CDNPUSHEXCEEDDAYLIMIT = "LimitExceeded.CdnPushExceedDayLimit"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  RESOURCEUNAVAILABLE_CDNHOSTISNOTONLINE = "ResourceUnavailable.CdnHostIsNotOnline"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) PushUrlsCacheWithContext(ctx context.Context, request *PushUrlsCacheRequest) (response *PushUrlsCacheResponse, err error) {
    if request == nil {
        request = NewPushUrlsCacheRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PushUrlsCache require credential")
    }

    request.SetContext(ctx)
    
    response = NewPushUrlsCacheResponse()
    err = c.Send(request, response)
    return
}

func NewSearchClsLogRequest() (request *SearchClsLogRequest) {
    request = &SearchClsLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "SearchClsLog")
    
    
    return
}

func NewSearchClsLogResponse() (response *SearchClsLogResponse) {
    response = &SearchClsLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SearchClsLog
// SearchClsLog 用于 CLS 日志检索。支持检索今天，24小时（可选近7中的某一天），近7天的日志数据。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSINVALIDPARAM = "InvalidParameter.ClsInvalidParam"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) SearchClsLog(request *SearchClsLogRequest) (response *SearchClsLogResponse, err error) {
    return c.SearchClsLogWithContext(context.Background(), request)
}

// SearchClsLog
// SearchClsLog 用于 CLS 日志检索。支持检索今天，24小时（可选近7中的某一天），近7天的日志数据。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSINVALIDPARAM = "InvalidParameter.ClsInvalidParam"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) SearchClsLogWithContext(ctx context.Context, request *SearchClsLogRequest) (response *SearchClsLogResponse, err error) {
    if request == nil {
        request = NewSearchClsLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchClsLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchClsLogResponse()
    err = c.Send(request, response)
    return
}

func NewStartCdnDomainRequest() (request *StartCdnDomainRequest) {
    request = &StartCdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "StartCdnDomain")
    
    
    return
}

func NewStartCdnDomainResponse() (response *StartCdnDomainResponse) {
    response = &StartCdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartCdnDomain
// StartCdnDomain 用于启用已停用域名的加速服务
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNCERTINFONOTFOUND = "InvalidParameter.CdnCertInfoNotFound"
//  INVALIDPARAMETER_CDNHOSTHASSPECIALCONFIG = "InvalidParameter.CdnHostHasSpecialConfig"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  RESOURCEUNAVAILABLE_CDNHOSTNOICP = "ResourceUnavailable.CdnHostNoIcp"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTINICPBLACKLIST = "UnauthorizedOperation.CdnHostInIcpBlacklist"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERINVALIDCREDENTIAL = "UnauthorizedOperation.CdnUserInvalidCredential"
//  UNAUTHORIZEDOPERATION_CDNUSERISISOLATED = "UnauthorizedOperation.CdnUserIsIsolated"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNSUPPORTEDOPERATION_OPNOAUTH = "UnsupportedOperation.OpNoAuth"
func (c *Client) StartCdnDomain(request *StartCdnDomainRequest) (response *StartCdnDomainResponse, err error) {
    return c.StartCdnDomainWithContext(context.Background(), request)
}

// StartCdnDomain
// StartCdnDomain 用于启用已停用域名的加速服务
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNCERTINFONOTFOUND = "InvalidParameter.CdnCertInfoNotFound"
//  INVALIDPARAMETER_CDNHOSTHASSPECIALCONFIG = "InvalidParameter.CdnHostHasSpecialConfig"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  RESOURCEUNAVAILABLE_CDNHOSTNOICP = "ResourceUnavailable.CdnHostNoIcp"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTINICPBLACKLIST = "UnauthorizedOperation.CdnHostInIcpBlacklist"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERINVALIDCREDENTIAL = "UnauthorizedOperation.CdnUserInvalidCredential"
//  UNAUTHORIZEDOPERATION_CDNUSERISISOLATED = "UnauthorizedOperation.CdnUserIsIsolated"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNSUPPORTEDOPERATION_OPNOAUTH = "UnsupportedOperation.OpNoAuth"
func (c *Client) StartCdnDomainWithContext(ctx context.Context, request *StartCdnDomainRequest) (response *StartCdnDomainResponse, err error) {
    if request == nil {
        request = NewStartCdnDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartCdnDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartCdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewStartScdnDomainRequest() (request *StartScdnDomainRequest) {
    request = &StartScdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "StartScdnDomain")
    
    
    return
}

func NewStartScdnDomainResponse() (response *StartScdnDomainResponse) {
    response = &StartScdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartScdnDomain
// StartScdnDomain 用于开启域名的安全防护配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNHOSTHASSPECIALCONFIG = "InvalidParameter.CdnHostHasSpecialConfig"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CDNUSERTOOMANYHOSTS = "ResourceNotFound.CdnUserTooManyHosts"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  RESOURCEUNAVAILABLE_SCDNUSERNOPACKAGE = "ResourceUnavailable.ScdnUserNoPackage"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
func (c *Client) StartScdnDomain(request *StartScdnDomainRequest) (response *StartScdnDomainResponse, err error) {
    return c.StartScdnDomainWithContext(context.Background(), request)
}

// StartScdnDomain
// StartScdnDomain 用于开启域名的安全防护配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNHOSTHASSPECIALCONFIG = "InvalidParameter.CdnHostHasSpecialConfig"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CDNUSERTOOMANYHOSTS = "ResourceNotFound.CdnUserTooManyHosts"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  RESOURCEUNAVAILABLE_SCDNUSERNOPACKAGE = "ResourceUnavailable.ScdnUserNoPackage"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
func (c *Client) StartScdnDomainWithContext(ctx context.Context, request *StartScdnDomainRequest) (response *StartScdnDomainResponse, err error) {
    if request == nil {
        request = NewStartScdnDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartScdnDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartScdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewStopCdnDomainRequest() (request *StopCdnDomainRequest) {
    request = &StopCdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "StopCdnDomain")
    
    
    return
}

func NewStopCdnDomainResponse() (response *StopCdnDomainResponse) {
    response = &StopCdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopCdnDomain
// StopCdnDomain 用于停止域名的加速服务。
//
// 注意：停止加速服务后，访问至加速节点的请求将会直接返回 404。为避免对您的业务造成影响，请在停止加速服务前将解析切走。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNHOSTHASSPECIALCONFIG = "InvalidParameter.CdnHostHasSpecialConfig"
//  INVALIDPARAMETER_CDNKEYRULESEXCLUDECUSTOMREQUIRESFULLLEGO = "InvalidParameter.CdnKeyRulesExcludeCustomRequiresFullLego"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNSUPPORTEDOPERATION_OPNOAUTH = "UnsupportedOperation.OpNoAuth"
func (c *Client) StopCdnDomain(request *StopCdnDomainRequest) (response *StopCdnDomainResponse, err error) {
    return c.StopCdnDomainWithContext(context.Background(), request)
}

// StopCdnDomain
// StopCdnDomain 用于停止域名的加速服务。
//
// 注意：停止加速服务后，访问至加速节点的请求将会直接返回 404。为避免对您的业务造成影响，请在停止加速服务前将解析切走。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNHOSTHASSPECIALCONFIG = "InvalidParameter.CdnHostHasSpecialConfig"
//  INVALIDPARAMETER_CDNKEYRULESEXCLUDECUSTOMREQUIRESFULLLEGO = "InvalidParameter.CdnKeyRulesExcludeCustomRequiresFullLego"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNSUPPORTEDOPERATION_OPNOAUTH = "UnsupportedOperation.OpNoAuth"
func (c *Client) StopCdnDomainWithContext(ctx context.Context, request *StopCdnDomainRequest) (response *StopCdnDomainResponse, err error) {
    if request == nil {
        request = NewStopCdnDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopCdnDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopCdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewStopScdnDomainRequest() (request *StopScdnDomainRequest) {
    request = &StopScdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "StopScdnDomain")
    
    
    return
}

func NewStopScdnDomainResponse() (response *StopScdnDomainResponse) {
    response = &StopScdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopScdnDomain
// StopScdnDomain 用于关闭域名的安全防护配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNHOSTHASSPECIALCONFIG = "InvalidParameter.CdnHostHasSpecialConfig"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CDNUSERTOOMANYHOSTS = "ResourceNotFound.CdnUserTooManyHosts"
//  RESOURCEUNAVAILABLE_SCDNUSERNOPACKAGE = "ResourceUnavailable.ScdnUserNoPackage"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
func (c *Client) StopScdnDomain(request *StopScdnDomainRequest) (response *StopScdnDomainResponse, err error) {
    return c.StopScdnDomainWithContext(context.Background(), request)
}

// StopScdnDomain
// StopScdnDomain 用于关闭域名的安全防护配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNHOSTHASSPECIALCONFIG = "InvalidParameter.CdnHostHasSpecialConfig"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CDNUSERTOOMANYHOSTS = "ResourceNotFound.CdnUserTooManyHosts"
//  RESOURCEUNAVAILABLE_SCDNUSERNOPACKAGE = "ResourceUnavailable.ScdnUserNoPackage"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
func (c *Client) StopScdnDomainWithContext(ctx context.Context, request *StopScdnDomainRequest) (response *StopScdnDomainResponse, err error) {
    if request == nil {
        request = NewStopScdnDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopScdnDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopScdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDomainConfigRequest() (request *UpdateDomainConfigRequest) {
    request = &UpdateDomainConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "UpdateDomainConfig")
    
    
    return
}

func NewUpdateDomainConfigResponse() (response *UpdateDomainConfigResponse) {
    response = &UpdateDomainConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateDomainConfig
// UpdateDomainConfig 用于修改内容分发网络加速域名配置信息
//
// 注意：如果需要更新复杂类型的配置项，必须传递整个对象的所有属性，未传递的属性将使用默认值，建议通过查询接口获取配置属性后，直接修改后传递给本接口。Https配置由于证书的特殊性，更新时不用传递证书和密钥字段。
//
// 云审计相关：接口的入参可能包含密钥等敏感信息，所以此接口的入参不会上报到云审计。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_INVALIDERRORCODE = "InternalError.InvalidErrorCode"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_ACCESSPORTOPENEDHTTPS = "InvalidParameter.AccessPortOpenedHttps"
//  INVALIDPARAMETER_BANDLIMITREQUIREDMAINLAND = "InvalidParameter.BandLimitRequiredMainland"
//  INVALIDPARAMETER_BANDWIDTHALERTCOUNTERMEASURECONFLICTORIGINTYPE = "InvalidParameter.BandwidthAlertCounterMeasureConflictOriginType"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNCERTINFONOTFOUND = "InvalidParameter.CdnCertInfoNotFound"
//  INVALIDPARAMETER_CDNCONFIGINVALIDCACHE = "InvalidParameter.CdnConfigInvalidCache"
//  INVALIDPARAMETER_CDNHOSTINVALIDMIDDLECONFIG = "InvalidParameter.CdnHostInvalidMiddleConfig"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNINVALIDPARAMSERVICETYPE = "InvalidParameter.CdnInvalidParamServiceType"
//  INVALIDPARAMETER_CDNKEYRULESEXCLUDECUSTOMREQUIRESFULLLEGO = "InvalidParameter.CdnKeyRulesExcludeCustomRequiresFullLego"
//  INVALIDPARAMETER_CDNKEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.CdnKeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  INVALIDPARAMETER_PATHREGEXTOOMANYSUBPATTERN = "InvalidParameter.PathRegexTooManySubPattern"
//  INVALIDPARAMETER_REMOTEAUTHINVALIDPLATFORM = "InvalidParameter.RemoteAuthInvalidPlatform"
//  INVALIDPARAMETER_REMOTEAUTHINVALIDPROTOCOL = "InvalidParameter.RemoteAuthInvalidProtocol"
//  LIMITEXCEEDED_CDNCONFIGTOOMANYCACHERULES = "LimitExceeded.CdnConfigTooManyCacheRules"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCEINUSE_CDNHOSTEXISTS = "ResourceInUse.CdnHostExists"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTBELONGSTOOTHERSINMAINLAND = "ResourceUnavailable.CdnHostBelongsToOthersInMainland"
//  RESOURCEUNAVAILABLE_CDNHOSTBELONGSTOOTHERSINOVERSEAS = "ResourceUnavailable.CdnHostBelongsToOthersInOverseas"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  RESOURCEUNAVAILABLE_CDNHOSTISNOTOFFLINE = "ResourceUnavailable.CdnHostIsNotOffline"
//  RESOURCEUNAVAILABLE_CDNHOSTNOICP = "ResourceUnavailable.CdnHostNoIcp"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTINICPBLACKLIST = "UnauthorizedOperation.CdnHostInIcpBlacklist"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNTAGUNAUTHORIZED = "UnauthorizedOperation.CdnTagUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISISOLATED = "UnauthorizedOperation.CdnUserIsIsolated"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) UpdateDomainConfig(request *UpdateDomainConfigRequest) (response *UpdateDomainConfigResponse, err error) {
    return c.UpdateDomainConfigWithContext(context.Background(), request)
}

// UpdateDomainConfig
// UpdateDomainConfig 用于修改内容分发网络加速域名配置信息
//
// 注意：如果需要更新复杂类型的配置项，必须传递整个对象的所有属性，未传递的属性将使用默认值，建议通过查询接口获取配置属性后，直接修改后传递给本接口。Https配置由于证书的特殊性，更新时不用传递证书和密钥字段。
//
// 云审计相关：接口的入参可能包含密钥等敏感信息，所以此接口的入参不会上报到云审计。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_INVALIDERRORCODE = "InternalError.InvalidErrorCode"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_ACCESSPORTOPENEDHTTPS = "InvalidParameter.AccessPortOpenedHttps"
//  INVALIDPARAMETER_BANDLIMITREQUIREDMAINLAND = "InvalidParameter.BandLimitRequiredMainland"
//  INVALIDPARAMETER_BANDWIDTHALERTCOUNTERMEASURECONFLICTORIGINTYPE = "InvalidParameter.BandwidthAlertCounterMeasureConflictOriginType"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNCERTINFONOTFOUND = "InvalidParameter.CdnCertInfoNotFound"
//  INVALIDPARAMETER_CDNCONFIGINVALIDCACHE = "InvalidParameter.CdnConfigInvalidCache"
//  INVALIDPARAMETER_CDNHOSTINVALIDMIDDLECONFIG = "InvalidParameter.CdnHostInvalidMiddleConfig"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNINVALIDPARAMSERVICETYPE = "InvalidParameter.CdnInvalidParamServiceType"
//  INVALIDPARAMETER_CDNKEYRULESEXCLUDECUSTOMREQUIRESFULLLEGO = "InvalidParameter.CdnKeyRulesExcludeCustomRequiresFullLego"
//  INVALIDPARAMETER_CDNKEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.CdnKeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  INVALIDPARAMETER_PATHREGEXTOOMANYSUBPATTERN = "InvalidParameter.PathRegexTooManySubPattern"
//  INVALIDPARAMETER_REMOTEAUTHINVALIDPLATFORM = "InvalidParameter.RemoteAuthInvalidPlatform"
//  INVALIDPARAMETER_REMOTEAUTHINVALIDPROTOCOL = "InvalidParameter.RemoteAuthInvalidProtocol"
//  LIMITEXCEEDED_CDNCONFIGTOOMANYCACHERULES = "LimitExceeded.CdnConfigTooManyCacheRules"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCEINUSE_CDNHOSTEXISTS = "ResourceInUse.CdnHostExists"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTBELONGSTOOTHERSINMAINLAND = "ResourceUnavailable.CdnHostBelongsToOthersInMainland"
//  RESOURCEUNAVAILABLE_CDNHOSTBELONGSTOOTHERSINOVERSEAS = "ResourceUnavailable.CdnHostBelongsToOthersInOverseas"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  RESOURCEUNAVAILABLE_CDNHOSTISNOTOFFLINE = "ResourceUnavailable.CdnHostIsNotOffline"
//  RESOURCEUNAVAILABLE_CDNHOSTNOICP = "ResourceUnavailable.CdnHostNoIcp"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTINICPBLACKLIST = "UnauthorizedOperation.CdnHostInIcpBlacklist"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNTAGUNAUTHORIZED = "UnauthorizedOperation.CdnTagUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISISOLATED = "UnauthorizedOperation.CdnUserIsIsolated"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) UpdateDomainConfigWithContext(ctx context.Context, request *UpdateDomainConfigRequest) (response *UpdateDomainConfigResponse, err error) {
    if request == nil {
        request = NewUpdateDomainConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateDomainConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateDomainConfigResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateImageConfigRequest() (request *UpdateImageConfigRequest) {
    request = &UpdateImageConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "UpdateImageConfig")
    
    
    return
}

func NewUpdateImageConfigResponse() (response *UpdateImageConfigResponse) {
    response = &UpdateImageConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateImageConfig
// UpdateImageConfig 用于更新控制台图片优化的相关配置，支持Webp、TPG、 Guetzli 和 Avif。 
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) UpdateImageConfig(request *UpdateImageConfigRequest) (response *UpdateImageConfigResponse, err error) {
    return c.UpdateImageConfigWithContext(context.Background(), request)
}

// UpdateImageConfig
// UpdateImageConfig 用于更新控制台图片优化的相关配置，支持Webp、TPG、 Guetzli 和 Avif。 
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) UpdateImageConfigWithContext(ctx context.Context, request *UpdateImageConfigRequest) (response *UpdateImageConfigResponse, err error) {
    if request == nil {
        request = NewUpdateImageConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateImageConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateImageConfigResponse()
    err = c.Send(request, response)
    return
}

func NewUpdatePayTypeRequest() (request *UpdatePayTypeRequest) {
    request = &UpdatePayTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "UpdatePayType")
    
    
    return
}

func NewUpdatePayTypeResponse() (response *UpdatePayTypeResponse) {
    response = &UpdatePayTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdatePayType
// 本接口(UpdatePayType)用于修改账号计费类型，暂不支持月结用户或子账号修改。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISISOLATED = "UnauthorizedOperation.CdnUserIsIsolated"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) UpdatePayType(request *UpdatePayTypeRequest) (response *UpdatePayTypeResponse, err error) {
    return c.UpdatePayTypeWithContext(context.Background(), request)
}

// UpdatePayType
// 本接口(UpdatePayType)用于修改账号计费类型，暂不支持月结用户或子账号修改。
//
// 可能返回的错误码:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISISOLATED = "UnauthorizedOperation.CdnUserIsIsolated"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) UpdatePayTypeWithContext(ctx context.Context, request *UpdatePayTypeRequest) (response *UpdatePayTypeResponse, err error) {
    if request == nil {
        request = NewUpdatePayTypeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdatePayType require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdatePayTypeResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateScdnDomainRequest() (request *UpdateScdnDomainRequest) {
    request = &UpdateScdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "UpdateScdnDomain")
    
    
    return
}

func NewUpdateScdnDomainResponse() (response *UpdateScdnDomainResponse) {
    response = &UpdateScdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateScdnDomain
// UpdateScdnDomain 用于修改 SCDN 加速域名安全相关配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_SCDNUSERNOPACKAGE = "InternalError.ScdnUserNoPackage"
//  INTERNALERROR_SCDNUSERSUSPEND = "InternalError.ScdnUserSuspend"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  RESOURCEUNAVAILABLE_SCDNUSERNOPACKAGE = "ResourceUnavailable.ScdnUserNoPackage"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
func (c *Client) UpdateScdnDomain(request *UpdateScdnDomainRequest) (response *UpdateScdnDomainResponse, err error) {
    return c.UpdateScdnDomainWithContext(context.Background(), request)
}

// UpdateScdnDomain
// UpdateScdnDomain 用于修改 SCDN 加速域名安全相关配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_SCDNUSERNOPACKAGE = "InternalError.ScdnUserNoPackage"
//  INTERNALERROR_SCDNUSERSUSPEND = "InternalError.ScdnUserSuspend"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  RESOURCEUNAVAILABLE_SCDNUSERNOPACKAGE = "ResourceUnavailable.ScdnUserNoPackage"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
func (c *Client) UpdateScdnDomainWithContext(ctx context.Context, request *UpdateScdnDomainRequest) (response *UpdateScdnDomainResponse, err error) {
    if request == nil {
        request = NewUpdateScdnDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateScdnDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateScdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyDomainRecordRequest() (request *VerifyDomainRecordRequest) {
    request = &VerifyDomainRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "VerifyDomainRecord")
    
    
    return
}

func NewVerifyDomainRecordResponse() (response *VerifyDomainRecordResponse) {
    response = &VerifyDomainRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// VerifyDomainRecord
// 验证域名解析值
//
// 可能返回的错误码:
//  INVALIDPARAMETER_CDNCONFIGINVALIDHOST = "InvalidParameter.CdnConfigInvalidHost"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNTXTRECORDVALUENOTMATCH = "UnauthorizedOperation.CdnTxtRecordValueNotMatch"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) VerifyDomainRecord(request *VerifyDomainRecordRequest) (response *VerifyDomainRecordResponse, err error) {
    return c.VerifyDomainRecordWithContext(context.Background(), request)
}

// VerifyDomainRecord
// 验证域名解析值
//
// 可能返回的错误码:
//  INVALIDPARAMETER_CDNCONFIGINVALIDHOST = "InvalidParameter.CdnConfigInvalidHost"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNTXTRECORDVALUENOTMATCH = "UnauthorizedOperation.CdnTxtRecordValueNotMatch"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) VerifyDomainRecordWithContext(ctx context.Context, request *VerifyDomainRecordRequest) (response *VerifyDomainRecordResponse, err error) {
    if request == nil {
        request = NewVerifyDomainRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyDomainRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyDomainRecordResponse()
    err = c.Send(request, response)
    return
}
