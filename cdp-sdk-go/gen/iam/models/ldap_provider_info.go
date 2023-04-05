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

// LdapProviderInfo Information about an LDAP provider connector.
//
// swagger:model LdapProviderInfo
type LdapProviderInfo struct {

	// bindDN
	BindDn string `json:"bindDn,omitempty"`

	// The LDAP email attribute name, will be used as the user's email.
	EmailMappingAttribute string `json:"emailMappingAttribute,omitempty"`

	// The LDAP attribute used as firstname , will be used as the user's firstname.
	FirstNameMappingAttribute string `json:"firstNameMappingAttribute,omitempty"`

	// The property of user object to use in {{dn}} interpolation of groupSearchFilter.
	GroupDnProperty string `json:"groupDnProperty,omitempty"`

	// The LDAP group attribute name, will be used as for user's group.
	GroupNameMappingAttribute string `json:"groupNameMappingAttribute,omitempty"`

	// Part of the directory tree under which group searches should be performed.
	// Required: true
	GroupSearchBase *string `json:"groupSearchBase"`

	// Filter which is used to search for group membership.
	// Required: true
	GroupSearchFilter *string `json:"groupSearchFilter"`

	// The LDAP sn attribute name, will be used as the user's lastname.
	LastNameMappingAttribute string `json:"lastNameMappingAttribute,omitempty"`

	// boolean value to indicate whether a start TLS request should be initiated on connect.
	StartTLS bool `json:"startTls,omitempty"`

	// If your ldaps:// server uses a self-signed SSL certificate or a certificate issued by a private Certificate Authority (CA), you need to provide the trusted certificates that can be used to validate the LDAP server certificate.
	TLSCaCertificates []string `json:"tlsCaCertificates"`

	// The URL of the LDAP server.
	// Required: true
	URL *string `json:"url"`

	// Property of the LDAP user object to use when binding to the LDAP directory.
	UserBindProperty string `json:"userBindProperty,omitempty"`

	// Part of the directory tree under which to search for users.
	// Required: true
	UserSearchBase *string `json:"userSearchBase"`

	// The search filter to use for finding users.
	// Required: true
	UserSearchFilter *string `json:"userSearchFilter"`

	// The LDAP displayName attribute name, will be used as the user's name.
	UsernameMappingAttribute string `json:"usernameMappingAttribute,omitempty"`
}

// Validate validates this ldap provider info
func (m *LdapProviderInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupSearchBase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupSearchFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserSearchBase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserSearchFilter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapProviderInfo) validateGroupSearchBase(formats strfmt.Registry) error {

	if err := validate.Required("groupSearchBase", "body", m.GroupSearchBase); err != nil {
		return err
	}

	return nil
}

func (m *LdapProviderInfo) validateGroupSearchFilter(formats strfmt.Registry) error {

	if err := validate.Required("groupSearchFilter", "body", m.GroupSearchFilter); err != nil {
		return err
	}

	return nil
}

func (m *LdapProviderInfo) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

func (m *LdapProviderInfo) validateUserSearchBase(formats strfmt.Registry) error {

	if err := validate.Required("userSearchBase", "body", m.UserSearchBase); err != nil {
		return err
	}

	return nil
}

func (m *LdapProviderInfo) validateUserSearchFilter(formats strfmt.Registry) error {

	if err := validate.Required("userSearchFilter", "body", m.UserSearchFilter); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this ldap provider info based on context it is used
func (m *LdapProviderInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LdapProviderInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapProviderInfo) UnmarshalBinary(b []byte) error {
	var res LdapProviderInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
