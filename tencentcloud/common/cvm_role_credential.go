package common

import (
	"log"
	"sync"
	"time"
)

const ExpiredTimeout = 300

type CvmRoleCredential struct {
	mu           sync.RWMutex
	roleName     string
	expiredTime  int64
	tmpSecretId  string
	tmpSecretKey string
	token        string
	source       *CvmRoleProvider
}

func (c *CvmRoleCredential) GetSecretId() string {
	c.mu.RLock()
	if c.needRefresh() {
		c.mu.RUnlock()
		c.refresh()
		c.mu.RLock()
	}
	defer c.mu.RUnlock()
	return c.tmpSecretId
}

func (c *CvmRoleCredential) GetToken() string {
	c.mu.RLock()
	if c.needRefresh() {
		c.mu.RUnlock()
		c.refresh()
		c.mu.RLock()
	}
	defer c.mu.RUnlock()
	return c.token
}

func (c *CvmRoleCredential) GetSecretKey() string {
	c.mu.RLock()
	if c.needRefresh() {
		c.mu.RUnlock()
		c.refresh()
		c.mu.RLock()
	}
	defer c.mu.RUnlock()
	return c.tmpSecretKey
}

func (c *CvmRoleCredential) GetCredential() (string, string, string) {
	c.mu.RLock()
	if c.needRefresh() {
		c.mu.RUnlock()
		c.refresh()
		c.mu.RLock()
	}
	defer c.mu.RUnlock()
	return c.tmpSecretId, c.tmpSecretKey, c.token
}

func (c *CvmRoleCredential) needRefresh() bool {
	if c.tmpSecretId == "" || c.tmpSecretKey == "" || c.token == "" || c.expiredTime-ExpiredTimeout <= time.Now().Unix() {
		return true
	}
	return false
}

func (c *CvmRoleCredential) refresh() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if !c.needRefresh() {
		return
	}
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
