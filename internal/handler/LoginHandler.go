package handler

import (
	"fmt"
	"net/http"
)

type LoginHandler struct {
}

func (L LoginHandler) Login(writer http.ResponseWriter, request *http.Request) {
	data := "you are logged in"
	_, err := fmt.Fprintln(writer, data)
	if err != nil {
		return
	}
}
