package main

import (
	"golang-mysql/mysql"
	service "golang-mysql/static"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	go func() {
		mysql.MysqlInit()
	}()
	r.GET("/", service.ShowAllUser)
	r.POST("/create", service.CreateUser)
	r.GET("/update/:id", service.UpdateUser)
	r.GET("/delete/:id", service.DeleteUser)
	r.Run(":8081")
}
