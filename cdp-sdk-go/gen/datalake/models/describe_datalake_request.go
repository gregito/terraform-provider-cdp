// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DescribeDatalakeRequest Request object for describe datalake request.
//
// swagger:model DescribeDatalakeRequest
type DescribeDatalakeRequest struct {

	// The name or CRN of the datalake.
	// Required: true
	DatalakeName *string `json:"datalakeName"`
}

// Validate validates this describe datalake request
func (m *DescribeDatalakeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatalakeName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DescribeDatalakeRequest) validateDatalakeName(formats strfmt.Registry) error {

	if err := validate.Required("datalakeName", "body", m.DatalakeName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DescribeDatalakeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DescribeDatalakeRequest) UnmarshalBinary(b []byte) error {
	var res DescribeDatalakeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
