package album

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func InitializeDb(database *sql.DB) {
	db = database
}

func GetAlbums(ginC *gin.Context) {
	rows, err := db.Query("SELECT Id, Title, Artist, Price from Albums")
	if err != nil {
		ginC.IndentedJSON(http.StatusInternalServerError, gin.H{"message:": "Error trying to get albums data"})
		return
	}
	defer rows.Close()

	var fetchedAlbums []album

	for rows.Next() {
		var a album
		if err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price); err != nil {
			ginC.IndentedJSON(http.StatusInternalServerError, gin.H{"message:": "Error scanning albums on database"})
			return
		}
		fetchedAlbums = append(fetchedAlbums, a)
	}

	ginC.IndentedJSON(http.StatusOK, fetchedAlbums)
}

func PostAlbums(ginC *gin.Context) {
	var newAlbum album

	if err := ginC.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	ginC.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumByID(ginC *gin.Context) {
	id := ginC.Query("id")

	if id == "" {
		ginC.IndentedJSON(http.StatusBadRequest, gin.H{"message:": "empty id on url"})
		return
	}

	for _, a := range albums {
		if a.ID == id {
			ginC.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	ginC.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
