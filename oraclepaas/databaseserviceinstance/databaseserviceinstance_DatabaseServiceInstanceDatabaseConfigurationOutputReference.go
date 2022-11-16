package databaseserviceinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/v2/jsii"

	"github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/v2/databaseserviceinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabaseServiceInstanceDatabaseConfigurationOutputReference interface {
	cdktf.ComplexObject
	AdminPassword() *string
	SetAdminPassword(val *string)
	AdminPasswordInput() *string
	BackupDestination() *string
	SetBackupDestination(val *string)
	BackupDestinationInput() *string
	BackupStorageVolumeSize() *float64
	SetBackupStorageVolumeSize(val *float64)
	BackupStorageVolumeSizeInput() *float64
	CharacterSet() *string
	SetCharacterSet(val *string)
	CharacterSetInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DataStorageVolumeSize() *float64
	SetDataStorageVolumeSize(val *float64)
	DataStorageVolumeSizeInput() *float64
	DbDemo() *string
	SetDbDemo(val *string)
	DbDemoInput() *string
	DisasterRecovery() interface{}
	SetDisasterRecovery(val interface{})
	DisasterRecoveryInput() interface{}
	FailoverDatabase() interface{}
	SetFailoverDatabase(val interface{})
	FailoverDatabaseInput() interface{}
	// Experimental.
	Fqn() *string
	GoldenGate() interface{}
	SetGoldenGate(val interface{})
	GoldenGateInput() interface{}
	InternalValue() *DatabaseServiceInstanceDatabaseConfiguration
	SetInternalValue(val *DatabaseServiceInstanceDatabaseConfiguration)
	IsRac() interface{}
	SetIsRac(val interface{})
	IsRacInput() interface{}
	NationalCharacterSet() *string
	SetNationalCharacterSet(val *string)
	NationalCharacterSetInput() *string
	PdbName() *string
	SetPdbName(val *string)
	PdbNameInput() *string
	Sid() *string
	SetSid(val *string)
	SidInput() *string
	SnapshotName() *string
	SetSnapshotName(val *string)
	SnapshotNameInput() *string
	SourceServiceName() *string
	SetSourceServiceName(val *string)
	SourceServiceNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timezone() *string
	SetTimezone(val *string)
	TimezoneInput() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UsableStorage() *float64
	SetUsableStorage(val *float64)
	UsableStorageInput() *float64
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetBackupDestination()
	ResetBackupStorageVolumeSize()
	ResetCharacterSet()
	ResetDataStorageVolumeSize()
	ResetDbDemo()
	ResetDisasterRecovery()
	ResetFailoverDatabase()
	ResetGoldenGate()
	ResetIsRac()
	ResetNationalCharacterSet()
	ResetPdbName()
	ResetSid()
	ResetSnapshotName()
	ResetSourceServiceName()
	ResetTimezone()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatabaseServiceInstanceDatabaseConfigurationOutputReference
type jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) AdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) AdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) BackupDestination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) BackupDestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) BackupStorageVolumeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupStorageVolumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) BackupStorageVolumeSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupStorageVolumeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) CharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) CharacterSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) DataStorageVolumeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageVolumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) DataStorageVolumeSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageVolumeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) DbDemo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbDemo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) DbDemoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbDemoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) DisasterRecovery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disasterRecovery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) DisasterRecoveryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disasterRecoveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) FailoverDatabase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failoverDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) FailoverDatabaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failoverDatabaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) GoldenGate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"goldenGate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) GoldenGateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"goldenGateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) InternalValue() *DatabaseServiceInstanceDatabaseConfiguration {
	var returns *DatabaseServiceInstanceDatabaseConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) IsRac() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isRac",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) IsRacInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isRacInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) NationalCharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nationalCharacterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) NationalCharacterSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nationalCharacterSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) PdbName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pdbName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) PdbNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pdbNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) Sid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) SidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) SnapshotName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) SnapshotNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) SourceServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceServiceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) SourceServiceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceServiceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) UsableStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"usableStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) UsableStorageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"usableStorageInput",
		&returns,
	)
	return returns
}


func NewDatabaseServiceInstanceDatabaseConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatabaseServiceInstanceDatabaseConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewDatabaseServiceInstanceDatabaseConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.databaseServiceInstance.DatabaseServiceInstanceDatabaseConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatabaseServiceInstanceDatabaseConfigurationOutputReference_Override(d DatabaseServiceInstanceDatabaseConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.databaseServiceInstance.DatabaseServiceInstanceDatabaseConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference)SetAdminPassword(val *string) {
	if err := j.validateSetAdminPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminPassword",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference)SetBackupDestination(val *string) {
	if err := j.validateSetBackupDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupDestination",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference)SetBackupStorageVolumeSize(val *float64) {
	if err := j.validateSetBackupStorageVolumeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupStorageVolumeSize",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference)SetCharacterSet(val *string) {
	if err := j.validateSetCharacterSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"characterSet",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference)SetDataStorageVolumeSize(val *float64) {
	if err := j.validateSetDataStorageVolumeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataStorageVolumeSize",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference)SetDbDemo(val *string) {
	if err := j.validateSetDbDemoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbDemo",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference)SetDisasterRecovery(val interface{}) {
	if err := j.validateSetDisasterRecoveryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disasterRecovery",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference)SetFailoverDatabase(val interface{}) {
	if err := j.validateSetFailoverDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failoverDatabase",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference)SetGoldenGate(val interface{}) {
	if err := j.validateSetGoldenGateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"goldenGate",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference)SetInternalValue(val *DatabaseServiceInstanceDatabaseConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference)SetIsRac(val interface{}) {
	if err := j.validateSetIsRacParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isRac",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference)SetNationalCharacterSet(val *string) {
	if err := j.validateSetNationalCharacterSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nationalCharacterSet",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference)SetPdbName(val *string) {
	if err := j.validateSetPdbNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pdbName",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference)SetSid(val *string) {
	if err := j.validateSetSidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sid",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference)SetSnapshotName(val *string) {
	if err := j.validateSetSnapshotNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotName",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference)SetSourceServiceName(val *string) {
	if err := j.validateSetSourceServiceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceServiceName",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference)SetUsableStorage(val *float64) {
	if err := j.validateSetUsableStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usableStorage",
		val,
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) ResetBackupDestination() {
	_jsii_.InvokeVoid(
		d,
		"resetBackupDestination",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) ResetBackupStorageVolumeSize() {
	_jsii_.InvokeVoid(
		d,
		"resetBackupStorageVolumeSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) ResetCharacterSet() {
	_jsii_.InvokeVoid(
		d,
		"resetCharacterSet",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) ResetDataStorageVolumeSize() {
	_jsii_.InvokeVoid(
		d,
		"resetDataStorageVolumeSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) ResetDbDemo() {
	_jsii_.InvokeVoid(
		d,
		"resetDbDemo",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) ResetDisasterRecovery() {
	_jsii_.InvokeVoid(
		d,
		"resetDisasterRecovery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) ResetFailoverDatabase() {
	_jsii_.InvokeVoid(
		d,
		"resetFailoverDatabase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) ResetGoldenGate() {
	_jsii_.InvokeVoid(
		d,
		"resetGoldenGate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) ResetIsRac() {
	_jsii_.InvokeVoid(
		d,
		"resetIsRac",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) ResetNationalCharacterSet() {
	_jsii_.InvokeVoid(
		d,
		"resetNationalCharacterSet",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) ResetPdbName() {
	_jsii_.InvokeVoid(
		d,
		"resetPdbName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) ResetSid() {
	_jsii_.InvokeVoid(
		d,
		"resetSid",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) ResetSnapshotName() {
	_jsii_.InvokeVoid(
		d,
		"resetSnapshotName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) ResetSourceServiceName() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceServiceName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) ResetTimezone() {
	_jsii_.InvokeVoid(
		d,
		"resetTimezone",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		d,
		"resetType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseServiceInstanceDatabaseConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

