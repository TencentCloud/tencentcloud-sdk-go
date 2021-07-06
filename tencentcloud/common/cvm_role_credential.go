package common

import (
	"time"
)

const ExpiredTimeout = 300

type CvmRoleCredential struct {
	RoleName     string
	ExpiredTime  int64
	tmpSecretId  string
	tmpSecretKey string
	token        string
	source       *CvmRoleProvider
}

func (c *CvmRoleCredential) GetSecretId() string {
	if c.needRefresh() {
		c.refresh()
	}
	return c.tmpSecretId
}

func (c *CvmRoleCredential) GetToken() string {
	if c.needRefresh() {
		c.refresh()
	}
	return c.token
}

func (c *CvmRoleCredential) GetCredentialParams() map[string]string {
	p := map[string]string{
		"SecretId": c.GetSecretId(),
	}
	if c.token != "" {
		p["Token"] = c.GetToken()
	}
	return p
}

func (c *CvmRoleCredential) GetSecretKey() string {
	if c.needRefresh() {
		c.refresh()
	}
	return c.tmpSecretKey
}

func (c *CvmRoleCredential) needRefresh() bool {
	if c.tmpSecretId == "" || c.tmpSecretKey == "" || c.token == "" || c.ExpiredTime-ExpiredTimeout <= time.Now().Unix() {
		return true
	}
	return false
}

func (c *CvmRoleCredential) refresh() {
	newCre, err := c.source.GetCredential()
	if err != nil {
		//todo: how to handle this err?
		panic(err)
	}
	//todo: maybe not safe
	*c = *newCre.(*CvmRoleCredential)
}
