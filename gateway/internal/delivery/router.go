package delivery

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter(h *HttpHandler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/repo", h.GetRepository)
	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	return mux
}
