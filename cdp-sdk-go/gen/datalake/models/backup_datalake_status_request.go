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

// BackupDatalakeStatusRequest Request object to get the status of datalake backup. Returns the status of the latest backup that matches the provided input.
//
// swagger:model BackupDatalakeStatusRequest
type BackupDatalakeStatusRequest struct {

	// Unique identifier of the backup performed. When provided, the status request will get the status entry that has the backupid provided.
	BackupID string `json:"backupId,omitempty"`

	// The name of the backup. When provided, the status request will get the status of the latest backup performed with the given backup name on the given datalake.
	BackupName string `json:"backupName,omitempty"`

	// The name of the datalake. When backupName and backupId are not provided, status request will get the status of the latest backup operation performed on the given datalake.
	// Required: true
	DatalakeName *string `json:"datalakeName"`
}

// Validate validates this backup datalake status request
func (m *BackupDatalakeStatusRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatalakeName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupDatalakeStatusRequest) validateDatalakeName(formats strfmt.Registry) error {

	if err := validate.Required("datalakeName", "body", m.DatalakeName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this backup datalake status request based on context it is used
func (m *BackupDatalakeStatusRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BackupDatalakeStatusRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupDatalakeStatusRequest) UnmarshalBinary(b []byte) error {
	var res BackupDatalakeStatusRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
