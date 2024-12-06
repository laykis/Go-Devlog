package models

import "time"

type User struct {
	Id           int       `json:"id"`
	UserName     string    `json:"userName" binding:"required"`
	UserPassword string    `json:"userPassword" binding:"required"`
	RegisterDate time.Time `json:"registerDate"`
	UseYn        string    `json:"useYn"`
}

func (User) TableName() string {
	return "users" // 원하는 테이블 이름 지정
}
