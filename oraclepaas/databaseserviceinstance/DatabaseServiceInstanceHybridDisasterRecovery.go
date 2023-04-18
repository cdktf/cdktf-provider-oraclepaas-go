package databaseserviceinstance


type DatabaseServiceInstanceHybridDisasterRecovery struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#cloud_storage_container DatabaseServiceInstance#cloud_storage_container}.
	CloudStorageContainer *string `field:"required" json:"cloudStorageContainer" yaml:"cloudStorageContainer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#cloud_storage_password DatabaseServiceInstance#cloud_storage_password}.
	CloudStoragePassword *string `field:"optional" json:"cloudStoragePassword" yaml:"cloudStoragePassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_service_instance#cloud_storage_username DatabaseServiceInstance#cloud_storage_username}.
	CloudStorageUsername *string `field:"optional" json:"cloudStorageUsername" yaml:"cloudStorageUsername"`
}

