package country_dto

import (
	"backPet0/models"
	"encoding/json"
	"io"
)

type CountryDTO struct {
	Name       string
	Area       *float64
	Population *int
}

func NewFromJSON(r io.Reader) (CountryDTO, error) {
	d := json.NewDecoder(r)

	res := CountryDTO{}

	err := d.Decode(&res)
	if err != nil {
		return CountryDTO{}, err
	}

	return res, nil
}

func NewFromModel(m models.Country) CountryDTO {
	return CountryDTO{
		Name:       m.Name,
		Area:       m.Area,
		Population: m.Population,
	}
}

func (dto CountryDTO) ToModel() models.Country {
	return models.Country{
		Id:         0,
		Name:       dto.Name,
		Area:       dto.Area,
		Population: dto.Population,
	}
}
