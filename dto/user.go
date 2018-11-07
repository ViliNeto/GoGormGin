package dto

type User struct {
	Username        string
	Userpassword    string `gorm:"default:''"`
	Userpasswordb64 string `gorm:"default:''"`
}
