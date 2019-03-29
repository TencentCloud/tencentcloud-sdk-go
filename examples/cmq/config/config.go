package config

import "fmt"

var (
	// os.Getenv("TENCENTCLOUD_SECRET_ID"),
	// os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	SecretID              = "AKIDAC8vVtJ4mEHB6okgvvgBV3QaUklpdhOz"
	SecretKey             = "DGoIbi5NUWAe4JMAnFUopQsAH8jVGo8C"
	QueueEndpoint         = fmt.Sprintf("cmq-queue-%s.api.qcloud.com", "bj")
	TopicEndpoint         = fmt.Sprintf("cmq-topic-%s.api.qcloud.com", "bj")
	IntranetQueueEndpoint = fmt.Sprintf("cmq-queue-%s.api.tencentyun.com", "bj")
	IntranetTopicEndpoint = fmt.Sprintf("cmq-topic-%s.api.tencentyun.com", "bj")
)
