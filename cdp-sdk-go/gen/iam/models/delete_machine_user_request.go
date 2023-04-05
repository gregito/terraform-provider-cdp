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

// DeleteMachineUserRequest Request object for delete machine user request.
//
// swagger:model DeleteMachineUserRequest
type DeleteMachineUserRequest struct {

	// The name or CRN of the machine user to delete.
	// Required: true
	MachineUserName *string `json:"machineUserName"`

	// Whether to delete the machine user immediately (and in rare cases risk orphaning associated resources) or spend more time to fully clean up.
	UnsafeDelete bool `json:"unsafeDelete,omitempty"`
}

// Validate validates this delete machine user request
func (m *DeleteMachineUserRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMachineUserName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteMachineUserRequest) validateMachineUserName(formats strfmt.Registry) error {

	if err := validate.Required("machineUserName", "body", m.MachineUserName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this delete machine user request based on context it is used
func (m *DeleteMachineUserRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteMachineUserRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteMachineUserRequest) UnmarshalBinary(b []byte) error {
	var res DeleteMachineUserRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
