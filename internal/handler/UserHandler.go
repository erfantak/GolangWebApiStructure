package handler

import (
	"encoding/json"
	"erfantak/goWebApiStructure/internal/Service"
	"fmt"
	"net/http"
)

type UserHandler struct {
}

func (U UserHandler) GetUser(writer http.ResponseWriter, request *http.Request) {
	id := request.PathValue("id")
	data := Service.UserService{}.GetUser(id)
	_, err := fmt.Fprintln(writer, data)
	if err != nil {
		return
	}
}

func (U UserHandler) StoreUser(writer http.ResponseWriter, request *http.Request) {
	Service.UserService{}.StoreUser()
	_, err := fmt.Fprintln(writer, "store user successfully")
	if err != nil {
		return
	}
}

func (U UserHandler) StoreManyUser(writer http.ResponseWriter, request *http.Request) {
	res := Service.UserService{}.StoreUsers()
	j, err := json.Marshal(res)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err2 := writer.Write(j)
	if err2 != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	//_, err := fmt.Fprintln(writer, res["companyl"])
	//if err != nil {
	//	return
	//}
}
