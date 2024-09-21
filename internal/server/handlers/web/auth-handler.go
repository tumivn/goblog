package web

import (
	"github.com/labstack/echo/v4"
	"github.com/tumivn/goblog/internal/domain/auth/dtos"
	"github.com/tumivn/goblog/internal/domain/auth/services"
	"github.com/tumivn/goblog/internal/server"
	services2 "github.com/tumivn/goblog/internal/server/services"
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
	var tick = time.Now().Format("20060102150405")
	user, _ := services2.GetCurrentUser(c, *h.server)
	if user != nil {
		return c.Redirect(http.StatusMovedPermanently, "/")
	}
	return c.Render(http.StatusOK, "sign-up.html", map[any]interface{}{
		"name":         "Login",
		"currentTime":  tick,
		"email":        "",
		"firstname":    "",
		"lastname":     "",
		"username":     "",
		"password":     "",
		"errorMessage": "",
	})
}

func (h *AuthHandler) PostSignUp(c echo.Context) error {
	user, _ := services2.GetCurrentUser(c, *h.server)
	if user != nil {
		return c.Redirect(http.StatusMovedPermanently, "/")
	}

	var tick = time.Now().Format("20060102150405")
	email := c.FormValue("email")
	password := c.FormValue("password")
	errorMessage := ""
	var u = new(dtos.CreateUserRequest)
	{
		u.Email = email
		u.Password = password
		u.Username = c.FormValue("username")
		u.Firstname = c.FormValue("firstname")
		u.Lastname = c.FormValue("lastname")
	}

	err := u.Validate()
	if err != nil {
		errorMessage = err.Error()
		return c.Render(http.StatusOK, "sign-up.html", map[any]interface{}{
			"name":         "Sign Up",
			"currentTime":  tick,
			"email":        u.Email,
			"password":     "",
			"username":     u.Username,
			"firstname":    u.Firstname,
			"lastname":     u.Lastname,
			"errorMessage": errorMessage,
		})
	}

	//Create user
	newUser, err := services.CreatUser(*u)

	if err != nil {
		errorMessage = err.Error()
		return c.Render(http.StatusOK, "sign-up.html", map[any]interface{}{
			"name":         "Sign Up",
			"currentTime":  tick,
			"email":        u.Email,
			"password":     "",
			"username":     u.Username,
			"firstname":    u.Firstname,
			"lastname":     u.Lastname,
			"errorMessage": errorMessage,
		})
	}

	return c.Render(http.StatusOK, "sign-up-success.html", newUser)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func (h *AuthHandler) GetSignIn(c echo.Context) error {
	var tick = time.Now().Format("20060102150405")
	user, _ := services2.GetCurrentUser(c, *h.server)
	if user != nil {
		return c.Redirect(http.StatusMovedPermanently, "/")
	}
	return c.Render(http.StatusOK, "sign-in.html", map[any]interface{}{
		"name":         "Login",
		"currentTime":  tick,
		"email":        "",
		"password":     "",
		"errorMessage": "",
	})
}

func (h *AuthHandler) PostSignIn(c echo.Context) error {
	var tick = time.Now().Format("20060102150405")
	user, _ := services2.GetCurrentUser(c, *h.server)
	if user != nil {
		return c.Redirect(http.StatusMovedPermanently, "/")
	}
	email := c.FormValue("email")
	password := c.FormValue("password")
	errorMessage := ""
	var u = new(dtos.LoginRequest)
	{
		u.Email = email
		u.Password = password
	}

	err := u.Validate()
	if err != nil {
		errorMessage = err.Error()
		println(errorMessage)
		return c.Render(http.StatusOK, "sign-in.html", map[any]interface{}{
			"name":         "Login",
			"currentTime":  tick,
			"email":        email,
			"password":     "",
			"errorMessage": errorMessage,
		})
	}

	res, err := services.AuthenticateUser(u, h.server.Config.JwtSecret)

	if err != nil {
		errorMessage = err.Error()
		println(errorMessage)
		return c.Render(http.StatusOK, "sign-in.html", map[any]interface{}{
			"name":         "Login",
			"currentTime":  tick,
			"email":        email,
			"password":     "",
			"errorMessage": errorMessage,
		})
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

func (h *AuthHandler) GetSignOut(c echo.Context) error {
	user, _ := services2.GetCurrentUser(c, *h.server)
	if user != nil {
		//logout
		c.SetCookie(&http.Cookie{
			Name:    "token",
			Value:   "",
			Expires: time.Now().Add(7 * time.Hour),
			Path:    "/",
		})

	}
	return c.Redirect(http.StatusMovedPermanently, "/")
}
