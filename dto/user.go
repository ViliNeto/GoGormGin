package dto

type User struct {
	Usersid          int64
	Username        string
	Userpassword    string `gorm:"default:''"`
	Userpasswordb64 string `gorm:"default:''"`
}
