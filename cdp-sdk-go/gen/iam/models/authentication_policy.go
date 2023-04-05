// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AuthenticationPolicy Information about the authentication policy for an account.
//
// swagger:model AuthenticationPolicy
type AuthenticationPolicy struct {

	// The expiration, in seconds, of the access key. The value of '0' indicates the system default expiration (which is 12 hours).
	AccessKeyExpirationSec int32 `json:"accessKeyExpirationSec,omitempty"`

	// The inactivity duration, in seconds, of the access key, which would invalidate the access key due to no activity. The value of '0' indicates default inactivity duration (which is 1 hour normally and 15 minutes for Cloudera for Government). There's no access key invalidation from no activity if the value is greater or equal to expiration.
	AccessKeyInactivityDurationSec int32 `json:"accessKeyInactivityDurationSec,omitempty"`

	// The expiration, in seconds, of the UI session token. The value of '0' indicates the system default expiration (which is 12 hours).
	SessionTokenExpirationSec int32 `json:"sessionTokenExpirationSec,omitempty"`

	// The inactivity duration, in seconds, of the UI session token, which would invalidate the session token due to no activity. The value of '0' indicates default inactivity duration (which is 1 hour normally and 15 minutes for Cloudera for Government). There's no session token invalidation from no activity if the value is greater or equal to expiration.
	SessionTokenInactivityDurationSec int32 `json:"sessionTokenInactivityDurationSec,omitempty"`

	// The expiration, in seconds, of the workload authentication token. The value of '0' indicates the system default expiration (which is 1 hour).
	WorkloadAuthTokenExpirationSec int32 `json:"workloadAuthTokenExpirationSec,omitempty"`
}

// Validate validates this authentication policy
func (m *AuthenticationPolicy) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this authentication policy based on context it is used
func (m *AuthenticationPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuthenticationPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthenticationPolicy) UnmarshalBinary(b []byte) error {
	var res AuthenticationPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
