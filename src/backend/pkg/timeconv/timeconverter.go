package timeconv

import (
	"strconv"
	"strings"
	"time"
)

func IsDateValid(day, month, year int) bool {
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

func MonthNameToInt(s string) int {
	monthTable := map[string]int{
		"jan": 1, "january": 1, "januari": 1,
		"feb": 2, "february": 2, "februari": 2,
		"mar": 3, "march": 3, "faret": 3,
		"apr": 4, "april": 4,
		"may": 5, "mei": 5,
		"jun": 6, "june": 6, "juni": 6,
		"jul": 7, "july": 7, "juli": 7,
		"aug": 8, "august": 8, "agustus": 8,
		"sep": 9, "sept": 9, "september": 9,
		"oct": 10, "okt": 10, "october": 10, "oktober": 10,
		"nov": 11, "november": 11,
		"dec": 12, "des": 12, "december": 12, "desember": 12,
	}

	return monthTable[strings.ToLower(s)]
}

func DateToString(year int, month time.Month, day int) string {
	return strconv.Itoa(day) + " " + month.String() + " " + strconv.Itoa(year)
}

func DateToYYYYMMDD(year int, month time.Month, day int) string {
	return strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)
}
