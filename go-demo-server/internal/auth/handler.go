package auth

import (
	"goadvancedserver/configs"
	"goadvancedserver/pkg/request"
	"goadvancedserver/pkg/response"
	"net/http"
)

type AuthHandlerDependencies struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDependencies) {
	handler := AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := request.Json[LoginRequest](r)
		if err != nil {
			response.Error(w, err, http.StatusForbidden)
			return
		}
		resp := LoginResponse{
			Token: "123",
		}
		response.Json(w, resp, http.StatusOK)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := request.Json[RegisterRequest](r)
		if err != nil {
			response.Error(w, err, http.StatusBadRequest)
			return
		}
		resp := RegisterResponse{
			Token: "123",
		}
		response.Json(w, resp, http.StatusCreated)
	}
}
