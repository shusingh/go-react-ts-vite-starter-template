package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	// replace with your actual module path
	"backend/internal/db"
)

// NewRouter sets up routes and middleware.
func NewRouter(q *db.Queries) *chi.Mux {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    // Health check
    r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"status":"ok"}`))
    })

    // Example user endpoint
    r.Get("/users/{id}", getUserHandler(q))

    return r
}

// getUserHandler returns a handler that fetches a user by ID.
func getUserHandler(q *db.Queries) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        idParam := chi.URLParam(r, "id")
        id, err := strconv.ParseInt(idParam, 10, 32)
        if err != nil {
            http.Error(w, "invalid user id", http.StatusBadRequest)
            return
        }

        user, err := q.GetUser(r.Context(), int32(id))
        if err != nil {
            http.Error(w, "user not found", http.StatusNotFound)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(user)
    }
}
