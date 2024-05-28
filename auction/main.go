package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

func (as *AuctionService) HandleNewAdAuction(w http.ResponseWriter, r *http.Request) {
	id := as.GenerateNewAdId()
	log.Printf("New Ad spot auction started for ad: %s", id)

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	bidServicesURL := []string{
		"http://bidding:8081",
		"http://bidding-2:8081",
		"http://bidding-3:8081",
	}

	var wg sync.WaitGroup
	for _, url := range bidServicesURL {
		wg.Add(1)
		go as.fetchBid(ctx, url, id.String(), &wg)
	}

	wg.Wait()
	maxBid := as.GetMaxBid()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(maxBid)
}

type Bid struct {
	Id     string `json:"id"`
	Bid    int    `json:"price"`
	Bidder string `json:"bidder"`
}

func (as *AuctionService) fetchBid(ctx context.Context, baseURL string, id string, wg *sync.WaitGroup) {
	log.Printf("Fetching bid for %s", baseURL)
	defer wg.Done()

	// Set Query params for bid
	url, _ := url.Parse(baseURL)
	query := url.Query()
	query.Set("id", id)
	url.RawQuery = query.Encode()

	// Create an HTTP bid request with the context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Printf("Error fetching bid to %s: %v", url.String(), err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return
	}

	// Read response from bidding service
	var bid Bid
	if err := json.NewDecoder(resp.Body).Decode(&bid); err != nil {
		log.Printf("Error reading response from %s: %v", url.String(), err)
		return
	}

	log.Printf("Received bid %d for %s from %s", bid.Bid, bid.Id, bid.Bidder)
	as.Bids[bid.Bidder] = bid.Bid
}

func main() {

	as := NewAuctionService()

	http.HandleFunc("/", as.HandleNewAdAuction)

	log.Println("Auction server up and running on port 8080...")

	http.ListenAndServe(":8080", nil)
}
