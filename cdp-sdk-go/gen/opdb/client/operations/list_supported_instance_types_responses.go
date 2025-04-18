// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/opdb/models"
)

// ListSupportedInstanceTypesReader is a Reader for the ListSupportedInstanceTypes structure.
type ListSupportedInstanceTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSupportedInstanceTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSupportedInstanceTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListSupportedInstanceTypesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListSupportedInstanceTypesOK creates a ListSupportedInstanceTypesOK with default headers values
func NewListSupportedInstanceTypesOK() *ListSupportedInstanceTypesOK {
	return &ListSupportedInstanceTypesOK{}
}

/*
ListSupportedInstanceTypesOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type ListSupportedInstanceTypesOK struct {
	Payload *models.ListSupportedInstanceTypesResponse
}

// IsSuccess returns true when this list supported instance types o k response has a 2xx status code
func (o *ListSupportedInstanceTypesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list supported instance types o k response has a 3xx status code
func (o *ListSupportedInstanceTypesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list supported instance types o k response has a 4xx status code
func (o *ListSupportedInstanceTypesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list supported instance types o k response has a 5xx status code
func (o *ListSupportedInstanceTypesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list supported instance types o k response a status code equal to that given
func (o *ListSupportedInstanceTypesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list supported instance types o k response
func (o *ListSupportedInstanceTypesOK) Code() int {
	return 200
}

func (o *ListSupportedInstanceTypesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/opdb/listSupportedInstanceTypes][%d] listSupportedInstanceTypesOK %s", 200, payload)
}

func (o *ListSupportedInstanceTypesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/opdb/listSupportedInstanceTypes][%d] listSupportedInstanceTypesOK %s", 200, payload)
}

func (o *ListSupportedInstanceTypesOK) GetPayload() *models.ListSupportedInstanceTypesResponse {
	return o.Payload
}

func (o *ListSupportedInstanceTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListSupportedInstanceTypesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSupportedInstanceTypesDefault creates a ListSupportedInstanceTypesDefault with default headers values
func NewListSupportedInstanceTypesDefault(code int) *ListSupportedInstanceTypesDefault {
	return &ListSupportedInstanceTypesDefault{
		_statusCode: code,
	}
}

/*
ListSupportedInstanceTypesDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ListSupportedInstanceTypesDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this list supported instance types default response has a 2xx status code
func (o *ListSupportedInstanceTypesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list supported instance types default response has a 3xx status code
func (o *ListSupportedInstanceTypesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list supported instance types default response has a 4xx status code
func (o *ListSupportedInstanceTypesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list supported instance types default response has a 5xx status code
func (o *ListSupportedInstanceTypesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list supported instance types default response a status code equal to that given
func (o *ListSupportedInstanceTypesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list supported instance types default response
func (o *ListSupportedInstanceTypesDefault) Code() int {
	return o._statusCode
}

func (o *ListSupportedInstanceTypesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/opdb/listSupportedInstanceTypes][%d] listSupportedInstanceTypes default %s", o._statusCode, payload)
}

func (o *ListSupportedInstanceTypesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/opdb/listSupportedInstanceTypes][%d] listSupportedInstanceTypes default %s", o._statusCode, payload)
}

func (o *ListSupportedInstanceTypesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListSupportedInstanceTypesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
