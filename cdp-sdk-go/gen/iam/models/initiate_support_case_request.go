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

// InitiateSupportCaseRequest Request object to initiate a support case.
//
// swagger:model InitiateSupportCaseRequest
type InitiateSupportCaseRequest struct {

	// Any additional data to be attached to this case, for example a client-side exception. In general data not to be found in server logs.
	ClientData string `json:"clientData,omitempty"`

	// Name of the component. The component name must be supported in SFDC.
	// Required: true
	Component *string `json:"component"`

	// CRN of a resource related to this support case, such as a cluster CRN. If this argument is provided it has to be a valid CRN.
	ResourceCrn string `json:"resourceCrn,omitempty"`

	// Name of the sub-component. The component subcomponent combination must be supported in SFDC.
	// Required: true
	SubComponent *string `json:"subComponent"`
}

// Validate validates this initiate support case request
func (m *InitiateSupportCaseRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComponent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubComponent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InitiateSupportCaseRequest) validateComponent(formats strfmt.Registry) error {

	if err := validate.Required("component", "body", m.Component); err != nil {
		return err
	}

	return nil
}

func (m *InitiateSupportCaseRequest) validateSubComponent(formats strfmt.Registry) error {

	if err := validate.Required("subComponent", "body", m.SubComponent); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this initiate support case request based on context it is used
func (m *InitiateSupportCaseRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InitiateSupportCaseRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InitiateSupportCaseRequest) UnmarshalBinary(b []byte) error {
	var res InitiateSupportCaseRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
