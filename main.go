package main

import (
	"encoding/json"
	"go-tokenomics/models"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

// Assuming you have a map to keep track of members and a mutex for concurrent access.
var (
	membersMap   = make(map[string]models.Member)
	membersMutex = &sync.Mutex{}
)

func main() {
	// Set up your routes
	http.HandleFunc("/members", AddMember)
	http.HandleFunc("/login", AuthMiddleware(Login))
	http.HandleFunc("/stake", AuthMiddleware(StakeTokens))
	http.HandleFunc("/scheduleGame", AuthMiddleware(ScheduleGame))

	// Start the server
	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}

func AddMember(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests for adding a member
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Unmarshal the body into a Member struct
	var member models.Member
	if err := json.Unmarshal(body, &member); err != nil {
		http.Error(w, "Error unmarshalling request body", http.StatusBadRequest)
		return
	}

	// Add the initial tokens to the member's CeptorWallet
	member.CeptorWallet = models.CeptorWallet{
		GamesXP:     0,
		ArtXP:       0,
		TechXP:      100,
		ArtTokens:   1,
		GamesTokens: 10,
		TechTokens:  20,
	}

	// Add member to the map, protected by a mutex
	membersMutex.Lock()
	membersMap[member.Username] = member
	membersMutex.Unlock()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Member added successfully"))
}

// Login handler function updated with mock login logic.
func Login(w http.ResponseWriter, r *http.Request) {
	// Mock login simply writes a successful login message.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Mock login successful"))
}

// StakeTokens handler function stub.
func StakeTokens(w http.ResponseWriter, r *http.Request) {
	// TODO: implement staking logic
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("StakeTokens not implemented"))
}

// ScheduleGame handler function stub.
func ScheduleGame(w http.ResponseWriter, r *http.Request) {
	// TODO: implement game scheduling logic
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ScheduleGame not implemented"))
}

// AuthMiddleware function updated with simple authentication logic.
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the secret phrase and number from headers
		secretPhrase := r.Header.Get("X-Secret-Phrase")
		secretNumber := r.Header.Get("X-Secret-Number")

		// Check if the secret phrase and number match your pre-defined values
		if secretPhrase == "HootyTooty" && secretNumber == "42" {
			// If they match, call the next handler
			next(w, r)
		} else {
			// If they don't match, return an HTTP 403 Forbidden status
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Forbidden: Incorrect secret phrase or number"))
		}
	}
}
