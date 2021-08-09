package main

import (
	"fmt"
	"os"
	"time"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	pb "github.com/golang/protobuf/proto"
	"github.com/pierrec/lz4"
	"github.com/tencentcloud/tencentcloud-sdk-go/examples/cls/proto"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
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
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"))
	cpf := profile.NewClientProfile()

	cpf.HttpProfile.Endpoint = "cls.tencentcloudapi.com"
	cpf.HttpProfile.ReqMethod = "POST"
	//创建common client
	client := common.NewCommonClient(credential, regions.GuangzhouOpen, cpf)
	// 创建common request
	headers := map[string]string{
		"X-CLS-TopicId":      "e621fdb8-16f4-41cf-bc73-5aead0b75a03",
		"X-CLS-HashKey":      "0fffffffffffffffffffffffffffffff",
	}
	commpresstype := ""//压缩算法类型， 目前只支持lz4
	body := GenPbBody()
	length := lz4.CompressBlockBound(len(body)) + 1
	compressbody := make([]byte, length)
	n, err := lz4.CompressBlock(body, compressbody, nil)
	if err == nil && n > 0 {
		commpresstype = "lz4"
		body = compressbody[0:n]
	}
	headers["X-CLS-CompressType"] = commpresstype

	request := tchttp.NewCommonRequest("cls", "2020-10-16", "UploadLog")
	request.SetOctetStreamParameters(headers, body)
	//创建common response
	response := tchttp.NewCommonResponse()

	//发送请求
	err = client.SendOctetStream(request, response)
	if err != nil {
		fmt.Println(fmt.Sprintf("fail to invoke api: %v", err))
		return
	}
	fmt.Println(string(response.GetBody()))
}
