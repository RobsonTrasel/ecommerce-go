package routes

import (
	"ecommerce-go/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incRoutes *gin.Engine) {
	incRoutes.POST("/users/signup", controllers.SignUp())
	incRoutes.POST("/users/login", controllers.Login())
	incRoutes.POST("/admin/addproduct", controllers.ProductViewAdmin())
	incRoutes.GET("/users/productview", controllers.SearchProduct())
	incRoutes.GET("/users/search", controllers.SearchProductByView())
}
