package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	autarizationHeader = "Authorization"
	userCtx            = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(autarizationHeader)
	if header == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	userId, err := h.services.Auth.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "invalid user id")
		return 0, errors.New("invalid user id")
	}

	idInt, ok := id.(int)
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "invalid user id")
		return 0, errors.New("invalid type of user id")
	}

	return idInt, nil
}
