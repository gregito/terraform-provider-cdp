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

// GetAuditEventsReader is a Reader for the GetAuditEvents structure.
type GetAuditEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuditEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuditEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAuditEventsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAuditEventsOK creates a GetAuditEventsOK with default headers values
func NewGetAuditEventsOK() *GetAuditEventsOK {
	return &GetAuditEventsOK{}
}

/*GetAuditEventsOK handles this case with default header values.

Expected response to a valid request.
*/
type GetAuditEventsOK struct {
	Payload *models.GetAuditEventsResponse
}

func (o *GetAuditEventsOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/ml/getAuditEvents][%d] getAuditEventsOK  %+v", 200, o.Payload)
}

func (o *GetAuditEventsOK) GetPayload() *models.GetAuditEventsResponse {
	return o.Payload
}

func (o *GetAuditEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetAuditEventsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditEventsDefault creates a GetAuditEventsDefault with default headers values
func NewGetAuditEventsDefault(code int) *GetAuditEventsDefault {
	return &GetAuditEventsDefault{
		_statusCode: code,
	}
}

/*GetAuditEventsDefault handles this case with default header values.

The default response on an error.
*/
type GetAuditEventsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get audit events default response
func (o *GetAuditEventsDefault) Code() int {
	return o._statusCode
}

func (o *GetAuditEventsDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/ml/getAuditEvents][%d] getAuditEvents default  %+v", o._statusCode, o.Payload)
}

func (o *GetAuditEventsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAuditEventsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
