package mysqlserviceinstance


type MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#em_agent_password MysqlServiceInstance#em_agent_password}.
	EmAgentPassword *string `field:"optional" json:"emAgentPassword" yaml:"emAgentPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#em_agent_username MysqlServiceInstance#em_agent_username}.
	EmAgentUsername *string `field:"optional" json:"emAgentUsername" yaml:"emAgentUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#em_password MysqlServiceInstance#em_password}.
	EmPassword *string `field:"optional" json:"emPassword" yaml:"emPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#em_port MysqlServiceInstance#em_port}.
	EmPort *float64 `field:"optional" json:"emPort" yaml:"emPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#em_username MysqlServiceInstance#em_username}.
	EmUsername *string `field:"optional" json:"emUsername" yaml:"emUsername"`
}

