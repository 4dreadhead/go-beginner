package link

import (
	"fmt"
	"goadvancedserver/configs"
	"goadvancedserver/pkg/middleware"
	"goadvancedserver/pkg/request"
	"goadvancedserver/pkg/response"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type HandlerDeps struct {
	*configs.Config
	Repository *Repository
}

type Handler struct {
	*configs.Config
	Repository *Repository
}

func NewLinkHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := Handler{
		Config:     deps.Config,
		Repository: deps.Repository,
	}
	router.HandleFunc("POST /links/create", handler.Create())
	router.Handle("POST /links/{id}", middleware.Authorize(handler.Update()))
	router.HandleFunc("DELETE /links/{id}", handler.Delete())
	router.HandleFunc("GET /s/{hash}", handler.RedirectTo())
}

const maxHashTries = 10

func (handler *Handler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.Json[CreateRequest](r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		link := NewLink(body.Url)

		for i := range maxHashTries {
			existedLink, err := handler.Repository.GetByHash(link.Hash)
			if existedLink == nil {
				fmt.Println(err.Error())
				break
			}
			if i+1 >= maxHashTries {
				http.Error(w, "Creation Unavailable", http.StatusConflict)
				break
			}
			link.GenerateHash()
		}

		savedLink, err := handler.Repository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.Json(w, savedLink, http.StatusCreated)
	}
}

func (handler *Handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.Json[UpdateRequest](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		idStr := r.PathValue("id")
		id, err := strconv.ParseUint(idStr, 10, 32)

		if err != nil {
			http.Error(w, "Invalid id", http.StatusBadRequest)
			return
		}
		link, err := handler.Repository.Update(&Link{
			Model: gorm.Model{ID: uint(id)},
			Url:   body.Url,
			Hash:  body.Hash,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.Json(w, link, http.StatusOK)
	}
}

func (handler *Handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			http.Error(w, "Invalid id", http.StatusBadRequest)
			return
		}
		idUint := uint(id)
		_, err = handler.Repository.Find(idUint)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		err = handler.Repository.Delete(idUint)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Json(w, nil, http.StatusOK)
	}
}

func (handler *Handler) RedirectTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")

		link, err := handler.Repository.GetByHash(hash)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		http.Redirect(w, r, link.Url, http.StatusFound)
	}
}
