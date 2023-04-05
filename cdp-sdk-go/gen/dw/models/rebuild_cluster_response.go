// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RebuildClusterResponse Response object for the rebuildCluster method.
//
// swagger:model RebuildClusterResponse
type RebuildClusterResponse struct {

	// The ID of the cluster to rebuild.
	ClusterID string `json:"clusterId,omitempty"`
}

// Validate validates this rebuild cluster response
func (m *RebuildClusterResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this rebuild cluster response based on context it is used
func (m *RebuildClusterResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RebuildClusterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RebuildClusterResponse) UnmarshalBinary(b []byte) error {
	var res RebuildClusterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
