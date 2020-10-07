package jubilee_test

import (
	"testing"
	"time"

	"github.com/stillwondering/jubilee"
)

func TestProcessorWithOnlyAgeFilter(t *testing.T) {
	year := 2020
	minAge := 60
	filter := jubilee.AgeFilter(year, minAge)
	processor := jubilee.NewProcessor([]jubilee.FilterFunc{filter})

	tables := []struct {
		records           []jubilee.BirthdayRecord
		expectedBirthdays []string
	}{
		{
			[]jubilee.BirthdayRecord{recordFromStrings("Tester", "1961-01-01"), recordFromStrings("Tester", "1960-01-01"), recordFromStrings("Tester", "1960-12-31")},
			[]string{"1960-01-01", "1960-12-31"},
		},
	}

	for _, table := range tables {
		filtered := processor.Process(table.records)

		var birthdays []string
		for _, record := range filtered {
			birthdays = append(birthdays, record.DateOfBirth.Format(jubilee.DateFormat))
		}

		if len(birthdays) != len(table.expectedBirthdays) {
			t.Errorf("for year %d, minimum age %d got: %v, expected: %v", year, minAge, birthdays, table.expectedBirthdays)

		}
	}
}

func TestAgeInFunction(t *testing.T) {
	bday := "1993-04-12"
	record := recordFromStrings("Tester", bday)

	tables := []struct {
		year        int
		expectedAge int
	}{
		{2020, 27},
		{1901, 0},
		{1993, 0},
	}

	for _, table := range tables {
		age := record.AgeIn(table.year)

		if age != table.expectedAge {
			t.Errorf("for date of birth %s, and year %d got: %d, expected: %d", bday, table.year, age, table.expectedAge)
		}
	}
}

func recordFromStrings(name string, bday string) jubilee.BirthdayRecord {
	dateOfBirth, _ := time.Parse(jubilee.DateFormat, bday)

	return jubilee.BirthdayRecord{
		Name:        name,
		DateOfBirth: dateOfBirth,
	}
}
