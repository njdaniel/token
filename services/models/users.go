package models

// User login credentials
type User struct {
	UUID     string `json:"uuid" form:"-"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
