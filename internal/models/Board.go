package models

import "time"

type Board struct {
	Id           int       `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	BoardName    string    `json:"boardName" binding:"required" gorm:"column:board_name"`
	RegisterDate time.Time `json:"registerDate" gorm:"column:register_date"`
	UseYn        string    `json:"useYn" gorm:"column:use_yn"`
}

func (Board) TableName() string {
	return "boards" // 원하는 테이블 이름 지정
}
