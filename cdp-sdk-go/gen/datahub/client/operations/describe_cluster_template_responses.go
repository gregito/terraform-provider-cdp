// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
)

// DescribeClusterTemplateReader is a Reader for the DescribeClusterTemplate structure.
type DescribeClusterTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeClusterTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeClusterTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDescribeClusterTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeClusterTemplateOK creates a DescribeClusterTemplateOK with default headers values
func NewDescribeClusterTemplateOK() *DescribeClusterTemplateOK {
	return &DescribeClusterTemplateOK{}
}

/*
DescribeClusterTemplateOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type DescribeClusterTemplateOK struct {
	Payload *models.DescribeClusterTemplateResponse
}

// IsSuccess returns true when this describe cluster template o k response has a 2xx status code
func (o *DescribeClusterTemplateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe cluster template o k response has a 3xx status code
func (o *DescribeClusterTemplateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe cluster template o k response has a 4xx status code
func (o *DescribeClusterTemplateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe cluster template o k response has a 5xx status code
func (o *DescribeClusterTemplateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe cluster template o k response a status code equal to that given
func (o *DescribeClusterTemplateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the describe cluster template o k response
func (o *DescribeClusterTemplateOK) Code() int {
	return 200
}

func (o *DescribeClusterTemplateOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/datahub/describeClusterTemplate][%d] describeClusterTemplateOK  %+v", 200, o.Payload)
}

func (o *DescribeClusterTemplateOK) String() string {
	return fmt.Sprintf("[POST /api/v1/datahub/describeClusterTemplate][%d] describeClusterTemplateOK  %+v", 200, o.Payload)
}

func (o *DescribeClusterTemplateOK) GetPayload() *models.DescribeClusterTemplateResponse {
	return o.Payload
}

func (o *DescribeClusterTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DescribeClusterTemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeClusterTemplateDefault creates a DescribeClusterTemplateDefault with default headers values
func NewDescribeClusterTemplateDefault(code int) *DescribeClusterTemplateDefault {
	return &DescribeClusterTemplateDefault{
		_statusCode: code,
	}
}

/*
DescribeClusterTemplateDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type DescribeClusterTemplateDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this describe cluster template default response has a 2xx status code
func (o *DescribeClusterTemplateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this describe cluster template default response has a 3xx status code
func (o *DescribeClusterTemplateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this describe cluster template default response has a 4xx status code
func (o *DescribeClusterTemplateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this describe cluster template default response has a 5xx status code
func (o *DescribeClusterTemplateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this describe cluster template default response a status code equal to that given
func (o *DescribeClusterTemplateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the describe cluster template default response
func (o *DescribeClusterTemplateDefault) Code() int {
	return o._statusCode
}

func (o *DescribeClusterTemplateDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/datahub/describeClusterTemplate][%d] describeClusterTemplate default  %+v", o._statusCode, o.Payload)
}

func (o *DescribeClusterTemplateDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/datahub/describeClusterTemplate][%d] describeClusterTemplate default  %+v", o._statusCode, o.Payload)
}

func (o *DescribeClusterTemplateDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DescribeClusterTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
