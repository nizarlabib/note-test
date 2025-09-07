package controllers

import (
	"net/http"
	model "sidita-be/models"
	"sidita-be/utils/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Summary Get Users
// @Description Get Users
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/user/get [get]
func GetUsers(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1")) 
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10")) 
	if err != nil || limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	users, err := model.GetUsers(limit, offset)
	if err != nil || len(users) == 0 {
		helper.SendResponse(c, http.StatusInternalServerError, "User not found", nil)
		return
	}

	total, err := model.CountAllUsers()
	if err != nil {
		total = 0
	}

	response := map[string]interface{}{
		"page":  page,
		"limit": limit,
		"users":  users,
		"total": total,
	}

	helper.SendResponse(c, http.StatusOK, "Users found", response)
}

// GetUserByID godoc
// @Summary Get User By ID
// @Description Get User By ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/user/{id} [get]
func GetUserByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		helper.SendResponse(c, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	user, err := model.GetUserByID(id)
	if err != nil || user.ID == 0 {
		helper.SendResponse(c, http.StatusNotFound, "User not found", nil)
		return
	}

	helper.SendResponse(c, http.StatusOK, "User found", user)
}