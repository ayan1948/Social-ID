package models

type User struct {
	Id       uint
	Name     string
	Email    string `gorm:"unique"`
	Password []byte
	Social   Social `gorm:"foreignKey:UserID"`
}

type Social struct {
	Facebook  string
	LinkedIn  string
	Instagram string
	WhatsApp  []byte
	UserID    uint
}
