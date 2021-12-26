package routes

import (
	"github.com/Delaram-Gholampoor-Sagha/Golang-Ecommerce-Project/controllers"
	"github.com/gin-gonic/gin"
)

func UserRputes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview", controllers.SearchProduct())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())
}
