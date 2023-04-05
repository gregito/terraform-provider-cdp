// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
)

// AssignUserResourceRoleReader is a Reader for the AssignUserResourceRole structure.
type AssignUserResourceRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssignUserResourceRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssignUserResourceRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAssignUserResourceRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAssignUserResourceRoleOK creates a AssignUserResourceRoleOK with default headers values
func NewAssignUserResourceRoleOK() *AssignUserResourceRoleOK {
	return &AssignUserResourceRoleOK{}
}

/*
AssignUserResourceRoleOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type AssignUserResourceRoleOK struct {
	Payload models.AssignUserResourceRoleResponse
}

// IsSuccess returns true when this assign user resource role o k response has a 2xx status code
func (o *AssignUserResourceRoleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this assign user resource role o k response has a 3xx status code
func (o *AssignUserResourceRoleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign user resource role o k response has a 4xx status code
func (o *AssignUserResourceRoleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this assign user resource role o k response has a 5xx status code
func (o *AssignUserResourceRoleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this assign user resource role o k response a status code equal to that given
func (o *AssignUserResourceRoleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the assign user resource role o k response
func (o *AssignUserResourceRoleOK) Code() int {
	return 200
}

func (o *AssignUserResourceRoleOK) Error() string {
	return fmt.Sprintf("[POST /iam/assignUserResourceRole][%d] assignUserResourceRoleOK  %+v", 200, o.Payload)
}

func (o *AssignUserResourceRoleOK) String() string {
	return fmt.Sprintf("[POST /iam/assignUserResourceRole][%d] assignUserResourceRoleOK  %+v", 200, o.Payload)
}

func (o *AssignUserResourceRoleOK) GetPayload() models.AssignUserResourceRoleResponse {
	return o.Payload
}

func (o *AssignUserResourceRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignUserResourceRoleDefault creates a AssignUserResourceRoleDefault with default headers values
func NewAssignUserResourceRoleDefault(code int) *AssignUserResourceRoleDefault {
	return &AssignUserResourceRoleDefault{
		_statusCode: code,
	}
}

/*
AssignUserResourceRoleDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type AssignUserResourceRoleDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this assign user resource role default response has a 2xx status code
func (o *AssignUserResourceRoleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this assign user resource role default response has a 3xx status code
func (o *AssignUserResourceRoleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this assign user resource role default response has a 4xx status code
func (o *AssignUserResourceRoleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this assign user resource role default response has a 5xx status code
func (o *AssignUserResourceRoleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this assign user resource role default response a status code equal to that given
func (o *AssignUserResourceRoleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the assign user resource role default response
func (o *AssignUserResourceRoleDefault) Code() int {
	return o._statusCode
}

func (o *AssignUserResourceRoleDefault) Error() string {
	return fmt.Sprintf("[POST /iam/assignUserResourceRole][%d] assignUserResourceRole default  %+v", o._statusCode, o.Payload)
}

func (o *AssignUserResourceRoleDefault) String() string {
	return fmt.Sprintf("[POST /iam/assignUserResourceRole][%d] assignUserResourceRole default  %+v", o._statusCode, o.Payload)
}

func (o *AssignUserResourceRoleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *AssignUserResourceRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
