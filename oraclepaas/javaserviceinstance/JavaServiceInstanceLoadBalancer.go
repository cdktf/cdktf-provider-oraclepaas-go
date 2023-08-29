// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package javaserviceinstance


type JavaServiceInstanceLoadBalancer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#load_balancing_policy JavaServiceInstance#load_balancing_policy}.
	LoadBalancingPolicy *string `field:"optional" json:"loadBalancingPolicy" yaml:"loadBalancingPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#subnets JavaServiceInstance#subnets}.
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

