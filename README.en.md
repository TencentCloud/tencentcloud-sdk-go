# Introduction

Welcome to Tencent Cloud Developer Tools Suite (SDK) 3.0, which is a companion tool to the Cloud API 3.0 platform. All subsequent cloud service products will be accessed. The new version of the SDK has been unified, with the same usage of the SDK in each language, the same way of calling the interface, the unified error code and the return packet format.

To make it easier for GO developers to debug and access the Tencent Cloud Product API, here is the Tencent Cloud Development Kit for GO and a simple example of using the development kit for the first time. Let you quickly get the Tencent Cloud GO SDK and start calling.

# Dependencies

1. Go 1.9 version and above, and set the necessary environment variables such as GOPATH.
2. Before using the related products, you need to open the corresponding products in Tencent Cloud Console.
3. Get SecretID and SecretKey on the Tencent Cloud Console [Access Management](https://console.cloud.tencent.com/cam/capi) page.

# Installation

Obtain security credentials before installing the Go SDK. Before using the Cloud API for the first time, users first need to apply for security credentials on the Tencent Cloud Console. The security credentials include SecretID and SecretKey. SecretID is used to identify the identity of the API caller. SecretKey is used to encrypt the signature string and server. The key that verifies the signature string. SecretKey must be kept strictly to avoid disclosure.

## Install by go get (recommended)

It is recommended to install the SDK using the tools that come with the language:

    go get -u github.com/tencentcloud/tencentcloud-sdk-go


## Install from source

Go to [Github Code Address](https://github.com/tencentcloud/tencentcloud-sdk-go) to download the latest code, unzip it and install it in the $GOPATH/src/github.com/tencentcloud directory.

# Example

Each interface has a corresponding Request structure and a Response structure. For example, the cloud server's query instance list interface DescribeInstances has a corresponding request structure DescribeInstancesRequest and a return structure DescribeInstancesResponse .

The following uses the cloud server query instance list interface as an example to introduce the basic usage of the SDK. For the purpose of the demonstration, some non-essential content has been added to try to show the functions commonly used by the SDK, but it is also bloated. When you actually write code using the SDK, you should try to be as simple as possible.

```
Package main

import (
        "fmt"
        "os"

        "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
        "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
        "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
        "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
        Cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func main() {
        // necessary steps:
        // Instantiate an authentication object. The incoming key needs to be sent to the Tencent cloud account key pair secretId, secretKey.
        // This is the way to read from environment variables, you need to set these two values ​​in the environment variable first.
        // You can also write a dead key pair directly in the code, but be careful not to copy, upload or share the code to others.
        // Avoid leaking key pairs that endanger your account.
        Credential := common.NewCredential(
                os.Getenv("TENCENTCLOUD_SECRET_ID"),
                os.Getenv("TENCENTCLOUD_SECRET_KEY"),
        )

        // non-essential steps
        / / Instantiate a client configuration object, you can specify the timeout and other configuration
        cpf := profile.NewClientProfile()
        // The SDK uses the POST method by default.
        // If you must use the GET method, you can set it here. The GET method cannot handle some large requests.
        cpf.HttpProfile.ReqMethod = "GET"
        // The SDK has a default timeout. Please do not adjust if it is not necessary.
        // Check if you need it in the code to get the latest defaults.
        cpf.HttpProfile.ReqTimeout = 10
        // The SDK will automatically specify the domain name. Usually you don't need to specify a domain name specifically, but if you are visiting a financial district service,
        // You must manually specify the domain name, such as the Shanghai financial region name of the cloud server: cvm.ap-shanghai-fsi.tencentcloudapi.com
        cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
        // The SDK is signed by HmacSHA256 by default, which is safer but slightly degrades performance.
        // Do not modify this field if it is not necessary.
        cpf.SignMethod = "HmacSHA1"

        / / Instantiate the client object to request the product (using cvm as an example)
        // The second parameter is the geographic information, you can directly fill in the string ap-guangzhou, or reference the preset constant
        client, _ := cvm.NewClient(credential, regions.Guangzhou, cpf)
        / / Instantiate a request object, according to the called interface and the actual situation, you can further set the request parameters
        // You can directly query the SDK source to determine which properties of DescribeInstancesRequest can be set.
        // The attribute may be a primitive type or it may reference another data structure.
        // It is recommended to use the IDE for development, which can be easily accessed to view the documentation of each interface and data structure.
        request := cvm.NewDescribeInstancesRequest()

        // Basic type settings.
        // This interface allows you to set the number of instances returned. Specify here to return only one.
        // The SDK uses pointer style specified parameters, even for basic types you need to use pointers to assign values ​​to the parameters.
        // SDK provides pointer reference wrapper function for basic types
        request.Limit = common.Int64Ptr(1)

        // The setting of the array type.
        // This interface allows the specified instance ID to be filtered, but is commented out first because it conflicts with the Filter parameter to be demonstrated next.
        // request.InstanceIds = common.StringPtrs([]string{"ins-r8hr2upy"})

        // Settings for complex objects.
        // In this interface, Filters is an array, the elements of the array are complex objects Filter, and the members of Filter are Value arrays.
        request.Filters = []*cvm.Filter{
            &cvm.Filter{
                Name: common.StringPtr("zone"),
                Values: common.StringPtrs([]string{"ap-guangzhou-1"}),
            },
        }

        // Set a request using the json string, note that the request is actually updated, ie Limit=1 will be retained.
        // The zone for the filter condition will change to ap-guangzhou-2.
        // If you need a new request, you need to create it with cvm.NewDescribeInstancesRequest().
        err := request.FromJsonString(`{"Filters":[{"Name":"zone","Values":["ap-guangzhou-2"]}]}`)
        if err != nil {
                Panic(err)
        }
        / / Through the client object call the interface you want to access, you need to pass in the request object
        response, err := client.DescribeInstances(request)
        / / Handling exceptions
        if _, ok := err.(*errors.TencentCloudSDKError); ok {
                fmt.Printf("An API error has returned: %s", err)
                Return
        }
        // Non-SDK exception, direct failure. Other processing can be added to the actual code.
        if err != nil {
                Panic(err)
        }
        // print the returned json string
        fmt.Printf("%s", response.ToJsonString())
}
```

See the [examples](https://github.com/TencentCloud/tencentcloud-sdk-go/tree/master/examples) directory for more examples. For examples of Request initialization for complex interfaces, see examples/cvm/v20170312/run_instances.go . For examples of initializing Requests using json strings, see examples/cvm/v20170312/describe_instances.go .

# Related configuration

## Proxy

If it is a proxy environment, you need to set the system environment variable `https_proxy`. Otherwise, it may not be called normally, and the connection timeout exception will be thrown.

# Supported product list

| Package Name | Product English Name |
|------------|------|
| aai | Intelligent Voice Service |
| as | flexible stretching |
| batch | Batch Calculation |
| billing | Billing |
| bm | Blackstone Physics Server |
| bmeip | Blackrock Elastic IP |
| cbs | Cloud Hard Drive |
| cdb | TencentDB MySQL |
| cdn | Content Distribution Network |
| cis | Container Instance Service |
| clb | Cloud Load Balancing |
| cr | Collection Robot |
| cvm | cloud server |
| cws | Web Vulnerability Scan |
| dc | Private line access |
| dcdb | Distributed Database DCDB |
| drm | Digital Rights Management |
| ds | Electronic Contract Services |
| dts | Data Transfer Service DTS |
| es | Elasticsearch Service |
| facefusion | Face Fusion |
| faceid | Face ID |
| habo | Sample Intelligence Analysis Platform |
| hcm | Math Job Correction |
| iai | Facial Recognition |
| iot | Accelerate IoT Suite |
| iotcloud | Internet of Things Communications |
| live | Cloud Live |
| mariadb | TencentDB MariaDB(TDSQL) |
| mongodb | TencentDB MongoDB |
| monitor | Cloud Monitoring |
| ms | Application Security |
| msp | Migration Service Platform |
| partners | Channel Partners |
| postgres | TencentDB PostgreSQL |
| redis | TencentDB Redis |
| scf | Serverless Cloud Function |
| soe | Sense of listening |
| sqlserver | TencentDB SQL Server |
| sts | Access Management |
| tbaas | TBaaS |
| tbm | Tencent Excellent Review |
| tmt | Tencent Machine Translation |
| vod | Cloud on Demand |
| vpc | Private Network |
| youmall | Tencent Excellent Mall |
| yunjing | Host Security |
