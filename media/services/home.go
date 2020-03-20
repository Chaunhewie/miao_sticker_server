package services

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"miao_sticker_server/media/models"
	"net/http"
	"os"

	"miao_sticker_server/index/logger"
)

type HomeHandler struct {
	RepoFilePath string
	RepoInfo *models.RepoInfo
	Router       *gin.Engine
	FilePrePath  string
	Exit         chan bool
}

func (h *HomeHandler) Get(c *gin.Context) {
	logger.Info("Into Get().")
	if err := h.GetRepoInfo(); err != nil{
		c.JSON(http.StatusOK, gin.H{
			"success":          false,
			"stargazers_count": -1,
			"watchers_count":   -1,
			"forks":            -1,
		})
		logger.Error("Out Get() With Error: %v.", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success":          true,
		"stargazers_count": h.RepoInfo.StargazersCount,
		"watchers_count":   h.RepoInfo.WatchersCount,
		"forks_count":      h.RepoInfo.Forks,
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

func (h *HomeHandler) GetRepoInfo() error {
	f, err := os.OpenFile(h.RepoFilePath, os.O_RDONLY, 0666)
	if err != nil {
		logger.Error("Open Repo File Error: %v", err.Error())
		return err
	}
	defer f.Close()
	var repoInfoBytes []byte
	if repoInfoBytes, err = ioutil.ReadAll(f); err != nil {
		logger.Error("Read Repo File Error: %v", err.Error())
		return err
	}
	logger.Info("%v", string(repoInfoBytes))
	if err = json.Unmarshal(repoInfoBytes, h.RepoInfo); err != nil {
		logger.Error("Unmarshal Repo Info Error: %v", err.Error())
		return err
	}
	return nil
}