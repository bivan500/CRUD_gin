package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createUsersReadList(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}
func (h *Handler) getUsersReadListByID(c *gin.Context) {

}
func (h *Handler) getUsersReadLists(c *gin.Context) {

}
func (h *Handler) updateUsersReadList(c *gin.Context) {

}
func (h *Handler) deleteUsersReadList(c *gin.Context) {

}
