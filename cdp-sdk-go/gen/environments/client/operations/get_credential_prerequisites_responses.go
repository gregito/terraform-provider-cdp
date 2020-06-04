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

// GetCredentialPrerequisitesReader is a Reader for the GetCredentialPrerequisites structure.
type GetCredentialPrerequisitesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCredentialPrerequisitesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCredentialPrerequisitesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCredentialPrerequisitesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCredentialPrerequisitesOK creates a GetCredentialPrerequisitesOK with default headers values
func NewGetCredentialPrerequisitesOK() *GetCredentialPrerequisitesOK {
	return &GetCredentialPrerequisitesOK{}
}

/*GetCredentialPrerequisitesOK handles this case with default header values.

Expected response to a valid request.
*/
type GetCredentialPrerequisitesOK struct {
	Payload *models.GetCredentialPrerequisitesResponse
}

func (o *GetCredentialPrerequisitesOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/getCredentialPrerequisites][%d] getCredentialPrerequisitesOK  %+v", 200, o.Payload)
}

func (o *GetCredentialPrerequisitesOK) GetPayload() *models.GetCredentialPrerequisitesResponse {
	return o.Payload
}

func (o *GetCredentialPrerequisitesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCredentialPrerequisitesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCredentialPrerequisitesDefault creates a GetCredentialPrerequisitesDefault with default headers values
func NewGetCredentialPrerequisitesDefault(code int) *GetCredentialPrerequisitesDefault {
	return &GetCredentialPrerequisitesDefault{
		_statusCode: code,
	}
}

/*GetCredentialPrerequisitesDefault handles this case with default header values.

The default response on an error.
*/
type GetCredentialPrerequisitesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get credential prerequisites default response
func (o *GetCredentialPrerequisitesDefault) Code() int {
	return o._statusCode
}

func (o *GetCredentialPrerequisitesDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/getCredentialPrerequisites][%d] getCredentialPrerequisites default  %+v", o._statusCode, o.Payload)
}

func (o *GetCredentialPrerequisitesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCredentialPrerequisitesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
