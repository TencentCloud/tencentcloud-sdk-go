package main

/**
 * 运行步骤：
 * 1 下载测试音频数据到本地
 *   bike.mp3:                                     https://soe-common-1255701415.cos.ap-guangzhou.myqcloud.com/bike.mp3
 *   cn.mp3:                                       https://soe-common-1255701415.cos.ap-guangzhou.myqcloud.com/cn.mp3
 *   Rex was a big dog.wav :                       https://soe-common-1255701415.cos.ap-guangzhou.myqcloud.com/Rex%20was%20a%20big%20dog.wav
 *   I see sixteen eggs and eighteen apples.wav :  https://soe-common-1255701415.cos.ap-guangzhou.myqcloud.com/I%20see%20sixteen%20eggs%20and%20eighteen%20apples.wav
 * 2 配置测试音频数据目录常量 AUDIO_FOR_ONCE、AUDIO_FOR_CN、AUDIO_FOR_SEQ、AUDIO_FOR_ASYNC 为音频下载目录
 * 3 修改密钥常量 SECRET_ID、SECRET_KEY 为腾讯云帐号密钥
 * 4 运行示例程序
 */

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	soe "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/soe/v20180724"
	"io/ioutil"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	PKG_SIZE = 16 * 1024
)

const (
	SECRET_ID  = ""
	SECRET_KEY = ""
)

const (
	AUDIO_FOR_ONCE  = "./bike.mp3"
	AUDIO_FOR_CN    = "./cn.mp3"
	AUDIO_FOR_SEQ   = "./Rex was a big dog.wav"
	AUDIO_FOR_ASYNC = "./I see sixteen eggs and eighteen apples.wav"
)

const (
	FINISHED   = "Finished"
	FAILED     = "Failed"
	EVALUATING = "Evaluating"
)

const (
	EVAL_MODE_WORD     = 0
	EVAL_MODE_SENTENCE = 1
	EVAL_MODE_PARA     = 2
	EVAL_MODE_FREETALK = 3
)

const (
	WORK_MODE_STREAM = 0
	WORK_MODE_ONCE   = 1
)

const (
	MP3 = 3
	WAV = 2
	RAW = 1
)

const (
	CN = 1
	EN = 0
)

func getSessionId() string {
	return fmt.Sprintf("%d%d", rand.Intn(1000), time.Now().Nanosecond())
}

/**
 * 一次性评估示例,用于较短时间音频的评估，例如单词或者短句，评估一次性发布所有音频数据
 *
 * @param client
 */
func evaluationOnce(client *soe.Client) {
	fmt.Println("--------一次性评测-----")
	sessionId := getSessionId()
	// 实例化一个请求对象，根据调用的接口和实际情况，可以进一步设置请求参数
	// 你可以直接查询SDK源码确定InitOralProcessRequest有哪些属性可以设置，
	// 属性可能是基本类型，也可能引用了另一个数据结构。
	// 推荐使用IDE进行开发，可以方便的跳转查阅各个接口和数据结构的文档说明。
	request := soe.NewInitOralProcessRequest()

	// 基本类型的设置。
	// 此接口允许设置返回的实例数量。此处指定为只返回一个。
	// SDK采用的是指针风格指定参数，即使对于基本类型你也需要用指针来对参数赋值。
	// SDK提供对基本类型的指针引用封装函数
	request.WorkMode = common.Int64Ptr(WORK_MODE_ONCE)
	request.EvalMode = common.Int64Ptr(EVAL_MODE_WORD)
	request.ScoreCoeff = common.Float64Ptr(2.0)
	request.SessionId = common.StringPtr(sessionId)
	request.RefText = common.StringPtr("bike")

	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := client.InitOralProcess(request)
	// 处理异常
	fmt.Println(err)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		panic(err)
	}

	// 打印返回的json字符串
	fmt.Printf("%s", response.ToJsonString())

	// 实例化一个请求对象，根据调用的接口和实际情况，可以进一步设置请求参数
	// 你可以直接查询SDK源码确定TransmitOralProcessRequest有哪些属性可以设置，
	// 属性可能是基本类型，也可能引用了另一个数据结构。
	// 推荐使用IDE进行开发，可以方便的跳转查阅各个接口和数据结构的文档说明。
	tRequest := soe.NewTransmitOralProcessRequest()

	// 基本类型的设置。
	// 此接口允许设置返回的实例数量。此处指定为只返回一个。
	// SDK采用的是指针风格指定参数，即使对于基本类型你也需要用指针来对参数赋值。
	// SDK提供对基本类型的指针引用封装函数
	tRequest.SeqId = common.Int64Ptr(1)
	tRequest.IsEnd = common.Int64Ptr(1)
	tRequest.SessionId = common.StringPtr(sessionId)
	tRequest.VoiceFileType = common.Int64Ptr(MP3)
	tRequest.VoiceEncodeType = common.Int64Ptr(1)

	buf, err := ioutil.ReadFile(AUDIO_FOR_ONCE)
	if err != nil {
		fmt.Printf("read file error: %s", err)
		return
	}
	base64Str := base64.StdEncoding.EncodeToString(buf)
	tRequest.UserVoiceData = common.StringPtr(base64Str)
	tRequest.SetHttpMethod("POST")
	tResponse, err := client.TransmitOralProcess(tRequest)
	// 处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		panic(err)
	}
	// 打印返回的json字符串
	b, _ := json.Marshal(tResponse.Response)
	fmt.Printf("\n评测结果：%s\n", b)
	fmt.Println("--------结束一次性评测-----")

}

/**
 * 中文评估示例，除了init参数 ServerType 以外，和英文评估完全一样
 *
 * @param client
 */
func evaluationCn(client *soe.Client) {

	fmt.Println("--------中文评测-----")
	sessionId := getSessionId()
	// 实例化一个请求对象，根据调用的接口和实际情况，可以进一步设置请求参数
	// 你可以直接查询SDK源码确定InitOralProcessRequest有哪些属性可以设置，
	// 属性可能是基本类型，也可能引用了另一个数据结构。
	// 推荐使用IDE进行开发，可以方便的跳转查阅各个接口和数据结构的文档说明。
	request := soe.NewInitOralProcessRequest()

	// 基本类型的设置。
	// 此接口允许设置返回的实例数量。此处指定为只返回一个。
	// SDK采用的是指针风格指定参数，即使对于基本类型你也需要用指针来对参数赋值。
	// SDK提供对基本类型的指针引用封装函数
	request.WorkMode = common.Int64Ptr(WORK_MODE_ONCE)
	request.EvalMode = common.Int64Ptr(EVAL_MODE_SENTENCE)
	request.ScoreCoeff = common.Float64Ptr(2.0)
	request.SessionId = common.StringPtr(sessionId)
	request.RefText = common.StringPtr("轻轻的我走了，正如我轻轻的来")
	request.ServerType = common.Int64Ptr(CN)

	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := client.InitOralProcess(request)
	// 处理异常
	fmt.Println(err)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		panic(err)
	}

	// 打印返回的json字符串
	fmt.Printf("%s", response.ToJsonString())

	// 实例化一个请求对象，根据调用的接口和实际情况，可以进一步设置请求参数
	// 你可以直接查询SDK源码确定TransmitOralProcessRequest有哪些属性可以设置，
	// 属性可能是基本类型，也可能引用了另一个数据结构。
	// 推荐使用IDE进行开发，可以方便的跳转查阅各个接口和数据结构的文档说明。
	tRequest := soe.NewTransmitOralProcessRequest()

	// 基本类型的设置。
	// 此接口允许设置返回的实例数量。此处指定为只返回一个。
	// SDK采用的是指针风格指定参数，即使对于基本类型你也需要用指针来对参数赋值。
	// SDK提供对基本类型的指针引用封装函数
	tRequest.SeqId = common.Int64Ptr(1)
	tRequest.IsEnd = common.Int64Ptr(1)
	tRequest.SessionId = common.StringPtr(sessionId)
	tRequest.VoiceFileType = common.Int64Ptr(MP3)
	tRequest.VoiceEncodeType = common.Int64Ptr(1)

	buf, err := ioutil.ReadFile(AUDIO_FOR_CN)
	if err != nil {
		fmt.Printf("read file error: %s", err)
		return
	}
	base64Str := base64.StdEncoding.EncodeToString(buf)
	tRequest.UserVoiceData = common.StringPtr(base64Str)
	tRequest.SetHttpMethod("POST")
	tResponse, err := client.TransmitOralProcess(tRequest)
	// 处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		panic(err)
	}
	// 打印返回的json字符串
	b, _ := json.Marshal(tResponse.Response)
	fmt.Printf("%s\n", b)
	fmt.Println("--------结束中文评测-----")
}

/**
* 流式分片传输评估示例,中等长度的音频或者边录边传模式，将音频分片传输到评估服务
*
* @param client
 */
func evaluationMutipleSeq(client *soe.Client) {
	fmt.Println("\n--------顺序评测-----")
	sessionId := getSessionId()
	// 实例化一个请求对象，根据调用的接口和实际情况，可以进一步设置请求参数
	// 你可以直接查询SDK源码确定InitOralProcessRequest有哪些属性可以设置，
	// 属性可能是基本类型，也可能引用了另一个数据结构。
	// 推荐使用IDE进行开发，可以方便的跳转查阅各个接口和数据结构的文档说明。
	request := soe.NewInitOralProcessRequest()

	// 基本类型的设置。
	// 此接口允许设置返回的实例数量。此处指定为只返回一个。
	// SDK采用的是指针风格指定参数，即使对于基本类型你也需要用指针来对参数赋值。
	// SDK提供对基本类型的指针引用封装函数
	request.WorkMode = common.Int64Ptr(WORK_MODE_STREAM)
	request.EvalMode = common.Int64Ptr(EVAL_MODE_SENTENCE)
	request.ScoreCoeff = common.Float64Ptr(2.0)
	request.SessionId = common.StringPtr(sessionId)
	request.RefText = common.StringPtr("Rex was a big dog")

	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := client.InitOralProcess(request)
	// 处理异常
	fmt.Println(err)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		panic(err)
	}
	// 打印返回的json字符串
	fmt.Printf("%s", response.ToJsonString())
	data, err := ioutil.ReadFile(AUDIO_FOR_SEQ)
	if err != nil {
		fmt.Printf("read file error: %s", err)
		return
	}

	timeBefore := time.Now().UnixNano() / 1e6
	defer func() {
		timeAfter := time.Now().UnixNano() / 1e6
		fmt.Printf("costTime:%d", timeAfter-timeBefore)
	}()
	pkgNum := int(math.Ceil(float64(len(data)) / float64(PKG_SIZE)))
	for i := 1; i <= pkgNum; i++ {
		lastIndex := i * PKG_SIZE
		if i == pkgNum {
			lastIndex = len(data)
		}
		buf := data[(i-1)*PKG_SIZE : lastIndex]
		base64Str := base64.StdEncoding.EncodeToString(buf)

		tRequest := soe.NewTransmitOralProcessRequest()

		// 基本类型的设置。
		// 此接口允许设置返回的实例数量。此处指定为只返回一个。
		// SDK采用的是指针风格指定参数，即使对于基本类型你也需要用指针来对参数赋值。
		// SDK提供对基本类型的指针引用封装函数
		tRequest.SeqId = common.Int64Ptr(int64(i))
		tRequest.IsEnd = common.Int64Ptr(0)
		if i == pkgNum {
			tRequest.IsEnd = common.Int64Ptr(1)
		}
		tRequest.SessionId = common.StringPtr(sessionId)
		tRequest.VoiceFileType = common.Int64Ptr(WAV)
		tRequest.VoiceEncodeType = common.Int64Ptr(1)
		tRequest.UserVoiceData = common.StringPtr(base64Str)
		tRequest.SetHttpMethod("POST")
		tResponse, err := client.TransmitOralProcess(tRequest)
		// 处理异常
		if _, ok := err.(*errors.TencentCloudSDKError); ok {
			fmt.Printf("An API error has returned: %s", err)
			return
		}
		// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
		if err != nil {
			panic(err)
		}
		b, _ := json.Marshal(tResponse.Response)
		fmt.Printf("\n评测结果：%s\n", b)
	}

	tRequest := soe.NewTransmitOralProcessRequest()
	tRequest.SeqId = common.Int64Ptr(1)
	tRequest.IsEnd = common.Int64Ptr(1)
	tRequest.SessionId = common.StringPtr(sessionId)
	tRequest.VoiceFileType = common.Int64Ptr(1)
	// 设置请求为查询接口
	tRequest.IsQuery = common.Int64Ptr(1)
	tRequest.VoiceEncodeType = common.Int64Ptr(WAV)
	// 填充空值
	tRequest.UserVoiceData = common.StringPtr(" ")
	tRequest.SetHttpMethod("POST")
	tResponse, err := client.TransmitOralProcess(tRequest)
	// 处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		panic(err)
	}
	b, _ := json.Marshal(tResponse.Response)
	fmt.Printf("\n查询结果：%s\n", b)
	fmt.Println("--------结束评测-----")
}

/**
 * 乱序流式评估示例,分片请求支持乱序
 *
 * @param client
 */
func evaluationMultipleSeqDisorder(client *soe.Client) {
	fmt.Println("--------乱序评测-----")
	sessionId := getSessionId()
	// 实例化一个请求对象，根据调用的接口和实际情况，可以进一步设置请求参数
	// 你可以直接查询SDK源码确定InitOralProcessRequest有哪些属性可以设置，
	// 属性可能是基本类型，也可能引用了另一个数据结构。
	// 推荐使用IDE进行开发，可以方便的跳转查阅各个接口和数据结构的文档说明。
	request := soe.NewInitOralProcessRequest()

	// 基本类型的设置。
	// 此接口允许设置返回的实例数量。此处指定为只返回一个。
	// SDK采用的是指针风格指定参数，即使对于基本类型你也需要用指针来对参数赋值。
	// SDK提供对基本类型的指针引用封装函数
	request.WorkMode = common.Int64Ptr(WORK_MODE_STREAM)
	request.EvalMode = common.Int64Ptr(EVAL_MODE_SENTENCE)
	request.ScoreCoeff = common.Float64Ptr(2.0)
	request.SessionId = common.StringPtr(sessionId)
	request.RefText = common.StringPtr("Rex was a big dog.")

	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := client.InitOralProcess(request)
	// 处理异常
	fmt.Println(err)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		panic(err)
	}

	// 打印返回的json字符串
	fmt.Printf("%s", response.ToJsonString())

	data, err := ioutil.ReadFile(AUDIO_FOR_SEQ)
	if err != nil {
		fmt.Printf("read file error: %s", err)
		return
	}

	pkgNum := int(math.Ceil(float64(len(data)) / float64(PKG_SIZE)))

	timeBefore := time.Now().UnixNano() / 1e6
	defer func() {
		timeAfter := time.Now().UnixNano() / 1e6
		fmt.Printf("costTime:%d", timeAfter-timeBefore)
	}()

	wg := &sync.WaitGroup{}
	for i := 1; i <= pkgNum; i++ {
		lastIndex := i * PKG_SIZE
		if i == pkgNum {
			lastIndex = len(data)
		}
		buf := data[(i-1)*PKG_SIZE : lastIndex]
		base64Str := base64.StdEncoding.EncodeToString(buf)

		tRequest := soe.NewTransmitOralProcessRequest()

		// 基本类型的设置。
		// 此接口允许设置返回的实例数量。此处指定为只返回一个。
		// SDK采用的是指针风格指定参数，即使对于基本类型你也需要用指针来对参数赋值。
		// SDK提供对基本类型的指针引用封装函数
		tRequest.SeqId = common.Int64Ptr(int64(i))
		tRequest.IsEnd = common.Int64Ptr(0)
		if i == pkgNum {
			tRequest.IsEnd = common.Int64Ptr(1)
		}
		tRequest.SessionId = common.StringPtr(sessionId)
		tRequest.VoiceFileType = common.Int64Ptr(WAV)
		tRequest.VoiceEncodeType = common.Int64Ptr(1)
		tRequest.UserVoiceData = common.StringPtr(base64Str)
		tRequest.SetHttpMethod("POST")

		wg.Add(1)
		go func(g *sync.WaitGroup, tReq *soe.TransmitOralProcessRequest) {

			defer g.Done()
			tResponse, err := client.TransmitOralProcess(tReq)
			if _, ok := err.(*errors.TencentCloudSDKError); ok {
				fmt.Printf("An API error has returned: %s", err)
				return
			}
			b, _ := json.Marshal(tResponse.Response)
			fmt.Printf("\n评测结果：%s\n", b)
			if err != nil {
				fmt.Printf("trans error: %s", err)
			}
		}(wg, tRequest)
	}
	wg.Wait()

	tRequest := soe.NewTransmitOralProcessRequest()
	tRequest.SeqId = common.Int64Ptr(1)
	tRequest.IsEnd = common.Int64Ptr(1)
	tRequest.SessionId = common.StringPtr(sessionId)
	tRequest.VoiceFileType = common.Int64Ptr(1)
	// 设置请求为查询接口
	tRequest.IsQuery = common.Int64Ptr(1)
	tRequest.VoiceEncodeType = common.Int64Ptr(1)
	// 填充空值
	tRequest.UserVoiceData = common.StringPtr(" ")
	tRequest.SetHttpMethod("POST")
	tResponse, err := client.TransmitOralProcess(tRequest)
	// 处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		panic(err)
	}
	b, _ := json.Marshal(tResponse.Response)
	fmt.Printf("\n查询结果：%s\n", b)

	fmt.Println("结束评测")
}

/**
 * 异步评估示例,较长音频的评估，短时间内获取不了结果(评估时间超过 20s )，采用异步模式，主动轮询评估服务得到评估结果，可用于段落或者离线评估
 *
 * @param client
 */
func evaluationMutipleAsync(client *soe.Client) {
	fmt.Println("--------开始异步评测-----")
	sessionId := getSessionId()
	// 实例化一个请求对象，根据调用的接口和实际情况，可以进一步设置请求参数
	// 你可以直接查询SDK源码确定InitOralProcessRequest有哪些属性可以设置，
	// 属性可能是基本类型，也可能引用了另一个数据结构。
	// 推荐使用IDE进行开发，可以方便的跳转查阅各个接口和数据结构的文档说明。
	request := soe.NewInitOralProcessRequest()

	// 基本类型的设置。
	// 此接口允许设置返回的实例数量。此处指定为只返回一个。
	// SDK采用的是指针风格指定参数，即使对于基本类型你也需要用指针来对参数赋值。
	// SDK提供对基本类型的指针引用封装函数
	request.WorkMode = common.Int64Ptr(WORK_MODE_STREAM)
	request.EvalMode = common.Int64Ptr(EVAL_MODE_SENTENCE)
	request.ScoreCoeff = common.Float64Ptr(2.0)
	request.SessionId = common.StringPtr(sessionId)
	request.RefText = common.StringPtr("I see sixteen eggs and eighteen apples.wav")
	request.IsAsync = common.Int64Ptr(1)

	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := client.InitOralProcess(request)
	// 处理异常
	fmt.Println(err)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		panic(err)
	}

	// 打印返回的json字符串
	fmt.Printf("%s", response.ToJsonString())

	data, err := ioutil.ReadFile(AUDIO_FOR_ASYNC)
	if err != nil {
		fmt.Printf("read file error: %s", err)
		return
	}

	pkgNum := int(math.Ceil(float64(len(data)) / float64(PKG_SIZE)))
	wg := &sync.WaitGroup{}
	for i := 1; i <= pkgNum; i++ {
		lastIndex := i * PKG_SIZE
		if i == pkgNum {
			lastIndex = len(data)
		}
		buf := data[(i-1)*PKG_SIZE : lastIndex]
		base64Str := base64.StdEncoding.EncodeToString(buf)

		tRequest := soe.NewTransmitOralProcessRequest()

		// 基本类型的设置。
		// 此接口允许设置返回的实例数量。此处指定为只返回一个。
		// SDK采用的是指针风格指定参数，即使对于基本类型你也需要用指针来对参数赋值。
		// SDK提供对基本类型的指针引用封装函数
		tRequest.SeqId = common.Int64Ptr(int64(i))
		tRequest.IsEnd = common.Int64Ptr(0)
		if i == pkgNum {
			tRequest.IsEnd = common.Int64Ptr(1)
		}
		tRequest.SessionId = common.StringPtr(sessionId)
		tRequest.VoiceFileType = common.Int64Ptr(WAV)
		tRequest.VoiceEncodeType = common.Int64Ptr(1)
		tRequest.UserVoiceData = common.StringPtr(base64Str)
		tRequest.SetHttpMethod("POST")

		wg.Add(1)
		go func(g *sync.WaitGroup, tReq *soe.TransmitOralProcessRequest) {

			defer g.Done()
			tResponse, err := client.TransmitOralProcess(tReq)
			if _, ok := err.(*errors.TencentCloudSDKError); ok {
				fmt.Printf("An API error has returned: %s", err)
				return
			}
			b, _ := json.Marshal(tResponse.Response)
			fmt.Printf("\n评测结果：%s\n", b)
			if err != nil {
				fmt.Printf("trans error: %s", err)
			}
		}(wg, tRequest)
	}
	wg.Wait()

	tRequest := soe.NewTransmitOralProcessRequest()
	tRequest.SeqId = common.Int64Ptr(1)
	tRequest.IsEnd = common.Int64Ptr(1)
	tRequest.SessionId = common.StringPtr(sessionId)

	// 设置请求为查询接口
	tRequest.IsQuery = common.Int64Ptr(1)
	tRequest.VoiceFileType = common.Int64Ptr(WAV)
	tRequest.VoiceEncodeType = common.Int64Ptr(1)
	// 填充空值
	tRequest.UserVoiceData = common.StringPtr(" ")
	tRequest.SetHttpMethod("POST")

	for i := 0; i < 100; i++ {
		tResponse, err := client.TransmitOralProcess(tRequest)
		// 处理异常
		if _, ok := err.(*errors.TencentCloudSDKError); ok {
			fmt.Printf("An API error has returned: %s", err)
			return
		}
		// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
		if err != nil {
			panic(err)
		}
		b, _ := json.Marshal(tResponse.Response)
		fmt.Printf("\n查询结果：%s\n", b)
		status := *tResponse.Response.Status
		if status == FAILED || status == FINISHED {
			fmt.Println("异步评测结束")
			return

		}
		time.Sleep(10 * time.Second)
	}
}

/**
 * 传输数据带初始化评估示例，除了init参数 ServerType 以外，和英文评估完全一样
 *
 * @param client
 */
func evaluationTransmitOralProcessWithInit(client *soe.Client) {
	fmt.Println("--------带初始化数据传输评测-----")
	sessionId := getSessionId()
	// 实例化一个请求对象，根据调用的接口和实际情况，可以进一步设置请求参数
	// 你可以直接查询SDK源码确定InitOralProcessRequest有哪些属性可以设置，
	// 属性可能是基本类型，也可能引用了另一个数据结构。
	// 推荐使用IDE进行开发，可以方便的跳转查阅各个接口和数据结构的文档说明。
	request := soe.NewTransmitOralProcessWithInitRequest()

	// 基本类型的设置。
	// 此接口允许设置返回的实例数量。此处指定为只返回一个。
	// SDK采用的是指针风格指定参数，即使对于基本类型你也需要用指针来对参数赋值。
	// SDK提供对基本类型的指针引用封装函数
	request.WorkMode = common.Int64Ptr(WORK_MODE_ONCE)
	request.EvalMode = common.Int64Ptr(EVAL_MODE_WORD)
	request.ScoreCoeff = common.Float64Ptr(2.0)
	request.SessionId = common.StringPtr(sessionId)
	request.RefText = common.StringPtr("bike")
	request.SeqId = common.Int64Ptr(1)
	request.IsEnd = common.Int64Ptr(1)
	request.SessionId = common.StringPtr(sessionId)
	request.VoiceFileType = common.Int64Ptr(MP3)
	request.VoiceEncodeType = common.Int64Ptr(1)

	buf, err := ioutil.ReadFile(AUDIO_FOR_ONCE)
	if err != nil {
		fmt.Printf("read file error: %s", err)
		return
	}
	base64Str := base64.StdEncoding.EncodeToString(buf)
	request.UserVoiceData = common.StringPtr(base64Str)
	request.SetHttpMethod("POST")
	response, err := client.TransmitOralProcessWithInit(request)
	// 处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		panic(err)
	}
	// 打印返回的json字符串
	b, _ := json.Marshal(response.Response)
	fmt.Printf("\n评测结果：%s\n", b)
	fmt.Println("--------结束带初始化数据传输评测-----")

}

func main() {

	credential := common.NewCredential(
		SECRET_ID,
		SECRET_KEY,
	)

	// 非必要步骤
	// 实例化一个客户端配置对象，可以指定超时时间等配置
	cpf := profile.NewClientProfile()
	// SDK默认使用POST方法。
	// 如果你一定要使用GET方法，可以在这里设置。GET方法无法处理一些较大的请求。
	cpf.HttpProfile.ReqMethod = "POST"
	// SDK有默认的超时时间，非必要请不要进行调整。
	// 如有需要请在代码中查阅以获取最新的默认值。
	cpf.HttpProfile.ReqTimeout = 30
	/**
	 *  设置访问域名，如果需要就近部署，可以使用 soe-tencentcloudapi.com, 腾讯云将根据访问的地域解析到合适的服务器上，如果调用服务已确定地域，如华南地区
	 *  可以直接使用地域域名，加快访问速度
	 */
	cpf.HttpProfile.Endpoint = "soe.ap-shanghai.tencentcloudapi.com"

	// 实例化要请求产品的client对象
	// 第二个参数是地域信息
	client, _ := soe.NewClient(credential, "ap-guangzhou", cpf)

	evaluationOnce(client)
	evaluationCn(client)
	evaluationMutipleSeq(client)
	evaluationMultipleSeqDisorder(client)
	evaluationMutipleAsync(client)
	evaluationTransmitOralProcessWithInit(client)

}
