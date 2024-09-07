package routes

import (
	"CRUD/go/controller"
	"github.com/gin-gonic/gin"
)

//根据请求url进行转发

func SetRouter() *gin.Engine {
	r := gin.Default()

	stuGroup := r.Group("/stu")
	{
		stuGroup.POST("/", controller.CreateStudent)
		stuGroup.GET("/", controller.GetStuList)
		stuGroup.PUT("/:id", controller.UpdateStudent)
		stuGroup.DELETE("/:id", controller.DeleteStudent)
		stuGroup.GET("/:id", controller.GetStudentByID)
	}

	return r
}
