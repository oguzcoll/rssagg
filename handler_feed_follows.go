package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"github.com/oguzcoll/rssagg/internal/database"
	"net/http"
	"time"
)

func (apiCfg *apiConfig) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Invalid parsing JSON: %s", err))
		return
	}
	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Cannot create feed follow: %s", err))
	}
	respondWithJson(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
}
