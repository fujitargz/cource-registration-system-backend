package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, cource-registration-system-backend")
}

type HelloJSONHandler struct{}

func (h *HelloJSONHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v := struct {
		Msg string `json:"msg"`
	}{Msg: "hello"}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println("HelloJSONHandler: ", err)
	}
}
