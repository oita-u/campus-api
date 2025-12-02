package service

import "time"

func resolveYearSemester(year int, semester string) (int, string) {
	if year == 0 {
		year = time.Now().Year()
	}

	if semester == "" {
		semester = "1st"
	}

	return year, semester
}
