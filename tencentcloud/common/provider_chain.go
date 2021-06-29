package common

import (
	"errors"
)

type Chain struct {
	Providers []Provider
}

func NewProviderChain(providers []Provider) Provider {
	return &Chain{
		Providers: providers,
	}
}

func DefaultChain() Provider {
	return NewProviderChain([]Provider{DefaultEnvProvider(), DefaultProfileProvider(), DefaultCvmRoleProvider()})
}

func (c *Chain) GetCredential() (CredentialIface, error) {
	for _, provider := range c.Providers {
		cred, err := provider.GetCredential()
		if err != nil {
			return nil, err
		} else if cred == nil {
			continue
		}
		return cred, err
	}
	return nil, errors.New("no credential found")

}
