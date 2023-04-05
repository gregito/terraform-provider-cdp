// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/ml/models"
)

// ListCdswServicesReader is a Reader for the ListCdswServices structure.
type ListCdswServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCdswServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCdswServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListCdswServicesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListCdswServicesOK creates a ListCdswServicesOK with default headers values
func NewListCdswServicesOK() *ListCdswServicesOK {
	return &ListCdswServicesOK{}
}

/*
ListCdswServicesOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type ListCdswServicesOK struct {
	Payload *models.ListCdswServicesResponse
}

// IsSuccess returns true when this list cdsw services o k response has a 2xx status code
func (o *ListCdswServicesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list cdsw services o k response has a 3xx status code
func (o *ListCdswServicesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list cdsw services o k response has a 4xx status code
func (o *ListCdswServicesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list cdsw services o k response has a 5xx status code
func (o *ListCdswServicesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list cdsw services o k response a status code equal to that given
func (o *ListCdswServicesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list cdsw services o k response
func (o *ListCdswServicesOK) Code() int {
	return 200
}

func (o *ListCdswServicesOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/ml/listCdswServices][%d] listCdswServicesOK  %+v", 200, o.Payload)
}

func (o *ListCdswServicesOK) String() string {
	return fmt.Sprintf("[POST /api/v1/ml/listCdswServices][%d] listCdswServicesOK  %+v", 200, o.Payload)
}

func (o *ListCdswServicesOK) GetPayload() *models.ListCdswServicesResponse {
	return o.Payload
}

func (o *ListCdswServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListCdswServicesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCdswServicesDefault creates a ListCdswServicesDefault with default headers values
func NewListCdswServicesDefault(code int) *ListCdswServicesDefault {
	return &ListCdswServicesDefault{
		_statusCode: code,
	}
}

/*
ListCdswServicesDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ListCdswServicesDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this list cdsw services default response has a 2xx status code
func (o *ListCdswServicesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list cdsw services default response has a 3xx status code
func (o *ListCdswServicesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list cdsw services default response has a 4xx status code
func (o *ListCdswServicesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list cdsw services default response has a 5xx status code
func (o *ListCdswServicesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list cdsw services default response a status code equal to that given
func (o *ListCdswServicesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list cdsw services default response
func (o *ListCdswServicesDefault) Code() int {
	return o._statusCode
}

func (o *ListCdswServicesDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/ml/listCdswServices][%d] listCdswServices default  %+v", o._statusCode, o.Payload)
}

func (o *ListCdswServicesDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/ml/listCdswServices][%d] listCdswServices default  %+v", o._statusCode, o.Payload)
}

func (o *ListCdswServicesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListCdswServicesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
