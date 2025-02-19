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

// DbcSummary A Database Catalog.
//
// swagger:model DbcSummary
type DbcSummary struct {

	// Creation date of Database Catalog.
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creationDate,omitempty"`

	// The creator of the Database Catalog.
	Creator *ActorResponse `json:"creator,omitempty"`

	// The CRN of the Database Catalog.
	Crn string `json:"crn,omitempty"`

	// The ID of the Database Catalog.
	ID string `json:"id,omitempty"`

	// The name of the Database Catalog.
	Name string `json:"name,omitempty"`

	// Status of the Database Catalog. Possible values are: Creating, Created, Accepted, Starting, Running, Stopping, Stopped, Updating, PreUpdate, Upgrading, PreUpgrade, Restarting, Deleting, Waiting, Failed, Error.
	Status string `json:"status,omitempty"`

	// Timestamp of the last status change of the Database Catalog.
	// Format: date-time
	StatusChangedAt strfmt.DateTime `json:"statusChangedAt,omitempty"`
}

// Validate validates this dbc summary
func (m *DbcSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusChangedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DbcSummary) validateCreationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DbcSummary) validateCreator(formats strfmt.Registry) error {
	if swag.IsZero(m.Creator) { // not required
		return nil
	}

	if m.Creator != nil {
		if err := m.Creator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("creator")
			}
			return err
		}
	}

	return nil
}

func (m *DbcSummary) validateStatusChangedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusChangedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("statusChangedAt", "body", "date-time", m.StatusChangedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this dbc summary based on the context it is used
func (m *DbcSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreator(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DbcSummary) contextValidateCreator(ctx context.Context, formats strfmt.Registry) error {

	if m.Creator != nil {

		if swag.IsZero(m.Creator) { // not required
			return nil
		}

		if err := m.Creator.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("creator")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DbcSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DbcSummary) UnmarshalBinary(b []byte) error {
	var res DbcSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
