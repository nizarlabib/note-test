package controllers

import (
	"net/http"
	model "sidita-be/models"
	"sidita-be/utils/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetWorklogs godoc
// @Summary Get Worklogs
// @Description Get Worklogs
// @Tags Worklog
// @Accept json
// @Produce json
// @Param page query string false "Page number"
// @Param limit query string false "Limit"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/worklog/get [get]
func GetWorklogs(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1")) 
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10")) 
	if err != nil || limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	worklogs, err := model.GetWorklogs(limit, offset)
	if err != nil {
		helper.SendResponse(c, http.StatusNotFound, "Failed to get worklogs", nil)
		return
	}

	total, err := model.CountAllWorklogs()
	if err != nil {
		total = 0
	}

	response := map[string]interface{}{
		"worklogs":  worklogs,
		"page":  page,
		"limit": limit,
		"total": total,
	}

	helper.SendResponse(c, http.StatusOK, "Worklogs found", response)
}

// GetWorklogsByUserId godoc
// @Summary Get Worklogs By User ID
// @Description Get Worklogs By User ID
// @Tags Worklog
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param page query string false "Page number"
// @Param limit query string false "Limit"
// @Security BearerAuth	
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}	
// @Router /api/worklog/get/user/{user_id} [get]
func GetWorklogsByUserId(c *gin.Context) {
	idParam := c.Param("user_id")

	uid, err := strconv.Atoi(idParam)
	if err != nil {
		helper.SendResponse(c, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1")) 
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10")) 
	if err != nil || limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	worklogs, err := model.GetWorkLogsByUserID(uid, limit, offset)
	if err != nil || len(worklogs) == 0 {
		helper.SendResponse(c, http.StatusNotFound, "Worklog not found", nil)
		return
	}

	total, err := model.CountWorklogsUser(uid)
	if err != nil {
		total = 0
	}

	response := map[string]interface{}{
		"worklogs":  worklogs,
		"page":  page,
		"limit": limit,
		"total": total,
	}

	helper.SendResponse(c, http.StatusOK, "Worklogs found", response)
}

type WorkLogInput struct {
	UserID      uint      `json:"user_id"`
	ProjectID uint   `json:"project_id"`
	WorkDate  string `json:"work_date"`
	HoursWorked int `json:"hours_worked"`
}

// InsertWorklog godoc
// @Summary Add Worklog
// @Description Add New Worklog
// @Tags Worklog
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body WorkLogInput true "Input Data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/worklog/add [post]
func InsertWorklog(c *gin.Context) {
	var input WorkLogInput

	if err := c.ShouldBindJSON(&input); err != nil {
		helper.SendResponse(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	if input.HoursWorked > 8 || input.HoursWorked < 0 {
		helper.SendResponse(c, http.StatusBadRequest, "Invalid hours worked", nil)
		return
	}

	w := model.Worklog{}

	w.UserID = input.UserID
	w.ProjectID = input.ProjectID
	w.WorkDate = input.WorkDate
	w.HoursWorked = input.HoursWorked

	data, err := w.SaveWorklog()

	if err != nil{
		helper.SendResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helper.SendResponse(c, http.StatusCreated, "Worklog created", data)
}

// DeleteWorklog godoc
// @Summary Delete Worklog
// @Description Delete Worklog
// @Tags Worklog
// @Accept json
// @Produce json
// @Param id path int true "Worklog ID"
// @Security BearerAuth	
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/worklog/delete/{id} [delete]
func DeleteWorklog(c *gin.Context) {

	id := c.Param("id")
	
	worklog, err := model.GetWorklogByID(id)
	if err != nil {
		helper.SendResponse(c, http.StatusNotFound, "Worklog not found", nil)
		return
	}
	
	if err := model.DeleteWorklog(worklog.ID); err != nil {
		helper.SendResponse(c, http.StatusInternalServerError, "Failed to delete worklog", nil)
		return
	}
	helper.SendResponse(c, http.StatusOK, "Worklog deleted", nil)
}

func GetTotalUserHoursWorkedByProject(c *gin.Context) {
	idParam := c.Param("user_id")

	uid, err := strconv.Atoi(idParam)
	if err != nil {
		helper.SendResponse(c, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}
	hours, err := model.CountUserHoursWorked(int(uid))
	if err != nil {
		helper.SendResponse(c, http.StatusNotFound, "Hours worked not found", nil)
		return
	}
	helper.SendResponse(c, http.StatusOK, "Hours worked", hours)
}
