package databaseserviceinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/databaseserviceinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance oraclepaas_database_service_instance}.
type DatabaseServiceInstance interface {
	cdktf.TerraformResource
	AvailabilityDomain() *string
	SetAvailabilityDomain(val *string)
	AvailabilityDomainInput() *string
	Backups() DatabaseServiceInstanceBackupsOutputReference
	BackupsInput() *DatabaseServiceInstanceBackups
	BringYourOwnLicense() interface{}
	SetBringYourOwnLicense(val interface{})
	BringYourOwnLicenseInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudStorageContainer() *string
	ComputeSiteName() *string
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
	DatabaseConfiguration() DatabaseServiceInstanceDatabaseConfigurationOutputReference
	DatabaseConfigurationInput() *DatabaseServiceInstanceDatabaseConfiguration
	DbaasMonitorUrl() *string
	DefaultAccessRules() DatabaseServiceInstanceDefaultAccessRulesOutputReference
	DefaultAccessRulesInput() *DatabaseServiceInstanceDefaultAccessRules
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DesiredState() *string
	SetDesiredState(val *string)
	DesiredStateInput() *string
	Edition() *string
	SetEdition(val *string)
	EditionInput() *string
	EmUrl() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GlassfishUrl() *string
	HighPerformanceStorage() interface{}
	SetHighPerformanceStorage(val interface{})
	HighPerformanceStorageInput() interface{}
	HybridDisasterRecovery() DatabaseServiceInstanceHybridDisasterRecoveryOutputReference
	HybridDisasterRecoveryInput() *DatabaseServiceInstanceHybridDisasterRecovery
	Id() *string
	SetId(val *string)
	IdentityDomain() *string
	IdInput() *string
	InstantiateFromBackup() DatabaseServiceInstanceInstantiateFromBackupOutputReference
	InstantiateFromBackupInput() *DatabaseServiceInstanceInstantiateFromBackup
	IpNetwork() *string
	SetIpNetwork(val *string)
	IpNetworkInput() *string
	IpReservations() *[]*string
	SetIpReservations(val *[]*string)
	IpReservationsInput() *[]*string
	Level() *string
	SetLevel(val *string)
	LevelInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	Shape() *string
	SetShape(val *string)
	ShapeInput() *string
	SshPublicKey() *string
	SetSshPublicKey(val *string)
	SshPublicKeyInput() *string
	Standby() DatabaseServiceInstanceStandbyOutputReference
	StandbyInput() *DatabaseServiceInstanceStandby
	Status() *string
	Subnet() *string
	SetSubnet(val *string)
	SubnetInput() *string
	SubscriptionType() *string
	SetSubscriptionType(val *string)
	SubscriptionTypeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DatabaseServiceInstanceTimeoutsOutputReference
	TimeoutsInput() interface{}
	Uri() *string
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutBackups(value *DatabaseServiceInstanceBackups)
	PutDatabaseConfiguration(value *DatabaseServiceInstanceDatabaseConfiguration)
	PutDefaultAccessRules(value *DatabaseServiceInstanceDefaultAccessRules)
	PutHybridDisasterRecovery(value *DatabaseServiceInstanceHybridDisasterRecovery)
	PutInstantiateFromBackup(value *DatabaseServiceInstanceInstantiateFromBackup)
	PutStandby(value *DatabaseServiceInstanceStandby)
	PutTimeouts(value *DatabaseServiceInstanceTimeouts)
	ResetAvailabilityDomain()
	ResetBackups()
	ResetBringYourOwnLicense()
	ResetDefaultAccessRules()
	ResetDescription()
	ResetDesiredState()
	ResetHighPerformanceStorage()
	ResetHybridDisasterRecovery()
	ResetId()
	ResetInstantiateFromBackup()
	ResetIpNetwork()
	ResetIpReservations()
	ResetLevel()
	ResetNotificationEmail()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetStandby()
	ResetSubnet()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DatabaseServiceInstance
type jsiiProxy_DatabaseServiceInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DatabaseServiceInstance) AvailabilityDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) AvailabilityDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) Backups() DatabaseServiceInstanceBackupsOutputReference {
	var returns DatabaseServiceInstanceBackupsOutputReference
	_jsii_.Get(
		j,
		"backups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) BackupsInput() *DatabaseServiceInstanceBackups {
	var returns *DatabaseServiceInstanceBackups
	_jsii_.Get(
		j,
		"backupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) BringYourOwnLicense() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bringYourOwnLicense",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) BringYourOwnLicenseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bringYourOwnLicenseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) CloudStorageContainer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStorageContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) ComputeSiteName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeSiteName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) DatabaseConfiguration() DatabaseServiceInstanceDatabaseConfigurationOutputReference {
	var returns DatabaseServiceInstanceDatabaseConfigurationOutputReference
	_jsii_.Get(
		j,
		"databaseConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) DatabaseConfigurationInput() *DatabaseServiceInstanceDatabaseConfiguration {
	var returns *DatabaseServiceInstanceDatabaseConfiguration
	_jsii_.Get(
		j,
		"databaseConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) DbaasMonitorUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbaasMonitorUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) DefaultAccessRules() DatabaseServiceInstanceDefaultAccessRulesOutputReference {
	var returns DatabaseServiceInstanceDefaultAccessRulesOutputReference
	_jsii_.Get(
		j,
		"defaultAccessRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) DefaultAccessRulesInput() *DatabaseServiceInstanceDefaultAccessRules {
	var returns *DatabaseServiceInstanceDefaultAccessRules
	_jsii_.Get(
		j,
		"defaultAccessRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) DesiredState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) DesiredStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) Edition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) EditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"editionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) EmUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) GlassfishUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"glassfishUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) HighPerformanceStorage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"highPerformanceStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) HighPerformanceStorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"highPerformanceStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) HybridDisasterRecovery() DatabaseServiceInstanceHybridDisasterRecoveryOutputReference {
	var returns DatabaseServiceInstanceHybridDisasterRecoveryOutputReference
	_jsii_.Get(
		j,
		"hybridDisasterRecovery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) HybridDisasterRecoveryInput() *DatabaseServiceInstanceHybridDisasterRecovery {
	var returns *DatabaseServiceInstanceHybridDisasterRecovery
	_jsii_.Get(
		j,
		"hybridDisasterRecoveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) IdentityDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) InstantiateFromBackup() DatabaseServiceInstanceInstantiateFromBackupOutputReference {
	var returns DatabaseServiceInstanceInstantiateFromBackupOutputReference
	_jsii_.Get(
		j,
		"instantiateFromBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) InstantiateFromBackupInput() *DatabaseServiceInstanceInstantiateFromBackup {
	var returns *DatabaseServiceInstanceInstantiateFromBackup
	_jsii_.Get(
		j,
		"instantiateFromBackupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) IpNetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) IpNetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) IpReservations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipReservations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) IpReservationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipReservationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) LevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"levelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) NotificationEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) NotificationEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) Shape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) ShapeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shapeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) SshPublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshPublicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) SshPublicKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshPublicKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) Standby() DatabaseServiceInstanceStandbyOutputReference {
	var returns DatabaseServiceInstanceStandbyOutputReference
	_jsii_.Get(
		j,
		"standby",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) StandbyInput() *DatabaseServiceInstanceStandby {
	var returns *DatabaseServiceInstanceStandby
	_jsii_.Get(
		j,
		"standbyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) Subnet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) SubnetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) SubscriptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) SubscriptionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) Timeouts() DatabaseServiceInstanceTimeoutsOutputReference {
	var returns DatabaseServiceInstanceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstance) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance oraclepaas_database_service_instance} Resource.
func NewDatabaseServiceInstance(scope constructs.Construct, id *string, config *DatabaseServiceInstanceConfig) DatabaseServiceInstance {
	_init_.Initialize()

	if err := validateNewDatabaseServiceInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseServiceInstance{}

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.databaseServiceInstance.DatabaseServiceInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/oraclepaas/r/database_service_instance oraclepaas_database_service_instance} Resource.
func NewDatabaseServiceInstance_Override(d DatabaseServiceInstance, scope constructs.Construct, id *string, config *DatabaseServiceInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.databaseServiceInstance.DatabaseServiceInstance",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DatabaseServiceInstance)SetAvailabilityDomain(val *string) {
	if err := j.validateSetAvailabilityDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityDomain",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstance)SetBringYourOwnLicense(val interface{}) {
	if err := j.validateSetBringYourOwnLicenseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bringYourOwnLicense",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstance)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstance)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstance)SetDesiredState(val *string) {
	if err := j.validateSetDesiredStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredState",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstance)SetEdition(val *string) {
	if err := j.validateSetEditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edition",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstance)SetHighPerformanceStorage(val interface{}) {
	if err := j.validateSetHighPerformanceStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"highPerformanceStorage",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstance)SetIpNetwork(val *string) {
	if err := j.validateSetIpNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipNetwork",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstance)SetIpReservations(val *[]*string) {
	if err := j.validateSetIpReservationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipReservations",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstance)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstance)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstance)SetNotificationEmail(val *string) {
	if err := j.validateSetNotificationEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationEmail",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstance)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstance)SetShape(val *string) {
	if err := j.validateSetShapeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shape",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstance)SetSshPublicKey(val *string) {
	if err := j.validateSetSshPublicKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshPublicKey",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstance)SetSubnet(val *string) {
	if err := j.validateSetSubnetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnet",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstance)SetSubscriptionType(val *string) {
	if err := j.validateSetSubscriptionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subscriptionType",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstance)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
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
func DatabaseServiceInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseServiceInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-oraclepaas.databaseServiceInstance.DatabaseServiceInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DatabaseServiceInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-oraclepaas.databaseServiceInstance.DatabaseServiceInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DatabaseServiceInstance) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DatabaseServiceInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabaseServiceInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseServiceInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabaseServiceInstance) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabaseServiceInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabaseServiceInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabaseServiceInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabaseServiceInstance) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabaseServiceInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabaseServiceInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseServiceInstance) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DatabaseServiceInstance) PutBackups(value *DatabaseServiceInstanceBackups) {
	if err := d.validatePutBackupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBackups",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseServiceInstance) PutDatabaseConfiguration(value *DatabaseServiceInstanceDatabaseConfiguration) {
	if err := d.validatePutDatabaseConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDatabaseConfiguration",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseServiceInstance) PutDefaultAccessRules(value *DatabaseServiceInstanceDefaultAccessRules) {
	if err := d.validatePutDefaultAccessRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDefaultAccessRules",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseServiceInstance) PutHybridDisasterRecovery(value *DatabaseServiceInstanceHybridDisasterRecovery) {
	if err := d.validatePutHybridDisasterRecoveryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHybridDisasterRecovery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseServiceInstance) PutInstantiateFromBackup(value *DatabaseServiceInstanceInstantiateFromBackup) {
	if err := d.validatePutInstantiateFromBackupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInstantiateFromBackup",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseServiceInstance) PutStandby(value *DatabaseServiceInstanceStandby) {
	if err := d.validatePutStandbyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStandby",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseServiceInstance) PutTimeouts(value *DatabaseServiceInstanceTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseServiceInstance) ResetAvailabilityDomain() {
	_jsii_.InvokeVoid(
		d,
		"resetAvailabilityDomain",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstance) ResetBackups() {
	_jsii_.InvokeVoid(
		d,
		"resetBackups",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstance) ResetBringYourOwnLicense() {
	_jsii_.InvokeVoid(
		d,
		"resetBringYourOwnLicense",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstance) ResetDefaultAccessRules() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultAccessRules",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstance) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstance) ResetDesiredState() {
	_jsii_.InvokeVoid(
		d,
		"resetDesiredState",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstance) ResetHighPerformanceStorage() {
	_jsii_.InvokeVoid(
		d,
		"resetHighPerformanceStorage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstance) ResetHybridDisasterRecovery() {
	_jsii_.InvokeVoid(
		d,
		"resetHybridDisasterRecovery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstance) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstance) ResetInstantiateFromBackup() {
	_jsii_.InvokeVoid(
		d,
		"resetInstantiateFromBackup",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstance) ResetIpNetwork() {
	_jsii_.InvokeVoid(
		d,
		"resetIpNetwork",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstance) ResetIpReservations() {
	_jsii_.InvokeVoid(
		d,
		"resetIpReservations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstance) ResetLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstance) ResetNotificationEmail() {
	_jsii_.InvokeVoid(
		d,
		"resetNotificationEmail",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstance) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstance) ResetStandby() {
	_jsii_.InvokeVoid(
		d,
		"resetStandby",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstance) ResetSubnet() {
	_jsii_.InvokeVoid(
		d,
		"resetSubnet",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseServiceInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseServiceInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseServiceInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

