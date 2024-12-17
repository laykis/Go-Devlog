package handlers

import (
	"devlog/internal/constant"
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
		response := constant.NewApiResponse().BadReqResp(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	result := h.DB.Where("user_name = ? AND use_yn = ?", user.UserName, constant.USE_YN_Y).Find(&searchUser)

	if result.Error != nil {
		response := constant.NewApiResponse().InternalDbErrorResp(result.Error)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if result.RowsAffected > 0 {
		response := constant.NewApiResponse().UserExistResp()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user.RegisterDate = time.Now()
	user.UseYn = constant.USE_YN_Y

	if err := h.DB.Create(&user).Error; err != nil {
		response := constant.NewApiResponse().InternalDbErrorResp(err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := constant.NewApiResponse().OkResp()
	c.JSON(http.StatusCreated, response)
}

func (h *UserHandler) Login(c *gin.Context) {
	var user models.User
	var searchUser models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		response := constant.NewApiResponse().RequireIdPasswordResp()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	result := h.DB.Where("user_name = ? AND use_yn = ?", user.UserName, constant.USE_YN_Y).Find(&searchUser)

	if result.RowsAffected == 0 {
		response := constant.NewApiResponse().UserNotExistResp()
		c.JSON(http.StatusUnauthorized, response)
	}
	
	if result.Error != nil {
		response := constant.NewApiResponse().InternalDbErrorResp(result.Error)
		c.JSON(http.StatusInternalServerError, response)
	} else {
		if searchUser.UserName != "" {
			if searchUser.UserName != user.UserName || searchUser.UserPassword != user.UserPassword {
				response := constant.NewApiResponse().LoginFailResp()
				c.JSON(http.StatusUnauthorized, response)
			} else {
				response := constant.NewApiResponse().LoginSuccessResp(searchUser.UserName)
				c.JSON(http.StatusOK, response)
			}
		}
	}
}
