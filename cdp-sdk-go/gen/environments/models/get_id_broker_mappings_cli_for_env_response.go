// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetIDBrokerMappingsCliForEnvResponse Response object which contains the IDBroker mapping creation CLI command
//
// swagger:model GetIdBrokerMappingsCliForEnvResponse
type GetIDBrokerMappingsCliForEnvResponse struct {

	// cdp cli command string
	Command string `json:"command,omitempty"`
}

// Validate validates this get Id broker mappings cli for env response
func (m *GetIDBrokerMappingsCliForEnvResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get Id broker mappings cli for env response based on context it is used
func (m *GetIDBrokerMappingsCliForEnvResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetIDBrokerMappingsCliForEnvResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetIDBrokerMappingsCliForEnvResponse) UnmarshalBinary(b []byte) error {
	var res GetIDBrokerMappingsCliForEnvResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
