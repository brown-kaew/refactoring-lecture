package main

import (
	"fmt"
	"math"
)

type PlayDetails struct {
	Name string
	Type string
}

type Play map[string]PlayDetails

type Performance struct {
	PlayID   string `json:"playID"`
	Audience int    `json:"audience"`
}

type Invoice struct {
	Customer     string        `json:"customer"`
	Performances []Performance `json:"performances"`
}

func statement(invoice Invoice, plays Play) string {
	totalAmount := 0.0
	volumeCredits := 0.0
	result := fmt.Sprintf("Statement for %s\n", invoice.Customer)

	for _, perf := range invoice.Performances {
		play := plays[perf.PlayID]
		thisAmount := 0.0

		switch play.Type {
		case "tragedy":
			thisAmount = 40000
			if perf.Audience > 30 {
				thisAmount += 1000 * (float64(perf.Audience - 30))
			}
		case "comedy":
			thisAmount = 30000
			if perf.Audience > 20 {
				thisAmount += 10000 + 500*(float64(perf.Audience-20))
			}
			thisAmount += 300 * float64(perf.Audience)
		default:
			panic(fmt.Sprintf("unknow type: %s", play.Type))
		}

		// add volume credits
		volumeCredits += math.Max(float64(perf.Audience-30), 0)
		// add extra credit for every ten comedy attendees
		if "comedy" == play.Type {
			volumeCredits += math.Floor(float64(perf.Audience / 5))
		}

		// print line for this order
		result += fmt.Sprintf("  %s: $%.2f (%d seats)\n", play.Name, thisAmount/100, perf.Audience)
		totalAmount += thisAmount
	}
	result += fmt.Sprintf("Amount owed is $%.2f\n", totalAmount/100)
	result += fmt.Sprintf("you earned %.0f credits\n", volumeCredits)
	return result
}

func main() {
	inv := Invoice{
		Customer: "Bigco",
		Performances: []Performance{
			{PlayID: "hamlet", Audience: 55},
			{PlayID: "as-like", Audience: 35},
			{PlayID: "othello", Audience: 40},
		}}
	plays := map[string]PlayDetails{
		"hamlet":  {Name: "Hamlet", Type: "tragedy"},
		"as-like": {Name: "As You Like It", Type: "comedy"},
		"othello": {Name: "Othello", Type: "tragedy"},
	}

	bill := statement(inv, plays)
	fmt.Println(bill)
}
