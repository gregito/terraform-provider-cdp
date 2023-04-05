// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

// RepairFreeipaReader is a Reader for the RepairFreeipa structure.
type RepairFreeipaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepairFreeipaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepairFreeipaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRepairFreeipaDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRepairFreeipaOK creates a RepairFreeipaOK with default headers values
func NewRepairFreeipaOK() *RepairFreeipaOK {
	return &RepairFreeipaOK{}
}

/*
RepairFreeipaOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type RepairFreeipaOK struct {
	Payload *models.RepairFreeipaResponse
}

// IsSuccess returns true when this repair freeipa o k response has a 2xx status code
func (o *RepairFreeipaOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repair freeipa o k response has a 3xx status code
func (o *RepairFreeipaOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repair freeipa o k response has a 4xx status code
func (o *RepairFreeipaOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repair freeipa o k response has a 5xx status code
func (o *RepairFreeipaOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repair freeipa o k response a status code equal to that given
func (o *RepairFreeipaOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the repair freeipa o k response
func (o *RepairFreeipaOK) Code() int {
	return 200
}

func (o *RepairFreeipaOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/repairFreeipa][%d] repairFreeipaOK  %+v", 200, o.Payload)
}

func (o *RepairFreeipaOK) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/repairFreeipa][%d] repairFreeipaOK  %+v", 200, o.Payload)
}

func (o *RepairFreeipaOK) GetPayload() *models.RepairFreeipaResponse {
	return o.Payload
}

func (o *RepairFreeipaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RepairFreeipaResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepairFreeipaDefault creates a RepairFreeipaDefault with default headers values
func NewRepairFreeipaDefault(code int) *RepairFreeipaDefault {
	return &RepairFreeipaDefault{
		_statusCode: code,
	}
}

/*
RepairFreeipaDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type RepairFreeipaDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this repair freeipa default response has a 2xx status code
func (o *RepairFreeipaDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this repair freeipa default response has a 3xx status code
func (o *RepairFreeipaDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this repair freeipa default response has a 4xx status code
func (o *RepairFreeipaDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this repair freeipa default response has a 5xx status code
func (o *RepairFreeipaDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this repair freeipa default response a status code equal to that given
func (o *RepairFreeipaDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the repair freeipa default response
func (o *RepairFreeipaDefault) Code() int {
	return o._statusCode
}

func (o *RepairFreeipaDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/repairFreeipa][%d] repairFreeipa default  %+v", o._statusCode, o.Payload)
}

func (o *RepairFreeipaDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/repairFreeipa][%d] repairFreeipa default  %+v", o._statusCode, o.Payload)
}

func (o *RepairFreeipaDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *RepairFreeipaDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
