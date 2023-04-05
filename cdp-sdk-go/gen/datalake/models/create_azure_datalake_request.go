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
	"github.com/go-openapi/validate"
)

// CreateAzureDatalakeRequest Request object for create Azure datalake request.
//
// swagger:model CreateAzureDatalakeRequest
type CreateAzureDatalakeRequest struct {

	// Azure configuration.
	// Required: true
	CloudProviderConfiguration *AzureConfigurationRequest `json:"cloudProviderConfiguration"`

	// Configure custom properties on an instance group level.
	CustomInstanceGroups []*SdxInstanceGroupRequest `json:"customInstanceGroups"`

	// This is an optional field. This is for QE testing purposes and internal use only. QE can pass this to modify database availability type.
	DatabaseAvailabilityType DatabaseAvailabilityType `json:"databaseAvailabilityType,omitempty"`

	// This is an optional field to set database engine version for the external database.
	DatabaseEngineVersion string `json:"databaseEngineVersion,omitempty"`

	// The datalake name. This name must be unique, must have between 5 and 100 characters, and must contain only lowercase letters, numbers and hyphens. Names are case-sensitive.
	// Required: true
	// Max Length: 100
	// Min Length: 5
	DatalakeName *string `json:"datalakeName"`

	// The datalake template to use for internal datalake creation.
	DatalakeTemplate string `json:"datalakeTemplate,omitempty"`

	// Used to enable or disable cloud load balancers on the Data Lake. By default they are enabled.
	EnableLoadBalancer bool `json:"enableLoadBalancer,omitempty"`

	// Whether to enable Ranger RAZ for the datalake. Defaults to not being enabled.
	EnableRangerRaz bool `json:"enableRangerRaz,omitempty"`

	// The environment name or CRN.
	// Required: true
	EnvironmentName *string `json:"environmentName"`

	// The image to use for the datalake. This must not be set if the runtime parameter is provided.
	Image *ImageRequest `json:"image,omitempty"`

	// Configure the major version of Java on the cluster.
	JavaVersion int32 `json:"javaVersion,omitempty"`

	// The SKU for the datalake load balancer. Allowed values are "BASIC", "STANDARD", or "NONE".
	LoadBalancerSku DatalakeLoadBalancerSkuType `json:"loadBalancerSku,omitempty"`

	// Additional recipes that will be attached on the datalake instances (by instance groups, most common ones are like 'master' or 'idbroker').
	Recipes []*InstanceGroupRecipeRequest `json:"recipes"`

	// Cloudera Runtime version.
	Runtime string `json:"runtime,omitempty"`

	// The scale of the datalake. Allowed values are "LIGHT_DUTY" or "MEDIUM_DUTY_HA". Defaults to "LIGHT_DUTY" if not set.
	Scale DatalakeScaleType `json:"scale,omitempty"`

	// Tags to be added to Data Lake related resources.
	Tags []*DatalakeResourceTagRequest `json:"tags"`
}

// Validate validates this create azure datalake request
func (m *CreateAzureDatalakeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudProviderConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomInstanceGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatabaseAvailabilityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatalakeName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoadBalancerSku(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScale(formats); err != nil {
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

func (m *CreateAzureDatalakeRequest) validateCloudProviderConfiguration(formats strfmt.Registry) error {

	if err := validate.Required("cloudProviderConfiguration", "body", m.CloudProviderConfiguration); err != nil {
		return err
	}

	if m.CloudProviderConfiguration != nil {
		if err := m.CloudProviderConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudProviderConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloudProviderConfiguration")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAzureDatalakeRequest) validateCustomInstanceGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomInstanceGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.CustomInstanceGroups); i++ {
		if swag.IsZero(m.CustomInstanceGroups[i]) { // not required
			continue
		}

		if m.CustomInstanceGroups[i] != nil {
			if err := m.CustomInstanceGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customInstanceGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customInstanceGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateAzureDatalakeRequest) validateDatabaseAvailabilityType(formats strfmt.Registry) error {
	if swag.IsZero(m.DatabaseAvailabilityType) { // not required
		return nil
	}

	if err := m.DatabaseAvailabilityType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("databaseAvailabilityType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("databaseAvailabilityType")
		}
		return err
	}

	return nil
}

func (m *CreateAzureDatalakeRequest) validateDatalakeName(formats strfmt.Registry) error {

	if err := validate.Required("datalakeName", "body", m.DatalakeName); err != nil {
		return err
	}

	if err := validate.MinLength("datalakeName", "body", *m.DatalakeName, 5); err != nil {
		return err
	}

	if err := validate.MaxLength("datalakeName", "body", *m.DatalakeName, 100); err != nil {
		return err
	}

	return nil
}

func (m *CreateAzureDatalakeRequest) validateEnvironmentName(formats strfmt.Registry) error {

	if err := validate.Required("environmentName", "body", m.EnvironmentName); err != nil {
		return err
	}

	return nil
}

func (m *CreateAzureDatalakeRequest) validateImage(formats strfmt.Registry) error {
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

func (m *CreateAzureDatalakeRequest) validateLoadBalancerSku(formats strfmt.Registry) error {
	if swag.IsZero(m.LoadBalancerSku) { // not required
		return nil
	}

	if err := m.LoadBalancerSku.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("loadBalancerSku")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("loadBalancerSku")
		}
		return err
	}

	return nil
}

func (m *CreateAzureDatalakeRequest) validateRecipes(formats strfmt.Registry) error {
	if swag.IsZero(m.Recipes) { // not required
		return nil
	}

	for i := 0; i < len(m.Recipes); i++ {
		if swag.IsZero(m.Recipes[i]) { // not required
			continue
		}

		if m.Recipes[i] != nil {
			if err := m.Recipes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recipes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("recipes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateAzureDatalakeRequest) validateScale(formats strfmt.Registry) error {
	if swag.IsZero(m.Scale) { // not required
		return nil
	}

	if err := m.Scale.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("scale")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("scale")
		}
		return err
	}

	return nil
}

func (m *CreateAzureDatalakeRequest) validateTags(formats strfmt.Registry) error {
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

// ContextValidate validate this create azure datalake request based on the context it is used
func (m *CreateAzureDatalakeRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCloudProviderConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustomInstanceGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDatabaseAvailabilityType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLoadBalancerSku(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecipes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScale(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAzureDatalakeRequest) contextValidateCloudProviderConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if m.CloudProviderConfiguration != nil {
		if err := m.CloudProviderConfiguration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudProviderConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloudProviderConfiguration")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAzureDatalakeRequest) contextValidateCustomInstanceGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CustomInstanceGroups); i++ {

		if m.CustomInstanceGroups[i] != nil {
			if err := m.CustomInstanceGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customInstanceGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customInstanceGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateAzureDatalakeRequest) contextValidateDatabaseAvailabilityType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.DatabaseAvailabilityType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("databaseAvailabilityType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("databaseAvailabilityType")
		}
		return err
	}

	return nil
}

func (m *CreateAzureDatalakeRequest) contextValidateImage(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CreateAzureDatalakeRequest) contextValidateLoadBalancerSku(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LoadBalancerSku.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("loadBalancerSku")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("loadBalancerSku")
		}
		return err
	}

	return nil
}

func (m *CreateAzureDatalakeRequest) contextValidateRecipes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Recipes); i++ {

		if m.Recipes[i] != nil {
			if err := m.Recipes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recipes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("recipes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateAzureDatalakeRequest) contextValidateScale(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Scale.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("scale")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("scale")
		}
		return err
	}

	return nil
}

func (m *CreateAzureDatalakeRequest) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *CreateAzureDatalakeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAzureDatalakeRequest) UnmarshalBinary(b []byte) error {
	var res CreateAzureDatalakeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
