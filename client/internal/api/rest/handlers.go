package rest

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"html/template"
	"net/http"
)

type ClientHandler struct {
	logger    *zerolog.Logger
	templates *template.Template
	// clientService api.CleintService
}

func NewClientHandler(
	logger *zerolog.Logger,
) (*ClientHandler, error) {
	templates, err := template.ParseGlob("static/templates/*.html")
	if err != nil {
		return nil, fmt.Errorf("[NewClientHandler] parse templates: %w]")
	}
	return &ClientHandler{
		logger:    logger,
		templates: templates,
	}, nil
}

func (c *ClientHandler) MainPage(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()

	// Проверка авторизации (пока заглушка)
	isAuthenticated := true

	if !isAuthenticated {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	err := c.templates.ExecuteTemplate(w, "main.html", nil)
	if err != nil {
		c.logger.Error().Err(err).Msg("Failed to render main page")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (c *ClientHandler) Login(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()

	err := c.templates.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		c.logger.Error().Err(err).Msg("Failed to render login page")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (c *ClientHandler) Register(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()

	// Рендеринг страницы регистрации
	err := c.templates.ExecuteTemplate(w, "register.html", nil)
	if err != nil {
		c.logger.Error().Err(err).Msg("Failed to render register page")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (c *ClientHandler) Check(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()

	// Заглушка для проверки авторизации
	isAuthenticated := true

	response := map[string]bool{
		"authenticated": isAuthenticated,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
