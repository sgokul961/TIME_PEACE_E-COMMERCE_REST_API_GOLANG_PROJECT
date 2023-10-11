package http

import (
	"github.com/gin-gonic/gin"
	"gokul.go/pkg/api/handler"
	"gokul.go/pkg/routes"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(adminHandler *handler.AdminHandler, userHandler *handler.UserHandler, otpHandler *handler.OtpHandler, inventoryHandler *handler.InventoryHandler, categoryHandler *handler.CategoryHandler, cartHandler *handler.CartHandler, orderHandler *handler.OrdeHandler) *ServerHTTP {
	engine := gin.New()

	//use loggger from gin

	engine.Use(gin.Logger())

	routes.AdminRoutes(engine.Group("/admin"), adminHandler, inventoryHandler, categoryHandler, orderHandler)
	routes.UserRoutes(engine.Group("/user"), userHandler, otpHandler, inventoryHandler, cartHandler, orderHandler)

	return &ServerHTTP{engine: engine}

}
func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
