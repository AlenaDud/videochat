package rest

import (
	"client/config"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"net/http"
)

func NewRestApi(cfg *config.Config, logger *zerolog.Logger) error {
	clientRestHandler, err := NewClientHandler(logger)
	if err != nil {
		return fmt.Errorf("[NewRestApi] create handler: %w]")
	}

	router := mux.NewRouter()

	router.HandleFunc("/", clientRestHandler.MainPage).Methods(http.MethodGet)
	router.HandleFunc("/check-auth", clientRestHandler.Check).Methods(http.MethodPost)
	router.HandleFunc("/login", clientRestHandler.Login).Methods(http.MethodPost)
	router.HandleFunc("/register", clientRestHandler.Register).Methods(http.MethodPost)

	appAddr := fmt.Sprintf("%s:%s", cfg.REST.RESTHost, cfg.REST.RESTPort)
	logger.Info().Msgf("running REST server at '%s'", appAddr)
	if err := http.ListenAndServe(appAddr, router); err != nil {
		return fmt.Errorf("[NewRestApi] listen and serve: %w", err)
	}

	return nil

}
