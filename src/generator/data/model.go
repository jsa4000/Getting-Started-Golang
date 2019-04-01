package data

import (
	"time"

	fake "github.com/brianvoe/gofakeit"
)

// Client structs with random injected data
type Client struct {
	ID               string
	FirstName        string
	LastName         string
	FullName         string
	Gender           string
	Email            string
	Phone            string
	Company          string
	CreditCardNumber int
	JobTitle         string
	Birth            time.Time
	StartDate        time.Time
	EndDate          time.Time
	Address          *fake.AddressInfo
	Vehicle          *fake.VehicleInfo
}

// NewClient creates a new client with random information
func NewClient() *Client {
	return &Client{
		ID:               fake.UUID(),
		FirstName:        fake.FirstName(),
		LastName:         fake.LastName(),
		FullName:         fake.Name(),
		Gender:           fake.Gender(),
		Email:            fake.Email(),
		Phone:            fake.Phone(),
		Company:          fake.Company(),
		CreditCardNumber: fake.CreditCardNumber(),
		JobTitle:         fake.JobTitle(),
		Birth:            fake.Date(),
		StartDate:        fake.Date(),
		EndDate:          fake.Date(),
		Address:          fake.Address(),
		Vehicle:          fake.Vehicle(),
	}
}
