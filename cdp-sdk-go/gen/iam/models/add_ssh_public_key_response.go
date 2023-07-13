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

// AddSSHPublicKeyResponse Response object for add ssh public key.
//
// swagger:model AddSshPublicKeyResponse
type AddSSHPublicKeyResponse struct {

	// Information about the SSH public key.
	// Required: true
	SSHPublicKey *SSHPublicKey `json:"sshPublicKey"`
}

// Validate validates this add Ssh public key response
func (m *AddSSHPublicKeyResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSSHPublicKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddSSHPublicKeyResponse) validateSSHPublicKey(formats strfmt.Registry) error {

	if err := validate.Required("sshPublicKey", "body", m.SSHPublicKey); err != nil {
		return err
	}

	if m.SSHPublicKey != nil {
		if err := m.SSHPublicKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sshPublicKey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sshPublicKey")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this add Ssh public key response based on the context it is used
func (m *AddSSHPublicKeyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSSHPublicKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddSSHPublicKeyResponse) contextValidateSSHPublicKey(ctx context.Context, formats strfmt.Registry) error {

	if m.SSHPublicKey != nil {

		if err := m.SSHPublicKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sshPublicKey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sshPublicKey")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddSSHPublicKeyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddSSHPublicKeyResponse) UnmarshalBinary(b []byte) error {
	var res AddSSHPublicKeyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
