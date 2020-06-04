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

// UpdateLdapProviderRequest Update LDAP Provider for account.
//
// swagger:model UpdateLdapProviderRequest
type UpdateLdapProviderRequest struct {

	// bind DN e.g. uid=myapp,ou=users,dc=example,dc=org.Optional. Required if bind is not anonymous.
	BindDn string `json:"bindDn,omitempty"`

	// The password of the bind user.
	BindPassword string `json:"bindPassword,omitempty"`

	// The LDAP email attribute name, will be used as the user's email. e.g. mail.
	EmailMappingAttribute string `json:"emailMappingAttribute,omitempty"`

	// The LDAP attribute used as firstname , will be used as the user's firstname. e.g. givenName.
	FirstNameMappingAttribute string `json:"firstNameMappingAttribute,omitempty"`

	// The property of user object to use in {{dn}} interpolation of groupSearchFilter.
	GroupDnProperty string `json:"groupDnProperty,omitempty"`

	// The LDAP group attribute name, will be used as for user's group. e.g. cn, name.
	GroupNameMappingAttribute string `json:"groupNameMappingAttribute,omitempty"`

	// The distinguished name indicating the path within the directory information tree to begin user searches from. e.g. ou=groups,dc=example,dc=com.
	// Required: true
	GroupSearchBase *string `json:"groupSearchBase"`

	// The search filter to use for finding groups for authorization of authenticated users.
	// Required: true
	GroupSearchFilter *string `json:"groupSearchFilter"`

	// The LDAP sn attribute name, will be used as the user's lastname. e.g. sn.
	LastNameMappingAttribute string `json:"lastNameMappingAttribute,omitempty"`

	// The name or CRN of LDAP provider to update.
	// Required: true
	LdapProviderName *string `json:"ldapProviderName"`

	// boolean value to control group sync. Can be omitted if no update is required.
	SkipGroupSyncOnLogin bool `json:"skipGroupSyncOnLogin,omitempty"`

	// If your ldaps:// server uses a self-signed SSL certificate or a certificate issued by a private Certificate Authority (CA), you need to provide the trusted certificates that can be used to validate the LDAP server certificate.
	TLSCaCertificates []string `json:"tlsCaCertificates"`

	// The URL of the LDAP server. The URL must be prefixed with ldap:// or ldaps://. The URL can optionally specify a custom port, for example e.g. ldaps://ldap.example.org:663.
	// Required: true
	URL *string `json:"url"`

	// Property of the LDAP user object to use when binding to verify the password. e.g. dn.
	UserBindProperty string `json:"userBindProperty,omitempty"`

	// Part of the directory tree under which to search for users. e.g. ou=users,dc=example,dc=org.
	// Required: true
	UserSearchBase *string `json:"userSearchBase"`

	// The search filter to use for finding users. e.g. (uid={{username}}).
	// Required: true
	UserSearchFilter *string `json:"userSearchFilter"`

	// The LDAP displayName attribute name, will be used as the user's name. e.g. uid, sAMAccountName.
	UsernameMappingAttribute string `json:"usernameMappingAttribute,omitempty"`
}

// Validate validates this update ldap provider request
func (m *UpdateLdapProviderRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupSearchBase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupSearchFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLdapProviderName(formats); err != nil {
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

func (m *UpdateLdapProviderRequest) validateGroupSearchBase(formats strfmt.Registry) error {

	if err := validate.Required("groupSearchBase", "body", m.GroupSearchBase); err != nil {
		return err
	}

	return nil
}

func (m *UpdateLdapProviderRequest) validateGroupSearchFilter(formats strfmt.Registry) error {

	if err := validate.Required("groupSearchFilter", "body", m.GroupSearchFilter); err != nil {
		return err
	}

	return nil
}

func (m *UpdateLdapProviderRequest) validateLdapProviderName(formats strfmt.Registry) error {

	if err := validate.Required("ldapProviderName", "body", m.LdapProviderName); err != nil {
		return err
	}

	return nil
}

func (m *UpdateLdapProviderRequest) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

func (m *UpdateLdapProviderRequest) validateUserSearchBase(formats strfmt.Registry) error {

	if err := validate.Required("userSearchBase", "body", m.UserSearchBase); err != nil {
		return err
	}

	return nil
}

func (m *UpdateLdapProviderRequest) validateUserSearchFilter(formats strfmt.Registry) error {

	if err := validate.Required("userSearchFilter", "body", m.UserSearchFilter); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateLdapProviderRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateLdapProviderRequest) UnmarshalBinary(b []byte) error {
	var res UpdateLdapProviderRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
