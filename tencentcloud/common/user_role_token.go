package common

import (
	"log"
	"time"
)

type UserRoleTokenCredential struct {
	expiredTime  uint64
	expiration   string
	tmpSecretId  string
	tmpSecretKey string
	token        string
	source       *UserRoleTokenProvider
}

func (c *UserRoleTokenCredential) GetSecretId() string {
	if c.needRefresh() {
		c.refresh()
	}
	return c.tmpSecretId
}

func (c *UserRoleTokenCredential) GetToken() string {
	if c.needRefresh() {
		c.refresh()
	}
	return c.token
}

func (c *UserRoleTokenCredential) GetSecretKey() string {
	if c.needRefresh() {
		c.refresh()
	}
	return c.tmpSecretKey
}

func (c *UserRoleTokenCredential) GetCredential() (string, string, string) {
	if c.needRefresh() {
		c.refresh()
	}
	return c.tmpSecretId, c.tmpSecretKey, c.token
}

func (c *UserRoleTokenCredential) needRefresh() bool {
	if c.tmpSecretId == "" || c.tmpSecretKey == "" || c.token == "" || int64(c.expiredTime-ExpiredTimeout) <= time.Now().Unix() {
		return true
	}
	return false
}

func (c *UserRoleTokenCredential) refresh() {
	newCre, err := c.source.GetCredential()
	if err != nil {
		log.Println(err)
		return
	}
	*c = *newCre.(*UserRoleTokenCredential)
}
