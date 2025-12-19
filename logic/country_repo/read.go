package country_repo

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
	return "There is no item with id=" + strconv.Itoa(err.id)
}

func get(query string) ([]models.Country, error) {
	rows, err := db.Db.Query(query)
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {

		}
	}(rows)
	if err != nil {
		return nil, err
	}

	var res []models.Country
	for rows.Next() {
		country := models.Country{}
		err = rows.Scan(&country.Id, &country.Name, &country.Area, &country.Population)
		if err != nil {
			return nil, err
		}
		res = append(res, country)
	}

	return res, nil
}

func GetById(id int) (models.Country, error) {
	query := "select * from countries where id = " + strconv.Itoa(id) + " limit 1"
	logger.Log(query)

	countries, err := get(query)
	if err != nil {
		return models.Country{}, err
	}

	if len(countries) == 0 {
		return models.Country{}, NotFoundError{id: id}
	}

	return countries[0], nil
}

func GetAll() ([]models.Country, error) {
	query := "select * from countries"
	logger.Log(query)

	return get(query)
}
