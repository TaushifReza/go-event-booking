package controllers

import (
	"net/http"

	"github.com/TaushifReza/go-event-booking-api/dto"
	"github.com/TaushifReza/go-event-booking-api/logger"
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
			ctx.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors))
			return
		}
	
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body", err))
		return
	}

	user, err := c.UserService.Register(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Registration failed.", err))
		return
	}

	res := dto.UserResponse{
		ID: user.ID,
		Email: user.Email,
	}

	ctx.JSON(http.StatusCreated, utils.SuccessResponse("User created successfully.", res))
}

func (c *UserController) Login(ctx *gin.Context){
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil{
		validationErrors := utils.FormatValidationErrors(err)
		if len(validationErrors) > 0{
			ctx.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors))
			return
		}

		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body", err))
		return
	}

	res, err := c.UserService.LoginUser(&req)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Login failed.", err))
		return
	}
	logger.Info("User login success")
	ctx.JSON(http.StatusOK, utils.SuccessResponse("Login success.", res))
}