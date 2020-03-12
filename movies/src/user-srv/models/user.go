package models

type User struct {
	UserId   int64  `json:"user_id" db:"user_id"`
	Name string `json:"user_name" db:"name"`
	Password string `json:"password" db:"password"`
	CreateTime string `json:"create_at" db:"create_time"`
	Email    string `json:"email" db:"email"`
	Phone    string `json:"phone" db:"phone"`
}


