package media

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"miao_sticker_server/media/services"
)

var handler *services.HomeHandler

func NewHandler(router *gin.Engine) *services.HomeHandler {
	handler = &services.HomeHandler{
		Router:      router,
		FilePrePath: strings.Replace(os.Getenv("GOPATH"), "\\", "/", -1) + "/src/miao_sticker/media/resources/",
		Exit:        make(chan bool),
	}
	return handler
}

func InitRouters(router *gin.Engine) {
	// 注册路由策略
	router.GET("/GET", handler.Get)
	router.GET("/POST", handler.Post)
	router.GET("/DELETE", handler.Delete)
	router.POST("/PUT", handler.Put)
}
