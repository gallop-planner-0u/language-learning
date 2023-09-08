package http

import (
	"language-learning/cmd/dictionary"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc dictionary.Usecase
}

func New(uc dictionary.Usecase) *Handler {
	return &Handler{
		uc: uc,
	}
}

func (h *Handler) GetByID(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (h *Handler) GetByWord(c *gin.Context) {}

func (h *Handler) Create(c *gin.Context) {}

func (h *Handler) Update(c *gin.Context) {}

func (h *Handler) Delete(c *gin.Context) {}

func RegisterHTTPRoutes(r *gin.Engine, h *Handler, authMiddleware gin.HandlerFunc) {
	d := r.Group("/dictionary")
	d.Use(authMiddleware)

	d.GET("/:id", h.GetByID)
	d.GET("/word/:value", h.GetByWord)
	d.POST("/", h.Create)
	d.PUT("/:id", h.Update)
	d.DELETE("/:id", h.Delete)
}
