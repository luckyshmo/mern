package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary Get User By Id
// @Security ApiKeyAuth
// @Tags user
// @Description get user by id
// @ID get-user-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user/:id [get]
func (h *Handler) getUser(c *gin.Context) {

	id := c.Param("id") //id in request

	uuid, err := uuid.Parse(id)
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.User.GetById(uuid)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendStatusResponse(c, http.StatusOK, user)
}

func (h *Handler) getUserList(c *gin.Context) {

	users, err := h.services.User.GetAll()
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendStatusResponse(c, http.StatusOK, users)
}
