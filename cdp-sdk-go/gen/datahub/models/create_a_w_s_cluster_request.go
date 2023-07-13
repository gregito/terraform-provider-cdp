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

// CreateAWSClusterRequest Request object for create AWS cluster request.
//
// swagger:model CreateAWSClusterRequest
type CreateAWSClusterRequest struct {

	// The name or CRN of the cluster definition to use for cluster creation.
	ClusterDefinition string `json:"clusterDefinition,omitempty"`

	// Deprecated, the name or CRN of the cluster definition to use for cluster creation.
	ClusterDefinitionName string `json:"clusterDefinitionName,omitempty"`

	// Cluster extensions for the given Data Hub cluster.
	ClusterExtension *ClusterExtension `json:"clusterExtension,omitempty"`

	// The name of the cluster. This name must be unique, must have between 5 and 40 characters, and must contain only lowercase letters, numbers and hyphens. Names are case-sensitive.
	// Max Length: 40
	// Min Length: 5
	ClusterName string `json:"clusterName,omitempty"`

	// Name or CRN of the cluster template to use for cluster creation.
	ClusterTemplate string `json:"clusterTemplate,omitempty"`

	// Deprecated, name or CRN of the cluster template to use for cluster creation.
	ClusterTemplateName string `json:"clusterTemplateName,omitempty"`

	// The name of the custom configurations to use for cluster creation.
	CustomConfigurationsName string `json:"customConfigurationsName,omitempty"`

	// Database type for datahub. Currently supported values: NONE, NON_HA, HA
	DatahubDatabase DatahubDatabaseType `json:"datahubDatabase,omitempty"`

	// Flag that decides whether to provision a load-balancer to front various service endpoints for the given datahub. This will typically be used for HA cluster shapes.
	EnableLoadBalancer bool `json:"enableLoadBalancer,omitempty"`

	// Name or CRN of the environment to use when creating the cluster. The environment must be an AWS environment.
	Environment string `json:"environment,omitempty"`

	// Deprecated, name or CRN of the environment to use when creating the cluster. The environment must be an AWS environment.
	EnvironmentName string `json:"environmentName,omitempty"`

	// The image to be used for cluster creation.
	Image *ImageRequest `json:"image,omitempty"`

	// Instance group details.
	InstanceGroups []*InstanceGroupRequest `json:"instanceGroups"`

	// Configure the major version of Java on the cluster.
	JavaVersion int32 `json:"javaVersion,omitempty"`

	// Flag that toggles the multi availability zone for the given datahub cluster when you are not sure what subnet IDs can be used. This way the subnet IDs will be used what the environment suggests.
	MultiAz bool `json:"multiAz,omitempty"`

	// JSON template to use for cluster creation. This is different from cluster template and would be removed in the future.
	RequestTemplate string `json:"requestTemplate,omitempty"`

	// The subnet ID.
	SubnetID string `json:"subnetId,omitempty"`

	// List of subnet IDs in case of multi availability zone setup.
	SubnetIds []string `json:"subnetIds"`

	// Tags to be added to Datahub related resources.
	Tags []*DatahubResourceTagRequest `json:"tags"`
}

// Validate validates this create a w s cluster request
func (m *CreateAWSClusterRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterExtension(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatahubDatabase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceGroups(formats); err != nil {
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

func (m *CreateAWSClusterRequest) validateClusterExtension(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterExtension) { // not required
		return nil
	}

	if m.ClusterExtension != nil {
		if err := m.ClusterExtension.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterExtension")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterExtension")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAWSClusterRequest) validateClusterName(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterName) { // not required
		return nil
	}

	if err := validate.MinLength("clusterName", "body", m.ClusterName, 5); err != nil {
		return err
	}

	if err := validate.MaxLength("clusterName", "body", m.ClusterName, 40); err != nil {
		return err
	}

	return nil
}

func (m *CreateAWSClusterRequest) validateDatahubDatabase(formats strfmt.Registry) error {
	if swag.IsZero(m.DatahubDatabase) { // not required
		return nil
	}

	if err := m.DatahubDatabase.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("datahubDatabase")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("datahubDatabase")
		}
		return err
	}

	return nil
}

func (m *CreateAWSClusterRequest) validateImage(formats strfmt.Registry) error {
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

func (m *CreateAWSClusterRequest) validateInstanceGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.InstanceGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.InstanceGroups); i++ {
		if swag.IsZero(m.InstanceGroups[i]) { // not required
			continue
		}

		if m.InstanceGroups[i] != nil {
			if err := m.InstanceGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instanceGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("instanceGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateAWSClusterRequest) validateTags(formats strfmt.Registry) error {
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

// ContextValidate validate this create a w s cluster request based on the context it is used
func (m *CreateAWSClusterRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusterExtension(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDatahubDatabase(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInstanceGroups(ctx, formats); err != nil {
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

func (m *CreateAWSClusterRequest) contextValidateClusterExtension(ctx context.Context, formats strfmt.Registry) error {

	if m.ClusterExtension != nil {

		if swag.IsZero(m.ClusterExtension) { // not required
			return nil
		}

		if err := m.ClusterExtension.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterExtension")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterExtension")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAWSClusterRequest) contextValidateDatahubDatabase(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.DatahubDatabase) { // not required
		return nil
	}

	if err := m.DatahubDatabase.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("datahubDatabase")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("datahubDatabase")
		}
		return err
	}

	return nil
}

func (m *CreateAWSClusterRequest) contextValidateImage(ctx context.Context, formats strfmt.Registry) error {

	if m.Image != nil {

		if swag.IsZero(m.Image) { // not required
			return nil
		}

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

func (m *CreateAWSClusterRequest) contextValidateInstanceGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InstanceGroups); i++ {

		if m.InstanceGroups[i] != nil {

			if swag.IsZero(m.InstanceGroups[i]) { // not required
				return nil
			}

			if err := m.InstanceGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instanceGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("instanceGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateAWSClusterRequest) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {

			if swag.IsZero(m.Tags[i]) { // not required
				return nil
			}

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
func (m *CreateAWSClusterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAWSClusterRequest) UnmarshalBinary(b []byte) error {
	var res CreateAWSClusterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
