// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LoadBalancer The load balancer configuration properties.
//
// swagger:model LoadBalancer
type LoadBalancer struct {

	// The AWS resource ID for the load balancer.
	AwsResourceID string `json:"awsResourceID,omitempty"`

	// The Azure resource ID for the load balancer.
	AzureResourceID string `json:"azureResourceID,omitempty"`

	// The generated DNS name for the load balancer.
	CloudDNS string `json:"cloudDns,omitempty"`

	// The registered FQDN of the load balancer.
	Fqdn string `json:"fqdn,omitempty"`

	// The GCP resource ID for the load balancer.
	GcpResourceID string `json:"gcpResourceID,omitempty"`

	// The frontend IP address of the load balancer.
	IP string `json:"ip,omitempty"`

	// Whether the load balancer is internet-facing (public), or only accessible over private endpoints.
	LoadBalancerType LoadBalancerType `json:"loadBalancerType,omitempty"`

	// The target instances the load balancer routes traffic to.
	Targets []*TargetGroup `json:"targets"`
}

// Validate validates this load balancer
func (m *LoadBalancer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLoadBalancerType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoadBalancer) validateLoadBalancerType(formats strfmt.Registry) error {
	if swag.IsZero(m.LoadBalancerType) { // not required
		return nil
	}

	if err := m.LoadBalancerType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("loadBalancerType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("loadBalancerType")
		}
		return err
	}

	return nil
}

func (m *LoadBalancer) validateTargets(formats strfmt.Registry) error {
	if swag.IsZero(m.Targets) { // not required
		return nil
	}

	for i := 0; i < len(m.Targets); i++ {
		if swag.IsZero(m.Targets[i]) { // not required
			continue
		}

		if m.Targets[i] != nil {
			if err := m.Targets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("targets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("targets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this load balancer based on the context it is used
func (m *LoadBalancer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLoadBalancerType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTargets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoadBalancer) contextValidateLoadBalancerType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LoadBalancerType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("loadBalancerType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("loadBalancerType")
		}
		return err
	}

	return nil
}

func (m *LoadBalancer) contextValidateTargets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Targets); i++ {

		if m.Targets[i] != nil {
			if err := m.Targets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("targets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("targets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoadBalancer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoadBalancer) UnmarshalBinary(b []byte) error {
	var res LoadBalancer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
