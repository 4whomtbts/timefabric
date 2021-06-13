package model

type User struct {
	Id int32 `json:"id" db:"id"`
	LoginId string `json:"login_id" db:"login_id"`
	Password string `json:"password" db:"password"`
	UserName string `json:"user_name" db:"user_name"`
	Phone string `json:"phone" db:"phone"`
	Email string `json:"email" db:"email"`
}