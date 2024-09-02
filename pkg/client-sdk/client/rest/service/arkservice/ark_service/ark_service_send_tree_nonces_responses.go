// Code generated by go-swagger; DO NOT EDIT.

package ark_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ark-network/ark/pkg/client-sdk/client/rest/service/models"
)

// ArkServiceSendTreeNoncesReader is a Reader for the ArkServiceSendTreeNonces structure.
type ArkServiceSendTreeNoncesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ArkServiceSendTreeNoncesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewArkServiceSendTreeNoncesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewArkServiceSendTreeNoncesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewArkServiceSendTreeNoncesOK creates a ArkServiceSendTreeNoncesOK with default headers values
func NewArkServiceSendTreeNoncesOK() *ArkServiceSendTreeNoncesOK {
	return &ArkServiceSendTreeNoncesOK{}
}

/*
ArkServiceSendTreeNoncesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ArkServiceSendTreeNoncesOK struct {
	Payload models.V1SendTreeNoncesResponse
}

// IsSuccess returns true when this ark service send tree nonces o k response has a 2xx status code
func (o *ArkServiceSendTreeNoncesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ark service send tree nonces o k response has a 3xx status code
func (o *ArkServiceSendTreeNoncesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ark service send tree nonces o k response has a 4xx status code
func (o *ArkServiceSendTreeNoncesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ark service send tree nonces o k response has a 5xx status code
func (o *ArkServiceSendTreeNoncesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ark service send tree nonces o k response a status code equal to that given
func (o *ArkServiceSendTreeNoncesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ark service send tree nonces o k response
func (o *ArkServiceSendTreeNoncesOK) Code() int {
	return 200
}

func (o *ArkServiceSendTreeNoncesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/payment/tree/nonces][%d] arkServiceSendTreeNoncesOK %s", 200, payload)
}

func (o *ArkServiceSendTreeNoncesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/payment/tree/nonces][%d] arkServiceSendTreeNoncesOK %s", 200, payload)
}

func (o *ArkServiceSendTreeNoncesOK) GetPayload() models.V1SendTreeNoncesResponse {
	return o.Payload
}

func (o *ArkServiceSendTreeNoncesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArkServiceSendTreeNoncesDefault creates a ArkServiceSendTreeNoncesDefault with default headers values
func NewArkServiceSendTreeNoncesDefault(code int) *ArkServiceSendTreeNoncesDefault {
	return &ArkServiceSendTreeNoncesDefault{
		_statusCode: code,
	}
}

/*
ArkServiceSendTreeNoncesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ArkServiceSendTreeNoncesDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this ark service send tree nonces default response has a 2xx status code
func (o *ArkServiceSendTreeNoncesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ark service send tree nonces default response has a 3xx status code
func (o *ArkServiceSendTreeNoncesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ark service send tree nonces default response has a 4xx status code
func (o *ArkServiceSendTreeNoncesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ark service send tree nonces default response has a 5xx status code
func (o *ArkServiceSendTreeNoncesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ark service send tree nonces default response a status code equal to that given
func (o *ArkServiceSendTreeNoncesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ark service send tree nonces default response
func (o *ArkServiceSendTreeNoncesDefault) Code() int {
	return o._statusCode
}

func (o *ArkServiceSendTreeNoncesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/payment/tree/nonces][%d] ArkService_SendTreeNonces default %s", o._statusCode, payload)
}

func (o *ArkServiceSendTreeNoncesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/payment/tree/nonces][%d] ArkService_SendTreeNonces default %s", o._statusCode, payload)
}

func (o *ArkServiceSendTreeNoncesDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *ArkServiceSendTreeNoncesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}