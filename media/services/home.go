package services

import (
	"github.com/gin-gonic/gin"

	"code.tianchanghao.org/index/logger"
)

type HomeHandler struct {
	Router         *gin.Engine
	FilePrePath    string
	Exit           chan bool
}

func (h *HomeHandler) Get(c *gin.Context) {
	logger.Info("Into Get().")
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

