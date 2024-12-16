package handlers

import (
	"devlog/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

type BoardHandler struct {
	DB *gorm.DB
}

func NewBoardHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{DB: db}
}

func (b *BoardHandler) BoardCreate(c *gin.Context) {
	var board models.Board

	if err := c.ShouldBindJSON(&board); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	
	if err := b.DB.Create(&board).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (b *BoardHandler) BoardDelete(c *gin.Context) {

}

func (b *BoardHandler) BoardList(c *gin.Context) {

}
