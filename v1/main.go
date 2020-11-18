package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
)

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

type Play struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// Though it looks ugly, I follow the JSON format the book described
type Plays struct {
	Hamlet struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"hamlet"`
	Aslike struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"aslike"`
	Othello struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"othello"`
}

type Performance struct {
	PlayID   string `json:"playID"`
	Audience int    `json:"audience"`
}

type Invoice struct {
	Customer     string `json:"customer"`
	Performances []struct {
		PlayID   string `json:"playID"`
		Audience int    `json:"audience"`
	} `json:"performances"`
}

func getPlayStruct(plays Plays, perf Performance) Play {
	switch perf.PlayID {
	case "hamlet":
		return plays.Hamlet

	case "aslike":
		return plays.Aslike

	case "othello":
		return plays.Othello

	default:
		return Play{}
	}
}

func statement(playsJSON, invoiceJSON string) string {
	var (
		totalAmount   int
		volumeCredits float64
		result        = `Statement for ${invoice.customer}\n`
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

	fmt.Printf("plays:\n%+v\n", plays)
	fmt.Printf("invoice:\n%+v\n", invoices)

	// const format = new Intl.NumberFormat(
	// 	"enUS",
	// 	{
	// 		style: "currency",
	// 		currency: "USD",
	// 		minimumFractionDigits: 2
	// 	}
	// ).format

	// for each invoice performance
	for _, perf := range invoices[0].Performances {

		play := getPlayStruct(plays, perf) // return Play struct

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
		volumeCredits += math.Max(float64(perf.Audience-30), 0)
		// add extra credit for every ten comedy attendees
		if "comedy" == play.Type {
			volumeCredits += math.Floor(float64(perf.Audience))
		}
		// print line for this order
		result += ` ${play.name}: ${format(thisAmount/100)} (${perf.audience} seats)`
		totalAmount += thisAmount
	}
	result += `Amount owed is ${format(totalAmount/100)}\n`
	result += `You earned ${volumeCredits} credits\n`
	return result

}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println(statement(playsJSON, invoiceJSON))
}
