// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package javaserviceinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JavaServiceInstanceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// backups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#backups JavaServiceInstance#backups}
	Backups *JavaServiceInstanceBackups `field:"required" json:"backups" yaml:"backups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#edition JavaServiceInstance#edition}.
	Edition *string `field:"required" json:"edition" yaml:"edition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#name JavaServiceInstance#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#ssh_public_key JavaServiceInstance#ssh_public_key}.
	SshPublicKey *string `field:"required" json:"sshPublicKey" yaml:"sshPublicKey"`
	// weblogic_server block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#weblogic_server JavaServiceInstance#weblogic_server}
	WeblogicServer *JavaServiceInstanceWeblogicServer `field:"required" json:"weblogicServer" yaml:"weblogicServer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#assign_public_ip JavaServiceInstance#assign_public_ip}.
	AssignPublicIp interface{} `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#availability_domain JavaServiceInstance#availability_domain}.
	AvailabilityDomain *string `field:"optional" json:"availabilityDomain" yaml:"availabilityDomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#backup_destination JavaServiceInstance#backup_destination}.
	BackupDestination *string `field:"optional" json:"backupDestination" yaml:"backupDestination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#bring_your_own_license JavaServiceInstance#bring_your_own_license}.
	BringYourOwnLicense interface{} `field:"optional" json:"bringYourOwnLicense" yaml:"bringYourOwnLicense"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#description JavaServiceInstance#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#desired_state JavaServiceInstance#desired_state}.
	DesiredState *string `field:"optional" json:"desiredState" yaml:"desiredState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#enable_admin_console JavaServiceInstance#enable_admin_console}.
	EnableAdminConsole interface{} `field:"optional" json:"enableAdminConsole" yaml:"enableAdminConsole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#force_delete JavaServiceInstance#force_delete}.
	ForceDelete interface{} `field:"optional" json:"forceDelete" yaml:"forceDelete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#id JavaServiceInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#ip_network JavaServiceInstance#ip_network}.
	IpNetwork *string `field:"optional" json:"ipNetwork" yaml:"ipNetwork"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#level JavaServiceInstance#level}.
	Level *string `field:"optional" json:"level" yaml:"level"`
	// load_balancer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#load_balancer JavaServiceInstance#load_balancer}
	LoadBalancer *JavaServiceInstanceLoadBalancer `field:"optional" json:"loadBalancer" yaml:"loadBalancer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#metering_frequency JavaServiceInstance#metering_frequency}.
	MeteringFrequency *string `field:"optional" json:"meteringFrequency" yaml:"meteringFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#notification_email JavaServiceInstance#notification_email}.
	NotificationEmail *string `field:"optional" json:"notificationEmail" yaml:"notificationEmail"`
	// oracle_traffic_director block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#oracle_traffic_director JavaServiceInstance#oracle_traffic_director}
	OracleTrafficDirector *JavaServiceInstanceOracleTrafficDirector `field:"optional" json:"oracleTrafficDirector" yaml:"oracleTrafficDirector"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#region JavaServiceInstance#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#service_version JavaServiceInstance#service_version}.
	ServiceVersion *string `field:"optional" json:"serviceVersion" yaml:"serviceVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#snapshot_name JavaServiceInstance#snapshot_name}.
	SnapshotName *string `field:"optional" json:"snapshotName" yaml:"snapshotName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#source_service_name JavaServiceInstance#source_service_name}.
	SourceServiceName *string `field:"optional" json:"sourceServiceName" yaml:"sourceServiceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#subnet JavaServiceInstance#subnet}.
	Subnet *string `field:"optional" json:"subnet" yaml:"subnet"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#timeouts JavaServiceInstance#timeouts}
	Timeouts *JavaServiceInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#use_identity_service JavaServiceInstance#use_identity_service}.
	UseIdentityService interface{} `field:"optional" json:"useIdentityService" yaml:"useIdentityService"`
}

