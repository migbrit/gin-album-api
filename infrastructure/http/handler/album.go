package handler

import (
	"database/sql"
	"fmt"
	"gin-simple-api/application/domain"
	"gin-simple-api/application/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Db *sql.DB

func InitializeDb(database *sql.DB) {
	Db = database
}

func PostAlbums(ginC *gin.Context) {
	var newAlbums []domain.Album

	if err := ginC.BindJSON(&newAlbums); err != nil {
		ginC.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON data"})
		return
	}

	for i, newAlbum := range newAlbums {
		result, err := repository.CreateAlbum(newAlbum, Db)
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

	rows, err := repository.GetAlbumByID(id, Db)

	fmt.Println("err: ", err)

	if err != nil {
		ginC.IndentedJSON(http.StatusNotFound, gin.H{"message:": "Album not found with the specified id."})
		return
	}

	defer rows.Close()

	if !rows.Next() {
		ginC.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found with the specified id."})
		return
	}

	var a domain.Album
	if err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price); err != nil {
		ginC.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error scanning albums on database"})
		return
	}

	ginC.IndentedJSON(http.StatusOK, a)

}

func GetAlbums(ginC *gin.Context) {
	rows, err := repository.GetAlbums(Db)
	if err != nil {
		ginC.IndentedJSON(http.StatusInternalServerError, gin.H{"message:": "Error trying to get albums data"})
		return
	}
	defer rows.Close()

	var fetchedAlbums []domain.Album

	for rows.Next() {
		var a domain.Album
		if err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price); err != nil {
			ginC.IndentedJSON(http.StatusInternalServerError, gin.H{"message:": "Error scanning albums on database"})
			return
		}
		fetchedAlbums = append(fetchedAlbums, a)
	}

	ginC.IndentedJSON(http.StatusOK, fetchedAlbums)
}
