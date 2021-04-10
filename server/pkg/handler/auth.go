package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luckyshmo/api-example/models"
)

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signUp(c *gin.Context) {

	var input models.User

	if err := c.BindJSON(&input); err != nil {
		sendErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendStatusResponse(c, http.StatusOK, map[string]interface{}{
		"id": id, //JSON body
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendStatusResponse(c, http.StatusOK, map[string]interface{}{
		"token": token, //JSON body
	})
}
