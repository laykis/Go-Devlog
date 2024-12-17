package handlers

import (
	"devlog/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

type UserHandler struct {
	DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{DB: db}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	var searchUser models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := h.DB.Where("user_name = ? AND use_yn = ?", user.UserName, "Y").Find(&searchUser)

	if result.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	user.RegisterDate = time.Now()
	user.UseYn = "Y"

	if err := h.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func (h *UserHandler) Login(c *gin.Context) {
	var user models.User
	var searchUser models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "아이디와 패스워드 모두 입력해주세요."})
		return
	}
	result := h.DB.Where("user_name = ? AND use_yn = ?", user.UserName, "Y").Find(&searchUser)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "존재하지 않는 아이디입니다."})
	} else {
		if searchUser.UserName != "" {
			if searchUser.UserName != user.UserName || searchUser.UserPassword != user.UserPassword {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "아이디 혹은 패스워드 오류입니다."})
			} else {
				c.JSON(http.StatusOK, gin.H{"user": searchUser.UserName + "님 환영합니다."})
			}
		}
	}
}
