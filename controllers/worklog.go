package controllers

import (
	"net/http"
	model "sidita-be/models"
	"sidita-be/utils/helper"

	"github.com/gin-gonic/gin"
)

// GetWorklogs godoc
// @Summary Get Worklogs
// @Description Get Worklogs
// @Tags Worklog
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/worklog/get [get]
func GetWorklogs(c *gin.Context) {
	worklogs, err := model.GetWorklogs()
	if err != nil {
		helper.SendResponse(c, http.StatusInternalServerError, "Failed to get worklogs", nil)
		return
	}
	helper.SendResponse(c, http.StatusOK, "Worklogs found", worklogs)
}

// GetWorklogsByUserID godoc
// @Summary Get Worklogs By User ID
// @Description Get Worklogs By User ID
// @Tags Worklog
// @Accept json
// @Produce json
// @Param uid path string true "User ID"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/worklog/{uid} [get]
func GetWorklogsByUserID(c *gin.Context) {
	uid := c.Param("uid")

	worklogs, err := model.GetWorkLogsByUserID(uid)
	if err != nil {
		helper.SendResponse(c, http.StatusInternalServerError, "Failed to get worklogs", nil)
		return
	}
	helper.SendResponse(c, http.StatusOK, "Worklogs found", worklogs)
}