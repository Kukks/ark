// Code generated by go-swagger; DO NOT EDIT.

package admin_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ark-network/ark-sdk/rest/admin/models"
)

// AdminServiceGetRoundsReader is a Reader for the AdminServiceGetRounds structure.
type AdminServiceGetRoundsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminServiceGetRoundsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAdminServiceGetRoundsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewAdminServiceGetRoundsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAdminServiceGetRoundsOK creates a AdminServiceGetRoundsOK with default headers values
func NewAdminServiceGetRoundsOK() *AdminServiceGetRoundsOK {
	return &AdminServiceGetRoundsOK{}
}

/*AdminServiceGetRoundsOK handles this case with default header values.

A successful response.
*/
type AdminServiceGetRoundsOK struct {
	Payload *models.V1GetRoundsResponse
}

func (o *AdminServiceGetRoundsOK) Error() string {
	return fmt.Sprintf("[POST /v1/admin/rounds][%d] adminServiceGetRoundsOK  %+v", 200, o.Payload)
}

func (o *AdminServiceGetRoundsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetRoundsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminServiceGetRoundsDefault creates a AdminServiceGetRoundsDefault with default headers values
func NewAdminServiceGetRoundsDefault(code int) *AdminServiceGetRoundsDefault {
	return &AdminServiceGetRoundsDefault{
		_statusCode: code,
	}
}

/*AdminServiceGetRoundsDefault handles this case with default header values.

An unexpected error response.
*/
type AdminServiceGetRoundsDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the admin service get rounds default response
func (o *AdminServiceGetRoundsDefault) Code() int {
	return o._statusCode
}

func (o *AdminServiceGetRoundsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/admin/rounds][%d] AdminService_GetRounds default  %+v", o._statusCode, o.Payload)
}

func (o *AdminServiceGetRoundsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
