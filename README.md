# Auction-Bid Microservices system

## Descripition

This project is an auction-bid microservices system written in Go. The auction service gets all the bids from the bidding service and return back the bid with the best price.

## Installation

`make run` the project to run.

## Usage

Run the following command to make a request to the auction server:
`curl http://localhost:8000/`

## Test

Run the test with `make test`

## Requirements

Design two services (Bidding & Auction) one which is bidding on an adrequest and another which is
performing the auction from multiple bidders.

### Bidding Service

1. Receives an AdRequest i.e http request for an ad object with an AdPlacementid, a unique string
   identifying an Ad Slot or Ad Spot.
2. Every AdRequest should be responded with an AdObject that should contain at least an AdID (to
   uniquely identify an ad) and bidptice (random price for the AdPlacementid in USD).
3. Incase the bidding service does not want to buy an ad slot, the service call should return a 204, this
   can also be in random order
4. Incase the service bids for an AdRequest, it should return the AdObject on a 200 status code.

### Auction Service

1. The auction service shall call multiple bidding services at the same time (header bidding)
2. The auction service gets all the bids available from associated Bidding Services as 200 status
   code responses for valid bids.
3. Auction service should accept an AdPlacementid in an externally exposed API.
4. The auction service selects the bid for an Ad Placement with the highest bidprice amongst
   different services
5. Incase no eligible bids by the bidding service in an auction, the auction service returns a 204 status
   code
6. The auction service should have a safety circuit to prevent bad bidding services from increasing
   latencies
7. If a bidding service does not respond within 200ms, the auction service should not accept the bid
   for auction from that bidding service

Your code will be tested by making a call to an auction API exposed out. Since the circuit has been
applied, the exposed API should always respond within 200ms
You will be tested on code simplicity, design and concurrency. Write relevant test cases. Think about
scalability, concurrency and functionality while running designing your code base. Create a docker
compose file to get the services up and running.
