package service

import (
	"fmt"
	"miniprogram/pkg/model"
	"strings"

	"github.com/xops-infra/noop/log"
)

type WebhookService struct {
	Io   model.WebHookIo
	Conf []model.Condition
}

func NewWebhookService(webhookio model.WebHookIo, conf []model.Condition) *WebhookService {
	return &WebhookService{
		Io:   webhookio,
		Conf: conf,
	}
}

func (w *WebhookService) CallBack(signature, timestamp, nonce string, input model.WebHookRequest) (model.WebHookResponse, error) {
	plainText, res, err := w.Io.CallBack(signature, timestamp, nonce, input)
	if err != nil {
		return res, err
	}
	// 解析plainText 处理业务逻辑
	for _, condition := range w.Conf {
		if strings.Contains(plainText, string(condition.EventType)) {
			// 依据不同的事件类型，执行不同的操作
			var req model.ServerApiReq
			switch condition.EventType {
			case model.BPMS_INSTANCE_CHANGE:
				var bpmsInstanceEvent model.BpmsInstanceEvent
				err = bpmsInstanceEvent.UnmarshalJSON([]byte(plainText))
				if err != nil {
					return res, err
				}
				err = w.Io.BpmsEvent(bpmsInstanceEvent, condition)
				if err != nil {
					log.Errorf("bpms event failed: %s", err)
					continue
				}
				req.Result = bpmsInstanceEvent.Result
				req.ProcessInstanceId = bpmsInstanceEvent.ProcessInstanceId
			default:
				return res, fmt.Errorf("not support key: %s", condition.EventType)
			}

			// 执行操作
			err = w.Io.Execute(req, condition.Actions)
			if err != nil {
				log.Errorf("execute actions failed: %+v", condition.Actions)
			}
			log.Infof("execute actions success: %+v", condition.Actions)
		}
	}
	return res, nil
}
