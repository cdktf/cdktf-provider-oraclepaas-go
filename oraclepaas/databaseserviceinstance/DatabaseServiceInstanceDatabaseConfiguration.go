// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databaseserviceinstance


type DatabaseServiceInstanceDatabaseConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#admin_password DatabaseServiceInstance#admin_password}.
	AdminPassword *string `field:"required" json:"adminPassword" yaml:"adminPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#usable_storage DatabaseServiceInstance#usable_storage}.
	UsableStorage *float64 `field:"required" json:"usableStorage" yaml:"usableStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#backup_destination DatabaseServiceInstance#backup_destination}.
	BackupDestination *string `field:"optional" json:"backupDestination" yaml:"backupDestination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#backup_storage_volume_size DatabaseServiceInstance#backup_storage_volume_size}.
	BackupStorageVolumeSize *float64 `field:"optional" json:"backupStorageVolumeSize" yaml:"backupStorageVolumeSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#character_set DatabaseServiceInstance#character_set}.
	CharacterSet *string `field:"optional" json:"characterSet" yaml:"characterSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#data_storage_volume_size DatabaseServiceInstance#data_storage_volume_size}.
	DataStorageVolumeSize *float64 `field:"optional" json:"dataStorageVolumeSize" yaml:"dataStorageVolumeSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#db_demo DatabaseServiceInstance#db_demo}.
	DbDemo *string `field:"optional" json:"dbDemo" yaml:"dbDemo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#disaster_recovery DatabaseServiceInstance#disaster_recovery}.
	DisasterRecovery interface{} `field:"optional" json:"disasterRecovery" yaml:"disasterRecovery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#failover_database DatabaseServiceInstance#failover_database}.
	FailoverDatabase interface{} `field:"optional" json:"failoverDatabase" yaml:"failoverDatabase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#golden_gate DatabaseServiceInstance#golden_gate}.
	GoldenGate interface{} `field:"optional" json:"goldenGate" yaml:"goldenGate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#is_rac DatabaseServiceInstance#is_rac}.
	IsRac interface{} `field:"optional" json:"isRac" yaml:"isRac"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#national_character_set DatabaseServiceInstance#national_character_set}.
	NationalCharacterSet *string `field:"optional" json:"nationalCharacterSet" yaml:"nationalCharacterSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#pdb_name DatabaseServiceInstance#pdb_name}.
	PdbName *string `field:"optional" json:"pdbName" yaml:"pdbName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#sid DatabaseServiceInstance#sid}.
	Sid *string `field:"optional" json:"sid" yaml:"sid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#snapshot_name DatabaseServiceInstance#snapshot_name}.
	SnapshotName *string `field:"optional" json:"snapshotName" yaml:"snapshotName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#source_service_name DatabaseServiceInstance#source_service_name}.
	SourceServiceName *string `field:"optional" json:"sourceServiceName" yaml:"sourceServiceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#timezone DatabaseServiceInstance#timezone}.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#type DatabaseServiceInstance#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

