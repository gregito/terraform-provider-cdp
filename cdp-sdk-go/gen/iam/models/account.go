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

// Account Information about a Cloudera CDP account.
//
// swagger:model Account
type Account struct {

	// The authentication policy object. May be omitted if no such policy was defined.
	AuthenticationPolicy *AuthenticationPolicy `json:"authenticationPolicy,omitempty"`

	// Whether login is enabled for Cloudera SSO users. It can only be set by Cloudera upon request and disables interactive login through Cloudera SSO. Note that restricting Cloudera SSO login will prevent account administrators from logging in interactively. Its default value is 'true'. When it is 'true', the Cloudera SSO interactive login behavior is controlled according to the existing `clouderaSSOLoginEnabled` flag. When it is 'false', it overrides the setting for `clouderaSSOLoginEnabled`.
	// Required: true
	ClouderaSSOAllLoginEnabled *bool `json:"clouderaSSOAllLoginEnabled"`

	// Whether interactive login using Cloudera SSO is enabled for users who are not account administrators. Its default value is 'true'. When it is 'true', the account administrators, as well as non-administrator users can login through Cloudera SSO. When it is 'false', Cloudera SSO users who are not account administrators will not be able to login.
	// Required: true
	ClouderaSSOLoginEnabled *bool `json:"clouderaSSOLoginEnabled"`

	// The machine user workload password policy object. May be omitted if no such policy was defined.
	MachineUserWorkloadPasswordPolicy *WorkloadPasswordPolicy `json:"machineUserWorkloadPasswordPolicy,omitempty"`

	// The workload password policy object.
	// Required: true
	WorkloadPasswordPolicy *WorkloadPasswordPolicy `json:"workloadPasswordPolicy"`
}

// Validate validates this account
func (m *Account) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthenticationPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClouderaSSOAllLoginEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClouderaSSOLoginEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachineUserWorkloadPasswordPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloadPasswordPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Account) validateAuthenticationPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthenticationPolicy) { // not required
		return nil
	}

	if m.AuthenticationPolicy != nil {
		if err := m.AuthenticationPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authenticationPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authenticationPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Account) validateClouderaSSOAllLoginEnabled(formats strfmt.Registry) error {

	if err := validate.Required("clouderaSSOAllLoginEnabled", "body", m.ClouderaSSOAllLoginEnabled); err != nil {
		return err
	}

	return nil
}

func (m *Account) validateClouderaSSOLoginEnabled(formats strfmt.Registry) error {

	if err := validate.Required("clouderaSSOLoginEnabled", "body", m.ClouderaSSOLoginEnabled); err != nil {
		return err
	}

	return nil
}

func (m *Account) validateMachineUserWorkloadPasswordPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.MachineUserWorkloadPasswordPolicy) { // not required
		return nil
	}

	if m.MachineUserWorkloadPasswordPolicy != nil {
		if err := m.MachineUserWorkloadPasswordPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("machineUserWorkloadPasswordPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("machineUserWorkloadPasswordPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Account) validateWorkloadPasswordPolicy(formats strfmt.Registry) error {

	if err := validate.Required("workloadPasswordPolicy", "body", m.WorkloadPasswordPolicy); err != nil {
		return err
	}

	if m.WorkloadPasswordPolicy != nil {
		if err := m.WorkloadPasswordPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workloadPasswordPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workloadPasswordPolicy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this account based on the context it is used
func (m *Account) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthenticationPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMachineUserWorkloadPasswordPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWorkloadPasswordPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Account) contextValidateAuthenticationPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.AuthenticationPolicy != nil {

		if swag.IsZero(m.AuthenticationPolicy) { // not required
			return nil
		}

		if err := m.AuthenticationPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authenticationPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authenticationPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Account) contextValidateMachineUserWorkloadPasswordPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.MachineUserWorkloadPasswordPolicy != nil {

		if swag.IsZero(m.MachineUserWorkloadPasswordPolicy) { // not required
			return nil
		}

		if err := m.MachineUserWorkloadPasswordPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("machineUserWorkloadPasswordPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("machineUserWorkloadPasswordPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Account) contextValidateWorkloadPasswordPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.WorkloadPasswordPolicy != nil {

		if err := m.WorkloadPasswordPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workloadPasswordPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workloadPasswordPolicy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Account) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Account) UnmarshalBinary(b []byte) error {
	var res Account
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
