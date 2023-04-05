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

// GenerateWorkloadAuthTokenResponse Response object for GenerateWorkloadAuthToken method.
//
// swagger:model GenerateWorkloadAuthTokenResponse
type GenerateWorkloadAuthTokenResponse struct {

	// The workload endpoint URL
	EndpointURL string `json:"endpointUrl,omitempty"`

	// When the information should expire
	// Format: date-time
	ExpireAt strfmt.DateTime `json:"expireAt,omitempty"`

	// The authentication token
	Token string `json:"token,omitempty"`
}

// Validate validates this generate workload auth token response
func (m *GenerateWorkloadAuthTokenResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpireAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GenerateWorkloadAuthTokenResponse) validateExpireAt(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpireAt) { // not required
		return nil
	}

	if err := validate.FormatOf("expireAt", "body", "date-time", m.ExpireAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this generate workload auth token response based on context it is used
func (m *GenerateWorkloadAuthTokenResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GenerateWorkloadAuthTokenResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenerateWorkloadAuthTokenResponse) UnmarshalBinary(b []byte) error {
	var res GenerateWorkloadAuthTokenResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
