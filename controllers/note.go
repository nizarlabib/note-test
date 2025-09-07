package controllers

import (
	"net/http"
	"strconv"
	"sidita-be/models"
	"sidita-be/utils/helper"
	"sidita-be/utils/token"

	"github.com/gin-gonic/gin"
)

// CreateNote godoc
// @Summary Create a new note
// @Description Create a new note for the authenticated user
// @Tags Notes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body models.Note true "Note input"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/note/add [post]
func CreateNote(c *gin.Context) {
	var input models.Note

	if err := c.ShouldBindJSON(&input); err != nil {
		helper.SendResponse(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		helper.SendResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	input.UserID = user_id

	note, err := models.CreateNote(&input)
	if err != nil {
		helper.SendResponse(c, http.StatusInternalServerError, "Failed to create note", nil)
		return
	}
	
	helper.SendResponse(c, http.StatusOK, "Note created successfully", note)
}

// GetAllNote godoc
// @Summary Get all notes with pagination
// @Description Get list of notes (paginated) for the authenticated user
// @Tags Notes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number (default is 1)"
// @Param limit query int false "Number of items per page (default is 10)"
// @Success 200 {object} helper.PaginationResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/note/get [get]
func GetAllNote(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1")) 
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10")) 
	if err != nil {
		limit = 10
	}

	// bikin 1 instance pagination saja
	pagination := &helper.Pagination{
		Limit: limit,
		Page:  page,
	}

	// pass ke model
	if err := models.GetAllNote(&models.Note{}, pagination); err != nil {
		helper.SendResponse(c, http.StatusInternalServerError, "Failed to get notes", nil)
		return
	}

	// sekarang Rows dan TotalRows sudah terisi
	helper.SuccessDataPaginate(
		"success",
		"notes",
		pagination.Rows,
		&helper.Page{
			TotalData:   pagination.TotalRows,
			Limit:       pagination.GetLimit(),
			CurrentPage: pagination.GetPage(),
			TotalPage:   pagination.TotalPages,
		},
		c.Writer,
	)
}


// GetNoteByID godoc
// @Summary Get note by ID
// @Description Get a single note by its ID
// @Tags Notes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Note ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/note/{id} [get]
func GetNoteByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		helper.SendResponse(c, http.StatusBadRequest, "Invalid note ID", nil)
		return
	}

	note, err := models.GetNoteByID(id)
	if err != nil || note.ID == 0 {
		helper.SendResponse(c, http.StatusNotFound, "Note not found", nil)
		return
	}

	helper.SendResponse(c, http.StatusOK, "Note found", note)
}

// UpdateNote godoc
// @Summary Update note
// @Description Update a note by its ID
// @Tags Notes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Note ID"
// @Param body body models.Note true "Updated note data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/note/update/{id} [put]
func UpdateNote(c *gin.Context) {
	var input models.Note

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		helper.SendResponse(c, http.StatusBadRequest, "Invalid note ID", nil)
		return
	}

	input.ID = uint(id)

	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		helper.SendResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	input.UserID = user_id

	if err := c.ShouldBindJSON(&input); err != nil {
		helper.SendResponse(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	note, err := models.UpdateNote(&input)
	if err != nil {
		helper.SendResponse(c, http.StatusInternalServerError, "Failed to update note", nil)
		return
	}
	
	helper.SendResponse(c, http.StatusOK, "Note updated successfully", note)
}

// DeleteNote godoc
// @Summary Delete note
// @Description Delete a note by its ID
// @Tags Notes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Note ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/note/delete/{id} [delete]
func DeleteNote(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		helper.SendResponse(c, http.StatusBadRequest, "Invalid note ID", nil)
		return
	}

	err = models.DeleteNote(id)
	if err != nil {
		helper.SendResponse(c, http.StatusInternalServerError, "Failed to delete note", nil)
		return
	}
	
	helper.SendResponse(c, http.StatusOK, "Note deleted successfully", nil)
}

