package databaseserviceinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabaseServiceInstanceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// database_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#database_configuration DatabaseServiceInstance#database_configuration}
	DatabaseConfiguration *DatabaseServiceInstanceDatabaseConfiguration `field:"required" json:"databaseConfiguration" yaml:"databaseConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#edition DatabaseServiceInstance#edition}.
	Edition *string `field:"required" json:"edition" yaml:"edition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#name DatabaseServiceInstance#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#shape DatabaseServiceInstance#shape}.
	Shape *string `field:"required" json:"shape" yaml:"shape"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#ssh_public_key DatabaseServiceInstance#ssh_public_key}.
	SshPublicKey *string `field:"required" json:"sshPublicKey" yaml:"sshPublicKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#subscription_type DatabaseServiceInstance#subscription_type}.
	SubscriptionType *string `field:"required" json:"subscriptionType" yaml:"subscriptionType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#version DatabaseServiceInstance#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#availability_domain DatabaseServiceInstance#availability_domain}.
	AvailabilityDomain *string `field:"optional" json:"availabilityDomain" yaml:"availabilityDomain"`
	// backups block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#backups DatabaseServiceInstance#backups}
	Backups *DatabaseServiceInstanceBackups `field:"optional" json:"backups" yaml:"backups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#bring_your_own_license DatabaseServiceInstance#bring_your_own_license}.
	BringYourOwnLicense interface{} `field:"optional" json:"bringYourOwnLicense" yaml:"bringYourOwnLicense"`
	// default_access_rules block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#default_access_rules DatabaseServiceInstance#default_access_rules}
	DefaultAccessRules *DatabaseServiceInstanceDefaultAccessRules `field:"optional" json:"defaultAccessRules" yaml:"defaultAccessRules"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#description DatabaseServiceInstance#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#desired_state DatabaseServiceInstance#desired_state}.
	DesiredState *string `field:"optional" json:"desiredState" yaml:"desiredState"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#high_performance_storage DatabaseServiceInstance#high_performance_storage}.
	HighPerformanceStorage interface{} `field:"optional" json:"highPerformanceStorage" yaml:"highPerformanceStorage"`
	// hybrid_disaster_recovery block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#hybrid_disaster_recovery DatabaseServiceInstance#hybrid_disaster_recovery}
	HybridDisasterRecovery *DatabaseServiceInstanceHybridDisasterRecovery `field:"optional" json:"hybridDisasterRecovery" yaml:"hybridDisasterRecovery"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#id DatabaseServiceInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// instantiate_from_backup block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#instantiate_from_backup DatabaseServiceInstance#instantiate_from_backup}
	InstantiateFromBackup *DatabaseServiceInstanceInstantiateFromBackup `field:"optional" json:"instantiateFromBackup" yaml:"instantiateFromBackup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#ip_network DatabaseServiceInstance#ip_network}.
	IpNetwork *string `field:"optional" json:"ipNetwork" yaml:"ipNetwork"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#ip_reservations DatabaseServiceInstance#ip_reservations}.
	IpReservations *[]*string `field:"optional" json:"ipReservations" yaml:"ipReservations"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#level DatabaseServiceInstance#level}.
	Level *string `field:"optional" json:"level" yaml:"level"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#notification_email DatabaseServiceInstance#notification_email}.
	NotificationEmail *string `field:"optional" json:"notificationEmail" yaml:"notificationEmail"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#region DatabaseServiceInstance#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// standby block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#standby DatabaseServiceInstance#standby}
	Standby *DatabaseServiceInstanceStandby `field:"optional" json:"standby" yaml:"standby"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#subnet DatabaseServiceInstance#subnet}.
	Subnet *string `field:"optional" json:"subnet" yaml:"subnet"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#timeouts DatabaseServiceInstance#timeouts}
	Timeouts *DatabaseServiceInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

