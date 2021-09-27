package handler

import (
	"CRUD_GIN/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("signup", h.signup)
		auth.POST("signin", h.signin)
	}

	api := router.Group("/api", h.userIdentity)
	{
		UsersReadList := api.Group("/UsersReadList")
		{
			UsersReadList.GET("/", h.getUsersReadLists)
			UsersReadList.POST("/", h.createUsersReadList)
			UsersReadList.GET("/:id", h.getUsersReadListByID)
			UsersReadList.PUT("/:id", h.updateUsersReadList)
			UsersReadList.DELETE("/:id", h.deleteUsersReadList)
		}

		book := api.Group("/book")
		{
			book.GET("/", h.getBooks)
			book.POST("/", h.createBook)
			book.GET("/:id", h.getBookByID)
			book.PUT("/:id", h.updateBook)
			book.DELETE("/:id", h.deleteBook)
		}
	}

	return router

}
