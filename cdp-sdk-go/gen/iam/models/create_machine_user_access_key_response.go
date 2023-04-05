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

// CreateMachineUserAccessKeyResponse Response object for a create machine user access key request.
//
// swagger:model CreateMachineUserAccessKeyResponse
type CreateMachineUserAccessKeyResponse struct {

	// The access key that was created.
	// Required: true
	AccessKey *AccessKey `json:"accessKey"`

	// The private key associated with this access key. This string is the contents of a PEM file containing a PKCS#8 private key.
	// Required: true
	PrivateKey *string `json:"privateKey"`
}

// Validate validates this create machine user access key response
func (m *CreateMachineUserAccessKeyResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateMachineUserAccessKeyResponse) validateAccessKey(formats strfmt.Registry) error {

	if err := validate.Required("accessKey", "body", m.AccessKey); err != nil {
		return err
	}

	if m.AccessKey != nil {
		if err := m.AccessKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessKey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessKey")
			}
			return err
		}
	}

	return nil
}

func (m *CreateMachineUserAccessKeyResponse) validatePrivateKey(formats strfmt.Registry) error {

	if err := validate.Required("privateKey", "body", m.PrivateKey); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create machine user access key response based on the context it is used
func (m *CreateMachineUserAccessKeyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateMachineUserAccessKeyResponse) contextValidateAccessKey(ctx context.Context, formats strfmt.Registry) error {

	if m.AccessKey != nil {
		if err := m.AccessKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessKey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessKey")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateMachineUserAccessKeyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateMachineUserAccessKeyResponse) UnmarshalBinary(b []byte) error {
	var res CreateMachineUserAccessKeyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
