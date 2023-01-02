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
	frequentRenterPoints := getTotalPoints(customer.rentals)

	totalAmount := getTotalAmounts(customer.rentals)

	result := fmt.Sprintf("Rental Record for %v\n", customer.Name())
	for _, rental := range customer.rentals {
		result += fmt.Sprintf("\t%v\t%.1f\n", rental.Movie().Title(), rental.Charge())
	}
	result += fmt.Sprintf("Amount owed is %.1f\n", totalAmount)
	result += fmt.Sprintf("You earned %v frequent renter points", frequentRenterPoints)
	return result
}
