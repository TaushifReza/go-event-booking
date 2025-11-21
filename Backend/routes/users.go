package routes

import (
	"fmt"
	"net/http"

	"github.com/TaushifReza/go-event-booking-api/dto"
	"github.com/TaushifReza/go-event-booking-api/models"
	"github.com/TaushifReza/go-event-booking-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func signup(c *gin.Context){
	var user  models.User
	err := c.ShouldBind(&user)

	if err != nil{
		if errs, ok := err.(validator.ValidationErrors); ok {
			out := make(map[string]string)
			for _, e := range errs {
				out[e.Field()] = fmt.Sprintf("%s is %s", e.Field(), e.ActualTag())
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": out})
			return
		}
		
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	err = user.Save()

	if err != nil{
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong. Please try again.", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Register successfully."})
}

func login(c *gin.Context){
	var loginDto  dto.LoginRequest
	err := c.ShouldBind(&loginDto)

	if err != nil{
		if errs, ok := err.(validator.ValidationErrors); ok {
			out := make(map[string]string)
			for _, e := range errs {
				out[e.Field()] = fmt.Sprintf("%s is %s", e.Field(), e.ActualTag())
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": out})
			return
		}
		
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	result , err := models.GetUserByEmail(loginDto.Email)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email or password."})
		return
	}

	// validate password
	err = utils.CheckPassword(loginDto.Password, result.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email or password."})
		return
	}

	res := dto.UserResponse{
		ID: result.ID,
		Email: result.Email,
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login success", "result": res})
}