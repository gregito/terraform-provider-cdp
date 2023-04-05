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

// BackupWorkspaceRequest The request object for workspace backup.
//
// swagger:model BackupWorkspaceRequest
type BackupWorkspaceRequest struct {

	// The timeout(in minutes) to use for the execution of the backup jobs.
	BackupJobTimeoutMinutes int32 `json:"backupJobTimeoutMinutes,omitempty"`

	// Backup name.
	// Required: true
	BackupName *string `json:"backupName"`

	// Skip pre-flight validations if requested.
	SkipValidation bool `json:"skipValidation,omitempty"`

	// CRN of the workspace to backup.
	// Required: true
	WorkspaceCrn *string `json:"workspaceCrn"`

	// The experimental entitlements to pass to CML for testing purpose. This is not meant for external customers in any case
	XEntitlements []string `json:"xEntitlements"`
}

// Validate validates this backup workspace request
func (m *BackupWorkspaceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkspaceCrn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupWorkspaceRequest) validateBackupName(formats strfmt.Registry) error {

	if err := validate.Required("backupName", "body", m.BackupName); err != nil {
		return err
	}

	return nil
}

func (m *BackupWorkspaceRequest) validateWorkspaceCrn(formats strfmt.Registry) error {

	if err := validate.Required("workspaceCrn", "body", m.WorkspaceCrn); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this backup workspace request based on context it is used
func (m *BackupWorkspaceRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BackupWorkspaceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupWorkspaceRequest) UnmarshalBinary(b []byte) error {
	var res BackupWorkspaceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
