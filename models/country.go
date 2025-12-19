package models

type Country struct {
	Id         int
	Name       string
	Area       *float64
	Population *int
}

func CompareCountries(a, b Country) bool {
	if a.Id != b.Id {
		return false
	}

	if a.Name != b.Name {
		return false
	}

	if (a.Area == nil) != (b.Area == nil) {
		return false
	}

	if (a.Population == nil) != (b.Population == nil) {
		return false
	}

	if a.Area != nil && b.Area != nil {
		if *a.Area != *b.Area {
			return false
		}
	}

	if a.Population != nil && b.Population != nil {
		if *a.Population != *b.Population {
			return false
		}
	}

	return true
}
