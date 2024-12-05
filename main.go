package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID int32 `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float32 `json:"price"`
}


var albums = []album{
	{ID: 1, Title: "This is how tomorrow goes", Artist: "Beabadoobee", Price: 13.98},
	{ID: 2, Title: "Pandora", Artist: "Wisp", Price: 11.59},
	{ID: 3, Title: "D>E>A>T>H>M>E>T>A>L", Artist: "Panchiko", Price: 89.10},
}

func main(){
	router:= gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}


func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)
}

