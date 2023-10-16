package http

import (
	"github.com/gin-gonic/gin"
	"gokul.go/pkg/api/handler"
	"gokul.go/pkg/routes"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(adminHandler *handler.AdminHandler, userHandler *handler.UserHandler, otpHandler *handler.OtpHandler, inventoryHandler *handler.InventoryHandler, categoryHandler *handler.CategoryHandler, cartHandler *handler.CartHandler, orderHandler *handler.OrdeHandler, paymentHandler *handler.PaymentHandler, reporthandler *handler.SalesHandler) *ServerHTTP {
	engine := gin.New()
	engine.LoadHTMLGlob("templates/*.html")

	//use loggger from gin

	engine.Use(gin.Logger())

	routes.AdminRoutes(engine.Group("/admin"), adminHandler, inventoryHandler, categoryHandler, orderHandler, *reporthandler)
	routes.UserRoutes(engine.Group("/user"), userHandler, otpHandler, inventoryHandler, cartHandler, orderHandler, paymentHandler)

	return &ServerHTTP{engine: engine}

}
func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
