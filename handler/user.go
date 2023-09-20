package handler

import (
	"crowdfunding/auth"
	"crowdfunding/helper"
	"crowdfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service 
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var errors []string

		errors = helper.FormatValidationError(err)

		errorMessage := gin.H{ "errors" : errors }

		response := helper.APIResponse("Register user gagal di tambahkan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return 
	}

	 _, err = h.userService.IsEmailAvailable(input.Email)
	 if err != nil {
		errorMessage := gin.H{"errors": err.Error()} //memanggil error yang ada di service
		response := helper.APIResponse("Register user gagal di tambahkan", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response) 
		return
	 }

	 newUser, err := h.userService.RegisterUser(input)
	 if err != nil {
		response := helper.APIResponse("Register user gagal di tambahkan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return 
	 }

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.APIResponse("Register user gagal di tambahkan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, token)

	response := helper.APIResponse("Register User berhasil ditambahkan", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) Login(c *gin.Context) {

	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{ "errors" : errors}

		response := helper.APIResponse("Login Gagal", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.LoginUser(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()} //memanggil error yang ada di service
		response := helper.APIResponse("Login Gagal", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response) 
		return
	}

	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login Gagal", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	formatter := user.FormatUser(loggedinUser, token)

	response := helper.APIResponse("Login Berhasil", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)

	formatter := user.FormatUser(currentUser, "") //token nya di kosongin juga gk papa

	response := helper.APIResponse("Successfully fetch user datta", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}
