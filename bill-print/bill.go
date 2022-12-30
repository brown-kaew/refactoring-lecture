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

type Rate struct {
	Play     Play
	Amount   float64
	Audience int
}

type Bill struct {
	Customer           string
	Rates              []Rate
	TotalAmounts       float64
	TotalVolumeCredits float64
}

func amountFor(play Play, perf Performance) float64 {
	result := 0.0
	switch play.Type {
	case "tragedy":
		result = 40000
		if perf.Audience > 30 {
			result += 1000 * (float64(perf.Audience - 30))
		}
	case "comedy":
		result = 30000
		if perf.Audience > 20 {
			result += 10000 + 500*(float64(perf.Audience-20))
		}
		result += 300 * float64(perf.Audience)
	default:
		panic(fmt.Sprintf("unknow type: %s", play.Type))
	}

	return result
}

func playFor(plays Plays, perf Performance) Play {
	return plays[perf.PlayID]
}

func volumeCreditsFor(plays Plays, perf Performance) float64 {
	result := 0.0
	// add volume credits
	result += math.Max(float64(perf.Audience-30), 0)
	// add extra credit for every ten comedy attendees
	if "comedy" == playFor(plays, perf).Type {
		result += math.Floor(float64(perf.Audience / 5))
	}
	return result
}

func totalVolumeCredits(performances []Performance, plays Plays) float64 {
	result := 0.0
	for _, perf := range performances {
		result += volumeCreditsFor(plays, perf)
	}
	return result
}

func totalAmounts(performances []Performance, plays Plays) float64 {
	result := 0.0
	for _, perf := range performances {
		result += amountFor(playFor(plays, perf), perf)
	}
	return result
}

func statement(invoice Invoice, plays Plays) string {
	rates := []Rate{}
	for _, perf := range invoice.Performances {
		rate := Rate{
			Play:     playFor(plays, perf),
			Amount:   amountFor(playFor(plays, perf), perf),
			Audience: perf.Audience,
		}
		rates = append(rates, rate)
	}

	bill := Bill{
		Customer:           invoice.Customer,
		Rates:              rates,
		TotalAmounts:       totalAmounts(invoice.Performances, plays),
		TotalVolumeCredits: totalVolumeCredits(invoice.Performances, plays),
	}
	return renderPlainText(bill)
}

func renderPlainText(bill Bill) string {
	result := fmt.Sprintf("Statement for %s\n", bill.Customer)
	for _, rate := range bill.Rates {
		// print line for this order
		result += fmt.Sprintf("  %s: $%.2f (%d seats)\n", rate.Play.Name, rate.Amount/100, rate.Audience)
	}
	result += fmt.Sprintf("Amount owed is $%.2f\n", bill.TotalAmounts/100)
	result += fmt.Sprintf("you earned %.0f credits\n", bill.TotalVolumeCredits)
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
