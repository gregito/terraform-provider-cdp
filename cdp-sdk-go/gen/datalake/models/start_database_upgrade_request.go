// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StartDatabaseUpgradeRequest Request object for Data Lake database upgrade.
//
// swagger:model StartDatabaseUpgradeRequest
type StartDatabaseUpgradeRequest struct {

	// The name or CRN of the Data Lake.
	// Required: true
	Datalake *string `json:"datalake"`

	// The database engine major version to upgrade to.
	// Enum: [VERSION_11]
	TargetVersion string `json:"targetVersion,omitempty"`
}

// Validate validates this start database upgrade request
func (m *StartDatabaseUpgradeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatalake(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StartDatabaseUpgradeRequest) validateDatalake(formats strfmt.Registry) error {

	if err := validate.Required("datalake", "body", m.Datalake); err != nil {
		return err
	}

	return nil
}

var startDatabaseUpgradeRequestTypeTargetVersionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["VERSION_11"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		startDatabaseUpgradeRequestTypeTargetVersionPropEnum = append(startDatabaseUpgradeRequestTypeTargetVersionPropEnum, v)
	}
}

const (

	// StartDatabaseUpgradeRequestTargetVersionVERSION11 captures enum value "VERSION_11"
	StartDatabaseUpgradeRequestTargetVersionVERSION11 string = "VERSION_11"
)

// prop value enum
func (m *StartDatabaseUpgradeRequest) validateTargetVersionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, startDatabaseUpgradeRequestTypeTargetVersionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StartDatabaseUpgradeRequest) validateTargetVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetVersion) { // not required
		return nil
	}

	// value enum
	if err := m.validateTargetVersionEnum("targetVersion", "body", m.TargetVersion); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this start database upgrade request based on context it is used
func (m *StartDatabaseUpgradeRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StartDatabaseUpgradeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StartDatabaseUpgradeRequest) UnmarshalBinary(b []byte) error {
	var res StartDatabaseUpgradeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
