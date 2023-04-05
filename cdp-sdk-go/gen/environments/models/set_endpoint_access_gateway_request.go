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

// SetEndpointAccessGatewayRequest Request object to set endpoint access gateway settings.
//
// swagger:model SetEndpointAccessGatewayRequest
type SetEndpointAccessGatewayRequest struct {

	// The scheme for the endpoint gateway. PUBLIC creates an external endpoint that can be accessed over internet. Defaults to PRIVATE which restricts the traffic to be internal to the VPC / Vnet.
	// Required: true
	// Enum: [PUBLIC PRIVATE]
	EndpointAccessGatewayScheme *string `json:"endpointAccessGatewayScheme"`

	// The subnets to use for endpoint access gateway.
	EndpointAccessGatewaySubnetIds []string `json:"endpointAccessGatewaySubnetIds"`

	// The name or CRN of the environment. Empty to get system wide settings.
	// Required: true
	Environment *string `json:"environment"`
}

// Validate validates this set endpoint access gateway request
func (m *SetEndpointAccessGatewayRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndpointAccessGatewayScheme(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var setEndpointAccessGatewayRequestTypeEndpointAccessGatewaySchemePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PUBLIC","PRIVATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		setEndpointAccessGatewayRequestTypeEndpointAccessGatewaySchemePropEnum = append(setEndpointAccessGatewayRequestTypeEndpointAccessGatewaySchemePropEnum, v)
	}
}

const (

	// SetEndpointAccessGatewayRequestEndpointAccessGatewaySchemePUBLIC captures enum value "PUBLIC"
	SetEndpointAccessGatewayRequestEndpointAccessGatewaySchemePUBLIC string = "PUBLIC"

	// SetEndpointAccessGatewayRequestEndpointAccessGatewaySchemePRIVATE captures enum value "PRIVATE"
	SetEndpointAccessGatewayRequestEndpointAccessGatewaySchemePRIVATE string = "PRIVATE"
)

// prop value enum
func (m *SetEndpointAccessGatewayRequest) validateEndpointAccessGatewaySchemeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, setEndpointAccessGatewayRequestTypeEndpointAccessGatewaySchemePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SetEndpointAccessGatewayRequest) validateEndpointAccessGatewayScheme(formats strfmt.Registry) error {

	if err := validate.Required("endpointAccessGatewayScheme", "body", m.EndpointAccessGatewayScheme); err != nil {
		return err
	}

	// value enum
	if err := m.validateEndpointAccessGatewaySchemeEnum("endpointAccessGatewayScheme", "body", *m.EndpointAccessGatewayScheme); err != nil {
		return err
	}

	return nil
}

func (m *SetEndpointAccessGatewayRequest) validateEnvironment(formats strfmt.Registry) error {

	if err := validate.Required("environment", "body", m.Environment); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this set endpoint access gateway request based on context it is used
func (m *SetEndpointAccessGatewayRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SetEndpointAccessGatewayRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetEndpointAccessGatewayRequest) UnmarshalBinary(b []byte) error {
	var res SetEndpointAccessGatewayRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
