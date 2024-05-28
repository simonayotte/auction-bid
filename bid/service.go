package main

import (
	"math/rand"

	"github.com/google/uuid"
)

type BidService struct {
	Id uuid.UUID
}

func NewBidService() *BidService {
	id, _ := uuid.NewUUID()
	return &BidService{Id: id}
}

func (bs *BidService) PlaceBid() int {
	//  1/10 chance of not bidding on an ad
	n := rand.Intn(10) + 1
	if n != 1 {
		return rand.Intn(1000)
	}

	return 0
}
