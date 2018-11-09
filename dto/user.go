package dto

type User struct {
	Username    string
	Userpass    string `gorm:"column:userpassword"`
	Userpassb64 string `gorm:"column:userpasswordb64"`
}
