package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("数据区连接失败，请检查配置文件")
	}
	LoadServer(file)
	LoadDatabase(file)
}
func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
}
func LoadDatabase(file *ini.File) {
	Db = file.Section("databse").Key("Db").MustString("mysql")
	DbUser = file.Section("databse").Key("DbUser").MustString("root")
	DbPassword = file.Section("databse").Key("DbPassword").MustString("wyc1750346")
	DbHost = file.Section("databse").Key("DbHost").MustString("localhost")
	DbPort = file.Section("databse").Key("DbPort").MustString("3306")
	DbName = file.Section("databse").Key("DbName").MustString("ginblog")
}
