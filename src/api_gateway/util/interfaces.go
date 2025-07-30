package util

import "github.com/go-chi/chi/v5"

type Closable interface {
	Close()
}

type ClosableRouter interface {
	chi.Router
	Closable
}
