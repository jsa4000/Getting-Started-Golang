// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CreateRoleReader is a Reader for the CreateRole structure.
type CreateRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 405:
		result := NewCreateRoleMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateRoleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateRoleMethodNotAllowed creates a CreateRoleMethodNotAllowed with default headers values
func NewCreateRoleMethodNotAllowed() *CreateRoleMethodNotAllowed {
	return &CreateRoleMethodNotAllowed{}
}

/*CreateRoleMethodNotAllowed handles this case with default header values.

Invalid input
*/
type CreateRoleMethodNotAllowed struct {
}

func (o *CreateRoleMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /roles][%d] createRoleMethodNotAllowed ", 405)
}

func (o *CreateRoleMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRoleInternalServerError creates a CreateRoleInternalServerError with default headers values
func NewCreateRoleInternalServerError() *CreateRoleInternalServerError {
	return &CreateRoleInternalServerError{}
}

/*CreateRoleInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateRoleInternalServerError struct {
}

func (o *CreateRoleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /roles][%d] createRoleInternalServerError ", 500)
}

func (o *CreateRoleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
