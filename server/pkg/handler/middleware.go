package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	authHeader = "Authorization"
	userCtx    = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	userId, err := parseHeader(c.GetHeader(authHeader), h.services.Authorization.ParseToken)
	if err != nil {
		sendErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	//add userCtx to metadata
	c.Set(userCtx, userId)
}

func parseHeader(header string, parse func(string) (uuid.UUID, error)) (uuid.UUID, error) {
	if header == "" {
		return uuid.Nil, errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" { //Bearer authentication (google it)
		return uuid.Nil, errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return uuid.Nil, errors.New("token is empty")
	}

	return parse(headerParts[1])
}

//return userID from context
func getUserId(c *gin.Context) (uuid.UUID, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return uuid.Nil, errors.New("user not found")
	}

	uuID, ok := id.(uuid.UUID)
	if !ok {
		return uuid.Nil, errors.New("user id is of invalid type")
	}

	return uuID, nil
}
