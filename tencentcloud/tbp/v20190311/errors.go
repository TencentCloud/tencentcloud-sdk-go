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

package v20190311

const (
	// 此产品的特有错误码

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 调用元数据服务失败。
	INTERNALERROR_CALLMMSFAILED = "InternalError.CallMMSFailed"

	// 请求ASR服务失败。
	INTERNALERROR_ERRORASR = "InternalError.ErrorAsr"

	// 查询数据库失败，没有对应数据。
	INTERNALERROR_ERRORMMS = "InternalError.ErrorMms"

	// nlu处理失败。
	INTERNALERROR_ERRORNLU = "InternalError.ErrorNlu"

	// rpc调用失败。
	INTERNALERROR_ERRORRPC = "InternalError.ErrorRpc"

	// 请求TTS服务失败。
	INTERNALERROR_ERRORTTS = "InternalError.ErrorTts"

	// webHook处理失败。
	INTERNALERROR_ERRORWEBHOOK = "InternalError.ErrorWebHook"

	// 查询元数据失败，没有对应数据。
	INTERNALERROR_MMSINTERNALERROR = "InternalError.MMSInternalError"

	// 未开通相关应用访问权限。
	INTERNALERROR_NOAPPPRIVILEGE = "InternalError.NoAppPrivilege"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"
)
