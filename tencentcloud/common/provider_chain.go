package common

import (
	"errors"
)

type ProviderChain struct {
	Providers []Provider
}

func NewProviderChain(providers []Provider) Provider {
	return &ProviderChain{
		Providers: providers,
	}
}

func DefaultChain() Provider {
	return NewProviderChain([]Provider{DefaultEnvProvider(), DefaultProfileProvider(), DefaultCvmRoleProvider()})
}

func (c *ProviderChain) GetCredential() (CredentialIface, error) {
	for _, provider := range c.Providers {
		cred, err := provider.GetCredential()
		if err != nil {
			if errors.Is(err, ENVNOTSET) || errors.Is(err, FILEDOSENOTEXIST) || errors.Is(err, CVMNOROLE) {
				continue
			} else {
				return nil, err
			}
		}
		return cred, err
	}
	return nil, errors.New("no credential found in each provider")

}
