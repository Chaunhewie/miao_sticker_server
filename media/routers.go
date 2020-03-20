package media

import (
	"miao_sticker_server/media/constdef"
	"miao_sticker_server/media/models"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"miao_sticker_server/media/services"
)

var handler *services.HomeHandler

func NewHandler(router *gin.Engine) *services.HomeHandler {
	handler = &services.HomeHandler{
		RepoFilePath: strings.Replace(os.Getenv("GOPATH"), "\\", "/", -1) + constdef.RESOURCES_PATH +
			constdef.REPO_FILE_NAME,
		RepoInfo: &models.RepoInfo{
			StargazersCount: 0,
			WatchersCount:   0,
			Forks:           0,
		},
		Router:      router,
		FilePrePath: strings.Replace(os.Getenv("GOPATH"), "\\", "/", -1) + constdef.RESOURCES_PATH,
		Exit:        make(chan bool),
	}
	return handler
}

func InitRouters(router *gin.Engine) {
	// 注册路由策略
	router.GET("/", handler.Get)
	router.POST("/", handler.Post)
	router.DELETE("/", handler.Delete)
	router.PUT("/", handler.Put)
}
