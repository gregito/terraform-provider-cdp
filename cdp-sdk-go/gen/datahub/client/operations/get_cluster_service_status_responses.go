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

// GetClusterServiceStatusReader is a Reader for the GetClusterServiceStatus structure.
type GetClusterServiceStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterServiceStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterServiceStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetClusterServiceStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterServiceStatusOK creates a GetClusterServiceStatusOK with default headers values
func NewGetClusterServiceStatusOK() *GetClusterServiceStatusOK {
	return &GetClusterServiceStatusOK{}
}

/*GetClusterServiceStatusOK handles this case with default header values.

Expected response to a valid request.
*/
type GetClusterServiceStatusOK struct {
	Payload *models.GetClusterServiceStatusResponse
}

func (o *GetClusterServiceStatusOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/datahub/getClusterServiceStatus][%d] getClusterServiceStatusOK  %+v", 200, o.Payload)
}

func (o *GetClusterServiceStatusOK) GetPayload() *models.GetClusterServiceStatusResponse {
	return o.Payload
}

func (o *GetClusterServiceStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetClusterServiceStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterServiceStatusDefault creates a GetClusterServiceStatusDefault with default headers values
func NewGetClusterServiceStatusDefault(code int) *GetClusterServiceStatusDefault {
	return &GetClusterServiceStatusDefault{
		_statusCode: code,
	}
}

/*GetClusterServiceStatusDefault handles this case with default header values.

The default response on an error.
*/
type GetClusterServiceStatusDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get cluster service status default response
func (o *GetClusterServiceStatusDefault) Code() int {
	return o._statusCode
}

func (o *GetClusterServiceStatusDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/datahub/getClusterServiceStatus][%d] getClusterServiceStatus default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterServiceStatusDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetClusterServiceStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
