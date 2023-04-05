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

// DescribeConfigDiffRequest Request object for the describeConfigDiff method.
//
// swagger:model DescribeConfigDiffRequest
type DescribeConfigDiffRequest struct {

	// ID of the cluster.
	// Required: true
	ClusterID *string `json:"clusterId"`

	// ID of the old service configuration.
	// Required: true
	FromConfigID *string `json:"fromConfigId"`

	// ID of the new service configuration.
	// Required: true
	ToConfigID *string `json:"toConfigId"`
}

// Validate validates this describe config diff request
func (m *DescribeConfigDiffRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFromConfigID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToConfigID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DescribeConfigDiffRequest) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("clusterId", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *DescribeConfigDiffRequest) validateFromConfigID(formats strfmt.Registry) error {

	if err := validate.Required("fromConfigId", "body", m.FromConfigID); err != nil {
		return err
	}

	return nil
}

func (m *DescribeConfigDiffRequest) validateToConfigID(formats strfmt.Registry) error {

	if err := validate.Required("toConfigId", "body", m.ToConfigID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this describe config diff request based on context it is used
func (m *DescribeConfigDiffRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DescribeConfigDiffRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DescribeConfigDiffRequest) UnmarshalBinary(b []byte) error {
	var res DescribeConfigDiffRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
