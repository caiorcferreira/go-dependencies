package cmd

import (
	"fmt"
	"github.com/caiorcferreira/swapi/internals/swapi/api"
	"net/http"
	"os"
)

func StartServer() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := api.SetupRoutes()

	s := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler:router,
	}

	return s.ListenAndServe()
}
