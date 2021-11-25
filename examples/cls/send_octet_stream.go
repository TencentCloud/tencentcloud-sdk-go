package main

import (
	"fmt"
	"github.com/pierrec/lz4"
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
	compressType := "xxx" //压缩方式, 目前只支持lz4, 客户根据需要填写(空字符串意味不压缩)
	region := "xxx"//需要根据客户的实际地域自行填写
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
		"X-CLS-CompressType": compressType,
	}

	body := GenPbBody()
	if compressType == "lz4" {
		out := make([]byte, lz4.CompressBlockBound(len(body)) + 1)
		var hashTable [1 << 16]int
		n, err := lz4.CompressBlock(body, out, hashTable[:])
		if err != nil {
			fmt.Printf("compress failed. err:%v\n", err.Error())
			return
		}
		// copy incompressible data as lz4 format
		if n == 0 {
			n, _ = copyIncompressible(body, out)
		}
		body = out[:n]
	}

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

func copyIncompressible(src, dst []byte) (int, error) {
	lLen, dn := len(src), len(dst)
	di := 0
	if lLen < 0xF {
		dst[di] = byte(lLen << 4)
	} else {
		dst[di] = 0xF0
		if di++; di == dn {
			return di, nil
		}
		lLen -= 0xF
		for ; lLen >= 0xFF; lLen -= 0xFF {
			dst[di] = 0xFF
			if di++; di == dn {
				return di, nil
			}
		}
		dst[di] = byte(lLen)
	}
	if di++; di+len(src) > dn {
		return di, nil
	}
	di += copy(dst[di:], src)
	return di, nil
}