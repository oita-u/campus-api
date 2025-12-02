package model

type User struct {
	ID           int64  `db:"id" json:"id"`
	UserCode     string `db:"user_code" json:"user_code"`
	Email        string `db:"email" json:"email"`
	Password     string `db:"password" json:"-"`
	Name         string `db:"name" json:"name"`
	Grade        string `db:"grade" json:"grade"`
	DepartmentID int64  `db:"department_id" json:"department_id"`
}
