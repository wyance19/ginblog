package model

import (
	"ginblog/utils/errmsg"

	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

func CheckCate(name string) (code int) {
	var Cate Category
	db.Select("id").Where("name = ?", name).Find(&Cate)
	if Cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCESS
}

//新增用户
func CreateCate(data *Category) int {
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询用户列表
func GetCate(pageSize int, PageNum int) []Category {
	var data []Category
	//分页
	err = db.Limit(pageSize).Offset((PageNum - 1) * pageSize).Find(&data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return data
}

//删除用户
func DeleteCate(id int) int {
	var category Category
	err = db.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//编辑用户信息
func EditCate(id int, data *Category) int {
	var cata []Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(cata).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
