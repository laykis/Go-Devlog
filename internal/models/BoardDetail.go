package models

import "time"

type BoardDetail struct {
	Id           int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	UserId       int       `json:"userId" binding:"required" gorm:"column:user_id"`
	BoardId      int       `json:"boardId" binding:"required" gorm:"column:board_id"`
	BoardContent string    `json:"boardContent" binding:"required" gorm:"column:board_content"`
	RegisterDate time.Time `json:"registerDate" gorm:"column:register_date"`
	UpdateDate   time.Time `json:"updateDate" gorm:"column:update_date"`
	UpdateUserId int       `json:"updateUserId" gorm:"column:update_user_id"`
	UseYn        string    `json:"useYn" gorm:"column:use_yn"`
}

func (BoardDetail) TableName() string {
	return "board_detail" // 원하는 테이블 이름 지정
}
