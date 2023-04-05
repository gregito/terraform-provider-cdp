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

// InitiateSupportCaseResponse Response object for a new support case request.
//
// swagger:model InitiateSupportCaseResponse
type InitiateSupportCaseResponse struct {

	// To complete the support request user must submit the form using a web browser.
	// Required: true
	CaseFormURL *string `json:"caseFormUrl"`
}

// Validate validates this initiate support case response
func (m *InitiateSupportCaseResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCaseFormURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InitiateSupportCaseResponse) validateCaseFormURL(formats strfmt.Registry) error {

	if err := validate.Required("caseFormUrl", "body", m.CaseFormURL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this initiate support case response based on context it is used
func (m *InitiateSupportCaseResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InitiateSupportCaseResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InitiateSupportCaseResponse) UnmarshalBinary(b []byte) error {
	var res InitiateSupportCaseResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
