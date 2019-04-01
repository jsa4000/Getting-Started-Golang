package data

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

// Generate generates a file with test data with the given rows
func Generate(out string, rows int) error {
	file, err := os.Create(out)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for i := 1; i <= rows; i++ {
		client := NewClient()
		if err := writer.Write([]string{
			client.ID,
			client.FirstName,
			client.LastName,
			client.FullName,
			client.Gender,
			client.Email,
			client.Phone,
			client.Company,
			strconv.Itoa(client.CreditCardNumber),
			client.JobTitle,
			client.Birth.Format(time.RFC3339),
			client.StartDate.Format(time.RFC3339),
			client.EndDate.Format(time.RFC3339),
			client.Address.Address,
			client.Address.Street,
			client.Address.City,
			client.Address.Country,
			client.Address.State,
			client.Address.Zip,
		}); err != nil {
			return err
		}
	}
	return nil

}
