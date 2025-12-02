package model

type User struct {
	ID       string `db:"id" json:"id"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"-"` // 返さない
	Name     string `db:"name" json:"name"`
}
