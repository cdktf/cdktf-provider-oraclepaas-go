// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mysqlserviceinstance


type MysqlServiceInstanceMysqlConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#db_name MysqlServiceInstance#db_name}.
	DbName *string `field:"optional" json:"dbName" yaml:"dbName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#db_storage MysqlServiceInstance#db_storage}.
	DbStorage *float64 `field:"optional" json:"dbStorage" yaml:"dbStorage"`
	// enterprise_monitor_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#enterprise_monitor_configuration MysqlServiceInstance#enterprise_monitor_configuration}
	EnterpriseMonitorConfiguration *MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfiguration `field:"optional" json:"enterpriseMonitorConfiguration" yaml:"enterpriseMonitorConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#mysql_charset MysqlServiceInstance#mysql_charset}.
	MysqlCharset *string `field:"optional" json:"mysqlCharset" yaml:"mysqlCharset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#mysql_collation MysqlServiceInstance#mysql_collation}.
	MysqlCollation *string `field:"optional" json:"mysqlCollation" yaml:"mysqlCollation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#mysql_password MysqlServiceInstance#mysql_password}.
	MysqlPassword *string `field:"optional" json:"mysqlPassword" yaml:"mysqlPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#mysql_port MysqlServiceInstance#mysql_port}.
	MysqlPort *float64 `field:"optional" json:"mysqlPort" yaml:"mysqlPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#mysql_username MysqlServiceInstance#mysql_username}.
	MysqlUsername *string `field:"optional" json:"mysqlUsername" yaml:"mysqlUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#snapshot_name MysqlServiceInstance#snapshot_name}.
	SnapshotName *string `field:"optional" json:"snapshotName" yaml:"snapshotName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/oraclepaas/1.5.3/docs/resources/mysql_service_instance#source_service_name MysqlServiceInstance#source_service_name}.
	SourceServiceName *string `field:"optional" json:"sourceServiceName" yaml:"sourceServiceName"`
}

