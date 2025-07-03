package handler

import "net/http"

type HealthHandler struct {
	Handler
}

type HealthResponse struct {
	Status string `json:"status"`
}

func NewHealthHandler() *Handler {
	mux := http.NewServeMux()

	mux.Handle("/liveness", handleLiveness())
	mux.Handle("/readiness", handleReadiness())

	return &Handler{
		mux: mux,
	}
}

func handleLiveness() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := &HealthResponse{
			Status: "ok",
		}

		encode(w, r, http.StatusOK, resp)
	})
}

func handleReadiness() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := &HealthResponse{
			Status: "ok",
		}

		encode(w, r, http.StatusOK, resp)
	})
}
