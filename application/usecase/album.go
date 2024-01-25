package usecase

import (
	"database/sql"
	"gin-simple-api/application/domain"
	"gin-simple-api/application/repository"
)

func CreateUsers(newAlbums []domain.Album, Db *sql.DB) ([]domain.Album, error) {
	for i, newAlbum := range newAlbums {
		result, err := repository.CreateAlbum(newAlbum, Db)
		if err != nil {
			return nil, err
		}

		lastInsertID, err := result.LastInsertId()
		if err != nil {
			return nil, err
		}

		newAlbums[i].ID = lastInsertID
	}

	return newAlbums, nil
}
