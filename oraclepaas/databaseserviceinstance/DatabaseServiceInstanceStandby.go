// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databaseserviceinstance


type DatabaseServiceInstanceStandby struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#availability_domain DatabaseServiceInstance#availability_domain}.
	AvailabilityDomain *string `field:"required" json:"availabilityDomain" yaml:"availabilityDomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#subnet DatabaseServiceInstance#subnet}.
	Subnet *string `field:"required" json:"subnet" yaml:"subnet"`
}

