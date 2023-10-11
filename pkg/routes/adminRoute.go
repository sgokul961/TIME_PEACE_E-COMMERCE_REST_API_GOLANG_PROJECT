package routes

import (
	"github.com/gin-gonic/gin"
	"gokul.go/pkg/api/handler"
	"gokul.go/pkg/api/middleware"
)

func AdminRoutes(engine *gin.RouterGroup, adminHandler *handler.AdminHandler, inventoryHandler *handler.InventoryHandler, categoryHandler *handler.CategoryHandler, orderhandler *handler.OrdeHandler) {
	engine.POST("/adminlogin", adminHandler.LoginHandler)
	engine.Use(middleware.AdminAuthMiddleware)
	{
		usermanagement := engine.Group("/users")
		{
			usermanagement.POST("/block", adminHandler.BlockUser)
			usermanagement.POST("/unblock", adminHandler.UnblockUser)
			usermanagement.GET("/getusers", adminHandler.GetUsers)
		}
		inventoryManagement := engine.Group("/inventories")
		{

			inventoryManagement.POST("/add", inventoryHandler.AddInventory)
			inventoryManagement.PUT("/update", inventoryHandler.UpdateInventory)
			inventoryManagement.DELETE("/delete", inventoryHandler.DeleteInventory)
		}
		categorymanagement := engine.Group("/category")
		{
			categorymanagement.POST("/add", categoryHandler.AddCategory)
			categorymanagement.PUT("/update/:id", categoryHandler.UpdateCategory)
			categorymanagement.DELETE("/delete", categoryHandler.DeleteCategory)
		}
		orders := engine.Group("/orders")
		{
			orders.GET("/", orderhandler.AdminOrders)
			orders.PUT("/edit/status", orderhandler.EditOrderStatus)
		}

	}

}
