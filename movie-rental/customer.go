package rental

import "fmt"

type Customer struct {
	name    string
	rentals []Rental
}

type Receipt struct {
	name        string
	points      int
	totalAmount float64
	details     []MovieDetails
}

type MovieDetails struct {
	title  string
	amount float64
}

func NewCustomer(name string) Customer {
	return Customer{
		name:    name,
		rentals: []Rental{},
	}
}
func (c Customer) AddRental(arg Rental) {
	c.rentals = append(c.rentals, arg)
}

func (c Customer) Name() string {
	return c.name
}

func (rental Rental) Charge() float64 {
	return rental.Movie().Charge(rental.DaysRented())
}

func (rental Rental) GetPoint() int {
	if rental.Movie().PriceCode() == NEW_RELEASE && rental.DaysRented() > 1 {
		return 2
	}
	return 1
}

func getTotalPoints(rentals []Rental) int {
	result := 0
	for _, rental := range rentals {
		result += rental.GetPoint()
	}
	return result
}

func getTotalAmounts(rentals []Rental) float64 {
	result := 0.0
	for _, rental := range rentals {
		result += rental.Charge()
	}
	return result
}

func (customer Customer) Statement() string {
	receipt := Receipt{
		name:        customer.name,
		points:      getTotalPoints(customer.rentals),
		totalAmount: getTotalAmounts(customer.rentals),
		details:     []MovieDetails{},
	}

	for _, rental := range customer.rentals {
		receipt.details = append(receipt.details, MovieDetails{
			title:  rental.Movie().Title(),
			amount: rental.Charge(),
		})
	}
	return renderPlainText(receipt)
}

func renderPlainText(receipt Receipt) string {
	result := fmt.Sprintf("Rental Record for %v\n", receipt.name)
	for _, d := range receipt.details {
		result += fmt.Sprintf("\t%v\t%.1f\n", d.title, d.amount)
	}
	result += fmt.Sprintf("Amount owed is %.1f\n", receipt.totalAmount)
	result += fmt.Sprintf("You earned %v frequent renter points", receipt.points)
	return result
}
