package mysqlserviceinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/jsii"

	"github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/mysqlserviceinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MysqlServiceInstanceMysqlConfigurationOutputReference interface {
	cdktf.ComplexObject
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
	ConnectString() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DbName() *string
	SetDbName(val *string)
	DbNameInput() *string
	DbStorage() *float64
	SetDbStorage(val *float64)
	DbStorageInput() *float64
	EnterpriseMonitorConfiguration() MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference
	EnterpriseMonitorConfigurationInput() *MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfiguration
	// Experimental.
	Fqn() *string
	InternalValue() *MysqlServiceInstanceMysqlConfiguration
	SetInternalValue(val *MysqlServiceInstanceMysqlConfiguration)
	IpAddress() *string
	MysqlCharset() *string
	SetMysqlCharset(val *string)
	MysqlCharsetInput() *string
	MysqlCollation() *string
	SetMysqlCollation(val *string)
	MysqlCollationInput() *string
	MysqlPassword() *string
	SetMysqlPassword(val *string)
	MysqlPasswordInput() *string
	MysqlPort() *float64
	SetMysqlPort(val *float64)
	MysqlPortInput() *float64
	MysqlUsername() *string
	SetMysqlUsername(val *string)
	MysqlUsernameInput() *string
	PublicIpAddress() *string
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
	PutEnterpriseMonitorConfiguration(value *MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfiguration)
	ResetDbName()
	ResetDbStorage()
	ResetEnterpriseMonitorConfiguration()
	ResetMysqlCharset()
	ResetMysqlCollation()
	ResetMysqlPassword()
	ResetMysqlPort()
	ResetMysqlUsername()
	ResetSnapshotName()
	ResetSourceServiceName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MysqlServiceInstanceMysqlConfigurationOutputReference
type jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) ConnectString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) DbName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) DbNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) DbStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) DbStorageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) EnterpriseMonitorConfiguration() MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference {
	var returns MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference
	_jsii_.Get(
		j,
		"enterpriseMonitorConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) EnterpriseMonitorConfigurationInput() *MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfiguration {
	var returns *MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfiguration
	_jsii_.Get(
		j,
		"enterpriseMonitorConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) InternalValue() *MysqlServiceInstanceMysqlConfiguration {
	var returns *MysqlServiceInstanceMysqlConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) MysqlCharset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mysqlCharset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) MysqlCharsetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mysqlCharsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) MysqlCollation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mysqlCollation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) MysqlCollationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mysqlCollationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) MysqlPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mysqlPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) MysqlPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mysqlPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) MysqlPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mysqlPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) MysqlPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mysqlPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) MysqlUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mysqlUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) MysqlUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mysqlUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) PublicIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) SnapshotName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) SnapshotNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) SourceServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceServiceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) SourceServiceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceServiceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMysqlServiceInstanceMysqlConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MysqlServiceInstanceMysqlConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewMysqlServiceInstanceMysqlConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.mysqlServiceInstance.MysqlServiceInstanceMysqlConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMysqlServiceInstanceMysqlConfigurationOutputReference_Override(m MysqlServiceInstanceMysqlConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.mysqlServiceInstance.MysqlServiceInstanceMysqlConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference)SetDbName(val *string) {
	if err := j.validateSetDbNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbName",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference)SetDbStorage(val *float64) {
	if err := j.validateSetDbStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbStorage",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference)SetInternalValue(val *MysqlServiceInstanceMysqlConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference)SetMysqlCharset(val *string) {
	if err := j.validateSetMysqlCharsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mysqlCharset",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference)SetMysqlCollation(val *string) {
	if err := j.validateSetMysqlCollationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mysqlCollation",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference)SetMysqlPassword(val *string) {
	if err := j.validateSetMysqlPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mysqlPassword",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference)SetMysqlPort(val *float64) {
	if err := j.validateSetMysqlPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mysqlPort",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference)SetMysqlUsername(val *string) {
	if err := j.validateSetMysqlUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mysqlUsername",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference)SetSnapshotName(val *string) {
	if err := j.validateSetSnapshotNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotName",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference)SetSourceServiceName(val *string) {
	if err := j.validateSetSourceServiceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceServiceName",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) PutEnterpriseMonitorConfiguration(value *MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfiguration) {
	if err := m.validatePutEnterpriseMonitorConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEnterpriseMonitorConfiguration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) ResetDbName() {
	_jsii_.InvokeVoid(
		m,
		"resetDbName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) ResetDbStorage() {
	_jsii_.InvokeVoid(
		m,
		"resetDbStorage",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) ResetEnterpriseMonitorConfiguration() {
	_jsii_.InvokeVoid(
		m,
		"resetEnterpriseMonitorConfiguration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) ResetMysqlCharset() {
	_jsii_.InvokeVoid(
		m,
		"resetMysqlCharset",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) ResetMysqlCollation() {
	_jsii_.InvokeVoid(
		m,
		"resetMysqlCollation",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) ResetMysqlPassword() {
	_jsii_.InvokeVoid(
		m,
		"resetMysqlPassword",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) ResetMysqlPort() {
	_jsii_.InvokeVoid(
		m,
		"resetMysqlPort",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) ResetMysqlUsername() {
	_jsii_.InvokeVoid(
		m,
		"resetMysqlUsername",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) ResetSnapshotName() {
	_jsii_.InvokeVoid(
		m,
		"resetSnapshotName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) ResetSourceServiceName() {
	_jsii_.InvokeVoid(
		m,
		"resetSourceServiceName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

