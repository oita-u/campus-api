package model

type User struct {
	ID           string `db:"id" json:"id"`
	Email        string `db:"email" json:"email"`
	Password     string `db:"password" json:"-"`
	Name         string `db:"name" json:"name"`
	UserCode     string `db:"user_code" json:"user_code"`
	Grade        string `db:"grade" json:"grade"`
	DepartmentID int64  `db:"department_id" json:"department_id"`
}
