package models

import "im/database"

type Users struct {
	Id          int    `gorm:"primaryKey" json:"id"`
	Nickname    string `gorm:"nickname" json:"nickname"`
	Avatar      string `gorm:"avatar" json:"avatar"`
	PhoneNumber string `gorm:"phone_number" json:"phone_number"`
	CreateTime  int    `gorm:"create_time" json:"create_time"`
}

func (u *Users) Find(id int) *Users {
	user := &Users{}
	database.Db.Find(user, id)
	return user
}
