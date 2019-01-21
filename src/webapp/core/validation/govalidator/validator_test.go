package govalidator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var validator *Validator

func init() {
	validator = New()
}

func TestDataWithoutValidations(t *testing.T) {

	type PostRequest struct {
		Title    string
		Message  string
		AuthorIP string
		Date     string
	}

	request := PostRequest{
		Title:    "My Example Post",
		Message:  "duck",
		AuthorIP: "123.234.54.3",
	}

	_, err := validator.Validate(&request)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, nil, err)
}

func TestDataWithValidations(t *testing.T) {

	type PostRequest struct {
		Title    string `valid:"ascii,required"`
		Message  string `valid:"ascii"`
		Email    string `valid:"email"`
		AuthorIP string `valid:"ipv4"`
		Date     string `valid:"-"`
	}

	request := PostRequest{
		Title:    "My Example Post",
		Message:  "duck",
		Email:    "test@example.com",
		AuthorIP: "123.234.54.3",
	}

	_, err := validator.Validate(&request)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, nil, err)
}

func TestDataWithErrorInValidation(t *testing.T) {

	type PostRequest struct {
		Title    string `valid:"alphanum,required"`
		Message  string `valid:"ascii"`
		AuthorIP string `valid:"ipv4"`
		Date     string `valid:"-"`
	}

	request := PostRequest{
		Title:    "My Example Post",
		Message:  "duck",
		AuthorIP: "123.234.54.3",
	}

	_, err := validator.Validate(&request)
	if err != nil {
		fmt.Println(err)
	}
	assert.NotEqual(t, nil, err)
}

func TestDataAdvancedValidations(t *testing.T) {

	type PostRequest struct {
		Title    string  `valid:"length(0|255),required"`
		Message  string  `valid:"ascii,required"`
		AuthorIP string  `valid:"ipv4"`
		Year     int     `valid:"range(1900|10000)"`
		rate     float64 `valid:"-"`
		Status   bool    `valid:"in(true|True|False|false)"`
	}

	request := PostRequest{
		Title:    "My Example Post",
		Message:  "duck",
		AuthorIP: "123.234.54.3",
		Year:     2018,
		rate:     6.23,
	}

	_, err := validator.Validate(&request)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, nil, err)
}

func TestDataLengthErrorValidation(t *testing.T) {

	type PostRequest struct {
		Title    string  `valid:"length(0|10),required"`
		Message  string  `valid:"ascii,required"`
		AuthorIP string  `valid:"ipv4"`
		Year     int     `valid:"range(1900|10000)"`
		rate     float64 `valid:"-"`
		Status   bool    `valid:"in(true|True|False|false)"`
	}

	request := PostRequest{
		Title:    "My Example Post",
		Message:  "duck",
		AuthorIP: "123.234.54.3",
		Year:     2018,
		rate:     6.23,
	}

	_, err := validator.Validate(&request)
	if err != nil {
		fmt.Println(err)
	}
	assert.NotEqual(t, nil, err)
}

func TestDataRangeErrorValidation(t *testing.T) {

	type PostRequest struct {
		Title    string  `valid:"length(0|255),required"`
		Message  string  `valid:"length(0|1024),required"`
		AuthorIP string  `valid:"ipv4"`
		Year     int     `valid:"range(1900|10000)"`
		rate     float64 `valid:"-"`
		Status   bool    `valid:"in(true|True|False|false)"`
	}

	request := PostRequest{
		Title:    "My Example Post",
		Message:  "duck",
		AuthorIP: "123.234.54.3",
		Year:     1200,
		rate:     6.23,
	}

	_, err := validator.Validate(&request)
	if err != nil {
		fmt.Println(err)
	}
	assert.NotEqual(t, nil, err)
}

func TestDataEmailErrorValidation(t *testing.T) {

	type PostRequest struct {
		Title    string  `valid:"length(0|255),required"`
		Message  string  `valid:"length(0|1024),required"`
		AuthorIP string  `valid:"ipv4"`
		Email    string  `valid:"email"`
		Year     int     `valid:"range(1900|10000)"`
		rate     float64 `valid:"-"`
		Status   bool    `valid:"in(true|True|False|false)"`
	}

	request := PostRequest{
		Title:    "My Example Post",
		Message:  "duck",
		AuthorIP: "123.234.54.3",
		Email:    "testexample.com",
		Year:     1983,
		rate:     6.23,
	}

	_, err := validator.Validate(&request)
	if err != nil {
		fmt.Println(err)
	}
	assert.NotEqual(t, nil, err)
}

func TestDataRequiredFieldErrorValidation(t *testing.T) {

	type PostRequest struct {
		Title    string  `valid:"length(0|255),required"`
		Message  string  `valid:"length(0|1024),required"`
		AuthorIP string  `valid:"ipv4"`
		Year     int     `valid:"range(1900|10000)"`
		rate     float64 `valid:"-"`
		Status   bool    `valid:"in(true|True|False|false)"`
	}

	request := PostRequest{
		Title: "My Example Post",
		rate:  6.23,
	}

	_, err := validator.Validate(&request)
	if err != nil {
		fmt.Println(err)
	}
	assert.NotEqual(t, nil, err)
}

func TestDataRequiredEmptyFieldErrorValidation(t *testing.T) {

	type PostRequest struct {
		Title    string  `valid:"length(0|255),required"`
		Message  string  `valid:"length(0|1024),required"`
		AuthorIP string  `valid:"ipv4"`
		Year     int     `valid:"range(1900|10000)"`
		rate     float64 `valid:"-"`
		Status   bool    `valid:"in(true|True|False|false)"`
	}

	request := PostRequest{
		Title:   "My Example Post",
		Message: "",
		rate:    6.23,
	}

	_, err := validator.Validate(&request)
	if err != nil {
		fmt.Println(err)
	}
	assert.NotEqual(t, nil, err)
}
