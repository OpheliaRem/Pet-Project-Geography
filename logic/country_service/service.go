package country_service

import (
	"backPet0/dtos/country_dto"
	country_repos2 "backPet0/logic/country_repo"
)

func Save(dto country_dto.CountryDTO) error {
	return country_repos2.Save(dto.ToModel())
}

func GetById(id int) (country_dto.CountryDTO, error) {
	country, err := country_repos2.GetById(id)
	return country_dto.NewFromModel(country), err
}

func GetAll() ([]country_dto.CountryDTO, error) {
	countries, err := country_repos2.GetAll()
	res := make([]country_dto.CountryDTO, len(countries))
	for i := range countries {
		res[i] = country_dto.NewFromModel(countries[i])
	}

	return res, err
}

func Remove(id int) error {
	return country_repos2.Remove(id)
}

func Update(id int, dto country_dto.CountryDTO) error {
	return country_repos2.Update(id, dto.ToModel())
}
