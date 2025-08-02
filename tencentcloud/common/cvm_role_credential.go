package common

import (
	"log"
	"sync"
	"time"
)

const ExpiredTimeout = 300

type CvmRoleCredential struct {
	roleName     string
	expiredTime  int64
	tmpSecretId  string
	tmpSecretKey string
	token        string
	source       *CvmRoleProvider
	mu           sync.Mutex
}

func (c *CvmRoleCredential) GetSecretId() string {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.needRefresh() {
		c.refresh()
	}
	return c.tmpSecretId
}

func (c *CvmRoleCredential) GetToken() string {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.needRefresh() {
		c.refresh()
	}
	return c.token
}

func (c *CvmRoleCredential) GetSecretKey() string {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.needRefresh() {
		c.refresh()
	}
	return c.tmpSecretKey
}

func (c *CvmRoleCredential) GetCredential() (string, string, string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.needRefresh() {
		c.refresh()
	}
	return c.tmpSecretId, c.tmpSecretKey, c.token
}

func (c *CvmRoleCredential) needRefresh() bool {
	if c.tmpSecretId == "" || c.tmpSecretKey == "" || c.token == "" || c.expiredTime-ExpiredTimeout <= time.Now().Unix() {
		return true
	}
	return false
}

func (c *CvmRoleCredential) refresh() {
	newCre, err := c.source.GetCredential()
	if err != nil {
		log.Println(err)
		return
	}

	newCred := newCre.(*CvmRoleCredential)
	c.roleName = newCred.roleName
	c.expiredTime = newCred.expiredTime
	c.tmpSecretId = newCred.tmpSecretId
	c.tmpSecretKey = newCred.tmpSecretKey
	c.token = newCred.token
	c.source = newCred.source
}
