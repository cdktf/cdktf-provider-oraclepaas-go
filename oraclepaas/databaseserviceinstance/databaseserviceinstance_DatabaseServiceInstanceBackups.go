package databaseserviceinstance


type DatabaseServiceInstanceBackups struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#cloud_storage_container DatabaseServiceInstance#cloud_storage_container}.
	CloudStorageContainer *string `field:"required" json:"cloudStorageContainer" yaml:"cloudStorageContainer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#cloud_storage_password DatabaseServiceInstance#cloud_storage_password}.
	CloudStoragePassword *string `field:"optional" json:"cloudStoragePassword" yaml:"cloudStoragePassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#cloud_storage_username DatabaseServiceInstance#cloud_storage_username}.
	CloudStorageUsername *string `field:"optional" json:"cloudStorageUsername" yaml:"cloudStorageUsername"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#create_if_missing DatabaseServiceInstance#create_if_missing}.
	CreateIfMissing interface{} `field:"optional" json:"createIfMissing" yaml:"createIfMissing"`
}

