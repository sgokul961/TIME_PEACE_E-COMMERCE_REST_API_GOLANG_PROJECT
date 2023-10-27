package routes

import (
	"github.com/gin-gonic/gin"
	"gokul.go/pkg/api/handler"
	"gokul.go/pkg/api/middleware"
)

func UserRoutes(engine *gin.RouterGroup, userHandler *handler.UserHandler, otpHandler *handler.OtpHandler, inventoryHandler *handler.InventoryHandler, cartHandler *handler.CartHandler, orderHandler *handler.OrdeHandler, paymentHandler *handler.PaymentHandler) {
	engine.POST("/signup", userHandler.UserSignUp)
	engine.POST("/login", userHandler.LoginHandler)
	engine.GET("/forgot-password", userHandler.ForgotPasswordSend)
	engine.POST("/forgot-passwordnew", userHandler.VarifyForgotPasswordAndChange)

	engine.POST("/otplogin", otpHandler.SendOTP)
	engine.POST("/verifyotp", otpHandler.VerifyOTP)
	payment := engine.Group("/payment")
	{
		payment.GET("/razorpay", paymentHandler.MakePaymentRazorPay)
		payment.GET("/update_status", paymentHandler.VerifyPayment)
	}

	engine.Use(middleware.UserAuthMiddleware)

	{
		home := engine.Group("/home")
		{
			home.GET("/product", inventoryHandler.ListProducts)
			home.GET("/products_show", inventoryHandler.ShowIndividualProducts)
			home.POST("/add-to-cart", cartHandler.AddToCart)

		}

		profile := engine.Group("/profile")
		{
			profile.GET("details", userHandler.GetUserDetails)
			profile.GET("/address", userHandler.GetAddress)
			profile.POST("/addaddress", userHandler.AddAddress)
			profile.GET("get_link", userHandler.GetMyReferanceLink)

			security := profile.Group("/security")
			{
				security.PUT("/change-password", userHandler.ChangePassword)
			}

			orders := engine.Group("/orders")
			{
				orders.GET("/", orderHandler.GetOrders)
				orders.DELETE("/", orderHandler.CancelOrder)
			}

		}

		// checkouts := engine.Group("/check")//command this
		// {
		// 	checkouts.GET("/", cartHandler.CheckOut)
		// }
		cart := engine.Group("/cart")
		{
			cart.GET("/", userHandler.GetCart)
			cart.DELETE("/remove", userHandler.RemoveFromCart)
			cart.PUT("/updatequantity-add", userHandler.UpdateQuantityAdd)
			cart.PUT("/updatequantity-minus", userHandler.UpdateQuantityMinus)
		}
		checkout := engine.Group("/check-out")
		{
			checkout.POST("/order", orderHandler.OrderItemsFromCart)
			checkout.GET("/invoice", orderHandler.GenerateInvoice)
		}

	}

}
