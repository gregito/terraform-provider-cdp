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

// GetCliForDatalakeRequest Request object for generating a create datalake CLI command.
//
// swagger:model GetCliForDatalakeRequest
type GetCliForDatalakeRequest struct {

	// cloud provider for the command
	// Required: true
	CommandCloudProvider *CloudProviderType `json:"commandCloudProvider"`

	// cloudbreak datalake request object of the private API, encoded in base64
	CommandRequestInput string `json:"commandRequestInput,omitempty"`

	// cloudbreak datalake response object of the private API, provided by describing an existing datalake, encoded in base64
	CommandResponseInput string `json:"commandResponseInput,omitempty"`
}

// Validate validates this get cli for datalake request
func (m *GetCliForDatalakeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommandCloudProvider(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCliForDatalakeRequest) validateCommandCloudProvider(formats strfmt.Registry) error {

	if err := validate.Required("commandCloudProvider", "body", m.CommandCloudProvider); err != nil {
		return err
	}

	if err := validate.Required("commandCloudProvider", "body", m.CommandCloudProvider); err != nil {
		return err
	}

	if m.CommandCloudProvider != nil {
		if err := m.CommandCloudProvider.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commandCloudProvider")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commandCloudProvider")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get cli for datalake request based on the context it is used
func (m *GetCliForDatalakeRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCommandCloudProvider(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCliForDatalakeRequest) contextValidateCommandCloudProvider(ctx context.Context, formats strfmt.Registry) error {

	if m.CommandCloudProvider != nil {
		if err := m.CommandCloudProvider.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commandCloudProvider")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commandCloudProvider")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCliForDatalakeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCliForDatalakeRequest) UnmarshalBinary(b []byte) error {
	var res GetCliForDatalakeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
