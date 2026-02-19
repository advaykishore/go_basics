package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

type answer struct {
	Name    string `json:"name"`
	Project string `json:"project"`
}

// -------- Interface --------
type UserRepo interface {
	GetUsers(ctx context.Context) ([]answer, error)
}

// -------- Neo4j Repo --------
type Neo4jRepo struct {
	driver neo4j.DriverWithContext
}

func NewNeo4jRepo(uri, user, pass string) (*Neo4jRepo, error) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(user, pass, ""))
	if err != nil {
		return nil, err
	}
	return &Neo4jRepo{driver: driver}, nil
}

func (r *Neo4jRepo) GetUsers(ctx context.Context) ([]answer, error) {
	result, err := neo4j.ExecuteQuery(ctx, r.driver, `
		MATCH (U:User)-[:WORKS_ON]->(P:Project)
		RETURN U.name AS name, P.name as project
	`, nil, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}

	var users []answer
	for _, record := range result.Records {
		n, _ := record.Get("name")
		p, _ := record.Get("project")
		users = append(users, answer{Name: n.(string), Project: p.(string)})
	}
	return users, nil
}

// -------- Handler --------
func readuser(repo UserRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := repo.GetUsers(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, users)
	}
}

func main() {
	r := gin.Default()
	err := godotenv.Load() // loads .env file
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	pass := os.Getenv("PASS")
	fmt.Println("this is the password", pass)
	repo, _ := NewNeo4jRepo("bolt://localhost:7687", "neo4j", pass)
	r.GET("/neo4j", readuser(repo))
	r.Run()
}
