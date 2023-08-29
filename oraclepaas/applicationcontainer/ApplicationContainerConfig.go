// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationcontainer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationContainerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#name ApplicationContainer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#archive_url ApplicationContainer#archive_url}.
	ArchiveUrl *string `field:"optional" json:"archiveUrl" yaml:"archiveUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#auth_type ApplicationContainer#auth_type}.
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#availability_domain ApplicationContainer#availability_domain}.
	AvailabilityDomain *[]*string `field:"optional" json:"availabilityDomain" yaml:"availabilityDomain"`
	// deployment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#deployment ApplicationContainer#deployment}
	Deployment *ApplicationContainerDeployment `field:"optional" json:"deployment" yaml:"deployment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#deployment_file ApplicationContainer#deployment_file}.
	DeploymentFile *string `field:"optional" json:"deploymentFile" yaml:"deploymentFile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#git_password ApplicationContainer#git_password}.
	GitPassword *string `field:"optional" json:"gitPassword" yaml:"gitPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#git_repository ApplicationContainer#git_repository}.
	GitRepository *string `field:"optional" json:"gitRepository" yaml:"gitRepository"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#git_username ApplicationContainer#git_username}.
	GitUsername *string `field:"optional" json:"gitUsername" yaml:"gitUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#id ApplicationContainer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#load_balancer_subnets ApplicationContainer#load_balancer_subnets}.
	LoadBalancerSubnets *[]*string `field:"optional" json:"loadBalancerSubnets" yaml:"loadBalancerSubnets"`
	// manifest block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#manifest ApplicationContainer#manifest}
	Manifest *ApplicationContainerManifest `field:"optional" json:"manifest" yaml:"manifest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#manifest_file ApplicationContainer#manifest_file}.
	ManifestFile *string `field:"optional" json:"manifestFile" yaml:"manifestFile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#notes ApplicationContainer#notes}.
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#notification_email ApplicationContainer#notification_email}.
	NotificationEmail *string `field:"optional" json:"notificationEmail" yaml:"notificationEmail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#region ApplicationContainer#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#runtime ApplicationContainer#runtime}.
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#subscription_type ApplicationContainer#subscription_type}.
	SubscriptionType *string `field:"optional" json:"subscriptionType" yaml:"subscriptionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#tags ApplicationContainer#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#timeouts ApplicationContainer#timeouts}
	Timeouts *ApplicationContainerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

