package model

type User struct {
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
	Age      int8   `json:"age"`
}
