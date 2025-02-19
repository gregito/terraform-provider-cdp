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

// InstanceGroup The type of the instance group which also contains the actual instance(s)
//
// swagger:model InstanceGroup
type InstanceGroup struct {

	// List of availability zones that this instance group is associated with.
	AvailabilityZones []string `json:"availabilityZones"`

	// List of instances in this instance group.
	// Required: true
	Instances []*Instance `json:"instances"`

	// The name of the instance group where the given instance is located.
	// Required: true
	Name *string `json:"name"`

	// The list of subnet IDs in case of multi-availability zone setup
	SubnetIds []string `json:"subnetIds"`
}

// Validate validates this instance group
func (m *InstanceGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstances(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceGroup) validateInstances(formats strfmt.Registry) error {

	if err := validate.Required("instances", "body", m.Instances); err != nil {
		return err
	}

	for i := 0; i < len(m.Instances); i++ {
		if swag.IsZero(m.Instances[i]) { // not required
			continue
		}

		if m.Instances[i] != nil {
			if err := m.Instances[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instances" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("instances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InstanceGroup) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this instance group based on the context it is used
func (m *InstanceGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInstances(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceGroup) contextValidateInstances(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Instances); i++ {

		if m.Instances[i] != nil {

			if swag.IsZero(m.Instances[i]) { // not required
				return nil
			}

			if err := m.Instances[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instances" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("instances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstanceGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceGroup) UnmarshalBinary(b []byte) error {
	var res InstanceGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
