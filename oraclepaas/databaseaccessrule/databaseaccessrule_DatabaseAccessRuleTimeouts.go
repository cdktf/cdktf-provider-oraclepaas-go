package databaseaccessrule


type DatabaseAccessRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_access_rule#create DatabaseAccessRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_access_rule#delete DatabaseAccessRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

