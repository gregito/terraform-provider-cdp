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

// ListLdapProvidersResponse Response object for listing LDAP provider.
//
// swagger:model ListLdapProvidersResponse
type ListLdapProvidersResponse struct {

	// List of LDAP providers.
	// Required: true
	LdapProviders []*LdapProvider `json:"ldapProviders"`

	// The token to use when requesting the next set of results. If not present, there are no additional results.
	NextToken string `json:"nextToken,omitempty"`
}

// Validate validates this list ldap providers response
func (m *ListLdapProvidersResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLdapProviders(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListLdapProvidersResponse) validateLdapProviders(formats strfmt.Registry) error {

	if err := validate.Required("ldapProviders", "body", m.LdapProviders); err != nil {
		return err
	}

	for i := 0; i < len(m.LdapProviders); i++ {
		if swag.IsZero(m.LdapProviders[i]) { // not required
			continue
		}

		if m.LdapProviders[i] != nil {
			if err := m.LdapProviders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ldapProviders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListLdapProvidersResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListLdapProvidersResponse) UnmarshalBinary(b []byte) error {
	var res ListLdapProvidersResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
