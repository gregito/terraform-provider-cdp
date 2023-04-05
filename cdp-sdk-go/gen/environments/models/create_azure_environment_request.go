// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateAzureEnvironmentRequest Request object for a create Azure environment request.
//
// swagger:model CreateAzureEnvironmentRequest
type CreateAzureEnvironmentRequest struct {

	// CCM V2 TLS connectivity types
	CcmV2TLSType CcmV2TLSType `json:"ccmV2TlsType,omitempty"`

	// When this is enabled, logs from the VMs will end up on the pre-defined cloud storage (enabled by default).
	CloudStorageLogging bool `json:"cloudStorageLogging,omitempty"`

	// When this is enabled, then Azure Postgres will be configured with Private Endpoint and a Private DNS Zone. When this is disabled, then Azure Service Endpoints will be created. The default value is disabled.
	CreatePrivateEndpoints bool `json:"createPrivateEndpoints,omitempty"`

	// Name of the credential to use for the environment.
	// Required: true
	CredentialName *string `json:"credentialName"`

	// When this is enabled, a dedicated storage account will be used in a pre-defined resource group for storing the images in each region.
	DedicatedStorageAccount bool `json:"dedicatedStorageAccount,omitempty"`

	// An description of the environment.
	Description string `json:"description,omitempty"`

	// Whether or not Knox and Oozie load balancers are enabled for all Data Lakes and Data Hubs in the environment. This will override the load balancer creation mode at the cluster level. The default behavior is to create load balancers on all Data Lakes, and on HA Data Hubs.
	EnableLoadBalancers bool `json:"enableLoadBalancers,omitempty"`

	// Whether or not outbound load balancers should be created for Azure environments. The default behavior is to not create the outbound load balancer.
	EnableOutboundLoadBalancer bool `json:"enableOutboundLoadBalancer,omitempty"`

	// Whether to enable SSH tunneling for the environment.
	EnableTunnel bool `json:"enableTunnel,omitempty"`

	// When this is enabled, diagnostic information about job and query execution is sent to Workload Manager for Data Hub clusters created within this environment.
	EnableWorkloadAnalytics bool `json:"enableWorkloadAnalytics,omitempty"`

	// Name of the existing Azure resource group hosting the Azure Key Vault containing customer managed key which will be used to encrypt the Azure Managed Disks. It is required only when the entitlement is granted and the resource group of the key vault is different from the resource group in which the environment is to be created. Omitting it implies that, the key vault containing the encryption key is present in the same resource group where the environment would be created.
	EncryptionKeyResourceGroupName string `json:"encryptionKeyResourceGroupName,omitempty"`

	// URL of the key which will be used to encrypt the Azure Managed Disks, if entitlement has been granted.
	EncryptionKeyURL string `json:"encryptionKeyUrl,omitempty"`

	// The scheme for the endpoint gateway. PUBLIC creates an external endpoint that can be accessed over the Internet. Defaults to PRIVATE which restricts the traffic to be internal to the VNet.
	// Enum: [PUBLIC PRIVATE]
	EndpointAccessGatewayScheme string `json:"endpointAccessGatewayScheme,omitempty"`

	// The name of the environment. Must contain only lowercase letters, numbers and hyphens.
	// Required: true
	EnvironmentName *string `json:"environmentName"`

	// Parameters needed to use an existing VNet and Subnets.
	ExistingNetworkParams *ExistingAzureNetworkRequest `json:"existingNetworkParams,omitempty"`

	// The FreeIPA creation request for the environment
	FreeIpa *AzureFreeIpaCreationRequest `json:"freeIpa,omitempty"`

	// This is an optional field. This is for QE testing purposes and internal use only.
	IDBrokerMappingSource string `json:"idBrokerMappingSource,omitempty"`

	// The FreeIPA image request for the environment
	Image *FreeIpaImageRequest `json:"image,omitempty"`

	// Azure storage configuration for cluster and audit logs.
	// Required: true
	LogStorage *AzureLogStorageRequest `json:"logStorage"`

	// new network params
	NewNetworkParams *CreateAzureEnvironmentRequestNewNetworkParams `json:"newNetworkParams,omitempty"`

	// Name of the proxy config to use for the environment.
	ProxyConfigName string `json:"proxyConfigName,omitempty"`

	// Public SSH key string. The associated private key can be used to get root-level access to the Data Lake instance and Data Hub cluster instances.
	// Required: true
	PublicKey *string `json:"publicKey"`

	// The region of the environment.
	// Required: true
	Region *string `json:"region"`

	// When true, this will report additional diagnostic information back to Cloudera.
	ReportDeploymentLogs bool `json:"reportDeploymentLogs,omitempty"`

	// Name of an existing Azure resource group to be used for the environment. If it is not specified then new resource groups will be generated.
	ResourceGroupName string `json:"resourceGroupName,omitempty"`

	// Security control for FreeIPA and Data Lake deployment.
	// Required: true
	SecurityAccess *SecurityAccessRequest `json:"securityAccess"`

	// Tags associated with the resources.
	Tags []*TagRequest `json:"tags"`

	// The CCM version that will be used for tunneling
	TunnelType TunnelType `json:"tunnelType,omitempty"`

	// Whether to associate public ip's to the resources within the network.
	// Required: true
	UsePublicIP *bool `json:"usePublicIp"`

	// When this is enabled, diagnostic information about job and query execution is sent to Workload Manager for Data Hub clusters created within this environment.
	WorkloadAnalytics bool `json:"workloadAnalytics,omitempty"`
}

// Validate validates this create azure environment request
func (m *CreateAzureEnvironmentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCcmV2TLSType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndpointAccessGatewayScheme(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExistingNetworkParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFreeIpa(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewNetworkParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTunnelType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsePublicIP(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAzureEnvironmentRequest) validateCcmV2TLSType(formats strfmt.Registry) error {
	if swag.IsZero(m.CcmV2TLSType) { // not required
		return nil
	}

	if err := m.CcmV2TLSType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ccmV2TlsType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ccmV2TlsType")
		}
		return err
	}

	return nil
}

func (m *CreateAzureEnvironmentRequest) validateCredentialName(formats strfmt.Registry) error {

	if err := validate.Required("credentialName", "body", m.CredentialName); err != nil {
		return err
	}

	return nil
}

var createAzureEnvironmentRequestTypeEndpointAccessGatewaySchemePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PUBLIC","PRIVATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createAzureEnvironmentRequestTypeEndpointAccessGatewaySchemePropEnum = append(createAzureEnvironmentRequestTypeEndpointAccessGatewaySchemePropEnum, v)
	}
}

const (

	// CreateAzureEnvironmentRequestEndpointAccessGatewaySchemePUBLIC captures enum value "PUBLIC"
	CreateAzureEnvironmentRequestEndpointAccessGatewaySchemePUBLIC string = "PUBLIC"

	// CreateAzureEnvironmentRequestEndpointAccessGatewaySchemePRIVATE captures enum value "PRIVATE"
	CreateAzureEnvironmentRequestEndpointAccessGatewaySchemePRIVATE string = "PRIVATE"
)

// prop value enum
func (m *CreateAzureEnvironmentRequest) validateEndpointAccessGatewaySchemeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createAzureEnvironmentRequestTypeEndpointAccessGatewaySchemePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateAzureEnvironmentRequest) validateEndpointAccessGatewayScheme(formats strfmt.Registry) error {
	if swag.IsZero(m.EndpointAccessGatewayScheme) { // not required
		return nil
	}

	// value enum
	if err := m.validateEndpointAccessGatewaySchemeEnum("endpointAccessGatewayScheme", "body", m.EndpointAccessGatewayScheme); err != nil {
		return err
	}

	return nil
}

func (m *CreateAzureEnvironmentRequest) validateEnvironmentName(formats strfmt.Registry) error {

	if err := validate.Required("environmentName", "body", m.EnvironmentName); err != nil {
		return err
	}

	return nil
}

func (m *CreateAzureEnvironmentRequest) validateExistingNetworkParams(formats strfmt.Registry) error {
	if swag.IsZero(m.ExistingNetworkParams) { // not required
		return nil
	}

	if m.ExistingNetworkParams != nil {
		if err := m.ExistingNetworkParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("existingNetworkParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("existingNetworkParams")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAzureEnvironmentRequest) validateFreeIpa(formats strfmt.Registry) error {
	if swag.IsZero(m.FreeIpa) { // not required
		return nil
	}

	if m.FreeIpa != nil {
		if err := m.FreeIpa.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("freeIpa")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("freeIpa")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAzureEnvironmentRequest) validateImage(formats strfmt.Registry) error {
	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if m.Image != nil {
		if err := m.Image.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAzureEnvironmentRequest) validateLogStorage(formats strfmt.Registry) error {

	if err := validate.Required("logStorage", "body", m.LogStorage); err != nil {
		return err
	}

	if m.LogStorage != nil {
		if err := m.LogStorage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logStorage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("logStorage")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAzureEnvironmentRequest) validateNewNetworkParams(formats strfmt.Registry) error {
	if swag.IsZero(m.NewNetworkParams) { // not required
		return nil
	}

	if m.NewNetworkParams != nil {
		if err := m.NewNetworkParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("newNetworkParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("newNetworkParams")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAzureEnvironmentRequest) validatePublicKey(formats strfmt.Registry) error {

	if err := validate.Required("publicKey", "body", m.PublicKey); err != nil {
		return err
	}

	return nil
}

func (m *CreateAzureEnvironmentRequest) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

func (m *CreateAzureEnvironmentRequest) validateSecurityAccess(formats strfmt.Registry) error {

	if err := validate.Required("securityAccess", "body", m.SecurityAccess); err != nil {
		return err
	}

	if m.SecurityAccess != nil {
		if err := m.SecurityAccess.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityAccess")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("securityAccess")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAzureEnvironmentRequest) validateTags(formats strfmt.Registry) error {
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateAzureEnvironmentRequest) validateTunnelType(formats strfmt.Registry) error {
	if swag.IsZero(m.TunnelType) { // not required
		return nil
	}

	if err := m.TunnelType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tunnelType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("tunnelType")
		}
		return err
	}

	return nil
}

func (m *CreateAzureEnvironmentRequest) validateUsePublicIP(formats strfmt.Registry) error {

	if err := validate.Required("usePublicIp", "body", m.UsePublicIP); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create azure environment request based on the context it is used
func (m *CreateAzureEnvironmentRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCcmV2TLSType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExistingNetworkParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFreeIpa(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLogStorage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNewNetworkParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityAccess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTunnelType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAzureEnvironmentRequest) contextValidateCcmV2TLSType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.CcmV2TLSType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ccmV2TlsType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ccmV2TlsType")
		}
		return err
	}

	return nil
}

func (m *CreateAzureEnvironmentRequest) contextValidateExistingNetworkParams(ctx context.Context, formats strfmt.Registry) error {

	if m.ExistingNetworkParams != nil {
		if err := m.ExistingNetworkParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("existingNetworkParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("existingNetworkParams")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAzureEnvironmentRequest) contextValidateFreeIpa(ctx context.Context, formats strfmt.Registry) error {

	if m.FreeIpa != nil {
		if err := m.FreeIpa.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("freeIpa")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("freeIpa")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAzureEnvironmentRequest) contextValidateImage(ctx context.Context, formats strfmt.Registry) error {

	if m.Image != nil {
		if err := m.Image.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAzureEnvironmentRequest) contextValidateLogStorage(ctx context.Context, formats strfmt.Registry) error {

	if m.LogStorage != nil {
		if err := m.LogStorage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logStorage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("logStorage")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAzureEnvironmentRequest) contextValidateNewNetworkParams(ctx context.Context, formats strfmt.Registry) error {

	if m.NewNetworkParams != nil {
		if err := m.NewNetworkParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("newNetworkParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("newNetworkParams")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAzureEnvironmentRequest) contextValidateSecurityAccess(ctx context.Context, formats strfmt.Registry) error {

	if m.SecurityAccess != nil {
		if err := m.SecurityAccess.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityAccess")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("securityAccess")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAzureEnvironmentRequest) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateAzureEnvironmentRequest) contextValidateTunnelType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.TunnelType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tunnelType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("tunnelType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateAzureEnvironmentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAzureEnvironmentRequest) UnmarshalBinary(b []byte) error {
	var res CreateAzureEnvironmentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CreateAzureEnvironmentRequestNewNetworkParams Parameteres needed to automatically create VNet and Subnets.
//
// swagger:model CreateAzureEnvironmentRequestNewNetworkParams
type CreateAzureEnvironmentRequestNewNetworkParams struct {

	// The range of private IPv4 addresses that resources will use under the created VNet.
	// Required: true
	NetworkCidr *string `json:"networkCidr"`
}

// Validate validates this create azure environment request new network params
func (m *CreateAzureEnvironmentRequestNewNetworkParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworkCidr(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAzureEnvironmentRequestNewNetworkParams) validateNetworkCidr(formats strfmt.Registry) error {

	if err := validate.Required("newNetworkParams"+"."+"networkCidr", "body", m.NetworkCidr); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create azure environment request new network params based on context it is used
func (m *CreateAzureEnvironmentRequestNewNetworkParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateAzureEnvironmentRequestNewNetworkParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAzureEnvironmentRequestNewNetworkParams) UnmarshalBinary(b []byte) error {
	var res CreateAzureEnvironmentRequestNewNetworkParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
