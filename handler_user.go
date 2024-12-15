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

func (apiCfg *apiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Invalid parsing JSON: %s", err))
		return
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Cannot create user: %s", err))
	}
	respondWithJson(w, 200, databaseUserToUser(user))
}
