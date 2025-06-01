# Release v1.0.1178

## 消息队列 MQTT 版(mqtt) 版本：2024-05-16

### 第 15 次发布

发布时间：2025-05-30 14:32:54

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeMessageByTopic](https://cloud.tencent.com/document/api/1778/118936)

修改接口：

* [ModifyInstance](https://cloud.tencent.com/document/api/1778/116039)

* [ModifyInstanceCertBinding](https://cloud.tencent.com/document/api/1778/116203)

	* <font color="#dd0000">**修改入参**：</font>SSLServerCertId, SSLCaCertId

* [ModifyJWTAuthenticator](https://cloud.tencent.com/document/api/1778/111086)

	* 新增入参：Status


新增数据结构：

* [MQTTMessage](https://cloud.tencent.com/document/api/1778/111031#MQTTMessage)

修改数据结构：

* [MQTTMessageItem](https://cloud.tencent.com/document/api/1778/111031#MQTTMessageItem)

* [PublicAccessRule](https://cloud.tencent.com/document/api/1778/111031#PublicAccessRule)

	* <font color="#dd0000">**修改成员**：</font>IpRule, Allow




