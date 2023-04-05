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

// DeleteBackupRequest The request object for DeleteBackup operation.
//
// swagger:model DeleteBackupRequest
type DeleteBackupRequest struct {

	// The CRN of the backup to be deleted.
	// Required: true
	BackupCrn *string `json:"backupCrn"`

	// Force deleting a workspace backup even if datalake does not exist. This option does not guarantee that the backed-up resources in the cloud account would be cleaned up.
	Force bool `json:"force,omitempty"`

	// Skip pre-flight validations if requested.
	SkipValidation bool `json:"skipValidation,omitempty"`

	// The experimental entitlements to pass to CML for testing purpose. This is not meant for external customers in any case
	XEntitlements []string `json:"xEntitlements"`
}

// Validate validates this delete backup request
func (m *DeleteBackupRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupCrn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteBackupRequest) validateBackupCrn(formats strfmt.Registry) error {

	if err := validate.Required("backupCrn", "body", m.BackupCrn); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this delete backup request based on context it is used
func (m *DeleteBackupRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteBackupRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteBackupRequest) UnmarshalBinary(b []byte) error {
	var res DeleteBackupRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
