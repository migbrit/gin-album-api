package album

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

type album struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func InitializeDb(database *sql.DB) {
	db = database
}

func PostAlbums(ginC *gin.Context) {
	var newAlbums []album

	if err := ginC.BindJSON(&newAlbums); err != nil {
		ginC.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON data"})
		return
	}

	for i, newAlbum := range newAlbums {
		result, err := db.Exec("INSERT INTO Albums (Title, Artist, Price) VALUES (?, ?, ?)", newAlbum.Title, newAlbum.Artist, newAlbum.Price)
		if err != nil {
			ginC.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error inserting new album into the database"})
			return
		}

		lastInsertID, err := result.LastInsertId()
		if err != nil {
			ginC.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error getting last insert ID"})
			return
		}

		newAlbums[i].ID = lastInsertID
	}

	ginC.IndentedJSON(http.StatusCreated, newAlbums)
}

func GetAlbumByID(ginC *gin.Context) {
	id := ginC.Query("id")

	if id == "" {
		ginC.IndentedJSON(http.StatusBadRequest, gin.H{"message:": "Empty id on url"})
		return
	}

	println("id: ", id)

	rows, err := db.Query("SELECT Id, Title, Artist, Price from Albums WHERE Id = ?", id)
	if err != nil {
		ginC.IndentedJSON(http.StatusNotFound, gin.H{"message:": "Album not found with the specified id."})
	}

	defer rows.Close()

	if !rows.Next() {
		ginC.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found with the specified id."})
		return
	}

	var a album
	if err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price); err != nil {
		ginC.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error scanning albums on database"})
		return
	}

	ginC.IndentedJSON(http.StatusOK, a)

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
