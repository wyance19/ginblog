package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Category   Category `gorm:"foreignKey:CategoryID"`
	Cid        int      `gorm:"type:int;not null" json:"cid"`
	Title      string   `gorm:"type:varchar(100);not null" json:"title"`
	Desc       string   `gorm:"type:varchar(200);not null" json:"desc"`
	Content    string   `gorm:"type:longtext;not null" json:"content"`
	Img        string   `gorm:"type:varchar(100);not null" json:"img"`
	CategoryID int
}
