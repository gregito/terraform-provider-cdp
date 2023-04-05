// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExistingDatabaseConfig Configurations for bringing an existing database for model metrics
//
// swagger:model ExistingDatabaseConfig
type ExistingDatabaseConfig struct {

	// Optionally provide a Postgresql database host to export model metrics to.
	ExistingDatabaseHost string `json:"existingDatabaseHost,omitempty"`

	// Optionally provide a Postgresql database name to export model metrics to.
	ExistingDatabaseName string `json:"existingDatabaseName,omitempty"`

	// Optionally provide a Postgresql database password to use when exporting model metrics.
	ExistingDatabasePassword string `json:"existingDatabasePassword,omitempty"`

	// Optionally provide a Postgresql database port to export model metrics to.
	ExistingDatabasePort string `json:"existingDatabasePort,omitempty"`

	// Optionally provide a Postgresql database user to use when exporting model metrics.
	ExistingDatabaseUser string `json:"existingDatabaseUser,omitempty"`
}

// Validate validates this existing database config
func (m *ExistingDatabaseConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this existing database config based on context it is used
func (m *ExistingDatabaseConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExistingDatabaseConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExistingDatabaseConfig) UnmarshalBinary(b []byte) error {
	var res ExistingDatabaseConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
