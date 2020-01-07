package action

import (
	g "github.com/gin-gonic/gin"
)

// 获取公钥私钥
func GenerateRSAKey(c *g.Context) {
	generateRSAKey(c)
}

// 获取公钥私钥
func GenerateRSAKeyforPem(c *g.Context) {
	generateRSAKeyforPem(c)
}

// encryption
// 获取公钥私钥
func Encryption(c *g.Context) {
	encryption(c)
}

// 测试get

func TestGet(c *g.Context) {
	testGet(c)
}
func TestPost(c *g.Context) {
	testPost(c)
}
