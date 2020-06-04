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

// AzureCredentialPrerequisitesResponse Response object for getting Azure credential prerequisites.
//
// swagger:model AzureCredentialPrerequisitesResponse
type AzureCredentialPrerequisitesResponse struct {

	// Azure CLI command to create Azure AD Application.
	// Required: true
	AppCreationCommand *string `json:"appCreationCommand"`

	// The related role definition json encoded in base64
	// Required: true
	RoleDefitionJSON *string `json:"roleDefitionJson"`
}

// Validate validates this azure credential prerequisites response
func (m *AzureCredentialPrerequisitesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppCreationCommand(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleDefitionJSON(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureCredentialPrerequisitesResponse) validateAppCreationCommand(formats strfmt.Registry) error {

	if err := validate.Required("appCreationCommand", "body", m.AppCreationCommand); err != nil {
		return err
	}

	return nil
}

func (m *AzureCredentialPrerequisitesResponse) validateRoleDefitionJSON(formats strfmt.Registry) error {

	if err := validate.Required("roleDefitionJson", "body", m.RoleDefitionJSON); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AzureCredentialPrerequisitesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureCredentialPrerequisitesResponse) UnmarshalBinary(b []byte) error {
	var res AzureCredentialPrerequisitesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
