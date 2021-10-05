package handler

import (
	crudApp "CRUD_GIN"
	"net/http"
	"strconv"

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

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	list, err := h.services.GetListsById(listId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
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
		return
	}
	c.JSON(http.StatusOK, lists)
}

func (h *Handler) updateUsersReadList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	CurrenListId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var input crudApp.ReadList
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	_, err = h.services.UpdateList(userId, CurrenListId, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]bool{
		"result status": true,
	})
}
func (h *Handler) deleteUsersReadList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	ListId, err := strconv.Atoi(c.Param("id"))
	err = h.services.DeleteList(userId, ListId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "List deleted")
}
