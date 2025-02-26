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

// UpdateGcpAvailabilityZonesRequest Request object to update Availability Zones for GCP environment.
//
// swagger:model UpdateGcpAvailabilityZonesRequest
type UpdateGcpAvailabilityZonesRequest struct {

	// List of availability zones for the environment.
	// Required: true
	AvailabilityZones []string `json:"availabilityZones"`

	// The name or CRN of the environment.
	// Required: true
	Environment *string `json:"environment"`
}

// Validate validates this update gcp availability zones request
func (m *UpdateGcpAvailabilityZonesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailabilityZones(formats); err != nil {
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

func (m *UpdateGcpAvailabilityZonesRequest) validateAvailabilityZones(formats strfmt.Registry) error {

	if err := validate.Required("availabilityZones", "body", m.AvailabilityZones); err != nil {
		return err
	}

	return nil
}

func (m *UpdateGcpAvailabilityZonesRequest) validateEnvironment(formats strfmt.Registry) error {

	if err := validate.Required("environment", "body", m.Environment); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update gcp availability zones request based on context it is used
func (m *UpdateGcpAvailabilityZonesRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateGcpAvailabilityZonesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateGcpAvailabilityZonesRequest) UnmarshalBinary(b []byte) error {
	var res UpdateGcpAvailabilityZonesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
