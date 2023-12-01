package model

import (
	"encoding/json"
)

// https://open.dingtalk.com/document/orgapp/push-events?spm=a2q3p.21071111.0.0.89c41cfal50m0f

type WebHookIo interface {
	CallBack(signature, timestamp, nonce string, input WebHookRequest) (string, WebHookResponse, error)
	BpmsEvent(event BpmsInstanceEvent, condition Condition) error
	Execute(req ServerApiReq, actions []Action) error
}

type WebHookRequest struct {
	// encrypt
	Encrypt string `json:"encrypt"`
}

type WebHookResponse struct {
	// encrypt
	Encrypt string `json:"encrypt"`
	// msg_signature
	MsgSignature string `json:"msg_signature"`
	// timeStamp
	Timestamp string `json:"timeStamp"`
	// nonce
	Nonce string `json:"nonce"`
}

// 订阅事件结构体 bpms_instance_change
type BpmsInstanceEvent struct {
	EventType         string `json:"EventType"`
	ProcessInstanceId string `json:"processInstanceId"`
	CorpId            string `json:"corpId"`
	CreateTime        int64  `json:"createTime"`
	Title             string `json:"title"`
	Type              string `json:"type"`
	StaffId           string `json:"staffId"`
	URL               string `json:"url"`
	ProcessCode       string `json:"processCode"`
	// result
	Result string `json:"result"`
	// finishTime
	FinishTime int64 `json:"finishTime"`
}

// UnmarshalJSON
func (b *BpmsInstanceEvent) UnmarshalJSON(data []byte) error {
	type Alias BpmsInstanceEvent
	tmp := &struct {
		*Alias
	}{
		Alias: (*Alias)(b),
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	return nil
}
