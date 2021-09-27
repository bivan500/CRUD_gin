package handler

import (
	crudApp "CRUD_GIN"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signup(c *gin.Context) {
	var input crudApp.User
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Auth.CreateUser(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}


func (h *Handler) signin(c *gin.Context) {
	var input crudApp.User
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Auth.GenerateToken(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
