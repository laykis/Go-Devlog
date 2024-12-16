package models

import "time"

type Board struct {
	Id           int       `json:"id"`
	BoardName    string    `json:"boardName" binding:"required"`
	RegisterDate time.Time `json:"registerDate"`
	UseYn        string    `json:"useYn"`
}

func (Board) TableName() string {
	return "boards" // 원하는 테이블 이름 지정
}
