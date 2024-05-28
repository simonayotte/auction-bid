package main

import (
	"github.com/google/uuid"
)

type AdId uuid.UUID

// Mapping bidding service -> bid
type AuctionService struct {
	ActiveAdId AdId
	Bids       map[string]int
}

func NewAuctionService() *AuctionService {
	bids := make(map[string]int)
	id, _ := uuid.NewUUID()
	return &AuctionService{ActiveAdId: AdId(id), Bids: bids}
}

func (as *AuctionService) GenerateNewAdId() uuid.UUID {
	id, _ := uuid.NewUUID()
	return id
}

func (as *AuctionService) GetMaxBid() int {
	var max int
	for _, v := range as.Bids {
		if v > max {
			max = v
		}
	}
	return max
}
