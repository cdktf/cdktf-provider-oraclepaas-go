package applicationcontainer


type ApplicationContainerTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#create ApplicationContainer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#delete ApplicationContainer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

