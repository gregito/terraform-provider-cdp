// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VolumeEncryptionRequest Configurations for volume encryption.
//
// swagger:model VolumeEncryptionRequest
type VolumeEncryptionRequest struct {

	// Enable encyrption for all volumes in the instance group. Default is false.
	EnableEncryption *bool `json:"enableEncryption,omitempty"`

	// The ARN of the encryption key to use. If nothing is specified, the default key will be used.
	EncryptionKey string `json:"encryptionKey,omitempty"`
}

// Validate validates this volume encryption request
func (m *VolumeEncryptionRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VolumeEncryptionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeEncryptionRequest) UnmarshalBinary(b []byte) error {
	var res VolumeEncryptionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
