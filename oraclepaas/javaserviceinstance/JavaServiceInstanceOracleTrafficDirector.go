// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package javaserviceinstance


type JavaServiceInstanceOracleTrafficDirector struct {
	// admin block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#admin JavaServiceInstance#admin}
	Admin *JavaServiceInstanceOracleTrafficDirectorAdmin `field:"required" json:"admin" yaml:"admin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#shape JavaServiceInstance#shape}.
	Shape *string `field:"required" json:"shape" yaml:"shape"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#high_availability JavaServiceInstance#high_availability}.
	HighAvailability interface{} `field:"optional" json:"highAvailability" yaml:"highAvailability"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#ip_reservations JavaServiceInstance#ip_reservations}.
	IpReservations *[]*string `field:"optional" json:"ipReservations" yaml:"ipReservations"`
	// listener block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#listener JavaServiceInstance#listener}
	Listener *JavaServiceInstanceOracleTrafficDirectorListener `field:"optional" json:"listener" yaml:"listener"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#load_balancing_policy JavaServiceInstance#load_balancing_policy}.
	LoadBalancingPolicy *string `field:"optional" json:"loadBalancingPolicy" yaml:"loadBalancingPolicy"`
}

