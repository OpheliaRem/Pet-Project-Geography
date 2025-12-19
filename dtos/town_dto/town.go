package town_dto

type TownDtoForRead struct {
	Name             string
	IsCountryCapital bool
	CountryName      *string
}

type TownDtoForCreate struct {
	Name             string
	IsCountryCapital bool
	CountryId        *int
}
