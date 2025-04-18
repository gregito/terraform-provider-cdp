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

// QuotaConfig Quota configuration for quota management.
//
// swagger:model QuotaConfig
type QuotaConfig struct {

	// The quota for CPU resource.
	// Required: true
	CPUQuota *string `json:"cpuQuota"`

	// The quota for gpu resource.
	GpuQuota string `json:"gpuQuota,omitempty"`

	// The quota for memory resource.
	// Required: true
	MemoryQuota *string `json:"memoryQuota"`
}

// Validate validates this quota config
func (m *QuotaConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPUQuota(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemoryQuota(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuotaConfig) validateCPUQuota(formats strfmt.Registry) error {

	if err := validate.Required("cpuQuota", "body", m.CPUQuota); err != nil {
		return err
	}

	return nil
}

func (m *QuotaConfig) validateMemoryQuota(formats strfmt.Registry) error {

	if err := validate.Required("memoryQuota", "body", m.MemoryQuota); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this quota config based on context it is used
func (m *QuotaConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QuotaConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuotaConfig) UnmarshalBinary(b []byte) error {
	var res QuotaConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
