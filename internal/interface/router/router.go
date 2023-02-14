package router

import (
	"net/http"

	"github.com/fujitargz/cource-registration-system-backend/internal/interface/handler"
)

func NewRouter(h handler.AppHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			panic("not implemented")
		case "POST":
			h.CreateUser(w, r)
		}
	})
	mux.Handle("/users/", http.StripPrefix("/users", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			h.FindUserById(w, r)
		case "PATCH":
			panic("not implemented")
		case "DELETE":
			h.DeleteUser(w, r)
		}
	})))
	return mux
}
