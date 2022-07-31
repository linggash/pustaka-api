package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello",helloHendler)
	router.GET("/books/:id/:title", booksHandler)
	router.GET("/query", queryHandler)

	router.POST("/books", postBooksHandler)

	router.Run(":8888")
}

func rootHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"name": "Lingga Sempana",
		"bio" : "A Software Engineer & Content Creator",
	})	
}

func helloHendler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"title": "Hello World",
		"subtitle" : "Belajar Golang sendiri anjayyy",
	})
}

func booksHandler(c *gin.Context){
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func queryHandler(c *gin.Context){
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{"title": title, "price" : price})
}

func postBooksHandler(c *gin.Context){
	//title, price
	
}