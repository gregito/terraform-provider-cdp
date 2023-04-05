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

// GetCliForEnvironmentReader is a Reader for the GetCliForEnvironment structure.
type GetCliForEnvironmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCliForEnvironmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCliForEnvironmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCliForEnvironmentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCliForEnvironmentOK creates a GetCliForEnvironmentOK with default headers values
func NewGetCliForEnvironmentOK() *GetCliForEnvironmentOK {
	return &GetCliForEnvironmentOK{}
}

/*
GetCliForEnvironmentOK describes a response with status code 200, with default header values.

Create environment CLI command for valid request.
*/
type GetCliForEnvironmentOK struct {
	Payload *models.GetCliForEnvironmentResponse
}

// IsSuccess returns true when this get cli for environment o k response has a 2xx status code
func (o *GetCliForEnvironmentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cli for environment o k response has a 3xx status code
func (o *GetCliForEnvironmentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cli for environment o k response has a 4xx status code
func (o *GetCliForEnvironmentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cli for environment o k response has a 5xx status code
func (o *GetCliForEnvironmentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cli for environment o k response a status code equal to that given
func (o *GetCliForEnvironmentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get cli for environment o k response
func (o *GetCliForEnvironmentOK) Code() int {
	return 200
}

func (o *GetCliForEnvironmentOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/getCliForEnvironment][%d] getCliForEnvironmentOK  %+v", 200, o.Payload)
}

func (o *GetCliForEnvironmentOK) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/getCliForEnvironment][%d] getCliForEnvironmentOK  %+v", 200, o.Payload)
}

func (o *GetCliForEnvironmentOK) GetPayload() *models.GetCliForEnvironmentResponse {
	return o.Payload
}

func (o *GetCliForEnvironmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCliForEnvironmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCliForEnvironmentDefault creates a GetCliForEnvironmentDefault with default headers values
func NewGetCliForEnvironmentDefault(code int) *GetCliForEnvironmentDefault {
	return &GetCliForEnvironmentDefault{
		_statusCode: code,
	}
}

/*
GetCliForEnvironmentDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type GetCliForEnvironmentDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get cli for environment default response has a 2xx status code
func (o *GetCliForEnvironmentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get cli for environment default response has a 3xx status code
func (o *GetCliForEnvironmentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get cli for environment default response has a 4xx status code
func (o *GetCliForEnvironmentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get cli for environment default response has a 5xx status code
func (o *GetCliForEnvironmentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get cli for environment default response a status code equal to that given
func (o *GetCliForEnvironmentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get cli for environment default response
func (o *GetCliForEnvironmentDefault) Code() int {
	return o._statusCode
}

func (o *GetCliForEnvironmentDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/getCliForEnvironment][%d] getCliForEnvironment default  %+v", o._statusCode, o.Payload)
}

func (o *GetCliForEnvironmentDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/getCliForEnvironment][%d] getCliForEnvironment default  %+v", o._statusCode, o.Payload)
}

func (o *GetCliForEnvironmentDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCliForEnvironmentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
