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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
)

// StopClusterReader is a Reader for the StopCluster structure.
type StopClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStopClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStopClusterOK creates a StopClusterOK with default headers values
func NewStopClusterOK() *StopClusterOK {
	return &StopClusterOK{}
}

/*
StopClusterOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type StopClusterOK struct {
	Payload *models.StopClusterResponse
}

// IsSuccess returns true when this stop cluster o k response has a 2xx status code
func (o *StopClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stop cluster o k response has a 3xx status code
func (o *StopClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop cluster o k response has a 4xx status code
func (o *StopClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop cluster o k response has a 5xx status code
func (o *StopClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stop cluster o k response a status code equal to that given
func (o *StopClusterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the stop cluster o k response
func (o *StopClusterOK) Code() int {
	return 200
}

func (o *StopClusterOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/stopCluster][%d] stopClusterOK %s", 200, payload)
}

func (o *StopClusterOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/stopCluster][%d] stopClusterOK %s", 200, payload)
}

func (o *StopClusterOK) GetPayload() *models.StopClusterResponse {
	return o.Payload
}

func (o *StopClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StopClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopClusterDefault creates a StopClusterDefault with default headers values
func NewStopClusterDefault(code int) *StopClusterDefault {
	return &StopClusterDefault{
		_statusCode: code,
	}
}

/*
StopClusterDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type StopClusterDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this stop cluster default response has a 2xx status code
func (o *StopClusterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this stop cluster default response has a 3xx status code
func (o *StopClusterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this stop cluster default response has a 4xx status code
func (o *StopClusterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this stop cluster default response has a 5xx status code
func (o *StopClusterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this stop cluster default response a status code equal to that given
func (o *StopClusterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the stop cluster default response
func (o *StopClusterDefault) Code() int {
	return o._statusCode
}

func (o *StopClusterDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/stopCluster][%d] stopCluster default %s", o._statusCode, payload)
}

func (o *StopClusterDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/stopCluster][%d] stopCluster default %s", o._statusCode, payload)
}

func (o *StopClusterDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *StopClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
