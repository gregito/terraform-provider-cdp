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

// RetryClusterReader is a Reader for the RetryCluster structure.
type RetryClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetryClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRetryClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRetryClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRetryClusterOK creates a RetryClusterOK with default headers values
func NewRetryClusterOK() *RetryClusterOK {
	return &RetryClusterOK{}
}

/*RetryClusterOK handles this case with default header values.

Expected response to a valid request.
*/
type RetryClusterOK struct {
	Payload models.RetryClusterResponse
}

func (o *RetryClusterOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/datahub/retryCluster][%d] retryClusterOK  %+v", 200, o.Payload)
}

func (o *RetryClusterOK) GetPayload() models.RetryClusterResponse {
	return o.Payload
}

func (o *RetryClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetryClusterDefault creates a RetryClusterDefault with default headers values
func NewRetryClusterDefault(code int) *RetryClusterDefault {
	return &RetryClusterDefault{
		_statusCode: code,
	}
}

/*RetryClusterDefault handles this case with default header values.

The default response on an error.
*/
type RetryClusterDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the retry cluster default response
func (o *RetryClusterDefault) Code() int {
	return o._statusCode
}

func (o *RetryClusterDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/datahub/retryCluster][%d] retryCluster default  %+v", o._statusCode, o.Payload)
}

func (o *RetryClusterDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *RetryClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
