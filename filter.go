package jubilee

// AgeFilter creates a FilterFunc that decides whether or not a person turns at
// least the given minimum age in the given year.
func AgeFilter(year int, minAge int) FilterFunc {
	return func(record BirthdayRecord) bool {
		return record.AgeIn(year) >= minAge
	}
}

// AnniversariesFilter creates a FilterFunc
func AnniversariesFilter(year int, anniversaries []int) FilterFunc {
	return func(record BirthdayRecord) bool {
		ageToBe := record.AgeIn(year)
		for _, mod := range anniversaries {
			if ageToBe%mod == 0 {
				return true
			}
		}

		return false
	}
}
