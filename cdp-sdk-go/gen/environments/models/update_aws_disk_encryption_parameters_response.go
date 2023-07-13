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

// UpdateAwsDiskEncryptionParametersResponse Response object for an update AWS encryption parameters request.
//
// swagger:model UpdateAwsDiskEncryptionParametersResponse
type UpdateAwsDiskEncryptionParametersResponse struct {

	// Object containing details of encryption parameters for AWS cloud.
	AwsDiskEncryptionParameters *AwsDiskEncryptionParameters `json:"awsDiskEncryptionParameters,omitempty"`

	// The environment summary.
	// Required: true
	Environment *Environment `json:"environment"`
}

// Validate validates this update aws disk encryption parameters response
func (m *UpdateAwsDiskEncryptionParametersResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsDiskEncryptionParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateAwsDiskEncryptionParametersResponse) validateAwsDiskEncryptionParameters(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsDiskEncryptionParameters) { // not required
		return nil
	}

	if m.AwsDiskEncryptionParameters != nil {
		if err := m.AwsDiskEncryptionParameters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsDiskEncryptionParameters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsDiskEncryptionParameters")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateAwsDiskEncryptionParametersResponse) validateEnvironment(formats strfmt.Registry) error {

	if err := validate.Required("environment", "body", m.Environment); err != nil {
		return err
	}

	if m.Environment != nil {
		if err := m.Environment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("environment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("environment")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update aws disk encryption parameters response based on the context it is used
func (m *UpdateAwsDiskEncryptionParametersResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAwsDiskEncryptionParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnvironment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateAwsDiskEncryptionParametersResponse) contextValidateAwsDiskEncryptionParameters(ctx context.Context, formats strfmt.Registry) error {

	if m.AwsDiskEncryptionParameters != nil {

		if swag.IsZero(m.AwsDiskEncryptionParameters) { // not required
			return nil
		}

		if err := m.AwsDiskEncryptionParameters.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsDiskEncryptionParameters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsDiskEncryptionParameters")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateAwsDiskEncryptionParametersResponse) contextValidateEnvironment(ctx context.Context, formats strfmt.Registry) error {

	if m.Environment != nil {

		if err := m.Environment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("environment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("environment")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateAwsDiskEncryptionParametersResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateAwsDiskEncryptionParametersResponse) UnmarshalBinary(b []byte) error {
	var res UpdateAwsDiskEncryptionParametersResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
