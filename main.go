package main

import (
	ent "github.com/api-rest-go/commons/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

var albums = []ent.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Year: 1957},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Year: 1962},
	{ID: "3", Title: "Ray Bradbury", Artist: "451 Fahrenheit ", Year: 1962},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbums(c *gin.Context) {
	var newAlbum ent.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func putAlbum(c *gin.Context) {
	var updateAlbum ent.Album

	if err := c.BindJSON(&updateAlbum); err != nil {
		return
	}

	if (updateAlbum == ent.Album{} || updateAlbum.ID == "") {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid album"})
		return
	}

	var updatedAlbum ent.Album
	for _, a := range albums {
		if a.ID == updateAlbum.ID {
			updatedAlbum = a
			break
		}
	}

	if (updatedAlbum == ent.Album{}) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	updatedAlbum.Artist = updateAlbum.Artist
	updatedAlbum.Title = updateAlbum.Title
	updatedAlbum.Year = updateAlbum.Year

	for i, a := range albums {
		if a.ID == updatedAlbum.ID {
			albums[i] = updatedAlbum
			break
		}
	}

	c.IndentedJSON(http.StatusOK, updatedAlbum)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/album", postAlbums)
	router.PUT("/album", putAlbum)

	router.Run("localhost:8080")
}
