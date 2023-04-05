// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
)

// ListBackupsReader is a Reader for the ListBackups structure.
type ListBackupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListBackupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListBackupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListBackupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListBackupsOK creates a ListBackupsOK with default headers values
func NewListBackupsOK() *ListBackupsOK {
	return &ListBackupsOK{}
}

/*
ListBackupsOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type ListBackupsOK struct {
	Payload *models.ListBackupsResponse
}

// IsSuccess returns true when this list backups o k response has a 2xx status code
func (o *ListBackupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list backups o k response has a 3xx status code
func (o *ListBackupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list backups o k response has a 4xx status code
func (o *ListBackupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list backups o k response has a 5xx status code
func (o *ListBackupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list backups o k response a status code equal to that given
func (o *ListBackupsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list backups o k response
func (o *ListBackupsOK) Code() int {
	return 200
}

func (o *ListBackupsOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/dw/listBackups][%d] listBackupsOK  %+v", 200, o.Payload)
}

func (o *ListBackupsOK) String() string {
	return fmt.Sprintf("[POST /api/v1/dw/listBackups][%d] listBackupsOK  %+v", 200, o.Payload)
}

func (o *ListBackupsOK) GetPayload() *models.ListBackupsResponse {
	return o.Payload
}

func (o *ListBackupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListBackupsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBackupsDefault creates a ListBackupsDefault with default headers values
func NewListBackupsDefault(code int) *ListBackupsDefault {
	return &ListBackupsDefault{
		_statusCode: code,
	}
}

/*
ListBackupsDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ListBackupsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this list backups default response has a 2xx status code
func (o *ListBackupsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list backups default response has a 3xx status code
func (o *ListBackupsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list backups default response has a 4xx status code
func (o *ListBackupsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list backups default response has a 5xx status code
func (o *ListBackupsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list backups default response a status code equal to that given
func (o *ListBackupsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list backups default response
func (o *ListBackupsDefault) Code() int {
	return o._statusCode
}

func (o *ListBackupsDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/dw/listBackups][%d] listBackups default  %+v", o._statusCode, o.Payload)
}

func (o *ListBackupsDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/dw/listBackups][%d] listBackups default  %+v", o._statusCode, o.Payload)
}

func (o *ListBackupsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListBackupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
