package repository

import (
	"github.com/oita-u/campus-api/internal/db"
	"github.com/oita-u/campus-api/internal/model"
)

type StudentStatusChangeRepository struct{}

func (r *StudentStatusChangeRepository) GetByStudentNumber(studentNumber string) ([]model.StudentStatusChange, error) {
	query := `
        SELECT id, student_id, change_type, reason, approval_date, 
               scheduled_start_date, scheduled_end_date, actual_start_date, actual_end_date
        FROM student_status_changes
        WHERE student_id = $1
    `

	changes := []model.StudentStatusChange{}

	err := db.Conn.Select(&changes, query, studentNumber)
	return changes, err
}
