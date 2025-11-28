package repository

import (
	"github.com/oita-u/campus-api/internal/db"
	"github.com/oita-u/campus-api/internal/model"
)

type StudentRepository struct{}

func (r *StudentRepository) GetByID(id int) (*model.Student, error) {
	var s model.Student
	err := db.Conn.Get(&s, `SELECT * FROM students WHERE student_id = $1`, id)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *StudentRepository) GetByStudentNumber(num string) (*model.Student, error) {
	var s model.Student
	err := db.Conn.Get(&s, `SELECT * FROM students WHERE student_number = $1`, num)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *StudentRepository) List() ([]model.Student, error) {
	list := []model.Student{}
	err := db.Conn.Select(&list, `SELECT * FROM students ORDER BY student_id`)
	return list, err
}
