package javaserviceinstance


type JavaServiceInstanceWeblogicServerPorts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#content_port JavaServiceInstance#content_port}.
	ContentPort *float64 `field:"optional" json:"contentPort" yaml:"contentPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#deployment_channel_port JavaServiceInstance#deployment_channel_port}.
	DeploymentChannelPort *float64 `field:"optional" json:"deploymentChannelPort" yaml:"deploymentChannelPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#privileged_content_port JavaServiceInstance#privileged_content_port}.
	PrivilegedContentPort *float64 `field:"optional" json:"privilegedContentPort" yaml:"privilegedContentPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/java_service_instance#privileged_secured_content_port JavaServiceInstance#privileged_secured_content_port}.
	PrivilegedSecuredContentPort *float64 `field:"optional" json:"privilegedSecuredContentPort" yaml:"privilegedSecuredContentPort"`
}

