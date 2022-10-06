package javaserviceinstance


type JavaServiceInstanceLoadBalancer struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/java_service_instance#load_balancing_policy JavaServiceInstance#load_balancing_policy}.
	LoadBalancingPolicy *string `field:"optional" json:"loadBalancingPolicy" yaml:"loadBalancingPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/java_service_instance#subnets JavaServiceInstance#subnets}.
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

