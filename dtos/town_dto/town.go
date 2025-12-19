package town_dto

import (
	"encoding/json"
	"io"
)

type TownDto struct {
	Name             string
	IsCountryCapital bool
	CountryName      *string
}

func NewFromJSON(r io.Reader) (TownDto, error) {
	res := TownDto{}
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&res)
	if err != nil {
		return TownDto{}, err
	}

	return res, nil
}
