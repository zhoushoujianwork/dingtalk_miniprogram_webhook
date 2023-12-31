package app

import (
	"os"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/spf13/viper"
	"github.com/xops-infra/noop/log"

	"miniprogram/pkg/model"
)

type Config struct {
	Conditions []model.Condition `json:"conditions" mapstructure:"conditions"`
}

var Conf *Config

func InitConfig(path string) {
	Conf = &Config{}
	// 判断文件是否存在
	_, err := os.ReadFile(path)
	if err != nil {
		log.Panicf(err.Error())
	}
	// 读取配置文件
	viper.SetConfigFile(path)
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(Conf)
	if err != nil {
		panic(err)
	}
	log.Debugf(tea.Prettify(Conf))
}
