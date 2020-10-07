package jubilee

import (
	"time"
)

const (
	// TenYearsAnniversary can be used to say that we only treat birthdays like
	// 60, 70, ...
	TenYearsAnniversary int = 10
	// FiveYearsAnniversary can be used to say that we only treat birthdays like
	// 65, 75, ...
	//
	// Of course this includes all ten years anniversaries.
	FiveYearsAnniversary int = 5
	// DateFormat is the default date format we'll use in this application.
	DateFormat string = "2006-01-02"
)

// A BirthdayRecord stores information about a person's birthday.
type BirthdayRecord struct {
	Name        string
	DateOfBirth time.Time
}

// AgeIn calculates the age a person will turn in the given year.
func (b BirthdayRecord) AgeIn(year int) int {
	age := year - b.DateOfBirth.Year()
	if age < 0 {
		return 0
	}

	return age
}

// NewProcessor creates a new instance of Processor.
func NewProcessor(filters []FilterFunc) *Processor {
	return &Processor{
		filters: filters,
	}
}

// The Processor is responsible for deciding whether or not a birthday should be
// included in the output.
type Processor struct {
	filters []FilterFunc
}

// Process takes a collection of birthday recrods and filters based on the
// configuration of the Processor.
func (p *Processor) Process(input []BirthdayRecord) []BirthdayRecord {
	var filtered []BirthdayRecord

	for _, record := range input {
		belongs := true
		for _, filter := range p.filters {
			if !filter(record) {
				belongs = false
			}
		}

		if !belongs {
			continue
		}

		filtered = append(filtered, record)
	}

	return filtered
}

// A FilterFunc is used to decide whether or not a birthday meets a certain
// criteria.
type FilterFunc func(BirthdayRecord) bool
