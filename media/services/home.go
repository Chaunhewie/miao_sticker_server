package services

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"miao_sticker_server/media/constdef"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"miao_sticker_server/index/logger"
	"miao_sticker_server/media/models"
	"miao_sticker_server/media/utils"
)

type HomeHandler struct {
	RepoFilePath string
	RepoInfo     *models.RepoInfo

	UserInfoOpenIdCache map[string]*models.UserInfo
}

func (h *HomeHandler) Get(c *gin.Context) {
	logger.Info("Into Get().")
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

func (h *HomeHandler) PostLogin(c *gin.Context) {
	logger.Info("Into PostLogin().")

	// 获取登录所携带的data字段
	loginData := &models.ReqLoginData{}
	if err := c.BindJSON(&loginData); err != nil {
		logger.Error("BindJSON loginData Error: %v.", err.Error())
	}
	logger.Info("BindJSON body: type:[%v] code:[%v] nickName:[%v] avatarUrl:[%v] gender:[%v].", loginData.Type,
		loginData.Code, loginData.NickName, loginData.AvatarUrl, loginData.Gender)

	// 获取用户唯一OpenId
	var respUserInfoData *models.RespUserInfoData
	var err error
	if respUserInfoData, err = h.UserLogin(loginData); err != nil {
		logger.Error("UserLogin failed: %v.", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "登录失败，请稍后再试！",
		})
		logger.Info("Out PostLogin() with Error.")
		return
	}
	userInfo := &models.UserInfo{
		NickName:   loginData.NickName,
		AvatarUrl:  loginData.AvatarUrl,
		Gender:     loginData.Gender,
		Code:       loginData.Code,
		OpenId:     respUserInfoData.OpenId,
		SessionKey: respUserInfoData.Sessionkey,
		UnionId:    respUserInfoData.UnionId,
	}
	h.UserInfoOpenIdCache[userInfo.OpenId] = userInfo
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     "登录成功",
	})
	logger.Info("Out PostLogin().")
	return
}

func (h *HomeHandler) UserLogin(loginData *models.ReqLoginData) (respUserInfoData *models.RespUserInfoData, err error) {
	var respBytes []byte
	respBytes, err = utils.GetCode2Session(loginData.Code)
	if err != nil {
		logger.Error("GetCode2Session Error: %v.", err.Error())
		return nil, err
	}
	respUserInfoData = &models.RespUserInfoData{}
	if err = json.Unmarshal(respBytes, respUserInfoData); err != nil {
		logger.Error("Unmarshal respUserInfoData Error: %v.", err.Error())
		return nil, err
	}
	return respUserInfoData, nil
}

// 从文件读取RepoInfo至h.RepoInfo结构中
func (h *HomeHandler) UpdateRepoInfo() error {
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

func (h *HomeHandler) FetchRepoInfoLoop() {
	for true {
		// 获取 github API 数据
		resp, err := utils.Fetch(constdef.REPO_URL, nil)
		if err != nil {
			logger.Error("Fetch Repo Info Error: %v", err.Error())
			time.Sleep(constdef.DURATION_REPO_REQ_FAILED)
			continue
		}

		// 提取信息
		h.RepoInfo.StargazersCount = int32(gjson.Get(string(resp), "stargazers_count").Num)
		h.RepoInfo.WatchersCount = int32(gjson.Get(string(resp), "watchers_count").Num)
		h.RepoInfo.Forks = int32(gjson.Get(string(resp), "forks_count").Num)
		logger.Info("UpdateRepoInfo Success: %v", h.RepoInfo)

		// 写入文件备份
		repoInfoBytes, err := json.Marshal(h.RepoInfo)
		if err != nil {
			logger.Error("Marshal Repo Info Error: %v", err.Error())
			time.Sleep(constdef.DURATION_REPO_REQ_FAILED)
			continue
		}
		f, err := os.OpenFile(h.RepoFilePath+"__", os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			logger.Error("Open Repo Info Tmp File Error: %v", err.Error())
			time.Sleep(constdef.DURATION_REPO_REQ_FAILED)
			continue
		}
		if _, err = f.Write(repoInfoBytes); err != nil {
			logger.Error("Write Repo Info Error: %v", err.Error())
			time.Sleep(constdef.DURATION_REPO_REQ_FAILED)
			continue
		}
		if err = f.Close(); err != nil {
			logger.Error("Close Repo Info Tmp File Error: %v", err.Error())
			time.Sleep(constdef.DURATION_REPO_REQ_FAILED)
			continue
		}
		if err = os.Rename(h.RepoFilePath+"__", h.RepoFilePath); err != nil {
			logger.Error("Rename Repo Info File Error: %v", err.Error())
			time.Sleep(constdef.DURATION_REPO_REQ_FAILED)
			continue
		}
		logger.Info("Fetch Repo Info Success")
		time.Sleep(constdef.DURATION_REPO_REQ_SUCCESS)
	}
}
