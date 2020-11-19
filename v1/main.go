package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
)

// reference for code: https://feel5ny.github.io/2020/04/04/Refactoring_002/

var (
	playsJSON string = `{
		"hamlet":  {"name": "Hamlet", "type": "tragedy"},
		"aslike":  {"name": "As You Like It", "type": "comedy"},
		"othello": {"name": "Othello", "type": "tragedy"}
	}`

	invoiceJSON string = `[
		{
			"customer": "BigCo",
			"performances": [
				{
					"playID": "hamlet",
					"audience": 55
				},
				{
					"playID": "aslike",
					"audience": 35
				},
				{
					"playID": "othello",
					"audience": 40
				}
			]
		}
	]`
)

// Play is a struct with Name and Type of the play
type Play struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// Plays is a map of the play
type Plays map[string]Play

// Performance has PlayID and the amount of the Audience
type Performance struct {
	PlayID   string `json:"playID"`
	Audience int    `json:"audience"`
}

// Invoice shows the customer name and performances info
type Invoice struct {
	Customer     string `json:"customer"`
	Performances []struct {
		PlayID   string `json:"playID"`
		Audience int    `json:"audience"`
	} `json:"performances"`
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func statement(playsJSON, invoiceJSON string) string {
	var (
		totalAmount   int
		volumeCredits int
		result        string
	)

	// unmarshal playsJSON, invoiceJSON
	var plays Plays
	var invoices []Invoice

	if err := json.Unmarshal([]byte(playsJSON), &plays); err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal([]byte(invoiceJSON), &invoices); err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("plays:\n%+v\n", plays)
	// fmt.Printf("invoice:\n%+v\n", invoices)

	result = fmt.Sprintf("Statement for %s\n", invoices[0].Customer)

	// not easy in Go
	/*
		const format = new Intl.NumberFormat(
			"enUS",
			{
				style: "currency",
				currency: "USD",
				minimumFractionDigits: 2
			}
		).format
	*/
	// for each invoice performance
	for _, perf := range invoices[0].Performances {

		// play := getPlayStruct(plays, perf) // return Play struct
		play, found := plays[perf.PlayID] // return Play struct
		if !found {
			log.Printf("not found the information of the %s", perf.PlayID)
			continue
		}

		var thisAmount int
		switch play.Type {
		case "tragedy":
			thisAmount = 40000
			if perf.Audience > 30 {
				thisAmount += 1000 * (perf.Audience - 30)
			}

		case "comedy":
			thisAmount = 30000
			if perf.Audience > 20 {
				thisAmount += 10000 + 500*(perf.Audience-20)
			}
			thisAmount += 300 * perf.Audience

		default:
			log.Printf("`unknown type: %v", play.Type)
		}

		// add volume credits
		volumeCredits += maxInt((perf.Audience - 30), 0)

		// add extra credit for every ten comedy attendees
		if "comedy" == play.Type {
			volumeCredits += int(math.Floor(float64(perf.Audience) / 5))
		}

		// print line for this order
		result += fmt.Sprintf("  %s: $%.2f (%d seats)\n", play.Name, float64(thisAmount/100), perf.Audience)
		totalAmount += thisAmount
	}
	result += fmt.Sprintf("Amount owed is $%.2f\n", float64(totalAmount/100))
	result += fmt.Sprintf("You earned %d credits\n", volumeCredits)
	return result
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println(statement(playsJSON, invoiceJSON))
}

/*
Statement for BigCo
  Hamlet: $650.00 (55 seats)
  As You Like It: $580.00 (35 seats)
  Othello: $500.00 (40 seats)
Amount owed is $1730.00
You earned 47 credits

Statement for BigCo
	Hamlet: $650.00 (55 seats)
	As You Like It: $580.00 (35 seats)
	Othello: $500.00 (40 seats)
Amount owed is $1,730.00
You earned 47 credits
*/
