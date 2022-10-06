package model

import (
	"ginblog/utils/errmsg"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Category Category `gorm:"foreignKey:cid"`
	Cid      int      `gorm:"type:int;not null" json:"cid"`
	Title    string   `gorm:"type:varchar(100);not null" json:"title"`
	Desc     string   `gorm:"type:varchar(200);not null" json:"desc"`
	Content  string   `gorm:"type:longtext;not null" json:"content"`
	Img      string   `gorm:"type:varchar(100);not null" json:"img"`
}

//新增文章
func CreateArt(data *Article) int {
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询分类下的所有文章

// 查询单个文章

// 查询文章列表
func GetArt(pageSize int, PageNum int) ([]Article, int) {
	var data []Article
	//分页
	err = db.Preload("Category").Limit(pageSize).Offset((PageNum - 1) * pageSize).Find(&data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return data, errmsg.SUCCESS
}

//删除文章
func DeleteArt(id int) int {
	var article Article
	err = db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//编辑文章
func EditArt(id int, data *Article) int {
	var Art []Article
	var maps = make(map[string]interface{})
	maps["cid"] = data.Cid
	maps["title"] = data.Title
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err = db.Model(Art).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
