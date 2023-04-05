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

// GetLogsRequest GetLogsRequestfor getting logs for a request ID.
//
// swagger:model GetLogsRequest
type GetLogsRequest struct {

	// Flag to fetch all logs.
	FetchAll bool `json:"fetchAll,omitempty"`

	// offset from which the logs should be fetched.
	Offset int32 `json:"offset,omitempty"`

	// Limit the number of logs.
	PageSize int32 `json:"pageSize,omitempty"`

	// Unique Key to identify a set of logs.
	// Required: true
	RequestID *string `json:"requestID"`

	// WorkspaceCrn the requestID belongs to.
	// Required: true
	WorkspaceCrn *string `json:"workspaceCrn"`
}

// Validate validates this get logs request
func (m *GetLogsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequestID(formats); err != nil {
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

func (m *GetLogsRequest) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("requestID", "body", m.RequestID); err != nil {
		return err
	}

	return nil
}

func (m *GetLogsRequest) validateWorkspaceCrn(formats strfmt.Registry) error {

	if err := validate.Required("workspaceCrn", "body", m.WorkspaceCrn); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get logs request based on context it is used
func (m *GetLogsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetLogsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetLogsRequest) UnmarshalBinary(b []byte) error {
	var res GetLogsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
