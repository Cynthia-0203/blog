package flash

import (
	"encoding/gob"

	"github.com/Cynthia/goblog/pkg/session"
	"github.com/gin-gonic/gin"
)


type Flashes map[string]interface{}

// 存入会话数据里的 key
var flashKey = "_flashes"

func init() {
    // 在 gorilla/sessions 中存储 map 和 struct 数据需
    // 要提前注册 gob，方便后续 gob 序列化编码、解码
    gob.Register(Flashes{})
}


func Info(c *gin.Context,message string) {
    addFlash(c,"info", message)
}


func Warning(c *gin.Context,message string) {
    addFlash(c,"warning", message)
}


func Success(c *gin.Context,message string) {
    addFlash(c,"success", message)
}


func Danger(c *gin.Context,message string) {
    addFlash(c,"danger", message)
}


func All(c *gin.Context) Flashes {
    val := session.Get(c,flashKey)
    
    flashMessages, ok := val.(Flashes)
    if !ok {
        return nil
    }
    
    session.Forget(c,flashKey)
    return flashMessages
}

func addFlash(c *gin.Context,key string, message string) {
    flashes := Flashes{}
    flashes[key] = message
    session.Put(c,flashKey, flashes)
    session.Save(c)
}