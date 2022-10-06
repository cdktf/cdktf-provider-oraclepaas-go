package javaaccessrule


type JavaAccessRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/java_access_rule#create JavaAccessRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/java_access_rule#delete JavaAccessRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

