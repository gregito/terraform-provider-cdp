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

// AddUserToGroupReader is a Reader for the AddUserToGroup structure.
type AddUserToGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddUserToGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddUserToGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddUserToGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddUserToGroupOK creates a AddUserToGroupOK with default headers values
func NewAddUserToGroupOK() *AddUserToGroupOK {
	return &AddUserToGroupOK{}
}

/*
AddUserToGroupOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type AddUserToGroupOK struct {
	Payload models.AddUserToGroupResponse
}

// IsSuccess returns true when this add user to group o k response has a 2xx status code
func (o *AddUserToGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add user to group o k response has a 3xx status code
func (o *AddUserToGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add user to group o k response has a 4xx status code
func (o *AddUserToGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add user to group o k response has a 5xx status code
func (o *AddUserToGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add user to group o k response a status code equal to that given
func (o *AddUserToGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the add user to group o k response
func (o *AddUserToGroupOK) Code() int {
	return 200
}

func (o *AddUserToGroupOK) Error() string {
	return fmt.Sprintf("[POST /iam/addUserToGroup][%d] addUserToGroupOK  %+v", 200, o.Payload)
}

func (o *AddUserToGroupOK) String() string {
	return fmt.Sprintf("[POST /iam/addUserToGroup][%d] addUserToGroupOK  %+v", 200, o.Payload)
}

func (o *AddUserToGroupOK) GetPayload() models.AddUserToGroupResponse {
	return o.Payload
}

func (o *AddUserToGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserToGroupDefault creates a AddUserToGroupDefault with default headers values
func NewAddUserToGroupDefault(code int) *AddUserToGroupDefault {
	return &AddUserToGroupDefault{
		_statusCode: code,
	}
}

/*
AddUserToGroupDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type AddUserToGroupDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this add user to group default response has a 2xx status code
func (o *AddUserToGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this add user to group default response has a 3xx status code
func (o *AddUserToGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this add user to group default response has a 4xx status code
func (o *AddUserToGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this add user to group default response has a 5xx status code
func (o *AddUserToGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this add user to group default response a status code equal to that given
func (o *AddUserToGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the add user to group default response
func (o *AddUserToGroupDefault) Code() int {
	return o._statusCode
}

func (o *AddUserToGroupDefault) Error() string {
	return fmt.Sprintf("[POST /iam/addUserToGroup][%d] addUserToGroup default  %+v", o._statusCode, o.Payload)
}

func (o *AddUserToGroupDefault) String() string {
	return fmt.Sprintf("[POST /iam/addUserToGroup][%d] addUserToGroup default  %+v", o._statusCode, o.Payload)
}

func (o *AddUserToGroupDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddUserToGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
