package model

type Condition struct {
	EventType  EventType `json:"event_type" mapstructure:"event_type"`
	JudgeField []Field   `json:"judge_field" mapstructure:"judge_field"`
	Actions    []Action  `json:"action" mapstructure:"actions"`
}

type Field struct {
	Name  string `json:"name" mapstructure:"name"`
	Value string `json:"value" mapstructure:"value"`
}

type Action struct {
	Type  ActionType `json:"type" mapstructure:"type"`
	Value string     `json:"value" mapstructure:"value"`
}

type ActionType string

type EventType string

const (
	// http Post请求
	HTTP_POST ActionType = "http_post"

	// 事件类型
	BPMS_INSTANCE_CHANGE EventType = "bpms_instance_change"
)
