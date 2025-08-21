# Changelog

### [v1.1.0] - 2025-07-29
- fix: 自动补全 Request 对象的初始化

### [v1.0.1116] - 2025-03-07
- fix: 解决了 HTTP 请求体因被消耗而导致重试失败的问题
- fix: SSE 消息解析器 (如果值以 U+0020 空格字符开头，则将其从值中移除)

### [v1.0.887] - 2024-03-27
- feat: 使用 OmitBehaviour 来控制 json 序列化的 omitnil 行为

### [v1.0.877] - 2024-03-14
- fix: omitnil 和 omitempty 的向后兼容性问题

### [v1.0.864] - 2024-02-28
- feat: 支持 OIDCRoleArnProvider 的 beforeRefresh

### [v1.0.863] - 2024-02-28
- fix: 移除 common/json 中的 omitempty (以便在 models.go 中添加 omitempty 标签同时保持代码兼容性)

### [v1.0.829] - 2023-12-29
- fix: 使 STS 凭证更加一致

### [v1.0.819] - 2023-12-14
- feat: OIDCRoleArnProvider 支持 endpoint

### [v1.0.812] - 2023-12-06
- feat: UnsafeRetryOnConnectionFailure

### [v1.0.877] - 2024-03-14
- feat: 支持 apigw endpoint

### [v1.0.768] - 2023-10-18
- feat: 支持 SSE

### [v1.0.738] - 2023-09-01
- feat: json 包支持 omitnil 标签

### [v1.0.636] - 2023-04-12
- feat: 支持自定义日志记录器

### [v1.0.606] - 2023-03-01
- fix: 修复了启用 debug 标志时 HTTP 请求失败导致程序崩溃的问题

### [v1.0.595] - 2023-03-14
- feat: 启用 debug 时记录响应

### [v1.0.507] - 2022-09-28
- fix: 支持 httpmock

### [v1.0.499] - 2022-09-19
- feat: 支持自定义 http client

### [v1.0.494] - 2022-09-12
- feat: 支持 Content-Encoding gzip

### [v1.0.486] - 2022-08-31
- feat: OIDC role provider

### [v1.0.479] - 2022-08-18
- feat: 支持跳过签名

### [v1.0.469] - 2022-08-10
- feat: http proxy

### [v1.0.379] - 2022-04-06
- 支持自定义 header

### [v1.0.324] - 2022-01-05
- feat: common/client 添加 context 支持

### [v1.0.315] - 2022-12-17
- deprecated: ClientProfile.BackupEndPoint 
- fix: BaseRequest.SetPath

### [v1.0.309] - 2021-12-09
- feat: 支持 octet-stream
- fix: use crypto/rand for clienttoken

### [v1.0.300] - 2021-11-26
- feat: 支持lz4压缩上传日志

### [v1.0.238] - 2021-08-27
- fix : 将 common request 更改为 common client

### [v1.0.227] - 2021-08-12
- feat: 支持地域熔断

### [v1.0.211] - 2021-07-19
- SDK增强特性，客户端重试

### [v1.0.210] - 2021-07-16
- feat: 通过 common request 发送 octet stream

### [v1.0.189] - 2021-06-18
- feat: 通过 common client 发送请求

### [v1.0.180] - 2020-05-27
- 添加 SetScheme 和 SetRootDomain 支持
