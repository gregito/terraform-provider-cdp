// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DescribeLdapProviderResponse response object for DescribeLdapProvider request.
//
// swagger:model DescribeLdapProviderResponse
type DescribeLdapProviderResponse struct {

	// The LDAP identity provider connector configuration details.
	// Required: true
	LdapProvider *LdapProvider `json:"ldapProvider"`
}

// Validate validates this describe ldap provider response
func (m *DescribeLdapProviderResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLdapProvider(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DescribeLdapProviderResponse) validateLdapProvider(formats strfmt.Registry) error {

	if err := validate.Required("ldapProvider", "body", m.LdapProvider); err != nil {
		return err
	}

	if m.LdapProvider != nil {
		if err := m.LdapProvider.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ldapProvider")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DescribeLdapProviderResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DescribeLdapProviderResponse) UnmarshalBinary(b []byte) error {
	var res DescribeLdapProviderResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
