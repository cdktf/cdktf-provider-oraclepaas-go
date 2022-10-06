package applicationcontainer


type ApplicationContainerManifestRelease struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/application_container#build ApplicationContainer#build}.
	BuildAttribute *string `field:"optional" json:"buildAttribute" yaml:"buildAttribute"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/application_container#commit ApplicationContainer#commit}.
	Commit *string `field:"optional" json:"commit" yaml:"commit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/application_container#version ApplicationContainer#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

