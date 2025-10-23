package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"uniqueIndex;size:50;not null"`
	Password string `json:"password" gorm:"size:125;not null"`
	AvatarId string `json:"avatar_id" gorm:"size:32;default:''"`
}

func FindUserByUsername(username string) (User, error) {
	var u User
	err := DB.Where("username=?", username).First(&u).Error
	return u, err
}

func CreateUser(username, passwordHash, avatarId string) (User, error) {
	u := User{Username: username, Password: passwordHash, AvatarId: avatarId}
	return u, DB.Create(&u).Error
}
