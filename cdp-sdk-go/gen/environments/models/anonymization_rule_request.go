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

// AnonymizationRuleRequest Anonymization rule request object that is applied on logs that are sent to Cloudera.
//
// swagger:model AnonymizationRuleRequest
type AnonymizationRuleRequest struct {

	// If rule pattern (value) matches, that will be replaced for this string (default [REDACTED])
	Replacement *string `json:"replacement,omitempty"`

	// Pattern of the rule that should be redacted.
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this anonymization rule request
func (m *AnonymizationRuleRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnonymizationRuleRequest) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this anonymization rule request based on context it is used
func (m *AnonymizationRuleRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AnonymizationRuleRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnonymizationRuleRequest) UnmarshalBinary(b []byte) error {
	var res AnonymizationRuleRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
