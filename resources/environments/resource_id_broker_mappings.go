// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package environments

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_             resource.ResourceWithConfigure   = &idBrokerMappingsResource{}
	_             resource.ResourceWithImportState = &idBrokerMappingsResource{}
	emptyMappings                                  = true
)

type idBrokerMappingsResource struct {
	client *cdp.Client
}

type idBrokerMapping struct {
	AccessorCrn types.String `tfsdk:"accessor_crn"`

	Role types.String `tfsdk:"role"`
}

func (r *idBrokerMappingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func NewIDBrokerMappingsResource() resource.Resource {
	return &idBrokerMappingsResource{}
}

func (r *idBrokerMappingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environments_id_broker_mappings"
}

func (r *idBrokerMappingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = IDBrokerMappingSchema
}

func (r *idBrokerMappingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *idBrokerMappingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var state idBrokerMappingsResourceModel
	diags := req.Plan.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Got Error while trying to set plan")
		return
	}

	client := r.client.Environments

	params := operations.NewSetIDBrokerMappingsParamsWithContext(ctx)
	params.WithInput(toSetIDBrokerMappingsRequest(ctx, &state, &resp.Diagnostics))
	responseOk, err := client.Operations.SetIDBrokerMappings(params)
	if err != nil {
		if isSetIDBEnvNotFoundError(err) {
			resp.Diagnostics.AddError(
				"Error applying ID Broker mappings",
				"Environment not found: "+state.EnvironmentCrn.ValueString(),
			)
			return
		}
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "create ID Broker mapping")
		return
	}

	idBrokerResp := responseOk.Payload
	state.ID = state.EnvironmentCrn
	state.RangerCloudAccessAuthorizerRole = types.StringValue(idBrokerResp.RangerCloudAccessAuthorizerRole)
	state.MappingsVersion = types.Int64PointerValue(idBrokerResp.MappingsVersion)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *idBrokerMappingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state idBrokerMappingsResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	client := r.client.Environments

	if err := queryEnvironment(ctx, client, state.EnvironmentName.ValueString(), &state); isEnvNotFoundError(err) {
		removeResourceFromState(ctx, &resp.Diagnostics, &resp.State, state)
		return
	}

	params := operations.NewGetIDBrokerMappingsParamsWithContext(ctx)
	params.WithInput(&environmentsmodels.GetIDBrokerMappingsRequest{
		EnvironmentName: state.EnvironmentName.ValueStringPointer(),
	})
	responseOk, err := client.Operations.GetIDBrokerMappings(params)
	if err != nil {
		if envErr, ok := err.(*operations.GetIDBrokerMappingsDefault); ok {
			if cdp.IsEnvironmentsError(envErr.GetPayload(), "NOT_FOUND", "") {
				resp.Diagnostics.AddWarning("Resource not found on provider", "Environment not found, removing from state.")
				tflog.Warn(ctx, "Environment not found, removing from state", map[string]interface{}{
					"id": state.ID.ValueString(),
				})
				resp.State.RemoveResource(ctx)
				return
			}
		}
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "read ID Broker mapping")
		return
	}

	idBrokerResp := responseOk.Payload
	toIdBrokerMappingsResourceModel(ctx, idBrokerResp, &state, &resp.Diagnostics)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *idBrokerMappingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state idBrokerMappingsResourceModel
	diags := req.Plan.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Got Error while trying to set plan")
		return
	}

	client := r.client.Environments

	if err := queryEnvironment(ctx, client, state.EnvironmentName.ValueString(), &state); isEnvNotFoundError(err) {
		removeResourceFromState(ctx, &resp.Diagnostics, &resp.State, state)
		return
	}

	params := operations.NewSetIDBrokerMappingsParamsWithContext(ctx)
	params.WithInput(toSetIDBrokerMappingsRequest(ctx, &state, &resp.Diagnostics))
	responseOk, err := client.Operations.SetIDBrokerMappings(params)
	if err != nil {
		if isSetIDBEnvNotFoundError(err) {
			resp.Diagnostics.AddError(
				"Error applying ID Broker mappings",
				"Environment not found: "+state.EnvironmentCrn.ValueString(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Error setting ID Broker mappings",
			"Got the following error setting ID Broker mappings: "+err.Error(),
		)
		return
	}

	idBrokerResp := responseOk.Payload
	state.MappingsVersion = types.Int64PointerValue(idBrokerResp.MappingsVersion)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *idBrokerMappingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state idBrokerMappingsResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	client := r.client.Environments

	if err := queryEnvironment(ctx, client, state.EnvironmentName.ValueString(), &state); isEnvNotFoundError(err) {
		removeResourceFromState(ctx, &resp.Diagnostics, &resp.State, state)
		return
	}

	params := operations.NewSetIDBrokerMappingsParamsWithContext(ctx)
	input := &environmentsmodels.SetIDBrokerMappingsRequest{}
	input.EnvironmentName = state.EnvironmentName.ValueStringPointer()
	input.DataAccessRole = state.DataAccessRole.ValueStringPointer()
	input.RangerAuditRole = state.RangerAuditRole.ValueString()
	input.Mappings = make([]*environmentsmodels.IDBrokerMappingRequest, 0)
	input.SetEmptyMappings = &emptyMappings
	params.WithInput(input)
	_, err := client.Operations.SetIDBrokerMappings(params)
	if err != nil {
		if isSetIDBEnvNotFoundError(err) {
			resp.Diagnostics.AddError(
				"Error deleting ID Broker mappings",
				"Environment not found: "+state.EnvironmentCrn.ValueString(),
			)
			return
		}
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "delete ID Broker mapping")
		return
	}
}

func queryEnvironment(ctx context.Context, client *client.Environments, envName string, state *idBrokerMappingsResourceModel) error {
	envParams := operations.NewDescribeEnvironmentParamsWithContext(ctx)
	envParams.WithInput(&environmentsmodels.DescribeEnvironmentRequest{
		EnvironmentName: &envName,
	})
	_, err := client.Operations.DescribeEnvironment(envParams)
	return err
}

func isSetIDBEnvNotFoundError(err error) bool {
	if envErr, ok := err.(*operations.SetIDBrokerMappingsDefault); ok {
		if cdp.IsEnvironmentsError(envErr.GetPayload(), "NOT_FOUND", "") {
			return true
		}
	}
	return false
}

func removeResourceFromState(ctx context.Context, diag *diag.Diagnostics, state *tfsdk.State, model idBrokerMappingsResourceModel) {
	diag.AddWarning("Resource not found on provider", "Environment not found, removing ID Broker mapping from state.")
	tflog.Warn(ctx, "Environment not found, removing ID Broker mapping from state", map[string]interface{}{
		"id": model.ID.ValueString(),
	})
	state.RemoveResource(ctx)
}

func toIdBrokerMappingsResourceModel(ctx context.Context, mapping *environmentsmodels.GetIDBrokerMappingsResponse, out *idBrokerMappingsResourceModel, diags *diag.Diagnostics) {
	out.DataAccessRole = types.StringPointerValue(mapping.DataAccessRole)
	out.MappingsVersion = types.Int64PointerValue(mapping.MappingsVersion)
	out.RangerAuditRole = types.StringPointerValue(mapping.RangerAuditRole)
	out.RangerCloudAccessAuthorizerRole = types.StringValue(mapping.RangerCloudAccessAuthorizerRole)
	if len(mapping.Mappings) == 0 {
		out.Mappings = types.SetNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"accessor_crn": types.StringType,
				"role":         types.StringType,
			},
		})
	} else {
		mappings := make([]*idBrokerMapping, len(mapping.Mappings))
		for i, v := range mapping.Mappings {
			mappings[i] = &idBrokerMapping{
				AccessorCrn: types.StringPointerValue(v.AccessorCrn),
				Role:        types.StringPointerValue(v.Role),
			}
		}
		var mappingsDiags diag.Diagnostics
		out.Mappings, mappingsDiags = types.SetValueFrom(ctx, types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"accessor_crn": types.StringType,
				"role":         types.StringType,
			},
		}, mappings)
		diags.Append(mappingsDiags...)
	}
}

func toSetIDBrokerMappingsRequest(ctx context.Context, model *idBrokerMappingsResourceModel, diag *diag.Diagnostics) *environmentsmodels.SetIDBrokerMappingsRequest {
	resp := &environmentsmodels.SetIDBrokerMappingsRequest{}
	resp.DataAccessRole = model.DataAccessRole.ValueStringPointer()
	resp.EnvironmentName = model.EnvironmentName.ValueStringPointer()
	resp.RangerAuditRole = model.RangerAuditRole.ValueString()
	resp.RangerCloudAccessAuthorizerRole = model.RangerCloudAccessAuthorizerRole.ValueString()
	resp.SetEmptyMappings = model.SetEmptyMappings.ValueBoolPointer()
	mappings := make([]*idBrokerMapping, len(model.Mappings.Elements()))
	diag.Append(model.Mappings.ElementsAs(ctx, &mappings, false)...)
	resp.Mappings = make([]*environmentsmodels.IDBrokerMappingRequest, len(mappings))
	for i, v := range mappings {
		resp.Mappings[i] = &environmentsmodels.IDBrokerMappingRequest{
			AccessorCrn: v.AccessorCrn.ValueStringPointer(),
			Role:        v.Role.ValueStringPointer(),
		}
	}
	return resp
}
