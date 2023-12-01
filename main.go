package main

import (
	"flag"
	"miniprogram/app"
	"miniprogram/pkg/api"

	"github.com/patsnapops/noop/log"
)

var (
	debug       bool
	port        string
	log_file    string // file path
	config_file string // config path
	version     string // version
)

func init() {
	flag.BoolVar(&debug, "debug", false, "debug mode")
	flag.StringVar(&port, "port", "8080", "port")
	flag.StringVar(&log_file, "log_file", "/var/log/miniprogram.log", "log path")
	flag.StringVar(&config_file, "config_file", "./config.yaml", "config path")
	flag.Parse()
	if debug {
		log.Default().WithLevel(log.DebugLevel).WithFilename(log_file).Init()
		log.Debugf("debug mode")
	} else {
		log.Default().WithLevel(log.InfoLevel).WithFilename(log_file).Init()
	}
}

// @title           mini program webhook API
// @version         2023-03
// @description     Patsnap OPS API spec.
// @termsOfService  http://swagger.io/terms/
// @contact.name    DevOps Team
// @host            localhost:8080
// @BasePath
func main() {
	log.Infof("mini program webhook API %s", version)
	app.InitConfig(config_file)
	app.InitServices()
	gin := api.InitGin(debug)
	err := gin.Run(":" + port)
	if err != nil {
		log.Panicf("gin start failed %s", err.Error())
	}
}
