package main

import (
	"encoding/json"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ame "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ame/v20190916"
)

func main()  {
	// 必要步骤：
	// 实例化一个认证对象，入参需要传入腾讯云账户密钥对secretId，secretKey。
	// 这里采用的是从环境变量读取的方式，需要在环境变量中先设置这两个值。
	// 你也可以直接在代码中写死密钥对，但是小心不要将代码复制、上传或者分享给他人，
	// 以免泄露密钥对危及你的财产安全。
	credential := common.NewCredential(
		"",
		"",
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	cpf.HttpProfile.ReqTimeout = 5
	cpf.SignMethod = "HmacSHA1"

	client, _ := ame.NewClient(credential, "ap-guangzhou", cpf)
	// 获取分类
	request := ame.NewDescribeStationsRequest()
	request.Limit = common.Uint64Ptr(30)
	request.Offset = common.Uint64Ptr(0)

	// get response structure
	response, err := client.DescribeStations(request)
	// API errors
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	// unexpected errors
	if err != nil {
		panic(err)
	}

	b, _ := json.Marshal(response.Response)
	fmt.Printf("%s\n", b)

	// 获取分类音乐列表
	itemReq := ame.NewDescribeItemsRequest()
	itemReq.Limit = common.Uint64Ptr(10)
	// Offset计算方法：Offset = Offset + Limit
	itemReq.Offset = common.Uint64Ptr(0)
	// 这里只是例子，取其中一个CategoryID
	itemReq.CategoryId = response.Response.Stations[0].CategoryID
	itemRsp, err := client.DescribeItems(itemReq)
	if err != nil {
		panic(err)
	}
	b, _ = json.Marshal(itemRsp.Response)
	fmt.Printf("%s\n", b)
	
	// 歌曲ID，这里取一个做示例
	itemID := itemRsp.Response.Items[0].ItemID

	// 获取歌曲信息
	musicReq := ame.NewDescribeMusicRequest()
	// 这里只是例子，取其中一个ItemID
	musicReq.ItemId = itemID
	// 在应用前端播放音乐C端用户的唯一标识。无需是账户信息，用户唯一标识即可。
	userId := "xcd323dasd1"
	musicReq.IdentityId = common.StringPtr(userId)
	musicRsp, err := client.DescribeMusic(musicReq)
	if err != nil {
		panic(err)
	}
	// 获取音乐播放路径，前提是您已经在腾讯云AME控制台添加过域名
	// 添加域名：https://console.cloud.tencent.com/ame/acc
	println(*musicRsp.Response.Music.FullUrl)

	// 获取歌词信息
	lyricReq := ame.NewDescribeLyricRequest()
	lyricReq.ItemId = itemID
	lyricRsp, err := client.DescribeLyric(lyricReq)
	if err != nil {
		panic(err)
	}
	b, _ = json.Marshal(lyricRsp.Response)
	fmt.Printf("%s\n", b)
}
