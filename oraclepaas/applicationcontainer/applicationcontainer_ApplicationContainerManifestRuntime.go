package applicationcontainer


type ApplicationContainerManifestRuntime struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/application_container#major_version ApplicationContainer#major_version}.
	MajorVersion *string `field:"required" json:"majorVersion" yaml:"majorVersion"`
}

