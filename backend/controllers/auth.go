package controllers

import (
	"zibon/common/response"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required,min=4"`
	Password string `json:"password" binding:"required,min=8"`
	Phone    string `json:"phone"`
}

func (h handler) RegisterUser(c *gin.Context) {
	resp := response.New()
	data := RegisterRequest{}

	if err := c.ShouldBindJSON(&data); err != nil {
		resp.SetError(response.VALIDATION_ERROR, err)
		c.JSON(200, resp)

		return
	}

	resp.SetData(data)
	c.JSON(200, resp)
}
