package models

type User struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email" gorm:"unique"`
	Password   []byte `json:"-"`
	MediaLinks *MediaLinks
}

type MediaLinks struct {
	Facebook  string `json:"facebook"`
	Twitter   string `json:"twitter"`
	Instagram string `json:"instagram"`
	Telegram  string `json:"telegram"`
}
