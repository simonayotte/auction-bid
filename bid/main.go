package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Bid struct {
	Id     string `json:"id"`
	Bid    int    `json:"price"`
	Bidder string `json:"bidder"`
}

func (bs *BidService) HandleBidRequest(w http.ResponseWriter, r *http.Request) {

	// Read query params
	id := r.URL.Query().Get("id")

	log.Println("Request received for ad placement: ", id)

	bidValue := bs.PlaceBid()
	if bidValue == 0 || id == "" {
		w.WriteHeader(http.StatusNoContent)
	}

	bid := Bid{Id: id, Bid: bidValue, Bidder: bs.Id.String()}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bid)
}

func main() {
	bs := NewBidService()
	http.HandleFunc("/", bs.HandleBidRequest)
	log.Printf("Bid server %s up and running on port 8081...", bs.Id)
	http.ListenAndServe(":8081", nil)
}
