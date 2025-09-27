package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalManager []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalManager: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globalManager = append(mngr.globalManager, middlewares...)
}

func (mngr *Manager) WrapMux(mux http.Handler, middlewareList ...Middleware) http.Handler {
	h := mux

	for _, middleware := range middlewareList {
		h = middleware(h)
	}
	return h
}
