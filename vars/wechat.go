package vars

import "fmt"

var (
	MiniProgramAppId            = "wx099c8f4930db46b3"
	MiniProgramAppSecret        = "7a0a812f0d5e1c8d39af97246bc745f2"
	MiniProgramHost             = "https://api.weixin.qq.com"
	MiniProgramHostCode2Session = fmt.Sprintf("%s/sns/jscode2session", MiniProgramHost)
)
