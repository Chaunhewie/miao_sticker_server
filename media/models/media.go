package models

/************************服务端结构定义****************************/
type UserInfo struct {
	NickName   string `json:"nick_name"`
	AvatarUrl  string `json:"avatar_url"`
	Gender     int32  `json:"gender"`
	Code       string `json:"code"`
	OpenId     string `json:"openid"`  // 针对此小程序的唯一ID
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"` // 针对所有小程序的唯一ID
}

type RepoInfo struct {
	StargazersCount int32 `json:"stargazers_count"`
	WatchersCount   int32 `json:"watchers_count"`
	Forks           int32 `json:"forks_count"`
}

/************************前端请求接口定义****************************/
type ReqLoginData struct {
	Type      string `form:"type" json:"type" binding:"required"`
	Code      string `form:"code" json:"code" binding:"required"`
	NickName  string `form:"nick_name" json:"nick_name" binding:"required"`
	AvatarUrl string `form:"avatar_url" json:"avatar_url" binding:"required"`
	Gender    int32  `form:"gender" json:"gender" binding:"required"`
}

/************************后端请求服务，返回结构定义****************************/
type RespUserInfoData struct {
	OpenId     string `json:"openid"`
	Sessionkey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    string `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}
