package country_repo

import (
	"backPet0/db"
	"backPet0/logger"
	"backPet0/models"
	"strconv"
)

func buildInsert(country models.Country) string {
	cols := "insert into countries (name"
	vals := " values ('" + country.Name + "'"
	if country.Area != nil {
		cols += ", area"
		vals += ", " + strconv.FormatFloat(*country.Area, 'f', 2, 64)
	}
	if country.Population != nil {
		cols += ", population"
		vals += ", " + strconv.Itoa(*country.Population)
	}

	cols += ")"
	vals += ")"

	return cols + vals
}

func Save(country models.Country) error {
	query := buildInsert(country)

	logger.Log(query)

	_, err := db.Db.Exec(query)

	return err
}
