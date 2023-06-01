package routes

import (
	"E-Commerce/controllers"

	"github.com/gin-gonic/gin"
)

func AddressRoutes(router *gin.Engine) {
	router.POST("/addaddress", controllers.AddAddress())
	router.PUT("/edithomeaddress", controllers.EditHomeAddress())
	router.PUT("/editworkaddress", controllers.EditWorkAddress())
	router.DELETE("/deleteaddresses", controllers.DeleteAddress())
}
