package cmd

import (
	"fmt"
	"github.com/caiorcferreira/swapi/internals/swapi/api"
	"net/http"
)

func StartServer(port int) error {
	router := api.SetupRoutes()

	s := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler:router,
	}

	return s.ListenAndServe()
}
