package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Receipt represents a store receipt with retailer name, date, time, total, and items purchased.
type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Total        string `json:"total"`
	Items        []Item `json:"items"`
}

// Item represents a purchased item with a short description and price.
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

// ResponseID is used to return the generated receipt ID.
type ResponseID struct {
	ID string `json:"id"`
}

// ResponsePoints is used to return the awarded points for a receipt.
type ResponsePoints struct {
	Points int `json:"points"`
}

// Global maps to store receipts and their corresponding points (in-memory storage).
var receiptStore = make(map[string]Receipt)
var pointsStore = make(map[string]int)

// processReceipt handles the POST request to store receipts and generate an ID.
func processReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt

	// Decode the incoming JSON request into a Receipt struct.
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Generate a unique ID for the receipt.
	receiptID := uuid.New().String()

	// Store the receipt in memory.
	receiptStore[receiptID] = receipt

	// Calculate points and store them.
	pointsStore[receiptID] = calculatePoints(receipt)

	// Return the generated ID as a JSON response.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ResponseID{ID: receiptID})
}

// getPoints handles the GET request to retrieve points for a given receipt ID.
func getPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	receiptID := vars["id"]

	// Check if the receipt ID exists.
	if points, exists := pointsStore[receiptID]; exists {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ResponsePoints{Points: points})
	} else {
		http.Error(w, "No receipt found for that ID", http.StatusNotFound)
	}
}

// calculatePoints applies the given rules to calculate points for a receipt.
func calculatePoints(receipt Receipt) int {
	totalPoints := 0

	// Rule 1: One point per alphanumeric character in the retailer name.
	alnumRegex := regexp.MustCompile("[a-zA-Z0-9]")
	totalPoints += len(alnumRegex.FindAllString(receipt.Retailer, -1))

	// Convert total to float for calculations.
	totalAmount, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return totalPoints
	}

	// Rule 2: 50 points if the total is a whole number (no cents).
	if totalAmount == float64(int(totalAmount)) {
		totalPoints += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25.
	if math.Mod(totalAmount, 0.25) == 0 {
		totalPoints += 25
	}

	// Rule 4: 5 points for every two items on the receipt.
	totalPoints += (len(receipt.Items) / 2) * 5

	// Rule 5: If item description length is a multiple of 3, apply the price-based rule.
	for _, item := range receipt.Items {
		trimmedDesc := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDesc)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			totalPoints += int(math.Ceil(price * 0.2))
		}
	}

	// Rule 6: 6 points if the purchase day is an odd number.
	purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err == nil && purchaseDate.Day()%2 != 0 {
		totalPoints += 6
	}

	// Rule 7: 10 points if the purchase time is between 2:00 PM and 4:00 PM.
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err == nil && purchaseTime.Hour() == 14 {
		totalPoints += 10
	}

	return totalPoints
}

// main function initializes the HTTP server and routes.
func main() {
	r := mux.NewRouter()

	// Define API routes.
	r.HandleFunc("/receipts/process", processReceipt).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", getPoints).Methods("GET")

	// Start the server on port 8080.
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
