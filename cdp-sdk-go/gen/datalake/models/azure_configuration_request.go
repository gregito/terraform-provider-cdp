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

// AzureConfigurationRequest Request object for Azure configuration.
//
// swagger:model AzureConfigurationRequest
type AzureConfigurationRequest struct {

	// The managed identity to use. The assumer should have Virtual Machine Contributor and Managed Identity Operator roles on subscription level.
	// Required: true
	ManagedIdentity *string `json:"managedIdentity"`

	// The storage location to use. The location has to be in the following format abfs://filesystem@storage-account-name.dfs.core.windows.net. The filesystem must already exist and the storage account must be StorageV2.
	// Required: true
	StorageLocation *string `json:"storageLocation"`
}

// Validate validates this azure configuration request
func (m *AzureConfigurationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateManagedIdentity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureConfigurationRequest) validateManagedIdentity(formats strfmt.Registry) error {

	if err := validate.Required("managedIdentity", "body", m.ManagedIdentity); err != nil {
		return err
	}

	return nil
}

func (m *AzureConfigurationRequest) validateStorageLocation(formats strfmt.Registry) error {

	if err := validate.Required("storageLocation", "body", m.StorageLocation); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AzureConfigurationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureConfigurationRequest) UnmarshalBinary(b []byte) error {
	var res AzureConfigurationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
