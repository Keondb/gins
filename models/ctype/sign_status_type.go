package ctype

import "encoding/json"

type SignStatus int

const (
	SignQQ     SignStatus = 1 //qq
	SignGitee  SignStatus = 2 //gitee
	SignGithub SignStatus = 3 //github
	SignEmail  SignStatus = 4 //email
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {
	var str string
	switch s {
	case SignQQ:
		str = "qq"
	case SignGitee:
		str = "gitee"
	case SignGithub:
		str = "github"
	case SignEmail:
		str = "email"
	default:
		str = "其他"
	}
	return str
}
