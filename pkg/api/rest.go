package api

import (
	"miniprogram/app"
	"miniprogram/pkg/model"

	_ "miniprogram/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
	"github.com/xops-infra/ginx/middleware"
	hh "github.com/xops-infra/http-headers"
	"github.com/xops-infra/noop/log"
)

func InitGin(debug bool) *gin.Engine {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
		log.Debugf("gin start in debug mode")
	}

	ginEngine := gin.Default()
	middleware.AttachTo(ginEngine).
		WithCacheDisabled().
		WithCORS().
		WithRecover().
		WithRequestID(hh.XRequestID).
		WithSecurity()
	// add hc
	ginEngine.GET("/hc", HealthCheck)
	ginEngine.GET("/swagger/*any", func(c *gin.Context) {
		c.Next()
	}, ginswagger.WrapHandler(swaggerfiles.Handler))
	ApplyRoutes(ginEngine.Group("/v2023-03"))
	return ginEngine
}

func ApplyRoutes(routerGroup *gin.RouterGroup) {

	r := routerGroup.Group("/webhook")
	// callback
	r.POST("/callback", CallBack)

}

// HealthCheck godoc
func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

// @Summary webhook callback
// @Description webhook callback
// @Tags webhook
// @Accept  json
// @Produce  json
// @Param signature query string true "signature"
// @Param timestamp query string true "timestamp"
// @Param nonce query string true "nonce"
// @Param input body model.WebHookRequest true "input"
// @Success 200 {object} string
// @Failure 500 {object} string
// @Router /v2023-03/webhook/callback [post]
func CallBack(c *gin.Context) {
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	request := model.WebHookRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Errorf("bind json failed %s", err.Error())
		c.JSON(500, err.Error())
		return
	}
	resp, err := app.State.Webhook.CallBack(signature, timestamp, nonce, request)
	if err != nil {
		c.JSON(500, err.Error())
		log.Errorf("callback failed %s", err.Error())
		return
	}
	c.JSON(200, resp)
}
