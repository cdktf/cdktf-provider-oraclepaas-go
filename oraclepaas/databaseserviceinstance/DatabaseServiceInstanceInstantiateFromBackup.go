// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databaseserviceinstance


type DatabaseServiceInstanceInstantiateFromBackup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#cloud_storage_container DatabaseServiceInstance#cloud_storage_container}.
	CloudStorageContainer *string `field:"required" json:"cloudStorageContainer" yaml:"cloudStorageContainer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#database_id DatabaseServiceInstance#database_id}.
	DatabaseId *string `field:"required" json:"databaseId" yaml:"databaseId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#cloud_storage_password DatabaseServiceInstance#cloud_storage_password}.
	CloudStoragePassword *string `field:"optional" json:"cloudStoragePassword" yaml:"cloudStoragePassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#cloud_storage_username DatabaseServiceInstance#cloud_storage_username}.
	CloudStorageUsername *string `field:"optional" json:"cloudStorageUsername" yaml:"cloudStorageUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#decryption_key DatabaseServiceInstance#decryption_key}.
	DecryptionKey *string `field:"optional" json:"decryptionKey" yaml:"decryptionKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#on_premise DatabaseServiceInstance#on_premise}.
	OnPremise interface{} `field:"optional" json:"onPremise" yaml:"onPremise"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#service_id DatabaseServiceInstance#service_id}.
	ServiceId *string `field:"optional" json:"serviceId" yaml:"serviceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#wallet_file_content DatabaseServiceInstance#wallet_file_content}.
	WalletFileContent *string `field:"optional" json:"walletFileContent" yaml:"walletFileContent"`
}

