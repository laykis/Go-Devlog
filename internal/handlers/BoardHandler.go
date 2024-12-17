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

		response := constant.NewApiResponse().BadReqResp(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err := b.DB.Create(&board).Error; err != nil {

		response := constant.NewApiResponse().InternalDbErrorResp(err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := constant.NewApiResponse().OkResp()
	c.JSON(http.StatusOK, response)
}

func (b *BoardHandler) BoardDelete(c *gin.Context) {
	var board models.Board
	var searchBoard models.Board

	// 1. JSON 파싱
	if err := c.ShouldBindJSON(&board); err != nil {
		c.JSON(http.StatusBadRequest, constant.NewApiResponse().BadReqResp(err))
		return
	}

	// 2. DB에서 대상 레코드 조회
	result := b.DB.First(&searchBoard, "id = ?", board.Id)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			// 레코드가 존재하지 않는 경우
			c.JSON(http.StatusBadRequest, constant.NewApiResponse().BoardNotExistResp())
			return
		}
		// 기타 DB 에러
		c.JSON(http.StatusInternalServerError, constant.NewApiResponse().InternalDbErrorResp(result.Error))
		return
	}

	// 3. UseYn 업데이트
	if err := b.DB.Model(&searchBoard).Update("use_yn", "N").Error; err != nil {
		c.JSON(http.StatusInternalServerError, constant.NewApiResponse().InternalDbErrorResp(err))
		return
	}

	// 4. 성공 응답
	c.JSON(http.StatusOK, constant.NewApiResponse().OkResp())
}

func (b *BoardHandler) BoardList(c *gin.Context) {
	var board []models.Board

	if err := b.DB.Where("use_yn = ?", constant.USE_YN_Y).Find(&board).Error; err != nil {
		response := constant.NewApiResponse().InternalDbErrorResp(err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := constant.NewApiResponse().OkRespWithData(board)
	c.JSON(http.StatusOK, response)

}
