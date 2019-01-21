package goplayground

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
		Title    string `validate:"ascii,required"`
		Message  string `validate:"ascii"`
		Email    string `validate:"email"`
		AuthorIP string `validate:"ipv4"`
		Date     string `validate:"-"`
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
		Title    string `validate:"alphanum,required,omitempty"`
		Message  string `validate:"ascii"`
		AuthorIP string `validate:"ipv4"`
		Date     string `validate:"-"`
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
		Title    string  `validate:"min=0,max=255,required"`
		Message  string  `validate:"ascii,required"`
		AuthorIP string  `validate:"ipv4"`
		Year     int     `validate:"gte=1900,lte=10000"`
		rate     float64 `validate:"-"`
		colors   string  `validate:"oneof=color red yellow blue)"`
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
		Title    string  `validate:"min=0,max=10,required"`
		Message  string  `validate:"ascii,required"`
		AuthorIP string  `validate:"ipv4"`
		Year     int     `validate:"gte=1900,lte=10000"`
		rate     float64 `validate:"-"`
		Status   bool
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
		Title    string  `validate:"min=0,max=255,required"`
		Message  string  `validate:"min=0,max=1024,required"`
		AuthorIP string  `validate:"ipv4"`
		Year     int     `validate:"gte=1900,lte=10000"`
		rate     float64 `validate:"-"`
		Status   bool
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
		Title    string  `validate:"min=0,max=255,required"`
		Message  string  `validate:"min=0,max=1024,required"`
		AuthorIP string  `validate:"ipv4"`
		Email    string  `validate:"email"`
		Year     int     `validate:"gte=1900,lte=10000"`
		rate     float64 `validate:"-"`
		Status   bool
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
		Title    string  `validate:"min=0,max=255,required"`
		Message  string  `validate:"min=0,max=1024,required"`
		AuthorIP string  `validate:"ipv4"`
		Year     int     `validate:"gte=1900,lte=10000"`
		rate     float64 `validate:"-"`
		Status   bool
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
		Title    string  `validate:"min=0,max=255,required"`
		Message  string  `validate:"min=0,max=1024,required"`
		AuthorIP string  `validate:"ipv4"`
		Year     int     `validate:"gte=1900,lte=10000"`
		rate     float64 `validate:"-"`
		Status   bool
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
