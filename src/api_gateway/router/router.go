package router

import (
	user_router "api-gateway/router/user"
	"net/http"

	"api-gateway/util"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	r           http.Handler
	sub_routers []util.ClosableRouter
}

func (r *Router) GetRouter() http.Handler {
	return r.r
}

func (r *Router) Close() {
	for _, s_r := range r.sub_routers {
		s_r.Close()
	}
}

func CreateRouter() *Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	sub_routers := make([]util.ClosableRouter, 0)

	sub_routers = append(sub_routers, user_router.CreateUserRouter())
	r.Mount("/user", sub_routers[0])

	return &Router{r: r, sub_routers: sub_routers}
}
