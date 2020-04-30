package models

import (

)

// LoginResult 登录结果结构
type LoginResult struct {
	Token string `json:"token"`
	User
}
