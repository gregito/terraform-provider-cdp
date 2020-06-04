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

// SuspendWorkspaceReader is a Reader for the SuspendWorkspace structure.
type SuspendWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SuspendWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSuspendWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSuspendWorkspaceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSuspendWorkspaceOK creates a SuspendWorkspaceOK with default headers values
func NewSuspendWorkspaceOK() *SuspendWorkspaceOK {
	return &SuspendWorkspaceOK{}
}

/*SuspendWorkspaceOK handles this case with default header values.

Expected response to a valid request.
*/
type SuspendWorkspaceOK struct {
	Payload models.SuspendWorkspaceResponse
}

func (o *SuspendWorkspaceOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/ml/suspendWorkspace][%d] suspendWorkspaceOK  %+v", 200, o.Payload)
}

func (o *SuspendWorkspaceOK) GetPayload() models.SuspendWorkspaceResponse {
	return o.Payload
}

func (o *SuspendWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSuspendWorkspaceDefault creates a SuspendWorkspaceDefault with default headers values
func NewSuspendWorkspaceDefault(code int) *SuspendWorkspaceDefault {
	return &SuspendWorkspaceDefault{
		_statusCode: code,
	}
}

/*SuspendWorkspaceDefault handles this case with default header values.

The default response on an error.
*/
type SuspendWorkspaceDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the suspend workspace default response
func (o *SuspendWorkspaceDefault) Code() int {
	return o._statusCode
}

func (o *SuspendWorkspaceDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/ml/suspendWorkspace][%d] suspendWorkspace default  %+v", o._statusCode, o.Payload)
}

func (o *SuspendWorkspaceDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *SuspendWorkspaceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
