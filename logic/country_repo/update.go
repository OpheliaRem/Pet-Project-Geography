package country_repo

import (
	"backPet0/db"
	"backPet0/logger"
	"backPet0/models"
	"strconv"
)

func Update(id int, country models.Country) error {
	current, err := GetById(id)
	if err != nil {
		return err
	}

	if models.CompareCountries(current, country) {
		return nil
	}

	var area, population string
	if country.Area == nil {
		area = "null"
	} else {
		area = strconv.FormatFloat(*country.Area, 'f', 2, 64)
	}

	if country.Population == nil {
		population = "null"
	} else {
		population = strconv.Itoa(*country.Population)
	}

	query := "update countries set name='" + country.Name + "', area=" + area +
		", population=" + population + " where id=" + strconv.Itoa(id)
	logger.Log(query)

	_, err = db.Db.Exec(query)
	return err
}
