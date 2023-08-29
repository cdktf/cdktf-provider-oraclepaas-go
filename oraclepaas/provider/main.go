// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-oraclepaas.provider.OraclepaasProvider",
		reflect.TypeOf((*OraclepaasProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasInput", GoGetter: "AliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "applicationEndpoint", GoGetter: "ApplicationEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "applicationEndpointInput", GoGetter: "ApplicationEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "databaseEndpoint", GoGetter: "DatabaseEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "databaseEndpointInput", GoGetter: "DatabaseEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "identityDomain", GoGetter: "IdentityDomain"},
			_jsii_.MemberProperty{JsiiProperty: "identityDomainInput", GoGetter: "IdentityDomainInput"},
			_jsii_.MemberProperty{JsiiProperty: "insecure", GoGetter: "Insecure"},
			_jsii_.MemberProperty{JsiiProperty: "insecureInput", GoGetter: "InsecureInput"},
			_jsii_.MemberProperty{JsiiProperty: "javaEndpoint", GoGetter: "JavaEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "javaEndpointInput", GoGetter: "JavaEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetries", GoGetter: "MaxRetries"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetriesInput", GoGetter: "MaxRetriesInput"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "mysqlEndpoint", GoGetter: "MysqlEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "mysqlEndpointInput", GoGetter: "MysqlEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "password", GoGetter: "Password"},
			_jsii_.MemberProperty{JsiiProperty: "passwordInput", GoGetter: "PasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetApplicationEndpoint", GoMethod: "ResetApplicationEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetDatabaseEndpoint", GoMethod: "ResetDatabaseEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetInsecure", GoMethod: "ResetInsecure"},
			_jsii_.MemberMethod{JsiiMethod: "resetJavaEndpoint", GoMethod: "ResetJavaEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxRetries", GoMethod: "ResetMaxRetries"},
			_jsii_.MemberMethod{JsiiMethod: "resetMysqlEndpoint", GoMethod: "ResetMysqlEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "user", GoGetter: "User"},
			_jsii_.MemberProperty{JsiiProperty: "userInput", GoGetter: "UserInput"},
		},
		func() interface{} {
			j := jsiiProxy_OraclepaasProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-oraclepaas.provider.OraclepaasProviderConfig",
		reflect.TypeOf((*OraclepaasProviderConfig)(nil)).Elem(),
	)
}
