package repository

import (
	"github.com/oita-u/campus-api/internal/db"
	"github.com/oita-u/campus-api/internal/model"
)

type ProfileRepository struct{}

func NewProfileRepository() *ProfileRepository {
	return &ProfileRepository{}
}

func (r *ProfileRepository) GetByUserID(userID string) (*model.Profile, error) {
	var profile model.Profile
	query := `
		SELECT 
			u.name as name,
			COALESCE(s.student_number, '') as student_id,
			COALESCE(u.email, '') as email,
			'' as phone,
			COALESCE(d.name, s.department, '') as department,
			COALESCE(u.grade, s.grade, '') as grade,
			'' as enrollment_year,
			'' as birthday
		FROM users u
		LEFT JOIN students s ON u.user_code = s.student_number
		LEFT JOIN departments d ON u.department_id = d.id
		WHERE u.id::text = $1 OR u.user_code = $1
	`
	err := db.Conn.Get(&profile, query, userID)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}
