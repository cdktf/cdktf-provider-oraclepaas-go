package mysqlserviceinstance


type MysqlServiceInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#create MysqlServiceInstance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#delete MysqlServiceInstance#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

