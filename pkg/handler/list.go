package handler

import (
	crudApp "CRUD_GIN"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createUsersReadList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input crudApp.ReadList
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Create(userId, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}
func (h *Handler) getUsersReadListByID(c *gin.Context) {

}

type getUsersReadListsResponse struct {
	Data []crudApp.ReadList `json:"data"`
}

func (h *Handler) getUsersReadLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	lists, err := h.services.GetLists(userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, lists)
}
func (h *Handler) updateUsersReadList(c *gin.Context) {

}
func (h *Handler) deleteUsersReadList(c *gin.Context) {

}
