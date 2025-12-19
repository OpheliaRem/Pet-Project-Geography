package town_service

import (
	"backPet0/dtos/town_dto"
	"backPet0/logic/country_repo"
	"backPet0/logic/town_repo"
	"backPet0/models"
)

func modelToDtoForRead(town models.Town) (town_dto.TownDtoForRead, error) {
	var countryName *string
	if town.CountryId != nil {
		var country models.Country
		country, err := country_repo.GetById(*town.CountryId)
		if err != nil {
			return town_dto.TownDtoForRead{}, err
		}
		countryName = &country.Name
	} else {
		countryName = nil
	}

	return town_dto.TownDtoForRead{
		Name:             town.Name,
		IsCountryCapital: town.IsCountryCapital,
		CountryName:      countryName,
	}, nil
}

func GetAll() ([]town_dto.TownDtoForRead, error) {
	towns, err := town_repo.GetAll()
	if err != nil {
		return nil, err
	}

	res := make([]town_dto.TownDtoForRead, len(towns))
	for i := range towns {
		res[i], err = modelToDtoForRead(towns[i])
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func GetById(id int) (town_dto.TownDtoForRead, error) {
	town, err := town_repo.GetById(id)
	if err != nil {
		return town_dto.TownDtoForRead{}, err
	}

	dto, err := modelToDtoForRead(town)
	if err != nil {
		return town_dto.TownDtoForRead{}, err
	}

	return dto, nil
}

func Save(dto town_dto.TownDtoForCreate) error {
	return town_repo.Save(models.Town{
		Id:               0,
		Name:             dto.Name,
		IsCountryCapital: dto.IsCountryCapital,
		CountryId:        dto.CountryId,
	})
}
