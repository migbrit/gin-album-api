package repository

import (
	"database/sql"
	"gin-simple-api/application/domain"
)

func GetAlbumByID(id string, db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("SELECT Id, Title, Artist, Price from Albums WHERE Id = ?", id)

	return rows, err
}

func GetAlbums(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("SELECT Id, Title, Artist, Price from Albums")

	return rows, err
}

func CreateAlbum(newAlbum domain.Album, db *sql.DB) (sql.Result, error) {
	result, err := db.Exec("INSERT INTO Albums (Title, Artist, Price) VALUES (?, ?, ?)", newAlbum.Title, newAlbum.Artist, newAlbum.Price)

	return result, err
}
