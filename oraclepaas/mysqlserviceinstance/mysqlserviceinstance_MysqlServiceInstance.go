package mysqlserviceinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/v2/mysqlserviceinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/oraclepaas/r/mysql_service_instance oraclepaas_mysql_service_instance}.
type MysqlServiceInstance interface {
	cdktf.TerraformResource
	AvailabilityDomain() *string
	SetAvailabilityDomain(val *string)
	AvailabilityDomainInput() *string
	BackupDestination() *string
	SetBackupDestination(val *string)
	BackupDestinationInput() *string
	Backups() MysqlServiceInstanceBackupsOutputReference
	BackupsInput() *MysqlServiceInstanceBackups
	BaseReleaseVersion() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
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
	SetDescription(val *string)
	DescriptionInput() *string
	EmUrl() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IpNetwork() *string
	SetIpNetwork(val *string)
	IpNetworkInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MeteringFrequency() *string
	SetMeteringFrequency(val *string)
	MeteringFrequencyInput() *string
	MysqlConfiguration() MysqlServiceInstanceMysqlConfigurationOutputReference
	MysqlConfigurationInput() *MysqlServiceInstanceMysqlConfiguration
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NotificationEmail() *string
	SetNotificationEmail(val *string)
	NotificationEmailInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ReleaseVersion() *string
	ServiceVersion() *string
	Shape() *string
	SetShape(val *string)
	ShapeInput() *string
	SshPublicKey() *string
	SetSshPublicKey(val *string)
	SshPublicKeyInput() *string
	Subnet() *string
	SetSubnet(val *string)
	SubnetInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MysqlServiceInstanceTimeoutsOutputReference
	TimeoutsInput() interface{}
	VmUser() *string
	SetVmUser(val *string)
	VmUserInput() *string
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
	PutBackups(value *MysqlServiceInstanceBackups)
	PutMysqlConfiguration(value *MysqlServiceInstanceMysqlConfiguration)
	PutTimeouts(value *MysqlServiceInstanceTimeouts)
	ResetAvailabilityDomain()
	ResetBackupDestination()
	ResetBackups()
	ResetDescription()
	ResetId()
	ResetIpNetwork()
	ResetMeteringFrequency()
	ResetNotificationEmail()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetSubnet()
	ResetTimeouts()
	ResetVmUser()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for MysqlServiceInstance
type jsiiProxy_MysqlServiceInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MysqlServiceInstance) AvailabilityDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) AvailabilityDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) BackupDestination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) BackupDestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) Backups() MysqlServiceInstanceBackupsOutputReference {
	var returns MysqlServiceInstanceBackupsOutputReference
	_jsii_.Get(
		j,
		"backups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) BackupsInput() *MysqlServiceInstanceBackups {
	var returns *MysqlServiceInstanceBackups
	_jsii_.Get(
		j,
		"backupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) BaseReleaseVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseReleaseVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) EmUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) IpNetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) IpNetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) MeteringFrequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meteringFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) MeteringFrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meteringFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) MysqlConfiguration() MysqlServiceInstanceMysqlConfigurationOutputReference {
	var returns MysqlServiceInstanceMysqlConfigurationOutputReference
	_jsii_.Get(
		j,
		"mysqlConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) MysqlConfigurationInput() *MysqlServiceInstanceMysqlConfiguration {
	var returns *MysqlServiceInstanceMysqlConfiguration
	_jsii_.Get(
		j,
		"mysqlConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) NotificationEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) NotificationEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) ReleaseVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) ServiceVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) Shape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) ShapeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shapeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) SshPublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshPublicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) SshPublicKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshPublicKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) Subnet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) SubnetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) Timeouts() MysqlServiceInstanceTimeoutsOutputReference {
	var returns MysqlServiceInstanceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) VmUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstance) VmUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmUserInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/oraclepaas/r/mysql_service_instance oraclepaas_mysql_service_instance} Resource.
func NewMysqlServiceInstance(scope constructs.Construct, id *string, config *MysqlServiceInstanceConfig) MysqlServiceInstance {
	_init_.Initialize()

	if err := validateNewMysqlServiceInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MysqlServiceInstance{}

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.mysqlServiceInstance.MysqlServiceInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/oraclepaas/r/mysql_service_instance oraclepaas_mysql_service_instance} Resource.
func NewMysqlServiceInstance_Override(m MysqlServiceInstance, scope constructs.Construct, id *string, config *MysqlServiceInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.mysqlServiceInstance.MysqlServiceInstance",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MysqlServiceInstance)SetAvailabilityDomain(val *string) {
	if err := j.validateSetAvailabilityDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityDomain",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstance)SetBackupDestination(val *string) {
	if err := j.validateSetBackupDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupDestination",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstance)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstance)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstance)SetIpNetwork(val *string) {
	if err := j.validateSetIpNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipNetwork",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstance)SetMeteringFrequency(val *string) {
	if err := j.validateSetMeteringFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"meteringFrequency",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstance)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstance)SetNotificationEmail(val *string) {
	if err := j.validateSetNotificationEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationEmail",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstance)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstance)SetShape(val *string) {
	if err := j.validateSetShapeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shape",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstance)SetSshPublicKey(val *string) {
	if err := j.validateSetSshPublicKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshPublicKey",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstance)SetSubnet(val *string) {
	if err := j.validateSetSubnetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnet",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstance)SetVmUser(val *string) {
	if err := j.validateSetVmUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmUser",
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
func MysqlServiceInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMysqlServiceInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-oraclepaas.mysqlServiceInstance.MysqlServiceInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MysqlServiceInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMysqlServiceInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-oraclepaas.mysqlServiceInstance.MysqlServiceInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MysqlServiceInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMysqlServiceInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-oraclepaas.mysqlServiceInstance.MysqlServiceInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MysqlServiceInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-oraclepaas.mysqlServiceInstance.MysqlServiceInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MysqlServiceInstance) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MysqlServiceInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstance) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstance) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstance) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MysqlServiceInstance) PutBackups(value *MysqlServiceInstanceBackups) {
	if err := m.validatePutBackupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putBackups",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MysqlServiceInstance) PutMysqlConfiguration(value *MysqlServiceInstanceMysqlConfiguration) {
	if err := m.validatePutMysqlConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMysqlConfiguration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MysqlServiceInstance) PutTimeouts(value *MysqlServiceInstanceTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MysqlServiceInstance) ResetAvailabilityDomain() {
	_jsii_.InvokeVoid(
		m,
		"resetAvailabilityDomain",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstance) ResetBackupDestination() {
	_jsii_.InvokeVoid(
		m,
		"resetBackupDestination",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstance) ResetBackups() {
	_jsii_.InvokeVoid(
		m,
		"resetBackups",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstance) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstance) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstance) ResetIpNetwork() {
	_jsii_.InvokeVoid(
		m,
		"resetIpNetwork",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstance) ResetMeteringFrequency() {
	_jsii_.InvokeVoid(
		m,
		"resetMeteringFrequency",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstance) ResetNotificationEmail() {
	_jsii_.InvokeVoid(
		m,
		"resetNotificationEmail",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstance) ResetRegion() {
	_jsii_.InvokeVoid(
		m,
		"resetRegion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstance) ResetSubnet() {
	_jsii_.InvokeVoid(
		m,
		"resetSubnet",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstance) ResetVmUser() {
	_jsii_.InvokeVoid(
		m,
		"resetVmUser",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

