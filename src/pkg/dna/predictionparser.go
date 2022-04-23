package dna

import (
	"regexp"
	"strconv"
	"time"
)

func ParsePrediction(text string) (bool, time.Time, string) {
	pattern := `^(?P<date>(?P<day>\d{1,2})[\s\-\/](?P<month>\d{1,2}|[a-zA-Z]+)[\s\-\/](?P<year>\d{4}))?\s*(?P<disease>\w+)?$`
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
			month = monthNameToInt(result["month"])
		}
		year, _ := strconv.Atoi(result["year"])

		if !isDateValid(day, month, year) {
			return false, time.Time{}, ""
		}

		t = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	}

	return true, t, disease

}

func isDateValid(day, month, year int) bool {
	if day <= 0 || day > 31 || month <= 0 || month > 12 || year < 0 {
		return false
	}

	if month == 2 {
		if year%100 == 0 && year%400 != 0 {
			return day <= 28
		} else if year%4 == 0 {
			return day <= 29
		} else {
			return day <= 28
		}
	} else if month == 4 || month == 6 || month == 9 || month == 11 {
		return day <= 30
	} else {
		return true
	}
}

func monthNameToInt(s string) int {
	monthTable := map[string]int{
		"Jan": 1, "January": 1, "Januari": 1,
		"Feb": 2, "February": 2, "Februari": 2,
		"Mar": 3, "March": 3, "Maret": 3,
		"Apr": 4, "April": 4,
		"May": 5, "Mei": 5,
		"Jun": 6, "June": 6, "Juni": 6,
		"Jul": 7, "July": 7, "Juli": 7,
		"Aug": 8, "August": 8, "Agustus": 8,
		"Sep": 9, "Sept": 9, "September": 9,
		"Oct": 10, "Okt": 10, "October": 10, "Oktober": 10,
		"Nov": 11, "November": 11,
		"Dec": 12, "Des": 12, "December": 12, "Desember": 12,
	}

	return monthTable[s]
}
