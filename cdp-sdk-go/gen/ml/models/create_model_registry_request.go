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

// CreateModelRegistryRequest Request object for creating model registry.
//
// swagger:model CreateModelRegistryRequest
type CreateModelRegistryRequest struct {

	// The whitelist of CIDR blocks which can access the API server.
	AuthorizedIPRanges []string `json:"authorizedIPRanges"`

	// The CRN of the backup that this model registry is created from.
	BackupCrn string `json:"backupCrn,omitempty"`

	// The creator of model registry.
	CreatorCrn string `json:"creatorCrn,omitempty"`

	// The environment CRN of model registry.
	EnvironmentCrn string `json:"environmentCrn,omitempty"`

	// The environment for the model registry to create.
	EnvironmentName string `json:"environmentName,omitempty"`

	// The namespace of the model registry.
	Namespace string `json:"namespace,omitempty"`

	// Outbound Types provided for the cluster.
	OutboundTypes []OutboundTypes `json:"outboundTypes"`

	// Whether to create a private cluster.
	PrivateCluster bool `json:"privateCluster,omitempty"`

	// The request for Kubernetes cluster provision. Required in public cloud.
	ProvisionK8sRequest *ModelRegistryProvisionK8sRequest `json:"provisionK8sRequest,omitempty"`

	// The S3 access key of Ozone.
	S3AccessKey string `json:"s3AccessKey,omitempty"`

	// The s3Bucket of Ozone.
	S3Bucket string `json:"s3Bucket,omitempty"`

	// The endpoint of Ozone.
	S3Endpoint string `json:"s3Endpoint,omitempty"`

	// The S3 secret key of Ozone.
	S3SecretKey string `json:"s3SecretKey,omitempty"`

	// Skip pre-flight validations if requested.
	SkipValidation bool `json:"skipValidation,omitempty"`

	// The list of subnets used for the load balancer that CML creates.
	SubnetsForLoadBalancers []string `json:"subnetsForLoadBalancers"`

	// The boolean flag to request a public load balancer. By default, a private load balancer is used.
	UsePublicLoadBalancer bool `json:"usePublicLoadBalancer,omitempty"`

	// Whether to whitelist only authorizedIPRanges given or all public IPs
	WhitelistAuthorizedIPRanges bool `json:"whitelistAuthorizedIPRanges,omitempty"`
}

// Validate validates this create model registry request
func (m *CreateModelRegistryRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOutboundTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvisionK8sRequest(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateModelRegistryRequest) validateOutboundTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.OutboundTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.OutboundTypes); i++ {

		if err := m.OutboundTypes[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outboundTypes" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("outboundTypes" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *CreateModelRegistryRequest) validateProvisionK8sRequest(formats strfmt.Registry) error {
	if swag.IsZero(m.ProvisionK8sRequest) { // not required
		return nil
	}

	if m.ProvisionK8sRequest != nil {
		if err := m.ProvisionK8sRequest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provisionK8sRequest")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("provisionK8sRequest")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create model registry request based on the context it is used
func (m *CreateModelRegistryRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOutboundTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProvisionK8sRequest(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateModelRegistryRequest) contextValidateOutboundTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OutboundTypes); i++ {

		if swag.IsZero(m.OutboundTypes[i]) { // not required
			return nil
		}

		if err := m.OutboundTypes[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outboundTypes" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("outboundTypes" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *CreateModelRegistryRequest) contextValidateProvisionK8sRequest(ctx context.Context, formats strfmt.Registry) error {

	if m.ProvisionK8sRequest != nil {

		if swag.IsZero(m.ProvisionK8sRequest) { // not required
			return nil
		}

		if err := m.ProvisionK8sRequest.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provisionK8sRequest")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("provisionK8sRequest")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateModelRegistryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateModelRegistryRequest) UnmarshalBinary(b []byte) error {
	var res CreateModelRegistryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
