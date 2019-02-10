// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "models"
)

// GetRoleByIDReader is a Reader for the GetRoleByID structure.
type GetRoleByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoleByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRoleByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetRoleByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetRoleByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRoleByIDOK creates a GetRoleByIDOK with default headers values
func NewGetRoleByIDOK() *GetRoleByIDOK {
	return &GetRoleByIDOK{}
}

/*GetRoleByIDOK handles this case with default header values.

successful operation
*/
type GetRoleByIDOK struct {
	Payload *models.Role
}

func (o *GetRoleByIDOK) Error() string {
	return fmt.Sprintf("[GET /roles/{id}][%d] getRoleByIdOK  %+v", 200, o.Payload)
}

func (o *GetRoleByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoleByIDBadRequest creates a GetRoleByIDBadRequest with default headers values
func NewGetRoleByIDBadRequest() *GetRoleByIDBadRequest {
	return &GetRoleByIDBadRequest{}
}

/*GetRoleByIDBadRequest handles this case with default header values.

Invalid ID supplied
*/
type GetRoleByIDBadRequest struct {
}

func (o *GetRoleByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /roles/{id}][%d] getRoleByIdBadRequest ", 400)
}

func (o *GetRoleByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRoleByIDNotFound creates a GetRoleByIDNotFound with default headers values
func NewGetRoleByIDNotFound() *GetRoleByIDNotFound {
	return &GetRoleByIDNotFound{}
}

/*GetRoleByIDNotFound handles this case with default header values.

Role not found
*/
type GetRoleByIDNotFound struct {
}

func (o *GetRoleByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /roles/{id}][%d] getRoleByIdNotFound ", 404)
}

func (o *GetRoleByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
