package applicationcontainer


type ApplicationContainerDeployment struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/application_container#environment ApplicationContainer#environment}.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/application_container#instances ApplicationContainer#instances}.
	Instances *float64 `field:"optional" json:"instances" yaml:"instances"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/application_container#java_system_properties ApplicationContainer#java_system_properties}.
	JavaSystemProperties *map[string]*string `field:"optional" json:"javaSystemProperties" yaml:"javaSystemProperties"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/application_container#memory ApplicationContainer#memory}.
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/application_container#notes ApplicationContainer#notes}.
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/application_container#secure_environment ApplicationContainer#secure_environment}.
	SecureEnvironment *[]*string `field:"optional" json:"secureEnvironment" yaml:"secureEnvironment"`
	// services block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/application_container#services ApplicationContainer#services}
	Services interface{} `field:"optional" json:"services" yaml:"services"`
}

