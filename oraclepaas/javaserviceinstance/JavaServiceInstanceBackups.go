package javaserviceinstance


type JavaServiceInstanceBackups struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#cloud_storage_container JavaServiceInstance#cloud_storage_container}.
	CloudStorageContainer *string `field:"required" json:"cloudStorageContainer" yaml:"cloudStorageContainer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#auto_generate JavaServiceInstance#auto_generate}.
	AutoGenerate interface{} `field:"optional" json:"autoGenerate" yaml:"autoGenerate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#cloud_storage_password JavaServiceInstance#cloud_storage_password}.
	CloudStoragePassword *string `field:"optional" json:"cloudStoragePassword" yaml:"cloudStoragePassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#cloud_storage_username JavaServiceInstance#cloud_storage_username}.
	CloudStorageUsername *string `field:"optional" json:"cloudStorageUsername" yaml:"cloudStorageUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#use_oauth_for_storage JavaServiceInstance#use_oauth_for_storage}.
	UseOauthForStorage interface{} `field:"optional" json:"useOauthForStorage" yaml:"useOauthForStorage"`
}

