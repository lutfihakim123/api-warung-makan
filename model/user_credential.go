package model

type UserCredential struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"userPassword"`
	Email    string
}
