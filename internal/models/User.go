package models

import "time"

type User struct {
	Id           int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`                // DB 컬럼명 "id"
	UserName     string    `json:"userName" binding:"required" gorm:"column:user_name"`         // DB 컬럼명 "user_name"
	UserPassword string    `json:"userPassword" binding:"required" gorm:"column:user_password"` // DB 컬럼명 "user_password"
	RegisterDate time.Time `json:"registerDate" gorm:"column:register_date"`                    // DB 컬럼명 "register_date"
	UseYn        string    `json:"useYn" gorm:"column:use_yn"`                                  // DB 컬럼명 "use_yn"`
}

// TableName 메서드로 테이블명 지정
func (User) TableName() string {
	return "users" // 매핑할 테이블 이름
}
