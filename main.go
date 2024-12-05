package main

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     int32   `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

var albums = []album{
	{ID: 1, Title: "This is how tomorrow goes", Artist: "Beabadoobee", Price: 13.98},
	{ID: 2, Title: "Pandora", Artist: "Wisp", Price: 11.59},
	{ID: 3, Title: "D>E>A>T>H>M>E>T>A>L", Artist: "Panchiko", Price: 89.10},
}

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumbByID)

	router.POST("/postAlbums", postAlbums)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumbByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, a := range albums {
		if a.ID == int32(id) {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Album not found"})
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
