package v1

import (
	"first-project/internal/database/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initAuthRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", h.authRegister)
		auth.POST("/login", h.authLogin)
	}
}

type authRegisterInput struct {
	Email             string `json:"email" form:"email" binding:"required"`
	Password          string `json:"password" form:"password"`
	PasswordConfirmed string `json:"password_confirmed" form:"password_confirmed"`
}

func (h *Handler) authRegister(c *gin.Context) {
	var input authRegisterInput
	err := c.Bind(&input)
	if err != nil {
		h.response.ErrValidate(c, err)
	}
	user, err := h.service.Auth.Register(&models.User{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		h.response.Error(c, err)
	}
	//r.ParseMultipartForm(1)
	//fmt.Println(r.PostForm, r.MultipartForm, r.Form)
	//response.Success(w, r, "Register")
	h.response.Success(c, user)
}

type authLoginInput struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (h *Handler) authLogin(ctx *gin.Context) {
	var userInput authLoginInput
	err := ctx.Bind(&userInput)
	if err != nil {
		h.response.ErrValidate(ctx, err)
	}
	token, err := h.service.Auth.Login(models.User{
		Email:    userInput.Email,
		Password: userInput.Password,
	})
	if err != nil {
		h.response.Error(ctx, err)
	}
	h.response.Success(ctx, token)
}
