package main

import (
	"erfantak/goWebApiStructure/internal/Route"
	"erfantak/goWebApiStructure/internal/middleware"
	"fmt"
	"net/http"
)

func main() {
	privateRoute := http.NewServeMux()
	publicRoute := http.NewServeMux()

	Route.PrivateRouteDelegator(privateRoute)
	Route.PublicRouteDelegator(publicRoute)

	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server listening on http://localhost%s\n", addr)

	publicRoute.Handle("/", middleware.EnsureAdmin(privateRoute))

	stack := middleware.CreateStack(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    addr,
		Handler: stack(publicRoute),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
