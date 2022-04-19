package integration

import (
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
)

func TestProviderChain(t *testing.T) {
	// custom provider chain
	//
	//	provider1 := common.DefaultCvmRoleProvider()
	//	provider2 := common.DefaultEnvProvider()
	//	customProviders := []common.Provider{provider1, provider2}
	//	pc := common.NewProviderChain(customProviders)

	// default provider
	pc := common.DefaultProviderChain()
	_, err := pc.GetCredential()
	if err != nil {
		t.Error(err)
	}
}
