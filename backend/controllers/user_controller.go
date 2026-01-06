package controllers

import (
	"net/http"
	"strings"

	"github.com/TaushifReza/go-event-booking-api/dto"
	"github.com/TaushifReza/go-event-booking-api/internal/http/errors"
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
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))

	user, err := c.UserService.Register(&req)
	if err != nil {
		errors.HandleError(ctx, "Registration failed.", err)
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
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))

	res, err := c.UserService.LoginUser(&req)
	if err != nil{
		errors.HandleError(ctx, "Login failed.", err)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessResponse("Login success.", res))
}

func (c *UserController) GetUserInfo(ctx *gin.Context){
	email := ctx.GetString("email")

	res , err := c.UserService.GetUserInfo(email)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to get user data.", err))
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("User data fetch successfully", res))
}

func (c *UserController) RefreshToken(ctx *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationErrors := utils.FormatValidationErrors(err)
		if len(validationErrors) > 0{
			ctx.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors))
			return
		}

		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body", err))
		return
	}

	res, err := c.UserService.VerifyRefreshToken(req.RefreshToken)
	if err != nil{
		errors.HandleError(ctx, "something went wrong. please try again", err)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessResponse("Successfully generate new access token.", res))
}