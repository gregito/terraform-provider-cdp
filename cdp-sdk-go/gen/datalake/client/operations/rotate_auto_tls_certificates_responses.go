// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
)

// RotateAutoTLSCertificatesReader is a Reader for the RotateAutoTLSCertificates structure.
type RotateAutoTLSCertificatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RotateAutoTLSCertificatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRotateAutoTLSCertificatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRotateAutoTLSCertificatesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRotateAutoTLSCertificatesOK creates a RotateAutoTLSCertificatesOK with default headers values
func NewRotateAutoTLSCertificatesOK() *RotateAutoTLSCertificatesOK {
	return &RotateAutoTLSCertificatesOK{}
}

/*
RotateAutoTLSCertificatesOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type RotateAutoTLSCertificatesOK struct {
	Payload *models.RotateAutoTLSCertificatesResponse
}

// IsSuccess returns true when this rotate auto Tls certificates o k response has a 2xx status code
func (o *RotateAutoTLSCertificatesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this rotate auto Tls certificates o k response has a 3xx status code
func (o *RotateAutoTLSCertificatesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rotate auto Tls certificates o k response has a 4xx status code
func (o *RotateAutoTLSCertificatesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this rotate auto Tls certificates o k response has a 5xx status code
func (o *RotateAutoTLSCertificatesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this rotate auto Tls certificates o k response a status code equal to that given
func (o *RotateAutoTLSCertificatesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the rotate auto Tls certificates o k response
func (o *RotateAutoTLSCertificatesOK) Code() int {
	return 200
}

func (o *RotateAutoTLSCertificatesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/rotateAutoTlsCertificates][%d] rotateAutoTlsCertificatesOK %s", 200, payload)
}

func (o *RotateAutoTLSCertificatesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/rotateAutoTlsCertificates][%d] rotateAutoTlsCertificatesOK %s", 200, payload)
}

func (o *RotateAutoTLSCertificatesOK) GetPayload() *models.RotateAutoTLSCertificatesResponse {
	return o.Payload
}

func (o *RotateAutoTLSCertificatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RotateAutoTLSCertificatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRotateAutoTLSCertificatesDefault creates a RotateAutoTLSCertificatesDefault with default headers values
func NewRotateAutoTLSCertificatesDefault(code int) *RotateAutoTLSCertificatesDefault {
	return &RotateAutoTLSCertificatesDefault{
		_statusCode: code,
	}
}

/*
RotateAutoTLSCertificatesDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type RotateAutoTLSCertificatesDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this rotate auto Tls certificates default response has a 2xx status code
func (o *RotateAutoTLSCertificatesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this rotate auto Tls certificates default response has a 3xx status code
func (o *RotateAutoTLSCertificatesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this rotate auto Tls certificates default response has a 4xx status code
func (o *RotateAutoTLSCertificatesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this rotate auto Tls certificates default response has a 5xx status code
func (o *RotateAutoTLSCertificatesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this rotate auto Tls certificates default response a status code equal to that given
func (o *RotateAutoTLSCertificatesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the rotate auto Tls certificates default response
func (o *RotateAutoTLSCertificatesDefault) Code() int {
	return o._statusCode
}

func (o *RotateAutoTLSCertificatesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/rotateAutoTlsCertificates][%d] rotateAutoTlsCertificates default %s", o._statusCode, payload)
}

func (o *RotateAutoTLSCertificatesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/rotateAutoTlsCertificates][%d] rotateAutoTlsCertificates default %s", o._statusCode, payload)
}

func (o *RotateAutoTLSCertificatesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *RotateAutoTLSCertificatesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
