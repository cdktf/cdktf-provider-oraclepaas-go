// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationcontainer


type ApplicationContainerManifest struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#clustered ApplicationContainer#clustered}.
	Clustered interface{} `field:"optional" json:"clustered" yaml:"clustered"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#command ApplicationContainer#command}.
	Command *string `field:"optional" json:"command" yaml:"command"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#health_check_endpoint ApplicationContainer#health_check_endpoint}.
	HealthCheckEndpoint *string `field:"optional" json:"healthCheckEndpoint" yaml:"healthCheckEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#home ApplicationContainer#home}.
	Home *string `field:"optional" json:"home" yaml:"home"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#mode ApplicationContainer#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#notes ApplicationContainer#notes}.
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// release block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#release ApplicationContainer#release}
	Release *ApplicationContainerManifestRelease `field:"optional" json:"release" yaml:"release"`
	// runtime block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#runtime ApplicationContainer#runtime}
	Runtime *ApplicationContainerManifestRuntime `field:"optional" json:"runtime" yaml:"runtime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#shutdown_time ApplicationContainer#shutdown_time}.
	ShutdownTime *float64 `field:"optional" json:"shutdownTime" yaml:"shutdownTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#startup_time ApplicationContainer#startup_time}.
	StartupTime *float64 `field:"optional" json:"startupTime" yaml:"startupTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#type ApplicationContainer#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

