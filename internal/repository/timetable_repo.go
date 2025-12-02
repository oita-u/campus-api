package repository

import (
	"github.com/oita-u/campus-api/internal/db"
	"github.com/oita-u/campus-api/internal/model"
)

type TimetableRepository struct{}

func NewTimetableRepository() *TimetableRepository {
	return &TimetableRepository{}
}

func (r *TimetableRepository) GetTimetableByUserID(userID string, year int, semester string) ([]model.TimetableItem, error) {
	items := []model.TimetableItem{}

	query := `
		SELECT
			c.id AS course_id,
			c.course_name,
			c.instructor_name,
			c.course_type,
			c.day_of_week,
			c.period,
			r.status
		FROM registrations r
		JOIN courses c ON r.course_id = c.id
		WHERE
			r.user_id = $1
			AND r.status = 'registered'
			AND c.academic_year = $2
			AND c.semester = $3
		ORDER BY c.day_of_week, c.period, c.course_name
	`

	if err := db.Conn.Select(&items, query, userID, year, semester); err != nil {
		return nil, err
	}

	return items, nil
}
