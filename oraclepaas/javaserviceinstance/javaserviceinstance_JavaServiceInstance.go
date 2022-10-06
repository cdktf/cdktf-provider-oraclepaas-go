package javaserviceinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/javaserviceinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/oraclepaas/r/java_service_instance oraclepaas_java_service_instance}.
type JavaServiceInstance interface {
	cdktf.TerraformResource
	AssignPublicIp() interface{}
	SetAssignPublicIp(val interface{})
	AssignPublicIpInput() interface{}
	AvailabilityDomain() *string
	SetAvailabilityDomain(val *string)
	AvailabilityDomainInput() *string
	BackupDestination() *string
	SetBackupDestination(val *string)
	BackupDestinationInput() *string
	Backups() JavaServiceInstanceBackupsOutputReference
	BackupsInput() *JavaServiceInstanceBackups
	BringYourOwnLicense() interface{}
	SetBringYourOwnLicense(val interface{})
	BringYourOwnLicenseInput() interface{}
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
	DesiredState() *string
	SetDesiredState(val *string)
	DesiredStateInput() *string
	Edition() *string
	SetEdition(val *string)
	EditionInput() *string
	EnableAdminConsole() interface{}
	SetEnableAdminConsole(val interface{})
	EnableAdminConsoleInput() interface{}
	ForceDelete() interface{}
	SetForceDelete(val interface{})
	ForceDeleteInput() interface{}
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
	Level() *string
	SetLevel(val *string)
	LevelInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancer() JavaServiceInstanceLoadBalancerOutputReference
	LoadBalancerInput() *JavaServiceInstanceLoadBalancer
	MeteringFrequency() *string
	SetMeteringFrequency(val *string)
	MeteringFrequencyInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NotificationEmail() *string
	SetNotificationEmail(val *string)
	NotificationEmailInput() *string
	OracleTrafficDirector() JavaServiceInstanceOracleTrafficDirectorOutputReference
	OracleTrafficDirectorInput() *JavaServiceInstanceOracleTrafficDirector
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
	ServiceVersion() *string
	SetServiceVersion(val *string)
	ServiceVersionInput() *string
	SnapshotName() *string
	SetSnapshotName(val *string)
	SnapshotNameInput() *string
	SourceServiceName() *string
	SetSourceServiceName(val *string)
	SourceServiceNameInput() *string
	SshPublicKey() *string
	SetSshPublicKey(val *string)
	SshPublicKeyInput() *string
	Status() *string
	Subnet() *string
	SetSubnet(val *string)
	SubnetInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() JavaServiceInstanceTimeoutsOutputReference
	TimeoutsInput() interface{}
	UseIdentityService() interface{}
	SetUseIdentityService(val interface{})
	UseIdentityServiceInput() interface{}
	WeblogicServer() JavaServiceInstanceWeblogicServerOutputReference
	WeblogicServerInput() *JavaServiceInstanceWeblogicServer
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
	PutBackups(value *JavaServiceInstanceBackups)
	PutLoadBalancer(value *JavaServiceInstanceLoadBalancer)
	PutOracleTrafficDirector(value *JavaServiceInstanceOracleTrafficDirector)
	PutTimeouts(value *JavaServiceInstanceTimeouts)
	PutWeblogicServer(value *JavaServiceInstanceWeblogicServer)
	ResetAssignPublicIp()
	ResetAvailabilityDomain()
	ResetBackupDestination()
	ResetBringYourOwnLicense()
	ResetDescription()
	ResetDesiredState()
	ResetEnableAdminConsole()
	ResetForceDelete()
	ResetId()
	ResetIpNetwork()
	ResetLevel()
	ResetLoadBalancer()
	ResetMeteringFrequency()
	ResetNotificationEmail()
	ResetOracleTrafficDirector()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetServiceVersion()
	ResetSnapshotName()
	ResetSourceServiceName()
	ResetSubnet()
	ResetTimeouts()
	ResetUseIdentityService()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for JavaServiceInstance
type jsiiProxy_JavaServiceInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_JavaServiceInstance) AssignPublicIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assignPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) AssignPublicIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assignPublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) AvailabilityDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) AvailabilityDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) BackupDestination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) BackupDestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) Backups() JavaServiceInstanceBackupsOutputReference {
	var returns JavaServiceInstanceBackupsOutputReference
	_jsii_.Get(
		j,
		"backups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) BackupsInput() *JavaServiceInstanceBackups {
	var returns *JavaServiceInstanceBackups
	_jsii_.Get(
		j,
		"backupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) BringYourOwnLicense() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bringYourOwnLicense",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) BringYourOwnLicenseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bringYourOwnLicenseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) DesiredState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) DesiredStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) Edition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) EditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"editionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) EnableAdminConsole() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAdminConsole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) EnableAdminConsoleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAdminConsoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) ForceDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) ForceDeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) IpNetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) IpNetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) LevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"levelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) LoadBalancer() JavaServiceInstanceLoadBalancerOutputReference {
	var returns JavaServiceInstanceLoadBalancerOutputReference
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) LoadBalancerInput() *JavaServiceInstanceLoadBalancer {
	var returns *JavaServiceInstanceLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) MeteringFrequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meteringFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) MeteringFrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meteringFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) NotificationEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) NotificationEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) OracleTrafficDirector() JavaServiceInstanceOracleTrafficDirectorOutputReference {
	var returns JavaServiceInstanceOracleTrafficDirectorOutputReference
	_jsii_.Get(
		j,
		"oracleTrafficDirector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) OracleTrafficDirectorInput() *JavaServiceInstanceOracleTrafficDirector {
	var returns *JavaServiceInstanceOracleTrafficDirector
	_jsii_.Get(
		j,
		"oracleTrafficDirectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) ServiceVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) ServiceVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) SnapshotName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) SnapshotNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) SourceServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceServiceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) SourceServiceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceServiceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) SshPublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshPublicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) SshPublicKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshPublicKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) Subnet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) SubnetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) Timeouts() JavaServiceInstanceTimeoutsOutputReference {
	var returns JavaServiceInstanceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) UseIdentityService() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useIdentityService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) UseIdentityServiceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useIdentityServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) WeblogicServer() JavaServiceInstanceWeblogicServerOutputReference {
	var returns JavaServiceInstanceWeblogicServerOutputReference
	_jsii_.Get(
		j,
		"weblogicServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) WeblogicServerInput() *JavaServiceInstanceWeblogicServer {
	var returns *JavaServiceInstanceWeblogicServer
	_jsii_.Get(
		j,
		"weblogicServerInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/oraclepaas/r/java_service_instance oraclepaas_java_service_instance} Resource.
func NewJavaServiceInstance(scope constructs.Construct, id *string, config *JavaServiceInstanceConfig) JavaServiceInstance {
	_init_.Initialize()

	if err := validateNewJavaServiceInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_JavaServiceInstance{}

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.javaServiceInstance.JavaServiceInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/oraclepaas/r/java_service_instance oraclepaas_java_service_instance} Resource.
func NewJavaServiceInstance_Override(j JavaServiceInstance, scope constructs.Construct, id *string, config *JavaServiceInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.javaServiceInstance.JavaServiceInstance",
		[]interface{}{scope, id, config},
		j,
	)
}

func (j *jsiiProxy_JavaServiceInstance)SetAssignPublicIp(val interface{}) {
	if err := j.validateSetAssignPublicIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assignPublicIp",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstance)SetAvailabilityDomain(val *string) {
	if err := j.validateSetAvailabilityDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityDomain",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstance)SetBackupDestination(val *string) {
	if err := j.validateSetBackupDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupDestination",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstance)SetBringYourOwnLicense(val interface{}) {
	if err := j.validateSetBringYourOwnLicenseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bringYourOwnLicense",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstance)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstance)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstance)SetDesiredState(val *string) {
	if err := j.validateSetDesiredStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredState",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstance)SetEdition(val *string) {
	if err := j.validateSetEditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edition",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstance)SetEnableAdminConsole(val interface{}) {
	if err := j.validateSetEnableAdminConsoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAdminConsole",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstance)SetForceDelete(val interface{}) {
	if err := j.validateSetForceDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDelete",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstance)SetIpNetwork(val *string) {
	if err := j.validateSetIpNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipNetwork",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstance)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstance)SetMeteringFrequency(val *string) {
	if err := j.validateSetMeteringFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"meteringFrequency",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstance)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstance)SetNotificationEmail(val *string) {
	if err := j.validateSetNotificationEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationEmail",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstance)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstance)SetServiceVersion(val *string) {
	if err := j.validateSetServiceVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceVersion",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstance)SetSnapshotName(val *string) {
	if err := j.validateSetSnapshotNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotName",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstance)SetSourceServiceName(val *string) {
	if err := j.validateSetSourceServiceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceServiceName",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstance)SetSshPublicKey(val *string) {
	if err := j.validateSetSshPublicKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshPublicKey",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstance)SetSubnet(val *string) {
	if err := j.validateSetSubnetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnet",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstance)SetUseIdentityService(val interface{}) {
	if err := j.validateSetUseIdentityServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useIdentityService",
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
func JavaServiceInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJavaServiceInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-oraclepaas.javaServiceInstance.JavaServiceInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func JavaServiceInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-oraclepaas.javaServiceInstance.JavaServiceInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstance) AddOverride(path *string, value interface{}) {
	if err := j.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (j *jsiiProxy_JavaServiceInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := j.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := j.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := j.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		j,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstance) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := j.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := j.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		j,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := j.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		j,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := j.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		j,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstance) GetStringAttribute(terraformAttribute *string) *string {
	if err := j.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := j.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		j,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := j.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstance) OverrideLogicalId(newLogicalId *string) {
	if err := j.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (j *jsiiProxy_JavaServiceInstance) PutBackups(value *JavaServiceInstanceBackups) {
	if err := j.validatePutBackupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putBackups",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JavaServiceInstance) PutLoadBalancer(value *JavaServiceInstanceLoadBalancer) {
	if err := j.validatePutLoadBalancerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putLoadBalancer",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JavaServiceInstance) PutOracleTrafficDirector(value *JavaServiceInstanceOracleTrafficDirector) {
	if err := j.validatePutOracleTrafficDirectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putOracleTrafficDirector",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JavaServiceInstance) PutTimeouts(value *JavaServiceInstanceTimeouts) {
	if err := j.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JavaServiceInstance) PutWeblogicServer(value *JavaServiceInstanceWeblogicServer) {
	if err := j.validatePutWeblogicServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putWeblogicServer",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JavaServiceInstance) ResetAssignPublicIp() {
	_jsii_.InvokeVoid(
		j,
		"resetAssignPublicIp",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstance) ResetAvailabilityDomain() {
	_jsii_.InvokeVoid(
		j,
		"resetAvailabilityDomain",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstance) ResetBackupDestination() {
	_jsii_.InvokeVoid(
		j,
		"resetBackupDestination",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstance) ResetBringYourOwnLicense() {
	_jsii_.InvokeVoid(
		j,
		"resetBringYourOwnLicense",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstance) ResetDescription() {
	_jsii_.InvokeVoid(
		j,
		"resetDescription",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstance) ResetDesiredState() {
	_jsii_.InvokeVoid(
		j,
		"resetDesiredState",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstance) ResetEnableAdminConsole() {
	_jsii_.InvokeVoid(
		j,
		"resetEnableAdminConsole",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstance) ResetForceDelete() {
	_jsii_.InvokeVoid(
		j,
		"resetForceDelete",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstance) ResetId() {
	_jsii_.InvokeVoid(
		j,
		"resetId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstance) ResetIpNetwork() {
	_jsii_.InvokeVoid(
		j,
		"resetIpNetwork",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstance) ResetLevel() {
	_jsii_.InvokeVoid(
		j,
		"resetLevel",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstance) ResetLoadBalancer() {
	_jsii_.InvokeVoid(
		j,
		"resetLoadBalancer",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstance) ResetMeteringFrequency() {
	_jsii_.InvokeVoid(
		j,
		"resetMeteringFrequency",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstance) ResetNotificationEmail() {
	_jsii_.InvokeVoid(
		j,
		"resetNotificationEmail",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstance) ResetOracleTrafficDirector() {
	_jsii_.InvokeVoid(
		j,
		"resetOracleTrafficDirector",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		j,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstance) ResetRegion() {
	_jsii_.InvokeVoid(
		j,
		"resetRegion",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstance) ResetServiceVersion() {
	_jsii_.InvokeVoid(
		j,
		"resetServiceVersion",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstance) ResetSnapshotName() {
	_jsii_.InvokeVoid(
		j,
		"resetSnapshotName",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstance) ResetSourceServiceName() {
	_jsii_.InvokeVoid(
		j,
		"resetSourceServiceName",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstance) ResetSubnet() {
	_jsii_.InvokeVoid(
		j,
		"resetSubnet",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		j,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstance) ResetUseIdentityService() {
	_jsii_.InvokeVoid(
		j,
		"resetUseIdentityService",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		j,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		j,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

