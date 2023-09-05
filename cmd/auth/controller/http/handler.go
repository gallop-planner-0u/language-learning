package http

import (
	"language-learning/cmd/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc auth.Usecase
}

func New(uc auth.Usecase) *Handler {
	return &Handler{
		uc: uc,
	}
}

func (h *Handler) SignUp(c *gin.Context) {
	ctx := c.Request.Context()
	_, err := h.uc.SignUp(ctx, nil)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	c.AbortWithStatus(http.StatusCreated)

}

func (h *Handler) SignIn(c *gin.Context) {
	var credentials interface{}
	if err := c.BindJSON(&credentials); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	ctx := c.Request.Context()

	var username, password string
	accessToken, err := h.uc.SignIn(ctx, username, password)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, &accessToken)
}

func (h *Handler) AuthMiddleware(c *gin.Context) {
	accessToken := c.Request.Header.Get("Authorization")
	claims, err := h.uc.ParseToken(accessToken)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.Set("claims", claims)
}

func RegisterRoutes(r *gin.Engine, h *Handler) {
	group := r.Group("/auth")

	group.POST("/sign-up", h.SignUp)
	group.POST("/sign-in", h.SignIn)
}
