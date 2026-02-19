package main

import (
	"encoding/json"
	"net/http"
)

type Calcrequest struct {
	Number1  float64 `json:"number1"`
	Number2  float64 `json:"number2"`
	Operator string  `json:"operator"`
}
type Calcresponse struct {
	Result float64 `json:"result"`
}

func handleinput(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only post method is allowed ", http.StatusMethodNotAllowed)
		return
	}
	var req Calcrequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid input ", http.StatusBadRequest)
		return
	}
	var result float64
	switch req.Operator {
	case "+":
		result = req.Number1 + req.Number2
	case "-":
		result = req.Number1 - req.Number2
	case "*":
		result = req.Number1 * req.Number2
	case "/":
		if req.Number2 == 0 {
			http.Error(w, "No division by zero(zero division error) ", http.StatusBadRequest)
			return
		}
		result = req.Number1 / req.Number2

	default:
		http.Error(w, "Invalid operation", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Calcresponse{Result: result})

}
func main() {

	http.HandleFunc("/", handleinput)
	http.ListenAndServe(":8080", nil)

}
