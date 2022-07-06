package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/quesiasts/gin-school/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/:name", controllers.Introduction)
	r.POST("/students", controllers.CreateNewStudent)
	r.GET("/students", controllers.ListAllStudents)
	r.GET("/students/:id", controllers.SearchForID)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.EditStudent)
	r.GET("/students/document/:document", controllers.SearchForDocument)
	r.GET("/index", controllers.ShowIndexPage)
	r.NoRoute(controllers.RouteNotFound)
	r.Run()
}
