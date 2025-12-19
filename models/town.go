package models

type Town struct {
	Id               int
	Name             string
	IsCountryCapital bool
	CountryId        *int
}
