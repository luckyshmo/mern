package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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
