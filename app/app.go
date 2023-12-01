package app

import (
	"miniprogram/pkg/io"
	"miniprogram/pkg/service"
	"os"
)

type Services struct {
	Webhook *service.WebhookService
}

var (
	State *Services
)

// you should export token, encodingAESKey, suiteKey in your env
func InitServices() {
	token := os.Getenv("token")
	encodingAESKey := os.Getenv("encodingAESKey")
	suiteKey := os.Getenv("suiteKey")
	webhookIo := io.NewWebHookClient(token, encodingAESKey, suiteKey)
	State = &Services{
		Webhook: service.NewWebhookService(webhookIo, Conf.Conditions),
	}
}
