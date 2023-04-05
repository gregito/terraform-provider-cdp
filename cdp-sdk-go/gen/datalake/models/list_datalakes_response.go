// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ListDatalakesResponse Response object for list datalakes request.
//
// swagger:model ListDatalakesResponse
type ListDatalakesResponse struct {

	// The datalakes.
	// Required: true
	Datalakes []*Datalake `json:"datalakes"`
}

// Validate validates this list datalakes response
func (m *ListDatalakesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatalakes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListDatalakesResponse) validateDatalakes(formats strfmt.Registry) error {

	if err := validate.Required("datalakes", "body", m.Datalakes); err != nil {
		return err
	}

	for i := 0; i < len(m.Datalakes); i++ {
		if swag.IsZero(m.Datalakes[i]) { // not required
			continue
		}

		if m.Datalakes[i] != nil {
			if err := m.Datalakes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("datalakes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("datalakes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list datalakes response based on the context it is used
func (m *ListDatalakesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDatalakes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListDatalakesResponse) contextValidateDatalakes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Datalakes); i++ {

		if m.Datalakes[i] != nil {
			if err := m.Datalakes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("datalakes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("datalakes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListDatalakesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListDatalakesResponse) UnmarshalBinary(b []byte) error {
	var res ListDatalakesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
