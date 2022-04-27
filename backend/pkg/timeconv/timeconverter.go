package timeconv

import (
	"strconv"
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

func DateToString(year int, month time.Month, day int) string {
	return strconv.Itoa(day) + " " + month.String() + " " + strconv.Itoa(year)
}
