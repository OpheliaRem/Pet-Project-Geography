package town_repo

import (
	"backPet0/db"
	"backPet0/logger"
	"backPet0/logic/country_repo"
	"backPet0/models"
	"strconv"
)

func Save(town models.Town) error {
	if town.CountryId != nil {
		_, err := country_repo.GetById(*town.CountryId)
		if err != nil {
			return err
		}
	}

	var countryIdStr string
	if town.CountryId != nil {
		countryIdStr = strconv.Itoa(*town.CountryId)
	} else {
		countryIdStr = "null"
	}

	cols := "insert into towns (name, is_country_capital, country_id)"
	vals := " values('" + town.Name + "', " + strconv.FormatBool(town.IsCountryCapital) +
		", " + countryIdStr + ")"

	query := cols + vals
	logger.Log(query)

	_, err := db.Db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
