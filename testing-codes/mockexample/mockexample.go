package mockexample

import "github.com/pkg/errors"

type OutputTrip struct {
	Active bool
}

type TripService struct {
	Customer CustomerRepository
	Trip     TripRepository
}

func (s *TripService) Run(customerCode string) (OutputTrip, error) {
	customer, err := s.Customer.GetByCode(customerCode)
	if err != nil {
		return OutputTrip{}, errors.Wrapf(err, "TripService.Run(`%s`) got error", customerCode)
	}

	trip, err := s.Trip.GetByCustomerID(customer.GetID(), "reserved")
	if err != nil {
		return OutputTrip{}, errors.Wrapf(err, "TripService.Run(`%s`) got error", customerCode)
	}
	if !trip.IsReserved() {
		return OutputTrip{Active: false}, nil
	}
	if !customer.IsBan() {
		return OutputTrip{Active: false}, nil
	}

	return OutputTrip{Active: true}, nil
}

type CustomerRepository interface {
	GetByCode(code string) (Customer, error)
}

type Customer struct {
	ID int
}

func (c *Customer) GetID() int {
	return c.ID
}
func (c *Customer) IsBan() bool {
	return false
}

type TripRepository interface {
	GetByCustomerID(id int, status string) (Trip, error)
}

type Trip struct{}

func (t *Trip) IsReserved() bool {
	return true
}
