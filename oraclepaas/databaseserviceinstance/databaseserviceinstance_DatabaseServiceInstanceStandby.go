package databaseserviceinstance


type DatabaseServiceInstanceStandby struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#availability_domain DatabaseServiceInstance#availability_domain}.
	AvailabilityDomain *string `field:"required" json:"availabilityDomain" yaml:"availabilityDomain"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance#subnet DatabaseServiceInstance#subnet}.
	Subnet *string `field:"required" json:"subnet" yaml:"subnet"`
}

