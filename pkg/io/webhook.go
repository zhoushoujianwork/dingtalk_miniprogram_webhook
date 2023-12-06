package io

import (
	"fmt"
	"miniprogram/pkg/model"
	"net/http"

	"github.com/alibabacloud-go/tea/tea"
	go_dingtalk "github.com/icepy/go-dingtalk/src"
	"github.com/xops-infra/noop/log"
)

type webhookClient struct {
	crpyto *go_dingtalk.DingTalkCrypto
}

func NewWebHookClient(token, encodingAESKey, suiteKey string) model.WebHookIo {
	crpyto := go_dingtalk.NewDingTalkCrypto(token, encodingAESKey, suiteKey)
	return &webhookClient{
		crpyto: crpyto,
	}
}

func (w *webhookClient) CallBack(signature, timestamp, nonce string, input model.WebHookRequest) (string, model.WebHookResponse, error) {
	var output model.WebHookResponse
	// 解密
	plainText, err := w.crpyto.GetDecryptMsg(signature, timestamp, nonce, input.Encrypt)
	if err != nil {
		return plainText, output, err
	}
	log.Infof("get plainText: %s", plainText)
	encrypt, signatureOut, err := w.crpyto.GetEncryptMsg("success", timestamp, nonce)
	if err != nil {
		return plainText, output, err
	}
	output.Encrypt = encrypt
	output.MsgSignature = signatureOut
	output.Timestamp = timestamp
	output.Nonce = nonce
	return plainText, output, nil
}

func (w *webhookClient) BpmsEvent(bpmsInstanceEvent model.BpmsInstanceEvent, condition model.Condition) error {
	// 判断是否符合条件
	IsOk := true
	for _, v := range condition.JudgeField {
		switch v.Name {
		case "EventType":
			if bpmsInstanceEvent.EventType != v.Value {
				IsOk = false
			}
		case "processInstanceId":
			if bpmsInstanceEvent.ProcessInstanceId != v.Value {
				IsOk = false
			}
		case "corpId":
			if bpmsInstanceEvent.CorpId != v.Value {
				IsOk = false
			}
		case "title":
			if bpmsInstanceEvent.Title != v.Value {
				IsOk = false
			}
		case "type":
			if bpmsInstanceEvent.Type != v.Value {
				IsOk = false
			}
		case "staffId":
			if bpmsInstanceEvent.StaffId != v.Value {
				IsOk = false
			}
		case "processCode":
			if bpmsInstanceEvent.ProcessCode != v.Value {
				IsOk = false
			}
		case "result":
			if bpmsInstanceEvent.Result != v.Value {
				IsOk = false
			}
		default:
			return fmt.Errorf("not support field: %s", v.Name)
		}
	}
	if !IsOk {
		return fmt.Errorf("not match condition")
	}
	return nil
}

// 支持多个action
func (w *webhookClient) Execute(req model.ServerApiReq, actions []model.Action) error {
	for _, v := range actions {
		switch v.Type {
		case model.HTTP_POST:
			//TODO: 带上表单的详细信息 Post给约定的接口
			log.Infof("post to %s with form: %v\n", v.Type, v.Value)
			// req to io.Reader
			log.Debugf(tea.Prettify(req))
			httpClient := &http.Client{}
			req, err := http.NewRequest("POST", v.Value, req.ToReader())
			if err != nil {
				return err
			}
			req.Header.Set("Content-Type", "application/json")
			resp, err := httpClient.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				return fmt.Errorf("status code is not 200")
			}
		default:
			return fmt.Errorf("not support action: %s", v.Type)
		}
	}
	return nil
}
