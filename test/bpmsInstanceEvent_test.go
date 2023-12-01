package test

import (
	"fmt"
	"miniprogram/pkg/io"
	"miniprogram/pkg/model"
	"os"
	"testing"
	"time"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/patsnapops/noop/log"
)

var event_example = `{"processInstanceId":"kNpmUavZQwSHP_ZJvVdL7g07561689326476","finishTime":1689326506000,"corpId":"ding1b99eee8b9ef5edf","EventType":"bpms_instance_change","businessId":"202307141721000164728","title":"周守健提交的云管专用-勿走钉钉提交-对象存储同步","type":"finish","url":"https://aflow.dingtalk.com/dingtalk/mobile/homepage.htm?corpid=ding1b99eee8b9ef5edf&dd_share=false&showmenu=false&dd_progress=false&back=native&procInstId=kNpmUavZQwSHP_ZJvVdL7g07561689326476&taskId=&swfrom=isv&dinghash=approval&dtaction=os&dd_from=corp#approval","result":"refuse","createTime":1689326476000,"processCode":"PROC-B85623B4-A372-4684-BB61-1B7E046CE9A8","bizCategoryId":"","businessType":"","staffId":"29070242122092575562"}`

var bpmsCondition = model.Condition{
	EventType: model.BPMS_INSTANCE_CHANGE,
	JudgeField: []model.Field{
		{
			Name:  "type",
			Value: "finish",
		}, {
			Name:  "processCode",
			Value: "PROC-B85623B4-A372-4684-BB61-1B7E046CE9A8",
		},
	},
	Actions: []model.Action{
		{
			Type:  model.HTTP_POST,
			Value: "http://cbs.patsnap.info/api/v1/webhook",
		},
	},
}

var (
	webhookIo         model.WebHookIo
	bpmsInstanceEvent model.BpmsInstanceEvent
)

func init() {
	log.Default().WithLevel(log.DebugLevel).WithHumanTime(time.Local).Init()
	token := os.Getenv("token")
	encodingAESKey := os.Getenv("encodingAESKey")
	suiteKey := os.Getenv("suiteKey")
	webhookIo = io.NewWebHookClient(token, encodingAESKey, suiteKey)
	fmtBpms()
}

func fmtBpms() {
	err := bpmsInstanceEvent.UnmarshalJSON([]byte(event_example))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tea.Prettify(bpmsInstanceEvent))
}

func TestCondition(t *testing.T) {
	err := webhookIo.BpmsEvent(bpmsInstanceEvent, bpmsCondition)
	if err != nil {
		log.Infof("bpms event failed: %s", err)
		return
	}
	err = webhookIo.Execute(model.ServerApiReq{
		Result:            bpmsInstanceEvent.Result,
		ProcessInstanceId: bpmsInstanceEvent.ProcessInstanceId,
	}, bpmsCondition.Actions)
	if err != nil {
		t.Errorf("execute failed: %s", err)
	}
}

func TestNoCondition(t *testing.T) {
	bpmsCondition.JudgeField = append(bpmsCondition.JudgeField, model.Field{
		Name:  "type",
		Value: "123",
	})
	err := webhookIo.BpmsEvent(bpmsInstanceEvent, bpmsCondition)
	if err == nil {
		t.Errorf("bpms event failed: %s", err)
	}
}
