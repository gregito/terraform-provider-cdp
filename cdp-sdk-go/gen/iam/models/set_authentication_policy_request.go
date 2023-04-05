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

// SetAuthenticationPolicyRequest Request object for a set authentication policy request.
//
// swagger:model SetAuthenticationPolicyRequest
type SetAuthenticationPolicyRequest struct {

	// The expiration, in seconds, of the access key. Set the value to '0' to use system default expiration (which is 12 hours for CDP). The value will be set to '0' if not provided.
	// Minimum: 0
	AccessKeyExpirationSec *int32 `json:"accessKeyExpirationSec,omitempty"`

	// The inactivity duration, in seconds, of the access key, which would invalidate the access key due to inactivity. Set the value to '0' to use the system's default inactivity duration (which is 1 hour normally and 15 minutes for Cloudera for Government). If set to a value longer than the value for `accessKeyExpirationSec` then there will be no inactivity timeout. The value will be set to '0' if not provided.
	// Minimum: 0
	AccessKeyInactivityDurationSec *int32 `json:"accessKeyInactivityDurationSec,omitempty"`

	// The expiration, in seconds, of the UI session token. Set the value to '0' to use system default expiration (which is 12 hours for CDP). The value will be set to '0' if not provided.
	// Minimum: 0
	SessionTokenExpirationSec *int32 `json:"sessionTokenExpirationSec,omitempty"`

	// The inactivity duration, in seconds, of the UI session token, which would invalidate the session token due to inactivity. Set the value to '0' to use the system's default inactivity duration (which is 1 hour normally and 15 minutes for Cloudera for Government). If set to a value longer than the value for `sessionTokenExpirationSec` then there will be no inactivity timeout. The value will be set to '0' if not provided.
	// Minimum: 0
	SessionTokenInactivityDurationSec *int32 `json:"sessionTokenInactivityDurationSec,omitempty"`

	// The expiration, in seconds, of the workload authentication token. Set the value to '0' to use system default expiration (which is 1 hour for CDP). The value will be set to '0' if not provided.
	// Minimum: 0
	WorkloadAuthTokenExpirationSec *int32 `json:"workloadAuthTokenExpirationSec,omitempty"`
}

// Validate validates this set authentication policy request
func (m *SetAuthenticationPolicyRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessKeyExpirationSec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccessKeyInactivityDurationSec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionTokenExpirationSec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionTokenInactivityDurationSec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloadAuthTokenExpirationSec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SetAuthenticationPolicyRequest) validateAccessKeyExpirationSec(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessKeyExpirationSec) { // not required
		return nil
	}

	if err := validate.MinimumInt("accessKeyExpirationSec", "body", int64(*m.AccessKeyExpirationSec), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *SetAuthenticationPolicyRequest) validateAccessKeyInactivityDurationSec(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessKeyInactivityDurationSec) { // not required
		return nil
	}

	if err := validate.MinimumInt("accessKeyInactivityDurationSec", "body", int64(*m.AccessKeyInactivityDurationSec), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *SetAuthenticationPolicyRequest) validateSessionTokenExpirationSec(formats strfmt.Registry) error {
	if swag.IsZero(m.SessionTokenExpirationSec) { // not required
		return nil
	}

	if err := validate.MinimumInt("sessionTokenExpirationSec", "body", int64(*m.SessionTokenExpirationSec), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *SetAuthenticationPolicyRequest) validateSessionTokenInactivityDurationSec(formats strfmt.Registry) error {
	if swag.IsZero(m.SessionTokenInactivityDurationSec) { // not required
		return nil
	}

	if err := validate.MinimumInt("sessionTokenInactivityDurationSec", "body", int64(*m.SessionTokenInactivityDurationSec), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *SetAuthenticationPolicyRequest) validateWorkloadAuthTokenExpirationSec(formats strfmt.Registry) error {
	if swag.IsZero(m.WorkloadAuthTokenExpirationSec) { // not required
		return nil
	}

	if err := validate.MinimumInt("workloadAuthTokenExpirationSec", "body", int64(*m.WorkloadAuthTokenExpirationSec), 0, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this set authentication policy request based on context it is used
func (m *SetAuthenticationPolicyRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SetAuthenticationPolicyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetAuthenticationPolicyRequest) UnmarshalBinary(b []byte) error {
	var res SetAuthenticationPolicyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
