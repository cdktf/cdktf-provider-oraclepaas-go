package provider


type OraclepaasProviderConfig struct {
	// The OPAAS identity domain for API operations.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas#identity_domain OraclepaasProvider#identity_domain}
	IdentityDomain *string `field:"required" json:"identityDomain" yaml:"identityDomain"`
	// The user password for OPAAS API operations.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas#password OraclepaasProvider#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// The user name for OPAAS API operations.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas#user OraclepaasProvider#user}
	User *string `field:"required" json:"user" yaml:"user"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas#alias OraclepaasProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The HTTP endpoint for the Oracle Application operations.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas#application_endpoint OraclepaasProvider#application_endpoint}
	ApplicationEndpoint *string `field:"optional" json:"applicationEndpoint" yaml:"applicationEndpoint"`
	// The HTTP endpoint for Oracle Database operations.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas#database_endpoint OraclepaasProvider#database_endpoint}
	DatabaseEndpoint *string `field:"optional" json:"databaseEndpoint" yaml:"databaseEndpoint"`
	// Skip TLS Verification for self-signed certificates. Should only be used if absolutely required.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas#insecure OraclepaasProvider#insecure}
	Insecure interface{} `field:"optional" json:"insecure" yaml:"insecure"`
	// The HTTP endpoint for Oracle Java operations.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas#java_endpoint OraclepaasProvider#java_endpoint}
	JavaEndpoint *string `field:"optional" json:"javaEndpoint" yaml:"javaEndpoint"`
	// Maximum number retries to wait for a successful response when operating on resources within OPAAS (defaults to 1).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas#max_retries OraclepaasProvider#max_retries}
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// The HTTP endpoint for Oracle MySQL operations.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/oraclepaas#mysql_endpoint OraclepaasProvider#mysql_endpoint}
	MysqlEndpoint *string `field:"optional" json:"mysqlEndpoint" yaml:"mysqlEndpoint"`
}

