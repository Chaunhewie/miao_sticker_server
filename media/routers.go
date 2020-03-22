package media

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"miao_sticker_server/index/logger"
	"miao_sticker_server/media/constdef"
	"miao_sticker_server/media/models"
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
		UserInfoOpenIdCache: make(map[string]*models.UserInfo),
	}
	if err := handler.UpdateRepoInfo(); err != nil {
		logger.Error("UpdateRepoInfo Error: %v", err.Error())
	}
	return handler
}

func InitRouters(router *gin.Engine) {
	// 注册路由策略
	router.GET("/", handler.Get)
	router.POST("/", handler.Post)
	router.DELETE("/", handler.Delete)
	router.PUT("/", handler.Put)
	router.POST("/login", handler.PostLogin)
}
