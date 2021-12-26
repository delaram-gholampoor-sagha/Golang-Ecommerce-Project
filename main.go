package main

import (
	"log"
	"os"

	"github.com/Delaram-Gholampoor-Sagha/Golang-Ecommerce-Project/controllers"
	"github.com/Delaram-Gholampoor-Sagha/Golang-Ecommerce-Project/database"
	"github.com/Delaram-Gholampoor-Sagha/Golang-Ecommerce-Project/middleware"
	"github.com/Delaram-Gholampoor-Sagha/Golang-Ecommerce-Project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)

	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InsatnBuy())

	log.Fatal(router.Run(":" + port))
}
