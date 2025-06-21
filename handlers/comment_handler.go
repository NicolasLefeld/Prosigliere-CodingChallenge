package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"prosigliere-coding-challenge/models"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func CreateComment(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		postID, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid post ID", 400)
			return
		}

		var comment models.Comment
		if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
			http.Error(w, "Invalid body", 400)
			return
		}
		comment.PostID = uint(postID)

		if err := db.Create(&comment).Error; err != nil {
			http.Error(w, "Failed to create comment", 500)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(comment)
	}
}
