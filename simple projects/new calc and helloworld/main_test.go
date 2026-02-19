package main

import (
	"encoding/json"
	"strings"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}
func TestSingleCalculator(t *testing.T) {
	router := SetUpRouter()
	router.POST("/calculator", calculator)

	w := httptest.NewRecorder()

	input := calc{
		Number1:  20,
		Number2:  30,
		Operator: "*",
	}
	userJson, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/calculator", strings.NewReader(string(userJson)))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	expected := result{Result: 600}
	expectedJSON, _ := json.Marshal(expected)

	assert.JSONEq(t, string(expectedJSON), w.Body.String())
}
func TestMultipleNumbers(t *testing.T) {
	router := SetUpRouter()
	router.POST("/readnum", readnum)

	input := []number{
		{Number1: 20, Number2: 30, Number3: 50},
		{Number1: 20, Number2: 40, Number3: 80},
		{Number1: 10, Number2: 30, Number3: 50},
		{Number1: 20, Number2: 10, Number3: 50},
		{Number1: 20, Number2: 30, Number3: 10},
	}

	for i, tc := range input {
		w := httptest.NewRecorder()

		userJson, _ := json.Marshal(tc)
		req, _ := http.NewRequest("POST", "/readnum", strings.NewReader(string(userJson)))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)

		expected := []number{
			{Number1: 20, Number2: 30, Number3: 50},
			{Number1: 20, Number2: 40, Number3: 80},
			{Number1: 10, Number2: 30, Number3: 50},
			{Number1: 20, Number2: 10, Number3: 50},
			{Number1: 20, Number2: 30, Number3: 10},
		}
		expectedJSON, _ := json.Marshal(expected[i])

		assert.JSONEq(t, string(expectedJSON), w.Body.String())
	}
}
func TestMultipleCalculator(t *testing.T) {
	router := SetUpRouter()
	router.POST("/calculator", calculator)

	input := []calc{
		{Number1: 20,
			Number2:  30,
			Operator: "*"},
		{Number1: 20,
			Number2:  30,
			Operator: "+"},
	}
	for i, tc := range input {
		w := httptest.NewRecorder()
		userJson, _ := json.Marshal(tc)
		req, _ := http.NewRequest("POST", "/calculator", strings.NewReader(string(userJson)))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)

		expected := []result{{Result: 600}, {Result: 50}}
		expectedJSON, _ := json.Marshal(expected[i])

		assert.JSONEq(t, string(expectedJSON), w.Body.String())
	}
}
