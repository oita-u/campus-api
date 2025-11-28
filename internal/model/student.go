package model

import "time"

type Student struct {
	StudentID     int       `db:"student_id" json:"student_id"`
	StudentNumber string    `db:"student_number" json:"student_number"`
	Name          string    `db:"name" json:"name"`
	Department    string    `db:"department" json:"department"`
	Grade         string    `db:"grade" json:"grade"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}
