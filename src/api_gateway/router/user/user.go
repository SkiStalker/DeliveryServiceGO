package user_router

import (
	"api-gateway/handlers/user"
	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	chi.Router
	u_h *user_handlers.UserHandler
}

func (u_r UserRouter) Close() {
	u_r.u_h.Close()
}

func CreateUserRouter() *UserRouter {
	u_r := &UserRouter{Router: chi.NewRouter(), u_h: user_handlers.CreateUserHandler()}

	u_r.Get("/{user_id}", u_r.u_h.GetUser)
	return u_r
}
