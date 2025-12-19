package town_service

import (
	"backPet0/dtos/town_dto"
	"backPet0/logic/country_repo"
	"backPet0/logic/town_repo"
	"backPet0/models"
)

func modelToDto(town models.Town) (town_dto.TownDto, error) {
	var countryName *string
	if town.CountryId != nil {
		var country models.Country
		country, err := country_repo.GetById(*town.CountryId)
		if err != nil {
			return town_dto.TownDto{}, err
		}
		countryName = &country.Name
	} else {
		countryName = nil
	}

	return town_dto.TownDto{
		Name:             town.Name,
		IsCountryCapital: town.IsCountryCapital,
		CountryName:      countryName,
	}, nil
}

func GetAll() ([]town_dto.TownDto, error) {
	towns, err := town_repo.GetAll()
	if err != nil {
		return nil, err
	}

	res := make([]town_dto.TownDto, len(towns))
	for i := range towns {
		res[i], err = modelToDto(towns[i])
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func GetById(id int) (town_dto.TownDto, error) {
	town, err := town_repo.GetById(id)
	if err != nil {
		return town_dto.TownDto{}, err
	}

	dto, err := modelToDto(town)
	if err != nil {
		return town_dto.TownDto{}, err
	}

	return dto, nil
}
