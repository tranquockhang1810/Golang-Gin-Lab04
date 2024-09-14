package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tranquockhang1810/Golang-Gin-Lab04/controllers"
)

func StudentRoute(r *gin.Engine) {
	r.GET("/students", controllers.GetStudents)
	r.GET("/students/:id", controllers.GetStudentById)
	r.POST("/students", controllers.CreateStudent)
	r.PUT("/students/:id", controllers.UpdateStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
}