package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	c.JSON(http.StatusOK, "hej")

}

func main() {

	router := gin.Default()

	router.GET("/api/test", test)
	router.Run(":8080")
}
