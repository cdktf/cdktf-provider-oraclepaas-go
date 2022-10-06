package applicationcontainer


type ApplicationContainerTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/application_container#create ApplicationContainer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/application_container#delete ApplicationContainer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

