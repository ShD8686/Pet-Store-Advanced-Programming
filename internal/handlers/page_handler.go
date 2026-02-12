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
	http.ServeFile(w, r, "index.html")
}

func (h *PageHandler) InfoPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "info.html")
}

func (h *PageHandler) StatsPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "stats.html")
}

func (h *PageHandler) CreateAdPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "create-ad.html")
}
