package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonas-granbom/uppgift1/data"
	"gorm.io/gorm"
)

func test(c *gin.Context) {
	c.JSON(http.StatusOK, "hej")

}

func myName(c *gin.Context) {
	name := "Jonas"
	city := "Gothenburg"

	c.JSON(http.StatusOK, name+", "+city)
}
func seed(c *gin.Context) {
	data.SeedData()
	c.JSON(http.StatusOK, "db filled with random data")
}

func getTeams(c *gin.Context) {
	var teams []data.Team
	data.DB.Find(&teams)
	c.JSON(http.StatusOK, teams)

}

func getPlayerById(c *gin.Context) {
	var player data.Player
	id := c.Param("id")

	err := data.DB.First(&player, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found in database"})
	} else {
		c.JSON(http.StatusOK, player)
	}

}

func main() {
	data.ConnectDatabase()
	router := gin.Default()

	router.GET("/api/test", test)
	router.GET("/api/jonas", myName)
	router.GET("/api/seed", seed)
	router.GET("/api/getteams", getTeams)
	router.GET("/api/players/:id", getPlayerById)

	router.Run(":8080")
}
