package controllers

import (
	"net/http"
	"regexp"
	"sidita-be/models"

	"sidita-be/utils/helper"
	"sidita-be/utils/token"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login godoc
// @Summary Login User
// @Description Login User
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body LoginInput true "Input Data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/auth/login [post]
func Login(c *gin.Context) {
	
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		helper.SendResponse(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	u := models.User{}

	u.Email = input.Email
	u.Password = input.Password

	token, uid, err := models.LoginCheck(u.Email, u.Password)

	if err != nil {
		helper.SendResponse(c, http.StatusBadRequest, "Email or password is incorrect.", nil)
		return
	}

	err = models.CreateLog(&models.Log{
		EndPoint:  c.FullPath(),
		Method:    c.Request.Method,
		UserID:    *uid,
	})
	if err != nil {
		helper.SendResponse(c, http.StatusInternalServerError, "Failed to create log", nil)
		return
	}
	if err != nil {
		helper.SendResponse(c, http.StatusInternalServerError, "Failed to create note", nil)
		return
	}
	
	helper.SendResponse(c, http.StatusOK, "Login success", gin.H{"token":token})

}

type RegisterInput struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Register godoc
// @Summary Register User
// @Description Register User
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body RegisterInput true "Input Data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/auth/register [post]
func Register(c *gin.Context){
	
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		helper.SendResponse(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(input.Email) {
		helper.SendResponse(c, http.StatusBadRequest, "Invalid email format", nil)
		return
	}

	u := models.User{}

	u.Email = input.Email
	u.Name = input.Name
	u.Password = input.Password

	data, err := u.SaveUser()

	if err != nil{
		helper.SendResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err = models.CreateLog(&models.Log{
		EndPoint:  c.FullPath(),
		Method:    c.Request.Method,
		UserID:    u.ID,
	})
	if err != nil {
		helper.SendResponse(c, http.StatusInternalServerError, "Failed to create log", nil)
		return
	}

	helper.SendResponse(c, http.StatusOK, "Registration success", data)
}

// CurrentUser godoc
// @Summary Current User
// @Description Get details of the currently authenticated user
// @Tags Auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/me [get]
func CurrentUser(c *gin.Context){

	user_id, err := token.ExtractTokenID(c)
	
	if err != nil {
		helper.SendResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	
	u,err := models.GetUserByID(int(user_id))
	
	if err != nil {
		helper.SendResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	helper.SendResponse(c, http.StatusOK, "Success get user", u)
}