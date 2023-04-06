package v1

import (
	"first-project/internal/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service  services.Services
	router   *gin.Engine
	response *response
}

func NewHandler(service services.Services, router *gin.Engine) *Handler {
	return &Handler{
		service:  service,
		router:   router,
		response: NewResponse(),
	}
}

func (h *Handler) Init() {
	v1 := h.router.Group("/api/v1")
	{
		h.initAuthRoutes(v1)
	}
	//h.router.Group("/api/v1", func(r chi.Router) {
	//	h.initAuthRoutes(r)
	//})
	//return router
}
