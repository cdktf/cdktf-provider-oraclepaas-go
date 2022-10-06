package mysqlserviceinstance


type MysqlServiceInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/mysql_service_instance#create MysqlServiceInstance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas/r/mysql_service_instance#delete MysqlServiceInstance#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

