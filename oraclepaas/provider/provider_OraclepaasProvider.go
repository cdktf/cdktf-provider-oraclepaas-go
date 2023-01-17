package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/v3/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/oraclepaas oraclepaas}.
type OraclepaasProvider interface {
	cdktf.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	ApplicationEndpoint() *string
	SetApplicationEndpoint(val *string)
	ApplicationEndpointInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	DatabaseEndpoint() *string
	SetDatabaseEndpoint(val *string)
	DatabaseEndpointInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IdentityDomain() *string
	SetIdentityDomain(val *string)
	IdentityDomainInput() *string
	Insecure() interface{}
	SetInsecure(val interface{})
	InsecureInput() interface{}
	JavaEndpoint() *string
	SetJavaEndpoint(val *string)
	JavaEndpointInput() *string
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	MaxRetriesInput() *float64
	// Experimental.
	MetaAttributes() *map[string]interface{}
	MysqlEndpoint() *string
	SetMysqlEndpoint(val *string)
	MysqlEndpointInput() *string
	// The tree node.
	Node() constructs.Node
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	User() *string
	SetUser(val *string)
	UserInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAlias()
	ResetApplicationEndpoint()
	ResetDatabaseEndpoint()
	ResetInsecure()
	ResetJavaEndpoint()
	ResetMaxRetries()
	ResetMysqlEndpoint()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for OraclepaasProvider
type jsiiProxy_OraclepaasProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_OraclepaasProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OraclepaasProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OraclepaasProvider) ApplicationEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OraclepaasProvider) ApplicationEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OraclepaasProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OraclepaasProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OraclepaasProvider) DatabaseEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OraclepaasProvider) DatabaseEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OraclepaasProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OraclepaasProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OraclepaasProvider) IdentityDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OraclepaasProvider) IdentityDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OraclepaasProvider) Insecure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OraclepaasProvider) InsecureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OraclepaasProvider) JavaEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OraclepaasProvider) JavaEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OraclepaasProvider) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OraclepaasProvider) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OraclepaasProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OraclepaasProvider) MysqlEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mysqlEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OraclepaasProvider) MysqlEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mysqlEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OraclepaasProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OraclepaasProvider) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OraclepaasProvider) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OraclepaasProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OraclepaasProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OraclepaasProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OraclepaasProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OraclepaasProvider) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OraclepaasProvider) UserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/oraclepaas oraclepaas} Resource.
func NewOraclepaasProvider(scope constructs.Construct, id *string, config *OraclepaasProviderConfig) OraclepaasProvider {
	_init_.Initialize()

	if err := validateNewOraclepaasProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OraclepaasProvider{}

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.provider.OraclepaasProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/oraclepaas oraclepaas} Resource.
func NewOraclepaasProvider_Override(o OraclepaasProvider, scope constructs.Construct, id *string, config *OraclepaasProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.provider.OraclepaasProvider",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OraclepaasProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_OraclepaasProvider)SetApplicationEndpoint(val *string) {
	_jsii_.Set(
		j,
		"applicationEndpoint",
		val,
	)
}

func (j *jsiiProxy_OraclepaasProvider)SetDatabaseEndpoint(val *string) {
	_jsii_.Set(
		j,
		"databaseEndpoint",
		val,
	)
}

func (j *jsiiProxy_OraclepaasProvider)SetIdentityDomain(val *string) {
	_jsii_.Set(
		j,
		"identityDomain",
		val,
	)
}

func (j *jsiiProxy_OraclepaasProvider)SetInsecure(val interface{}) {
	if err := j.validateSetInsecureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecure",
		val,
	)
}

func (j *jsiiProxy_OraclepaasProvider)SetJavaEndpoint(val *string) {
	_jsii_.Set(
		j,
		"javaEndpoint",
		val,
	)
}

func (j *jsiiProxy_OraclepaasProvider)SetMaxRetries(val *float64) {
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_OraclepaasProvider)SetMysqlEndpoint(val *string) {
	_jsii_.Set(
		j,
		"mysqlEndpoint",
		val,
	)
}

func (j *jsiiProxy_OraclepaasProvider)SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_OraclepaasProvider)SetUser(val *string) {
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func OraclepaasProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOraclepaasProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-oraclepaas.provider.OraclepaasProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OraclepaasProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOraclepaasProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-oraclepaas.provider.OraclepaasProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OraclepaasProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOraclepaasProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-oraclepaas.provider.OraclepaasProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OraclepaasProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-oraclepaas.provider.OraclepaasProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OraclepaasProvider) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OraclepaasProvider) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OraclepaasProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		o,
		"resetAlias",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OraclepaasProvider) ResetApplicationEndpoint() {
	_jsii_.InvokeVoid(
		o,
		"resetApplicationEndpoint",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OraclepaasProvider) ResetDatabaseEndpoint() {
	_jsii_.InvokeVoid(
		o,
		"resetDatabaseEndpoint",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OraclepaasProvider) ResetInsecure() {
	_jsii_.InvokeVoid(
		o,
		"resetInsecure",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OraclepaasProvider) ResetJavaEndpoint() {
	_jsii_.InvokeVoid(
		o,
		"resetJavaEndpoint",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OraclepaasProvider) ResetMaxRetries() {
	_jsii_.InvokeVoid(
		o,
		"resetMaxRetries",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OraclepaasProvider) ResetMysqlEndpoint() {
	_jsii_.InvokeVoid(
		o,
		"resetMysqlEndpoint",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OraclepaasProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OraclepaasProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OraclepaasProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OraclepaasProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OraclepaasProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

