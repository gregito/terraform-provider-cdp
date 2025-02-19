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

// UpdateClusterRequest Request object for the updateCluster method.
//
// swagger:model UpdateClusterRequest
type UpdateClusterRequest struct {

	// aws update
	AwsUpdate *UpdateClusterRequestAwsUpdate `json:"awsUpdate,omitempty"`

	// azure update
	AzureUpdate *UpdateClusterRequestAzureUpdate `json:"azureUpdate,omitempty"`

	// The ID of the cluster to update.
	// Required: true
	ClusterID *string `json:"clusterId"`

	// Cluster description.
	Description string `json:"description,omitempty"`

	// observability config
	ObservabilityConfig *UpdateClusterRequestObservabilityConfig `json:"observabilityConfig,omitempty"`

	// List of IP address CIDRs to whitelist. NOTE: CDW is in process of rolling out a new feature to whitelist IP CIDR separately for Kubernetes Clusters and Loadbalancers on CDP Public Cloud. For an existing cluster, if different IP CIDR has been already applied to LoadBalancer and the Kubernetes cluster through the DWX UI, then updating the IP CIDR of such cluster is not supported from CLI. In the upcoming release, the CLI will support this feature. Please make use of UI for the same.
	WhitelistIPCIDRs []string `json:"whitelistIpCIDRs"`

	// List of IP address CIDRs to whitelist for kubernetes cluster access.
	WhitelistK8sClusterAccessIPCIDRs []string `json:"whitelistK8sClusterAccessIpCIDRs"`

	// List of IP address CIDRs to whitelist for workload access.
	WhitelistWorkloadAccessIPCIDRs []string `json:"whitelistWorkloadAccessIpCIDRs"`
}

// Validate validates this update cluster request
func (m *UpdateClusterRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsUpdate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureUpdate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObservabilityConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateClusterRequest) validateAwsUpdate(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsUpdate) { // not required
		return nil
	}

	if m.AwsUpdate != nil {
		if err := m.AwsUpdate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsUpdate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsUpdate")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateClusterRequest) validateAzureUpdate(formats strfmt.Registry) error {
	if swag.IsZero(m.AzureUpdate) { // not required
		return nil
	}

	if m.AzureUpdate != nil {
		if err := m.AzureUpdate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureUpdate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azureUpdate")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateClusterRequest) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("clusterId", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *UpdateClusterRequest) validateObservabilityConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.ObservabilityConfig) { // not required
		return nil
	}

	if m.ObservabilityConfig != nil {
		if err := m.ObservabilityConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("observabilityConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("observabilityConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update cluster request based on the context it is used
func (m *UpdateClusterRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAwsUpdate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzureUpdate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateObservabilityConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateClusterRequest) contextValidateAwsUpdate(ctx context.Context, formats strfmt.Registry) error {

	if m.AwsUpdate != nil {

		if swag.IsZero(m.AwsUpdate) { // not required
			return nil
		}

		if err := m.AwsUpdate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsUpdate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsUpdate")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateClusterRequest) contextValidateAzureUpdate(ctx context.Context, formats strfmt.Registry) error {

	if m.AzureUpdate != nil {

		if swag.IsZero(m.AzureUpdate) { // not required
			return nil
		}

		if err := m.AzureUpdate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureUpdate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azureUpdate")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateClusterRequest) contextValidateObservabilityConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.ObservabilityConfig != nil {

		if swag.IsZero(m.ObservabilityConfig) { // not required
			return nil
		}

		if err := m.ObservabilityConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("observabilityConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("observabilityConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateClusterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateClusterRequest) UnmarshalBinary(b []byte) error {
	var res UpdateClusterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UpdateClusterRequestAwsUpdate Additional properties for AWS clusters.
//
// swagger:model UpdateClusterRequestAwsUpdate
type UpdateClusterRequestAwsUpdate struct {

	// External bucket definition.
	ExternalBuckets map[string]ExternalBucketAccessInfo `json:"externalBuckets,omitempty"`
}

// Validate validates this update cluster request aws update
func (m *UpdateClusterRequestAwsUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExternalBuckets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateClusterRequestAwsUpdate) validateExternalBuckets(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalBuckets) { // not required
		return nil
	}

	for k := range m.ExternalBuckets {

		if err := validate.Required("awsUpdate"+"."+"externalBuckets"+"."+k, "body", m.ExternalBuckets[k]); err != nil {
			return err
		}
		if val, ok := m.ExternalBuckets[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("awsUpdate" + "." + "externalBuckets" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("awsUpdate" + "." + "externalBuckets" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update cluster request aws update based on the context it is used
func (m *UpdateClusterRequestAwsUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExternalBuckets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateClusterRequestAwsUpdate) contextValidateExternalBuckets(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.ExternalBuckets {

		if val, ok := m.ExternalBuckets[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateClusterRequestAwsUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateClusterRequestAwsUpdate) UnmarshalBinary(b []byte) error {
	var res UpdateClusterRequestAwsUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UpdateClusterRequestAzureUpdate Additional properties for Azure clusters.
//
// swagger:model UpdateClusterRequestAzureUpdate
type UpdateClusterRequestAzureUpdate struct {

	// Renew Azure cluster certificate.
	RenewCertificate *bool `json:"renewCertificate,omitempty"`
}

// Validate validates this update cluster request azure update
func (m *UpdateClusterRequestAzureUpdate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update cluster request azure update based on context it is used
func (m *UpdateClusterRequestAzureUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateClusterRequestAzureUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateClusterRequestAzureUpdate) UnmarshalBinary(b []byte) error {
	var res UpdateClusterRequestAzureUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UpdateClusterRequestObservabilityConfig Update the cluster observability configurations. You can forward logs from environments activated in Cloudera Data Warehouse (CDW) to observability and monitoring systems such as Datadog, New Relic, or Splun. Please refer to the following Cloudera documentation for more info. https://docs.cloudera.com/data-warehouse/cloud/monitoring/topics/dw-observability-log-forwarding.html
//
// swagger:model UpdateClusterRequestObservabilityConfig
type UpdateClusterRequestObservabilityConfig struct {

	// Create the log forwarding configuration in a valid fluentd format. Then that configuration is later inserted into a larger fluentd configuration.
	LogsForwardingConfig string `json:"logsForwardingConfig,omitempty"`

	// Set the proxy CA certificates (PEM Bundle). If you use a TLS-terminating proxy server to inspect outbound internet traffic, you need to provide the proxy server's CA certificates bundle in PEM bundle format when you configure log forwarding.
	ProxyCABundle string `json:"proxyCABundle,omitempty"`
}

// Validate validates this update cluster request observability config
func (m *UpdateClusterRequestObservabilityConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update cluster request observability config based on context it is used
func (m *UpdateClusterRequestObservabilityConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateClusterRequestObservabilityConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateClusterRequestObservabilityConfig) UnmarshalBinary(b []byte) error {
	var res UpdateClusterRequestObservabilityConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
