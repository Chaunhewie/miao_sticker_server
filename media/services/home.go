package services

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"miao_sticker_server/index/logger"
)

type HomeHandler struct {
	Router         *gin.Engine
	FilePrePath    string
	Exit           chan bool
}

func (h *HomeHandler) Get(c *gin.Context) {
	logger.Info("Into Get().")
	print(c.Params)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"extra":   "数据上传失败，文件写入失败！",
	})
	logger.Info("Out Get().")
	return
}

func (h *HomeHandler) Post(c *gin.Context) {
	logger.Info("Into Post().")
	logger.Info("Out Post().")
	return
}

func (h *HomeHandler) Delete(c *gin.Context) {
	logger.Info("Into Delete().")
	logger.Info("Out Delete().")
	return
}

func (h *HomeHandler) Put(c *gin.Context) {
	logger.Info("Into Put().")
	logger.Info("Out Put().")
	return
}

