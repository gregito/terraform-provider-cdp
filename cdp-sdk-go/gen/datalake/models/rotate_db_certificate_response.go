// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RotateDbCertificateResponse Response object to rotate ssl certificate a datalake.
//
// swagger:model RotateDbCertificateResponse
type RotateDbCertificateResponse struct {

	// Unique operation ID assigned to this command execution. Use this identifier with 'get-operation' to track status and retrieve detailed results.
	OperationID string `json:"operationId,omitempty"`
}

// Validate validates this rotate db certificate response
func (m *RotateDbCertificateResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this rotate db certificate response based on context it is used
func (m *RotateDbCertificateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RotateDbCertificateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RotateDbCertificateResponse) UnmarshalBinary(b []byte) error {
	var res RotateDbCertificateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
