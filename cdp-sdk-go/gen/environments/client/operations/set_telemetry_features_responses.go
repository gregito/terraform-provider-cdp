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

// SetTelemetryFeaturesReader is a Reader for the SetTelemetryFeatures structure.
type SetTelemetryFeaturesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetTelemetryFeaturesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetTelemetryFeaturesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSetTelemetryFeaturesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSetTelemetryFeaturesOK creates a SetTelemetryFeaturesOK with default headers values
func NewSetTelemetryFeaturesOK() *SetTelemetryFeaturesOK {
	return &SetTelemetryFeaturesOK{}
}

/*
SetTelemetryFeaturesOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type SetTelemetryFeaturesOK struct {
	Payload models.SetTelemetryFeaturesResponse
}

// IsSuccess returns true when this set telemetry features o k response has a 2xx status code
func (o *SetTelemetryFeaturesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this set telemetry features o k response has a 3xx status code
func (o *SetTelemetryFeaturesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set telemetry features o k response has a 4xx status code
func (o *SetTelemetryFeaturesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this set telemetry features o k response has a 5xx status code
func (o *SetTelemetryFeaturesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this set telemetry features o k response a status code equal to that given
func (o *SetTelemetryFeaturesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the set telemetry features o k response
func (o *SetTelemetryFeaturesOK) Code() int {
	return 200
}

func (o *SetTelemetryFeaturesOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/setTelemetryFeatures][%d] setTelemetryFeaturesOK  %+v", 200, o.Payload)
}

func (o *SetTelemetryFeaturesOK) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/setTelemetryFeatures][%d] setTelemetryFeaturesOK  %+v", 200, o.Payload)
}

func (o *SetTelemetryFeaturesOK) GetPayload() models.SetTelemetryFeaturesResponse {
	return o.Payload
}

func (o *SetTelemetryFeaturesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetTelemetryFeaturesDefault creates a SetTelemetryFeaturesDefault with default headers values
func NewSetTelemetryFeaturesDefault(code int) *SetTelemetryFeaturesDefault {
	return &SetTelemetryFeaturesDefault{
		_statusCode: code,
	}
}

/*
SetTelemetryFeaturesDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type SetTelemetryFeaturesDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this set telemetry features default response has a 2xx status code
func (o *SetTelemetryFeaturesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this set telemetry features default response has a 3xx status code
func (o *SetTelemetryFeaturesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this set telemetry features default response has a 4xx status code
func (o *SetTelemetryFeaturesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this set telemetry features default response has a 5xx status code
func (o *SetTelemetryFeaturesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this set telemetry features default response a status code equal to that given
func (o *SetTelemetryFeaturesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the set telemetry features default response
func (o *SetTelemetryFeaturesDefault) Code() int {
	return o._statusCode
}

func (o *SetTelemetryFeaturesDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/setTelemetryFeatures][%d] setTelemetryFeatures default  %+v", o._statusCode, o.Payload)
}

func (o *SetTelemetryFeaturesDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/setTelemetryFeatures][%d] setTelemetryFeatures default  %+v", o._statusCode, o.Payload)
}

func (o *SetTelemetryFeaturesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetTelemetryFeaturesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
