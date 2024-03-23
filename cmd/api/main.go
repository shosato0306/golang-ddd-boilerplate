package main

import (
	"github.com/shosato0306/golang-ddd-boilerplate/internal/app/api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/example", handler.ExampleHandler)

	r.Run()
}
