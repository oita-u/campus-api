package model

type Profile struct {
	Name           string `db:"name" json:"name"`
	StudentID      string `db:"student_id" json:"studentId"`
	Email          string `db:"email" json:"email"`
	Phone          string `db:"phone" json:"phone"`
	Department     string `db:"department" json:"department"`
	Grade          string `db:"grade" json:"grade"`
	EnrollmentYear string `db:"enrollment_year" json:"enrollmentYear"`
	Birthday       string `db:"birthday" json:"birthday"`
}
