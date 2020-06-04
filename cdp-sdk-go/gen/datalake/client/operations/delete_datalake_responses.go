// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
)

// DeleteDatalakeReader is a Reader for the DeleteDatalake structure.
type DeleteDatalakeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDatalakeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDatalakeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteDatalakeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteDatalakeOK creates a DeleteDatalakeOK with default headers values
func NewDeleteDatalakeOK() *DeleteDatalakeOK {
	return &DeleteDatalakeOK{}
}

/*DeleteDatalakeOK handles this case with default header values.

Expected response to a valid request.
*/
type DeleteDatalakeOK struct {
	Payload models.DeleteDatalakeResponse
}

func (o *DeleteDatalakeOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/datalake/deleteDatalake][%d] deleteDatalakeOK  %+v", 200, o.Payload)
}

func (o *DeleteDatalakeOK) GetPayload() models.DeleteDatalakeResponse {
	return o.Payload
}

func (o *DeleteDatalakeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDatalakeDefault creates a DeleteDatalakeDefault with default headers values
func NewDeleteDatalakeDefault(code int) *DeleteDatalakeDefault {
	return &DeleteDatalakeDefault{
		_statusCode: code,
	}
}

/*DeleteDatalakeDefault handles this case with default header values.

The default response on an error.
*/
type DeleteDatalakeDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete datalake default response
func (o *DeleteDatalakeDefault) Code() int {
	return o._statusCode
}

func (o *DeleteDatalakeDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/datalake/deleteDatalake][%d] deleteDatalake default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteDatalakeDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDatalakeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
