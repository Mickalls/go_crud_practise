package main

import (
	"CRUD/go/dao"
	"CRUD/go/entity"
	"CRUD/go/routes"
)

func main() {
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	defer dao.Close()

	dao.SqlConn.AutoMigrate(&entity.Student{})
	r := routes.SetRouter()

	r.Run(":8081")
}
