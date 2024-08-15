package handler

import (
	"fmt"
	"net/http"
)

type UserHandler struct {
}

func (U UserHandler) GetUser(writer http.ResponseWriter, request *http.Request) {
	data := "User Data"
	_, err := fmt.Fprintln(writer, data)
	if err != nil {
		return
	}
}
