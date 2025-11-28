package model

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"` // 返さない
	Name     string `json:"name"`
}
