package controllers

import (
	"net/http"

	"github.com/TaushifReza/go-event-booking-api/dto"
	"github.com/TaushifReza/go-event-booking-api/services"
	"github.com/TaushifReza/go-event-booking-api/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(services *services.UserService) *UserController{
	return &UserController{UserService: services}
}

func (c *UserController) Register(ctx *gin.Context){
	var req dto.CreateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil{
		validationErrors := utils.FormatValidationErrors(err)
		if len(validationErrors) > 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
			return
		}
	
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.UserService.Create(&req)
	if err != nil {
		validationErrors := utils.FormatValidationErrors(err)
		if len(validationErrors) > 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	res := dto.UserResponse{
		ID: user.ID,
		Email: user.Email,
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully.", "data": res})
}

func (c *UserController) Login(ctx *gin.Context){
	var req dto.LoginRequest
	if err := ctx.ShouldBind(&req); err != nil{
		validationErrors := utils.FormatValidationErrors(err)
		if len(validationErrors) > 0{
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErrors})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	res, err := c.UserService.GetUserByEmail(&req)
	if err != nil{
		validationErrors := utils.FormatValidationErrors(err)
		if len(validationErrors) > 0{
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": validationErrors})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login success.", "data": res})
}