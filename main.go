package main

import (
	"ecommerce-go/controllers"
	"ecommerce-go/database"
	"ecommerce-go/middlewares"
	"ecommerce-go/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))
	router := gin.New()
	router.use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middlewares.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.instantbuy())

	log.Fatal(router.Run(":" + port))

}
