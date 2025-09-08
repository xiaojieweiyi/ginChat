package main

import (
	"ginchat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:12345678.@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connnect database")
	}
	db.AutoMigrate(&models.Message{})
	db.AutoMigrate(&models.GroupBasic{})
	db.AutoMigrate(&models.Contact{})
	//user := &models.UserBasic{}
	//user.Name = "项目1"
	//db.Create(user)

	//fmt.Println(db.First(user, 1))
	//db.Model(user).Update("PassWord", "1234")

}
