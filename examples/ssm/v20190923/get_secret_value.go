package main

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	ssm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssm/v20190923"
)

func main() {
	provider, err := common.DefaultTkeOIDCRoleArnProvider()
	if err != nil {
		panic(err)
	}

	credential, err := provider.GetCredential()
	if err != nil {
		panic(err)
	}

	cpf := profile.NewClientProfile()

	client, _ := ssm.NewClient(credential, regions.Guangzhou, cpf)

	request := ssm.NewGetSecretValueRequest()
	request.SecretName = common.StringPtr("SECRET_NAME")
	request.VersionId = common.StringPtr("1.0.0")

	response, err := client.GetSecretValue(request)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.ToJsonString())
}
