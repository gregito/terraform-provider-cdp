// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AuditLogs AuditLogs contains all the logs for a given request id.
//
// swagger:model AuditLogs
type AuditLogs struct {

	// Total number of logs associated with the request id.
	// Required: true
	Count *int32 `json:"count"`

	// Contains all the logs for a given request id.
	// Required: true
	Logs []*AuditLog `json:"logs"`

	// The request ID associated with a long-running operation to fetch the logs for.
	// Required: true
	RequestID *string `json:"requestID"`
}

// Validate validates this audit logs
func (m *AuditLogs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditLogs) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("count", "body", m.Count); err != nil {
		return err
	}

	return nil
}

func (m *AuditLogs) validateLogs(formats strfmt.Registry) error {

	if err := validate.Required("logs", "body", m.Logs); err != nil {
		return err
	}

	for i := 0; i < len(m.Logs); i++ {
		if swag.IsZero(m.Logs[i]) { // not required
			continue
		}

		if m.Logs[i] != nil {
			if err := m.Logs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("logs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AuditLogs) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("requestID", "body", m.RequestID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuditLogs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditLogs) UnmarshalBinary(b []byte) error {
	var res AuditLogs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
