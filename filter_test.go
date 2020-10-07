package jubilee_test

import (
	"testing"

	"github.com/stillwondering/jubilee"
)

func TestAgeFilter(t *testing.T) {
	year := 2020
	minAge := 60
	filter := jubilee.AgeFilter(year, minAge)

	tables := []struct {
		record   jubilee.BirthdayRecord
		expected bool
	}{
		{recordFromStrings("Tester", "2020-01-01"), false},
		{recordFromStrings("Tester", "1960-01-01"), true},
		{recordFromStrings("Tester", "1960-12-31"), true},
	}

	for _, table := range tables {
		result := filter(table.record)

		if result != table.expected {
			t.Errorf("for year %d, minimum age %d, and date of birth %s got: %v, expected: %v", year, minAge, table.record.DateOfBirth.Format(jubilee.DateFormat), result, table.expected)
		}
	}
}

func TestAnniversaryFilter(t *testing.T) {
	year := 2020
	anniversaries := []int{jubilee.TenYearsAnniversary, jubilee.FiveYearsAnniversary}
	filter := jubilee.AnniversariesFilter(year, anniversaries)

	tables := []struct {
		record   jubilee.BirthdayRecord
		expected bool
	}{
		// yeah well, 0 is still an anniversary
		{recordFromStrings("Tester", "2020-01-01"), true},
		{recordFromStrings("Tester", "1960-01-01"), true},
		{recordFromStrings("Tester", "1960-12-31"), true},
		{recordFromStrings("Tester", "1965-12-31"), true},
		{recordFromStrings("Tester", "1994-12-31"), false},
	}

	for _, table := range tables {
		result := filter(table.record)

		if result != table.expected {
			t.Errorf("for year %d, anniversaries %v, and date of birth %s got: %v, expected: %v", year, anniversaries, table.record.DateOfBirth.Format(jubilee.DateFormat), result, table.expected)
		}
	}
}
