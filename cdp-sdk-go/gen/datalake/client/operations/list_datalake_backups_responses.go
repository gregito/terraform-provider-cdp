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

// ListDatalakeBackupsReader is a Reader for the ListDatalakeBackups structure.
type ListDatalakeBackupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDatalakeBackupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListDatalakeBackupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListDatalakeBackupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListDatalakeBackupsOK creates a ListDatalakeBackupsOK with default headers values
func NewListDatalakeBackupsOK() *ListDatalakeBackupsOK {
	return &ListDatalakeBackupsOK{}
}

/*ListDatalakeBackupsOK handles this case with default header values.

Expected response to a valid request.
*/
type ListDatalakeBackupsOK struct {
	Payload *models.ListDatalakeBackupsResponse
}

func (o *ListDatalakeBackupsOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/datalake/listDatalakeBackups][%d] listDatalakeBackupsOK  %+v", 200, o.Payload)
}

func (o *ListDatalakeBackupsOK) GetPayload() *models.ListDatalakeBackupsResponse {
	return o.Payload
}

func (o *ListDatalakeBackupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListDatalakeBackupsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDatalakeBackupsDefault creates a ListDatalakeBackupsDefault with default headers values
func NewListDatalakeBackupsDefault(code int) *ListDatalakeBackupsDefault {
	return &ListDatalakeBackupsDefault{
		_statusCode: code,
	}
}

/*ListDatalakeBackupsDefault handles this case with default header values.

The default response on an error.
*/
type ListDatalakeBackupsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list datalake backups default response
func (o *ListDatalakeBackupsDefault) Code() int {
	return o._statusCode
}

func (o *ListDatalakeBackupsDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/datalake/listDatalakeBackups][%d] listDatalakeBackups default  %+v", o._statusCode, o.Payload)
}

func (o *ListDatalakeBackupsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListDatalakeBackupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
