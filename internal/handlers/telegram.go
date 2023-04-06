package handlers

import (
	"first-project/pkg/e"
	"first-project/pkg/response"
	telegramPkg "first-project/pkg/telegram"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type telegram struct {
	Telegram *telegramPkg.Client
}

func newTelegram() *telegram {
	telApiId, _ := strconv.Atoi(os.Getenv("TELEGRAM_API_ID"))
	telHash := os.Getenv("TELEGRAM_API_HASH")
	return &telegram{
		Telegram: telegramPkg.New(telApiId, telHash),
	}
}

type authStruct struct {
	Phone string `json:"phone" validate:"e164"`
}

func (t *telegram) AuthSendCode(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			response.Error(w, r, err)
		}
	}()
	err = r.ParseForm()
	if err != nil {
		return
	}
	// Remove extra spaces
	phone := strings.Replace(strings.TrimSpace(r.Form.Get("phone")), " ", "", -1)

	data := &authStruct{
		Phone: phone,
	}
	err = validate.Struct(data)
	if err != nil {
		return
	}

	user, err := t.Telegram.NewFlow(data.Phone)
	if err != nil {
		response.Error(w, r, err)
	}
	e.Info("Telegram authSendCode response:", user)

	response.MakeResponse(w, r, http.StatusOK, "", user)
}

type authSingIn struct {
	authStruct
	Code string `json:"code" validate:"required"`
}

func (t *telegram) SingIn(w http.ResponseWriter, r *http.Request) {
	phone, phone_hash, code := r.FormValue("phone"), r.FormValue("phone_hash"), r.FormValue("code")

	_, err := t.Telegram.SingIn(phone, phone_hash, code)
	if err != nil {
		response.Error(w, r, err)
	}
}

func (t *telegram) LogOut(w http.ResponseWriter, r *http.Request) {
	t.Telegram.LogOut()
	//if err != nil {
	//	e.Error("Telegram logout:", err)
	//}
	//e.Info("LogOut", logout)
}

// Login Авторизация по номеру телефона
func (t *telegram) Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	phone := r.FormValue("phone")
	user, err := t.Telegram.NewFlow(phone)
	if err != nil {
		response.Error(w, r, err)
	}
	response.JSON(w, r, map[string]interface{}{
		"user": user,
	})
}

func (t *telegram) GetSelf(w http.ResponseWriter, r *http.Request) {
	self, err := t.Telegram.GetSelf()
	if err != nil {
		response.Error(w, r, err)
	}
	response.JSON(w, r, self)
}

// Тестовый метод для дебага
func (t *telegram) Test(w http.ResponseWriter, r *http.Request) {
	t.Telegram.Test()
}
