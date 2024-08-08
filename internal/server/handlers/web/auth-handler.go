package web

import (
	"github.com/labstack/echo/v4"
	"github.com/tumivn/goblog/internal/domain/auth/dtos"
	"github.com/tumivn/goblog/internal/domain/auth/services"
	"github.com/tumivn/goblog/internal/server"
	services2 "github.com/tumivn/goblog/internal/server/services"
	"github.com/tumivn/goblog/internal/server/views/auth"
	"log"
	"net/http"
	"time"
)

type AuthHandler struct {
	server *server.Server
}

func NewAuthHandler(s *server.Server) *AuthHandler {
	return &AuthHandler{
		server: s,
	}
}

func (h *AuthHandler) GetSignUp(c echo.Context) error {
	user, _ := services2.GetCurrentUser(c, *h.server)
	if user != nil {
		return c.Redirect(http.StatusMovedPermanently, "/")
	}
	//TODO: SignUpVM
	data := new(auth.SignUpViewModel)
	hi := auth.SignInIndex("Sign In", auth.SignUp(*data))

	return RenderComponent(c, http.StatusOK, hi)
}

func (h *AuthHandler) GetSignIn(c echo.Context) error {
	user, _ := services2.GetCurrentUser(c, *h.server)
	if user != nil {
		return c.Redirect(http.StatusMovedPermanently, "/")
	}
	data := new(auth.SignInViewModel)
	hi := auth.SignInIndex("Sign In", auth.SignIn(*data))

	return RenderComponent(c, http.StatusOK, hi)
}

func (h *AuthHandler) PostSignIn(c echo.Context) error {

	var u = new(dtos.LoginRequest)
	{
		u.Email = c.FormValue("email")
		u.Password = c.FormValue("password")
	}
	data := new(auth.SignInViewModel)
	data.Email = u.Email

	err := u.Validate()
	if err != nil {
		log.Println(err.Error())

		data.ErrorMessage = err.Error()
		hi := auth.SignInIndex("Sign In", auth.SignIn(*data))
		return RenderComponent(c, http.StatusOK, hi)
	}

	res, err := services.AuthenticateUser(u, h.server.Config.JwtSecret)

	if err != nil {
		data.ErrorMessage = err.Error()
		hi := auth.SignInIndex("Sign In", auth.SignIn(*data))
		return RenderComponent(c, http.StatusOK, hi)
	}

	//Signed in successfully
	c.SetCookie(&http.Cookie{
		Name:     "token",
		Value:    res.Token,
		Expires:  res.Expires,
		HttpOnly: true,
		Path:     "/",
	})

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func (h *AuthHandler) GetSignButton(c echo.Context) error {
	user, _ := services2.GetCurrentUser(c, *h.server)
	return RenderComponent(c, http.StatusOK, auth.SignInButton(user))
}

func (h *AuthHandler) GetSignOut(c echo.Context) error {
	user, _ := services2.GetCurrentUser(c, *h.server)
	if user != nil {
		//logout
		c.SetCookie(&http.Cookie{
			Name:    "token",
			Value:   "",
			Expires: time.Now().Add(1 * time.Hour),
			Path:    "/",
		})

	}
	return c.Redirect(http.StatusMovedPermanently, "/")
}
