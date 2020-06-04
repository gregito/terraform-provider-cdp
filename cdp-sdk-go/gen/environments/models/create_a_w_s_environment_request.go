// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateAWSEnvironmentRequest Request object for a create AWS environment request.
//
// swagger:model CreateAWSEnvironmentRequest
type CreateAWSEnvironmentRequest struct {

	// SSH authentication information for accessing cluster node instances. Users with access to this authentication information have root level access to the Data Lake and Data Hub cluster instances.
	// Required: true
	Authentication *AuthenticationRequest `json:"authentication"`

	// Whether to create private subnets or not.
	CreatePrivateSubnets bool `json:"createPrivateSubnets,omitempty"`

	// Name of the credential to use for the environment.
	// Required: true
	CredentialName *string `json:"credentialName"`

	// An description of the environment.
	Description string `json:"description,omitempty"`

	// Whether to enable SSH tunnelling for the environment.
	EnableTunnel bool `json:"enableTunnel,omitempty"`

	// When this is enabled, diagnostic information about job and query execution is sent to Workload Manager for Data Hub clusters created within this environment.
	EnableWorkloadAnalytics bool `json:"enableWorkloadAnalytics,omitempty"`

	// The name of the environment. Must contain only lowercase letters, numbers and hyphens.
	// Required: true
	EnvironmentName *string `json:"environmentName"`

	// The FreeIPA creation request for the environment
	FreeIpa *AWSFreeIpaCreationRequest `json:"freeIpa,omitempty"`

	// AWS storage configuration for cluster and audit logs.
	// Required: true
	LogStorage *AwsLogStorageRequest `json:"logStorage"`

	// The network CIDR. This will create a VPC along with subnets in multiple Availability Zones.
	NetworkCidr string `json:"networkCidr,omitempty"`

	// Name of the proxy config to use for the environment.
	ProxyConfigName string `json:"proxyConfigName,omitempty"`

	// The region of the environment.
	// Required: true
	Region *string `json:"region"`

	// When true, this will report additional diagnostic information back to Cloudera.
	ReportDeploymentLogs bool `json:"reportDeploymentLogs,omitempty"`

	// The name for the DynamoDB table backing S3Guard.
	S3GuardTableName string `json:"s3GuardTableName,omitempty"`

	// Security control for FreeIPA and Datalake deployment.
	// Required: true
	SecurityAccess *SecurityAccessRequest `json:"securityAccess"`

	// One or more subnet ids within the VPC. Mutually exclusive with networkCidr.
	// Unique: true
	SubnetIds []string `json:"subnetIds"`

	// Tags associated with the resources.
	Tags []*TagRequest `json:"tags"`

	// The Amazon VPC ID. Mutually exclusive with networkCidr.
	VpcID string `json:"vpcId,omitempty"`

	// When this is enabled, diagnostic information about job and query execution is sent to Workload Manager for Data Hub clusters created within this environment.
	WorkloadAnalytics bool `json:"workloadAnalytics,omitempty"`
}

// Validate validates this create a w s environment request
func (m *CreateAWSEnvironmentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFreeIpa(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnetIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAWSEnvironmentRequest) validateAuthentication(formats strfmt.Registry) error {

	if err := validate.Required("authentication", "body", m.Authentication); err != nil {
		return err
	}

	if m.Authentication != nil {
		if err := m.Authentication.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authentication")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAWSEnvironmentRequest) validateCredentialName(formats strfmt.Registry) error {

	if err := validate.Required("credentialName", "body", m.CredentialName); err != nil {
		return err
	}

	return nil
}

func (m *CreateAWSEnvironmentRequest) validateEnvironmentName(formats strfmt.Registry) error {

	if err := validate.Required("environmentName", "body", m.EnvironmentName); err != nil {
		return err
	}

	return nil
}

func (m *CreateAWSEnvironmentRequest) validateFreeIpa(formats strfmt.Registry) error {

	if swag.IsZero(m.FreeIpa) { // not required
		return nil
	}

	if m.FreeIpa != nil {
		if err := m.FreeIpa.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("freeIpa")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAWSEnvironmentRequest) validateLogStorage(formats strfmt.Registry) error {

	if err := validate.Required("logStorage", "body", m.LogStorage); err != nil {
		return err
	}

	if m.LogStorage != nil {
		if err := m.LogStorage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logStorage")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAWSEnvironmentRequest) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

func (m *CreateAWSEnvironmentRequest) validateSecurityAccess(formats strfmt.Registry) error {

	if err := validate.Required("securityAccess", "body", m.SecurityAccess); err != nil {
		return err
	}

	if m.SecurityAccess != nil {
		if err := m.SecurityAccess.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityAccess")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAWSEnvironmentRequest) validateSubnetIds(formats strfmt.Registry) error {

	if swag.IsZero(m.SubnetIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("subnetIds", "body", m.SubnetIds); err != nil {
		return err
	}

	return nil
}

func (m *CreateAWSEnvironmentRequest) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateAWSEnvironmentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAWSEnvironmentRequest) UnmarshalBinary(b []byte) error {
	var res CreateAWSEnvironmentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
