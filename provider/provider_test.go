// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package provider

import (
	"context"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func createCdpProviderModel() *CdpProviderModel {
	return &CdpProviderModel{
		CdpAccessKeyId:           types.StringValue("cdp-access-key"),
		CdpPrivateKey:            types.StringValue("cdp-private-key"),
		Profile:                  types.StringValue("profile"),
		CdpRegion:                types.StringValue("region"),
		AltusEndpointUrl:         types.StringValue("altus-endpoint-url"),
		CdpEndpointUrl:           types.StringValue("cdp-endpoint-url"),
		CdpConfigFile:            types.StringValue("cdp-config-file"),
		CdpSharedCredentialsFile: types.StringValue("cdp-shared-credentials-file"),
		LocalEnvironment:         types.BoolValue(false),
	}
}

func TestProviderOverridesUserAgent(t *testing.T) {
	model := createCdpProviderModel()

	config := getCdpConfig(context.Background(), model, "0.1.0", "v1.4.2")
	userAgent := config.GetUserAgentOrDefault()

	r, _ := regexp.Compile(`^CDPTFPROVIDER/.+ Terraform/.+ Go/.+ .+_.+$`)
	if !r.MatchString(userAgent) {
		t.Fatalf("Failed to match the User-Agent regex: %v", userAgent)
	}
}

func TestProviderClientApplicationName(t *testing.T) {
	model := createCdpProviderModel()

	config := getCdpConfig(context.Background(), model, "0.1.0", "v1.4.2")
	clientApplicationName := config.ClientApplicationName

	if clientApplicationName != "terraform-provider-cdp" {
		t.Fatalf("Terraform provider should have set client application name. Got: %v", clientApplicationName)
	}
}

func TestProviderCdpRegion(t *testing.T) {
	model := createCdpProviderModel()

	config := getCdpConfig(context.Background(), model, "0.1.0", "v1.4.2")

	cdpRegion, err := config.GetCdpRegion()
	if err != nil {
		t.Fatalf("Error getting cdp region: %v", err)
	}
	if cdpRegion != "region" {
		t.Fatalf("Terraform provider should have set cdp region. Got: %v", cdpRegion)
	}
}
