package controllers

import (
	"net/http"
	model "sidita-be/models"
	"sidita-be/utils/helper"

	"github.com/gin-gonic/gin"
)

// GetProjects godoc
// @Summary Get Projects
// @Description Get Projects
// @Tags Project
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/project/get [get]
func GetProjects(c *gin.Context) {
	projects, err := model.GetProjects()
	if err != nil {
		helper.SendResponse(c, http.StatusInternalServerError, "Failed to get projects", nil)
		return
	}
	helper.SendResponse(c, http.StatusOK, "Projects found", projects)
}

// GetProjectByID godoc
// @Summary Get Project By ID
// @Description Get Project By ID
// @Tags Project
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/project/{id} [get]
func GetProjectByID(c *gin.Context) {
	id := c.Param("id")

	project, err := model.GetProjectByID(id)
	if err != nil || project.ID == 0 {
		helper.SendResponse(c, http.StatusNotFound, "Project not found", nil)
		return
	}

	helper.SendResponse(c, http.StatusOK, "Project found", project)
}


// func CreateProject(c *gin.Context) {
// 	var project model.Project
// 	if err := c.ShouldBindJSON(&project); err != nil {
// 		utils.SendResponse(c, http.StatusBadRequest, "Invalid input", nil)
// 		return
// 	}

// 	err := model.CreateProject(&project)
// 	if err != nil {
// 		utils.SendResponse(c, http.StatusInternalServerError, "Failed to create project", nil)
// 		return
// 	}
// 	utils.SendResponse(c, http.StatusCreated, "Project created", project)
// }

func GetAllProjectsRecap(c *gin.Context) {
	var projects []model.Project
	var worklogs []model.Worklog

	worklogs, err := model.GetWorklogsNotPaginated()
	
	helper.SendResponse(c, http.StatusOK, "Ok", worklogs)
	for _, worklog := range worklogs {
		if worklog.Project.Name != "" {
			projects = append(projects, worklog.Project)
		}
	}
	if projects, err = model.GetProjects(); err != nil {
		helper.SendResponse(c, http.StatusInternalServerError, "Failed to get projects", nil)
		return
	}
	helper.SendResponse(c, http.StatusOK, "Projects fetched successfully", projects)
}