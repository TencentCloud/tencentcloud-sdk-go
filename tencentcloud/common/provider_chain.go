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
			if errors.Is(err, ENVNOTSET) || errors.Is(err, FILEDOSENOTEXIST) || errors.Is(err, CVMNOROLE) {
				continue
			} else {
				return nil, err
			}
		}
		return cred, err
	}
	return nil, errors.New("no credential found")

}
