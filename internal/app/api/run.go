package api

import (
	"github.com/gin-gonic/gin"

	"github.com/shosato0306/golang-ddd-boilerplate/internal/app/api/handler"
)

func Run() {
	r := gin.Default()

	r.GET("/example", handler.ExampleHandler)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
