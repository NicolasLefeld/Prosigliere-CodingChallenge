package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"prosigliere-coding-challenge/models"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func GetAllPosts(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var posts []models.BlogPost
		if err := db.Preload("Comments").Find(&posts).Error; err != nil {
			http.Error(w, "Failed to fetch posts", 500)
			return
		}

		type PostWithCount struct {
			ID           uint   `json:"id"`
			Title        string `json:"title"`
			CommentCount int    `json:"comment_count"`
		}
		var response []PostWithCount
		for _, p := range posts {
			response = append(response, PostWithCount{
				ID:           p.ID,
				Title:        p.Title,
				CommentCount: len(p.Comments),
			})
		}
		json.NewEncoder(w).Encode(response)
	}
}

func CreatePost(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var post models.BlogPost
		if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
			http.Error(w, "Invalid body", 400)
			return
		}
		if err := db.Create(&post).Error; err != nil {
			http.Error(w, "Failed to create post", 500)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(post)
	}
}

func GetPostByID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", 400)
			return
		}
		var post models.BlogPost
		if err := db.Preload("Comments").First(&post, id).Error; err != nil {
			http.Error(w, "Post not found", 404)
			return
		}
		json.NewEncoder(w).Encode(post)
	}
}
