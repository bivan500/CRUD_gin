package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("signup", h.signup)
		auth.POST("signin", h.signin)
	}

	api := router.Group("/api")
	{
		list := api.Group("/list")
		{
			list.GET("/", h.getLists)
			list.POST("/", h.createList)
			list.GET("/:id", h.getListByID)
			list.PUT("/:id", h.updateList)
			list.DELETE("/:id", h.deleteList)
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
