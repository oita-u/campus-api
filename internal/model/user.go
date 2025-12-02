package model

type User struct {
	ID       string `db:"id" json:"id"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"-"`
	Name     string `db:"name" json:"name"`
}
