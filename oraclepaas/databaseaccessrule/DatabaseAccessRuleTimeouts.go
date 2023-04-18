package databaseaccessrule


type DatabaseAccessRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_access_rule#create DatabaseAccessRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/database_access_rule#delete DatabaseAccessRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

