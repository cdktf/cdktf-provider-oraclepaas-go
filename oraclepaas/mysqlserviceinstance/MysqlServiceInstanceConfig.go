package mysqlserviceinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MysqlServiceInstanceConfig struct {
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
	// mysql_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#mysql_configuration MysqlServiceInstance#mysql_configuration}
	MysqlConfiguration *MysqlServiceInstanceMysqlConfiguration `field:"required" json:"mysqlConfiguration" yaml:"mysqlConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#name MysqlServiceInstance#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#shape MysqlServiceInstance#shape}.
	Shape *string `field:"required" json:"shape" yaml:"shape"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#ssh_public_key MysqlServiceInstance#ssh_public_key}.
	SshPublicKey *string `field:"required" json:"sshPublicKey" yaml:"sshPublicKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#availability_domain MysqlServiceInstance#availability_domain}.
	AvailabilityDomain *string `field:"optional" json:"availabilityDomain" yaml:"availabilityDomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#backup_destination MysqlServiceInstance#backup_destination}.
	BackupDestination *string `field:"optional" json:"backupDestination" yaml:"backupDestination"`
	// backups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#backups MysqlServiceInstance#backups}
	Backups *MysqlServiceInstanceBackups `field:"optional" json:"backups" yaml:"backups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#description MysqlServiceInstance#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#id MysqlServiceInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#ip_network MysqlServiceInstance#ip_network}.
	IpNetwork *string `field:"optional" json:"ipNetwork" yaml:"ipNetwork"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#metering_frequency MysqlServiceInstance#metering_frequency}.
	MeteringFrequency *string `field:"optional" json:"meteringFrequency" yaml:"meteringFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#notification_email MysqlServiceInstance#notification_email}.
	NotificationEmail *string `field:"optional" json:"notificationEmail" yaml:"notificationEmail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#region MysqlServiceInstance#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#subnet MysqlServiceInstance#subnet}.
	Subnet *string `field:"optional" json:"subnet" yaml:"subnet"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#timeouts MysqlServiceInstance#timeouts}
	Timeouts *MysqlServiceInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#vm_user MysqlServiceInstance#vm_user}.
	VmUser *string `field:"optional" json:"vmUser" yaml:"vmUser"`
}

