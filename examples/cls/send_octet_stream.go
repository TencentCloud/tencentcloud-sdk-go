package main

import (
	"fmt"
	"os"
	"time"

	pb "github.com/golang/protobuf/proto"
	"github.com/tencentcloud/tencentcloud-sdk-go/examples/cls/proto"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)
func GenPbBody() []byte {
	nowtimestamp := int64(time.Now().Unix())
	key := "system1,log"
	value := "start1,report"
	log := proto.Log{
		Time: &nowtimestamp,
		Contents: []*proto.Log_Content{
			&proto.Log_Content{
				Key: &key,
				Value: &value,
			},
		},
	}

	logGroup := proto.LogGroup{
		Logs:  []*proto.Log{&log},
	}
	logGroupList := proto.LogGroupList{
		LogGroupList: []*proto.LogGroup{
			&logGroup,
		},
	}
	data, _ := pb.Marshal(&logGroupList)
	return data
}

func main() {
	region := "xxxx"//需要根据客户的实际地域自行填写
	topicId := "xxxxxx-xxxxxx-xxxxxx-xxxxxx"//这里需要使用客户实际的topicId，不能输入topicname
	hashKey := ""//可选参数，具体参考官方文档：https://cloud.tencent.com/document/product/614/59470
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"))
	cpf := profile.NewClientProfile()

	cpf.HttpProfile.Endpoint = "cls.tencentcloudapi.com"
	cpf.HttpProfile.ReqMethod = "POST"

	//创建common client
	client := common.NewCommonClient(credential, region, cpf)
	// 创建common request
	headers := map[string]string{
		"X-CLS-TopicId": topicId,
		"X-CLS-HashKey": hashKey,
	}

	body := GenPbBody()
	request := tchttp.NewCommonRequest("cls", "2020-10-16", "UploadLog")
	request.SetOctetStreamParameters(headers, body)
	//创建common response
	response := tchttp.NewCommonResponse()

	//发送请求
	err := client.SendOctetStream(request, response)
	if err != nil {
		fmt.Println(fmt.Sprintf("fail to invoke api: %v", err))
		return
	}
	fmt.Println(string(response.GetBody()))
}
