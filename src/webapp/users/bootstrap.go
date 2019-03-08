package users

import "encoding/json"

const bootstrapData = `[
{
	"name": "root",
	"email": "root@example.com",
	"password": "$2a$10$iJL5k6j.fLHPIefUQHXNbeBf/FlMIqvB3x8SfbFzZpyDk396TGBMa",
	"roles": [
		"READ",
		"WRITE",
		"ADMIN"
	]
},
{
	"name": "user",
	"email": "user@example.com",
	"password": "$2a$10$vVtepm/gwOd6rWtSTlqpbeLzAM0I59uqTnN1PmlBKA5pKtg2dtyg2",
	"roles": [
		"READ",
		"WRITE"
	]
},
{
	"name": "viewer",
	"email": "viewer@example.com",
	"password": "$2a$10$H/T0Rzn6WPos6MnNC/IuxeS5gaNe6ckeF9uwfYSSpWJy4g2XO0Y6m",
	"roles": [
		"READ"
	]
}
]`

// BootstrapData Get the boot strapdata deserialized
func BootstrapData() ([]User, error) {
	var result []User
	err := json.Unmarshal([]byte(bootstrapData), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
