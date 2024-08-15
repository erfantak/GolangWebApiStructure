package route

import (
	"erfantak/goWebApiStructure/internal/handler"
	"net/http"
)

func PublicRouteDelegator(router *http.ServeMux) {
	// ===================== PUBLIC ROUTES =====================
	router.HandleFunc("POST /login", handler.LoginHandler{}.Login)

}

func PrivateRouteDelegator(router *http.ServeMux) {
	// ===================== PRIVATE ROUTES ====================
	router.HandleFunc("GET /user", handler.UserHandler{}.GetUser)

}
