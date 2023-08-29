// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package javaserviceinstance


type JavaServiceInstanceOracleTrafficDirectorListener struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#port JavaServiceInstance#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#privileged_port JavaServiceInstance#privileged_port}.
	PrivilegedPort *float64 `field:"optional" json:"privilegedPort" yaml:"privilegedPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#privileged_secured_port JavaServiceInstance#privileged_secured_port}.
	PrivilegedSecuredPort *float64 `field:"optional" json:"privilegedSecuredPort" yaml:"privilegedSecuredPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#secured_port JavaServiceInstance#secured_port}.
	SecuredPort *float64 `field:"optional" json:"securedPort" yaml:"securedPort"`
}

