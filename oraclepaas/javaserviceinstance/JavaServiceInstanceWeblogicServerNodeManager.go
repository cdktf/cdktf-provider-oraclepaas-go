package javaserviceinstance


type JavaServiceInstanceWeblogicServerNodeManager struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#password JavaServiceInstance#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#port JavaServiceInstance#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#username JavaServiceInstance#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

