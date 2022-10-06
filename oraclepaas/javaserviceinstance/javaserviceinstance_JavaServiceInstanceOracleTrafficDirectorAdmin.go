package javaserviceinstance


type JavaServiceInstanceOracleTrafficDirectorAdmin struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/java_service_instance#password JavaServiceInstance#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/java_service_instance#username JavaServiceInstance#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/java_service_instance#port JavaServiceInstance#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

