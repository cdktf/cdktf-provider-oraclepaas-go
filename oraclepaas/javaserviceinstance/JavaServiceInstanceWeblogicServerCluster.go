package javaserviceinstance


type JavaServiceInstanceWeblogicServerCluster struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#name JavaServiceInstance#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#type JavaServiceInstance#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#path_prefixes JavaServiceInstance#path_prefixes}.
	PathPrefixes *[]*string `field:"optional" json:"pathPrefixes" yaml:"pathPrefixes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#server_count JavaServiceInstance#server_count}.
	ServerCount *float64 `field:"optional" json:"serverCount" yaml:"serverCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#servers_per_node JavaServiceInstance#servers_per_node}.
	ServersPerNode *float64 `field:"optional" json:"serversPerNode" yaml:"serversPerNode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#shape JavaServiceInstance#shape}.
	Shape *string `field:"optional" json:"shape" yaml:"shape"`
}

