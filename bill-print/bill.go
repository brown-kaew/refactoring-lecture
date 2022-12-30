package main

import (
	"fmt"
	"math"
)

type Play struct {
	Name string
	Type string
}

type Plays map[string]Play

type Performance struct {
	PlayID   string `json:"playID"`
	Audience int    `json:"audience"`
}

type Invoice struct {
	Customer     string        `json:"customer"`
	Performances []Performance `json:"performances"`
}

func calculateAmount(play Play, perf Performance) float64 {
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

	return thisAmount
}

func playFor(plays Plays, perf Performance) Play {
	return plays[perf.PlayID]
}

func volumeCreditsFor(plays Plays, perf Performance) float64 {
	volumeCredits := 0.0
	// add volume credits
	volumeCredits += math.Max(float64(perf.Audience-30), 0)
	// add extra credit for every ten comedy attendees
	if "comedy" == playFor(plays, perf).Type {
		volumeCredits += math.Floor(float64(perf.Audience / 5))
	}

	return volumeCredits
}

func totalVolumeCredits(invoice Invoice, plays Plays) float64 {
	volumeCredits := 0.0
	for _, perf := range invoice.Performances {
		volumeCredits += volumeCreditsFor(plays, perf)
	}
	return volumeCredits
}

func totalAmounts(invoice Invoice, plays Plays) float64 {
	totalAmount := 0.0
	for _, perf := range invoice.Performances {
		totalAmount += calculateAmount(playFor(plays, perf), perf)
	}
	return totalAmount
}

func statement(invoice Invoice, plays Plays) string {
	result := fmt.Sprintf("Statement for %s\n", invoice.Customer)
	for _, perf := range invoice.Performances {
		// print line for this order
		result += fmt.Sprintf("  %s: $%.2f (%d seats)\n", playFor(plays, perf).Name, calculateAmount(playFor(plays, perf), perf)/100, perf.Audience)
	}
	result += fmt.Sprintf("Amount owed is $%.2f\n", totalAmounts(invoice, plays)/100)
	result += fmt.Sprintf("you earned %.0f credits\n", totalVolumeCredits(invoice, plays))
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
	plays := map[string]Play{
		"hamlet":  {Name: "Hamlet", Type: "tragedy"},
		"as-like": {Name: "As You Like It", Type: "comedy"},
		"othello": {Name: "Othello", Type: "tragedy"},
	}

	bill := statement(inv, plays)
	fmt.Println(bill)
}
