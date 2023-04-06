package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
}
type any interface{}

func NewResponse() *response {
	return &response{}
}
func (r *response) MakeResponse(c *gin.Context, code int, data any) {
	c.JSON(code, data)
}
func (r *response) ErrValidate(c *gin.Context, err error) error {
	return c.AbortWithError(http.StatusForbidden, err)
}

func (r *response) Error(c *gin.Context, err error) error {
	return c.AbortWithError(http.StatusInternalServerError, err)
}

func (r *response) Success(c *gin.Context, data any) {
	r.MakeResponse(c, http.StatusOK, data)
}

func (r *response) Created(c *gin.Context, data any) {
	r.MakeResponse(c, http.StatusCreated, data)
}
