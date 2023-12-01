package model

import (
	"bytes"
	"encoding/json"
	"io"
)

// 提交 ID 和结果给其他服务 API 接口
type ServerApiReq struct {
	ProcessInstanceId string `json:"processInstanceId" binding:"required"`
	Result            string `json:"result" binding:"required"`
}

// req to io.Reader
func (s *ServerApiReq) ToReader() io.Reader {
	b, _ := json.Marshal(s)
	return bytes.NewReader(b)
}
