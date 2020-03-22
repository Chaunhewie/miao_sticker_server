package constdef

import "time"

const (
	REPO_URL       = "https://api.github.com/repos/chaunhewie/miao_sticker" // 获取仓库信息的URL
	RESOURCES_PATH = "/src/miao_sticker_server/media/resources/"            // 存储仓库信息的文件路径
	REPO_FILE_NAME = "repo_info"                                            // 存储仓库信息的文件路径

	DURATION_REPO_REQ_FAILED  = time.Duration(10) * time.Second   // 获取仓库信息失败后睡眠的时间
	DURATION_REPO_REQ_SUCCESS = time.Duration(1800) * time.Second // 获取仓库信息成功后睡眠的时间

	Code2Session_URL = "https://api.weixin.qq.com/sns/jscode2session" // 微信获取授权URL
	APPID            = "wxc11e2201bad6ccfd"                           // 小程序ID
	APPSecret        = "bdaaeeb7e102bb88688be2ef2c794043"             // 小程序密钥
	GRANT_TYPE       = "authorization_code"                           // 授权类型
)
