package routes

import (
	"net/http"

	"prosigliere-coding-challenge/handlers"

	auth "prosigliere-coding-challenge/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(auth.APIKeyAuth)

	r.Route("/api/posts", func(r chi.Router) {
		r.Get("/", handlers.GetAllPosts(db))
		r.Post("/", handlers.CreatePost(db))
		r.Get("/{id}", handlers.GetPostByID(db))
		r.Post("/{id}/comments", handlers.CreateComment(db))
	})

	return r
}
