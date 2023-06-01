package routes

import (
	"E-Commerce/controllers"
	"E-Commerce/database"

	"github.com/gin-gonic/gin"
)

func CartRoutes(router *gin.Engine) {

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())
	router.GET("/listcart", controllers.GetItemFromCart())
}
