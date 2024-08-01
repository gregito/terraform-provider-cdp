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

// GetFreeipaUpgradeOptionsRequest The request object for retrieving FreeIPA upgrade candidates.
//
// swagger:model GetFreeipaUpgradeOptionsRequest
type GetFreeipaUpgradeOptionsRequest struct {

	// Allows the upgrade to a subsequent major OS version in the series.
	AllowMajorOsUpgrade bool `json:"allowMajorOsUpgrade,omitempty"`

	// The URL of the source image catalog. If not specify this option we'll use image catalog of the current image.
	Catalog string `json:"catalog,omitempty"`

	// The name or the CRN of the environment.
	// Required: true
	Environment *string `json:"environment"`
}

// Validate validates this get freeipa upgrade options request
func (m *GetFreeipaUpgradeOptionsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetFreeipaUpgradeOptionsRequest) validateEnvironment(formats strfmt.Registry) error {

	if err := validate.Required("environment", "body", m.Environment); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get freeipa upgrade options request based on context it is used
func (m *GetFreeipaUpgradeOptionsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetFreeipaUpgradeOptionsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetFreeipaUpgradeOptionsRequest) UnmarshalBinary(b []byte) error {
	var res GetFreeipaUpgradeOptionsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
