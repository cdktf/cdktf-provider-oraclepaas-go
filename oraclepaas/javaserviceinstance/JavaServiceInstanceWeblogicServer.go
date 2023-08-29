// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package javaserviceinstance


type JavaServiceInstanceWeblogicServer struct {
	// admin block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#admin JavaServiceInstance#admin}
	Admin *JavaServiceInstanceWeblogicServerAdmin `field:"required" json:"admin" yaml:"admin"`
	// database block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#database JavaServiceInstance#database}
	Database *JavaServiceInstanceWeblogicServerDatabase `field:"required" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#shape JavaServiceInstance#shape}.
	Shape *string `field:"required" json:"shape" yaml:"shape"`
	// application_database block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#application_database JavaServiceInstance#application_database}
	ApplicationDatabase interface{} `field:"optional" json:"applicationDatabase" yaml:"applicationDatabase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#backup_volume_size JavaServiceInstance#backup_volume_size}.
	BackupVolumeSize *string `field:"optional" json:"backupVolumeSize" yaml:"backupVolumeSize"`
	// cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#cluster JavaServiceInstance#cluster}
	Cluster interface{} `field:"optional" json:"cluster" yaml:"cluster"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#cluster_name JavaServiceInstance#cluster_name}.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#connect_string JavaServiceInstance#connect_string}.
	ConnectString *string `field:"optional" json:"connectString" yaml:"connectString"`
	// domain block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#domain JavaServiceInstance#domain}
	Domain *JavaServiceInstanceWeblogicServerDomain `field:"optional" json:"domain" yaml:"domain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#ip_reservations JavaServiceInstance#ip_reservations}.
	IpReservations *[]*string `field:"optional" json:"ipReservations" yaml:"ipReservations"`
	// managed_servers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#managed_servers JavaServiceInstance#managed_servers}
	ManagedServers *JavaServiceInstanceWeblogicServerManagedServers `field:"optional" json:"managedServers" yaml:"managedServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#middleware_volume_size JavaServiceInstance#middleware_volume_size}.
	MiddlewareVolumeSize *string `field:"optional" json:"middlewareVolumeSize" yaml:"middlewareVolumeSize"`
	// node_manager block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#node_manager JavaServiceInstance#node_manager}
	NodeManager *JavaServiceInstanceWeblogicServerNodeManager `field:"optional" json:"nodeManager" yaml:"nodeManager"`
	// ports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#ports JavaServiceInstance#ports}
	Ports *JavaServiceInstanceWeblogicServerPorts `field:"optional" json:"ports" yaml:"ports"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#upper_stack_product_name JavaServiceInstance#upper_stack_product_name}.
	UpperStackProductName *string `field:"optional" json:"upperStackProductName" yaml:"upperStackProductName"`
}

