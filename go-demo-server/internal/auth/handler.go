package auth

import (
	"goadvancedserver/configs"
	"goadvancedserver/pkg/jwt"
	"goadvancedserver/pkg/request"
	"goadvancedserver/pkg/response"
	"net/http"
)

type HandlerDependencies struct {
	*configs.Config
	*Service
}

type Handler struct {
	*configs.Config
	*Service
}

func NewAuthHandler(router *http.ServeMux, deps HandlerDependencies) {
	handler := Handler{
		Config:  deps.Config,
		Service: deps.Service,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *Handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.Json[LoginRequest](r)
		if err != nil {
			response.Error(w, err, http.StatusForbidden)
			return
		}
		email, err := handler.Service.Login(body.Email, body.Password)
		if err != nil {
			response.Error(w, err, http.StatusUnauthorized)
			return
		}
		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(email)
		if err != nil {
			response.Error(w, err, http.StatusInternalServerError)
			return
		}
		resp := LoginResponse{
			Token: token,
		}
		response.Json(w, resp, http.StatusOK)
	}
}

func (handler *Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.Json[RegisterRequest](r)
		if err != nil {
			response.Error(w, err, http.StatusBadRequest)
			return
		}
		email, err := handler.Service.Register(body.Email, body.Password, body.Name)
		if err != nil {
			response.Error(w, err, http.StatusUnauthorized)
			return
		}
		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(email)
		if err != nil {
			response.Error(w, err, http.StatusInternalServerError)
			return
		}
		resp := RegisterResponse{
			Token: token,
		}
		response.Json(w, resp, http.StatusOK)
	}
}
