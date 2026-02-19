package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"web-go/db"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	session := db.Driver.NewSession(ctx, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})
	defer session.Close(ctx)

	result, err := session.Run(
		ctx,
		"MATCH (u:User) RETURN id(u) AS id, u.name AS name",
		nil,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	users := []User{}

	for result.Next(ctx) {
		record := result.Record()

		idVal, ok := record.Get("id")
		if !ok || idVal == nil {
			continue
		}

		nameVal, ok := record.Get("name")
		if !ok || nameVal == nil {
			continue
		}

		id, ok := idVal.(int64)
		if !ok {
			continue
		}

		name, ok := nameVal.(string)
		if !ok {
			continue
		}

		users = append(users, User{
			ID:   id,
			Name: name,
		})
	}

	// check for iteration errors
	if err = result.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
