package main

import (
	"fmt"

	ame "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ame/v20190916"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

func main() {
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

	// 云音乐相关接口调用示例
	CloudMusicDemo(client)

	// 曲库包相关接口调用示例
	PackageMusicDemo(client)
}

// 云音乐相关接口调用示例
func CloudMusicDemo(client *ame.Client) {
	// 1. 获取授权项目列表
	request := ame.NewDescribeAuthInfoRequest()
	request.Limit = common.Int64Ptr(10)
	request.Offset = common.Int64Ptr(0)

	response, err := client.DescribeAuthInfo(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	// unexpected errors
	if err != nil {
		panic(err)
	}

	println(response.ToJsonString())
	if len(response.Response.AuthInfo) <= 0 {
		println("暂未创建授权项目")
		return
	}

	// 2. 获取授权项目已购云音乐列表
	authInfoId := response.Response.AuthInfo[0].Id
	request1 := ame.NewDescribeCloudMusicPurchasedRequest()
	request1.AuthInfoId = authInfoId
	response1, err := client.DescribeCloudMusicPurchased(request1)
	if err != nil {
		panic(err)
	}

	println(response1.ToJsonString())
	if len(response1.Response.MusicOpenDetail) <= 0 {
		println("授权项目下暂无购买音乐")
		return
	}

	// 3. 获取云音乐播放信息
	musicId := response1.Response.MusicOpenDetail[0].MusicId
	request2 := ame.NewDescribeCloudMusicRequest()
	request2.MusicId = musicId
	// 该值为固定值，必填，不填返回带水印的试听音乐播放链接
	request2.MusicType = common.StringPtr("MP3-320K-FTD")
	response2, err := client.DescribeCloudMusic(request2)
	if err != nil {
		panic(err)
	}
	// 打印音乐播放链接
	println(*response2.Response.MusicUrl)
}

// 曲库包相关接口调用示例
func PackageMusicDemo(client *ame.Client) {
	// 1. 获取已购曲库包列表
	request := ame.NewDescribePackagesRequest()
	request.Length = common.Uint64Ptr(10)
	request.Offset = common.Uint64Ptr(0)
	response, err := client.DescribePackages(request)
	if err != nil {
		panic(err)
	}
	println(response.ToJsonString())
	if len(response.Response.Packages) <= 0 {
		println("暂未购买任何曲库包")
		return
	}

	// 2. 获取曲库包已核销歌曲列表
	request1 := ame.NewDescribePackageItemsRequest()
	request1.Length = common.Uint64Ptr(10)
	request1.Offset = common.Uint64Ptr(0)
	request1.OrderId = response.Response.Packages[0].OrderId
	response1, err := client.DescribePackageItems(request1)
	if err != nil {
		panic(err)
	}
	println(response1.ToJsonString())
	if len(response1.Response.PackageItems) <= 0 {
		println("曲库包下未核销任何歌曲")
		return
	}

	// 3. 获取曲库包歌曲播放信息
	itemId := response1.Response.PackageItems[0].ItemID
	request2 := ame.NewDescribeMusicRequest()
	request2.ItemId = itemId
	// 该值为固定值
	request2.SubItemType = common.StringPtr("MP3-320K-FTD")
	// 在应用前端播放音乐C端用户的唯一标识。无需是账户信息，用户唯一标识即可。
	request2.IdentityId = common.StringPtr("userid")
	request2.Ssl = common.StringPtr("Y")
	response2, err := client.DescribeMusic(request2)
	if err != nil {
		panic(err)
	}
	println(*response2.Response.Music.FullUrl)

	// 4. 获取歌词信息
	request3 := ame.NewDescribeLyricRequest()
	request3.ItemId = itemId
	request3.SubItemType = common.StringPtr("LRC-LRC")
	response3, err := client.DescribeLyric(request3)
	if err != nil {
		panic(err)
	}
	println(*response3.Response.Lyric.Url)
}
