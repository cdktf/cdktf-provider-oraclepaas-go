package javaserviceinstance


type JavaServiceInstanceWeblogicServerManagedServers struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/java_service_instance#initial_heap_size JavaServiceInstance#initial_heap_size}.
	InitialHeapSize *float64 `field:"optional" json:"initialHeapSize" yaml:"initialHeapSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/java_service_instance#initial_permanent_generation JavaServiceInstance#initial_permanent_generation}.
	InitialPermanentGeneration *float64 `field:"optional" json:"initialPermanentGeneration" yaml:"initialPermanentGeneration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/java_service_instance#jvm_args JavaServiceInstance#jvm_args}.
	JvmArgs *string `field:"optional" json:"jvmArgs" yaml:"jvmArgs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/java_service_instance#max_heap_size JavaServiceInstance#max_heap_size}.
	MaxHeapSize *float64 `field:"optional" json:"maxHeapSize" yaml:"maxHeapSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/java_service_instance#max_permanent_generation JavaServiceInstance#max_permanent_generation}.
	MaxPermanentGeneration *float64 `field:"optional" json:"maxPermanentGeneration" yaml:"maxPermanentGeneration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/java_service_instance#overwrite_jvm_args JavaServiceInstance#overwrite_jvm_args}.
	OverwriteJvmArgs interface{} `field:"optional" json:"overwriteJvmArgs" yaml:"overwriteJvmArgs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/java_service_instance#server_count JavaServiceInstance#server_count}.
	ServerCount *float64 `field:"optional" json:"serverCount" yaml:"serverCount"`
}

