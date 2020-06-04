// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AutoScalePolicyRequest An individual AutoScale Policy
//
// swagger:model AutoScalePolicyRequest
type AutoScalePolicyRequest struct {

	// A single hostgroup, or a comma separated list of hostGroups to which the rules apply
	// Required: true
	// Max Length: 200
	// Min Length: 1
	HostGroups *string `json:"hostGroups"`

	// Load based policy
	LoadBasdePolicy *AutoScaleLoadRequest `json:"loadBasdePolicy,omitempty"`

	// Scheduled based policy
	ScheduleBasedPolicy *AutoScaleScheduleRequest `json:"scheduleBasedPolicy,omitempty"`
}

// Validate validates this auto scale policy request
func (m *AutoScalePolicyRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoadBasdePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduleBasedPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoScalePolicyRequest) validateHostGroups(formats strfmt.Registry) error {

	if err := validate.Required("hostGroups", "body", m.HostGroups); err != nil {
		return err
	}

	if err := validate.MinLength("hostGroups", "body", string(*m.HostGroups), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("hostGroups", "body", string(*m.HostGroups), 200); err != nil {
		return err
	}

	return nil
}

func (m *AutoScalePolicyRequest) validateLoadBasdePolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.LoadBasdePolicy) { // not required
		return nil
	}

	if m.LoadBasdePolicy != nil {
		if err := m.LoadBasdePolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("loadBasdePolicy")
			}
			return err
		}
	}

	return nil
}

func (m *AutoScalePolicyRequest) validateScheduleBasedPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.ScheduleBasedPolicy) { // not required
		return nil
	}

	if m.ScheduleBasedPolicy != nil {
		if err := m.ScheduleBasedPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scheduleBasedPolicy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AutoScalePolicyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutoScalePolicyRequest) UnmarshalBinary(b []byte) error {
	var res AutoScalePolicyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
