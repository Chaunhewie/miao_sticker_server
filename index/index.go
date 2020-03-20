package index

import (
	"os"
	"strconv"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"miao_sticker_server/index/constdef"
	"miao_sticker_server/index/logger"
	"miao_sticker_server/media"
	mediaServices "miao_sticker_server/media/services"
)

type MyApp struct {
	ProjectPath  string
	Router       *gin.Engine
	HomeHandler  *mediaServices.HomeHandler
}

func (app *MyApp) Init() {
	app.ProjectPath = strings.Replace(os.Getenv("GOPATH"), "\\", "/", -1) + "/src/miao_sticker_server"
	// 设置模式
	if constdef.APP_DEBUG_MODE {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	// 注册一个默认的路由器
	app.Router = gin.Default()
	// 设置使用session
	store := cookie.NewStore([]byte(constdef.COOKIE_NAME))
	app.Router.Use(sessions.Sessions(constdef.SESSION_NAME, store))
	// 注册路由规则
	app.registerRouters()
}

func (app *MyApp) registerRouters() {
	logger.Info("Register Routers...")
	app.HomeHandler = media.NewHandler(app.Router)
	media.InitRouters(app.Router)
}

func (app *MyApp) Run() {
	// 绑定端口是8080
	logger.Info("Begin to run...")
	if err := app.Router.Run(":"+strconv.Itoa(constdef.PORT)); err != nil {
		logger.Error("Run Error: %v", err)
	}
}
