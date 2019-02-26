package main

import (
	"fmt"
	"log"
	"os"

	"encoding/base64"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	scf "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/scf/v20180416"
)

func main() {
	// 从环境变量获取SID和SKEY实例化认证对象
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)
	// 实例化客户端配置对象
	cpf := profile.NewClientProfile()
	// 实例化SCF的client对象
	client, _ := scf.NewClient(credential, regions.Guangzhou, cpf)

	// API -- CreateFunction:创建函数
	createFunctionRequest := scf.NewCreateFunctionRequest()
	// 创建的函数名称，函数名称支持26个英文字母大小写、数字、连接符和下划线，第一个字符只能以字母开头，最后一个字符不能为连接符或者下划线，名称长度2-60
	createFunctionRequest.FunctionName = common.StringPtr("sum")
	// 包含函数代码文件及其依赖项的 zip 格式文件，使用该接口时要求将 zip 文件的内容转成 base64 编码，最大支持20M
	// zip包在打包时需要加上-j参数去掉文件夹名字： zip -rj sum.zip sum
	zipFile, err := getZipFileBase64Content("./sum.zip")
	if err != nil {
		log.Fatal(err)
	}
	createFunctionRequest.Code = &scf.Code{
		ZipFile: zipFile,
	}
	// 非必填。函数处理方法名称，名称格式支持 "文件名称.方法名称" 形式，文件名称和函数名称之间以"."隔开，文件名称和函数名称要求以字母开始和结尾，中间允许插入字母、数字、下划线和连接符，文件名称和函数名字的长度要求是 2-60 个字符
	createFunctionRequest.Handler = common.StringPtr("index.main_handler")
	// 非必填。函数描述,最大支持 1000 个英文字母、数字、空格、逗号、换行符和英文句号，支持中文
	createFunctionRequest.Description = common.StringPtr("执行key1和key2参数的sum操作")
	// 非必填。函数运行时内存大小，默认为 128M，可选范围 128MB-1536MB，并且以 128MB 为阶梯
	createFunctionRequest.MemorySize = common.Int64Ptr(128)
	// 非必填。函数最长执行时间，单位为秒，可选值范围 1-300 秒，默认为 3 秒
	createFunctionRequest.Timeout = common.Int64Ptr(3)
	// 非必填。函数的环境变量
	createFunctionRequest.Environment = &scf.Environment{
		Variables: []*scf.Variable{
			&scf.Variable{
				common.StringPtr("ENV_VAR_1"),
				common.StringPtr("ENV_VAL_1"),
			},
			&scf.Variable{
				common.StringPtr("ENV_VAR_2"),
				common.StringPtr("ENV_VAL_2"),
			},
		},
	}
	// 非必填。函数运行环境，目前仅支持 Python2.7，Python3.6，Nodejs6.10， PHP5， PHP7，Golang1 和 Java8，默认Python2.7
	createFunctionRequest.Runtime = common.StringPtr("Python2.7")
	// 非必填。函数的私有网络配置
	createFunctionRequest.VpcConfig = &scf.VpcConfig{
		VpcId:    common.StringPtr(os.Getenv("TENCENTCLOUD_VPCID")),
		SubnetId: common.StringPtr(os.Getenv("TENCENTCLOUD_VPC_SUBNETID")),
	}
	createFunctionResponse, err := client.CreateFunction(createFunctionRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(createFunctionResponse.ToJsonString())
}

func getZipFileBase64Content(zipFilePath string) (b64Content *string, err error) {
	// 读取给定路径下打包的zip代码文件内容并转为base64编码
	zipFile, err := os.Open(zipFilePath)
	if err != nil {
		return nil, err
	}
	defer zipFile.Close()
	fileInfo, err := zipFile.Stat()
	if err != nil {
		return nil, err
	}
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	_, err = zipFile.Read(buffer)
	if err != nil {
		return nil, err
	}
	b64Content = common.StringPtr(base64.StdEncoding.EncodeToString(buffer))
	return b64Content, err
}
