package api

import "github.com/gin-gonic/gin"

func routing(d Dependency) error {
	router := gin.Default()

	router.GET("/user/:id", func(c *gin.Context) {
		d.apiHandler.GetUser(c)
	})

	err := router.Run(":8080")
	if err != nil {
		return err
	}

	return nil
}
