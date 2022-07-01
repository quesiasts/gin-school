package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/quesiasts/gin-school/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/:name", controllers.Introduction)
	r.Run()
}
