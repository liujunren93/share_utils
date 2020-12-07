package userStore

import (
	"encoding/json"
	"github.com/liujunren93/share_utils/metadata"
)

type Permission struct {
	Name   string `json:"name"`
	Code   string `json:"code"`
	NameEn string `json:"name_en"`
	Method string `json:"method"`
	Path   string `json:"path"`
}

//LoginInfo 用户登录信息
type LoginInfo struct {
	metadata.UserAgent
	CreateAt int64 `json:"create_at"`
}

//encode 编码
func encode(data *LoginInfo) (string, error) {
	marshal, err := json.Marshal(data)
	return string(marshal), err
}

//decode 解码
func decode(str []byte) (*LoginInfo, error) {
	var data LoginInfo
	err := json.Unmarshal(str, &data)
	return &data, err
}
