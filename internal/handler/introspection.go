package handler

import "net/http"

type IntrospectionHandler struct {
	Handler
}

type IntrospectionResponse struct {
	Active bool `json:"active"`
}

func NewIntrospectionHandler() *Handler {
	mux := http.NewServeMux()

	mux.Handle("/token/introspect", handleIntrospection())

	return &Handler{
		mux: mux,
	}
}

func handleIntrospection() http.Handler {
	resp := IntrospectionResponse{
		Active: false,
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encode(w, r, http.StatusOK, resp)
	})
}
