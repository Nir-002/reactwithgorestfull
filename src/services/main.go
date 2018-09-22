package main

import (
	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
func PostHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Posted",
	})
}

func QueryString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func PathParameterString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func main() {
	r := gin.Default()

	r.GET("/", HomePage)
	r.GET("/query", QueryString)
	r.GET("/path/:name/:age", PathParameterString)
	r.POST("/", PostHomePage)
	r.Run("2000")
}
