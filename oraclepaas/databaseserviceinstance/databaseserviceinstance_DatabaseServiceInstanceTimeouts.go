package databaseserviceinstance


type DatabaseServiceInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#create DatabaseServiceInstance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#delete DatabaseServiceInstance#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#update DatabaseServiceInstance#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

