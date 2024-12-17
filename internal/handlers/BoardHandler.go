package handlers

import (
	"devlog/internal/constant"
	"devlog/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type BoardHandler struct {
	DB *gorm.DB
}

func NewBoardHandler(db *gorm.DB) *BoardHandler {
	return &BoardHandler{DB: db}
}

func (b *BoardHandler) BoardCreate(c *gin.Context) {
	var board models.Board

	if err := c.ShouldBindJSON(&board); err != nil {

		response := constant.NewApiResponse()
		response.Code = constant.STATUS_BAD_REQUEST_CODE
		response.Message = constant.STATUS_BAD_REQUEST_MSG
		response.Data = err.Error()
		
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err := b.DB.Create(&board).Error; err != nil {

		response := constant.NewApiResponse()
		response.Code = constant.STATUS_INTERNAL_DB_ERROR_CODE
		response.Message = constant.STATUS_INTERNAL_DB_ERROR_MSG
		response.Data = err.Error()

		c.JSON(http.StatusInternalServerError, response)
		return
	}
}

func (b *BoardHandler) BoardDelete(c *gin.Context) {

}

func (b *BoardHandler) BoardList(c *gin.Context) {

}
