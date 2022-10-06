package model

import (
	"encoding/base64"
	"ginblog/utils/errmsg"
	"log"

	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

//查询用户是否存在
func CheckUser(data string) (code int) {
	var users User
	db.Select("id").Where("username = ?", data).Find(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

//新增用户
func CreateUser(data *User) int {
	data.Password = ScryptPw(data.Password)
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询用户列表
func GetUsers(pageSize int, PageNum int) []User {
	var users []User
	//分页
	err = db.Limit(pageSize).Offset((PageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

//密码加密
func ScryptPw(password string) string {
	const Keylen = 4
	var salt = make([]byte, 8)
	salt = []byte{2, 4, 15, 74, 88, 34, 234, 1}
	HashPw, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, Keylen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}

//删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//编辑用户信息
func EditUser(id int, data *User) int {
	var user []User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
