package dna

import (
	"regexp"
	"strconv"
	"time"

	"github.com/weslygio/Tubes3_13520071/src/backend/pkg/timeconv"
)

func ParsePrediction(text string) (bool, time.Time, string) {
	pattern := `^(?P<date>(?P<day>\d{1,2})[\s\-\/](?P<month>\d{1,2}|[a-zA-Z]+)[\s\-\/](?P<year>\d{4}))?\s*(?P<disease>.+)?$`
	re := regexp.MustCompile(pattern)
	result := make(map[string]string)
	match := re.FindStringSubmatch(text)

	if len(match) == 0 {
		return false, time.Time{}, ""
	}

	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	t := time.Time{}
	disease := result["disease"]

	if result["date"] != "" {
		day, _ := strconv.Atoi(result["day"])
		month, err := strconv.Atoi(result["month"])
		if err != nil {
			month = timeconv.MonthNameToInt(result["month"])
		}
		year, _ := strconv.Atoi(result["year"])

		if !timeconv.IsDateValid(day, month, year) {
			return false, time.Time{}, ""
		}

		t = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	}

	return true, t, disease

}
