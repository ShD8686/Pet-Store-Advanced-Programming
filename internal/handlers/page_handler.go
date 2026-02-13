package handlers

import (
	"net/http"
)

type PageHandler struct{}

func NewPageHandler() *PageHandler {
	return &PageHandler{}
}

func (h *PageHandler) IndexPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "web/templates/index.html")
}

func (h *PageHandler) InfoPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/templates/info.html")
}

func (h *PageHandler) StatsPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/templates/stats.html")
}

func (h *PageHandler) CreateAdPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/templates/create-ad.html")
}

func (h *PageHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/templates/login.html")
}

func (h *PageHandler) RegisterPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/templates/register.html")
}

func (h *PageHandler) AdminPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/templates/admin.html")
}
