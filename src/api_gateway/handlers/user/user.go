package user_handlers

import (
	"api-gateway/clients/user"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	user_clt user_client.UserServiceClient
}

func CreateUserHandler() *UserHandler {
	return &UserHandler{user_clt: *user_client.CreateUserServiceClient()}
}

func (u_h UserHandler) Close() {
	u_h.user_clt.Close()
}

func (u_h UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	if user_id := chi.URLParam(r, "user_id"); user_id != "" {
		resp, err := u_h.user_clt.GetUser(r.Context(), user_id)
		if err != nil {
			if grpc_err, ok := status.FromError(err); ok && grpc_err.Code() == codes.NotFound {
				http.Error(w, grpc_err.Message(), http.StatusNotFound)
			} else {
				http.Error(w, fmt.Sprintf("Internal error : %v", err), http.StatusInternalServerError)
			}
		} else {
			json.NewEncoder(w).Encode(resp)
		}
	} else {
		http.Error(w, "user_id url parameter doesn't specified", http.StatusBadRequest)
	}
}
