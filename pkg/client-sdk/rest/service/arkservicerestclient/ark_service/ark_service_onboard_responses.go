// Code generated by go-swagger; DO NOT EDIT.

package ark_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ark-network/ark-sdk/rest/service/models"
)

// ArkServiceOnboardReader is a Reader for the ArkServiceOnboard structure.
type ArkServiceOnboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ArkServiceOnboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewArkServiceOnboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewArkServiceOnboardDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewArkServiceOnboardOK creates a ArkServiceOnboardOK with default headers values
func NewArkServiceOnboardOK() *ArkServiceOnboardOK {
	return &ArkServiceOnboardOK{}
}

/*ArkServiceOnboardOK handles this case with default header values.

A successful response.
*/
type ArkServiceOnboardOK struct {
	Payload models.V1OnboardResponse
}

func (o *ArkServiceOnboardOK) Error() string {
	return fmt.Sprintf("[POST /v1/onboard][%d] arkServiceOnboardOK  %+v", 200, o.Payload)
}

func (o *ArkServiceOnboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArkServiceOnboardDefault creates a ArkServiceOnboardDefault with default headers values
func NewArkServiceOnboardDefault(code int) *ArkServiceOnboardDefault {
	return &ArkServiceOnboardDefault{
		_statusCode: code,
	}
}

/*ArkServiceOnboardDefault handles this case with default header values.

An unexpected error response.
*/
type ArkServiceOnboardDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the ark service onboard default response
func (o *ArkServiceOnboardDefault) Code() int {
	return o._statusCode
}

func (o *ArkServiceOnboardDefault) Error() string {
	return fmt.Sprintf("[POST /v1/onboard][%d] ArkService_Onboard default  %+v", o._statusCode, o.Payload)
}

func (o *ArkServiceOnboardDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
