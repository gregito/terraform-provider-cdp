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

// SetAWSAuditCredentialReader is a Reader for the SetAWSAuditCredential structure.
type SetAWSAuditCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetAWSAuditCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetAWSAuditCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSetAWSAuditCredentialDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSetAWSAuditCredentialOK creates a SetAWSAuditCredentialOK with default headers values
func NewSetAWSAuditCredentialOK() *SetAWSAuditCredentialOK {
	return &SetAWSAuditCredentialOK{}
}

/*
SetAWSAuditCredentialOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type SetAWSAuditCredentialOK struct {
	Payload *models.SetAWSAuditCredentialResponse
}

// IsSuccess returns true when this set a w s audit credential o k response has a 2xx status code
func (o *SetAWSAuditCredentialOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this set a w s audit credential o k response has a 3xx status code
func (o *SetAWSAuditCredentialOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set a w s audit credential o k response has a 4xx status code
func (o *SetAWSAuditCredentialOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this set a w s audit credential o k response has a 5xx status code
func (o *SetAWSAuditCredentialOK) IsServerError() bool {
	return false
}

// IsCode returns true when this set a w s audit credential o k response a status code equal to that given
func (o *SetAWSAuditCredentialOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the set a w s audit credential o k response
func (o *SetAWSAuditCredentialOK) Code() int {
	return 200
}

func (o *SetAWSAuditCredentialOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/setAWSAuditCredential][%d] setAWSAuditCredentialOK  %+v", 200, o.Payload)
}

func (o *SetAWSAuditCredentialOK) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/setAWSAuditCredential][%d] setAWSAuditCredentialOK  %+v", 200, o.Payload)
}

func (o *SetAWSAuditCredentialOK) GetPayload() *models.SetAWSAuditCredentialResponse {
	return o.Payload
}

func (o *SetAWSAuditCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SetAWSAuditCredentialResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetAWSAuditCredentialDefault creates a SetAWSAuditCredentialDefault with default headers values
func NewSetAWSAuditCredentialDefault(code int) *SetAWSAuditCredentialDefault {
	return &SetAWSAuditCredentialDefault{
		_statusCode: code,
	}
}

/*
SetAWSAuditCredentialDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type SetAWSAuditCredentialDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this set a w s audit credential default response has a 2xx status code
func (o *SetAWSAuditCredentialDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this set a w s audit credential default response has a 3xx status code
func (o *SetAWSAuditCredentialDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this set a w s audit credential default response has a 4xx status code
func (o *SetAWSAuditCredentialDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this set a w s audit credential default response has a 5xx status code
func (o *SetAWSAuditCredentialDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this set a w s audit credential default response a status code equal to that given
func (o *SetAWSAuditCredentialDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the set a w s audit credential default response
func (o *SetAWSAuditCredentialDefault) Code() int {
	return o._statusCode
}

func (o *SetAWSAuditCredentialDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/setAWSAuditCredential][%d] setAWSAuditCredential default  %+v", o._statusCode, o.Payload)
}

func (o *SetAWSAuditCredentialDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/setAWSAuditCredential][%d] setAWSAuditCredential default  %+v", o._statusCode, o.Payload)
}

func (o *SetAWSAuditCredentialDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetAWSAuditCredentialDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
