package common

import (
	"log"
	"sync"
	"time"
)

type RoleArnCredential struct {
	roleArn         string
	roleSessionName string
	durationSeconds int64
	expiredTime     int64
	token           string
	tmpSecretId     string
	tmpSecretKey    string
	source          Provider
	mu              sync.Mutex
}

func (c *RoleArnCredential) GetSecretId() string {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.needRefresh() {
		c.refresh()
	}
	return c.tmpSecretId
}

func (c *RoleArnCredential) GetSecretKey() string {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.needRefresh() {
		c.refresh()
	}
	return c.tmpSecretKey

}

func (c *RoleArnCredential) GetToken() string {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.needRefresh() {
		c.refresh()
	}
	return c.token
}

func (c *RoleArnCredential) GetCredential() (string, string, string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.needRefresh() {
		c.refresh()
	}
	return c.tmpSecretId, c.tmpSecretKey, c.token
}

func (c *RoleArnCredential) needRefresh() bool {
	if c.tmpSecretKey == "" || c.tmpSecretId == "" || c.token == "" || c.expiredTime <= time.Now().Unix() {
		return true
	}
	return false
}

func (c *RoleArnCredential) refresh() {
	newCre, err := c.source.GetCredential()
	if err != nil {
		log.Println(err)
		return
	}

	newCred := newCre.(*RoleArnCredential)
	c.roleArn = newCred.roleArn
	c.roleSessionName = newCred.roleSessionName
	c.durationSeconds = newCred.durationSeconds
	c.expiredTime = newCred.expiredTime
	c.token = newCred.token
	c.tmpSecretId = newCred.tmpSecretId
	c.tmpSecretKey = newCred.tmpSecretKey
	c.source = newCred.source
}
