package rental

import "fmt"

type Customer struct {
	name    string
	rentals []Rental
}

func NewCustomer(name string) Customer {
	return Customer{
		name:    name,
		rentals: []Rental{},
	}
}
func (rcvr Customer) AddRental(arg Rental) {
	rcvr.rentals = append(rcvr.rentals, arg)
}
func (rcvr Customer) Name() string {
	return rcvr.name
}

func (rental Rental) Charge() float64 {
	result := 0.0
	switch rental.Movie().PriceCode() {
	case REGULAR:
		result += 2
		if rental.DaysRented() > 2 {
			result += float64(rental.DaysRented()-2) * 1.5
		}
	case NEW_RELEASE:
		result += float64(rental.DaysRented()) * 3.0
	case CHILDRENS:
		result += 1.5
		if rental.DaysRented() > 3 {
			result += float64(rental.DaysRented()-3) * 1.5
		}
	}
	return result
}

func (rcvr Customer) Statement() string {
	totalAmount := 0.0
	frequentRenterPoints := 0
	result := fmt.Sprintf("Rental Record for %v\n", rcvr.Name())
	for _, rental := range rcvr.rentals {
		thisAmount := rental.Charge()

		frequentRenterPoints++
		if rental.Movie().PriceCode() == NEW_RELEASE && rental.DaysRented() > 1 {
			frequentRenterPoints++
		}
		result += fmt.Sprintf("\t%v\t%.1f\n", rental.Movie().Title(), thisAmount)
		totalAmount += thisAmount
	}
	result += fmt.Sprintf("Amount owed is %.1f\n", totalAmount)
	result += fmt.Sprintf("You earned %v frequent renter points", frequentRenterPoints)
	return result
}
