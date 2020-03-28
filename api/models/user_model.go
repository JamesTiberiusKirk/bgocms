package models

// User for user table
type User struct {
	ID    string `json:"id"`
	Uname string `json:"uname"`
	Pass  string `json:"pass"`
}

//Users for multiple users
type UsersResponce struct {
	Users []User `json:"users"`
  Total int    `json:"total"`
}
