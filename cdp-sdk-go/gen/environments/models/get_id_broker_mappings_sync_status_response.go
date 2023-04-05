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

// GetIDBrokerMappingsSyncStatusResponse Response object for getting ID Broker mappings sync status.
//
// swagger:model GetIdBrokerMappingsSyncStatusResponse
type GetIDBrokerMappingsSyncStatusResponse struct {

	// The overall mappings sync status for all datalake clusters in the environment.
	// Required: true
	GlobalStatus *SyncStatus `json:"globalStatus"`

	// Map of datalake cluster CRN to mappings sync status for each datalake cluster in the environment.
	// Required: true
	Statuses map[string]IDBrokerSyncStatus `json:"statuses"`

	// Whether a sync is needed to bring in-cluster mappings up-to-date.
	// Required: true
	SyncNeeded *bool `json:"syncNeeded"`
}

// Validate validates this get Id broker mappings sync status response
func (m *GetIDBrokerMappingsSyncStatusResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGlobalStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatuses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSyncNeeded(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetIDBrokerMappingsSyncStatusResponse) validateGlobalStatus(formats strfmt.Registry) error {

	if err := validate.Required("globalStatus", "body", m.GlobalStatus); err != nil {
		return err
	}

	if err := validate.Required("globalStatus", "body", m.GlobalStatus); err != nil {
		return err
	}

	if m.GlobalStatus != nil {
		if err := m.GlobalStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("globalStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("globalStatus")
			}
			return err
		}
	}

	return nil
}

func (m *GetIDBrokerMappingsSyncStatusResponse) validateStatuses(formats strfmt.Registry) error {

	if err := validate.Required("statuses", "body", m.Statuses); err != nil {
		return err
	}

	for k := range m.Statuses {

		if err := validate.Required("statuses"+"."+k, "body", m.Statuses[k]); err != nil {
			return err
		}
		if val, ok := m.Statuses[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("statuses" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("statuses" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetIDBrokerMappingsSyncStatusResponse) validateSyncNeeded(formats strfmt.Registry) error {

	if err := validate.Required("syncNeeded", "body", m.SyncNeeded); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get Id broker mappings sync status response based on the context it is used
func (m *GetIDBrokerMappingsSyncStatusResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGlobalStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatuses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetIDBrokerMappingsSyncStatusResponse) contextValidateGlobalStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.GlobalStatus != nil {
		if err := m.GlobalStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("globalStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("globalStatus")
			}
			return err
		}
	}

	return nil
}

func (m *GetIDBrokerMappingsSyncStatusResponse) contextValidateStatuses(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.Required("statuses", "body", m.Statuses); err != nil {
		return err
	}

	for k := range m.Statuses {

		if val, ok := m.Statuses[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetIDBrokerMappingsSyncStatusResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetIDBrokerMappingsSyncStatusResponse) UnmarshalBinary(b []byte) error {
	var res GetIDBrokerMappingsSyncStatusResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
