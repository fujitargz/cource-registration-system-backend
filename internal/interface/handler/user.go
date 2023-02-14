package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/fujitargz/cource-registration-system-backend/internal/usecase"
)

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	FindUserById(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	u usecase.UserUsecase
}

func NewUserHandler(u usecase.UserUsecase) *userHandler {
	return &userHandler{u}
}

func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	input := &struct {
		ID       string `json:"id"`
		Password string `json:"password"`
		IsAdmin  bool   `json:"is_admin"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(input); err != nil {
		log.Println("CreateUser: ", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if err := h.u.Create(input.ID, input.Password, input.IsAdmin); err != nil {
		log.Println("CreateUser: ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	if len(path) != 2 {
		http.NotFound(w, r)
		return
	}
	ID := path[1]
	if err := h.u.Delete(ID); err != nil {
		log.Println("DeleteUser: ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *userHandler) FindUserById(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	if len(path) != 2 {
		http.NotFound(w, r)
		return
	}
	ID := path[1]
	user, err := h.u.FindByID(ID)
	if err != nil {
		log.Println("FindUserById: ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Println("FindUserById: ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}
