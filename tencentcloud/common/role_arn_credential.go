package common

import (
	"time"
)

type RoleArnCredential struct {
	RoleArn         string
	RoleSessionName string
	DurationSeconds int64
	ExpiredTime     int64
	token           string
	tmpSecretId     string
	tmpSecretKey    string
	source          *RoleArnProvider
}

func (c *RoleArnCredential) GetSecretId() string {
	if c.needRefresh() {
		c.refresh()
	}
	return c.tmpSecretId
}

func (c *RoleArnCredential) GetSecretKey() string {
	if c.needRefresh() {
		c.refresh()
	}
	return c.tmpSecretKey

}

func (c *RoleArnCredential) GetToken() string {
	if c.needRefresh() {
		c.refresh()
	}
	return c.token
}

func (c *RoleArnCredential) GetCredentialParams() map[string]string {
	p := map[string]string{
		"SecretId": c.GetSecretId(),
	}
	if c.token != "" {
		p["token"] = c.GetToken()
	}
	return p
}

func (c *RoleArnCredential) needRefresh() bool {
	if c.tmpSecretKey == "" || c.tmpSecretId == "" || c.token == "" || c.ExpiredTime <= time.Now().Unix() {
		return true
	}
	return false
}

func (c *RoleArnCredential) refresh() {
	newCre, err := c.source.GetCredential()
	if err != nil {
		//todo: how to handle this err?
		panic(err)
	}
	*c = *newCre.(*RoleArnCredential)
}
