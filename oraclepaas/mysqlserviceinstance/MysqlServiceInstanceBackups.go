// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mysqlserviceinstance


type MysqlServiceInstanceBackups struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#cloud_storage_container MysqlServiceInstance#cloud_storage_container}.
	CloudStorageContainer *string `field:"required" json:"cloudStorageContainer" yaml:"cloudStorageContainer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#cloud_storage_password MysqlServiceInstance#cloud_storage_password}.
	CloudStoragePassword *string `field:"optional" json:"cloudStoragePassword" yaml:"cloudStoragePassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#cloud_storage_username MysqlServiceInstance#cloud_storage_username}.
	CloudStorageUsername *string `field:"optional" json:"cloudStorageUsername" yaml:"cloudStorageUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#create_if_missing MysqlServiceInstance#create_if_missing}.
	CreateIfMissing interface{} `field:"optional" json:"createIfMissing" yaml:"createIfMissing"`
}

