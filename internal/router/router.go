package router

import (
	"first-project/internal/handlers/v1"
	"first-project/internal/services"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Services services.Services
}

func NewRouter(services services.Services) *Router {
	return &Router{Services: services}
}

func (r *Router) Init() (*gin.Engine, error) {
	router := gin.Default()
	r.InitAPI(router)
	return router, nil
}

func (r *Router) InitAPI(router *gin.Engine) {
	v1.NewHandler(r.Services, router).Init()
}
