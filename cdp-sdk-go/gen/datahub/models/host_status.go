// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HostStatus Information about cluster host status.
//
// swagger:model HostStatus
type HostStatus struct {

	// The host health summary.
	HealthSummary string `json:"healthSummary,omitempty"`

	// Unique identifier of the cluster host given by Cloudera Manager.
	Hostid string `json:"hostid,omitempty"`

	// The cluster hostname.
	Hostname string `json:"hostname,omitempty"`
}

// Validate validates this host status
func (m *HostStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HostStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostStatus) UnmarshalBinary(b []byte) error {
	var res HostStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
