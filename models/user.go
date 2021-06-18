package models

type User struct {
	Id         uint
	Name       string
	Email      string `gorm:"unique"`
	Password   []byte
	MediaLinks *MediaLinks `json:"mediaLinks"`
}

type MediaLinks struct {
	Facebook  string `json:"facebook"`
	Twitter   string `json:"twitter"`
	Instagram string `json:"instagram"`
	Telegram  string `json:"telegram"`
}
