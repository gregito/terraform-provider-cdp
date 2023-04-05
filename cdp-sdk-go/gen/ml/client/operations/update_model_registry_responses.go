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

// UpdateModelRegistryReader is a Reader for the UpdateModelRegistry structure.
type UpdateModelRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateModelRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateModelRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateModelRegistryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateModelRegistryOK creates a UpdateModelRegistryOK with default headers values
func NewUpdateModelRegistryOK() *UpdateModelRegistryOK {
	return &UpdateModelRegistryOK{}
}

/*
UpdateModelRegistryOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type UpdateModelRegistryOK struct {
	Payload models.UpdateModelRegistryResponse
}

// IsSuccess returns true when this update model registry o k response has a 2xx status code
func (o *UpdateModelRegistryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update model registry o k response has a 3xx status code
func (o *UpdateModelRegistryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update model registry o k response has a 4xx status code
func (o *UpdateModelRegistryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update model registry o k response has a 5xx status code
func (o *UpdateModelRegistryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update model registry o k response a status code equal to that given
func (o *UpdateModelRegistryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update model registry o k response
func (o *UpdateModelRegistryOK) Code() int {
	return 200
}

func (o *UpdateModelRegistryOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/ml/updateModelRegistry][%d] updateModelRegistryOK  %+v", 200, o.Payload)
}

func (o *UpdateModelRegistryOK) String() string {
	return fmt.Sprintf("[POST /api/v1/ml/updateModelRegistry][%d] updateModelRegistryOK  %+v", 200, o.Payload)
}

func (o *UpdateModelRegistryOK) GetPayload() models.UpdateModelRegistryResponse {
	return o.Payload
}

func (o *UpdateModelRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateModelRegistryDefault creates a UpdateModelRegistryDefault with default headers values
func NewUpdateModelRegistryDefault(code int) *UpdateModelRegistryDefault {
	return &UpdateModelRegistryDefault{
		_statusCode: code,
	}
}

/*
UpdateModelRegistryDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type UpdateModelRegistryDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update model registry default response has a 2xx status code
func (o *UpdateModelRegistryDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update model registry default response has a 3xx status code
func (o *UpdateModelRegistryDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update model registry default response has a 4xx status code
func (o *UpdateModelRegistryDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update model registry default response has a 5xx status code
func (o *UpdateModelRegistryDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update model registry default response a status code equal to that given
func (o *UpdateModelRegistryDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update model registry default response
func (o *UpdateModelRegistryDefault) Code() int {
	return o._statusCode
}

func (o *UpdateModelRegistryDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/ml/updateModelRegistry][%d] updateModelRegistry default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateModelRegistryDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/ml/updateModelRegistry][%d] updateModelRegistry default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateModelRegistryDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateModelRegistryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
