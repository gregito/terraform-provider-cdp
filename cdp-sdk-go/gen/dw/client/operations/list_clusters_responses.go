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

// ListClustersReader is a Reader for the ListClusters structure.
type ListClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListClustersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListClustersOK creates a ListClustersOK with default headers values
func NewListClustersOK() *ListClustersOK {
	return &ListClustersOK{}
}

/*
ListClustersOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type ListClustersOK struct {
	Payload *models.ListClustersResponse
}

// IsSuccess returns true when this list clusters o k response has a 2xx status code
func (o *ListClustersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list clusters o k response has a 3xx status code
func (o *ListClustersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list clusters o k response has a 4xx status code
func (o *ListClustersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list clusters o k response has a 5xx status code
func (o *ListClustersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list clusters o k response a status code equal to that given
func (o *ListClustersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list clusters o k response
func (o *ListClustersOK) Code() int {
	return 200
}

func (o *ListClustersOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/dw/listClusters][%d] listClustersOK  %+v", 200, o.Payload)
}

func (o *ListClustersOK) String() string {
	return fmt.Sprintf("[POST /api/v1/dw/listClusters][%d] listClustersOK  %+v", 200, o.Payload)
}

func (o *ListClustersOK) GetPayload() *models.ListClustersResponse {
	return o.Payload
}

func (o *ListClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListClustersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClustersDefault creates a ListClustersDefault with default headers values
func NewListClustersDefault(code int) *ListClustersDefault {
	return &ListClustersDefault{
		_statusCode: code,
	}
}

/*
ListClustersDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ListClustersDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this list clusters default response has a 2xx status code
func (o *ListClustersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list clusters default response has a 3xx status code
func (o *ListClustersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list clusters default response has a 4xx status code
func (o *ListClustersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list clusters default response has a 5xx status code
func (o *ListClustersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list clusters default response a status code equal to that given
func (o *ListClustersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list clusters default response
func (o *ListClustersDefault) Code() int {
	return o._statusCode
}

func (o *ListClustersDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/dw/listClusters][%d] listClusters default  %+v", o._statusCode, o.Payload)
}

func (o *ListClustersDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/dw/listClusters][%d] listClusters default  %+v", o._statusCode, o.Payload)
}

func (o *ListClustersDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListClustersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
