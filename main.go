package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	c.JSON(http.StatusOK, "hej")

}

func myName(c *gin.Context) {
	name := "Jonas"
	city := "Gothenburg"

	c.JSON(http.StatusOK, name+", "+city)
}

func main() {

	router := gin.Default()

	router.GET("/api/test", test)
	router.GET("/api/jonas", myName)

	router.Run(":8080")
}
