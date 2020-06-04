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

// AssignServicePrincipalAzureCloudIdentityReader is a Reader for the AssignServicePrincipalAzureCloudIdentity structure.
type AssignServicePrincipalAzureCloudIdentityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssignServicePrincipalAzureCloudIdentityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssignServicePrincipalAzureCloudIdentityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAssignServicePrincipalAzureCloudIdentityDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAssignServicePrincipalAzureCloudIdentityOK creates a AssignServicePrincipalAzureCloudIdentityOK with default headers values
func NewAssignServicePrincipalAzureCloudIdentityOK() *AssignServicePrincipalAzureCloudIdentityOK {
	return &AssignServicePrincipalAzureCloudIdentityOK{}
}

/*AssignServicePrincipalAzureCloudIdentityOK handles this case with default header values.

Expected response to a valid request.
*/
type AssignServicePrincipalAzureCloudIdentityOK struct {
	Payload models.AssignServicePrincipalAzureCloudIdentityResponse
}

func (o *AssignServicePrincipalAzureCloudIdentityOK) Error() string {
	return fmt.Sprintf("[POST /iam/assignServicePrincipalAzureCloudIdentity][%d] assignServicePrincipalAzureCloudIdentityOK  %+v", 200, o.Payload)
}

func (o *AssignServicePrincipalAzureCloudIdentityOK) GetPayload() models.AssignServicePrincipalAzureCloudIdentityResponse {
	return o.Payload
}

func (o *AssignServicePrincipalAzureCloudIdentityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignServicePrincipalAzureCloudIdentityDefault creates a AssignServicePrincipalAzureCloudIdentityDefault with default headers values
func NewAssignServicePrincipalAzureCloudIdentityDefault(code int) *AssignServicePrincipalAzureCloudIdentityDefault {
	return &AssignServicePrincipalAzureCloudIdentityDefault{
		_statusCode: code,
	}
}

/*AssignServicePrincipalAzureCloudIdentityDefault handles this case with default header values.

The default response on an error.
*/
type AssignServicePrincipalAzureCloudIdentityDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the assign service principal azure cloud identity default response
func (o *AssignServicePrincipalAzureCloudIdentityDefault) Code() int {
	return o._statusCode
}

func (o *AssignServicePrincipalAzureCloudIdentityDefault) Error() string {
	return fmt.Sprintf("[POST /iam/assignServicePrincipalAzureCloudIdentity][%d] assignServicePrincipalAzureCloudIdentity default  %+v", o._statusCode, o.Payload)
}

func (o *AssignServicePrincipalAzureCloudIdentityDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *AssignServicePrincipalAzureCloudIdentityDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
