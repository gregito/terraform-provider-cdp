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

// SyncIDBrokerMappingsReader is a Reader for the SyncIDBrokerMappings structure.
type SyncIDBrokerMappingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SyncIDBrokerMappingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSyncIDBrokerMappingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSyncIDBrokerMappingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSyncIDBrokerMappingsOK creates a SyncIDBrokerMappingsOK with default headers values
func NewSyncIDBrokerMappingsOK() *SyncIDBrokerMappingsOK {
	return &SyncIDBrokerMappingsOK{}
}

/*
SyncIDBrokerMappingsOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type SyncIDBrokerMappingsOK struct {
	Payload models.SyncIDBrokerMappingsResponse
}

// IsSuccess returns true when this sync Id broker mappings o k response has a 2xx status code
func (o *SyncIDBrokerMappingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this sync Id broker mappings o k response has a 3xx status code
func (o *SyncIDBrokerMappingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sync Id broker mappings o k response has a 4xx status code
func (o *SyncIDBrokerMappingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this sync Id broker mappings o k response has a 5xx status code
func (o *SyncIDBrokerMappingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this sync Id broker mappings o k response a status code equal to that given
func (o *SyncIDBrokerMappingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the sync Id broker mappings o k response
func (o *SyncIDBrokerMappingsOK) Code() int {
	return 200
}

func (o *SyncIDBrokerMappingsOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/syncIdBrokerMappings][%d] syncIdBrokerMappingsOK  %+v", 200, o.Payload)
}

func (o *SyncIDBrokerMappingsOK) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/syncIdBrokerMappings][%d] syncIdBrokerMappingsOK  %+v", 200, o.Payload)
}

func (o *SyncIDBrokerMappingsOK) GetPayload() models.SyncIDBrokerMappingsResponse {
	return o.Payload
}

func (o *SyncIDBrokerMappingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSyncIDBrokerMappingsDefault creates a SyncIDBrokerMappingsDefault with default headers values
func NewSyncIDBrokerMappingsDefault(code int) *SyncIDBrokerMappingsDefault {
	return &SyncIDBrokerMappingsDefault{
		_statusCode: code,
	}
}

/*
SyncIDBrokerMappingsDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type SyncIDBrokerMappingsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this sync Id broker mappings default response has a 2xx status code
func (o *SyncIDBrokerMappingsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this sync Id broker mappings default response has a 3xx status code
func (o *SyncIDBrokerMappingsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this sync Id broker mappings default response has a 4xx status code
func (o *SyncIDBrokerMappingsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this sync Id broker mappings default response has a 5xx status code
func (o *SyncIDBrokerMappingsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this sync Id broker mappings default response a status code equal to that given
func (o *SyncIDBrokerMappingsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the sync Id broker mappings default response
func (o *SyncIDBrokerMappingsDefault) Code() int {
	return o._statusCode
}

func (o *SyncIDBrokerMappingsDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/syncIdBrokerMappings][%d] syncIdBrokerMappings default  %+v", o._statusCode, o.Payload)
}

func (o *SyncIDBrokerMappingsDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/syncIdBrokerMappings][%d] syncIdBrokerMappings default  %+v", o._statusCode, o.Payload)
}

func (o *SyncIDBrokerMappingsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *SyncIDBrokerMappingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
