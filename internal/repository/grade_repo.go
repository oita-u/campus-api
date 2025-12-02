package repository

import (
	"github.com/oita-u/campus-api/internal/db"
	"github.com/oita-u/campus-api/internal/model"
)

type GradeRepository struct{}

func NewGradeRepository() *GradeRepository {
	return &GradeRepository{}
}

// GetGradesByUserID は指定したユーザー・年度・学期の成績一覧を取得する
func (r *GradeRepository) GetGradesByUserID(userID string, year int, semester string) ([]model.GradeItem, error) {
	items := []model.GradeItem{}

	query := `
		SELECT
			c.id AS course_id,
			c.course_code,
			c.course_name,
			c.instructor_name,
			c.credits,
			c.semester,
			c.academic_year,
			r.status,
			r.grade
		FROM registrations r
		JOIN courses c ON r.course_id = c.id
		WHERE
			r.user_id = $1
			AND c.academic_year = $2
			AND c.semester = $3
			AND r.status = 'completed'
		ORDER BY c.course_code, c.course_name
	`

	if err := db.Conn.Select(&items, query, userID, year, semester); err != nil {
		return nil, err
	}

	return items, nil
}


