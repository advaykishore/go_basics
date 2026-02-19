package db

import (
	"context"
	"log"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"

	"github.com/joho/godotenv"
)

var Driver neo4j.DriverWithContext

func InitNeo4j() {
	err1 := godotenv.Load() // loads .env file
	if err1 != nil {
		log.Fatal("Error loading .env file")
	}
	var err error
	pass := os.Getenv("PASS")

	Driver, err = neo4j.NewDriverWithContext(
		"neo4j://localhost:7687",
		neo4j.BasicAuth("neo4j", pass, ""),
	)

	if err != nil {
		log.Fatal("Failed to create driver:", err)
	}

	err = Driver.VerifyConnectivity(context.Background())
	if err != nil {
		log.Fatal("Neo4j not reachable:", err)
	}

	log.Println("Connected to Neo4j ")
}
