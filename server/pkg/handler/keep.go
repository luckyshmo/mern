package handler

import "github.com/gin-gonic/gin"

func (h *Handler) keep(c *gin.Context) {
	h.services.GoogleKeep.GetAll()
}
