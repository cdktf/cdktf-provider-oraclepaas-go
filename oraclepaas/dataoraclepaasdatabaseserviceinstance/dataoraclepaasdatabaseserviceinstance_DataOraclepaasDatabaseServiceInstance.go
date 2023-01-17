package dataoraclepaasdatabaseserviceinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/v3/dataoraclepaasdatabaseserviceinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/oraclepaas/d/database_service_instance oraclepaas_database_service_instance}.
type DataOraclepaasDatabaseServiceInstance interface {
	cdktf.TerraformDataSource
	ApexUrl() *string
	AvailabilityDomain() *string
	BackupDestination() *string
	BringYourOwnLicense() cdktf.IResolvable
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CharacterSet() *string
	CloudStorageContainer() *string
	ComputeSiteName() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	Edition() *string
	EnterpriseManagerUrl() *string
	FailoverDatabase() cdktf.IResolvable
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GlassfishUrl() *string
	HighPerformanceStorage() cdktf.IResolvable
	HybridDisasterRecoveryIp() *string
	Id() *string
	SetId(val *string)
	IdentityDomain() *string
	IdInput() *string
	IpNetwork() *string
	IpReservations() *string
	Level() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ListenerPort() *float64
	MonitorUrl() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NationalCharacterSet() *string
	// The tree node.
	Node() constructs.Node
	PluggableDatabaseName() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	Shape() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Uri() *string
	Version() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetId()
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

// The jsii proxy struct for DataOraclepaasDatabaseServiceInstance
type jsiiProxy_DataOraclepaasDatabaseServiceInstance struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) ApexUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apexUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) AvailabilityDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) BackupDestination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) BringYourOwnLicense() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"bringYourOwnLicense",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) CharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) CloudStorageContainer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStorageContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) ComputeSiteName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeSiteName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) Edition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) EnterpriseManagerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enterpriseManagerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) FailoverDatabase() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"failoverDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) GlassfishUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"glassfishUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) HighPerformanceStorage() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"highPerformanceStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) HybridDisasterRecoveryIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hybridDisasterRecoveryIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) IdentityDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) IpNetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) IpReservations() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipReservations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) ListenerPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"listenerPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) MonitorUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitorUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) NationalCharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nationalCharacterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) PluggableDatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluggableDatabaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) Shape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/oraclepaas/d/database_service_instance oraclepaas_database_service_instance} Data Source.
func NewDataOraclepaasDatabaseServiceInstance(scope constructs.Construct, id *string, config *DataOraclepaasDatabaseServiceInstanceConfig) DataOraclepaasDatabaseServiceInstance {
	_init_.Initialize()

	if err := validateNewDataOraclepaasDatabaseServiceInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataOraclepaasDatabaseServiceInstance{}

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.dataOraclepaasDatabaseServiceInstance.DataOraclepaasDatabaseServiceInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/oraclepaas/d/database_service_instance oraclepaas_database_service_instance} Data Source.
func NewDataOraclepaasDatabaseServiceInstance_Override(d DataOraclepaasDatabaseServiceInstance, scope constructs.Construct, id *string, config *DataOraclepaasDatabaseServiceInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.dataOraclepaasDatabaseServiceInstance.DataOraclepaasDatabaseServiceInstance",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataOraclepaasDatabaseServiceInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
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
func DataOraclepaasDatabaseServiceInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOraclepaasDatabaseServiceInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-oraclepaas.dataOraclepaasDatabaseServiceInstance.DataOraclepaasDatabaseServiceInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataOraclepaasDatabaseServiceInstance_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOraclepaasDatabaseServiceInstance_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-oraclepaas.dataOraclepaasDatabaseServiceInstance.DataOraclepaasDatabaseServiceInstance",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataOraclepaasDatabaseServiceInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOraclepaasDatabaseServiceInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-oraclepaas.dataOraclepaasDatabaseServiceInstance.DataOraclepaasDatabaseServiceInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataOraclepaasDatabaseServiceInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-oraclepaas.dataOraclepaasDatabaseServiceInstance.DataOraclepaasDatabaseServiceInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataOraclepaasDatabaseServiceInstance) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataOraclepaasDatabaseServiceInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOraclepaasDatabaseServiceInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOraclepaasDatabaseServiceInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOraclepaasDatabaseServiceInstance) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOraclepaasDatabaseServiceInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOraclepaasDatabaseServiceInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOraclepaasDatabaseServiceInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOraclepaasDatabaseServiceInstance) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOraclepaasDatabaseServiceInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOraclepaasDatabaseServiceInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOraclepaasDatabaseServiceInstance) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataOraclepaasDatabaseServiceInstance) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOraclepaasDatabaseServiceInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOraclepaasDatabaseServiceInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOraclepaasDatabaseServiceInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOraclepaasDatabaseServiceInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOraclepaasDatabaseServiceInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

