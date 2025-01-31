package controllers

import (
	"math"
	"net/http"
	model "sidita-be/models"
	"sidita-be/utils/helper"
	"strconv"
	"strings"
	"time"

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


// GetUserAbsentRecap godoc
// @Summary Get User Absent Recap
// @Description Get User Absent Recap
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/user/get/absent/recap [get]
func GetUserAbsentRecap(c *gin.Context) {
	users, err := model.GetUsersNotPaginated() 
	if err != nil {
		helper.SendResponse(c, http.StatusNotFound, "Users not found", nil)
		return
	}

	userAbsentRecap := make(map[uint]map[string]interface{})

	for i := 0; i < len(users); i++ {
		userAbsentRecap[users[i].ID] = map[string]interface{}{
			"id":               users[i].ID,
			"name":             users[i].Name,
			"email":            users[i].Email,
			"total_hours_worked": float64(0),
			"total_days_worked": float64(0),
			"recaps_per_month": map[string]map[string]float64{},
		}
	}

	allWorklogs, err := model.GetWorklogsNotPaginated()

	if err != nil {
		
	}


	for i := 0; i < len(allWorklogs); i++ {
		if user, exists := userAbsentRecap[allWorklogs[i].UserID]; exists {
			currentTotal, _ := user["total_hours_worked"].(float64)
	
			newTotalHours := currentTotal + float64(allWorklogs[i].HoursWorked)
			userAbsentRecap[allWorklogs[i].UserID]["total_hours_worked"] = newTotalHours
			userAbsentRecap[allWorklogs[i].UserID]["total_days_worked"] = newTotalHours / 8
	
			yearMonth := allWorklogs[i].WorkDate[:7]
	
			if _, exists := userAbsentRecap[allWorklogs[i].UserID]["recaps_per_month"]; !exists {
				userAbsentRecap[allWorklogs[i].UserID]["recaps_per_month"] = make(map[string]map[string]float64)
			}
	
			averageDaysMap := userAbsentRecap[allWorklogs[i].UserID]["recaps_per_month"].(map[string]map[string]float64)
	
			if _, exists := averageDaysMap[yearMonth]; !exists {
				daysInMonth := float64(GetDaysInMonth(yearMonth))
				averageDaysMap[yearMonth] = map[string]float64{
					"total_days_worked":  0,
					"total_absent_days":   0,
					"days_worked_percent":   0,
					"total_days_in_month": daysInMonth,
				}
			}
	
			averageDaysMap[yearMonth]["total_days_worked"] += float64(allWorklogs[i].HoursWorked) / 8
			averageDaysMap[yearMonth]["total_absent_days"] = averageDaysMap[yearMonth]["total_days_in_month"] - averageDaysMap[yearMonth]["total_days_worked"]
			percentage := float64((averageDaysMap[yearMonth]["total_days_worked"] / averageDaysMap[yearMonth]["total_days_in_month"]) * 100)
			averageDaysMap[yearMonth]["days_worked_percent"] = math.Round(percentage*100) / 100
		}
	}

	var response []map[string]interface{}

	for _, user := range userAbsentRecap {
		response = append(response, user)
	}

	helper.SendResponse(c, http.StatusOK, "Data found", response)
}

// GetUserAbsentRecapByUID godoc
// @Summary Get User Absent Recap By User ID
// @Description Get User Absent Recap By User ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/user/get/absent/recap/byuser/{id} [get]
func GetUserAbsentRecapByUID(c *gin.Context) {
	idParam := c.Param("id")

	uid, err := strconv.Atoi(idParam)
	if err != nil {
		helper.SendResponse(c, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	user, err := model.GetUserByID(uid)
	if err != nil || user.ID == 0 {
		helper.SendResponse(c, http.StatusNotFound, "User not found", nil)
		return
	}

	userAbsentRecap := map[string]interface{}{
		"id":               user.ID,
		"name":             user.Name,
		"email":            user.Email,
		"total_hours_worked": float64(0),
		"total_days_worked": float64(0),
		"total_absent_days": float64(0),
		"recaps_per_month": map[string]map[string]float64{},
	}

	userWorklogs, err := model.GetWorkLogsByUserIDNotPaginated(uid)

	if err != nil {
		
	}


	for i := 0; i < len(userWorklogs); i++ {
		currentTotal, _ := userAbsentRecap["total_hours_worked"].(float64)

		newTotalHours := currentTotal + float64(userWorklogs[i].HoursWorked)
		userAbsentRecap["total_hours_worked"] = newTotalHours
		userAbsentRecap["total_days_worked"] = newTotalHours / 8

		yearMonth := userWorklogs[i].WorkDate[:7]

		if _, exists := userAbsentRecap["recaps_per_month"]; !exists {
			userAbsentRecap["recaps_per_month"] = make(map[string]map[string]float64)
		}

		averageDaysMap := userAbsentRecap["recaps_per_month"].(map[string]map[string]float64)

		if _, exists := averageDaysMap[yearMonth]; !exists {
			daysInMonth := float64(GetDaysInMonth(yearMonth))
			averageDaysMap[yearMonth] = map[string]float64{
				"total_days_worked":  0,
				"total_absent_days":   0,
				"days_worked_percent":   0,
				"total_days_in_month": daysInMonth,
			}
		}

		averageDaysMap[yearMonth]["total_days_worked"] += float64(userWorklogs[i].HoursWorked) / 8
		averageDaysMap[yearMonth]["total_absent_days"] = averageDaysMap[yearMonth]["total_days_in_month"] - averageDaysMap[yearMonth]["total_days_worked"]
		percentage := float64((averageDaysMap[yearMonth]["total_days_worked"] / averageDaysMap[yearMonth]["total_days_in_month"]) * 100)
		averageDaysMap[yearMonth]["days_worked_percent"] = math.Round(percentage*100) / 100
	}


    totalAbsentDays := 0.0

    recapsPerMonth, ok := userAbsentRecap["recaps_per_month"].(map[string]map[string]float64)
    if ok {
        for _, monthData := range recapsPerMonth {
            totalAbsentDays += monthData["total_absent_days"]
        }
    }

	if totalAbsentDays > 0 {
		userAbsentRecap["total_absent_days"] = totalAbsentDays
	}


	helper.SendResponse(c, http.StatusOK, "Data found", userAbsentRecap)
}


func GetDaysInMonth(dateStr string) (float64) {
	parts := strings.Split(dateStr, "-")
	

	year, err := strconv.Atoi(parts[0])
	if err != nil {
	}

	monthInt, err := strconv.Atoi(parts[1])
	if err != nil {
	}

	days := time.Date(year, time.Month(monthInt+1), 0, 0, 0, 0, 0, time.UTC).Day()
	return float64(days)
}