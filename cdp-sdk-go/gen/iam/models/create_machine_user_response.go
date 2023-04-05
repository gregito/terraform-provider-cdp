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

// CreateMachineUserResponse Response object for create machine user request.
//
// swagger:model CreateMachineUserResponse
type CreateMachineUserResponse struct {

	// Information about the machine user.
	// Required: true
	MachineUser *MachineUser `json:"machineUser"`
}

// Validate validates this create machine user response
func (m *CreateMachineUserResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMachineUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateMachineUserResponse) validateMachineUser(formats strfmt.Registry) error {

	if err := validate.Required("machineUser", "body", m.MachineUser); err != nil {
		return err
	}

	if m.MachineUser != nil {
		if err := m.MachineUser.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("machineUser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("machineUser")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create machine user response based on the context it is used
func (m *CreateMachineUserResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMachineUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateMachineUserResponse) contextValidateMachineUser(ctx context.Context, formats strfmt.Registry) error {

	if m.MachineUser != nil {
		if err := m.MachineUser.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("machineUser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("machineUser")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateMachineUserResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateMachineUserResponse) UnmarshalBinary(b []byte) error {
	var res CreateMachineUserResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
