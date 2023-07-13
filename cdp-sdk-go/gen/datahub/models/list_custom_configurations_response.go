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

// ListCustomConfigurationsResponse The response object for the list custom configurations request.
//
// swagger:model ListCustomConfigurationsResponse
type ListCustomConfigurationsResponse struct {

	// The list of custom configurations.
	// Required: true
	CustomConfigurations []*CustomConfigurations `json:"customConfigurations"`

	// The token to use when requesting the next set of results. If not present, there are no additional results.
	NextToken string `json:"nextToken,omitempty"`
}

// Validate validates this list custom configurations response
func (m *ListCustomConfigurationsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomConfigurations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListCustomConfigurationsResponse) validateCustomConfigurations(formats strfmt.Registry) error {

	if err := validate.Required("customConfigurations", "body", m.CustomConfigurations); err != nil {
		return err
	}

	for i := 0; i < len(m.CustomConfigurations); i++ {
		if swag.IsZero(m.CustomConfigurations[i]) { // not required
			continue
		}

		if m.CustomConfigurations[i] != nil {
			if err := m.CustomConfigurations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customConfigurations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customConfigurations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list custom configurations response based on the context it is used
func (m *ListCustomConfigurationsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomConfigurations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListCustomConfigurationsResponse) contextValidateCustomConfigurations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CustomConfigurations); i++ {

		if m.CustomConfigurations[i] != nil {

			if swag.IsZero(m.CustomConfigurations[i]) { // not required
				return nil
			}

			if err := m.CustomConfigurations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customConfigurations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customConfigurations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListCustomConfigurationsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListCustomConfigurationsResponse) UnmarshalBinary(b []byte) error {
	var res ListCustomConfigurationsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
