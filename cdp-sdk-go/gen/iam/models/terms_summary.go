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

// TermsSummary Information about a set of terms.
//
// swagger:model TermsSummary
type TermsSummary struct {

	// The date of the last time the terms have been acceptance. May be empty if the terms have not been accepted.
	// Format: date-time
	AcceptanceDate strfmt.DateTime `json:"acceptanceDate,omitempty"`

	// The acceptance state.
	// Required: true
	AcceptanceState *TermsAcceptanceState `json:"acceptanceState"`

	// The CRN of the last user who accepted the terms. May be empty if the terms have not been accepted.
	Acceptor string `json:"acceptor,omitempty"`

	// The terms acceptance expiry date. Value is not set if the terms acceptance never expire.
	// Format: date-time
	ExpiryDate strfmt.DateTime `json:"expiryDate,omitempty"`

	// The name of the terms.
	// Required: true
	TermsName *string `json:"termsName"`
}

// Validate validates this terms summary
func (m *TermsSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcceptanceDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAcceptanceState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiryDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTermsName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TermsSummary) validateAcceptanceDate(formats strfmt.Registry) error {
	if swag.IsZero(m.AcceptanceDate) { // not required
		return nil
	}

	if err := validate.FormatOf("acceptanceDate", "body", "date-time", m.AcceptanceDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TermsSummary) validateAcceptanceState(formats strfmt.Registry) error {

	if err := validate.Required("acceptanceState", "body", m.AcceptanceState); err != nil {
		return err
	}

	if err := validate.Required("acceptanceState", "body", m.AcceptanceState); err != nil {
		return err
	}

	if m.AcceptanceState != nil {
		if err := m.AcceptanceState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("acceptanceState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("acceptanceState")
			}
			return err
		}
	}

	return nil
}

func (m *TermsSummary) validateExpiryDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpiryDate) { // not required
		return nil
	}

	if err := validate.FormatOf("expiryDate", "body", "date-time", m.ExpiryDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TermsSummary) validateTermsName(formats strfmt.Registry) error {

	if err := validate.Required("termsName", "body", m.TermsName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this terms summary based on the context it is used
func (m *TermsSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAcceptanceState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TermsSummary) contextValidateAcceptanceState(ctx context.Context, formats strfmt.Registry) error {

	if m.AcceptanceState != nil {
		if err := m.AcceptanceState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("acceptanceState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("acceptanceState")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TermsSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TermsSummary) UnmarshalBinary(b []byte) error {
	var res TermsSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
