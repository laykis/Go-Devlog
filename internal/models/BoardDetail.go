package models

import "time"

type BoardDetail struct {
	Id           int       `json:"id"`
	UserId       int       `json:"userId" binding:"required"`
	BoardId      int       `json:"boardId" binding:"required"`
	BoardContent string    `json:"boardContent" binding:"required"`
	RegisterDate time.Time `json:"registerDate"`
	UpdateDate   time.Time `json:"updateDate"`
	UpdateUserId int       `json:"updateUserId"`
	UseYn        string    `json:"useYn"`
}

func (BoardDetail) TableName() string {
	return "board_detail" // 원하는 테이블 이름 지정
}
