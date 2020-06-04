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

// BackupDatalakeReader is a Reader for the BackupDatalake structure.
type BackupDatalakeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BackupDatalakeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBackupDatalakeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewBackupDatalakeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBackupDatalakeOK creates a BackupDatalakeOK with default headers values
func NewBackupDatalakeOK() *BackupDatalakeOK {
	return &BackupDatalakeOK{}
}

/*BackupDatalakeOK handles this case with default header values.

Expected response to a valid request.
*/
type BackupDatalakeOK struct {
	Payload *models.BackupDatalakeResponse
}

func (o *BackupDatalakeOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/datalake/backupDatalake][%d] backupDatalakeOK  %+v", 200, o.Payload)
}

func (o *BackupDatalakeOK) GetPayload() *models.BackupDatalakeResponse {
	return o.Payload
}

func (o *BackupDatalakeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BackupDatalakeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupDatalakeDefault creates a BackupDatalakeDefault with default headers values
func NewBackupDatalakeDefault(code int) *BackupDatalakeDefault {
	return &BackupDatalakeDefault{
		_statusCode: code,
	}
}

/*BackupDatalakeDefault handles this case with default header values.

The default response on an error.
*/
type BackupDatalakeDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the backup datalake default response
func (o *BackupDatalakeDefault) Code() int {
	return o._statusCode
}

func (o *BackupDatalakeDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/datalake/backupDatalake][%d] backupDatalake default  %+v", o._statusCode, o.Payload)
}

func (o *BackupDatalakeDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *BackupDatalakeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
