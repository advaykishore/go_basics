package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// -------- MOCK REPO --------
type MockRepo struct{}

func (m *MockRepo) GetUsers(ctx context.Context) ([]answer, error) {
	return []answer{
		{Name: "Advay", Project: "infragraph"},
	}, nil
}

func TestReadUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := &MockRepo{}
	router := gin.Default()
	router.GET("/neo4j", readuser(mockRepo))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/neo4j", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `[{"name":"Advay","project":"infragraph"}]`, w.Body.String())
}
