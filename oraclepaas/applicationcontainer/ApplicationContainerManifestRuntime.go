package applicationcontainer


type ApplicationContainerManifestRuntime struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/application_container#major_version ApplicationContainer#major_version}.
	MajorVersion *string `field:"required" json:"majorVersion" yaml:"majorVersion"`
}

