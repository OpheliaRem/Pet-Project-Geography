package town_repo

import (
	"backPet0/db"
	"backPet0/logger"
	"backPet0/models"
	"database/sql"
	"strconv"
)

type NotFoundError struct {
	id int
}

func (err NotFoundError) Error() string {
	return "Town with id=" + strconv.Itoa(err.id) + " was not found"
}

func get(query string) ([]models.Town, error) {
	rows, err := db.Db.Query(query)
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)
	if err != nil {
		return nil, err
	}

	var res []models.Town
	for rows.Next() {
		town := models.Town{}
		err = rows.Scan(&town.Id, &town.Name, &town.IsCountryCapital, &town.CountryId)
		if err != nil {
			return nil, err
		}
		res = append(res, town)
	}

	return res, err
}

func GetAll() ([]models.Town, error) {
	query := "select * from towns"
	logger.Log(query)

	return get(query)
}

func GetById(id int) (models.Town, error) {
	query := "select * from towns where id=" + strconv.Itoa(id)
	logger.Log(query)

	towns, err := get(query)

	if err != nil {
		return models.Town{}, err
	}

	if len(towns) == 0 {
		return models.Town{}, NotFoundError{id: id}
	}

	return towns[0], nil
}
