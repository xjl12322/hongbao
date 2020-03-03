package models

type User struct {
	UserId int64 `json:"user_id" db:"id"`
	UserName string `json:"user_name" db:"name"`
	Password int64 `json:"password" db:"password"`
	CreateAt int64 `json:"create_at" db:"create_time"`
	Email string `json:"email" db:"email"`
	Phone    string `json:"phone" db:"phone"`

}


