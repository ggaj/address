package routers

import (
	"github.com/ggaj/address/src/api/controllers"
	"github.com/gin-gonic/gin"
)

func RouteInit(app controllers.AppController) {
	r := gin.Default()

	r.POST("address", app.AddressController.CreateAddress)

	r.Run()
}
