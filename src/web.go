package src

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// SolveRequest Define the request structure
type SolveRequest struct {
	Grid       [][]int `json:"grid"`
	Colors     []int   `json:"colors"`
	Iterations int     `json:"iterations"`
}

// SolveResponse Define the response structure
type SolveResponse struct {
	MinRounds int `json:"minRounds"`
}

// SolveGridHandler HTTP handler function to solve the grid with CORS enabled
func SolveGridHandler(w http.ResponseWriter, r *http.Request) {

	var request SolveRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Failed to parse request", http.StatusBadRequest)
		return
	}

	// Call the function to solve the grid
	minRounds := findMinRoundsWithGrowth(request.Grid, request.Colors, request.Iterations)

	// Return the result as JSON
	response := SolveResponse{
		MinRounds: minRounds,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// RegisterRoutes registers the src-specific routes to the provided router
func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/solve", SolveGridHandler).Methods("POST")
}
