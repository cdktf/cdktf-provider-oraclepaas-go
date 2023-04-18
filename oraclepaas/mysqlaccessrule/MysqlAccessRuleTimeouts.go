package mysqlaccessrule


type MysqlAccessRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_access_rule#create MysqlAccessRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_access_rule#delete MysqlAccessRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

