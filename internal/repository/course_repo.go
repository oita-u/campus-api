package repository

import (
	"fmt"

	"github.com/oita-u/campus-api/internal/db"
	"github.com/oita-u/campus-api/internal/model"
)

type CourseRepository struct{}

func NewCourseRepository() *CourseRepository {
	return &CourseRepository{}
}

func (r *CourseRepository) SearchCourses(year int, semester, keyword string) ([]model.Course, error) {
	courses := []model.Course{}

	query := `
		SELECT
			id,
			course_name,
			instructor_name,
			course_type,
			day_of_week,
			period,
			academic_year,
			semester,
			course_code,
			credits,
			department_id,
			created_at,
			updated_at
		FROM courses
		WHERE 1=1
	`

	args := []interface{}{}
	argPos := 1

	if year != 0 {
		query += fmt.Sprintf(" AND academic_year = $%d", argPos)
		args = append(args, year)
		argPos++
	}

	if semester != "" {
		query += fmt.Sprintf(" AND semester = $%d", argPos)
		args = append(args, semester)
		argPos++
	}

	if keyword != "" {
		like := "%" + keyword + "%"
		query += fmt.Sprintf(`
			AND (
				course_name ILIKE $%d OR
				instructor_name ILIKE $%d OR
				course_code ILIKE $%d
			)
		`, argPos, argPos, argPos)
		args = append(args, like)
		argPos++
	}

	query += " ORDER BY course_name"

	if err := db.Conn.Select(&courses, query, args...); err != nil {
		return nil, err
	}

	return courses, nil
}
