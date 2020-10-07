package csv

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/stillwondering/jubilee"
)

const csvSeparator = ","

func NewProvider(reader io.Reader) *Provider {
	return &Provider{
		input: reader,
	}
}

type Provider struct {
	input io.Reader
}

func (p *Provider) Fetch() ([]jubilee.BirthdayRecord, error) {
	reader := bufio.NewReader(p.input)
	var records []jubilee.BirthdayRecord

	for {
		input, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}

		line := strings.TrimSpace(string(input))

		// Disregard empty lines or comments
		if "" == line || '#' == line[0] {
			continue
		}

		record, err := recordFromLine(line)
		if err != nil {
			return records, fmt.Errorf("corrupted file: %v", err)
		}

		records = append(records, record)
	}

	return records, nil
}

func recordFromLine(line string) (jubilee.BirthdayRecord, error) {
	fields := strings.Split(line, csvSeparator)

	if len(fields) < 2 {
		return jubilee.BirthdayRecord{}, fmt.Errorf("invalid line '%s'", line)
	}

	name := fields[0]
	dateString := fields[1]

	bday, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return jubilee.BirthdayRecord{}, fmt.Errorf("invalid line '%s': cannot parse date", line)
	}

	return jubilee.BirthdayRecord{
		Name:        name,
		DateOfBirth: bday,
	}, nil
}
