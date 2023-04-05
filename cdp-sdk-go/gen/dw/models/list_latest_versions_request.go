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

// ListLatestVersionsRequest Request object of the listLatestVersions call.
//
// swagger:model ListLatestVersionsRequest
type ListLatestVersionsRequest struct {

	// The ID of the cluster from which the version information must be collected.
	// Required: true
	ClusterID *string `json:"clusterId"`
}

// Validate validates this list latest versions request
func (m *ListLatestVersionsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListLatestVersionsRequest) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("clusterId", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this list latest versions request based on context it is used
func (m *ListLatestVersionsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListLatestVersionsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListLatestVersionsRequest) UnmarshalBinary(b []byte) error {
	var res ListLatestVersionsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
