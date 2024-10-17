package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var sessionKey = "session_name"

// SetupSessionMiddleware 设置 Gin 的 Session 中间件
func SetupSessionMiddleware(secret string) gin.HandlerFunc {
	// 使用 CookieStore 作为会话存储方式
	store := cookie.NewStore([]byte(secret))
	return sessions.Sessions(sessionKey, store)
}

// Put 将键值对存入会话
func Put(c *gin.Context, key string, value interface{}) {
	session := sessions.Default(c)
	session.Set(key, value)
	session.Save()
}

// Get 从会话中获取值
func Get(c *gin.Context, key string) interface{} {
	session := sessions.Default(c)
	return session.Get(key)
}

// Forget 从会话中删除键
func Forget(c *gin.Context, key string) {
	session := sessions.Default(c)
	session.Delete(key)
	session.Save()
}

func Save(c *gin.Context){
	session := sessions.Default(c)
	session.Save()
}
// Flush 清空会话
func Flush(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}
