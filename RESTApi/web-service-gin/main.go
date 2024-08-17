package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
)

//Structu Albuns, the anotation after each type is to make the json more ✨pretty✨
type album struct {
	ID	string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

//List of albuns
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

//Main
func main () {
	router := gin.Default()

	//definindo Uso de memória maximo
	router.MaxMultipartMemory = 8 << 20

	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/gatinhos", uploadPhoto)

	router.Run("localhost:8080")
}


//get albuns will respond with a list of all the albuns
func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context){
	var newAlbum album

	err := c.BindJSON(&newAlbum)
	if err != nil {
		log.Println(err)
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusOK, newAlbum)
}

func getAlbumById(c *gin.Context){
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id{
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func uploadPhoto(c *gin.Context){
	//Getting the file
	
	file, err := c.FormFile("file")

	if err != nil {
		log.Println("RECIVING: ", err)
		c.String(http.StatusNotAcceptable, "File is not being acepted\n")
		return
	}

	log.Println(file.Filename)

	err = c.SaveUploadedFile(file, "./photos/")
	if err != nil {
		log.Println("SAVING: ",err)
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}