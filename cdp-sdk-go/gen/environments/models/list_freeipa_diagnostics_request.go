// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ListFreeipaDiagnosticsRequest Request object for listing recent FreeIPA diagnostics collections.
//
// swagger:model ListFreeipaDiagnosticsRequest
type ListFreeipaDiagnosticsRequest struct {

	// The name or CRN of the environment.
	// Required: true
	EnvironmentName *string `json:"environmentName"`
}

// Validate validates this list freeipa diagnostics request
func (m *ListFreeipaDiagnosticsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironmentName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListFreeipaDiagnosticsRequest) validateEnvironmentName(formats strfmt.Registry) error {

	if err := validate.Required("environmentName", "body", m.EnvironmentName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this list freeipa diagnostics request based on context it is used
func (m *ListFreeipaDiagnosticsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListFreeipaDiagnosticsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListFreeipaDiagnosticsRequest) UnmarshalBinary(b []byte) error {
	var res ListFreeipaDiagnosticsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
