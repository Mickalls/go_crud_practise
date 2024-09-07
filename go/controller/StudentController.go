package controller

import (
	"CRUD/go/entity"
	"CRUD/go/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateStudent(c *gin.Context) {
	var stu entity.Student

	//将调用后端的request请求中的body数据根据json格式解析到User结构变量中
	c.BindJSON(&stu)

	//调用Service层的创建Student函数
	err := service.CreateStudent(&stu)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": stu,
	})
}

func GetStuList(c *gin.Context) {
	StuList, err := service.GetAllStudent()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": StuList,
	})
}

func UpdateStudent(c *gin.Context) {
	var stu entity.Student
	c.BindJSON(&stu)

	idStr := c.Param("id") // 获取URL中的"id"参数，而不是用stu的id

	//调用service层的修改函数
	err := service.UpdateStudentByID(idStr, &stu)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

func DeleteStudent(c *gin.Context) {
	studentID := c.Param("id")

	//调用service层的删除函数
	err := service.DeleteStudentByID(studentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

func GetStudentByID(c *gin.Context) {
	studentID := c.Param("id")

	//调用service层id查询函数
	stu, err := service.GetStudentByID(studentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": stu,
	})
}
